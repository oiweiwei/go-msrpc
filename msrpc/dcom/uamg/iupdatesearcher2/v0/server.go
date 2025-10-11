package iupdatesearcher2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatesearcher "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatesearcher/v0"
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
	_ = iupdatesearcher.GoPackage
)

// IUpdateSearcher2 server interface.
type UpdateSearcher2Server interface {

	// IUpdateSearcher base class.
	iupdatesearcher.UpdateSearcherServer

	// The IUpdateSearcher2::IgnoreDownloadPriority (opnum 26) method retrieves whether
	// the update download priority is ignored.
	//
	// The IUpdateSearcher2::IgnoreDownloadPriority (opnum 27) method sets whether the update
	// download priority is ignored.
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
	// This method SHOULD return the value of the IgnoreDownloadPriority ADM element.
	GetIgnoreDownloadPriority(context.Context, *GetIgnoreDownloadPriorityRequest) (*GetIgnoreDownloadPriorityResponse, error)

	// The IUpdateSearcher2::IgnoreDownloadPriority (opnum 26) method retrieves whether
	// the update download priority is ignored.
	//
	// The IUpdateSearcher2::IgnoreDownloadPriority (opnum 27) method sets whether the update
	// download priority is ignored.
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
	// This method SHOULD return the value of the IgnoreDownloadPriority ADM element.
	SetIgnoreDownloadPriority(context.Context, *SetIgnoreDownloadPriorityRequest) (*SetIgnoreDownloadPriorityResponse, error)
}

func RegisterUpdateSearcher2Server(conn dcerpc.Conn, o UpdateSearcher2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSearcher2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSearcher2SyntaxV0_0))...)
}

func NewUpdateSearcher2ServerHandle(o UpdateSearcher2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSearcher2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSearcher2ServerHandle(ctx context.Context, o UpdateSearcher2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 25 {
		// IUpdateSearcher base method.
		return iupdatesearcher.UpdateSearcherServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 25: // IgnoreDownloadPriority
		op := &xxx_GetIgnoreDownloadPriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIgnoreDownloadPriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIgnoreDownloadPriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // IgnoreDownloadPriority
		op := &xxx_SetIgnoreDownloadPriorityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetIgnoreDownloadPriorityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetIgnoreDownloadPriority(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSearcher2
type UnimplementedUpdateSearcher2Server struct {
	iupdatesearcher.UnimplementedUpdateSearcherServer
}

func (UnimplementedUpdateSearcher2Server) GetIgnoreDownloadPriority(context.Context, *GetIgnoreDownloadPriorityRequest) (*GetIgnoreDownloadPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcher2Server) SetIgnoreDownloadPriority(context.Context, *SetIgnoreDownloadPriorityRequest) (*SetIgnoreDownloadPriorityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSearcher2Server = (*UnimplementedUpdateSearcher2Server)(nil)
