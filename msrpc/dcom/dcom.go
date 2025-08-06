// The dcom package implements the DCOM client protocol.
//
// # Introduction
//
// The Distributed Component Object Model (DCOM) Remote Protocol is a protocol for exposing
// application objects by way of remote procedure calls (RPCs). The protocol consists
// of a set of extensions layered on Microsoft Remote Procedure Call Protocol Extensions
// as specified in [MS-RPCE].
//
// # Overview
//
// The Distributed Component Object Model (DCOM) Remote Protocol extends the Component
// Object Model (COM) over a network by providing facilities for creating and activating
// objects, and for managing object references, object lifetimes, and object interface
// queries. The DCOM Remote Protocol is built on top of Remote Procedure Call Protocol
// Extensions, as specified in [MS-RPCE], and relies on its authentication, authorization,
// and message integrity capabilities. The DCOM Remote Protocol is also referred to
// as Object RPC or ORPC. The following diagram shows the layering of the protocol stack.
package dcom

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
	GoPackage = "dcom"
)

// MinActpropLimit represents the MIN_ACTPROP_LIMIT RPC constant
const MinActpropLimit = 0x00000001

// MaxActpropLimit represents the MAX_ACTPROP_LIMIT RPC constant
const MaxActpropLimit = 0x0000000A

// ClassID structure represents CLSID RPC structure.
type ClassID dtyp.GUID

func (o *ClassID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ClassID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ClassID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClassID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IID structure represents IID RPC structure.
type IID dtyp.GUID

func (o *IID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *IID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// IPID structure represents IPID RPC structure.
type IPID dtyp.GUID

func (o *IPID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *IPID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *IPID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *IPID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CID structure represents CID RPC structure.
type CID dtyp.GUID

func (o *CID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *CID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// COMVersion structure represents COMVERSION RPC structure.
//
// The COMVERSION structure is used to specify the major and minor version of either
// the client or the server DCOM Remote Protocol implementation.
type COMVersion struct {
	// MajorVersion:  This MUST contain the major version of the DCOM Remote Protocol.
	MajorVersion uint16 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion:  This MUST contain the minor version of the DCOM Remote Protocol.
	//
	// The following table specifies the capabilities introduced in each DCOM version.<4>
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| VERSION |                                      CHANGE                                      |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.1 | Initial DCOM Remote Protocol release.                                            |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.2 | Added ResolveOxid2 to the IObjectExporter interface; see section 3.1.2.5.1.5.    |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.3 | MUST NOT be used.                                                                |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.4 | Update in the marshaling of arrays of interface pointers. Update in the          |
	//	|         | marshaling of conformant embedded structures.                                    |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.5 | Unused. This is to avoid having a DCOM version with matching major and minor     |
	//	|         | version numbers.                                                                 |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.6 | Added OBJREF_HANDLER and OBJREF_EXTENDED to the OBJREF type. Added               |
	//	|         | IRemoteSCMActivator interface methods (see section 3.1.2.5.2.2). Added           |
	//	|         | IObjectExporter::ServerAlive2 (Opnum 5) method to IObjectExporter interface.     |
	//	|         | Added IRemUnknown2 interface.                                                    |
	//	+---------+----------------------------------------------------------------------------------+
	//	|     5.7 | No DCOM changes from 5.6.<5>                                                     |
	//	+---------+----------------------------------------------------------------------------------+
	MinorVersion uint16 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *COMVersion) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *COMVersion) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	return nil
}
func (o *COMVersion) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	return nil
}

// ORPCExtent structure represents ORPC_EXTENT RPC structure.
//
// ORPC_EXTENT is a binary large object (BLOB) of data whose format is identified by
// a GUID. It is used on DCOM Remote Protocol calls to pass arbitrary out-of-band data
// that is not part of the explicit method signature. Unless otherwise specified, clients
// and servers MUST ignore ORPC_EXTENTs whose format they do not recognize.<6>
type ORPCExtent struct {
	// id:  This MUST contain a GUID that identifies the format of the opaque data in the
	// data field.
	ID *dtyp.GUID `idl:"name:id" json:"id"`
	// size:  This MUST specify the size, in bytes, in the data field excluding any padding
	// bytes that were added to round the array size to a multiple of 8.
	Size uint32 `idl:"name:size" json:"size"`
	// data:  This MUST contain an array of bytes that form the extent data. The array size
	// MUST be a multiple of 8 for alignment reasons.
	Data []byte `idl:"name:data;size_is:(((size+7)&(^uint32(7))))" json:"data"`
}

func (o *ORPCExtent) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ORPCExtent) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(((o.Size + 7) & (^uint32(7))))
	return []uint64{
		dimSize1,
	}
}
func (o *ORPCExtent) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.ID != nil {
		if err := o.ID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Size); err != nil {
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
func (o *ORPCExtent) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if o.ID == nil {
		o.ID = &dtyp.GUID{}
	}
	if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if ((o.Size+7)&(^uint32(7))) > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(((o.Size + 7) & (^uint32(7))))
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

// ORPCExtentArray structure represents ORPC_EXTENT_ARRAY RPC structure.
//
// ORPC_EXTENT_ARRAY is an array of ORPC_EXTENT structures.
type ORPCExtentArray struct {
	// size:   This MUST specify the number of non-NULL elements in the extent field.
	Size uint32 `idl:"name:size" json:"size"`
	// reserved:  This MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:reserved"`
	// extent:   This MUST be an array of ORPC_EXTENTs. The array size MUST be a multiple
	// of 2 for alignment reasons.
	Extent []*ORPCExtent `idl:"name:extent;size_is:(((size+1)&(^uint32(1))), );pointer:unique" json:"extent"`
}

func (o *ORPCExtentArray) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Extent != nil && o.Size == 0 {
		o.Size = uint32(len(o.Extent))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ORPCExtentArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.Extent != nil || ((o.Size+1)&(^uint32(1))) > 0 {
		_ptr_extent := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(((o.Size + 1) & (^uint32(1))))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Extent {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Extent[i1] != nil {
					_ptr_extent := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Extent[i1] != nil {
							if err := o.Extent[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&ORPCExtent{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Extent[i1], _ptr_extent); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Extent); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extent, _ptr_extent); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ORPCExtentArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint32
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	_ptr_extent := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if ((o.Size+1)&(^uint32(1))) > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(((o.Size + 1) & (^uint32(1))))
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Extent", sizeInfo[0])
		}
		o.Extent = make([]*ORPCExtent, sizeInfo[0])
		for i1 := range o.Extent {
			i1 := i1
			_ptr_extent := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Extent[i1] == nil {
					o.Extent[i1] = &ORPCExtent{}
				}
				if err := o.Extent[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_extent := func(ptr interface{}) { o.Extent[i1] = *ptr.(**ORPCExtent) }
			if err := w.ReadPointer(&o.Extent[i1], _s_extent, _ptr_extent); err != nil {
				return err
			}
		}
		return nil
	})
	_s_extent := func(ptr interface{}) { o.Extent = *ptr.(*[]*ORPCExtent) }
	if err := w.ReadPointer(&o.Extent, _s_extent, _ptr_extent); err != nil {
		return err
	}
	return nil
}

// ORPCThis structure represents ORPCTHIS RPC structure.
//
// The ORPCTHIS structure is the first (implicit) argument sent in an ORPC request PDU
// and is used to send ORPC extension data to the server. The ORPCTHIS structure is
// also sent as an explicit argument in activation RPC requests.
type ORPCThis struct {
	// version:   A COMVERSION structure that MUST contain the version number of the client.
	// For details, see section 2.2.11.
	Version *COMVersion `idl:"name:version" json:"version"`
	// flags:   When the ORPCTHIS structure is used as a parameter in ORPC invocations (as
	// specified in section 3.2.4.2), this MUST be set to 0x00000000. When the ORPCTHIS
	// structure is used as a parameter in IActivation::RemoteActivation, IRemoteSCMActivator::RemoteGetClassObject
	// and IRemoteSCMActivator::RemoteCreateInstance methods (section 3.1.2.5.2.3), this
	// can be set to any arbitrary value when sent and MUST be ignored on receipt.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// reserved1:   This MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:reserved1"`
	// cid:  This MUST contain a CID for the ORPC call. For details, see section 1.3.5.
	CID *CID `idl:"name:cid" json:"cid"`
	// extensions:   If non-NULL, this MUST be a pointer to an ORPC_EXTENT_ARRAY structure.
	Extensions *ORPCExtentArray `idl:"name:extensions;pointer:unique" json:"extensions"`
}

func (o *ORPCThis) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ORPCThis) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Version != nil {
		if err := o.Version.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&COMVersion{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved reserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.CID != nil {
		if err := o.CID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Extensions != nil {
		_ptr_extensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extensions != nil {
				if err := o.Extensions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ORPCExtentArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extensions, _ptr_extensions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ORPCThis) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Version == nil {
		o.Version = &COMVersion{}
	}
	if err := o.Version.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved reserved1
	var _reserved1 uint32
	if err := w.ReadData(&_reserved1); err != nil {
		return err
	}
	if o.CID == nil {
		o.CID = &CID{}
	}
	if err := o.CID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_extensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extensions == nil {
			o.Extensions = &ORPCExtentArray{}
		}
		if err := o.Extensions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_extensions := func(ptr interface{}) { o.Extensions = *ptr.(**ORPCExtentArray) }
	if err := w.ReadPointer(&o.Extensions, _s_extensions, _ptr_extensions); err != nil {
		return err
	}
	return nil
}

// ORPCThat structure represents ORPCTHAT RPC structure.
//
// The ORPCTHAT structure is the first (implicit) argument returned in an ORPC response
// PDU, and is used to return ORPC extension data to the client. The ORPCTHAT structure
// is also returned as an explicit argument from an activation request.
type ORPCThat struct {
	// flags:  This can be set to any arbitrary value and MUST be ignored on receipt.
	Flags uint32 `idl:"name:flags" json:"flags"`
	// extensions:   If non-NULL, this field MUST contain an ORPC_EXTENT_ARRAY.
	Extensions *ORPCExtentArray `idl:"name:extensions;pointer:unique" json:"extensions"`
}

func (o *ORPCThat) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ORPCThat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Extensions != nil {
		_ptr_extensions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Extensions != nil {
				if err := o.Extensions.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ORPCExtentArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Extensions, _ptr_extensions); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ORPCThat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_extensions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Extensions == nil {
			o.Extensions = &ORPCExtentArray{}
		}
		if err := o.Extensions.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_extensions := func(ptr interface{}) { o.Extensions = *ptr.(**ORPCExtentArray) }
	if err := w.ReadPointer(&o.Extensions, _s_extensions, _ptr_extensions); err != nil {
		return err
	}
	return nil
}

// DualStringArray structure represents DUALSTRINGARRAY RPC structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| wNumEntries                                                   | wSecurityOffset                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| StringBinding (variable)                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nullterm1                                                     | SecBinding (variable)                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nullterm2                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// In certain cases in the DCOM Remote Protocol, a DUALSTRINGARRAY is passed or returned
// as a parameter in an RPC call. For example, see section 3.1.2.5.2.3.1. In all such
// cases, the IDL definition that the DCOM Remote Protocol uses is as follows.
type DualStringArray struct {
	// wNumEntries (2 bytes):  The (unsigned) number of unsigned shorts (that is, 2-octet
	// units) from the first entry in the StringBinding array to the end of the buffer.
	//
	// wNumEntries:   This MUST be set to the number of unsigned shorts in the aStringArray
	// field.
	EntriesLength uint16 `idl:"name:wNumEntries" json:"entries_length"`
	// wSecurityOffset (2 bytes):  The (unsigned) number of unsigned shorts from the first
	// entry in the StringBinding array to the first entry in the SecBinding array.
	//
	// wSecurityOffset:   This MUST be set to the number of unsigned shorts from the beginning
	// of the aStringArray array to the beginning of the first security binding within the
	// array. For details, see section 2.2.19.1.
	SecurityOffset uint16 `idl:"name:wSecurityOffset" json:"security_offset"`
	// aStringArray:  This MUST be an array of wNumEntries unsigned shorts. This field MUST
	// be interpreted to contain a sequence of STRINGBINDING entries followed by a sequence
	// of SECURITYBINDING entries, in the same syntax as defined in section 2.2.19.1.
	StringArray []uint16 `idl:"name:aStringArray;size_is:(wNumEntries)" json:"string_array"`
}

func (o *DualStringArray) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.StringArray != nil && o.EntriesLength == 0 {
		o.EntriesLength = uint16(len(o.StringArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DualStringArray) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntriesLength)
	return []uint64{
		dimSize1,
	}
}
func (o *DualStringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesLength); err != nil {
		return err
	}
	if err := w.WriteData(o.SecurityOffset); err != nil {
		return err
	}
	for i1 := range o.StringArray {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.StringArray[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.StringArray); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *DualStringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecurityOffset); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.EntriesLength > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.EntriesLength)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.StringArray", sizeInfo[0])
	}
	o.StringArray = make([]uint16, sizeInfo[0])
	for i1 := range o.StringArray {
		i1 := i1
		if err := w.ReadData(&o.StringArray[i1]); err != nil {
			return err
		}
	}
	return nil
}

// InterfacePointer structure represents MInterfacePointer RPC structure.
//
// MInterfacePointer is an NDR-marshaled structure that MUST contain a hand-marshaled
// OBJREF.
type InterfacePointer struct {
	// ulCntData:  This MUST specify the size, in bytes, of the abData parameter.
	DataCount uint32 `idl:"name:ulCntData" json:"data_count"`
	// abData:   An array of bytes that MUST contain an OBJREF.
	Data []byte `idl:"name:abData;size_is:(ulCntData)" json:"data"`
}

func (o *InterfacePointer) xxx_PreparePayload(ctx context.Context) error {
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

func (o *InterfacePointer) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *InterfacePointer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *InterfacePointer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ErrorObjectData structure represents ErrorObjectData RPC structure.
//
// This section defines the format of an OBJREF_CUSTOM that, depending on the use of
// the DCOM Remote Protocol by an application or a higher-layer protocol, MAY be passed
// as an error information ORPC extension (see section 2.2.21.1). CLSID_ErrorObject
// (see section 1.9) is the unmarshaler CLSID for this OBJREF_CUSTOM. The format of
// the pObjectData field of the OBJREF_CUSTOM for this CLSID is as follows.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwVersion                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwHelpContext                                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| iid (16 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwSourceSignature                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Source (variable)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwDescriptionSignature                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Description (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwHelpFileSignature                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| HelpFile (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ErrorObjectData struct {
	// dwVersion (4 bytes): This MUST be set to 0x00000000.
	Version uint32 `idl:"name:dwVersion" json:"version"`
	// dwHelpContext (4 bytes): An implementation-specific value that SHOULD be ignored
	// on receipt.<16>
	HelpContext uint32 `idl:"name:dwHelpContext" json:"help_context"`
	// iid (16 bytes): An IID that MUST be the IID of the interface returning the error.
	IID *IID `idl:"name:iid" json:"iid"`
	// dwSourceSignature (4 bytes): This MUST be set to 0xFFFFFFFF if the Source field is
	// present and MUST be set to 0x00000000 if the Source field is not present.
	SourceSignature uint32 `idl:"name:dwSourceSignature" json:"source_signature"`
	// Source (variable): This MUST contain an ErrorInfoString (see ErrorInfoString (section
	// 2.2.21.3)) if the dwSourceSignature field is set to 0xFFFFFFFF and MUST NOT be present
	// if the dwSourceSignature field is set to 0x00000000.<17>
	Source *ErrorObjectDataString `idl:"name:Source;switch_is:dwSourceSignature" json:"source"`
	// dwDescriptionSignature (4 bytes): This MUST be set to 0xFFFFFFFF if the Description
	// field is present and that MUST be set to 0x00000000 if the Description field is not
	// present.
	DescriptionSignature uint32 `idl:"name:dwDescriptionSignature" json:"description_signature"`
	// Description (variable): This MUST contain an ErrorInfoString (see ErrorInfoString
	// (section 2.2.21.3)) if the dwDescriptionSignature field is set to 0xFFFFFFFF and
	// MUST NOT be present if the dwDescriptionSignature field is set to 0x00000000.<18>
	Description *ErrorObjectDataString `idl:"name:Description;switch_is:dwDescriptionSignature" json:"description"`
	// dwHelpFileSignature (4 bytes): This MUST be set to 0xFFFFFFFF if the HelpFile field
	// is present and MUST be set to 0x00000000 if the HelpFile field is not present.
	HelpFileSignature uint32 `idl:"name:dwHelpFileSignature" json:"help_file_signature"`
	// HelpFile (variable): This MUST contain an ErrorInfoString (see ErrorInfoString (section
	// 2.2.21.3)) if the dwHelpFileSignature field is set to 0xFFFFFFFF and MUST NOT be
	// present if the dwHelpFileSignature field is set to 0x00000000.<19>
	HelpFile *ErrorObjectDataString `idl:"name:HelpFile;switch_is:dwHelpFileSignature" json:"help_file"`
}

func (o *ErrorObjectData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ErrorObjectData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.HelpContext); err != nil {
		return err
	}
	if o.IID != nil {
		if err := o.IID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SourceSignature); err != nil {
		return err
	}
	if o.Source != nil {
		_ptr_Source := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swSource := uint32(o.SourceSignature)
			if o.Source != nil {
				if err := o.Source.MarshalUnionNDR(ctx, w, _swSource); err != nil {
					return err
				}
			} else {
				if err := (&ErrorObjectDataString{}).MarshalUnionNDR(ctx, w, _swSource); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Source, _ptr_Source); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DescriptionSignature); err != nil {
		return err
	}
	if o.Description != nil {
		_ptr_Description := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swDescription := uint32(o.DescriptionSignature)
			if o.Description != nil {
				if err := o.Description.MarshalUnionNDR(ctx, w, _swDescription); err != nil {
					return err
				}
			} else {
				if err := (&ErrorObjectDataString{}).MarshalUnionNDR(ctx, w, _swDescription); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_Description); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HelpFileSignature); err != nil {
		return err
	}
	if o.HelpFile != nil {
		_ptr_HelpFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			_swHelpFile := uint32(o.HelpFileSignature)
			if o.HelpFile != nil {
				if err := o.HelpFile.MarshalUnionNDR(ctx, w, _swHelpFile); err != nil {
					return err
				}
			} else {
				if err := (&ErrorObjectDataString{}).MarshalUnionNDR(ctx, w, _swHelpFile); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HelpFile, _ptr_HelpFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ErrorObjectData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.HelpContext); err != nil {
		return err
	}
	if o.IID == nil {
		o.IID = &IID{}
	}
	if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.SourceSignature); err != nil {
		return err
	}
	_ptr_Source := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Source == nil {
			o.Source = &ErrorObjectDataString{}
		}
		_swSource := uint32(o.SourceSignature)
		if err := o.Source.UnmarshalUnionNDR(ctx, w, _swSource); err != nil {
			return err
		}
		return nil
	})
	_s_Source := func(ptr interface{}) { o.Source = *ptr.(**ErrorObjectDataString) }
	if err := w.ReadPointer(&o.Source, _s_Source, _ptr_Source); err != nil {
		return err
	}
	if err := w.ReadData(&o.DescriptionSignature); err != nil {
		return err
	}
	_ptr_Description := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Description == nil {
			o.Description = &ErrorObjectDataString{}
		}
		_swDescription := uint32(o.DescriptionSignature)
		if err := o.Description.UnmarshalUnionNDR(ctx, w, _swDescription); err != nil {
			return err
		}
		return nil
	})
	_s_Description := func(ptr interface{}) { o.Description = *ptr.(**ErrorObjectDataString) }
	if err := w.ReadPointer(&o.Description, _s_Description, _ptr_Description); err != nil {
		return err
	}
	if err := w.ReadData(&o.HelpFileSignature); err != nil {
		return err
	}
	_ptr_HelpFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.HelpFile == nil {
			o.HelpFile = &ErrorObjectDataString{}
		}
		_swHelpFile := uint32(o.HelpFileSignature)
		if err := o.HelpFile.UnmarshalUnionNDR(ctx, w, _swHelpFile); err != nil {
			return err
		}
		return nil
	})
	_s_HelpFile := func(ptr interface{}) { o.HelpFile = *ptr.(**ErrorObjectDataString) }
	if err := w.ReadPointer(&o.HelpFile, _s_HelpFile, _ptr_HelpFile); err != nil {
		return err
	}
	return nil
}

// ErrorObjectDataString structure represents ErrorObjectDataString RPC union.
type ErrorObjectDataString struct {
	// Types that are assignable to Value
	//
	// *ErrorObjectDataString_String
	Value is_ErrorObjectDataString `json:"value"`
}

func (o *ErrorObjectDataString) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ErrorObjectDataString_String:
		if value != nil {
			return value.String
		}
	}
	return nil
}

type is_ErrorObjectDataString interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ErrorObjectDataString()
}

func (o *ErrorObjectDataString) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ErrorObjectDataString_String:
		return uint32(4294967295)
	}
	return uint32(0)
}

func (o *ErrorObjectDataString) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(4294967295):
		_o, _ := o.Value.(*ErrorObjectDataString_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ErrorObjectDataString_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ErrorObjectDataString) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(4294967295):
		o.Value = &ErrorObjectDataString_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ErrorObjectDataString_String structure represents ErrorObjectDataString RPC union arm.
//
// It has following labels: 4294967295
type ErrorObjectDataString_String struct {
	String *ErrorInfoString `idl:"name:String" json:"string"`
}

func (*ErrorObjectDataString_String) is_ErrorObjectDataString() {}

func (o *ErrorObjectDataString_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != nil {
		_ptr_String := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.String != nil {
				if err := o.String.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ErrorInfoString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_String); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ErrorObjectDataString_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_String := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.String == nil {
			o.String = &ErrorInfoString{}
		}
		if err := o.String.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_String := func(ptr interface{}) { o.String = *ptr.(**ErrorInfoString) }
	if err := w.ReadPointer(&o.String, _s_String, _ptr_String); err != nil {
		return err
	}
	return nil
}

// ErrorInfoString structure represents ErrorInfoString RPC structure.
//
// This packet specifies the format of the string data that is contained in a Custom-Marshaled
// Error Info Format (section 2.2.21.2) packet.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwMax                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwOffSet                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwActual                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Name (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ErrorInfoString struct {
	// dwMax (4 bytes): The unsigned number of Unicode characters in the Name array including
	// the NULL terminator.
	Max uint32 `idl:"name:dwMax" json:"max"`
	// dwOffSet (4 bytes): This MUST be set to zero.
	OffSet uint32 `idl:"name:dwOffSet" json:"off_set"`
	// dwActual (4 bytes): This MUST be set to the value of the dwMax field.
	Actual uint32 `idl:"name:dwActual" json:"actual"`
	// Name (variable):  This MUST contain an implementation-specific NULL-terminated Unicode
	// string and SHOULD be ignored on receipt.
	Name string `idl:"name:Name;size_is:(dwMax);string;pointer:unique" json:"name"`
}

func (o *ErrorInfoString) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Name != "" && o.Max == 0 {
		o.Max = uint32(ndr.UTF16NLen(o.Name))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ErrorInfoString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Max); err != nil {
		return err
	}
	if err := w.WriteData(o.OffSet); err != nil {
		return err
	}
	if err := w.WriteData(o.Actual); err != nil {
		return err
	}
	if o.Name != "" || o.Max > 0 {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Max)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.Name)
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			_Name_buf := utf16.Encode([]rune(o.Name))
			if uint64(len(_Name_buf)) > sizeInfo[0]-1 {
				_Name_buf = _Name_buf[:sizeInfo[0]-1]
			}
			if o.Name != ndr.ZeroString {
				_Name_buf = append(_Name_buf, uint16(0))
			}
			for i1 := range _Name_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Name_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Name_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_Name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ErrorInfoString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Max); err != nil {
		return err
	}
	if err := w.ReadData(&o.OffSet); err != nil {
		return err
	}
	if err := w.ReadData(&o.Actual); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Name_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Name_buf", sizeInfo[0])
		}
		_Name_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Name_buf {
			i1 := i1
			if err := w.ReadData(&_Name_buf[i1]); err != nil {
				return err
			}
		}
		o.Name = strings.TrimRight(string(utf16.Decode(_Name_buf)), ndr.ZeroString)
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// ContextORPCExtension structure represents ContextORPCExtension RPC structure.
//
// CLSID_CONTEXT_EXTENSION is the identifying GUID for this ORPC extension, which is
// marshaled into the data field of an ORPC_EXTENT using the structure format that follows.
//
// A context ORPC extension is used when a client or server passes data associated with
// a context property along with and in response to an ORPC call.
//
// All fields MUST be marshaled with the same endianness as the encompassing RPC PDU.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Signature                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Version                                                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cPolicies                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbBuffer                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbSize                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| hr                                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| hrServer                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| reserved                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| EntryHeader (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| PolicyData (variable)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ContextORPCExtension struct {
	// Signature (4 bytes):  This MUST be set to 0x414E554B.
	Signature []byte `idl:"name:Signature" json:"signature"`
	// Version (4 bytes):  This MUST be set to 0x00010000.
	Version uint32 `idl:"name:Version" json:"version"`
	// cPolicies (4 bytes): The unsigned number of elements in the EntryHeader array. This
	// MUST also be equal to the number of elements in the PolicyData array.
	PoliciesCount uint32 `idl:"name:cPolicies" json:"policies_count"`
	// cbBuffer (4 bytes):  An implementation-specific value that MUST be ignored on receipt.<20>
	BufferLength uint32 `idl:"name:cbBuffer" json:"buffer_length"`
	// cbSize (4 bytes): The unsigned size (in bytes) from the beginning of this structure
	// to the end of the last EntryHeader.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// hr (4 bytes):  An HRESULT that MUST be set to zero and MUST be ignored on receipt.
	HResult uint32 `idl:"name:hr" json:"hresult"`
	// hrServer (4 bytes):  An HRESULT that MUST be set to zero when sent from client to
	// server and MUST be ignored on receipt. When sent from server to client, this field
	// MAY contain an implementation-specific error value.<21>
	ServerHResult uint32 `idl:"name:hrServer" json:"server_hresult"`
	// reserved (4 bytes):  This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:reserved"`
	// EntryHeader (variable):  An array of EntryHeader structures. The number of elements
	// in the array MUST be specified in cPolicies.
	EntryHeader []*EntryHeader `idl:"name:EntryHeader;size_is:(cPolicies)" json:"entry_header"`
	// PolicyData (variable):  An array of byte arrays. The number of elements in the outer
	// array MUST be specified in cPolicies. The number of elements in the byte array MUST
	// be specified in the cbEHBuffer field (see section 2.2.21.5) in the corresponding
	// element in the EntryHeader array.
	//
	// Both the EntryHeader and PolicyData arrays MUST have sizes that are integral multiples
	// of eight, and all padding bytes MUST be initialized to zero and MUST be ignored on
	// receipt.
	PolicyData []byte `idl:"name:PolicyData" json:"policy_data"`
}

func (o *ContextORPCExtension) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.EntryHeader != nil && o.PoliciesCount == 0 {
		o.PoliciesCount = uint32(len(o.EntryHeader))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ContextORPCExtension) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	for i1 := range o.Signature {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Signature[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Signature); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.PoliciesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.HResult); err != nil {
		return err
	}
	if err := w.WriteData(o.ServerHResult); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.EntryHeader != nil || o.PoliciesCount > 0 {
		_ptr_EntryHeader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.PoliciesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.EntryHeader {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.EntryHeader[i1] != nil {
					if err := o.EntryHeader[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EntryHeader{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.EntryHeader); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&EntryHeader{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryHeader, _ptr_EntryHeader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.PolicyData != nil {
		_ptr_PolicyData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.PolicyData {
				i1 := i1
				if err := w.WriteData(o.PolicyData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PolicyData, _ptr_PolicyData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextORPCExtension) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	o.Signature = make([]byte, 4)
	for i1 := range o.Signature {
		i1 := i1
		if err := w.ReadData(&o.Signature[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.PoliciesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.HResult); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServerHResult); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint32
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	_ptr_EntryHeader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.PoliciesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.PoliciesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.EntryHeader", sizeInfo[0])
		}
		o.EntryHeader = make([]*EntryHeader, sizeInfo[0])
		for i1 := range o.EntryHeader {
			i1 := i1
			if o.EntryHeader[i1] == nil {
				o.EntryHeader[i1] = &EntryHeader{}
			}
			if err := o.EntryHeader[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_EntryHeader := func(ptr interface{}) { o.EntryHeader = *ptr.(*[]*EntryHeader) }
	if err := w.ReadPointer(&o.EntryHeader, _s_EntryHeader, _ptr_EntryHeader); err != nil {
		return err
	}
	_ptr_PolicyData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.PolicyData = append(o.PolicyData, uint8(0))
			if err := w.ReadData(&o.PolicyData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_PolicyData := func(ptr interface{}) { o.PolicyData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.PolicyData, _s_PolicyData, _ptr_PolicyData); err != nil {
		return err
	}
	return nil
}

// EntryHeader structure represents EntryHeader RPC structure.
//
// The EntryHeader structure is used to describe an opaque array of bytes associated
// with a context property within the context ORPC extension.
//
// All fields MUST be marshaled with the same endianness as the encompassing RPC PDU.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Signature                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbEHBuffer                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbSize                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| reserved                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| policyID (16 bytes)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type EntryHeader struct {
	// Signature (4 bytes): This MUST be set to 0x494E414E.
	Signature []byte `idl:"name:Signature" json:"signature"`
	// cbEHBuffer (4 bytes):  The unsigned size (in bytes) of the opaque policy data buffer
	// corresponding to this EntryHeader. This MUST NOT be zero.
	BufferLength uint32 `idl:"name:cbEHBuffer" json:"buffer_length"`
	// cbSize (4 bytes): The unsigned size (in bytes) of the offset from the beginning of
	// the context ORPC extension buffer to the beginning of the opaque policy data buffer
	// corresponding to this EntryHeader.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// reserved (4 bytes):  This MUST be set to 0x00000000 and MUST be ignored on receipt.
	_ uint32 `idl:"name:reserved"`
	// policyID (16 bytes):  A GUID that MUST contain a context property identifier.
	PolicyID *dtyp.GUID `idl:"name:policyID" json:"policy_id"`
}

func (o *EntryHeader) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *EntryHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Signature {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Signature[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Signature); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	// reserved reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.PolicyID != nil {
		if err := o.PolicyID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *EntryHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Signature = make([]byte, 4)
	for i1 := range o.Signature {
		i1 := i1
		if err := w.ReadData(&o.Signature[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	// reserved reserved
	var _reserved uint32
	if err := w.ReadData(&_reserved); err != nil {
		return err
	}
	if o.PolicyID == nil {
		o.PolicyID = &dtyp.GUID{}
	}
	if err := o.PolicyID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// StdObjectReference structure represents STDOBJREF RPC structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cPublicRefs                                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| oxid                                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| oid                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ipid (16 bytes)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//
// The parameter meanings for this structure are identical to those defined in section
// 2.2.18.2.
type StdObjectReference struct {
	// flags (4 bytes):  This can be one of the following values. Any other value MUST
	// be ignored by the client.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|         VALUE          |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000             | The client is requested to perform garbage collection pinging (see section       |
	//	|                        | 3.2.6.1) for this object identifier (OID).                                       |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| SORF_NOPING 0x00001000 | The client is requested to not perform garbage collection pinging (see section   |
	//	|                        | 3.2.6.1) for this object identifier (OID).<7>                                    |
	//	+------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// cPublicRefs (4 bytes): The number of public references on the server object, which
	// MUST be released later. For more information, see section 3.2.4.4.2.
	PublicReferencesCount uint32 `idl:"name:cPublicRefs" json:"public_references_count"`
	// oxid (8 bytes):  This MUST be an OXID identifying the object exporter that contains
	// the object.
	OXID uint64 `idl:"name:oxid" json:"oxid"`
	// oid (8 bytes):  This MUST be an OID identifying the object.
	OID uint64 `idl:"name:oid" json:"oid"`
	// ipid (16 bytes):  This MUST be an IPID identifying a specific interface on the object.
	IPID *IPID `idl:"name:ipid" json:"ipid"`
}

func (o *StdObjectReference) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *StdObjectReference) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.PublicReferencesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.OXID); err != nil {
		return err
	}
	if err := w.WriteData(o.OID); err != nil {
		return err
	}
	if o.IPID != nil {
		if err := o.IPID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *StdObjectReference) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicReferencesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.OXID); err != nil {
		return err
	}
	if err := w.ReadData(&o.OID); err != nil {
		return err
	}
	if o.IPID == nil {
		o.IPID = &IPID{}
	}
	if err := o.IPID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ObjectReferenceType type represents OBJREF_TYPE RPC enumeration.
type ObjectReferenceType uint16

var (
	ObjectReferenceTypeStandard ObjectReferenceType = 1
	ObjectReferenceTypeHandler  ObjectReferenceType = 2
	ObjectReferenceTypeCustom   ObjectReferenceType = 4
	ObjectReferenceTypeExtended ObjectReferenceType = 8
)

func (o ObjectReferenceType) String() string {
	switch o {
	case ObjectReferenceTypeStandard:
		return "ObjectReferenceTypeStandard"
	case ObjectReferenceTypeHandler:
		return "ObjectReferenceTypeHandler"
	case ObjectReferenceTypeCustom:
		return "ObjectReferenceTypeCustom"
	case ObjectReferenceTypeExtended:
		return "ObjectReferenceTypeExtended"
	}
	return "Invalid"
}

// ObjectReferenceStandard structure represents OBJREF_STANDARD RPC structure.
//
// This form of OBJREF is the simplest, combining an STDOBJREF structure with a DUALSTRINGARRAY
// structure. It is used when there is no need to utilize the extra fields offered by
// the other OBJREF formats.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| std (40 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| saResAddr (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReferenceStandard struct {
	// std (40 bytes):  This MUST be an STDOBJREF.
	Std *StdObjectReference `idl:"name:std" json:"std"`
	// saResAddr (variable):  A DUALSTRINGARRAY that MUST contain the network and security
	// bindings for the object resolver service on the server.
	ResolverAddr *DualStringArray `idl:"name:saResAddr" json:"resolver_addr"`
}

func (o *ObjectReferenceStandard) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectReferenceStandard) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Std != nil {
		if err := o.Std.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StdObjectReference{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ResolverAddr != nil {
		_ptr_saResAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResolverAddr != nil {
				if err := o.ResolverAddr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DualStringArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResolverAddr, _ptr_saResAddr); err != nil {
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
func (o *ObjectReferenceStandard) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Std == nil {
		o.Std = &StdObjectReference{}
	}
	if err := o.Std.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_saResAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResolverAddr == nil {
			o.ResolverAddr = &DualStringArray{}
		}
		if err := o.ResolverAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_saResAddr := func(ptr interface{}) { o.ResolverAddr = *ptr.(**DualStringArray) }
	if err := w.ReadPointer(&o.ResolverAddr, _s_saResAddr, _ptr_saResAddr); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ObjectReferenceHandler structure represents OBJREF_HANDLER RPC structure.
//
// This form of OBJREF is used by the server object to provide an identifier for a helper
// object on the client, which the client can use as a proxy for the server object.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| std (40 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| clsid (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| saResAddr (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReferenceHandler struct {
	// std (40 bytes): This MUST specify an STDOBJREF.
	Std *StdObjectReference `idl:"name:std" json:"std"`
	// clsid (16 bytes): This MUST specify a CLSID identifying an object class on the client
	// that the client uses as a handler for the interface identified by the iid field of
	// the containing OBJREF.
	ClassID *ClassID `idl:"name:clsid" json:"class_id"`
	// saResAddr (variable):  This MUST specify a DUALSTRINGARRAY that MUST contain the
	// network and security bindings for the object resolver service on the server.
	ResolverAddr *DualStringArray `idl:"name:saResAddr" json:"resolver_addr"`
}

func (o *ObjectReferenceHandler) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectReferenceHandler) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Std != nil {
		if err := o.Std.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StdObjectReference{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ResolverAddr != nil {
		_ptr_saResAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResolverAddr != nil {
				if err := o.ResolverAddr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DualStringArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResolverAddr, _ptr_saResAddr); err != nil {
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
func (o *ObjectReferenceHandler) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Std == nil {
		o.Std = &StdObjectReference{}
	}
	if err := o.Std.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &ClassID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_saResAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResolverAddr == nil {
			o.ResolverAddr = &DualStringArray{}
		}
		if err := o.ResolverAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_saResAddr := func(ptr interface{}) { o.ResolverAddr = *ptr.(**DualStringArray) }
	if err := w.ReadPointer(&o.ResolverAddr, _s_saResAddr, _ptr_saResAddr); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ObjectReferenceCustom structure represents OBJREF_CUSTOM RPC structure.
//
// This form of OBJREF is used by a server object to marshal itself into an opaque BLOB
// using a custom marshaler. The custom marshaler is a COM object that can marshal and
// unmarshal the data contained in the BLOB. The CLSID of the custom marshaler object's
// object class is specified within the OBJREF.
//
// If the interface specified by the iid field of the OBJREF structure contained in
// the OBJREF_CUSTOM has the local IDL attribute (section 2.2.27), the OBJREF_CUSTOM
// MUST represent an object that is local to the client that unmarshals the object.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| clsid (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbExtension                                                                                                                   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| reserved                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| pObjectData (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReferenceCustom struct {
	// clsid (16 bytes): This MUST specify a CLSID, supplied by an application or higher-layer
	// protocol, identifying an object class associated with the data in the pObjectData
	// field.<8>
	ClassID *ClassID `idl:"name:clsid" json:"class_id"`
	// cbExtension (4 bytes): This MUST be set to zero when sent and MUST be ignored on
	// receipt.
	ExtensionLength uint32 `idl:"name:cbExtension" json:"extension_length"`
	Size            uint32 `idl:"name:size" json:"size"`
	// pObjectData (variable): This MUST be an array of bytes containing data supplied by
	// an application or higher-layer protocol.
	ObjectData []byte `idl:"name:pObjectData" json:"object_data"`
}

func (o *ObjectReferenceCustom) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectReferenceCustom) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ExtensionLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.ObjectData != nil {
		_ptr_pObjectData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			for i1 := range o.ObjectData {
				i1 := i1
				if err := w.WriteData(o.ObjectData[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectData, _ptr_pObjectData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReferenceCustom) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &ClassID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtensionLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pObjectData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		for i1 := 0; w.Len() > 0; i1++ {
			i1 := i1
			o.ObjectData = append(o.ObjectData, uint8(0))
			if err := w.ReadData(&o.ObjectData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pObjectData := func(ptr interface{}) { o.ObjectData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ObjectData, _s_pObjectData, _ptr_pObjectData); err != nil {
		return err
	}
	return nil
}

// ObjectReferenceExtended structure represents OBJREF_EXTENDED RPC structure.
//
// The OBJREF_EXTENDED format is used when the server returns a marshaled envoy context
// to the client.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| std (40 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Signature1                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| saResAddr (variable)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| nElms                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Signature2                                                                                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ElmArray (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReferenceExtended struct {
	// std (40 bytes):  This MUST contain an STDOBJREF structure.
	Std *StdObjectReference `idl:"name:std" json:"std"`
	// Signature1 (4 bytes): This MUST be set to 0x4E535956.
	Signature1 []byte `idl:"name:Signature1" json:"signature1"`
	// saResAddr (variable): This MUST contain a DUALSTRINGARRAY structure containing network
	// and security bindings for the object resolver service on the server.
	ResolverAddr *DualStringArray `idl:"name:saResAddr" json:"resolver_addr"`
	// nElms (4 bytes): The number of elements in the ElmArray field. This field MUST be
	// set to 0x00000001. (Note that while this protocol supports only a single element,
	// for historical reasons the protocol uses an array of one element.)
	Elements uint32 `idl:"name:nElms" json:"elements"`
	// Signature2 (4 bytes): This MUST be set to 0x4E535956.
	Signature2 []byte `idl:"name:Signature2" json:"signature2"`
	// ElmArray (variable): This MUST be a DATAELEMENT entry.
	ElementArray []*DataElement `idl:"name:ElmArray;size_is:(nElms)" json:"element_array"`
}

func (o *ObjectReferenceExtended) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ElementArray != nil && o.Elements == 0 {
		o.Elements = uint32(len(o.ElementArray))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectReferenceExtended) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.Std != nil {
		if err := o.Std.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StdObjectReference{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Signature1 {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Signature1[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Signature1); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.ResolverAddr != nil {
		_ptr_saResAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResolverAddr != nil {
				if err := o.ResolverAddr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DualStringArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResolverAddr, _ptr_saResAddr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Elements); err != nil {
		return err
	}
	for i1 := range o.Signature2 {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Signature2[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Signature2); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if o.ElementArray != nil || o.Elements > 0 {
		_ptr_ElmArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Elements)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ElementArray {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ElementArray[i1] != nil {
					if err := o.ElementArray[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DataElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ElementArray); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DataElement{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ElementArray, _ptr_ElmArray); err != nil {
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
func (o *ObjectReferenceExtended) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.Std == nil {
		o.Std = &StdObjectReference{}
	}
	if err := o.Std.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	o.Signature1 = make([]byte, 4)
	for i1 := range o.Signature1 {
		i1 := i1
		if err := w.ReadData(&o.Signature1[i1]); err != nil {
			return err
		}
	}
	_ptr_saResAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResolverAddr == nil {
			o.ResolverAddr = &DualStringArray{}
		}
		if err := o.ResolverAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_saResAddr := func(ptr interface{}) { o.ResolverAddr = *ptr.(**DualStringArray) }
	if err := w.ReadPointer(&o.ResolverAddr, _s_saResAddr, _ptr_saResAddr); err != nil {
		return err
	}
	if err := w.ReadData(&o.Elements); err != nil {
		return err
	}
	o.Signature2 = make([]byte, 4)
	for i1 := range o.Signature2 {
		i1 := i1
		if err := w.ReadData(&o.Signature2[i1]); err != nil {
			return err
		}
	}
	_ptr_ElmArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Elements > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Elements)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ElementArray", sizeInfo[0])
		}
		o.ElementArray = make([]*DataElement, sizeInfo[0])
		for i1 := range o.ElementArray {
			i1 := i1
			if o.ElementArray[i1] == nil {
				o.ElementArray[i1] = &DataElement{}
			}
			if err := o.ElementArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ElmArray := func(ptr interface{}) { o.ElementArray = *ptr.(*[]*DataElement) }
	if err := w.ReadPointer(&o.ElementArray, _s_ElmArray, _ptr_ElmArray); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// DataElement structure represents DATAELEMENT RPC structure.
//
// The DATAELEMENT structure is used to identify and marshal an envoy context as part
// of a larger OBJREF_EXTENDED structure.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dataID (16 bytes)                                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbSize                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbRounded                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Data (variable)                                                                                                               |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type DataElement struct {
	// dataID (16 bytes):  This MUST specify a context identifier for the marshaled context
	// (1). This MUST NOT be set to GUID_NULL.
	DataID *dtyp.GUID `idl:"name:dataID" json:"data_id"`
	// cbSize (4 bytes):  The unsigned number of bytes present in the Data field, excluding
	// any padding bytes at the end of the Data field that were added to round the array
	// size to an integral multiple of eight bytes. This MUST NOT be 0.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// cbRounded (4 bytes):  The unsigned size, in bytes, of the Data field. The cbRounded
	// value MUST equal the cbSize value, rounded up to a multiple of eight.
	RoundedLength uint32 `idl:"name:cbRounded" json:"rounded_length"`
	// Data (variable):  An array of cbRounded bytes that MUST contain a marshaled envoy
	// context; see section 2.2.20.
	Data []byte `idl:"name:Data;size_is:(((cbSize+7)&(^uint32(7))))" json:"data"`
}

func (o *DataElement) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.Length == 0 {
		o.Length = uint32(len(o.Data))
	}
	if o.RoundedLength == uint32(0) {
		o.RoundedLength = uint32(((o.Length + 7) & (^uint32(7))))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *DataElement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.DataID != nil {
		if err := o.DataID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.RoundedLength); err != nil {
		return err
	}
	if o.Data != nil || ((o.Length+7)&(^uint32(7))) > 0 {
		_ptr_Data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(((o.Length + 7) & (^uint32(7))))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		})
		if err := w.WritePointer(&o.Data, _ptr_Data); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DataElement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.DataID == nil {
		o.DataID = &dtyp.GUID{}
	}
	if err := o.DataID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoundedLength); err != nil {
		return err
	}
	_ptr_Data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if ((o.Length+7)&(^uint32(7))) > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(((o.Length + 7) & (^uint32(7))))
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
	})
	_s_Data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_Data, _ptr_Data); err != nil {
		return err
	}
	return nil
}

// ObjectReference structure represents OBJREF RPC structure.
//
// OBJREF is the marshaled format for a DCOM Remote Protocol object reference. There
// are four different formats for an OBJREF, which are specified by different definitions
// of the u_objref field. This section defines the initial header information. The following
// sections define substructures found in the u_objref field.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| signature                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| iid (16 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| u_objref (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReference struct {
	// signature (4 bytes): This MUST be set to the value 0x574f454d.
	Signature []byte `idl:"name:signature" json:"signature"`
	// flags (4 bytes): This MUST be set to ONE of the following values.
	//
	//	+----------------------------+-------------------------------------------+
	//	|                            |                                           |
	//	|           VALUE            |                  MEANING                  |
	//	|                            |                                           |
	//	+----------------------------+-------------------------------------------+
	//	+----------------------------+-------------------------------------------+
	//	| OBJREF_STANDARD 0x00000001 | u_objref MUST contain an OBJREF_STANDARD. |
	//	+----------------------------+-------------------------------------------+
	//	| OBJREF_HANDLER 0x00000002  | u_objref MUST contain an OBJREF_HANDLER.  |
	//	+----------------------------+-------------------------------------------+
	//	| OBJREF_CUSTOM 0x00000004   | u_objref MUST contain an OBJREF_CUSTOM.   |
	//	+----------------------------+-------------------------------------------+
	//	| OBJREF_EXTENDED 0x00000008 | u_objref MUST contain an OBJREF_EXTENDED. |
	//	+----------------------------+-------------------------------------------+
	Flags ObjectReferenceType `idl:"name:flags" json:"flags"`
	// iid (16 bytes):  The IID for which this OBJREF was marshaled; this MUST NOT be set
	// to GUID_NULL.
	IID *IID `idl:"name:iid" json:"iid"`
	// u_objref (variable):  A structure specified by the value of the preceding flags.
	ObjectReference *ObjectReference_ObjectReference `idl:"name:u_objref;switch_is:Flags" json:"object_reference"`
}

func (o *ObjectReference) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectReference) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	for i1 := range o.Signature {
		i1 := i1
		if uint64(i1) >= 4 {
			break
		}
		if err := w.WriteData(o.Signature[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Signature); uint64(i1) < 4; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint16(o.Flags)); err != nil {
		return err
	}
	if o.IID != nil {
		if err := o.IID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	_swObjectReference := uint32(o.Flags)
	if o.ObjectReference != nil {
		if err := o.ObjectReference.MarshalUnionNDR(ctx, w, _swObjectReference); err != nil {
			return err
		}
	} else {
		if err := (&ObjectReference_ObjectReference{}).MarshalUnionNDR(ctx, w, _swObjectReference); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReference) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	o.Signature = make([]byte, 4)
	for i1 := range o.Signature {
		i1 := i1
		if err := w.ReadData(&o.Signature[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadEnum((*uint16)(&o.Flags)); err != nil {
		return err
	}
	if o.IID == nil {
		o.IID = &IID{}
	}
	if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ObjectReference == nil {
		o.ObjectReference = &ObjectReference_ObjectReference{}
	}
	_swObjectReference := uint32(o.Flags)
	if err := o.ObjectReference.UnmarshalUnionNDR(ctx, w, _swObjectReference); err != nil {
		return err
	}
	return nil
}

// ObjectReference_ObjectReference structure represents OBJREF union anonymous member.
//
// OBJREF is the marshaled format for a DCOM Remote Protocol object reference. There
// are four different formats for an OBJREF, which are specified by different definitions
// of the u_objref field. This section defines the initial header information. The following
// sections define substructures found in the u_objref field.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| signature                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| iid (16 bytes)                                                                                                                |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| u_objref (variable)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ObjectReference_ObjectReference struct {
	// Types that are assignable to Value
	//
	// *ObjectReference_Standard
	// *ObjectReference_Handler
	// *ObjectReference_Custom
	// *ObjectReference_Extended
	Value is_ObjectReference_ObjectReference `json:"value"`
}

func (o *ObjectReference_ObjectReference) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ObjectReference_Standard:
		if value != nil {
			return value.Standard
		}
	case *ObjectReference_Handler:
		if value != nil {
			return value.Handler
		}
	case *ObjectReference_Custom:
		if value != nil {
			return value.Custom
		}
	case *ObjectReference_Extended:
		if value != nil {
			return value.Extended
		}
	}
	return nil
}

type is_ObjectReference_ObjectReference interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ObjectReference_ObjectReference()
}

func (o *ObjectReference_ObjectReference) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ObjectReference_Standard:
		return uint32(1)
	case *ObjectReference_Handler:
		return uint32(2)
	case *ObjectReference_Custom:
		return uint32(4)
	case *ObjectReference_Extended:
		return uint32(8)
	}
	return uint32(0)
}

func (o *ObjectReference_ObjectReference) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
	case uint32(1):
		_o, _ := o.Value.(*ObjectReference_Standard)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectReference_Standard{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*ObjectReference_Handler)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectReference_Handler{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*ObjectReference_Custom)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectReference_Custom{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*ObjectReference_Extended)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectReference_Extended{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ObjectReference_ObjectReference) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
	case uint32(1):
		o.Value = &ObjectReference_Standard{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &ObjectReference_Handler{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &ObjectReference_Custom{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &ObjectReference_Extended{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ObjectReference_Standard structure represents ObjectReference_ObjectReference RPC union arm.
//
// It has following labels: 1
type ObjectReference_Standard struct {
	Standard *ObjectReferenceStandard `idl:"name:standard" json:"standard"`
}

func (*ObjectReference_Standard) is_ObjectReference_ObjectReference() {}

func (o *ObjectReference_Standard) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Standard != nil {
		if err := o.Standard.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectReferenceStandard{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReference_Standard) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Standard == nil {
		o.Standard = &ObjectReferenceStandard{}
	}
	if err := o.Standard.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectReference_Handler structure represents ObjectReference_ObjectReference RPC union arm.
//
// It has following labels: 2
type ObjectReference_Handler struct {
	Handler *ObjectReferenceHandler `idl:"name:handler" json:"handler"`
}

func (*ObjectReference_Handler) is_ObjectReference_ObjectReference() {}

func (o *ObjectReference_Handler) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Handler != nil {
		if err := o.Handler.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectReferenceHandler{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReference_Handler) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Handler == nil {
		o.Handler = &ObjectReferenceHandler{}
	}
	if err := o.Handler.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectReference_Custom structure represents ObjectReference_ObjectReference RPC union arm.
//
// It has following labels: 4
type ObjectReference_Custom struct {
	Custom *ObjectReferenceCustom `idl:"name:custom" json:"custom"`
}

func (*ObjectReference_Custom) is_ObjectReference_ObjectReference() {}

func (o *ObjectReference_Custom) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Custom != nil {
		if err := o.Custom.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectReferenceCustom{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReference_Custom) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Custom == nil {
		o.Custom = &ObjectReferenceCustom{}
	}
	if err := o.Custom.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ObjectReference_Extended structure represents ObjectReference_ObjectReference RPC union arm.
//
// It has following labels: 8
type ObjectReference_Extended struct {
	Extended *ObjectReferenceExtended `idl:"name:extended" json:"extended"`
}

func (*ObjectReference_Extended) is_ObjectReference_ObjectReference() {}

func (o *ObjectReference_Extended) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Extended != nil {
		if err := o.Extended.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectReferenceExtended{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectReference_Extended) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Extended == nil {
		o.Extended = &ObjectReferenceExtended{}
	}
	if err := o.Extended.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ContextMarshalFlag type represents CTXMSHLFLAGS RPC enumeration.
type ContextMarshalFlag uint16

var (
	ContextMarshalFlagByValue ContextMarshalFlag = 2
)

func (o ContextMarshalFlag) String() string {
	switch o {
	case ContextMarshalFlagByValue:
		return "ContextMarshalFlagByValue"
	}
	return "Invalid"
}

// Context structure represents Context RPC structure.
//
// This is the marshaled representation of a context (1). It contains an array of marshaled
// context properties, each of which is represented by a PROPMARSHALHEADER.
//
// There are three types of context (1) structures:
//
// * A client context ( ba4c4d80-ef81-49b4-848f-9714d72b5c01#gt_519bff3c-1c9f-4d5a-aa88-a3c820a4ff3a
// ). This type of context (1) is sent in an activation request in the *pIFDClientCtx*
// field of the ActivationContextInfoData ( 5892b550-cd9e-4277-9644-4886d3b6d754 ) structure
// (section 2.2.22.2.5). The context (1) structure MUST be marshaled into an OBJREF_CUSTOM
// ( 88fdb261-7c9a-41ea-bfac-31583bdc65e6 ) structure (section 2.2.18.6) whose *clsid*
// field is set to CLSID_ContextMarshaler (section 1.9 ( c25391af-f59e-40da-885e-cc84076673e4
// ) ). The *iid* field of the OBJREF ( fe6c5e46-adf8-4e34-a8de-3f756c875f31 ) portion
// of the OBJREF_CUSTOM structure MUST be set to IID_IContext (section 1.9). An implementation
// MAY use this value as the IID ( ba4c4d80-ef81-49b4-848f-9714d72b5c01#gt_76ad3105-3f05-479d-a40c-c9c8fa2ebd83
// ) of an interface with the local IDL attribute (section 2.2.27 ( 100ac689-eb4e-4b63-8354-6e064a485d51
// ) ). <12> ( acc42954-4073-4f05-b850-efd562022077#Appendix_A_12 )
//
// * A prototype context ( ba4c4d80-ef81-49b4-848f-9714d72b5c01#gt_3b138442-c9c6-4f0f-bb59-66c783ed4d65
// ). An application or a higher-layer protocol can instruct a DCOM client to send this
// type of context (1) in an activation request. The prototype context is sent in the
// *pIFDPrototypeCtx* field of the ActivationContextInfoData structure (section 2.2.22.2.5).
// The context (1) structure MUST be marshaled into an OBJREF_CUSTOM structure (section
// 2.2.18.6) whose *clsid* field is set to CLSID_ContextMarshaler (section 1.9). The
// *iid* field of the OBJREF portion of the OBJREF_CUSTOM structure MUST be set to IID_IContext
// (see section 1.9). An implementation MAY use this value as the IID of an interface
// with the local IDL attribute (section 2.2.27). <13> ( acc42954-4073-4f05-b850-efd562022077#Appendix_A_13
// )
//
// * An envoy context ( ba4c4d80-ef81-49b4-848f-9714d72b5c01#gt_a61f46c9-edea-481b-b494-517f3b05e88a
// ). An application or a higher-layer protocol can instruct a DCOM server to send this
// type of context (1) when marshaling an object. The envoy context is sent in the *Data*
// field of the *ElmArray* field (section 2.2.18.8 ( 33ae0690-a3f7-4591-8443-b9df4567984d
// ) ) of an OBJREF_EXTENDED ( 06220154-0372-495a-bfc6-a96b804884c6 ) structure (section
// 2.2.18.7).
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| MajorVersion                                                  | MinVersion                                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ContextId (16 bytes)                                                                                                          |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Reserved                                                                                                                      |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| dwNumExtents                                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cbExtents                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| MshlFlags                                                                                                                     |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Count                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Frozen                                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| PropMarshalHeader (variable)                                                                                                  |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type Context struct {
	// MajorVersion (2 bytes):  The major version of this context marshaled format. This
	// MUST be set to 0x0001.
	MajorVersion uint16 `idl:"name:MajorVersion" json:"major_version"`
	// MinVersion (2 bytes): The minor version of this context (1) marshaled format. This
	// MUST be set to 0x0001.
	MinVersion uint16 `idl:"name:MinVersion" json:"min_version"`
	// ContextId (16 bytes):  A GUID identifying the marshaled context (1).
	ContextID *dtyp.GUID `idl:"name:ContextId" json:"context_id"`
	// Flags (4 bytes): This MUST be set to the following value.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| CTXMSHLFLAGS_BYVAL 0x00000002 | The context is marshaled by value (this is the only representation that is valid |
	//	|                               | on the wire).                                                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// Reserved (4 bytes):  This MUST be set to 0x00000000 and MUST be ignored on receipt.
	_ uint32 `idl:"name:Reserved"`
	// dwNumExtents (4 bytes):  This MUST be set to 0x00000000.
	ExtensionsCount uint32 `idl:"name:dwNumExtents" json:"extensions_count"`
	// cbExtents (4 bytes):  This MUST be set to 0x00000000.
	ExtentsLength uint32 `idl:"name:cbExtents" json:"extents_length"`
	// MshlFlags (4 bytes): This MUST contain an implementation-specific value that MUST
	// be ignored on receipt.<14>
	MarshalFlags uint32 `idl:"name:MshlFlags" json:"marshal_flags"`
	// Count (4 bytes):  The unsigned number of elements in the PropMarshalHeader array.
	Count uint32 `idl:"name:Count" json:"count"`
	// Frozen (4 bytes):  A Boolean that MUST be set to TRUE (0x00000001) and that MUST
	// be ignored on receipt.
	Frozen bool `idl:"name:Frozen" json:"frozen"`
	// PropMarshalHeader (variable):  This MUST be an array of PROPMARSHALHEADER entries.
	ContextProperties []*ContextProperty `idl:"name:PropMarshalHeader;size_is:(Count)" json:"context_properties"`
}

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ContextProperties != nil && o.Count == 0 {
		o.Count = uint32(len(o.ContextProperties))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinVersion); err != nil {
		return err
	}
	if o.ContextID != nil {
		if err := o.ContextID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtensionsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtentsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.MarshalFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if !o.Frozen {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.ContextProperties != nil || o.Count > 0 {
		_ptr_PropMarshalHeader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ContextProperties {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ContextProperties[i1] != nil {
					if err := o.ContextProperties[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ContextProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ContextProperties); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ContextProperty{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ContextProperties, _ptr_PropMarshalHeader); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinVersion); err != nil {
		return err
	}
	if o.ContextID == nil {
		o.ContextID = &dtyp.GUID{}
	}
	if err := o.ContextID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint32
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtensionsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtentsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.MarshalFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	var _bFrozen int32
	if err := w.ReadData(&_bFrozen); err != nil {
		return err
	}
	o.Frozen = _bFrozen != 0
	_ptr_PropMarshalHeader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ContextProperties", sizeInfo[0])
		}
		o.ContextProperties = make([]*ContextProperty, sizeInfo[0])
		for i1 := range o.ContextProperties {
			i1 := i1
			if o.ContextProperties[i1] == nil {
				o.ContextProperties[i1] = &ContextProperty{}
			}
			if err := o.ContextProperties[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_PropMarshalHeader := func(ptr interface{}) { o.ContextProperties = *ptr.(*[]*ContextProperty) }
	if err := w.ReadPointer(&o.ContextProperties, _s_PropMarshalHeader, _ptr_PropMarshalHeader); err != nil {
		return err
	}
	return nil
}

// ContextPropertyFlag type represents CPFLAG RPC enumeration.
type ContextPropertyFlag uint16

var (
	ContextPropertyFlagPropagate ContextPropertyFlag = 1
	ContextPropertyFlagExpose    ContextPropertyFlag = 2
	ContextPropertyFlagEnvoy     ContextPropertyFlag = 4
)

func (o ContextPropertyFlag) String() string {
	switch o {
	case ContextPropertyFlagPropagate:
		return "ContextPropertyFlagPropagate"
	case ContextPropertyFlagExpose:
		return "ContextPropertyFlagExpose"
	case ContextPropertyFlagEnvoy:
		return "ContextPropertyFlagEnvoy"
	}
	return "Invalid"
}

// ContextProperty structure represents PROPMARSHALHEADER RPC structure.
//
// PROPMARSHALHEADER is the marshaled representation of a context property. It contains
// a context property identifier and a context property data buffer supplied by an application
// or higher-layer protocol. The data buffer contains either an OBJREF that can be of
// any valid OBJREF format representing a client or prototype context property, or it
// contains opaque data representing an envoy context property.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| clsid (16 bytes)                                                                                                              |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| policyId (16 bytes)                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| flags                                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| cb                                                                                                                            |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ctxProperty (variable)                                                                                                        |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type ContextProperty struct {
	// clsid (16 bytes):  This field MUST be either GUID_NULL or a CLSID supplied by the
	// application or higher-layer protocol identifying an object class associated with
	// the data in the ctxProperty field. If this field is GUID_NULL, ctxProperty MUST contain
	// an OBJREF representing the marshaled client or prototype context property; otherwise,
	// ctxProperty MUST contain opaque data representing the marshaled envoy context property.
	ClassID *ClassID `idl:"name:clsid" json:"class_id"`
	// policyId (16 bytes):  A GUID supplied by the application or higher-layer protocol
	// containing a context property identifier for the marshaled context property in ctxProperty.
	PolicyID *dtyp.GUID `idl:"name:policyId" json:"policy_id"`
	// flags (4 bytes): This MUST be set to one of the following values.
	//
	//	+-----------------------------+-------------------------------------------------------+
	//	|                             |                                                       |
	//	|            VALUE            |                        MEANING                        |
	//	|                             |                                                       |
	//	+-----------------------------+-------------------------------------------------------+
	//	+-----------------------------+-------------------------------------------------------+
	//	| CPFLAG_PROPAGATE 0x00000001 | This context property is part of a prototype context. |
	//	+-----------------------------+-------------------------------------------------------+
	//	| CPFLAG_EXPOSE 0x00000002    | This context property is part of a client context.    |
	//	+-----------------------------+-------------------------------------------------------+
	//	| CPFLAG_ENVOY 0x00000004     | This context property is part of an envoy context.    |
	//	+-----------------------------+-------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// cb (4 bytes): The unsigned size, in bytes, of the ctxProperty field.
	Length uint32 `idl:"name:cb" json:"length"`
	// ctxProperty (variable):  A buffer of cb bytes supplied by the application or higher-layer
	// protocol. This buffer MUST contain an OBJREF representing the marshaled client or
	// prototype context property if clsid is GUID_NULL; otherwise, it MUST contain opaque
	// data representing the marshaled envoy context property.
	ContextProperty []byte `idl:"name:ctxProperty;size_is:(cb)" json:"context_property"`
}

func (o *ContextProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ContextProperty != nil && o.Length == 0 {
		o.Length = uint32(len(o.ContextProperty))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ContextProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PolicyID != nil {
		if err := o.PolicyID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.ContextProperty != nil || o.Length > 0 {
		_ptr_ctxProperty := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ContextProperty {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.ContextProperty[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.ContextProperty); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ContextProperty, _ptr_ctxProperty); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &ClassID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PolicyID == nil {
		o.PolicyID = &dtyp.GUID{}
	}
	if err := o.PolicyID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_ctxProperty := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ContextProperty", sizeInfo[0])
		}
		o.ContextProperty = make([]byte, sizeInfo[0])
		for i1 := range o.ContextProperty {
			i1 := i1
			if err := w.ReadData(&o.ContextProperty[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ctxProperty := func(ptr interface{}) { o.ContextProperty = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.ContextProperty, _s_ctxProperty, _ptr_ctxProperty); err != nil {
		return err
	}
	return nil
}

// ActivationPropertiesBlob structure represents ActivationPropertiesBLOB RPC structure.
type ActivationPropertiesBlob struct {
	Size uint32 `idl:"name:dwSize" json:"size"`
	_    uint32 `idl:"name:dwReserved"`
	Data []byte `idl:"name:data;size_is:(dwSize)" json:"data"`
}

func (o *ActivationPropertiesBlob) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ActivationPropertiesBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.Data != nil || o.Size > 0 {
		_ptr_data := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		})
		if err := w.WritePointer(&o.Data, _ptr_data); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ActivationPropertiesBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	_ptr_data := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
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
	})
	_s_data := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_data, _ptr_data); err != nil {
		return err
	}
	return nil
}

// RemoteQueryInterfaceResult structure represents REMQIRESULT RPC structure.
//
// The REMQIRESULT structure is passed as an output parameter from IRemUnknown::RemQueryInterface.
// It contains the result of the RemQueryInterface and the STDOBJREF containing the
// object reference for the queried interface.
type RemoteQueryInterfaceResult struct {
	// hResult:   This MUST contain zero if the QueryInterface operation was successful.
	// Otherwise, this MUST contain a negative value to indicate failure; see section 3.1.1.5.6.1.1.
	HResult int32 `idl:"name:hResult" json:"hresult"`
	// std:   If hResult is zero, this MUST contain a STDOBJREF instance that the client
	// can unmarshal and use to make calls on the interface. If hResult contains an error
	// value, this field MUST be ignored.
	Std *StdObjectReference `idl:"name:std" json:"std"`
}

func (o *RemoteQueryInterfaceResult) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RemoteQueryInterfaceResult) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.HResult); err != nil {
		return err
	}
	if o.Std != nil {
		if err := o.Std.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StdObjectReference{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteQueryInterfaceResult) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.HResult); err != nil {
		return err
	}
	if o.Std == nil {
		o.Std = &StdObjectReference{}
	}
	if err := o.Std.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RemoteInterfaceReference structure represents REMINTERFACEREF RPC structure.
//
// The REMINTERFACEREF structure is passed as a parameter to either IRemUnknown::RemAddRef
// (Opnum 4) or IRemUnknown::RemRelease (Opnum 5). It specifies the number and type
// of references that the client requests to be added to (or subtracted from) an interface
// reference count.
type RemoteInterfaceReference struct {
	// ipid:  This MUST be the IPID of the interface reference count to be modified.
	IPID *IPID `idl:"name:ipid" json:"ipid"`
	// cPublicRefs:   This MUST be the number of public references (see section 1.3.6) on
	// the interface identified by IPID being requested by the client.
	PublicReferencesCount uint32 `idl:"name:cPublicRefs" json:"public_references_count"`
	// cPrivateRefs:  This MUST be the number of private references (see section 1.3.6)
	// on the interface identified by IPID being requested by the client.
	PrivateReferencesCount uint32 `idl:"name:cPrivateRefs" json:"private_references_count"`
}

func (o *RemoteInterfaceReference) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *RemoteInterfaceReference) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.IPID != nil {
		if err := o.IPID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PublicReferencesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.PrivateReferencesCount); err != nil {
		return err
	}
	return nil
}
func (o *RemoteInterfaceReference) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.IPID == nil {
		o.IPID = &IPID{}
	}
	if err := o.IPID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PublicReferencesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrivateReferencesCount); err != nil {
		return err
	}
	return nil
}

// COMServerInfo structure represents COSERVERINFO RPC structure.
//
// The COSERVERINFO structure SHOULD NOT be sent and MUST be ignored on receipt.
type COMServerInfo struct {
	// dwReserved1:  This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved1"`
	// pwszName:  This SHOULD be set to NULL and MUST be ignored by servers.<36>
	Name string `idl:"name:pwszName;string" json:"name"`
	// pdwReserved:   This MUST be set to NULL and MUST be ignored on receipt.
	_ uint32 `idl:"name:pdwReserved"`
	// dwReserved2:  This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved2"`
}

func (o *COMServerInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *COMServerInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_pwszName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_pwszName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved pdwReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	// reserved dwReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *COMServerInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	_ptr_pwszName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_pwszName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_pwszName, _ptr_pwszName); err != nil {
		return err
	}
	// reserved pdwReserved
	var _pdwReserved uint32
	_ptr_pdwReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_pdwReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pdwReserved := func(ptr interface{}) { _pdwReserved = *ptr.(*uint32) }
	if err := w.ReadPointer(&_pdwReserved, _s_pdwReserved, _ptr_pdwReserved); err != nil {
		return err
	}
	// reserved dwReserved2
	var _dwReserved2 uint32
	if err := w.ReadData(&_dwReserved2); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// CustomRemoteRequestSCMInfo structure represents customREMOTE_REQUEST_SCM_INFO RPC structure.
//
// The customREMOTE_REQUEST_SCM_INFO structure specifies the protocol sequence identifiers
// supported by the client.
type CustomRemoteRequestSCMInfo struct {
	// ClientImpLevel:   This MUST contain an implementation-specific value that MUST be
	// ignored on receipt.<33>
	ClientImpLevel uint32 `idl:"name:ClientImpLevel" json:"client_imp_level"`
	// cRequestedProtseqs:   This MUST contain the number of elements in the pRequestedProtseqs
	// array and SHOULD be at least 1.
	RequestedProtocolSequencesCount uint16 `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	// pRequestedProtseqs:  This MUST contain an array of RPC protocol sequence identifiers
	// supported by the client.
	RequestedProtocolSequences []uint16 `idl:"name:pRequestedProtseqs;size_is:(cRequestedProtseqs)" json:"requested_protocol_sequences"`
}

func (o *CustomRemoteRequestSCMInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.RequestedProtocolSequences != nil && o.RequestedProtocolSequencesCount == 0 {
		o.RequestedProtocolSequencesCount = uint16(len(o.RequestedProtocolSequences))
	}
	if o.RequestedProtocolSequencesCount > uint16(32768) {
		return fmt.Errorf("RequestedProtocolSequencesCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CustomRemoteRequestSCMInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientImpLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestedProtocolSequencesCount); err != nil {
		return err
	}
	if o.RequestedProtocolSequences != nil || o.RequestedProtocolSequencesCount > 0 {
		_ptr_pRequestedProtseqs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.RequestedProtocolSequencesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RequestedProtocolSequences {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.RequestedProtocolSequences[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.RequestedProtocolSequences); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RequestedProtocolSequences, _ptr_pRequestedProtseqs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CustomRemoteRequestSCMInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientImpLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestedProtocolSequencesCount); err != nil {
		return err
	}
	_ptr_pRequestedProtseqs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.RequestedProtocolSequencesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.RequestedProtocolSequencesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RequestedProtocolSequences", sizeInfo[0])
		}
		o.RequestedProtocolSequences = make([]uint16, sizeInfo[0])
		for i1 := range o.RequestedProtocolSequences {
			i1 := i1
			if err := w.ReadData(&o.RequestedProtocolSequences[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRequestedProtseqs := func(ptr interface{}) { o.RequestedProtocolSequences = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.RequestedProtocolSequences, _s_pRequestedProtseqs, _ptr_pRequestedProtseqs); err != nil {
		return err
	}
	return nil
}

// CustomRemoteReplySCMInfo structure represents customREMOTE_REPLY_SCM_INFO RPC structure.
//
// The customREMOTE_REPLY_SCM_INFO structure is used to return information about the
// object exporter, specifically the OXID, RPC bindings, COMVERSION, and IPID of the
// IRemUnknown interface and the authentication hint of the object exporter.
type CustomRemoteReplySCMInfo struct {
	// Oxid:   This MUST contain the OXID identifier for the object exporter.
	OXID uint64 `idl:"name:Oxid" json:"oxid"`
	// pdsaOxidBindings:   This MUST specify the string and security bindings supported
	// by the object exporter and MUST NOT be NULL. The returned string bindings SHOULD
	// contain endpoints.
	OXIDBindings *DualStringArray `idl:"name:pdsaOxidBindings" json:"oxid_bindings"`
	// ipidRemUnknown:   This MUST specify the IPID of the object exporter's Remote Unknown
	// object.
	IPIDRemoteUnknown *IPID `idl:"name:ipidRemUnknown" json:"ipid_remote_unknown"`
	// authnHint:  This SHOULD contain an RPC authentication level (see [MS-RPCE] section
	// 2.2.1.1.8) that denotes the minimum authentication level supported by the object
	// exporter. This field MAY be ignored by the client.<37>
	AuthnHint uint32 `idl:"name:authnHint" json:"authn_hint"`
	// serverVersion:  This MUST contain the COMVERSION of the server.
	ServerVersion *COMVersion `idl:"name:serverVersion" json:"server_version"`
}

func (o *CustomRemoteReplySCMInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CustomRemoteReplySCMInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.OXID); err != nil {
		return err
	}
	if o.OXIDBindings != nil {
		_ptr_pdsaOxidBindings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.OXIDBindings != nil {
				if err := o.OXIDBindings.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&DualStringArray{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.OXIDBindings, _ptr_pdsaOxidBindings); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IPIDRemoteUnknown != nil {
		if err := o.IPIDRemoteUnknown.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&IPID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AuthnHint); err != nil {
		return err
	}
	if o.ServerVersion != nil {
		if err := o.ServerVersion.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&COMVersion{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *CustomRemoteReplySCMInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.OXID); err != nil {
		return err
	}
	_ptr_pdsaOxidBindings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.OXIDBindings == nil {
			o.OXIDBindings = &DualStringArray{}
		}
		if err := o.OXIDBindings.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pdsaOxidBindings := func(ptr interface{}) { o.OXIDBindings = *ptr.(**DualStringArray) }
	if err := w.ReadPointer(&o.OXIDBindings, _s_pdsaOxidBindings, _ptr_pdsaOxidBindings); err != nil {
		return err
	}
	if o.IPIDRemoteUnknown == nil {
		o.IPIDRemoteUnknown = &IPID{}
	}
	if err := o.IPIDRemoteUnknown.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthnHint); err != nil {
		return err
	}
	if o.ServerVersion == nil {
		o.ServerVersion = &COMVersion{}
	}
	if err := o.ServerVersion.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// ActivateFlags type represents ACTVFLAGS RPC enumeration.
type ActivateFlags uint16

var (
	ActivateFlagsDisableAAA          ActivateFlags = 2
	ActivateFlagsActivate32BitServer ActivateFlags = 4
	ActivateFlagsActivate64BitServer ActivateFlags = 8
	ActivateFlagsNoFailureLog        ActivateFlags = 32
)

func (o ActivateFlags) String() string {
	switch o {
	case ActivateFlagsDisableAAA:
		return "ActivateFlagsDisableAAA"
	case ActivateFlagsActivate32BitServer:
		return "ActivateFlagsActivate32BitServer"
	case ActivateFlagsActivate64BitServer:
		return "ActivateFlagsActivate64BitServer"
	case ActivateFlagsNoFailureLog:
		return "ActivateFlagsNoFailureLog"
	}
	return "Invalid"
}

// InstantiationInfoData structure represents InstantiationInfoData RPC structure.
//
// The client uses this structure to specify basic details of the object to be activated,
// including the identifying object CLSID and one or more requested object interfaces.
//
// CLSID_InstantiationInfo (section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type InstantiationInfoData struct {
	// classId:  The CLSID of the COM object class that the client activates.
	ClassID *ClassID `idl:"name:classId" json:"class_id"`
	// classCtx:  An implementation-specific value that SHOULD be ignored on receipt.<23>
	ClassContext uint32 `idl:"name:classCtx" json:"class_context"`
	// actvflags:  0x00000000 or any combination of the following bit values.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                             |                                                                                  |
	//	|                    VALUE                    |                                     MEANING                                      |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACTVFLAGS_DISABLE_AAA 0x00000002            | The object resolver is requested to not execute the object exporter under the    |
	//	|                                             | client's identity.                                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACTVFLAGS_ACTIVATE_32_BIT_SERVER 0x00000004 | The object resolver is requested to execute the object exporter in the 32-bit    |
	//	|                                             | address space.                                                                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACTVFLAGS_ACTIVATE_64_BIT_SERVER 0x00000008 | The object resolver is requested to execute the object exporter in the 64-bit    |
	//	|                                             | address space.                                                                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| ACTVFLAGS_NO_FAILURE_LOG 0x00000020         | The object resolver is requested to not log an error if a failure occurs during  |
	//	|                                             | the activation request.                                                          |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	ActivateFlags uint32 `idl:"name:actvflags" json:"activate_flags"`
	// fIsSurrogate:   This MUST be set to FALSE (0x00000000) and MUST be ignored on receipt.
	IsSurrogate int32 `idl:"name:fIsSurrogate" json:"is_surrogate"`
	// cIID:  The number of interfaces in the pIID array. This value MUST be between 1 and
	// MAX_REQUESTED_INTERFACES (see section 2.2.28.1).
	IIDCount uint32 `idl:"name:cIID" json:"iid_count"`
	// instFlag:   This MUST be set to zero and MUST be ignored on receipt.
	InstanceFlag uint32 `idl:"name:instFlag" json:"instance_flag"`
	// pIID:  An array of IIDs identifying the interfaces that the client requests from
	// the server.
	IID []*IID `idl:"name:pIID;size_is:(cIID)" json:"iid"`
	// thisSize:  The size (in bytes) of this structure, as marshaled by the NDR Type Serialization
	// 1 engine (as specified in [MS-RPCE] section 2.2.6). It SHOULD be ignored on receipt.
	ThisSize uint32 `idl:"name:thisSize" json:"this_size"`
	// clientCOMVersion:   The COMVERSION of the client. This MUST be ignored on receipt.
	ClientCOMVersion *COMVersion `idl:"name:clientCOMVersion" json:"client_com_version"`
}

func (o *InstantiationInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.IID != nil && o.IIDCount == 0 {
		o.IIDCount = uint32(len(o.IID))
	}
	if o.IIDCount < uint32(1) || o.IIDCount > uint32(32768) {
		return fmt.Errorf("IIDCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InstantiationInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ClassID != nil {
		if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ClassContext); err != nil {
		return err
	}
	if err := w.WriteData(o.ActivateFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.IsSurrogate); err != nil {
		return err
	}
	if err := w.WriteData(o.IIDCount); err != nil {
		return err
	}
	if err := w.WriteData(o.InstanceFlag); err != nil {
		return err
	}
	if o.IID != nil || o.IIDCount > 0 {
		_ptr_pIID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.IIDCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.IID {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.IID[i1] != nil {
					if err := o.IID[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.IID); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IID, _ptr_pIID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ThisSize); err != nil {
		return err
	}
	if o.ClientCOMVersion != nil {
		if err := o.ClientCOMVersion.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&COMVersion{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *InstantiationInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.ClassID == nil {
		o.ClassID = &ClassID{}
	}
	if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClassContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.ActivateFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.IsSurrogate); err != nil {
		return err
	}
	if err := w.ReadData(&o.IIDCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.InstanceFlag); err != nil {
		return err
	}
	_ptr_pIID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.IIDCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.IIDCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.IID", sizeInfo[0])
		}
		o.IID = make([]*IID, sizeInfo[0])
		for i1 := range o.IID {
			i1 := i1
			if o.IID[i1] == nil {
				o.IID[i1] = &IID{}
			}
			if err := o.IID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pIID := func(ptr interface{}) { o.IID = *ptr.(*[]*IID) }
	if err := w.ReadPointer(&o.IID, _s_pIID, _ptr_pIID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ThisSize); err != nil {
		return err
	}
	if o.ClientCOMVersion == nil {
		o.ClientCOMVersion = &COMVersion{}
	}
	if err := o.ClientCOMVersion.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// LocationInfoData structure represents LocationInfoData RPC structure.
//
// The LocationInfoData structure MUST be present in the Activation Properties BLOB
// structure. The server MUST ignore this structure.
//
// CLSID_ServerLocationInfo (see section 1.9) is used to identify this property in the
// CustomHeader.pclsid array.
type LocationInfoData struct {
	// machineName:  This MUST be set to NULL and MUST be ignored on receipt.
	MachineName string `idl:"name:machineName;string" json:"machine_name"`
	// processId:   This MUST be set to 0 and MUST be ignored on receipt.
	ProcessID uint32 `idl:"name:processId" json:"process_id"`
	// apartmentId:   This MUST be set to 0 and MUST be ignored on receipt.
	ApartmentID uint32 `idl:"name:apartmentId" json:"apartment_id"`
	// contextId:   This MUST be set to 0 and MUST be ignored on receipt.
	ContextID uint32 `idl:"name:contextId" json:"context_id"`
}

func (o *LocationInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *LocationInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.MachineName != "" {
		_ptr_machineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.MachineName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MachineName, _ptr_machineName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if err := w.WriteData(o.ApartmentID); err != nil {
		return err
	}
	if err := w.WriteData(o.ContextID); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *LocationInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_machineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.MachineName); err != nil {
			return err
		}
		return nil
	})
	_s_machineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
	if err := w.ReadPointer(&o.MachineName, _s_machineName, _ptr_machineName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ApartmentID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ContextID); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ActivationContextInfoData structure represents ActivationContextInfoData RPC structure.
//
// The ActivationContextInfoData structure passes a client context and optionally a
// prototype context to the server as part of an activation request.
//
// CLSID_ActivationContextInfo (see section 1.9) is used to identify this property in
// the CustomHeader.pclsid array.
type ActivationContextInfoData struct {
	// clientOK:  This MUST be set to FALSE (0x00000000) and MUST be ignored on receipt.
	ClientOK int32 `idl:"name:clientOK" json:"client_ok"`
	// bReserved1:   This MUST be set to FALSE (0x00000000) and MUST be ignored on receipt.
	_ int32 `idl:"name:bReserved1"`
	// dwReserved1:  This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved1"`
	// dwReserved2:   This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved2"`
	// pIFDClientCtx:  This MUST contain an OBJREF specifying a marshaled client context
	// as specified in section 2.2.20. The server MUST return RPC_E_INVALID_OBJREF (as specified
	// in [MS-ERREF] section 2.1) if the OBJREF is NULL or invalid.
	IfdClientContext *InterfacePointer `idl:"name:pIFDClientCtx" json:"ifd_client_context"`
	// pIFDPrototypeCtx:  If an application or a higher-layer protocol instructs the client
	// to send prototype context properties, this MUST contain an OBJREF specifying a marshaled
	// prototype context as specified in section 2.2.20. Otherwise, this MUST be set to
	// NULL.
	IfdPrototypeContext *InterfacePointer `idl:"name:pIFDPrototypeCtx" json:"ifd_prototype_context"`
}

func (o *ActivationContextInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ActivationContextInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientOK); err != nil {
		return err
	}
	// reserved bReserved1
	if err := w.WriteData(int32(0)); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if o.IfdClientContext != nil {
		_ptr_pIFDClientCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IfdClientContext != nil {
				if err := o.IfdClientContext.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IfdClientContext, _ptr_pIFDClientCtx); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IfdPrototypeContext != nil {
		_ptr_pIFDPrototypeCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IfdPrototypeContext != nil {
				if err := o.IfdPrototypeContext.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IfdPrototypeContext, _ptr_pIFDPrototypeCtx); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ActivationContextInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientOK); err != nil {
		return err
	}
	// reserved bReserved1
	var _bReserved1 int32
	if err := w.ReadData(&_bReserved1); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	// reserved dwReserved2
	var _dwReserved2 uint32
	if err := w.ReadData(&_dwReserved2); err != nil {
		return err
	}
	_ptr_pIFDClientCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IfdClientContext == nil {
			o.IfdClientContext = &InterfacePointer{}
		}
		if err := o.IfdClientContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pIFDClientCtx := func(ptr interface{}) { o.IfdClientContext = *ptr.(**InterfacePointer) }
	if err := w.ReadPointer(&o.IfdClientContext, _s_pIFDClientCtx, _ptr_pIFDClientCtx); err != nil {
		return err
	}
	_ptr_pIFDPrototypeCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IfdPrototypeContext == nil {
			o.IfdPrototypeContext = &InterfacePointer{}
		}
		if err := o.IfdPrototypeContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pIFDPrototypeCtx := func(ptr interface{}) { o.IfdPrototypeContext = *ptr.(**InterfacePointer) }
	if err := w.ReadPointer(&o.IfdPrototypeContext, _s_pIFDPrototypeCtx, _ptr_pIFDPrototypeCtx); err != nil {
		return err
	}
	return nil
}

// CustomHeader structure represents CustomHeader RPC structure.
//
// The CustomHeader structure is used to identify the format and ordering of the properties
// in the activation properties BLOB.
type CustomHeader struct {
	// totalSize:   This MUST be the total size (in bytes) from the beginning of the CustomHeader
	// to the end of the last entry in the subsequent Property array of the activation properties
	// BLOB.
	TotalSize uint32 `idl:"name:totalSize" json:"total_size"`
	// headerSize:  This MUST be the total size (in bytes) of the CustomHeader as marshaled
	// by the NDR Type Serialization 1 engine (as specified in [MS-RPCE] section 2.2.6).
	HeaderSize uint32 `idl:"name:headerSize" json:"header_size"`
	// dwReserved:  This MUST be set to zero and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved"`
	// destCtx:  This MUST contain an implementation-specific value that SHOULD be ignored
	// on receipt.<22>
	DestinationContext uint32 `idl:"name:destCtx" json:"destination_context"`
	// cIfs:  This MUST be the total number of entries in the subsequent Property array
	// of the activation properties BLOB. The value MUST be between MIN_ACTPROP_LIMIT and
	// MAX_ACTPROP_LIMIT (see section 2.2.28.1).
	InterfacesCount uint32 `idl:"name:cIfs" json:"interfaces_count"`
	// classInfoClsid:  This MUST be set to GUID_NULL.
	ClassInfoClassID *ClassID `idl:"name:classInfoClsid" json:"class_info_class_id"`
	// pclsid:  This MUST specify an array of cIfs CLSIDs; the Nth entry identifies the
	// Nth entry in the Property array of the activation properties BLOB. Each CLSID is
	// used to uniquely identify an activation property. The valid CLSID values are defined
	// in section 1.9.
	ClassIDs []*ClassID `idl:"name:pclsid;size_is:(cIfs)" json:"class_ids"`
	// pSizes:  This MUST specify an array of cIfs DWORDs, each containing the size (in
	// bytes) of the corresponding property following the CustomHeader in the buffer.
	Sizes []uint32 `idl:"name:pSizes;size_is:(cIfs)" json:"sizes"`
	// pdwReserved:  This MUST be set to NULL and MUST be ignored on receipt.
	_ uint32 `idl:"name:pdwReserved"`
}

func (o *CustomHeader) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ClassIDs != nil && o.InterfacesCount == 0 {
		o.InterfacesCount = uint32(len(o.ClassIDs))
	}
	if o.Sizes != nil && o.InterfacesCount == 0 {
		o.InterfacesCount = uint32(len(o.Sizes))
	}
	if o.InterfacesCount < uint32(1) || o.InterfacesCount > uint32(10) {
		return fmt.Errorf("InterfacesCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *CustomHeader) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalSize); err != nil {
		return err
	}
	if err := w.WriteData(o.HeaderSize); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.DestinationContext); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfacesCount); err != nil {
		return err
	}
	if o.ClassInfoClassID != nil {
		if err := o.ClassInfoClassID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ClassIDs != nil || o.InterfacesCount > 0 {
		_ptr_pclsid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfacesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ClassIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ClassIDs[i1] != nil {
					if err := o.ClassIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ClassID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ClassIDs, _ptr_pclsid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Sizes != nil || o.InterfacesCount > 0 {
		_ptr_pSizes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfacesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Sizes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Sizes[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Sizes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Sizes, _ptr_pSizes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved pdwReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *CustomHeader) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.HeaderSize); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.DestinationContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfacesCount); err != nil {
		return err
	}
	if o.ClassInfoClassID == nil {
		o.ClassInfoClassID = &ClassID{}
	}
	if err := o.ClassInfoClassID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_pclsid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfacesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfacesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ClassIDs", sizeInfo[0])
		}
		o.ClassIDs = make([]*ClassID, sizeInfo[0])
		for i1 := range o.ClassIDs {
			i1 := i1
			if o.ClassIDs[i1] == nil {
				o.ClassIDs[i1] = &ClassID{}
			}
			if err := o.ClassIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pclsid := func(ptr interface{}) { o.ClassIDs = *ptr.(*[]*ClassID) }
	if err := w.ReadPointer(&o.ClassIDs, _s_pclsid, _ptr_pclsid); err != nil {
		return err
	}
	_ptr_pSizes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfacesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfacesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Sizes", sizeInfo[0])
		}
		o.Sizes = make([]uint32, sizeInfo[0])
		for i1 := range o.Sizes {
			i1 := i1
			if err := w.ReadData(&o.Sizes[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSizes := func(ptr interface{}) { o.Sizes = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Sizes, _s_pSizes, _ptr_pSizes); err != nil {
		return err
	}
	// reserved pdwReserved
	var _pdwReserved uint32
	_ptr_pdwReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_pdwReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pdwReserved := func(ptr interface{}) { _pdwReserved = *ptr.(*uint32) }
	if err := w.ReadPointer(&_pdwReserved, _s_pdwReserved, _ptr_pdwReserved); err != nil {
		return err
	}
	return nil
}

// PropertiesOutInfo structure represents PropsOutInfo RPC structure.
//
// The PropsOutInfo structure represents a collection of interfaces that the object
// implements and that are returned to the client. If the object does not support a
// particular interface requested by the client, it also sends an error back using this
// structure.
//
// CLSID_PropsOutInfo (see section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type PropertiesOutInfo struct {
	// cIfs:  This MUST contain the number of interfaces being returned by the server. This
	// value MUST be between 1 and MAX_REQUESTED_INTERFACES (see section 2.2.28.1).
	InterfacesCount uint32 `idl:"name:cIfs" json:"interfaces_count"`
	// piid:  This MUST be an array of IIDs identifying the interfaces returned by the server.
	IIDs []*IID `idl:"name:piid;size_is:(cIfs)" json:"iids"`
	// phresults:   This MUST be an array of status codes indicating the success or failure
	// of each attempt to return an interface requested by the client. For each array location
	// containing a zero value, a non-NULL MInterfacePointer pointer MUST be present in
	// the corresponding location in the ppIntfData array. For each array location containing
	// a negative value, a NULL MUST be present in the corresponding location in the ppIntfData
	// array.
	HResults []int32 `idl:"name:phresults;size_is:(cIfs)" json:"hresults"`
	// ppIntfData:   This MUST be an array of MInterfacePointer pointers containing the
	// OBJREFs for the interfaces returned by the server.
	InterfaceData []*InterfacePointer `idl:"name:ppIntfData;size_is:(cIfs)" json:"interface_data"`
}

func (o *PropertiesOutInfo) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.IIDs != nil && o.InterfacesCount == 0 {
		o.InterfacesCount = uint32(len(o.IIDs))
	}
	if o.HResults != nil && o.InterfacesCount == 0 {
		o.InterfacesCount = uint32(len(o.HResults))
	}
	if o.InterfaceData != nil && o.InterfacesCount == 0 {
		o.InterfacesCount = uint32(len(o.InterfaceData))
	}
	if o.InterfacesCount < uint32(1) || o.InterfacesCount > uint32(32768) {
		return fmt.Errorf("InterfacesCount is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *PropertiesOutInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.InterfacesCount); err != nil {
		return err
	}
	if o.IIDs != nil || o.InterfacesCount > 0 {
		_ptr_piid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfacesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.IIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.IIDs[i1] != nil {
					if err := o.IIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.IIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IIDs, _ptr_piid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.HResults != nil || o.InterfacesCount > 0 {
		_ptr_phresults := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfacesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.HResults {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.HResults[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.HResults); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(int32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HResults, _ptr_phresults); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.InterfaceData != nil || o.InterfacesCount > 0 {
		_ptr_ppIntfData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InterfacesCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.InterfaceData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.InterfaceData[i1] != nil {
					_ptr_ppIntfData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.InterfaceData[i1] != nil {
							if err := o.InterfaceData[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.InterfaceData[i1], _ptr_ppIntfData); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.InterfaceData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InterfaceData, _ptr_ppIntfData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertiesOutInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.InterfacesCount); err != nil {
		return err
	}
	_ptr_piid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfacesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfacesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.IIDs", sizeInfo[0])
		}
		o.IIDs = make([]*IID, sizeInfo[0])
		for i1 := range o.IIDs {
			i1 := i1
			if o.IIDs[i1] == nil {
				o.IIDs[i1] = &IID{}
			}
			if err := o.IIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_piid := func(ptr interface{}) { o.IIDs = *ptr.(*[]*IID) }
	if err := w.ReadPointer(&o.IIDs, _s_piid, _ptr_piid); err != nil {
		return err
	}
	_ptr_phresults := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfacesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfacesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.HResults", sizeInfo[0])
		}
		o.HResults = make([]int32, sizeInfo[0])
		for i1 := range o.HResults {
			i1 := i1
			if err := w.ReadData(&o.HResults[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_phresults := func(ptr interface{}) { o.HResults = *ptr.(*[]int32) }
	if err := w.ReadPointer(&o.HResults, _s_phresults, _ptr_phresults); err != nil {
		return err
	}
	_ptr_ppIntfData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.InterfacesCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.InterfacesCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceData", sizeInfo[0])
		}
		o.InterfaceData = make([]*InterfacePointer, sizeInfo[0])
		for i1 := range o.InterfaceData {
			i1 := i1
			_ptr_ppIntfData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.InterfaceData[i1] == nil {
					o.InterfaceData[i1] = &InterfacePointer{}
				}
				if err := o.InterfaceData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppIntfData := func(ptr interface{}) { o.InterfaceData[i1] = *ptr.(**InterfacePointer) }
			if err := w.ReadPointer(&o.InterfaceData[i1], _s_ppIntfData, _ptr_ppIntfData); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ppIntfData := func(ptr interface{}) { o.InterfaceData = *ptr.(*[]*InterfacePointer) }
	if err := w.ReadPointer(&o.InterfaceData, _s_ppIntfData, _ptr_ppIntfData); err != nil {
		return err
	}
	return nil
}

// SecurityInfoData structure represents SecurityInfoData RPC structure.
//
// The SecurityInfoData structure SHOULD NOT be sent and MUST be ignored on receipt.<34>
//
// CLSID_SecurityInfo (see section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type SecurityInfoData struct {
	// dwAuthnFlags:  This MUST be set to zero and MUST be ignored on receipt.
	AuthnFlags uint32 `idl:"name:dwAuthnFlags" json:"authn_flags"`
	// pServerInfo:   This SHOULD be NULL and MUST be ignored on receipt. For details, see
	// section 2.2.22.2.7.1.<35>
	ServerInfo *COMServerInfo `idl:"name:pServerInfo" json:"server_info"`
	// pdwReserved:   This MUST be set to NULL and MUST be ignored on receipt.
	_ uint32 `idl:"name:pdwReserved"`
}

func (o *SecurityInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SecurityInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.AuthnFlags); err != nil {
		return err
	}
	if o.ServerInfo != nil {
		_ptr_pServerInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ServerInfo != nil {
				if err := o.ServerInfo.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&COMServerInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerInfo, _ptr_pServerInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved pdwReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	return nil
}
func (o *SecurityInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuthnFlags); err != nil {
		return err
	}
	_ptr_pServerInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ServerInfo == nil {
			o.ServerInfo = &COMServerInfo{}
		}
		if err := o.ServerInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pServerInfo := func(ptr interface{}) { o.ServerInfo = *ptr.(**COMServerInfo) }
	if err := w.ReadPointer(&o.ServerInfo, _s_pServerInfo, _ptr_pServerInfo); err != nil {
		return err
	}
	// reserved pdwReserved
	var _pdwReserved uint32
	_ptr_pdwReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_pdwReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pdwReserved := func(ptr interface{}) { _pdwReserved = *ptr.(*uint32) }
	if err := w.ReadPointer(&_pdwReserved, _s_pdwReserved, _ptr_pdwReserved); err != nil {
		return err
	}
	return nil
}

// SCMRequestInfoData structure represents ScmRequestInfoData RPC structure.
//
// The ScmRequestInfoData structure contains a customREMOTE_REQUEST_SCM_INFO structure.
//
// CLSID_ScmRequestInfo (see section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type SCMRequestInfoData struct {
	// pdwReserved:   This MUST be set to NULL and MUST be ignored on receipt.
	_ uint32 `idl:"name:pdwReserved"`
	// remoteRequest:   This MUST specify a customREMOTE_REQUEST_SCM_INFO structure. This
	// field MUST NOT be NULL.
	RemoteRequest *CustomRemoteRequestSCMInfo `idl:"name:remoteRequest" json:"remote_request"`
}

func (o *SCMRequestInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SCMRequestInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	// reserved pdwReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if o.RemoteRequest != nil {
		_ptr_remoteRequest := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RemoteRequest != nil {
				if err := o.RemoteRequest.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CustomRemoteRequestSCMInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteRequest, _ptr_remoteRequest); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SCMRequestInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	// reserved pdwReserved
	var _pdwReserved uint32
	_ptr_pdwReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_pdwReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pdwReserved := func(ptr interface{}) { _pdwReserved = *ptr.(*uint32) }
	if err := w.ReadPointer(&_pdwReserved, _s_pdwReserved, _ptr_pdwReserved); err != nil {
		return err
	}
	_ptr_remoteRequest := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RemoteRequest == nil {
			o.RemoteRequest = &CustomRemoteRequestSCMInfo{}
		}
		if err := o.RemoteRequest.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_remoteRequest := func(ptr interface{}) { o.RemoteRequest = *ptr.(**CustomRemoteRequestSCMInfo) }
	if err := w.ReadPointer(&o.RemoteRequest, _s_remoteRequest, _ptr_remoteRequest); err != nil {
		return err
	}
	return nil
}

// SCMReplyInfoData structure represents ScmReplyInfoData RPC structure.
//
// The ScmReplyInfoData structure contains a customREMOTE_REPLY_SCM_INFO structure.
//
// CLSID_ScmReplyInfo (see section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type SCMReplyInfoData struct {
	// pdwReserved:   This MUST be set to NULL and MUST be ignored on receipt.
	_ uint32 `idl:"name:pdwReserved"`
	// remoteReply:  This MUST specify the customREMOTE_REPLY_SCM_INFO for the object exporter
	// of the server object.
	RemoteReply *CustomRemoteReplySCMInfo `idl:"name:remoteReply" json:"remote_reply"`
}

func (o *SCMReplyInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SCMReplyInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	// reserved pdwReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if o.RemoteReply != nil {
		_ptr_remoteReply := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RemoteReply != nil {
				if err := o.RemoteReply.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&CustomRemoteReplySCMInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RemoteReply, _ptr_remoteReply); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SCMReplyInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	// reserved pdwReserved
	var _pdwReserved uint32
	_ptr_pdwReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&_pdwReserved); err != nil {
			return err
		}
		return nil
	})
	_s_pdwReserved := func(ptr interface{}) { _pdwReserved = *ptr.(*uint32) }
	if err := w.ReadPointer(&_pdwReserved, _s_pdwReserved, _ptr_pdwReserved); err != nil {
		return err
	}
	_ptr_remoteReply := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RemoteReply == nil {
			o.RemoteReply = &CustomRemoteReplySCMInfo{}
		}
		if err := o.RemoteReply.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_remoteReply := func(ptr interface{}) { o.RemoteReply = *ptr.(**CustomRemoteReplySCMInfo) }
	if err := w.ReadPointer(&o.RemoteReply, _s_remoteReply, _ptr_remoteReply); err != nil {
		return err
	}
	return nil
}

// InstanceInfoData structure represents InstanceInfoData RPC structure.
//
// The InstanceInfoData structure contains data related to persistent activations; that
// is, object activations in which the newly created object is immediately initialized
// with state from a previously persisted instance of the object. For more information,
// see [MSDN-COM], [MSDN-SS], and [MSDN-IPersistFile].
//
// CLSID_InstanceInfo (see section 1.9) is used to identify this property in the CustomHeader.pclsid
// array.
type InstanceInfoData struct {
	// fileName:   This MAY contain a string to be used to initialize the object.<30>
	FileName string `idl:"name:fileName;string" json:"file_name"`
	// mode:   This MUST contain an implementation-specific value and MAY be ignored on
	// receipt.<31>
	Mode uint32 `idl:"name:mode" json:"mode"`
	// ifdROT:   The pointer MUST be set to NULL and MUST be ignored on receipt.
	ROT *InterfacePointer `idl:"name:ifdROT" json:"rot"`
	// ifdStg:  This MAY contain a marshaled OBJREF to be used to initialize the object.<32>
	Storage *InterfacePointer `idl:"name:ifdStg" json:"storage"`
}

func (o *InstanceInfoData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *InstanceInfoData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.FileName != "" {
		_ptr_fileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FileName, _ptr_fileName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Mode); err != nil {
		return err
	}
	if o.ROT != nil {
		_ptr_ifdROT := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ROT != nil {
				if err := o.ROT.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ROT, _ptr_ifdROT); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Storage != nil {
		_ptr_ifdStg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Storage != nil {
				if err := o.Storage.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Storage, _ptr_ifdStg); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InstanceInfoData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_fileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
			return err
		}
		return nil
	})
	_s_fileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FileName, _s_fileName, _ptr_fileName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mode); err != nil {
		return err
	}
	_ptr_ifdROT := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ROT == nil {
			o.ROT = &InterfacePointer{}
		}
		if err := o.ROT.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ifdROT := func(ptr interface{}) { o.ROT = *ptr.(**InterfacePointer) }
	if err := w.ReadPointer(&o.ROT, _s_ifdROT, _ptr_ifdROT); err != nil {
		return err
	}
	_ptr_ifdStg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Storage == nil {
			o.Storage = &InterfacePointer{}
		}
		if err := o.Storage.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ifdStg := func(ptr interface{}) { o.Storage = *ptr.(**InterfacePointer) }
	if err := w.ReadPointer(&o.Storage, _s_ifdStg, _ptr_ifdStg); err != nil {
		return err
	}
	return nil
}

// SPDFlags type represents SPD_FLAGS RPC enumeration.
type SPDFlags uint16

var (
	SPDFlagsUseConsoleSession    SPDFlags = 1
	SPDFlagsUseDefaultAuthnLevel SPDFlags = 2
)

func (o SPDFlags) String() string {
	switch o {
	case SPDFlagsUseConsoleSession:
		return "SPDFlagsUseConsoleSession"
	case SPDFlagsUseDefaultAuthnLevel:
		return "SPDFlagsUseDefaultAuthnLevel"
	}
	return "Invalid"
}

// SpecialPropertiesData structure represents SpecialPropertiesData RPC structure.
//
// The SpecialPropertiesData structure contains miscellaneous parameters specified by
// the client for an activation request.
//
// CLSID_SpecialSystemProperties (see section 1.9) is used to identify this property
// in the CustomHeader.pclsid array.
type SpecialPropertiesData struct {
	// dwSessionId:  A value that uniquely identifies a logon session on the server. The
	// value 0xFFFFFFFF indicates that any logon session is acceptable to the client.
	SessionID uint32 `idl:"name:dwSessionId" json:"session_id"`
	// fRemoteThisSessionId:   This MUST be set to TRUE (0x00000001) if dwSessionId is not
	// 0xFFFFFFFF; otherwise this MUST be set to FALSE (0x00000000). This field MUST be
	// ignored on receipt.
	RemoteThisSessionID int32 `idl:"name:fRemoteThisSessionId" json:"remote_this_session_id"`
	// fClientImpersonating:  This SHOULD be set to FALSE (0x00000000) and MUST be ignored
	// on receipt.<24>
	ClientImpersonating int32 `idl:"name:fClientImpersonating" json:"client_impersonating"`
	// fPartitionIDPresent:  This MUST contain an implementation-specific value and MAY
	// be ignored on receipt.<25>
	PartitionIDPresent int32 `idl:"name:fPartitionIDPresent" json:"partition_id_present"`
	// dwDefaultAuthnLvl:  This MUST contain an implementation-specific value and MUST be
	// ignored on receipt.<26>
	DefaultAuthnLevel uint32 `idl:"name:dwDefaultAuthnLvl" json:"default_authn_level"`
	// guidPartition:   This MUST contain an implementation-specific value specified by
	// higher-layer protocols and MAY be ignored on receipt.<27>
	Partition *dtyp.GUID `idl:"name:guidPartition" json:"partition"`
	// dwPRTFlags:  This MUST be set to zero and MUST be ignored on receipt.
	PartitionFlags uint32 `idl:"name:dwPRTFlags" json:"partition_flags"`
	// dwOrigClsctx:  This MUST contain an implementation-specific value and SHOULD be ignored
	// on receipt.<28>
	OrigClassContext uint32 `idl:"name:dwOrigClsctx" json:"orig_class_context"`
	// dwFlags:  This is a set of bitflags, defined as follows.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| SPD_FLAG_USE_CONSOLE_SESSION 0x00000001 | If this bit is set, the object resolver is requested to create the object        |
	//	|                                         | exporter in the console logon session. If this bit is not set, the object        |
	//	|                                         | resolver is requested to create the object exporter in the logon session         |
	//	|                                         | specified in the dwSessionID field.                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Object resolvers MUST ignore any other bits that are set in the dwFlags field.
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// Reserved1:  This MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint32 `idl:"name:Reserved1"`
	// Reserved2:  This MUST be set to zero when sent and MUST be ignored on receipt.
	_ uint64 `idl:"name:Reserved2"`
	// Reserved3:  This can contain arbitrary values and MUST be ignored on receipt.
	//
	// This structure has an alternate definition that is specified as follows.
	//
	// All the fields have the same meaning as the corresponding fields in the first structure.
	// A DCOM server MUST accept as valid both definitions. A DCOM client SHOULD<29> use
	// SpecialPropertiesData in activation requests.
	_ []uint32 `idl:"name:Reserved3"`
}

func (o *SpecialPropertiesData) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SpecialPropertiesData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteThisSessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientImpersonating); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionIDPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAuthnLevel); err != nil {
		return err
	}
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PartitionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.OrigClassContext); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved Reserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved Reserved2
	if err := w.WriteData(uint64(0)); err != nil {
		return err
	}
	// reserved Reserved3
	for i1 := 0; uint64(i1) < 5; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(8); err != nil {
		return err
	}
	return nil
}
func (o *SpecialPropertiesData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteThisSessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientImpersonating); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionIDPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAuthnLevel); err != nil {
		return err
	}
	if o.Partition == nil {
		o.Partition = &dtyp.GUID{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.OrigClassContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 uint32
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	// reserved Reserved2
	var _Reserved2 uint64
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 []uint32
	_Reserved3 = make([]uint32, 5)
	for i1 := range _Reserved3 {
		i1 := i1
		if err := w.ReadData(&_Reserved3[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(8); err != nil {
		return err
	}
	return nil
}

// SpecialPropertiesDataAlternate structure represents SpecialPropertiesData_Alternate RPC structure.
type SpecialPropertiesDataAlternate struct {
	SessionID           uint32     `idl:"name:dwSessionId" json:"session_id"`
	RemoteThisSessionID int32      `idl:"name:fRemoteThisSessionId" json:"remote_this_session_id"`
	ClientImpersonating int32      `idl:"name:fClientImpersonating" json:"client_impersonating"`
	PartitionIDPresent  int32      `idl:"name:fPartitionIDPresent" json:"partition_id_present"`
	DefaultAuthnLevel   uint32     `idl:"name:dwDefaultAuthnLvl" json:"default_authn_level"`
	Partition           *dtyp.GUID `idl:"name:guidPartition" json:"partition"`
	PartitionFlags      uint32     `idl:"name:dwPRTFlags" json:"partition_flags"`
	OrigClassContext    uint32     `idl:"name:dwOrigClsctx" json:"orig_class_context"`
	Flags               uint32     `idl:"name:dwFlags" json:"flags"`
	_                   []uint32   `idl:"name:Reserved3"`
}

func (o *SpecialPropertiesDataAlternate) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SpecialPropertiesDataAlternate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.RemoteThisSessionID); err != nil {
		return err
	}
	if err := w.WriteData(o.ClientImpersonating); err != nil {
		return err
	}
	if err := w.WriteData(o.PartitionIDPresent); err != nil {
		return err
	}
	if err := w.WriteData(o.DefaultAuthnLevel); err != nil {
		return err
	}
	if o.Partition != nil {
		if err := o.Partition.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PartitionFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.OrigClassContext); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved Reserved3
	for i1 := 0; uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *SpecialPropertiesDataAlternate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.RemoteThisSessionID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClientImpersonating); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionIDPresent); err != nil {
		return err
	}
	if err := w.ReadData(&o.DefaultAuthnLevel); err != nil {
		return err
	}
	if o.Partition == nil {
		o.Partition = &dtyp.GUID{}
	}
	if err := o.Partition.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PartitionFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.OrigClassContext); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 []uint32
	_Reserved3 = make([]uint32, 8)
	for i1 := range _Reserved3 {
		i1 := i1
		if err := w.ReadData(&_Reserved3[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RemoteUnknown2 structure represents IRemUnknown2 RPC structure.
type RemoteUnknown2 InterfacePointer

func (o *RemoteUnknown2) InterfacePointer() *InterfacePointer { return (*InterfacePointer)(o) }

func (o *RemoteUnknown2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RemoteUnknown2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteUnknown2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteUnknown2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
type Unknown InterfacePointer

func (o *Unknown) InterfacePointer() *InterfacePointer { return (*InterfacePointer)(o) }

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

// RemoteUnknown structure represents IRemUnknown RPC structure.
type RemoteUnknown InterfacePointer

func (o *RemoteUnknown) InterfacePointer() *InterfacePointer { return (*InterfacePointer)(o) }

func (o *RemoteUnknown) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RemoteUnknown) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteUnknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteUnknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
