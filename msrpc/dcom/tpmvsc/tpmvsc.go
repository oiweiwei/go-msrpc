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

// TpmvscDefaultAdminAlgorithmID represents the TPMVSC_DEFAULT_ADMIN_ALGORITHM_ID RPC constant
var TpmvscDefaultAdminAlgorithmID = 130

// TpmvscAttestationType type represents TPMVSC_ATTESTATION_TYPE RPC enumeration.
type TpmvscAttestationType uint32

var (
	TpmvscAttestationTypeNone              TpmvscAttestationType = 0
	TpmvscAttestationTypeAikOnly           TpmvscAttestationType = 1
	TpmvscAttestationTypeAikAndCertificate TpmvscAttestationType = 2
)

func (o TpmvscAttestationType) String() string {
	switch o {
	case TpmvscAttestationTypeNone:
		return "TpmvscAttestationTypeNone"
	case TpmvscAttestationTypeAikOnly:
		return "TpmvscAttestationTypeAikOnly"
	case TpmvscAttestationTypeAikAndCertificate:
		return "TpmvscAttestationTypeAikAndCertificate"
	}
	return "Invalid"
}

// TpmvscmgrStatus type represents TPMVSCMGR_STATUS RPC enumeration.
type TpmvscmgrStatus uint32

var (
	TpmvscmgrStatusVtpmsmartcardInitializing  TpmvscmgrStatus = 0
	TpmvscmgrStatusVtpmsmartcardCreating      TpmvscmgrStatus = 1
	TpmvscmgrStatusVtpmsmartcardDestroying    TpmvscmgrStatus = 2
	TpmvscmgrStatusVgidssimulatorInitializing TpmvscmgrStatus = 3
	TpmvscmgrStatusVgidssimulatorCreating     TpmvscmgrStatus = 4
	TpmvscmgrStatusVgidssimulatorDestroying   TpmvscmgrStatus = 5
	TpmvscmgrStatusVreaderInitializing        TpmvscmgrStatus = 6
	TpmvscmgrStatusVreaderCreating            TpmvscmgrStatus = 7
	TpmvscmgrStatusVreaderDestroying          TpmvscmgrStatus = 8
	TpmvscmgrStatusGenerateWaiting            TpmvscmgrStatus = 9
	TpmvscmgrStatusGenerateAuthenticating     TpmvscmgrStatus = 10
	TpmvscmgrStatusGenerateRunning            TpmvscmgrStatus = 11
	TpmvscmgrStatusCardCreated                TpmvscmgrStatus = 12
	TpmvscmgrStatusCardDestroyed              TpmvscmgrStatus = 13
)

func (o TpmvscmgrStatus) String() string {
	switch o {
	case TpmvscmgrStatusVtpmsmartcardInitializing:
		return "TpmvscmgrStatusVtpmsmartcardInitializing"
	case TpmvscmgrStatusVtpmsmartcardCreating:
		return "TpmvscmgrStatusVtpmsmartcardCreating"
	case TpmvscmgrStatusVtpmsmartcardDestroying:
		return "TpmvscmgrStatusVtpmsmartcardDestroying"
	case TpmvscmgrStatusVgidssimulatorInitializing:
		return "TpmvscmgrStatusVgidssimulatorInitializing"
	case TpmvscmgrStatusVgidssimulatorCreating:
		return "TpmvscmgrStatusVgidssimulatorCreating"
	case TpmvscmgrStatusVgidssimulatorDestroying:
		return "TpmvscmgrStatusVgidssimulatorDestroying"
	case TpmvscmgrStatusVreaderInitializing:
		return "TpmvscmgrStatusVreaderInitializing"
	case TpmvscmgrStatusVreaderCreating:
		return "TpmvscmgrStatusVreaderCreating"
	case TpmvscmgrStatusVreaderDestroying:
		return "TpmvscmgrStatusVreaderDestroying"
	case TpmvscmgrStatusGenerateWaiting:
		return "TpmvscmgrStatusGenerateWaiting"
	case TpmvscmgrStatusGenerateAuthenticating:
		return "TpmvscmgrStatusGenerateAuthenticating"
	case TpmvscmgrStatusGenerateRunning:
		return "TpmvscmgrStatusGenerateRunning"
	case TpmvscmgrStatusCardCreated:
		return "TpmvscmgrStatusCardCreated"
	case TpmvscmgrStatusCardDestroyed:
		return "TpmvscmgrStatusCardDestroyed"
	}
	return "Invalid"
}

// TpmvscmgrError type represents TPMVSCMGR_ERROR RPC enumeration.
type TpmvscmgrError uint32

var (
	TpmvscmgrErrorImpersonation               TpmvscmgrError = 0
	TpmvscmgrErrorPinComplexity               TpmvscmgrError = 1
	TpmvscmgrErrorReaderCountLimit            TpmvscmgrError = 2
	TpmvscmgrErrorTerminalServicesSession     TpmvscmgrError = 3
	TpmvscmgrErrorVtpmsmartcardInitialize     TpmvscmgrError = 4
	TpmvscmgrErrorVtpmsmartcardCreate         TpmvscmgrError = 5
	TpmvscmgrErrorVtpmsmartcardDestroy        TpmvscmgrError = 6
	TpmvscmgrErrorVgidssimulatorInitialize    TpmvscmgrError = 7
	TpmvscmgrErrorVgidssimulatorCreate        TpmvscmgrError = 8
	TpmvscmgrErrorVgidssimulatorDestroy       TpmvscmgrError = 9
	TpmvscmgrErrorVgidssimulatorWriteProperty TpmvscmgrError = 10
	TpmvscmgrErrorVgidssimulatorReadProperty  TpmvscmgrError = 11
	TpmvscmgrErrorVreaderInitialize           TpmvscmgrError = 12
	TpmvscmgrErrorVreaderCreate               TpmvscmgrError = 13
	TpmvscmgrErrorVreaderDestroy              TpmvscmgrError = 14
	TpmvscmgrErrorGenerateLocateReader        TpmvscmgrError = 15
	TpmvscmgrErrorGenerateFilesystem          TpmvscmgrError = 16
	TpmvscmgrErrorCardCreate                  TpmvscmgrError = 17
	TpmvscmgrErrorCardDestroy                 TpmvscmgrError = 18
)

func (o TpmvscmgrError) String() string {
	switch o {
	case TpmvscmgrErrorImpersonation:
		return "TpmvscmgrErrorImpersonation"
	case TpmvscmgrErrorPinComplexity:
		return "TpmvscmgrErrorPinComplexity"
	case TpmvscmgrErrorReaderCountLimit:
		return "TpmvscmgrErrorReaderCountLimit"
	case TpmvscmgrErrorTerminalServicesSession:
		return "TpmvscmgrErrorTerminalServicesSession"
	case TpmvscmgrErrorVtpmsmartcardInitialize:
		return "TpmvscmgrErrorVtpmsmartcardInitialize"
	case TpmvscmgrErrorVtpmsmartcardCreate:
		return "TpmvscmgrErrorVtpmsmartcardCreate"
	case TpmvscmgrErrorVtpmsmartcardDestroy:
		return "TpmvscmgrErrorVtpmsmartcardDestroy"
	case TpmvscmgrErrorVgidssimulatorInitialize:
		return "TpmvscmgrErrorVgidssimulatorInitialize"
	case TpmvscmgrErrorVgidssimulatorCreate:
		return "TpmvscmgrErrorVgidssimulatorCreate"
	case TpmvscmgrErrorVgidssimulatorDestroy:
		return "TpmvscmgrErrorVgidssimulatorDestroy"
	case TpmvscmgrErrorVgidssimulatorWriteProperty:
		return "TpmvscmgrErrorVgidssimulatorWriteProperty"
	case TpmvscmgrErrorVgidssimulatorReadProperty:
		return "TpmvscmgrErrorVgidssimulatorReadProperty"
	case TpmvscmgrErrorVreaderInitialize:
		return "TpmvscmgrErrorVreaderInitialize"
	case TpmvscmgrErrorVreaderCreate:
		return "TpmvscmgrErrorVreaderCreate"
	case TpmvscmgrErrorVreaderDestroy:
		return "TpmvscmgrErrorVreaderDestroy"
	case TpmvscmgrErrorGenerateLocateReader:
		return "TpmvscmgrErrorGenerateLocateReader"
	case TpmvscmgrErrorGenerateFilesystem:
		return "TpmvscmgrErrorGenerateFilesystem"
	case TpmvscmgrErrorCardCreate:
		return "TpmvscmgrErrorCardCreate"
	case TpmvscmgrErrorCardDestroy:
		return "TpmvscmgrErrorCardDestroy"
	}
	return "Invalid"
}

// TpmVirtualSmartCardManager structure represents ITpmVirtualSmartCardManager RPC structure.
type TpmVirtualSmartCardManager dcom.InterfacePointer

func (o *TpmVirtualSmartCardManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TpmVirtualSmartCardManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TpmVirtualSmartCardManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TpmVirtualSmartCardManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TpmVirtualSmartCardManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TpmVirtualSmartCardManager2 structure represents ITpmVirtualSmartCardManager2 RPC structure.
type TpmVirtualSmartCardManager2 dcom.InterfacePointer

func (o *TpmVirtualSmartCardManager2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TpmVirtualSmartCardManager2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TpmVirtualSmartCardManager2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TpmVirtualSmartCardManager2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TpmVirtualSmartCardManager2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TpmVirtualSmartCardManager3 structure represents ITpmVirtualSmartCardManager3 RPC structure.
type TpmVirtualSmartCardManager3 dcom.InterfacePointer

func (o *TpmVirtualSmartCardManager3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TpmVirtualSmartCardManager3) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TpmVirtualSmartCardManager3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TpmVirtualSmartCardManager3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TpmVirtualSmartCardManager3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TpmVirtualSmartCardManagerStatusCallback structure represents ITpmVirtualSmartCardManagerStatusCallback RPC structure.
type TpmVirtualSmartCardManagerStatusCallback dcom.InterfacePointer

func (o *TpmVirtualSmartCardManagerStatusCallback) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TpmVirtualSmartCardManagerStatusCallback) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TpmVirtualSmartCardManagerStatusCallback) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TpmVirtualSmartCardManagerStatusCallback) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TpmVirtualSmartCardManagerStatusCallback) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
