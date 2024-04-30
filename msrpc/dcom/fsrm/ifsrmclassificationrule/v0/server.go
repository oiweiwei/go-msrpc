package ifsrmclassificationrule

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmrule "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmrule/v0"
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
	_ = ifsrmrule.GoPackage
)

// IFsrmClassificationRule server interface.
type ClassificationRuleServer interface {

	// IFsrmRule base class.
	ifsrmrule.RuleServer

	// ExecutionOption operation.
	GetExecutionOption(context.Context, *GetExecutionOptionRequest) (*GetExecutionOptionResponse, error)

	// ExecutionOption operation.
	SetExecutionOption(context.Context, *SetExecutionOptionRequest) (*SetExecutionOptionResponse, error)

	// PropertyAffected operation.
	GetPropertyAffected(context.Context, *GetPropertyAffectedRequest) (*GetPropertyAffectedResponse, error)

	// PropertyAffected operation.
	SetPropertyAffected(context.Context, *SetPropertyAffectedRequest) (*SetPropertyAffectedResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)
}

func RegisterClassificationRuleServer(conn dcerpc.Conn, o ClassificationRuleServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClassificationRuleServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClassificationRuleSyntaxV0_0))...)
}

func NewClassificationRuleServerHandle(o ClassificationRuleServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClassificationRuleServerHandle(ctx, o, opNum, r)
	}
}

func ClassificationRuleServerHandle(ctx context.Context, o ClassificationRuleServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 24 {
		// IFsrmRule base method.
		return ifsrmrule.RuleServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 24: // ExecutionOption
		in := &GetExecutionOptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExecutionOption(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // ExecutionOption
		in := &SetExecutionOptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExecutionOption(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // PropertyAffected
		in := &GetPropertyAffectedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyAffected(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // PropertyAffected
		in := &SetPropertyAffectedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPropertyAffected(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Value
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
