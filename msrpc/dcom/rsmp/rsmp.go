// The rsmp package implements the RSMP client protocol.
//
// # Introduction
//
// This document specifies the Removable Storage Manager (RSM) Remote Protocol.
//
// The RSM Remote Protocol is a set of distributed component object model (DCOM) [MS-DCOM]
// interfaces for applications to manage robotic changers, media libraries, and tape
// drives. The RSM Remote Protocol deals with detailed low-level operating system and
// storage concepts. Although the necessary basic concepts are outlined in this specification,
// this specification assumes reader familiarity with these technologies.
//
// # Overview
//
// The RSM Remote Protocol provides a mechanism for the remote configuration and management
// of removable storage devices such as robotic changers, media libraries, and tape
// drives. It allows multiple clients to manage removable media within a single-server
// system, and share local robotic media libraries, tape drives, and disk drives. The
// protocol also enables clients to obtain notifications of changes to these storage
// objects.
//
// Two entities are involved in the RSM Remote Protocol: the server, whose storage is
// configured, and the client, which accesses and requests changes to the server's storage
// configuration.
//
// The RSM Remote Protocol is expressed as a set of DCOM interfaces [MS-DCOM].
//
// The client end of the protocol invokes method calls on the interface to perform various
// tasks with the removable storage on the server. The client also implements some DCOM
// interfaces to get notifications for changes in the removable storage.
//
// The server end of the protocol implements DCOM interfaces to provide the following
// functions:<1>
//
// * *Session Management*
//
// This interface is used to open and close sessions. Establishing a session is a prerequisite
// to using the other functions of the RSM Remote Protocol.
//
// * *Media Library Management*
//
// The Media Library Management interface provides functions that:
//
// * Eject ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_a22dff61-81ae-4414-8242-3d28f4c70c31
// ) or inject ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_5e9a10ec-4dc8-457c-bac6-4df3af609f40
// ) media from a library.
//
// * Reserve or release a slot ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_a1f0ab90-f3c0-46b8-9c77-d0d7508ede94
// ) for cleaning.
//
// * Clean the drive.
//
// * Eject or inject a cleaner ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_95e53eb6-39e3-43ed-905e-ff1f427446d2
// ).
//
// * *Object Management*
//
// During the initialization process, the server performs an inventory ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_38cb8216-53ff-4936-a5c4-4d8d7672950a
// ) of media libraries, tape drives, robotic changers, and so on. The object management
// functions allow a client to create, delete, modify, or enumerate these objects. The
// server also maintains a record of all configured objects in the RSM database, which
// can be used across sessions.
//
// * *Media Management*
//
// The media management functions enable a client to perform any of the following functions:
//
// * Create or delete a media pool.
//
// * Mount ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_cc98b88f-7fd4-40f6-b315-212dc4d8378b
// ) or dismount ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_a8ebf8eb-a5f7-406a-b211-9f9c08ee15d3
// ) media.
//
// * Allocate ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_0e546e86-71c4-45cc-9ef1-4a78f6cb425a
// ) , deallocate ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_a8986e06-b078-4c42-86bc-178390427f04
// ) , or decommission ( 8d988f06-72be-47cf-9e09-9060aa8ed897#gt_d753ed1c-ff6a-4179-81ea-10ba369e1cca
// ) media.
package rsmp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
//
// The NTMS_LIBRARYINFORMATION structure defines properties specific to a library object.
type LibraryInformation struct {
	// LibraryType:  The library type object. This MUST be one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYTYPE_UNKNOWN 0x00000000    | The library type cannot be determined.                                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYTYPE_OFFLINE 0x00000001    | The library is not accessible.                                                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYTYPE_ONLINE 0x00000002     | A robotic element that automates the mounting and dismounting of media into one  |
	//	|                                        | or more drives.                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYTYPE_STANDALONE 0x00000003 | A stand-alone drive that is modeled as a library with one drive in RSM.          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	LibraryType uint32 `idl:"name:LibraryType" json:"library_type"`
	// CleanerSlot:  Specifies, for each library, the slot that was assigned to the cleaner
	// cartridge. If no cleaner slot is defined for this library, this member MUST be NULL.
	CleanerSlot *GUID `idl:"name:CleanerSlot" json:"cleaner_slot"`
	// CleanerSlotDefault:  Specifies a library's default or preferred cleaner slot. If
	// there is no preferred slot, this MUST be NULL.
	CleanerSlotDefault *GUID `idl:"name:CleanerSlotDefault" json:"cleaner_slot_default"`
	// LibrarySupportsDriveCleaning:  Used by drives requiring cleaning under automated
	// control. If TRUE, automatic drive cleaning operations are enabled; otherwise, cleaning
	// operations are not enabled.
	LibrarySupportsDriveCleaning bool `idl:"name:LibrarySupportsDriveCleaning" json:"library_supports_drive_cleaning"`
	// BarCodeReaderInstalled:  This MUST return TRUE if a bar code reader is installed
	// in a library; otherwise, it MUST return FALSE.
	BarCodeReaderInstalled bool `idl:"name:BarCodeReaderInstalled" json:"bar_code_reader_installed"`
	// InventoryMethod:  A default or user-selected method for performing an inventory of
	// this library. This MUST be one of the following values.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|             VALUE              |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_INVENTORY_NONE 0x00000000 | An inventory MUST NOT be performed after the library door is closed. An          |
	//	|                                | inventory might be required if a mount label check fails.                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_INVENTORY_FAST 0x00000001 | If the library has a bar code reader installed, a bar code inventory MUST        |
	//	|                                | be performed. If the library does not have a bar code reader, a differential     |
	//	|                                | inventory MUST be performed (slots that transitioned from empty to full are      |
	//	|                                | added).                                                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_INVENTORY_OMID 0x00000002 | A full inventory MUST be performed. A full inventory involves mounting each side |
	//	|                                | in a library and reading the on-media identification from the media.             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	InventoryMethod uint32 `idl:"name:InventoryMethod" json:"inventory_method"`
	// dwCleanerUsesRemaining:  The number of uses remaining on the cleaner in the library.
	// This member MUST be 0 if no cleaner is present, or if the library does not support
	// cleaning.
	CleanerUsesRemaining uint32 `idl:"name:dwCleanerUsesRemaining" json:"cleaner_uses_remaining"`
	// FirstDriveNumber:  The number of the first drive in the library.
	FirstDriveNumber uint32 `idl:"name:FirstDriveNumber" json:"first_drive_number"`
	// dwNumberOfDrives:  The number of drives in the library.
	NumberOfDrives uint32 `idl:"name:dwNumberOfDrives" json:"number_of_drives"`
	// FirstSlotNumber:  The number of the first slot in the library.
	FirstSlotNumber uint32 `idl:"name:FirstSlotNumber" json:"first_slot_number"`
	// dwNumberOfSlots:  The number of slots in the library.
	NumberOfSlots uint32 `idl:"name:dwNumberOfSlots" json:"number_of_slots"`
	// FirstDoorNumber:  The number of the first access door in the library.
	FirstDoorNumber uint32 `idl:"name:FirstDoorNumber" json:"first_door_number"`
	// dwNumberOfDoors:  The number of access doors in the library.
	NumberOfDoors uint32 `idl:"name:dwNumberOfDoors" json:"number_of_doors"`
	// FirstPortNumber:  The number of the first IE port in the library.
	FirstPortNumber uint32 `idl:"name:FirstPortNumber" json:"first_port_number"`
	// dwNumberOfPorts:  The number of IE ports in the library.
	NumberOfPorts uint32 `idl:"name:dwNumberOfPorts" json:"number_of_ports"`
	// FirstChangerNumber:  The number of the first changer in the library.
	FirstChangerNumber uint32 `idl:"name:FirstChangerNumber" json:"first_changer_number"`
	// dwNumberOfChangers:  The number of changers in the library.
	NumberOfChangers uint32 `idl:"name:dwNumberOfChangers" json:"number_of_changers"`
	// dwNumberOfMedia:  The number of media in the online or offline library.
	NumberOfMedia uint32 `idl:"name:dwNumberOfMedia" json:"number_of_media"`
	// dwNumberOfMediaTypes:  The number of media types that the library supports.
	NumberOfMediaTypes uint32 `idl:"name:dwNumberOfMediaTypes" json:"number_of_media_types"`
	// dwNumberOfLibRequests:  The number of current library requests.
	NumberOfLibRequests uint32 `idl:"name:dwNumberOfLibRequests" json:"number_of_lib_requests"`
	// Reserved:  This MUST be 0 and MUST be ignored on receipt.
	_ *dtyp.GUID `idl:"name:Reserved"`
	// AutoRecovery:  If the mount operation fails and this member is TRUE, a full inventory
	// MUST be performed. If this member is FALSE, a full inventory MUST NOT be performed.
	// The failure can be either a hardware or a label mismatch. For ATAPI CD libraries,
	// this member MUST NOT be set to FALSE. The default value is TRUE.
	AutoRecovery bool `idl:"name:AutoRecovery" json:"auto_recovery"`
	// dwFlags:  This member MUST be one or more of the following values.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                  |                                                                                  |
	//	|                      VALUE                       |                                     MEANING                                      |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYFLAG_FIXEDOFFLINE 0x01               | The library is an offline library, not a library that is not present.            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYFLAG_CLEANERPRESENT 0x02             | A cleaner is present in the changer.                                             |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYFLAG_IGNORECLEANERUSESREMAINING 0x08 | The cleaner cartridge MUST be used until it no longer cleans the drive, instead  |
	//	|                                                  | of keeping track of the number of cleanings left. This flag MUST NOT be set      |
	//	|                                                  | by the client. The server MUST set the flag if dwCleanerUsesRemaining is         |
	//	|                                                  | 0xFFFFFFFF, and the server MUST clear the flag otherwise.                        |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_LIBRARYFLAG_RECOGNIZECLEANERBARCODE 0x10    | Bar-coded cartridges that have CLN as a prefix MUST be treated as cleaner        |
	//	|                                                  | cartridges, instead of mounting them in the drive to identify them.              |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The NTMS_LIBRARYINFORMATION structure defines properties specific to a library object.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
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
//
// The SECURITY_ATTRIBUTES_NTMS structure contains the security descriptor for an object.
type SecurityAttributesNTMS struct {
	// nLength:  The size, in bytes, of the particular instance of the structure containing
	// this field.
	Length uint32 `idl:"name:nLength" json:"length"`
	// lpSecurityDescriptor:  A pointer to a security descriptor for the object that controls
	// the sharing of that object. Security descriptors are specified in [MS-DTYP].
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(nDescriptorLength)" json:"security_descriptor"`
	// bInheritHandle:  If set to TRUE, the new process MUST inherit the handle; if set
	// to FALSE, the handle MUST NOT be inherited.
	InheritHandle bool `idl:"name:bInheritHandle" json:"inherit_handle"`
	// nDescriptorLength:  The size, in bytes, of the descriptor.
	DescriptorLength uint32 `idl:"name:nDescriptorLength" json:"descriptor_length"`
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
//
// The NTMS_ALLOCATION_INFORMATION structure contains information about the source media
// pool from which a medium was taken.
type AllocationInformation struct {
	// dwSize:  The size, in bytes, of the structure.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// lpReserved:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	_ uint8 `idl:"name:lpReserved"`
	// AllocatedFrom:  The GUID of the media source (that is, an import pool or any other
	// user-defined pool).
	AllocatedFrom *GUID `idl:"name:AllocatedFrom" json:"allocated_from"`
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
//
// The NTMS_ASYNC_IO structure defines the state of an asynchronous request.
type AsyncIO struct {
	// OperationId:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	OperationID *GUID `idl:"name:OperationId" json:"operation_id"`
	// EventId:  The NTMS_GUID which is used by the server to notify the client using the
	// INtmsNotifySink::OnNotify (section 3.1.5.2.2.2) method.
	EventID *GUID `idl:"name:EventId" json:"event_id"`
	// dwOperationType:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	OperationType uint32 `idl:"name:dwOperationType" json:"operation_type"`
	// dwResult:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	Result uint32 `idl:"name:dwResult" json:"result"`
	// dwAsyncState:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	AsyncState uint32 `idl:"name:dwAsyncState" json:"async_state"`
	// hEvent:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	Event uint64 `idl:"name:hEvent" json:"event"`
	// bOnStateChange:  Indicates whether or not to signal on every status change. FALSE
	// means to signal only upon completion of the request.
	OnStateChange bool `idl:"name:bOnStateChange" json:"on_state_change"`
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
//
// The NTMS_MOUNT_INFORMATION structure defines mount information for the management
// of removable storage libraries.
type MountInformation struct {
	// dwSize:  The size, in bytes, of the structure.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// lpReserved:  Unused. This value MUST be NULL and MUST be ignored on receipt.
	_ *AsyncIO `idl:"name:lpReserved;pointer:ptr"`
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
//
// The NTMS_CHANGERINFORMATIONA structure describes the properties of a changer object
// as a sequence of ASCII characters.
type ChangerInformationA struct {
	// Number:  The number of the changer within the online library.
	Number uint32 `idl:"name:Number" json:"number"`
	// ChangerType:  An identifier of the type object for the changer.
	ChangerType *GUID `idl:"name:ChangerType" json:"changer_type"`
	// szSerialNumber:   A serial number for the changer in a null-terminated ASCII-character
	// string. Devices that do not support serial numbers MUST report NULL for this member.
	SerialNumber []byte `idl:"name:szSerialNumber" json:"serial_number"`
	// szRevision:   A null-terminated sequence of ASCII characters specifying the revision
	// of the changer.
	Revision []byte `idl:"name:szRevision" json:"revision"`
	// szDeviceName:   A null-terminated sequence of ASCII characters specifying the name
	// of the device used to access the changer.
	DeviceName []byte `idl:"name:szDeviceName" json:"device_name"`
	// ScsiPort:  The small computer system interface (SCSI) [ANSI-131-1994] host adapter
	// to which the changer is connected.
	SCSIPort uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	// ScsiBus:  The SCSI bus to which the changer is connected.
	SCSIBus uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	// ScsiTarget:  The SCSI target identifier of the changer.
	SCSITarget uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	// ScsiLun:  The SCSI logical unit identifier of the changer.
	SCSILUN uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	// Library:  The identifier of the library that contains the changer.
	Library *GUID `idl:"name:Library" json:"library"`
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
//
// The NTMS_CHANGERINFORMATIONW structure describes the properties of a changer object
// as a sequence of Unicode [UNICODE] characters.
type ChangerInformationW struct {
	// Number:  The number of the changer within the online library.
	Number uint32 `idl:"name:Number" json:"number"`
	// ChangerType:  The identifier of the type object for the changer.
	ChangerType *GUID `idl:"name:ChangerType" json:"changer_type"`
	// szSerialNumber:   The serial number for the changer in a null-terminated string.
	// Devices that do not support serial numbers MUST report NULL for this member.
	SerialNumber string `idl:"name:szSerialNumber;string" json:"serial_number"`
	// szRevision:   A null-terminated sequence of Unicode characters specifying the revision
	// of the changer.
	Revision string `idl:"name:szRevision;string" json:"revision"`
	// szDeviceName:   A null-terminated sequence of Unicode characters specifying the name
	// of the device used to access the changer.
	DeviceName string `idl:"name:szDeviceName;string" json:"device_name"`
	// ScsiPort:  The SCSI [ANSI-131-1994] host adapter to which the changer is connected.
	SCSIPort uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	// ScsiBus:  The SCSI bus to which the changer is connected.
	SCSIBus uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	// ScsiTarget:  The SCSI target identifier of the changer.
	SCSITarget uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	// ScsiLun:  The SCSI logical unit identifier of the changer.
	SCSILUN uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	// Library:  The identifier of the library that contains the changer.
	Library *GUID `idl:"name:Library" json:"library"`
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
//
// The NTMS_CHANGERTYPEINFORMATIONA structure describes the properties specific to a
// type of changer, as a sequence of ASCII characters.
type ChangerTypeInformationA struct {
	// szVendor:   A null-terminated sequence of ASCII characters specifying the name of
	// the changer vendor, acquired from device inquiry data. If no name is available, this
	// MUST contain an empty string.
	Vendor []byte `idl:"name:szVendor" json:"vendor"`
	// szProduct:   A null-terminated sequence of ASCII characters specifying the name of
	// the changer product, acquired through SCSI commands. If no name is available, this
	// MUST contain an empty string.
	Product []byte `idl:"name:szProduct" json:"product"`
	// DeviceType:  The following SCSI device type [ANSI-131-1994] acquired from device
	// inquiry data.
	//
	//	+--------------------------------+----------------------+
	//	|                                |                      |
	//	|             VALUE              |       MEANING        |
	//	|                                |                      |
	//	+--------------------------------+----------------------+
	//	+--------------------------------+----------------------+
	//	| FILE_DEVICE_CHANGER 0x00000030 | Device is a changer. |
	//	+--------------------------------+----------------------+
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
//
// The NTMS_CHANGERTYPEINFORMATIONW structure describes the properties specific to a
// type of changer, in Unicode.
type ChangerTypeInformationW struct {
	// szVendor:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the changer vendor, acquired from device inquiry data. If no name is available,
	// this MUST contain an empty string.
	Vendor string `idl:"name:szVendor;string" json:"vendor"`
	// szProduct:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the changer product, acquired from device inquiry data. If no name is available,
	// this MUST contain an empty string.
	Product string `idl:"name:szProduct;string" json:"product"`
	// DeviceType:  The following SCSI device type [ANSI-131-1994] acquired through SCSI
	// commands.
	//
	//	+--------------------------------+----------------------+
	//	|                                |                      |
	//	|             VALUE              |       MEANING        |
	//	|                                |                      |
	//	+--------------------------------+----------------------+
	//	+--------------------------------+----------------------+
	//	| FILE_DEVICE_CHANGER 0x00000030 | Device is a changer. |
	//	+--------------------------------+----------------------+
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
//
// The NTMS_DRIVEINFORMATIONA structure describes the properties of a drive object,
// as a sequence of ASCII characters.
type DriveInformationA struct {
	// Number:  The number of the drive in the library. Some changers assign the number
	// 0 to the first of their drives; other changers number the first drive 1.
	Number uint32 `idl:"name:Number" json:"number"`
	// State:  A value from the NtmsDriveState enumeration specifying the state of the drive.
	State uint32 `idl:"name:State" json:"state"`
	// DriveType:  The identifier of the type object for the drive.
	DriveType *GUID `idl:"name:DriveType" json:"drive_type"`
	// szDeviceName:   A null-terminated sequence of ASCII characters specifying the name
	// of the device path to access the drive.
	DeviceName []byte `idl:"name:szDeviceName" json:"device_name"`
	// szSerialNumber:   The null-terminated serial number of the drive.
	SerialNumber []byte `idl:"name:szSerialNumber" json:"serial_number"`
	// szRevision:   A null-terminated sequence of ASCII characters specifying the revision
	// of the drive.
	Revision []byte `idl:"name:szRevision" json:"revision"`
	// ScsiPort:  The SCSI [ANSI-131-1994] host adapter to which the drive is connected.
	SCSIPort uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	// ScsiBus:  The SCSI bus to which the drive is connected.
	SCSIBus uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	// ScsiTarget:  The SCSI target identifier of the drive.
	SCSITarget uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	// ScsiLun:  The SCSI logical unit identifier of the drive.
	SCSILUN uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	// dwMountCount:  The number of times the drive has had media mounted in it. If the
	// drive supports the reporting of a unique serial number, this value MUST be the number
	// of times the drive has been mounted since RSM began managing this drive. If the drive
	// does not support unique serial numbers, this member MUST indicate the number of mounts
	// to all the drives at that mount location.
	MountCount uint32 `idl:"name:dwMountCount" json:"mount_count"`
	// LastCleanedTs:  A SYSTEMTIME structure that specifies the last time the drive was
	// cleaned.
	LastCleanedTS *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	// SavedPartitionId:  The partition identifier of the medium in the drive. If this value
	// is NULL and the drive is marked as full, the medium was loaded by a user, and it
	// MUST be identified and given a partition identifier.
	SavedPartitionID *GUID `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	// Library:  The identifier of the library that contains the drive.
	Library *GUID `idl:"name:Library" json:"library"`
	// Reserved:  This parameter is unused. It MUST be 0 and MUST be ignored on receipt.
	_ *dtyp.GUID `idl:"name:Reserved"`
	// dwDeferDismountDelay:  Minimum number of seconds that media will remain in the drive
	// of an online library after a deferred dismount is performed; the default is 5 minutes.
	// This member does not apply to stand-alone libraries.
	DeferDismountDelay uint32 `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
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
//
// The NTMS_DRIVEINFORMATIONW structure describes the properties of a drive object,
// as a sequence of Unicode characters.
type DriveInformationW struct {
	// Number:  The number of the drive in the library. Some changers assign the number
	// 0 to the first drive; other changers number the first drive 1.
	Number uint32 `idl:"name:Number" json:"number"`
	// State:  A value from the NtmsDriveState enumeration that specifies the state of the
	// drive.
	State uint32 `idl:"name:State" json:"state"`
	// DriveType:  The identifier of the type object for the drive.
	DriveType *GUID `idl:"name:DriveType" json:"drive_type"`
	// szDeviceName:   A null-terminated sequence of Unicode UTF-16 characters specifying
	// the name of the device path to access the drive.
	DeviceName string `idl:"name:szDeviceName;string" json:"device_name"`
	// szSerialNumber:   The null-terminated serial number of the drive.
	SerialNumber string `idl:"name:szSerialNumber;string" json:"serial_number"`
	// szRevision:   A null-terminated sequence of Unicode UTF-16 characters specifying
	// the revision of the drive.
	Revision string `idl:"name:szRevision;string" json:"revision"`
	// ScsiPort:  The SCSI [ANSI-131-1994] host adapter to which the drive is connected.
	SCSIPort uint16 `idl:"name:ScsiPort" json:"scsi_port"`
	// ScsiBus:  The SCSI bus to which the drive is connected.
	SCSIBus uint16 `idl:"name:ScsiBus" json:"scsi_bus"`
	// ScsiTarget:  The SCSI target identifier of the drive.
	SCSITarget uint16 `idl:"name:ScsiTarget" json:"scsi_target"`
	// ScsiLun:  The SCSI logical unit identifier of the drive.
	SCSILUN uint16 `idl:"name:ScsiLun" json:"scsi_lun"`
	// dwMountCount:  The number of times the drive has had media mounted in it. If the
	// drive supports the reporting of a unique serial number, this value MUST be the number
	// of times the drive has been mounted since the RSM began managing this drive. If the
	// drive does not support unique serial numbers, this member MUST indicate the number
	// of mounts to all the drives at that mount location.
	MountCount uint32 `idl:"name:dwMountCount" json:"mount_count"`
	// LastCleanedTs:  A SYSTEMTIME structure specifying the last time the drive was cleaned.
	LastCleanedTS *dtyp.SystemTime `idl:"name:LastCleanedTs" json:"last_cleaned_ts"`
	// SavedPartitionId:  The partition identifier of the media in the drive. If this value
	// is NULL and the drive is marked as full, the medium was loaded by a user, and MUST
	// be identified and given a partition identifier.
	SavedPartitionID *GUID `idl:"name:SavedPartitionId" json:"saved_partition_id"`
	// Library:  The identifier of the library that contains the drive.
	Library *GUID `idl:"name:Library" json:"library"`
	// Reserved:  This parameter is unused. It MUST be 0 and MUST be ignored on receipt.
	_ *dtyp.GUID `idl:"name:Reserved"`
	// dwDeferDismountDelay:  The minimum number of seconds that media MUST remain in the
	// drive of an online library after a deferred dismount is performed; the default MUST
	// be five minutes. This member MUST NOT apply to stand-alone libraries.
	DeferDismountDelay uint32 `idl:"name:dwDeferDismountDelay" json:"defer_dismount_delay"`
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
//
// The NTMS_DRIVETYPEINFORMATIONA structure describes the properties specific to a type
// of drive, in ASCII.
type DriveTypeInformationA struct {
	// szVendor:   A null-terminated sequence of ASCII characters specifying the name of
	// the vendor of the drive, acquired from device inquiry data. If this information is
	// not available, the member MUST specify an empty string.
	Vendor []byte `idl:"name:szVendor" json:"vendor"`
	// szProduct:   A null-terminated sequence of ASCII characters specifying the name of
	// the product of the drive, acquired from device inquiry data. If this information
	// is not available, the member MUST specify an empty string.
	Product []byte `idl:"name:szProduct" json:"product"`
	// NumberOfHeads:  This parameter is currently unused. It MUST be NULL and MUST be ignored
	// on  receipt.
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	// DeviceType:  One of the following SCSI device types [ANSI-131-1994] acquired from
	// device inquiry data.
	//
	//	+-------------------------------+--------------------------------------+
	//	|                               |                                      |
	//	|             VALUE             |               MEANING                |
	//	|                               |                                      |
	//	+-------------------------------+--------------------------------------+
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_CD_ROM 0x00000002 | Device is a CD-ROM.                  |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_DISK 0x00000007   | Device is a direct-access drive.     |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_DVD 0x00000033    | Device is a DVD.                     |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_TAPE 0x0000001F   | Device is a sequential-access drive. |
	//	+-------------------------------+--------------------------------------+
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
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
//
// The NTMS_DRIVETYPEINFORMATIONW structure describes the properties specific to a type
// of drive, in Unicode.
type DriveTypeInformationW struct {
	// szVendor:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the vendor of the drive, acquired from device inquiry data. If this information
	// is not available, the member MUST specify an empty string.
	Vendor string `idl:"name:szVendor;string" json:"vendor"`
	// szProduct:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the product of the drive, acquired from device inquiry data. If this information
	// is not available, the member MUST specify an empty string.
	Product string `idl:"name:szProduct;string" json:"product"`
	// NumberOfHeads:  This parameter is currently unused. It MUST be zero and MUST be ignored
	// on receipt.
	NumberOfHeads uint32 `idl:"name:NumberOfHeads" json:"number_of_heads"`
	// DeviceType:  One of the following SCSI device types [ANSI-131-1994] acquired from
	// device inquiry data.
	//
	//	+-------------------------------+--------------------------------------+
	//	|                               |                                      |
	//	|             VALUE             |               MEANING                |
	//	|                               |                                      |
	//	+-------------------------------+--------------------------------------+
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_CD_ROM 0x00000002 | Device is a CD-ROM.                  |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_DISK 0x00000007   | Device is a direct-access drive.     |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_DVD 0x00000033    | Device is a DVD.                     |
	//	+-------------------------------+--------------------------------------+
	//	| FILE_DEVICE_TAPE 0x0000001F   | Device is a sequential-access drive. |
	//	+-------------------------------+--------------------------------------+
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
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
//
// The NTMS_LIBREQUESTINFORMATIONA structure describes the properties of a work request,
// in ASCII.
type LibraryRequestInformationA struct {
	// OperationCode:  A value from the NtmsLmOperation enumeration specifying the type
	// of operation requested.
	OperationCode uint32 `idl:"name:OperationCode" json:"operation_code"`
	// OperationOption:  Options specific to a library request. The following table shows
	// the meanings if OperationCode is set to LM_MOUNT, LM_DISMOUNT, or LM_EJECT.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| IMMEDIATE 0x00000000      | The operation MUST be completed immediately.                                     |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| DEFERRED 0x00000001       | The operation MUST be completed only when the slot is later required for an      |
	//	|                           | operation.                                                                       |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| FORCEIMMEDIATE 0x00000002 | The operation MUST be completed immediately. The operation will complete even if |
	//	|                           | there are open handles to the medium.                                            |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| FORCEDEFERRED 0x00000003  | The operation MUST be completed only when the slot is later required for an      |
	//	|                           | operation. The operation MUST complete even if there are open handles to the     |
	//	|                           | medium.                                                                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| INJECTMANY 0x00000004     | The operation applies to multiple slots.                                         |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// The following table shows the meanings if OperationCode is set to LM_INVENTORY.
	//
	//	+---------------------------+-------------------------------------------------------+
	//	|                           |                                                       |
	//	|           VALUE           |                        MEANING                        |
	//	|                           |                                                       |
	//	+---------------------------+-------------------------------------------------------+
	//	+---------------------------+-------------------------------------------------------+
	//	| FULL_INVENTORY 0x00000000 | A complete inventory of media MUST be done.           |
	//	+---------------------------+-------------------------------------------------------+
	//	| SLOTS_ONLY 0x00000001     | Only the media loaded into slots MUST be inventoried. |
	//	+---------------------------+-------------------------------------------------------+
	OperationOption uint32 `idl:"name:OperationOption" json:"operation_option"`
	// State:  A value from the NtmsLmState (section 2.2.1.10) enumeration specifying the
	// state of the work request.
	State uint32 `idl:"name:State" json:"state"`
	// PartitionId:  The identifier of a side for which the request is submitted to the
	// server.
	PartitionID *GUID `idl:"name:PartitionId" json:"partition_id"`
	// DriveId:  The identifier of a drive that is being serviced.
	DriveID *GUID `idl:"name:DriveId" json:"drive_id"`
	// PhysMediaId:  The identifier of a piece of physical media that is being serviced.
	PhysicalMediaID *GUID `idl:"name:PhysMediaId" json:"physical_media_id"`
	// Library:  The identifier of the library for the request.
	Library *GUID `idl:"name:Library" json:"library"`
	// SlotId:  The identifier of the slot of the piece of physical media that is being
	// serviced.
	SlotID *GUID `idl:"name:SlotId" json:"slot_id"`
	// TimeQueued:  A SYSTEMTIME structure specifying the time at which the request was
	// submitted to the server.
	TimeQueued *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	// TimeCompleted:  A SYSTEMTIME structure specifying the time at which the request was
	// completed.
	TimeCompleted *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	// szApplication:   A null-terminated sequence of ASCII characters specifying the name
	// of the application that submitted the operator request.<7>
	Application []byte `idl:"name:szApplication" json:"application"`
	// szUser:   A null-terminated sequence of ASCII characters specifying the name of the
	// interactive user who submitted the operator request.<8>
	User []byte `idl:"name:szUser" json:"user"`
	// szComputer:   A null-terminated sequence of ASCII characters specifying the name
	// of the computer that submitted the operator request.<9>
	Computer []byte `idl:"name:szComputer" json:"computer"`
	// dwErrorCode:  An implementation-specific nonzero error code that returns with State
	// set to the NTMS_LM_FAILED value.<10>
	ErrorCode uint32 `idl:"name:dwErrorCode" json:"error_code"`
	// WorkItemId:  The associated identifier for the request, which was assigned by a server
	// when it received a request from a client to perform an operation on a library.
	WorkItemID *GUID `idl:"name:WorkItemId" json:"work_item_id"`
	// dwPriority:  The priority of the request.
	Priority uint32 `idl:"name:dwPriority" json:"priority"`
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
//
// The NTMS_LIBREQUESTINFORMATIONW structure describes the properties of a work request,
// in Unicode.
type LibraryRequestInformationW struct {
	// OperationCode:  A value from the NtmsLmOperation enumeration specifying the type
	// of operation requested.
	OperationCode uint32 `idl:"name:OperationCode" json:"operation_code"`
	// OperationOption:  Options specific to a library request. The following table shows
	// the meanings if OperationCode is set to LM_MOUNT, LM_DISMOUNT, or LM_EJECT.
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|                           |                                                                                  |
	//	|           VALUE           |                                     MEANING                                      |
	//	|                           |                                                                                  |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| IMMEDIATE 0x00000000      | The operation MUST be completed immediately.                                     |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| DEFERRED 0x00000001       | The operation MUST be completed only when the slot is later required for an      |
	//	|                           | operation.                                                                       |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| FORCEIMMEDIATE 0x00000002 | The operation MUST be completed immediately. The operation will complete even if |
	//	|                           | there are open handles to the medium.                                            |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| FORCEDEFERRED 0x00000003  | The operation MUST be completed only when the slot is later required for an      |
	//	|                           | operation. The operation MUST complete even if there are open handles to the     |
	//	|                           | medium.                                                                          |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| INJECTMANY 0x00000004     | The operation applies to multiple slots.                                         |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// The following table shows the meanings if OperationCode is set to LM_INVENTORY.
	//
	//	+---------------------------+-------------------------------------------------------+
	//	|                           |                                                       |
	//	|           VALUE           |                        MEANING                        |
	//	|                           |                                                       |
	//	+---------------------------+-------------------------------------------------------+
	//	+---------------------------+-------------------------------------------------------+
	//	| FULL_INVENTORY 0x00000000 | A complete inventory of media MUST be done.           |
	//	+---------------------------+-------------------------------------------------------+
	//	| SLOTS_ONLY 0x00000001     | Only the media loaded into slots MUST be inventoried. |
	//	+---------------------------+-------------------------------------------------------+
	OperationOption uint32 `idl:"name:OperationOption" json:"operation_option"`
	// State:  A value from the NtmsLmState (section 2.2.1.10) enumeration specifying the
	// state of the work request.
	State uint32 `idl:"name:State" json:"state"`
	// PartitionId:  The identifier of a side for which the request is submitted to the
	// server.
	PartitionID *GUID `idl:"name:PartitionId" json:"partition_id"`
	// DriveId:  The identifier of a drive that is being serviced.
	DriveID *GUID `idl:"name:DriveId" json:"drive_id"`
	// PhysMediaId:  The identifier of a piece of physical media that is being serviced.
	PhysicalMediaID *GUID `idl:"name:PhysMediaId" json:"physical_media_id"`
	// Library:  The identifier of the library for the request.
	Library *GUID `idl:"name:Library" json:"library"`
	// SlotId:  The identifier of the slot of the piece of physical media that is being
	// serviced.
	SlotID *GUID `idl:"name:SlotId" json:"slot_id"`
	// TimeQueued:  A SYSTEMTIME structure specifying the time at which the request was
	// queued.
	TimeQueued *dtyp.SystemTime `idl:"name:TimeQueued" json:"time_queued"`
	// TimeCompleted:  A SYSTEMTIME structure specifying the time at which the request was
	// completed.
	TimeCompleted *dtyp.SystemTime `idl:"name:TimeCompleted" json:"time_completed"`
	// szApplication:   A null-terminated sequence of Unicode UTF-16 characters specifying
	// the name of the application that submitted the operator request.<11>
	Application string `idl:"name:szApplication;string" json:"application"`
	// szUser:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the interactive user who submitted the operator request.
	User string `idl:"name:szUser;string" json:"user"`
	// szComputer:   A null-terminated sequence of Unicode UTF-16 characters specifying
	// the name of the computer that submitted the operator request.
	Computer string `idl:"name:szComputer;string" json:"computer"`
	// dwErrorCode:  An implementation-specific nonzero error code for requests that return
	// with State set to the NTMS_LM_FAILED value.<12>
	ErrorCode uint32 `idl:"name:dwErrorCode" json:"error_code"`
	// WorkItemId:  The associated identifier for the request, which is assigned by a server
	// when it receives a request from a client to perform an operation on a library.
	WorkItemID *GUID `idl:"name:WorkItemId" json:"work_item_id"`
	// dwPriority:  The priority of the request.
	Priority uint32 `idl:"name:dwPriority" json:"priority"`
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
//
// The NTMS_MEDIAPOOLINFORMATION structure defines the properties specific to a media
// pool object.
type MediaPoolInformation struct {
	// PoolType:  An NTMS-supported media pool type.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLTYPE_UNKNOWN 0x00000000    | Unknown pool type.                                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLTYPE_SCRATCH 0x00000001    | Media that are available to other applications.                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLTYPE_FOREIGN 0x00000002    | Media that have been written to and that do not contain a recognizable on-media  |
	//	|                                     | identifier label type or label ID.                                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLTYPE_IMPORT 0x00000003     | Media that have been written to and that have a recognizable on-media identifier |
	//	|                                     | label type but an unrecognizable label ID.                                       |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLTYPE_APPLICATION 0x00003E8 | A media pool that is created by an application. One or more application media    |
	//	|                                     | pools can be created per system.                                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	PoolType uint32 `idl:"name:PoolType" json:"pool_type"`
	// MediaType:  A single media type that makes up each media pool.
	MediaType *GUID `idl:"name:MediaType" json:"media_type"`
	// Parent:  A parent media pool or NULL.
	Parent *GUID `idl:"name:Parent" json:"parent"`
	// AllocationPolicy:  A bitfield that specifies the action at allocation time. This
	// MUST be the following value, or left as 0.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_ALLOCATE_FROMSCRATCH 0x00000001 | Draw media from the free pool if none is available in the pool. The default is   |
	//	|                                      | not to draw from free pool. Return media to free when available. The default is  |
	//	|                                      | not to return to free.                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	AllocationPolicy uint32 `idl:"name:AllocationPolicy" json:"allocation_policy"`
	// DeallocationPolicy:  A bitfield that specifies action at deallocation time. This
	// member can be the following value or left as 0.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------+
	//	|                                      |                                                                            |
	//	|                VALUE                 |                                  MEANING                                   |
	//	|                                      |                                                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------+
	//	| NTMS_DEALLOCATE_TOSCRATCH 0x00000001 | Return media to free when available. The default is not to return to free. |
	//	+--------------------------------------+----------------------------------------------------------------------------+
	DeallocationPolicy uint32 `idl:"name:DeallocationPolicy" json:"deallocation_policy"`
	// dwMaxAllocates:  The maximum number of times the medium can be allocated and deallocated.
	MaxAllocates uint32 `idl:"name:dwMaxAllocates" json:"max_allocates"`
	// dwNumberOfPhysicalMedia:  The number of physical media in this media pool.
	NumberOfPhysicalMedia uint32 `idl:"name:dwNumberOfPhysicalMedia" json:"number_of_physical_media"`
	// dwNumberOfLogicalMedia:  The number of logical media in this media pool.
	NumberOfLogicalMedia uint32 `idl:"name:dwNumberOfLogicalMedia" json:"number_of_logical_media"`
	// dwNumberOfMediaPools:  The number of media pools in this media pool.
	NumberOfMediaPools uint32 `idl:"name:dwNumberOfMediaPools" json:"number_of_media_pools"`
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
//
// The NTMS_MEDIATYPEINFORMATION structure defines the properties specific to a type
// of media supported by RSM.
type MediaTypeInformation struct {
	// MediaType:  Each disk or tape driver reports the media type enumeration value of
	// the medium that is currently mounted in the drive. This media type value MUST be
	// unique, and is mapped to a human-readable string in the object szName member of a
	// NTMS_OBJECTINFORMATIONA or NTMS_OBJECTINFORMATIONW structure.
	//
	// MediaType MUST be one of the following values.
	//
	//	+-------------------------------+----------------------------------------------------------+
	//	|                               |                                                          |
	//	|             VALUE             |                         MEANING                          |
	//	|                               |                                                          |
	//	+-------------------------------+----------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------+
	//	| DDS_4mm 0x00000020            | Tape - DAT, DDS1, DDS2, and so on (all vendors)          |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MiniQic 0x00000021            | Tape - miniQIC tape                                      |
	//	+-------------------------------+----------------------------------------------------------+
	//	| Travan 0x00000022             | Tape - Travan tape (TR-1, TR-2, TR-3, and so on)         |
	//	+-------------------------------+----------------------------------------------------------+
	//	| QIC 0x00000023                | Tape - QIC tape                                          |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MP_8mm 0x00000024             | Tape - 8 mm Exabyte metal particle tape                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| AME_8mm 0x00000025            | Tape - 8 mm Exabyte advanced metal evaporative tape      |
	//	+-------------------------------+----------------------------------------------------------+
	//	| AIT1_8mm 0x00000026           | Tape - 8 mm Sony AIT                                     |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DLT 0x00000027                | Tape - DLT compact tape (IIIxt or IV)                    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| NCTP 0x00000028               | Tape - Philips NCTP tape                                 |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IBM_3480 0x00000029           | Tape - IBM 3480 tape                                     |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IBM_3490E 0x0000002a          | Tape - IBM 3490E tape                                    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IBM_Magstar_3590 0x0000002b   | Tape - IBM Magstar 3590 tape                             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IBM_Magstar_MP 0x0000002c     | Tape - IBM Magstar MP tape                               |
	//	+-------------------------------+----------------------------------------------------------+
	//	| STK_DATA_D3 0x0000002d        | Tape - STK Data D3 tape                                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SONY_DTF 0x0000002e           | Tape - Sony DTF tape                                     |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DV_6mm 0x0000002f             | Tape - 6 mm digital video tape                           |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DMI 0x00000030                | Tape - Exabyte DMI tape and compatibles                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SONY_D2 0x00000031            | Tape - Sony D2S and D2L tape                             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| CLEANER_CARTRIDGE 0x00000032  | Cleaner (all drive types that support drive cleaners)    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| CD_ROM 0x00000033             | Optical disk - CD                                        |
	//	+-------------------------------+----------------------------------------------------------+
	//	| CD_R 0x00000034               | Optical disk - CD-Recordable (write once)                |
	//	+-------------------------------+----------------------------------------------------------+
	//	| CD_RW 0x00000035              | Optical disk - CD-Rewritable                             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DVD_ROM 0x00000036            | Optical disk - DVD-ROM                                   |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DVD_R 0x00000037              | Optical disk - DVD-Recordable (write once)               |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DVD_RW 0x00000038             | Optical disk - DVD-Rewritable                            |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MO_3_RW 0x00000039            | Optical disk - 3.5 inch rewritable MO disk               |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MO_5_WO 0x0000003a            | Optical disk - MO 5.25 inch write once                   |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MO_5_RW 0x0000003b            | Optical disk - MO 5.25 inch rewritable (not LIMDOW)      |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MO_5_LIMDOW 0x0000003c        | Optical disk - MO 5.25 inch rewritable (LIMDOW)          |
	//	+-------------------------------+----------------------------------------------------------+
	//	| PC_5_WO 0x0000003d            | Optical disk - Phase change 5.25 inch write once optical |
	//	+-------------------------------+----------------------------------------------------------+
	//	| PC_5_RW 0x0000003e            | Optical disk - Phase change 5.25 inch rewritable         |
	//	+-------------------------------+----------------------------------------------------------+
	//	| PD_5_RW 0x0000003f            | Optical disk - Phase change dual rewritable              |
	//	+-------------------------------+----------------------------------------------------------+
	//	| ABL_5_WO 0x00000040           | Optical disk - Ablative 5.25 inch write once optical     |
	//	+-------------------------------+----------------------------------------------------------+
	//	| PINNACLE_APEX_5_RW 0x00000041 | Optical disk - Pinnacle Apex 4.6GB rewritable optical    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SONY_12_WO 0x00000042         | Optical disk - Sony 12 inch write once                   |
	//	+-------------------------------+----------------------------------------------------------+
	//	| PHILIPS_12_WO 0x00000043      | Optical disk - Philips/LMS 12 inch write once            |
	//	+-------------------------------+----------------------------------------------------------+
	//	| HITACHI_12_WO 0x00000044      | Optical disk - Hitachi 12 inch write once                |
	//	+-------------------------------+----------------------------------------------------------+
	//	| CYGNET_12_WO 0x00000045       | Optical disk - Cygnet/ATG 12 inch write once             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| KODAK_14_WO 0x00000046        | Optical disk - Kodak 14 inch write once                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MO_NFR_525 0x00000047         | Optical disk - Near field recording (Terastor)           |
	//	+-------------------------------+----------------------------------------------------------+
	//	| NIKON_12_RW 0x00000048        | Optical disk - Nikon 12 inch rewritable                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IOMEGA_ZIP 0x00000049         | Magnetic disk - Iomega Zip                               |
	//	+-------------------------------+----------------------------------------------------------+
	//	| IOMEGA_JAZ 0x0000004a         | Magnetic disk - Iomega Jaz                               |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SYQUEST_EZ135 0x0000004b      | Magnetic disk - Syquest EZ135                            |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SYQUEST_EZFLYER 0x0000004c    | Magnetic disk - Syquest EzFlyer                          |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SYQUEST_SYJET 0x0000004d      | Magnetic disk - Syquest SyJet                            |
	//	+-------------------------------+----------------------------------------------------------+
	//	| AVATAR_F2 0x0000004e          | Magnetic disk - 2.5 inch floppy                          |
	//	+-------------------------------+----------------------------------------------------------+
	//	| MP2_8mm 0x0000004f            | Tape - 8 millimeter Hitachi tape                         |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DST_S 0x00000050              | Ampex DST small tapes                                    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DST_M 0x00000051              | Ampex DST medium tapes                                   |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DST_L 0x00000052              | Ampex DST large tapes                                    |
	//	+-------------------------------+----------------------------------------------------------+
	//	| VXATape_1 0x00000053          | Ecrix 8 millimeter tape                                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| VXATape_2 0x00000054          | Ecrix 8 millimeter tape                                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| STK_9840 0x00000055           | STK 9840                                                 |
	//	+-------------------------------+----------------------------------------------------------+
	//	| LTO_Ultrium 0x00000056        | IBM, HP, Seagate LTO Ultrium                             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| LTO_Accelis 0x00000057        | IBM, HP, Seagate LTO Accelis                             |
	//	+-------------------------------+----------------------------------------------------------+
	//	| DVD_RAM 0x00000058            | Optical disk - DVD-RAM                                   |
	//	+-------------------------------+----------------------------------------------------------+
	//	| AIT_8mm 0x00000059            | AIT2 or higher                                           |
	//	+-------------------------------+----------------------------------------------------------+
	//	| ADR_1 0x0000005a              | OnStream ADR Mediatypes                                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| ADR_2 0x0000005b              | OnStream ADR Mediatypes                                  |
	//	+-------------------------------+----------------------------------------------------------+
	//	| STK_9940 0x0000005c           | STK 9940                                                 |
	//	+-------------------------------+----------------------------------------------------------+
	//	| SAIT 0x0000005d               | SAIT tapes                                               |
	//	+-------------------------------+----------------------------------------------------------+
	MediaType uint32 `idl:"name:MediaType" json:"media_type"`
	// NumberOfSides:  The number of sides on the media.
	NumberOfSides uint32 `idl:"name:NumberOfSides" json:"number_of_sides"`
	// ReadWriteCharacteristics:  Identifies the read/write characteristics of the media
	// type. This MUST be one of the following values.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_MEDIARW_UNKNOWN 0x00000000    | Unknown media characteristics. This value can be used to initialize              |
	//	|                                    | ReadWriteCharacteristics before a final value is assigned.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_MEDIARW_REWRITABLE 0x00000001 | Media that can be written to more than once. This includes magnetic tape,        |
	//	|                                    | magnetic disk, and some optical disk media.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_MEDIARW_WRITEONCE 0x00000002  | Media that can be written to only one time. Some optical media (for example,     |
	//	|                                    | 5.25-inch, 12-inch, 14-inch WORM, and CD-R) are designed to be write-once.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_MEDIARW_READONLY 0x00000003   | Media that cannot be written to, such as a CD-ROM and a DVD-ROM.                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ReadWriteCharacteristics uint32 `idl:"name:ReadWriteCharacteristics" json:"read_write_characteristics"`
	// DeviceType:  The SCSI device type [ANSI-131-1994] as reported from device inquiry
	// data. This MUST be one of the following values.
	//
	//	+-------------------------------+---------------------------+
	//	|                               |                           |
	//	|             VALUE             |          MEANING          |
	//	|                               |                           |
	//	+-------------------------------+---------------------------+
	//	+-------------------------------+---------------------------+
	//	| FILE_DEVICE_CD_ROM 0x00000002 | CD-ROM device.            |
	//	+-------------------------------+---------------------------+
	//	| FILE_DEVICE_DISK 0x00000007   | Direct-access device.     |
	//	+-------------------------------+---------------------------+
	//	| FILE_DEVICE_TAPE 0x0000001F   | Sequential-access device. |
	//	+-------------------------------+---------------------------+
	DeviceType uint32 `idl:"name:DeviceType" json:"device_type"`
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
//
// The NTMS_STORAGESLOTINFORMATION structure defines properties specific to a storage
// slot object.
type StorageSlotInformation struct {
	// Number:  The number of the slot in the library.
	Number uint32 `idl:"name:Number" json:"number"`
	// State:  The current state of the slot. This MUST be one of the following values.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                          |                                                                                  |
	//	|                  VALUE                   |                                     MEANING                                      |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_SLOTSTATE_UNKNOWN 0x00000000        | The slot state cannot be determined.                                             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_SLOTSTATE_FULL 0x00000001           | The slot is present and contains physical media.                                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_SLOTSTATE_EMPTY 0x00000002          | The slot is present but does not contain physical media.                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_SLOTSTATE_NOTPRESENT 0x00000003     | The slot is not present. If the library contains magazines, this value is        |
	//	|                                          | reported for each slot when the associated magazine is missing.                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_SLOTSTATE_NEEDSINVENTORY 0x00000004 | The slot needs inventory.                                                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// Library:  The library that contains the slot.
	Library *GUID `idl:"name:Library" json:"library"`
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
//
// The NTMS_IEDOORINFORMATION structure defines properties specific to an insert/eject
// door object.
type IEDoorInformation struct {
	// Number:  The number of the doors in the library. Libraries typically have one door.
	Number uint32 `idl:"name:Number" json:"number"`
	// State:  The state of the door. This MUST be one of the following values.
	//
	//	+-----------------------------------+--------------------------------------+
	//	|                                   |                                      |
	//	|               VALUE               |               MEANING                |
	//	|                                   |                                      |
	//	+-----------------------------------+--------------------------------------+
	//	+-----------------------------------+--------------------------------------+
	//	| NTMS_DOORSTATE_CLOSED 0x00000001  | The library door is closed.          |
	//	+-----------------------------------+--------------------------------------+
	//	| NTMS_DOORSTATE_OPEN 0x00000002    | The library door is open.            |
	//	+-----------------------------------+--------------------------------------+
	//	| NTMS_DOORSTATE_UNKNOWN 0x00000000 | The state of the library is unknown. |
	//	+-----------------------------------+--------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// MaxOpenSecs:  The maximum number of seconds the door is to remain open.
	MaxOpenSecs uint16 `idl:"name:MaxOpenSecs" json:"max_open_secs"`
	// Library:  The library that contains this door.
	Library *GUID `idl:"name:Library" json:"library"`
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
//
// The NTMS_IEPORTINFORMATION structure defines properties specific to an IE port object.
type IEPortInformation struct {
	// Number:   The library port number.
	Number uint32 `idl:"name:Number" json:"number"`
	// Content:  The full/empty state of the NTMS_IEPORT object. This MUST be one of the
	// following values.
	//
	//	+-------------------------------------+---------------------------------------+
	//	|                                     |                                       |
	//	|                VALUE                |                MEANING                |
	//	|                                     |                                       |
	//	+-------------------------------------+---------------------------------------+
	//	+-------------------------------------+---------------------------------------+
	//	| NTMS_PORTCONTENT_UNKNOWN 0x00000000 | The content of the port is not known. |
	//	+-------------------------------------+---------------------------------------+
	//	| NTMS_PORTCONTENT_FULL 0x00000001    | The port is full.                     |
	//	+-------------------------------------+---------------------------------------+
	//	| NTMS_PORTCONTENT_EMPTY 0x00000002   | The port is empty.                    |
	//	+-------------------------------------+---------------------------------------+
	Content uint32 `idl:"name:Content" json:"content"`
	// Position:  The position of the NTMS_IEPORT object. This MUST be one of the following
	// values.
	//
	//	+----------------------------------------+-------------------------------+
	//	|                                        |                               |
	//	|                 VALUE                  |            MEANING            |
	//	|                                        |                               |
	//	+----------------------------------------+-------------------------------+
	//	+----------------------------------------+-------------------------------+
	//	| NTMS_PORTPOSITION_UNKNOWN 0x00000000   | The port position is unknown. |
	//	+----------------------------------------+-------------------------------+
	//	| NTMS_PORTPOSITION_EXTENDED 0x00000001  | The port is extended.         |
	//	+----------------------------------------+-------------------------------+
	//	| NTMS_PORTPOSITION_RETRACTED 0x00000002 | The port is retracted.        |
	//	+----------------------------------------+-------------------------------+
	Position uint32 `idl:"name:Position" json:"position"`
	// MaxExtendSecs:  The maximum number of seconds the port is allowed to remain open
	// before an operator request is issued. Valid values are between zero and 65,535 seconds.
	MaxExtendSecs uint16 `idl:"name:MaxExtendSecs" json:"max_extend_secs"`
	// Library:  The library that contains the port.
	Library *GUID `idl:"name:Library" json:"library"`
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
//
// The NTMS_LMIDINFORMATION structure defines the properties specific to a logical media
// object.
type LMIDInformation struct {
	// MediaPool:  The unique identifier of the media pool that contains the logical media.
	MediaPool *GUID `idl:"name:MediaPool" json:"media_pool"`
	// dwNumberOfPartitions:  The number of sides in the media object.
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
//
// The NTMS_COMPUTERINFORMATION structure defines the properties specific to the RSM
// server.
type ComputerInformation struct {
	// dwLibRequestPurgeTime:  The total number of seconds that the completed library requests
	// are maintained in the work queue.<15>
	LibRequestPurgeTime uint32 `idl:"name:dwLibRequestPurgeTime" json:"lib_request_purge_time"`
	// dwOpRequestPurgeTime:  The total number of seconds that the completed operator requests
	// are maintained in the operator request queue.<16>
	OperationRequestPurgeTime uint32 `idl:"name:dwOpRequestPurgeTime" json:"operation_request_purge_time"`
	// dwLibRequestFlags:  The library request options. This member contains the following
	// flag fields.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | F | A |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| A NTMS_LIBREQFLAGS_NOAUTOPURGE   | Library requests are not purged from the work queue. This MUST be set to NULL by |
	//	|                                  | default.                                                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| F NTMS_LIBREQFLAGS_NOFAILEDPURGE | Failed work items are not purged from the work queue. This MUST be set to NULL   |
	//	|                                  | by default.                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	LibRequestFlags uint32 `idl:"name:dwLibRequestFlags" json:"lib_request_flags"`
	// dwOpRequestFlags:  The operator request options. Possible values include the following.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQFLAGS_NOAUTOPURGE 0x01   | Operator requests MUST NOT be purged from the work queue. This MUST be set to    |
	//	|                                    | NULL by default.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQFLAGS_NOFAILEDPURGE 0x02 | Failed operator requests MUST NOT be purged from the queue. This MUST be set to  |
	//	|                                    | NULL by default.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQFLAGS_NOALERTS 0x10      | The alert for operator requests MUST be disabled.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_OPREQFLAGS_NOTRAYICON 0x20    | The taskbar icon for operator requests MUST be disabled.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	OperationRequestFlags uint32 `idl:"name:dwOpRequestFlags" json:"operation_request_flags"`
	// dwMediaPoolPolicy:  Media pool policies. Possible values include the following.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                          |                                                                                  |
	//	|                  VALUE                   |                                     MEANING                                      |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLPOLICY_PURGEOFFLINESCRATCH 0x01 | Any free media that is ejected MUST be automatically deleted. This MUST be set   |
	//	|                                          | to NULL by default.                                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| NTMS_POOLPOLICY_KEEPOFFLINEIMPORT 0x02   | Any import media that is ejected MUST NOT be deleted automatically. This MUST be |
	//	|                                          | set to NULL by default.                                                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	MediaPoolPolicy uint32 `idl:"name:dwMediaPoolPolicy" json:"media_pool_policy"`
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
//
// The NTMS_OPREQUESTINFORMATIONA structure describes the properties of an operator
// request, in ASCII.
type OperationRequestInformationA struct {
	// Request:  The value from the NtmsOpreqCommand enumeration that specifies the type
	// of the operator request.
	Request uint32 `idl:"name:Request" json:"request"`
	// Submitted:  A SYSTEMTIME structure that specifies the time at which the request was
	// submitted.
	Submitted *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	// State:  A value from the NtmsOpreqState enumeration that specifies the state of the
	// operator request.
	State uint32 `idl:"name:State" json:"state"`
	// szMessage:   A null-terminated string that contains operator message text. For example,
	// "The cleaner cartridge in %s (slot #%d) has reached its maximum usage and needs to
	// be replaced."
	Message []byte `idl:"name:szMessage" json:"message"`
	// Arg1Type:  A value from the NtmsObjectsTypes enumeration that specifies the type
	// of object in Arg1.
	Arg1Type uint32 `idl:"name:Arg1Type" json:"arg1_type"`
	// Arg1:  This parameter MUST be set based on the value of Request.
	//
	// If Request = NTMS_OPREQ_NEWMEDIA (0x00000001), Arg1 MUST be set to the identifier
	// of the media pool requiring new media.
	//
	// If Request = NTMS_OPREQ_CLEANER (0x00000002), Arg1 MUST be set to the identifier
	// of the library requiring the cleaner cartridge.
	//
	// If Request = NTMS_OPREQ_DEVICESERVICE (0x00000003), Arg1 MUST be set to the identifier
	// of the device requiring service.
	//
	// If Request = NTMS_OPREQ_MOVEMEDIA (0x00000004), Arg1 MUST be set to the identifier
	// of the physical medium to move.
	Arg1 *GUID `idl:"name:Arg1" json:"arg1"`
	// Arg2Type:  A value from the NtmsObjectsTypes enumeration that specifies the type
	// of object in Arg2.
	Arg2Type uint32 `idl:"name:Arg2Type" json:"arg2_type"`
	// Arg2:  This parameter MUST be set based on the value of Request.
	//
	// If Request = NTMS_OPREQ_NEWMEDIA (0x00000001), Arg2 MAY be set to the identifier
	// of the library in which the new media MUST be placed.
	//
	// If Request = NTMS_OPREQ_MOVEMEDIA (0x00000004), Arg2 MUST be set to the identifier
	// of the library to which the physical medium MUST be moved.
	Arg2 *GUID `idl:"name:Arg2" json:"arg2"`
	// szApplication:   A null-terminated sequence of Unicode characters that specifies
	// the name of the application that submitted the operator request.
	Application []byte `idl:"name:szApplication" json:"application"`
	// szUser:   A null-terminated sequence of Unicode characters that specifies the name
	// of the interactive user who submitted the operator request.
	User []byte `idl:"name:szUser" json:"user"`
	// szComputer:   A null-terminated sequence of Unicode characters that specifies the
	// name of the computer that submitted the operator request.
	Computer []byte `idl:"name:szComputer" json:"computer"`
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
//
// The NTMS_OPREQUESTINFORMATIONW structure describes the properties of an operator
// request, in Unicode.
type OperationRequestInformationW struct {
	// Request:  A value from the NtmsOpreqCommand enumeration that specifies the type of
	// the operator request.
	Request uint32 `idl:"name:Request" json:"request"`
	// Submitted:  A SYSTEMTIME structure that specifies the time at which the request was
	// submitted.
	Submitted *dtyp.SystemTime `idl:"name:Submitted" json:"submitted"`
	// State:  A value from the NtmsOpreqState enumeration that specifies the state of the
	// operator request.
	State uint32 `idl:"name:State" json:"state"`
	// szMessage:   A null-terminated operator message text.
	Message string `idl:"name:szMessage;string" json:"message"`
	// Arg1Type:  A value from the NtmsObjectsTypes enumeration that specifies the type
	// of object in Arg1.
	Arg1Type uint32 `idl:"name:Arg1Type" json:"arg1_type"`
	// Arg1:  This parameter MUST be set based on the value of Request.
	//
	// If Request = NTMS_OPREQ_NEWMEDIA (0x00000001), Arg1 MUST be set to the identifier
	// of the media pool requiring new media.
	//
	// If Request = NTMS_OPREQ_CLEANER (0x00000002), Arg1 MUST be set to the identifier
	// of the library requiring the cleaner cartridge.
	//
	// If Request = NTMS_OPREQ_DEVICESERVICE (0x00000003), Arg1 MUST be set to the identifier
	// of the device requiring service.
	Arg1 *GUID `idl:"name:Arg1" json:"arg1"`
	// Arg2Type:  A value from the NtmsObjectsTypes enumeration that specifies the type
	// of object in Arg2.
	Arg2Type uint32 `idl:"name:Arg2Type" json:"arg2_type"`
	// Arg2:  This parameter MUST be set based on the value of Request.
	//
	// If Request = NTMS_OPREQ_NEWMEDIA (0x00000001), Arg2 MAY be set to the identifier
	// of the library in which the new media MUST be placed.
	Arg2 *GUID `idl:"name:Arg2" json:"arg2"`
	// szApplication:   A null-terminated sequence of Unicode characters that specifies
	// the name of the application that submitted the operator request.
	Application string `idl:"name:szApplication;string" json:"application"`
	// szUser:   A null-terminated sequence of Unicode characters that specifies the name
	// of the interactive user who submitted the operator request.
	User string `idl:"name:szUser;string" json:"user"`
	// szComputer:   A null-terminated sequence of Unicode characters that specifies the
	// name of the computer that submitted the operator request.
	Computer string `idl:"name:szComputer;string" json:"computer"`
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
//
// The NTMS_PARTITIONINFORMATIONA structure describes the properties of a media side
// object, as a sequence of ASCII characters.
type PartitionInformationA struct {
	// PhysicalMedia:  The identifier of the medium that contains the side.
	PhysicalMedia *GUID `idl:"name:PhysicalMedia" json:"physical_media"`
	// LogicalMedia:  The identifier of the piece of logical media that contains the side.
	// This MUST be set to 0 if the side is not allocated.
	//
	//	+-------+----------------------------+
	//	|       |                            |
	//	| VALUE |          MEANING           |
	//	|       |                            |
	//	+-------+----------------------------+
	//	+-------+----------------------------+
	//	|     0 | The side is not allocated. |
	//	+-------+----------------------------+
	LogicalMedia *GUID `idl:"name:LogicalMedia" json:"logical_media"`
	// State:  The value from the NtmsPartitionState (section 2.2.4.7) enumeration describing
	// the state of the side.
	State uint32 `idl:"name:State" json:"state"`
	// Side:  A zero-relative value that indicates which side of a multisided media this
	// is. For single-sided media, this value MUST be 0. For dual-sided media, one NTMS_PARTITIONINFORMATIONA
	// record MUST have its member set to 0, and a second NTMS_PARTITIONINFORMATIONA record
	// MUST have its member set to 1.
	//
	//	+-------------+-----------------------------------------------------------------------------+
	//	|             |                                                                             |
	//	| MEDIA\VALUE |                                   MEANING                                   |
	//	|             |                                                                             |
	//	+-------------+-----------------------------------------------------------------------------+
	//	+-------------+-----------------------------------------------------------------------------+
	//	|           0 | The only side of single-sided media, or the first side of dual-sided media. |
	//	+-------------+-----------------------------------------------------------------------------+
	//	|           1 | The second side of dual-sided media.                                        |
	//	+-------------+-----------------------------------------------------------------------------+
	Side uint16 `idl:"name:Side" json:"side"`
	// dwOmidLabelIdLength:  The length of the label identification string of the on-media
	// identifier.
	OMIDLabelIDLength uint32 `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	// OmidLabelId:   The label identifier of the on-media identifier.
	OMIDLabelID []byte `idl:"name:OmidLabelId" json:"omid_label_id"`
	// szOmidLabelType:   A null-terminated sequence of ASCII characters specifying the
	// label type of the on-media identifier.
	OMIDLabelType []byte `idl:"name:szOmidLabelType" json:"omid_label_type"`
	// szOmidLabelInfo:   A null-terminated sequence of ASCII characters specifying the
	// label information of the on-media identifier.
	OMIDLabelInfo []byte `idl:"name:szOmidLabelInfo" json:"omid_label_info"`
	// dwMountCount:  The number of times the medium has been mounted into a drive since
	// being initialized for this server. This member MUST be initialized to 0 when the
	// object is created in the database.
	MountCount uint32 `idl:"name:dwMountCount" json:"mount_count"`
	// dwAllocateCount:  The number of times the medium has been allocated since being initialized
	// for this server.
	AllocateCount uint32 `idl:"name:dwAllocateCount" json:"allocate_count"`
	// Capacity:  The number of bytes available on this side.
	Capacity *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
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
//
// The NTMS_PARTITIONINFORMATIONW structure describes the properties of a media side
// object, as a sequence of Unicode characters. Unicode encoding is specified in [UNICODE].
type PartitionInformationW struct {
	// PhysicalMedia:  The identifier of the medium that contains the side.
	PhysicalMedia *GUID `idl:"name:PhysicalMedia" json:"physical_media"`
	// LogicalMedia:  The identifier of the piece of logical media that contains the side.
	// This MUST be set to GUID_NULL if the side is not allocated.
	//
	//	+-------------+----------------------------+
	//	|             |                            |
	//	|    VALUE    |          MEANING           |
	//	|             |                            |
	//	+-------------+----------------------------+
	//	+-------------+----------------------------+
	//	| "GUID_NULL" | The side is not allocated. |
	//	+-------------+----------------------------+
	LogicalMedia *GUID `idl:"name:LogicalMedia" json:"logical_media"`
	// State:  A value from the NtmsPartitionState (section 2.2.4.7) enumeration describing
	// the state of the side.
	State uint32 `idl:"name:State" json:"state"`
	// Side:  A zero-relative value that indicates which side of a multisided media this
	// is. For single-sided media, this value MUST be 0. For dual-sided media, one NTMS_PARTITIONINFORMATIONW
	// record MUST have its member set to 0, and a second NTMS_PARTITIONINFORMATIONW record
	// MUST have its member set to 1.
	//
	//	+-------------+----------------------------------------------------------------------------+
	//	|             |                                                                            |
	//	| MEDIA\VALUE |                                  MEANING                                   |
	//	|             |                                                                            |
	//	+-------------+----------------------------------------------------------------------------+
	//	+-------------+----------------------------------------------------------------------------+
	//	|           0 | The only side of single-sided media or the first side of dual-sided media. |
	//	+-------------+----------------------------------------------------------------------------+
	//	|           1 | The second side of dual-sided media.                                       |
	//	+-------------+----------------------------------------------------------------------------+
	Side uint16 `idl:"name:Side" json:"side"`
	// dwOmidLabelIdLength:  The length of the label identification string of the on-media
	// identifier.
	OMIDLabelIDLength uint32 `idl:"name:dwOmidLabelIdLength" json:"omid_label_id_length"`
	// OmidLabelId:   The label identifier of the on-media identifier.
	OMIDLabelID []byte `idl:"name:OmidLabelId" json:"omid_label_id"`
	// szOmidLabelType:   A null-terminated sequence of Unicode UTF-16 characters specifying
	// the label type of the on-media identifier.
	OMIDLabelType string `idl:"name:szOmidLabelType;string" json:"omid_label_type"`
	// szOmidLabelInfo:   A null-terminated sequence of Unicode characters specifying the
	// label information of the on-media identifier.
	OMIDLabelInfo string `idl:"name:szOmidLabelInfo;string" json:"omid_label_info"`
	// dwMountCount:  The number of times the medium has been mounted into a drive. This
	// member is initialized to 0 when the object is created in the server information database.
	MountCount uint32 `idl:"name:dwMountCount" json:"mount_count"`
	// dwAllocateCount:  The number of times the medium has been allocated.
	AllocateCount uint32 `idl:"name:dwAllocateCount" json:"allocate_count"`
	// Capacity:  The number of bytes available on this side.
	Capacity *dtyp.LargeInteger `idl:"name:Capacity" json:"capacity"`
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
//
// The NTMS_PMIDINFORMATIONA structure describes the properties of a physical media
// object, as a sequence of ASCII characters.
type PMIDInformationA struct {
	// CurrentLibrary:  The identifier of the library in which the media is contained.
	CurrentLibrary *GUID `idl:"name:CurrentLibrary" json:"current_library"`
	// MediaPool:  The identifier of the media pool to which the media is assigned.
	MediaPool *GUID `idl:"name:MediaPool" json:"media_pool"`
	// Location:  The identifier of the physical location object for the media.
	Location *GUID `idl:"name:Location" json:"location"`
	// LocationType:  Any of the NTMS_STORAGESLOT, NTMS_DRIVE, or NTMS_IEPORT values defined
	// in the NtmsObjectsTypes (section 2.2.1.6) enumeration, specifying the type of the
	// current location for a piece of physical media.
	LocationType uint32 `idl:"name:LocationType" json:"location_type"`
	// MediaType:  The identifier of the media type object for the medium.
	MediaType *GUID `idl:"name:MediaType" json:"media_type"`
	// HomeSlot:  The identifier of the library storage slot in which the medium is stored.
	HomeSlot *GUID `idl:"name:HomeSlot" json:"home_slot"`
	// szBarCode:   The null-terminated Unicode string specifying the numeric value of the
	// bar code of the media. If the bar code is not available, BarCodeState MUST be set
	// to NTMS_BARCODESTATE_UNREADABLE. For more information, see section 2.2.4.1.
	BarCode []byte `idl:"name:szBarCode" json:"bar_code"`
	// BarCodeState:  A value from the NtmsBarCodeState (section 2.2.4.1) enumeration specifying
	// the state of the bar code.
	BarCodeState uint32 `idl:"name:BarCodeState" json:"bar_code_state"`
	// szSequenceNumber:   A sequential number assigned to the specified medium as a human-readable
	// value.
	SequenceNumber []byte `idl:"name:szSequenceNumber" json:"sequence_number"`
	// MediaState:  The value from the NtmsMediaState (section 2.2.4.4) enumeration describing
	// the state of the media.
	MediaState uint32 `idl:"name:MediaState" json:"media_state"`
	// dwNumberOfPartitions:  The number of sides on the medium.
	NumberOfPartitions uint32 `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	// dwMediaTypeCode:  The SCSI [ANSI-131-1994] type code of the medium.
	MediaTypeCode uint32 `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	// dwDensityCode:  The SCSI density code of the medium.
	DensityCode uint32 `idl:"name:dwDensityCode" json:"density_code"`
	// MountedPartition:  The identifier of the media side that is currently mounted.
	MountedPartition *GUID `idl:"name:MountedPartition" json:"mounted_partition"`
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
//
// The NTMS_PMIDINFORMATIONW structure describes the properties of a physical media
// object, as a sequence of Unicode characters.
type PMIDInformationW struct {
	// CurrentLibrary:  The identifier of the library in which the medium is contained.
	CurrentLibrary *GUID `idl:"name:CurrentLibrary" json:"current_library"`
	// MediaPool:  The identifier of the media pool to which the medium is assigned.
	MediaPool *GUID `idl:"name:MediaPool" json:"media_pool"`
	// Location:  The identifier of the physical location object for the medium.
	Location *GUID `idl:"name:Location" json:"location"`
	// LocationType:  Any of the NTMS_STORAGESLOT, NTMS_DRIVE, or NTMS_IEPORT values defined
	// in the NtmsObjectsTypes (section 2.2.1.6) enumeration, specifying the type of the
	// current location for a piece of physical media.
	LocationType uint32 `idl:"name:LocationType" json:"location_type"`
	// MediaType:  The identifier of the media type object for the medium.
	MediaType *GUID `idl:"name:MediaType" json:"media_type"`
	// HomeSlot:  The identifier of the library storage slot in which the medium is stored.
	HomeSlot *GUID `idl:"name:HomeSlot" json:"home_slot"`
	// szBarCode:   The null-terminated bar code of the medium.
	BarCode string `idl:"name:szBarCode;string" json:"bar_code"`
	// BarCodeState:  The value from the NtmsBarCodeState (section 2.2.4.1) enumeration
	// specifying the state of the bar code.
	BarCodeState uint32 `idl:"name:BarCodeState" json:"bar_code_state"`
	// szSequenceNumber:   Sequential number assigned to the specified medium as a human-readable
	// value. This value MUST be transcribed by a user on the medium so it can be located
	// in an offline library.
	SequenceNumber string `idl:"name:szSequenceNumber;string" json:"sequence_number"`
	// MediaState:  The value from the NtmsMediaState (section 2.2.4.4) enumeration describing
	// the state of the media.
	MediaState uint32 `idl:"name:MediaState" json:"media_state"`
	// dwNumberOfPartitions:  The number of sides on the medium.
	NumberOfPartitions uint32 `idl:"name:dwNumberOfPartitions" json:"number_of_partitions"`
	// dwMediaTypeCode:  The SCSI [ANSI-131-1994] type code of the medium.
	MediaTypeCode uint32 `idl:"name:dwMediaTypeCode" json:"media_type_code"`
	// dwDensityCode:  The SCSI density code of the medium.
	DensityCode uint32 `idl:"name:dwDensityCode" json:"density_code"`
	// MountedPartition:  The identifier of the media side that is currently mounted.
	MountedPartition *GUID `idl:"name:MountedPartition" json:"mounted_partition"`
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
//
// The RSM_MESSAGE structure describes a message.
type RSMMessage struct {
	// lpguidOperation:  A pointer to the identifier of the operation to which the message
	// refers.
	LpguidOperation *dtyp.GUID `idl:"name:lpguidOperation;pointer:unique" json:"lpguid_operation"`
	// dwNtmsType:  A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of object to which the message refers.
	NTMSType uint32 `idl:"name:dwNtmsType" json:"ntms_type"`
	// dwState:  A value from the NtmsLmState (section 2.2.1.10) enumeration specifying
	// the state of the operation to which the message refers.
	State uint32 `idl:"name:dwState" json:"state"`
	// dwFlags:  This parameter is unused. It MUST be 0 and MUST be ignored on receipt.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwPriority:  The priority of the message.
	Priority uint32 `idl:"name:dwPriority" json:"priority"`
	// dwErrorCode:  An implementation-specific, nonzero error code.
	ErrorCode uint32 `idl:"name:dwErrorCode" json:"error_code"`
	// lpszComputerName:  A null-terminated sequence of Unicode characters specifying the
	// name of the computer from which the message was sent.
	ComputerName string `idl:"name:lpszComputerName;string;pointer:unique" json:"computer_name"`
	// lpszApplication:  A null-terminated sequence of Unicode characters specifying the
	// name of the application sending the message.
	Application string `idl:"name:lpszApplication;string" json:"application"`
	// lpszUser:  A null-terminated sequence of Unicode characters specifying the name of
	// the user sending the message.
	User string `idl:"name:lpszUser;string" json:"user"`
	// lpszTimeSubmitted:  The null-terminated time at which the message was created.
	TimeSubmitted string `idl:"name:lpszTimeSubmitted;string" json:"time_submitted"`
	// lpszMessage:  The null-terminated description of the message.
	Message string `idl:"name:lpszMessage;string" json:"message"`
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
//
// The NTMS_OBJECTINFORMATIONA structure describes the properties of RSM objects, in
// ASCII.
type ObjectInformationA struct {
	// dwSize:  The size, in bytes, of the structure.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// dwType:  A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// Created:  A SYSTEMTIME structure specifying the time when the object was created.
	Created *dtyp.SystemTime `idl:"name:Created" json:"created"`
	// Modified:  A SYSTEMTIME structure specifying the time when the object was last modified.
	Modified *dtyp.SystemTime `idl:"name:Modified" json:"modified"`
	// ObjectGuid:  The identifier of the object.
	ObjectGUID *GUID `idl:"name:ObjectGuid" json:"object_guid"`
	// Enabled:  If set to TRUE, the object MUST be enabled; if set to FALSE, the object
	// MUST NOT be enabled.
	Enabled bool `idl:"name:Enabled" json:"enabled"`
	// dwOperationalState:  A value from the NtmsOperationalState (section 2.2.4.5) enumeration
	// specifying the operation state of the object.
	OperationalState uint32 `idl:"name:dwOperationalState" json:"operational_state"`
	// szName:   A null-terminated sequence of ASCII characters specifying the name of the
	// object.
	Name []byte `idl:"name:szName" json:"name"`
	// szDescription:   The null-terminated description of the object.<13>
	Description []byte `idl:"name:szDescription" json:"description"`
	// Info:  A device or system control object information that is specific to the value
	// of dwType.
	Info *ObjectInformationA_Info `idl:"name:Info;switch_is:dwType" json:"info"`
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
//
// The NTMS_OBJECTINFORMATIONA structure describes the properties of RSM objects, in
// ASCII.
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
	// Drive:   An NTMS_DRIVEINFORMATIONA structure that describes the properties of a drive.
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
	// DriveType:  An NTMS_DRIVETYPEINFORMATIONA structure that describes the properties
	// specific to a type of drive.
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
	// Library:  An NTMS_LIBRARYINFORMATION structure that describes the properties of a
	// media library.
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
	// Changer:  An NTMS_CHANGERINFORMATIONA structure that describes the properties of
	// a changer object.
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
	// ChangerType:  An NTMS_CHANGERTYPEINFORMATIONA structure that describes the properties
	// specific to a type of changer object.
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
	// StorageSlot:  An NTMS_STORAGESLOTINFORMATION structure that describes the properties
	// of a media storage slot.
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
	// IEDoor:  An NTMS_IEDOORINFORMATION structure that describes the properties of an
	// access door.
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
	// IEPort:  An NTMS_IEPORTINFORMATION structure that describes the properties of an
	// inject/eject port.
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
	// PhysicalMedia:  An NTMS_PMIDINFORMATIONA structure that describes the properties
	// of a physical media object.
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
	// LogicalMedia:  An NTMS_LMIDINFORMATION structure that describes the properties of
	// a logical media object.
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
	// Partition:  An NTMS_PARTITIONINFORMATIONA structure that describes the properties
	// of a media-side object.
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
	// MediaPool:  An NTMS_MEDIAPOOLINFORMATION structure that describes the properties
	// of a media pool.
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
	// MediaType:  An NTMS_MEDIATYPEINFORMATION structure that describes the properties
	// specific to a type of media.
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
	// LibRequest:  An NTMS_LIBREQUESTINFORMATIONA structure that describes the properties
	// of a library request.
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
	// OpRequest:  An NTMS_OPREQUESTINFORMATIONA structure that describes the properties
	// of an operator request.
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
	// Computer:  An NTMS_COMPUTERINFORMATION structure that describes the properties of
	// a computer.
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
//
// The NTMS_OBJECTINFORMATIONW structure describes the properties of RSM objects, in
// Unicode.
type ObjectInformationW struct {
	// dwSize:  The size, in bytes, of the structure.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// dwType:  A value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// Created:  A SYSTEMTIME structure specifying the time when the object was created.
	Created *dtyp.SystemTime `idl:"name:Created" json:"created"`
	// Modified:  A SYSTEMTIME structure specifying the time when the object was last modified.
	Modified *dtyp.SystemTime `idl:"name:Modified" json:"modified"`
	// ObjectGuid:  The identifier of the object.
	ObjectGUID *GUID `idl:"name:ObjectGuid" json:"object_guid"`
	// Enabled:  If set to TRUE, the object MUST be enabled; if set to FALSE, the object
	// MUST NOT be enabled.
	Enabled bool `idl:"name:Enabled" json:"enabled"`
	// dwOperationalState:  A value from the NtmsOperationalState (section 2.2.4.5) enumeration
	// specifying the operation state of the object.
	OperationalState uint32 `idl:"name:dwOperationalState" json:"operational_state"`
	// szName:   A null-terminated sequence of Unicode UTF-16 characters specifying the
	// name of the object.
	Name string `idl:"name:szName;string" json:"name"`
	// szDescription:   The null-terminated description of the object.<14>
	Description string `idl:"name:szDescription;string" json:"description"`
	// Info:  A device or system control object information that is specific to the value
	// of dwType.
	Info *ObjectInformationW_Info `idl:"name:Info;switch_is:dwType" json:"info"`
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
//
// The NTMS_OBJECTINFORMATIONW structure describes the properties of RSM objects, in
// Unicode.
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
	// Drive:   An NTMS_DRIVEINFORMATIONW structure that describes the properties of a drive.
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
	// DriveType:  An NTMS_DRIVETYPEINFORMATIONW structure that describes the properties
	// specific to a type of drive.
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
	// Library:  An NTMS_LIBRARYINFORMATION structure that describes the properties of a
	// media library.
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
	// Changer:  An NTMS_CHANGERINFORMATIONW structure that describes the properties of
	// a changer object.
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
	// ChangerType:  An NTMS_CHANGERTYPEINFORMATIONW structure that describes the properties
	// specific to a type of changer object.
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
	// StorageSlot:  An NTMS_STORAGESLOTINFORMATION structure that describes the properties
	// of a media storage slot.
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
	// IEDoor:  An NTMS_IEDOORINFORMATION structure that describes the properties of an
	// access door.
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
	// IEPort:  An NTMS_IEPORTINFORMATION structure that describes the properties of an
	// IE port.
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
	// PhysicalMedia:  An NTMS_PMIDINFORMATIONW structure that describes the properties
	// of a physical media object.
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
	// LogicalMedia:  An NTMS_LMIDINFORMATION structure that describes the properties of
	// a logical media object.
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
	// Partition:  An NTMS_PARTITIONINFORMATIONW structure that describes the properties
	// of a media-side object.
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
	// MediaPool:  An NTMS_MEDIAPOOLINFORMATION structure that describes the properties
	// of a media pool.
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
	// MediaType:  An NTMS_MEDIATYPEINFORMATION structure that describes the properties
	// specific to a type of media.
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
	// LibRequest:  An NTMS_LIBREQUESTINFORMATIONW structure that describes the properties
	// of a library request.
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
	// OpRequest:  An NTMS_OPREQUESTINFORMATIONW structure that describes the properties
	// of an operator request.
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
	// Computer:  An NTMS_COMPUTERINFORMATION structure that describes the properties of
	// a computer.
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
