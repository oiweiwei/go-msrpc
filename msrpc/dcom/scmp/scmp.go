// The scmp package implements the SCMP client protocol.
//
// # Introduction
//
// The Shadow Copy Management Protocol is used to programmatically enumerate shadow
// copies and configure shadow copy storage on remote machines. The protocol uses a
// set of Distributed Component Object Model (DCOM) interfaces to query shadow copies
// and manage shadow copy storage on a remote machine.
//
// This specification describes storage concepts, including volume storage concepts,
// in the Windows operating system. Although this specification outlines some basic
// storage concepts, it assumes that the reader has familiarity with these technologies.
// For background information about storage, disk, and volume concepts, see [MSDN-STC]
// and [MSDN-VOLMAN].
//
// This protocol documentation is intended for use together with publicly available
// standard specifications, networking programming art, and Microsoft distributed systems
// concepts. It assumes that the reader is either familiar with this material or has
// immediate access to it.
//
// A protocol specification does not require the use of Microsoft programming tools
// or programming environments for a Licensee to develop an implementation. Licensees
// who have access to Microsoft programming tools and environments are free to take
// advantage of them.
//
// # Overview
//
// The Shadow Copy Management Protocol provides a mechanism for remote configuration
// of shadow copies. Through the Shadow Copy Management Protocol, a client performs
// operations to enumerate shadow copies and configure the storage size and location
// that are used to maintain the shadow copies on the server.
//
// The Shadow Copy Management Protocol is expressed as a set of DCOM interfaces. The
// server end of the protocol implements support for the DCOM interfaces to manage shadow
// copy configuration objects. The client end of the protocol invokes method calls on
// the interfaces to perform shadow copy configuration tasks on the server.<1> Specifically,
// the protocol is used for the following purposes:
//
// * Enumerating the volumes ( 726a3dc9-844e-44c8-9527-236c3bd52dff#gt_9a876829-33a1-4f0b-8b81-8552b7e5561c
// ) on the server that can be shadow copied.
//
// * Enumerating the shadow copies that are currently available on the server and that
// are point-in-time copies of a specified original volume ( 726a3dc9-844e-44c8-9527-236c3bd52dff#gt_57484dae-5eef-4485-bfe4-db22b9cd90d6
// ).
//
// * Enumerating the volumes on the server that can be used as shadow copy storage (
// 726a3dc9-844e-44c8-9527-236c3bd52dff#gt_d85fc09f-c375-4e90-952d-7c95a8e244dd ).
//
// * Creating, modifying, enumerating, and deleting the shadow copy storage association
// ( 726a3dc9-844e-44c8-9527-236c3bd52dff#gt_34a368ce-08be-44b8-8d15-cfa0d4ac176e )
// objects that define the location and size of shadow copy storage for specific original
// volumes.
//
// * Querying all the shadow copy storage association objects on the server that provide
// shadow copy storage for a specified original volume.
//
// * Querying all the shadow copy storage association objects on the server that are
// located on a specified shadow copy storage volume ( 726a3dc9-844e-44c8-9527-236c3bd52dff#gt_9abab753-55aa-4455-bb21-7b51a32654fa
// ).
package scmp

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
	GoPackage = "dcom/scmp"
)

// ObjectType type represents VSS_OBJECT_TYPE RPC enumeration.
//
// The VSS_OBJECT_TYPE enumeration defines the types of objects that can be queried
// by the IVssEnumObject interface.
type ObjectType uint32

var (
	// VSS_OBJECT_UNKNOWN:  The object is of an unknown type of shadow copy.
	ObjectTypeUnknown ObjectType = 0
	// VSS_OBJECT_NONE:  This value MUST NOT be used and MUST be ignored upon receipt.
	ObjectTypeNone ObjectType = 1
	// VSS_OBJECT_SNAPSHOT_SET:  The object is a shadow copy set.
	ObjectTypeSnapshotSet ObjectType = 2
	// VSS_OBJECT_SNAPSHOT:  The object is a shadow copy.
	ObjectTypeSnapshot ObjectType = 3
	// VSS_OBJECT_PROVIDER:  This value is not used by the Shadow Copy Management Protocol
	// and MUST NOT be referenced. It MUST be ignored on receipt.
	ObjectTypeProvider ObjectType = 4
	// VSS_OBJECT_TYPE_COUNT:  This value is the number of VSS_OBJECT_TYPE values in the
	// enumeration.
	ObjectTypeCount ObjectType = 5
)

func (o ObjectType) String() string {
	switch o {
	case ObjectTypeUnknown:
		return "ObjectTypeUnknown"
	case ObjectTypeNone:
		return "ObjectTypeNone"
	case ObjectTypeSnapshotSet:
		return "ObjectTypeSnapshotSet"
	case ObjectTypeSnapshot:
		return "ObjectTypeSnapshot"
	case ObjectTypeProvider:
		return "ObjectTypeProvider"
	case ObjectTypeCount:
		return "ObjectTypeCount"
	}
	return "Invalid"
}

// SnapshotState type represents VSS_SNAPSHOT_STATE RPC enumeration.
//
// The VSS_SNAPSHOT_STATE enumeration defines the set of valid states of a shadow copy.
type SnapshotState uint32

var (
	// VSS_SS_UNKNOWN:  The shadow copy state is unknown. This is a restricted shadow copy
	// state. Shadow copies that are managed with this protocol MUST NOT appear in this
	// state.
	SnapshotStateUnknown SnapshotState = 0
	// VSS_SS_CREATED:  The shadow copy is created.
	SnapshotStateCreated SnapshotState = 12
)

func (o SnapshotState) String() string {
	switch o {
	case SnapshotStateUnknown:
		return "SnapshotStateUnknown"
	case SnapshotStateCreated:
		return "SnapshotStateCreated"
	}
	return "Invalid"
}

// VolumeSnapshotAttributes type represents VSS_VOLUME_SNAPSHOT_ATTRIBUTES RPC enumeration.
//
// The VSS_VOLUME_SNAPSHOT_ATTRIBUTES enumeration defines the set of valid attribute
// flags for a shadow copy.
type VolumeSnapshotAttributes uint32

var (
	// VSS_VOLSNAP_ATTR_PERSISTENT:  The shadow copy persists on the system despite rebooting
	// the machine.
	VolumeSnapshotAttributesPersistent VolumeSnapshotAttributes = 1
	// VSS_VOLSNAP_ATTR_NO_AUTORECOVERY:  The shadow copy is created as read-only. Applications
	// are not provided an opportunity to modify its contents.
	VolumeSnapshotAttributesNoAutorecovery VolumeSnapshotAttributes = 2
	// VSS_VOLSNAP_ATTR_CLIENT_ACCESSIBLE:  The shadow copy is of a specific type that
	// can be exposed remotely through the SMB Protocol [MS-SMB].
	VolumeSnapshotAttributesClientAccessible VolumeSnapshotAttributes = 4
	// VSS_VOLSNAP_ATTR_NO_AUTO_RELEASE:  The shadow copy is not deleted after the client
	// releases all references to the local interface that is used to create the shadow
	// copy.
	VolumeSnapshotAttributesNoAutoRelease VolumeSnapshotAttributes = 8
	// VSS_VOLSNAP_ATTR_NO_WRITERS:  The shadow copy is created without any application-specific
	// participation.
	VolumeSnapshotAttributesNoWriters VolumeSnapshotAttributes = 16
)

func (o VolumeSnapshotAttributes) String() string {
	switch o {
	case VolumeSnapshotAttributesPersistent:
		return "VolumeSnapshotAttributesPersistent"
	case VolumeSnapshotAttributesNoAutorecovery:
		return "VolumeSnapshotAttributesNoAutorecovery"
	case VolumeSnapshotAttributesClientAccessible:
		return "VolumeSnapshotAttributesClientAccessible"
	case VolumeSnapshotAttributesNoAutoRelease:
		return "VolumeSnapshotAttributesNoAutoRelease"
	case VolumeSnapshotAttributesNoWriters:
		return "VolumeSnapshotAttributesNoWriters"
	}
	return "Invalid"
}

// ManagementObjectType type represents VSS_MGMT_OBJECT_TYPE RPC enumeration.
//
// The VSS_MGMT_OBJECT_TYPE enumeration defines the types of objects that can be queried
// by the IVssEnumMgmtType interface.
type ManagementObjectType uint32

var (
	// VSS_MGMT_OBJECT_UNKNOWN:  The object is of an unknown type.
	ManagementObjectTypeUnknown ManagementObjectType = 0
	// VSS_MGMT_OBJECT_VOLUME:  The object is an original volume.
	ManagementObjectTypeVolume ManagementObjectType = 1
	// VSS_MGMT_OBJECT_DIFF_VOLUME:  The object is a shadow copy storage volume.
	ManagementObjectTypeDiffVolume ManagementObjectType = 2
	// VSS_MGMT_OBJECT_DIFF_AREA:  The object is shadow copy storage.
	ManagementObjectTypeDiffArea ManagementObjectType = 3
)

func (o ManagementObjectType) String() string {
	switch o {
	case ManagementObjectTypeUnknown:
		return "ManagementObjectTypeUnknown"
	case ManagementObjectTypeVolume:
		return "ManagementObjectTypeVolume"
	case ManagementObjectTypeDiffVolume:
		return "ManagementObjectTypeDiffVolume"
	case ManagementObjectTypeDiffArea:
		return "ManagementObjectTypeDiffArea"
	}
	return "Invalid"
}

// ProviderType type represents VSS_PROVIDER_TYPE RPC enumeration.
//
// The VSS_PROVIDER_TYPE enumeration defines the set of valid shadow copy provider types.
// This enumeration is not used by the Shadow Copy Management Protocol; it MUST NOT
// be referenced and MUST be ignored on receipt.
type ProviderType uint32

var (
	// VSS_PROV_UNKNOWN:  The shadow copy provider type is unknown.
	ProviderTypeUnknown ProviderType = 0
)

func (o ProviderType) String() string {
	switch o {
	case ProviderTypeUnknown:
		return "ProviderTypeUnknown"
	}
	return "Invalid"
}

// ID structure represents VSS_ID RPC structure.
type ID dtyp.GUID

func (o *ID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// SnapshotProperty structure represents VSS_SNAPSHOT_PROP RPC structure.
//
// The VSS_SNAPSHOT_PROP structure provides information about a shadow copy object.
type SnapshotProperty struct {
	// m_SnapshotId:  The VSS_ID (section 2.2.1.1) that identifies this shadow copy object.
	SnapshotID *ID `idl:"name:m_SnapshotId" json:"snapshot_id"`
	// m_SnapshotSetId:  The VSS_ID that identifies the shadow copy set of which this shadow
	// copy object is a member. All shadow copy objects in the same snapshot set MUST have
	// the same value for m_SnapshotSetId.
	SnapshotSetID *ID `idl:"name:m_SnapshotSetId" json:"snapshot_set_id"`
	// m_lSnapshotsCount:  The number of shadow copies in the shadow copy set when it was
	// originally created. It is possible that individual shadow copies that make up the
	// shadow copy set are deleted so that, at any time, it is possible that the number
	// of shadow copies currently in the snapshot set is less than m_lSnapshotCount.
	SnapshotsCount int32 `idl:"name:m_lSnapshotsCount" json:"snapshots_count"`
	// m_pwszSnapshotDeviceObject:  The null-terminated character string that contains the
	// name of the volume device for the shadow copy volume object on the server.<5>
	SnapshotDeviceObject string `idl:"name:m_pwszSnapshotDeviceObject" json:"snapshot_device_object"`
	// m_pwszOriginalVolumeName:  The null-terminated character string that contains the
	// volume mount name of the volume from which a shadow copy was obtained in order to
	// generate this shadow copy object.
	OriginalVolumeName string `idl:"name:m_pwszOriginalVolumeName" json:"original_volume_name"`
	// m_pwszOriginatingMachine:  The null-terminated character string that contains the
	// name of the machine that hosts the original volume. The server MUST populate this
	// string with the fully qualified domain name (FQDN) of the server machine. For this
	// protocol, the value of m_pwszOriginatingMachine and m_pwszServiceMachine MUST be
	// the same.
	OriginatingMachine string `idl:"name:m_pwszOriginatingMachine" json:"originating_machine"`
	// m_pwszServiceMachine:  The null-terminated character string that contains the name
	// of the machine on which the shadow copy was created. The server MUST populate this
	// string with the FQDN of the server machine. For this protocol, the value of m_pwszOriginatingMachine
	// and m_pwszServiceMachine MUST be the same.
	ServiceMachine string `idl:"name:m_pwszServiceMachine" json:"service_machine"`
	// m_pwszExposedName:  The null-terminated character string that contains the drive
	// letter, mount point, or SMB share name if the shadow copy is exposed on the server.
	// For this protocol, the server MUST set this value to NULL.
	ExposedName string `idl:"name:m_pwszExposedName" json:"exposed_name"`
	// m_pwszExposedPath:  The null-terminated character string that contains the full,
	// root-relative path to a folder on the shadow copy that is to be exposed as an SMB
	// share. For this protocol, the server MUST set this value to NULL.
	ExposedPath string `idl:"name:m_pwszExposedPath" json:"exposed_path"`
	// m_ProviderId:  The VSS_ID of the VSS provider that was used to create the shadow
	// copy.
	ProviderID *ID `idl:"name:m_ProviderId" json:"provider_id"`
	// m_lSnapshotAttributes:  The attributes of the shadow copy. The value of this LONG
	// value is a combination of the values that are defined in VSS_VOLUME_SNAPSHOT_ATTRIBUTES.
	SnapshotAttributes int32 `idl:"name:m_lSnapshotAttributes" json:"snapshot_attributes"`
	// m_tsCreationTimestamp:  The time stamp that defines when the shadow copy was created.
	CreationTimestamp int64 `idl:"name:m_tsCreationTimestamp" json:"creation_timestamp"`
	// m_eStatus:  A value from the VSS_SNAPSHOT_STATE enumeration (section 2.2.2.4) that
	// defines the state of the snapshot. For this protocol, the value of m_eStatus MUST
	// be VSS_SS_CREATED.
	Status SnapshotState `idl:"name:m_eStatus" json:"status"`
}

func (o *SnapshotProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SnapshotProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.SnapshotID != nil {
		if err := o.SnapshotID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SnapshotSetID != nil {
		if err := o.SnapshotSetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SnapshotsCount); err != nil {
		return err
	}
	if o.SnapshotDeviceObject != "" {
		_ptr_m_pwszSnapshotDeviceObject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.SnapshotDeviceObject); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.SnapshotDeviceObject, _ptr_m_pwszSnapshotDeviceObject); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OriginalVolumeName != "" {
		_ptr_m_pwszOriginalVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.OriginalVolumeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.OriginalVolumeName, _ptr_m_pwszOriginalVolumeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OriginatingMachine != "" {
		_ptr_m_pwszOriginatingMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.OriginatingMachine); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.OriginatingMachine, _ptr_m_pwszOriginatingMachine); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServiceMachine != "" {
		_ptr_m_pwszServiceMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ServiceMachine); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServiceMachine, _ptr_m_pwszServiceMachine); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ExposedName != "" {
		_ptr_m_pwszExposedName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ExposedName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ExposedName, _ptr_m_pwszExposedName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ExposedPath != "" {
		_ptr_m_pwszExposedPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ExposedPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ExposedPath, _ptr_m_pwszExposedPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProviderID != nil {
		if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SnapshotAttributes); err != nil {
		return err
	}
	if err := w.WriteData(o.CreationTimestamp); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.Status)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *SnapshotProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.SnapshotID == nil {
		o.SnapshotID = &ID{}
	}
	if err := o.SnapshotID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SnapshotSetID == nil {
		o.SnapshotSetID = &ID{}
	}
	if err := o.SnapshotSetID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SnapshotsCount); err != nil {
		return err
	}
	_ptr_m_pwszSnapshotDeviceObject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.SnapshotDeviceObject); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszSnapshotDeviceObject := func(ptr interface{}) { o.SnapshotDeviceObject = *ptr.(*string) }
	if err := w.ReadPointer(&o.SnapshotDeviceObject, _s_m_pwszSnapshotDeviceObject, _ptr_m_pwszSnapshotDeviceObject); err != nil {
		return err
	}
	_ptr_m_pwszOriginalVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.OriginalVolumeName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszOriginalVolumeName := func(ptr interface{}) { o.OriginalVolumeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.OriginalVolumeName, _s_m_pwszOriginalVolumeName, _ptr_m_pwszOriginalVolumeName); err != nil {
		return err
	}
	_ptr_m_pwszOriginatingMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.OriginatingMachine); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszOriginatingMachine := func(ptr interface{}) { o.OriginatingMachine = *ptr.(*string) }
	if err := w.ReadPointer(&o.OriginatingMachine, _s_m_pwszOriginatingMachine, _ptr_m_pwszOriginatingMachine); err != nil {
		return err
	}
	_ptr_m_pwszServiceMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServiceMachine); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszServiceMachine := func(ptr interface{}) { o.ServiceMachine = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServiceMachine, _s_m_pwszServiceMachine, _ptr_m_pwszServiceMachine); err != nil {
		return err
	}
	_ptr_m_pwszExposedName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ExposedName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszExposedName := func(ptr interface{}) { o.ExposedName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ExposedName, _s_m_pwszExposedName, _ptr_m_pwszExposedName); err != nil {
		return err
	}
	_ptr_m_pwszExposedPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ExposedPath); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszExposedPath := func(ptr interface{}) { o.ExposedPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.ExposedPath, _s_m_pwszExposedPath, _ptr_m_pwszExposedPath); err != nil {
		return err
	}
	if o.ProviderID == nil {
		o.ProviderID = &ID{}
	}
	if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SnapshotAttributes); err != nil {
		return err
	}
	if err := w.ReadData(&o.CreationTimestamp); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Status)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ProviderProperty structure represents VSS_PROVIDER_PROP RPC structure.
//
// The VSS_PROVIDER_PROP structure provides information about a shadow copy provider.
// This structure is not used by this protocol. It MUST NOT be referenced and MUST be
// ignored on receipt.
type ProviderProperty struct {
	ProviderID        *ID           `idl:"name:m_ProviderId" json:"provider_id"`
	ProviderName      string        `idl:"name:m_pwszProviderName" json:"provider_name"`
	ProviderType      ProviderType  `idl:"name:m_eProviderType" json:"provider_type"`
	ProviderVersion   string        `idl:"name:m_pwszProviderVersion" json:"provider_version"`
	ProviderVersionID *ID           `idl:"name:m_ProviderVersionId" json:"provider_version_id"`
	ClassID           *dcom.ClassID `idl:"name:m_ClassId" json:"class_id"`
}

func (o *ProviderProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ProviderProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ProviderID != nil {
		if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProviderName != "" {
		_ptr_m_pwszProviderName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProviderName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderName, _ptr_m_pwszProviderName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint32(o.ProviderType)); err != nil {
		return err
	}
	if o.ProviderVersion != "" {
		_ptr_m_pwszProviderVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ProviderVersion); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ProviderVersion, _ptr_m_pwszProviderVersion); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ProviderVersionID != nil {
		if err := o.ProviderVersionID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dcom.ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ProviderProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ProviderID == nil {
		o.ProviderID = &ID{}
	}
	if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_m_pwszProviderName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProviderName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszProviderName := func(ptr interface{}) { o.ProviderName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProviderName, _s_m_pwszProviderName, _ptr_m_pwszProviderName); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.ProviderType)); err != nil {
		return err
	}
	_ptr_m_pwszProviderVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ProviderVersion); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszProviderVersion := func(ptr interface{}) { o.ProviderVersion = *ptr.(*string) }
	if err := w.ReadPointer(&o.ProviderVersion, _s_m_pwszProviderVersion, _ptr_m_pwszProviderVersion); err != nil {
		return err
	}
	if o.ProviderVersionID == nil {
		o.ProviderVersionID = &ID{}
	}
	if err := o.ProviderVersionID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &dcom.ClassID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ObjectUnion structure represents VSS_OBJECT_UNION RPC union.
//
// The VSS_OBJECT_UNION defines the union of object types that can be defined by the
// VSS_OBJECT_PROP structure (section 2.2.3.2).
type ObjectUnion struct {
	// Types that are assignable to Value
	//
	// *ObjectUnion_Snap
	// *ObjectUnion_Provider
	Value is_ObjectUnion `json:"value"`
}

func (o *ObjectUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ObjectUnion_Snap:
		if value != nil {
			return value.Snap
		}
	case *ObjectUnion_Provider:
		if value != nil {
			return value.Provider
		}
	}
	return nil
}

type is_ObjectUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ObjectUnion()
}

func (o *ObjectUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ObjectUnion_Snap:
		return uint32(3)
	case *ObjectUnion_Provider:
		return uint32(4)
	}
	return uint32(0)
}

func (o *ObjectUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint32(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(3):
		_o, _ := o.Value.(*ObjectUnion_Snap)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectUnion_Snap{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*ObjectUnion_Provider)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectUnion_Provider{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ObjectUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint32)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(3):
		o.Value = &ObjectUnion_Snap{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &ObjectUnion_Provider{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ObjectUnion_Snap structure represents VSS_OBJECT_UNION RPC union arm.
//
// It has following labels: 3
type ObjectUnion_Snap struct {
	// Snap:  The structure specifies a shadow copy object as a VSS_SNAPSHOT_PROP structure
	// (section 2.2.3.3).
	Snap *SnapshotProperty `idl:"name:Snap" json:"snap"`
}

func (*ObjectUnion_Snap) is_ObjectUnion() {}

func (o *ObjectUnion_Snap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Snap != nil {
		if err := o.Snap.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SnapshotProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectUnion_Snap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Snap == nil {
		o.Snap = &SnapshotProperty{}
	}
	if err := o.Snap.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectUnion_Provider structure represents VSS_OBJECT_UNION RPC union arm.
//
// It has following labels: 4
type ObjectUnion_Provider struct {
	// Prov:  The structure specifies a VSS provider object. The Shadow Copy Management
	// Protocol is not used to manage VSS provider objects; therefore, this member MUST
	// NOT be referenced and MUST be ignored on receipt.
	Provider *ProviderProperty `idl:"name:Prov" json:"provider"`
}

func (*ObjectUnion_Provider) is_ObjectUnion() {}

func (o *ObjectUnion_Provider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Provider != nil {
		if err := o.Provider.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ProviderProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectUnion_Provider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Provider == nil {
		o.Provider = &ProviderProperty{}
	}
	if err := o.Provider.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectProperty structure represents VSS_OBJECT_PROP RPC structure.
//
// The VSS_OBJECT_PROP structure specifies the union of object types that can be enumerated
// by the IVssEnumObject interface.
type ObjectProperty struct {
	// Type:  A value defined in the VSS_OBJECT_TYPE enumeration (section 2.2.2.1) that
	// specifies the type of object that is contained in the Obj union structure.
	Type ObjectType `idl:"name:Type" json:"type"`
	// Obj:  A VSS_OBJECT_UNION structure (section 2.2.3.1).
	Object *ObjectUnion `idl:"name:Obj;switch_is:Type" json:"object"`
}

func (o *ObjectProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.Type)); err != nil {
		return err
	}
	_swObject := uint32(o.Type)
	if o.Object != nil {
		if err := o.Object.MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	} else {
		if err := (&ObjectUnion{}).MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Type)); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &ObjectUnion{}
	}
	_swObject := uint32(o.Type)
	if err := o.Object.UnmarshalUnionNDR(ctx, w, _swObject); err != nil {
		return err
	}
	return nil
}

// VolumeProperty structure represents VSS_VOLUME_PROP RPC structure.
//
// The VSS_VOLUME_PROP structure defines properties of a volume.
type VolumeProperty struct {
	// m_pwszVolumeName:  A null-terminated character string that contains the volume mount
	// name of the volume.
	VolumeName string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	// m_pwszVolumeDisplayName:  A null-terminated character string that contains a mount
	// point path for the volume. If the volume has no mount points, the string MUST be
	// equal to m_pwszVolumeName.
	VolumeDisplayName string `idl:"name:m_pwszVolumeDisplayName" json:"volume_display_name"`
}

func (o *VolumeProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VolumeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.VolumeName != "" {
		_ptr_m_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeName, _ptr_m_pwszVolumeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VolumeDisplayName != "" {
		_ptr_m_pwszVolumeDisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VolumeDisplayName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeDisplayName, _ptr_m_pwszVolumeDisplayName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VolumeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_m_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VolumeName, _s_m_pwszVolumeName, _ptr_m_pwszVolumeName); err != nil {
		return err
	}
	_ptr_m_pwszVolumeDisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeDisplayName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszVolumeDisplayName := func(ptr interface{}) { o.VolumeDisplayName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VolumeDisplayName, _s_m_pwszVolumeDisplayName, _ptr_m_pwszVolumeDisplayName); err != nil {
		return err
	}
	return nil
}

// DiffVolumeProperty structure represents VSS_DIFF_VOLUME_PROP RPC structure.
//
// The VSS_DIFF_VOLUME_PROP structure defines the properties of a shadow copy storage
// volume.
type DiffVolumeProperty struct {
	// m_pwszVolumeName:  A null-terminated character string that contains the volume mount
	// name of the volume.
	VolumeName string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	// m_pwszVolumeDisplayName:  A null-terminated character string that contains one of
	// the mount point paths for the volume. If the volume has no mount points, the string
	// MUST be equal to m_pwszVolumeName.
	VolumeDisplayName string `idl:"name:m_pwszVolumeDisplayName" json:"volume_display_name"`
	// m_llVolumeFreeSpace:  The amount of free space, in BYTEs, on the volume.
	VolumeFreeSpace int64 `idl:"name:m_llVolumeFreeSpace" json:"volume_free_space"`
	// m_llVolumeTotalSpace:  The total size, in BYTEs, of the volume.
	VolumeTotalSpace int64 `idl:"name:m_llVolumeTotalSpace" json:"volume_total_space"`
}

func (o *DiffVolumeProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DiffVolumeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.VolumeName != "" {
		_ptr_m_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeName, _ptr_m_pwszVolumeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VolumeDisplayName != "" {
		_ptr_m_pwszVolumeDisplayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VolumeDisplayName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeDisplayName, _ptr_m_pwszVolumeDisplayName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VolumeFreeSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.VolumeTotalSpace); err != nil {
		return err
	}
	return nil
}
func (o *DiffVolumeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_m_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VolumeName, _s_m_pwszVolumeName, _ptr_m_pwszVolumeName); err != nil {
		return err
	}
	_ptr_m_pwszVolumeDisplayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeDisplayName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszVolumeDisplayName := func(ptr interface{}) { o.VolumeDisplayName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VolumeDisplayName, _s_m_pwszVolumeDisplayName, _ptr_m_pwszVolumeDisplayName); err != nil {
		return err
	}
	if err := w.ReadData(&o.VolumeFreeSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.VolumeTotalSpace); err != nil {
		return err
	}
	return nil
}

// DiffAreaProperty structure represents VSS_DIFF_AREA_PROP RPC structure.
//
// The VSS_DIFF_AREA_PROP structure defines a shadow copy storage association and the
// current sizes of the shadow copy storage.
type DiffAreaProperty struct {
	// m_pwszVolumeName:  A null-terminated character string that contains the volume mount
	// name of the original volume that is or will be shadow copied.
	VolumeName string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	// m_pwszDiffAreaVolumeName:  A null-terminated character string that contains the volume
	// mount name of the shadow copy storage volume where shadow copy differential data
	// will be located for the volume specified in m_pwszVolumeName.
	DiffAreaVolumeName string `idl:"name:m_pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	// m_llMaximumDiffSpace:  The maximum number of BYTEs that will be consumed on the shadow
	// copy storage volume to maintain shadow copies.
	MaximumDiffSpace int64 `idl:"name:m_llMaximumDiffSpace" json:"maximum_diff_space"`
	// m_llAllocatedDiffSpace:   The number of BYTEs currently allocated for shadow copy
	// storage space. This value MUST be less than or equal to m_llMaximumDiffSpace.
	AllocatedDiffSpace int64 `idl:"name:m_llAllocatedDiffSpace" json:"allocated_diff_space"`
	// m_llUsedDiffSpace:   The number of BYTEs currently in use on the shadow copy storage
	// volume to maintain shadow copies. This value MUST be less than or equal to m_llAllocatedDiffSpace.
	UsedDiffSpace int64 `idl:"name:m_llUsedDiffSpace" json:"used_diff_space"`
}

func (o *DiffAreaProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DiffAreaProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.VolumeName != "" {
		_ptr_m_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.VolumeName, _ptr_m_pwszVolumeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DiffAreaVolumeName != "" {
		_ptr_m_pwszDiffAreaVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DiffAreaVolumeName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DiffAreaVolumeName, _ptr_m_pwszDiffAreaVolumeName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MaximumDiffSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.AllocatedDiffSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.UsedDiffSpace); err != nil {
		return err
	}
	return nil
}
func (o *DiffAreaProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	_ptr_m_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.VolumeName, _s_m_pwszVolumeName, _ptr_m_pwszVolumeName); err != nil {
		return err
	}
	_ptr_m_pwszDiffAreaVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DiffAreaVolumeName); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwszDiffAreaVolumeName := func(ptr interface{}) { o.DiffAreaVolumeName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DiffAreaVolumeName, _s_m_pwszDiffAreaVolumeName, _ptr_m_pwszDiffAreaVolumeName); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumDiffSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.AllocatedDiffSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.UsedDiffSpace); err != nil {
		return err
	}
	return nil
}

// ManagementObjectUnion structure represents VSS_MGMT_OBJECT_UNION RPC union.
//
// The VSS_MGMT_OBJECT_UNION specifies the union of object types that can be defined
// by the VSS_MGMT_OBJECT_PROP structure (section 2.2.3.6).
type ManagementObjectUnion struct {
	// Types that are assignable to Value
	//
	// *ManagementObjectUnion_Volume
	// *ManagementObjectUnion_DiffVolume
	// *ManagementObjectUnion_DiffArea
	Value is_ManagementObjectUnion `json:"value"`
}

func (o *ManagementObjectUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ManagementObjectUnion_Volume:
		if value != nil {
			return value.Volume
		}
	case *ManagementObjectUnion_DiffVolume:
		if value != nil {
			return value.DiffVolume
		}
	case *ManagementObjectUnion_DiffArea:
		if value != nil {
			return value.DiffArea
		}
	}
	return nil
}

type is_ManagementObjectUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ManagementObjectUnion()
}

func (o *ManagementObjectUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ManagementObjectUnion_Volume:
		return uint32(1)
	case *ManagementObjectUnion_DiffVolume:
		return uint32(2)
	case *ManagementObjectUnion_DiffArea:
		return uint32(3)
	}
	return uint32(0)
}

func (o *ManagementObjectUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint32(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ManagementObjectUnion_Volume)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObjectUnion_Volume{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*ManagementObjectUnion_DiffVolume)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObjectUnion_DiffVolume{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*ManagementObjectUnion_DiffArea)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObjectUnion_DiffArea{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ManagementObjectUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint32)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ManagementObjectUnion_Volume{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &ManagementObjectUnion_DiffVolume{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &ManagementObjectUnion_DiffArea{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ManagementObjectUnion_Volume structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 1
type ManagementObjectUnion_Volume struct {
	// Vol:  The structure specifies an original volume object as a VSS_VOLUME_PROP structure
	// (section 2.2.3.7).
	Volume *VolumeProperty `idl:"name:Vol" json:"volume"`
}

func (*ManagementObjectUnion_Volume) is_ManagementObjectUnion() {}

func (o *ManagementObjectUnion_Volume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Volume != nil {
		if err := o.Volume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VolumeProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObjectUnion_Volume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Volume == nil {
		o.Volume = &VolumeProperty{}
	}
	if err := o.Volume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ManagementObjectUnion_DiffVolume structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 2
type ManagementObjectUnion_DiffVolume struct {
	// DiffVol:  The structure specifies a shadow copy storage volume as a VSS_DIFF_VOLUME_PROP
	// structure.
	DiffVolume *DiffVolumeProperty `idl:"name:DiffVol" json:"diff_volume"`
}

func (*ManagementObjectUnion_DiffVolume) is_ManagementObjectUnion() {}

func (o *ManagementObjectUnion_DiffVolume) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DiffVolume != nil {
		if err := o.DiffVolume.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DiffVolumeProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObjectUnion_DiffVolume) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiffVolume == nil {
		o.DiffVolume = &DiffVolumeProperty{}
	}
	if err := o.DiffVolume.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ManagementObjectUnion_DiffArea structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 3
type ManagementObjectUnion_DiffArea struct {
	// DiffArea:  The structure specifies a shadow copy storage object as a VSS_DIFF_AREA_PROP.
	DiffArea *DiffAreaProperty `idl:"name:DiffArea" json:"diff_area"`
}

func (*ManagementObjectUnion_DiffArea) is_ManagementObjectUnion() {}

func (o *ManagementObjectUnion_DiffArea) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DiffArea != nil {
		if err := o.DiffArea.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DiffAreaProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObjectUnion_DiffArea) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiffArea == nil {
		o.DiffArea = &DiffAreaProperty{}
	}
	if err := o.DiffArea.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ManagementObjectProperty structure represents VSS_MGMT_OBJECT_PROP RPC structure.
//
// The VSS_MGMT_OBJECT_PROP structure defines the union of object types that can be
// enumerated by the IVssEnumMgmtObject interface.
type ManagementObjectProperty struct {
	// Type:  A value that is defined in the VSS_MGMT_OBJECT_TYPE enumeration that specifies
	// the type of object that is contained in the Obj union structure.
	Type ManagementObjectType `idl:"name:Type" json:"type"`
	// Obj:  A VSS_MGMT_OBJECT_UNION structure.
	Object *ManagementObjectUnion `idl:"name:Obj;switch_is:Type" json:"object"`
}

func (o *ManagementObjectProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ManagementObjectProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.Type)); err != nil {
		return err
	}
	_swObject := uint32(o.Type)
	if o.Object != nil {
		if err := o.Object.MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	} else {
		if err := (&ManagementObjectUnion{}).MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObjectProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Type)); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &ManagementObjectUnion{}
	}
	_swObject := uint32(o.Type)
	if err := o.Object.UnmarshalUnionNDR(ctx, w, _swObject); err != nil {
		return err
	}
	return nil
}

// SnapshotManagement structure represents IVssSnapshotMgmt RPC structure.
type SnapshotManagement dcom.InterfacePointer

func (o *SnapshotManagement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *SnapshotManagement) xxx_PreparePayload(ctx context.Context) error {
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

func (o *SnapshotManagement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SnapshotManagement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SnapshotManagement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EnumManagementObject structure represents IVssEnumMgmtObject RPC structure.
type EnumManagementObject dcom.InterfacePointer

func (o *EnumManagementObject) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EnumManagementObject) xxx_PreparePayload(ctx context.Context) error {
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

func (o *EnumManagementObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumManagementObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EnumManagementObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EnumObject structure represents IVssEnumObject RPC structure.
type EnumObject dcom.InterfacePointer

func (o *EnumObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EnumObject) xxx_PreparePayload(ctx context.Context) error {
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

func (o *EnumObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EnumObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DifferentialSoftwareSnapshotManagement structure represents IVssDifferentialSoftwareSnapshotMgmt RPC structure.
type DifferentialSoftwareSnapshotManagement dcom.InterfacePointer

func (o *DifferentialSoftwareSnapshotManagement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DifferentialSoftwareSnapshotManagement) xxx_PreparePayload(ctx context.Context) error {
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

func (o *DifferentialSoftwareSnapshotManagement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DifferentialSoftwareSnapshotManagement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DifferentialSoftwareSnapshotManagement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
