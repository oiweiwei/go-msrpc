package conv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = conv.GoPackage
)

var (
	// import guard
	GoPackage = "conv"
)

var (
	// Syntax UUID
	ConvSyntaxUUID = &uuid.UUID{TimeLow: 0x333a2276, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xd, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x80, 0x9c, 0x0, 0x0, 0x0}}
	// Syntax ID
	ConvSyntaxV3_0 = &dcerpc.SyntaxID{IfUUID: ConvSyntaxUUID, IfVersionMajor: 3, IfVersionMinor: 0}
)

// conv interface.
type ConvClient interface {

	// conv_who_are_you operation.
	WhoAreYou(context.Context, *WhoAreYouRequest, ...dcerpc.CallOption) (*WhoAreYouResponse, error)

	// conv_who_are_you2 operation.
	WhoAreYou2(context.Context, *WhoAreYou2Request, ...dcerpc.CallOption) (*WhoAreYou2Response, error)

	// conv_are_you_there operation.
	AreYouThere(context.Context, *AreYouThereRequest, ...dcerpc.CallOption) (*AreYouThereResponse, error)

	// conv_who_are_you_auth operation.
	WhoAreYouAuth(context.Context, *WhoAreYouAuthRequest, ...dcerpc.CallOption) (*WhoAreYouAuthResponse, error)

	// conv_who_are_you_auth_more operation.
	WhoAreYouAuthMore(context.Context, *WhoAreYouAuthMoreRequest, ...dcerpc.CallOption) (*WhoAreYouAuthMoreResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultConvClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultConvClient) WhoAreYou(ctx context.Context, in *WhoAreYouRequest, opts ...dcerpc.CallOption) (*WhoAreYouResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WhoAreYouResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvClient) WhoAreYou2(ctx context.Context, in *WhoAreYou2Request, opts ...dcerpc.CallOption) (*WhoAreYou2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WhoAreYou2Response{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvClient) AreYouThere(ctx context.Context, in *AreYouThereRequest, opts ...dcerpc.CallOption) (*AreYouThereResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AreYouThereResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvClient) WhoAreYouAuth(ctx context.Context, in *WhoAreYouAuthRequest, opts ...dcerpc.CallOption) (*WhoAreYouAuthResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WhoAreYouAuthResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvClient) WhoAreYouAuthMore(ctx context.Context, in *WhoAreYouAuthMoreRequest, opts ...dcerpc.CallOption) (*WhoAreYouAuthMoreResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WhoAreYouAuthMoreResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultConvClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultConvClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewConvClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ConvClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ConvSyntaxV3_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultConvClient{cc: cc}, nil
}

// xxx_WhoAreYouOperation structure represents the conv_who_are_you operation
type xxx_WhoAreYouOperation struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
	Seq         uint32     `idl:"name:seq" json:"seq"`
	Status      uint32     `idl:"name:st" json:"status"`
}

func (o *xxx_WhoAreYouOperation) OpNum() int { return 0 }

func (o *xxx_WhoAreYouOperation) OpName() string { return "/conv/v3/conv_who_are_you" }

func (o *xxx_WhoAreYouOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID != nil {
			if err := o.ActivityUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID == nil {
			o.ActivityUID = &conv.UUID{}
		}
		if err := o.ActivityUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Seq); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Seq); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// WhoAreYouRequest structure represents the conv_who_are_you operation request
type WhoAreYouRequest struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
}

func (o *WhoAreYouRequest) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouOperation) *xxx_WhoAreYouOperation {
	if op == nil {
		op = &xxx_WhoAreYouOperation{}
	}
	if o == nil {
		return op
	}
	op.ActivityUID = o.ActivityUID
	op.BootTime = o.BootTime
	return op
}

func (o *WhoAreYouRequest) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouOperation) {
	if o == nil {
		return
	}
	o.ActivityUID = op.ActivityUID
	o.BootTime = op.BootTime
}
func (o *WhoAreYouRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WhoAreYouRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WhoAreYouResponse structure represents the conv_who_are_you operation response
type WhoAreYouResponse struct {
	Seq    uint32 `idl:"name:seq" json:"seq"`
	Status uint32 `idl:"name:st" json:"status"`
}

func (o *WhoAreYouResponse) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouOperation) *xxx_WhoAreYouOperation {
	if op == nil {
		op = &xxx_WhoAreYouOperation{}
	}
	if o == nil {
		return op
	}
	op.Seq = o.Seq
	op.Status = o.Status
	return op
}

func (o *WhoAreYouResponse) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouOperation) {
	if o == nil {
		return
	}
	o.Seq = op.Seq
	o.Status = op.Status
}
func (o *WhoAreYouResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WhoAreYouResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WhoAreYou2Operation structure represents the conv_who_are_you2 operation
type xxx_WhoAreYou2Operation struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
	Seq         uint32     `idl:"name:seq" json:"seq"`
	CasUUID     *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
	Status      uint32     `idl:"name:st" json:"status"`
}

func (o *xxx_WhoAreYou2Operation) OpNum() int { return 1 }

func (o *xxx_WhoAreYou2Operation) OpName() string { return "/conv/v3/conv_who_are_you2" }

func (o *xxx_WhoAreYou2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYou2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID != nil {
			if err := o.ActivityUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYou2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID == nil {
			o.ActivityUID = &conv.UUID{}
		}
		if err := o.ActivityUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYou2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYou2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Seq); err != nil {
			return err
		}
	}
	// cas_uuid {out} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
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
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYou2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Seq); err != nil {
			return err
		}
	}
	// cas_uuid {out} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.CasUUID == nil {
			o.CasUUID = &conv.UUID{}
		}
		if err := o.CasUUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// WhoAreYou2Request structure represents the conv_who_are_you2 operation request
type WhoAreYou2Request struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
}

func (o *WhoAreYou2Request) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYou2Operation) *xxx_WhoAreYou2Operation {
	if op == nil {
		op = &xxx_WhoAreYou2Operation{}
	}
	if o == nil {
		return op
	}
	op.ActivityUID = o.ActivityUID
	op.BootTime = o.BootTime
	return op
}

func (o *WhoAreYou2Request) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYou2Operation) {
	if o == nil {
		return
	}
	o.ActivityUID = op.ActivityUID
	o.BootTime = op.BootTime
}
func (o *WhoAreYou2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WhoAreYou2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYou2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WhoAreYou2Response structure represents the conv_who_are_you2 operation response
type WhoAreYou2Response struct {
	Seq     uint32     `idl:"name:seq" json:"seq"`
	CasUUID *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
	Status  uint32     `idl:"name:st" json:"status"`
}

func (o *WhoAreYou2Response) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYou2Operation) *xxx_WhoAreYou2Operation {
	if op == nil {
		op = &xxx_WhoAreYou2Operation{}
	}
	if o == nil {
		return op
	}
	op.Seq = o.Seq
	op.CasUUID = o.CasUUID
	op.Status = o.Status
	return op
}

func (o *WhoAreYou2Response) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYou2Operation) {
	if o == nil {
		return
	}
	o.Seq = op.Seq
	o.CasUUID = op.CasUUID
	o.Status = op.Status
}
func (o *WhoAreYou2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WhoAreYou2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYou2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AreYouThereOperation structure represents the conv_are_you_there operation
type xxx_AreYouThereOperation struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
	Status      uint32     `idl:"name:st" json:"status"`
}

func (o *xxx_AreYouThereOperation) OpNum() int { return 2 }

func (o *xxx_AreYouThereOperation) OpName() string { return "/conv/v3/conv_are_you_there" }

func (o *xxx_AreYouThereOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AreYouThereOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID != nil {
			if err := o.ActivityUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AreYouThereOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID == nil {
			o.ActivityUID = &conv.UUID{}
		}
		if err := o.ActivityUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.BootTime); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AreYouThereOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AreYouThereOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AreYouThereOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// AreYouThereRequest structure represents the conv_are_you_there operation request
type AreYouThereRequest struct {
	ActivityUID *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime    uint32     `idl:"name:boot_time" json:"boot_time"`
}

func (o *AreYouThereRequest) xxx_ToOp(ctx context.Context, op *xxx_AreYouThereOperation) *xxx_AreYouThereOperation {
	if op == nil {
		op = &xxx_AreYouThereOperation{}
	}
	if o == nil {
		return op
	}
	op.ActivityUID = o.ActivityUID
	op.BootTime = o.BootTime
	return op
}

func (o *AreYouThereRequest) xxx_FromOp(ctx context.Context, op *xxx_AreYouThereOperation) {
	if o == nil {
		return
	}
	o.ActivityUID = op.ActivityUID
	o.BootTime = op.BootTime
}
func (o *AreYouThereRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AreYouThereRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AreYouThereOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AreYouThereResponse structure represents the conv_are_you_there operation response
type AreYouThereResponse struct {
	Status uint32 `idl:"name:st" json:"status"`
}

func (o *AreYouThereResponse) xxx_ToOp(ctx context.Context, op *xxx_AreYouThereOperation) *xxx_AreYouThereOperation {
	if op == nil {
		op = &xxx_AreYouThereOperation{}
	}
	if o == nil {
		return op
	}
	op.Status = o.Status
	return op
}

func (o *AreYouThereResponse) xxx_FromOp(ctx context.Context, op *xxx_AreYouThereOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
}
func (o *AreYouThereResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AreYouThereResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AreYouThereOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WhoAreYouAuthOperation structure represents the conv_who_are_you_auth operation
type xxx_WhoAreYouAuthOperation struct {
	ActivityUID  *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime     uint32     `idl:"name:boot_time" json:"boot_time"`
	InData       []byte     `idl:"name:in_data;size_is:(in_len)" json:"in_data"`
	InLength     int32      `idl:"name:in_len" json:"in_length"`
	OutMaxLength int32      `idl:"name:out_max_len" json:"out_max_length"`
	Seq          uint32     `idl:"name:seq" json:"seq"`
	CasUUID      *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
	OutData      []byte     `idl:"name:out_data;size_is:(out_max_len);length_is:(out_len)" json:"out_data"`
	OutLength    int32      `idl:"name:out_len" json:"out_length"`
	Status       uint32     `idl:"name:st" json:"status"`
}

func (o *xxx_WhoAreYouAuthOperation) OpNum() int { return 3 }

func (o *xxx_WhoAreYouAuthOperation) OpName() string { return "/conv/v3/conv_who_are_you_auth" }

func (o *xxx_WhoAreYouAuthOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InData != nil && o.InLength == 0 {
		o.InLength = int32(len(o.InData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID != nil {
			if err := o.ActivityUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.BootTime); err != nil {
			return err
		}
	}
	// in_data {in} (1:[dim:0,size_is=in_len](uint8))
	{
		dimSize1 := uint64(o.InLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.InData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.InData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.InData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// in_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.InLength); err != nil {
			return err
		}
	}
	// out_max_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.OutMaxLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID == nil {
			o.ActivityUID = &conv.UUID{}
		}
		if err := o.ActivityUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.BootTime); err != nil {
			return err
		}
	}
	// in_data {in} (1:[dim:0,size_is=in_len](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InData", sizeInfo[0])
		}
		o.InData = make([]byte, sizeInfo[0])
		for i1 := range o.InData {
			i1 := i1
			if err := w.ReadData(&o.InData[i1]); err != nil {
				return err
			}
		}
	}
	// in_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.InLength); err != nil {
			return err
		}
	}
	// out_max_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.OutMaxLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutData != nil && o.OutLength == 0 {
		o.OutLength = int32(len(o.OutData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Seq); err != nil {
			return err
		}
	}
	// cas_uuid {out} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
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
	// out_data {out} (1:[dim:0,size_is=out_max_len,length_is=out_len](uint8))
	{
		dimSize1 := uint64(o.OutMaxLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// out_len {out} (1:{pointer=ref}*(1))(2:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// seq {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Seq); err != nil {
			return err
		}
	}
	// cas_uuid {out} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.CasUUID == nil {
			o.CasUUID = &conv.UUID{}
		}
		if err := o.CasUUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// out_data {out} (1:[dim:0,size_is=out_max_len,length_is=out_len](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutData", sizeInfo[0])
		}
		o.OutData = make([]byte, sizeInfo[0])
		for i1 := range o.OutData {
			i1 := i1
			if err := w.ReadData(&o.OutData[i1]); err != nil {
				return err
			}
		}
	}
	// out_len {out} (1:{pointer=ref}*(1))(2:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// WhoAreYouAuthRequest structure represents the conv_who_are_you_auth operation request
type WhoAreYouAuthRequest struct {
	ActivityUID  *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime     uint32     `idl:"name:boot_time" json:"boot_time"`
	InData       []byte     `idl:"name:in_data;size_is:(in_len)" json:"in_data"`
	InLength     int32      `idl:"name:in_len" json:"in_length"`
	OutMaxLength int32      `idl:"name:out_max_len" json:"out_max_length"`
}

func (o *WhoAreYouAuthRequest) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouAuthOperation) *xxx_WhoAreYouAuthOperation {
	if op == nil {
		op = &xxx_WhoAreYouAuthOperation{}
	}
	if o == nil {
		return op
	}
	op.ActivityUID = o.ActivityUID
	op.BootTime = o.BootTime
	op.InData = o.InData
	op.InLength = o.InLength
	op.OutMaxLength = o.OutMaxLength
	return op
}

func (o *WhoAreYouAuthRequest) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouAuthOperation) {
	if o == nil {
		return
	}
	o.ActivityUID = op.ActivityUID
	o.BootTime = op.BootTime
	o.InData = op.InData
	o.InLength = op.InLength
	o.OutMaxLength = op.OutMaxLength
}
func (o *WhoAreYouAuthRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WhoAreYouAuthRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouAuthOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WhoAreYouAuthResponse structure represents the conv_who_are_you_auth operation response
type WhoAreYouAuthResponse struct {
	// XXX: out_max_len is an implicit input depedency for output parameters
	OutMaxLength int32 `idl:"name:out_max_len" json:"out_max_length"`

	Seq       uint32     `idl:"name:seq" json:"seq"`
	CasUUID   *conv.UUID `idl:"name:cas_uuid" json:"cas_uuid"`
	OutData   []byte     `idl:"name:out_data;size_is:(out_max_len);length_is:(out_len)" json:"out_data"`
	OutLength int32      `idl:"name:out_len" json:"out_length"`
	Status    uint32     `idl:"name:st" json:"status"`
}

func (o *WhoAreYouAuthResponse) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouAuthOperation) *xxx_WhoAreYouAuthOperation {
	if op == nil {
		op = &xxx_WhoAreYouAuthOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.OutMaxLength == int32(0) {
		op.OutMaxLength = o.OutMaxLength
	}

	op.Seq = o.Seq
	op.CasUUID = o.CasUUID
	op.OutData = o.OutData
	op.OutLength = o.OutLength
	op.Status = o.Status
	return op
}

func (o *WhoAreYouAuthResponse) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouAuthOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.OutMaxLength = op.OutMaxLength

	o.Seq = op.Seq
	o.CasUUID = op.CasUUID
	o.OutData = op.OutData
	o.OutLength = op.OutLength
	o.Status = op.Status
}
func (o *WhoAreYouAuthResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WhoAreYouAuthResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouAuthOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WhoAreYouAuthMoreOperation structure represents the conv_who_are_you_auth_more operation
type xxx_WhoAreYouAuthMoreOperation struct {
	ActivityUID  *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime     uint32     `idl:"name:boot_time" json:"boot_time"`
	Index        int32      `idl:"name:index" json:"index"`
	OutMaxLength int32      `idl:"name:out_max_len" json:"out_max_length"`
	OutData      []byte     `idl:"name:out_data;size_is:(out_max_len);length_is:(out_len)" json:"out_data"`
	OutLength    int32      `idl:"name:out_len" json:"out_length"`
	Status       uint32     `idl:"name:st" json:"status"`
}

func (o *xxx_WhoAreYouAuthMoreOperation) OpNum() int { return 4 }

func (o *xxx_WhoAreYouAuthMoreOperation) OpName() string {
	return "/conv/v3/conv_who_are_you_auth_more"
}

func (o *xxx_WhoAreYouAuthMoreOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthMoreOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID != nil {
			if err := o.ActivityUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&conv.UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.BootTime); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// out_max_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.OutMaxLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthMoreOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// actuid {in} (1:{pointer=ref}*(1))(2:{alias=uuid_t, names=GUID}(struct))
	{
		if o.ActivityUID == nil {
			o.ActivityUID = &conv.UUID{}
		}
		if err := o.ActivityUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// boot_time {in} (1:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.BootTime); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// out_max_len {in} (1:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.OutMaxLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthMoreOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutData != nil && o.OutLength == 0 {
		o.OutLength = int32(len(o.OutData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthMoreOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// out_data {out} (1:[dim:0,size_is=out_max_len,length_is=out_len](uint8))
	{
		dimSize1 := uint64(o.OutMaxLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutLength)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// out_len {out} (1:{pointer=ref}*(1))(2:{alias=signed32}(int32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WhoAreYouAuthMoreOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// out_data {out} (1:[dim:0,size_is=out_max_len,length_is=out_len](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutData", sizeInfo[0])
		}
		o.OutData = make([]byte, sizeInfo[0])
		for i1 := range o.OutData {
			i1 := i1
			if err := w.ReadData(&o.OutData[i1]); err != nil {
				return err
			}
		}
	}
	// out_len {out} (1:{pointer=ref}*(1))(2:{alias=signed32}(int32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
			return err
		}
	}
	// st {out} (1:{pointer=ref}*(1))(2:{alias=unsigned32}(uint32))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	return nil
}

// WhoAreYouAuthMoreRequest structure represents the conv_who_are_you_auth_more operation request
type WhoAreYouAuthMoreRequest struct {
	ActivityUID  *conv.UUID `idl:"name:actuid" json:"activity_uid"`
	BootTime     uint32     `idl:"name:boot_time" json:"boot_time"`
	Index        int32      `idl:"name:index" json:"index"`
	OutMaxLength int32      `idl:"name:out_max_len" json:"out_max_length"`
}

func (o *WhoAreYouAuthMoreRequest) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouAuthMoreOperation) *xxx_WhoAreYouAuthMoreOperation {
	if op == nil {
		op = &xxx_WhoAreYouAuthMoreOperation{}
	}
	if o == nil {
		return op
	}
	op.ActivityUID = o.ActivityUID
	op.BootTime = o.BootTime
	op.Index = o.Index
	op.OutMaxLength = o.OutMaxLength
	return op
}

func (o *WhoAreYouAuthMoreRequest) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouAuthMoreOperation) {
	if o == nil {
		return
	}
	o.ActivityUID = op.ActivityUID
	o.BootTime = op.BootTime
	o.Index = op.Index
	o.OutMaxLength = op.OutMaxLength
}
func (o *WhoAreYouAuthMoreRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WhoAreYouAuthMoreRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouAuthMoreOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WhoAreYouAuthMoreResponse structure represents the conv_who_are_you_auth_more operation response
type WhoAreYouAuthMoreResponse struct {
	// XXX: out_max_len is an implicit input depedency for output parameters
	OutMaxLength int32 `idl:"name:out_max_len" json:"out_max_length"`

	OutData   []byte `idl:"name:out_data;size_is:(out_max_len);length_is:(out_len)" json:"out_data"`
	OutLength int32  `idl:"name:out_len" json:"out_length"`
	Status    uint32 `idl:"name:st" json:"status"`
}

func (o *WhoAreYouAuthMoreResponse) xxx_ToOp(ctx context.Context, op *xxx_WhoAreYouAuthMoreOperation) *xxx_WhoAreYouAuthMoreOperation {
	if op == nil {
		op = &xxx_WhoAreYouAuthMoreOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.OutMaxLength == int32(0) {
		op.OutMaxLength = o.OutMaxLength
	}

	op.OutData = o.OutData
	op.OutLength = o.OutLength
	op.Status = o.Status
	return op
}

func (o *WhoAreYouAuthMoreResponse) xxx_FromOp(ctx context.Context, op *xxx_WhoAreYouAuthMoreOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.OutMaxLength = op.OutMaxLength

	o.OutData = op.OutData
	o.OutLength = op.OutLength
	o.Status = op.Status
}
func (o *WhoAreYouAuthMoreResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WhoAreYouAuthMoreResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WhoAreYouAuthMoreOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
