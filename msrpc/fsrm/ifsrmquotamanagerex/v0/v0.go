package ifsrmquotamanagerex

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmquotamanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotamanager/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = ifsrmquotamanager.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaManagerEx interface identifier 4846cb01-d430-494f-abb4-b1054999fb09
	QuotaManagerExIID = &dcom.IID{Data1: 0x4846cb01, Data2: 0xd430, Data3: 0x494f, Data4: []byte{0xab, 0xb4, 0xb1, 0x05, 0x49, 0x99, 0xfb, 0x09}}
	// Syntax UUID
	QuotaManagerExSyntaxUUID = &uuid.UUID{TimeLow: 0x4846cb01, TimeMid: 0xd430, TimeHiAndVersion: 0x494f, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xb4, Node: [6]uint8{0xb1, 0x5, 0x49, 0x99, 0xfb, 0x9}}
	// Syntax ID
	QuotaManagerExSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaManagerExSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaManagerEx interface.
type QuotaManagerExClient interface {

	// IFsrmQuotaManager retrieval method.
	QuotaManager() ifsrmquotamanager.QuotaManagerClient

	// The IsAffectedByQuota method retrieves a value that determines whether a specified
	// path is subject to a Persisted Directory Quota (section 3.2.1.2.1.1).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The affected parameter is NULL. |
	//	|                         | The options parameter is not a valid FsrmEnumOptions (section 2.2.1.2.5) value.  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004003 E_POINTER    | The path parameter is NULL.                                                      |
	//	+-------------------------+----------------------------------------------------------------------------------+
	IsAffectedByQuota(context.Context, *IsAffectedByQuotaRequest, ...dcerpc.CallOption) (*IsAffectedByQuotaResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaManagerExClient
}

type xxx_DefaultQuotaManagerExClient struct {
	ifsrmquotamanager.QuotaManagerClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaManagerExClient) QuotaManager() ifsrmquotamanager.QuotaManagerClient {
	return o.QuotaManagerClient
}

func (o *xxx_DefaultQuotaManagerExClient) IsAffectedByQuota(ctx context.Context, in *IsAffectedByQuotaRequest, opts ...dcerpc.CallOption) (*IsAffectedByQuotaResponse, error) {
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
	out := &IsAffectedByQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaManagerExClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaManagerExClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaManagerExClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaManagerExClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaManagerExClient{
		QuotaManagerClient: o.QuotaManagerClient.IPID(ctx, ipid),
		cc:                 o.cc,
		ipid:               ipid,
	}
}

func NewQuotaManagerExClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaManagerExClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaManagerExSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmquotamanager.NewQuotaManagerClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaManagerExClient{
		QuotaManagerClient: base,
		cc:                 cc,
		ipid:               ipid,
	}, nil
}

// xxx_IsAffectedByQuotaOperation structure represents the IsAffectedByQuota operation
type xxx_IsAffectedByQuotaOperation struct {
	This     *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Path     *oaut.String     `idl:"name:path" json:"path"`
	Options  fsrm.EnumOptions `idl:"name:options" json:"options"`
	Affected int16            `idl:"name:affected" json:"affected"`
	Return   int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_IsAffectedByQuotaOperation) OpNum() int { return 19 }

func (o *xxx_IsAffectedByQuotaOperation) OpName() string {
	return "/IFsrmQuotaManagerEx/v0/IsAffectedByQuota"
}

func (o *xxx_IsAffectedByQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsAffectedByQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsAffectedByQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsAffectedByQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsAffectedByQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// affected {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Affected); err != nil {
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

func (o *xxx_IsAffectedByQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// affected {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Affected); err != nil {
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

// IsAffectedByQuotaRequest structure represents the IsAffectedByQuota operation request
type IsAffectedByQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: The local directory path to determine whether a Persisted Directory Quota applies.
	// The maximum length of this string MUST be 260 characters.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: The options to use when checking for a Persisted Directory Quota. For possible
	// values, see the FsrmEnumOptions (section 2.2.1.2.5) enumeration. It is ignored on
	// receipt.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *IsAffectedByQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_IsAffectedByQuotaOperation) *xxx_IsAffectedByQuotaOperation {
	if op == nil {
		op = &xxx_IsAffectedByQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	op.Options = o.Options
	return op
}

func (o *IsAffectedByQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_IsAffectedByQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *IsAffectedByQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsAffectedByQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsAffectedByQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsAffectedByQuotaResponse structure represents the IsAffectedByQuota operation response
type IsAffectedByQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// affected: Pointer to a Boolean variable that returns whether a specified path is
	// subject to a Persisted Directory Quota.
	Affected int16 `idl:"name:affected" json:"affected"`
	// Return: The IsAffectedByQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsAffectedByQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_IsAffectedByQuotaOperation) *xxx_IsAffectedByQuotaOperation {
	if op == nil {
		op = &xxx_IsAffectedByQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Affected = o.Affected
	op.Return = o.Return
	return op
}

func (o *IsAffectedByQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_IsAffectedByQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Affected = op.Affected
	o.Return = op.Return
}
func (o *IsAffectedByQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsAffectedByQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsAffectedByQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
