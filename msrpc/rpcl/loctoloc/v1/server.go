package loctoloc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
)

// LocToLoc server interface.
type LocToLocServer interface {

	// The I_nsi_lookup_begin method is invoked by a client locator to enumerate the binding
	// information for a set of RPC servers that satisfy a given set of criteria. The Microsoft
	// Interface Definition Language (MIDL) syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions might be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupBegin(context.Context, *LookupBeginRequest) (*LookupBeginResponse, error)

	// The I_nsi_lookup_done method is invoked to free any resources associated with the
	// context handle returned by a preceding call to the I_nsi_lookup_begin method. The
	// MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupDone(context.Context, *LookupDoneRequest) (*LookupDoneResponse, error)

	// The I_nsi_lookup_next method is invoked to continue an enumeration of binding vectors
	// that satisfy the criteria specified in a call to the I_nsi_lookup_begin method. The
	// number of bindings in the binding_vector is limited by the parameter binding_max_count
	// specified in the call to the I_nsi_lookup_begin method. The MIDL syntax of this method
	// is specified as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	LookupNext(context.Context, *LookupNextRequest) (*LookupNextResponse, error)

	// The I_nsi_entry_object_inq_next method is invoked to continue an enumeration initiated
	// by a previous call to the I_nsi_entry_object_inq_begin method. The MIDL syntax of
	// the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireNext(context.Context, *EntryObjectInquireNextRequest) (*EntryObjectInquireNextResponse, error)

	// The I_nsi_ping_locator method is invoked by the client to determine if the target
	// computer is available as a master locator. The MIDL syntax of the method is specified
	// as follows.
	//
	// Return Values: This method does not return any values.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	PingLocator(context.Context, *PingLocatorRequest) (*PingLocatorResponse, error)

	// The I_nsi_entry_object_inq_done method is invoked to free any resources associated
	// with the context handle returned by a preceding call to the I_nsi_entry_object_inq_begin
	// method. The MIDL syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireDone(context.Context, *EntryObjectInquireDoneRequest) (*EntryObjectInquireDoneResponse, error)

	// The I_nsi_entry_object_inq_begin method is invoked to enumerate the object UUIDs
	// on a name service entry. The MIDL syntax of the method is specified as follows.
	//
	// Return Values: This method does not return any values. RPC exceptions can be thrown
	// from this method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	EntryObjectInquireBegin(context.Context, *EntryObjectInquireBeginRequest) (*EntryObjectInquireBeginResponse, error)
}

func RegisterLocToLocServer(conn dcerpc.Conn, o LocToLocServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLocToLocServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LocToLocSyntaxV1_0))...)
}

func NewLocToLocServerHandle(o LocToLocServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LocToLocServerHandle(ctx, o, opNum, r)
	}
}

func LocToLocServerHandle(ctx context.Context, o LocToLocServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // I_nsi_lookup_begin
		in := &LookupBeginRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupBegin(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // I_nsi_lookup_done
		in := &LookupDoneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupDone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // I_nsi_lookup_next
		in := &LookupNextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupNext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // I_nsi_entry_object_inq_next
		in := &EntryObjectInquireNextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EntryObjectInquireNext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // I_nsi_ping_locator
		in := &PingLocatorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PingLocator(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // I_nsi_entry_object_inq_done
		in := &EntryObjectInquireDoneRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EntryObjectInquireDone(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // I_nsi_entry_object_inq_begin
		in := &EntryObjectInquireBeginRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EntryObjectInquireBegin(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
