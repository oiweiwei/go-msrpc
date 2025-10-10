// The frs2 package implements the FRS2 client protocol.
//
// # Introduction
//
// # Overview
package frs2

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
	GoPackage = "frs2"
)

// CommunicationProtocolVersionW2K3R2 represents the FRS_COMMUNICATION_PROTOCOL_VERSION_W2K3R2 RPC constant
var CommunicationProtocolVersionW2K3R2 = 327680

// CommunicationProtocolVersionLonghornServer represents the FRS_COMMUNICATION_PROTOCOL_VERSION_LONGHORN_SERVER RPC constant
var CommunicationProtocolVersionLonghornServer = 327682

// CommunicationProtocolWin8Server represents the FRS_COMMUNICATION_PROTOCOL_WIN8_SERVER RPC constant
var CommunicationProtocolWin8Server = 327683

// CommunicationProtocolWinBlueServer represents the FRS_COMMUNICATION_PROTOCOL_WINBLUE_SERVER RPC constant
var CommunicationProtocolWinBlueServer = 327684

// ConfigRDCVersion represents the CONFIG_RDC_VERSION RPC constant
var ConfigRDCVersion = 1

// ConfigRDCVersionCompatible represents the CONFIG_RDC_VERSION_COMPATIBLE RPC constant
var ConfigRDCVersionCompatible = 1

// ConfigFileHashDataSize represents the CONFIG_FILEHASH_DATASIZE RPC constant
var ConfigFileHashDataSize = 20

// ConfigRDCSimilarityDataSize represents the CONFIG_RDC_SIMILARITY_DATASIZE RPC constant
var ConfigRDCSimilarityDataSize = 16

// ConfigRDCHorizonSizeMin represents the CONFIG_RDC_HORIZONSIZE_MIN RPC constant
var ConfigRDCHorizonSizeMin = 128

// ConfigRDCHorizonSizeMax represents the CONFIG_RDC_HORIZONSIZE_MAX RPC constant
var ConfigRDCHorizonSizeMax = 16384

// ConfigRDCHashWindowSizeMin represents the CONFIG_RDC_HASHWINDOWSIZE_MIN RPC constant
var ConfigRDCHashWindowSizeMin = 2

// ConfigRDCHashWindowSizeMax represents the CONFIG_RDC_HASHWINDOWSIZE_MAX RPC constant
var ConfigRDCHashWindowSizeMax = 96

// ConfigRDCMaxLevels represents the CONFIG_RDC_MAX_LEVELS RPC constant
var ConfigRDCMaxLevels = 8

// ConfigRDCMaxNeedLength represents the CONFIG_RDC_MAX_NEEDLENGTH RPC constant
var ConfigRDCMaxNeedLength = 65536

// ConfigTransportMaxBufferSize represents the CONFIG_TRANSPORT_MAX_BUFFER_SIZE RPC constant
var ConfigTransportMaxBufferSize = 262144

// ConfigRDCNeedQueueSize represents the CONFIG_RDC_NEED_QUEUE_SIZE RPC constant
var ConfigRDCNeedQueueSize = 20

// UpdateFlagGhostedHeader represents the FRS_UPDATE_FLAG_GHOSTED_HEADER RPC constant
var UpdateFlagGhostedHeader = 4

// UpdateFlagData represents the FRS_UPDATE_FLAG_DATA RPC constant
var UpdateFlagData = 8

// UpdateFlagClockDecremented represents the FRS_UPDATE_FLAG_CLOCK_DECREMENTED RPC constant
var UpdateFlagClockDecremented = 16

// ReplicaSetID structure represents FRS_REPLICA_SET_ID RPC structure.
type ReplicaSetID dtyp.GUID

func (o *ReplicaSetID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ReplicaSetID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ReplicaSetID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ReplicaSetID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ContentSetID structure represents FRS_CONTENT_SET_ID RPC structure.
type ContentSetID dtyp.GUID

func (o *ContentSetID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ContentSetID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ContentSetID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ContentSetID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DatabaseID structure represents FRS_DATABASE_ID RPC structure.
type DatabaseID dtyp.GUID

func (o *DatabaseID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *DatabaseID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DatabaseID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DatabaseID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MemberID structure represents FRS_MEMBER_ID RPC structure.
type MemberID dtyp.GUID

func (o *MemberID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *MemberID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *MemberID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MemberID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ConnectionID structure represents FRS_CONNECTION_ID RPC structure.
type ConnectionID dtyp.GUID

func (o *ConnectionID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ConnectionID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ConnectionID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ConnectionID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Epoque structure represents EPOQUE RPC structure.
type Epoque dtyp.SystemTime

func (o *Epoque) SystemTime() *dtyp.SystemTime { return (*dtyp.SystemTime)(o) }

func (o *Epoque) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Epoque) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.Year); err != nil {
		return err
	}
	if err := w.WriteData(o.Month); err != nil {
		return err
	}
	if err := w.WriteData(o.DayOfWeek); err != nil {
		return err
	}
	if err := w.WriteData(o.Day); err != nil {
		return err
	}
	if err := w.WriteData(o.Hour); err != nil {
		return err
	}
	if err := w.WriteData(o.Minute); err != nil {
		return err
	}
	if err := w.WriteData(o.Second); err != nil {
		return err
	}
	if err := w.WriteData(o.Milliseconds); err != nil {
		return err
	}
	return nil
}
func (o *Epoque) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Year); err != nil {
		return err
	}
	if err := w.ReadData(&o.Month); err != nil {
		return err
	}
	if err := w.ReadData(&o.DayOfWeek); err != nil {
		return err
	}
	if err := w.ReadData(&o.Day); err != nil {
		return err
	}
	if err := w.ReadData(&o.Hour); err != nil {
		return err
	}
	if err := w.ReadData(&o.Minute); err != nil {
		return err
	}
	if err := w.ReadData(&o.Second); err != nil {
		return err
	}
	if err := w.ReadData(&o.Milliseconds); err != nil {
		return err
	}
	return nil
}

// VersionVector structure represents FRS_VERSION_VECTOR RPC structure.
type VersionVector struct {
	DBGUID *dtyp.GUID `idl:"name:dbGuid" json:"db_guid"`
	Low    uint64     `idl:"name:low" json:"low"`
	High   uint64     `idl:"name:high" json:"high"`
}

func (o *VersionVector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *VersionVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.DBGUID != nil {
		if err := o.DBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Low); err != nil {
		return err
	}
	if err := w.WriteData(o.High); err != nil {
		return err
	}
	return nil
}
func (o *VersionVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.DBGUID == nil {
		o.DBGUID = &dtyp.GUID{}
	}
	if err := o.DBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Low); err != nil {
		return err
	}
	if err := w.ReadData(&o.High); err != nil {
		return err
	}
	return nil
}

// EpoqueVector structure represents FRS_EPOQUE_VECTOR RPC structure.
type EpoqueVector struct {
	Machine *dtyp.GUID `idl:"name:machine" json:"machine"`
	Epoque  *Epoque    `idl:"name:epoque" json:"epoque"`
}

func (o *EpoqueVector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EpoqueVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.Machine != nil {
		if err := o.Machine.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Epoque != nil {
		if err := o.Epoque.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Epoque{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *EpoqueVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Machine == nil {
		o.Machine = &dtyp.GUID{}
	}
	if err := o.Machine.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Epoque == nil {
		o.Epoque = &Epoque{}
	}
	if err := o.Epoque.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// IDGVSN structure represents FRS_ID_GVSN RPC structure.
type IDGVSN struct {
	UIDDBGUID   *dtyp.GUID `idl:"name:uidDbGuid" json:"uid_db_guid"`
	UIDVersion  uint64     `idl:"name:uidVersion" json:"uid_version"`
	GVSNDBGUID  *dtyp.GUID `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	GVSNVersion uint64     `idl:"name:gvsnVersion" json:"gvsn_version"`
}

func (o *IDGVSN) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IDGVSN) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.UIDDBGUID != nil {
		if err := o.UIDDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UIDVersion); err != nil {
		return err
	}
	if o.GVSNDBGUID != nil {
		if err := o.GVSNDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GVSNVersion); err != nil {
		return err
	}
	return nil
}
func (o *IDGVSN) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.UIDDBGUID == nil {
		o.UIDDBGUID = &dtyp.GUID{}
	}
	if err := o.UIDDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UIDVersion); err != nil {
		return err
	}
	if o.GVSNDBGUID == nil {
		o.GVSNDBGUID = &dtyp.GUID{}
	}
	if err := o.GVSNDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GVSNVersion); err != nil {
		return err
	}
	return nil
}

// Update structure represents FRS_UPDATE RPC structure.
type Update struct {
	Present       int32          `idl:"name:present" json:"present"`
	Conflict      int32          `idl:"name:nameConflict" json:"conflict"`
	Attributes    uint32         `idl:"name:attributes" json:"attributes"`
	Fence         *dtyp.Filetime `idl:"name:fence" json:"fence"`
	Clock         *dtyp.Filetime `idl:"name:clock" json:"clock"`
	CreateTime    *dtyp.Filetime `idl:"name:createTime" json:"create_time"`
	ContentSetID  *ContentSetID  `idl:"name:contentSetId" json:"content_set_id"`
	Hash          []byte         `idl:"name:hash" json:"hash"`
	Similarity    []byte         `idl:"name:rdcSimilarity" json:"similarity"`
	UIDDBGUID     *dtyp.GUID     `idl:"name:uidDbGuid" json:"uid_db_guid"`
	UIDVersion    uint64         `idl:"name:uidVersion" json:"uid_version"`
	GVSNDBGUID    *dtyp.GUID     `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	GVSNVersion   uint64         `idl:"name:gvsnVersion" json:"gvsn_version"`
	ParentDBGUID  *dtyp.GUID     `idl:"name:parentDbGuid" json:"parent_db_guid"`
	ParentVersion uint64         `idl:"name:parentVersion" json:"parent_version"`
	Name          string         `idl:"name:name;string" json:"name"`
	Flags         int32          `idl:"name:flags" json:"flags"`
}

func (o *Update) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Update) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Present); err != nil {
		return err
	}
	if err := w.WriteData(o.Conflict); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.Fence != nil {
		if err := o.Fence.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Clock != nil {
		if err := o.Clock.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CreateTime != nil {
		if err := o.CreateTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ContentSetID != nil {
		if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ContentSetID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Hash {
		i1 := i1
		if uint64(i1) >= 20 {
			break
		}
		if err := w.WriteData(o.Hash[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Hash); uint64(i1) < 20; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Similarity {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Similarity[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Similarity); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.UIDDBGUID != nil {
		if err := o.UIDDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UIDVersion); err != nil {
		return err
	}
	if o.GVSNDBGUID != nil {
		if err := o.GVSNDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GVSNVersion); err != nil {
		return err
	}
	if o.ParentDBGUID != nil {
		if err := o.ParentDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ParentVersion); err != nil {
		return err
	}
	_Name_buf := utf16.Encode([]rune(o.Name))
	if uint64(len(_Name_buf)) > 261-1 {
		_Name_buf = _Name_buf[:261-1]
	}
	if o.Name != ndr.ZeroString {
		_Name_buf = append(_Name_buf, uint16(0))
	}
	for i1 := range _Name_buf {
		i1 := i1
		if uint64(i1) >= 261 {
			break
		}
		if err := w.WriteData(_Name_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Name_buf); uint64(i1) < 261; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *Update) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Present); err != nil {
		return err
	}
	if err := w.ReadData(&o.Conflict); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.Fence == nil {
		o.Fence = &dtyp.Filetime{}
	}
	if err := o.Fence.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.Clock == nil {
		o.Clock = &dtyp.Filetime{}
	}
	if err := o.Clock.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CreateTime == nil {
		o.CreateTime = &dtyp.Filetime{}
	}
	if err := o.CreateTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ContentSetID == nil {
		o.ContentSetID = &ContentSetID{}
	}
	if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.Hash = make([]byte, 20)
	for i1 := range o.Hash {
		i1 := i1
		if err := w.ReadData(&o.Hash[i1]); err != nil {
			return err
		}
	}
	o.Similarity = make([]byte, 16)
	for i1 := range o.Similarity {
		i1 := i1
		if err := w.ReadData(&o.Similarity[i1]); err != nil {
			return err
		}
	}
	if o.UIDDBGUID == nil {
		o.UIDDBGUID = &dtyp.GUID{}
	}
	if err := o.UIDDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.UIDVersion); err != nil {
		return err
	}
	if o.GVSNDBGUID == nil {
		o.GVSNDBGUID = &dtyp.GUID{}
	}
	if err := o.GVSNDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GVSNVersion); err != nil {
		return err
	}
	if o.ParentDBGUID == nil {
		o.ParentDBGUID = &dtyp.GUID{}
	}
	if err := o.ParentDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParentVersion); err != nil {
		return err
	}
	var _Name_buf []uint16
	_Name_buf = make([]uint16, 261)
	for i1 := range _Name_buf {
		i1 := i1
		if err := w.ReadData(&_Name_buf[i1]); err != nil {
			return err
		}
	}
	o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// UpdateCancelData structure represents FRS_UPDATE_CANCEL_DATA RPC structure.
type UpdateCancelData struct {
	BlockingUpdate   *Update       `idl:"name:blockingUpdate" json:"blocking_update"`
	ContentSetID     *ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	GVSNDatabaseID   *DatabaseID   `idl:"name:gvsnDatabaseId" json:"gvsn_database_id"`
	UIDDatabaseID    *DatabaseID   `idl:"name:uidDatabaseId" json:"uid_database_id"`
	ParentDatabaseID *DatabaseID   `idl:"name:parentDatabaseId" json:"parent_database_id"`
	GVSNVersion      uint64        `idl:"name:gvsnVersion" json:"gvsn_version"`
	UIDVersion       uint64        `idl:"name:uidVersion" json:"uid_version"`
	ParentVersion    uint64        `idl:"name:parentVersion" json:"parent_version"`
	CancelType       uint32        `idl:"name:cancelType" json:"cancel_type"`
	IsUIDValid       int32         `idl:"name:isUidValid" json:"is_uid_valid"`
	IsParentUIDValid int32         `idl:"name:isParentUidValid" json:"is_parent_uid_valid"`
	IsBlockerValid   int32         `idl:"name:isBlockerValid" json:"is_blocker_valid"`
}

func (o *UpdateCancelData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *UpdateCancelData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.BlockingUpdate != nil {
		if err := o.BlockingUpdate.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Update{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ContentSetID != nil {
		if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ContentSetID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.GVSNDatabaseID != nil {
		if err := o.GVSNDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UIDDatabaseID != nil {
		if err := o.UIDDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ParentDatabaseID != nil {
		if err := o.ParentDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GVSNVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.UIDVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ParentVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.CancelType); err != nil {
		return err
	}
	if err := w.WriteData(o.IsUIDValid); err != nil {
		return err
	}
	if err := w.WriteData(o.IsParentUIDValid); err != nil {
		return err
	}
	if err := w.WriteData(o.IsBlockerValid); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *UpdateCancelData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.BlockingUpdate == nil {
		o.BlockingUpdate = &Update{}
	}
	if err := o.BlockingUpdate.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ContentSetID == nil {
		o.ContentSetID = &ContentSetID{}
	}
	if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.GVSNDatabaseID == nil {
		o.GVSNDatabaseID = &DatabaseID{}
	}
	if err := o.GVSNDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UIDDatabaseID == nil {
		o.UIDDatabaseID = &DatabaseID{}
	}
	if err := o.UIDDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ParentDatabaseID == nil {
		o.ParentDatabaseID = &DatabaseID{}
	}
	if err := o.ParentDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GVSNVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.UIDVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParentVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.CancelType); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsUIDValid); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsParentUIDValid); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsBlockerValid); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// SourceNeed structure represents FRS_RDC_SOURCE_NEED RPC structure.
type SourceNeed struct {
	NeedOffset uint64 `idl:"name:needOffset" json:"need_offset"`
	NeedSize   uint64 `idl:"name:needSize" json:"need_size"`
}

func (o *SourceNeed) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SourceNeed) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.NeedOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.NeedSize); err != nil {
		return err
	}
	return nil
}
func (o *SourceNeed) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.NeedOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.NeedSize); err != nil {
		return err
	}
	return nil
}

// TransportFlags type represents TransportFlags RPC enumeration.
type TransportFlags uint16

var (
	TransportFlagsSupportsRDCSimilarity TransportFlags = 1
)

func (o TransportFlags) String() string {
	switch o {
	case TransportFlagsSupportsRDCSimilarity:
		return "TransportFlagsSupportsRDCSimilarity"
	}
	return "Invalid"
}

// FileCompressionTypes type represents RDC_FILE_COMPRESSION_TYPES RPC enumeration.
type FileCompressionTypes uint16

var (
	FileCompressionTypesUncompressed FileCompressionTypes = 0
	FileCompressionTypesXPress       FileCompressionTypes = 1
)

func (o FileCompressionTypes) String() string {
	switch o {
	case FileCompressionTypesUncompressed:
		return "FileCompressionTypesUncompressed"
	case FileCompressionTypesXPress:
		return "FileCompressionTypesXPress"
	}
	return "Invalid"
}

// ChunkerAlgorithm type represents RDC_CHUNKER_ALGORITHM RPC enumeration.
type ChunkerAlgorithm uint16

var (
	ChunkerAlgorithmFiltergeneric ChunkerAlgorithm = 0
	ChunkerAlgorithmFilterMax     ChunkerAlgorithm = 1
	ChunkerAlgorithmFilterPoint   ChunkerAlgorithm = 2
	ChunkerAlgorithmMaxalgorithm  ChunkerAlgorithm = 3
)

func (o ChunkerAlgorithm) String() string {
	switch o {
	case ChunkerAlgorithmFiltergeneric:
		return "ChunkerAlgorithmFiltergeneric"
	case ChunkerAlgorithmFilterMax:
		return "ChunkerAlgorithmFilterMax"
	case ChunkerAlgorithmFilterPoint:
		return "ChunkerAlgorithmFilterPoint"
	case ChunkerAlgorithmMaxalgorithm:
		return "ChunkerAlgorithmMaxalgorithm"
	}
	return "Invalid"
}

// UpdateRequestType type represents UPDATE_REQUEST_TYPE RPC enumeration.
type UpdateRequestType uint16

var (
	UpdateRequestTypeAll        UpdateRequestType = 0
	UpdateRequestTypeTombstones UpdateRequestType = 1
	UpdateRequestTypeLive       UpdateRequestType = 2
)

func (o UpdateRequestType) String() string {
	switch o {
	case UpdateRequestTypeAll:
		return "UpdateRequestTypeAll"
	case UpdateRequestTypeTombstones:
		return "UpdateRequestTypeTombstones"
	case UpdateRequestTypeLive:
		return "UpdateRequestTypeLive"
	}
	return "Invalid"
}

// UpdateStatus type represents UPDATE_STATUS RPC enumeration.
type UpdateStatus uint16

var (
	UpdateStatusDone UpdateStatus = 2
	UpdateStatusMore UpdateStatus = 3
)

func (o UpdateStatus) String() string {
	switch o {
	case UpdateStatusDone:
		return "UpdateStatusDone"
	case UpdateStatusMore:
		return "UpdateStatusMore"
	}
	return "Invalid"
}

// RecordsStatus type represents RECORDS_STATUS RPC enumeration.
type RecordsStatus uint16

var (
	RecordsStatusDone RecordsStatus = 0
	RecordsStatusMore RecordsStatus = 1
)

func (o RecordsStatus) String() string {
	switch o {
	case RecordsStatusDone:
		return "RecordsStatusDone"
	case RecordsStatusMore:
		return "RecordsStatusMore"
	}
	return "Invalid"
}

// VersionRequestType type represents VERSION_REQUEST_TYPE RPC enumeration.
type VersionRequestType uint16

var (
	VersionRequestTypeNormalSync      VersionRequestType = 0
	VersionRequestTypeSlowSync        VersionRequestType = 1
	VersionRequestTypeSubordinateSync VersionRequestType = 2
)

func (o VersionRequestType) String() string {
	switch o {
	case VersionRequestTypeNormalSync:
		return "VersionRequestTypeNormalSync"
	case VersionRequestTypeSlowSync:
		return "VersionRequestTypeSlowSync"
	case VersionRequestTypeSubordinateSync:
		return "VersionRequestTypeSubordinateSync"
	}
	return "Invalid"
}

// VersionChangeType type represents VERSION_CHANGE_TYPE RPC enumeration.
type VersionChangeType uint16

var (
	VersionChangeTypeNotify VersionChangeType = 0
	VersionChangeTypeAll    VersionChangeType = 2
)

func (o VersionChangeType) String() string {
	switch o {
	case VersionChangeTypeNotify:
		return "VersionChangeTypeNotify"
	case VersionChangeTypeAll:
		return "VersionChangeTypeAll"
	}
	return "Invalid"
}

// RequestedStagingPolicy type represents FRS_REQUESTED_STAGING_POLICY RPC enumeration.
type RequestedStagingPolicy uint16

var (
	RequestedStagingPolicyServerDefault     RequestedStagingPolicy = 0
	RequestedStagingPolicyRequired          RequestedStagingPolicy = 1
	RequestedStagingPolicyRestagingRequired RequestedStagingPolicy = 2
)

func (o RequestedStagingPolicy) String() string {
	switch o {
	case RequestedStagingPolicyServerDefault:
		return "RequestedStagingPolicyServerDefault"
	case RequestedStagingPolicyRequired:
		return "RequestedStagingPolicyRequired"
	case RequestedStagingPolicyRestagingRequired:
		return "RequestedStagingPolicyRestagingRequired"
	}
	return "Invalid"
}

// ParametersFilterMax structure represents FRS_RDC_PARAMETERS_FILTERMAX RPC structure.
type ParametersFilterMax struct {
	HorizonSize uint16 `idl:"name:horizonSize" json:"horizon_size"`
	WindowSize  uint16 `idl:"name:windowSize" json:"window_size"`
}

func (o *ParametersFilterMax) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.HorizonSize < uint16(128) || o.HorizonSize > uint16(16384) {
		return fmt.Errorf("HorizonSize is out of range")
	}
	if o.WindowSize < uint16(2) || o.WindowSize > uint16(96) {
		return fmt.Errorf("WindowSize is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ParametersFilterMax) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.HorizonSize); err != nil {
		return err
	}
	if err := w.WriteData(o.WindowSize); err != nil {
		return err
	}
	return nil
}
func (o *ParametersFilterMax) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.HorizonSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.WindowSize); err != nil {
		return err
	}
	return nil
}

// ParametersFilterPoint structure represents FRS_RDC_PARAMETERS_FILTERPOINT RPC structure.
type ParametersFilterPoint struct {
	MinChunkSize uint16 `idl:"name:minChunkSize" json:"min_chunk_size"`
	MaxChunkSize uint16 `idl:"name:maxChunkSize" json:"max_chunk_size"`
}

func (o *ParametersFilterPoint) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ParametersFilterPoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.MinChunkSize); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxChunkSize); err != nil {
		return err
	}
	return nil
}
func (o *ParametersFilterPoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinChunkSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxChunkSize); err != nil {
		return err
	}
	return nil
}

// ParametersGeneric structure represents FRS_RDC_PARAMETERS_GENERIC RPC structure.
type ParametersGeneric struct {
	ChunkerType       uint16 `idl:"name:chunkerType" json:"chunker_type"`
	ChunkerParameters []byte `idl:"name:chunkerParameters" json:"chunker_parameters"`
}

func (o *ParametersGeneric) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ParametersGeneric) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ChunkerType); err != nil {
		return err
	}
	for i1 := range o.ChunkerParameters {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.ChunkerParameters[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ChunkerParameters); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(2); err != nil {
		return err
	}
	return nil
}
func (o *ParametersGeneric) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ChunkerType); err != nil {
		return err
	}
	o.ChunkerParameters = make([]byte, 64)
	for i1 := range o.ChunkerParameters {
		i1 := i1
		if err := w.ReadData(&o.ChunkerParameters[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(2); err != nil {
		return err
	}
	return nil
}

// Parameters structure represents FRS_RDC_PARAMETERS RPC structure.
type Parameters struct {
	ChunkerAlgorithm uint16            `idl:"name:rdcChunkerAlgorithm" json:"chunker_algorithm"`
	Union            *Parameters_Union `idl:"name:u;switch_is:rdcChunkerAlgorithm" json:"union"`
}

func (o *Parameters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Parameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ChunkerAlgorithm); err != nil {
		return err
	}
	_swUnion := uint16(o.ChunkerAlgorithm)
	if o.Union != nil {
		if err := o.Union.MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	} else {
		if err := (&Parameters_Union{}).MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *Parameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ChunkerAlgorithm); err != nil {
		return err
	}
	if o.Union == nil {
		o.Union = &Parameters_Union{}
	}
	_swUnion := uint16(o.ChunkerAlgorithm)
	if err := o.Union.UnmarshalUnionNDR(ctx, w, _swUnion); err != nil {
		return err
	}
	return nil
}

// Parameters_Union structure represents FRS_RDC_PARAMETERS union anonymous member.
type Parameters_Union struct {
	// Types that are assignable to Value
	//
	// *Parameters_Union_FilterGeneric
	// *Parameters_Union_FilterMax
	// *Parameters_Union_FilterPoint
	Value is_Parameters_Union `json:"value"`
}

func (o *Parameters_Union) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Parameters_Union_FilterGeneric:
		if value != nil {
			return value.FilterGeneric
		}
	case *Parameters_Union_FilterMax:
		if value != nil {
			return value.FilterMax
		}
	case *Parameters_Union_FilterPoint:
		if value != nil {
			return value.FilterPoint
		}
	}
	return nil
}

type is_Parameters_Union interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Parameters_Union()
}

func (o *Parameters_Union) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Parameters_Union_FilterGeneric:
		return uint16(0)
	case *Parameters_Union_FilterMax:
		return uint16(1)
	case *Parameters_Union_FilterPoint:
		return uint16(2)
	}
	return uint16(0)
}

func (o *Parameters_Union) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(2); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(2); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		_o, _ := o.Value.(*Parameters_Union_FilterGeneric)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Parameters_Union_FilterGeneric{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*Parameters_Union_FilterMax)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Parameters_Union_FilterMax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*Parameters_Union_FilterPoint)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Parameters_Union_FilterPoint{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Parameters_Union) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(2); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(2); err != nil {
		return err
	}
	switch sw {
	case uint16(0):
		o.Value = &Parameters_Union_FilterGeneric{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &Parameters_Union_FilterMax{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &Parameters_Union_FilterPoint{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Parameters_Union_FilterGeneric structure represents Parameters_Union RPC union arm.
//
// It has following labels: 0
type Parameters_Union_FilterGeneric struct {
	FilterGeneric *ParametersGeneric `idl:"name:filterGeneric" json:"filter_generic"`
}

func (*Parameters_Union_FilterGeneric) is_Parameters_Union() {}

func (o *Parameters_Union_FilterGeneric) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterGeneric != nil {
		if err := o.FilterGeneric.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ParametersGeneric{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Parameters_Union_FilterGeneric) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterGeneric == nil {
		o.FilterGeneric = &ParametersGeneric{}
	}
	if err := o.FilterGeneric.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Parameters_Union_FilterMax structure represents Parameters_Union RPC union arm.
//
// It has following labels: 1
type Parameters_Union_FilterMax struct {
	FilterMax *ParametersFilterMax `idl:"name:filterMax" json:"filter_max"`
}

func (*Parameters_Union_FilterMax) is_Parameters_Union() {}

func (o *Parameters_Union_FilterMax) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterMax != nil {
		if err := o.FilterMax.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ParametersFilterMax{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Parameters_Union_FilterMax) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterMax == nil {
		o.FilterMax = &ParametersFilterMax{}
	}
	if err := o.FilterMax.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Parameters_Union_FilterPoint structure represents Parameters_Union RPC union arm.
//
// It has following labels: 2
type Parameters_Union_FilterPoint struct {
	FilterPoint *ParametersFilterPoint `idl:"name:filterPoint" json:"filter_point"`
}

func (*Parameters_Union_FilterPoint) is_Parameters_Union() {}

func (o *Parameters_Union_FilterPoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterPoint != nil {
		if err := o.FilterPoint.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ParametersFilterPoint{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Parameters_Union_FilterPoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterPoint == nil {
		o.FilterPoint = &ParametersFilterPoint{}
	}
	if err := o.FilterPoint.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FileInfo structure represents FRS_RDC_FILEINFO RPC structure.
type FileInfo struct {
	OnDiskFileSize           uint64               `idl:"name:onDiskFileSize" json:"on_disk_file_size"`
	FileSizeEstimate         uint64               `idl:"name:fileSizeEstimate" json:"file_size_estimate"`
	Version                  uint16               `idl:"name:rdcVersion" json:"version"`
	MinimumCompatibleVersion uint16               `idl:"name:rdcMinimumCompatibleVersion" json:"minimum_compatible_version"`
	SignatureLevels          uint8                `idl:"name:rdcSignatureLevels" json:"signature_levels"`
	CompressionAlgorithm     FileCompressionTypes `idl:"name:compressionAlgorithm" json:"compression_algorithm"`
	FilterParameters         []*Parameters        `idl:"name:rdcFilterParameters;size_is:(rdcSignatureLevels)" json:"filter_parameters"`
}

func (o *FileInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.FilterParameters != nil && o.SignatureLevels == 0 {
		o.SignatureLevels = uint8(len(o.FilterParameters))
	}
	if o.SignatureLevels > uint8(8) {
		return fmt.Errorf("SignatureLevels is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *FileInfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.SignatureLevels)
	return []uint64{
		dimSize1,
	}
}
func (o *FileInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.OnDiskFileSize); err != nil {
		return err
	}
	if err := w.WriteData(o.FileSizeEstimate); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.MinimumCompatibleVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.SignatureLevels); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.CompressionAlgorithm)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	for i1 := range o.FilterParameters {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.FilterParameters[i1] != nil {
			if err := o.FilterParameters[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Parameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.FilterParameters); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&Parameters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FileInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.OnDiskFileSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileSizeEstimate); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinimumCompatibleVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.SignatureLevels); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.CompressionAlgorithm)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.SignatureLevels > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.SignatureLevels)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.FilterParameters", sizeInfo[0])
	}
	o.FilterParameters = make([]*Parameters, sizeInfo[0])
	for i1 := range o.FilterParameters {
		i1 := i1
		if o.FilterParameters[i1] == nil {
			o.FilterParameters[i1] = &Parameters{}
		}
		if err := o.FilterParameters[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// AsyncVersionVectorResponse structure represents FRS_ASYNC_VERSION_VECTOR_RESPONSE RPC structure.
type AsyncVersionVectorResponse struct {
	Generation         uint64           `idl:"name:vvGeneration" json:"generation"`
	VersionVectorCount uint32           `idl:"name:versionVectorCount" json:"version_vector_count"`
	VersionVector      []*VersionVector `idl:"name:versionVector;size_is:(versionVectorCount)" json:"version_vector"`
	EpoqueVectorCount  uint32           `idl:"name:epoqueVectorCount" json:"epoque_vector_count"`
	EpoqueVector       []*EpoqueVector  `idl:"name:epoqueVector;size_is:(epoqueVectorCount)" json:"epoque_vector"`
}

func (o *AsyncVersionVectorResponse) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.VersionVector != nil && o.VersionVectorCount == 0 {
		o.VersionVectorCount = uint32(len(o.VersionVector))
	}
	if o.EpoqueVector != nil && o.EpoqueVectorCount == 0 {
		o.EpoqueVectorCount = uint32(len(o.EpoqueVector))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *AsyncVersionVectorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Generation); err != nil {
		return err
	}
	if err := w.WriteData(o.VersionVectorCount); err != nil {
		return err
	}
	if o.VersionVector != nil || o.VersionVectorCount > 0 {
		_ptr_versionVector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.VersionVectorCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.VersionVector {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.VersionVector[i1] != nil {
					if err := o.VersionVector[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&VersionVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.VersionVector); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&VersionVector{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VersionVector, _ptr_versionVector); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EpoqueVectorCount); err != nil {
		return err
	}
	if o.EpoqueVector != nil || o.EpoqueVectorCount > 0 {
		_ptr_epoqueVector := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EpoqueVectorCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EpoqueVector {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.EpoqueVector[i1] != nil {
					if err := o.EpoqueVector[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EpoqueVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EpoqueVector); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&EpoqueVector{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EpoqueVector, _ptr_epoqueVector); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *AsyncVersionVectorResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Generation); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersionVectorCount); err != nil {
		return err
	}
	_ptr_versionVector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.VersionVectorCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.VersionVectorCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VersionVector", sizeInfo[0])
		}
		o.VersionVector = make([]*VersionVector, sizeInfo[0])
		for i1 := range o.VersionVector {
			i1 := i1
			if o.VersionVector[i1] == nil {
				o.VersionVector[i1] = &VersionVector{}
			}
			if err := o.VersionVector[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_versionVector := func(ptr interface{}) { o.VersionVector = *ptr.(*[]*VersionVector) }
	if err := w.ReadPointer(&o.VersionVector, _s_versionVector, _ptr_versionVector); err != nil {
		return err
	}
	if err := w.ReadData(&o.EpoqueVectorCount); err != nil {
		return err
	}
	_ptr_epoqueVector := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EpoqueVectorCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EpoqueVectorCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EpoqueVector", sizeInfo[0])
		}
		o.EpoqueVector = make([]*EpoqueVector, sizeInfo[0])
		for i1 := range o.EpoqueVector {
			i1 := i1
			if o.EpoqueVector[i1] == nil {
				o.EpoqueVector[i1] = &EpoqueVector{}
			}
			if err := o.EpoqueVector[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_epoqueVector := func(ptr interface{}) { o.EpoqueVector = *ptr.(*[]*EpoqueVector) }
	if err := w.ReadPointer(&o.EpoqueVector, _s_epoqueVector, _ptr_epoqueVector); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// AsyncResponseContext structure represents FRS_ASYNC_RESPONSE_CONTEXT RPC structure.
type AsyncResponseContext struct {
	SequenceNumber uint32                      `idl:"name:sequenceNumber" json:"sequence_number"`
	Status         uint32                      `idl:"name:status" json:"status"`
	Result         *AsyncVersionVectorResponse `idl:"name:result" json:"result"`
}

func (o *AsyncResponseContext) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *AsyncResponseContext) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SequenceNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	if o.Result != nil {
		if err := o.Result.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&AsyncVersionVectorResponse{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *AsyncResponseContext) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SequenceNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	if o.Result == nil {
		o.Result = &AsyncVersionVectorResponse{}
	}
	if err := o.Result.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// BytePipe type represents BYTE_PIPE RPC pipe.
type BytePipe interface {
	Pipe() <-chan []byte
	Send(v []byte)
	Append(v []byte)
	Recv() ([]byte, bool)
	Close() error
}

// Pipe provides a channel-based interface to send data to the RPC server.
// Note: pipe must be closed after use to avoid resource leaks.
type xxx_BytePipe struct {
	p    chan []byte
	recv [][]byte
}

// Newxxx_BytePipe creates a new BytePipe instance.
func NewBytePipe() BytePipe {
	return &xxx_BytePipe{
		p: make(chan []byte),
	}
}

func (o *xxx_BytePipe) Pipe() <-chan []byte {
	return o.p
}

// Send sends data to the RPC server via the pipe.
func (o *xxx_BytePipe) Send(v []byte) {
	if o != nil && o.p != nil {
		o.p <- v
	}
}

// Append appends data to the internal buffer of the pipe.
func (o *xxx_BytePipe) Append(v []byte) {
	if o != nil {
		o.recv = append(o.recv, v)
	}
}

// Recv receives data from the internal buffer of the pipe.
func (o *xxx_BytePipe) Recv() ([]byte, bool) {
	if o != nil {
		if len(o.recv) > 0 {
			v := o.recv[0]
			o.recv = o.recv[1:]
			return v, true
		}
	}
	return nil, false
}
func (o *xxx_BytePipe) Close() error {
	if o != nil && o.p != nil {
		close(o.p)
	}
	return nil
}
