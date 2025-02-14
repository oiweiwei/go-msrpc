package iwbemcallresult

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

// IWbemCallResult server interface.
type CallResultServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// When the IWbemCallResult::GetResultObject method is called, the server MUST attempt
	// to retrieve a CIM object from a previous semisynchronous operation call to the IWbemServices::GetObject
	// method or the IWbemServices::ExecMethod method. The entry in WbemCallResultTable
	// with WbemCallResultPointer pointing to IWbemCallResult is used to identify the previous
	// semisynchronous call.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultObject(context.Context, *GetResultObjectRequest) (*GetResultObjectResponse, error)

	// When the IWbemCallResult::GetResultString method is called, the server MUST return
	// the assigned CIM path of a CIM instance that was created by the IWbemServices::PutInstance
	// method that returned IWbemCallResult in the ppCallResult parameter.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultString(context.Context, *GetResultStringRequest) (*GetResultStringResponse, error)

	// When the IWbemCallResult::GetResultServices method is called, the server MUST retrieve
	// a pointer to the IWbemServices interface that results from a semisynchronous call
	// to the IWbemServices::OpenNamespace method.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetResultServices(context.Context, *GetResultServicesRequest) (*GetResultServicesResponse, error)

	// When the IWbemCallResult::GetCallStatus method is invoked, the server MUST return
	// the status of the current outstanding semisynchronous call.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetCallStatus(context.Context, *GetCallStatusRequest) (*GetCallStatusResponse, error)
}

func RegisterCallResultServer(conn dcerpc.Conn, o CallResultServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCallResultServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CallResultSyntaxV0_0))...)
}

func NewCallResultServerHandle(o CallResultServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CallResultServerHandle(ctx, o, opNum, r)
	}
}

func CallResultServerHandle(ctx context.Context, o CallResultServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetResultObject
		op := &xxx_GetResultObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResultObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetResultString
		op := &xxx_GetResultStringOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultStringRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResultString(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetResultServices
		op := &xxx_GetResultServicesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultServicesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResultServices(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetCallStatus
		op := &xxx_GetCallStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCallStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCallStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemCallResult
type UnimplementedCallResultServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCallResultServer) GetResultObject(context.Context, *GetResultObjectRequest) (*GetResultObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCallResultServer) GetResultString(context.Context, *GetResultStringRequest) (*GetResultStringResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCallResultServer) GetResultServices(context.Context, *GetResultServicesRequest) (*GetResultServicesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCallResultServer) GetCallStatus(context.Context, *GetCallStatusRequest) (*GetCallStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CallResultServer = (*UnimplementedCallResultServer)(nil)
