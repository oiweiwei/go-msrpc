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

// VSSObjectType type represents VSS_OBJECT_TYPE RPC enumeration.
type VSSObjectType uint32

var (
	VSSObjectTypeUnknown     VSSObjectType = 0
	VSSObjectTypeNone        VSSObjectType = 1
	VSSObjectTypeSnapshotSet VSSObjectType = 2
	VSSObjectTypeSnapshot    VSSObjectType = 3
	VSSObjectTypeProvider    VSSObjectType = 4
	VSSObjectTypeCount       VSSObjectType = 5
)

func (o VSSObjectType) String() string {
	switch o {
	case VSSObjectTypeUnknown:
		return "VSSObjectTypeUnknown"
	case VSSObjectTypeNone:
		return "VSSObjectTypeNone"
	case VSSObjectTypeSnapshotSet:
		return "VSSObjectTypeSnapshotSet"
	case VSSObjectTypeSnapshot:
		return "VSSObjectTypeSnapshot"
	case VSSObjectTypeProvider:
		return "VSSObjectTypeProvider"
	case VSSObjectTypeCount:
		return "VSSObjectTypeCount"
	}
	return "Invalid"
}

// VSSSnapshotState type represents VSS_SNAPSHOT_STATE RPC enumeration.
type VSSSnapshotState uint32

var (
	VSSSnapshotStateSsUnknown VSSSnapshotState = 0
	VSSSnapshotStateSsCreated VSSSnapshotState = 12
)

func (o VSSSnapshotState) String() string {
	switch o {
	case VSSSnapshotStateSsUnknown:
		return "VSSSnapshotStateSsUnknown"
	case VSSSnapshotStateSsCreated:
		return "VSSSnapshotStateSsCreated"
	}
	return "Invalid"
}

// VSSVolumeSnapshotAttributes type represents VSS_VOLUME_SNAPSHOT_ATTRIBUTES RPC enumeration.
type VSSVolumeSnapshotAttributes uint32

var (
	VSSVolumeSnapshotAttributesVolsnapAttributePersistent       VSSVolumeSnapshotAttributes = 1
	VSSVolumeSnapshotAttributesVolsnapAttributeNoAutorecovery   VSSVolumeSnapshotAttributes = 2
	VSSVolumeSnapshotAttributesVolsnapAttributeClientAccessible VSSVolumeSnapshotAttributes = 4
	VSSVolumeSnapshotAttributesVolsnapAttributeNoAutoRelease    VSSVolumeSnapshotAttributes = 8
	VSSVolumeSnapshotAttributesVolsnapAttributeNoWriters        VSSVolumeSnapshotAttributes = 16
)

func (o VSSVolumeSnapshotAttributes) String() string {
	switch o {
	case VSSVolumeSnapshotAttributesVolsnapAttributePersistent:
		return "VSSVolumeSnapshotAttributesVolsnapAttributePersistent"
	case VSSVolumeSnapshotAttributesVolsnapAttributeNoAutorecovery:
		return "VSSVolumeSnapshotAttributesVolsnapAttributeNoAutorecovery"
	case VSSVolumeSnapshotAttributesVolsnapAttributeClientAccessible:
		return "VSSVolumeSnapshotAttributesVolsnapAttributeClientAccessible"
	case VSSVolumeSnapshotAttributesVolsnapAttributeNoAutoRelease:
		return "VSSVolumeSnapshotAttributesVolsnapAttributeNoAutoRelease"
	case VSSVolumeSnapshotAttributesVolsnapAttributeNoWriters:
		return "VSSVolumeSnapshotAttributesVolsnapAttributeNoWriters"
	}
	return "Invalid"
}

// VSSManagementObjectType type represents VSS_MGMT_OBJECT_TYPE RPC enumeration.
type VSSManagementObjectType uint32

var (
	VSSManagementObjectTypeUnknown    VSSManagementObjectType = 0
	VSSManagementObjectTypeVolume     VSSManagementObjectType = 1
	VSSManagementObjectTypeDiffVolume VSSManagementObjectType = 2
	VSSManagementObjectTypeDiffArea   VSSManagementObjectType = 3
)

func (o VSSManagementObjectType) String() string {
	switch o {
	case VSSManagementObjectTypeUnknown:
		return "VSSManagementObjectTypeUnknown"
	case VSSManagementObjectTypeVolume:
		return "VSSManagementObjectTypeVolume"
	case VSSManagementObjectTypeDiffVolume:
		return "VSSManagementObjectTypeDiffVolume"
	case VSSManagementObjectTypeDiffArea:
		return "VSSManagementObjectTypeDiffArea"
	}
	return "Invalid"
}

// VSSProviderType type represents VSS_PROVIDER_TYPE RPC enumeration.
type VSSProviderType uint32

var (
	VSSProviderTypeUnknown VSSProviderType = 0
)

func (o VSSProviderType) String() string {
	switch o {
	case VSSProviderTypeUnknown:
		return "VSSProviderTypeUnknown"
	}
	return "Invalid"
}

// VSSID structure represents VSS_ID RPC structure.
type VSSID dtyp.GUID

func (o *VSSID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *VSSID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VSSSnapshotProperty structure represents VSS_SNAPSHOT_PROP RPC structure.
type VSSSnapshotProperty struct {
	SnapshotID           *VSSID           `idl:"name:m_SnapshotId" json:"snapshot_id"`
	SnapshotSetID        *VSSID           `idl:"name:m_SnapshotSetId" json:"snapshot_set_id"`
	SnapshotsCount       int32            `idl:"name:m_lSnapshotsCount" json:"snapshots_count"`
	SnapshotDeviceObject string           `idl:"name:m_pwszSnapshotDeviceObject" json:"snapshot_device_object"`
	OriginalVolumeName   string           `idl:"name:m_pwszOriginalVolumeName" json:"original_volume_name"`
	OriginatingMachine   string           `idl:"name:m_pwszOriginatingMachine" json:"originating_machine"`
	ServiceMachine       string           `idl:"name:m_pwszServiceMachine" json:"service_machine"`
	ExposedName          string           `idl:"name:m_pwszExposedName" json:"exposed_name"`
	ExposedPath          string           `idl:"name:m_pwszExposedPath" json:"exposed_path"`
	ProviderID           *VSSID           `idl:"name:m_ProviderId" json:"provider_id"`
	SnapshotAttributes   int32            `idl:"name:m_lSnapshotAttributes" json:"snapshot_attributes"`
	CreationTimestamp    int64            `idl:"name:m_tsCreationTimestamp" json:"creation_timestamp"`
	Status               VSSSnapshotState `idl:"name:m_eStatus" json:"status"`
}

func (o *VSSSnapshotProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSSnapshotProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&VSSID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SnapshotSetID != nil {
		if err := o.SnapshotSetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&VSSID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *VSSSnapshotProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.SnapshotID == nil {
		o.SnapshotID = &VSSID{}
	}
	if err := o.SnapshotID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SnapshotSetID == nil {
		o.SnapshotSetID = &VSSID{}
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
		o.ProviderID = &VSSID{}
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

// VSSProviderProperty structure represents VSS_PROVIDER_PROP RPC structure.
type VSSProviderProperty struct {
	ProviderID        *VSSID          `idl:"name:m_ProviderId" json:"provider_id"`
	ProviderName      string          `idl:"name:m_pwszProviderName" json:"provider_name"`
	ProviderType      VSSProviderType `idl:"name:m_eProviderType" json:"provider_type"`
	ProviderVersion   string          `idl:"name:m_pwszProviderVersion" json:"provider_version"`
	ProviderVersionID *VSSID          `idl:"name:m_ProviderVersionId" json:"provider_version_id"`
	ClassID           *dcom.ClassID   `idl:"name:m_ClassId" json:"class_id"`
}

func (o *VSSProviderProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSProviderProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&VSSID{}).MarshalNDR(ctx, w); err != nil {
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
		if err := (&VSSID{}).MarshalNDR(ctx, w); err != nil {
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
func (o *VSSProviderProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ProviderID == nil {
		o.ProviderID = &VSSID{}
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
		o.ProviderVersionID = &VSSID{}
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

// VSSObjectUnion structure represents VSS_OBJECT_UNION RPC union.
type VSSObjectUnion struct {
	// Types that are assignable to Value
	//
	// *VSSObjectUnion_Snap
	// *VSSObjectUnion_Prov
	Value is_VSSObjectUnion `json:"value"`
}

func (o *VSSObjectUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *VSSObjectUnion_Snap:
		if value != nil {
			return value.Snap
		}
	case *VSSObjectUnion_Prov:
		if value != nil {
			return value.Prov
		}
	}
	return nil
}

type is_VSSObjectUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_VSSObjectUnion()
}

func (o *VSSObjectUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *VSSObjectUnion_Snap:
		return uint32(3)
	case *VSSObjectUnion_Prov:
		return uint32(4)
	}
	return uint32(0)
}

func (o *VSSObjectUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*VSSObjectUnion_Snap)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VSSObjectUnion_Snap{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*VSSObjectUnion_Prov)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VSSObjectUnion_Prov{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *VSSObjectUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &VSSObjectUnion_Snap{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &VSSObjectUnion_Prov{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// VSSObjectUnion_Snap structure represents VSS_OBJECT_UNION RPC union arm.
//
// It has following labels: 3
type VSSObjectUnion_Snap struct {
	Snap *VSSSnapshotProperty `idl:"name:Snap" json:"snap"`
}

func (*VSSObjectUnion_Snap) is_VSSObjectUnion() {}

func (o *VSSObjectUnion_Snap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Snap != nil {
		if err := o.Snap.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSSnapshotProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSObjectUnion_Snap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Snap == nil {
		o.Snap = &VSSSnapshotProperty{}
	}
	if err := o.Snap.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VSSObjectUnion_Prov structure represents VSS_OBJECT_UNION RPC union arm.
//
// It has following labels: 4
type VSSObjectUnion_Prov struct {
	Prov *VSSProviderProperty `idl:"name:Prov" json:"prov"`
}

func (*VSSObjectUnion_Prov) is_VSSObjectUnion() {}

func (o *VSSObjectUnion_Prov) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Prov != nil {
		if err := o.Prov.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSProviderProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSObjectUnion_Prov) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Prov == nil {
		o.Prov = &VSSProviderProperty{}
	}
	if err := o.Prov.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VSSObjectProperty structure represents VSS_OBJECT_PROP RPC structure.
type VSSObjectProperty struct {
	Type   VSSObjectType   `idl:"name:Type" json:"type"`
	Object *VSSObjectUnion `idl:"name:Obj;switch_is:Type" json:"object"`
}

func (o *VSSObjectProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSObjectProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&VSSObjectUnion{}).MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSObjectProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Type)); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &VSSObjectUnion{}
	}
	_swObject := uint32(o.Type)
	if err := o.Object.UnmarshalUnionNDR(ctx, w, _swObject); err != nil {
		return err
	}
	return nil
}

// VSSVolumeProperty structure represents VSS_VOLUME_PROP RPC structure.
type VSSVolumeProperty struct {
	VolumeName        string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	VolumeDisplayName string `idl:"name:m_pwszVolumeDisplayName" json:"volume_display_name"`
}

func (o *VSSVolumeProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSVolumeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSVolumeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VSSDiffVolumeProperty structure represents VSS_DIFF_VOLUME_PROP RPC structure.
type VSSDiffVolumeProperty struct {
	VolumeName         string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	VolumeDisplayName  string `idl:"name:m_pwszVolumeDisplayName" json:"volume_display_name"`
	LlVolumeFreeSpace  int64  `idl:"name:m_llVolumeFreeSpace" json:"ll_volume_free_space"`
	LlVolumeTotalSpace int64  `idl:"name:m_llVolumeTotalSpace" json:"ll_volume_total_space"`
}

func (o *VSSDiffVolumeProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSDiffVolumeProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.LlVolumeFreeSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.LlVolumeTotalSpace); err != nil {
		return err
	}
	return nil
}
func (o *VSSDiffVolumeProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.LlVolumeFreeSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.LlVolumeTotalSpace); err != nil {
		return err
	}
	return nil
}

// VSSDiffAreaProperty structure represents VSS_DIFF_AREA_PROP RPC structure.
type VSSDiffAreaProperty struct {
	VolumeName           string `idl:"name:m_pwszVolumeName" json:"volume_name"`
	DiffAreaVolumeName   string `idl:"name:m_pwszDiffAreaVolumeName" json:"diff_area_volume_name"`
	LlMaximumDiffSpace   int64  `idl:"name:m_llMaximumDiffSpace" json:"ll_maximum_diff_space"`
	LlAllocatedDiffSpace int64  `idl:"name:m_llAllocatedDiffSpace" json:"ll_allocated_diff_space"`
	LlUsedDiffSpace      int64  `idl:"name:m_llUsedDiffSpace" json:"ll_used_diff_space"`
}

func (o *VSSDiffAreaProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSDiffAreaProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.LlMaximumDiffSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.LlAllocatedDiffSpace); err != nil {
		return err
	}
	if err := w.WriteData(o.LlUsedDiffSpace); err != nil {
		return err
	}
	return nil
}
func (o *VSSDiffAreaProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.LlMaximumDiffSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.LlAllocatedDiffSpace); err != nil {
		return err
	}
	if err := w.ReadData(&o.LlUsedDiffSpace); err != nil {
		return err
	}
	return nil
}

// VSSManagementObjectUnion structure represents VSS_MGMT_OBJECT_UNION RPC union.
type VSSManagementObjectUnion struct {
	// Types that are assignable to Value
	//
	// *VSSManagementObjectUnion_Vol
	// *VSSManagementObjectUnion_DiffVol
	// *VSSManagementObjectUnion_DiffArea
	Value is_VSSManagementObjectUnion `json:"value"`
}

func (o *VSSManagementObjectUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *VSSManagementObjectUnion_Vol:
		if value != nil {
			return value.Vol
		}
	case *VSSManagementObjectUnion_DiffVol:
		if value != nil {
			return value.DiffVol
		}
	case *VSSManagementObjectUnion_DiffArea:
		if value != nil {
			return value.DiffArea
		}
	}
	return nil
}

type is_VSSManagementObjectUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_VSSManagementObjectUnion()
}

func (o *VSSManagementObjectUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *VSSManagementObjectUnion_Vol:
		return uint32(1)
	case *VSSManagementObjectUnion_DiffVol:
		return uint32(2)
	case *VSSManagementObjectUnion_DiffArea:
		return uint32(3)
	}
	return uint32(0)
}

func (o *VSSManagementObjectUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*VSSManagementObjectUnion_Vol)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VSSManagementObjectUnion_Vol{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*VSSManagementObjectUnion_DiffVol)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VSSManagementObjectUnion_DiffVol{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*VSSManagementObjectUnion_DiffArea)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VSSManagementObjectUnion_DiffArea{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *VSSManagementObjectUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &VSSManagementObjectUnion_Vol{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &VSSManagementObjectUnion_DiffVol{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &VSSManagementObjectUnion_DiffArea{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// VSSManagementObjectUnion_Vol structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 1
type VSSManagementObjectUnion_Vol struct {
	Vol *VSSVolumeProperty `idl:"name:Vol" json:"vol"`
}

func (*VSSManagementObjectUnion_Vol) is_VSSManagementObjectUnion() {}

func (o *VSSManagementObjectUnion_Vol) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Vol != nil {
		if err := o.Vol.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSVolumeProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSManagementObjectUnion_Vol) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Vol == nil {
		o.Vol = &VSSVolumeProperty{}
	}
	if err := o.Vol.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VSSManagementObjectUnion_DiffVol structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 2
type VSSManagementObjectUnion_DiffVol struct {
	DiffVol *VSSDiffVolumeProperty `idl:"name:DiffVol" json:"diff_vol"`
}

func (*VSSManagementObjectUnion_DiffVol) is_VSSManagementObjectUnion() {}

func (o *VSSManagementObjectUnion_DiffVol) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DiffVol != nil {
		if err := o.DiffVol.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSDiffVolumeProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSManagementObjectUnion_DiffVol) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiffVol == nil {
		o.DiffVol = &VSSDiffVolumeProperty{}
	}
	if err := o.DiffVol.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VSSManagementObjectUnion_DiffArea structure represents VSS_MGMT_OBJECT_UNION RPC union arm.
//
// It has following labels: 3
type VSSManagementObjectUnion_DiffArea struct {
	DiffArea *VSSDiffAreaProperty `idl:"name:DiffArea" json:"diff_area"`
}

func (*VSSManagementObjectUnion_DiffArea) is_VSSManagementObjectUnion() {}

func (o *VSSManagementObjectUnion_DiffArea) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DiffArea != nil {
		if err := o.DiffArea.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&VSSDiffAreaProperty{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSManagementObjectUnion_DiffArea) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DiffArea == nil {
		o.DiffArea = &VSSDiffAreaProperty{}
	}
	if err := o.DiffArea.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VSSManagementObjectProperty structure represents VSS_MGMT_OBJECT_PROP RPC structure.
type VSSManagementObjectProperty struct {
	Type   VSSManagementObjectType   `idl:"name:Type" json:"type"`
	Object *VSSManagementObjectUnion `idl:"name:Obj;switch_is:Type" json:"object"`
}

func (o *VSSManagementObjectProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VSSManagementObjectProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&VSSManagementObjectUnion{}).MarshalUnionNDR(ctx, w, _swObject); err != nil {
			return err
		}
	}
	return nil
}
func (o *VSSManagementObjectProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Type)); err != nil {
		return err
	}
	if o.Object == nil {
		o.Object = &VSSManagementObjectUnion{}
	}
	_swObject := uint32(o.Type)
	if err := o.Object.UnmarshalUnionNDR(ctx, w, _swObject); err != nil {
		return err
	}
	return nil
}

// VSSSnapshotManagement structure represents IVssSnapshotMgmt RPC structure.
type VSSSnapshotManagement dcom.InterfacePointer

func (o *VSSSnapshotManagement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VSSSnapshotManagement) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VSSSnapshotManagement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VSSSnapshotManagement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSSnapshotManagement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VSSEnumManagementObject structure represents IVssEnumMgmtObject RPC structure.
type VSSEnumManagementObject dcom.InterfacePointer

func (o *VSSEnumManagementObject) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VSSEnumManagementObject) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VSSEnumManagementObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VSSEnumManagementObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSEnumManagementObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VSSEnumObject structure represents IVssEnumObject RPC structure.
type VSSEnumObject dcom.InterfacePointer

func (o *VSSEnumObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *VSSEnumObject) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VSSEnumObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VSSEnumObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSEnumObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// VSSDifferentialSoftwareSnapshotManagement structure represents IVssDifferentialSoftwareSnapshotMgmt RPC structure.
type VSSDifferentialSoftwareSnapshotManagement dcom.InterfacePointer

func (o *VSSDifferentialSoftwareSnapshotManagement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *VSSDifferentialSoftwareSnapshotManagement) xxx_PreparePayload(ctx context.Context) error {
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

func (o *VSSDifferentialSoftwareSnapshotManagement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VSSDifferentialSoftwareSnapshotManagement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *VSSDifferentialSoftwareSnapshotManagement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
