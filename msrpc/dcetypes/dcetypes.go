// The dcetypes package implements the DCETYPES client protocol.
package dcetypes

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
	GoPackage = "dcetypes"
)

// NDRCIntBigEndian represents the ndr_c_int_big_endian RPC constant
var NDRCIntBigEndian = 0

// NDRCIntLittleEndian represents the ndr_c_int_little_endian RPC constant
var NDRCIntLittleEndian = 1

// NDRCFloatIeee represents the ndr_c_float_ieee RPC constant
var NDRCFloatIeee = 0

// NDRCFloatVax represents the ndr_c_float_vax RPC constant
var NDRCFloatVax = 1

// NDRCFloatCray represents the ndr_c_float_cray RPC constant
var NDRCFloatCray = 2

// NDRCFloatIBM represents the ndr_c_float_ibm RPC constant
var NDRCFloatIBM = 3

// NDRCCharASCII represents the ndr_c_char_ascii RPC constant
var NDRCCharASCII = 0

// NDRCCharEbcdic represents the ndr_c_char_ebcdic RPC constant
var NDRCCharEbcdic = 1

// Tower structure represents twr_t RPC structure.
type Tower struct {
	TowerLength      uint32 `idl:"name:tower_length" json:"tower_length"`
	TowerOctetString []byte `idl:"name:tower_octet_string;size_is:(tower_length)" json:"tower_octet_string"`
}

func (o *Tower) xxx_PreparePayload(ctx context.Context) error {
	if o.TowerOctetString != nil && o.TowerLength == 0 {
		o.TowerLength = uint32(len(o.TowerOctetString))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *Tower) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.TowerLength)
	return []uint64{
		dimSize1,
	}
}
func (o *Tower) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.TowerLength); err != nil {
		return err
	}
	for i1 := range o.TowerOctetString {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.TowerOctetString[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.TowerOctetString); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Tower) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.TowerLength); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.TowerLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.TowerLength)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.TowerOctetString", sizeInfo[0])
	}
	o.TowerOctetString = make([]byte, sizeInfo[0])
	for i1 := range o.TowerOctetString {
		i1 := i1
		if err := w.ReadData(&o.TowerOctetString[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ContextHandle structure represents ndr_context_handle RPC structure.
type ContextHandle struct {
	Attributes uint32     `idl:"name:context_handle_attributes" json:"attributes"`
	UUID       *dtyp.GUID `idl:"name:context_handle_uuid" json:"uuid"`
}

func (o *ContextHandle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextHandle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextHandle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NDRFormat structure represents ndr_format_t RPC structure.
type NDRFormat struct {
	IntRep   uint8 `idl:"name:int_rep" json:"int_rep"`
	CharRep  uint8 `idl:"name:char_rep" json:"char_rep"`
	FloatRep uint8 `idl:"name:float_rep" json:"float_rep"`
	_        uint8 `idl:"name:reserved"`
}

func (o *NDRFormat) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NDRFormat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteData(o.IntRep); err != nil {
		return err
	}
	if err := w.WriteData(o.CharRep); err != nil {
		return err
	}
	if err := w.WriteData(o.FloatRep); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	return nil
}
func (o *NDRFormat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.IntRep); err != nil {
		return err
	}
	if err := w.ReadData(&o.CharRep); err != nil {
		return err
	}
	if err := w.ReadData(&o.FloatRep); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint8
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	return nil
}

// InterfaceID structure represents rpc_if_id_t RPC structure.
type InterfaceID struct {
	UUID      *dtyp.GUID `idl:"name:uuid" json:"uuid"`
	VersMajor uint16     `idl:"name:vers_major" json:"vers_major"`
	VersMinor uint16     `idl:"name:vers_minor" json:"vers_minor"`
}

func (o *InterfaceID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VersMajor); err != nil {
		return err
	}
	if err := w.WriteData(o.VersMinor); err != nil {
		return err
	}
	return nil
}
func (o *InterfaceID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersMajor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VersMinor); err != nil {
		return err
	}
	return nil
}

// InterfaceIDVector structure represents rpc_if_id_vector_t RPC structure.
type InterfaceIDVector struct {
	Count       uint32         `idl:"name:count" json:"count"`
	InterfaceID []*InterfaceID `idl:"name:if_id;size_is:(count)" json:"interface_id"`
}

func (o *InterfaceIDVector) xxx_PreparePayload(ctx context.Context) error {
	if o.InterfaceID != nil && o.Count == 0 {
		o.Count = uint32(len(o.InterfaceID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *InterfaceIDVector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Count)
	return []uint64{
		dimSize1,
	}
}
func (o *InterfaceIDVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	for i1 := range o.InterfaceID {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.InterfaceID[i1] != nil {
			_ptr_if_id := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceID[i1] != nil {
					if err := o.InterfaceID[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InterfaceID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceID[i1], _ptr_if_id); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.InterfaceID); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InterfaceIDVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Count > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Count)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.InterfaceID", sizeInfo[0])
	}
	o.InterfaceID = make([]*InterfaceID, sizeInfo[0])
	for i1 := range o.InterfaceID {
		i1 := i1
		_ptr_if_id := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceID[i1] == nil {
				o.InterfaceID[i1] = &InterfaceID{}
			}
			if err := o.InterfaceID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_if_id := func(ptr interface{}) { o.InterfaceID[i1] = *ptr.(**InterfaceID) }
		if err := w.ReadPointer(&o.InterfaceID[i1], _s_if_id, _ptr_if_id); err != nil {
			return err
		}
	}
	return nil
}

// StatsVector structure represents rpc_stats_vector_t RPC structure.
type StatsVector struct {
	Count uint32   `idl:"name:count" json:"count"`
	Stats []uint32 `idl:"name:stats" json:"stats"`
}

func (o *StatsVector) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatsVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	for i1 := range o.Stats {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Stats[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Stats); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatsVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	o.Stats = make([]uint32, 1)
	for i1 := range o.Stats {
		i1 := i1
		if err := w.ReadData(&o.Stats[i1]); err != nil {
			return err
		}
	}
	return nil
}
