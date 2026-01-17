package iwrmmachinegroup

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

// IWRMMachineGroup server interface.
type MachineGroupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The CreateMachineGroup method creates and initializes a machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                           RETURN                           |                                                                                  |
	//	|                         VALUE/CODE                         |                                   DESCRIPTION                                    |
	//	|                                                            |                                                                                  |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                            | Operation successful.                                                            |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                    | One or more arguments are invalid.                                               |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER                       | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                                            | be processed.<86>                                                                |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0385 WRM_ERR_MACHINE_GROUP_LIMIT_EXCEEDED            | The total number of machine groups has exceeded an implementation-defined        |
	//	|                                                            | limit.<87>                                                                       |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0387 WRM_ERR_MACHINES_LIMIT_IN_MACHINEGROUP_EXCEEDED | The total number of machines directly under a machine group has exceeded an      |
	//	|                                                            | implementation-defined limit.<88>                                                |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0388 WRM_ERR_MACHINEGROUP_ALREADY_EXISTS             | A machine group with the specified name in bstrMachineGroupInfo XML already      |
	//	|                                                            | exists in the entire WSRM configuration.                                         |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID                  | The specified parent machine group id does not exist.                            |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	CreateMachineGroup(context.Context, *CreateMachineGroupRequest) (*CreateMachineGroupResponse, error)

	// The GetMachineGroupInfo method returns information about a machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+--------------------------------------------+
	//	|                  RETURN                   |                                            |
	//	|                VALUE/CODE                 |                DESCRIPTION                 |
	//	|                                           |                                            |
	//	+-------------------------------------------+--------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                      |
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.         |
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID | The specified machine group id is invalid. |
	//	+-------------------------------------------+--------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	GetMachineGroupInfo(context.Context, *GetMachineGroupInfoRequest) (*GetMachineGroupInfoResponse, error)

	// The ModifyMachineGroup method modifies an existing machine group. The method replaces
	// or merges the machine group information according to the value specified in the enumMGMergeOptions
	// member.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                           RETURN                           |                                                                                  |
	//	|                         VALUE/CODE                         |                                   DESCRIPTION                                    |
	//	|                                                            |                                                                                  |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                            | Operation successful.                                                            |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                    | One or more arguments are invalid.<90>                                           |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER                       | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                                            | be processed.<91>                                                                |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0385 WRM_ERR_MACHINE_GROUP_LIMIT_EXCEEDED            | The total number of machine groups as specified in bstrMachineGroupInfo, has     |
	//	|                                                            | exceeded an implementation-defined limit.<92>                                    |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0387 WRM_ERR_MACHINES_LIMIT_IN_MACHINEGROUP_EXCEEDED | The machine group information could not be modified because the total number of  |
	//	|                                                            | machines directly under a machine group has exceeded an implementation-defined   |
	//	|                                                            | limit.<93>                                                                       |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0388 WRM_ERR_MACHINEGROUP_ALREADY_EXISTS             | A machine group with the specified name in bstrMachineGroupInfo XML already      |
	//	|                                                            | exists in the entire WSRM configuration. For example, if ModifyMachineGroup is   |
	//	|                                                            | used to modify a machine group ID that is identical to the existing group ID,    |
	//	|                                                            | this error will be generated.<94>                                                |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID                  | The specified machine group ID is invalid.                                       |
	//	+------------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	ModifyMachineGroup(context.Context, *ModifyMachineGroupRequest) (*ModifyMachineGroupResponse, error)

	// The DeleteMachineGroup method deletes an existing machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+--------------------------------------------+
	//	|                  RETURN                   |                                            |
	//	|                VALUE/CODE                 |                DESCRIPTION                 |
	//	|                                           |                                            |
	//	+-------------------------------------------+--------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                      |
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.         |
	//	+-------------------------------------------+--------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID | The specified machine group id is invalid. |
	//	+-------------------------------------------+--------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	DeleteMachineGroup(context.Context, *DeleteMachineGroupRequest) (*DeleteMachineGroupResponse, error)

	// The RenameMachineGroup method changes the name of an existing machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                | Operation successful.                                                            |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                        | One or more arguments are invalid.                                               |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0388 WRM_ERR_MACHINEGROUP_ALREADY_EXISTS | A machine group with the specified name already exists in the entire WSRM        |
	//	|                                                | configuration.                                                                   |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID      | The specified machine group id is invalid.                                       |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	RenameMachineGroup(context.Context, *RenameMachineGroupRequest) (*RenameMachineGroupResponse, error)

	// The AddMachine method adds a machine to a machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER      | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                           | be processed.<96>                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID | The specified machine group id is invalid.                                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF038B WRM_ERR_MACHINE_LIMIT_EXCEEDED | The total number of machines has exceeded an implementation-defined limit.<97>   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF038C WRM_ERR_MACHINE_ALREADY_EXISTS | A machine with the specified name already exists as the direct child in the      |
	//	|                                           | machine group.                                                                   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	AddMachine(context.Context, *AddMachineRequest) (*AddMachineResponse, error)

	// The GetMachineInfo method returns information about a machine. If more than one machine
	// with the specified bstrMachineId value exists in the hierarchy of machine groups,
	// the information can be returned from any of the machine groups.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation successful.                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG              | One or more arguments are invalid.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF038A WRM_ERR_MACHINEID_INVALID | The specified machine ID is not found in any machine group of the                |
	//	|                                      | configuration.<98>                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	GetMachineInfo(context.Context, *GetMachineInfoRequest) (*GetMachineInfoResponse, error)

	// The ModifyMachineInfo method modifies a machine in a machine group. This method modifies
	// the direct child machine with the specified machine ID under the specified parent
	// machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0070 WRM_ERR_TAGS_NOT_IN_ORDER      | The XML data that is maintained by the management service is invalid or cannot   |
	//	|                                           | be processed.<100>                                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID | The specified machine group ID is invalid.                                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF038A WRM_ERR_MACHINEID_INVALID      | The specified machine ID is not found.<101>                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF038C WRM_ERR_MACHINE_ALREADY_EXISTS | A machine with the specified machine name in the bstrMachineInfo XML already     |
	//	|                                           | exists as the direct child in the machine group.<102>                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	ModifyMachineInfo(context.Context, *ModifyMachineInfoRequest) (*ModifyMachineInfoResponse, error)

	// The DeleteMachine method deletes a machine from a machine group.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+---------------------------------------------+
	//	|                  RETURN                   |                                             |
	//	|                VALUE/CODE                 |                 DESCRIPTION                 |
	//	|                                           |                                             |
	//	+-------------------------------------------+---------------------------------------------+
	//	+-------------------------------------------+---------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                       |
	//	+-------------------------------------------+---------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.          |
	//	+-------------------------------------------+---------------------------------------------+
	//	| 0xC1FF0389 WRM_ERR_MACHINEGROUPID_INVALID | The specified machine group ID is invalid.  |
	//	+-------------------------------------------+---------------------------------------------+
	//	| 0xC1FF038A WRM_ERR_MACHINEID_INVALID      | The specified machine ID is not found.<104> |
	//	+-------------------------------------------+---------------------------------------------+
	//
	// Additional IWRMMachineGroup interface methods are specified in section 3.2.4.6.
	DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error)
}

func RegisterMachineGroupServer(conn dcerpc.Conn, o MachineGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMachineGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MachineGroupSyntaxV0_0))...)
}

func NewMachineGroupServerHandle(o MachineGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MachineGroupServerHandle(ctx, o, opNum, r)
	}
}

func MachineGroupServerHandle(ctx context.Context, o MachineGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateMachineGroup
		op := &xxx_CreateMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetMachineGroupInfo
		op := &xxx_GetMachineGroupInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineGroupInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachineGroupInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyMachineGroup
		op := &xxx_ModifyMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteMachineGroup
		op := &xxx_DeleteMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameMachineGroup
		op := &xxx_RenameMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // AddMachine
		op := &xxx_AddMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetMachineInfo
		op := &xxx_GetMachineInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachineInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ModifyMachineInfo
		op := &xxx_ModifyMachineInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyMachineInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyMachineInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // DeleteMachine
		op := &xxx_DeleteMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMMachineGroup
type UnimplementedMachineGroupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedMachineGroupServer) CreateMachineGroup(context.Context, *CreateMachineGroupRequest) (*CreateMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) GetMachineGroupInfo(context.Context, *GetMachineGroupInfoRequest) (*GetMachineGroupInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) ModifyMachineGroup(context.Context, *ModifyMachineGroupRequest) (*ModifyMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) DeleteMachineGroup(context.Context, *DeleteMachineGroupRequest) (*DeleteMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) RenameMachineGroup(context.Context, *RenameMachineGroupRequest) (*RenameMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) AddMachine(context.Context, *AddMachineRequest) (*AddMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) GetMachineInfo(context.Context, *GetMachineInfoRequest) (*GetMachineInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) ModifyMachineInfo(context.Context, *ModifyMachineInfoRequest) (*ModifyMachineInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MachineGroupServer = (*UnimplementedMachineGroupServer)(nil)
