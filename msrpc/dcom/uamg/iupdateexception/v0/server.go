package iupdateexception

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IUpdateException server interface.
type UpdateExceptionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IUpdateException::Message (opnum 8) method retrieves the message for this exception.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Message ADM element.
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)

	// The IUpdateException::HResult (opnum 9) method retrieves the HRESULT describing the
	// error.
	//
	// The IUpdateHistoryEntry::HResult (opnum 10) method retrieves the HRESULT from the
	// operation.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the HResult ADM element.
	GetHResult(context.Context, *GetHResultRequest) (*GetHResultResponse, error)

	// The IUpdateException::Context (opnum 10) method retrieves the context for the exception.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Context ADM element.
	GetContext(context.Context, *GetContextRequest) (*GetContextResponse, error)
}

func RegisterUpdateExceptionServer(conn dcerpc.Conn, o UpdateExceptionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateExceptionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateExceptionSyntaxV0_0))...)
}

func NewUpdateExceptionServerHandle(o UpdateExceptionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateExceptionServerHandle(ctx, o, opNum, r)
	}
}

func UpdateExceptionServerHandle(ctx context.Context, o UpdateExceptionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Message
		op := &xxx_GetMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // HResult
		op := &xxx_GetHResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Context
		op := &xxx_GetContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateException
type UnimplementedUpdateExceptionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateExceptionServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateExceptionServer) GetHResult(context.Context, *GetHResultRequest) (*GetHResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateExceptionServer) GetContext(context.Context, *GetContextRequest) (*GetContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateExceptionServer = (*UnimplementedUpdateExceptionServer)(nil)
