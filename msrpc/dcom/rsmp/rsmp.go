// The rsmp package implements the RSMP client protocol.
//
// # Introduction
//
// # Overview
package rsmp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

// GenericRead represents the NTMS_GENERIC_READ RPC constant
var GenericRead = 1

// GenericWrite represents the NTMS_GENERIC_WRITE RPC constant
var GenericWrite = 3

// GenericExecute represents the NTMS_GENERIC_EXECUTE RPC constant
var GenericExecute = 3

// GenericAll represents the NTMS_GENERIC_ALL RPC constant
var GenericAll = 3

// GUID structure represents NTMS_GUID RPC structure.
type GUID dtyp.GUID

func (o *GUID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *GUID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *GUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *GUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// LibraryInformation structure represents NTMS_LIBRARYINFORMATION RPC structure.
type LibraryInformation struct {
	LibraryType                  uint32     `idl:"name:LibraryType" json:"library_type"`
	CleanerSlot                  *GUID      `idl:"name:CleanerSlot" json:"cleaner_slot"`
	CleanerSlotDefault           *GUID      `idl:"name:CleanerSlotDefault" json:"cleaner_slot_default"`
	LibrarySupportsDriveCleaning bool       `idl:"name:LibrarySupportsDriveCleaning" json:"library_supports_drive_cleaning"`
	BarCodeReaderInstalled       bool       `idl:"name:BarCodeReaderInstalled" json:"bar_code_reader_installed"`
	InventoryMethod              uint32     `idl:"name:InventoryMethod" json:"inventory_method"`
	CleanerUsesRemaining         uint32     `idl:"name:dwCleanerUsesRemaining" json:"cleaner_uses_remaining"`
	FirstDriveNumber             uint32     `idl:"name:FirstDriveNumber" json:"first_drive_number"`
	NumberOfDrives               uint32     `idl:"name:dwNumberOfDrives" json:"number_of_drives"`
	FirstSlotNumber              uint32     `idl:"name:FirstSlotNumber" json:"first_slot_number"`
	NumberOfSlots                uint32     `idl:"name:dwNumberOfSlots" json:"number_of_slots"`
	FirstDoorNumber              uint32     `idl:"name:FirstDoorNumber" json:"first_door_number"`
	NumberOfDoors                uint32     `idl:"name:dwNumberOfDoors" json:"number_of_doors"`
	FirstPortNumber              uint32     `idl:"name:FirstPortNumber" json:"first_port_number"`
	NumberOfPorts                uint32     `idl:"name:dwNumberOfPorts" json:"number_of_ports"`
	FirstChangerNumber           uint32     `idl:"name:FirstChangerNumber" json:"first_changer_number"`
	NumberOfChangers             uint32     `idl:"name:dwNumberOfChangers" json:"number_of_changers"`
	NumberOfMedia                uint32     `idl:"name:dwNumberOfMedia" json:"number_of_media"`
	NumberOfMediaTypes           uint32     `idl:"name:dwNumberOfMediaTypes" json:"number_of_media_types"`
	NumberOfLibRequests          uint32     `idl:"name:dwNumberOfLibRequests" json:"number_of_lib_requests"`
	_                            *dtyp.GUID `idl:"name:Reserved"`
	AutoRecovery                 bool       `idl:"name:AutoRecovery" json:"auto_recovery"`
	Flags                        uint32     `idl:"name:dwFlags" json:"flags"`
}

func (o *LibraryInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *LibraryInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LibraryType); err != nil {
		return err
	}
	if o.CleanerSlot != nil {
		if err := o.CleanerSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CleanerSlotDefault != nil {
		if err := o.CleanerSlotDefault.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.LibrarySupportsDriveCleaning {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.BarCodeReaderInstalled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InventoryMethod); err != nil {
		return err
	}
	if err := w.WriteData(o.CleanerUsesRemaining); err != nil {
		return err
	}
	if err := w.WriteData(o.FirstDriveNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfDrives); err != nil {
		return err
	}
	if err := w.WriteData(o.FirstSlotNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfSlots); err != nil {
		return err
	}
	if err := w.WriteData(o.FirstDoorNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfDoors); err != nil {
		return err
	}
	if err := w.WriteData(o.FirstPortNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfPorts); err != nil {
		return err
	}
	if err := w.WriteData(o.FirstChangerNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfChangers); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfMedia); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfMediaTypes); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfLibRequests); err != nil {
		return err
	}
	// reserved Reserved
	if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
		return err
	}
	if !o.AutoRecovery {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *LibraryInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LibraryType); err != nil {
		return err
	}
	if o.CleanerSlot == nil {
		o.CleanerSlot = &GUID{}
	}
	if err := o.CleanerSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CleanerSlotDefault == nil {
		o.CleanerSlotDefault = &GUID{}
	}
	if err := o.CleanerSlotDefault.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bLibrarySupportsDriveCleaning int32
	if err := w.ReadData(&_bLibrarySupportsDriveCleaning); err != nil {
		return err
	}
	o.LibrarySupportsDriveCleaning = _bLibrarySupportsDriveCleaning != 0
	var _bBarCodeReaderInstalled int32
	if err := w.ReadData(&_bBarCodeReaderInstalled); err != nil {
		return err
	}
	o.BarCodeReaderInstalled = _bBarCodeReaderInstalled != 0
	if err := w.ReadData(&o.InventoryMethod); err != nil {
		return err
	}
	if err := w.ReadData(&o.CleanerUsesRemaining); err != nil {
		return err
	}
	if err := w.ReadData(&o.FirstDriveNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfDrives); err != nil {
		return err
	}
	if err := w.ReadData(&o.FirstSlotNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfSlots); err != nil {
		return err
	}
	if err := w.ReadData(&o.FirstDoorNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfDoors); err != nil {
		return err
	}
	if err := w.ReadData(&o.FirstPortNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPorts); err != nil {
		return err
	}
	if err := w.ReadData(&o.FirstChangerNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfChangers); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfMedia); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfMediaTypes); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfLibRequests); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved *dtyp.GUID
	if _Reserved == nil {
		_Reserved = &dtyp.GUID{}
	}
	if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bAutoRecovery int32
	if err := w.ReadData(&_bAutoRecovery); err != nil {
		return err
	}
	o.AutoRecovery = _bAutoRecovery != 0
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// SecurityAttributesNTMS structure represents SECURITY_ATTRIBUTES_NTMS RPC structure.
type SecurityAttributesNTMS struct {
	Length             uint32 `idl:"name:nLength" json:"length"`
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(nDescriptorLength)" json:"security_descriptor"`
	InheritHandle      bool   `idl:"name:bInheritHandle" json:"inherit_handle"`
	DescriptorLength   uint32 `idl:"name:nDescriptorLength" json:"descriptor_length"`
}

func (o *SecurityAttributesNTMS) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil && o.DescriptorLength == 0 {
		o.DescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributesNTMS) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.DescriptorLength > 0 {
		_ptr_lpSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.InheritHandle {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DescriptorLength); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributesNTMS) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_lpSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DescriptorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DescriptorLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_lpSecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
		return err
	}
	var _bInheritHandle int32
	if err := w.ReadData(&_bInheritHandle); err != nil {
		return err
	}
	o.InheritHandle = _bInheritHandle != 0
	if err := w.ReadData(&o.DescriptorLength); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// AllocationInformation structure represents NTMS_ALLOCATION_INFORMATION RPC structure.
type AllocationInformation struct {
	Size          uint32 `idl:"name:dwSize" json:"size"`
	_             uint8  `idl:"name:lpReserved"`
	AllocatedFrom *GUID  `idl:"name:AllocatedFrom" json:"allocated_from"`
}

func (o *AllocationInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *AllocationInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	// reserved lpReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if o.AllocatedFrom != nil {
		if err := o.AllocatedFrom.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *AllocationInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved lpReserved
	var _lpReserved uint8
	_ptr_lpReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_lpReserved); err != nil {
			return err
		}
		return nil
	})
	_s_lpReserved := func(ptr interface{}) { _lpReserved = *ptr.(*uint8) }
	if err := w.ReadPointer(&_lpReserved, _s_lpReserved, _ptr_lpReserved); err != nil {
		return err
	}
	if o.AllocatedFrom == nil {
		o.AllocatedFrom = &GUID{}
	}
	if err := o.AllocatedFrom.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// AsyncIO structure represents NTMS_ASYNC_IO RPC structure.
type AsyncIO struct {
	OperationID   *GUID  `idl:"name:OperationId" json:"operation_id"`
	EventID       *GUID  `idl:"name:EventId" json:"event_id"`
	OperationType uint32 `idl:"name:dwOperationType" json:"operation_type"`
	Result        uint32 `idl:"name:dwResult" json:"result"`
	AsyncState    uint32 `idl:"name:dwAsyncState" json:"async_state"`
	Event         uint64 `idl:"name:hEvent" json:"event"`
	OnStateChange bool   `idl:"name:bOnStateChange" json:"on_state_change"`
}

func (o *AsyncIO) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *AsyncIO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.OperationID != nil {
		if err := o.OperationID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EventID != nil {
		if err := o.EventID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.OperationType); err != nil {
		return err
	}
	if err := w.WriteData(o.Result); err != nil {
		return err
	}
	if err := w.WriteData(o.AsyncState); err != nil {
		return err
	}
	if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
		return err
	}
	if !o.OnStateChange {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncIO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.OperationID == nil {
		o.OperationID = &GUID{}
	}
	if err := o.OperationID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EventID == nil {
		o.EventID = &GUID{}
	}
	if err := o.EventID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Result); err != nil {
		return err
	}
	if err := w.ReadData(&o.AsyncState); err != nil {
		return err
	}
	if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
		return err
	}
	var _bOnStateChange int32
	if err := w.ReadData(&_bOnStateChange); err != nil {
		return err
	}
	o.OnStateChange = _bOnStateChange != 0
	return nil
}

// MountInformation structure represents NTMS_MOUNT_INFORMATION RPC structure.
type MountInformation struct {
	Size uint32   `idl:"name:dwSize" json:"size"`
	_    *AsyncIO `idl:"name:lpReserved;pointer:ptr"`
}

func (o *MountInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MountInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	// reserved lpReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *MountInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved lpReserved
	var _lpReserved *AsyncIO
	_ptr_lpReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if _lpReserved == nil {
			_lpReserved = &AsyncIO{}
		}
		if err := _lpReserved.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpReserved := func(ptr interface{}) { _lpReserved = *ptr.(**AsyncIO) }
	if err := w.ReadPointer(&_lpReserved, _s_lpReserved, _ptr_lpReserved); err != nil {
		return err
	}
	return nil
}

// ChangerInformationA structure represents NTMS_CHANGERINFORMATIONA RPC structure.
type ChangerInformationA struct {
	Number       uint32 `idl:"name:Number" json:"number"`
	ChangerType  *GUID  `idl:"name:ChangerType" json:"changer_type"`
	SerialNumber []byte `idl:"name:szSerialNumber" json:"serial_number"`
	Revision     []byte `idl:"name:szRevision" json:"revision"`
	DeviceName   []byte `idl:"name:szDeviceName" json:"device_name"`
	SCSIPort     uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus      uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget   uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN      uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	Library      *GUID  `idl:"name:Library" json:"library"`
}

func (o *ChangerInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ChangerInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.SerialNumber {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.SerialNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SerialNumber); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Revision {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.Revision[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Revision); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SCSIPort); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSIBus); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSITarget); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSILUN); err != nil {
		return err
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ChangerInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if o.ChangerType == nil {
		o.ChangerType = &GUID{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.SerialNumber = make([]byte, 32)
	for i1 := range o.SerialNumber {
		i1 := i1
		if err := w.ReadData(&o.SerialNumber[i1]); err != nil {
			return err
		}
	}
	o.Revision = make([]byte, 32)
	for i1 := range o.Revision {
		i1 := i1
		if err := w.ReadData(&o.Revision[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = make([]byte, 64)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.SCSIPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSIBus); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSITarget); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSILUN); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ChangerInformationW structure represents NTMS_CHANGERINFORMATIONW RPC structure.
type ChangerInformationW struct {
	Number       uint32 `idl:"name:Number" json:"number"`
	ChangerType  *GUID  `idl:"name:ChangerType" json:"changer_type"`
	SerialNumber string `idl:"name:szSerialNumber;string" json:"serial_number"`
	Revision     string `idl:"name:szRevision;string" json:"revision"`
	DeviceName   string `idl:"name:szDeviceName;string" json:"device_name"`
	SCSIPort     uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus      uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget   uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN      uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	Library      *GUID  `idl:"name:Library" json:"library"`
}

func (o *ChangerInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ChangerInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_SerialNumber_buf := utf16.Encode([]rune(o.SerialNumber))
	if uint64(len(_SerialNumber_buf)) > 32-1 {
		_SerialNumber_buf = _SerialNumber_buf[:32-1]
	}
	if o.SerialNumber != ndr.ZeroString {
		_SerialNumber_buf = append(_SerialNumber_buf, uint16(0))
	}
	for i1 := range _SerialNumber_buf {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(_SerialNumber_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_SerialNumber_buf); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Revision_buf := utf16.Encode([]rune(o.Revision))
	if uint64(len(_Revision_buf)) > 32-1 {
		_Revision_buf = _Revision_buf[:32-1]
	}
	if o.Revision != ndr.ZeroString {
		_Revision_buf = append(_Revision_buf, uint16(0))
	}
	for i1 := range _Revision_buf {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(_Revision_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Revision_buf); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_DeviceName_buf := utf16.Encode([]rune(o.DeviceName))
	if uint64(len(_DeviceName_buf)) > 64-1 {
		_DeviceName_buf = _DeviceName_buf[:64-1]
	}
	if o.DeviceName != ndr.ZeroString {
		_DeviceName_buf = append(_DeviceName_buf, uint16(0))
	}
	for i1 := range _DeviceName_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_DeviceName_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_DeviceName_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SCSIPort); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSIBus); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSITarget); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSILUN); err != nil {
		return err
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ChangerInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if o.ChangerType == nil {
		o.ChangerType = &GUID{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _SerialNumber_buf []uint16
	_SerialNumber_buf = make([]uint16, 32)
	for i1 := range _SerialNumber_buf {
		i1 := i1
		if err := w.ReadData(&_SerialNumber_buf[i1]); err != nil {
			return err
		}
	}
	o.SerialNumber = strings.TrimRight(string(utf16.Decode(_SerialNumber_buf)), ndr.ZeroString)
	var _Revision_buf []uint16
	_Revision_buf = make([]uint16, 32)
	for i1 := range _Revision_buf {
		i1 := i1
		if err := w.ReadData(&_Revision_buf[i1]); err != nil {
			return err
		}
	}
	o.Revision = strings.TrimRight(string(utf16.Decode(_Revision_buf)), ndr.ZeroString)
	var _DeviceName_buf []uint16
	_DeviceName_buf = make([]uint16, 64)
	for i1 := range _DeviceName_buf {
		i1 := i1
		if err := w.ReadData(&_DeviceName_buf[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = strings.TrimRight(string(utf16.Decode(_DeviceName_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.SCSIPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSIBus); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSITarget); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSILUN); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ChangerTypeInformationA structure represents NTMS_CHANGERTYPEINFORMATIONA RPC structure.
type ChangerTypeInformationA struct {
	Vendor     []byte `idl:"name:szVendor" json:"vendor"`
	Product    []byte `idl:"name:szProduct" json:"product"`
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *ChangerTypeInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ChangerTypeInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Vendor {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(o.Vendor[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Vendor); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Product {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(o.Product[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Product); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	return nil
}
func (o *ChangerTypeInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Vendor = make([]byte, 128)
	for i1 := range o.Vendor {
		i1 := i1
		if err := w.ReadData(&o.Vendor[i1]); err != nil {
			return err
		}
	}
	o.Product = make([]byte, 128)
	for i1 := range o.Product {
		i1 := i1
		if err := w.ReadData(&o.Product[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	return nil
}

// ChangerTypeInformationW structure represents NTMS_CHANGERTYPEINFORMATIONW RPC structure.
type ChangerTypeInformationW struct {
	Vendor     string `idl:"name:szVendor;string" json:"vendor"`
	Product    string `idl:"name:szProduct;string" json:"product"`
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *ChangerTypeInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ChangerTypeInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	_Vendor_buf := utf16.Encode([]rune(o.Vendor))
	if uint64(len(_Vendor_buf)) > 128-1 {
		_Vendor_buf = _Vendor_buf[:128-1]
	}
	if o.Vendor != ndr.ZeroString {
		_Vendor_buf = append(_Vendor_buf, uint16(0))
	}
	for i1 := range _Vendor_buf {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(_Vendor_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Vendor_buf); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Product_buf := utf16.Encode([]rune(o.Product))
	if uint64(len(_Product_buf)) > 128-1 {
		_Product_buf = _Product_buf[:128-1]
	}
	if o.Product != ndr.ZeroString {
		_Product_buf = append(_Product_buf, uint16(0))
	}
	for i1 := range _Product_buf {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(_Product_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Product_buf); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ChangerTypeInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	var _Vendor_buf []uint16
	_Vendor_buf = make([]uint16, 128)
	for i1 := range _Vendor_buf {
		i1 := i1
		if err := w.ReadData(&_Vendor_buf[i1]); err != nil {
			return err
		}
	}
	o.Vendor = strings.TrimRight(string(utf16.Decode(_Vendor_buf)), ndr.ZeroString)
	var _Product_buf []uint16
	_Product_buf = make([]uint16, 128)
	for i1 := range _Product_buf {
		i1 := i1
		if err := w.ReadData(&_Product_buf[i1]); err != nil {
			return err
		}
	}
	o.Product = strings.TrimRight(string(utf16.Decode(_Product_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// DriveInformationA structure represents NTMS_DRIVEINFORMATIONA RPC structure.
type DriveInformationA struct {
	Number             uint32           `idl:"name:Number" json:"number"`
	State              uint32           `idl:"name:State" json:"state"`
	DriveType          *GUID            `idl:"name:DriveType" json:"drive_type"`
	DeviceName         []byte           `idl:"name:szDeviceName" json:"device_name"`
	SerialNumber       []byte           `idl:"name:szSerialNumber" json:"serial_number"`
	Revision           []byte           `idl:"name:szRevision" json:"revision"`
	SCSIPort           uint16           `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus            uint16           `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget         uint16           `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN            uint16           `idl:"name:ScsiLun" json:"scsi_lun"`
	MountCount         uint32           `idl:"name:dwMountCount" json:"mount_count"`
	LastCleanedTS      *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	SavedPartitionID   *GUID            `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	Library            *GUID            `idl:"name:Library" json:"library"`
	_                  *dtyp.GUID       `idl:"name:Reserved"`
	DeferDismountDelay uint32           `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
}

func (o *DriveInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DriveInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.DeviceName {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.DeviceName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DeviceName); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.SerialNumber {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.SerialNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SerialNumber); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Revision {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.Revision[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Revision); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SCSIPort); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSIBus); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSITarget); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSILUN); err != nil {
		return err
	}
	if err := w.WriteData(o.MountCount); err != nil {
		return err
	}
	if o.LastCleanedTS != nil {
		if err := o.LastCleanedTS.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SavedPartitionID != nil {
		if err := o.SavedPartitionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved Reserved
	if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.WriteData(o.DeferDismountDelay); err != nil {
		return err
	}
	return nil
}
func (o *DriveInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.DriveType == nil {
		o.DriveType = &GUID{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.DeviceName = make([]byte, 64)
	for i1 := range o.DeviceName {
		i1 := i1
		if err := w.ReadData(&o.DeviceName[i1]); err != nil {
			return err
		}
	}
	o.SerialNumber = make([]byte, 32)
	for i1 := range o.SerialNumber {
		i1 := i1
		if err := w.ReadData(&o.SerialNumber[i1]); err != nil {
			return err
		}
	}
	o.Revision = make([]byte, 32)
	for i1 := range o.Revision {
		i1 := i1
		if err := w.ReadData(&o.Revision[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.SCSIPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSIBus); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSITarget); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSILUN); err != nil {
		return err
	}
	if err := w.ReadData(&o.MountCount); err != nil {
		return err
	}
	if o.LastCleanedTS == nil {
		o.LastCleanedTS = &dtyp.SystemTime{}
	}
	if err := o.LastCleanedTS.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SavedPartitionID == nil {
		o.SavedPartitionID = &GUID{}
	}
	if err := o.SavedPartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved *dtyp.GUID
	if _Reserved == nil {
		_Reserved = &dtyp.GUID{}
	}
	if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeferDismountDelay); err != nil {
		return err
	}
	return nil
}

// DriveInformationW structure represents NTMS_DRIVEINFORMATIONW RPC structure.
type DriveInformationW struct {
	Number             uint32           `idl:"name:Number" json:"number"`
	State              uint32           `idl:"name:State" json:"state"`
	DriveType          *GUID            `idl:"name:DriveType" json:"drive_type"`
	DeviceName         string           `idl:"name:szDeviceName;string" json:"device_name"`
	SerialNumber       string           `idl:"name:szSerialNumber;string" json:"serial_number"`
	Revision           string           `idl:"name:szRevision;string" json:"revision"`
	SCSIPort           uint16           `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus            uint16           `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget         uint16           `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN            uint16           `idl:"name:ScsiLun" json:"scsi_lun"`
	MountCount         uint32           `idl:"name:dwMountCount" json:"mount_count"`
	LastCleanedTS      *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	SavedPartitionID   *GUID            `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	Library            *GUID            `idl:"name:Library" json:"library"`
	_                  *dtyp.GUID       `idl:"name:Reserved"`
	DeferDismountDelay uint32           `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
}

func (o *DriveInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DriveInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_DeviceName_buf := utf16.Encode([]rune(o.DeviceName))
	if uint64(len(_DeviceName_buf)) > 64-1 {
		_DeviceName_buf = _DeviceName_buf[:64-1]
	}
	if o.DeviceName != ndr.ZeroString {
		_DeviceName_buf = append(_DeviceName_buf, uint16(0))
	}
	for i1 := range _DeviceName_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_DeviceName_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_DeviceName_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_SerialNumber_buf := utf16.Encode([]rune(o.SerialNumber))
	if uint64(len(_SerialNumber_buf)) > 32-1 {
		_SerialNumber_buf = _SerialNumber_buf[:32-1]
	}
	if o.SerialNumber != ndr.ZeroString {
		_SerialNumber_buf = append(_SerialNumber_buf, uint16(0))
	}
	for i1 := range _SerialNumber_buf {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(_SerialNumber_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_SerialNumber_buf); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Revision_buf := utf16.Encode([]rune(o.Revision))
	if uint64(len(_Revision_buf)) > 32-1 {
		_Revision_buf = _Revision_buf[:32-1]
	}
	if o.Revision != ndr.ZeroString {
		_Revision_buf = append(_Revision_buf, uint16(0))
	}
	for i1 := range _Revision_buf {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(_Revision_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Revision_buf); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SCSIPort); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSIBus); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSITarget); err != nil {
		return err
	}
	if err := w.WriteData(o.SCSILUN); err != nil {
		return err
	}
	if err := w.WriteData(o.MountCount); err != nil {
		return err
	}
	if o.LastCleanedTS != nil {
		if err := o.LastCleanedTS.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SavedPartitionID != nil {
		if err := o.SavedPartitionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved Reserved
	if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.WriteData(o.DeferDismountDelay); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *DriveInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.DriveType == nil {
		o.DriveType = &GUID{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _DeviceName_buf []uint16
	_DeviceName_buf = make([]uint16, 64)
	for i1 := range _DeviceName_buf {
		i1 := i1
		if err := w.ReadData(&_DeviceName_buf[i1]); err != nil {
			return err
		}
	}
	o.DeviceName = strings.TrimRight(string(utf16.Decode(_DeviceName_buf)), ndr.ZeroString)
	var _SerialNumber_buf []uint16
	_SerialNumber_buf = make([]uint16, 32)
	for i1 := range _SerialNumber_buf {
		i1 := i1
		if err := w.ReadData(&_SerialNumber_buf[i1]); err != nil {
			return err
		}
	}
	o.SerialNumber = strings.TrimRight(string(utf16.Decode(_SerialNumber_buf)), ndr.ZeroString)
	var _Revision_buf []uint16
	_Revision_buf = make([]uint16, 32)
	for i1 := range _Revision_buf {
		i1 := i1
		if err := w.ReadData(&_Revision_buf[i1]); err != nil {
			return err
		}
	}
	o.Revision = strings.TrimRight(string(utf16.Decode(_Revision_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.SCSIPort); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSIBus); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSITarget); err != nil {
		return err
	}
	if err := w.ReadData(&o.SCSILUN); err != nil {
		return err
	}
	if err := w.ReadData(&o.MountCount); err != nil {
		return err
	}
	if o.LastCleanedTS == nil {
		o.LastCleanedTS = &dtyp.SystemTime{}
	}
	if err := o.LastCleanedTS.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SavedPartitionID == nil {
		o.SavedPartitionID = &GUID{}
	}
	if err := o.SavedPartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved *dtyp.GUID
	if _Reserved == nil {
		_Reserved = &dtyp.GUID{}
	}
	if err := _Reserved.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeferDismountDelay); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// DriveTypeInformationA structure represents NTMS_DRIVETYPEINFORMATIONA RPC structure.
type DriveTypeInformationA struct {
	Vendor        []byte `idl:"name:szVendor" json:"vendor"`
	Product       []byte `idl:"name:szProduct" json:"product"`
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	DeviceType    uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *DriveTypeInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DriveTypeInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Vendor {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(o.Vendor[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Vendor); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Product {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(o.Product[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Product); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfHeads); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	return nil
}
func (o *DriveTypeInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Vendor = make([]byte, 128)
	for i1 := range o.Vendor {
		i1 := i1
		if err := w.ReadData(&o.Vendor[i1]); err != nil {
			return err
		}
	}
	o.Product = make([]byte, 128)
	for i1 := range o.Product {
		i1 := i1
		if err := w.ReadData(&o.Product[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.NumberOfHeads); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	return nil
}

// DriveTypeInformationW structure represents NTMS_DRIVETYPEINFORMATIONW RPC structure.
type DriveTypeInformationW struct {
	Vendor        string `idl:"name:szVendor;string" json:"vendor"`
	Product       string `idl:"name:szProduct;string" json:"product"`
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	DeviceType    uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *DriveTypeInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DriveTypeInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	_Vendor_buf := utf16.Encode([]rune(o.Vendor))
	if uint64(len(_Vendor_buf)) > 128-1 {
		_Vendor_buf = _Vendor_buf[:128-1]
	}
	if o.Vendor != ndr.ZeroString {
		_Vendor_buf = append(_Vendor_buf, uint16(0))
	}
	for i1 := range _Vendor_buf {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(_Vendor_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Vendor_buf); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Product_buf := utf16.Encode([]rune(o.Product))
	if uint64(len(_Product_buf)) > 128-1 {
		_Product_buf = _Product_buf[:128-1]
	}
	if o.Product != ndr.ZeroString {
		_Product_buf = append(_Product_buf, uint16(0))
	}
	for i1 := range _Product_buf {
		i1 := i1
		if uint64(i1) >= 128 {
			break
		}
		if err := w.WriteData(_Product_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Product_buf); uint64(i1) < 128; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfHeads); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *DriveTypeInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	var _Vendor_buf []uint16
	_Vendor_buf = make([]uint16, 128)
	for i1 := range _Vendor_buf {
		i1 := i1
		if err := w.ReadData(&_Vendor_buf[i1]); err != nil {
			return err
		}
	}
	o.Vendor = strings.TrimRight(string(utf16.Decode(_Vendor_buf)), ndr.ZeroString)
	var _Product_buf []uint16
	_Product_buf = make([]uint16, 128)
	for i1 := range _Product_buf {
		i1 := i1
		if err := w.ReadData(&_Product_buf[i1]); err != nil {
			return err
		}
	}
	o.Product = strings.TrimRight(string(utf16.Decode(_Product_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.NumberOfHeads); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// LibraryRequestInformationA structure represents NTMS_LIBREQUESTINFORMATIONA RPC structure.
type LibraryRequestInformationA struct {
	OperationCode   uint32           `idl:"name:OperationCode" json:"operation_code"`
	OperationOption uint32           `idl:"name:OperationOption" json:"operation_option"`
	State           uint32           `idl:"name:State" json:"state"`
	PartitionID     *GUID            `idl:"name:PartitionId" json:"partition_id"`
	DriveID         *GUID            `idl:"name:DriveId" json:"drive_id"`
	PhysicalMediaID *GUID            `idl:"name:PhysMediaId" json:"physical_media_id"`
	Library         *GUID            `idl:"name:Library" json:"library"`
	SlotID          *GUID            `idl:"name:SlotId" json:"slot_id"`
	TimeQueued      *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	TimeCompleted   *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	Application     []byte           `idl:"name:szApplication" json:"application"`
	User            []byte           `idl:"name:szUser" json:"user"`
	Computer        []byte           `idl:"name:szComputer" json:"computer"`
	ErrorCode       uint32           `idl:"name:dwErrorCode" json:"error_code"`
	WorkItemID      *GUID            `idl:"name:WorkItemId" json:"work_item_id"`
	Priority        uint32           `idl:"name:dwPriority" json:"priority"`
}

func (o *LibraryRequestInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *LibraryRequestInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationCode); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationOption); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.PartitionID != nil {
		if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DriveID != nil {
		if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PhysicalMediaID != nil {
		if err := o.PhysicalMediaID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SlotID != nil {
		if err := o.SlotID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TimeQueued != nil {
		if err := o.TimeQueued.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TimeCompleted != nil {
		if err := o.TimeCompleted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Application {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Application[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Application); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.User {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.User[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.User); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Computer {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Computer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Computer); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ErrorCode); err != nil {
		return err
	}
	if o.WorkItemID != nil {
		if err := o.WorkItemID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	return nil
}
func (o *LibraryRequestInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationOption); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.PartitionID == nil {
		o.PartitionID = &GUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DriveID == nil {
		o.DriveID = &GUID{}
	}
	if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PhysicalMediaID == nil {
		o.PhysicalMediaID = &GUID{}
	}
	if err := o.PhysicalMediaID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SlotID == nil {
		o.SlotID = &GUID{}
	}
	if err := o.SlotID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TimeQueued == nil {
		o.TimeQueued = &dtyp.SystemTime{}
	}
	if err := o.TimeQueued.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TimeCompleted == nil {
		o.TimeCompleted = &dtyp.SystemTime{}
	}
	if err := o.TimeCompleted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.Application = make([]byte, 64)
	for i1 := range o.Application {
		i1 := i1
		if err := w.ReadData(&o.Application[i1]); err != nil {
			return err
		}
	}
	o.User = make([]byte, 64)
	for i1 := range o.User {
		i1 := i1
		if err := w.ReadData(&o.User[i1]); err != nil {
			return err
		}
	}
	o.Computer = make([]byte, 64)
	for i1 := range o.Computer {
		i1 := i1
		if err := w.ReadData(&o.Computer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.ErrorCode); err != nil {
		return err
	}
	if o.WorkItemID == nil {
		o.WorkItemID = &GUID{}
	}
	if err := o.WorkItemID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	return nil
}

// LibraryRequestInformationW structure represents NTMS_LIBREQUESTINFORMATIONW RPC structure.
type LibraryRequestInformationW struct {
	OperationCode   uint32           `idl:"name:OperationCode" json:"operation_code"`
	OperationOption uint32           `idl:"name:OperationOption" json:"operation_option"`
	State           uint32           `idl:"name:State" json:"state"`
	PartitionID     *GUID            `idl:"name:PartitionId" json:"partition_id"`
	DriveID         *GUID            `idl:"name:DriveId" json:"drive_id"`
	PhysicalMediaID *GUID            `idl:"name:PhysMediaId" json:"physical_media_id"`
	Library         *GUID            `idl:"name:Library" json:"library"`
	SlotID          *GUID            `idl:"name:SlotId" json:"slot_id"`
	TimeQueued      *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	TimeCompleted   *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	Application     string           `idl:"name:szApplication;string" json:"application"`
	User            string           `idl:"name:szUser;string" json:"user"`
	Computer        string           `idl:"name:szComputer;string" json:"computer"`
	ErrorCode       uint32           `idl:"name:dwErrorCode" json:"error_code"`
	WorkItemID      *GUID            `idl:"name:WorkItemId" json:"work_item_id"`
	Priority        uint32           `idl:"name:dwPriority" json:"priority"`
}

func (o *LibraryRequestInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *LibraryRequestInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationCode); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationOption); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.PartitionID != nil {
		if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DriveID != nil {
		if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PhysicalMediaID != nil {
		if err := o.PhysicalMediaID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SlotID != nil {
		if err := o.SlotID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TimeQueued != nil {
		if err := o.TimeQueued.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TimeCompleted != nil {
		if err := o.TimeCompleted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_Application_buf := utf16.Encode([]rune(o.Application))
	if uint64(len(_Application_buf)) > 64-1 {
		_Application_buf = _Application_buf[:64-1]
	}
	if o.Application != ndr.ZeroString {
		_Application_buf = append(_Application_buf, uint16(0))
	}
	for i1 := range _Application_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_Application_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Application_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_User_buf := utf16.Encode([]rune(o.User))
	if uint64(len(_User_buf)) > 64-1 {
		_User_buf = _User_buf[:64-1]
	}
	if o.User != ndr.ZeroString {
		_User_buf = append(_User_buf, uint16(0))
	}
	for i1 := range _User_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_User_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_User_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Computer_buf := utf16.Encode([]rune(o.Computer))
	if uint64(len(_Computer_buf)) > 64-1 {
		_Computer_buf = _Computer_buf[:64-1]
	}
	if o.Computer != ndr.ZeroString {
		_Computer_buf = append(_Computer_buf, uint16(0))
	}
	for i1 := range _Computer_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_Computer_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Computer_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ErrorCode); err != nil {
		return err
	}
	if o.WorkItemID != nil {
		if err := o.WorkItemID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *LibraryRequestInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationOption); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.PartitionID == nil {
		o.PartitionID = &GUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DriveID == nil {
		o.DriveID = &GUID{}
	}
	if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PhysicalMediaID == nil {
		o.PhysicalMediaID = &GUID{}
	}
	if err := o.PhysicalMediaID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SlotID == nil {
		o.SlotID = &GUID{}
	}
	if err := o.SlotID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TimeQueued == nil {
		o.TimeQueued = &dtyp.SystemTime{}
	}
	if err := o.TimeQueued.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TimeCompleted == nil {
		o.TimeCompleted = &dtyp.SystemTime{}
	}
	if err := o.TimeCompleted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _Application_buf []uint16
	_Application_buf = make([]uint16, 64)
	for i1 := range _Application_buf {
		i1 := i1
		if err := w.ReadData(&_Application_buf[i1]); err != nil {
			return err
		}
	}
	o.Application = strings.TrimRight(string(utf16.Decode(_Application_buf)), ndr.ZeroString)
	var _User_buf []uint16
	_User_buf = make([]uint16, 64)
	for i1 := range _User_buf {
		i1 := i1
		if err := w.ReadData(&_User_buf[i1]); err != nil {
			return err
		}
	}
	o.User = strings.TrimRight(string(utf16.Decode(_User_buf)), ndr.ZeroString)
	var _Computer_buf []uint16
	_Computer_buf = make([]uint16, 64)
	for i1 := range _Computer_buf {
		i1 := i1
		if err := w.ReadData(&_Computer_buf[i1]); err != nil {
			return err
		}
	}
	o.Computer = strings.TrimRight(string(utf16.Decode(_Computer_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.ErrorCode); err != nil {
		return err
	}
	if o.WorkItemID == nil {
		o.WorkItemID = &GUID{}
	}
	if err := o.WorkItemID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// MediaPoolInformation structure represents NTMS_MEDIAPOOLINFORMATION RPC structure.
type MediaPoolInformation struct {
	PoolType              uint32 `idl:"name:PoolType" json:"pool_type"`
	MediaType             *GUID  `idl:"name:MediaType" json:"media_type"`
	Parent                *GUID  `idl:"name:Parent" json:"parent"`
	AllocationPolicy      uint32 `idl:"name:AllocationPolicy" json:"allocation_policy"`
	DeallocationPolicy    uint32 `idl:"name:DeallocationPolicy" json:"deallocation_policy"`
	MaxAllocates          uint32 `idl:"name:dwMaxAllocates" json:"max_allocates"`
	NumberOfPhysicalMedia uint32 `idl:"name:dwNumberOfPhysicalMedia" json:"number_of_physical_media"`
	NumberOfLogicalMedia  uint32 `idl:"name:dwNumberOfLogicalMedia" json:"number_of_logical_media"`
	NumberOfMediaPools    uint32 `idl:"name:dwNumberOfMediaPools" json:"number_of_media_pools"`
}

func (o *MediaPoolInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MediaPoolInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PoolType); err != nil {
		return err
	}
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Parent != nil {
		if err := o.Parent.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AllocationPolicy); err != nil {
		return err
	}
	if err := w.WriteData(o.DeallocationPolicy); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAllocates); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfPhysicalMedia); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfLogicalMedia); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfMediaPools); err != nil {
		return err
	}
	return nil
}
func (o *MediaPoolInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PoolType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &GUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Parent == nil {
		o.Parent = &GUID{}
	}
	if err := o.Parent.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocationPolicy); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeallocationPolicy); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAllocates); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPhysicalMedia); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfLogicalMedia); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfMediaPools); err != nil {
		return err
	}
	return nil
}

// MediaTypeInformation structure represents NTMS_MEDIATYPEINFORMATION RPC structure.
type MediaTypeInformation struct {
	MediaType                uint32 `idl:"name:MediaType" json:"media_type"`
	NumberOfSides            uint32 `idl:"name:NumberOfSides" json:"number_of_sides"`
	ReadWriteCharacteristics uint32 `idl:"name:ReadWriteCharacteristics" json:"read_write_characteristics"`
	DeviceType               uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *MediaTypeInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MediaTypeInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaType); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfSides); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadWriteCharacteristics); err != nil {
		return err
	}
	if err := w.WriteData(o.DeviceType); err != nil {
		return err
	}
	return nil
}
func (o *MediaTypeInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaType); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfSides); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadWriteCharacteristics); err != nil {
		return err
	}
	if err := w.ReadData(&o.DeviceType); err != nil {
		return err
	}
	return nil
}

// StorageSlotInformation structure represents NTMS_STORAGESLOTINFORMATION RPC structure.
type StorageSlotInformation struct {
	Number  uint32 `idl:"name:Number" json:"number"`
	State   uint32 `idl:"name:State" json:"state"`
	Library *GUID  `idl:"name:Library" json:"library"`
}

func (o *StorageSlotInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StorageSlotInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageSlotInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IEDoorInformation structure represents NTMS_IEDOORINFORMATION RPC structure.
type IEDoorInformation struct {
	Number      uint32 `idl:"name:Number" json:"number"`
	State       uint32 `idl:"name:State" json:"state"`
	MaxOpenSecs uint16 `idl:"name:MaxOpenSecs" json:"max_open_secs"`
	Library     *GUID  `idl:"name:Library" json:"library"`
}

func (o *IEDoorInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IEDoorInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxOpenSecs); err != nil {
		return err
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IEDoorInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxOpenSecs); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// IEPortInformation structure represents NTMS_IEPORTINFORMATION RPC structure.
type IEPortInformation struct {
	Number        uint32 `idl:"name:Number" json:"number"`
	Content       uint32 `idl:"name:Content" json:"content"`
	Position      uint32 `idl:"name:Position" json:"position"`
	MaxExtendSecs uint16 `idl:"name:MaxExtendSecs" json:"max_extend_secs"`
	Library       *GUID  `idl:"name:Library" json:"library"`
}

func (o *IEPortInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IEPortInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Number); err != nil {
		return err
	}
	if err := w.WriteData(o.Content); err != nil {
		return err
	}
	if err := w.WriteData(o.Position); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxExtendSecs); err != nil {
		return err
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IEPortInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if err := w.ReadData(&o.Content); err != nil {
		return err
	}
	if err := w.ReadData(&o.Position); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxExtendSecs); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &GUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// LMIDInformation structure represents NTMS_LMIDINFORMATION RPC structure.
type LMIDInformation struct {
	MediaPool          *GUID  `idl:"name:MediaPool" json:"media_pool"`
	NumberOfPartitions uint32 `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
}

func (o *LMIDInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *LMIDInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfPartitions); err != nil {
		return err
	}
	return nil
}
func (o *LMIDInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &GUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPartitions); err != nil {
		return err
	}
	return nil
}

// ComputerInformation structure represents NTMS_COMPUTERINFORMATION RPC structure.
type ComputerInformation struct {
	LibRequestPurgeTime       uint32 `idl:"name:dwLibRequestPurgeTime" json:"lib_request_purge_time"`
	OperationRequestPurgeTime uint32 `idl:"name:dwOpRequestPurgeTime" json:"operation_request_purge_time"`
	LibRequestFlags           uint32 `idl:"name:dwLibRequestFlags" json:"lib_request_flags"`
	OperationRequestFlags     uint32 `idl:"name:dwOpRequestFlags" json:"operation_request_flags"`
	MediaPoolPolicy           uint32 `idl:"name:dwMediaPoolPolicy" json:"media_pool_policy"`
}

func (o *ComputerInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ComputerInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.LibRequestPurgeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationRequestPurgeTime); err != nil {
		return err
	}
	if err := w.WriteData(o.LibRequestFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.OperationRequestFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaPoolPolicy); err != nil {
		return err
	}
	return nil
}
func (o *ComputerInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LibRequestPurgeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationRequestPurgeTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.LibRequestFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.OperationRequestFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaPoolPolicy); err != nil {
		return err
	}
	return nil
}

// OperationRequestInformationA structure represents NTMS_OPREQUESTINFORMATIONA RPC structure.
type OperationRequestInformationA struct {
	Request     uint32           `idl:"name:Request" json:"request"`
	Submitted   *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	State       uint32           `idl:"name:State" json:"state"`
	Message     []byte           `idl:"name:szMessage" json:"message"`
	Arg1Type    uint32           `idl:"name:Arg1Type" json:"arg1_type"`
	Arg1        *GUID            `idl:"name:Arg1" json:"arg1"`
	Arg2Type    uint32           `idl:"name:Arg2Type" json:"arg2_type"`
	Arg2        *GUID            `idl:"name:Arg2" json:"arg2"`
	Application []byte           `idl:"name:szApplication" json:"application"`
	User        []byte           `idl:"name:szUser" json:"user"`
	Computer    []byte           `idl:"name:szComputer" json:"computer"`
}

func (o *OperationRequestInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OperationRequestInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Request); err != nil {
		return err
	}
	if o.Submitted != nil {
		if err := o.Submitted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	for i1 := range o.Message {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.Message[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Message); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Arg1Type); err != nil {
		return err
	}
	if o.Arg1 != nil {
		if err := o.Arg1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 != nil {
		if err := o.Arg2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Application {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Application[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Application); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.User {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.User[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.User); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Computer {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Computer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Computer); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *OperationRequestInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Request); err != nil {
		return err
	}
	if o.Submitted == nil {
		o.Submitted = &dtyp.SystemTime{}
	}
	if err := o.Submitted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	o.Message = make([]byte, 256)
	for i1 := range o.Message {
		i1 := i1
		if err := w.ReadData(&o.Message[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Arg1Type); err != nil {
		return err
	}
	if o.Arg1 == nil {
		o.Arg1 = &GUID{}
	}
	if err := o.Arg1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 == nil {
		o.Arg2 = &GUID{}
	}
	if err := o.Arg2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.Application = make([]byte, 64)
	for i1 := range o.Application {
		i1 := i1
		if err := w.ReadData(&o.Application[i1]); err != nil {
			return err
		}
	}
	o.User = make([]byte, 64)
	for i1 := range o.User {
		i1 := i1
		if err := w.ReadData(&o.User[i1]); err != nil {
			return err
		}
	}
	o.Computer = make([]byte, 64)
	for i1 := range o.Computer {
		i1 := i1
		if err := w.ReadData(&o.Computer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// OperationRequestInformationW structure represents NTMS_OPREQUESTINFORMATIONW RPC structure.
type OperationRequestInformationW struct {
	Request     uint32           `idl:"name:Request" json:"request"`
	Submitted   *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	State       uint32           `idl:"name:State" json:"state"`
	Message     string           `idl:"name:szMessage;string" json:"message"`
	Arg1Type    uint32           `idl:"name:Arg1Type" json:"arg1_type"`
	Arg1        *GUID            `idl:"name:Arg1" json:"arg1"`
	Arg2Type    uint32           `idl:"name:Arg2Type" json:"arg2_type"`
	Arg2        *GUID            `idl:"name:Arg2" json:"arg2"`
	Application string           `idl:"name:szApplication;string" json:"application"`
	User        string           `idl:"name:szUser;string" json:"user"`
	Computer    string           `idl:"name:szComputer;string" json:"computer"`
}

func (o *OperationRequestInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *OperationRequestInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Request); err != nil {
		return err
	}
	if o.Submitted != nil {
		if err := o.Submitted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	_Message_buf := utf16.Encode([]rune(o.Message))
	if uint64(len(_Message_buf)) > 256-1 {
		_Message_buf = _Message_buf[:256-1]
	}
	if o.Message != ndr.ZeroString {
		_Message_buf = append(_Message_buf, uint16(0))
	}
	for i1 := range _Message_buf {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(_Message_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Message_buf); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Arg1Type); err != nil {
		return err
	}
	if o.Arg1 != nil {
		if err := o.Arg1.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 != nil {
		if err := o.Arg2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_Application_buf := utf16.Encode([]rune(o.Application))
	if uint64(len(_Application_buf)) > 64-1 {
		_Application_buf = _Application_buf[:64-1]
	}
	if o.Application != ndr.ZeroString {
		_Application_buf = append(_Application_buf, uint16(0))
	}
	for i1 := range _Application_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_Application_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Application_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_User_buf := utf16.Encode([]rune(o.User))
	if uint64(len(_User_buf)) > 64-1 {
		_User_buf = _User_buf[:64-1]
	}
	if o.User != ndr.ZeroString {
		_User_buf = append(_User_buf, uint16(0))
	}
	for i1 := range _User_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_User_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_User_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	sizeInfo := []uint64{
		0,
	}
	dimLength1 := ndr.UTF16NLen(o.Computer)
	sizeInfo[0] = dimLength1
	if err := w.WriteSize(0); err != nil {
		return err
	}
	if err := w.WriteSize(dimLength1); err != nil {
		return err
	}
	_Computer_buf := utf16.Encode([]rune(o.Computer))
	if uint64(len(_Computer_buf)) > 64-1 {
		_Computer_buf = _Computer_buf[:64-1]
	}
	if o.Computer != ndr.ZeroString {
		_Computer_buf = append(_Computer_buf, uint16(0))
	}
	for i1 := range _Computer_buf {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(_Computer_buf[i1]); err != nil {
			return err
		}
	}
	return nil
}
func (o *OperationRequestInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Request); err != nil {
		return err
	}
	if o.Submitted == nil {
		o.Submitted = &dtyp.SystemTime{}
	}
	if err := o.Submitted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	var _Message_buf []uint16
	_Message_buf = make([]uint16, 256)
	for i1 := range _Message_buf {
		i1 := i1
		if err := w.ReadData(&_Message_buf[i1]); err != nil {
			return err
		}
	}
	o.Message = strings.TrimRight(string(utf16.Decode(_Message_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.Arg1Type); err != nil {
		return err
	}
	if o.Arg1 == nil {
		o.Arg1 = &GUID{}
	}
	if err := o.Arg1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 == nil {
		o.Arg2 = &GUID{}
	}
	if err := o.Arg2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _Application_buf []uint16
	_Application_buf = make([]uint16, 64)
	for i1 := range _Application_buf {
		i1 := i1
		if err := w.ReadData(&_Application_buf[i1]); err != nil {
			return err
		}
	}
	o.Application = strings.TrimRight(string(utf16.Decode(_Application_buf)), ndr.ZeroString)
	var _User_buf []uint16
	_User_buf = make([]uint16, 64)
	for i1 := range _User_buf {
		i1 := i1
		if err := w.ReadData(&_User_buf[i1]); err != nil {
			return err
		}
	}
	o.User = strings.TrimRight(string(utf16.Decode(_User_buf)), ndr.ZeroString)
	sizeInfo := []uint64{
		0,
	}
	for sz1 := range sizeInfo {
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
		if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
			return err
		}
	}
	var _Computer_buf []uint16
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Computer_buf", sizeInfo[0])
	}
	_Computer_buf = make([]uint16, sizeInfo[0])
	for i1 := range _Computer_buf {
		i1 := i1
		if err := w.ReadData(&_Computer_buf[i1]); err != nil {
			return err
		}
	}
	o.Computer = strings.TrimRight(string(utf16.Decode(_Computer_buf)), ndr.ZeroString)
	return nil
}

// PartitionInformationA structure represents NTMS_PARTITIONINFORMATIONA RPC structure.
type PartitionInformationA struct {
	PhysicalMedia     *GUID              `idl:"name:PhysicalMedia" json:"physical_media"`
	LogicalMedia      *GUID              `idl:"name:LogicalMedia" json:"logical_media"`
	State             uint32             `idl:"name:State" json:"state"`
	Side              uint16             `idl:"name:Side" json:"side"`
	OMIDLabelIDLength uint32             `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	OMIDLabelID       []byte             `idl:"name:OmidLabelId" json:"omid_label_id"`
	OMIDLabelType     []byte             `idl:"name:szOmidLabelType" json:"omid_label_type"`
	OMIDLabelInfo     []byte             `idl:"name:szOmidLabelInfo" json:"omid_label_info"`
	MountCount        uint32             `idl:"name:dwMountCount" json:"mount_count"`
	AllocateCount     uint32             `idl:"name:dwAllocateCount" json:"allocate_count"`
	Capacity          *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
}

func (o *PartitionInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PartitionInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Side); err != nil {
		return err
	}
	if err := w.WriteData(o.OMIDLabelIDLength); err != nil {
		return err
	}
	for i1 := range o.OMIDLabelID {
		i1 := i1
		if uint64(i1) >= 255 {
			break
		}
		if err := w.WriteData(o.OMIDLabelID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OMIDLabelID); uint64(i1) < 255; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.OMIDLabelType {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.OMIDLabelType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OMIDLabelType); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.OMIDLabelInfo {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.OMIDLabelInfo[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OMIDLabelInfo); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MountCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocateCount); err != nil {
		return err
	}
	if o.Capacity != nil {
		if err := o.Capacity.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &GUID{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogicalMedia == nil {
		o.LogicalMedia = &GUID{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Side); err != nil {
		return err
	}
	if err := w.ReadData(&o.OMIDLabelIDLength); err != nil {
		return err
	}
	o.OMIDLabelID = make([]byte, 255)
	for i1 := range o.OMIDLabelID {
		i1 := i1
		if err := w.ReadData(&o.OMIDLabelID[i1]); err != nil {
			return err
		}
	}
	o.OMIDLabelType = make([]byte, 64)
	for i1 := range o.OMIDLabelType {
		i1 := i1
		if err := w.ReadData(&o.OMIDLabelType[i1]); err != nil {
			return err
		}
	}
	o.OMIDLabelInfo = make([]byte, 256)
	for i1 := range o.OMIDLabelInfo {
		i1 := i1
		if err := w.ReadData(&o.OMIDLabelInfo[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.MountCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocateCount); err != nil {
		return err
	}
	if o.Capacity == nil {
		o.Capacity = &dtyp.LargeInteger{}
	}
	if err := o.Capacity.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PartitionInformationW structure represents NTMS_PARTITIONINFORMATIONW RPC structure.
type PartitionInformationW struct {
	PhysicalMedia     *GUID              `idl:"name:PhysicalMedia" json:"physical_media"`
	LogicalMedia      *GUID              `idl:"name:LogicalMedia" json:"logical_media"`
	State             uint32             `idl:"name:State" json:"state"`
	Side              uint16             `idl:"name:Side" json:"side"`
	OMIDLabelIDLength uint32             `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	OMIDLabelID       []byte             `idl:"name:OmidLabelId" json:"omid_label_id"`
	OMIDLabelType     string             `idl:"name:szOmidLabelType;string" json:"omid_label_type"`
	OMIDLabelInfo     string             `idl:"name:szOmidLabelInfo;string" json:"omid_label_info"`
	MountCount        uint32             `idl:"name:dwMountCount" json:"mount_count"`
	AllocateCount     uint32             `idl:"name:dwAllocateCount" json:"allocate_count"`
	Capacity          *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
}

func (o *PartitionInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PartitionInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Side); err != nil {
		return err
	}
	if err := w.WriteData(o.OMIDLabelIDLength); err != nil {
		return err
	}
	for i1 := range o.OMIDLabelID {
		i1 := i1
		if uint64(i1) >= 255 {
			break
		}
		if err := w.WriteData(o.OMIDLabelID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OMIDLabelID); uint64(i1) < 255; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	_OMIDLabelType_buf := utf16.Encode([]rune(o.OMIDLabelType))
	if uint64(len(_OMIDLabelType_buf)) > 64-1 {
		_OMIDLabelType_buf = _OMIDLabelType_buf[:64-1]
	}
	if o.OMIDLabelType != ndr.ZeroString {
		_OMIDLabelType_buf = append(_OMIDLabelType_buf, uint16(0))
	}
	for i1 := range _OMIDLabelType_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_OMIDLabelType_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_OMIDLabelType_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_OMIDLabelInfo_buf := utf16.Encode([]rune(o.OMIDLabelInfo))
	if uint64(len(_OMIDLabelInfo_buf)) > 256-1 {
		_OMIDLabelInfo_buf = _OMIDLabelInfo_buf[:256-1]
	}
	if o.OMIDLabelInfo != ndr.ZeroString {
		_OMIDLabelInfo_buf = append(_OMIDLabelInfo_buf, uint16(0))
	}
	for i1 := range _OMIDLabelInfo_buf {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(_OMIDLabelInfo_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_OMIDLabelInfo_buf); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MountCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocateCount); err != nil {
		return err
	}
	if o.Capacity != nil {
		if err := o.Capacity.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PartitionInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &GUID{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogicalMedia == nil {
		o.LogicalMedia = &GUID{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Side); err != nil {
		return err
	}
	if err := w.ReadData(&o.OMIDLabelIDLength); err != nil {
		return err
	}
	o.OMIDLabelID = make([]byte, 255)
	for i1 := range o.OMIDLabelID {
		i1 := i1
		if err := w.ReadData(&o.OMIDLabelID[i1]); err != nil {
			return err
		}
	}
	var _OMIDLabelType_buf []uint16
	_OMIDLabelType_buf = make([]uint16, 64)
	for i1 := range _OMIDLabelType_buf {
		i1 := i1
		if err := w.ReadData(&_OMIDLabelType_buf[i1]); err != nil {
			return err
		}
	}
	o.OMIDLabelType = strings.TrimRight(string(utf16.Decode(_OMIDLabelType_buf)), ndr.ZeroString)
	var _OMIDLabelInfo_buf []uint16
	_OMIDLabelInfo_buf = make([]uint16, 256)
	for i1 := range _OMIDLabelInfo_buf {
		i1 := i1
		if err := w.ReadData(&_OMIDLabelInfo_buf[i1]); err != nil {
			return err
		}
	}
	o.OMIDLabelInfo = strings.TrimRight(string(utf16.Decode(_OMIDLabelInfo_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.MountCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocateCount); err != nil {
		return err
	}
	if o.Capacity == nil {
		o.Capacity = &dtyp.LargeInteger{}
	}
	if err := o.Capacity.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PMIDInformationA structure represents NTMS_PMIDINFORMATIONA RPC structure.
type PMIDInformationA struct {
	CurrentLibrary     *GUID  `idl:"name:CurrentLibrary" json:"current_library"`
	MediaPool          *GUID  `idl:"name:MediaPool" json:"media_pool"`
	Location           *GUID  `idl:"name:Location" json:"location"`
	LocationType       uint32 `idl:"name:LocationType" json:"location_type"`
	MediaType          *GUID  `idl:"name:MediaType" json:"media_type"`
	HomeSlot           *GUID  `idl:"name:HomeSlot" json:"home_slot"`
	BarCode            []byte `idl:"name:szBarCode" json:"bar_code"`
	BarCodeState       uint32 `idl:"name:BarCodeState" json:"bar_code_state"`
	SequenceNumber     []byte `idl:"name:szSequenceNumber" json:"sequence_number"`
	MediaState         uint32 `idl:"name:MediaState" json:"media_state"`
	NumberOfPartitions uint32 `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	MediaTypeCode      uint32 `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	DensityCode        uint32 `idl:"name:dwDensityCode" json:"density_code"`
	MountedPartition   *GUID  `idl:"name:MountedPartition" json:"mounted_partition"`
}

func (o *PMIDInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PMIDInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.CurrentLibrary != nil {
		if err := o.CurrentLibrary.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Location != nil {
		if err := o.Location.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocationType); err != nil {
		return err
	}
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeSlot != nil {
		if err := o.HomeSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.BarCode {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.BarCode[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.BarCode); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BarCodeState); err != nil {
		return err
	}
	for i1 := range o.SequenceNumber {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(o.SequenceNumber[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SequenceNumber); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MediaState); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfPartitions); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaTypeCode); err != nil {
		return err
	}
	if err := w.WriteData(o.DensityCode); err != nil {
		return err
	}
	if o.MountedPartition != nil {
		if err := o.MountedPartition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PMIDInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.CurrentLibrary == nil {
		o.CurrentLibrary = &GUID{}
	}
	if err := o.CurrentLibrary.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &GUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Location == nil {
		o.Location = &GUID{}
	}
	if err := o.Location.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocationType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &GUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeSlot == nil {
		o.HomeSlot = &GUID{}
	}
	if err := o.HomeSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.BarCode = make([]byte, 64)
	for i1 := range o.BarCode {
		i1 := i1
		if err := w.ReadData(&o.BarCode[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.BarCodeState); err != nil {
		return err
	}
	o.SequenceNumber = make([]byte, 32)
	for i1 := range o.SequenceNumber {
		i1 := i1
		if err := w.ReadData(&o.SequenceNumber[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.MediaState); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPartitions); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaTypeCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.DensityCode); err != nil {
		return err
	}
	if o.MountedPartition == nil {
		o.MountedPartition = &GUID{}
	}
	if err := o.MountedPartition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PMIDInformationW structure represents NTMS_PMIDINFORMATIONW RPC structure.
type PMIDInformationW struct {
	CurrentLibrary     *GUID  `idl:"name:CurrentLibrary" json:"current_library"`
	MediaPool          *GUID  `idl:"name:MediaPool" json:"media_pool"`
	Location           *GUID  `idl:"name:Location" json:"location"`
	LocationType       uint32 `idl:"name:LocationType" json:"location_type"`
	MediaType          *GUID  `idl:"name:MediaType" json:"media_type"`
	HomeSlot           *GUID  `idl:"name:HomeSlot" json:"home_slot"`
	BarCode            string `idl:"name:szBarCode;string" json:"bar_code"`
	BarCodeState       uint32 `idl:"name:BarCodeState" json:"bar_code_state"`
	SequenceNumber     string `idl:"name:szSequenceNumber;string" json:"sequence_number"`
	MediaState         uint32 `idl:"name:MediaState" json:"media_state"`
	NumberOfPartitions uint32 `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	MediaTypeCode      uint32 `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	DensityCode        uint32 `idl:"name:dwDensityCode" json:"density_code"`
	MountedPartition   *GUID  `idl:"name:MountedPartition" json:"mounted_partition"`
}

func (o *PMIDInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PMIDInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.CurrentLibrary != nil {
		if err := o.CurrentLibrary.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Location != nil {
		if err := o.Location.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocationType); err != nil {
		return err
	}
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeSlot != nil {
		if err := o.HomeSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_BarCode_buf := utf16.Encode([]rune(o.BarCode))
	if uint64(len(_BarCode_buf)) > 64-1 {
		_BarCode_buf = _BarCode_buf[:64-1]
	}
	if o.BarCode != ndr.ZeroString {
		_BarCode_buf = append(_BarCode_buf, uint16(0))
	}
	for i1 := range _BarCode_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_BarCode_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_BarCode_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BarCodeState); err != nil {
		return err
	}
	_SequenceNumber_buf := utf16.Encode([]rune(o.SequenceNumber))
	if uint64(len(_SequenceNumber_buf)) > 32-1 {
		_SequenceNumber_buf = _SequenceNumber_buf[:32-1]
	}
	if o.SequenceNumber != ndr.ZeroString {
		_SequenceNumber_buf = append(_SequenceNumber_buf, uint16(0))
	}
	for i1 := range _SequenceNumber_buf {
		i1 := i1
		if uint64(i1) >= 32 {
			break
		}
		if err := w.WriteData(_SequenceNumber_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_SequenceNumber_buf); uint64(i1) < 32; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MediaState); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfPartitions); err != nil {
		return err
	}
	if err := w.WriteData(o.MediaTypeCode); err != nil {
		return err
	}
	if err := w.WriteData(o.DensityCode); err != nil {
		return err
	}
	if o.MountedPartition != nil {
		if err := o.MountedPartition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *PMIDInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.CurrentLibrary == nil {
		o.CurrentLibrary = &GUID{}
	}
	if err := o.CurrentLibrary.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &GUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Location == nil {
		o.Location = &GUID{}
	}
	if err := o.Location.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocationType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &GUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeSlot == nil {
		o.HomeSlot = &GUID{}
	}
	if err := o.HomeSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _BarCode_buf []uint16
	_BarCode_buf = make([]uint16, 64)
	for i1 := range _BarCode_buf {
		i1 := i1
		if err := w.ReadData(&_BarCode_buf[i1]); err != nil {
			return err
		}
	}
	o.BarCode = strings.TrimRight(string(utf16.Decode(_BarCode_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.BarCodeState); err != nil {
		return err
	}
	var _SequenceNumber_buf []uint16
	_SequenceNumber_buf = make([]uint16, 32)
	for i1 := range _SequenceNumber_buf {
		i1 := i1
		if err := w.ReadData(&_SequenceNumber_buf[i1]); err != nil {
			return err
		}
	}
	o.SequenceNumber = strings.TrimRight(string(utf16.Decode(_SequenceNumber_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.MediaState); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPartitions); err != nil {
		return err
	}
	if err := w.ReadData(&o.MediaTypeCode); err != nil {
		return err
	}
	if err := w.ReadData(&o.DensityCode); err != nil {
		return err
	}
	if o.MountedPartition == nil {
		o.MountedPartition = &GUID{}
	}
	if err := o.MountedPartition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RSMMessage structure represents RSM_MESSAGE RPC structure.
type RSMMessage struct {
	LpguidOperation *dtyp.GUID `idl:"name:lpguidOperation;pointer:unique" json:"lpguid_operation"`
	NTMSType        uint32     `idl:"name:dwNtmsType" json:"ntms_type"`
	State           uint32     `idl:"name:dwState" json:"state"`
	Flags           uint32     `idl:"name:dwFlags" json:"flags"`
	Priority        uint32     `idl:"name:dwPriority" json:"priority"`
	ErrorCode       uint32     `idl:"name:dwErrorCode" json:"error_code"`
	ComputerName    string     `idl:"name:lpszComputerName;string;pointer:unique" json:"computer_name"`
	Application     string     `idl:"name:lpszApplication;string" json:"application"`
	User            string     `idl:"name:lpszUser;string" json:"user"`
	TimeSubmitted   string     `idl:"name:lpszTimeSubmitted;string" json:"time_submitted"`
	Message         string     `idl:"name:lpszMessage;string" json:"message"`
}

func (o *RSMMessage) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RSMMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LpguidOperation != nil {
		_ptr_lpguidOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LpguidOperation != nil {
				if err := o.LpguidOperation.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LpguidOperation, _ptr_lpguidOperation); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NTMSType); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	if err := w.WriteData(o.ErrorCode); err != nil {
		return err
	}
	if o.ComputerName != "" {
		_ptr_lpszComputerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ComputerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ComputerName, _ptr_lpszComputerName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Application != "" {
		_ptr_lpszApplication := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Application); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Application, _ptr_lpszApplication); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.User != "" {
		_ptr_lpszUser := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.User); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.User, _ptr_lpszUser); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TimeSubmitted != "" {
		_ptr_lpszTimeSubmitted := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TimeSubmitted); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TimeSubmitted, _ptr_lpszTimeSubmitted); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Message != "" {
		_ptr_lpszMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Message); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Message, _ptr_lpszMessage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RSMMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_lpguidOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LpguidOperation == nil {
			o.LpguidOperation = &dtyp.GUID{}
		}
		if err := o.LpguidOperation.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpguidOperation := func(ptr interface{}) { o.LpguidOperation = *ptr.(**dtyp.GUID) }
	if err := w.ReadPointer(&o.LpguidOperation, _s_lpguidOperation, _ptr_lpguidOperation); err != nil {
		return err
	}
	if err := w.ReadData(&o.NTMSType); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	if err := w.ReadData(&o.ErrorCode); err != nil {
		return err
	}
	_ptr_lpszComputerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ComputerName); err != nil {
			return err
		}
		return nil
	})
	_s_lpszComputerName := func(ptr interface{}) { o.ComputerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ComputerName, _s_lpszComputerName, _ptr_lpszComputerName); err != nil {
		return err
	}
	_ptr_lpszApplication := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Application); err != nil {
			return err
		}
		return nil
	})
	_s_lpszApplication := func(ptr interface{}) { o.Application = *ptr.(*string) }
	if err := w.ReadPointer(&o.Application, _s_lpszApplication, _ptr_lpszApplication); err != nil {
		return err
	}
	_ptr_lpszUser := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.User); err != nil {
			return err
		}
		return nil
	})
	_s_lpszUser := func(ptr interface{}) { o.User = *ptr.(*string) }
	if err := w.ReadPointer(&o.User, _s_lpszUser, _ptr_lpszUser); err != nil {
		return err
	}
	_ptr_lpszTimeSubmitted := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TimeSubmitted); err != nil {
			return err
		}
		return nil
	})
	_s_lpszTimeSubmitted := func(ptr interface{}) { o.TimeSubmitted = *ptr.(*string) }
	if err := w.ReadPointer(&o.TimeSubmitted, _s_lpszTimeSubmitted, _ptr_lpszTimeSubmitted); err != nil {
		return err
	}
	_ptr_lpszMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Message); err != nil {
			return err
		}
		return nil
	})
	_s_lpszMessage := func(ptr interface{}) { o.Message = *ptr.(*string) }
	if err := w.ReadPointer(&o.Message, _s_lpszMessage, _ptr_lpszMessage); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA structure represents NTMS_OBJECTINFORMATIONA RPC structure.
type ObjectInformationA struct {
	Size             uint32                   `idl:"name:dwSize" json:"size"`
	Type             uint32                   `idl:"name:dwType" json:"type"`
	Created          *dtyp.SystemTime         `idl:"name:Created" json:"created"`
	Modified         *dtyp.SystemTime         `idl:"name:Modified" json:"modified"`
	ObjectGUID       *GUID                    `idl:"name:ObjectGuid" json:"object_guid"`
	Enabled          bool                     `idl:"name:Enabled" json:"enabled"`
	OperationalState uint32                   `idl:"name:dwOperationalState" json:"operational_state"`
	Name             []byte                   `idl:"name:szName" json:"name"`
	Description      []byte                   `idl:"name:szDescription" json:"description"`
	Info             *ObjectInformationA_Info `idl:"name:Info;switch_is:dwType" json:"info"`
}

func (o *ObjectInformationA) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectInformationA) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if o.Created != nil {
		if err := o.Created.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Modified != nil {
		if err := o.Modified.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ObjectGUID != nil {
		if err := o.ObjectGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.OperationalState); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Description {
		i1 := i1
		if uint64(i1) >= 127 {
			break
		}
		if err := w.WriteData(o.Description[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Description); uint64(i1) < 127; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	_swInfo := uint32(o.Type)
	if o.Info != nil {
		if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	} else {
		if err := (&ObjectInformationA_Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if o.Created == nil {
		o.Created = &dtyp.SystemTime{}
	}
	if err := o.Created.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Modified == nil {
		o.Modified = &dtyp.SystemTime{}
	}
	if err := o.Modified.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ObjectGUID == nil {
		o.ObjectGUID = &GUID{}
	}
	if err := o.ObjectGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadData(&o.OperationalState); err != nil {
		return err
	}
	o.Name = make([]byte, 64)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	o.Description = make([]byte, 127)
	for i1 := range o.Description {
		i1 := i1
		if err := w.ReadData(&o.Description[i1]); err != nil {
			return err
		}
	}
	if o.Info == nil {
		o.Info = &ObjectInformationA_Info{}
	}
	_swInfo := uint32(o.Type)
	if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_Info structure represents NTMS_OBJECTINFORMATIONA union anonymous member.
type ObjectInformationA_Info struct {
	// Types that are assignable to Value
	//
	// *ObjectInformationA_Drive
	// *ObjectInformationA_DriveType
	// *ObjectInformationA_Library
	// *ObjectInformationA_Changer
	// *ObjectInformationA_ChangerType
	// *ObjectInformationA_StorageSlot
	// *ObjectInformationA_IEDoor
	// *ObjectInformationA_IEPort
	// *ObjectInformationA_PhysicalMedia
	// *ObjectInformationA_LogicalMedia
	// *ObjectInformationA_Partition
	// *ObjectInformationA_MediaPool
	// *ObjectInformationA_MediaType
	// *ObjectInformationA_LibRequest
	// *ObjectInformationA_OperationRequest
	// *ObjectInformationA_Computer
	Value is_ObjectInformationA_Info `json:"value"`
}

func (o *ObjectInformationA_Info) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ObjectInformationA_Drive:
		if value != nil {
			return value.Drive
		}
	case *ObjectInformationA_DriveType:
		if value != nil {
			return value.DriveType
		}
	case *ObjectInformationA_Library:
		if value != nil {
			return value.Library
		}
	case *ObjectInformationA_Changer:
		if value != nil {
			return value.Changer
		}
	case *ObjectInformationA_ChangerType:
		if value != nil {
			return value.ChangerType
		}
	case *ObjectInformationA_StorageSlot:
		if value != nil {
			return value.StorageSlot
		}
	case *ObjectInformationA_IEDoor:
		if value != nil {
			return value.IEDoor
		}
	case *ObjectInformationA_IEPort:
		if value != nil {
			return value.IEPort
		}
	case *ObjectInformationA_PhysicalMedia:
		if value != nil {
			return value.PhysicalMedia
		}
	case *ObjectInformationA_LogicalMedia:
		if value != nil {
			return value.LogicalMedia
		}
	case *ObjectInformationA_Partition:
		if value != nil {
			return value.Partition
		}
	case *ObjectInformationA_MediaPool:
		if value != nil {
			return value.MediaPool
		}
	case *ObjectInformationA_MediaType:
		if value != nil {
			return value.MediaType
		}
	case *ObjectInformationA_LibRequest:
		if value != nil {
			return value.LibRequest
		}
	case *ObjectInformationA_OperationRequest:
		if value != nil {
			return value.OperationRequest
		}
	case *ObjectInformationA_Computer:
		if value != nil {
			return value.Computer
		}
	}
	return nil
}

type is_ObjectInformationA_Info interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ObjectInformationA_Info()
}

func (o *ObjectInformationA_Info) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ObjectInformationA_Drive:
		return uint32(5)
	case *ObjectInformationA_DriveType:
		return uint32(6)
	case *ObjectInformationA_Library:
		return uint32(9)
	case *ObjectInformationA_Changer:
		return uint32(2)
	case *ObjectInformationA_ChangerType:
		return uint32(3)
	case *ObjectInformationA_StorageSlot:
		return uint32(16)
	case *ObjectInformationA_IEDoor:
		return uint32(7)
	case *ObjectInformationA_IEPort:
		return uint32(8)
	case *ObjectInformationA_PhysicalMedia:
		return uint32(15)
	case *ObjectInformationA_LogicalMedia:
		return uint32(11)
	case *ObjectInformationA_Partition:
		return uint32(14)
	case *ObjectInformationA_MediaPool:
		return uint32(12)
	case *ObjectInformationA_MediaType:
		return uint32(13)
	case *ObjectInformationA_LibRequest:
		return uint32(10)
	case *ObjectInformationA_OperationRequest:
		return uint32(17)
	case *ObjectInformationA_Computer:
		return uint32(4)
	}
	return uint32(0)
}

func (o *ObjectInformationA_Info) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(5):
		_o, _ := o.Value.(*ObjectInformationA_Drive)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_Drive{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*ObjectInformationA_DriveType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_DriveType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*ObjectInformationA_Library)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_Library{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*ObjectInformationA_Changer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_Changer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*ObjectInformationA_ChangerType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_ChangerType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*ObjectInformationA_StorageSlot)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_StorageSlot{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*ObjectInformationA_IEDoor)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_IEDoor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*ObjectInformationA_IEPort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_IEPort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(15):
		_o, _ := o.Value.(*ObjectInformationA_PhysicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_PhysicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*ObjectInformationA_LogicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_LogicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*ObjectInformationA_Partition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_Partition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*ObjectInformationA_MediaPool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_MediaPool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*ObjectInformationA_MediaType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_MediaType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*ObjectInformationA_LibRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_LibRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17):
		_o, _ := o.Value.(*ObjectInformationA_OperationRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_OperationRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*ObjectInformationA_Computer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationA_Computer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ObjectInformationA_Info) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(5):
		o.Value = &ObjectInformationA_Drive{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &ObjectInformationA_DriveType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &ObjectInformationA_Library{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &ObjectInformationA_Changer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &ObjectInformationA_ChangerType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &ObjectInformationA_StorageSlot{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &ObjectInformationA_IEDoor{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &ObjectInformationA_IEPort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(15):
		o.Value = &ObjectInformationA_PhysicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &ObjectInformationA_LogicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &ObjectInformationA_Partition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &ObjectInformationA_MediaPool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &ObjectInformationA_MediaType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &ObjectInformationA_LibRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17):
		o.Value = &ObjectInformationA_OperationRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &ObjectInformationA_Computer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ObjectInformationA_Drive structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 5
type ObjectInformationA_Drive struct {
	Drive *DriveInformationA `idl:"name:Drive" json:"drive"`
}

func (*ObjectInformationA_Drive) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_Drive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Drive != nil {
		if err := o.Drive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DriveInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_Drive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Drive == nil {
		o.Drive = &DriveInformationA{}
	}
	if err := o.Drive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_DriveType structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 6
type ObjectInformationA_DriveType struct {
	DriveType *DriveTypeInformationA `idl:"name:DriveType" json:"drive_type"`
}

func (*ObjectInformationA_DriveType) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_DriveType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DriveTypeInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_DriveType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DriveType == nil {
		o.DriveType = &DriveTypeInformationA{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_Library structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 9
type ObjectInformationA_Library struct {
	Library *LibraryInformation `idl:"name:Library" json:"library"`
}

func (*ObjectInformationA_Library) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_Library) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LibraryInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_Library) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Library == nil {
		o.Library = &LibraryInformation{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_Changer structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 2
type ObjectInformationA_Changer struct {
	Changer *ChangerInformationA `idl:"name:Changer" json:"changer"`
}

func (*ObjectInformationA_Changer) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_Changer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Changer != nil {
		if err := o.Changer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangerInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_Changer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Changer == nil {
		o.Changer = &ChangerInformationA{}
	}
	if err := o.Changer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_ChangerType structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 3
type ObjectInformationA_ChangerType struct {
	ChangerType *ChangerTypeInformationA `idl:"name:ChangerType" json:"changer_type"`
}

func (*ObjectInformationA_ChangerType) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_ChangerType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangerTypeInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_ChangerType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ChangerType == nil {
		o.ChangerType = &ChangerTypeInformationA{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_StorageSlot structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 16
type ObjectInformationA_StorageSlot struct {
	StorageSlot *StorageSlotInformation `idl:"name:StorageSlot" json:"storage_slot"`
}

func (*ObjectInformationA_StorageSlot) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_StorageSlot) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StorageSlot != nil {
		if err := o.StorageSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StorageSlotInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_StorageSlot) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StorageSlot == nil {
		o.StorageSlot = &StorageSlotInformation{}
	}
	if err := o.StorageSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_IEDoor structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 7
type ObjectInformationA_IEDoor struct {
	IEDoor *IEDoorInformation `idl:"name:IEDoor" json:"ie_door"`
}

func (*ObjectInformationA_IEDoor) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_IEDoor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IEDoor != nil {
		if err := o.IEDoor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IEDoorInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_IEDoor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IEDoor == nil {
		o.IEDoor = &IEDoorInformation{}
	}
	if err := o.IEDoor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_IEPort structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 8
type ObjectInformationA_IEPort struct {
	IEPort *IEPortInformation `idl:"name:IEPort" json:"ie_port"`
}

func (*ObjectInformationA_IEPort) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_IEPort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IEPort != nil {
		if err := o.IEPort.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IEPortInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_IEPort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IEPort == nil {
		o.IEPort = &IEPortInformation{}
	}
	if err := o.IEPort.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_PhysicalMedia structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 15
type ObjectInformationA_PhysicalMedia struct {
	PhysicalMedia *PMIDInformationA `idl:"name:PhysicalMedia" json:"physical_media"`
}

func (*ObjectInformationA_PhysicalMedia) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_PhysicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PMIDInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_PhysicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &PMIDInformationA{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_LogicalMedia structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 11
type ObjectInformationA_LogicalMedia struct {
	LogicalMedia *LMIDInformation `idl:"name:LogicalMedia" json:"logical_media"`
}

func (*ObjectInformationA_LogicalMedia) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_LogicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LMIDInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_LogicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LogicalMedia == nil {
		o.LogicalMedia = &LMIDInformation{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_Partition structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 14
type ObjectInformationA_Partition struct {
	Partition *PartitionInformationA `idl:"name:Partition" json:"partition"`
}

func (*ObjectInformationA_Partition) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_Partition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PartitionInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_Partition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Partition == nil {
		o.Partition = &PartitionInformationA{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_MediaPool structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 12
type ObjectInformationA_MediaPool struct {
	MediaPool *MediaPoolInformation `idl:"name:MediaPool" json:"media_pool"`
}

func (*ObjectInformationA_MediaPool) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_MediaPool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MediaPoolInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_MediaPool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaPool == nil {
		o.MediaPool = &MediaPoolInformation{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_MediaType structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 13
type ObjectInformationA_MediaType struct {
	MediaType *MediaTypeInformation `idl:"name:MediaType" json:"media_type"`
}

func (*ObjectInformationA_MediaType) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_MediaType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MediaTypeInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_MediaType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaType == nil {
		o.MediaType = &MediaTypeInformation{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_LibRequest structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 10
type ObjectInformationA_LibRequest struct {
	LibRequest *LibraryRequestInformationA `idl:"name:LibRequest" json:"lib_request"`
}

func (*ObjectInformationA_LibRequest) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_LibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LibRequest != nil {
		if err := o.LibRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LibraryRequestInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_LibRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LibRequest == nil {
		o.LibRequest = &LibraryRequestInformationA{}
	}
	if err := o.LibRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_OperationRequest structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 17
type ObjectInformationA_OperationRequest struct {
	OperationRequest *OperationRequestInformationA `idl:"name:OpRequest" json:"operation_request"`
}

func (*ObjectInformationA_OperationRequest) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_OperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OperationRequest != nil {
		if err := o.OperationRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OperationRequestInformationA{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_OperationRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OperationRequest == nil {
		o.OperationRequest = &OperationRequestInformationA{}
	}
	if err := o.OperationRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationA_Computer structure represents ObjectInformationA_Info RPC union arm.
//
// It has following labels: 4
type ObjectInformationA_Computer struct {
	Computer *ComputerInformation `idl:"name:Computer" json:"computer"`
}

func (*ObjectInformationA_Computer) is_ObjectInformationA_Info() {}

func (o *ObjectInformationA_Computer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Computer != nil {
		if err := o.Computer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ComputerInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationA_Computer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Computer == nil {
		o.Computer = &ComputerInformation{}
	}
	if err := o.Computer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW structure represents NTMS_OBJECTINFORMATIONW RPC structure.
type ObjectInformationW struct {
	Size             uint32                   `idl:"name:dwSize" json:"size"`
	Type             uint32                   `idl:"name:dwType" json:"type"`
	Created          *dtyp.SystemTime         `idl:"name:Created" json:"created"`
	Modified         *dtyp.SystemTime         `idl:"name:Modified" json:"modified"`
	ObjectGUID       *GUID                    `idl:"name:ObjectGuid" json:"object_guid"`
	Enabled          bool                     `idl:"name:Enabled" json:"enabled"`
	OperationalState uint32                   `idl:"name:dwOperationalState" json:"operational_state"`
	Name             string                   `idl:"name:szName;string" json:"name"`
	Description      string                   `idl:"name:szDescription;string" json:"description"`
	Info             *ObjectInformationW_Info `idl:"name:Info;switch_is:dwType" json:"info"`
}

func (o *ObjectInformationW) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectInformationW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if o.Created != nil {
		if err := o.Created.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Modified != nil {
		if err := o.Modified.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.SystemTime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ObjectGUID != nil {
		if err := o.ObjectGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if !o.Enabled {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.OperationalState); err != nil {
		return err
	}
	_Name_buf := utf16.Encode([]rune(o.Name))
	if uint64(len(_Name_buf)) > 64-1 {
		_Name_buf = _Name_buf[:64-1]
	}
	if o.Name != ndr.ZeroString {
		_Name_buf = append(_Name_buf, uint16(0))
	}
	for i1 := range _Name_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_Name_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Name_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_Description_buf := utf16.Encode([]rune(o.Description))
	if uint64(len(_Description_buf)) > 127-1 {
		_Description_buf = _Description_buf[:127-1]
	}
	if o.Description != ndr.ZeroString {
		_Description_buf = append(_Description_buf, uint16(0))
	}
	for i1 := range _Description_buf {
		i1 := i1
		if uint64(i1) >= 127 {
			break
		}
		if err := w.WriteData(_Description_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Description_buf); uint64(i1) < 127; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_swInfo := uint32(o.Type)
	if o.Info != nil {
		if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	} else {
		if err := (&ObjectInformationW_Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if o.Created == nil {
		o.Created = &dtyp.SystemTime{}
	}
	if err := o.Created.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Modified == nil {
		o.Modified = &dtyp.SystemTime{}
	}
	if err := o.Modified.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ObjectGUID == nil {
		o.ObjectGUID = &GUID{}
	}
	if err := o.ObjectGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	var _bEnabled int32
	if err := w.ReadData(&_bEnabled); err != nil {
		return err
	}
	o.Enabled = _bEnabled != 0
	if err := w.ReadData(&o.OperationalState); err != nil {
		return err
	}
	var _Name_buf []uint16
	_Name_buf = make([]uint16, 64)
	for i1 := range _Name_buf {
		i1 := i1
		if err := w.ReadData(&_Name_buf[i1]); err != nil {
			return err
		}
	}
	o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
	var _Description_buf []uint16
	_Description_buf = make([]uint16, 127)
	for i1 := range _Description_buf {
		i1 := i1
		if err := w.ReadData(&_Description_buf[i1]); err != nil {
			return err
		}
	}
	o.Description = strings.TrimRight(string(utf16.Decode(_Description_buf)), ndr.ZeroString)
	if o.Info == nil {
		o.Info = &ObjectInformationW_Info{}
	}
	_swInfo := uint32(o.Type)
	if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_Info structure represents NTMS_OBJECTINFORMATIONW union anonymous member.
type ObjectInformationW_Info struct {
	// Types that are assignable to Value
	//
	// *ObjectInformationW_Drive
	// *ObjectInformationW_DriveType
	// *ObjectInformationW_Library
	// *ObjectInformationW_Changer
	// *ObjectInformationW_ChangerType
	// *ObjectInformationW_StorageSlot
	// *ObjectInformationW_IEDoor
	// *ObjectInformationW_IEPort
	// *ObjectInformationW_PhysicalMedia
	// *ObjectInformationW_LogicalMedia
	// *ObjectInformationW_Partition
	// *ObjectInformationW_MediaPool
	// *ObjectInformationW_MediaType
	// *ObjectInformationW_LibRequest
	// *ObjectInformationW_OperationRequest
	// *ObjectInformationW_Computer
	Value is_ObjectInformationW_Info `json:"value"`
}

func (o *ObjectInformationW_Info) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ObjectInformationW_Drive:
		if value != nil {
			return value.Drive
		}
	case *ObjectInformationW_DriveType:
		if value != nil {
			return value.DriveType
		}
	case *ObjectInformationW_Library:
		if value != nil {
			return value.Library
		}
	case *ObjectInformationW_Changer:
		if value != nil {
			return value.Changer
		}
	case *ObjectInformationW_ChangerType:
		if value != nil {
			return value.ChangerType
		}
	case *ObjectInformationW_StorageSlot:
		if value != nil {
			return value.StorageSlot
		}
	case *ObjectInformationW_IEDoor:
		if value != nil {
			return value.IEDoor
		}
	case *ObjectInformationW_IEPort:
		if value != nil {
			return value.IEPort
		}
	case *ObjectInformationW_PhysicalMedia:
		if value != nil {
			return value.PhysicalMedia
		}
	case *ObjectInformationW_LogicalMedia:
		if value != nil {
			return value.LogicalMedia
		}
	case *ObjectInformationW_Partition:
		if value != nil {
			return value.Partition
		}
	case *ObjectInformationW_MediaPool:
		if value != nil {
			return value.MediaPool
		}
	case *ObjectInformationW_MediaType:
		if value != nil {
			return value.MediaType
		}
	case *ObjectInformationW_LibRequest:
		if value != nil {
			return value.LibRequest
		}
	case *ObjectInformationW_OperationRequest:
		if value != nil {
			return value.OperationRequest
		}
	case *ObjectInformationW_Computer:
		if value != nil {
			return value.Computer
		}
	}
	return nil
}

type is_ObjectInformationW_Info interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ObjectInformationW_Info()
}

func (o *ObjectInformationW_Info) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ObjectInformationW_Drive:
		return uint32(5)
	case *ObjectInformationW_DriveType:
		return uint32(6)
	case *ObjectInformationW_Library:
		return uint32(9)
	case *ObjectInformationW_Changer:
		return uint32(2)
	case *ObjectInformationW_ChangerType:
		return uint32(3)
	case *ObjectInformationW_StorageSlot:
		return uint32(16)
	case *ObjectInformationW_IEDoor:
		return uint32(7)
	case *ObjectInformationW_IEPort:
		return uint32(8)
	case *ObjectInformationW_PhysicalMedia:
		return uint32(15)
	case *ObjectInformationW_LogicalMedia:
		return uint32(11)
	case *ObjectInformationW_Partition:
		return uint32(14)
	case *ObjectInformationW_MediaPool:
		return uint32(12)
	case *ObjectInformationW_MediaType:
		return uint32(13)
	case *ObjectInformationW_LibRequest:
		return uint32(10)
	case *ObjectInformationW_OperationRequest:
		return uint32(17)
	case *ObjectInformationW_Computer:
		return uint32(4)
	}
	return uint32(0)
}

func (o *ObjectInformationW_Info) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(5):
		_o, _ := o.Value.(*ObjectInformationW_Drive)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_Drive{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*ObjectInformationW_DriveType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_DriveType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*ObjectInformationW_Library)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_Library{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*ObjectInformationW_Changer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_Changer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*ObjectInformationW_ChangerType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_ChangerType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*ObjectInformationW_StorageSlot)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_StorageSlot{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*ObjectInformationW_IEDoor)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_IEDoor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*ObjectInformationW_IEPort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_IEPort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(15):
		_o, _ := o.Value.(*ObjectInformationW_PhysicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_PhysicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*ObjectInformationW_LogicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_LogicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*ObjectInformationW_Partition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_Partition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*ObjectInformationW_MediaPool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_MediaPool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*ObjectInformationW_MediaType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_MediaType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*ObjectInformationW_LibRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_LibRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17):
		_o, _ := o.Value.(*ObjectInformationW_OperationRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_OperationRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*ObjectInformationW_Computer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectInformationW_Computer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ObjectInformationW_Info) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(5):
		o.Value = &ObjectInformationW_Drive{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &ObjectInformationW_DriveType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &ObjectInformationW_Library{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &ObjectInformationW_Changer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &ObjectInformationW_ChangerType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &ObjectInformationW_StorageSlot{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &ObjectInformationW_IEDoor{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &ObjectInformationW_IEPort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(15):
		o.Value = &ObjectInformationW_PhysicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &ObjectInformationW_LogicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &ObjectInformationW_Partition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &ObjectInformationW_MediaPool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &ObjectInformationW_MediaType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &ObjectInformationW_LibRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17):
		o.Value = &ObjectInformationW_OperationRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &ObjectInformationW_Computer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ObjectInformationW_Drive structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 5
type ObjectInformationW_Drive struct {
	Drive *DriveInformationW `idl:"name:Drive" json:"drive"`
}

func (*ObjectInformationW_Drive) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_Drive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Drive != nil {
		if err := o.Drive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DriveInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_Drive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Drive == nil {
		o.Drive = &DriveInformationW{}
	}
	if err := o.Drive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_DriveType structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 6
type ObjectInformationW_DriveType struct {
	DriveType *DriveTypeInformationW `idl:"name:DriveType" json:"drive_type"`
}

func (*ObjectInformationW_DriveType) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_DriveType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DriveTypeInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_DriveType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DriveType == nil {
		o.DriveType = &DriveTypeInformationW{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_Library structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 9
type ObjectInformationW_Library struct {
	Library *LibraryInformation `idl:"name:Library" json:"library"`
}

func (*ObjectInformationW_Library) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_Library) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LibraryInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_Library) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Library == nil {
		o.Library = &LibraryInformation{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_Changer structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 2
type ObjectInformationW_Changer struct {
	Changer *ChangerInformationW `idl:"name:Changer" json:"changer"`
}

func (*ObjectInformationW_Changer) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_Changer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Changer != nil {
		if err := o.Changer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangerInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_Changer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Changer == nil {
		o.Changer = &ChangerInformationW{}
	}
	if err := o.Changer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_ChangerType structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 3
type ObjectInformationW_ChangerType struct {
	ChangerType *ChangerTypeInformationW `idl:"name:ChangerType" json:"changer_type"`
}

func (*ObjectInformationW_ChangerType) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_ChangerType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ChangerTypeInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_ChangerType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ChangerType == nil {
		o.ChangerType = &ChangerTypeInformationW{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_StorageSlot structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 16
type ObjectInformationW_StorageSlot struct {
	StorageSlot *StorageSlotInformation `idl:"name:StorageSlot" json:"storage_slot"`
}

func (*ObjectInformationW_StorageSlot) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_StorageSlot) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StorageSlot != nil {
		if err := o.StorageSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StorageSlotInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_StorageSlot) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StorageSlot == nil {
		o.StorageSlot = &StorageSlotInformation{}
	}
	if err := o.StorageSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_IEDoor structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 7
type ObjectInformationW_IEDoor struct {
	IEDoor *IEDoorInformation `idl:"name:IEDoor" json:"ie_door"`
}

func (*ObjectInformationW_IEDoor) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_IEDoor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IEDoor != nil {
		if err := o.IEDoor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IEDoorInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_IEDoor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IEDoor == nil {
		o.IEDoor = &IEDoorInformation{}
	}
	if err := o.IEDoor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_IEPort structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 8
type ObjectInformationW_IEPort struct {
	IEPort *IEPortInformation `idl:"name:IEPort" json:"ie_port"`
}

func (*ObjectInformationW_IEPort) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_IEPort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IEPort != nil {
		if err := o.IEPort.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IEPortInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_IEPort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IEPort == nil {
		o.IEPort = &IEPortInformation{}
	}
	if err := o.IEPort.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_PhysicalMedia structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 15
type ObjectInformationW_PhysicalMedia struct {
	PhysicalMedia *PMIDInformationW `idl:"name:PhysicalMedia" json:"physical_media"`
}

func (*ObjectInformationW_PhysicalMedia) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_PhysicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PMIDInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_PhysicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &PMIDInformationW{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_LogicalMedia structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 11
type ObjectInformationW_LogicalMedia struct {
	LogicalMedia *LMIDInformation `idl:"name:LogicalMedia" json:"logical_media"`
}

func (*ObjectInformationW_LogicalMedia) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_LogicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LMIDInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_LogicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LogicalMedia == nil {
		o.LogicalMedia = &LMIDInformation{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_Partition structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 14
type ObjectInformationW_Partition struct {
	Partition *PartitionInformationW `idl:"name:Partition" json:"partition"`
}

func (*ObjectInformationW_Partition) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_Partition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PartitionInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_Partition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Partition == nil {
		o.Partition = &PartitionInformationW{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_MediaPool structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 12
type ObjectInformationW_MediaPool struct {
	MediaPool *MediaPoolInformation `idl:"name:MediaPool" json:"media_pool"`
}

func (*ObjectInformationW_MediaPool) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_MediaPool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MediaPoolInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_MediaPool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaPool == nil {
		o.MediaPool = &MediaPoolInformation{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_MediaType structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 13
type ObjectInformationW_MediaType struct {
	MediaType *MediaTypeInformation `idl:"name:MediaType" json:"media_type"`
}

func (*ObjectInformationW_MediaType) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_MediaType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MediaTypeInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_MediaType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaType == nil {
		o.MediaType = &MediaTypeInformation{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_LibRequest structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 10
type ObjectInformationW_LibRequest struct {
	LibRequest *LibraryRequestInformationW `idl:"name:LibRequest" json:"lib_request"`
}

func (*ObjectInformationW_LibRequest) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_LibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LibRequest != nil {
		if err := o.LibRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LibraryRequestInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_LibRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LibRequest == nil {
		o.LibRequest = &LibraryRequestInformationW{}
	}
	if err := o.LibRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_OperationRequest structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 17
type ObjectInformationW_OperationRequest struct {
	OperationRequest *OperationRequestInformationW `idl:"name:OpRequest" json:"operation_request"`
}

func (*ObjectInformationW_OperationRequest) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_OperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OperationRequest != nil {
		if err := o.OperationRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&OperationRequestInformationW{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_OperationRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OperationRequest == nil {
		o.OperationRequest = &OperationRequestInformationW{}
	}
	if err := o.OperationRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectInformationW_Computer structure represents ObjectInformationW_Info RPC union arm.
//
// It has following labels: 4
type ObjectInformationW_Computer struct {
	Computer *ComputerInformation `idl:"name:Computer" json:"computer"`
}

func (*ObjectInformationW_Computer) is_ObjectInformationW_Info() {}

func (o *ObjectInformationW_Computer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Computer != nil {
		if err := o.Computer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ComputerInformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectInformationW_Computer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Computer == nil {
		o.Computer = &ComputerInformation{}
	}
	if err := o.Computer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectManagement2 structure represents INtmsObjectManagement2 RPC structure.
type ObjectManagement2 dcom.InterfacePointer

func (o *ObjectManagement2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ObjectManagement2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ObjectManagement2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ObjectManagement2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ObjectManagement2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ObjectManagement3 structure represents INtmsObjectManagement3 RPC structure.
type ObjectManagement3 dcom.InterfacePointer

func (o *ObjectManagement3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ObjectManagement3) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ObjectManagement3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ObjectManagement3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ObjectManagement3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// LibraryControl1 structure represents INtmsLibraryControl1 RPC structure.
type LibraryControl1 dcom.InterfacePointer

func (o *LibraryControl1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *LibraryControl1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *LibraryControl1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *LibraryControl1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *LibraryControl1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ObjectManagement1 structure represents INtmsObjectManagement1 RPC structure.
type ObjectManagement1 dcom.InterfacePointer

func (o *ObjectManagement1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ObjectManagement1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ObjectManagement1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ObjectManagement1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ObjectManagement1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MediaServices1 structure represents INtmsMediaServices1 RPC structure.
type MediaServices1 dcom.InterfacePointer

func (o *MediaServices1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *MediaServices1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *MediaServices1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *MediaServices1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MediaServices1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClientSink structure represents IClientSink RPC structure.
type ClientSink dcom.InterfacePointer

func (o *ClientSink) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClientSink) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ClientSink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClientSink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClientSink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RobustNTMSMediaServices1 structure represents IRobustNtmsMediaServices1 RPC structure.
type RobustNTMSMediaServices1 dcom.InterfacePointer

func (o *RobustNTMSMediaServices1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RobustNTMSMediaServices1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RobustNTMSMediaServices1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RobustNTMSMediaServices1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RobustNTMSMediaServices1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Session1 structure represents INtmsSession1 RPC structure.
type Session1 dcom.InterfacePointer

func (o *Session1) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Session1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Session1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Session1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Session1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Unknown structure represents IUnknown RPC structure.
type Unknown dcom.InterfacePointer

func (o *Unknown) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Unknown) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Unknown) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Unknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Unknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NotifySink structure represents INtmsNotifySink RPC structure.
type NotifySink dcom.InterfacePointer

func (o *NotifySink) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *NotifySink) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NotifySink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NotifySink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NotifySink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Messenger structure represents IMessenger RPC structure.
type Messenger dcom.InterfacePointer

func (o *Messenger) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Messenger) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Messenger) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Messenger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Messenger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ObjectInfo1 structure represents INtmsObjectInfo1 RPC structure.
type ObjectInfo1 dcom.InterfacePointer

func (o *ObjectInfo1) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ObjectInfo1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ObjectInfo1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ObjectInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ObjectInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// LibraryControl2 structure represents INtmsLibraryControl2 RPC structure.
type LibraryControl2 dcom.InterfacePointer

func (o *LibraryControl2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *LibraryControl2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *LibraryControl2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *LibraryControl2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *LibraryControl2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CNtmsSvr class identifier d61a27c6-8f53-11d0-bfa0-00a024151983
var NTMSServerClassID = &dcom.ClassID{Data1: 0xd61a27c6, Data2: 0x8f53, Data3: 0x11d0, Data4: []byte{0xbf, 0xa0, 0x00, 0xa0, 0x24, 0x15, 0x19, 0x83}}
