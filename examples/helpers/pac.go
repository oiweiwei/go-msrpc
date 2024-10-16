package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/oiweiwei/go-msrpc/msrpc/pac"
)

var (
	format string
	input  string
)

func init() {
	flag.StringVar(&format, "format", "base64", "input format")
	flag.StringVar(&input, "input", "-", "input file")
	flag.Parse()
}

func main() {

	var (
		f   io.Reader
		err error
	)

	if f = os.Stdin; input != "-" && input != "" {
		if _, err := os.Stat(input); !os.IsNotExist(err) {
			f = bytes.NewReader([]byte(input))
		} else if f, err = os.Open(input); err != nil {
			log.Fatal(err)
		}
	}

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "base64":
		if b, err = base64.StdEncoding.DecodeString(string(b)); err != nil {
			log.Fatal(err)
		}
	case "hex":
		if b, err = hex.DecodeString(string(b)); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown format: %s", format)
	}

	p := &pac.PAC{}

	if err := p.Unmarshal(b); err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, string(b))
}
