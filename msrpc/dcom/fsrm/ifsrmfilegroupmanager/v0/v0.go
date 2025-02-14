package ifsrmfilegroupmanager

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
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = fsrm.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmFileGroupManager interface identifier 426677d5-018c-485c-8a51-20b86d00bdc4
	FileGroupManagerIID = &dcom.IID{Data1: 0x426677d5, Data2: 0x018c, Data3: 0x485c, Data4: []byte{0x8a, 0x51, 0x20, 0xb8, 0x6d, 0x00, 0xbd, 0xc4}}
	// Syntax UUID
	FileGroupManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x426677d5, TimeMid: 0x18c, TimeHiAndVersion: 0x485c, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x51, Node: [6]uint8{0x20, 0xb8, 0x6d, 0x0, 0xbd, 0xc4}}
	// Syntax ID
	FileGroupManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: FileGroupManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmFileGroupManager interface.
type FileGroupManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The CreateFileGroup method creates a blank Non-Persisted File Group Instance (section
	// 3.2.1.3.4.2).
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
	//	| 0x80070057 E_INVALIDARG | The fileGroup parameter is NULL. |
	//	+-------------------------+----------------------------------+
	CreateFileGroup(context.Context, *CreateFileGroupRequest, ...dcerpc.CallOption) (*CreateFileGroupResponse, error)

	// The GetFileGroup method returns a pointer to the file group with the specified Name
	// property from the List of Persisted File Groups (section 3.2.1.3).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | The specified file group could not be found.                                     |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045308 FSRM_E_INVALID_NAME | The specified name is not valid.                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530D FSRM_E_OUT_OF_RANGE | The content of the name parameter exceeds the maximum length of 4,000            |
	//	|                                | characters.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The fileGroup parameter is NULL.                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetFileGroup(context.Context, *GetFileGroupRequest, ...dcerpc.CallOption) (*GetFileGroupResponse, error)

	// The EnumFileGroups method returns all the file groups from the List of Persisted
	// File Groups (section 3.2.1.3) of the server.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+-----------------------------------------------------------------+
	//	|             RETURN              |                                                                 |
	//	|           VALUE/CODE            |                           DESCRIPTION                           |
	//	|                                 |                                                                 |
	//	+---------------------------------+-----------------------------------------------------------------+
	//	+---------------------------------+-----------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions values. |
	//	+---------------------------------+-----------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileGroups parameter is NULL.                               |
	//	+---------------------------------+-----------------------------------------------------------------+
	EnumFileGroups(context.Context, *EnumFileGroupsRequest, ...dcerpc.CallOption) (*EnumFileGroupsResponse, error)

	// The ExportFileGroups method exports an XML string representation of the File Server
	// Resource Manager Protocol file groups from the List of Persisted File Groups (section
	// 3.2.1.3).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified file group could not be found.                                     |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | This code is returned for the following reasons: The serializedFileGroups        |
	//	|                             | parameter is NULL. The fileGroupNamesArray parameter is not a variant array of   |
	//	|                             | BSTRs.                                                                           |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	ExportFileGroups(context.Context, *ExportFileGroupsRequest, ...dcerpc.CallOption) (*ExportFileGroupsResponse, error)

	// The ImportFileGroups method imports file groups from the XML string of file groups.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND              | This specified file group could not be found.                                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004530B FSRM_E_INVALID_IMPORT_VERSION | The version of the imported objects is not valid.                                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | This code is returned for the following reasons: The serializedFileGroups        |
	//	|                                          | parameter is NULL. The fileGroups parameter is NULL. The fileGroupNamesArray     |
	//	|                                          | parameter is not a variant array of BSTRs.                                       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	ImportFileGroups(context.Context, *ImportFileGroupsRequest, ...dcerpc.CallOption) (*ImportFileGroupsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileGroupManagerClient
}

type xxx_DefaultFileGroupManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileGroupManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultFileGroupManagerClient) CreateFileGroup(ctx context.Context, in *CreateFileGroupRequest, opts ...dcerpc.CallOption) (*CreateFileGroupResponse, error) {
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
	out := &CreateFileGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileGroupManagerClient) GetFileGroup(ctx context.Context, in *GetFileGroupRequest, opts ...dcerpc.CallOption) (*GetFileGroupResponse, error) {
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
	out := &GetFileGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileGroupManagerClient) EnumFileGroups(ctx context.Context, in *EnumFileGroupsRequest, opts ...dcerpc.CallOption) (*EnumFileGroupsResponse, error) {
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
	out := &EnumFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileGroupManagerClient) ExportFileGroups(ctx context.Context, in *ExportFileGroupsRequest, opts ...dcerpc.CallOption) (*ExportFileGroupsResponse, error) {
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
	out := &ExportFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileGroupManagerClient) ImportFileGroups(ctx context.Context, in *ImportFileGroupsRequest, opts ...dcerpc.CallOption) (*ImportFileGroupsResponse, error) {
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
	out := &ImportFileGroupsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileGroupManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileGroupManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileGroupManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) FileGroupManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileGroupManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewFileGroupManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileGroupManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileGroupManagerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultFileGroupManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CreateFileGroupOperation structure represents the CreateFileGroup operation
type xxx_CreateFileGroupOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	FileGroup *fsrm.FileGroup `idl:"name:fileGroup" json:"file_group"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFileGroupOperation) OpNum() int { return 7 }

func (o *xxx_CreateFileGroupOperation) OpName() string {
	return "/IFsrmFileGroupManager/v0/CreateFileGroup"
}

func (o *xxx_CreateFileGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateFileGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateFileGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileGroup {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileGroup}(interface))
	{
		if o.FileGroup != nil {
			_ptr_fileGroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroup != nil {
					if err := o.FileGroup.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileGroup{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroup, _ptr_fileGroup); err != nil {
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

func (o *xxx_CreateFileGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileGroup {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileGroup}(interface))
	{
		_ptr_fileGroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroup == nil {
				o.FileGroup = &fsrm.FileGroup{}
			}
			if err := o.FileGroup.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroup := func(ptr interface{}) { o.FileGroup = *ptr.(**fsrm.FileGroup) }
		if err := w.ReadPointer(&o.FileGroup, _s_fileGroup, _ptr_fileGroup); err != nil {
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

// CreateFileGroupRequest structure represents the CreateFileGroup operation request
type CreateFileGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateFileGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateFileGroupOperation) *xxx_CreateFileGroupOperation {
	if op == nil {
		op = &xxx_CreateFileGroupOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *CreateFileGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFileGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateFileGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateFileGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFileGroupResponse structure represents the CreateFileGroup operation response
type CreateFileGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileGroup: Pointer to an IFsrmFileGroup interface pointer (section 3.2.4.2.23) that
	// upon completion points to a blank file group. To have the server add the file group
	// to its List of Persisted File Groups (section 3.2.1.3), the caller MUST call Commit
	// (section 3.2.4.2.23.1) on the file group. The caller MUST release the file group
	// received when it is done with it.
	FileGroup *fsrm.FileGroup `idl:"name:fileGroup" json:"file_group"`
	// Return: The CreateFileGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFileGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateFileGroupOperation) *xxx_CreateFileGroupOperation {
	if op == nil {
		op = &xxx_CreateFileGroupOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileGroup = op.FileGroup
	o.Return = op.Return
	return op
}

func (o *CreateFileGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFileGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileGroup = op.FileGroup
	o.Return = op.Return
}
func (o *CreateFileGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateFileGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileGroupOperation structure represents the GetFileGroup operation
type xxx_GetFileGroupOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Name      *oaut.String    `idl:"name:name" json:"name"`
	FileGroup *fsrm.FileGroup `idl:"name:fileGroup" json:"file_group"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileGroupOperation) OpNum() int { return 8 }

func (o *xxx_GetFileGroupOperation) OpName() string { return "/IFsrmFileGroupManager/v0/GetFileGroup" }

func (o *xxx_GetFileGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetFileGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileGroup {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileGroup}(interface))
	{
		if o.FileGroup != nil {
			_ptr_fileGroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroup != nil {
					if err := o.FileGroup.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileGroup{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroup, _ptr_fileGroup); err != nil {
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

func (o *xxx_GetFileGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileGroup {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileGroup}(interface))
	{
		_ptr_fileGroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroup == nil {
				o.FileGroup = &fsrm.FileGroup{}
			}
			if err := o.FileGroup.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroup := func(ptr interface{}) { o.FileGroup = *ptr.(**fsrm.FileGroup) }
		if err := w.ReadPointer(&o.FileGroup, _s_fileGroup, _ptr_fileGroup); err != nil {
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

// GetFileGroupRequest structure represents the GetFileGroup operation request
type GetFileGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// name: Contains the Name of the file group to return. The maximum length of this string
	// MUST be 4,000 characters.
	Name *oaut.String `idl:"name:name" json:"name"`
}

func (o *GetFileGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileGroupOperation) *xxx_GetFileGroupOperation {
	if op == nil {
		op = &xxx_GetFileGroupOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Name = op.Name
	return op
}

func (o *GetFileGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *GetFileGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileGroupResponse structure represents the GetFileGroup operation response
type GetFileGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileGroup: Pointer to an IFsrmFileGroup interface pointer (section 3.2.4.2.23) that
	// upon completion points to the file group with the specified Name property. The caller
	// MUST release the file group when it is done with it.
	FileGroup *fsrm.FileGroup `idl:"name:fileGroup" json:"file_group"`
	// Return: The GetFileGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileGroupOperation) *xxx_GetFileGroupOperation {
	if op == nil {
		op = &xxx_GetFileGroupOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileGroup = op.FileGroup
	o.Return = op.Return
	return op
}

func (o *GetFileGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileGroup = op.FileGroup
	o.Return = op.Return
}
func (o *GetFileGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFileGroupsOperation structure represents the EnumFileGroups operation
type xxx_EnumFileGroupsOperation struct {
	This       *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Options    fsrm.EnumOptions            `idl:"name:options" json:"options"`
	FileGroups *fsrm.CommittableCollection `idl:"name:fileGroups" json:"file_groups"`
	Return     int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFileGroupsOperation) OpNum() int { return 9 }

func (o *xxx_EnumFileGroupsOperation) OpName() string {
	return "/IFsrmFileGroupManager/v0/EnumFileGroups"
}

func (o *xxx_EnumFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.FileGroups != nil {
			_ptr_fileGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroups != nil {
					if err := o.FileGroups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroups, _ptr_fileGroups); err != nil {
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

func (o *xxx_EnumFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_fileGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroups == nil {
				o.FileGroups = &fsrm.CommittableCollection{}
			}
			if err := o.FileGroups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroups := func(ptr interface{}) { o.FileGroups = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.FileGroups, _s_fileGroups, _ptr_fileGroups); err != nil {
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

// EnumFileGroupsRequest structure represents the EnumFileGroups operation request
type EnumFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the filegroups.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumFileGroupsOperation) *xxx_EnumFileGroupsOperation {
	if op == nil {
		op = &xxx_EnumFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Options = op.Options
	return op
}

func (o *EnumFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *EnumFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFileGroupsResponse structure represents the EnumFileGroups operation response
type EnumFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileGroups: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion SHOULD contain a pointer to every file group on the server.
	// The caller MUST release the collection when it is done with it.
	FileGroups *fsrm.CommittableCollection `idl:"name:fileGroups" json:"file_groups"`
	// Return: The EnumFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumFileGroupsOperation) *xxx_EnumFileGroupsOperation {
	if op == nil {
		op = &xxx_EnumFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileGroups = op.FileGroups
	o.Return = op.Return
	return op
}

func (o *EnumFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileGroups = op.FileGroups
	o.Return = op.Return
}
func (o *EnumFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportFileGroupsOperation structure represents the ExportFileGroups operation
type xxx_ExportFileGroupsOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	FileGroupNamesArray  *oaut.Variant  `idl:"name:fileGroupNamesArray" json:"file_group_names_array"`
	SerializedFileGroups *oaut.String   `idl:"name:serializedFileGroups" json:"serialized_file_groups"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportFileGroupsOperation) OpNum() int { return 10 }

func (o *xxx_ExportFileGroupsOperation) OpName() string {
	return "/IFsrmFileGroupManager/v0/ExportFileGroups"
}

func (o *xxx_ExportFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fileGroupNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.FileGroupNamesArray != nil {
			_ptr_fileGroupNamesArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroupNamesArray != nil {
					if err := o.FileGroupNamesArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroupNamesArray, _ptr_fileGroupNamesArray); err != nil {
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

func (o *xxx_ExportFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fileGroupNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_fileGroupNamesArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroupNamesArray == nil {
				o.FileGroupNamesArray = &oaut.Variant{}
			}
			if err := o.FileGroupNamesArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroupNamesArray := func(ptr interface{}) { o.FileGroupNamesArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.FileGroupNamesArray, _s_fileGroupNamesArray, _ptr_fileGroupNamesArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// serializedFileGroups {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedFileGroups != nil {
			_ptr_serializedFileGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedFileGroups != nil {
					if err := o.SerializedFileGroups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedFileGroups, _ptr_serializedFileGroups); err != nil {
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

func (o *xxx_ExportFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// serializedFileGroups {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedFileGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedFileGroups == nil {
				o.SerializedFileGroups = &oaut.String{}
			}
			if err := o.SerializedFileGroups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedFileGroups := func(ptr interface{}) { o.SerializedFileGroups = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedFileGroups, _s_serializedFileGroups, _ptr_serializedFileGroups); err != nil {
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

// ExportFileGroupsRequest structure represents the ExportFileGroups operation request
type ExportFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// fileGroupNamesArray: Pointer to a SAFEARRAY that contains the names of file groups
	// to export.
	FileGroupNamesArray *oaut.Variant `idl:"name:fileGroupNamesArray" json:"file_group_names_array"`
}

func (o *ExportFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportFileGroupsOperation) *xxx_ExportFileGroupsOperation {
	if op == nil {
		op = &xxx_ExportFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.FileGroupNamesArray = op.FileGroupNamesArray
	return op
}

func (o *ExportFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FileGroupNamesArray = op.FileGroupNamesArray
}
func (o *ExportFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportFileGroupsResponse structure represents the ExportFileGroups operation response
type ExportFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// serializedFileGroups: Pointer to a variable that upon completion contains the XML
	// string representation of all the specified file groups.
	SerializedFileGroups *oaut.String `idl:"name:serializedFileGroups" json:"serialized_file_groups"`
	// Return: The ExportFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportFileGroupsOperation) *xxx_ExportFileGroupsOperation {
	if op == nil {
		op = &xxx_ExportFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.SerializedFileGroups = op.SerializedFileGroups
	o.Return = op.Return
	return op
}

func (o *ExportFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SerializedFileGroups = op.SerializedFileGroups
	o.Return = op.Return
}
func (o *ExportFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportFileGroupsOperation structure represents the ImportFileGroups operation
type xxx_ImportFileGroupsOperation struct {
	This                 *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat              `idl:"name:That" json:"that"`
	SerializedFileGroups *oaut.String                `idl:"name:serializedFileGroups" json:"serialized_file_groups"`
	FileGroupNamesArray  *oaut.Variant               `idl:"name:fileGroupNamesArray" json:"file_group_names_array"`
	FileGroups           *fsrm.CommittableCollection `idl:"name:fileGroups" json:"file_groups"`
	Return               int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportFileGroupsOperation) OpNum() int { return 11 }

func (o *xxx_ImportFileGroupsOperation) OpName() string {
	return "/IFsrmFileGroupManager/v0/ImportFileGroups"
}

func (o *xxx_ImportFileGroupsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFileGroupsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serializedFileGroups {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedFileGroups != nil {
			_ptr_serializedFileGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedFileGroups != nil {
					if err := o.SerializedFileGroups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedFileGroups, _ptr_serializedFileGroups); err != nil {
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
	// fileGroupNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.FileGroupNamesArray != nil {
			_ptr_fileGroupNamesArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroupNamesArray != nil {
					if err := o.FileGroupNamesArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroupNamesArray, _ptr_fileGroupNamesArray); err != nil {
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

func (o *xxx_ImportFileGroupsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serializedFileGroups {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedFileGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedFileGroups == nil {
				o.SerializedFileGroups = &oaut.String{}
			}
			if err := o.SerializedFileGroups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedFileGroups := func(ptr interface{}) { o.SerializedFileGroups = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedFileGroups, _s_serializedFileGroups, _ptr_serializedFileGroups); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fileGroupNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_fileGroupNamesArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroupNamesArray == nil {
				o.FileGroupNamesArray = &oaut.Variant{}
			}
			if err := o.FileGroupNamesArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroupNamesArray := func(ptr interface{}) { o.FileGroupNamesArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.FileGroupNamesArray, _s_fileGroupNamesArray, _ptr_fileGroupNamesArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFileGroupsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportFileGroupsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.FileGroups != nil {
			_ptr_fileGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileGroups != nil {
					if err := o.FileGroups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileGroups, _ptr_fileGroups); err != nil {
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

func (o *xxx_ImportFileGroupsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_fileGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileGroups == nil {
				o.FileGroups = &fsrm.CommittableCollection{}
			}
			if err := o.FileGroups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileGroups := func(ptr interface{}) { o.FileGroups = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.FileGroups, _s_fileGroups, _ptr_fileGroups); err != nil {
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

// ImportFileGroupsRequest structure represents the ImportFileGroups operation request
type ImportFileGroupsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// serializedFileGroups: Contains the XML string representation of a list of file groups.
	// There is no maximum length for this string.
	SerializedFileGroups *oaut.String `idl:"name:serializedFileGroups" json:"serialized_file_groups"`
	// fileGroupNamesArray: Pointer to a SAFEARRAY that contains the names of file groups
	// to import.
	FileGroupNamesArray *oaut.Variant `idl:"name:fileGroupNamesArray" json:"file_group_names_array"`
}

func (o *ImportFileGroupsRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportFileGroupsOperation) *xxx_ImportFileGroupsOperation {
	if op == nil {
		op = &xxx_ImportFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.SerializedFileGroups = op.SerializedFileGroups
	o.FileGroupNamesArray = op.FileGroupNamesArray
	return op
}

func (o *ImportFileGroupsRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportFileGroupsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SerializedFileGroups = op.SerializedFileGroups
	o.FileGroupNamesArray = op.FileGroupNamesArray
}
func (o *ImportFileGroupsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportFileGroupsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFileGroupsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportFileGroupsResponse structure represents the ImportFileGroups operation response
type ImportFileGroupsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileGroups: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that upon completion contains an IFsrmFileGroup interface pointer (section 3.2.4.2.23)
	// for each of the imported file groups. The caller MUST release the collection when
	// it is done with it.
	FileGroups *fsrm.CommittableCollection `idl:"name:fileGroups" json:"file_groups"`
	// Return: The ImportFileGroups return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportFileGroupsResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportFileGroupsOperation) *xxx_ImportFileGroupsOperation {
	if op == nil {
		op = &xxx_ImportFileGroupsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileGroups = op.FileGroups
	o.Return = op.Return
	return op
}

func (o *ImportFileGroupsResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportFileGroupsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileGroups = op.FileGroups
	o.Return = op.Return
}
func (o *ImportFileGroupsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportFileGroupsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportFileGroupsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
