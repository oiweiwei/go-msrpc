package iwbemobjectsink

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

// IWbemObjectSink server interface.
type ObjectSinkServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// When the IWbemObjectSink::Indicate method is called, the apObjArray parameter MUST
	// contain additional result objects as an array of an IWbemClassObject, sent by the
	// client to the server. The IWbemObjectSink::Indicate method has the following syntax,
	// expressed in Microsoft Interface Definition Language (MIDL).
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call.
	//
	// When the IWbemObjectSink::Indicate method is called for the first time, the server
	// that implements the ObjectArray structure MUST return WBEM_S_NEW_STYLE if the execution
	// of the method succeeds. If a server does not implement the ObjectArray structure,
	// it MUST return WBEM_S_NO_ERROR for successful execution of the method.
	//
	// If the server implements the ObjectArray structure and WBEM_S_NEW_STYLE is returned
	// during the first call to the IWbemObjectSink::Indicate method, the server MUST support
	// subsequent calls to the IWbemObjectSink::Indicate method by using both the DCOM Remote
	// Protocol marshaling and the ObjectArray structure as specified in section 2.2.14.
	Indicate(context.Context, *IndicateRequest) (*IndicateResponse, error)

	// When the IWbemObjectSink::SetStatus method is called, the parameter MUST contain
	// the result of the operation or the progress status information.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	SetStatus(context.Context, *SetStatusRequest) (*SetStatusResponse, error)
}

func RegisterObjectSinkServer(conn dcerpc.Conn, o ObjectSinkServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectSinkServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectSinkSyntaxV0_0))...)
}

func NewObjectSinkServerHandle(o ObjectSinkServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectSinkServerHandle(ctx, o, opNum, r)
	}
}

func ObjectSinkServerHandle(ctx context.Context, o ObjectSinkServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Indicate
		op := &xxx_IndicateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IndicateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Indicate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetStatus
		op := &xxx_SetStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemObjectSink
type UnimplementedObjectSinkServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedObjectSinkServer) Indicate(context.Context, *IndicateRequest) (*IndicateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectSinkServer) SetStatus(context.Context, *SetStatusRequest) (*SetStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ObjectSinkServer = (*UnimplementedObjectSinkServer)(nil)
