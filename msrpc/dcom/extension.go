package dcom

import (
	"bytes"
	"fmt"

	ndr "github.com/oiweiwei/go-msrpc/ndr"

	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

const ErrorStringSignature uint32 = 0xFFFFFFFF

const ContextORPCExtensionSignature string = "ANUK"

const EntryHeaderSignature string = "INAN"

// Extension represents the ORPC Extension object.
type Extension interface {
	// ExtensionID returns the ClassID of the extension.
	ExtensionID() *ClassID
	// Data returns the binary data of the extension.
	Data(...any) ([]byte, error)
	// Parse parses the binary data to the extension.
	Parse([]byte, ...any) error
}

// Extensions represents the ORPC Extensions object used
// to parse and generate the ORPC Extensions.
type Extensions struct {
	// DataRepresentation is needed for ContextExtension.
	DataRepresentation ndr.DataRepresentation
	// Entries is the list of extensions.
	Entries []Extension
}

// ParseWithDataRepresentation parses the ORPCExtentArray with the given data representation.
func (o *Extensions) ParseWithDataRepresentation(array *ORPCExtentArray, drep ndr.DataRepresentation) error {
	o.DataRepresentation = drep
	return o.Parse(array)
}

// Parse parses the ORPCExtentArray.
func (o *Extensions) Parse(array *ORPCExtentArray) error {

	for i := range array.Extent {
		if array.Extent[i] == nil {
			continue
		}

		var ext Extension

		switch id := (*ClassID)(array.Extent[i].ID); {
		case id.Equal(ErrorExtensionClassID):
			ext = &ErrorExtension{}
		case id.Equal(ContextExtensionClassID):
			ext = &ContextExtension{}
		default:
			continue
		}

		if err := ext.Parse(array.Extent[i].Data, o.DataRepresentation); err != nil {
			return fmt.Errorf("parse extension error: %w", err)
		}

		o.Entries = append(o.Entries, ext)
	}

	return nil
}

// ORPCExtentArray function marshals the extensions list to ORPCExtentArray.
func (o *Extensions) ORPCExtentArray() (*ORPCExtentArray, error) {

	ret := make([]*ORPCExtent, 0)

	for i := range o.Entries {
		if o.Entries[i] == nil {
			continue
		}

		data, err := o.Entries[i].Data(o.DataRepresentation)
		if err != nil {
			return nil, fmt.Errorf("data error: %w", err)
		}

		ret = append(ret, &ORPCExtent{
			ID:   o.Entries[i].ExtensionID().GUID(),
			Data: data,
		})
	}

	return &ORPCExtentArray{
		Extent: ret,
	}, nil
}

// ErrorExtension represents the Error Extension object.
// (see ErrorObjectData).
type ErrorExtension struct {
	HelpContext uint32
	IID         *IID
	Source      string
	Description string
	HelpFile    string
}

func (ErrorExtension) ExtensionID() *ClassID {
	return ErrorExtensionClassID
}

func (o *ErrorExtension) Data(_ ...any) ([]byte, error) {

	ref, err := o.ObjectReference()
	if err != nil {
		return nil, fmt.Errorf("object reference error: %w", err)
	}

	blob, err := ndr.Marshal(ref, ndr.Opaque)
	if err != nil {
		return nil, fmt.Errorf("marshal object reference error: %w", err)
	}

	return blob, nil
}

func (o *ErrorExtension) Parse(b []byte, _ ...any) error {

	var obj ObjectReference

	if err := ndr.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("unmarshal object reference error: %w", err)
	}

	cref, ok := obj.ObjectReference.GetValue().(*ObjectReferenceCustom)
	if !ok {
		return fmt.Errorf("object reference is invalid")
	}

	if !cref.ClassID.Equal(ErrorObjectClassID) {
		return fmt.Errorf("object reference is not error")
	}

	if err := o.Unmarshal(cref.ObjectData); err != nil {
		return fmt.Errorf("unmarshal error object: %w", err)
	}

	return nil
}

func (o *ErrorExtension) ObjectReference() (*ObjectReference, error) {

	blob, err := o.Marshal()
	if err != nil {
		return nil, fmt.Errorf("marshal error object: %w", err)
	}

	return &ObjectReference{
		Signature: ([]byte)(ObjectReferenceCustomSignature),
		Flags:     ObjectReferenceTypeCustom,
		IID:       o.IID,
		ObjectReference: &ObjectReference_ObjectReference{
			Value: &ObjectReference_Custom{
				Custom: &ObjectReferenceCustom{
					ClassID:    ErrorObjectClassID,
					Size:       uint32(len(blob)),
					ObjectData: blob,
				},
			},
		},
	}, nil

}

func (o *ErrorExtension) Marshal() ([]byte, error) {

	obj := &ErrorObjectData{
		HelpContext: o.HelpContext,
		IID:         o.IID,
	}

	if o.Source != "" {
		obj.SourceSignature, obj.Source = ErrorStringSignature, &ErrorObjectDataString{
			Value: &ErrorObjectDataString_String{
				String: &ErrorInfoString{Name: o.Source},
			},
		}
	}

	if o.Description != "" {
		obj.DescriptionSignature, obj.Description = ErrorStringSignature, &ErrorObjectDataString{
			Value: &ErrorObjectDataString_String{
				String: &ErrorInfoString{Name: o.Description},
			},
		}
	}

	if o.HelpFile != "" {
		obj.HelpFileSignature, obj.HelpFile = ErrorStringSignature, &ErrorObjectDataString{
			Value: &ErrorObjectDataString_String{
				String: &ErrorInfoString{Name: o.HelpFile},
			},
		}
	}

	b, err := ndr.Marshal(obj, ndr.Opaque)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (o *ErrorExtension) Unmarshal(b []byte) error {

	var obj ErrorObjectData

	if err := ndr.Unmarshal(b, &obj, ndr.Opaque); err != nil {
		return err
	}

	if val, ok := obj.Source.GetValue().(*ErrorInfoString); ok {
		o.Source = val.Name
	}

	if val, ok := obj.Description.GetValue().(*ErrorInfoString); ok {
		o.Description = val.Name
	}

	if val, ok := obj.HelpFile.GetValue().(*ErrorInfoString); ok {
		o.HelpFile = val.Name
	}

	o.HelpContext = obj.HelpContext
	o.IID = obj.IID

	return nil
}

// ContextEntry represents the ORPC Context Extension entry.
type ContextEntry struct {
	// ID is the PolicyID.
	ID *dtyp.GUID
	// Data is the binary data.
	Data []byte
}

// ContextExtension represents the ORPC Context Extension object.
type ContextExtension struct {
	ServerHResult uint32
	Entries       []*ContextEntry
}

func (ContextExtension) ExtensionID() *ClassID {
	return ContextExtensionClassID
}

func (o *ContextExtension) Data(opts ...any) ([]byte, error) {

	ctx := &ContextORPCExtension{
		Signature:     ([]byte)(ContextORPCExtensionSignature),
		Version:       0x00010000,
		ServerHResult: o.ServerHResult,
		Length:        uint32(32 + 32*len(o.Entries)), /* header size + entry header sizes */
	}

	buf := bytes.NewBuffer(nil)

	for i1 := range o.Entries {

		ctx.EntryHeader = append(ctx.EntryHeader, &EntryHeader{
			Signature:    ([]byte)(EntryHeaderSignature),
			BufferLength: uint32(len(o.Entries[i1].Data)),
			Length:       ctx.Length,
			PolicyID:     o.Entries[i1].ID,
		})

		// padded to 16 bytes.
		ctx.Length += uint32(len(o.Entries[i1].Data)+15) & (^uint32(15))
		// write entry data to buffer.
		if _, err := buf.Write(o.Entries[i1].Data); err != nil {
			return nil, err
		}
		// write pad.
		for i2 := uint32(len(o.Entries[i1].Data)); i2 < ctx.Length; i2++ {
			if err := buf.WriteByte(0); err != nil {
				return nil, err
			}
		}
	}

	ctx.PolicyData = buf.Bytes()
	ctx.BufferLength = uint32(len(ctx.PolicyData))

	b, err := ndr.Marshal(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (o *ContextExtension) Parse(b []byte, opts ...any) error {

	var ctx ContextORPCExtension

	if err := ndr.Unmarshal(b, &ctx, opts...); err != nil {
		return fmt.Errorf("unmarshal context extension error: %w", err)
	}

	o.ServerHResult = ctx.ServerHResult

	hdrSize := 32 + 32*len(ctx.EntryHeader)

	for i1 := range ctx.EntryHeader {

		eOft, eLen := int(ctx.EntryHeader[i1].Length)-(hdrSize), int(ctx.EntryHeader[i1].BufferLength)

		if eOft <= 0 || eOft+eLen > len(ctx.PolicyData) {
			return fmt.Errorf("entry data is invalid")
		}
		o.Entries = append(o.Entries, &ContextEntry{
			ID:   ctx.EntryHeader[i1].PolicyID,
			Data: ctx.PolicyData[eOft : eOft+eLen],
		})
	}

	return nil
}
