//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"reflect"
	"text/template"

	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"
)

var o string

func init() {
	flag.StringVar(&o, "o", "", "output")
	flag.Parse()
}

func main() {

	cli := reflect.TypeOf(logon.SecureChannel_T)

	var tT = struct {
		Methods []string
		Package string
	}{
		Package: "logon",
	}

	for i := 0; i < cli.NumMethod(); i++ {
		f := cli.Method(i).Func.Type()
		if f.NumIn() != 4 || f.NumOut() != 2 || f.In(2).Elem().Kind() != reflect.Struct {
			continue
		}
		if _, ok := f.In(2).Elem().FieldByName("ReturnAuthenticator"); !ok {
			continue
		}

		tT.Methods = append(tT.Methods, cli.Method(i).Name)
	}

	buf := bytes.NewBuffer(nil)

	if err := GenT.Execute(buf, tT); err != nil {
		fmt.Fprintln(os.Stderr, "gen_t: execute", err)
		os.Exit(1)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, "format: source", err)
		os.Exit(1)
	}

	f := os.Stdout
	if o != "" {
		f, err = os.OpenFile(o, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "open:", o, err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(f, "%s", string(b))
	return
}

var GenT = template.Must(template.New("secure_channel.tpl").Parse(`
package {{ .Package }}

import (
	"fmt"
	"context"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
)
{{ range $m := .Methods }}

func (o *xxx_SecureChannelClient) {{ $m }}(ctx context.Context, in *{{ $m }}Request, opts ...dcerpc.CallOption) (*{{ $m }}Response, error) {

	if in == nil {
		in = &{{ $m }}Request{}
	}

	if err := o.SetAuthenticators(ctx, &in.Authenticator, &in.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx).OpName(), err)
	}

	ret, err := o.LogonClient.{{ $m }}(ctx, in, opts...)
	if err != nil {
		return nil, err
	}

	if err := o.VerifyAuthenticator(ctx, ret.ReturnAuthenticator); err != nil {
		return nil, fmt.Errorf("%s: %v", in.xxx_ToOp(ctx).OpName(), err)
	}

	return ret, nil
}
{{- end }}
`))
