package idispatch

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IDispatch server interface.
type DispatchServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetTypeInfoCount method provides the number of automation type descriptions in
	// the type information table.
	//
	// The GetTypeInfoCount method specifies whether the automation server provides Type
	// description information.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeInfoCount(context.Context, *GetTypeInfoCountRequest) (*GetTypeInfoCountResponse, error)

	// The GetTypeInfo method provides access to the Type description information exposed
	// by the automation server.
	//
	// The GetTypeInfo method retrieves the automation type description that has the specified
	// ordinal position within the type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) DWORD does not match the value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches the value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002000B
	//
	// DISP_E_BADINDEX
	//
	// SHOULD be returned when the value of the passed in iTInfo argument was not 0. See
	// [MS-ERREF].
	GetTypeInfo(context.Context, *GetTypeInfoRequest) (*GetTypeInfoResponse, error)

	// The GetIDsOfNames method maps a single member (method or property) name, and an optional
	// set of argument names, to a corresponding set of integer DISPIDs, which can be used
	// on subsequent calls to IDispatch::Invoke.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, as
	// defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies the
	// following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x80020006
	//
	// DISP_E_UNKNOWNNAME
	//
	// One or more names were not known. The returned array of DISPIDs MUST contain at least
	// one DISPID_UNKNOWN, and there MUST be one DISPID_UNKNOWN for each entry that corresponds
	// to an unknown name. See [MS-ERREF].
	//
	// 0x80020001
	//
	// DISP_E_UNKNOWNINTERFACE
	//
	// The interface identifier passed in riid is not IID_NULL. See [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown from this method except those that are
	// thrown by the underlying RPC Protocol specified in [MS-RPCE].
	//
	// When GetIDsOfNames is called with more than one name, the first name (rgszNames[0])
	// corresponds to the member name, and subsequent names correspond to the names of member
	// parameters.
	//
	// The same name can map to different DISPIDs, depending on context. For example, a
	// name can have a DISPID when it is used as: a member name with a particular interface,
	// a different ID as a member of a different interface, or a different mapping for each
	// time it appears as a parameter.
	//
	// The implementation of GetIDsOfNames MUST be case-insensitive.
	//
	// An implementation of the OLE Automation Protocol MAY<55> choose to implement a mapping
	// for the parameter names that map to the index of the parameter in the member parameter
	// list.
	GetIDsOfNames(context.Context, *GetIDsOfNamesRequest) (*GetIDsOfNamesResponse, error)

	// The Invoke method provides access to properties and methods exposed by the automation
	// server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x80020009
	//
	// DISP_E_EXCEPTION
	//
	// The application needs to raise an exception. In this case, the structure passed in
	// pExcepInfo MUST be filled in with a nonzero error code. See [MS-ERREF].
	//
	// 0x80020004
	//
	// DISP_E_PARAMNOTFOUND
	//
	// One of the parameter DISPIDs does not correspond to a parameter on the method. In
	// this case, pArgErr MUST be set to the first argument that contains the error. See
	// [MS-ERREF].
	//
	// 0x80020005
	//
	// DISP_E_TYPEMISMATCH
	//
	// One or more of the arguments could not be coerced into the type of the corresponding
	// formal parameter. The index within rgvarg of the first parameter with the incorrect
	// type MUST be returned in the pArgErr parameter. For more information, see section
	// 3.1.4.4.4 ( 5c01ab3c-f719-44cc-bb45-d36850cf4c5b ) and [MS-ERREF].
	//
	// 0x8002000E
	//
	// DISP_E_BADPARAMCOUNT
	//
	// The number of elements provided to DISPPARAMS is different from the number of arguments
	// accepted by the method or property. See [MS-ERREF].
	//
	// 0x80020008
	//
	// DISP_E_BADVARTYPE
	//
	// One of the arguments in rgvarg is not a valid variant type. See [MS-ERREF].
	//
	// 0x80020003
	//
	// DISP_E_MEMBERNOTFOUND
	//
	// The requested member does not exist, or the call to Invoke tried to set the value
	// of a read-only property. See [MS-ERREF].
	//
	// 0x80020007
	//
	// DISP_E_NONAMEDARGS
	//
	// At least one named argument was provided for methods with a vararg parameter (see
	// section 3.1.4.4.3 ( be6e35f6-9327-4164-9bde-ffcd0fa0e07d ) ), for which named arguments
	// are illegal. See [MS-ERREF].
	//
	// 0x8002000A
	//
	// DISP_E_OVERFLOW
	//
	// One of the arguments in rgvarg could not be coerced to the type of its corresponding
	// formal parameter. See [MS-ERREF].
	//
	// 0x80020001
	//
	// DISP_E_UNKNOWNINTERFACE
	//
	// The interface identifier passed in riid is not IID_NULL. See [MS-ERREF].
	//
	// 0x8002000F
	//
	// DISP_E_PARAMNOTOPTIONAL
	//
	// A required parameter was omitted. See [MS-ERREF].
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC Protocol, as specified in [MS-RPCE].
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
}

func RegisterDispatchServer(conn dcerpc.Conn, o DispatchServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDispatchServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DispatchSyntaxV0_0))...)
}

func NewDispatchServerHandle(o DispatchServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DispatchServerHandle(ctx, o, opNum, r)
	}
}

func DispatchServerHandle(ctx context.Context, o DispatchServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetTypeInfoCount
		op := &xxx_GetTypeInfoCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeInfoCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeInfoCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetTypeInfo
		op := &xxx_GetTypeInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetIDsOfNames
		op := &xxx_GetIDsOfNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIDsOfNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIDsOfNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // Invoke
		op := &xxx_InvokeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InvokeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Invoke(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDispatch
type UnimplementedDispatchServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedDispatchServer) GetTypeInfoCount(context.Context, *GetTypeInfoCountRequest) (*GetTypeInfoCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDispatchServer) GetTypeInfo(context.Context, *GetTypeInfoRequest) (*GetTypeInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDispatchServer) GetIDsOfNames(context.Context, *GetIDsOfNamesRequest) (*GetIDsOfNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDispatchServer) Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DispatchServer = (*UnimplementedDispatchServer)(nil)
