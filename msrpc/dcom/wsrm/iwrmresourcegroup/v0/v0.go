package iwrmresourcegroup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMResourceGroup interface identifier bc681469-9dd9-4bf4-9b3d-709f69efe431
	ResourceGroupIID = &dcom.IID{Data1: 0xbc681469, Data2: 0x9dd9, Data3: 0x4bf4, Data4: []byte{0x9b, 0x3d, 0x70, 0x9f, 0x69, 0xef, 0xe4, 0x31}}
	// Syntax UUID
	ResourceGroupSyntaxUUID = &uuid.UUID{TimeLow: 0xbc681469, TimeMid: 0x9dd9, TimeHiAndVersion: 0x4bf4, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x3d, Node: [6]uint8{0x70, 0x9f, 0x69, 0xef, 0xe4, 0x31}}
	// Syntax ID
	ResourceGroupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ResourceGroupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMResourceGroup interface.
type ResourceGroupClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetResourceGroupInfo method gets information about the resource group with the
	// specified ID. If the ID is "\", this method returns all selection criteria.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | Operation successful.                                                            |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                    | One or more arguments are invalid.                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                | The specified name contains characters that are invalid. The name cannot start   |
	//	|                                            | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                            | following characters; "\" cannot be used with other characters. / ? * | : < > "  |
	//	|                                            | , ;                                                                              |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID | The request has been aborted because the specified resource group does not       |
	//	|                                            | exist.                                                                           |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMResourceGroup interface methods are specified in section 3.2.4.10.
	GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest, ...dcerpc.CallOption) (*GetResourceGroupInfoResponse, error)

	// The ModifyResourceGroup method modifies an existing resource group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                     |                                                                                  |
	//	|                  VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                               | The call was successful.                                                         |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0069 WRM_ERR_OLD_INFORMATION            | The XML timestamp is out-of-date.                                                |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006A WRM_ERR_NO_TIMESTAMP_PRESENT       | The specified resource group could not be modified because the XML timestamp in  |
	//	|                                               | the bstrResourceGroupInfo parameter was not found.                               |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                   | The specified name of the ProcessMatchingCriteria element contains characters    |
	//	|                                               | that are not valid. The name cannot start with a hyphen ("-"), cannot contain    |
	//	|                                               | spaces, and cannot contain any of the following characters: \ / ? * | : < > " ,  |
	//	|                                               | ;                                                                                |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER          | The XML data that is maintained by the management service is not valid or cannot |
	//	|                                               | be processed.<115>                                                               |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012D WRM_ERR_TOO_LONG_RESOURCE_GROUP_ID | The resource group name has exceeded an implementation-defined<116> limit. The   |
	//	|                                               | modify request has been aborted.                                                 |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID    | The specified resource group does not exist, or "\" was specified. "\" is a      |
	//	|                                               | reserved name for resource group identifiers.                                    |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0132 WRM_ERR_PATH_LIMIT_EXCEEDED        | The command-line length has exceeded an implementation-defined limit.<117>       |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0134 WRM_ERR_USER_LIMIT_EXCEEDED        | The user name or group value has exceeded an implementation-defined limit.<118>  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0135 WRM_ERR_RULE_LIMIT_EXCEEDED        | The specified resource group could not be modified because a resource group      |
	//	|                                               | cannot have more than 64 rules.                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0138 WRM_ERR_RESERVED_RESOURCEGROUP     | The specified resource group is built-in. It cannot be modified.<119>            |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMResourceGroup interface methods are specified in section 3.2.4.10.
	ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest, ...dcerpc.CallOption) (*ModifyResourceGroupResponse, error)

	// The CreateResourceGroup function creates a new resource group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                          |                                                                                  |
	//	|                       VALUE/CODE                        |                                   DESCRIPTION                                    |
	//	|                                                         |                                                                                  |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                         | The operation was successfully done.                                             |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                 | One or more arguments are invalid.                                               |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                             | The specified name contains characters that are not valid. The name cannot       |
	//	|                                                         | start with a hyphen ("-"), cannot contain spaces, and cannot contain any of the  |
	//	|                                                         | following characters: \ / ? * | : < > " , ;                                      |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012F WRM_ERR_RESOURCEGROUPID_ALREADY_EXISTS       | The request has been aborted because the resource group with the specified name  |
	//	|                                                         | already exists.                                                                  |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0136 WRM_ERR_PMC_LIMIT_EXCEEDED                   | The total number of resource groups has exceeded an implementation-defined       |
	//	|                                                         | limit.<120>.                                                                     |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0139 WRM_ERR_CANNOT_CREATE_RESERVED_RESOURCEGROUP | A user-created resource group cannot have the same name as that of a built-in    |
	//	|                                                         | resource group.                                                                  |
	//	+---------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMResourceGroup interface methods are specified in section 3.2.4.10.
	CreateResourceGroup(context.Context, *CreateResourceGroupRequest, ...dcerpc.CallOption) (*CreateResourceGroupResponse, error)

	// The DeleteResourceGroup deletes a resource group with the specified identifier.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | Operation successful.                                                            |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                    | One or more arguments are invalid.                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                | The specified name contains characters that are invalid. The name cannot start   |
	//	|                                            | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                            | following characters: \ / ? * | : < > " , ;                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID | The specified resource group does not exist.                                     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0137 WRM_ERR_DELETING_RESOURCE_GROUP | resource group that are members of one or more RAPs cannot be deleted.           |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0138 WRM_ERR_RESERVED_RESOURCEGROUP  | The specified resource group is built-in. It cannot be modified.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMResourceGroup interface methods are specified in section 3.2.4.10.
	DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest, ...dcerpc.CallOption) (*DeleteResourceGroupResponse, error)

	// The RenameResourceGroup renames an existing resource group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                       |                                                                                  |
	//	|                    VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                        | The call was successful.                                                         |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                           | One or more arguments are invalid.                                               |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                       | The specified name contains characters that are valid. The name cannot start     |
	//	|                                                   | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                                   | following characters: \ / ? * | : < > " , ;                                      |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012D WRM_ERR_TOO_LONG_RESOURCE_GROUP_ID     | The resource group name has exceeded an implementation-defined limit. <121>.     |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID        | The specified resource group does not exist.                                     |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0Xc1FF012F WRM_ERR_RESOURCEGROUPID_ALREADY_EXISTS | A resource group with the specified name already exists.                         |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0130 WRM_ERR_RESOURCEGROUPID_RESERVED_WORD  | "\" is a reserved name for resource group identifiers.                           |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0138 WRM_ERR_RESERVED_RESOURCEGROUP         | The specified resource group is built-in. It cannot be modified.                 |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMResourceGroup interface methods are specified in section 3.2.4.10.
	RenameResourceGroup(context.Context, *RenameResourceGroupRequest, ...dcerpc.CallOption) (*RenameResourceGroupResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ResourceGroupClient
}

type xxx_DefaultResourceGroupClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultResourceGroupClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultResourceGroupClient) GetResourceGroupInfo(ctx context.Context, in *GetResourceGroupInfoRequest, opts ...dcerpc.CallOption) (*GetResourceGroupInfoResponse, error) {
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
	out := &GetResourceGroupInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceGroupClient) ModifyResourceGroup(ctx context.Context, in *ModifyResourceGroupRequest, opts ...dcerpc.CallOption) (*ModifyResourceGroupResponse, error) {
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
	out := &ModifyResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceGroupClient) CreateResourceGroup(ctx context.Context, in *CreateResourceGroupRequest, opts ...dcerpc.CallOption) (*CreateResourceGroupResponse, error) {
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
	out := &CreateResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceGroupClient) DeleteResourceGroup(ctx context.Context, in *DeleteResourceGroupRequest, opts ...dcerpc.CallOption) (*DeleteResourceGroupResponse, error) {
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
	out := &DeleteResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceGroupClient) RenameResourceGroup(ctx context.Context, in *RenameResourceGroupRequest, opts ...dcerpc.CallOption) (*RenameResourceGroupResponse, error) {
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
	out := &RenameResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultResourceGroupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultResourceGroupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultResourceGroupClient) IPID(ctx context.Context, ipid *dcom.IPID) ResourceGroupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultResourceGroupClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewResourceGroupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ResourceGroupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ResourceGroupSyntaxV0_0))...)
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
	return &xxx_DefaultResourceGroupClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetResourceGroupInfoOperation structure represents the GetResourceGroupInfo operation
type xxx_GetResourceGroupInfoOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ResourceGroupInfo *oaut.String   `idl:"name:pbstrResourceGroupInfo" json:"resource_group_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourceGroupInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetResourceGroupInfoOperation) OpName() string {
	return "/IWRMResourceGroup/v0/GetResourceGroupInfo"
}

func (o *xxx_GetResourceGroupInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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

func (o *xxx_GetResourceGroupInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrResourceGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_pbstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_pbstrResourceGroupInfo); err != nil {
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

func (o *xxx_GetResourceGroupInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrResourceGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_pbstrResourceGroupInfo, _ptr_pbstrResourceGroupInfo); err != nil {
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

// GetResourceGroupInfoRequest structure represents the GetResourceGroupInfo operation request
type GetResourceGroupInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrResourceGroupName: A string that specifies the name of the selection criteria.
	ResourceGroupName *oaut.String `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
}

func (o *GetResourceGroupInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) *xxx_GetResourceGroupInfoOperation {
	if op == nil {
		op = &xxx_GetResourceGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupName = o.ResourceGroupName
	return op
}

func (o *GetResourceGroupInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupName = op.ResourceGroupName
}
func (o *GetResourceGroupInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResourceGroupInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceGroupInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourceGroupInfoResponse structure represents the GetResourceGroupInfo operation response
type GetResourceGroupInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrResourceGroupInfo: A pointer to a string that returns the selection criteria,
	// in the form of a ProcessMatchingCriteria element (section 2.2.5.24). For an example,
	// see the ProcessMatchingCriteria example (section 4.2.20).
	ResourceGroupInfo *oaut.String `idl:"name:pbstrResourceGroupInfo" json:"resource_group_info"`
	// Return: The GetResourceGroupInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResourceGroupInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) *xxx_GetResourceGroupInfoOperation {
	if op == nil {
		op = &xxx_GetResourceGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ResourceGroupInfo = o.ResourceGroupInfo
	op.Return = o.Return
	return op
}

func (o *GetResourceGroupInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResourceGroupInfo = op.ResourceGroupInfo
	o.Return = op.Return
}
func (o *GetResourceGroupInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResourceGroupInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceGroupInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyResourceGroupOperation structure represents the ModifyResourceGroup operation
type xxx_ModifyResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	Overwrite         bool           `idl:"name:bOverwrite" json:"overwrite"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyResourceGroupOperation) OpNum() int { return 8 }

func (o *xxx_ModifyResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/ModifyResourceGroup"
}

func (o *xxx_ModifyResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_bstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
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
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		if !o.Overwrite {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_bstrResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwrite int32
		if err := w.ReadData(&_bOverwrite); err != nil {
			return err
		}
		o.Overwrite = _bOverwrite != 0
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyResourceGroupRequest structure represents the ModifyResourceGroup operation request
type ModifyResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrResourceGroupInfo: A string that contains the modified resource group, in the
	// form of a ProcessMatchingCriteria element (section 2.2.5.24). Sample XML is provided
	// in ProcessMatchingCriteria example (section 4.2.20).
	ResourceGroupInfo *oaut.String `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	// bOverwrite: A Boolean value that specifies whether to ignore the timestamp of the
	// resource group when validating.
	//
	// A timestamp MUST be defined inside a common node at the root level of an XML element,
	// as shown in the Calendar example (section 4.2.6). The format of a timestamp is specified
	// in section 2.2.1.4.
	//
	//	+------------------+----------------------------------------------------------------------------------+
	//	|                  |                                                                                  |
	//	|      VALUE       |                                     MEANING                                      |
	//	|                  |                                                                                  |
	//	+------------------+----------------------------------------------------------------------------------+
	//	+------------------+----------------------------------------------------------------------------------+
	//	| FALSE 0x00000000 | The timestamp of the new resource group MUST specify a time that is equal to     |
	//	|                  | or later than the timestamp of any modifications made to the PMC object on the   |
	//	|                  | server. Otherwise, the modification SHOULD fail, and WRM_ERR_OLD_INFORMATION     |
	//	|                  | SHOULD be returned.                                                              |
	//	+------------------+----------------------------------------------------------------------------------+
	//	| TRUE 0x00000001  | The resource group is validated and modified without checking the timestamp.     |
	//	+------------------+----------------------------------------------------------------------------------+
	Overwrite bool `idl:"name:bOverwrite" json:"overwrite"`
}

func (o *ModifyResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) *xxx_ModifyResourceGroupOperation {
	if op == nil {
		op = &xxx_ModifyResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupInfo = o.ResourceGroupInfo
	op.Overwrite = o.Overwrite
	return op
}

func (o *ModifyResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupInfo = op.ResourceGroupInfo
	o.Overwrite = op.Overwrite
}
func (o *ModifyResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyResourceGroupResponse structure represents the ModifyResourceGroup operation response
type ModifyResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) *xxx_ModifyResourceGroupOperation {
	if op == nil {
		op = &xxx_ModifyResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateResourceGroupOperation structure represents the CreateResourceGroup operation
type xxx_CreateResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateResourceGroupOperation) OpNum() int { return 9 }

func (o *xxx_CreateResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/CreateResourceGroup"
}

func (o *xxx_CreateResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_bstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
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

func (o *xxx_CreateResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_bstrResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateResourceGroupRequest structure represents the CreateResourceGroup operation request
type CreateResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrResourceGroupInfo: A string that specifies a new resource group, in the form
	// of a ProcessMatchingCriteria (section 2.2.5.24) element (section 2.2.5.24). For an
	// example, see ProcessMatchingCriteria example (section 4.2.20).
	ResourceGroupInfo *oaut.String `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
}

func (o *CreateResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) *xxx_CreateResourceGroupOperation {
	if op == nil {
		op = &xxx_CreateResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupInfo = o.ResourceGroupInfo
	return op
}

func (o *CreateResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupInfo = op.ResourceGroupInfo
}
func (o *CreateResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResourceGroupResponse structure represents the CreateResourceGroup operation response
type CreateResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) *xxx_CreateResourceGroupOperation {
	if op == nil {
		op = &xxx_CreateResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteResourceGroupOperation structure represents the DeleteResourceGroup operation
type xxx_DeleteResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteResourceGroupOperation) OpNum() int { return 10 }

func (o *xxx_DeleteResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/DeleteResourceGroup"
}

func (o *xxx_DeleteResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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

func (o *xxx_DeleteResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteResourceGroupRequest structure represents the DeleteResourceGroup operation request
type DeleteResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrResourceGroupName: A string that specifies the name of the resource group that
	// is to be deleted.
	ResourceGroupName *oaut.String `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
}

func (o *DeleteResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) *xxx_DeleteResourceGroupOperation {
	if op == nil {
		op = &xxx_DeleteResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupName = o.ResourceGroupName
	return op
}

func (o *DeleteResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupName = op.ResourceGroupName
}
func (o *DeleteResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResourceGroupResponse structure represents the DeleteResourceGroup operation response
type DeleteResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) *xxx_DeleteResourceGroupOperation {
	if op == nil {
		op = &xxx_DeleteResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameResourceGroupOperation structure represents the RenameResourceGroup operation
type xxx_RenameResourceGroupOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	NewResourceGroupName *oaut.String   `idl:"name:bstrNewResourceGroupName" json:"new_resource_group_name"`
	OldResourceGroupName *oaut.String   `idl:"name:bstrOldResourceGroupName" json:"old_resource_group_name"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameResourceGroupOperation) OpNum() int { return 11 }

func (o *xxx_RenameResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/RenameResourceGroup"
}

func (o *xxx_RenameResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrNewResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewResourceGroupName != nil {
			_ptr_bstrNewResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewResourceGroupName != nil {
					if err := o.NewResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewResourceGroupName, _ptr_bstrNewResourceGroupName); err != nil {
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
	// bstrOldResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldResourceGroupName != nil {
			_ptr_bstrOldResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldResourceGroupName != nil {
					if err := o.OldResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldResourceGroupName, _ptr_bstrOldResourceGroupName); err != nil {
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

func (o *xxx_RenameResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrNewResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewResourceGroupName == nil {
				o.NewResourceGroupName = &oaut.String{}
			}
			if err := o.NewResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewResourceGroupName := func(ptr interface{}) { o.NewResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewResourceGroupName, _s_bstrNewResourceGroupName, _ptr_bstrNewResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrOldResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldResourceGroupName == nil {
				o.OldResourceGroupName = &oaut.String{}
			}
			if err := o.OldResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldResourceGroupName := func(ptr interface{}) { o.OldResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldResourceGroupName, _s_bstrOldResourceGroupName, _ptr_bstrOldResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameResourceGroupRequest structure represents the RenameResourceGroup operation request
type RenameResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrNewResourceGroupName: A string that specifies the new name of the resource group.
	NewResourceGroupName *oaut.String `idl:"name:bstrNewResourceGroupName" json:"new_resource_group_name"`
	// bstrOldResourceGroupName: A string that specifies the name of the resource group
	// to be renamed.
	OldResourceGroupName *oaut.String `idl:"name:bstrOldResourceGroupName" json:"old_resource_group_name"`
}

func (o *RenameResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) *xxx_RenameResourceGroupOperation {
	if op == nil {
		op = &xxx_RenameResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NewResourceGroupName = o.NewResourceGroupName
	op.OldResourceGroupName = o.OldResourceGroupName
	return op
}

func (o *RenameResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NewResourceGroupName = op.NewResourceGroupName
	o.OldResourceGroupName = op.OldResourceGroupName
}
func (o *RenameResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameResourceGroupResponse structure represents the RenameResourceGroup operation response
type RenameResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) *xxx_RenameResourceGroupOperation {
	if op == nil {
		op = &xxx_RenameResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
