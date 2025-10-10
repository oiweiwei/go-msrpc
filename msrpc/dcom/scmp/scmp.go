// The scmp package implements the SCMP client protocol.
//
// # Introduction
//
// # Overview
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
type ObjectType uint32

var (
	ObjectTypeUnknown     ObjectType = 0
	ObjectTypeNone        ObjectType = 1
	ObjectTypeSnapshotSet ObjectType = 2
	ObjectTypeSnapshot    ObjectType = 3
	ObjectTypeProvider    ObjectType = 4
	ObjectTypeCount       ObjectType = 5
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
type SnapshotState uint32

var (
	SnapshotStateUnknown SnapshotState = 0
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
type VolumeSnapshotAttributes uint32

var (
	VolumeSnapshotAttributesPersistent       VolumeSnapshotAttributes = 1
	VolumeSnapshotAttributesNoAutorecovery   VolumeSnapshotAttributes = 2
	VolumeSnapshotAttributesClientAccessible VolumeSnapshotAttributes = 4
	VolumeSnapshotAttributesNoAutoRelease    VolumeSnapshotAttributes = 8
	VolumeSnapshotAttributesNoWriters        VolumeSnapshotAttributes = 16
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
type ManagementObjectType uint32

var (
	ManagementObjectTypeUnknown    ManagementObjectType = 0
	ManagementObjectTypeVolume     ManagementObjectType = 1
	ManagementObjectTypeDiffVolume ManagementObjectType = 2
	ManagementObjectTypeDiffArea   ManagementObjectType = 3
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
type ProviderType uint32

var (
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
type SnapshotProperty struct {
	SnapshotID           *ID           `idl:"name:m_SnapshotId" json:"snapshot_id"`
	SnapshotSetID        *ID           `idl:"name:m_SnapshotSetId" json:"snapshot_set_id"`
	SnapshotsCount       int32         `idl:"name:m_lSnapshotsCount" json:"snapshots_count"`
	SnapshotDeviceObject string        `idl:"name:m_pwszSnapshotDeviceObject" json:"snapshot_device_object"`
	OriginalVolumeName   string        `idl:"name:m_pwszOriginalVolumeName" json:"original_volume_name"`
	OriginatingMachine   string        `idl:"name:m_pwszOriginatingMachine" json:"originating_machine"`
	ServiceMachine       string        `idl:"name:m_pwszServiceMachine" json:"service_machine"`
	ExposedName          string        `idl:"name:m_pwszExposedName" json:"exposed_name"`
	ExposedPath          string        `idl:"name:m_pwszExposedPath" json:"exposed_path"`
	ProviderID           *ID           `idl:"name:m_ProviderId" json:"provider_id"`
	SnapshotAttributes   int32         `idl:"name:m_lSnapshotAttributes" json:"snapshot_attributes"`
	CreationTimestamp    int64         `idl:"name:m_tsCreationTimestamp" json:"creation_timestamp"`
	Status               SnapshotState `idl:"name:m_eStatus" json:"status"`
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
type ObjectProperty struct {
	Type   ObjectType   `idl:"name:Type" json:"type"`
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
type VolumeProperty struct {
	VolumeName        string `idl:"name:m_pwszVolumeName" json:"volume_name"`
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
type DiffVolumeProperty struct {
	VolumeName        string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	VolumeDisplayName string `idl:"name:m_pwszVolumeDisplayName" json:"volume_display_name"`
	VolumeFreeSpace   int64  `idl:"name:m_llVolumeFreeSpace" json:"volume_free_space"`
	VolumeTotalSpace  int64  `idl:"name:m_llVolumeTotalSpace" json:"volume_total_space"`
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
type DiffAreaProperty struct {
	VolumeName         string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	DiffAreaVolumeName string `idl:"name:m_pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	MaximumDiffSpace   int64  `idl:"name:m_llMaximumDiffSpace" json:"maximum_diff_space"`
	AllocatedDiffSpace int64  `idl:"name:m_llAllocatedDiffSpace" json:"allocated_diff_space"`
	UsedDiffSpace      int64  `idl:"name:m_llUsedDiffSpace" json:"used_diff_space"`
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
type ManagementObjectProperty struct {
	Type   ManagementObjectType   `idl:"name:Type" json:"type"`
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
