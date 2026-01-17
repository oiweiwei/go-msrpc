package iupdatedownloadcontent2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdatedownloadcontent "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatedownloadcontent/v0"
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
	_ = dcom.GoPackage
	_ = iupdatedownloadcontent.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateDownloadContent2 interface identifier c97ad11b-f257-420b-9d9f-377f733f6f68
	UpdateDownloadContent2IID = &dcom.IID{Data1: 0xc97ad11b, Data2: 0xf257, Data3: 0x420b, Data4: []byte{0x9d, 0x9f, 0x37, 0x7f, 0x73, 0x3f, 0x6f, 0x68}}
	// Syntax UUID
	UpdateDownloadContent2SyntaxUUID = &uuid.UUID{TimeLow: 0xc97ad11b, TimeMid: 0xf257, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x9f, Node: [6]uint8{0x37, 0x7f, 0x73, 0x3f, 0x6f, 0x68}}
	// Syntax ID
	UpdateDownloadContent2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateDownloadContent2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateDownloadContent2 interface.
type UpdateDownloadContent2Client interface {

	// IUpdateDownloadContent retrieval method.
	UpdateDownloadContent() iupdatedownloadcontent.UpdateDownloadContentClient

	// The IUpdateDownloadContent2::IsDeltaCompressedContent (opnum 9) method retrieves
	// whether the content is delta-compressed.
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
	// This method SHOULD return the value of the IsDeltaCompressedContent ADM element.
	GetIsDeltaCompressedContent(context.Context, *GetIsDeltaCompressedContentRequest, ...dcerpc.CallOption) (*GetIsDeltaCompressedContentResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateDownloadContent2Client
}

type xxx_DefaultUpdateDownloadContent2Client struct {
	iupdatedownloadcontent.UpdateDownloadContentClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateDownloadContent2Client) UpdateDownloadContent() iupdatedownloadcontent.UpdateDownloadContentClient {
	return o.UpdateDownloadContentClient
}

func (o *xxx_DefaultUpdateDownloadContent2Client) GetIsDeltaCompressedContent(ctx context.Context, in *GetIsDeltaCompressedContentRequest, opts ...dcerpc.CallOption) (*GetIsDeltaCompressedContentResponse, error) {
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
	out := &GetIsDeltaCompressedContentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateDownloadContent2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateDownloadContent2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateDownloadContent2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateDownloadContent2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateDownloadContent2Client{
		UpdateDownloadContentClient: o.UpdateDownloadContentClient.IPID(ctx, ipid),
		cc:                          o.cc,
		ipid:                        ipid,
	}
}

func NewUpdateDownloadContent2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateDownloadContent2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateDownloadContent2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdatedownloadcontent.NewUpdateDownloadContentClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateDownloadContent2Client{
		UpdateDownloadContentClient: base,
		cc:                          cc,
		ipid:                        ipid,
	}, nil
}

// xxx_GetIsDeltaCompressedContentOperation structure represents the IsDeltaCompressedContent operation
type xxx_GetIsDeltaCompressedContentOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsDeltaCompressedContentOperation) OpNum() int { return 8 }

func (o *xxx_GetIsDeltaCompressedContentOperation) OpName() string {
	return "/IUpdateDownloadContent2/v0/IsDeltaCompressedContent"
}

func (o *xxx_GetIsDeltaCompressedContentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDeltaCompressedContentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsDeltaCompressedContentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsDeltaCompressedContentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDeltaCompressedContentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_GetIsDeltaCompressedContentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// GetIsDeltaCompressedContentRequest structure represents the IsDeltaCompressedContent operation request
type GetIsDeltaCompressedContentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsDeltaCompressedContentRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsDeltaCompressedContentOperation) *xxx_GetIsDeltaCompressedContentOperation {
	if op == nil {
		op = &xxx_GetIsDeltaCompressedContentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsDeltaCompressedContentRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsDeltaCompressedContentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsDeltaCompressedContentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsDeltaCompressedContentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDeltaCompressedContentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsDeltaCompressedContentResponse structure represents the IsDeltaCompressedContent operation response
type GetIsDeltaCompressedContentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be VARIANT_TRUE if the content is delta-compressed or VARIANT_FALSE
	// if not.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The IsDeltaCompressedContent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsDeltaCompressedContentResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsDeltaCompressedContentOperation) *xxx_GetIsDeltaCompressedContentOperation {
	if op == nil {
		op = &xxx_GetIsDeltaCompressedContentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetIsDeltaCompressedContentResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsDeltaCompressedContentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetIsDeltaCompressedContentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsDeltaCompressedContentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDeltaCompressedContentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
