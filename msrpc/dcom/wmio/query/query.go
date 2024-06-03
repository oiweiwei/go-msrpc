package query

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ndr"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemclassobject/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi/iwbemservices/v0"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom/wmio"
)

type builder struct {
	ctx    context.Context
	ver    *dcom.COMVersion
	wmi    iwbemservices.ServicesClient
	obj    *wmio.Object
	cls, m string
	err    error
}

func NewBuilder(ctx context.Context, wmi iwbemservices.ServicesClient, ver *dcom.COMVersion) *builder {
	return &builder{ctx: ctx, ver: ver, wmi: wmi}
}

func (b *builder) clone(obj *wmio.Object) *builder {
	return &builder{ctx: b.ctx, ver: b.ver /* err: b.err, */, wmi: b.wmi, obj: obj, m: b.m, cls: b.cls}
}

func (b *builder) withErrf(frmt string, args ...interface{}) *builder {
	if b.err == nil {
		b.err = fmt.Errorf(frmt, args...)
	}
	return b
}

func (b *builder) withErr(err error) *builder {
	b.err = err
	return b
}

func (b *builder) withClassObject(cls *wmi.ClassObject) *builder {

	ref := &dcom.ObjectReference{}

	if err := ndr.Unmarshal(cls.Data, ref, ndr.Opaque); err != nil {
		return b.withErrf("with_reference: ndr: %v", err)
	}

	c, ok := ref.ObjectReference.GetValue().(*dcom.ObjectReferenceCustom)
	if !ok || len(c.ObjectData) == 0 {
		return b.withErrf("with_reference: custom: ref is empty")
	}

	obj, err := wmio.Unmarshal(c.ObjectData)
	if err != nil {
		return b.withErrf("with_reference: unmarshal: %v", err)
	}

	return b.clone(obj)
}

func (b *builder) WithObject(obj any) *builder {

	switch obj := obj.(type) {
	case *wmi.ClassObject:
		return b.withClassObject(obj)
	}
	return b
}

func (b *builder) Spawn(cls string) *builder {

	obj, err := b.wmi.GetObject(b.ctx, &iwbemservices.GetObjectRequest{
		This:       &dcom.ORPCThis{Version: b.ver},
		ObjectPath: &oaut.String{Data: cls},
		Object:     &wmi.ClassObject{},
	})

	if err != nil {
		return b.withErr(err)
	}

	b.cls = cls

	return b.WithObject(obj.Object)
}

// Method function selects the method.
func (b *builder) Method(m string) *builder {

	if b.obj == nil {
		return b.withErrf("Methods: object is nil")
	}

	in, _, err := b.obj.Method(m)
	if err != nil {
		return b.withErrf("Method: %w", err)
	}

	b.m = m

	return b.clone(in)
}

// Values function sets the parameters for the object.
func (b *builder) Values(values wmio.Values, convert ...func(any, wmio.CIMType) (any, bool)) *builder {

	if b.obj == nil {
		return b.withErrf("Values: object is nil")
	}

	params, err := b.obj.New(values, convert...)
	if err != nil {
		return b.withErrf("Values: %v", params)
	}

	return b.clone(params)
}

// Exec function executes the method.
func (b *builder) Exec() *builder {

	in, err := b.ClassObject()
	if err != nil {
		return b.withErr(err)
	}

	exec, err := b.wmi.ExecMethod(b.ctx, &iwbemservices.ExecMethodRequest{
		This:       &dcom.ORPCThis{Version: b.ver},
		ObjectPath: &oaut.String{Data: b.cls},
		MethodName: &oaut.String{Data: b.m},
		InParams:   in,
		OutParams:  &wmi.ClassObject{},
	})

	if err != nil {
		return b.withErr(err)
	}

	return b.withClassObject(exec.OutParams)
}

// Object function returns the WMIO Object.
func (b *builder) Object() (*wmio.Object, error) {
	return b.obj, b.err
}

// ClassObject function returns the WMI ClassObject representation
// for the WMIO Object.
func (b *builder) ClassObject() (*wmi.ClassObject, error) {

	if b.obj == nil {
		return nil, fmt.Errorf("ClassObject: object is nil")
	}

	bb, err := wmio.Marshal(b.obj)
	if err != nil {
		return nil, fmt.Errorf("ClassObject: marshal_object: %w", err)
	}

	ref := &dcom.ObjectReference{
		Signature: []byte{0x4d, 0x45, 0x4f, 0x57},
		Flags:     dcom.ObjectReferenceTypeCustom,
		IID:       iwbemclassobject.ClassObjectIID,
		ObjectReference: &dcom.ObjectReference_ObjectReference{
			Value: &dcom.ObjectReference_Custom{
				Custom: &dcom.ObjectReferenceCustom{
					ClassID:    wmi.ClassObjectUnmarshalClassID,
					ObjectData: bb,
				},
			},
		},
	}

	if bb, err = ndr.Marshal(ref, ndr.Opaque); err != nil {
		return nil, fmt.Errorf("ClassObject: marshal reference: %w", err)
	}

	return &wmi.ClassObject{Data: bb}, nil
}
