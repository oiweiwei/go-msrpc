package dssetup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dssp"
)

var (
	// Syntax UUID
	DssetupSyntaxUUID = &uuid.UUID{TimeLow: 0x3919286a, TimeMid: 0xb10c, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0xa8, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x2e, 0xf5}}
	// Syntax ID
	DssetupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DssetupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// dssetup interface.
type DssetupClient interface {

	// The DsRolerGetPrimaryDomainInformation (Opnum 0) method returns the requested information
	// about the current configuration or state of the computer on which the server is running.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code as specified in [MS-ERREF]. Specifically, in addition to any other error
	// codes, the server MUST return the following error codes for the following error conditions.
	// Any other values transmitted in this field are implementation-specific. All nonzero
	// values MUST be treated the same for protocol purposes.
	//
	//	+------------------------------------+---------------------------------------+
	//	|               RETURN               |                                       |
	//	|             VALUE/CODE             |              DESCRIPTION              |
	//	|                                    |                                       |
	//	+------------------------------------+---------------------------------------+
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One or more parameters are invalid.   |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | A memory allocation failure occurred. |
	//	+------------------------------------+---------------------------------------+
	GetPrimaryDomainInformation(context.Context, *GetPrimaryDomainInformationRequest, ...dcerpc.CallOption) (*GetPrimaryDomainInformationResponse, error)

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MachineRole type represents DSROLE_MACHINE_ROLE RPC enumeration.
//
// The DSROLE_MACHINE_ROLE enumeration specifies the current role of the computer.
type MachineRole uint16

var (
	// DsRole_RoleStandaloneWorkstation:  The computer is a stand-alone workstation.
	MachineRoleStandaloneWorkstation MachineRole = 0
	// DsRole_RoleMemberWorkstation:  The computer is a workstation that is joined to a
	// domain.
	MachineRoleMemberWorkstation MachineRole = 1
	// DsRole_RoleStandaloneServer:  The computer is a stand-alone server.
	MachineRoleStandaloneServer MachineRole = 2
	// DsRole_RoleMemberServer:  The computer is a server that is joined to a domain.
	MachineRoleMemberServer MachineRole = 3
	// DsRole_RoleBackupDomainController:  The computer is a server that is a backup domain
	// controller or a read-only domain controller.<3>
	MachineRoleBackupDomainController MachineRole = 4
	// DsRole_RolePrimaryDomainController:  The computer is a server that is the primary
	// domain controller emulator.
	MachineRolePrimaryDomainController MachineRole = 5
)

func (o MachineRole) String() string {
	switch o {
	case MachineRoleStandaloneWorkstation:
		return "MachineRoleStandaloneWorkstation"
	case MachineRoleMemberWorkstation:
		return "MachineRoleMemberWorkstation"
	case MachineRoleStandaloneServer:
		return "MachineRoleStandaloneServer"
	case MachineRoleMemberServer:
		return "MachineRoleMemberServer"
	case MachineRoleBackupDomainController:
		return "MachineRoleBackupDomainController"
	case MachineRolePrimaryDomainController:
		return "MachineRolePrimaryDomainController"
	}
	return "Invalid"
}

// ServerState type represents DSROLE_SERVER_STATE RPC enumeration.
//
// The DSROLE_SERVER_STATE enumeration specifies the role of the computer prior to the
// upgrade.
type ServerState uint16

var (
	// DsRoleServerUnknown:  The previous role of the computer is unknown.
	ServerStateUnknown ServerState = 0
	// DsRoleServerPrimary:  The previous role of the computer was primary domain controller
	// in a legacy domain.
	ServerStatePrimary ServerState = 1
	// DsRoleServerBackup:  The previous role of the computer was backup domain controller
	// in a legacy domain.
	ServerStateBackup ServerState = 2
)

func (o ServerState) String() string {
	switch o {
	case ServerStateUnknown:
		return "ServerStateUnknown"
	case ServerStatePrimary:
		return "ServerStatePrimary"
	case ServerStateBackup:
		return "ServerStateBackup"
	}
	return "Invalid"
}

// PrimaryDomainInfoLevel type represents DSROLE_PRIMARY_DOMAIN_INFO_LEVEL RPC enumeration.
//
// The DSROLE_PRIMARY_DOMAIN_INFO_LEVEL enumeration defines the information level that
// the client requests.
type PrimaryDomainInfoLevel uint16

var (
	// DsRolePrimaryDomainInfoBasic:  Request for information about the domain to which
	// the computer belongs.
	PrimaryDomainInfoLevelBasic PrimaryDomainInfoLevel = 1
	// DsRoleUpgradeStatus:  Request for computer operating system upgrade status.
	PrimaryDomainInfoLevelUpgradeStatus PrimaryDomainInfoLevel = 2
	// DsRoleOperationState:  Request for computer operation state.
	PrimaryDomainInfoLevelOperationState PrimaryDomainInfoLevel = 3
)

func (o PrimaryDomainInfoLevel) String() string {
	switch o {
	case PrimaryDomainInfoLevelBasic:
		return "PrimaryDomainInfoLevelBasic"
	case PrimaryDomainInfoLevelUpgradeStatus:
		return "PrimaryDomainInfoLevelUpgradeStatus"
	case PrimaryDomainInfoLevelOperationState:
		return "PrimaryDomainInfoLevelOperationState"
	}
	return "Invalid"
}

// UpgradeStatusInfo structure represents DSROLE_UPGRADE_STATUS_INFO RPC structure.
//
// The DSROLE_UPGRADE_STATUS_INFO structure contains information about the status of
// a pending operating system upgrade, if any, for the computer. This structure is intended
// to store only the status of an operating system upgrade of a legacy domain controller.
type UpgradeStatusInfo struct {
	// OperationState:  The current status of the upgrade. Valid values are shown in the
	// following table.<5>
	//
	//	+---------------------------------------+--------------------------------------+
	//	|                                       |                                      |
	//	|                 VALUE                 |               MEANING                |
	//	|                                       |                                      |
	//	+---------------------------------------+--------------------------------------+
	//	+---------------------------------------+--------------------------------------+
	//	| 0x00000000                            | No upgrade is currently in progress. |
	//	+---------------------------------------+--------------------------------------+
	//	| DSROLE_UPGRADE_IN_PROGRESS 0x00000004 | An upgrade is currently in progress. |
	//	+---------------------------------------+--------------------------------------+
	OperationState uint32 `idl:"name:OperationState" json:"operation_state"`
	// PreviousServerState:  The role of the computer prior to the upgrade. The value of
	// this member is valid only if an upgrade is in progress (that is, if the OperationState
	// member is set to DSROLE_UPGRADE_IN_PROGRESS).
	PreviousServerState ServerState `idl:"name:PreviousServerState" json:"previous_server_state"`
}

func (o *UpgradeStatusInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UpgradeStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationState); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.PreviousServerState)); err != nil {
		return err
	}
	return nil
}
func (o *UpgradeStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationState); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.PreviousServerState)); err != nil {
		return err
	}
	return nil
}

// OperationState type represents DSROLE_OPERATION_STATE RPC enumeration.
//
// The DSROLE_OPERATION_STATE enumeration specifies values that determine whether a
// DC promotion or demotion operation is currently being performed on a computer.<4>
type OperationState uint16

var (
	// DsRoleOperationIdle:  No promotion or demotion operation is currently being performed
	// on the computer.
	OperationStateIdle OperationState = 0
	// DsRoleOperationActive:  A promotion or demotion operation is in progress.
	OperationStateActive OperationState = 1
	// DsRoleOperationNeedReboot:  A promotion or demotion operation has been performed.
	// The computer MUST be restarted to function in the new role.
	OperationStateNeedReboot OperationState = 2
)

func (o OperationState) String() string {
	switch o {
	case OperationStateIdle:
		return "OperationStateIdle"
	case OperationStateActive:
		return "OperationStateActive"
	case OperationStateNeedReboot:
		return "OperationStateNeedReboot"
	}
	return "Invalid"
}

// OperationStateInfo structure represents DSROLE_OPERATION_STATE_INFO RPC structure.
//
// The DSROLE_OPERATION_STATE_INFO structure contains the status of a pending domain
// controller (DC) domain membership role change operation, if any, for the computer.
type OperationStateInfo struct {
	// OperationState:   The domain membership role change status of the computer, as specified
	// by a DSROLE_OPERATION_STATE enumeration.
	OperationState OperationState `idl:"name:OperationState" json:"operation_state"`
}

func (o *OperationStateInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OperationStateInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.OperationState)); err != nil {
		return err
	}
	return nil
}
func (o *OperationStateInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.OperationState)); err != nil {
		return err
	}
	return nil
}

// PrimaryDomainInfoBasic structure represents DSROLER_PRIMARY_DOMAIN_INFO_BASIC RPC structure.
//
// The DSROLER_PRIMARY_DOMAIN_INFO_BASIC structure contains basic information, including
// the role of the computer, domain name, and GUID of the domain.
type PrimaryDomainInfoBasic struct {
	// MachineRole:  The current role of the computer, expressed as a DSROLE_MACHINE_ROLE
	// data type.
	MachineRole MachineRole `idl:"name:MachineRole" json:"machine_role"`
	// Flags:  The value that indicates the state of the directory service and validity
	// of the information contained in the DomainGuid member. The value of this parameter
	// MUST be zero or a combination of one or more individual flags in the following table.
	// The combination is the result of a bitwise OR of the flags that apply to the computer
	// for which information is being retrieved. All undefined bits MUST be 0.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DSROLE_PRIMARY_DS_RUNNING 0x00000001          | The directory service is running on this computer. If this flag is not set, the  |
	//	|                                               | directory service is not running on this computer.                               |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DSROLE_PRIMARY_DS_MIXED_MODE 0x00000002       | The directory service is running in mixed mode. This flag is valid only if the   |
	//	|                                               | DSROLE_PRIMARY_DS_RUNNING flag is set and the DSROLE_PRIMARY_DS_READONLY flag is |
	//	|                                               | not set.                                                                         |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DSROLE_PRIMARY_DS_READONLY 0x00000008         | The computer holds a read-only copy of the directory. This flag is valid only if |
	//	|                                               | the DSROLE_PRIMARY_DS_RUNNING flag is set and the DSROLE_PRIMARY_DS_MIXED_MODE   |
	//	|                                               | flag is not set.                                                                 |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DSROLE_PRIMARY_DOMAIN_GUID_PRESENT 0x01000000 | The DomainGuid member contains a valid domain GUID. If this bit is not set, the  |
	//	|                                               | value in DomainGuid member is undefined.                                         |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// DomainNameFlat:  The NetBIOS name of the domain or non-domain workgroup to which
	// the computer belongs.
	DomainNameFlat string `idl:"name:DomainNameFlat;string;pointer:unique" json:"domain_name_flat"`
	// DomainNameDns:  The domain name of the computer. This member MUST be NULL if the
	// MachineRole member is DsRole_RoleStandaloneWorkstation or DsRole_RoleStandaloneServer
	// and MUST NOT be NULL otherwise.
	DomainNameDNS string `idl:"name:DomainNameDns;string;pointer:unique" json:"domain_name_dns"`
	// DomainForestName:  The name of the forest to which the computer belongs. This member
	// MUST be NULL, if the computer is a stand-alone workstation or server.
	DomainForestName string `idl:"name:DomainForestName;string;pointer:unique" json:"domain_forest_name"`
	// DomainGuid:   The UUID of the domain to which the computer belongs. The value of
	// this member is valid only if the DSROLE_PRIMARY_DOMAIN_GUID_PRESENT flag is set.
	DomainGUID *dtyp.GUID `idl:"name:DomainGuid" json:"domain_guid"`
}

func (o *PrimaryDomainInfoBasic) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrimaryDomainInfoBasic) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.MachineRole)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.DomainNameFlat != "" {
		_ptr_DomainNameFlat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DomainNameFlat); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainNameFlat, _ptr_DomainNameFlat); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainNameDNS != "" {
		_ptr_DomainNameDns := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DomainNameDNS); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainNameDNS, _ptr_DomainNameDns); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainForestName != "" {
		_ptr_DomainForestName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DomainForestName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainForestName, _ptr_DomainForestName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainGUID != nil {
		if err := o.DomainGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrimaryDomainInfoBasic) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.MachineRole)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_DomainNameFlat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainNameFlat); err != nil {
			return err
		}
		return nil
	})
	_s_DomainNameFlat := func(ptr interface{}) { o.DomainNameFlat = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainNameFlat, _s_DomainNameFlat, _ptr_DomainNameFlat); err != nil {
		return err
	}
	_ptr_DomainNameDns := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainNameDNS); err != nil {
			return err
		}
		return nil
	})
	_s_DomainNameDns := func(ptr interface{}) { o.DomainNameDNS = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainNameDNS, _s_DomainNameDns, _ptr_DomainNameDns); err != nil {
		return err
	}
	_ptr_DomainForestName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainForestName); err != nil {
			return err
		}
		return nil
	})
	_s_DomainForestName := func(ptr interface{}) { o.DomainForestName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainForestName, _s_DomainForestName, _ptr_DomainForestName); err != nil {
		return err
	}
	if o.DomainGUID == nil {
		o.DomainGUID = &dtyp.GUID{}
	}
	if err := o.DomainGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PrimaryDomainInformation structure represents DSROLER_PRIMARY_DOMAIN_INFORMATION RPC union.
type PrimaryDomainInformation struct {
	// Types that are assignable to Value
	//
	// *PrimaryDomainInformation_DomainInfoBasic
	// *PrimaryDomainInformation_UpgradeStatusInfo
	// *PrimaryDomainInformation_OperationStateInfo
	Value is_PrimaryDomainInformation `json:"value"`
}

func (o *PrimaryDomainInformation) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PrimaryDomainInformation_DomainInfoBasic:
		if value != nil {
			return value.DomainInfoBasic
		}
	case *PrimaryDomainInformation_UpgradeStatusInfo:
		if value != nil {
			return value.UpgradeStatusInfo
		}
	case *PrimaryDomainInformation_OperationStateInfo:
		if value != nil {
			return value.OperationStateInfo
		}
	}
	return nil
}

type is_PrimaryDomainInformation interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PrimaryDomainInformation()
}

func (o *PrimaryDomainInformation) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PrimaryDomainInformation_DomainInfoBasic:
		return uint16(1)
	case *PrimaryDomainInformation_UpgradeStatusInfo:
		return uint16(2)
	case *PrimaryDomainInformation_OperationStateInfo:
		return uint16(3)
	}
	return uint16(0)
}

func (o *PrimaryDomainInformation) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*PrimaryDomainInformation_DomainInfoBasic)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PrimaryDomainInformation_DomainInfoBasic{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*PrimaryDomainInformation_UpgradeStatusInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PrimaryDomainInformation_UpgradeStatusInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*PrimaryDomainInformation_OperationStateInfo)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PrimaryDomainInformation_OperationStateInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PrimaryDomainInformation) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &PrimaryDomainInformation_DomainInfoBasic{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &PrimaryDomainInformation_UpgradeStatusInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &PrimaryDomainInformation_OperationStateInfo{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PrimaryDomainInformation_DomainInfoBasic structure represents DSROLER_PRIMARY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 1
type PrimaryDomainInformation_DomainInfoBasic struct {
	DomainInfoBasic *PrimaryDomainInfoBasic `idl:"name:DomainInfoBasic" json:"domain_info_basic"`
}

func (*PrimaryDomainInformation_DomainInfoBasic) is_PrimaryDomainInformation() {}

func (o *PrimaryDomainInformation_DomainInfoBasic) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DomainInfoBasic != nil {
		if err := o.DomainInfoBasic.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PrimaryDomainInfoBasic{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrimaryDomainInformation_DomainInfoBasic) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DomainInfoBasic == nil {
		o.DomainInfoBasic = &PrimaryDomainInfoBasic{}
	}
	if err := o.DomainInfoBasic.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PrimaryDomainInformation_UpgradeStatusInfo structure represents DSROLER_PRIMARY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 2
type PrimaryDomainInformation_UpgradeStatusInfo struct {
	UpgradeStatusInfo *UpgradeStatusInfo `idl:"name:UpgradStatusInfo" json:"upgrade_status_info"`
}

func (*PrimaryDomainInformation_UpgradeStatusInfo) is_PrimaryDomainInformation() {}

func (o *PrimaryDomainInformation_UpgradeStatusInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UpgradeStatusInfo != nil {
		if err := o.UpgradeStatusInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UpgradeStatusInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrimaryDomainInformation_UpgradeStatusInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UpgradeStatusInfo == nil {
		o.UpgradeStatusInfo = &UpgradeStatusInfo{}
	}
	if err := o.UpgradeStatusInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PrimaryDomainInformation_OperationStateInfo structure represents DSROLER_PRIMARY_DOMAIN_INFORMATION RPC union arm.
//
// It has following labels: 3
type PrimaryDomainInformation_OperationStateInfo struct {
	OperationStateInfo *OperationStateInfo `idl:"name:OperationStateInfo" json:"operation_state_info"`
}

func (*PrimaryDomainInformation_OperationStateInfo) is_PrimaryDomainInformation() {}

func (o *PrimaryDomainInformation_OperationStateInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OperationStateInfo != nil {
		if err := o.OperationStateInfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OperationStateInfo{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PrimaryDomainInformation_OperationStateInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OperationStateInfo == nil {
		o.OperationStateInfo = &OperationStateInfo{}
	}
	if err := o.OperationStateInfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultDssetupClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDssetupClient) GetPrimaryDomainInformation(ctx context.Context, in *GetPrimaryDomainInformationRequest, opts ...dcerpc.CallOption) (*GetPrimaryDomainInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPrimaryDomainInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDssetupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDssetupClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDssetupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DssetupClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DssetupSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDssetupClient{cc: cc}, nil
}

// xxx_GetPrimaryDomainInformationOperation structure represents the DsRolerGetPrimaryDomainInformation operation
type xxx_GetPrimaryDomainInformationOperation struct {
	InfoLevel  PrimaryDomainInfoLevel    `idl:"name:InfoLevel" json:"info_level"`
	DomainInfo *PrimaryDomainInformation `idl:"name:DomainInfo;switch_is:InfoLevel" json:"domain_info"`
	Return     uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPrimaryDomainInformationOperation) OpNum() int { return 0 }

func (o *xxx_GetPrimaryDomainInformationOperation) OpName() string {
	return "/dssetup/v0/DsRolerGetPrimaryDomainInformation"
}

func (o *xxx_GetPrimaryDomainInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrimaryDomainInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// InfoLevel {in} (1:{alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.InfoLevel)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrimaryDomainInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// InfoLevel {in} (1:{alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InfoLevel)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrimaryDomainInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrimaryDomainInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DomainInfo {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum), alias=PDSROLER_PRIMARY_DOMAIN_INFORMATION}*(1))(3:{switch_type={alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum), alias=DSROLER_PRIMARY_DOMAIN_INFORMATION}(union))
	{
		if o.DomainInfo != nil {
			_ptr_DomainInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				_swDomainInfo := uint16(o.InfoLevel)
				if o.DomainInfo != nil {
					if err := o.DomainInfo.MarshalUnionNDR(ctx, w, _swDomainInfo); err != nil {
						return err
					}
				} else {
					if err := (&PrimaryDomainInformation{}).MarshalUnionNDR(ctx, w, _swDomainInfo); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainInfo, _ptr_DomainInfo); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrimaryDomainInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DomainInfo {out} (1:{pointer=ref}*(2))(2:{switch_type={alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum), alias=PDSROLER_PRIMARY_DOMAIN_INFORMATION,pointer=ref}*(1))(3:{switch_type={alias=DSROLE_PRIMARY_DOMAIN_INFO_LEVEL}(enum), alias=DSROLER_PRIMARY_DOMAIN_INFORMATION}(union))
	{
		_ptr_DomainInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainInfo == nil {
				o.DomainInfo = &PrimaryDomainInformation{}
			}
			_swDomainInfo := uint16(o.InfoLevel)
			if err := o.DomainInfo.UnmarshalUnionNDR(ctx, w, _swDomainInfo); err != nil {
				return err
			}
			return nil
		})
		_s_DomainInfo := func(ptr interface{}) { o.DomainInfo = *ptr.(**PrimaryDomainInformation) }
		if err := w.ReadPointer(&o.DomainInfo, _s_DomainInfo, _ptr_DomainInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPrimaryDomainInformationRequest structure represents the DsRolerGetPrimaryDomainInformation operation request
type GetPrimaryDomainInformationRequest struct {
	// InfoLevel: The type of data requested by the client. For possible values in this
	// enumeration, see section 2.2.7.
	InfoLevel PrimaryDomainInfoLevel `idl:"name:InfoLevel" json:"info_level"`
}

func (o *GetPrimaryDomainInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPrimaryDomainInformationOperation) *xxx_GetPrimaryDomainInformationOperation {
	if op == nil {
		op = &xxx_GetPrimaryDomainInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.InfoLevel = o.InfoLevel
	return op
}

func (o *GetPrimaryDomainInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPrimaryDomainInformationOperation) {
	if o == nil {
		return
	}
	o.InfoLevel = op.InfoLevel
}
func (o *GetPrimaryDomainInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPrimaryDomainInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrimaryDomainInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPrimaryDomainInformationResponse structure represents the DsRolerGetPrimaryDomainInformation operation response
type GetPrimaryDomainInformationResponse struct {
	// XXX: InfoLevel is an implicit input depedency for output parameters
	InfoLevel PrimaryDomainInfoLevel `idl:"name:InfoLevel" json:"info_level"`

	// DomainInfo: The requested information that the server provides to the client. The
	// value of the InfoLevel parameter indicates the type of information that is requested;
	// information is returned in the corresponding member of the DSROLER_PRIMARY_DOMAIN_INFORMATION
	// union.
	DomainInfo *PrimaryDomainInformation `idl:"name:DomainInfo;switch_is:InfoLevel" json:"domain_info"`
	// Return: The DsRolerGetPrimaryDomainInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPrimaryDomainInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPrimaryDomainInformationOperation) *xxx_GetPrimaryDomainInformationOperation {
	if op == nil {
		op = &xxx_GetPrimaryDomainInformationOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.InfoLevel == PrimaryDomainInfoLevel(0) {
		op.InfoLevel = o.InfoLevel
	}

	op.DomainInfo = o.DomainInfo
	op.Return = o.Return
	return op
}

func (o *GetPrimaryDomainInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPrimaryDomainInformationOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.InfoLevel = op.InfoLevel

	o.DomainInfo = op.DomainInfo
	o.Return = op.Return
}
func (o *GetPrimaryDomainInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPrimaryDomainInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrimaryDomainInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
