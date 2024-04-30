package dcom

import "github.com/oiweiwei/go-msrpc/ndr"

func (o *InterfacePointer) IPID() *IPID {
	if std := o.GetStandardObjectReference().Std; std != nil && std.IPID != nil {
		return std.IPID
	}
	return &IPID{}
}

func (o *InterfacePointer) GetObjectReference() *ObjectReference {
	ref := &ObjectReference{}
	if len(o.Data) > 0 {
		ndr.Unmarshal(o.Data, ref, ndr.Opaque)
	}
	return ref
}

func (o *InterfacePointer) GetStandardObjectReference() *ObjectReferenceStandard {
	std, _ := o.GetObjectReference().ObjectReference.GetValue().(*ObjectReferenceStandard)
	if std != nil {
		return std
	}
	return &ObjectReferenceStandard{}
}

func (o *InterfacePointer) GetCustomObjectReference() *ObjectReferenceCustom {
	custom, _ := o.GetObjectReference().ObjectReference.GetValue().(*ObjectReferenceCustom)
	if custom != nil {
		return custom
	}
	return &ObjectReferenceCustom{}
}
