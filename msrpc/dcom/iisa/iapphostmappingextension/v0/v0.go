package iapphostmappingextension

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iisa "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
	_ = iisa.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostMappingExtension interface identifier 31a83ea0-c0e4-4a2c-8a01-353cc2a4c60a
	AppHostMappingExtensionIID = &dcom.IID{Data1: 0x31a83ea0, Data2: 0xc0e4, Data3: 0x4a2c, Data4: []byte{0x8a, 0x01, 0x35, 0x3c, 0xc2, 0xa4, 0xc6, 0x0a}}
	// Syntax UUID
	AppHostMappingExtensionSyntaxUUID = &uuid.UUID{TimeLow: 0x31a83ea0, TimeMid: 0xc0e4, TimeHiAndVersion: 0x4a2c, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x1, Node: [6]uint8{0x35, 0x3c, 0xc2, 0xa4, 0xc6, 0xa}}
	// Syntax ID
	AppHostMappingExtensionSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostMappingExtensionSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostMappingExtension interface.
type AppHostMappingExtensionClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetSiteNameFromSiteId operation.
	GetSiteNameFromSiteID(context.Context, *GetSiteNameFromSiteIDRequest, ...dcerpc.CallOption) (*GetSiteNameFromSiteIDResponse, error)

	// GetSiteIdFromSiteName operation.
	GetSiteIDFromSiteName(context.Context, *GetSiteIDFromSiteNameRequest, ...dcerpc.CallOption) (*GetSiteIDFromSiteNameResponse, error)

	// GetSiteElementFromSiteId operation.
	GetSiteElementFromSiteID(context.Context, *GetSiteElementFromSiteIDRequest, ...dcerpc.CallOption) (*GetSiteElementFromSiteIDResponse, error)

	// MapPath operation.
	MapPath(context.Context, *MapPathRequest, ...dcerpc.CallOption) (*MapPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostMappingExtensionClient
}

type xxx_DefaultAppHostMappingExtensionClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostMappingExtensionClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostMappingExtensionClient) GetSiteNameFromSiteID(ctx context.Context, in *GetSiteNameFromSiteIDRequest, opts ...dcerpc.CallOption) (*GetSiteNameFromSiteIDResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetSiteNameFromSiteIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostMappingExtensionClient) GetSiteIDFromSiteName(ctx context.Context, in *GetSiteIDFromSiteNameRequest, opts ...dcerpc.CallOption) (*GetSiteIDFromSiteNameResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetSiteIDFromSiteNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostMappingExtensionClient) GetSiteElementFromSiteID(ctx context.Context, in *GetSiteElementFromSiteIDRequest, opts ...dcerpc.CallOption) (*GetSiteElementFromSiteIDResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetSiteElementFromSiteIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostMappingExtensionClient) MapPath(ctx context.Context, in *MapPathRequest, opts ...dcerpc.CallOption) (*MapPathResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &MapPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostMappingExtensionClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostMappingExtensionClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostMappingExtensionClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostMappingExtensionClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewAppHostMappingExtensionClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostMappingExtensionClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostMappingExtensionSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostMappingExtensionClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetSiteNameFromSiteIDOperation structure represents the GetSiteNameFromSiteId operation
type xxx_GetSiteNameFromSiteIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SiteID   uint32         `idl:"name:dwSiteId" json:"site_id"`
	SiteName *oaut.String   `idl:"name:pbstrSiteName" json:"site_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSiteNameFromSiteIDOperation) OpNum() int { return 3 }

func (o *xxx_GetSiteNameFromSiteIDOperation) OpName() string {
	return "/IAppHostMappingExtension/v0/GetSiteNameFromSiteId"
}

func (o *xxx_GetSiteNameFromSiteIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteNameFromSiteIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwSiteId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SiteID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteNameFromSiteIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwSiteId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SiteID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteNameFromSiteIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteNameFromSiteIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrSiteName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SiteName != nil {
			_ptr_pbstrSiteName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteName != nil {
					if err := o.SiteName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteName, _ptr_pbstrSiteName); err != nil {
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

func (o *xxx_GetSiteNameFromSiteIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSiteName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSiteName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SiteName == nil {
				o.SiteName = &oaut.String{}
			}
			if err := o.SiteName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSiteName := func(ptr interface{}) { o.SiteName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SiteName, _s_pbstrSiteName, _ptr_pbstrSiteName); err != nil {
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

// GetSiteNameFromSiteIDRequest structure represents the GetSiteNameFromSiteId operation request
type GetSiteNameFromSiteIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	SiteID uint32         `idl:"name:dwSiteId" json:"site_id"`
}

func (o *GetSiteNameFromSiteIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetSiteNameFromSiteIDOperation {
	if o == nil {
		return &xxx_GetSiteNameFromSiteIDOperation{}
	}
	return &xxx_GetSiteNameFromSiteIDOperation{
		This:   o.This,
		SiteID: o.SiteID,
	}
}

func (o *GetSiteNameFromSiteIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSiteNameFromSiteIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SiteID = op.SiteID
}
func (o *GetSiteNameFromSiteIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSiteNameFromSiteIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteNameFromSiteIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSiteNameFromSiteIDResponse structure represents the GetSiteNameFromSiteId operation response
type GetSiteNameFromSiteIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SiteName *oaut.String   `idl:"name:pbstrSiteName" json:"site_name"`
	// Return: The GetSiteNameFromSiteId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSiteNameFromSiteIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetSiteNameFromSiteIDOperation {
	if o == nil {
		return &xxx_GetSiteNameFromSiteIDOperation{}
	}
	return &xxx_GetSiteNameFromSiteIDOperation{
		That:     o.That,
		SiteName: o.SiteName,
		Return:   o.Return,
	}
}

func (o *GetSiteNameFromSiteIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSiteNameFromSiteIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SiteName = op.SiteName
	o.Return = op.Return
}
func (o *GetSiteNameFromSiteIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSiteNameFromSiteIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteNameFromSiteIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSiteIDFromSiteNameOperation structure represents the GetSiteIdFromSiteName operation
type xxx_GetSiteIDFromSiteNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SiteName *oaut.String   `idl:"name:bstrSiteName" json:"site_name"`
	SiteID   uint32         `idl:"name:pdwSiteId" json:"site_id"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSiteIDFromSiteNameOperation) OpNum() int { return 4 }

func (o *xxx_GetSiteIDFromSiteNameOperation) OpName() string {
	return "/IAppHostMappingExtension/v0/GetSiteIdFromSiteName"
}

func (o *xxx_GetSiteIDFromSiteNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteIDFromSiteNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSiteName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SiteName != nil {
			_ptr_bstrSiteName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteName != nil {
					if err := o.SiteName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteName, _ptr_bstrSiteName); err != nil {
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

func (o *xxx_GetSiteIDFromSiteNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSiteName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSiteName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SiteName == nil {
				o.SiteName = &oaut.String{}
			}
			if err := o.SiteName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSiteName := func(ptr interface{}) { o.SiteName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SiteName, _s_bstrSiteName, _ptr_bstrSiteName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteIDFromSiteNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteIDFromSiteNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwSiteId {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SiteID); err != nil {
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

func (o *xxx_GetSiteIDFromSiteNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwSiteId {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SiteID); err != nil {
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

// GetSiteIDFromSiteNameRequest structure represents the GetSiteIdFromSiteName operation request
type GetSiteIDFromSiteNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	SiteName *oaut.String   `idl:"name:bstrSiteName" json:"site_name"`
}

func (o *GetSiteIDFromSiteNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetSiteIDFromSiteNameOperation {
	if o == nil {
		return &xxx_GetSiteIDFromSiteNameOperation{}
	}
	return &xxx_GetSiteIDFromSiteNameOperation{
		This:     o.This,
		SiteName: o.SiteName,
	}
}

func (o *GetSiteIDFromSiteNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSiteIDFromSiteNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SiteName = op.SiteName
}
func (o *GetSiteIDFromSiteNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSiteIDFromSiteNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteIDFromSiteNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSiteIDFromSiteNameResponse structure represents the GetSiteIdFromSiteName operation response
type GetSiteIDFromSiteNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	SiteID uint32         `idl:"name:pdwSiteId" json:"site_id"`
	// Return: The GetSiteIdFromSiteName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSiteIDFromSiteNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetSiteIDFromSiteNameOperation {
	if o == nil {
		return &xxx_GetSiteIDFromSiteNameOperation{}
	}
	return &xxx_GetSiteIDFromSiteNameOperation{
		That:   o.That,
		SiteID: o.SiteID,
		Return: o.Return,
	}
}

func (o *GetSiteIDFromSiteNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSiteIDFromSiteNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SiteID = op.SiteID
	o.Return = op.Return
}
func (o *GetSiteIDFromSiteNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSiteIDFromSiteNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteIDFromSiteNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSiteElementFromSiteIDOperation structure represents the GetSiteElementFromSiteId operation
type xxx_GetSiteElementFromSiteIDOperation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SiteID      uint32               `idl:"name:dwSiteId" json:"site_id"`
	SiteElement *iisa.AppHostElement `idl:"name:ppSiteElement" json:"site_element"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSiteElementFromSiteIDOperation) OpNum() int { return 5 }

func (o *xxx_GetSiteElementFromSiteIDOperation) OpName() string {
	return "/IAppHostMappingExtension/v0/GetSiteElementFromSiteId"
}

func (o *xxx_GetSiteElementFromSiteIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteElementFromSiteIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwSiteId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SiteID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteElementFromSiteIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwSiteId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SiteID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteElementFromSiteIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSiteElementFromSiteIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSiteElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.SiteElement != nil {
			_ptr_ppSiteElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteElement != nil {
					if err := o.SiteElement.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteElement, _ptr_ppSiteElement); err != nil {
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

func (o *xxx_GetSiteElementFromSiteIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSiteElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppSiteElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SiteElement == nil {
				o.SiteElement = &iisa.AppHostElement{}
			}
			if err := o.SiteElement.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSiteElement := func(ptr interface{}) { o.SiteElement = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.SiteElement, _s_ppSiteElement, _ptr_ppSiteElement); err != nil {
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

// GetSiteElementFromSiteIDRequest structure represents the GetSiteElementFromSiteId operation request
type GetSiteElementFromSiteIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	SiteID uint32         `idl:"name:dwSiteId" json:"site_id"`
}

func (o *GetSiteElementFromSiteIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetSiteElementFromSiteIDOperation {
	if o == nil {
		return &xxx_GetSiteElementFromSiteIDOperation{}
	}
	return &xxx_GetSiteElementFromSiteIDOperation{
		This:   o.This,
		SiteID: o.SiteID,
	}
}

func (o *GetSiteElementFromSiteIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSiteElementFromSiteIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SiteID = op.SiteID
}
func (o *GetSiteElementFromSiteIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSiteElementFromSiteIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteElementFromSiteIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSiteElementFromSiteIDResponse structure represents the GetSiteElementFromSiteId operation response
type GetSiteElementFromSiteIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SiteElement *iisa.AppHostElement `idl:"name:ppSiteElement" json:"site_element"`
	// Return: The GetSiteElementFromSiteId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSiteElementFromSiteIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetSiteElementFromSiteIDOperation {
	if o == nil {
		return &xxx_GetSiteElementFromSiteIDOperation{}
	}
	return &xxx_GetSiteElementFromSiteIDOperation{
		That:        o.That,
		SiteElement: o.SiteElement,
		Return:      o.Return,
	}
}

func (o *GetSiteElementFromSiteIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSiteElementFromSiteIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SiteElement = op.SiteElement
	o.Return = op.Return
}
func (o *GetSiteElementFromSiteIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSiteElementFromSiteIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSiteElementFromSiteIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MapPathOperation structure represents the MapPath operation
type xxx_MapPathOperation struct {
	This                    *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SiteName                *oaut.String         `idl:"name:bstrSiteName" json:"site_name"`
	VirtualPath             *oaut.String         `idl:"name:bstrVirtualPath" json:"virtual_path"`
	PhysicalPath            *oaut.String         `idl:"name:pbstrPhysicalPath" json:"physical_path"`
	VirtualDirectoryElement *iisa.AppHostElement `idl:"name:ppVirtualDirectoryElement" json:"virtual_directory_element"`
	ApplicationElement      *iisa.AppHostElement `idl:"name:ppApplicationElement" json:"application_element"`
	Return                  int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_MapPathOperation) OpNum() int { return 6 }

func (o *xxx_MapPathOperation) OpName() string { return "/IAppHostMappingExtension/v0/MapPath" }

func (o *xxx_MapPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSiteName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SiteName != nil {
			_ptr_bstrSiteName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SiteName != nil {
					if err := o.SiteName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SiteName, _ptr_bstrSiteName); err != nil {
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
	// bstrVirtualPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.VirtualPath != nil {
			_ptr_bstrVirtualPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VirtualPath != nil {
					if err := o.VirtualPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualPath, _ptr_bstrVirtualPath); err != nil {
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

func (o *xxx_MapPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSiteName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSiteName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SiteName == nil {
				o.SiteName = &oaut.String{}
			}
			if err := o.SiteName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSiteName := func(ptr interface{}) { o.SiteName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SiteName, _s_bstrSiteName, _ptr_bstrSiteName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrVirtualPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrVirtualPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VirtualPath == nil {
				o.VirtualPath = &oaut.String{}
			}
			if err := o.VirtualPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrVirtualPath := func(ptr interface{}) { o.VirtualPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.VirtualPath, _s_bstrVirtualPath, _ptr_bstrVirtualPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MapPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPhysicalPath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PhysicalPath != nil {
			_ptr_pbstrPhysicalPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PhysicalPath != nil {
					if err := o.PhysicalPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PhysicalPath, _ptr_pbstrPhysicalPath); err != nil {
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
	// ppVirtualDirectoryElement {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.VirtualDirectoryElement != nil {
			_ptr_ppVirtualDirectoryElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VirtualDirectoryElement != nil {
					if err := o.VirtualDirectoryElement.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualDirectoryElement, _ptr_ppVirtualDirectoryElement); err != nil {
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
	// ppApplicationElement {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.ApplicationElement != nil {
			_ptr_ppApplicationElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ApplicationElement != nil {
					if err := o.ApplicationElement.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ApplicationElement, _ptr_ppApplicationElement); err != nil {
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

func (o *xxx_MapPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPhysicalPath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPhysicalPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PhysicalPath == nil {
				o.PhysicalPath = &oaut.String{}
			}
			if err := o.PhysicalPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPhysicalPath := func(ptr interface{}) { o.PhysicalPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PhysicalPath, _s_pbstrPhysicalPath, _ptr_pbstrPhysicalPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppVirtualDirectoryElement {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppVirtualDirectoryElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VirtualDirectoryElement == nil {
				o.VirtualDirectoryElement = &iisa.AppHostElement{}
			}
			if err := o.VirtualDirectoryElement.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVirtualDirectoryElement := func(ptr interface{}) { o.VirtualDirectoryElement = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.VirtualDirectoryElement, _s_ppVirtualDirectoryElement, _ptr_ppVirtualDirectoryElement); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppApplicationElement {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppApplicationElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ApplicationElement == nil {
				o.ApplicationElement = &iisa.AppHostElement{}
			}
			if err := o.ApplicationElement.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppApplicationElement := func(ptr interface{}) { o.ApplicationElement = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.ApplicationElement, _s_ppApplicationElement, _ptr_ppApplicationElement); err != nil {
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

// MapPathRequest structure represents the MapPath operation request
type MapPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	SiteName    *oaut.String   `idl:"name:bstrSiteName" json:"site_name"`
	VirtualPath *oaut.String   `idl:"name:bstrVirtualPath" json:"virtual_path"`
}

func (o *MapPathRequest) xxx_ToOp(ctx context.Context) *xxx_MapPathOperation {
	if o == nil {
		return &xxx_MapPathOperation{}
	}
	return &xxx_MapPathOperation{
		This:        o.This,
		SiteName:    o.SiteName,
		VirtualPath: o.VirtualPath,
	}
}

func (o *MapPathRequest) xxx_FromOp(ctx context.Context, op *xxx_MapPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SiteName = op.SiteName
	o.VirtualPath = op.VirtualPath
}
func (o *MapPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MapPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MapPathResponse structure represents the MapPath operation response
type MapPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	PhysicalPath            *oaut.String         `idl:"name:pbstrPhysicalPath" json:"physical_path"`
	VirtualDirectoryElement *iisa.AppHostElement `idl:"name:ppVirtualDirectoryElement" json:"virtual_directory_element"`
	ApplicationElement      *iisa.AppHostElement `idl:"name:ppApplicationElement" json:"application_element"`
	// Return: The MapPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MapPathResponse) xxx_ToOp(ctx context.Context) *xxx_MapPathOperation {
	if o == nil {
		return &xxx_MapPathOperation{}
	}
	return &xxx_MapPathOperation{
		That:                    o.That,
		PhysicalPath:            o.PhysicalPath,
		VirtualDirectoryElement: o.VirtualDirectoryElement,
		ApplicationElement:      o.ApplicationElement,
		Return:                  o.Return,
	}
}

func (o *MapPathResponse) xxx_FromOp(ctx context.Context, op *xxx_MapPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PhysicalPath = op.PhysicalPath
	o.VirtualDirectoryElement = op.VirtualDirectoryElement
	o.ApplicationElement = op.ApplicationElement
	o.Return = op.Return
}
func (o *MapPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MapPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MapPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
