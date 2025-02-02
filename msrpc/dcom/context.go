package dcom

import (
	"fmt"

	"github.com/oiweiwei/go-msrpc/ndr"
)

// InterfacePointer function encodes the context object into an InterfacePointer.
func (o *Context) InterfacePointer() (*InterfacePointer, error) {

	blob, err := ndr.Marshal(o, ndr.Opaque)
	if err != nil {
		return nil, fmt.Errorf("marshal context: %w", err)
	}

	ref := &ObjectReference{
		Signature: ([]byte)(ObjectReferenceCustomSignature),
		Flags:     ObjectReferenceTypeCustom,
		IID:       ContextIID,
		ObjectReference: &ObjectReference_ObjectReference{
			Value: &ObjectReference_Custom{
				Custom: &ObjectReferenceCustom{
					ClassID:    ContextMarshalerClassID,
					Size:       uint32(len(blob)),
					ObjectData: blob,
				},
			},
		},
	}

	b, err := ndr.Marshal(ref, ndr.Opaque)
	if err != nil {
		return nil, fmt.Errorf("marshal object_reference: %w", err)
	}

	return &InterfacePointer{Data: b}, nil
}

// Parse function decodes the context object from an InterfacePointer.
func (o *Context) Parse(in *InterfacePointer) error {

	ref := in.GetCustomObjectReference()
	if ref == nil {
		return fmt.Errorf("invalid interface_pointer")
	}

	if err := ndr.Unmarshal(ref.ObjectData, o, ndr.Opaque); err != nil {
		return fmt.Errorf("unmarshal context: %w", err)
	}

	return nil
}
