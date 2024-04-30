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
		in := &GetReportTypesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetReportTypes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ReportTypes
		in := &SetReportTypesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetReportTypes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // MailTo
		in := &GetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // MailTo
		in := &SetMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
