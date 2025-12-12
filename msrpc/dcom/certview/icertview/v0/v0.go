package icertview

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	certview "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = certview.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/certview"
)

var (
	// ICertView interface identifier c3fac344-1e84-11d1-9bd6-00c04fb683fa
	CertViewIID = &dcom.IID{Data1: 0xc3fac344, Data2: 0x1e84, Data3: 0x11d1, Data4: []byte{0x9b, 0xd6, 0x00, 0xc0, 0x4f, 0xb6, 0x83, 0xfa}}
	// Syntax UUID
	CertViewSyntaxUUID = &uuid.UUID{TimeLow: 0xc3fac344, TimeMid: 0x1e84, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0xd6, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x83, 0xfa}}
	// Syntax ID
	CertViewSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertViewSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertView interface.
type CertViewClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// OpenConnection operation.
	OpenConnection(context.Context, *OpenConnectionRequest, ...dcerpc.CallOption) (*OpenConnectionResponse, error)

	// EnumCertViewColumn operation.
	EnumCertViewColumn(context.Context, *EnumCertViewColumnRequest, ...dcerpc.CallOption) (*EnumCertViewColumnResponse, error)

	// GetColumnCount operation.
	GetColumnCount(context.Context, *GetColumnCountRequest, ...dcerpc.CallOption) (*GetColumnCountResponse, error)

	// GetColumnIndex operation.
	GetColumnIndex(context.Context, *GetColumnIndexRequest, ...dcerpc.CallOption) (*GetColumnIndexResponse, error)

	// SetResultColumnCount operation.
	SetResultColumnCount(context.Context, *SetResultColumnCountRequest, ...dcerpc.CallOption) (*SetResultColumnCountResponse, error)

	// SetResultColumn operation.
	SetResultColumn(context.Context, *SetResultColumnRequest, ...dcerpc.CallOption) (*SetResultColumnResponse, error)

	// SetRestriction operation.
	SetRestriction(context.Context, *SetRestrictionRequest, ...dcerpc.CallOption) (*SetRestrictionResponse, error)

	// OpenView operation.
	OpenView(context.Context, *OpenViewRequest, ...dcerpc.CallOption) (*OpenViewResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertViewClient
}

type xxx_DefaultCertViewClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertViewClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultCertViewClient) OpenConnection(ctx context.Context, in *OpenConnectionRequest, opts ...dcerpc.CallOption) (*OpenConnectionResponse, error) {
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
	out := &OpenConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) EnumCertViewColumn(ctx context.Context, in *EnumCertViewColumnRequest, opts ...dcerpc.CallOption) (*EnumCertViewColumnResponse, error) {
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
	out := &EnumCertViewColumnResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) GetColumnCount(ctx context.Context, in *GetColumnCountRequest, opts ...dcerpc.CallOption) (*GetColumnCountResponse, error) {
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
	out := &GetColumnCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) GetColumnIndex(ctx context.Context, in *GetColumnIndexRequest, opts ...dcerpc.CallOption) (*GetColumnIndexResponse, error) {
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
	out := &GetColumnIndexResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) SetResultColumnCount(ctx context.Context, in *SetResultColumnCountRequest, opts ...dcerpc.CallOption) (*SetResultColumnCountResponse, error) {
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
	out := &SetResultColumnCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) SetResultColumn(ctx context.Context, in *SetResultColumnRequest, opts ...dcerpc.CallOption) (*SetResultColumnResponse, error) {
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
	out := &SetResultColumnResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) SetRestriction(ctx context.Context, in *SetRestrictionRequest, opts ...dcerpc.CallOption) (*SetRestrictionResponse, error) {
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
	out := &SetRestrictionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) OpenView(ctx context.Context, in *OpenViewRequest, opts ...dcerpc.CallOption) (*OpenViewResponse, error) {
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
	out := &OpenViewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertViewClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertViewClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCertViewClient) IPID(ctx context.Context, ipid *dcom.IPID) CertViewClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertViewClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewCertViewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertViewClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertViewSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCertViewClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_OpenConnectionOperation structure represents the OpenConnection operation
type xxx_OpenConnectionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Config *oaut.String   `idl:"name:strConfig" json:"config"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenConnectionOperation) OpNum() int { return 7 }

func (o *xxx_OpenConnectionOperation) OpName() string { return "/ICertView/v0/OpenConnection" }

func (o *xxx_OpenConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// strConfig {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Config != nil {
			_ptr_strConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Config != nil {
					if err := o.Config.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Config, _ptr_strConfig); err != nil {
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
	return nil
}

func (o *xxx_OpenConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// strConfig {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Config == nil {
				o.Config = &oaut.String{}
			}
			if err := o.Config.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strConfig := func(ptr interface{}) { o.Config = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Config, _s_strConfig, _ptr_strConfig); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OpenConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OpenConnectionRequest structure represents the OpenConnection operation request
type OpenConnectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	Config *oaut.String   `idl:"name:strConfig" json:"config"`
}

func (o *OpenConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenConnectionOperation) *xxx_OpenConnectionOperation {
	if op == nil {
		op = &xxx_OpenConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Config = o.Config
	return op
}

func (o *OpenConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenConnectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Config = op.Config
}
func (o *OpenConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenConnectionResponse structure represents the OpenConnection operation response
type OpenConnectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OpenConnection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenConnectionOperation) *xxx_OpenConnectionOperation {
	if op == nil {
		op = &xxx_OpenConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OpenConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenConnectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OpenConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumCertViewColumnOperation structure represents the EnumCertViewColumn operation
type xxx_EnumCertViewColumnOperation struct {
	This         *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat               `idl:"name:That" json:"that"`
	ResultColumn int32                        `idl:"name:fResultColumn" json:"result_column"`
	Enum         *certview.EnumCertViewColumn `idl:"name:ppenum" json:"enum"`
	Return       int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumCertViewColumnOperation) OpNum() int { return 8 }

func (o *xxx_EnumCertViewColumnOperation) OpName() string { return "/ICertView/v0/EnumCertViewColumn" }

func (o *xxx_EnumCertViewColumnOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumCertViewColumnOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ResultColumn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumCertViewColumnOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ResultColumn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumCertViewColumnOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumCertViewColumnOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppenum {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumCERTVIEWCOLUMN}(interface))
	{
		if o.Enum != nil {
			_ptr_ppenum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&certview.EnumCertViewColumn{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppenum); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumCertViewColumnOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppenum {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumCERTVIEWCOLUMN}(interface))
	{
		_ptr_ppenum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &certview.EnumCertViewColumn{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenum := func(ptr interface{}) { o.Enum = *ptr.(**certview.EnumCertViewColumn) }
		if err := w.ReadPointer(&o.Enum, _s_ppenum, _ptr_ppenum); err != nil {
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

// EnumCertViewColumnRequest structure represents the EnumCertViewColumn operation request
type EnumCertViewColumnRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResultColumn int32          `idl:"name:fResultColumn" json:"result_column"`
}

func (o *EnumCertViewColumnRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumCertViewColumnOperation) *xxx_EnumCertViewColumnOperation {
	if op == nil {
		op = &xxx_EnumCertViewColumnOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResultColumn = o.ResultColumn
	return op
}

func (o *EnumCertViewColumnRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumCertViewColumnOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResultColumn = op.ResultColumn
}
func (o *EnumCertViewColumnRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumCertViewColumnRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumCertViewColumnOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumCertViewColumnResponse structure represents the EnumCertViewColumn operation response
type EnumCertViewColumnResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Enum *certview.EnumCertViewColumn `idl:"name:ppenum" json:"enum"`
	// Return: The EnumCertViewColumn return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumCertViewColumnResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumCertViewColumnOperation) *xxx_EnumCertViewColumnOperation {
	if op == nil {
		op = &xxx_EnumCertViewColumnOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *EnumCertViewColumnResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumCertViewColumnOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *EnumCertViewColumnResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumCertViewColumnResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumCertViewColumnOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetColumnCountOperation structure represents the GetColumnCount operation
type xxx_GetColumnCountOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResultColumn int32          `idl:"name:fResultColumn" json:"result_column"`
	ColumnCount  int32          `idl:"name:pcColumn" json:"column_count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetColumnCountOperation) OpNum() int { return 9 }

func (o *xxx_GetColumnCountOperation) OpName() string { return "/ICertView/v0/GetColumnCount" }

func (o *xxx_GetColumnCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ResultColumn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ResultColumn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcColumn {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ColumnCount); err != nil {
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

func (o *xxx_GetColumnCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcColumn {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ColumnCount); err != nil {
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

// GetColumnCountRequest structure represents the GetColumnCount operation request
type GetColumnCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResultColumn int32          `idl:"name:fResultColumn" json:"result_column"`
}

func (o *GetColumnCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetColumnCountOperation) *xxx_GetColumnCountOperation {
	if op == nil {
		op = &xxx_GetColumnCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResultColumn = o.ResultColumn
	return op
}

func (o *GetColumnCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetColumnCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResultColumn = op.ResultColumn
}
func (o *GetColumnCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetColumnCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetColumnCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetColumnCountResponse structure represents the GetColumnCount operation response
type GetColumnCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnCount int32          `idl:"name:pcColumn" json:"column_count"`
	// Return: The GetColumnCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetColumnCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetColumnCountOperation) *xxx_GetColumnCountOperation {
	if op == nil {
		op = &xxx_GetColumnCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ColumnCount = o.ColumnCount
	op.Return = o.Return
	return op
}

func (o *GetColumnCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetColumnCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ColumnCount = op.ColumnCount
	o.Return = op.Return
}
func (o *GetColumnCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetColumnCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetColumnCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetColumnIndexOperation structure represents the GetColumnIndex operation
type xxx_GetColumnIndexOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResultColumn int32          `idl:"name:fResultColumn" json:"result_column"`
	ColumnName   *oaut.String   `idl:"name:strColumnName" json:"column_name"`
	ColumnIndex  int32          `idl:"name:pColumnIndex" json:"column_index"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetColumnIndexOperation) OpNum() int { return 10 }

func (o *xxx_GetColumnIndexOperation) OpName() string { return "/ICertView/v0/GetColumnIndex" }

func (o *xxx_GetColumnIndexOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnIndexOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ResultColumn); err != nil {
			return err
		}
	}
	// strColumnName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ColumnName != nil {
			_ptr_strColumnName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ColumnName != nil {
					if err := o.ColumnName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ColumnName, _ptr_strColumnName); err != nil {
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
	return nil
}

func (o *xxx_GetColumnIndexOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ResultColumn); err != nil {
			return err
		}
	}
	// strColumnName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_strColumnName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ColumnName == nil {
				o.ColumnName = &oaut.String{}
			}
			if err := o.ColumnName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_strColumnName := func(ptr interface{}) { o.ColumnName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ColumnName, _s_strColumnName, _ptr_strColumnName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnIndexOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetColumnIndexOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pColumnIndex {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ColumnIndex); err != nil {
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

func (o *xxx_GetColumnIndexOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pColumnIndex {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ColumnIndex); err != nil {
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

// GetColumnIndexRequest structure represents the GetColumnIndex operation request
type GetColumnIndexRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResultColumn int32          `idl:"name:fResultColumn" json:"result_column"`
	ColumnName   *oaut.String   `idl:"name:strColumnName" json:"column_name"`
}

func (o *GetColumnIndexRequest) xxx_ToOp(ctx context.Context, op *xxx_GetColumnIndexOperation) *xxx_GetColumnIndexOperation {
	if op == nil {
		op = &xxx_GetColumnIndexOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResultColumn = o.ResultColumn
	op.ColumnName = o.ColumnName
	return op
}

func (o *GetColumnIndexRequest) xxx_FromOp(ctx context.Context, op *xxx_GetColumnIndexOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResultColumn = op.ResultColumn
	o.ColumnName = op.ColumnName
}
func (o *GetColumnIndexRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetColumnIndexRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetColumnIndexOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetColumnIndexResponse structure represents the GetColumnIndex operation response
type GetColumnIndexResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnIndex int32          `idl:"name:pColumnIndex" json:"column_index"`
	// Return: The GetColumnIndex return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetColumnIndexResponse) xxx_ToOp(ctx context.Context, op *xxx_GetColumnIndexOperation) *xxx_GetColumnIndexOperation {
	if op == nil {
		op = &xxx_GetColumnIndexOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ColumnIndex = o.ColumnIndex
	op.Return = o.Return
	return op
}

func (o *GetColumnIndexResponse) xxx_FromOp(ctx context.Context, op *xxx_GetColumnIndexOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ColumnIndex = op.ColumnIndex
	o.Return = op.Return
}
func (o *GetColumnIndexResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetColumnIndexResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetColumnIndexOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResultColumnCountOperation structure represents the SetResultColumnCount operation
type xxx_SetResultColumnCountOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResultColumnCount int32          `idl:"name:cResultColumn" json:"result_column_count"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResultColumnCountOperation) OpNum() int { return 11 }

func (o *xxx_SetResultColumnCountOperation) OpName() string {
	return "/ICertView/v0/SetResultColumnCount"
}

func (o *xxx_SetResultColumnCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ResultColumnCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cResultColumn {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ResultColumnCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResultColumnCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetResultColumnCountRequest structure represents the SetResultColumnCount operation request
type SetResultColumnCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResultColumnCount int32          `idl:"name:cResultColumn" json:"result_column_count"`
}

func (o *SetResultColumnCountRequest) xxx_ToOp(ctx context.Context, op *xxx_SetResultColumnCountOperation) *xxx_SetResultColumnCountOperation {
	if op == nil {
		op = &xxx_SetResultColumnCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResultColumnCount = o.ResultColumnCount
	return op
}

func (o *SetResultColumnCountRequest) xxx_FromOp(ctx context.Context, op *xxx_SetResultColumnCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResultColumnCount = op.ResultColumnCount
}
func (o *SetResultColumnCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetResultColumnCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResultColumnCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResultColumnCountResponse structure represents the SetResultColumnCount operation response
type SetResultColumnCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetResultColumnCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetResultColumnCountResponse) xxx_ToOp(ctx context.Context, op *xxx_SetResultColumnCountOperation) *xxx_SetResultColumnCountOperation {
	if op == nil {
		op = &xxx_SetResultColumnCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetResultColumnCountResponse) xxx_FromOp(ctx context.Context, op *xxx_SetResultColumnCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetResultColumnCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetResultColumnCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResultColumnCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResultColumnOperation structure represents the SetResultColumn operation
type xxx_SetResultColumnOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnIndex int32          `idl:"name:ColumnIndex" json:"column_index"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResultColumnOperation) OpNum() int { return 12 }

func (o *xxx_SetResultColumnOperation) OpName() string { return "/ICertView/v0/SetResultColumn" }

func (o *xxx_SetResultColumnOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ColumnIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ColumnIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ColumnIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ColumnIndex); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResultColumnOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetResultColumnOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetResultColumnRequest structure represents the SetResultColumn operation request
type SetResultColumnRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ColumnIndex int32          `idl:"name:ColumnIndex" json:"column_index"`
}

func (o *SetResultColumnRequest) xxx_ToOp(ctx context.Context, op *xxx_SetResultColumnOperation) *xxx_SetResultColumnOperation {
	if op == nil {
		op = &xxx_SetResultColumnOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ColumnIndex = o.ColumnIndex
	return op
}

func (o *SetResultColumnRequest) xxx_FromOp(ctx context.Context, op *xxx_SetResultColumnOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ColumnIndex = op.ColumnIndex
}
func (o *SetResultColumnRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetResultColumnRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResultColumnOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResultColumnResponse structure represents the SetResultColumn operation response
type SetResultColumnResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetResultColumn return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetResultColumnResponse) xxx_ToOp(ctx context.Context, op *xxx_SetResultColumnOperation) *xxx_SetResultColumnOperation {
	if op == nil {
		op = &xxx_SetResultColumnOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetResultColumnResponse) xxx_FromOp(ctx context.Context, op *xxx_SetResultColumnOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetResultColumnResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetResultColumnResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResultColumnOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRestrictionOperation structure represents the SetRestriction operation
type xxx_SetRestrictionOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ColumnIndex  int32          `idl:"name:ColumnIndex" json:"column_index"`
	SeekOperator int32          `idl:"name:SeekOperator" json:"seek_operator"`
	SortOrder    int32          `idl:"name:SortOrder" json:"sort_order"`
	Value        *oaut.Variant  `idl:"name:pvarValue" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRestrictionOperation) OpNum() int { return 13 }

func (o *xxx_SetRestrictionOperation) OpName() string { return "/ICertView/v0/SetRestriction" }

func (o *xxx_SetRestrictionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRestrictionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ColumnIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ColumnIndex); err != nil {
			return err
		}
	}
	// SeekOperator {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.SeekOperator); err != nil {
			return err
		}
	}
	// SortOrder {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.SortOrder); err != nil {
			return err
		}
	}
	// pvarValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			_ptr_pvarValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pvarValue); err != nil {
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
	return nil
}

func (o *xxx_SetRestrictionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ColumnIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ColumnIndex); err != nil {
			return err
		}
	}
	// SeekOperator {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.SeekOperator); err != nil {
			return err
		}
	}
	// SortOrder {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.SortOrder); err != nil {
			return err
		}
	}
	// pvarValue {in} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Value, _s_pvarValue, _ptr_pvarValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRestrictionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRestrictionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetRestrictionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetRestrictionRequest structure represents the SetRestriction operation request
type SetRestrictionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	ColumnIndex  int32          `idl:"name:ColumnIndex" json:"column_index"`
	SeekOperator int32          `idl:"name:SeekOperator" json:"seek_operator"`
	SortOrder    int32          `idl:"name:SortOrder" json:"sort_order"`
	Value        *oaut.Variant  `idl:"name:pvarValue" json:"value"`
}

func (o *SetRestrictionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetRestrictionOperation) *xxx_SetRestrictionOperation {
	if op == nil {
		op = &xxx_SetRestrictionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ColumnIndex = o.ColumnIndex
	op.SeekOperator = o.SeekOperator
	op.SortOrder = o.SortOrder
	op.Value = o.Value
	return op
}

func (o *SetRestrictionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRestrictionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ColumnIndex = op.ColumnIndex
	o.SeekOperator = op.SeekOperator
	o.SortOrder = op.SortOrder
	o.Value = op.Value
}
func (o *SetRestrictionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetRestrictionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRestrictionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRestrictionResponse structure represents the SetRestriction operation response
type SetRestrictionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetRestriction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetRestrictionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetRestrictionOperation) *xxx_SetRestrictionOperation {
	if op == nil {
		op = &xxx_SetRestrictionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetRestrictionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRestrictionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetRestrictionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetRestrictionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRestrictionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenViewOperation structure represents the OpenView operation
type xxx_OpenViewOperation struct {
	This   *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Enum   *certview.EnumCertViewRow `idl:"name:ppenum" json:"enum"`
	Return int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenViewOperation) OpNum() int { return 14 }

func (o *xxx_OpenViewOperation) OpName() string { return "/ICertView/v0/OpenView" }

func (o *xxx_OpenViewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OpenViewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_OpenViewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppenum {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumCERTVIEWROW}(interface))
	{
		if o.Enum != nil {
			_ptr_ppenum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&certview.EnumCertViewRow{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppenum); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenViewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppenum {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumCERTVIEWROW}(interface))
	{
		_ptr_ppenum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &certview.EnumCertViewRow{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppenum := func(ptr interface{}) { o.Enum = *ptr.(**certview.EnumCertViewRow) }
		if err := w.ReadPointer(&o.Enum, _s_ppenum, _ptr_ppenum); err != nil {
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

// OpenViewRequest structure represents the OpenView operation request
type OpenViewRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *OpenViewRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenViewOperation) *xxx_OpenViewOperation {
	if op == nil {
		op = &xxx_OpenViewOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *OpenViewRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenViewOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *OpenViewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenViewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenViewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenViewResponse structure represents the OpenView operation response
type OpenViewResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Enum *certview.EnumCertViewRow `idl:"name:ppenum" json:"enum"`
	// Return: The OpenView return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenViewResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenViewOperation) *xxx_OpenViewOperation {
	if op == nil {
		op = &xxx_OpenViewOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *OpenViewResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenViewOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *OpenViewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenViewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenViewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
