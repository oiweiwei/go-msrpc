package itransaction

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// ITransaction interface identifier 0fb15084-af41-11ce-bd2b-204c4f4f5020
	ITransactionIID = &dcom.IID{Data1: 0x0fb15084, Data2: 0xaf41, Data3: 0x11ce, Data4: []byte{0xbd, 0x2b, 0x20, 0x4c, 0x4f, 0x4f, 0x50, 0x20}}
	// Syntax UUID
	ITransactionSyntaxUUID = &uuid.UUID{TimeLow: 0xfb15084, TimeMid: 0xaf41, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x2b, Node: [6]uint8{0x20, 0x4c, 0x4f, 0x4f, 0x50, 0x20}}
	// Syntax ID
	ITransactionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ITransactionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITransaction interface.
type ITransactionClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Commit operation.
	Commit(context.Context, *CommitRequest, ...dcerpc.CallOption) (*CommitResponse, error)

	// Abort operation.
	Abort(context.Context, *AbortRequest, ...dcerpc.CallOption) (*AbortResponse, error)

	// GetTransactionInfo operation.
	GetTransactionInfo(context.Context, *GetTransactionInfoRequest, ...dcerpc.CallOption) (*GetTransactionInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ITransactionClient
}

type xxx_DefaultITransactionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultITransactionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultITransactionClient) Commit(ctx context.Context, in *CommitRequest, opts ...dcerpc.CallOption) (*CommitResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultITransactionClient) Abort(ctx context.Context, in *AbortRequest, opts ...dcerpc.CallOption) (*AbortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultITransactionClient) GetTransactionInfo(ctx context.Context, in *GetTransactionInfoRequest, opts ...dcerpc.CallOption) (*GetTransactionInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTransactionInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultITransactionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultITransactionClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultITransactionClient) IPID(ctx context.Context, ipid *dcom.IPID) ITransactionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultITransactionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewITransactionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ITransactionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ITransactionSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultITransactionClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CommitOperation structure represents the Commit operation
type xxx_CommitOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Retaining int16          `idl:"name:fRetaining" json:"retaining"`
	TC        uint32         `idl:"name:grfTC" json:"tc"`
	RM        uint32         `idl:"name:grfRM" json:"rm"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitOperation) OpNum() int { return 3 }

func (o *xxx_CommitOperation) OpName() string { return "/ITransaction/v0/Commit" }

func (o *xxx_CommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fRetaining {in} (1:(int16))
	{
		if err := w.WriteData(o.Retaining); err != nil {
			return err
		}
	}
	// grfTC {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TC); err != nil {
			return err
		}
	}
	// grfRM {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RM); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRetaining {in} (1:(int16))
	{
		if err := w.ReadData(&o.Retaining); err != nil {
			return err
		}
	}
	// grfTC {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TC); err != nil {
			return err
		}
	}
	// grfRM {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RM); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CommitRequest structure represents the Commit operation request
type CommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Retaining int16          `idl:"name:fRetaining" json:"retaining"`
	TC        uint32         `idl:"name:grfTC" json:"tc"`
	RM        uint32         `idl:"name:grfRM" json:"rm"`
}

func (o *CommitRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Retaining = o.Retaining
	op.TC = o.TC
	op.RM = o.RM
	return op
}

func (o *CommitRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Retaining = op.Retaining
	o.TC = op.TC
	o.RM = op.RM
}
func (o *CommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitResponse structure represents the Commit operation response
type CommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Commit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitOperation) *xxx_CommitOperation {
	if op == nil {
		op = &xxx_CommitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CommitResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortOperation structure represents the Abort operation
type xxx_AbortOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PboidReason *mqac.Boid     `idl:"name:pboidReason;pointer:unique" json:"pboid_reason"`
	Retaining   int16          `idl:"name:fRetaining" json:"retaining"`
	Async       int16          `idl:"name:fAsync" json:"async"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortOperation) OpNum() int { return 4 }

func (o *xxx_AbortOperation) OpName() string { return "/ITransaction/v0/Abort" }

func (o *xxx_AbortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pboidReason {in} (1:{pointer=unique}*(1))(2:{alias=BOID}(struct))
	{
		if o.PboidReason != nil {
			_ptr_pboidReason := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PboidReason != nil {
					if err := o.PboidReason.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Boid{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PboidReason, _ptr_pboidReason); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fRetaining {in} (1:(int16))
	{
		if err := w.WriteData(o.Retaining); err != nil {
			return err
		}
	}
	// fAsync {in} (1:(int16))
	{
		if err := w.WriteData(o.Async); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pboidReason {in} (1:{pointer=unique}*(1))(2:{alias=BOID}(struct))
	{
		_ptr_pboidReason := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PboidReason == nil {
				o.PboidReason = &mqac.Boid{}
			}
			if err := o.PboidReason.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pboidReason := func(ptr interface{}) { o.PboidReason = *ptr.(**mqac.Boid) }
		if err := w.ReadPointer(&o.PboidReason, _s_pboidReason, _ptr_pboidReason); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fRetaining {in} (1:(int16))
	{
		if err := w.ReadData(&o.Retaining); err != nil {
			return err
		}
	}
	// fAsync {in} (1:(int16))
	{
		if err := w.ReadData(&o.Async); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AbortRequest structure represents the Abort operation request
type AbortRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	PboidReason *mqac.Boid     `idl:"name:pboidReason;pointer:unique" json:"pboid_reason"`
	Retaining   int16          `idl:"name:fRetaining" json:"retaining"`
	Async       int16          `idl:"name:fAsync" json:"async"`
}

func (o *AbortRequest) xxx_ToOp(ctx context.Context, op *xxx_AbortOperation) *xxx_AbortOperation {
	if op == nil {
		op = &xxx_AbortOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PboidReason = o.PboidReason
	op.Retaining = o.Retaining
	op.Async = o.Async
	return op
}

func (o *AbortRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PboidReason = op.PboidReason
	o.Retaining = op.Retaining
	o.Async = op.Async
}
func (o *AbortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AbortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortResponse structure represents the Abort operation response
type AbortResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Abort return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AbortResponse) xxx_ToOp(ctx context.Context, op *xxx_AbortOperation) *xxx_AbortOperation {
	if op == nil {
		op = &xxx_AbortOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AbortResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AbortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AbortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTransactionInfoOperation structure represents the GetTransactionInfo operation
type xxx_GetTransactionInfoOperation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Pinfo  *mqac.Xacttransinfo `idl:"name:pinfo" json:"pinfo"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTransactionInfoOperation) OpNum() int { return 5 }

func (o *xxx_GetTransactionInfoOperation) OpName() string {
	return "/ITransaction/v0/GetTransactionInfo"
}

func (o *xxx_GetTransactionInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pinfo {out} (1:{pointer=ref}*(1))(2:{alias=XACTTRANSINFO}(struct))
	{
		if o.Pinfo != nil {
			if err := o.Pinfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&mqac.Xacttransinfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pinfo {out} (1:{pointer=ref}*(1))(2:{alias=XACTTRANSINFO}(struct))
	{
		if o.Pinfo == nil {
			o.Pinfo = &mqac.Xacttransinfo{}
		}
		if err := o.Pinfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTransactionInfoRequest structure represents the GetTransactionInfo operation request
type GetTransactionInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTransactionInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionInfoOperation) *xxx_GetTransactionInfoOperation {
	if op == nil {
		op = &xxx_GetTransactionInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTransactionInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTransactionInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTransactionInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTransactionInfoResponse structure represents the GetTransactionInfo operation response
type GetTransactionInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Pinfo *mqac.Xacttransinfo `idl:"name:pinfo" json:"pinfo"`
	// Return: The GetTransactionInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTransactionInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionInfoOperation) *xxx_GetTransactionInfoOperation {
	if op == nil {
		op = &xxx_GetTransactionInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Pinfo = o.Pinfo
	op.Return = o.Return
	return op
}

func (o *GetTransactionInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pinfo = op.Pinfo
	o.Return = op.Return
}
func (o *GetTransactionInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTransactionInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
