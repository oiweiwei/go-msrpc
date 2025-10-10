// The tpmvsc package implements the TPMVSC client protocol.
//
// # Introduction
//
// # Overview
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
	StatusVTPMSmartCardInitializing  Status = 0
	StatusVTPMSmartCardCreating      Status = 1
	StatusVTPMSmartCardDestroying    Status = 2
	StatusVGIDSSimulatorInitializing Status = 3
	StatusVGIDSSimulatorCreating     Status = 4
	StatusVGIDSSimulatorDestroying   Status = 5
	StatusVReaderInitializing        Status = 6
	StatusVReaderCreating            Status = 7
	StatusVReaderDestroying          Status = 8
	StatusGenerateWaiting            Status = 9
	StatusGenerateAuthenticating     Status = 10
	StatusGenerateRunning            Status = 11
	StatusCardCreated                Status = 12
	StatusCardDestroyed              Status = 13
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
	ErrorImpersonation               Error = 0
	ErrorPINComplexity               Error = 1
	ErrorReaderCountLimit            Error = 2
	ErrorTerminalServicesSession     Error = 3
	ErrorVTPMSmartCardInitialize     Error = 4
	ErrorVTPMSmartCardCreate         Error = 5
	ErrorVTPMSmartCardDestroy        Error = 6
	ErrorVGIDSSimulatorInitialize    Error = 7
	ErrorVGIDSSimulatorCreate        Error = 8
	ErrorVGIDSSimulatorDestroy       Error = 9
	ErrorVGIDSSimulatorWriteProperty Error = 10
	ErrorVGIDSSimulatorReadProperty  Error = 11
	ErrorVReaderInitialize           Error = 12
	ErrorVReaderCreate               Error = 13
	ErrorVReaderDestroy              Error = 14
	ErrorGenerateLocateReader        Error = 15
	ErrorGenerateFilesystem          Error = 16
	ErrorCardCreate                  Error = 17
	ErrorCardDestroy                 Error = 18
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
