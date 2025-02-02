package dcom

import (
	"bytes"
	"fmt"

	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

// Activation property CLSID for ActivationContextInfoData. {000001a5-0000-0000-c000-000000000046}
var ActivationContextInfoClassID = &ClassID{Data1: 0x000001a5, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// OBJREF_CUSTOM unmarshaler CLSID for ActivationPropertiesIn. {00000338-0000-0000-c000-000000000046}
var ActivationPropertiesInClassID = &ClassID{Data1: 0x00000338, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// OBJREF_CUSTOM unmarshaler CLSID for ActivationPropertiesOut. {00000339-0000-0000-c000-000000000046}
var ActivationPropertiesOutClassID = &ClassID{Data1: 0x00000339, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// ORPC_EXTENT identifier for context (2) ORPC extension. {00000334-0000-0000-c000-000000000046}
var ContextExtensionClassID = &ClassID{Data1: 0x00000334, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// OBJREF_CUSTOM unmarshaler CLSID for contexts (2). {0000033b-0000-0000-c000-000000000046}
var ContextMarshalerClassID = &ClassID{Data1: 0x0000033b, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// ORPC_EXTENT identifier for Error information ORPC extension. {0000031c-0000-0000-c000-000000000046}
var ErrorExtensionClassID = &ClassID{Data1: 0x0000031c, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// OBJREF_CUSTOM unmarshaler CLSID for error information. {0000031b-0000-0000-c000-000000000046}
var ErrorObjectClassID = &ClassID{Data1: 0x0000031b, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for InstanceInfoData. {000001ad-0000-0000-c000-000000000046}
var InstanceInfoClassID = &ClassID{Data1: 0x000001ad, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for InstantiationInfoData.	{000001ab-0000-0000-c000-000000000046}
var InstantiationInfoClassID = &ClassID{Data1: 0x000001ab, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for PropsOutInfo. {00000339-0000-0000-c000-000000000046}
var PropertiesOutInfoClassID = &ClassID{Data1: 0x00000339, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for ScmReplyInfoData. {000001b6-0000-0000-c000-000000000046}
var SCMReplyInfoClassID = &ClassID{Data1: 0x000001b6, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for ScmRequestInfoData. {000001aa-0000-0000-c000-000000000046}
var SCMRequestInfoClassID = &ClassID{Data1: 0x000001aa, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for SecurityInfoData. {000001a6-0000-0000-c000-000000000046}
var SecurityInfoClassID = &ClassID{Data1: 0x000001a6, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for LocationInfoData. {000001a4-0000-0000-c000-000000000046}
var ServerLocationClassID = &ClassID{Data1: 0x000001a4, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property CLSID for SpecialPropertiesData. {000001b9-0000-0000-c000-000000000046}
var SpecialPropertiesClassID = &ClassID{Data1: 0x000001b9, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// Activation property interface.
type ActivationProperty interface {
	ndr.Marshaler
	ndr.Unmarshaler
	// ActivationPropertyClassID returns class id of the activation property.
	ActivationPropertyClassID() *ClassID
}

func (InstantiationInfoData) ActivationPropertyClassID() *ClassID {
	return InstantiationInfoClassID
}

func (SpecialPropertiesData) ActivationPropertyClassID() *ClassID {
	return SpecialPropertiesClassID
}

func (InstanceInfoData) ActivationPropertyClassID() *ClassID {
	return InstanceInfoClassID
}

func (SCMRequestInfoData) ActivationPropertyClassID() *ClassID {
	return SCMRequestInfoClassID
}

func (ActivationContextInfoData) ActivationPropertyClassID() *ClassID {
	return ActivationContextInfoClassID
}

func (LocationInfoData) ActivationPropertyClassID() *ClassID {
	return ServerLocationClassID
}

func (SecurityInfoData) ActivationPropertyClassID() *ClassID {
	return SecurityInfoClassID
}

func (SCMReplyInfoData) ActivationPropertyClassID() *ClassID {
	return SCMReplyInfoClassID
}

func (PropertiesOutInfo) ActivationPropertyClassID() *ClassID {
	return PropertiesOutInfoClassID
}

var (
	_ ActivationProperty = (*InstantiationInfoData)(nil)
	_ ActivationProperty = (*SpecialPropertiesData)(nil)
	_ ActivationProperty = (*InstanceInfoData)(nil)
	_ ActivationProperty = (*SCMRequestInfoData)(nil)
	_ ActivationProperty = (*ActivationContextInfoData)(nil)
	_ ActivationProperty = (*LocationInfoData)(nil)
	_ ActivationProperty = (*SecurityInfoData)(nil)
	_ ActivationProperty = (*SCMReplyInfoData)(nil)
	_ ActivationProperty = (*PropertiesOutInfo)(nil)
)

// ActivationProperties structures represents the activation properties.
type ActivationProperties struct {
	// DestinationContext is the destination context.
	DestinationContext uint32
	// Properties is the list of activation properties.
	Properties []ActivationProperty
}

func (o *ActivationProperties) InstantiationInfoData() *InstanceInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*InstanceInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) SpecialPropertiesData() *SpecialPropertiesData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*SpecialPropertiesData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) InstanceInfoData() *InstanceInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*InstanceInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) SCMRequestInfoData() *SCMRequestInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*SCMRequestInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) ActivationContextInfoData() *ActivationContextInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*ActivationContextInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) LocationInfoData() *LocationInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*LocationInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) SecurityInfoData() *SecurityInfoData {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*SecurityInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) SCMReplyInfoData() *SCMReplyInfoData {
	for i := range o.Properties {

		if v, ok := o.Properties[i].(*SCMReplyInfoData); ok {
			return v
		}
	}
	return nil
}

func (o *ActivationProperties) PropertiesOutInfo() *PropertiesOutInfo {
	for i := range o.Properties {
		if v, ok := o.Properties[i].(*PropertiesOutInfo); ok {
			return v
		}
	}
	return nil
}

// ActivationPropertiesIn function creates InterfacePointer with ActivationPropertiesIn class id.
func (o *ActivationProperties) ActivationPropertiesIn() (*InterfacePointer, error) {
	return o.InterfacePointerWithClassID(ActivationPropertiesInClassID)
}

// ActivationPropertiesOut function creates InterfacePointer with ActivationPropertiesOut class id.
func (o *ActivationProperties) ActivationPropertiesOut() (*InterfacePointer, error) {
	return o.InterfacePointerWithClassID(ActivationPropertiesOutClassID)
}

// InterfacePointerWithClassID function creates InterfacePointer with specified class id.
func (o *ActivationProperties) InterfacePointerWithClassID(cls *ClassID) (*InterfacePointer, error) {

	if !cls.Equal(ActivationPropertiesInClassID) && !cls.Equal(ActivationPropertiesOutClassID) {
		return nil, fmt.Errorf("interface_pointer: invalid class id")
	}

	blob, err := o.Marshal()
	if err != nil {
		return nil, err
	}

	ref := &ObjectReference{
		Signature: ([]byte)(ObjectReferenceCustomSignature),
		Flags:     ObjectReferenceTypeCustom,
		IID:       ActivationPropertiesInIID,
		ObjectReference: &ObjectReference_ObjectReference{
			Value: &ObjectReference_Custom{
				Custom: &ObjectReferenceCustom{
					ClassID:    ActivationPropertiesInClassID,
					Size:       uint32(len(blob)) + 8, /* XXX: extra 8 bytes for unknown reason */
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

// Parse function parses ActivationProperties object from InterfacePointer.
func (o *ActivationProperties) Parse(in *InterfacePointer) error {

	ref := in.GetCustomObjectReference()

	if ref == nil {
		return fmt.Errorf("parse_interface_pointer: invalid object reference")
	}

	if !ref.ClassID.Equal(ActivationPropertiesOutClassID) && !ref.ClassID.Equal(ActivationPropertiesInClassID) {
		return fmt.Errorf("parse_interface_pointer: invalid class id")
	}

	if ref.ObjectData == nil {
		return fmt.Errorf("parse_interface_pointer: invalid object data")
	}

	if err := o.Unmarshal(ref.ObjectData); err != nil {
		return fmt.Errorf("parse_interface_pointer: %w", err)
	}

	return nil
}

// Marshal function marshals ActivationProperties object into Activation Properties BLOB.
func (o *ActivationProperties) Marshal() ([]byte, error) {

	w := bytes.NewBuffer(nil)

	header := &CustomHeader{
		DestinationContext: o.DestinationContext,
	}

	for i := range o.Properties {
		b, err := ndr.MarshalWithTypeSerializationV1(o.Properties[i], ndr.Pad(8))
		if err != nil {
			return nil, err
		}
		if _, err := w.Write(b); err != nil {
			return nil, fmt.Errorf("marshal activation_property: %w", err)
		}
		header.Sizes = append(header.Sizes, uint32(len(b)))
		header.ClassIDs = append(header.ClassIDs, o.Properties[i].ActivationPropertyClassID())
	}

	b, err := ndr.MarshalWithTypeSerializationV1(header, ndr.Pad(8))
	if err != nil {
		return nil, fmt.Errorf("marshal custom_header: %w", err)
	}

	header.HeaderSize = uint32(len(b))
	header.TotalSize = header.HeaderSize + uint32(w.Len())

	b, err = ndr.MarshalWithTypeSerializationV1(header, ndr.Pad(8))
	if err != nil {
		return nil, err
	}

	blob := &ActivationPropertiesBlob{
		Data: append(b, w.Bytes()...),
	}

	b, err = ndr.Marshal(blob, ndr.Opaque)
	if err != nil {
		return nil, fmt.Errorf("marshal activation_properties_blob: %w", err)
	}

	return b, nil
}

// Unmarshal function unmarshals Activation Properties BLOB into ActivationProperties object.
func (o *ActivationProperties) Unmarshal(b []byte) error {

	blob := &ActivationPropertiesBlob{}

	if err := ndr.Unmarshal(b, blob, ndr.Opaque); err != nil {
		return fmt.Errorf("unmarshal activation_properties_blob: %w", err)
	}

	b = blob.Data

	header := &CustomHeader{}

	if err := ndr.UnmarshalWithTypeSerializationV1(b, header); err != nil {
		return fmt.Errorf("unmarshal custom_header: %w", err)
	}

	if header.TotalSize != uint32(len(b)) {
		return fmt.Errorf("unmarshal custom_header: invalid total size")
	}

	if len(header.Sizes) != len(header.ClassIDs) {
		return fmt.Errorf("unmarshal custom_header: sizes and class ids mismatch")
	}

	o.DestinationContext = header.DestinationContext

	if len(header.Sizes) < 1 || len(header.Sizes) > 10 {
		return fmt.Errorf("unmarshal custom_header: invalid sizes count")
	}

	o.Properties = make([]ActivationProperty, 0, len(header.Sizes))

	for i, oft := 0, header.HeaderSize; i < len(header.Sizes); i, oft = i+1, oft+header.Sizes[i] {

		var prop ActivationProperty

		switch cls := header.ClassIDs[i]; {
		case cls.Equal(InstantiationInfoClassID):
			prop = &InstantiationInfoData{}
		case cls.Equal(SpecialPropertiesClassID):
			prop = &SpecialPropertiesData{}
		case cls.Equal(InstanceInfoClassID):
			prop = &InstanceInfoData{}
		case cls.Equal(SCMRequestInfoClassID):
			prop = &SCMRequestInfoData{}
		case cls.Equal(ActivationContextInfoClassID):
			prop = &ActivationContextInfoData{}
		case cls.Equal(ServerLocationClassID):
			prop = &LocationInfoData{}
		case cls.Equal(SecurityInfoClassID):
			prop = &SecurityInfoData{}
		case cls.Equal(SCMReplyInfoClassID):
			prop = &SCMReplyInfoData{}
		case cls.Equal(PropertiesOutInfoClassID):
			prop = &PropertiesOutInfo{}
		default:
			return fmt.Errorf("unmarshal activation_property: unknown class id: %s", cls.GUID())
		}

		if oft+header.Sizes[i] > uint32(len(b)) {
			return fmt.Errorf("unmarshal activation_property: out of range")
		}

		if err := ndr.UnmarshalWithTypeSerializationV1(b[oft:oft+header.Sizes[i]], prop); err != nil {
			return fmt.Errorf("unmarshal activation_property: %w", err)
		}

		o.Properties = append(o.Properties, prop)
	}

	return nil
}
