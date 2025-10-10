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

// NTMSGenericRead represents the NTMS_GENERIC_READ RPC constant
var NTMSGenericRead = 1

// NTMSGenericWrite represents the NTMS_GENERIC_WRITE RPC constant
var NTMSGenericWrite = 3

// NTMSGenericExecute represents the NTMS_GENERIC_EXECUTE RPC constant
var NTMSGenericExecute = 3

// NTMSGenericAll represents the NTMS_GENERIC_ALL RPC constant
var NTMSGenericAll = 3

// NTMSGUID structure represents NTMS_GUID RPC structure.
type NTMSGUID dtyp.GUID

func (o *NTMSGUID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *NTMSGUID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSGUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSGUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSLibraryinformation structure represents NTMS_LIBRARYINFORMATION RPC structure.
type NTMSLibraryinformation struct {
	LibraryType                  uint32     `idl:"name:LibraryType" json:"library_type"`
	CleanerSlot                  *NTMSGUID  `idl:"name:CleanerSlot" json:"cleaner_slot"`
	CleanerSlotDefault           *NTMSGUID  `idl:"name:CleanerSlotDefault" json:"cleaner_slot_default"`
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

func (o *NTMSLibraryinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLibraryinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CleanerSlotDefault != nil {
		if err := o.CleanerSlotDefault.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSLibraryinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LibraryType); err != nil {
		return err
	}
	if o.CleanerSlot == nil {
		o.CleanerSlot = &NTMSGUID{}
	}
	if err := o.CleanerSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CleanerSlotDefault == nil {
		o.CleanerSlotDefault = &NTMSGUID{}
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

// NTMSAllocationInformation structure represents NTMS_ALLOCATION_INFORMATION RPC structure.
type NTMSAllocationInformation struct {
	Size          uint32    `idl:"name:dwSize" json:"size"`
	_             uint8     `idl:"name:lpReserved"`
	AllocatedFrom *NTMSGUID `idl:"name:AllocatedFrom" json:"allocated_from"`
}

func (o *NTMSAllocationInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSAllocationInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *NTMSAllocationInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.AllocatedFrom = &NTMSGUID{}
	}
	if err := o.AllocatedFrom.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// NTMSAsyncIO structure represents NTMS_ASYNC_IO RPC structure.
type NTMSAsyncIO struct {
	OperationID   *NTMSGUID `idl:"name:OperationId" json:"operation_id"`
	EventID       *NTMSGUID `idl:"name:EventId" json:"event_id"`
	OperationType uint32    `idl:"name:dwOperationType" json:"operation_type"`
	Result        uint32    `idl:"name:dwResult" json:"result"`
	AsyncState    uint32    `idl:"name:dwAsyncState" json:"async_state"`
	HEvent        uint64    `idl:"name:hEvent" json:"h_event"`
	OnStateChange bool      `idl:"name:bOnStateChange" json:"on_state_change"`
}

func (o *NTMSAsyncIO) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSAsyncIO) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EventID != nil {
		if err := o.EventID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
	if err := w.WriteData(ndr.Uint3264(o.HEvent)); err != nil {
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
func (o *NTMSAsyncIO) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.OperationID == nil {
		o.OperationID = &NTMSGUID{}
	}
	if err := o.OperationID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EventID == nil {
		o.EventID = &NTMSGUID{}
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
	if err := w.ReadData((*ndr.Uint3264)(&o.HEvent)); err != nil {
		return err
	}
	var _bOnStateChange int32
	if err := w.ReadData(&_bOnStateChange); err != nil {
		return err
	}
	o.OnStateChange = _bOnStateChange != 0
	return nil
}

// NTMSMountInformation structure represents NTMS_MOUNT_INFORMATION RPC structure.
type NTMSMountInformation struct {
	Size uint32       `idl:"name:dwSize" json:"size"`
	_    *NTMSAsyncIO `idl:"name:lpReserved;pointer:ptr"`
}

func (o *NTMSMountInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSMountInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSMountInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved lpReserved
	var _lpReserved *NTMSAsyncIO
	_ptr_lpReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if _lpReserved == nil {
			_lpReserved = &NTMSAsyncIO{}
		}
		if err := _lpReserved.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpReserved := func(ptr interface{}) { _lpReserved = *ptr.(**NTMSAsyncIO) }
	if err := w.ReadPointer(&_lpReserved, _s_lpReserved, _ptr_lpReserved); err != nil {
		return err
	}
	return nil
}

// NTMSChangerinformationa structure represents NTMS_CHANGERINFORMATIONA RPC structure.
type NTMSChangerinformationa struct {
	Number       uint32    `idl:"name:Number" json:"number"`
	ChangerType  *NTMSGUID `idl:"name:ChangerType" json:"changer_type"`
	SerialNumber []byte    `idl:"name:szSerialNumber" json:"serial_number"`
	Revision     []byte    `idl:"name:szRevision" json:"revision"`
	DeviceName   []byte    `idl:"name:szDeviceName" json:"device_name"`
	SCSIPort     uint16    `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus      uint16    `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget   uint16    `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN      uint16    `idl:"name:ScsiLun" json:"scsi_lun"`
	Library      *NTMSGUID `idl:"name:Library" json:"library"`
}

func (o *NTMSChangerinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSChangerinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSChangerinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if o.ChangerType == nil {
		o.ChangerType = &NTMSGUID{}
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
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSChangerinformationw structure represents NTMS_CHANGERINFORMATIONW RPC structure.
type NTMSChangerinformationw struct {
	Number       uint32    `idl:"name:Number" json:"number"`
	ChangerType  *NTMSGUID `idl:"name:ChangerType" json:"changer_type"`
	SerialNumber string    `idl:"name:szSerialNumber;string" json:"serial_number"`
	Revision     string    `idl:"name:szRevision;string" json:"revision"`
	DeviceName   string    `idl:"name:szDeviceName;string" json:"device_name"`
	SCSIPort     uint16    `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus      uint16    `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget   uint16    `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN      uint16    `idl:"name:ScsiLun" json:"scsi_lun"`
	Library      *NTMSGUID `idl:"name:Library" json:"library"`
}

func (o *NTMSChangerinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSChangerinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *NTMSChangerinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Number); err != nil {
		return err
	}
	if o.ChangerType == nil {
		o.ChangerType = &NTMSGUID{}
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
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// NTMSChangertypeinformationa structure represents NTMS_CHANGERTYPEINFORMATIONA RPC structure.
type NTMSChangertypeinformationa struct {
	Vendor     []byte `idl:"name:szVendor" json:"vendor"`
	Product    []byte `idl:"name:szProduct" json:"product"`
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *NTMSChangertypeinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSChangertypeinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSChangertypeinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSChangertypeinformationw structure represents NTMS_CHANGERTYPEINFORMATIONW RPC structure.
type NTMSChangertypeinformationw struct {
	Vendor     string `idl:"name:szVendor;string" json:"vendor"`
	Product    string `idl:"name:szProduct;string" json:"product"`
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *NTMSChangertypeinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSChangertypeinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSChangertypeinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSDriveinformationa structure represents NTMS_DRIVEINFORMATIONA RPC structure.
type NTMSDriveinformationa struct {
	Number             uint32           `idl:"name:Number" json:"number"`
	State              uint32           `idl:"name:State" json:"state"`
	DriveType          *NTMSGUID        `idl:"name:DriveType" json:"drive_type"`
	DeviceName         []byte           `idl:"name:szDeviceName" json:"device_name"`
	SerialNumber       []byte           `idl:"name:szSerialNumber" json:"serial_number"`
	Revision           []byte           `idl:"name:szRevision" json:"revision"`
	SCSIPort           uint16           `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus            uint16           `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget         uint16           `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN            uint16           `idl:"name:ScsiLun" json:"scsi_lun"`
	MountCount         uint32           `idl:"name:dwMountCount" json:"mount_count"`
	LastCleanedTS      *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	SavedPartitionID   *NTMSGUID        `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	Library            *NTMSGUID        `idl:"name:Library" json:"library"`
	_                  *dtyp.GUID       `idl:"name:Reserved"`
	DeferDismountDelay uint32           `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
}

func (o *NTMSDriveinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSDriveinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSDriveinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.DriveType = &NTMSGUID{}
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
		o.SavedPartitionID = &NTMSGUID{}
	}
	if err := o.SavedPartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &NTMSGUID{}
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

// NTMSDriveinformationw structure represents NTMS_DRIVEINFORMATIONW RPC structure.
type NTMSDriveinformationw struct {
	Number             uint32           `idl:"name:Number" json:"number"`
	State              uint32           `idl:"name:State" json:"state"`
	DriveType          *NTMSGUID        `idl:"name:DriveType" json:"drive_type"`
	DeviceName         string           `idl:"name:szDeviceName;string" json:"device_name"`
	SerialNumber       string           `idl:"name:szSerialNumber;string" json:"serial_number"`
	Revision           string           `idl:"name:szRevision;string" json:"revision"`
	SCSIPort           uint16           `idl:"name:ScsiPort" json:"scsi_port"`
	SCSIBus            uint16           `idl:"name:ScsiBus" json:"scsi_bus"`
	SCSITarget         uint16           `idl:"name:ScsiTarget" json:"scsi_target"`
	SCSILUN            uint16           `idl:"name:ScsiLun" json:"scsi_lun"`
	MountCount         uint32           `idl:"name:dwMountCount" json:"mount_count"`
	LastCleanedTS      *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	SavedPartitionID   *NTMSGUID        `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	Library            *NTMSGUID        `idl:"name:Library" json:"library"`
	_                  *dtyp.GUID       `idl:"name:Reserved"`
	DeferDismountDelay uint32           `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
}

func (o *NTMSDriveinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSDriveinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSDriveinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.DriveType = &NTMSGUID{}
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
		o.SavedPartitionID = &NTMSGUID{}
	}
	if err := o.SavedPartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &NTMSGUID{}
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

// NTMSDrivetypeinformationa structure represents NTMS_DRIVETYPEINFORMATIONA RPC structure.
type NTMSDrivetypeinformationa struct {
	Vendor        []byte `idl:"name:szVendor" json:"vendor"`
	Product       []byte `idl:"name:szProduct" json:"product"`
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	DeviceType    uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *NTMSDrivetypeinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSDrivetypeinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSDrivetypeinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSDrivetypeinformationw structure represents NTMS_DRIVETYPEINFORMATIONW RPC structure.
type NTMSDrivetypeinformationw struct {
	Vendor        string `idl:"name:szVendor;string" json:"vendor"`
	Product       string `idl:"name:szProduct;string" json:"product"`
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	DeviceType    uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *NTMSDrivetypeinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSDrivetypeinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSDrivetypeinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSLibrequestinformationa structure represents NTMS_LIBREQUESTINFORMATIONA RPC structure.
type NTMSLibrequestinformationa struct {
	OperationCode   uint32           `idl:"name:OperationCode" json:"operation_code"`
	OperationOption uint32           `idl:"name:OperationOption" json:"operation_option"`
	State           uint32           `idl:"name:State" json:"state"`
	PartitionID     *NTMSGUID        `idl:"name:PartitionId" json:"partition_id"`
	DriveID         *NTMSGUID        `idl:"name:DriveId" json:"drive_id"`
	PhysMediaID     *NTMSGUID        `idl:"name:PhysMediaId" json:"phys_media_id"`
	Library         *NTMSGUID        `idl:"name:Library" json:"library"`
	SlotID          *NTMSGUID        `idl:"name:SlotId" json:"slot_id"`
	TimeQueued      *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	TimeCompleted   *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	Application     []byte           `idl:"name:szApplication" json:"application"`
	User            []byte           `idl:"name:szUser" json:"user"`
	Computer        []byte           `idl:"name:szComputer" json:"computer"`
	ErrorCode       uint32           `idl:"name:dwErrorCode" json:"error_code"`
	WorkItemID      *NTMSGUID        `idl:"name:WorkItemId" json:"work_item_id"`
	Priority        uint32           `idl:"name:dwPriority" json:"priority"`
}

func (o *NTMSLibrequestinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLibrequestinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DriveID != nil {
		if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PhysMediaID != nil {
		if err := o.PhysMediaID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SlotID != nil {
		if err := o.SlotID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Priority); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLibrequestinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.PartitionID = &NTMSGUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DriveID == nil {
		o.DriveID = &NTMSGUID{}
	}
	if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PhysMediaID == nil {
		o.PhysMediaID = &NTMSGUID{}
	}
	if err := o.PhysMediaID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SlotID == nil {
		o.SlotID = &NTMSGUID{}
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
		o.WorkItemID = &NTMSGUID{}
	}
	if err := o.WorkItemID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Priority); err != nil {
		return err
	}
	return nil
}

// NTMSLibrequestinformationw structure represents NTMS_LIBREQUESTINFORMATIONW RPC structure.
type NTMSLibrequestinformationw struct {
	OperationCode   uint32           `idl:"name:OperationCode" json:"operation_code"`
	OperationOption uint32           `idl:"name:OperationOption" json:"operation_option"`
	State           uint32           `idl:"name:State" json:"state"`
	PartitionID     *NTMSGUID        `idl:"name:PartitionId" json:"partition_id"`
	DriveID         *NTMSGUID        `idl:"name:DriveId" json:"drive_id"`
	PhysMediaID     *NTMSGUID        `idl:"name:PhysMediaId" json:"phys_media_id"`
	Library         *NTMSGUID        `idl:"name:Library" json:"library"`
	SlotID          *NTMSGUID        `idl:"name:SlotId" json:"slot_id"`
	TimeQueued      *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	TimeCompleted   *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	Application     string           `idl:"name:szApplication;string" json:"application"`
	User            string           `idl:"name:szUser;string" json:"user"`
	Computer        string           `idl:"name:szComputer;string" json:"computer"`
	ErrorCode       uint32           `idl:"name:dwErrorCode" json:"error_code"`
	WorkItemID      *NTMSGUID        `idl:"name:WorkItemId" json:"work_item_id"`
	Priority        uint32           `idl:"name:dwPriority" json:"priority"`
}

func (o *NTMSLibrequestinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLibrequestinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.DriveID != nil {
		if err := o.DriveID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PhysMediaID != nil {
		if err := o.PhysMediaID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SlotID != nil {
		if err := o.SlotID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSLibrequestinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.PartitionID = &NTMSGUID{}
	}
	if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.DriveID == nil {
		o.DriveID = &NTMSGUID{}
	}
	if err := o.DriveID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PhysMediaID == nil {
		o.PhysMediaID = &NTMSGUID{}
	}
	if err := o.PhysMediaID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Library == nil {
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SlotID == nil {
		o.SlotID = &NTMSGUID{}
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
		o.WorkItemID = &NTMSGUID{}
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

// NTMSMediapoolinformation structure represents NTMS_MEDIAPOOLINFORMATION RPC structure.
type NTMSMediapoolinformation struct {
	PoolType              uint32    `idl:"name:PoolType" json:"pool_type"`
	MediaType             *NTMSGUID `idl:"name:MediaType" json:"media_type"`
	Parent                *NTMSGUID `idl:"name:Parent" json:"parent"`
	AllocationPolicy      uint32    `idl:"name:AllocationPolicy" json:"allocation_policy"`
	DeallocationPolicy    uint32    `idl:"name:DeallocationPolicy" json:"deallocation_policy"`
	MaxAllocates          uint32    `idl:"name:dwMaxAllocates" json:"max_allocates"`
	NumberOfPhysicalMedia uint32    `idl:"name:dwNumberOfPhysicalMedia" json:"number_of_physical_media"`
	NumberOfLogicalMedia  uint32    `idl:"name:dwNumberOfLogicalMedia" json:"number_of_logical_media"`
	NumberOfMediaPools    uint32    `idl:"name:dwNumberOfMediaPools" json:"number_of_media_pools"`
}

func (o *NTMSMediapoolinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSMediapoolinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Parent != nil {
		if err := o.Parent.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSMediapoolinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PoolType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &NTMSGUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Parent == nil {
		o.Parent = &NTMSGUID{}
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

// NTMSMediatypeinformation structure represents NTMS_MEDIATYPEINFORMATION RPC structure.
type NTMSMediatypeinformation struct {
	MediaType                uint32 `idl:"name:MediaType" json:"media_type"`
	NumberOfSides            uint32 `idl:"name:NumberOfSides" json:"number_of_sides"`
	ReadWriteCharacteristics uint32 `idl:"name:ReadWriteCharacteristics" json:"read_write_characteristics"`
	DeviceType               uint32 `idl:"name:DeviceType" json:"device_type"`
}

func (o *NTMSMediatypeinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSMediatypeinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSMediatypeinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSStorageslotinformation structure represents NTMS_STORAGESLOTINFORMATION RPC structure.
type NTMSStorageslotinformation struct {
	Number  uint32    `idl:"name:Number" json:"number"`
	State   uint32    `idl:"name:State" json:"state"`
	Library *NTMSGUID `idl:"name:Library" json:"library"`
}

func (o *NTMSStorageslotinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSStorageslotinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSStorageslotinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSIedoorinformation structure represents NTMS_IEDOORINFORMATION RPC structure.
type NTMSIedoorinformation struct {
	Number      uint32    `idl:"name:Number" json:"number"`
	State       uint32    `idl:"name:State" json:"state"`
	MaxOpenSecs uint16    `idl:"name:MaxOpenSecs" json:"max_open_secs"`
	Library     *NTMSGUID `idl:"name:Library" json:"library"`
}

func (o *NTMSIedoorinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSIedoorinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSIedoorinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSIeportinformation structure represents NTMS_IEPORTINFORMATION RPC structure.
type NTMSIeportinformation struct {
	Number        uint32    `idl:"name:Number" json:"number"`
	Content       uint32    `idl:"name:Content" json:"content"`
	Position      uint32    `idl:"name:Position" json:"position"`
	MaxExtendSecs uint16    `idl:"name:MaxExtendSecs" json:"max_extend_secs"`
	Library       *NTMSGUID `idl:"name:Library" json:"library"`
}

func (o *NTMSIeportinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSIeportinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSIeportinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Library = &NTMSGUID{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSLmidinformation structure represents NTMS_LMIDINFORMATION RPC structure.
type NTMSLmidinformation struct {
	MediaPool          *NTMSGUID `idl:"name:MediaPool" json:"media_pool"`
	NumberOfPartitions uint32    `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
}

func (o *NTMSLmidinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLmidinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfPartitions); err != nil {
		return err
	}
	return nil
}
func (o *NTMSLmidinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &NTMSGUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfPartitions); err != nil {
		return err
	}
	return nil
}

// NTMSComputerinformation structure represents NTMS_COMPUTERINFORMATION RPC structure.
type NTMSComputerinformation struct {
	LibRequestPurgeTime       uint32 `idl:"name:dwLibRequestPurgeTime" json:"lib_request_purge_time"`
	OperationRequestPurgeTime uint32 `idl:"name:dwOpRequestPurgeTime" json:"operation_request_purge_time"`
	LibRequestFlags           uint32 `idl:"name:dwLibRequestFlags" json:"lib_request_flags"`
	OperationRequestFlags     uint32 `idl:"name:dwOpRequestFlags" json:"operation_request_flags"`
	MediaPoolPolicy           uint32 `idl:"name:dwMediaPoolPolicy" json:"media_pool_policy"`
}

func (o *NTMSComputerinformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSComputerinformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSComputerinformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSOprequestinformationa structure represents NTMS_OPREQUESTINFORMATIONA RPC structure.
type NTMSOprequestinformationa struct {
	Request     uint32           `idl:"name:Request" json:"request"`
	Submitted   *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	State       uint32           `idl:"name:State" json:"state"`
	Message     []byte           `idl:"name:szMessage" json:"message"`
	Arg1Type    uint32           `idl:"name:Arg1Type" json:"arg1_type"`
	Arg1        *NTMSGUID        `idl:"name:Arg1" json:"arg1"`
	Arg2Type    uint32           `idl:"name:Arg2Type" json:"arg2_type"`
	Arg2        *NTMSGUID        `idl:"name:Arg2" json:"arg2"`
	Application []byte           `idl:"name:szApplication" json:"application"`
	User        []byte           `idl:"name:szUser" json:"user"`
	Computer    []byte           `idl:"name:szComputer" json:"computer"`
}

func (o *NTMSOprequestinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSOprequestinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSOprequestinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Arg1 = &NTMSGUID{}
	}
	if err := o.Arg1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 == nil {
		o.Arg2 = &NTMSGUID{}
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

// NTMSOprequestinformationw structure represents NTMS_OPREQUESTINFORMATIONW RPC structure.
type NTMSOprequestinformationw struct {
	Request     uint32           `idl:"name:Request" json:"request"`
	Submitted   *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	State       uint32           `idl:"name:State" json:"state"`
	Message     string           `idl:"name:szMessage;string" json:"message"`
	Arg1Type    uint32           `idl:"name:Arg1Type" json:"arg1_type"`
	Arg1        *NTMSGUID        `idl:"name:Arg1" json:"arg1"`
	Arg2Type    uint32           `idl:"name:Arg2Type" json:"arg2_type"`
	Arg2        *NTMSGUID        `idl:"name:Arg2" json:"arg2"`
	Application string           `idl:"name:szApplication;string" json:"application"`
	User        string           `idl:"name:szUser;string" json:"user"`
	Computer    string           `idl:"name:szComputer;string" json:"computer"`
}

func (o *NTMSOprequestinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSOprequestinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *NTMSOprequestinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Arg1 = &NTMSGUID{}
	}
	if err := o.Arg1.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Arg2Type); err != nil {
		return err
	}
	if o.Arg2 == nil {
		o.Arg2 = &NTMSGUID{}
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

// NTMSPartitioninformationa structure represents NTMS_PARTITIONINFORMATIONA RPC structure.
type NTMSPartitioninformationa struct {
	PhysicalMedia     *NTMSGUID          `idl:"name:PhysicalMedia" json:"physical_media"`
	LogicalMedia      *NTMSGUID          `idl:"name:LogicalMedia" json:"logical_media"`
	State             uint32             `idl:"name:State" json:"state"`
	Side              uint16             `idl:"name:Side" json:"side"`
	OmidLabelIDLength uint32             `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	OmidLabelID       []byte             `idl:"name:OmidLabelId" json:"omid_label_id"`
	OmidLabelType     []byte             `idl:"name:szOmidLabelType" json:"omid_label_type"`
	OmidLabelInfo     []byte             `idl:"name:szOmidLabelInfo" json:"omid_label_info"`
	MountCount        uint32             `idl:"name:dwMountCount" json:"mount_count"`
	AllocateCount     uint32             `idl:"name:dwAllocateCount" json:"allocate_count"`
	Capacity          *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
}

func (o *NTMSPartitioninformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSPartitioninformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Side); err != nil {
		return err
	}
	if err := w.WriteData(o.OmidLabelIDLength); err != nil {
		return err
	}
	for i1 := range o.OmidLabelID {
		i1 := i1
		if uint64(i1) >= 255 {
			break
		}
		if err := w.WriteData(o.OmidLabelID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OmidLabelID); uint64(i1) < 255; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.OmidLabelType {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.OmidLabelType[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OmidLabelType); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.OmidLabelInfo {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.OmidLabelInfo[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OmidLabelInfo); uint64(i1) < 256; i1++ {
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
func (o *NTMSPartitioninformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &NTMSGUID{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogicalMedia == nil {
		o.LogicalMedia = &NTMSGUID{}
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
	if err := w.ReadData(&o.OmidLabelIDLength); err != nil {
		return err
	}
	o.OmidLabelID = make([]byte, 255)
	for i1 := range o.OmidLabelID {
		i1 := i1
		if err := w.ReadData(&o.OmidLabelID[i1]); err != nil {
			return err
		}
	}
	o.OmidLabelType = make([]byte, 64)
	for i1 := range o.OmidLabelType {
		i1 := i1
		if err := w.ReadData(&o.OmidLabelType[i1]); err != nil {
			return err
		}
	}
	o.OmidLabelInfo = make([]byte, 256)
	for i1 := range o.OmidLabelInfo {
		i1 := i1
		if err := w.ReadData(&o.OmidLabelInfo[i1]); err != nil {
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

// NTMSPartitioninformationw structure represents NTMS_PARTITIONINFORMATIONW RPC structure.
type NTMSPartitioninformationw struct {
	PhysicalMedia     *NTMSGUID          `idl:"name:PhysicalMedia" json:"physical_media"`
	LogicalMedia      *NTMSGUID          `idl:"name:LogicalMedia" json:"logical_media"`
	State             uint32             `idl:"name:State" json:"state"`
	Side              uint16             `idl:"name:Side" json:"side"`
	OmidLabelIDLength uint32             `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	OmidLabelID       []byte             `idl:"name:OmidLabelId" json:"omid_label_id"`
	OmidLabelType     string             `idl:"name:szOmidLabelType;string" json:"omid_label_type"`
	OmidLabelInfo     string             `idl:"name:szOmidLabelInfo;string" json:"omid_label_info"`
	MountCount        uint32             `idl:"name:dwMountCount" json:"mount_count"`
	AllocateCount     uint32             `idl:"name:dwAllocateCount" json:"allocate_count"`
	Capacity          *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
}

func (o *NTMSPartitioninformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSPartitioninformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Side); err != nil {
		return err
	}
	if err := w.WriteData(o.OmidLabelIDLength); err != nil {
		return err
	}
	for i1 := range o.OmidLabelID {
		i1 := i1
		if uint64(i1) >= 255 {
			break
		}
		if err := w.WriteData(o.OmidLabelID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.OmidLabelID); uint64(i1) < 255; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	_OmidLabelType_buf := utf16.Encode([]rune(o.OmidLabelType))
	if uint64(len(_OmidLabelType_buf)) > 64-1 {
		_OmidLabelType_buf = _OmidLabelType_buf[:64-1]
	}
	if o.OmidLabelType != ndr.ZeroString {
		_OmidLabelType_buf = append(_OmidLabelType_buf, uint16(0))
	}
	for i1 := range _OmidLabelType_buf {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(_OmidLabelType_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_OmidLabelType_buf); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	_OmidLabelInfo_buf := utf16.Encode([]rune(o.OmidLabelInfo))
	if uint64(len(_OmidLabelInfo_buf)) > 256-1 {
		_OmidLabelInfo_buf = _OmidLabelInfo_buf[:256-1]
	}
	if o.OmidLabelInfo != ndr.ZeroString {
		_OmidLabelInfo_buf = append(_OmidLabelInfo_buf, uint16(0))
	}
	for i1 := range _OmidLabelInfo_buf {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(_OmidLabelInfo_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_OmidLabelInfo_buf); uint64(i1) < 256; i1++ {
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
func (o *NTMSPartitioninformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &NTMSGUID{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogicalMedia == nil {
		o.LogicalMedia = &NTMSGUID{}
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
	if err := w.ReadData(&o.OmidLabelIDLength); err != nil {
		return err
	}
	o.OmidLabelID = make([]byte, 255)
	for i1 := range o.OmidLabelID {
		i1 := i1
		if err := w.ReadData(&o.OmidLabelID[i1]); err != nil {
			return err
		}
	}
	var _OmidLabelType_buf []uint16
	_OmidLabelType_buf = make([]uint16, 64)
	for i1 := range _OmidLabelType_buf {
		i1 := i1
		if err := w.ReadData(&_OmidLabelType_buf[i1]); err != nil {
			return err
		}
	}
	o.OmidLabelType = strings.TrimRight(string(utf16.Decode(_OmidLabelType_buf)), ndr.ZeroString)
	var _OmidLabelInfo_buf []uint16
	_OmidLabelInfo_buf = make([]uint16, 256)
	for i1 := range _OmidLabelInfo_buf {
		i1 := i1
		if err := w.ReadData(&_OmidLabelInfo_buf[i1]); err != nil {
			return err
		}
	}
	o.OmidLabelInfo = strings.TrimRight(string(utf16.Decode(_OmidLabelInfo_buf)), ndr.ZeroString)
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

// NTMSPmidinformationa structure represents NTMS_PMIDINFORMATIONA RPC structure.
type NTMSPmidinformationa struct {
	CurrentLibrary     *NTMSGUID `idl:"name:CurrentLibrary" json:"current_library"`
	MediaPool          *NTMSGUID `idl:"name:MediaPool" json:"media_pool"`
	Location           *NTMSGUID `idl:"name:Location" json:"location"`
	LocationType       uint32    `idl:"name:LocationType" json:"location_type"`
	MediaType          *NTMSGUID `idl:"name:MediaType" json:"media_type"`
	HomeSlot           *NTMSGUID `idl:"name:HomeSlot" json:"home_slot"`
	BarCode            []byte    `idl:"name:szBarCode" json:"bar_code"`
	BarCodeState       uint32    `idl:"name:BarCodeState" json:"bar_code_state"`
	SequenceNumber     []byte    `idl:"name:szSequenceNumber" json:"sequence_number"`
	MediaState         uint32    `idl:"name:MediaState" json:"media_state"`
	NumberOfPartitions uint32    `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	MediaTypeCode      uint32    `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	DensityCode        uint32    `idl:"name:dwDensityCode" json:"density_code"`
	MountedPartition   *NTMSGUID `idl:"name:MountedPartition" json:"mounted_partition"`
}

func (o *NTMSPmidinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSPmidinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Location != nil {
		if err := o.Location.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeSlot != nil {
		if err := o.HomeSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSPmidinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.CurrentLibrary == nil {
		o.CurrentLibrary = &NTMSGUID{}
	}
	if err := o.CurrentLibrary.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &NTMSGUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Location == nil {
		o.Location = &NTMSGUID{}
	}
	if err := o.Location.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocationType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &NTMSGUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeSlot == nil {
		o.HomeSlot = &NTMSGUID{}
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
		o.MountedPartition = &NTMSGUID{}
	}
	if err := o.MountedPartition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSPmidinformationw structure represents NTMS_PMIDINFORMATIONW RPC structure.
type NTMSPmidinformationw struct {
	CurrentLibrary     *NTMSGUID `idl:"name:CurrentLibrary" json:"current_library"`
	MediaPool          *NTMSGUID `idl:"name:MediaPool" json:"media_pool"`
	Location           *NTMSGUID `idl:"name:Location" json:"location"`
	LocationType       uint32    `idl:"name:LocationType" json:"location_type"`
	MediaType          *NTMSGUID `idl:"name:MediaType" json:"media_type"`
	HomeSlot           *NTMSGUID `idl:"name:HomeSlot" json:"home_slot"`
	BarCode            string    `idl:"name:szBarCode;string" json:"bar_code"`
	BarCodeState       uint32    `idl:"name:BarCodeState" json:"bar_code_state"`
	SequenceNumber     string    `idl:"name:szSequenceNumber;string" json:"sequence_number"`
	MediaState         uint32    `idl:"name:MediaState" json:"media_state"`
	NumberOfPartitions uint32    `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	MediaTypeCode      uint32    `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	DensityCode        uint32    `idl:"name:dwDensityCode" json:"density_code"`
	MountedPartition   *NTMSGUID `idl:"name:MountedPartition" json:"mounted_partition"`
}

func (o *NTMSPmidinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSPmidinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Location != nil {
		if err := o.Location.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeSlot != nil {
		if err := o.HomeSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *NTMSPmidinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.CurrentLibrary == nil {
		o.CurrentLibrary = &NTMSGUID{}
	}
	if err := o.CurrentLibrary.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.MediaPool == nil {
		o.MediaPool = &NTMSGUID{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Location == nil {
		o.Location = &NTMSGUID{}
	}
	if err := o.Location.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocationType); err != nil {
		return err
	}
	if o.MediaType == nil {
		o.MediaType = &NTMSGUID{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeSlot == nil {
		o.HomeSlot = &NTMSGUID{}
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
		o.MountedPartition = &NTMSGUID{}
	}
	if err := o.MountedPartition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// RsmMessage structure represents RSM_MESSAGE RPC structure.
type RsmMessage struct {
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

func (o *RsmMessage) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RsmMessage) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RsmMessage) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSObjectinformationa structure represents NTMS_OBJECTINFORMATIONA RPC structure.
type NTMSObjectinformationa struct {
	Size             uint32                       `idl:"name:dwSize" json:"size"`
	Type             uint32                       `idl:"name:dwType" json:"type"`
	Created          *dtyp.SystemTime             `idl:"name:Created" json:"created"`
	Modified         *dtyp.SystemTime             `idl:"name:Modified" json:"modified"`
	ObjectGUID       *NTMSGUID                    `idl:"name:ObjectGuid" json:"object_guid"`
	Enabled          bool                         `idl:"name:Enabled" json:"enabled"`
	OperationalState uint32                       `idl:"name:dwOperationalState" json:"operational_state"`
	Name             []byte                       `idl:"name:szName" json:"name"`
	Description      []byte                       `idl:"name:szDescription" json:"description"`
	Info             *NTMSObjectinformationa_Info `idl:"name:Info;switch_is:dwType" json:"info"`
}

func (o *NTMSObjectinformationa) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSObjectinformationa) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSObjectinformationa_Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.ObjectGUID = &NTMSGUID{}
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
		o.Info = &NTMSObjectinformationa_Info{}
	}
	_swInfo := uint32(o.Type)
	if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info structure represents NTMS_OBJECTINFORMATIONA union anonymous member.
type NTMSObjectinformationa_Info struct {
	// Types that are assignable to Value
	//
	// *NTMSObjectinformationa_Info_Drive
	// *NTMSObjectinformationa_Info_DriveType
	// *NTMSObjectinformationa_Info_Library
	// *NTMSObjectinformationa_Info_Changer
	// *NTMSObjectinformationa_Info_ChangerType
	// *NTMSObjectinformationa_Info_StorageSlot
	// *NTMSObjectinformationa_Info_IeDoor
	// *NTMSObjectinformationa_Info_IePort
	// *NTMSObjectinformationa_Info_PhysicalMedia
	// *NTMSObjectinformationa_Info_LogicalMedia
	// *NTMSObjectinformationa_Info_Partition
	// *NTMSObjectinformationa_Info_MediaPool
	// *NTMSObjectinformationa_Info_MediaType
	// *NTMSObjectinformationa_Info_LibRequest
	// *NTMSObjectinformationa_Info_OperationRequest
	// *NTMSObjectinformationa_Info_Computer
	Value is_NTMSObjectinformationa_Info `json:"value"`
}

func (o *NTMSObjectinformationa_Info) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *NTMSObjectinformationa_Info_Drive:
		if value != nil {
			return value.Drive
		}
	case *NTMSObjectinformationa_Info_DriveType:
		if value != nil {
			return value.DriveType
		}
	case *NTMSObjectinformationa_Info_Library:
		if value != nil {
			return value.Library
		}
	case *NTMSObjectinformationa_Info_Changer:
		if value != nil {
			return value.Changer
		}
	case *NTMSObjectinformationa_Info_ChangerType:
		if value != nil {
			return value.ChangerType
		}
	case *NTMSObjectinformationa_Info_StorageSlot:
		if value != nil {
			return value.StorageSlot
		}
	case *NTMSObjectinformationa_Info_IeDoor:
		if value != nil {
			return value.IeDoor
		}
	case *NTMSObjectinformationa_Info_IePort:
		if value != nil {
			return value.IePort
		}
	case *NTMSObjectinformationa_Info_PhysicalMedia:
		if value != nil {
			return value.PhysicalMedia
		}
	case *NTMSObjectinformationa_Info_LogicalMedia:
		if value != nil {
			return value.LogicalMedia
		}
	case *NTMSObjectinformationa_Info_Partition:
		if value != nil {
			return value.Partition
		}
	case *NTMSObjectinformationa_Info_MediaPool:
		if value != nil {
			return value.MediaPool
		}
	case *NTMSObjectinformationa_Info_MediaType:
		if value != nil {
			return value.MediaType
		}
	case *NTMSObjectinformationa_Info_LibRequest:
		if value != nil {
			return value.LibRequest
		}
	case *NTMSObjectinformationa_Info_OperationRequest:
		if value != nil {
			return value.OperationRequest
		}
	case *NTMSObjectinformationa_Info_Computer:
		if value != nil {
			return value.Computer
		}
	}
	return nil
}

type is_NTMSObjectinformationa_Info interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_NTMSObjectinformationa_Info()
}

func (o *NTMSObjectinformationa_Info) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *NTMSObjectinformationa_Info_Drive:
		return uint32(5)
	case *NTMSObjectinformationa_Info_DriveType:
		return uint32(6)
	case *NTMSObjectinformationa_Info_Library:
		return uint32(9)
	case *NTMSObjectinformationa_Info_Changer:
		return uint32(2)
	case *NTMSObjectinformationa_Info_ChangerType:
		return uint32(3)
	case *NTMSObjectinformationa_Info_StorageSlot:
		return uint32(16)
	case *NTMSObjectinformationa_Info_IeDoor:
		return uint32(7)
	case *NTMSObjectinformationa_Info_IePort:
		return uint32(8)
	case *NTMSObjectinformationa_Info_PhysicalMedia:
		return uint32(15)
	case *NTMSObjectinformationa_Info_LogicalMedia:
		return uint32(11)
	case *NTMSObjectinformationa_Info_Partition:
		return uint32(14)
	case *NTMSObjectinformationa_Info_MediaPool:
		return uint32(12)
	case *NTMSObjectinformationa_Info_MediaType:
		return uint32(13)
	case *NTMSObjectinformationa_Info_LibRequest:
		return uint32(10)
	case *NTMSObjectinformationa_Info_OperationRequest:
		return uint32(17)
	case *NTMSObjectinformationa_Info_Computer:
		return uint32(4)
	}
	return uint32(0)
}

func (o *NTMSObjectinformationa_Info) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_Drive)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_Drive{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_DriveType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_DriveType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_Library)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_Library{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_Changer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_Changer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_ChangerType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_ChangerType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_StorageSlot)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_StorageSlot{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_IeDoor)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_IeDoor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_IePort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_IePort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(15):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_PhysicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_PhysicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_LogicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_LogicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_Partition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_Partition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_MediaPool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_MediaPool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_MediaType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_MediaType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_LibRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_LibRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_OperationRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_OperationRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*NTMSObjectinformationa_Info_Computer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationa_Info_Computer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *NTMSObjectinformationa_Info) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &NTMSObjectinformationa_Info_Drive{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &NTMSObjectinformationa_Info_DriveType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &NTMSObjectinformationa_Info_Library{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &NTMSObjectinformationa_Info_Changer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &NTMSObjectinformationa_Info_ChangerType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &NTMSObjectinformationa_Info_StorageSlot{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &NTMSObjectinformationa_Info_IeDoor{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &NTMSObjectinformationa_Info_IePort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(15):
		o.Value = &NTMSObjectinformationa_Info_PhysicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &NTMSObjectinformationa_Info_LogicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &NTMSObjectinformationa_Info_Partition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &NTMSObjectinformationa_Info_MediaPool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &NTMSObjectinformationa_Info_MediaType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &NTMSObjectinformationa_Info_LibRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17):
		o.Value = &NTMSObjectinformationa_Info_OperationRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &NTMSObjectinformationa_Info_Computer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// NTMSObjectinformationa_Info_Drive structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 5
type NTMSObjectinformationa_Info_Drive struct {
	Drive *NTMSDriveinformationa `idl:"name:Drive" json:"drive"`
}

func (*NTMSObjectinformationa_Info_Drive) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_Drive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Drive != nil {
		if err := o.Drive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSDriveinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_Drive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Drive == nil {
		o.Drive = &NTMSDriveinformationa{}
	}
	if err := o.Drive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_DriveType structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 6
type NTMSObjectinformationa_Info_DriveType struct {
	DriveType *NTMSDrivetypeinformationa `idl:"name:DriveType" json:"drive_type"`
}

func (*NTMSObjectinformationa_Info_DriveType) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_DriveType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSDrivetypeinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_DriveType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DriveType == nil {
		o.DriveType = &NTMSDrivetypeinformationa{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_Library structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 9
type NTMSObjectinformationa_Info_Library struct {
	Library *NTMSLibraryinformation `idl:"name:Library" json:"library"`
}

func (*NTMSObjectinformationa_Info_Library) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_Library) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLibraryinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_Library) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Library == nil {
		o.Library = &NTMSLibraryinformation{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_Changer structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 2
type NTMSObjectinformationa_Info_Changer struct {
	Changer *NTMSChangerinformationa `idl:"name:Changer" json:"changer"`
}

func (*NTMSObjectinformationa_Info_Changer) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_Changer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Changer != nil {
		if err := o.Changer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSChangerinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_Changer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Changer == nil {
		o.Changer = &NTMSChangerinformationa{}
	}
	if err := o.Changer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_ChangerType structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 3
type NTMSObjectinformationa_Info_ChangerType struct {
	ChangerType *NTMSChangertypeinformationa `idl:"name:ChangerType" json:"changer_type"`
}

func (*NTMSObjectinformationa_Info_ChangerType) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_ChangerType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSChangertypeinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_ChangerType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ChangerType == nil {
		o.ChangerType = &NTMSChangertypeinformationa{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_StorageSlot structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 16
type NTMSObjectinformationa_Info_StorageSlot struct {
	StorageSlot *NTMSStorageslotinformation `idl:"name:StorageSlot" json:"storage_slot"`
}

func (*NTMSObjectinformationa_Info_StorageSlot) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_StorageSlot) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StorageSlot != nil {
		if err := o.StorageSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSStorageslotinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_StorageSlot) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StorageSlot == nil {
		o.StorageSlot = &NTMSStorageslotinformation{}
	}
	if err := o.StorageSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_IeDoor structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 7
type NTMSObjectinformationa_Info_IeDoor struct {
	IeDoor *NTMSIedoorinformation `idl:"name:IEDoor" json:"ie_door"`
}

func (*NTMSObjectinformationa_Info_IeDoor) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_IeDoor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IeDoor != nil {
		if err := o.IeDoor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSIedoorinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_IeDoor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IeDoor == nil {
		o.IeDoor = &NTMSIedoorinformation{}
	}
	if err := o.IeDoor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_IePort structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 8
type NTMSObjectinformationa_Info_IePort struct {
	IePort *NTMSIeportinformation `idl:"name:IEPort" json:"ie_port"`
}

func (*NTMSObjectinformationa_Info_IePort) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_IePort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IePort != nil {
		if err := o.IePort.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSIeportinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_IePort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IePort == nil {
		o.IePort = &NTMSIeportinformation{}
	}
	if err := o.IePort.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_PhysicalMedia structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 15
type NTMSObjectinformationa_Info_PhysicalMedia struct {
	PhysicalMedia *NTMSPmidinformationa `idl:"name:PhysicalMedia" json:"physical_media"`
}

func (*NTMSObjectinformationa_Info_PhysicalMedia) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_PhysicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSPmidinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_PhysicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &NTMSPmidinformationa{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_LogicalMedia structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 11
type NTMSObjectinformationa_Info_LogicalMedia struct {
	LogicalMedia *NTMSLmidinformation `idl:"name:LogicalMedia" json:"logical_media"`
}

func (*NTMSObjectinformationa_Info_LogicalMedia) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_LogicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLmidinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_LogicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LogicalMedia == nil {
		o.LogicalMedia = &NTMSLmidinformation{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_Partition structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 14
type NTMSObjectinformationa_Info_Partition struct {
	Partition *NTMSPartitioninformationa `idl:"name:Partition" json:"partition"`
}

func (*NTMSObjectinformationa_Info_Partition) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_Partition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSPartitioninformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_Partition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Partition == nil {
		o.Partition = &NTMSPartitioninformationa{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_MediaPool structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 12
type NTMSObjectinformationa_Info_MediaPool struct {
	MediaPool *NTMSMediapoolinformation `idl:"name:MediaPool" json:"media_pool"`
}

func (*NTMSObjectinformationa_Info_MediaPool) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_MediaPool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSMediapoolinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_MediaPool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaPool == nil {
		o.MediaPool = &NTMSMediapoolinformation{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_MediaType structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 13
type NTMSObjectinformationa_Info_MediaType struct {
	MediaType *NTMSMediatypeinformation `idl:"name:MediaType" json:"media_type"`
}

func (*NTMSObjectinformationa_Info_MediaType) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_MediaType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSMediatypeinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_MediaType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaType == nil {
		o.MediaType = &NTMSMediatypeinformation{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_LibRequest structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 10
type NTMSObjectinformationa_Info_LibRequest struct {
	LibRequest *NTMSLibrequestinformationa `idl:"name:LibRequest" json:"lib_request"`
}

func (*NTMSObjectinformationa_Info_LibRequest) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_LibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LibRequest != nil {
		if err := o.LibRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLibrequestinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_LibRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LibRequest == nil {
		o.LibRequest = &NTMSLibrequestinformationa{}
	}
	if err := o.LibRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_OperationRequest structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 17
type NTMSObjectinformationa_Info_OperationRequest struct {
	OperationRequest *NTMSOprequestinformationa `idl:"name:OpRequest" json:"operation_request"`
}

func (*NTMSObjectinformationa_Info_OperationRequest) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_OperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OperationRequest != nil {
		if err := o.OperationRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSOprequestinformationa{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_OperationRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OperationRequest == nil {
		o.OperationRequest = &NTMSOprequestinformationa{}
	}
	if err := o.OperationRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationa_Info_Computer structure represents NTMSObjectinformationa_Info RPC union arm.
//
// It has following labels: 4
type NTMSObjectinformationa_Info_Computer struct {
	Computer *NTMSComputerinformation `idl:"name:Computer" json:"computer"`
}

func (*NTMSObjectinformationa_Info_Computer) is_NTMSObjectinformationa_Info() {}

func (o *NTMSObjectinformationa_Info_Computer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Computer != nil {
		if err := o.Computer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSComputerinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationa_Info_Computer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Computer == nil {
		o.Computer = &NTMSComputerinformation{}
	}
	if err := o.Computer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw structure represents NTMS_OBJECTINFORMATIONW RPC structure.
type NTMSObjectinformationw struct {
	Size             uint32                       `idl:"name:dwSize" json:"size"`
	Type             uint32                       `idl:"name:dwType" json:"type"`
	Created          *dtyp.SystemTime             `idl:"name:Created" json:"created"`
	Modified         *dtyp.SystemTime             `idl:"name:Modified" json:"modified"`
	ObjectGUID       *NTMSGUID                    `idl:"name:ObjectGuid" json:"object_guid"`
	Enabled          bool                         `idl:"name:Enabled" json:"enabled"`
	OperationalState uint32                       `idl:"name:dwOperationalState" json:"operational_state"`
	Name             string                       `idl:"name:szName;string" json:"name"`
	Description      string                       `idl:"name:szDescription;string" json:"description"`
	Info             *NTMSObjectinformationw_Info `idl:"name:Info;switch_is:dwType" json:"info"`
}

func (o *NTMSObjectinformationw) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *NTMSObjectinformationw) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&NTMSGUID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&NTMSObjectinformationw_Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.ObjectGUID = &NTMSGUID{}
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
		o.Info = &NTMSObjectinformationw_Info{}
	}
	_swInfo := uint32(o.Type)
	if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info structure represents NTMS_OBJECTINFORMATIONW union anonymous member.
type NTMSObjectinformationw_Info struct {
	// Types that are assignable to Value
	//
	// *NTMSObjectinformationw_Info_Drive
	// *NTMSObjectinformationw_Info_DriveType
	// *NTMSObjectinformationw_Info_Library
	// *NTMSObjectinformationw_Info_Changer
	// *NTMSObjectinformationw_Info_ChangerType
	// *NTMSObjectinformationw_Info_StorageSlot
	// *NTMSObjectinformationw_Info_IeDoor
	// *NTMSObjectinformationw_Info_IePort
	// *NTMSObjectinformationw_Info_PhysicalMedia
	// *NTMSObjectinformationw_Info_LogicalMedia
	// *NTMSObjectinformationw_Info_Partition
	// *NTMSObjectinformationw_Info_MediaPool
	// *NTMSObjectinformationw_Info_MediaType
	// *NTMSObjectinformationw_Info_LibRequest
	// *NTMSObjectinformationw_Info_OperationRequest
	// *NTMSObjectinformationw_Info_Computer
	Value is_NTMSObjectinformationw_Info `json:"value"`
}

func (o *NTMSObjectinformationw_Info) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *NTMSObjectinformationw_Info_Drive:
		if value != nil {
			return value.Drive
		}
	case *NTMSObjectinformationw_Info_DriveType:
		if value != nil {
			return value.DriveType
		}
	case *NTMSObjectinformationw_Info_Library:
		if value != nil {
			return value.Library
		}
	case *NTMSObjectinformationw_Info_Changer:
		if value != nil {
			return value.Changer
		}
	case *NTMSObjectinformationw_Info_ChangerType:
		if value != nil {
			return value.ChangerType
		}
	case *NTMSObjectinformationw_Info_StorageSlot:
		if value != nil {
			return value.StorageSlot
		}
	case *NTMSObjectinformationw_Info_IeDoor:
		if value != nil {
			return value.IeDoor
		}
	case *NTMSObjectinformationw_Info_IePort:
		if value != nil {
			return value.IePort
		}
	case *NTMSObjectinformationw_Info_PhysicalMedia:
		if value != nil {
			return value.PhysicalMedia
		}
	case *NTMSObjectinformationw_Info_LogicalMedia:
		if value != nil {
			return value.LogicalMedia
		}
	case *NTMSObjectinformationw_Info_Partition:
		if value != nil {
			return value.Partition
		}
	case *NTMSObjectinformationw_Info_MediaPool:
		if value != nil {
			return value.MediaPool
		}
	case *NTMSObjectinformationw_Info_MediaType:
		if value != nil {
			return value.MediaType
		}
	case *NTMSObjectinformationw_Info_LibRequest:
		if value != nil {
			return value.LibRequest
		}
	case *NTMSObjectinformationw_Info_OperationRequest:
		if value != nil {
			return value.OperationRequest
		}
	case *NTMSObjectinformationw_Info_Computer:
		if value != nil {
			return value.Computer
		}
	}
	return nil
}

type is_NTMSObjectinformationw_Info interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_NTMSObjectinformationw_Info()
}

func (o *NTMSObjectinformationw_Info) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *NTMSObjectinformationw_Info_Drive:
		return uint32(5)
	case *NTMSObjectinformationw_Info_DriveType:
		return uint32(6)
	case *NTMSObjectinformationw_Info_Library:
		return uint32(9)
	case *NTMSObjectinformationw_Info_Changer:
		return uint32(2)
	case *NTMSObjectinformationw_Info_ChangerType:
		return uint32(3)
	case *NTMSObjectinformationw_Info_StorageSlot:
		return uint32(16)
	case *NTMSObjectinformationw_Info_IeDoor:
		return uint32(7)
	case *NTMSObjectinformationw_Info_IePort:
		return uint32(8)
	case *NTMSObjectinformationw_Info_PhysicalMedia:
		return uint32(15)
	case *NTMSObjectinformationw_Info_LogicalMedia:
		return uint32(11)
	case *NTMSObjectinformationw_Info_Partition:
		return uint32(14)
	case *NTMSObjectinformationw_Info_MediaPool:
		return uint32(12)
	case *NTMSObjectinformationw_Info_MediaType:
		return uint32(13)
	case *NTMSObjectinformationw_Info_LibRequest:
		return uint32(10)
	case *NTMSObjectinformationw_Info_OperationRequest:
		return uint32(17)
	case *NTMSObjectinformationw_Info_Computer:
		return uint32(4)
	}
	return uint32(0)
}

func (o *NTMSObjectinformationw_Info) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_Drive)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_Drive{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_DriveType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_DriveType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_Library)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_Library{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_Changer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_Changer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_ChangerType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_ChangerType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_StorageSlot)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_StorageSlot{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_IeDoor)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_IeDoor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_IePort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_IePort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(15):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_PhysicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_PhysicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_LogicalMedia)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_LogicalMedia{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_Partition)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_Partition{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_MediaPool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_MediaPool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_MediaType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_MediaType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_LibRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_LibRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_OperationRequest)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_OperationRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*NTMSObjectinformationw_Info_Computer)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NTMSObjectinformationw_Info_Computer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *NTMSObjectinformationw_Info) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &NTMSObjectinformationw_Info_Drive{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &NTMSObjectinformationw_Info_DriveType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &NTMSObjectinformationw_Info_Library{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &NTMSObjectinformationw_Info_Changer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &NTMSObjectinformationw_Info_ChangerType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &NTMSObjectinformationw_Info_StorageSlot{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &NTMSObjectinformationw_Info_IeDoor{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &NTMSObjectinformationw_Info_IePort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(15):
		o.Value = &NTMSObjectinformationw_Info_PhysicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &NTMSObjectinformationw_Info_LogicalMedia{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &NTMSObjectinformationw_Info_Partition{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &NTMSObjectinformationw_Info_MediaPool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &NTMSObjectinformationw_Info_MediaType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &NTMSObjectinformationw_Info_LibRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17):
		o.Value = &NTMSObjectinformationw_Info_OperationRequest{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &NTMSObjectinformationw_Info_Computer{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// NTMSObjectinformationw_Info_Drive structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 5
type NTMSObjectinformationw_Info_Drive struct {
	Drive *NTMSDriveinformationw `idl:"name:Drive" json:"drive"`
}

func (*NTMSObjectinformationw_Info_Drive) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_Drive) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Drive != nil {
		if err := o.Drive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSDriveinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_Drive) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Drive == nil {
		o.Drive = &NTMSDriveinformationw{}
	}
	if err := o.Drive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_DriveType structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 6
type NTMSObjectinformationw_Info_DriveType struct {
	DriveType *NTMSDrivetypeinformationw `idl:"name:DriveType" json:"drive_type"`
}

func (*NTMSObjectinformationw_Info_DriveType) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_DriveType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DriveType != nil {
		if err := o.DriveType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSDrivetypeinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_DriveType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DriveType == nil {
		o.DriveType = &NTMSDrivetypeinformationw{}
	}
	if err := o.DriveType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_Library structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 9
type NTMSObjectinformationw_Info_Library struct {
	Library *NTMSLibraryinformation `idl:"name:Library" json:"library"`
}

func (*NTMSObjectinformationw_Info_Library) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_Library) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Library != nil {
		if err := o.Library.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLibraryinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_Library) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Library == nil {
		o.Library = &NTMSLibraryinformation{}
	}
	if err := o.Library.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_Changer structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 2
type NTMSObjectinformationw_Info_Changer struct {
	Changer *NTMSChangerinformationw `idl:"name:Changer" json:"changer"`
}

func (*NTMSObjectinformationw_Info_Changer) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_Changer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Changer != nil {
		if err := o.Changer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSChangerinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_Changer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Changer == nil {
		o.Changer = &NTMSChangerinformationw{}
	}
	if err := o.Changer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_ChangerType structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 3
type NTMSObjectinformationw_Info_ChangerType struct {
	ChangerType *NTMSChangertypeinformationw `idl:"name:ChangerType" json:"changer_type"`
}

func (*NTMSObjectinformationw_Info_ChangerType) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_ChangerType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ChangerType != nil {
		if err := o.ChangerType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSChangertypeinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_ChangerType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ChangerType == nil {
		o.ChangerType = &NTMSChangertypeinformationw{}
	}
	if err := o.ChangerType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_StorageSlot structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 16
type NTMSObjectinformationw_Info_StorageSlot struct {
	StorageSlot *NTMSStorageslotinformation `idl:"name:StorageSlot" json:"storage_slot"`
}

func (*NTMSObjectinformationw_Info_StorageSlot) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_StorageSlot) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StorageSlot != nil {
		if err := o.StorageSlot.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSStorageslotinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_StorageSlot) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StorageSlot == nil {
		o.StorageSlot = &NTMSStorageslotinformation{}
	}
	if err := o.StorageSlot.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_IeDoor structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 7
type NTMSObjectinformationw_Info_IeDoor struct {
	IeDoor *NTMSIedoorinformation `idl:"name:IEDoor" json:"ie_door"`
}

func (*NTMSObjectinformationw_Info_IeDoor) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_IeDoor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IeDoor != nil {
		if err := o.IeDoor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSIedoorinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_IeDoor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IeDoor == nil {
		o.IeDoor = &NTMSIedoorinformation{}
	}
	if err := o.IeDoor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_IePort structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 8
type NTMSObjectinformationw_Info_IePort struct {
	IePort *NTMSIeportinformation `idl:"name:IEPort" json:"ie_port"`
}

func (*NTMSObjectinformationw_Info_IePort) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_IePort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IePort != nil {
		if err := o.IePort.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSIeportinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_IePort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.IePort == nil {
		o.IePort = &NTMSIeportinformation{}
	}
	if err := o.IePort.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_PhysicalMedia structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 15
type NTMSObjectinformationw_Info_PhysicalMedia struct {
	PhysicalMedia *NTMSPmidinformationw `idl:"name:PhysicalMedia" json:"physical_media"`
}

func (*NTMSObjectinformationw_Info_PhysicalMedia) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_PhysicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PhysicalMedia != nil {
		if err := o.PhysicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSPmidinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_PhysicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PhysicalMedia == nil {
		o.PhysicalMedia = &NTMSPmidinformationw{}
	}
	if err := o.PhysicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_LogicalMedia structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 11
type NTMSObjectinformationw_Info_LogicalMedia struct {
	LogicalMedia *NTMSLmidinformation `idl:"name:LogicalMedia" json:"logical_media"`
}

func (*NTMSObjectinformationw_Info_LogicalMedia) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_LogicalMedia) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LogicalMedia != nil {
		if err := o.LogicalMedia.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLmidinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_LogicalMedia) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LogicalMedia == nil {
		o.LogicalMedia = &NTMSLmidinformation{}
	}
	if err := o.LogicalMedia.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_Partition structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 14
type NTMSObjectinformationw_Info_Partition struct {
	Partition *NTMSPartitioninformationw `idl:"name:Partition" json:"partition"`
}

func (*NTMSObjectinformationw_Info_Partition) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_Partition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSPartitioninformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_Partition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Partition == nil {
		o.Partition = &NTMSPartitioninformationw{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_MediaPool structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 12
type NTMSObjectinformationw_Info_MediaPool struct {
	MediaPool *NTMSMediapoolinformation `idl:"name:MediaPool" json:"media_pool"`
}

func (*NTMSObjectinformationw_Info_MediaPool) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_MediaPool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaPool != nil {
		if err := o.MediaPool.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSMediapoolinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_MediaPool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaPool == nil {
		o.MediaPool = &NTMSMediapoolinformation{}
	}
	if err := o.MediaPool.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_MediaType structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 13
type NTMSObjectinformationw_Info_MediaType struct {
	MediaType *NTMSMediatypeinformation `idl:"name:MediaType" json:"media_type"`
}

func (*NTMSObjectinformationw_Info_MediaType) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_MediaType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MediaType != nil {
		if err := o.MediaType.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSMediatypeinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_MediaType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MediaType == nil {
		o.MediaType = &NTMSMediatypeinformation{}
	}
	if err := o.MediaType.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_LibRequest structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 10
type NTMSObjectinformationw_Info_LibRequest struct {
	LibRequest *NTMSLibrequestinformationw `idl:"name:LibRequest" json:"lib_request"`
}

func (*NTMSObjectinformationw_Info_LibRequest) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_LibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LibRequest != nil {
		if err := o.LibRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSLibrequestinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_LibRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LibRequest == nil {
		o.LibRequest = &NTMSLibrequestinformationw{}
	}
	if err := o.LibRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_OperationRequest structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 17
type NTMSObjectinformationw_Info_OperationRequest struct {
	OperationRequest *NTMSOprequestinformationw `idl:"name:OpRequest" json:"operation_request"`
}

func (*NTMSObjectinformationw_Info_OperationRequest) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_OperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.OperationRequest != nil {
		if err := o.OperationRequest.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSOprequestinformationw{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_OperationRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.OperationRequest == nil {
		o.OperationRequest = &NTMSOprequestinformationw{}
	}
	if err := o.OperationRequest.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectinformationw_Info_Computer structure represents NTMSObjectinformationw_Info RPC union arm.
//
// It has following labels: 4
type NTMSObjectinformationw_Info_Computer struct {
	Computer *NTMSComputerinformation `idl:"name:Computer" json:"computer"`
}

func (*NTMSObjectinformationw_Info_Computer) is_NTMSObjectinformationw_Info() {}

func (o *NTMSObjectinformationw_Info_Computer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Computer != nil {
		if err := o.Computer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&NTMSComputerinformation{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTMSObjectinformationw_Info_Computer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Computer == nil {
		o.Computer = &NTMSComputerinformation{}
	}
	if err := o.Computer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NTMSObjectManagement2 structure represents INtmsObjectManagement2 RPC structure.
type NTMSObjectManagement2 dcom.InterfacePointer

func (o *NTMSObjectManagement2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSObjectManagement2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSObjectManagement2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSObjectManagement2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSObjectManagement2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSObjectManagement3 structure represents INtmsObjectManagement3 RPC structure.
type NTMSObjectManagement3 dcom.InterfacePointer

func (o *NTMSObjectManagement3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSObjectManagement3) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSObjectManagement3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSObjectManagement3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSObjectManagement3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSLibraryControl1 structure represents INtmsLibraryControl1 RPC structure.
type NTMSLibraryControl1 dcom.InterfacePointer

func (o *NTMSLibraryControl1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSLibraryControl1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSLibraryControl1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSLibraryControl1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSLibraryControl1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSObjectManagement1 structure represents INtmsObjectManagement1 RPC structure.
type NTMSObjectManagement1 dcom.InterfacePointer

func (o *NTMSObjectManagement1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSObjectManagement1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSObjectManagement1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSObjectManagement1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSObjectManagement1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSMediaServices1 structure represents INtmsMediaServices1 RPC structure.
type NTMSMediaServices1 dcom.InterfacePointer

func (o *NTMSMediaServices1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSMediaServices1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSMediaServices1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSMediaServices1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSMediaServices1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSSession1 structure represents INtmsSession1 RPC structure.
type NTMSSession1 dcom.InterfacePointer

func (o *NTMSSession1) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *NTMSSession1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSSession1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSSession1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSSession1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSNotifySink structure represents INtmsNotifySink RPC structure.
type NTMSNotifySink dcom.InterfacePointer

func (o *NTMSNotifySink) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSNotifySink) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSNotifySink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSNotifySink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSNotifySink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSObjectInfo1 structure represents INtmsObjectInfo1 RPC structure.
type NTMSObjectInfo1 dcom.InterfacePointer

func (o *NTMSObjectInfo1) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSObjectInfo1) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSObjectInfo1) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSObjectInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSObjectInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// NTMSLibraryControl2 structure represents INtmsLibraryControl2 RPC structure.
type NTMSLibraryControl2 dcom.InterfacePointer

func (o *NTMSLibraryControl2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *NTMSLibraryControl2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *NTMSLibraryControl2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *NTMSLibraryControl2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *NTMSLibraryControl2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
