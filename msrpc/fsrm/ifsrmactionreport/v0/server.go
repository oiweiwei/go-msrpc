package ifsrmactionreport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
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
	_ = ifsrmaction.GoPackage
)

// IFsrmActionReport server interface.
type ActionReportServer interface {

	// IFsrmAction base class.
	ifsrmaction.ActionServer

	// ReportTypes operation.
	GetReportTypes(context.Context, *GetReportTypesRequest) (*GetReportTypesResponse, error)

	// ReportTypes operation.
	SetReportTypes(context.Context, *SetReportTypesRequest) (*SetReportTypesResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error)
}

func RegisterActionReportServer(conn dcerpc.Conn, o ActionReportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionReportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionReportSyntaxV0_0))...)
}

func NewActionReportServerHandle(o ActionReportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionReportServerHandle(ctx, o, opNum, r)
	}
}

func ActionReportServerHandle(ctx context.Context, o ActionReportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmAction base method.
		return ifsrmaction.ActionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // ReportTypes
		op := &xxx_GetReportTypesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportTypesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReportTypes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ReportTypes
		op := &xxx_SetReportTypesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReportTypesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReportTypes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // MailTo
		op := &xxx_GetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // MailTo
		op := &xxx_SetMailToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMailToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMailTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmActionReport
type UnimplementedActionReportServer struct {
	ifsrmaction.UnimplementedActionServer
}

func (UnimplementedActionReportServer) GetReportTypes(context.Context, *GetReportTypesRequest) (*GetReportTypesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionReportServer) SetReportTypes(context.Context, *SetReportTypesRequest) (*SetReportTypesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionReportServer) GetMailTo(context.Context, *GetMailToRequest) (*GetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionReportServer) SetMailTo(context.Context, *SetMailToRequest) (*SetMailToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ActionReportServer = (*UnimplementedActionReportServer)(nil)
