package dcom

import (
	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

// Superclass option is an alias of the dcerpc.WithNoBind option.
func Superclass(cc dcerpc.Conn) dcerpc.Option {
	return dcerpc.WithNoBind(cc)
}

// IsSuperclass option is an alias of the dcerpc.HasNoBind
// function.
func IsSuperclass(opts any) bool {
	_, ok := dcerpc.HasNoBind(opts)
	return ok
}

// WithIPID option returns the ObjectUUIDOption.
func WithIPID(ipid *IPID) dcerpc.ObjectUUIDOption {
	if ipid == nil {
		ipid = &IPID{}
	}
	return dcerpc.WithObjectUUID(ipid.UUID())
}

// HasIPID function returns the ObjectUUID casted to IPID.
func HasIPID(opts any) (*IPID, bool) {
	if u, ok := dcerpc.HasObjectUUID(opts); ok {
		return (*IPID)(dtyp.GUIDFromUUID(u)), true
	}
	return nil, false
}
