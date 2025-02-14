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
		op := &xxx_GetExecutionOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExecutionOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExecutionOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // ExecutionOption
		op := &xxx_SetExecutionOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExecutionOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExecutionOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // PropertyAffected
		op := &xxx_GetPropertyAffectedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertyAffectedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertyAffected(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // PropertyAffected
		op := &xxx_SetPropertyAffectedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPropertyAffectedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPropertyAffected(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Value
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // Value
		op := &xxx_SetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmClassificationRule
type UnimplementedClassificationRuleServer struct {
	ifsrmrule.UnimplementedRuleServer
}

func (UnimplementedClassificationRuleServer) GetExecutionOption(context.Context, *GetExecutionOptionRequest) (*GetExecutionOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassificationRuleServer) SetExecutionOption(context.Context, *SetExecutionOptionRequest) (*SetExecutionOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassificationRuleServer) GetPropertyAffected(context.Context, *GetPropertyAffectedRequest) (*GetPropertyAffectedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassificationRuleServer) SetPropertyAffected(context.Context, *SetPropertyAffectedRequest) (*SetPropertyAffectedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassificationRuleServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClassificationRuleServer) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClassificationRuleServer = (*UnimplementedClassificationRuleServer)(nil)
