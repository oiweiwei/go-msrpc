package ifsrmquota

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ifsrmquotaobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotaobject/v0"
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
	_ = ifsrmquotaobject.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuota interface identifier 377f739d-9647-4b8e-97d2-5ffce6d759cd
	QuotaIID = &dcom.IID{Data1: 0x377f739d, Data2: 0x9647, Data3: 0x4b8e, Data4: []byte{0x97, 0xd2, 0x5f, 0xfc, 0xe6, 0xd7, 0x59, 0xcd}}
	// Syntax UUID
	QuotaSyntaxUUID = &uuid.UUID{TimeLow: 0x377f739d, TimeMid: 0x9647, TimeHiAndVersion: 0x4b8e, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xd2, Node: [6]uint8{0x5f, 0xfc, 0xe6, 0xd7, 0x59, 0xcd}}
	// Syntax ID
	QuotaSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuota interface.
type QuotaClient interface {

	// IFsrmQuotaObject retrieval method.
	QuotaObject() ifsrmquotaobject.QuotaObjectClient

	// The QuotaUsed (get) method returns the current, read-only quota usage for this quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------+
	//	|         RETURN          |                             |
	//	|       VALUE/CODE        |         DESCRIPTION         |
	//	|                         |                             |
	//	+-------------------------+-----------------------------+
	//	+-------------------------+-----------------------------+
	//	| 0x80070057 E_INVALIDARG | The used parameter is NULL. |
	//	+-------------------------+-----------------------------+
	GetQuotaUsed(context.Context, *GetQuotaUsedRequest, ...dcerpc.CallOption) (*GetQuotaUsedResponse, error)

	// The QuotaPeakUsage (get) method returns the peak quota usage of this quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------+
	//	|         RETURN          |                                  |
	//	|       VALUE/CODE        |           DESCRIPTION            |
	//	|                         |                                  |
	//	+-------------------------+----------------------------------+
	//	+-------------------------+----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The peakUsage parameter is NULL. |
	//	+-------------------------+----------------------------------+
	GetQuotaPeakUsage(context.Context, *GetQuotaPeakUsageRequest, ...dcerpc.CallOption) (*GetQuotaPeakUsageResponse, error)

	// The QuotaPeakUsageTime (get) method returns the peak quota usage time stamp of this
	// quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------------------+
	//	|         RETURN          |                                          |
	//	|       VALUE/CODE        |               DESCRIPTION                |
	//	|                         |                                          |
	//	+-------------------------+------------------------------------------+
	//	+-------------------------+------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The peakUsageDateTime parameter is NULL. |
	//	+-------------------------+------------------------------------------+
	GetQuotaPeakUsageTime(context.Context, *GetQuotaPeakUsageTimeRequest, ...dcerpc.CallOption) (*GetQuotaPeakUsageTimeResponse, error)

	// The ResetPeakUsage method resets the peak quota usage of this quota to zero and returns
	// S_OK upon successful completion.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	// There are no parameters for this method.
	ResetPeakUsage(context.Context, *ResetPeakUsageRequest, ...dcerpc.CallOption) (*ResetPeakUsageResponse, error)

	// The RefreshUsageProperties method refreshes the quota usage information for the caller's
	// copy of the object.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+-----------------------------------------+
	//	|           RETURN            |                                         |
	//	|         VALUE/CODE          |               DESCRIPTION               |
	//	|                             |                                         |
	//	+-----------------------------+-----------------------------------------+
	//	+-----------------------------+-----------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified quota could not be found. |
	//	+-----------------------------+-----------------------------------------+
	//
	// There are no parameters for this method.
	//
	// If no Persisted Directory Quota exists that has the same Directory Quota.Folder path
	// property as Non-Persisted Directory Quota Instance, the server MUST return FSRM_E_NOT_FOUND.
	//
	// Otherwise, the server MUST reset the quota usage, quota peak usage, and quota peak
	// usage time of the Non-Persisted Directory Quota Instance to the current values stored
	// in the corresponding properties of the Persisted Directory Quota that has the same
	// Directory Quota.Folder path property as this Non-Persisted Directory Quota Instance.
	RefreshUsageProperties(context.Context, *RefreshUsagePropertiesRequest, ...dcerpc.CallOption) (*RefreshUsagePropertiesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaClient
}

type xxx_DefaultQuotaClient struct {
	ifsrmquotaobject.QuotaObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaClient) QuotaObject() ifsrmquotaobject.QuotaObjectClient {
	return o.QuotaObjectClient
}

func (o *xxx_DefaultQuotaClient) GetQuotaUsed(ctx context.Context, in *GetQuotaUsedRequest, opts ...dcerpc.CallOption) (*GetQuotaUsedResponse, error) {
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
	out := &GetQuotaUsedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaClient) GetQuotaPeakUsage(ctx context.Context, in *GetQuotaPeakUsageRequest, opts ...dcerpc.CallOption) (*GetQuotaPeakUsageResponse, error) {
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
	out := &GetQuotaPeakUsageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaClient) GetQuotaPeakUsageTime(ctx context.Context, in *GetQuotaPeakUsageTimeRequest, opts ...dcerpc.CallOption) (*GetQuotaPeakUsageTimeResponse, error) {
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
	out := &GetQuotaPeakUsageTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaClient) ResetPeakUsage(ctx context.Context, in *ResetPeakUsageRequest, opts ...dcerpc.CallOption) (*ResetPeakUsageResponse, error) {
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
	out := &ResetPeakUsageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaClient) RefreshUsageProperties(ctx context.Context, in *RefreshUsagePropertiesRequest, opts ...dcerpc.CallOption) (*RefreshUsagePropertiesResponse, error) {
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
	out := &RefreshUsagePropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaClient{
		QuotaObjectClient: o.QuotaObjectClient.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewQuotaClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmquotaobject.NewQuotaObjectClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaClient{
		QuotaObjectClient: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_GetQuotaUsedOperation structure represents the QuotaUsed operation
type xxx_GetQuotaUsedOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Used   *oaut.Variant  `idl:"name:used" json:"used"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaUsedOperation) OpNum() int { return 28 }

func (o *xxx_GetQuotaUsedOperation) OpName() string { return "/IFsrmQuota/v0/QuotaUsed" }

func (o *xxx_GetQuotaUsedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaUsedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaUsedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaUsedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaUsedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// used {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Used != nil {
			_ptr_used := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Used != nil {
					if err := o.Used.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Used, _ptr_used); err != nil {
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

func (o *xxx_GetQuotaUsedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// used {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_used := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Used == nil {
				o.Used = &oaut.Variant{}
			}
			if err := o.Used.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_used := func(ptr interface{}) { o.Used = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Used, _s_used, _ptr_used); err != nil {
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

// GetQuotaUsedRequest structure represents the QuotaUsed operation request
type GetQuotaUsedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaUsedRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaUsedOperation) *xxx_GetQuotaUsedOperation {
	if op == nil {
		op = &xxx_GetQuotaUsedOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetQuotaUsedRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaUsedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaUsedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaUsedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaUsedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaUsedResponse structure represents the QuotaUsed operation response
type GetQuotaUsedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// used: Pointer to a variable, which upon completion, contains the quota usage for
	// this quota.
	Used *oaut.Variant `idl:"name:used" json:"used"`
	// Return: The QuotaUsed return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaUsedResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaUsedOperation) *xxx_GetQuotaUsedOperation {
	if op == nil {
		op = &xxx_GetQuotaUsedOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Used = op.Used
	o.Return = op.Return
	return op
}

func (o *GetQuotaUsedResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaUsedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Used = op.Used
	o.Return = op.Return
}
func (o *GetQuotaUsedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaUsedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaUsedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuotaPeakUsageOperation structure represents the QuotaPeakUsage operation
type xxx_GetQuotaPeakUsageOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	PeakUsage *oaut.Variant  `idl:"name:peakUsage" json:"peak_usage"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaPeakUsageOperation) OpNum() int { return 29 }

func (o *xxx_GetQuotaPeakUsageOperation) OpName() string { return "/IFsrmQuota/v0/QuotaPeakUsage" }

func (o *xxx_GetQuotaPeakUsageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaPeakUsageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaPeakUsageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaPeakUsageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaPeakUsageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// peakUsage {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PeakUsage != nil {
			_ptr_peakUsage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PeakUsage != nil {
					if err := o.PeakUsage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PeakUsage, _ptr_peakUsage); err != nil {
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

func (o *xxx_GetQuotaPeakUsageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// peakUsage {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_peakUsage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PeakUsage == nil {
				o.PeakUsage = &oaut.Variant{}
			}
			if err := o.PeakUsage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_peakUsage := func(ptr interface{}) { o.PeakUsage = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PeakUsage, _s_peakUsage, _ptr_peakUsage); err != nil {
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

// GetQuotaPeakUsageRequest structure represents the QuotaPeakUsage operation request
type GetQuotaPeakUsageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaPeakUsageRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaPeakUsageOperation) *xxx_GetQuotaPeakUsageOperation {
	if op == nil {
		op = &xxx_GetQuotaPeakUsageOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetQuotaPeakUsageRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaPeakUsageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaPeakUsageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaPeakUsageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaPeakUsageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaPeakUsageResponse structure represents the QuotaPeakUsage operation response
type GetQuotaPeakUsageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// peakUsage: Pointer to a variable that upon completion contains the peak quota usage
	// of this quota.
	PeakUsage *oaut.Variant `idl:"name:peakUsage" json:"peak_usage"`
	// Return: The QuotaPeakUsage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaPeakUsageResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaPeakUsageOperation) *xxx_GetQuotaPeakUsageOperation {
	if op == nil {
		op = &xxx_GetQuotaPeakUsageOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.PeakUsage = op.PeakUsage
	o.Return = op.Return
	return op
}

func (o *GetQuotaPeakUsageResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaPeakUsageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PeakUsage = op.PeakUsage
	o.Return = op.Return
}
func (o *GetQuotaPeakUsageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaPeakUsageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaPeakUsageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuotaPeakUsageTimeOperation structure represents the QuotaPeakUsageTime operation
type xxx_GetQuotaPeakUsageTimeOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	PeakUsageDateTime float64        `idl:"name:peakUsageDateTime" json:"peak_usage_date_time"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaPeakUsageTimeOperation) OpNum() int { return 30 }

func (o *xxx_GetQuotaPeakUsageTimeOperation) OpName() string {
	return "/IFsrmQuota/v0/QuotaPeakUsageTime"
}

func (o *xxx_GetQuotaPeakUsageTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaPeakUsageTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaPeakUsageTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaPeakUsageTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaPeakUsageTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// peakUsageDateTime {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.WriteData(o.PeakUsageDateTime); err != nil {
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

func (o *xxx_GetQuotaPeakUsageTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// peakUsageDateTime {out, retval} (1:{pointer=ref}*(1))(2:{alias=DATE}(float64))
	{
		if err := w.ReadData(&o.PeakUsageDateTime); err != nil {
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

// GetQuotaPeakUsageTimeRequest structure represents the QuotaPeakUsageTime operation request
type GetQuotaPeakUsageTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaPeakUsageTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaPeakUsageTimeOperation) *xxx_GetQuotaPeakUsageTimeOperation {
	if op == nil {
		op = &xxx_GetQuotaPeakUsageTimeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetQuotaPeakUsageTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaPeakUsageTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaPeakUsageTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaPeakUsageTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaPeakUsageTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaPeakUsageTimeResponse structure represents the QuotaPeakUsageTime operation response
type GetQuotaPeakUsageTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// peakUsageDateTime: Pointer to a variable that upon completion contains peak quota
	// usage time stamp of this quota.
	PeakUsageDateTime float64 `idl:"name:peakUsageDateTime" json:"peak_usage_date_time"`
	// Return: The QuotaPeakUsageTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaPeakUsageTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaPeakUsageTimeOperation) *xxx_GetQuotaPeakUsageTimeOperation {
	if op == nil {
		op = &xxx_GetQuotaPeakUsageTimeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.PeakUsageDateTime = op.PeakUsageDateTime
	o.Return = op.Return
	return op
}

func (o *GetQuotaPeakUsageTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaPeakUsageTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PeakUsageDateTime = op.PeakUsageDateTime
	o.Return = op.Return
}
func (o *GetQuotaPeakUsageTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaPeakUsageTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaPeakUsageTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResetPeakUsageOperation structure represents the ResetPeakUsage operation
type xxx_ResetPeakUsageOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetPeakUsageOperation) OpNum() int { return 31 }

func (o *xxx_ResetPeakUsageOperation) OpName() string { return "/IFsrmQuota/v0/ResetPeakUsage" }

func (o *xxx_ResetPeakUsageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetPeakUsageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetPeakUsageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetPeakUsageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetPeakUsageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetPeakUsageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResetPeakUsageRequest structure represents the ResetPeakUsage operation request
type ResetPeakUsageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetPeakUsageRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetPeakUsageOperation) *xxx_ResetPeakUsageOperation {
	if op == nil {
		op = &xxx_ResetPeakUsageOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *ResetPeakUsageRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetPeakUsageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetPeakUsageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetPeakUsageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetPeakUsageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetPeakUsageResponse structure represents the ResetPeakUsage operation response
type ResetPeakUsageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ResetPeakUsage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetPeakUsageResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetPeakUsageOperation) *xxx_ResetPeakUsageOperation {
	if op == nil {
		op = &xxx_ResetPeakUsageOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ResetPeakUsageResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetPeakUsageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetPeakUsageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetPeakUsageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetPeakUsageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshUsagePropertiesOperation structure represents the RefreshUsageProperties operation
type xxx_RefreshUsagePropertiesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshUsagePropertiesOperation) OpNum() int { return 32 }

func (o *xxx_RefreshUsagePropertiesOperation) OpName() string {
	return "/IFsrmQuota/v0/RefreshUsageProperties"
}

func (o *xxx_RefreshUsagePropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshUsagePropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshUsagePropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RefreshUsagePropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshUsagePropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshUsagePropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RefreshUsagePropertiesRequest structure represents the RefreshUsageProperties operation request
type RefreshUsagePropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshUsagePropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_RefreshUsagePropertiesOperation) *xxx_RefreshUsagePropertiesOperation {
	if op == nil {
		op = &xxx_RefreshUsagePropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *RefreshUsagePropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshUsagePropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshUsagePropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RefreshUsagePropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshUsagePropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshUsagePropertiesResponse structure represents the RefreshUsageProperties operation response
type RefreshUsagePropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RefreshUsageProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshUsagePropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_RefreshUsagePropertiesOperation) *xxx_RefreshUsagePropertiesOperation {
	if op == nil {
		op = &xxx_RefreshUsagePropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *RefreshUsagePropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshUsagePropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshUsagePropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RefreshUsagePropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshUsagePropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
