package ivdsservicesw

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
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsServiceSw interface identifier 15fc031c-0652-4306-b2c3-f558b8f837e2
	ServiceSwIID = &dcom.IID{Data1: 0x15fc031c, Data2: 0x0652, Data3: 0x4306, Data4: []byte{0xb2, 0xc3, 0xf5, 0x58, 0xb8, 0xf8, 0x37, 0xe2}}
	// Syntax UUID
	ServiceSwSyntaxUUID = &uuid.UUID{TimeLow: 0x15fc031c, TimeMid: 0x652, TimeHiAndVersion: 0x4306, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xc3, Node: [6]uint8{0xf5, 0x58, 0xb8, 0xf8, 0x37, 0xe2}}
	// Syntax ID
	ServiceSwSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceSwSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceSw interface.
type ServiceSwClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetDiskObject method<75> returns the disk for the given PnP Device ID string.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskObject(context.Context, *GetDiskObjectRequest, ...dcerpc.CallOption) (*GetDiskObjectResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceSwClient
}

type xxx_DefaultServiceSwClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceSwClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceSwClient) GetDiskObject(ctx context.Context, in *GetDiskObjectRequest, opts ...dcerpc.CallOption) (*GetDiskObjectResponse, error) {
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
	out := &GetDiskObjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceSwClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceSwClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceSwClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceSwClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewServiceSwClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceSwClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceSwSyntaxV0_0))...)
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
	return &xxx_DefaultServiceSwClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetDiskObjectOperation structure represents the GetDiskObject operation
type xxx_GetDiskObjectOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DeviceID    string         `idl:"name:pwszDeviceID;string" json:"device_id"`
	DiskUnknown *dcom.Unknown  `idl:"name:ppDiskUnk" json:"disk_unknown"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDiskObjectOperation) OpNum() int { return 3 }

func (o *xxx_GetDiskObjectOperation) OpName() string { return "/IVdsServiceSw/v0/GetDiskObject" }

func (o *xxx_GetDiskObjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskObjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszDeviceID {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DeviceID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskObjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszDeviceID {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskObjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskObjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppDiskUnk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.DiskUnknown != nil {
			_ptr_ppDiskUnk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DiskUnknown != nil {
					if err := o.DiskUnknown.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DiskUnknown, _ptr_ppDiskUnk); err != nil {
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

func (o *xxx_GetDiskObjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppDiskUnk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppDiskUnk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DiskUnknown == nil {
				o.DiskUnknown = &dcom.Unknown{}
			}
			if err := o.DiskUnknown.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDiskUnk := func(ptr interface{}) { o.DiskUnknown = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.DiskUnknown, _s_ppDiskUnk, _ptr_ppDiskUnk); err != nil {
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

// GetDiskObjectRequest structure represents the GetDiskObject operation request
type GetDiskObjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszDeviceID: The null-terminated Unicode device path that the operating system uses
	// to identify the device for the disk.
	DeviceID string `idl:"name:pwszDeviceID;string" json:"device_id"`
}

func (o *GetDiskObjectRequest) xxx_ToOp(ctx context.Context) *xxx_GetDiskObjectOperation {
	if o == nil {
		return &xxx_GetDiskObjectOperation{}
	}
	return &xxx_GetDiskObjectOperation{
		This:     o.This,
		DeviceID: o.DeviceID,
	}
}

func (o *GetDiskObjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDiskObjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DeviceID = op.DeviceID
}
func (o *GetDiskObjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDiskObjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskObjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDiskObjectResponse structure represents the GetDiskObject operation response
type GetDiskObjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskUnknown *dcom.Unknown  `idl:"name:ppDiskUnk" json:"disk_unknown"`
	// Return: The GetDiskObject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDiskObjectResponse) xxx_ToOp(ctx context.Context) *xxx_GetDiskObjectOperation {
	if o == nil {
		return &xxx_GetDiskObjectOperation{}
	}
	return &xxx_GetDiskObjectOperation{
		That:        o.That,
		DiskUnknown: o.DiskUnknown,
		Return:      o.Return,
	}
}

func (o *GetDiskObjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDiskObjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskUnknown = op.DiskUnknown
	o.Return = op.Return
}
func (o *GetDiskObjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDiskObjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskObjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
