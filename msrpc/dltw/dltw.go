// The dltw package implements the DLTW client protocol.
//
// # Introduction
//
// This document specifies the Distributed Link Tracking: Workstation Protocol.
//
// Distributed Link Tracking (DLT) consists of two protocols that work together to discover
// the new location of a file that has moved. DLT can determine if the file has moved
// on a mass-storage device, within a computer, or between computers in a network. The
// Distributed Link Tracking: Workstation Protocol helps a computer locate files that
// have been moved within a computer or between computers in a computer network.
//
// In addition to the Distributed Link Tracking: Workstation Protocol, DLT includes
// the Distributed Link Tracking: Central Manager Protocol that keeps track of file
// and volume moves as well as other relevant information from participating computers
// so it can provide this information in response to workstation queries. Both DLT protocols
// are remote procedure call (RPC) interfaces.
//
// # Overview
//
// The Distributed Link Tracking: Workstation Protocol is based on the RPC runtime,
// as specified in [C706] and [MS-RPCE], and on the server message block (SMB) protocol
// and extensions, as specified in [MS-SMB] and [MS-SMB2].
//
// This protocol is used by a client to get a file's identity and location on the server
// computer as a MachineID, FileID, FileLocation, and Universal Naming Convention (UNC)
// name. If a client contacts a server that previously stored the file, but the file
// has been moved to a new computer, the server might be able to return the MachineID
// of the computer to which the file was moved, so that the client can contact the DLT
// Workstation server on the new computer to get the file's current UNC and FileLocation,
// or another referral. This process of following referrals continues until a server
// returns the file's UNC name and FileLocation, or an error.
//
// Rather than following referrals in this manner, a client can use the Distributed
// Link Tracking: Central Manager Protocol to determine a file's current MachineID and
// FileLocation, and then use that information to initiate a call to the DLT Workstation
// server on the computer indicated by that MachineID. For more information on using
// the Distributed Link Tracking: Central Manager Protocol in combination with the Distributed
// Link Tracking: Workstation Protocol, see [MS-DLTM].
//
// The following is a scenario that describes the DLT protocols working together:
//
// *
//
// A file is created on computer M1. M1 assigns identifiers, specifically FileID and
// FileLocation, to the file.
//
// *
//
// Computer M0 takes note of the file, locally storing its identifiers.
//
// *
//
// The file is moved from computer M1 to M2 and from there to M3. In concert with these
// moves, the file maintains its FileID but gets a new FileLocation assigned.
//
// *
//
// If the Distributed Link Tracking: Central Manager Protocol is used, clients on computers
// M1 and M2 notify the server that the file has been moved, indicating the file's FileID
// and its old and new FileLocation values.
//
// *
//
// Computer M0 finds the file in its new location in one of two ways:
//
// *
//
// Using only the Distributed Link Tracking: Workstation Protocol:
//
// *
//
// M0 contacts M1, using the identifiers stored previously, and learns that the file
// was moved to M2.
//
// *
//
// M0 contacts M2 and learns that the file was moved to M3.
//
// *
//
// M0 contacts M3 and learns the file's new name and location.
//
// *
//
// Using both the Distributed Link Tracking: Workstation Protocol and the Distributed
// Link Tracking: Central Manager Protocol:
//
// *
//
// M0 contacts a DLT Central Manager server to query the current location of the file.
//
// *
//
// The server queries its tables and determines that the file is currently on computer
// M3.
//
// *
//
// M0 contacts the DLT Workstation client on M3 and learns the file's new name and location.
//
// The following is an example of a file being moved between computers, and shows in
// more detail how to use the Distributed Link Tracking: Workstation Protocol to determine
// the file's new location. In this example, only the Distributed Link Tracking: Workstation
// Protocol is used, without using the Distributed Link Tracking: Central Manager Protocol.
//
// *
//
// In the initial state, a file is located on a computer named M1. Assume that the file
// is named "F1.txt", and can be located via the UNC "\\M1\share1\F1.txt".
//
// *
//
// Before the file is moved, a user on computer M0 requests that information about the
// file be saved, <1> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_1 ) so that
// its location can be determined after it has been moved. As a result, M0 stores the
// UNC, the MachineID, the FileLocation, and the FileID.
//
// *
//
// The file is moved from machine M1 to machine M2; for example, to the UNC "\\M2\share2\F2.txt".
// M1 stores the file's old FileLocation and FileID, as well as the file's new FileLocation
// and MachineID.
//
// *
//
// When M0 attempts to open the file <2> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_2
// ) by using the UNC "\\M1\share1\F1.txt", it receives a file-not-found error message.
// M0 then initiates a call to the DLT Workstation server on M1 with the previously
// stored FileID and FileLocation of the file.
//
// *
//
// The DLT Workstation server on M1 returns to the DLT Workstation client on M0 that
// the file was moved, and specifies the MachineID of the file's new location (M2),
// as well as the file's new FileLocation value.
//
// *
//
// M0 then repeats the call, but this time to the DLT Workstation server on M2 with
// the new FileLocation value.
//
// *
//
// The DLT Workstation server on M2 returns to the DLT Workstation client on M0 the
// file's new UNC, "\\M2\share2\F2.txt".
//
// *
//
// The DLT Workstation client on M0 then updates its stored values with the updated
// UNC and FileLocation values.
package dltw

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
	GoPackage = "dltw"
)

// ObjectID structure represents CObjId RPC structure.
//
// The CObjId type is used to represent an ObjectID, which is a unique identifier that
// represents the identity of a file within a file system volume.
type ObjectID struct {
	// _object:  This field MUST contain an identifier for a file. It is unique within the
	// volume on which the file was created. Restrictions on the value of a VolumeID are
	// defined in section 3.1.1.
	Object *dtyp.GUID `idl:"name:_object" json:"object"`
}

func (o *ObjectID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Object != nil {
		if err := o.Object.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &dtyp.GUID{}
	}
	if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VolumeID structure represents CVolumeId RPC structure.
//
// The CVolumeId type is used to represent a VolumeID, which is a unique identifier
// that represents the identity of a file system volume.
type VolumeID struct {
	// _volume:  This field MUST contain a GUID for a volume. The lowest-order bit of this
	// value MUST be zero. Further restrictions on the value of a VolumeID are defined in
	// section 3.1.1.
	Volume *dtyp.GUID `idl:"name:_volume" json:"volume"`
}

func (o *VolumeID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VolumeID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Volume != nil {
		if err := o.Volume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Volume == nil {
		o.Volume = &dtyp.GUID{}
	}
	if err := o.Volume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// MachineID structure represents CMachineId RPC structure.
//
// The CMachineId structure is used to represent a MachineID, which is a unique identifier
// that represents the identity of a computer.
type MachineID struct {
	// _szMachine:  This member MUST be a NetBIOS name, as specified in [RFC1088]. This
	// name MUST be terminated with a zero byte, and any remaining bytes MUST also be zero.
	Machine []byte `idl:"name:_szMachine" json:"machine"`
}

func (o *MachineID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MachineID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Machine {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Machine[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Machine); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *MachineID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Machine = make([]byte, 16)
	for i1 := range o.Machine {
		i1 := i1
		if err := w.ReadData(&o.Machine[i1]); err != nil {
			return err
		}
	}
	return nil
}

// DomainRelativeObjectID structure represents CDomainRelativeObjId RPC structure.
//
// The CDomainRelativeObjId structure is used to represent a file's FileLocation or
// its FileID, which is the FileLocation of a file at the time it was originally created.
type DomainRelativeObjectID struct {
	// _volume:  A volume identifier for the volume that contains the file, as specified
	// in section 2.2.4.
	Volume *VolumeID `idl:"name:_volume" json:"volume"`
	// _object:  A file identifier for the file, as specified in section 2.2.5.
	Object *ObjectID `idl:"name:_object" json:"object"`
}

func (o *DomainRelativeObjectID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DomainRelativeObjectID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Volume != nil {
		if err := o.Volume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VolumeID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Object != nil {
		if err := o.Object.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainRelativeObjectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Volume == nil {
		o.Volume = &VolumeID{}
	}
	if err := o.Volume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &ObjectID{}
	}
	if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}
