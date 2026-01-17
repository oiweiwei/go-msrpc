package convc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	conv "github.com/oiweiwei/go-msrpc/msrpc/conv"
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
	_ = conv.GoPackage
)

var (
	// import guard
	GoPackage = "conv"
)

var (
	// Syntax UUID
	ConvcSyntaxUUID = &uuid.UUID{TimeLow: 0x4a967f14, TimeMid: 0x3000, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xd, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x28, 0x82, 0x0, 0x0, 0x0}}
	// Syntax ID
	ConvcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ConvcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// convc interface.
type ConvcClient interface {

	// convc_indy operation.
	Indy(context.Context, *IndyRequest, ...dcerpc.CallOption) (*IndyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultConvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultConvcClient) Indy(ctx context.Context, in *IndyRequest, opts ...dcerpc.CallOption) (*IndyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IndyResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewConvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConvcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultConvcClient{cc: cc}, nil
}

// xxx_IndyOperation structure represents the convc_indy operation
type xxx_IndyOperation struct {
	CasUUID *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
}

func (o *xxx_IndyOperation) OpNum() int { return 0 }

func (o *xxx_IndyOperation) OpName() string { return "/convc/v1/convc_indy" }

func (o *xxx_IndyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// cas_uuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.CasUUID != nil {
			if err := o.CasUUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_IndyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// cas_uuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.CasUUID == nil {
			o.CasUUID = &conv.UUID{}
		}
		if err := o.CasUUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IndyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_IndyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	return nil
}

// IndyRequest structure represents the convc_indy operation request
type IndyRequest struct {
	CasUUID *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
}

func (o *IndyRequest) xxx_ToOp(ctx context.Context, op *xxx_IndyOperation) *xxx_IndyOperation {
	if op == nil {
		op = &xxx_IndyOperation{}
	}
	if o == nil {
		return op
	}
	op.CasUUID = o.CasUUID
	return op
}

func (o *IndyRequest) xxx_FromOp(ctx context.Context, op *xxx_IndyOperation) {
	if o == nil {
		return
	}
	o.CasUUID = op.CasUUID
}
func (o *IndyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IndyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IndyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IndyResponse structure represents the convc_indy operation response
type IndyResponse struct {
}

func (o *IndyResponse) xxx_ToOp(ctx context.Context, op *xxx_IndyOperation) *xxx_IndyOperation {
	if op == nil {
		op = &xxx_IndyOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *IndyResponse) xxx_FromOp(ctx context.Context, op *xxx_IndyOperation) {
	if o == nil {
		return
	}
}
func (o *IndyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IndyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IndyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
