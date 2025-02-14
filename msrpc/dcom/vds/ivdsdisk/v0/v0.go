package ivdsdisk

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
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsDisk interface identifier 07e5c822-f00c-47a1-8fce-b244da56fd06
	DiskIID = &dcom.IID{Data1: 0x07e5c822, Data2: 0xf00c, Data3: 0x47a1, Data4: []byte{0x8f, 0xce, 0xb2, 0x44, 0xda, 0x56, 0xfd, 0x06}}
	// Syntax UUID
	DiskSyntaxUUID = &uuid.UUID{TimeLow: 0x7e5c822, TimeMid: 0xf00c, TimeHiAndVersion: 0x47a1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xce, Node: [6]uint8{0xb2, 0x44, 0xda, 0x56, 0xfd, 0x6}}
	// Syntax ID
	DiskSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DiskSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsDisk interface.
type DiskClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// GetPack operation.
	GetPack(context.Context, *GetPackRequest, ...dcerpc.CallOption) (*GetPackResponse, error)

	// The GetIdentificationData method retrieves information that uniquely identifies a
	// disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetIdentificationData(context.Context, *GetIdentificationDataRequest, ...dcerpc.CallOption) (*GetIdentificationDataResponse, error)

	// QueryExtents operation.
	QueryExtents(context.Context, *QueryExtentsRequest, ...dcerpc.CallOption) (*QueryExtentsResponse, error)

	// The ConvertStyle method converts a disk's partitioning format.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ConvertStyle(context.Context, *ConvertStyleRequest, ...dcerpc.CallOption) (*ConvertStyleResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest, ...dcerpc.CallOption) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest, ...dcerpc.CallOption) (*ClearFlagsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DiskClient
}

type xxx_DefaultDiskClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDiskClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDiskClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) GetPack(ctx context.Context, in *GetPackRequest, opts ...dcerpc.CallOption) (*GetPackResponse, error) {
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
	out := &GetPackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) GetIdentificationData(ctx context.Context, in *GetIdentificationDataRequest, opts ...dcerpc.CallOption) (*GetIdentificationDataResponse, error) {
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
	out := &GetIdentificationDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) QueryExtents(ctx context.Context, in *QueryExtentsRequest, opts ...dcerpc.CallOption) (*QueryExtentsResponse, error) {
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
	out := &QueryExtentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) ConvertStyle(ctx context.Context, in *ConvertStyleRequest, opts ...dcerpc.CallOption) (*ConvertStyleResponse, error) {
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
	out := &ConvertStyleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) SetFlags(ctx context.Context, in *SetFlagsRequest, opts ...dcerpc.CallOption) (*SetFlagsResponse, error) {
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
	out := &SetFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) ClearFlags(ctx context.Context, in *ClearFlagsRequest, opts ...dcerpc.CallOption) (*ClearFlagsResponse, error) {
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
	out := &ClearFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDiskClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDiskClient) IPID(ctx context.Context, ipid *dcom.IPID) DiskClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDiskClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDiskClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DiskClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DiskSyntaxV0_0))...)
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
	return &xxx_DefaultDiskClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This           *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat    `idl:"name:That" json:"that"`
	DiskProperties *vds.DiskProperty `idl:"name:pDiskProperties" json:"disk_properties"`
	Return         int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsDisk/v0/GetProperties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pDiskProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DISK_PROP}(struct))
	{
		if o.DiskProperties != nil {
			if err := o.DiskProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.DiskProperty{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pDiskProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DISK_PROP}(struct))
	{
		if o.DiskProperties == nil {
			o.DiskProperties = &vds.DiskProperty{}
		}
		if err := o.DiskProperties.UnmarshalNDR(ctx, w); err != nil {
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

// GetPropertiesRequest structure represents the GetProperties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the GetProperties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat    `idl:"name:That" json:"that"`
	DiskProperties *vds.DiskProperty `idl:"name:pDiskProperties" json:"disk_properties"`
	// Return: The GetProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
	o.Return = op.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPackOperation structure represents the GetPack operation
type xxx_GetPackOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pack   *vds.Pack      `idl:"name:ppPack" json:"pack"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPackOperation) OpNum() int { return 4 }

func (o *xxx_GetPackOperation) OpName() string { return "/IVdsDisk/v0/GetPack" }

func (o *xxx_GetPackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		if o.Pack != nil {
			_ptr_ppPack := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pack != nil {
					if err := o.Pack.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Pack{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pack, _ptr_ppPack); err != nil {
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

func (o *xxx_GetPackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		_ptr_ppPack := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pack == nil {
				o.Pack = &vds.Pack{}
			}
			if err := o.Pack.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPack := func(ptr interface{}) { o.Pack = *ptr.(**vds.Pack) }
		if err := w.ReadPointer(&o.Pack, _s_ppPack, _ptr_ppPack); err != nil {
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

// GetPackRequest structure represents the GetPack operation request
type GetPackRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPackRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPackOperation) *xxx_GetPackOperation {
	if op == nil {
		op = &xxx_GetPackOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetPackRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPackOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPackResponse structure represents the GetPack operation response
type GetPackResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pack *vds.Pack      `idl:"name:ppPack" json:"pack"`
	// Return: The GetPack return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPackResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPackOperation) *xxx_GetPackOperation {
	if op == nil {
		op = &xxx_GetPackOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
	return op
}

func (o *GetPackResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPackOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
}
func (o *GetPackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIdentificationDataOperation structure represents the GetIdentificationData operation
type xxx_GetIdentificationDataOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	LUNInfo *vds.LUNInformation `idl:"name:pLunInfo" json:"lun_info"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIdentificationDataOperation) OpNum() int { return 5 }

func (o *xxx_GetIdentificationDataOperation) OpName() string {
	return "/IVdsDisk/v0/GetIdentificationData"
}

func (o *xxx_GetIdentificationDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIdentificationDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIdentificationDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIdentificationDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIdentificationDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pLunInfo {out} (1:{pointer=ref}*(1))(2:{alias=VDS_LUN_INFORMATION}(struct))
	{
		if o.LUNInfo != nil {
			if err := o.LUNInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.LUNInformation{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetIdentificationDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pLunInfo {out} (1:{pointer=ref}*(1))(2:{alias=VDS_LUN_INFORMATION}(struct))
	{
		if o.LUNInfo == nil {
			o.LUNInfo = &vds.LUNInformation{}
		}
		if err := o.LUNInfo.UnmarshalNDR(ctx, w); err != nil {
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

// GetIdentificationDataRequest structure represents the GetIdentificationData operation request
type GetIdentificationDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIdentificationDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIdentificationDataOperation) *xxx_GetIdentificationDataOperation {
	if op == nil {
		op = &xxx_GetIdentificationDataOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetIdentificationDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIdentificationDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIdentificationDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIdentificationDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIdentificationDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIdentificationDataResponse structure represents the GetIdentificationData operation response
type GetIdentificationDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pLunInfo: A pointer to a VDS_LUN_INFORMATION structure that, if the operation is
	// successfully completed, receives the LUN information for the disk.
	LUNInfo *vds.LUNInformation `idl:"name:pLunInfo" json:"lun_info"`
	// Return: The GetIdentificationData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIdentificationDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIdentificationDataOperation) *xxx_GetIdentificationDataOperation {
	if op == nil {
		op = &xxx_GetIdentificationDataOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.LUNInfo = op.LUNInfo
	o.Return = op.Return
	return op
}

func (o *GetIdentificationDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIdentificationDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LUNInfo = op.LUNInfo
	o.Return = op.Return
}
func (o *GetIdentificationDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIdentificationDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIdentificationDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryExtentsOperation structure represents the QueryExtents operation
type xxx_QueryExtentsOperation struct {
	This            *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ExtentArray     []*vds.DiskExtent `idl:"name:ppExtentArray;size_is:(, plNumberOfExtents)" json:"extent_array"`
	NumberOfExtents int32             `idl:"name:plNumberOfExtents" json:"number_of_extents"`
	Return          int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryExtentsOperation) OpNum() int { return 6 }

func (o *xxx_QueryExtentsOperation) OpName() string { return "/IVdsDisk/v0/QueryExtents" }

func (o *xxx_QueryExtentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryExtentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryExtentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryExtentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ExtentArray != nil && o.NumberOfExtents == 0 {
		o.NumberOfExtents = int32(len(o.ExtentArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryExtentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_EXTENT}[dim:0,size_is=plNumberOfExtents](struct))
	{
		if o.ExtentArray != nil || o.NumberOfExtents > 0 {
			_ptr_ppExtentArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfExtents)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ExtentArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ExtentArray[i1] != nil {
						if err := o.ExtentArray[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.DiskExtent{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ExtentArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.DiskExtent{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExtentArray, _ptr_ppExtentArray); err != nil {
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
	// plNumberOfExtents {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfExtents); err != nil {
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

func (o *xxx_QueryExtentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_EXTENT}[dim:0,size_is=plNumberOfExtents](struct))
	{
		_ptr_ppExtentArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ExtentArray", sizeInfo[0])
			}
			o.ExtentArray = make([]*vds.DiskExtent, sizeInfo[0])
			for i1 := range o.ExtentArray {
				i1 := i1
				if o.ExtentArray[i1] == nil {
					o.ExtentArray[i1] = &vds.DiskExtent{}
				}
				if err := o.ExtentArray[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppExtentArray := func(ptr interface{}) { o.ExtentArray = *ptr.(*[]*vds.DiskExtent) }
		if err := w.ReadPointer(&o.ExtentArray, _s_ppExtentArray, _ptr_ppExtentArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfExtents {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfExtents); err != nil {
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

// QueryExtentsRequest structure represents the QueryExtents operation request
type QueryExtentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryExtentsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryExtentsOperation) *xxx_QueryExtentsOperation {
	if op == nil {
		op = &xxx_QueryExtentsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *QueryExtentsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryExtentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryExtentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryExtentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryExtentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryExtentsResponse structure represents the QueryExtents operation response
type QueryExtentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ExtentArray     []*vds.DiskExtent `idl:"name:ppExtentArray;size_is:(, plNumberOfExtents)" json:"extent_array"`
	NumberOfExtents int32             `idl:"name:plNumberOfExtents" json:"number_of_extents"`
	// Return: The QueryExtents return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryExtentsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryExtentsOperation) *xxx_QueryExtentsOperation {
	if op == nil {
		op = &xxx_QueryExtentsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.ExtentArray = op.ExtentArray
	o.NumberOfExtents = op.NumberOfExtents
	o.Return = op.Return
	return op
}

func (o *QueryExtentsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryExtentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExtentArray = op.ExtentArray
	o.NumberOfExtents = op.NumberOfExtents
	o.Return = op.Return
}
func (o *QueryExtentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryExtentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryExtentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConvertStyleOperation structure represents the ConvertStyle operation
type xxx_ConvertStyleOperation struct {
	This     *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat     `idl:"name:That" json:"that"`
	NewStyle vds.PartitionStyle `idl:"name:NewStyle" json:"new_style"`
	Return   int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_ConvertStyleOperation) OpNum() int { return 7 }

func (o *xxx_ConvertStyleOperation) OpName() string { return "/IVdsDisk/v0/ConvertStyle" }

func (o *xxx_ConvertStyleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertStyleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// NewStyle {in} (1:{alias=VDS_PARTITION_STYLE}(enum))
	{
		if err := w.WriteEnum(uint16(o.NewStyle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertStyleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// NewStyle {in} (1:{alias=VDS_PARTITION_STYLE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.NewStyle)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertStyleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConvertStyleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ConvertStyleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ConvertStyleRequest structure represents the ConvertStyle operation request
type ConvertStyleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// NewStyle: A value from the VDS_PARTITION_STYLE enumeration that indicates the new
	// partitioning format.
	NewStyle vds.PartitionStyle `idl:"name:NewStyle" json:"new_style"`
}

func (o *ConvertStyleRequest) xxx_ToOp(ctx context.Context, op *xxx_ConvertStyleOperation) *xxx_ConvertStyleOperation {
	if op == nil {
		op = &xxx_ConvertStyleOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.NewStyle = op.NewStyle
	return op
}

func (o *ConvertStyleRequest) xxx_FromOp(ctx context.Context, op *xxx_ConvertStyleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NewStyle = op.NewStyle
}
func (o *ConvertStyleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConvertStyleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConvertStyleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConvertStyleResponse structure represents the ConvertStyle operation response
type ConvertStyleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConvertStyle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConvertStyleResponse) xxx_ToOp(ctx context.Context, op *xxx_ConvertStyleOperation) *xxx_ConvertStyleOperation {
	if op == nil {
		op = &xxx_ConvertStyleOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ConvertStyleResponse) xxx_FromOp(ctx context.Context, op *xxx_ConvertStyleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConvertStyleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConvertStyleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConvertStyleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFlagsOperation structure represents the SetFlags operation
type xxx_SetFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFlagsOperation) OpNum() int { return 8 }

func (o *xxx_SetFlagsOperation) OpName() string { return "/IVdsDisk/v0/SetFlags" }

func (o *xxx_SetFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFlagsRequest structure represents the SetFlags operation request
type SetFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags uint32         `idl:"name:ulFlags" json:"flags"`
}

func (o *SetFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFlagsOperation) *xxx_SetFlagsOperation {
	if op == nil {
		op = &xxx_SetFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Flags = op.Flags
	return op
}

func (o *SetFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *SetFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFlagsResponse structure represents the SetFlags operation response
type SetFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFlagsOperation) *xxx_SetFlagsOperation {
	if op == nil {
		op = &xxx_SetFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearFlagsOperation structure represents the ClearFlags operation
type xxx_ClearFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearFlagsOperation) OpNum() int { return 9 }

func (o *xxx_ClearFlagsOperation) OpName() string { return "/IVdsDisk/v0/ClearFlags" }

func (o *xxx_ClearFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearFlagsRequest structure represents the ClearFlags operation request
type ClearFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags uint32         `idl:"name:ulFlags" json:"flags"`
}

func (o *ClearFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearFlagsOperation) *xxx_ClearFlagsOperation {
	if op == nil {
		op = &xxx_ClearFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Flags = op.Flags
	return op
}

func (o *ClearFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *ClearFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearFlagsResponse structure represents the ClearFlags operation response
type ClearFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearFlagsOperation) *xxx_ClearFlagsOperation {
	if op == nil {
		op = &xxx_ClearFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ClearFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
