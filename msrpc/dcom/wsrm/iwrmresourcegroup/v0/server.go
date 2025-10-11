package iwrmresourcegroup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = idispatch.GoPackage
)

// IWRMResourceGroup server interface.
type ResourceGroupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

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
	GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest) (*GetResourceGroupInfoResponse, error)

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
	ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest) (*ModifyResourceGroupResponse, error)

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
	CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error)

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
	DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest) (*DeleteResourceGroupResponse, error)

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
	RenameResourceGroup(context.Context, *RenameResourceGroupRequest) (*RenameResourceGroupResponse, error)
}

func RegisterResourceGroupServer(conn dcerpc.Conn, o ResourceGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewResourceGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ResourceGroupSyntaxV0_0))...)
}

func NewResourceGroupServerHandle(o ResourceGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ResourceGroupServerHandle(ctx, o, opNum, r)
	}
}

func ResourceGroupServerHandle(ctx context.Context, o ResourceGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetResourceGroupInfo
		op := &xxx_GetResourceGroupInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceGroupInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceGroupInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ModifyResourceGroup
		op := &xxx_ModifyResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CreateResourceGroup
		op := &xxx_CreateResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteResourceGroup
		op := &xxx_DeleteResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameResourceGroup
		op := &xxx_RenameResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMResourceGroup
type UnimplementedResourceGroupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedResourceGroupServer) GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest) (*GetResourceGroupInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceGroupServer) ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest) (*ModifyResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceGroupServer) CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceGroupServer) DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest) (*DeleteResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedResourceGroupServer) RenameResourceGroup(context.Context, *RenameResourceGroupRequest) (*RenameResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ResourceGroupServer = (*UnimplementedResourceGroupServer)(nil)
