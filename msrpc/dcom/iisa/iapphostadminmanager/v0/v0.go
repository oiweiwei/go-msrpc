package iapphostadminmanager

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
	// IAppHostAdminManager interface identifier 9be77978-73ed-4a9a-87fd-13f09fec1b13
	AppHostAdminManagerIID = &dcom.IID{Data1: 0x9be77978, Data2: 0x73ed, Data3: 0x4a9a, Data4: []byte{0x87, 0xfd, 0x13, 0xf0, 0x9f, 0xec, 0x1b, 0x13}}
	// Syntax UUID
	AppHostAdminManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x9be77978, TimeMid: 0x73ed, TimeHiAndVersion: 0x4a9a, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0xfd, Node: [6]uint8{0x13, 0xf0, 0x9f, 0xec, 0x1b, 0x13}}
	// Syntax ID
	AppHostAdminManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostAdminManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostAdminManager interface.
type AppHostAdminManagerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest, ...dcerpc.CallOption) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest, ...dcerpc.CallOption) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest, ...dcerpc.CallOption) (*SetMetadataResponse, error)

	// ConfigManager operation.
	GetConfigManager(context.Context, *GetConfigManagerRequest, ...dcerpc.CallOption) (*GetConfigManagerResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostAdminManagerClient
}

type xxx_DefaultAppHostAdminManagerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostAdminManagerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostAdminManagerClient) GetAdminSection(ctx context.Context, in *GetAdminSectionRequest, opts ...dcerpc.CallOption) (*GetAdminSectionResponse, error) {
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
	out := &GetAdminSectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostAdminManagerClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...dcerpc.CallOption) (*GetMetadataResponse, error) {
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
	out := &GetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostAdminManagerClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...dcerpc.CallOption) (*SetMetadataResponse, error) {
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
	out := &SetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostAdminManagerClient) GetConfigManager(ctx context.Context, in *GetConfigManagerRequest, opts ...dcerpc.CallOption) (*GetConfigManagerResponse, error) {
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
	out := &GetConfigManagerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostAdminManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostAdminManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostAdminManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostAdminManagerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewAppHostAdminManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostAdminManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostAdminManagerSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostAdminManagerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetAdminSectionOperation structure represents the GetAdminSection operation
type xxx_GetAdminSectionOperation struct {
	This         *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SectionName  *oaut.String         `idl:"name:bstrSectionName" json:"section_name"`
	Path         *oaut.String         `idl:"name:bstrPath" json:"path"`
	AdminSection *iisa.AppHostElement `idl:"name:ppAdminSection" json:"admin_section"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminSectionOperation) OpNum() int { return 3 }

func (o *xxx_GetAdminSectionOperation) OpName() string {
	return "/IAppHostAdminManager/v0/GetAdminSection"
}

func (o *xxx_GetAdminSectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_bstrPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_bstrPath); err != nil {
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

func (o *xxx_GetAdminSectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPath := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_bstrPath, _ptr_bstrPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAdminSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.AdminSection != nil {
			_ptr_ppAdminSection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AdminSection != nil {
					if err := o.AdminSection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminSection, _ptr_ppAdminSection); err != nil {
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

func (o *xxx_GetAdminSectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAdminSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppAdminSection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AdminSection == nil {
				o.AdminSection = &iisa.AppHostElement{}
			}
			if err := o.AdminSection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAdminSection := func(ptr interface{}) { o.AdminSection = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.AdminSection, _s_ppAdminSection, _ptr_ppAdminSection); err != nil {
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

// GetAdminSectionRequest structure represents the GetAdminSection operation request
type GetAdminSectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	SectionName *oaut.String   `idl:"name:bstrSectionName" json:"section_name"`
	Path        *oaut.String   `idl:"name:bstrPath" json:"path"`
}

func (o *GetAdminSectionRequest) xxx_ToOp(ctx context.Context) *xxx_GetAdminSectionOperation {
	if o == nil {
		return &xxx_GetAdminSectionOperation{}
	}
	return &xxx_GetAdminSectionOperation{
		This:        o.This,
		SectionName: o.SectionName,
		Path:        o.Path,
	}
}

func (o *GetAdminSectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAdminSectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SectionName = op.SectionName
	o.Path = op.Path
}
func (o *GetAdminSectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminSectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminSectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminSectionResponse structure represents the GetAdminSection operation response
type GetAdminSectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	AdminSection *iisa.AppHostElement `idl:"name:ppAdminSection" json:"admin_section"`
	// Return: The GetAdminSection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminSectionResponse) xxx_ToOp(ctx context.Context) *xxx_GetAdminSectionOperation {
	if o == nil {
		return &xxx_GetAdminSectionOperation{}
	}
	return &xxx_GetAdminSectionOperation{
		That:         o.That,
		AdminSection: o.AdminSection,
		Return:       o.Return,
	}
}

func (o *GetAdminSectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAdminSectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AdminSection = op.AdminSection
	o.Return = op.Return
}
func (o *GetAdminSectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminSectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminSectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMetadataOperation structure represents the GetMetadata operation
type xxx_GetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:pValue" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMetadataOperation) OpNum() int { return 4 }

func (o *xxx_GetMetadataOperation) OpName() string { return "/IAppHostAdminManager/v0/GetMetadata" }

func (o *xxx_GetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			_ptr_pValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Value, _ptr_pValue); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Value, _s_pValue, _ptr_pValue); err != nil {
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

// GetMetadataRequest structure represents the GetMetadata operation request
type GetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
}

func (o *GetMetadataRequest) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		This:         o.This,
		MetadataType: o.MetadataType,
	}
}

func (o *GetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
}
func (o *GetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMetadataResponse structure represents the GetMetadata operation response
type GetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.Variant  `idl:"name:pValue" json:"value"`
	// Return: The GetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMetadataResponse) xxx_ToOp(ctx context.Context) *xxx_GetMetadataOperation {
	if o == nil {
		return &xxx_GetMetadataOperation{}
	}
	return &xxx_GetMetadataOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *GetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMetadataOperation structure represents the SetMetadata operation
type xxx_SetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:value" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMetadataOperation) OpNum() int { return 5 }

func (o *xxx_SetMetadataOperation) OpName() string { return "/IAppHostAdminManager/v0/SetMetadata" }

func (o *xxx_SetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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
	// value {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			if err := o.Value.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Value == nil {
			o.Value = &oaut.Variant{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMetadataRequest structure represents the SetMetadata operation request
type SetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:value" json:"value"`
}

func (o *SetMetadataRequest) xxx_ToOp(ctx context.Context) *xxx_SetMetadataOperation {
	if o == nil {
		return &xxx_SetMetadataOperation{}
	}
	return &xxx_SetMetadataOperation{
		This:         o.This,
		MetadataType: o.MetadataType,
		Value:        o.Value,
	}
}

func (o *SetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
	o.Value = op.Value
}
func (o *SetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMetadataResponse structure represents the SetMetadata operation response
type SetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMetadataResponse) xxx_ToOp(ctx context.Context) *xxx_SetMetadataOperation {
	if o == nil {
		return &xxx_SetMetadataOperation{}
	}
	return &xxx_SetMetadataOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConfigManagerOperation structure represents the ConfigManager operation
type xxx_GetConfigManagerOperation struct {
	This          *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat             `idl:"name:That" json:"that"`
	ConfigManager *iisa.AppHostConfigManager `idl:"name:ppConfigManager" json:"config_manager"`
	Return        int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigManagerOperation) OpNum() int { return 6 }

func (o *xxx_GetConfigManagerOperation) OpName() string {
	return "/IAppHostAdminManager/v0/ConfigManager"
}

func (o *xxx_GetConfigManagerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigManagerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetConfigManagerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetConfigManagerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigManagerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppConfigManager {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigManager}(interface))
	{
		if o.ConfigManager != nil {
			_ptr_ppConfigManager := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigManager != nil {
					if err := o.ConfigManager.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostConfigManager{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigManager, _ptr_ppConfigManager); err != nil {
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

func (o *xxx_GetConfigManagerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppConfigManager {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigManager}(interface))
	{
		_ptr_ppConfigManager := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigManager == nil {
				o.ConfigManager = &iisa.AppHostConfigManager{}
			}
			if err := o.ConfigManager.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConfigManager := func(ptr interface{}) { o.ConfigManager = *ptr.(**iisa.AppHostConfigManager) }
		if err := w.ReadPointer(&o.ConfigManager, _s_ppConfigManager, _ptr_ppConfigManager); err != nil {
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

// GetConfigManagerRequest structure represents the ConfigManager operation request
type GetConfigManagerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetConfigManagerRequest) xxx_ToOp(ctx context.Context) *xxx_GetConfigManagerOperation {
	if o == nil {
		return &xxx_GetConfigManagerOperation{}
	}
	return &xxx_GetConfigManagerOperation{
		This: o.This,
	}
}

func (o *GetConfigManagerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigManagerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetConfigManagerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigManagerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigManagerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigManagerResponse structure represents the ConfigManager operation response
type GetConfigManagerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat             `idl:"name:That" json:"that"`
	ConfigManager *iisa.AppHostConfigManager `idl:"name:ppConfigManager" json:"config_manager"`
	// Return: The ConfigManager return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigManagerResponse) xxx_ToOp(ctx context.Context) *xxx_GetConfigManagerOperation {
	if o == nil {
		return &xxx_GetConfigManagerOperation{}
	}
	return &xxx_GetConfigManagerOperation{
		That:          o.That,
		ConfigManager: o.ConfigManager,
		Return:        o.Return,
	}
}

func (o *GetConfigManagerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigManagerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConfigManager = op.ConfigManager
	o.Return = op.Return
}
func (o *GetConfigManagerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigManagerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigManagerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
