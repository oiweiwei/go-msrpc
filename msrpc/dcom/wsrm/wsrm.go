// The wsrm package implements the WSRM client protocol.
//
// # Introduction
//
// This is a specification of the Windows System Resource Manager (WSRM) Protocol. It
// is based on the Remote Procedure Call (RPC) Protocol [C706] [MS-RPCE].
//
// WSRM exposes a set of Distributed Component Object Model (DCOM) Remote Protocol interfaces
// for managing processor and memory resources and accounting functions on a computer.
// Using these interfaces, the following operations can be performed:
//
// * Create, enumerate, modify, or delete rules governing resource consumption.
//
// * Retrieve, store, or delete long-term, persistent resource consumption histories.
//
// * Directly control resource management ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_43754e9e-2697-43b5-857d-2a81c17f010a
// ) enforcement.
//
// # Overview
//
// The Windows System Resource Manager (WSRM) Protocol provides tools for managing processor
// and memory resources on a computer. With WSRM, administrators can control how CPU
// resources are allocated to applications, services, and processes, and prevent applications
// from consuming more than their share, thereby preventing one application from starving
// other applications of CPU resources and memory.
//
// WSRM policies can be applied according to a date/time schedule. This allows administrators
// to free up the CPU and memory for maintenance applications during nonpeak hours and
// for mission-critical applications during peak hours.
//
// The WSRM accounting feature allows administrators to generate, store, view, and export
// resource utilization reports for systems management, as well as service-level agreement
// tracking and billing information.
//
// WSRM exposes a set of DCOM Protocol [MS-DCOM] interfaces that perform the following
// functions:
//
// * *Resource Management:* Performs administrative tasks and retrieves system information
// from the computer being managed.
//
// * *Accounting Management:* Performs configuration tasks on accounting data, including
// the following:
//
// * Specifying the database server.
//
// * Changing the location of database files.
//
// * Controlling how often the server creates accounting records.
//
// * Managing accounting data records.
//
// * *Calendar Management:* Schedules resource management ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_43754e9e-2697-43b5-857d-2a81c17f010a
// ) operations.
//
// * *Configuration:* Manages the global configuration of server resources.
//
// * *Machine Group Management:* Creates, manages, and deletes machine groups ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_f4153c8c-848d-4db5-b9cc-4113e5f1e406
// ).
//
// * *Object Management:* Imports and exports objects with smart import-conflict resolution
// ( e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_576cff61-4f0a-4a8a-9829-52ad99f6891c ).
//
// * *Policy Management:* Creates, edits, enumerates, and deletes objects that control
// the allocation of resources to user tasks.
//
// * *Protocol Management:* Manages client versions supported on the server.
//
// * *Remote Session Management:* Manages CPU quota allocation to different remote sessions.
//
// * *Resource Group Management:* Retrieves, modifies, and deletes resource groups (
// e371cc74-bf4c-4870-8afe-5062cc628b4f#gt_04c2dc43-6f4a-432d-bca6-ceeef2d5f2ae ).
package wsrm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

// ConfigType type represents CONFIGTYPE RPC enumeration.
//
// The CONFIGTYPE enumeration defines types of WSRM configuration objects on which operations
// can be performed. WSRM configuration objects are used in the IWRMConfig Interface
// (section 3.2.4.5).
type ConfigType uint32

var (
	// CONFIGTYPE_ACCOUNTING:  This operation targets the accounting configuration.
	ConfigTypeAccounting ConfigType = 1
	// CONFIGTYPE_NOTIFICATION:  This operation targets the notification configuration.
	ConfigTypeNotification ConfigType = 2
	// CONFIGTYPE_CALENDARING:  This operation targets the calendar configuration.
	ConfigTypeCalendaring ConfigType = 3
)

func (o ConfigType) String() string {
	switch o {
	case ConfigTypeAccounting:
		return "ConfigTypeAccounting"
	case ConfigTypeNotification:
		return "ConfigTypeNotification"
	case ConfigTypeCalendaring:
		return "ConfigTypeCalendaring"
	}
	return "Invalid"
}

// RestoreMode type represents RESTORE_MODE RPC enumeration.
//
// The RESTORE_MODE enumeration defines states that the WSRM configuration can have
// when it is restored. The WSRM configuration is restored by the RestoreXMLFiles method
// (section 3.2.4.1.5).
type RestoreMode uint32

var (
	// RESTORE_LAST_GOOD_STATE:  This value targets the last saved good state of the WSRM
	// configuration, which is defined as a consistent set of WSRM objects. WSRM objects
	// are stored in the XML files after successful completion of any method call of this
	// protocol.
	RestoreModeLastGoodState RestoreMode = 1
	// RESTORE_EMPTY_FILES:  This value targets the initial state of the WSRM configuration
	// with which the product is shipped. This means that only the out-of-box policies,
	// empty calendars, and conditional policies will be used in the restored state of the
	// object.
	RestoreModeEmptyFiles RestoreMode = 2
)

func (o RestoreMode) String() string {
	switch o {
	case RestoreModeLastGoodState:
		return "RestoreModeLastGoodState"
	case RestoreModeEmptyFiles:
		return "RestoreModeEmptyFiles"
	}
	return "Invalid"
}

// ObjectType type represents OBJECT_TYPE RPC enumeration.
//
// The OBJECT_TYPE enumeration defines types of objects on which an operation can be
// performed. The WSRM objects are used by the GetDependencies (section 3.2.4.1.6) and
// ExportObjects (section 3.2.4.2.1) methods.
type ObjectType uint32

var (
	// OBJECT_SELECTION_CRITERIA:  The target object is a process matching criteria (PMC).
	ObjectTypeSelectionCriteria ObjectType = 1
	// OBJECT_POLICY:  The target object is a policy.
	ObjectTypePolicy ObjectType = 2
	// OBJECT_SCHEDULE:  The target object is a calendar or a schedule object.
	ObjectTypeSchedule ObjectType = 3
)

func (o ObjectType) String() string {
	switch o {
	case ObjectTypeSelectionCriteria:
		return "ObjectTypeSelectionCriteria"
	case ObjectTypePolicy:
		return "ObjectTypePolicy"
	case ObjectTypeSchedule:
		return "ObjectTypeSchedule"
	}
	return "Invalid"
}

// ManagementType type represents MANAGEMENT_TYPE RPC enumeration.
//
// The MANAGEMENT_TYPE enumeration defines management modes for the WSRM Protocol server.
// Management modes are used by IWRMPolicy methods (section 3.2.4.7).
type ManagementType uint32

var (
	// MANUAL_ACTIVE_POLICY:  The server is in manual active policy mode. In this mode,
	// WSRM manages the CPU and memory allocation for different processes according to the
	// PMCs that are defined by the current resource policy. The current resource policy
	// is selected using the SetCurrentPolicy method (section 3.2.4.7.12).
	ManagementTypeManualActivePolicy ManagementType = 1
	// CALENDAR_POLICY:  The server is in calendar mode. In this mode, a resource policy
	// is managing the CPU and memory allocation for different processes according to the
	// PMCs defined by the policy. The policy was set and the management mode was set to
	// calendar mode as a result of a calendar event, which is created and managed by IWRMCalendar
	// methods (section 3.2.4.4).
	ManagementTypeCalendarPolicy ManagementType = 2
	// PROFILING:  The server is in profiling mode. In this mode, WSRM does not manage
	// the CPU and memory allocation for processes. If accounting is enabled, process properties
	// of all running processes are logged into the accounting database.
	ManagementTypeProfiling ManagementType = 3
)

func (o ManagementType) String() string {
	switch o {
	case ManagementTypeManualActivePolicy:
		return "ManagementTypeManualActivePolicy"
	case ManagementTypeCalendarPolicy:
		return "ManagementTypeCalendarPolicy"
	case ManagementTypeProfiling:
		return "ManagementTypeProfiling"
	}
	return "Invalid"
}

// ExclusionlistType type represents EXCLUSIONLIST_TYPE RPC enumeration.
//
// The EXCLUSIONLIST_TYPE enumeration defines types of server exclusion list. A WSRM
// exclusion list is in the form of an ExclusionList element (section 2.2.5.16) in XML.
// They are used by IWRMConfig Interface (section 3.2.4.5).
type ExclusionlistType uint32

var (
	// SYSTEM_EXCLUSION_LIST:  The exclusion list is system-defined.
	ExclusionlistTypeSystemExclusionList ExclusionlistType = 1
	// USER_EXCLUSION_LIST:  The exclusion list is user-defined.
	ExclusionlistTypeUserExclusionList ExclusionlistType = 2
	// DEFAULT_USER_EXCLUSION_LIST:  The exclusion list is the default user-defined exclusion
	// list.<9>
	ExclusionlistTypeDefaultUserExclusionList ExclusionlistType = 4
)

func (o ExclusionlistType) String() string {
	switch o {
	case ExclusionlistTypeSystemExclusionList:
		return "ExclusionlistTypeSystemExclusionList"
	case ExclusionlistTypeUserExclusionList:
		return "ExclusionlistTypeUserExclusionList"
	case ExclusionlistTypeDefaultUserExclusionList:
		return "ExclusionlistTypeDefaultUserExclusionList"
	}
	return "Invalid"
}

// ImportType type represents IMPORT_TYPE RPC enumeration.
//
// The IMPORT_TYPE enumeration defines modes for importing objects when conflicting
// objects are present in the configuration.
type ImportType uint32

var (
	// OVERWRITE_IMPORT:  Specifies that objects of a given type SHOULD be removed prior
	// to importing more objects of that type.
	ImportTypeOverwriteImport ImportType = 1
	// IGNORE_EXISTING_IMPORT:  Specifies that the import of objects that conflict with
	// existing objects in the configuration SHOULD be ignored.
	ImportTypeIgnoreExistingImport ImportType = 2
	// OVERRIDE_EXISTING_IMPORT:  Specifies that the import of objects that conflict with
	// existing objects in the configuration SHOULD cause the existing objects to be deleted.
	ImportTypeOverrideExistingImport ImportType = 3
	// SMART_MERGE_RENAME_EXISTING_IMPORT:  Specifies that existing objects in the configuration
	// that conflict with imported objects SHOULD be renamed prior to importing.
	ImportTypeSmartMergeRenameExistingImport ImportType = 4
	// SMART_MERGE_RENAME_IMPORTED_IMPORT:  Specifies that imported objects that conflict
	// with existing objects in the configuration SHOULD be renamed prior to importing.
	ImportTypeSmartMergeRenameImportedImport ImportType = 5
)

func (o ImportType) String() string {
	switch o {
	case ImportTypeOverwriteImport:
		return "ImportTypeOverwriteImport"
	case ImportTypeIgnoreExistingImport:
		return "ImportTypeIgnoreExistingImport"
	case ImportTypeOverrideExistingImport:
		return "ImportTypeOverrideExistingImport"
	case ImportTypeSmartMergeRenameExistingImport:
		return "ImportTypeSmartMergeRenameExistingImport"
	case ImportTypeSmartMergeRenameImportedImport:
		return "ImportTypeSmartMergeRenameImportedImport"
	}
	return "Invalid"
}

// MachineGroupMergeOptions type represents MACHINE_GROUP_MERGE_OPTIONS RPC enumeration.
//
// The MACHINE_GROUP_MERGE_OPTIONS enumeration defines options for machine group modification.
type MachineGroupMergeOptions uint32

var (
	// OVERWRITE_MG_MERGE_OPTION:  Specifies that the machine group configuration SHOULD
	// be overwritten.
	MachineGroupMergeOptionsOverwriteMGMergeOption MachineGroupMergeOptions = 1
	// OVERRIDE_MG_MERGE_OPTION:  Specifies that the machine group configuration SHOULD
	// be overridden. That means that if the same machine group name exists with the same
	// hierarchy then this (along with its children) SHOULD be replaced by the new machine
	// group node. Nonconflicting new nodes SHOULD be imported as is. Nonconflicting existing
	// nodes SHOULD remain as is.
	MachineGroupMergeOptionsOverrideMGMergeOption MachineGroupMergeOptions = 2
	// APPEND_MG_MERGE_OPTION:  Specifies that the machine group configuration SHOULD be
	// appended. This means that nonconflicting nodes SHOULD remain as is in the existing
	// configuration. Nonconflicting new nodes SHOULD be imported as is. Conflicting nodes
	// SHOULD not be modified.
	MachineGroupMergeOptionsAppendMGMergeOption MachineGroupMergeOptions = 3
	// SMART_MG_MERGE_OPTION:  Specifies that the machine group configuration SHOULD be
	// handled smartly. This means that nonconflicting nodes SHOULD remain as is in the
	// existing configuration. Nonconflicting new nodes SHOULD be imported as is. Conflicting
	// nodes SHOULD be merged in a way that the name of the imported node will be changed
	// to avoid name conflict. The new name of the conflicting object, which is being imported,
	// is built by appending its original name with the string "##@" followed by a number.
	// The numbers used are in the range 1-16384 and a number, once used, is not reused
	// until the object it was used in is either deleted or renamed. For example, an object
	// name "ObjectName" might become "ObjectName##@1" after using this option. This range
	// is sufficiently large and is not consumed completely because, in the WSRM configuration,
	// the maximum allowed number of objects of each type is 128. In addition, whenever
	// an object with a smart name is deleted or renamed, the number that was used in its
	// name is available for reuse.
	MachineGroupMergeOptionsSmartMGMergeOption MachineGroupMergeOptions = 4
)

func (o MachineGroupMergeOptions) String() string {
	switch o {
	case MachineGroupMergeOptionsOverwriteMGMergeOption:
		return "MachineGroupMergeOptionsOverwriteMGMergeOption"
	case MachineGroupMergeOptionsOverrideMGMergeOption:
		return "MachineGroupMergeOptionsOverrideMGMergeOption"
	case MachineGroupMergeOptionsAppendMGMergeOption:
		return "MachineGroupMergeOptionsAppendMGMergeOption"
	case MachineGroupMergeOptionsSmartMGMergeOption:
		return "MachineGroupMergeOptionsSmartMGMergeOption"
	}
	return "Invalid"
}

// ResourceManager2 structure represents IResourceManager2 RPC structure.
type ResourceManager2 dcom.InterfacePointer

func (o *ResourceManager2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ResourceManager2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ResourceManager2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ResourceManager2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResourceManager2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Policy structure represents IWRMPolicy RPC structure.
type Policy dcom.InterfacePointer

func (o *Policy) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Policy) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Policy) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Policy) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Policy) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ResourceManager structure represents IResourceManager RPC structure.
type ResourceManager dcom.InterfacePointer

func (o *ResourceManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ResourceManager) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ResourceManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ResourceManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResourceManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Accounting structure represents IWRMAccounting RPC structure.
type Accounting dcom.InterfacePointer

func (o *Accounting) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Accounting) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Accounting) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Accounting) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Accounting) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Config structure represents IWRMConfig RPC structure.
type Config dcom.InterfacePointer

func (o *Config) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Config) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Config) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Config) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Config) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Calendar structure represents IWRMCalendar RPC structure.
type Calendar dcom.InterfacePointer

func (o *Calendar) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Calendar) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Calendar) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Calendar) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Calendar) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RemoteSessionManagement structure represents IWRMRemoteSessionMgmt RPC structure.
type RemoteSessionManagement dcom.InterfacePointer

func (o *RemoteSessionManagement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteSessionManagement) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *RemoteSessionManagement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteSessionManagement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteSessionManagement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Protocol structure represents IWRMProtocol RPC structure.
type Protocol dcom.InterfacePointer

func (o *Protocol) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Protocol) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Protocol) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Protocol) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Protocol) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ResourceGroup structure represents IWRMResourceGroup RPC structure.
type ResourceGroup dcom.InterfacePointer

func (o *ResourceGroup) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ResourceGroup) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ResourceGroup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ResourceGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ResourceGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// MachineGroup structure represents IWRMMachineGroup RPC structure.
type MachineGroup dcom.InterfacePointer

func (o *MachineGroup) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *MachineGroup) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *MachineGroup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *MachineGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MachineGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ResourceManager class identifier e8bcffac-b864-4574-b2e8-f1fb21dfdc18
var ResourceManagerClassID = &dcom.ClassID{Data1: 0xe8bcffac, Data2: 0xb864, Data3: 0x4574, Data4: []byte{0xb2, 0xe8, 0xf1, 0xfb, 0x21, 0xdf, 0xdc, 0x18}}
