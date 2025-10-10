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

// FrsCommunicationProtocolVersionW2k3r2 represents the FRS_COMMUNICATION_PROTOCOL_VERSION_W2K3R2 RPC constant
var FrsCommunicationProtocolVersionW2k3r2 = 327680

// FrsCommunicationProtocolVersionLonghornServer represents the FRS_COMMUNICATION_PROTOCOL_VERSION_LONGHORN_SERVER RPC constant
var FrsCommunicationProtocolVersionLonghornServer = 327682

// FrsCommunicationProtocolWin8Server represents the FRS_COMMUNICATION_PROTOCOL_WIN8_SERVER RPC constant
var FrsCommunicationProtocolWin8Server = 327683

// FrsCommunicationProtocolWinblueServer represents the FRS_COMMUNICATION_PROTOCOL_WINBLUE_SERVER RPC constant
var FrsCommunicationProtocolWinblueServer = 327684

// ConfigRdcVersion represents the CONFIG_RDC_VERSION RPC constant
var ConfigRdcVersion = 1

// ConfigRdcVersionCompatible represents the CONFIG_RDC_VERSION_COMPATIBLE RPC constant
var ConfigRdcVersionCompatible = 1

// ConfigFilehashDatasize represents the CONFIG_FILEHASH_DATASIZE RPC constant
var ConfigFilehashDatasize = 20

// ConfigRdcSimilarityDatasize represents the CONFIG_RDC_SIMILARITY_DATASIZE RPC constant
var ConfigRdcSimilarityDatasize = 16

// ConfigRdcHorizonsizeMin represents the CONFIG_RDC_HORIZONSIZE_MIN RPC constant
var ConfigRdcHorizonsizeMin = 128

// ConfigRdcHorizonsizeMax represents the CONFIG_RDC_HORIZONSIZE_MAX RPC constant
var ConfigRdcHorizonsizeMax = 16384

// ConfigRdcHashwindowsizeMin represents the CONFIG_RDC_HASHWINDOWSIZE_MIN RPC constant
var ConfigRdcHashwindowsizeMin = 2

// ConfigRdcHashwindowsizeMax represents the CONFIG_RDC_HASHWINDOWSIZE_MAX RPC constant
var ConfigRdcHashwindowsizeMax = 96

// ConfigRdcMaxLevels represents the CONFIG_RDC_MAX_LEVELS RPC constant
var ConfigRdcMaxLevels = 8

// ConfigRdcMaxNeedlength represents the CONFIG_RDC_MAX_NEEDLENGTH RPC constant
var ConfigRdcMaxNeedlength = 65536

// ConfigTransportMaxBufferSize represents the CONFIG_TRANSPORT_MAX_BUFFER_SIZE RPC constant
var ConfigTransportMaxBufferSize = 262144

// ConfigRdcNeedQueueSize represents the CONFIG_RDC_NEED_QUEUE_SIZE RPC constant
var ConfigRdcNeedQueueSize = 20

// FrsUpdateFlagGhostedHeader represents the FRS_UPDATE_FLAG_GHOSTED_HEADER RPC constant
var FrsUpdateFlagGhostedHeader = 4

// FrsUpdateFlagData represents the FRS_UPDATE_FLAG_DATA RPC constant
var FrsUpdateFlagData = 8

// FrsUpdateFlagClockDecremented represents the FRS_UPDATE_FLAG_CLOCK_DECREMENTED RPC constant
var FrsUpdateFlagClockDecremented = 16

// FrsReplicaSetID structure represents FRS_REPLICA_SET_ID RPC structure.
type FrsReplicaSetID dtyp.GUID

func (o *FrsReplicaSetID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *FrsReplicaSetID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsReplicaSetID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsReplicaSetID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsContentSetID structure represents FRS_CONTENT_SET_ID RPC structure.
type FrsContentSetID dtyp.GUID

func (o *FrsContentSetID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *FrsContentSetID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsContentSetID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsContentSetID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsDatabaseID structure represents FRS_DATABASE_ID RPC structure.
type FrsDatabaseID dtyp.GUID

func (o *FrsDatabaseID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *FrsDatabaseID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsDatabaseID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsDatabaseID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsMemberID structure represents FRS_MEMBER_ID RPC structure.
type FrsMemberID dtyp.GUID

func (o *FrsMemberID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *FrsMemberID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsMemberID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsMemberID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsConnectionID structure represents FRS_CONNECTION_ID RPC structure.
type FrsConnectionID dtyp.GUID

func (o *FrsConnectionID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *FrsConnectionID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsConnectionID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsConnectionID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsVersionVector structure represents FRS_VERSION_VECTOR RPC structure.
type FrsVersionVector struct {
	DBGUID *dtyp.GUID `idl:"name:dbGuid" json:"db_guid"`
	Low    uint64     `idl:"name:low" json:"low"`
	High   uint64     `idl:"name:high" json:"high"`
}

func (o *FrsVersionVector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsVersionVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsVersionVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsEpoqueVector structure represents FRS_EPOQUE_VECTOR RPC structure.
type FrsEpoqueVector struct {
	Machine *dtyp.GUID `idl:"name:machine" json:"machine"`
	Epoque  *Epoque    `idl:"name:epoque" json:"epoque"`
}

func (o *FrsEpoqueVector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsEpoqueVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsEpoqueVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsIDGvsn structure represents FRS_ID_GVSN RPC structure.
type FrsIDGvsn struct {
	UIDDBGUID   *dtyp.GUID `idl:"name:uidDbGuid" json:"uid_db_guid"`
	UIDVersion  uint64     `idl:"name:uidVersion" json:"uid_version"`
	GvsnDBGUID  *dtyp.GUID `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	GvsnVersion uint64     `idl:"name:gvsnVersion" json:"gvsn_version"`
}

func (o *FrsIDGvsn) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsIDGvsn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.GvsnDBGUID != nil {
		if err := o.GvsnDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GvsnVersion); err != nil {
		return err
	}
	return nil
}
func (o *FrsIDGvsn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.GvsnDBGUID == nil {
		o.GvsnDBGUID = &dtyp.GUID{}
	}
	if err := o.GvsnDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GvsnVersion); err != nil {
		return err
	}
	return nil
}

// FrsUpdate structure represents FRS_UPDATE RPC structure.
type FrsUpdate struct {
	Present       int32            `idl:"name:present" json:"present"`
	Conflict      int32            `idl:"name:nameConflict" json:"conflict"`
	Attributes    uint32           `idl:"name:attributes" json:"attributes"`
	Fence         *dtyp.Filetime   `idl:"name:fence" json:"fence"`
	Clock         *dtyp.Filetime   `idl:"name:clock" json:"clock"`
	CreateTime    *dtyp.Filetime   `idl:"name:createTime" json:"create_time"`
	ContentSetID  *FrsContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	Hash          []byte           `idl:"name:hash" json:"hash"`
	RdcSimilarity []byte           `idl:"name:rdcSimilarity" json:"rdc_similarity"`
	UIDDBGUID     *dtyp.GUID       `idl:"name:uidDbGuid" json:"uid_db_guid"`
	UIDVersion    uint64           `idl:"name:uidVersion" json:"uid_version"`
	GvsnDBGUID    *dtyp.GUID       `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	GvsnVersion   uint64           `idl:"name:gvsnVersion" json:"gvsn_version"`
	ParentDBGUID  *dtyp.GUID       `idl:"name:parentDbGuid" json:"parent_db_guid"`
	ParentVersion uint64           `idl:"name:parentVersion" json:"parent_version"`
	Name          string           `idl:"name:name;string" json:"name"`
	Flags         int32            `idl:"name:flags" json:"flags"`
}

func (o *FrsUpdate) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsUpdate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&FrsContentSetID{}).MarshalNDR(ctx, w); err != nil {
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
	for i1 := range o.RdcSimilarity {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.RdcSimilarity[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.RdcSimilarity); uint64(i1) < 16; i1++ {
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
	if o.GvsnDBGUID != nil {
		if err := o.GvsnDBGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GvsnVersion); err != nil {
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
func (o *FrsUpdate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.ContentSetID = &FrsContentSetID{}
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
	o.RdcSimilarity = make([]byte, 16)
	for i1 := range o.RdcSimilarity {
		i1 := i1
		if err := w.ReadData(&o.RdcSimilarity[i1]); err != nil {
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
	if o.GvsnDBGUID == nil {
		o.GvsnDBGUID = &dtyp.GUID{}
	}
	if err := o.GvsnDBGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GvsnVersion); err != nil {
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

// FrsUpdateCancelData structure represents FRS_UPDATE_CANCEL_DATA RPC structure.
type FrsUpdateCancelData struct {
	BlockingUpdate   *FrsUpdate       `idl:"name:blockingUpdate" json:"blocking_update"`
	ContentSetID     *FrsContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	GvsnDatabaseID   *FrsDatabaseID   `idl:"name:gvsnDatabaseId" json:"gvsn_database_id"`
	UIDDatabaseID    *FrsDatabaseID   `idl:"name:uidDatabaseId" json:"uid_database_id"`
	ParentDatabaseID *FrsDatabaseID   `idl:"name:parentDatabaseId" json:"parent_database_id"`
	GvsnVersion      uint64           `idl:"name:gvsnVersion" json:"gvsn_version"`
	UIDVersion       uint64           `idl:"name:uidVersion" json:"uid_version"`
	ParentVersion    uint64           `idl:"name:parentVersion" json:"parent_version"`
	CancelType       uint32           `idl:"name:cancelType" json:"cancel_type"`
	IsUIDValid       int32            `idl:"name:isUidValid" json:"is_uid_valid"`
	IsParentUIDValid int32            `idl:"name:isParentUidValid" json:"is_parent_uid_valid"`
	IsBlockerValid   int32            `idl:"name:isBlockerValid" json:"is_blocker_valid"`
}

func (o *FrsUpdateCancelData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsUpdateCancelData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&FrsUpdate{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ContentSetID != nil {
		if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsContentSetID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.GvsnDatabaseID != nil {
		if err := o.GvsnDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsDatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UIDDatabaseID != nil {
		if err := o.UIDDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsDatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ParentDatabaseID != nil {
		if err := o.ParentDatabaseID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsDatabaseID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GvsnVersion); err != nil {
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
func (o *FrsUpdateCancelData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.BlockingUpdate == nil {
		o.BlockingUpdate = &FrsUpdate{}
	}
	if err := o.BlockingUpdate.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ContentSetID == nil {
		o.ContentSetID = &FrsContentSetID{}
	}
	if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.GvsnDatabaseID == nil {
		o.GvsnDatabaseID = &FrsDatabaseID{}
	}
	if err := o.GvsnDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.UIDDatabaseID == nil {
		o.UIDDatabaseID = &FrsDatabaseID{}
	}
	if err := o.UIDDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ParentDatabaseID == nil {
		o.ParentDatabaseID = &FrsDatabaseID{}
	}
	if err := o.ParentDatabaseID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.GvsnVersion); err != nil {
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

// FrsRdcSourceNeed structure represents FRS_RDC_SOURCE_NEED RPC structure.
type FrsRdcSourceNeed struct {
	NeedOffset uint64 `idl:"name:needOffset" json:"need_offset"`
	NeedSize   uint64 `idl:"name:needSize" json:"need_size"`
}

func (o *FrsRdcSourceNeed) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsRdcSourceNeed) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsRdcSourceNeed) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	TransportFlagsSupportsRdcSimilarity TransportFlags = 1
)

func (o TransportFlags) String() string {
	switch o {
	case TransportFlagsSupportsRdcSimilarity:
		return "TransportFlagsSupportsRdcSimilarity"
	}
	return "Invalid"
}

// RdcFileCompressionTypes type represents RDC_FILE_COMPRESSION_TYPES RPC enumeration.
type RdcFileCompressionTypes uint16

var (
	RdcFileCompressionTypesUncompressed RdcFileCompressionTypes = 0
	RdcFileCompressionTypesXPress       RdcFileCompressionTypes = 1
)

func (o RdcFileCompressionTypes) String() string {
	switch o {
	case RdcFileCompressionTypesUncompressed:
		return "RdcFileCompressionTypesUncompressed"
	case RdcFileCompressionTypesXPress:
		return "RdcFileCompressionTypesXPress"
	}
	return "Invalid"
}

// RdcChunkerAlgorithm type represents RDC_CHUNKER_ALGORITHM RPC enumeration.
type RdcChunkerAlgorithm uint16

var (
	RdcChunkerAlgorithmFiltergeneric RdcChunkerAlgorithm = 0
	RdcChunkerAlgorithmFiltermax     RdcChunkerAlgorithm = 1
	RdcChunkerAlgorithmFilterpoint   RdcChunkerAlgorithm = 2
	RdcChunkerAlgorithmMaxalgorithm  RdcChunkerAlgorithm = 3
)

func (o RdcChunkerAlgorithm) String() string {
	switch o {
	case RdcChunkerAlgorithmFiltergeneric:
		return "RdcChunkerAlgorithmFiltergeneric"
	case RdcChunkerAlgorithmFiltermax:
		return "RdcChunkerAlgorithmFiltermax"
	case RdcChunkerAlgorithmFilterpoint:
		return "RdcChunkerAlgorithmFilterpoint"
	case RdcChunkerAlgorithmMaxalgorithm:
		return "RdcChunkerAlgorithmMaxalgorithm"
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

// FrsRequestedStagingPolicy type represents FRS_REQUESTED_STAGING_POLICY RPC enumeration.
type FrsRequestedStagingPolicy uint16

var (
	FrsRequestedStagingPolicyServerDefault     FrsRequestedStagingPolicy = 0
	FrsRequestedStagingPolicyRequired          FrsRequestedStagingPolicy = 1
	FrsRequestedStagingPolicyRestagingRequired FrsRequestedStagingPolicy = 2
)

func (o FrsRequestedStagingPolicy) String() string {
	switch o {
	case FrsRequestedStagingPolicyServerDefault:
		return "FrsRequestedStagingPolicyServerDefault"
	case FrsRequestedStagingPolicyRequired:
		return "FrsRequestedStagingPolicyRequired"
	case FrsRequestedStagingPolicyRestagingRequired:
		return "FrsRequestedStagingPolicyRestagingRequired"
	}
	return "Invalid"
}

// FrsRdcParametersFiltermax structure represents FRS_RDC_PARAMETERS_FILTERMAX RPC structure.
type FrsRdcParametersFiltermax struct {
	HorizonSize uint16 `idl:"name:horizonSize" json:"horizon_size"`
	WindowSize  uint16 `idl:"name:windowSize" json:"window_size"`
}

func (o *FrsRdcParametersFiltermax) xxx_PreparePayload(ctx context.Context) error {
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
func (o *FrsRdcParametersFiltermax) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsRdcParametersFiltermax) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsRdcParametersFilterpoint structure represents FRS_RDC_PARAMETERS_FILTERPOINT RPC structure.
type FrsRdcParametersFilterpoint struct {
	MinChunkSize uint16 `idl:"name:minChunkSize" json:"min_chunk_size"`
	MaxChunkSize uint16 `idl:"name:maxChunkSize" json:"max_chunk_size"`
}

func (o *FrsRdcParametersFilterpoint) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsRdcParametersFilterpoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsRdcParametersFilterpoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsRdcParametersGeneric structure represents FRS_RDC_PARAMETERS_GENERIC RPC structure.
type FrsRdcParametersGeneric struct {
	ChunkerType       uint16 `idl:"name:chunkerType" json:"chunker_type"`
	ChunkerParameters []byte `idl:"name:chunkerParameters" json:"chunker_parameters"`
}

func (o *FrsRdcParametersGeneric) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsRdcParametersGeneric) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FrsRdcParametersGeneric) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FrsRdcParameters structure represents FRS_RDC_PARAMETERS RPC structure.
type FrsRdcParameters struct {
	RdcChunkerAlgorithm uint16                  `idl:"name:rdcChunkerAlgorithm" json:"rdc_chunker_algorithm"`
	Union               *FrsRdcParameters_Union `idl:"name:u;switch_is:rdcChunkerAlgorithm" json:"union"`
}

func (o *FrsRdcParameters) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsRdcParameters) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.RdcChunkerAlgorithm); err != nil {
		return err
	}
	_swUnion := uint16(o.RdcChunkerAlgorithm)
	if o.Union != nil {
		if err := o.Union.MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	} else {
		if err := (&FrsRdcParameters_Union{}).MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsRdcParameters) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.RdcChunkerAlgorithm); err != nil {
		return err
	}
	if o.Union == nil {
		o.Union = &FrsRdcParameters_Union{}
	}
	_swUnion := uint16(o.RdcChunkerAlgorithm)
	if err := o.Union.UnmarshalUnionNDR(ctx, w, _swUnion); err != nil {
		return err
	}
	return nil
}

// FrsRdcParameters_Union structure represents FRS_RDC_PARAMETERS union anonymous member.
type FrsRdcParameters_Union struct {
	// Types that are assignable to Value
	//
	// *FrsRdcParameters_Union_FilterGeneric
	// *FrsRdcParameters_Union_FilterMax
	// *FrsRdcParameters_Union_FilterPoint
	Value is_FrsRdcParameters_Union `json:"value"`
}

func (o *FrsRdcParameters_Union) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *FrsRdcParameters_Union_FilterGeneric:
		if value != nil {
			return value.FilterGeneric
		}
	case *FrsRdcParameters_Union_FilterMax:
		if value != nil {
			return value.FilterMax
		}
	case *FrsRdcParameters_Union_FilterPoint:
		if value != nil {
			return value.FilterPoint
		}
	}
	return nil
}

type is_FrsRdcParameters_Union interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_FrsRdcParameters_Union()
}

func (o *FrsRdcParameters_Union) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *FrsRdcParameters_Union_FilterGeneric:
		return uint16(0)
	case *FrsRdcParameters_Union_FilterMax:
		return uint16(1)
	case *FrsRdcParameters_Union_FilterPoint:
		return uint16(2)
	}
	return uint16(0)
}

func (o *FrsRdcParameters_Union) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
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
		_o, _ := o.Value.(*FrsRdcParameters_Union_FilterGeneric)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FrsRdcParameters_Union_FilterGeneric{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*FrsRdcParameters_Union_FilterMax)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FrsRdcParameters_Union_FilterMax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*FrsRdcParameters_Union_FilterPoint)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FrsRdcParameters_Union_FilterPoint{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *FrsRdcParameters_Union) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
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
		o.Value = &FrsRdcParameters_Union_FilterGeneric{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &FrsRdcParameters_Union_FilterMax{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &FrsRdcParameters_Union_FilterPoint{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// FrsRdcParameters_Union_FilterGeneric structure represents FrsRdcParameters_Union RPC union arm.
//
// It has following labels: 0
type FrsRdcParameters_Union_FilterGeneric struct {
	FilterGeneric *FrsRdcParametersGeneric `idl:"name:filterGeneric" json:"filter_generic"`
}

func (*FrsRdcParameters_Union_FilterGeneric) is_FrsRdcParameters_Union() {}

func (o *FrsRdcParameters_Union_FilterGeneric) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterGeneric != nil {
		if err := o.FilterGeneric.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsRdcParametersGeneric{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsRdcParameters_Union_FilterGeneric) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterGeneric == nil {
		o.FilterGeneric = &FrsRdcParametersGeneric{}
	}
	if err := o.FilterGeneric.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FrsRdcParameters_Union_FilterMax structure represents FrsRdcParameters_Union RPC union arm.
//
// It has following labels: 1
type FrsRdcParameters_Union_FilterMax struct {
	FilterMax *FrsRdcParametersFiltermax `idl:"name:filterMax" json:"filter_max"`
}

func (*FrsRdcParameters_Union_FilterMax) is_FrsRdcParameters_Union() {}

func (o *FrsRdcParameters_Union_FilterMax) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterMax != nil {
		if err := o.FilterMax.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsRdcParametersFiltermax{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsRdcParameters_Union_FilterMax) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterMax == nil {
		o.FilterMax = &FrsRdcParametersFiltermax{}
	}
	if err := o.FilterMax.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FrsRdcParameters_Union_FilterPoint structure represents FrsRdcParameters_Union RPC union arm.
//
// It has following labels: 2
type FrsRdcParameters_Union_FilterPoint struct {
	FilterPoint *FrsRdcParametersFilterpoint `idl:"name:filterPoint" json:"filter_point"`
}

func (*FrsRdcParameters_Union_FilterPoint) is_FrsRdcParameters_Union() {}

func (o *FrsRdcParameters_Union_FilterPoint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.FilterPoint != nil {
		if err := o.FilterPoint.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&FrsRdcParametersFilterpoint{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsRdcParameters_Union_FilterPoint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.FilterPoint == nil {
		o.FilterPoint = &FrsRdcParametersFilterpoint{}
	}
	if err := o.FilterPoint.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FrsRdcFileinfo structure represents FRS_RDC_FILEINFO RPC structure.
type FrsRdcFileinfo struct {
	OnDiskFileSize              uint64                  `idl:"name:onDiskFileSize" json:"on_disk_file_size"`
	FileSizeEstimate            uint64                  `idl:"name:fileSizeEstimate" json:"file_size_estimate"`
	RdcVersion                  uint16                  `idl:"name:rdcVersion" json:"rdc_version"`
	RdcMinimumCompatibleVersion uint16                  `idl:"name:rdcMinimumCompatibleVersion" json:"rdc_minimum_compatible_version"`
	RdcSignatureLevels          uint8                   `idl:"name:rdcSignatureLevels" json:"rdc_signature_levels"`
	CompressionAlgorithm        RdcFileCompressionTypes `idl:"name:compressionAlgorithm" json:"compression_algorithm"`
	RdcFilterParameters         []*FrsRdcParameters     `idl:"name:rdcFilterParameters;size_is:(rdcSignatureLevels)" json:"rdc_filter_parameters"`
}

func (o *FrsRdcFileinfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.RdcFilterParameters != nil && o.RdcSignatureLevels == 0 {
		o.RdcSignatureLevels = uint8(len(o.RdcFilterParameters))
	}
	if o.RdcSignatureLevels > uint8(8) {
		return fmt.Errorf("RdcSignatureLevels is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *FrsRdcFileinfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.RdcSignatureLevels)
	return []uint64{
		dimSize1,
	}
}
func (o *FrsRdcFileinfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.RdcVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.RdcMinimumCompatibleVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.RdcSignatureLevels); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.CompressionAlgorithm)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	for i1 := range o.RdcFilterParameters {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.RdcFilterParameters[i1] != nil {
			if err := o.RdcFilterParameters[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FrsRdcParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.RdcFilterParameters); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&FrsRdcParameters{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsRdcFileinfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.RdcVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.RdcMinimumCompatibleVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.RdcSignatureLevels); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.CompressionAlgorithm)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.RdcSignatureLevels > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.RdcSignatureLevels)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.RdcFilterParameters", sizeInfo[0])
	}
	o.RdcFilterParameters = make([]*FrsRdcParameters, sizeInfo[0])
	for i1 := range o.RdcFilterParameters {
		i1 := i1
		if o.RdcFilterParameters[i1] == nil {
			o.RdcFilterParameters[i1] = &FrsRdcParameters{}
		}
		if err := o.RdcFilterParameters[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// FrsAsyncVersionVectorResponse structure represents FRS_ASYNC_VERSION_VECTOR_RESPONSE RPC structure.
type FrsAsyncVersionVectorResponse struct {
	VvGeneration       uint64              `idl:"name:vvGeneration" json:"vv_generation"`
	VersionVectorCount uint32              `idl:"name:versionVectorCount" json:"version_vector_count"`
	VersionVector      []*FrsVersionVector `idl:"name:versionVector;size_is:(versionVectorCount)" json:"version_vector"`
	EpoqueVectorCount  uint32              `idl:"name:epoqueVectorCount" json:"epoque_vector_count"`
	EpoqueVector       []*FrsEpoqueVector  `idl:"name:epoqueVector;size_is:(epoqueVectorCount)" json:"epoque_vector"`
}

func (o *FrsAsyncVersionVectorResponse) xxx_PreparePayload(ctx context.Context) error {
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
func (o *FrsAsyncVersionVectorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.VvGeneration); err != nil {
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
					if err := (&FrsVersionVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.VersionVector); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&FrsVersionVector{}).MarshalNDR(ctx, w); err != nil {
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
					if err := (&FrsEpoqueVector{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EpoqueVector); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&FrsEpoqueVector{}).MarshalNDR(ctx, w); err != nil {
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
func (o *FrsAsyncVersionVectorResponse) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.VvGeneration); err != nil {
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
		o.VersionVector = make([]*FrsVersionVector, sizeInfo[0])
		for i1 := range o.VersionVector {
			i1 := i1
			if o.VersionVector[i1] == nil {
				o.VersionVector[i1] = &FrsVersionVector{}
			}
			if err := o.VersionVector[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_versionVector := func(ptr interface{}) { o.VersionVector = *ptr.(*[]*FrsVersionVector) }
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
		o.EpoqueVector = make([]*FrsEpoqueVector, sizeInfo[0])
		for i1 := range o.EpoqueVector {
			i1 := i1
			if o.EpoqueVector[i1] == nil {
				o.EpoqueVector[i1] = &FrsEpoqueVector{}
			}
			if err := o.EpoqueVector[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_epoqueVector := func(ptr interface{}) { o.EpoqueVector = *ptr.(*[]*FrsEpoqueVector) }
	if err := w.ReadPointer(&o.EpoqueVector, _s_epoqueVector, _ptr_epoqueVector); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// FrsAsyncResponseContext structure represents FRS_ASYNC_RESPONSE_CONTEXT RPC structure.
type FrsAsyncResponseContext struct {
	SequenceNumber uint32                         `idl:"name:sequenceNumber" json:"sequence_number"`
	Status         uint32                         `idl:"name:status" json:"status"`
	Result         *FrsAsyncVersionVectorResponse `idl:"name:result" json:"result"`
}

func (o *FrsAsyncResponseContext) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *FrsAsyncResponseContext) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
		if err := (&FrsAsyncVersionVectorResponse{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *FrsAsyncResponseContext) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Result = &FrsAsyncVersionVectorResponse{}
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
