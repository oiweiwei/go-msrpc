package intmsobjectmanagement3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	intmsobjectmanagement2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement2/v0"
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
	_ = intmsobjectmanagement2.GoPackage
)

// INtmsObjectManagement3 server interface.
type ObjectManagement3Server interface {

	// INtmsObjectManagement2 base class.
	intmsobjectmanagement2.ObjectManagement2Server

	// The GetNtmsObjectAttributeAR method retrieves private data from an object, with strings
	// encoded using ASCII.
	GetNTMSObjectAttributeAR(context.Context, *GetNTMSObjectAttributeARRequest) (*GetNTMSObjectAttributeARResponse, error)

	// The GetNtmsObjectAttributeWR method retrieves private data from an object, with strings
	// encoded using Unicode.
	GetNTMSObjectAttributeWR(context.Context, *GetNTMSObjectAttributeWRRequest) (*GetNTMSObjectAttributeWRResponse, error)
}

func RegisterObjectManagement3Server(conn dcerpc.Conn, o ObjectManagement3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewObjectManagement3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ObjectManagement3SyntaxV0_0))...)
}

func NewObjectManagement3ServerHandle(o ObjectManagement3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ObjectManagement3ServerHandle(ctx, o, opNum, r)
	}
}

func ObjectManagement3ServerHandle(ctx context.Context, o ObjectManagement3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// INtmsObjectManagement2 base method.
		return intmsobjectmanagement2.ObjectManagement2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // GetNtmsObjectAttributeAR
		op := &xxx_GetNTMSObjectAttributeAROperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeARRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeAR(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetNtmsObjectAttributeWR
		op := &xxx_GetNTMSObjectAttributeWROperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNTMSObjectAttributeWRRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNTMSObjectAttributeWR(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented INtmsObjectManagement3
type UnimplementedObjectManagement3Server struct {
	intmsobjectmanagement2.UnimplementedObjectManagement2Server
}

func (UnimplementedObjectManagement3Server) GetNTMSObjectAttributeAR(context.Context, *GetNTMSObjectAttributeARRequest) (*GetNTMSObjectAttributeARResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedObjectManagement3Server) GetNTMSObjectAttributeWR(context.Context, *GetNTMSObjectAttributeWRRequest) (*GetNTMSObjectAttributeWRResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ObjectManagement3Server = (*UnimplementedObjectManagement3Server)(nil)
