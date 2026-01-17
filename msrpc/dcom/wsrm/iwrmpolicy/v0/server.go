package iwrmpolicy

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IWRMPolicy server interface.
type PolicyServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The GetPolicyInfo method gets one or more specified resource allocation policies
	// (RAP).
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------+------------------------------------+
	//	|               RETURN                |                                    |
	//	|             VALUE/CODE              |            DESCRIPTION             |
	//	|                                     |                                    |
	//	+-------------------------------------+------------------------------------+
	//	+-------------------------------------+------------------------------------+
	//	| 0x00000000 S_OK                     | Operation successful.              |
	//	+-------------------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid. |
	//	+-------------------------------------+------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID | The specified RAP does not exist.  |
	//	+-------------------------------------+------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetPolicyInfo(context.Context, *GetPolicyInfoRequest) (*GetPolicyInfoResponse, error)

	// The CreatePolicy method creates a new resource allocation policy (RAP).
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                  | Operation successful.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                          | One or more arguments are invalid.                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                      | The specified name contains characters that are invalid. The name cannot         |
	//	|                                                  | start with a hyphen "-", cannot contain spaces, and cannot contain any of the    |
	//	|                                                  | following characters: \ / ? * | : < > " , ;                                      |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER             | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                                  | be processed.<105>                                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00CA WRM_ERR_POLICYID_ALREADY_EXISTS       | The request has been aborted because a RAP with the specified name already       |
	//	|                                                  | exists.                                                                          |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00DF WRM_ERR_POLICY_LIMIT_EXCEEDED         | The request has been aborted because the total number of RAPs has exceeded an    |
	//	|                                                  | implementation-defined limit.<106>                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E3 WRM_ERR_CANNOT_CREATE_RESERVED_POLICY | A user created policy cannot have the same name as that of a built-in policy.    |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID       | The request has been aborted because the process matching criteria (PMC) name    |
	//	|                                                  | could not be found.                                                              |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)

	// The ModifyPolicy method modifies an existing resource allocation policy (RAP).
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
	//	| 0xC1FF0069 WRM_ERR_OLD_INFORMATION         | The XML timestamp is out of date.                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE                | The specified name contains characters that are invalid. The name cannot start   |
	//	|                                            | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                            | following characters: \ / ? * | : < > " , ;                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER       | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                            | be processed.<107>                                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID        | The specified RAP does not exist.                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E0 WRM_ERR_WSRM_RESERVED_POLICY    | The specified policy is a built-in policy. It cannot be altered.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID | The request has been aborted because the process matching criteria (PMC) name    |
	//	|                                            | could not be found.                                                              |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	ModifyPolicy(context.Context, *ModifyPolicyRequest) (*ModifyPolicyResponse, error)

	// The DeletePolicy method deletes an existing resource policy.
	//
	// );
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                         | Operation successful.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                 | One or more arguments are invalid.                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE             | The specified name contains characters that are invalid. The name cannot start   |
	//	|                                         | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                         | following characters: \ / ? * | : < > " , ;                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00CF WRM_ERR_IS_CURRENT_POLICY    | This resource allocation policy (RAP) is being used by WSRM and cannot be        |
	//	|                                         | deleted.                                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E0 WRM_ERR_WSRM_RESERVED_POLICY | The specified policy is a built-in policy. It cannot be altered.                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E7 WRM_ERR_DELETING_POLICY      | The specified policy could not be deleted. A policy cannot be deleted if it is a |
	//	|                                         | member of one or more conditional policies.                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER    | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                         | be processed. Windows returns this value if the XML data is corrupt.             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID     | The specified RAP does not exist.                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)

	// The RenameAllocationPolicy method renames an existing resource allocation policy
	// (RAP).
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
	//	| 0xC1FF00C8 WRM_ERR_TOO_LONG_POLICY_ID      | The request has been aborted, because the RAP name has exceeded an               |
	//	|                                            | implementation-defined limit.<108>                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID        | The specified RAP does not exist.                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00CA WRM_ERR_POLICYID_ALREADY_EXISTS | A RAP with the specified name already exists.                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00CB WRM_ERR_POLICYID_RESERVED_WORD  | The specified RAP name is a reserved word used by WSRM and cannot be used as     |
	//	|                                            | a name. Reserved words for RAPs are "current", "none", "\" and "Residual".       |
	//	|                                            | Reserved words are case-insensitive.                                             |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E0 WRM_ERR_WSRM_RESERVED_POLICY    | The specified policy is a built-in policy. It cannot be altered.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER       | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                            | be processed.<109>                                                               |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00E5 WRM_ERR_RENAME_ACTIVE_POLICY    | The specified RAP is being used by WSRM and cannot be renamed.                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	RenameAllocationPolicy(context.Context, *RenameAllocationPolicyRequest) (*RenameAllocationPolicyResponse, error)

	// The MoveBefore method moves a specified resource group to a location just before
	// a reference resource group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                  |
	//	|                 VALUE/CODE                 |                           DESCRIPTION                            |
	//	|                                            |                                                                  |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | Operation successful.                                            |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID | The specified process matching criteria (PMC) does not exist.    |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID        | The specified resource allocation policy does not exist.         |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF00E0 WRM_ERR_WSRM_RESERVED_POLICY    | The specified policy is a built-in policy. It cannot be altered. |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                    | One or more arguments are invalid.                               |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	MoveBefore(context.Context, *MoveBeforeRequest) (*MoveBeforeResponse, error)

	// The MoveAfter method moves a specified resource group to a location just after a
	// reference resource group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                  |
	//	|                 VALUE/CODE                 |                           DESCRIPTION                            |
	//	|                                            |                                                                  |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | Operation successful.                                            |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                    | One or more arguments are invalid.                               |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF00E0 WRM_ERR_WSRM_RESERVED_POLICY    | The specified policy is a built-in policy. It cannot be altered. |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID        | The specified resource allocation policy does not exist.         |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//	| 0xC1FF012E WRM_ERR_RESOURCEGROUPID_INVALID | The specified process matching criteria (PMC) does not exist.    |
	//	+--------------------------------------------+------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	MoveAfter(context.Context, *MoveAfterRequest) (*MoveAfterResponse, error)

	// The SetCalDefaultPolicyName method stores the name of the default resource allocation
	// policy (RAP) in the registry.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | Operation successful.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID | The request has been aborted because the specified resource allocation policy    |
	//	|                                     | does not exist.                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE         | The specified name contains characters that are invalid. The name cannot start   |
	//	|                                     | with a hyphen ("-"), cannot contain spaces, and cannot contain any of the        |
	//	|                                     | following characters: \ / ? * | : < > " , ;                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	SetCALDefaultPolicyName(context.Context, *SetCALDefaultPolicyNameRequest) (*SetCALDefaultPolicyNameResponse, error)

	// The GetCalDefaultPolicyName method used to get the name of the default resource allocation
	// policy (RAP).
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------+--------------------------------------------------------------------------+
	//	|          RETURN           |                                                                          |
	//	|        VALUE/CODE         |                               DESCRIPTION                                |
	//	|                           |                                                                          |
	//	+---------------------------+--------------------------------------------------------------------------+
	//	+---------------------------+--------------------------------------------------------------------------+
	//	| 0xC1FF038F WRM_NO_DEF_RAP | The WSRM server does not have a default resource allocation policy.<110> |
	//	+---------------------------+--------------------------------------------------------------------------+
	//	| 0x00000000 S_OK           | Operation successful.                                                    |
	//	+---------------------------+--------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetCALDefaultPolicyName(context.Context, *GetCALDefaultPolicyNameRequest) (*GetCALDefaultPolicyNameResponse, error)

	// The GetProcessList method returns a list of processes for a specified policy.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | Operation successful.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE         | The specified name contains characters that are not valid. The name cannot       |
	//	|                                     | start with a hyphen ("-"), cannot contain spaces, and cannot contain any of the  |
	//	|                                     | following characters: \ / ? * | : < > " , ;                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID | The specified resource allocation policy does not exist.                         |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetProcessList(context.Context, *GetProcessListRequest) (*GetProcessListResponse, error)

	// The GetCurrentPolicy method returns the current resource policy.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetCurrentPolicy(context.Context, *GetCurrentPolicyRequest) (*GetCurrentPolicyResponse, error)

	// The SetCurrentPolicy method sets the current resource policy to a specified resource
	// policy by name.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                     | Operation successful.                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF006F WRM_ERR_ID_VALUE         | The specified name contains characters that are not valid. The name cannot       |
	//	|                                     | start with a hyphen ("-"), cannot contain spaces, and cannot contain any of the  |
	//	|                                     | following characters: \ / ? * | : < > " , ;                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG             | One or more arguments are invalid.                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF00C9 WRM_ERR_POLICYID_INVALID | The specified RAP does not exist.                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	SetCurrentPolicy(context.Context, *SetCurrentPolicyRequest) (*SetCurrentPolicyResponse, error)

	// The GetCurrentStateAndActivePolicyName method returns the name of the current resource
	// policy and the management state in which the WSRM management service is currently
	// available.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetCurrentStateAndActivePolicyName(context.Context, *GetCurrentStateAndActivePolicyNameRequest) (*GetCurrentStateAndActivePolicyNameResponse, error)

	// The GetConditionalPolicy function returns conditions for a specified conditional
	// policy.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                              |
	//	|       VALUE/CODE        |                                 DESCRIPTION                                  |
	//	|                         |                                                                              |
	//	+-------------------------+------------------------------------------------------------------------------+
	//	+-------------------------+------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.                                                        |
	//	+-------------------------+------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid or the specified policy name is not found. |
	//	+-------------------------+------------------------------------------------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	GetConditionalPolicy(context.Context, *GetConditionalPolicyRequest) (*GetConditionalPolicyResponse, error)

	// The SetConditionalPolicy method loads specified conditions into the conditional policy,
	// which contains the conditions and the respective SwitchToPolicy element values. When
	// a specified condition occurs and the management mode is MANUAL_ACTIVE_POLICY, WSRM
	// MUST start managing resources using the policy specified by the corresponding SwitchToPolicy
	// element value.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | Operation successful.              |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Additional IWRMPolicy interface methods are specified in section 3.2.4.7.
	SetConditionalPolicy(context.Context, *SetConditionalPolicyRequest) (*SetConditionalPolicyResponse, error)
}

func RegisterPolicyServer(conn dcerpc.Conn, o PolicyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPolicyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PolicySyntaxV0_0))...)
}

func NewPolicyServerHandle(o PolicyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PolicyServerHandle(ctx, o, opNum, r)
	}
}

func PolicyServerHandle(ctx context.Context, o PolicyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetPolicyInfo
		op := &xxx_GetPolicyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPolicyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPolicyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreatePolicy
		op := &xxx_CreatePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyPolicy
		op := &xxx_ModifyPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeletePolicy
		op := &xxx_DeletePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameAllocationPolicy
		op := &xxx_RenameAllocationPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameAllocationPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameAllocationPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // MoveBefore
		op := &xxx_MoveBeforeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveBeforeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveBefore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // MoveAfter
		op := &xxx_MoveAfterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveAfterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveAfter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SetCalDefaultPolicyName
		op := &xxx_SetCALDefaultPolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCALDefaultPolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCALDefaultPolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetCalDefaultPolicyName
		op := &xxx_GetCALDefaultPolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCALDefaultPolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCALDefaultPolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetProcessList
		op := &xxx_GetProcessListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProcessListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProcessList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // GetCurrentPolicy
		op := &xxx_GetCurrentPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetCurrentPolicy
		op := &xxx_SetCurrentPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCurrentPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCurrentPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // GetCurrentStateAndActivePolicyName
		op := &xxx_GetCurrentStateAndActivePolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentStateAndActivePolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentStateAndActivePolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // GetConditionalPolicy
		op := &xxx_GetConditionalPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConditionalPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConditionalPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // SetConditionalPolicy
		op := &xxx_SetConditionalPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConditionalPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConditionalPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMPolicy
type UnimplementedPolicyServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPolicyServer) GetPolicyInfo(context.Context, *GetPolicyInfoRequest) (*GetPolicyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) ModifyPolicy(context.Context, *ModifyPolicyRequest) (*ModifyPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) RenameAllocationPolicy(context.Context, *RenameAllocationPolicyRequest) (*RenameAllocationPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) MoveBefore(context.Context, *MoveBeforeRequest) (*MoveBeforeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) MoveAfter(context.Context, *MoveAfterRequest) (*MoveAfterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetCALDefaultPolicyName(context.Context, *SetCALDefaultPolicyNameRequest) (*SetCALDefaultPolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCALDefaultPolicyName(context.Context, *GetCALDefaultPolicyNameRequest) (*GetCALDefaultPolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetProcessList(context.Context, *GetProcessListRequest) (*GetProcessListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCurrentPolicy(context.Context, *GetCurrentPolicyRequest) (*GetCurrentPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetCurrentPolicy(context.Context, *SetCurrentPolicyRequest) (*SetCurrentPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCurrentStateAndActivePolicyName(context.Context, *GetCurrentStateAndActivePolicyNameRequest) (*GetCurrentStateAndActivePolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetConditionalPolicy(context.Context, *GetConditionalPolicyRequest) (*GetConditionalPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetConditionalPolicy(context.Context, *SetConditionalPolicyRequest) (*SetConditionalPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PolicyServer = (*UnimplementedPolicyServer)(nil)
