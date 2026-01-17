package lsacap

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	capr "github.com/oiweiwei/go-msrpc/msrpc/capr"
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
	_ = capr.GoPackage
)

var (
	// import guard
	GoPackage = "capr"
)

var (
	// Syntax UUID
	LsacapSyntaxUUID = &uuid.UUID{TimeLow: 0xafc07e2e, TimeMid: 0x311c, TimeHiAndVersion: 0x4435, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x8c, Node: [6]uint8{0xc4, 0x83, 0xff, 0xee, 0xc7, 0xc9}}
	// Syntax ID
	LsacapSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: LsacapSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// lsacap interface.
type LsacapClient interface {

	// This method returns a list of the CAPIDs of all the central access policies available
	// on the specified remote machine. These identifiers are equivalent to the SIDs of
	// the central access policy objects as they are stored in Active Directory.
	//
	// Return Values:
	//
	// If the method succeeds, the function MUST return 0x00000000 (ERROR_SUCCESS).
	GetAvailableCapIDs(context.Context, *GetAvailableCapIDsRequest, ...dcerpc.CallOption) (*GetAvailableCapIDsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultLsacapClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultLsacapClient) GetAvailableCapIDs(ctx context.Context, in *GetAvailableCapIDsRequest, opts ...dcerpc.CallOption) (*GetAvailableCapIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAvailableCapIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsacapClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLsacapClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewLsacapClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LsacapClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LsacapSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultLsacapClient{cc: cc}, nil
}

// xxx_GetAvailableCapIDsOperation structure represents the LsarGetAvailableCAPIDs operation
type xxx_GetAvailableCapIDsOperation struct {
	WrappedCapIDs *capr.WrappedCapidSet `idl:"name:WrappedCAPIDs" json:"wrapped_cap_ids"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAvailableCapIDsOperation) OpNum() int { return 0 }

func (o *xxx_GetAvailableCapIDsOperation) OpName() string { return "/lsacap/v1/LsarGetAvailableCAPIDs" }

func (o *xxx_GetAvailableCapIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAvailableCapIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetAvailableCapIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetAvailableCapIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAvailableCapIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// WrappedCAPIDs {out} (1:{pointer=ref}*(1))(2:{alias=LSAPR_WRAPPED_CAPID_SET}(struct))
	{
		if o.WrappedCapIDs != nil {
			if err := o.WrappedCapIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&capr.WrappedCapidSet{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAvailableCapIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// WrappedCAPIDs {out} (1:{pointer=ref}*(1))(2:{alias=LSAPR_WRAPPED_CAPID_SET}(struct))
	{
		if o.WrappedCapIDs == nil {
			o.WrappedCapIDs = &capr.WrappedCapidSet{}
		}
		if err := o.WrappedCapIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAvailableCapIDsRequest structure represents the LsarGetAvailableCAPIDs operation request
type GetAvailableCapIDsRequest struct {
}

func (o *GetAvailableCapIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAvailableCapIDsOperation) *xxx_GetAvailableCapIDsOperation {
	if op == nil {
		op = &xxx_GetAvailableCapIDsOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetAvailableCapIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAvailableCapIDsOperation) {
	if o == nil {
		return
	}
}
func (o *GetAvailableCapIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAvailableCapIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAvailableCapIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAvailableCapIDsResponse structure represents the LsarGetAvailableCAPIDs operation response
type GetAvailableCapIDsResponse struct {
	// WrappedCAPIDs: A pointer to LSAPR_WRAPPED_CAPID_SET, as defined in section 2.2.1.1.
	WrappedCapIDs *capr.WrappedCapidSet `idl:"name:WrappedCAPIDs" json:"wrapped_cap_ids"`
	// Return: The LsarGetAvailableCAPIDs return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAvailableCapIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAvailableCapIDsOperation) *xxx_GetAvailableCapIDsOperation {
	if op == nil {
		op = &xxx_GetAvailableCapIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.WrappedCapIDs = o.WrappedCapIDs
	op.Return = o.Return
	return op
}

func (o *GetAvailableCapIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAvailableCapIDsOperation) {
	if o == nil {
		return
	}
	o.WrappedCapIDs = op.WrappedCapIDs
	o.Return = op.Return
}
func (o *GetAvailableCapIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAvailableCapIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAvailableCapIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
