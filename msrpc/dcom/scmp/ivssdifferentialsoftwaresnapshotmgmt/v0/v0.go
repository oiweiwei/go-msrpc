package ivssdifferentialsoftwaresnapshotmgmt

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
	scmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp"
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
	_ = scmp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/scmp"
)

var (
	// IVssDifferentialSoftwareSnapshotMgmt interface identifier 214a0f28-b737-4026-b847-4f9e37d79529
	DifferentialSoftwareSnapshotManagementIID = &dcom.IID{Data1: 0x214a0f28, Data2: 0xb737, Data3: 0x4026, Data4: []byte{0xb8, 0x47, 0x4f, 0x9e, 0x37, 0xd7, 0x95, 0x29}}
	// Syntax UUID
	DifferentialSoftwareSnapshotManagementSyntaxUUID = &uuid.UUID{TimeLow: 0x214a0f28, TimeMid: 0xb737, TimeHiAndVersion: 0x4026, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x47, Node: [6]uint8{0x4f, 0x9e, 0x37, 0xd7, 0x95, 0x29}}
	// Syntax ID
	DifferentialSoftwareSnapshotManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DifferentialSoftwareSnapshotManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVssDifferentialSoftwareSnapshotMgmt interface.
type DifferentialSoftwareSnapshotManagementClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The AddDiffArea method creates a shadow copy storage association for a shadow copy.
	//
	// Return Values: The method MUST return the following error code for the specific conditions.
	//
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                          |                                                                                  |
	//	|                       VALUE/CODE                        |                                   DESCRIPTION                                    |
	//	|                                                         |                                                                                  |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230d VSS_E_OBJECT_ALREADY_EXISTS                  | The object already exists on the server.                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                 | R returned when pwszVolumeName or pwszDiffAreaVolumeName is NULL, or if          |
	//	|                                                         | llMaximumDiffSpace is 0.                                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230c VSS_E_VOLUME_NOT_SUPPORTED                   | Returned when the pwszVolumeName does not support shadow copies, or              |
	//	|                                                         | pwszDiffAreaVolumeName does not support shadow copy storage.                     |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231e VSS_E_MAXIMUM_DIFF-AREA_ASSOCIATIONS_REACHED | Returned when the maximum number of diff area associations for pwszVolumeName    |
	//	|                                                         | has been reached.                                                                |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042306 VSS_E_PROVIDER_VETO                          | Returned when the snapshot provider receives an expected error and tries to veto |
	//	|                                                         | the impending operation.                                                         |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                               | Returned when the user making the request does not have sufficient privileges to |
	//	|                                                         | perform the operation.                                                           |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// specified in [MS-RPCE].
	AddDiffArea(context.Context, *AddDiffAreaRequest, ...dcerpc.CallOption) (*AddDiffAreaResponse, error)

	// The ChangeDiffAreaMaximumSize method changes the maximum size of a shadow copy storage
	// association on the server.
	//
	// Return Values: The method MUST return the following error code for the specific conditions.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042308 VSS_E_OBJECT_NOT_FOUND     | The object does not exist on the server.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG               | Returned when pwszVolumeName or pwszDiffAreaVolume is NULL.                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231d VSS_E_VOLUME_IN_USE        | Returned when llMaximumDiffSpace is zero, and the diff area cannot be deleted    |
	//	|                                       | because shadow copies are still being stored.                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004231f VSS_E_INSUFFICIENT_STORAGE | Returned if a nonzero size is specified in llMaximumDiffSpace that is smaller    |
	//	|                                       | than the size required for storing a single shadow copy.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED             | Returned when the user making the request does not have sufficient privileges to |
	//	|                                       | perform the operation.                                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	ChangeDiffAreaMaximumSize(context.Context, *ChangeDiffAreaMaximumSizeRequest, ...dcerpc.CallOption) (*ChangeDiffAreaMaximumSizeResponse, error)

	// The QueryVolumesSupportedForDiffAreas method retrieves from the server the collection
	// of volumes that can be used as a shadow copy storage volume for a specified original
	// volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszOriginalVolumeName or ppEnum is NULL.                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryVolumesSupportedForDiffAreas(context.Context, *QueryVolumesSupportedForDiffAreasRequest, ...dcerpc.CallOption) (*QueryVolumesSupportedForDiffAreasResponse, error)

	// The QueryDiffAreasForVolume method retrieves from the server the collection of shadow
	// copy storage associations that are being used for shadow copy storage for a specified
	// original volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszVolumeName or ppEnum is NULL.                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryDiffAreasForVolume(context.Context, *QueryDiffAreasForVolumeRequest, ...dcerpc.CallOption) (*QueryDiffAreasForVolumeResponse, error)

	// The QueryDiffAreasOnVolume method retrieves from the server the collection of shadow
	// copy storage associations that are located on a specified volume.
	//
	// Return Values: The method MUST return zero when it has succeeded or an implementation-specific
	// nonzero error code on failure.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN           |                                                                                  |
	//	|        VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | Returned when pwszVolumeName or ppEnum is NULL.                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | Returned when the user making the request does not have sufficient privileges to |
	//	|                           | perform the operation.                                                           |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol
	// [MS-RPCE].
	QueryDiffAreasOnVolume(context.Context, *QueryDiffAreasOnVolumeRequest, ...dcerpc.CallOption) (*QueryDiffAreasOnVolumeResponse, error)

	// Opnum08NotUsedOnWire operation.
	// Opnum08NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DifferentialSoftwareSnapshotManagementClient
}

type xxx_DefaultDifferentialSoftwareSnapshotManagementClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) AddDiffArea(ctx context.Context, in *AddDiffAreaRequest, opts ...dcerpc.CallOption) (*AddDiffAreaResponse, error) {
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
	out := &AddDiffAreaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) ChangeDiffAreaMaximumSize(ctx context.Context, in *ChangeDiffAreaMaximumSizeRequest, opts ...dcerpc.CallOption) (*ChangeDiffAreaMaximumSizeResponse, error) {
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
	out := &ChangeDiffAreaMaximumSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) QueryVolumesSupportedForDiffAreas(ctx context.Context, in *QueryVolumesSupportedForDiffAreasRequest, opts ...dcerpc.CallOption) (*QueryVolumesSupportedForDiffAreasResponse, error) {
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
	out := &QueryVolumesSupportedForDiffAreasResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) QueryDiffAreasForVolume(ctx context.Context, in *QueryDiffAreasForVolumeRequest, opts ...dcerpc.CallOption) (*QueryDiffAreasForVolumeResponse, error) {
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
	out := &QueryDiffAreasForVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) QueryDiffAreasOnVolume(ctx context.Context, in *QueryDiffAreasOnVolumeRequest, opts ...dcerpc.CallOption) (*QueryDiffAreasOnVolumeResponse, error) {
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
	out := &QueryDiffAreasOnVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDifferentialSoftwareSnapshotManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) DifferentialSoftwareSnapshotManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDifferentialSoftwareSnapshotManagementClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDifferentialSoftwareSnapshotManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DifferentialSoftwareSnapshotManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DifferentialSoftwareSnapshotManagementSyntaxV0_0))...)
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
	return &xxx_DefaultDifferentialSoftwareSnapshotManagementClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_AddDiffAreaOperation structure represents the AddDiffArea operation
type xxx_AddDiffAreaOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeName         string         `idl:"name:pwszVolumeName" json:"volume_name"`
	DiffAreaVolumeName string         `idl:"name:pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	MaximumDiffSpace   int64          `idl:"name:llMaximumDiffSpace" json:"maximum_diff_space"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddDiffAreaOperation) OpNum() int { return 3 }

func (o *xxx_AddDiffAreaOperation) OpName() string {
	return "/IVssDifferentialSoftwareSnapshotMgmt/v0/AddDiffArea"
}

func (o *xxx_AddDiffAreaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiffAreaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.VolumeName != "" {
			_ptr_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeName, _ptr_pwszVolumeName); err != nil {
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
	// pwszDiffAreaVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.DiffAreaVolumeName != "" {
			_ptr_pwszDiffAreaVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DiffAreaVolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DiffAreaVolumeName, _ptr_pwszDiffAreaVolumeName); err != nil {
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
	// llMaximumDiffSpace {in} (1:{alias=LONGLONG}(int64))
	{
		if err := w.WriteData(o.MaximumDiffSpace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiffAreaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeName, _s_pwszVolumeName, _ptr_pwszVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszDiffAreaVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszDiffAreaVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DiffAreaVolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszDiffAreaVolumeName := func(ptr interface{}) { o.DiffAreaVolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DiffAreaVolumeName, _s_pwszDiffAreaVolumeName, _ptr_pwszDiffAreaVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// llMaximumDiffSpace {in} (1:{alias=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.MaximumDiffSpace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiffAreaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiffAreaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddDiffAreaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddDiffAreaRequest structure represents the AddDiffArea operation request
type AddDiffAreaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszVolumeName:  A null-terminated UNICODE string that contains the drive letter,
	// mount point, or volume mount name of the volume for which the shadow copy is made.
	// This is the original volume.
	VolumeName string `idl:"name:pwszVolumeName" json:"volume_name"`
	// pwszDiffAreaVolumeName: A null-terminated UNICODE string that contains the drive
	// letter, mount point, or volume mount name of the volume on which the shadow copy
	// storage is located for the volume that is specified in pwszVolumeName. This is the
	// shadow copy storage volume.
	DiffAreaVolumeName string `idl:"name:pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	// llMaximumDiffSpace: The maximum number of BYTEs that the shadow copy storage will
	// occupy. The server MAY automatically delete shadow copies based on an implementation-specific
	// algorithm that reclaims space for newer shadow copies.
	MaximumDiffSpace int64 `idl:"name:llMaximumDiffSpace" json:"maximum_diff_space"`
}

func (o *AddDiffAreaRequest) xxx_ToOp(ctx context.Context, op *xxx_AddDiffAreaOperation) *xxx_AddDiffAreaOperation {
	if op == nil {
		op = &xxx_AddDiffAreaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VolumeName = o.VolumeName
	op.DiffAreaVolumeName = o.DiffAreaVolumeName
	op.MaximumDiffSpace = o.MaximumDiffSpace
	return op
}

func (o *AddDiffAreaRequest) xxx_FromOp(ctx context.Context, op *xxx_AddDiffAreaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeName = op.VolumeName
	o.DiffAreaVolumeName = op.DiffAreaVolumeName
	o.MaximumDiffSpace = op.MaximumDiffSpace
}
func (o *AddDiffAreaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddDiffAreaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddDiffAreaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddDiffAreaResponse structure represents the AddDiffArea operation response
type AddDiffAreaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddDiffArea return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddDiffAreaResponse) xxx_ToOp(ctx context.Context, op *xxx_AddDiffAreaOperation) *xxx_AddDiffAreaOperation {
	if op == nil {
		op = &xxx_AddDiffAreaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddDiffAreaResponse) xxx_FromOp(ctx context.Context, op *xxx_AddDiffAreaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddDiffAreaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddDiffAreaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddDiffAreaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangeDiffAreaMaximumSizeOperation structure represents the ChangeDiffAreaMaximumSize operation
type xxx_ChangeDiffAreaMaximumSizeOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeName         string         `idl:"name:pwszVolumeName" json:"volume_name"`
	DiffAreaVolumeName string         `idl:"name:pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	MaximumDiffSpace   int64          `idl:"name:llMaximumDiffSpace" json:"maximum_diff_space"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) OpNum() int { return 4 }

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) OpName() string {
	return "/IVssDifferentialSoftwareSnapshotMgmt/v0/ChangeDiffAreaMaximumSize"
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.VolumeName != "" {
			_ptr_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeName, _ptr_pwszVolumeName); err != nil {
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
	// pwszDiffAreaVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.DiffAreaVolumeName != "" {
			_ptr_pwszDiffAreaVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DiffAreaVolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DiffAreaVolumeName, _ptr_pwszDiffAreaVolumeName); err != nil {
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
	// llMaximumDiffSpace {in} (1:{alias=LONGLONG}(int64))
	{
		if err := w.WriteData(o.MaximumDiffSpace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeName, _s_pwszVolumeName, _ptr_pwszVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszDiffAreaVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszDiffAreaVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DiffAreaVolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszDiffAreaVolumeName := func(ptr interface{}) { o.DiffAreaVolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DiffAreaVolumeName, _s_pwszDiffAreaVolumeName, _ptr_pwszDiffAreaVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// llMaximumDiffSpace {in} (1:{alias=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.MaximumDiffSpace); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ChangeDiffAreaMaximumSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ChangeDiffAreaMaximumSizeRequest structure represents the ChangeDiffAreaMaximumSize operation request
type ChangeDiffAreaMaximumSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszVolumeName: A null-terminated UNICODE string that contains the drive letter,
	// mount point, or volume mount name of the volume for which the shadow copy is made.
	// This is the original volume.
	VolumeName string `idl:"name:pwszVolumeName" json:"volume_name"`
	// pwszDiffAreaVolumeName: A null-terminated UNICODE string that contains the drive
	// letter, mount point, or volume mount name of the volume on which the shadow copy
	// storage is located for the volume specified in pwszVolumeName. This is the shadow
	// copy storage volume.
	DiffAreaVolumeName string `idl:"name:pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	// llMaximumDiffSpace: The maximum number of BYTEs that the shadow copy storage will
	// occupy. The server MAY automatically delete shadow copies based on an implementation-specific
	// algorithm that reclaims space for newer shadow copies.
	MaximumDiffSpace int64 `idl:"name:llMaximumDiffSpace" json:"maximum_diff_space"`
}

func (o *ChangeDiffAreaMaximumSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangeDiffAreaMaximumSizeOperation) *xxx_ChangeDiffAreaMaximumSizeOperation {
	if op == nil {
		op = &xxx_ChangeDiffAreaMaximumSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VolumeName = o.VolumeName
	op.DiffAreaVolumeName = o.DiffAreaVolumeName
	op.MaximumDiffSpace = o.MaximumDiffSpace
	return op
}

func (o *ChangeDiffAreaMaximumSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangeDiffAreaMaximumSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeName = op.VolumeName
	o.DiffAreaVolumeName = op.DiffAreaVolumeName
	o.MaximumDiffSpace = op.MaximumDiffSpace
}
func (o *ChangeDiffAreaMaximumSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangeDiffAreaMaximumSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeDiffAreaMaximumSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangeDiffAreaMaximumSizeResponse structure represents the ChangeDiffAreaMaximumSize operation response
type ChangeDiffAreaMaximumSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ChangeDiffAreaMaximumSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangeDiffAreaMaximumSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangeDiffAreaMaximumSizeOperation) *xxx_ChangeDiffAreaMaximumSizeOperation {
	if op == nil {
		op = &xxx_ChangeDiffAreaMaximumSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ChangeDiffAreaMaximumSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangeDiffAreaMaximumSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ChangeDiffAreaMaximumSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangeDiffAreaMaximumSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeDiffAreaMaximumSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryVolumesSupportedForDiffAreasOperation structure represents the QueryVolumesSupportedForDiffAreas operation
type xxx_QueryVolumesSupportedForDiffAreasOperation struct {
	This               *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat             `idl:"name:That" json:"that"`
	OriginalVolumeName string                     `idl:"name:pwszOriginalVolumeName" json:"original_volume_name"`
	Enum               *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	Return             int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) OpNum() int { return 5 }

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) OpName() string {
	return "/IVssDifferentialSoftwareSnapshotMgmt/v0/QueryVolumesSupportedForDiffAreas"
}

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszOriginalVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.OriginalVolumeName != "" {
			_ptr_pwszOriginalVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.OriginalVolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.OriginalVolumeName, _ptr_pwszOriginalVolumeName); err != nil {
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

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszOriginalVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszOriginalVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.OriginalVolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszOriginalVolumeName := func(ptr interface{}) { o.OriginalVolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.OriginalVolumeName, _s_pwszOriginalVolumeName, _ptr_pwszOriginalVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.EnumManagementObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryVolumesSupportedForDiffAreasOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.EnumManagementObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.EnumManagementObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryVolumesSupportedForDiffAreasRequest structure represents the QueryVolumesSupportedForDiffAreas operation request
type QueryVolumesSupportedForDiffAreasRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszOriginalVolumeName: A null-terminated UNICODE string that contains the drive
	// letter, mount point, or volume mount name of the original volume.
	OriginalVolumeName string `idl:"name:pwszOriginalVolumeName" json:"original_volume_name"`
}

func (o *QueryVolumesSupportedForDiffAreasRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesSupportedForDiffAreasOperation) *xxx_QueryVolumesSupportedForDiffAreasOperation {
	if op == nil {
		op = &xxx_QueryVolumesSupportedForDiffAreasOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OriginalVolumeName = o.OriginalVolumeName
	return op
}

func (o *QueryVolumesSupportedForDiffAreasRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesSupportedForDiffAreasOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OriginalVolumeName = op.OriginalVolumeName
}
func (o *QueryVolumesSupportedForDiffAreasRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryVolumesSupportedForDiffAreasRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesSupportedForDiffAreasOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryVolumesSupportedForDiffAreasResponse structure represents the QueryVolumesSupportedForDiffAreas operation response
type QueryVolumesSupportedForDiffAreasResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IVssEnumMgmtObject pointer that upon completion, contains
	// a collection of volumes that can be used to create shadow copy storage associations
	// with the specified original volume. Each element in the collection MUST be a VSS_DIFF_VOLUME_PROP
	// structure. A caller MUST release the ppEnum received when the caller is done with
	// it.
	Enum *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryVolumesSupportedForDiffAreas return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryVolumesSupportedForDiffAreasResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesSupportedForDiffAreasOperation) *xxx_QueryVolumesSupportedForDiffAreasOperation {
	if op == nil {
		op = &xxx_QueryVolumesSupportedForDiffAreasOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryVolumesSupportedForDiffAreasResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesSupportedForDiffAreasOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryVolumesSupportedForDiffAreasResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryVolumesSupportedForDiffAreasResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesSupportedForDiffAreasOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDiffAreasForVolumeOperation structure represents the QueryDiffAreasForVolume operation
type xxx_QueryDiffAreasForVolumeOperation struct {
	This       *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat             `idl:"name:That" json:"that"`
	VolumeName string                     `idl:"name:pwszVolumeName" json:"volume_name"`
	Enum       *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	Return     int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDiffAreasForVolumeOperation) OpNum() int { return 6 }

func (o *xxx_QueryDiffAreasForVolumeOperation) OpName() string {
	return "/IVssDifferentialSoftwareSnapshotMgmt/v0/QueryDiffAreasForVolume"
}

func (o *xxx_QueryDiffAreasForVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasForVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.VolumeName != "" {
			_ptr_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeName, _ptr_pwszVolumeName); err != nil {
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

func (o *xxx_QueryDiffAreasForVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeName, _s_pwszVolumeName, _ptr_pwszVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasForVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasForVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.EnumManagementObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryDiffAreasForVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.EnumManagementObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.EnumManagementObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryDiffAreasForVolumeRequest structure represents the QueryDiffAreasForVolume operation request
type QueryDiffAreasForVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszVolumeName: A null-terminated UNICODE string that contains the drive letter,
	// mount point, or volume mount name of the original volume for which the existing shadow
	// copy association collection is requested.
	VolumeName string `idl:"name:pwszVolumeName" json:"volume_name"`
}

func (o *QueryDiffAreasForVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryDiffAreasForVolumeOperation) *xxx_QueryDiffAreasForVolumeOperation {
	if op == nil {
		op = &xxx_QueryDiffAreasForVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VolumeName = o.VolumeName
	return op
}

func (o *QueryDiffAreasForVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDiffAreasForVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeName = op.VolumeName
}
func (o *QueryDiffAreasForVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryDiffAreasForVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDiffAreasForVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDiffAreasForVolumeResponse structure represents the QueryDiffAreasForVolume operation response
type QueryDiffAreasForVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IVssEnumMgmtObject pointer that upon completion, contains
	// a collection of shadow copy storage associations that are providing shadow copy storage
	// for the specified original volume. Each element in the collection MUST be a VSS_DIFF_AREA_PROP
	// structure. A caller MUST release the ppEnum received when the caller is done with
	// it.
	Enum *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryDiffAreasForVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDiffAreasForVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryDiffAreasForVolumeOperation) *xxx_QueryDiffAreasForVolumeOperation {
	if op == nil {
		op = &xxx_QueryDiffAreasForVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryDiffAreasForVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDiffAreasForVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryDiffAreasForVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryDiffAreasForVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDiffAreasForVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDiffAreasOnVolumeOperation structure represents the QueryDiffAreasOnVolume operation
type xxx_QueryDiffAreasOnVolumeOperation struct {
	This       *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat             `idl:"name:That" json:"that"`
	VolumeName string                     `idl:"name:pwszVolumeName" json:"volume_name"`
	Enum       *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	Return     int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDiffAreasOnVolumeOperation) OpNum() int { return 7 }

func (o *xxx_QueryDiffAreasOnVolumeOperation) OpName() string {
	return "/IVssDifferentialSoftwareSnapshotMgmt/v0/QueryDiffAreasOnVolume"
}

func (o *xxx_QueryDiffAreasOnVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasOnVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.VolumeName != "" {
			_ptr_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeName, _ptr_pwszVolumeName); err != nil {
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

func (o *xxx_QueryDiffAreasOnVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeName, _s_pwszVolumeName, _ptr_pwszVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasOnVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDiffAreasOnVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.EnumManagementObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryDiffAreasOnVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.EnumManagementObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.EnumManagementObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryDiffAreasOnVolumeRequest structure represents the QueryDiffAreasOnVolume operation request
type QueryDiffAreasOnVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszVolumeName: A null-terminated UNICODE string that contains the drive letter,
	// mount point, or volume mount name of the shadow copy storage volume on which the
	// existing shadow copy association collection is requested.
	VolumeName string `idl:"name:pwszVolumeName" json:"volume_name"`
}

func (o *QueryDiffAreasOnVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryDiffAreasOnVolumeOperation) *xxx_QueryDiffAreasOnVolumeOperation {
	if op == nil {
		op = &xxx_QueryDiffAreasOnVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VolumeName = o.VolumeName
	return op
}

func (o *QueryDiffAreasOnVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDiffAreasOnVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeName = op.VolumeName
}
func (o *QueryDiffAreasOnVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryDiffAreasOnVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDiffAreasOnVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDiffAreasOnVolumeResponse structure represents the QueryDiffAreasOnVolume operation response
type QueryDiffAreasOnVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IVssEnumMgmtObject pointer that upon completion, contains
	// a collection of shadow copy storage associations that are located on the specified
	// volume. Each element in the collection MUST be a VSS_DIFF_AREA_PROP structure. A
	// caller MUST release the ppEnum received when the caller is done with it.
	Enum *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryDiffAreasOnVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDiffAreasOnVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryDiffAreasOnVolumeOperation) *xxx_QueryDiffAreasOnVolumeOperation {
	if op == nil {
		op = &xxx_QueryDiffAreasOnVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryDiffAreasOnVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDiffAreasOnVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryDiffAreasOnVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryDiffAreasOnVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDiffAreasOnVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
