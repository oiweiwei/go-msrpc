package iapphostconfiglocation

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
	// IAppHostConfigLocation interface identifier 370af178-7758-4dad-8146-7391f6e18585
	AppHostConfigLocationIID = &dcom.IID{Data1: 0x370af178, Data2: 0x7758, Data3: 0x4dad, Data4: []byte{0x81, 0x46, 0x73, 0x91, 0xf6, 0xe1, 0x85, 0x85}}
	// Syntax UUID
	AppHostConfigLocationSyntaxUUID = &uuid.UUID{TimeLow: 0x370af178, TimeMid: 0x7758, TimeHiAndVersion: 0x4dad, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x46, Node: [6]uint8{0x73, 0x91, 0xf6, 0xe1, 0x85, 0x85}}
	// Syntax ID
	AppHostConfigLocationSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostConfigLocationSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostConfigLocation interface.
type AppHostConfigLocationClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Path operation.
	GetPath(context.Context, *GetPathRequest, ...dcerpc.CallOption) (*GetPathResponse, error)

	// Count operation.
	GetCount(context.Context, *GetCountRequest, ...dcerpc.CallOption) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest, ...dcerpc.CallOption) (*GetItemResponse, error)

	// AddConfigSection operation.
	AddConfigSection(context.Context, *AddConfigSectionRequest, ...dcerpc.CallOption) (*AddConfigSectionResponse, error)

	// DeleteConfigSection operation.
	DeleteConfigSection(context.Context, *DeleteConfigSectionRequest, ...dcerpc.CallOption) (*DeleteConfigSectionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostConfigLocationClient
}

type xxx_DefaultAppHostConfigLocationClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostConfigLocationClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostConfigLocationClient) GetPath(ctx context.Context, in *GetPathRequest, opts ...dcerpc.CallOption) (*GetPathResponse, error) {
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
	out := &GetPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigLocationClient) GetCount(ctx context.Context, in *GetCountRequest, opts ...dcerpc.CallOption) (*GetCountResponse, error) {
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
	out := &GetCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigLocationClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...dcerpc.CallOption) (*GetItemResponse, error) {
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
	out := &GetItemResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigLocationClient) AddConfigSection(ctx context.Context, in *AddConfigSectionRequest, opts ...dcerpc.CallOption) (*AddConfigSectionResponse, error) {
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
	out := &AddConfigSectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigLocationClient) DeleteConfigSection(ctx context.Context, in *DeleteConfigSectionRequest, opts ...dcerpc.CallOption) (*DeleteConfigSectionResponse, error) {
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
	out := &DeleteConfigSectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigLocationClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostConfigLocationClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostConfigLocationClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostConfigLocationClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostConfigLocationClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostConfigLocationClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostConfigLocationClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostConfigLocationSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostConfigLocationClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPathOperation structure represents the Path operation
type xxx_GetPathOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LocationPath *oaut.String   `idl:"name:pbstrLocationPath" json:"location_path"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPathOperation) OpNum() int { return 3 }

func (o *xxx_GetPathOperation) OpName() string { return "/IAppHostConfigLocation/v0/Path" }

func (o *xxx_GetPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrLocationPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LocationPath != nil {
			_ptr_pbstrLocationPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LocationPath != nil {
					if err := o.LocationPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LocationPath, _ptr_pbstrLocationPath); err != nil {
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

func (o *xxx_GetPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrLocationPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrLocationPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LocationPath == nil {
				o.LocationPath = &oaut.String{}
			}
			if err := o.LocationPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrLocationPath := func(ptr interface{}) { o.LocationPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LocationPath, _s_pbstrLocationPath, _ptr_pbstrLocationPath); err != nil {
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

// GetPathRequest structure represents the Path operation request
type GetPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPathRequest) xxx_ToOp(ctx context.Context) *xxx_GetPathOperation {
	if o == nil {
		return &xxx_GetPathOperation{}
	}
	return &xxx_GetPathOperation{
		This: o.This,
	}
}

func (o *GetPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPathResponse structure represents the Path operation response
type GetPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	LocationPath *oaut.String   `idl:"name:pbstrLocationPath" json:"location_path"`
	// Return: The Path return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPathResponse) xxx_ToOp(ctx context.Context) *xxx_GetPathOperation {
	if o == nil {
		return &xxx_GetPathOperation{}
	}
	return &xxx_GetPathOperation{
		That:         o.That,
		LocationPath: o.LocationPath,
		Return:       o.Return,
	}
}

func (o *GetPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LocationPath = op.LocationPath
	o.Return = op.Return
}
func (o *GetPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCountOperation structure represents the Count operation
type xxx_GetCountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count  uint32         `idl:"name:pcCount" json:"count"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCountOperation) OpNum() int { return 4 }

func (o *xxx_GetCountOperation) OpName() string { return "/IAppHostConfigLocation/v0/Count" }

func (o *xxx_GetCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
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

func (o *xxx_GetCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcCount {out, retval} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
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

// GetCountRequest structure represents the Count operation request
type GetCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCountRequest) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		This: o.This,
	}
}

func (o *GetCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCountResponse structure represents the Count operation response
type GetCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Count uint32         `idl:"name:pcCount" json:"count"`
	// Return: The Count return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCountResponse) xxx_ToOp(ctx context.Context) *xxx_GetCountOperation {
	if o == nil {
		return &xxx_GetCountOperation{}
	}
	return &xxx_GetCountOperation{
		That:   o.That,
		Count:  o.Count,
		Return: o.Return,
	}
}

func (o *GetCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetItemOperation structure represents the Item operation
type xxx_GetItemOperation struct {
	This    *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Index   *oaut.Variant        `idl:"name:cIndex" json:"index"`
	Section *iisa.AppHostElement `idl:"name:ppSection" json:"section"`
	Return  int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetItemOperation) OpNum() int { return 5 }

func (o *xxx_GetItemOperation) OpName() string { return "/IAppHostConfigLocation/v0/Item" }

func (o *xxx_GetItemOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cIndex {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index != nil {
			if err := o.Index.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cIndex {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index == nil {
			o.Index = &oaut.Variant{}
		}
		if err := o.Index.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetItemOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.Section != nil {
			_ptr_ppSection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Section != nil {
					if err := o.Section.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Section, _ptr_ppSection); err != nil {
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

func (o *xxx_GetItemOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppSection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Section == nil {
				o.Section = &iisa.AppHostElement{}
			}
			if err := o.Section.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSection := func(ptr interface{}) { o.Section = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.Section, _s_ppSection, _ptr_ppSection); err != nil {
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

// GetItemRequest structure represents the Item operation request
type GetItemRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index *oaut.Variant  `idl:"name:cIndex" json:"index"`
}

func (o *GetItemRequest) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetItemRequest) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetItemRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetItemRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetItemResponse structure represents the Item operation response
type GetItemResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Section *iisa.AppHostElement `idl:"name:ppSection" json:"section"`
	// Return: The Item return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetItemResponse) xxx_ToOp(ctx context.Context) *xxx_GetItemOperation {
	if o == nil {
		return &xxx_GetItemOperation{}
	}
	return &xxx_GetItemOperation{
		That:    o.That,
		Section: o.Section,
		Return:  o.Return,
	}
}

func (o *GetItemResponse) xxx_FromOp(ctx context.Context, op *xxx_GetItemOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Section = op.Section
	o.Return = op.Return
}
func (o *GetItemResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetItemResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetItemOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddConfigSectionOperation structure represents the AddConfigSection operation
type xxx_AddConfigSectionOperation struct {
	This         *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SectionName  *oaut.String         `idl:"name:bstrSectionName" json:"section_name"`
	AdminElement *iisa.AppHostElement `idl:"name:ppAdminElement" json:"admin_element"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_AddConfigSectionOperation) OpNum() int { return 6 }

func (o *xxx_AddConfigSectionOperation) OpName() string {
	return "/IAppHostConfigLocation/v0/AddConfigSection"
}

func (o *xxx_AddConfigSectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddConfigSectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SectionName != nil {
			_ptr_bstrSectionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SectionName != nil {
					if err := o.SectionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SectionName, _ptr_bstrSectionName); err != nil {
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

func (o *xxx_AddConfigSectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSectionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SectionName == nil {
				o.SectionName = &oaut.String{}
			}
			if err := o.SectionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSectionName := func(ptr interface{}) { o.SectionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SectionName, _s_bstrSectionName, _ptr_bstrSectionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddConfigSectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddConfigSectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAdminElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.AdminElement != nil {
			_ptr_ppAdminElement := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AdminElement != nil {
					if err := o.AdminElement.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminElement, _ptr_ppAdminElement); err != nil {
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

func (o *xxx_AddConfigSectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAdminElement {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppAdminElement := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AdminElement == nil {
				o.AdminElement = &iisa.AppHostElement{}
			}
			if err := o.AdminElement.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAdminElement := func(ptr interface{}) { o.AdminElement = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.AdminElement, _s_ppAdminElement, _ptr_ppAdminElement); err != nil {
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

// AddConfigSectionRequest structure represents the AddConfigSection operation request
type AddConfigSectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	SectionName *oaut.String   `idl:"name:bstrSectionName" json:"section_name"`
}

func (o *AddConfigSectionRequest) xxx_ToOp(ctx context.Context) *xxx_AddConfigSectionOperation {
	if o == nil {
		return &xxx_AddConfigSectionOperation{}
	}
	return &xxx_AddConfigSectionOperation{
		This:        o.This,
		SectionName: o.SectionName,
	}
}

func (o *AddConfigSectionRequest) xxx_FromOp(ctx context.Context, op *xxx_AddConfigSectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SectionName = op.SectionName
}
func (o *AddConfigSectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddConfigSectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddConfigSectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddConfigSectionResponse structure represents the AddConfigSection operation response
type AddConfigSectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	AdminElement *iisa.AppHostElement `idl:"name:ppAdminElement" json:"admin_element"`
	// Return: The AddConfigSection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddConfigSectionResponse) xxx_ToOp(ctx context.Context) *xxx_AddConfigSectionOperation {
	if o == nil {
		return &xxx_AddConfigSectionOperation{}
	}
	return &xxx_AddConfigSectionOperation{
		That:         o.That,
		AdminElement: o.AdminElement,
		Return:       o.Return,
	}
}

func (o *AddConfigSectionResponse) xxx_FromOp(ctx context.Context, op *xxx_AddConfigSectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AdminElement = op.AdminElement
	o.Return = op.Return
}
func (o *AddConfigSectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddConfigSectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddConfigSectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteConfigSectionOperation structure represents the DeleteConfigSection operation
type xxx_DeleteConfigSectionOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index  *oaut.Variant  `idl:"name:cIndex" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteConfigSectionOperation) OpNum() int { return 7 }

func (o *xxx_DeleteConfigSectionOperation) OpName() string {
	return "/IAppHostConfigLocation/v0/DeleteConfigSection"
}

func (o *xxx_DeleteConfigSectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigSectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// cIndex {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index != nil {
			if err := o.Index.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigSectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// cIndex {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Index == nil {
			o.Index = &oaut.Variant{}
		}
		if err := o.Index.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigSectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigSectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteConfigSectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteConfigSectionRequest structure represents the DeleteConfigSection operation request
type DeleteConfigSectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Index *oaut.Variant  `idl:"name:cIndex" json:"index"`
}

func (o *DeleteConfigSectionRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteConfigSectionOperation {
	if o == nil {
		return &xxx_DeleteConfigSectionOperation{}
	}
	return &xxx_DeleteConfigSectionOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *DeleteConfigSectionRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteConfigSectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *DeleteConfigSectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteConfigSectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteConfigSectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteConfigSectionResponse structure represents the DeleteConfigSection operation response
type DeleteConfigSectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteConfigSection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteConfigSectionResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteConfigSectionOperation {
	if o == nil {
		return &xxx_DeleteConfigSectionOperation{}
	}
	return &xxx_DeleteConfigSectionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteConfigSectionResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteConfigSectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteConfigSectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteConfigSectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteConfigSectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
