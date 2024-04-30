package ivdsvolumemf3

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
	// IVdsVolumeMF3 interface identifier 6788faf9-214e-4b85-ba59-266953616e09
	VolumeMF3IID = &dcom.IID{Data1: 0x6788faf9, Data2: 0x214e, Data3: 0x4b85, Data4: []byte{0xba, 0x59, 0x26, 0x69, 0x53, 0x61, 0x6e, 0x09}}
	// Syntax UUID
	VolumeMF3SyntaxUUID = &uuid.UUID{TimeLow: 0x6788faf9, TimeMid: 0x214e, TimeHiAndVersion: 0x4b85, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x59, Node: [6]uint8{0x26, 0x69, 0x53, 0x61, 0x6e, 0x9}}
	// Syntax ID
	VolumeMF3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeMF3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolumeMF3 interface.
type VolumeMF3Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The QueryVolumeGuidPathnames method returns a volume's volume GUID path names.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVolumeGUIDPathnames(context.Context, *QueryVolumeGUIDPathnamesRequest, ...dcerpc.CallOption) (*QueryVolumeGUIDPathnamesResponse, error)

	// The FormatEx2 method formats a file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatEx2(context.Context, *FormatEx2Request, ...dcerpc.CallOption) (*FormatEx2Response, error)

	// The OfflineVolume method offlines a volume. An offline volume will fail data IO.
	// The volume can be opened for configuration.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	OfflineVolume(context.Context, *OfflineVolumeRequest, ...dcerpc.CallOption) (*OfflineVolumeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeMF3Client
}

type xxx_DefaultVolumeMF3Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeMF3Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeMF3Client) QueryVolumeGUIDPathnames(ctx context.Context, in *QueryVolumeGUIDPathnamesRequest, opts ...dcerpc.CallOption) (*QueryVolumeGUIDPathnamesResponse, error) {
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
	out := &QueryVolumeGUIDPathnamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF3Client) FormatEx2(ctx context.Context, in *FormatEx2Request, opts ...dcerpc.CallOption) (*FormatEx2Response, error) {
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
	out := &FormatEx2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF3Client) OfflineVolume(ctx context.Context, in *OfflineVolumeRequest, opts ...dcerpc.CallOption) (*OfflineVolumeResponse, error) {
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
	out := &OfflineVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMF3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeMF3Client) IPID(ctx context.Context, ipid *dcom.IPID) VolumeMF3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeMF3Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewVolumeMF3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeMF3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeMF3SyntaxV0_0))...)
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
	return &xxx_DefaultVolumeMF3Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryVolumeGUIDPathnamesOperation structure represents the QueryVolumeGuidPathnames operation
type xxx_QueryVolumeGUIDPathnamesOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathArray     []string       `idl:"name:pwszPathArray;size_is:(, pulNumberOfPaths);string" json:"path_array"`
	NumberOfPaths uint32         `idl:"name:pulNumberOfPaths" json:"number_of_paths"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryVolumeGUIDPathnamesOperation) OpNum() int { return 3 }

func (o *xxx_QueryVolumeGUIDPathnamesOperation) OpName() string {
	return "/IVdsVolumeMF3/v0/QueryVolumeGuidPathnames"
}

func (o *xxx_QueryVolumeGUIDPathnamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumeGUIDPathnamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryVolumeGUIDPathnamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryVolumeGUIDPathnamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PathArray != nil && o.NumberOfPaths == 0 {
		o.NumberOfPaths = uint32(len(o.PathArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumeGUIDPathnamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pwszPathArray {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pulNumberOfPaths]*(1)[dim:0,string,null](wchar))
	{
		if o.PathArray != nil || o.NumberOfPaths > 0 {
			_ptr_pwszPathArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfPaths)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PathArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PathArray[i1] != "" {
						_ptr_pwszPathArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.PathArray[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.PathArray[i1], _ptr_pwszPathArray); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PathArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PathArray, _ptr_pwszPathArray); err != nil {
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
	// pulNumberOfPaths {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.NumberOfPaths); err != nil {
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

func (o *xxx_QueryVolumeGUIDPathnamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pwszPathArray {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pulNumberOfPaths]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPathArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PathArray", sizeInfo[0])
			}
			o.PathArray = make([]string, sizeInfo[0])
			for i1 := range o.PathArray {
				i1 := i1
				_ptr_pwszPathArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.PathArray[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pwszPathArray := func(ptr interface{}) { o.PathArray[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.PathArray[i1], _s_pwszPathArray, _ptr_pwszPathArray); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pwszPathArray := func(ptr interface{}) { o.PathArray = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.PathArray, _s_pwszPathArray, _ptr_pwszPathArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pulNumberOfPaths {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.NumberOfPaths); err != nil {
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

// QueryVolumeGUIDPathnamesRequest structure represents the QueryVolumeGuidPathnames operation request
type QueryVolumeGUIDPathnamesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryVolumeGUIDPathnamesRequest) xxx_ToOp(ctx context.Context) *xxx_QueryVolumeGUIDPathnamesOperation {
	if o == nil {
		return &xxx_QueryVolumeGUIDPathnamesOperation{}
	}
	return &xxx_QueryVolumeGUIDPathnamesOperation{
		This: o.This,
	}
}

func (o *QueryVolumeGUIDPathnamesRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumeGUIDPathnamesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryVolumeGUIDPathnamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryVolumeGUIDPathnamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumeGUIDPathnamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryVolumeGUIDPathnamesResponse structure represents the QueryVolumeGuidPathnames operation response
type QueryVolumeGUIDPathnamesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pwszPathArray: Returns an array of null-terminated Unicode strings; one string for
	// each volume GUID path name associated with the volume.
	PathArray []string `idl:"name:pwszPathArray;size_is:(, pulNumberOfPaths);string" json:"path_array"`
	// pulNumberOfPaths: Returns the number of volume GUID path names returned in pwszPathArray.
	NumberOfPaths uint32 `idl:"name:pulNumberOfPaths" json:"number_of_paths"`
	// Return: The QueryVolumeGuidPathnames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryVolumeGUIDPathnamesResponse) xxx_ToOp(ctx context.Context) *xxx_QueryVolumeGUIDPathnamesOperation {
	if o == nil {
		return &xxx_QueryVolumeGUIDPathnamesOperation{}
	}
	return &xxx_QueryVolumeGUIDPathnamesOperation{
		That:          o.That,
		PathArray:     o.PathArray,
		NumberOfPaths: o.NumberOfPaths,
		Return:        o.Return,
	}
}

func (o *QueryVolumeGUIDPathnamesResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumeGUIDPathnamesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PathArray = op.PathArray
	o.NumberOfPaths = op.NumberOfPaths
	o.Return = op.Return
}
func (o *QueryVolumeGUIDPathnamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryVolumeGUIDPathnamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumeGUIDPathnamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatEx2Operation structure represents the FormatEx2 operation
type xxx_FormatEx2Operation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileSystemTypeName        string         `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	FileSystemRevision        uint16         `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	DesiredUnitAllocationSize uint32         `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	Label                     string         `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	Options                   uint32         `idl:"name:Options" json:"options"`
	Async                     *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatEx2Operation) OpNum() int { return 4 }

func (o *xxx_FormatEx2Operation) OpName() string { return "/IVdsVolumeMF3/v0/FormatEx2" }

func (o *xxx_FormatEx2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatEx2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// usFileSystemRevision {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.WriteData(o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// Options {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatEx2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// usFileSystemRevision {in} (1:{alias=USHORT}(uint16))
	{
		if err := w.ReadData(&o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// Options {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatEx2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatEx2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatEx2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FormatEx2Request structure represents the FormatEx2 operation request
type FormatEx2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszFileSystemTypeName: A null-terminated Unicode string that contains the name of
	// the file systems to format the volume with.
	FileSystemTypeName string `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	// usFileSystemRevision: A 16-bit, binary-coded decimal number that indicates the revision
	// of the file system, if any. The first two (most significant) digits (8 bits) indicate
	// the major revision, and the last two (least significant) digits (8 bits) indicate
	// the minor revision. For example, 0x0250 represents revision 2.50.
	FileSystemRevision uint16 `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	// ulDesiredUnitAllocationSize: The size, in bytes, of the allocation unit for the file
	// system. The value MUST be a power of 2. If the value is 0, a default allocation unit
	// that is determined by the file system type is used. The allocation unit range is
	// file system-dependent.
	DesiredUnitAllocationSize uint32 `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	// pwszLabel: A null-terminated Unicode string to assign to the new file system. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	// Options: The combination of any values, by using a bitwise OR operator, that are
	// defined in the VDS_FORMAT_OPTION_FLAGS enumeration.
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *FormatEx2Request) xxx_ToOp(ctx context.Context) *xxx_FormatEx2Operation {
	if o == nil {
		return &xxx_FormatEx2Operation{}
	}
	return &xxx_FormatEx2Operation{
		This:                      o.This,
		FileSystemTypeName:        o.FileSystemTypeName,
		FileSystemRevision:        o.FileSystemRevision,
		DesiredUnitAllocationSize: o.DesiredUnitAllocationSize,
		Label:                     o.Label,
		Options:                   o.Options,
	}
}

func (o *FormatEx2Request) xxx_FromOp(ctx context.Context, op *xxx_FormatEx2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileSystemTypeName = op.FileSystemTypeName
	o.FileSystemRevision = op.FileSystemRevision
	o.DesiredUnitAllocationSize = op.DesiredUnitAllocationSize
	o.Label = op.Label
	o.Options = op.Options
}
func (o *FormatEx2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FormatEx2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatEx2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatEx2Response structure represents the FormatEx2 operation response
type FormatEx2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The FormatEx2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatEx2Response) xxx_ToOp(ctx context.Context) *xxx_FormatEx2Operation {
	if o == nil {
		return &xxx_FormatEx2Operation{}
	}
	return &xxx_FormatEx2Operation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *FormatEx2Response) xxx_FromOp(ctx context.Context, op *xxx_FormatEx2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatEx2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FormatEx2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatEx2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OfflineVolumeOperation structure represents the OfflineVolume operation
type xxx_OfflineVolumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OfflineVolumeOperation) OpNum() int { return 5 }

func (o *xxx_OfflineVolumeOperation) OpName() string { return "/IVdsVolumeMF3/v0/OfflineVolume" }

func (o *xxx_OfflineVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OfflineVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_OfflineVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OfflineVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OfflineVolumeRequest structure represents the OfflineVolume operation request
type OfflineVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *OfflineVolumeRequest) xxx_ToOp(ctx context.Context) *xxx_OfflineVolumeOperation {
	if o == nil {
		return &xxx_OfflineVolumeOperation{}
	}
	return &xxx_OfflineVolumeOperation{
		This: o.This,
	}
}

func (o *OfflineVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_OfflineVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *OfflineVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OfflineVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OfflineVolumeResponse structure represents the OfflineVolume operation response
type OfflineVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OfflineVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OfflineVolumeResponse) xxx_ToOp(ctx context.Context) *xxx_OfflineVolumeOperation {
	if o == nil {
		return &xxx_OfflineVolumeOperation{}
	}
	return &xxx_OfflineVolumeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *OfflineVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_OfflineVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OfflineVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OfflineVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
