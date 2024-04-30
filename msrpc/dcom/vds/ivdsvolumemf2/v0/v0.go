package ivdsvolumemf2

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
	// IVdsVolumeMF2 interface identifier 4dbcee9a-6343-4651-b85f-5e75d74d983c
	VolumeMF2IID = &dcom.IID{Data1: 0x4dbcee9a, Data2: 0x6343, Data3: 0x4651, Data4: []byte{0xb8, 0x5f, 0x5e, 0x75, 0xd7, 0x4d, 0x98, 0x3c}}
	// Syntax UUID
	VolumeMF2SyntaxUUID = &uuid.UUID{TimeLow: 0x4dbcee9a, TimeMid: 0x6343, TimeHiAndVersion: 0x4651, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x5f, Node: [6]uint8{0x5e, 0x75, 0xd7, 0x4d, 0x98, 0x3c}}
	// Syntax ID
	VolumeMF2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeMF2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolumeMF2 interface.
type VolumeMF2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetFileSystemTypeName method retrieves the name of the file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetFileSystemTypeName(context.Context, *GetFileSystemTypeNameRequest, ...dcerpc.CallOption) (*GetFileSystemTypeNameResponse, error)

	// The QueryFileSystemFormatSupport method retrieves the properties of the file systems
	// that are supported for formatting a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryFileSystemFormatSupport(context.Context, *QueryFileSystemFormatSupportRequest, ...dcerpc.CallOption) (*QueryFileSystemFormatSupportResponse, error)

	// The FormatEx method formats a file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatEx(context.Context, *FormatExRequest, ...dcerpc.CallOption) (*FormatExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeMF2Client
}

type xxx_DefaultVolumeMF2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeMF2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeMF2Client) GetFileSystemTypeName(ctx context.Context, in *GetFileSystemTypeNameRequest, opts ...dcerpc.CallOption) (*GetFileSystemTypeNameResponse, error) {
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
	out := &GetFileSystemTypeNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF2Client) QueryFileSystemFormatSupport(ctx context.Context, in *QueryFileSystemFormatSupportRequest, opts ...dcerpc.CallOption) (*QueryFileSystemFormatSupportResponse, error) {
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
	out := &QueryFileSystemFormatSupportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF2Client) FormatEx(ctx context.Context, in *FormatExRequest, opts ...dcerpc.CallOption) (*FormatExResponse, error) {
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
	out := &FormatExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeMF2Client) IPID(ctx context.Context, ipid *dcom.IPID) VolumeMF2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeMF2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewVolumeMF2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeMF2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeMF2SyntaxV0_0))...)
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
	return &xxx_DefaultVolumeMF2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetFileSystemTypeNameOperation structure represents the GetFileSystemTypeName operation
type xxx_GetFileSystemTypeNameOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileSystemTypeName string         `idl:"name:ppwszFileSystemTypeName;string" json:"file_system_type_name"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileSystemTypeNameOperation) OpNum() int { return 3 }

func (o *xxx_GetFileSystemTypeNameOperation) OpName() string {
	return "/IVdsVolumeMF2/v0/GetFileSystemTypeName"
}

func (o *xxx_GetFileSystemTypeNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileSystemTypeNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileSystemTypeNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileSystemTypeNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileSystemTypeNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppwszFileSystemTypeName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.FileSystemTypeName != "" {
			_ptr_ppwszFileSystemTypeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileSystemTypeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemTypeName, _ptr_ppwszFileSystemTypeName); err != nil {
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

func (o *xxx_GetFileSystemTypeNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppwszFileSystemTypeName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ppwszFileSystemTypeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileSystemTypeName); err != nil {
				return err
			}
			return nil
		})
		_s_ppwszFileSystemTypeName := func(ptr interface{}) { o.FileSystemTypeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileSystemTypeName, _s_ppwszFileSystemTypeName, _ptr_ppwszFileSystemTypeName); err != nil {
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

// GetFileSystemTypeNameRequest structure represents the GetFileSystemTypeName operation request
type GetFileSystemTypeNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileSystemTypeNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetFileSystemTypeNameOperation {
	if o == nil {
		return &xxx_GetFileSystemTypeNameOperation{}
	}
	return &xxx_GetFileSystemTypeNameOperation{
		This: o.This,
	}
}

func (o *GetFileSystemTypeNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileSystemTypeNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileSystemTypeNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFileSystemTypeNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileSystemTypeNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileSystemTypeNameResponse structure represents the GetFileSystemTypeName operation response
type GetFileSystemTypeNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszFileSystemTypeName: A pointer that, if the operation is successfully completed,
	// receives a null-terminated Unicode string with the file system name.
	FileSystemTypeName string `idl:"name:ppwszFileSystemTypeName;string" json:"file_system_type_name"`
	// Return: The GetFileSystemTypeName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileSystemTypeNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetFileSystemTypeNameOperation {
	if o == nil {
		return &xxx_GetFileSystemTypeNameOperation{}
	}
	return &xxx_GetFileSystemTypeNameOperation{
		That:               o.That,
		FileSystemTypeName: o.FileSystemTypeName,
		Return:             o.Return,
	}
}

func (o *GetFileSystemTypeNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileSystemTypeNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemTypeName = op.FileSystemTypeName
	o.Return = op.Return
}
func (o *GetFileSystemTypeNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFileSystemTypeNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileSystemTypeNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryFileSystemFormatSupportOperation structure represents the QueryFileSystemFormatSupport operation
type xxx_QueryFileSystemFormatSupportOperation struct {
	This                        *dcom.ORPCThis                         `idl:"name:This" json:"this"`
	That                        *dcom.ORPCThat                         `idl:"name:That" json:"that"`
	FileSystemSupportProperties []*vds.FileSystemFormatSupportProperty `idl:"name:ppFileSystemSupportProps;size_is:(, plNumberOfFileSystems)" json:"file_system_support_properties"`
	NumberOfFileSystems         int32                                  `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	Return                      int32                                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryFileSystemFormatSupportOperation) OpNum() int { return 4 }

func (o *xxx_QueryFileSystemFormatSupportOperation) OpName() string {
	return "/IVdsVolumeMF2/v0/QueryFileSystemFormatSupport"
}

func (o *xxx_QueryFileSystemFormatSupportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemFormatSupportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryFileSystemFormatSupportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryFileSystemFormatSupportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FileSystemSupportProperties != nil && o.NumberOfFileSystems == 0 {
		o.NumberOfFileSystems = int32(len(o.FileSystemSupportProperties))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFileSystemFormatSupportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppFileSystemSupportProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		if o.FileSystemSupportProperties != nil || o.NumberOfFileSystems > 0 {
			_ptr_ppFileSystemSupportProps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfFileSystems)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FileSystemSupportProperties {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FileSystemSupportProperties[i1] != nil {
						if err := o.FileSystemSupportProperties[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.FileSystemFormatSupportProperty{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FileSystemSupportProperties); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.FileSystemFormatSupportProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemSupportProperties, _ptr_ppFileSystemSupportProps); err != nil {
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
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfFileSystems); err != nil {
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

func (o *xxx_QueryFileSystemFormatSupportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppFileSystemSupportProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		_ptr_ppFileSystemSupportProps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FileSystemSupportProperties", sizeInfo[0])
			}
			o.FileSystemSupportProperties = make([]*vds.FileSystemFormatSupportProperty, sizeInfo[0])
			for i1 := range o.FileSystemSupportProperties {
				i1 := i1
				if o.FileSystemSupportProperties[i1] == nil {
					o.FileSystemSupportProperties[i1] = &vds.FileSystemFormatSupportProperty{}
				}
				if err := o.FileSystemSupportProperties[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppFileSystemSupportProps := func(ptr interface{}) { o.FileSystemSupportProperties = *ptr.(*[]*vds.FileSystemFormatSupportProperty) }
		if err := w.ReadPointer(&o.FileSystemSupportProperties, _s_ppFileSystemSupportProps, _ptr_ppFileSystemSupportProps); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfFileSystems); err != nil {
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

// QueryFileSystemFormatSupportRequest structure represents the QueryFileSystemFormatSupport operation request
type QueryFileSystemFormatSupportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryFileSystemFormatSupportRequest) xxx_ToOp(ctx context.Context) *xxx_QueryFileSystemFormatSupportOperation {
	if o == nil {
		return &xxx_QueryFileSystemFormatSupportOperation{}
	}
	return &xxx_QueryFileSystemFormatSupportOperation{
		This: o.This,
	}
}

func (o *QueryFileSystemFormatSupportRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryFileSystemFormatSupportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryFileSystemFormatSupportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryFileSystemFormatSupportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileSystemFormatSupportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryFileSystemFormatSupportResponse structure represents the QueryFileSystemFormatSupport operation response
type QueryFileSystemFormatSupportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppFileSystemSupportProps: A pointer to an array of VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP
	// structures which, if the operation completes successfully, receives an array of properties
	// of each supported file-system.
	FileSystemSupportProperties []*vds.FileSystemFormatSupportProperty `idl:"name:ppFileSystemSupportProps;size_is:(, plNumberOfFileSystems)" json:"file_system_support_properties"`
	// plNumberOfFileSystems: A pointer to a variable which, if the operation completes
	// successfully, receives the total number of elements returned in ppFileSystemSupportProps.
	NumberOfFileSystems int32 `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	// Return: The QueryFileSystemFormatSupport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryFileSystemFormatSupportResponse) xxx_ToOp(ctx context.Context) *xxx_QueryFileSystemFormatSupportOperation {
	if o == nil {
		return &xxx_QueryFileSystemFormatSupportOperation{}
	}
	return &xxx_QueryFileSystemFormatSupportOperation{
		That:                        o.That,
		FileSystemSupportProperties: o.FileSystemSupportProperties,
		NumberOfFileSystems:         o.NumberOfFileSystems,
		Return:                      o.Return,
	}
}

func (o *QueryFileSystemFormatSupportResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryFileSystemFormatSupportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemSupportProperties = op.FileSystemSupportProperties
	o.NumberOfFileSystems = op.NumberOfFileSystems
	o.Return = op.Return
}
func (o *QueryFileSystemFormatSupportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryFileSystemFormatSupportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFileSystemFormatSupportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatExOperation structure represents the FormatEx operation
type xxx_FormatExOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileSystemTypeName        string         `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	FileSystemRevision        uint16         `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	DesiredUnitAllocationSize uint32         `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	Label                     string         `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	Force                     int32          `idl:"name:bForce" json:"force"`
	QuickFormat               int32          `idl:"name:bQuickFormat" json:"quick_format"`
	EnableCompression         int32          `idl:"name:bEnableCompression" json:"enable_compression"`
	Async                     *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatExOperation) OpNum() int { return 5 }

func (o *xxx_FormatExOperation) OpName() string { return "/IVdsVolumeMF2/v0/FormatEx" }

func (o *xxx_FormatExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.FileSystemTypeName != "" {
			_ptr_pwszFileSystemTypeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileSystemTypeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemTypeName, _ptr_pwszFileSystemTypeName); err != nil {
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
	// usFileSystemRevision {in} (1:(uint16))
	{
		if err := w.WriteData(o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Label != "" {
			_ptr_pwszLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Label); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_pwszLabel); err != nil {
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
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// bQuickFormat {in} (1:(int32))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	// bEnableCompression {in} (1:(int32))
	{
		if err := w.WriteData(o.EnableCompression); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszFileSystemTypeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileSystemTypeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszFileSystemTypeName := func(ptr interface{}) { o.FileSystemTypeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileSystemTypeName, _s_pwszFileSystemTypeName, _ptr_pwszFileSystemTypeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// usFileSystemRevision {in} (1:(uint16))
	{
		if err := w.ReadData(&o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Label); err != nil {
				return err
			}
			return nil
		})
		_s_pwszLabel := func(ptr interface{}) { o.Label = *ptr.(*string) }
		if err := w.ReadPointer(&o.Label, _s_pwszLabel, _ptr_pwszLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// bQuickFormat {in} (1:(int32))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	// bEnableCompression {in} (1:(int32))
	{
		if err := w.ReadData(&o.EnableCompression); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Async != nil {
					if err := o.Async.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Async{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Async, _ptr_ppAsync); err != nil {
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

func (o *xxx_FormatExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Async == nil {
				o.Async = &vds.Async{}
			}
			if err := o.Async.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAsync := func(ptr interface{}) { o.Async = *ptr.(**vds.Async) }
		if err := w.ReadPointer(&o.Async, _s_ppAsync, _ptr_ppAsync); err != nil {
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

// FormatExRequest structure represents the FormatEx operation request
type FormatExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszFileSystemTypeName: A null-terminated Unicode string that contains the name of
	// the file systems to format the volume with.
	FileSystemTypeName string `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	// usFileSystemRevision: A 16-bit, binary-coded decimal number that indicates the revision
	// of the file system, if any. The first two (most significant) digits (8-bits) indicate
	// the major revision, and the last two (least significant) digits (8-bits) indicate
	// the minor revision.
	FileSystemRevision uint16 `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	// ulDesiredUnitAllocationSize: The size, in bytes, of the allocation unit for the file
	// system. The value MUST be a power of 2. If the value is 0, a default allocation unit
	// that is determined by the file system type is used. The allocation unit range is
	// file system-dependent.
	DesiredUnitAllocationSize uint32 `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	// pwszLabel: A null-terminated Unicode string to assign to the new file system. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	// bForce: A Boolean that determines whether a file system format is forced, even if
	// the volume is in use.
	Force int32 `idl:"name:bForce" json:"force"`
	// bQuickFormat: A Boolean that determines whether a file system is quick formatted.
	// A quick format does not verify each sector on the volume.
	QuickFormat int32 `idl:"name:bQuickFormat" json:"quick_format"`
	// bEnableCompression: A Boolean that determines whether a file system is created with
	// compression enabled.<134>
	EnableCompression int32 `idl:"name:bEnableCompression" json:"enable_compression"`
}

func (o *FormatExRequest) xxx_ToOp(ctx context.Context) *xxx_FormatExOperation {
	if o == nil {
		return &xxx_FormatExOperation{}
	}
	return &xxx_FormatExOperation{
		This:                      o.This,
		FileSystemTypeName:        o.FileSystemTypeName,
		FileSystemRevision:        o.FileSystemRevision,
		DesiredUnitAllocationSize: o.DesiredUnitAllocationSize,
		Label:                     o.Label,
		Force:                     o.Force,
		QuickFormat:               o.QuickFormat,
		EnableCompression:         o.EnableCompression,
	}
}

func (o *FormatExRequest) xxx_FromOp(ctx context.Context, op *xxx_FormatExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileSystemTypeName = op.FileSystemTypeName
	o.FileSystemRevision = op.FileSystemRevision
	o.DesiredUnitAllocationSize = op.DesiredUnitAllocationSize
	o.Label = op.Label
	o.Force = op.Force
	o.QuickFormat = op.QuickFormat
	o.EnableCompression = op.EnableCompression
}
func (o *FormatExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FormatExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatExResponse structure represents the FormatEx operation response
type FormatExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The FormatEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatExResponse) xxx_ToOp(ctx context.Context) *xxx_FormatExOperation {
	if o == nil {
		return &xxx_FormatExOperation{}
	}
	return &xxx_FormatExOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *FormatExResponse) xxx_FromOp(ctx context.Context, op *xxx_FormatExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FormatExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
