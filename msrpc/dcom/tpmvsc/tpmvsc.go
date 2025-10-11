// The tpmvsc package implements the TPMVSC client protocol.
//
// # Introduction
//
// The DCOM Interfaces for Trusted Platform Module (TPM) Virtual Smart Card Management
// Protocol is used to manage virtual smart cards (VSCs) on a remote machine, such as
// those based on trusted platform modules (TPM). It provides methods for a protocol
// client to request creation and destruction of VSCs and to monitor the status of these
// operations.
//
// # Overview
//
// The DCOM Interfaces for the Trusted Platform Module (TPM) Virtual Smart Card Management
// Protocol provides a Distributed Component Object Model (DCOM) Remote Protocol [MS-DCOM]
// interface used for creating and destroying VSCs. Like all other DCOM interfaces,
// this protocol uses RPC [C706], with the extensions specified in [MS-RPCE], as its
// underlying protocol. A VSC is a device that presents a device interface complying
// with the PC/SC specification for PC-connected interface devices [PCSC3] to its host
// operating system (OS) platform. This protocol does not assume anything about the
// underlying implementation of VSC devices. In particular, while it is primarily intended
// for the management of VSCs based on TPMs, it can also be used to manage other types
// of VSCs. The protocol defines two interfaces: a primary interface which is used to
// request VSC operations on a target system, and a secondary interface which is used
// by that target system to return status and progress information to the requestor.
//
// In a typical scenario, this protocol is used by a requestor (generally an administrative
// workstation) to manage VSC devices on a target (generally an end-user workstation).
// The requestor, acting as a client, connects to the ITpmVirtualSmartCardManager, ITpmVirtualSmartCardManager2,
// or ITpmVirtualSmartCardManager3 interface on the target (which acts as the server)
// and requests the target to either create or destroy a VSC by passing appropriate
// parameters. These parameters include a reference to an ITpmVirtualSmartCardManagerStatusCallback
// DCOM interface on the requestor that can be used to provide status updates through
// callbacks.
//
// The principal difference between the ITpmVirtualSmartCardManager2 interface and the
// ITpmVirtualSmartCardManager3 interface is that the latter supports creation of attestation-capable
// virtual smart cards.
//
// The principal difference between the ITpmVirtualSmartCardManager interface and the
// ITpmVirtualSmartCardManager2 interface is that the latter supports policies to define
// valid values for the smart-card PIN.
//
// The target, after validating these parameters, starts executing the requested operation.
// It also opens a second connection back to the requestor over which it invokes the
// requestor’s ITpmVirtualSmartCardManagerStatusCallback interface as a client, and
// calls the appropriate functions of that interface to provide progress or error codes.
// When the operation is completed, the target closes this second connection and returns
// the result for the requestor’s original method invocation.
//
// This entire process is illustrated in Figure 1.
package tpmvsc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/tpmvsc"
)

// DefaultAdminAlgorithmID represents the TPMVSC_DEFAULT_ADMIN_ALGORITHM_ID RPC constant
var DefaultAdminAlgorithmID = 130

// AttestationType type represents TPMVSC_ATTESTATION_TYPE RPC enumeration.
type AttestationType uint32

var (
	AttestationTypeNone              AttestationType = 0
	AttestationTypeAikOnly           AttestationType = 1
	AttestationTypeAikAndCertificate AttestationType = 2
)

func (o AttestationType) String() string {
	switch o {
	case AttestationTypeNone:
		return "AttestationTypeNone"
	case AttestationTypeAikOnly:
		return "AttestationTypeAikOnly"
	case AttestationTypeAikAndCertificate:
		return "AttestationTypeAikAndCertificate"
	}
	return "Invalid"
}

// Status type represents TPMVSCMGR_STATUS RPC enumeration.
type Status uint32

var (
	// TPMVSCMGR_STATUS_VTPMSMARTCARD_INITIALIZING: Initializing the VSC component.
	StatusVTPMSmartCardInitializing Status = 0
	// TPMVSCMGR_STATUS_VTPMSMARTCARD_CREATING: Creating the VSC component.
	StatusVTPMSmartCardCreating Status = 1
	// TPMVSCMGR_STATUS_VTPMSMARTCARD_DESTROYING: Deleting the VSC component.
	StatusVTPMSmartCardDestroying Status = 2
	// TPMVSCMGR_STATUS_VGIDSSIMULATOR_INITIALIZING: Initializing the VSC simulator.
	StatusVGIDSSimulatorInitializing Status = 3
	// TPMVSCMGR_STATUS_VGIDSSIMULATOR_CREATING: Creating the VSC simulator.
	StatusVGIDSSimulatorCreating Status = 4
	// TPMVSCMGR_STATUS_VGIDSSIMULATOR_DESTROYING: Destroying the VSC simulator.
	StatusVGIDSSimulatorDestroying Status = 5
	// TPMVSCMGR_STATUS_VREADER_INITIALIZING: Initializing the VSC reader.
	StatusVReaderInitializing Status = 6
	// TPMVSCMGR_STATUS_VREADER_CREATING: Creating the VSC reader.
	StatusVReaderCreating Status = 7
	// TPMVSCMGR_STATUS_VREADER_DESTROYING: Destroying the VSC reader.
	StatusVReaderDestroying Status = 8
	// TPMVSCMGR_STATUS_GENERATE_WAITING: Waiting for the VSC device.
	StatusGenerateWaiting Status = 9
	// TPMVSCMGR_STATUS_GENERATE_AUTHENTICATING: Authenticating to the VSC.
	StatusGenerateAuthenticating Status = 10
	// TPMVSCMGR_STATUS_GENERATE_RUNNING: Generating the file system on the VSC.
	StatusGenerateRunning Status = 11
	// TPMVSCMGR_STATUS_CARD_CREATED: The VSC is created.
	StatusCardCreated Status = 12
	// TPMVSCMGR_STATUS_CARD_DESTROYED: The VSC is deleted.
	StatusCardDestroyed Status = 13
)

func (o Status) String() string {
	switch o {
	case StatusVTPMSmartCardInitializing:
		return "StatusVTPMSmartCardInitializing"
	case StatusVTPMSmartCardCreating:
		return "StatusVTPMSmartCardCreating"
	case StatusVTPMSmartCardDestroying:
		return "StatusVTPMSmartCardDestroying"
	case StatusVGIDSSimulatorInitializing:
		return "StatusVGIDSSimulatorInitializing"
	case StatusVGIDSSimulatorCreating:
		return "StatusVGIDSSimulatorCreating"
	case StatusVGIDSSimulatorDestroying:
		return "StatusVGIDSSimulatorDestroying"
	case StatusVReaderInitializing:
		return "StatusVReaderInitializing"
	case StatusVReaderCreating:
		return "StatusVReaderCreating"
	case StatusVReaderDestroying:
		return "StatusVReaderDestroying"
	case StatusGenerateWaiting:
		return "StatusGenerateWaiting"
	case StatusGenerateAuthenticating:
		return "StatusGenerateAuthenticating"
	case StatusGenerateRunning:
		return "StatusGenerateRunning"
	case StatusCardCreated:
		return "StatusCardCreated"
	case StatusCardDestroyed:
		return "StatusCardDestroyed"
	}
	return "Invalid"
}

// Error type represents TPMVSCMGR_ERROR RPC enumeration.
type Error uint32

var (
	// TPMVSCMGR_ERROR_IMPERSONATION: An error occurred during impersonation of the caller.
	ErrorImpersonation Error = 0
	// TPMVSCMGR_ERROR_PIN_COMPLEXITY: The user personal identification number (PIN) or
	// personal unblocking key (PUK) value does not meet the minimum length requirement.
	ErrorPINComplexity Error = 1
	// TPMVSCMGR_ERROR_READER_COUNT_LIMIT:  The limit on the number of Smart Card Readers
	// has been reached.
	ErrorReaderCountLimit Error = 2
	// TPMVSCMGR_ERROR_TERMINAL_SERVICES_SESSION: The TPM Virtual Smart Card Management
	// Protocol cannot be used within a Terminal Services session.
	ErrorTerminalServicesSession Error = 3
	// TPMVSCMGR_ERROR_VTPMSMARTCARD_INITIALIZE: An error occurred during initialization
	// of the VSC component.
	ErrorVTPMSmartCardInitialize Error = 4
	// TPMVSCMGR_ERROR_VTPMSMARTCARD_CREATE: An error occurred during creation of the VSC
	// component.
	ErrorVTPMSmartCardCreate Error = 5
	// TPMVSCMGR_ERROR_VTPMSMARTCARD_DESTROY: An error occurred during deletion of the VSC
	// component.
	ErrorVTPMSmartCardDestroy Error = 6
	// TPMVSCMGR_ERROR_VGIDSSIMULATOR_INITIALIZE: An error occurred during initialization
	// of the VSC simulator.
	ErrorVGIDSSimulatorInitialize Error = 7
	// TPMVSCMGR_ERROR_VGIDSSIMULATOR_CREATE: An error occurred during creation of the VSC
	// simulator.
	ErrorVGIDSSimulatorCreate Error = 8
	// TPMVSCMGR_ERROR_VGIDSSIMULATOR_DESTROY: An error occurred during deletion of the
	// VSC simulator.
	ErrorVGIDSSimulatorDestroy Error = 9
	// TPMVSCMGR_ERROR_VGIDSSIMULATOR_WRITE_PROPERTY: An error occurred during configuration
	// of the VSC simulator.
	ErrorVGIDSSimulatorWriteProperty Error = 10
	// TPMVSCMGR_ERROR_VGIDSSIMULATOR_READ_PROPERTY: An error occurred finding the VSC simulator.
	ErrorVGIDSSimulatorReadProperty Error = 11
	// TPMVSCMGR_ERROR_VREADER_INITIALIZE: An error occurred during the initialization of
	// the VSC reader.
	ErrorVReaderInitialize Error = 12
	// TPMVSCMGR_ERROR_VREADER_CREATE: An error occurred during creation of the VSC reader.
	ErrorVReaderCreate Error = 13
	// TPMVSCMGR_ERROR_VREADER_DESTROY: An error occurred during deletion of the VSC reader.
	ErrorVReaderDestroy Error = 14
	// TPMVSCMGR_ERROR_GENERATE_LOCATE_READER: An error occurred preventing connection to
	// the VSC reader.
	ErrorGenerateLocateReader Error = 15
	// TPMVSCMGR_ERROR_GENERATE_FILESYSTEM: An error occurred during generation of the file
	// system on the VSC.
	ErrorGenerateFilesystem Error = 16
	// TPMVSCMGR_ERROR_CARD_CREATE: An error occurred during creation of the VSC.
	ErrorCardCreate Error = 17
	// TPMVSCMGR_ERROR_CARD_DESTROY: An error occurred during deletion of the VSC.
	ErrorCardDestroy Error = 18
)

func (o Error) String() string {
	switch o {
	case ErrorImpersonation:
		return "ErrorImpersonation"
	case ErrorPINComplexity:
		return "ErrorPINComplexity"
	case ErrorReaderCountLimit:
		return "ErrorReaderCountLimit"
	case ErrorTerminalServicesSession:
		return "ErrorTerminalServicesSession"
	case ErrorVTPMSmartCardInitialize:
		return "ErrorVTPMSmartCardInitialize"
	case ErrorVTPMSmartCardCreate:
		return "ErrorVTPMSmartCardCreate"
	case ErrorVTPMSmartCardDestroy:
		return "ErrorVTPMSmartCardDestroy"
	case ErrorVGIDSSimulatorInitialize:
		return "ErrorVGIDSSimulatorInitialize"
	case ErrorVGIDSSimulatorCreate:
		return "ErrorVGIDSSimulatorCreate"
	case ErrorVGIDSSimulatorDestroy:
		return "ErrorVGIDSSimulatorDestroy"
	case ErrorVGIDSSimulatorWriteProperty:
		return "ErrorVGIDSSimulatorWriteProperty"
	case ErrorVGIDSSimulatorReadProperty:
		return "ErrorVGIDSSimulatorReadProperty"
	case ErrorVReaderInitialize:
		return "ErrorVReaderInitialize"
	case ErrorVReaderCreate:
		return "ErrorVReaderCreate"
	case ErrorVReaderDestroy:
		return "ErrorVReaderDestroy"
	case ErrorGenerateLocateReader:
		return "ErrorGenerateLocateReader"
	case ErrorGenerateFilesystem:
		return "ErrorGenerateFilesystem"
	case ErrorCardCreate:
		return "ErrorCardCreate"
	case ErrorCardDestroy:
		return "ErrorCardDestroy"
	}
	return "Invalid"
}

// VirtualSmartCardManager structure represents ITpmVirtualSmartCardManager RPC structure.
type VirtualSmartCardManager dcom.InterfacePointer

func (o *VirtualSmartCardManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VirtualSmartCardManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VirtualSmartCardManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VirtualSmartCardManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VirtualSmartCardManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VirtualSmartCardManager2 structure represents ITpmVirtualSmartCardManager2 RPC structure.
type VirtualSmartCardManager2 dcom.InterfacePointer

func (o *VirtualSmartCardManager2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VirtualSmartCardManager2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VirtualSmartCardManager2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VirtualSmartCardManager2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VirtualSmartCardManager2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VirtualSmartCardManager3 structure represents ITpmVirtualSmartCardManager3 RPC structure.
type VirtualSmartCardManager3 dcom.InterfacePointer

func (o *VirtualSmartCardManager3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VirtualSmartCardManager3) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VirtualSmartCardManager3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VirtualSmartCardManager3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VirtualSmartCardManager3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VirtualSmartCardManagerStatusCallback structure represents ITpmVirtualSmartCardManagerStatusCallback RPC structure.
type VirtualSmartCardManagerStatusCallback dcom.InterfacePointer

func (o *VirtualSmartCardManagerStatusCallback) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VirtualSmartCardManagerStatusCallback) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VirtualSmartCardManagerStatusCallback) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VirtualSmartCardManagerStatusCallback) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VirtualSmartCardManagerStatusCallback) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
