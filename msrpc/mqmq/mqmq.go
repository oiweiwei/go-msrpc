// The mqmq package implements the MQMQ client protocol.
//
// # Introduction
//
// Message Queuing (MSMQ): Data Structures contains common definitions and data structures
// that are used in various protocols in the set of Microsoft Message Queuing protocols.
// The documentation for individual protocols contains references to this document,
// as needed.
//
// # Overview
//
// The common definitions, naming formats, structures, data types, and error codes defined
// in this document are used by the member protocols of the Microsoft Message Queuing
// (MSMQ) protocol set.
package mqmq

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
	GoPackage = "mqmq"
)

// TransactionUOW structure represents XACTUOW RPC structure.
//
// An XACTUOW is a structure that serves as a unique identifier for a transactional
// unit of work. An XACTUOW contains 16 unsigned single-byte characters representing
// a GUID ([MS-DTYP] section 2.3.4) and is defined as follows:
type TransactionUOW struct {
	// rgb:  An array of unsigned single-byte characters that contains a globally unique
	// identifier (GUID).
	Data []byte `idl:"name:rgb" json:"data"`
}

func (o *TransactionUOW) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransactionUOW) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *TransactionUOW) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]byte, 16)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Blob structure represents BLOB RPC structure.
//
// The BLOB structure defines a counted array of unsigned characters.
type Blob struct {
	// cbSize:  A 32-bit unsigned integer that specifies the size of the array of unsigned
	// characters pointed to by pBlobData.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// pBlobData:  An array of 8-bit unsigned characters.
	BlobData []byte `idl:"name:pBlobData;size_is:(cbSize)" json:"blob_data"`
}

func (o *Blob) xxx_PreparePayload(ctx context.Context) error {
	if o.BlobData != nil && o.Length == 0 {
		o.Length = uint32(len(o.BlobData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Blob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.BlobData != nil || o.Length > 0 {
		_ptr_pBlobData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.BlobData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.BlobData[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.BlobData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.BlobData, _ptr_pBlobData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Blob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pBlobData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.BlobData", sizeInfo[0])
		}
		o.BlobData = make([]byte, sizeInfo[0])
		for i1 := range o.BlobData {
			i1 := i1
			if err := w.ReadData(&o.BlobData[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pBlobData := func(ptr interface{}) { o.BlobData = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.BlobData, _s_pBlobData, _ptr_pBlobData); err != nil {
		return err
	}
	return nil
}

// ByteArray structure represents CAUB RPC structure.
//
// The CAUB structure defines a counted array of unsigned characters.
type ByteArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of unsigned characters.
	Elems []byte `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *ByteArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Elems[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]byte, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if err := w.ReadData(&o.Elems[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// UshortArray structure represents CAUI RPC structure.
//
// The CAUI structure defines a counted array of unsigned short integers.
type UshortArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of unsigned short integers.
	Elems []uint16 `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *UshortArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UshortArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Elems[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UshortArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]uint16, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if err := w.ReadData(&o.Elems[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// LongArray structure represents CAL RPC structure.
//
// The CAL structure defines a counted array of 32-bit unsigned integers.
type LongArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of 32-bit unsigned integers.
	Elems []int32 `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *LongArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LongArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Elems[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(int32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *LongArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]int32, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if err := w.ReadData(&o.Elems[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]int32) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// Uint32Array structure represents CAUL RPC structure.
//
// The CAUL structure defines a counted array of 32-bit unsigned integers.
type Uint32Array struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of 32-bit unsigned integers.
	Elems []uint32 `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *Uint32Array) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Elems[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]uint32, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if err := w.ReadData(&o.Elems[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// UlargeIntegerArray structure represents CAUH RPC structure.
//
// The CAUH structure defines a counted array of ULARGE_INTEGER (section 2.2.17) values.
type UlargeIntegerArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of ULARGE_INTEGER (section 2.2.17) values.
	Elems []*dtyp.UlargeInteger `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *UlargeIntegerArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UlargeIntegerArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elems[i1] != nil {
					if err := o.Elems[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UlargeIntegerArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]*dtyp.UlargeInteger, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if o.Elems[i1] == nil {
				o.Elems[i1] = &dtyp.UlargeInteger{}
			}
			if err := o.Elems[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]*dtyp.UlargeInteger) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// ClassIDArray structure represents CACLSID RPC structure.
//
// The CACLSID structure defines a counted array of GUID (as specified in [MS-DTYP]
// section 2.3.4) values.
type ClassIDArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of GUID (as specified in [MS-DTYP] section 2.3.4) values.
	Elems []*dtyp.GUID `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *ClassIDArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elems[i1] != nil {
					if err := o.Elems[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if o.Elems[i1] == nil {
				o.Elems[i1] = &dtyp.GUID{}
			}
			if err := o.Elems[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]*dtyp.GUID) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// UnicodeStringArray structure represents CALPWSTR RPC structure.
//
// The CALPWSTR structure defines a counted array of wchar_t* values.
type UnicodeStringArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of wchar_t* values.
	Elems []string `idl:"name:pElems;size_is:(cElems);string" json:"elems"`
}

func (o *UnicodeStringArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeStringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elems[i1] != "" {
					_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.Elems[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Elems[i1], _ptr_pElems); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeStringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]string, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Elems[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_pElems := func(ptr interface{}) { o.Elems[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Elems[i1], _s_pElems, _ptr_pElems); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// PropertyVariantArray structure represents CAPROPVARIANT RPC structure.
//
// The CAPROPVARIANT structure defines a counted array of PROPVARIANT (section 2.2.13.2)
// values.
type PropertyVariantArray struct {
	// cElems:  MUST be set to the total number of elements of the array.
	ElemsCount uint32 `idl:"name:cElems" json:"elems_count"`
	// pElems:  An array of PROPVARIANT (section 2.2.13.2) values.
	Elems []*PropertyVariant `idl:"name:pElems;size_is:(cElems)" json:"elems"`
}

func (o *PropertyVariantArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Elems != nil && o.ElemsCount == 0 {
		o.ElemsCount = uint32(len(o.Elems))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariantArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ElemsCount); err != nil {
		return err
	}
	if o.Elems != nil || o.ElemsCount > 0 {
		_ptr_pElems := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ElemsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Elems {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Elems[i1] != nil {
					if err := o.Elems[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Elems); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Elems, _ptr_pElems); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariantArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElemsCount); err != nil {
		return err
	}
	_ptr_pElems := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ElemsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ElemsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Elems", sizeInfo[0])
		}
		o.Elems = make([]*PropertyVariant, sizeInfo[0])
		for i1 := range o.Elems {
			i1 := i1
			if o.Elems[i1] == nil {
				o.Elems[i1] = &PropertyVariant{}
			}
			if err := o.Elems[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pElems := func(ptr interface{}) { o.Elems = *ptr.(*[]*PropertyVariant) }
	if err := w.ReadPointer(&o.Elems, _s_pElems, _ptr_pElems); err != nil {
		return err
	}
	return nil
}

// VarEnum type represents VARENUM RPC enumeration.
//
// The following values are used in the discriminant field, vt, of the PROPVARIANT (section
// 2.2.13) type.
//
// The PROPVARIANT (section 2.2.13) type constants are defined in the VARENUM enumeration,
// as follows:
type VarEnum uint16

var (
	// VT_EMPTY:  (0x0000): The type of the contained field is undefined. When this flag
	// is specified, the PROPVARIANT (section 2.2.13) MUST NOT contain a data field.
	VarEmpty VarEnum = 0
	// VT_NULL:  (0x0001): The type of the contained field is NULL. When this flag is specified,
	// the PROPVARIANT (section 2.2.13) MUST NOT contain a data field.
	VarNull VarEnum = 1
	// VT_I2:  (0x0002): The type of the contained field MUST be a 2-byte signed integer.
	VarEnumI2 VarEnum = 2
	// VT_I4:  (0x0003): The type of the contained field MUST be a 4-byte signed integer.
	VarEnumI4 VarEnum = 3
	// VT_BOOL:  (0x000B): The type of the contained field MUST be VARIANT_BOOL (section
	// 2.2.14).
	VarEnumBool VarEnum = 11
	// VT_VARIANT:  (0x000C): The type of the contained field MUST be CAPROPVARIANT (section
	// 2.2.16.8). It MUST appear with the bit flag VT_VECTOR.
	VarEnumVariant VarEnum = 12
	// VT_I1:  (0x0010): The type of the contained field MUST be a 1-byte integer.
	VarEnumI1 VarEnum = 16
	// VT_UI1:  (0x0011): The type of the contained field MUST be a 1-byte unsigned integer.
	VarEnumUI1 VarEnum = 17
	// VT_UI2:  (0x0012): The type of the contained field MUST be a 2-byte unsigned integer.
	VarEnumUI2 VarEnum = 18
	// VT_UI4:  (0x0013): The type of the contained field MUST be a 4-byte unsigned integer.
	VarEnumUI4 VarEnum = 19
	// VT_I8:  (0x0014): The type of the contained field MUST be an 8-byte signed integer.
	VarEnumI8 VarEnum = 20
	// VT_UI8:  (0x0015): The type of the contained field MUST be an 8-byte unsigned integer.
	VarEnumUI8 VarEnum = 21
	// VT_LPWSTR:  (0x001F): The type of the contained field MUST be an LPWSTR (as specified
	// in [MS-DTYP] section 2.2.36), a null-terminated Unicode string.
	VarUnicodeString VarEnum = 31
	// VT_BLOB:  (0x0041): The type of the contained field MUST be a binary large object
	// (BLOB) (section 2.2.15).
	VarEnumBlob VarEnum = 65
	// VT_CLSID:  (0x0048): The type of the contained field MUST be a pointer to a GUID
	// (as specified in [MS-DTYP] section 2.3.4) value.
	VarEnumClassID VarEnum = 72
	// VT_VECTOR:   (0x1000): The type of the contained field MUST be combined with other
	// values by using the bitwise OR operation to indicate a counted field. The type of
	// the contained field MUST be a COUNTEDARRAY (section 2.2.16).
	VarEnumVector VarEnum = 4096
)

func (o VarEnum) String() string {
	switch o {
	case VarEmpty:
		return "VarEmpty"
	case VarNull:
		return "VarNull"
	case VarEnumI2:
		return "VarEnumI2"
	case VarEnumI4:
		return "VarEnumI4"
	case VarEnumBool:
		return "VarEnumBool"
	case VarEnumVariant:
		return "VarEnumVariant"
	case VarEnumI1:
		return "VarEnumI1"
	case VarEnumUI1:
		return "VarEnumUI1"
	case VarEnumUI2:
		return "VarEnumUI2"
	case VarEnumUI4:
		return "VarEnumUI4"
	case VarEnumI8:
		return "VarEnumI8"
	case VarEnumUI8:
		return "VarEnumUI8"
	case VarUnicodeString:
		return "VarUnicodeString"
	case VarEnumBlob:
		return "VarEnumBlob"
	case VarEnumClassID:
		return "VarEnumClassID"
	case VarEnumVector:
		return "VarEnumVector"
	}
	return "Invalid"
}

// PropertyVariant structure represents PROPVARIANT RPC structure.
type PropertyVariant struct {
	VT       uint16                    `idl:"name:vt" json:"vt"`
	_        uint8                     `idl:"name:wReserved1"`
	_        uint8                     `idl:"name:wReserved2"`
	_        uint32                    `idl:"name:wReserved3"`
	VarUnion *PropertyVariant_VarUnion `idl:"name:_varUnion;switch_is:vt" json:"var_union"`
}

func (o *PropertyVariant) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.VT); err != nil {
		return err
	}
	// reserved wReserved1
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	// reserved wReserved2
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	// reserved wReserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	_swVarUnion := uint16(o.VT)
	if o.VarUnion != nil {
		if err := o.VarUnion.MarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
			return err
		}
	} else {
		if err := (&PropertyVariant_VarUnion{}).MarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.VT); err != nil {
		return err
	}
	// reserved wReserved1
	var _wReserved1 uint8
	if err := w.ReadData(&_wReserved1); err != nil {
		return err
	}
	// reserved wReserved2
	var _wReserved2 uint8
	if err := w.ReadData(&_wReserved2); err != nil {
		return err
	}
	// reserved wReserved3
	var _wReserved3 uint32
	if err := w.ReadData(&_wReserved3); err != nil {
		return err
	}
	if o.VarUnion == nil {
		o.VarUnion = &PropertyVariant_VarUnion{}
	}
	_swVarUnion := uint16(o.VT)
	if err := o.VarUnion.UnmarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion structure represents PROPVARIANT union anonymous member.
type PropertyVariant_VarUnion struct {
	// Types that are assignable to Value
	//
	// *PropertyVariant_VarUnion_0
	// *PropertyVariant_VarUnion_Char
	// *PropertyVariant_VarUnion_Byte
	// *PropertyVariant_VarUnion_Short
	// *PropertyVariant_VarUnion_Ushort
	// *PropertyVariant_VarUnion_Long
	// *PropertyVariant_VarUnion_Ulong
	// *PropertyVariant_VarUnion_LargeInteger
	// *PropertyVariant_VarUnion_UlargeInteger
	// *PropertyVariant_VarUnion_Bool
	// *PropertyVariant_VarUnion_UUID
	// *PropertyVariant_VarUnion_Blob
	// *PropertyVariant_VarUnion_UnicodeString
	// *PropertyVariant_VarUnion_ByteArray
	// *PropertyVariant_VarUnion_UshortArray
	// *PropertyVariant_VarUnion_LongArray
	// *PropertyVariant_VarUnion_Uint32Array
	// *PropertyVariant_VarUnion_UlargeIntegerArray
	// *PropertyVariant_VarUnion_UUIDArray
	// *PropertyVariant_VarUnion_UnicodeStringArray
	// *PropertyVariant_VarUnion_PropertyVariantArray
	Value is_PropertyVariant_VarUnion `json:"value"`
}

func (o *PropertyVariant_VarUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *PropertyVariant_VarUnion_Char:
		if value != nil {
			return value.Char
		}
	case *PropertyVariant_VarUnion_Byte:
		if value != nil {
			return value.Byte
		}
	case *PropertyVariant_VarUnion_Short:
		if value != nil {
			return value.Short
		}
	case *PropertyVariant_VarUnion_Ushort:
		if value != nil {
			return value.Ushort
		}
	case *PropertyVariant_VarUnion_Long:
		if value != nil {
			return value.Long
		}
	case *PropertyVariant_VarUnion_Ulong:
		if value != nil {
			return value.Ulong
		}
	case *PropertyVariant_VarUnion_LargeInteger:
		if value != nil {
			return value.LargeInteger
		}
	case *PropertyVariant_VarUnion_UlargeInteger:
		if value != nil {
			return value.UlargeInteger
		}
	case *PropertyVariant_VarUnion_Bool:
		if value != nil {
			return value.Bool
		}
	case *PropertyVariant_VarUnion_UUID:
		if value != nil {
			return value.UUID
		}
	case *PropertyVariant_VarUnion_Blob:
		if value != nil {
			return value.Blob
		}
	case *PropertyVariant_VarUnion_UnicodeString:
		if value != nil {
			return value.UnicodeString
		}
	case *PropertyVariant_VarUnion_ByteArray:
		if value != nil {
			return value.ByteArray
		}
	case *PropertyVariant_VarUnion_UshortArray:
		if value != nil {
			return value.UshortArray
		}
	case *PropertyVariant_VarUnion_LongArray:
		if value != nil {
			return value.LongArray
		}
	case *PropertyVariant_VarUnion_Uint32Array:
		if value != nil {
			return value.Uint32Array
		}
	case *PropertyVariant_VarUnion_UlargeIntegerArray:
		if value != nil {
			return value.UlargeIntegerArray
		}
	case *PropertyVariant_VarUnion_UUIDArray:
		if value != nil {
			return value.UUIDArray
		}
	case *PropertyVariant_VarUnion_UnicodeStringArray:
		if value != nil {
			return value.UnicodeStringArray
		}
	case *PropertyVariant_VarUnion_PropertyVariantArray:
		if value != nil {
			return value.PropertyVariantArray
		}
	}
	return nil
}

type is_PropertyVariant_VarUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_PropertyVariant_VarUnion()
}

func (o *PropertyVariant_VarUnion) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *PropertyVariant_VarUnion_0:
		switch sw {
		case uint16(0),
			uint16(1):
			return sw
		}
		return uint16(0)
	case *PropertyVariant_VarUnion_Char:
		return uint16(16)
	case *PropertyVariant_VarUnion_Byte:
		return uint16(17)
	case *PropertyVariant_VarUnion_Short:
		return uint16(2)
	case *PropertyVariant_VarUnion_Ushort:
		return uint16(18)
	case *PropertyVariant_VarUnion_Long:
		return uint16(3)
	case *PropertyVariant_VarUnion_Ulong:
		return uint16(19)
	case *PropertyVariant_VarUnion_LargeInteger:
		return uint16(20)
	case *PropertyVariant_VarUnion_UlargeInteger:
		return uint16(21)
	case *PropertyVariant_VarUnion_Bool:
		return uint16(11)
	case *PropertyVariant_VarUnion_UUID:
		return uint16(72)
	case *PropertyVariant_VarUnion_Blob:
		return uint16(65)
	case *PropertyVariant_VarUnion_UnicodeString:
		return uint16(31)
	case *PropertyVariant_VarUnion_ByteArray:
		return uint16(4113)
	case *PropertyVariant_VarUnion_UshortArray:
		return uint16(4114)
	case *PropertyVariant_VarUnion_LongArray:
		return uint16(4099)
	case *PropertyVariant_VarUnion_Uint32Array:
		return uint16(4115)
	case *PropertyVariant_VarUnion_UlargeIntegerArray:
		return uint16(4117)
	case *PropertyVariant_VarUnion_UUIDArray:
		return uint16(4168)
	case *PropertyVariant_VarUnion_UnicodeStringArray:
		return uint16(4127)
	case *PropertyVariant_VarUnion_PropertyVariantArray:
		return uint16(4108)
	}
	return uint16(0)
}

func (o *PropertyVariant_VarUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0),
		uint16(1):
	case uint16(16):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Char)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Char{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(17):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Byte)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Byte{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Short)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Short{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(18):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Ushort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Ushort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Long)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Long{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(19):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Ulong)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Ulong{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(20):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_LargeInteger)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_LargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(21):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UlargeInteger)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(11):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Bool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Bool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(72):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UUID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(65):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Blob)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Blob{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(31):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UnicodeString)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4113):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_ByteArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_ByteArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4114):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UshortArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UshortArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4099):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_LongArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_LongArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4115):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_Uint32Array)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4117):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UlargeIntegerArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UlargeIntegerArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4168):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UUIDArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UUIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4127):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_UnicodeStringArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_UnicodeStringArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(4108):
		_o, _ := o.Value.(*PropertyVariant_VarUnion_PropertyVariantArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PropertyVariant_VarUnion_PropertyVariantArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *PropertyVariant_VarUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(0),
		uint16(1):
		o.Value = &PropertyVariant_VarUnion_0{}
	case uint16(16):
		o.Value = &PropertyVariant_VarUnion_Char{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(17):
		o.Value = &PropertyVariant_VarUnion_Byte{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2):
		o.Value = &PropertyVariant_VarUnion_Short{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(18):
		o.Value = &PropertyVariant_VarUnion_Ushort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &PropertyVariant_VarUnion_Long{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(19):
		o.Value = &PropertyVariant_VarUnion_Ulong{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(20):
		o.Value = &PropertyVariant_VarUnion_LargeInteger{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(21):
		o.Value = &PropertyVariant_VarUnion_UlargeInteger{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(11):
		o.Value = &PropertyVariant_VarUnion_Bool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(72):
		o.Value = &PropertyVariant_VarUnion_UUID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(65):
		o.Value = &PropertyVariant_VarUnion_Blob{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(31):
		o.Value = &PropertyVariant_VarUnion_UnicodeString{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4113):
		o.Value = &PropertyVariant_VarUnion_ByteArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4114):
		o.Value = &PropertyVariant_VarUnion_UshortArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4099):
		o.Value = &PropertyVariant_VarUnion_LongArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4115):
		o.Value = &PropertyVariant_VarUnion_Uint32Array{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4117):
		o.Value = &PropertyVariant_VarUnion_UlargeIntegerArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4168):
		o.Value = &PropertyVariant_VarUnion_UUIDArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4127):
		o.Value = &PropertyVariant_VarUnion_UnicodeStringArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(4108):
		o.Value = &PropertyVariant_VarUnion_PropertyVariantArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// PropertyVariant_VarUnion_0 structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 0, 1
type PropertyVariant_VarUnion_0 struct {
}

func (*PropertyVariant_VarUnion_0) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *PropertyVariant_VarUnion_0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// PropertyVariant_VarUnion_Char structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 16
type PropertyVariant_VarUnion_Char struct {
	Char uint8 `idl:"name:cVal" json:"char"`
}

func (*PropertyVariant_VarUnion_Char) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Char) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Char); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Char) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Char); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Byte structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 17
type PropertyVariant_VarUnion_Byte struct {
	Byte uint8 `idl:"name:bVal" json:"byte"`
}

func (*PropertyVariant_VarUnion_Byte) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Byte) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Byte); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Byte) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Byte); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Short structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 2
type PropertyVariant_VarUnion_Short struct {
	Short int16 `idl:"name:iVal" json:"short"`
}

func (*PropertyVariant_VarUnion_Short) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Short) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Short); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Short) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Short); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Ushort structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 18
type PropertyVariant_VarUnion_Ushort struct {
	Ushort uint16 `idl:"name:uiVal" json:"ushort"`
}

func (*PropertyVariant_VarUnion_Ushort) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Ushort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Ushort); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Ushort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Ushort); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Long structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 3
type PropertyVariant_VarUnion_Long struct {
	Long int32 `idl:"name:lVal" json:"long"`
}

func (*PropertyVariant_VarUnion_Long) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Long) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Long); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Long) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Long); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Ulong structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 19
type PropertyVariant_VarUnion_Ulong struct {
	Ulong uint32 `idl:"name:ulVal" json:"ulong"`
}

func (*PropertyVariant_VarUnion_Ulong) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Ulong) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Ulong); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Ulong) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Ulong); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_LargeInteger structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 20
type PropertyVariant_VarUnion_LargeInteger struct {
	LargeInteger *dtyp.LargeInteger `idl:"name:hVal" json:"large_integer"`
}

func (*PropertyVariant_VarUnion_LargeInteger) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_LargeInteger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LargeInteger != nil {
		if err := o.LargeInteger.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_LargeInteger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LargeInteger == nil {
		o.LargeInteger = &dtyp.LargeInteger{}
	}
	if err := o.LargeInteger.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UlargeInteger structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 21
type PropertyVariant_VarUnion_UlargeInteger struct {
	UlargeInteger *dtyp.UlargeInteger `idl:"name:uhVal" json:"ularge_integer"`
}

func (*PropertyVariant_VarUnion_UlargeInteger) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UlargeInteger) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UlargeInteger != nil {
		if err := o.UlargeInteger.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UlargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UlargeInteger) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UlargeInteger == nil {
		o.UlargeInteger = &dtyp.UlargeInteger{}
	}
	if err := o.UlargeInteger.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Bool structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 11
type PropertyVariant_VarUnion_Bool struct {
	Bool int16 `idl:"name:boolVal" json:"bool"`
}

func (*PropertyVariant_VarUnion_Bool) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Bool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Bool); err != nil {
		return err
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Bool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Bool); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UUID structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 72
type PropertyVariant_VarUnion_UUID struct {
	UUID *dtyp.GUID `idl:"name:puuid" json:"uuid"`
}

func (*PropertyVariant_VarUnion_UUID) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UUID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UUID != nil {
		_ptr_puuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
		})
		if err := w.WritePointer(&o.UUID, _ptr_puuid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UUID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_puuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UUID == nil {
			o.UUID = &dtyp.GUID{}
		}
		if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_puuid := func(ptr interface{}) { o.UUID = *ptr.(**dtyp.GUID) }
	if err := w.ReadPointer(&o.UUID, _s_puuid, _ptr_puuid); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Blob structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 65
type PropertyVariant_VarUnion_Blob struct {
	Blob *Blob `idl:"name:blob" json:"blob"`
}

func (*PropertyVariant_VarUnion_Blob) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Blob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Blob != nil {
		if err := o.Blob.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Blob{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Blob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Blob == nil {
		o.Blob = &Blob{}
	}
	if err := o.Blob.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UnicodeString structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 31
type PropertyVariant_VarUnion_UnicodeString struct {
	UnicodeString string `idl:"name:pwszVal;string" json:"unicode_string"`
}

func (*PropertyVariant_VarUnion_UnicodeString) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeString != "" {
		_ptr_pwszVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UnicodeString); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UnicodeString, _ptr_pwszVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pwszVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UnicodeString); err != nil {
			return err
		}
		return nil
	})
	_s_pwszVal := func(ptr interface{}) { o.UnicodeString = *ptr.(*string) }
	if err := w.ReadPointer(&o.UnicodeString, _s_pwszVal, _ptr_pwszVal); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_ByteArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4113
type PropertyVariant_VarUnion_ByteArray struct {
	ByteArray *ByteArray `idl:"name:caub" json:"byte_array"`
}

func (*PropertyVariant_VarUnion_ByteArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_ByteArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ByteArray != nil {
		if err := o.ByteArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_ByteArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ByteArray == nil {
		o.ByteArray = &ByteArray{}
	}
	if err := o.ByteArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UshortArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4114
type PropertyVariant_VarUnion_UshortArray struct {
	UshortArray *UshortArray `idl:"name:caui" json:"ushort_array"`
}

func (*PropertyVariant_VarUnion_UshortArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UshortArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UshortArray != nil {
		if err := o.UshortArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UshortArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UshortArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UshortArray == nil {
		o.UshortArray = &UshortArray{}
	}
	if err := o.UshortArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_LongArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4099
type PropertyVariant_VarUnion_LongArray struct {
	LongArray *LongArray `idl:"name:cal" json:"long_array"`
}

func (*PropertyVariant_VarUnion_LongArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_LongArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.LongArray != nil {
		if err := o.LongArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&LongArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_LongArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.LongArray == nil {
		o.LongArray = &LongArray{}
	}
	if err := o.LongArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_Uint32Array structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4115
type PropertyVariant_VarUnion_Uint32Array struct {
	Uint32Array *Uint32Array `idl:"name:caul" json:"uint32_array"`
}

func (*PropertyVariant_VarUnion_Uint32Array) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_Uint32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Uint32Array != nil {
		if err := o.Uint32Array.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_Uint32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Uint32Array == nil {
		o.Uint32Array = &Uint32Array{}
	}
	if err := o.Uint32Array.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UlargeIntegerArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4117
type PropertyVariant_VarUnion_UlargeIntegerArray struct {
	UlargeIntegerArray *UlargeIntegerArray `idl:"name:cauh" json:"ularge_integer_array"`
}

func (*PropertyVariant_VarUnion_UlargeIntegerArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UlargeIntegerArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UlargeIntegerArray != nil {
		if err := o.UlargeIntegerArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UlargeIntegerArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UlargeIntegerArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UlargeIntegerArray == nil {
		o.UlargeIntegerArray = &UlargeIntegerArray{}
	}
	if err := o.UlargeIntegerArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UUIDArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4168
type PropertyVariant_VarUnion_UUIDArray struct {
	UUIDArray *ClassIDArray `idl:"name:cauuid" json:"uuid_array"`
}

func (*PropertyVariant_VarUnion_UUIDArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UUIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UUIDArray != nil {
		if err := o.UUIDArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ClassIDArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UUIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UUIDArray == nil {
		o.UUIDArray = &ClassIDArray{}
	}
	if err := o.UUIDArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_UnicodeStringArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4127
type PropertyVariant_VarUnion_UnicodeStringArray struct {
	UnicodeStringArray *UnicodeStringArray `idl:"name:calpwstr" json:"unicode_string_array"`
}

func (*PropertyVariant_VarUnion_UnicodeStringArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_UnicodeStringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UnicodeStringArray != nil {
		if err := o.UnicodeStringArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UnicodeStringArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_UnicodeStringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.UnicodeStringArray == nil {
		o.UnicodeStringArray = &UnicodeStringArray{}
	}
	if err := o.UnicodeStringArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// PropertyVariant_VarUnion_PropertyVariantArray structure represents PropertyVariant_VarUnion RPC union arm.
//
// It has following labels: 4108
type PropertyVariant_VarUnion_PropertyVariantArray struct {
	PropertyVariantArray *PropertyVariantArray `idl:"name:capropvar" json:"property_variant_array"`
}

func (*PropertyVariant_VarUnion_PropertyVariantArray) is_PropertyVariant_VarUnion() {}

func (o *PropertyVariant_VarUnion_PropertyVariantArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PropertyVariantArray != nil {
		if err := o.PropertyVariantArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&PropertyVariantArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PropertyVariant_VarUnion_PropertyVariantArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PropertyVariantArray == nil {
		o.PropertyVariantArray = &PropertyVariantArray{}
	}
	if err := o.PropertyVariantArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// DLID structure represents DL_ID RPC structure.
//
// The DL_ID structure defines a distribution list queue identifier.
type DLID struct {
	// m_DlGuid:  The GUID (as specified in [MS-DTYP] section 2.3.4) of the distribution
	// list queue.
	DLGUID *dtyp.GUID `idl:"name:m_DlGuid" json:"dl_guid"`
	// m_pwzDomain:  The Active Directory domain of the distribution list queue. This field
	// MUST be a null-terminated Unicode string.
	Domain string `idl:"name:m_pwzDomain;string" json:"domain"`
}

func (o *DLID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DLID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.DLGUID != nil {
		if err := o.DLGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Domain != "" {
		_ptr_m_pwzDomain := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Domain); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Domain, _ptr_m_pwzDomain); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DLID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.DLGUID == nil {
		o.DLGUID = &dtyp.GUID{}
	}
	if err := o.DLGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_m_pwzDomain := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Domain); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwzDomain := func(ptr interface{}) { o.Domain = *ptr.(*string) }
	if err := w.ReadPointer(&o.Domain, _s_m_pwzDomain, _ptr_m_pwzDomain); err != nil {
		return err
	}
	return nil
}

// MulticastID structure represents MULTICAST_ID RPC structure.
//
// The MULTICAST_ID structure defines a multicast queue identifier.
type MulticastID struct {
	// m_address:  The IP address of the queue.
	Address uint32 `idl:"name:m_address" json:"address"`
	// m_port:  The port to which the queue is attached.
	Port uint32 `idl:"name:m_port" json:"port"`
}

func (o *MulticastID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *MulticastID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Address); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	return nil
}
func (o *MulticastID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Address); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	return nil
}

// ObjectID structure represents OBJECTID RPC structure.
//
// The OBJECTID structure is used to uniquely distinguish objects of the same type within
// the message queuing system. The structure has two identifiers: a group identifier
// and an object identifier.
type ObjectID struct {
	// Lineage:  A GUID (as specified in [MS-DTYP] section 2.3.4) value that identifies
	// the group to which an object belongs. A group is a protocol-specific concept. For
	// instance, it can be the identifier of the object owner, or it can be the identifier
	// of the source where the objects originate.
	Lineage *dtyp.GUID `idl:"name:Lineage" json:"lineage"`
	// Uniquifier:  A DWORD value that identifies the object within the group.
	Uniquifier uint32 `idl:"name:Uniquifier" json:"uniquifier"`
}

func (o *ObjectID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
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
	if o.Lineage != nil {
		if err := o.Lineage.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Uniquifier); err != nil {
		return err
	}
	return nil
}
func (o *ObjectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.Lineage == nil {
		o.Lineage = &dtyp.GUID{}
	}
	if err := o.Lineage.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Uniquifier); err != nil {
		return err
	}
	return nil
}

// QueueFormatType type represents QUEUE_FORMAT_TYPE RPC enumeration.
//
// The QUEUE_FORMAT_TYPE enumeration identifies the type of name format being used.
//
// QUEUE_FORMAT_TYPE_UNKNOWN:  The format type is unknown.
type QueueFormatType uint16

var (
	QueueFormatTypeUnknown QueueFormatType = 0
	// QUEUE_FORMAT_TYPE_PUBLIC:  The QUEUE_FORMAT (section 2.2.7) structure contains a
	// GUID (as specified in [MS-DTYP] section 2.3.4) that identifies a queue.
	QueueFormatTypePublic QueueFormatType = 1
	// QUEUE_FORMAT_TYPE_PRIVATE:  The QUEUE_FORMAT (section 2.2.7) structure contains
	// an OBJECTID structure that identifies a queue.
	QueueFormatTypePrivate QueueFormatType = 2
	// QUEUE_FORMAT_TYPE_DIRECT:  The QUEUE_FORMAT structure contains a direct format name
	// string that identifies a queue.
	QueueFormatTypeDirect QueueFormatType = 3
	// QUEUE_FORMAT_TYPE_MACHINE:  The QUEUE_FORMAT structure contains a GUID (as specified
	// in [MS-DTYP] section 2.3.4) that identifies a queue.
	QueueFormatTypeMachine QueueFormatType = 4
	// QUEUE_FORMAT_TYPE_CONNECTOR:  The QUEUE_FORMAT structure contains a GUID (as specified
	// in [MS-DTYP] section 2.3.4) that identifies a connector queue. This is not supported
	// by all protocols.
	QueueFormatTypeConnector QueueFormatType = 5
	// QUEUE_FORMAT_TYPE_DL:  The QUEUE_FORMAT structure contains a GUID (as specified
	// in [MS-DTYP] section 2.3.4) that identifies a distribution list (DL). This is not
	// supported by all protocols.
	QueueFormatTypeDL QueueFormatType = 6
	// QUEUE_FORMAT_TYPE_MULTICAST:  The QUEUE_FORMAT structure contains a MULTICAST_ID
	// (section 2.2.10) that identifies a multicast address. This is not supported by all
	// protocols.
	QueueFormatTypeMulticast QueueFormatType = 7
	// QUEUE_FORMAT_TYPE_SUBQUEUE:  The QUEUE_FORMAT structure contains a direct name string
	// that identifies a subqueue.
	QueueFormatTypeSubqueue QueueFormatType = 8
)

func (o QueueFormatType) String() string {
	switch o {
	case QueueFormatTypeUnknown:
		return "QueueFormatTypeUnknown"
	case QueueFormatTypePublic:
		return "QueueFormatTypePublic"
	case QueueFormatTypePrivate:
		return "QueueFormatTypePrivate"
	case QueueFormatTypeDirect:
		return "QueueFormatTypeDirect"
	case QueueFormatTypeMachine:
		return "QueueFormatTypeMachine"
	case QueueFormatTypeConnector:
		return "QueueFormatTypeConnector"
	case QueueFormatTypeDL:
		return "QueueFormatTypeDL"
	case QueueFormatTypeMulticast:
		return "QueueFormatTypeMulticast"
	case QueueFormatTypeSubqueue:
		return "QueueFormatTypeSubqueue"
	}
	return "Invalid"
}

// QueueFormat structure represents QUEUE_FORMAT RPC structure.
//
// The QUEUE_FORMAT structure describes the type of queue being managed and an identifier
// for that queue.
type QueueFormat struct {
	// m_qft:  The type of queue format name. It MUST be set to one of the values of QUEUE_FORMAT_TYPE.
	// It is used as a union discriminant in the QUEUE_FORMAT structure.
	QueueFormatType uint8 `idl:"name:m_qft" json:"queue_format_type"`
	// m_SuffixAndFlags:  This member is broken into two subfields: Suffix Type is located
	// in the 4 least-significant bits, and Flags is located in the 4 most-significant bits.
	//
	//	+---+---+---+---+---+---+---+---+
	//	|   |   |   |   |   |   |   |   |
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
	//	|   |   |   |   |   |   |   |   |
	//	+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+
	//	| Flags         | Suffix type   |
	//	+---+---+---+---+---+---+---+---+
	//
	//
	//	+-----------------------------------+--------------------------------------------+
	//	|                                   |                                            |
	//	|               FLAGS               |                  MEANING                   |
	//	|                                   |                                            |
	//	+-----------------------------------+--------------------------------------------+
	//	+-----------------------------------+--------------------------------------------+
	//	| QUEUE_FORMAT_FLAG_NOT_SYSTEM 0x00 | The specified queue is not a system queue. |
	//	+-----------------------------------+--------------------------------------------+
	//	| QUEUE_FORMAT_FLAG_SYSTEM 0x80     | The specified queue is a system queue.     |
	//	+-----------------------------------+--------------------------------------------+
	//
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              SUFFIX               |                                                                                  |
	//	|               TYPE                |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_NONE 0x00       | No suffix is specified. The Flags subfield MUST be set to 0x00. The m_qft member |
	//	|                                   | MUST NOT be set to 0x04.                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_JOURNAL 0x01    | A journal suffix. The Flags subfield MUST be set to 0x80. The m_qft member MUST  |
	//	|                                   | NOT be set to 0x05, 0x06, or 0x07.                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_DEADLETTER 0x02 | A dead-letter suffix. The Flags subfield MUST be set to 0x80. The m_qft member   |
	//	|                                   | MUST NOT be set to 0x01, 0x02, 0x05, 0x06, or 0x07.                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_DEADXACT 0x03   | A transacted dead-letter suffix. The Flags subfield MUST be set to 0x80. The     |
	//	|                                   | m_qft member MUST be set to 0x03 or 0x04.                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_XACTONLY 0x04   | A transaction-only suffix. The m_qft member MUST be set to 0x05.                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| QUEUE_SUFFIX_TYPE_SUBQUEUE 0x05   | A subqueue suffix. The Flags subfield MUST be 0x00. The m_qft member MUST be set |
	//	|                                   | to 0x08.                                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	SuffixAndFlags uint8 `idl:"name:m_SuffixAndFlags" json:"suffix_and_flags"`
	// m_reserved:  The integer value used for padding. The client SHOULD set this value
	// to 0. The server MUST not use it.
	//
	// (unnamed union):  Based on the value of m_qft.
	_           uint16                   `idl:"name:m_reserved"`
	QueueFormat *QueueFormat_QueueFormat `idl:"name:QueueFormat;switch_is:m_qft" json:"queue_format"`
}

func (o *QueueFormat) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.QueueFormatType); err != nil {
		return err
	}
	if err := w.WriteData(o.SuffixAndFlags); err != nil {
		return err
	}
	// reserved m_reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	_swQueueFormat := uint8(o.QueueFormatType)
	if o.QueueFormat != nil {
		if err := o.QueueFormat.MarshalUnionNDR(ctx, w, _swQueueFormat); err != nil {
			return err
		}
	} else {
		if err := (&QueueFormat_QueueFormat{}).MarshalUnionNDR(ctx, w, _swQueueFormat); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.QueueFormatType); err != nil {
		return err
	}
	if err := w.ReadData(&o.SuffixAndFlags); err != nil {
		return err
	}
	// reserved m_reserved
	var _m_reserved uint16
	if err := w.ReadData(&_m_reserved); err != nil {
		return err
	}
	if o.QueueFormat == nil {
		o.QueueFormat = &QueueFormat_QueueFormat{}
	}
	_swQueueFormat := uint8(o.QueueFormatType)
	if err := o.QueueFormat.UnmarshalUnionNDR(ctx, w, _swQueueFormat); err != nil {
		return err
	}
	return nil
}

// QueueFormat_QueueFormat structure represents QUEUE_FORMAT union anonymous member.
//
// The QUEUE_FORMAT structure describes the type of queue being managed and an identifier
// for that queue.
type QueueFormat_QueueFormat struct {
	// Types that are assignable to Value
	//
	// *QueueFormat_0
	// *QueueFormat_PublicID
	// *QueueFormat_PrivateID
	// *QueueFormat_DirectID
	// *QueueFormat_MachineID
	// *QueueFormat_ConnectorID
	// *QueueFormat_DLID
	// *QueueFormat_MulticastID
	// *QueueFormat_DirectSubqueueID
	Value is_QueueFormat_QueueFormat `json:"value"`
}

func (o *QueueFormat_QueueFormat) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *QueueFormat_PublicID:
		if value != nil {
			return value.PublicID
		}
	case *QueueFormat_PrivateID:
		if value != nil {
			return value.PrivateID
		}
	case *QueueFormat_DirectID:
		if value != nil {
			return value.DirectID
		}
	case *QueueFormat_MachineID:
		if value != nil {
			return value.MachineID
		}
	case *QueueFormat_ConnectorID:
		if value != nil {
			return value.ConnectorID
		}
	case *QueueFormat_DLID:
		if value != nil {
			return value.DLID
		}
	case *QueueFormat_MulticastID:
		if value != nil {
			return value.MulticastID
		}
	case *QueueFormat_DirectSubqueueID:
		if value != nil {
			return value.DirectSubqueueID
		}
	}
	return nil
}

type is_QueueFormat_QueueFormat interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_QueueFormat_QueueFormat()
}

func (o *QueueFormat_QueueFormat) NDRSwitchValue(sw uint8) uint8 {
	if o == nil {
		return uint8(0)
	}
	switch (interface{})(o.Value).(type) {
	case *QueueFormat_0:
		return uint8(0)
	case *QueueFormat_PublicID:
		return uint8(1)
	case *QueueFormat_PrivateID:
		return uint8(2)
	case *QueueFormat_DirectID:
		return uint8(3)
	case *QueueFormat_MachineID:
		return uint8(4)
	case *QueueFormat_ConnectorID:
		return uint8(5)
	case *QueueFormat_DLID:
		return uint8(6)
	case *QueueFormat_MulticastID:
		return uint8(7)
	case *QueueFormat_DirectSubqueueID:
		return uint8(8)
	}
	return uint8(0)
}

func (o *QueueFormat_QueueFormat) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint8) error {
	if err := w.WriteSwitch(uint8(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint8(0):
	case uint8(1):
		_o, _ := o.Value.(*QueueFormat_PublicID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_PublicID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(2):
		_o, _ := o.Value.(*QueueFormat_PrivateID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_PrivateID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(3):
		_o, _ := o.Value.(*QueueFormat_DirectID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_DirectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(4):
		_o, _ := o.Value.(*QueueFormat_MachineID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_MachineID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(5):
		_o, _ := o.Value.(*QueueFormat_ConnectorID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_ConnectorID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(6):
		_o, _ := o.Value.(*QueueFormat_DLID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_DLID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(7):
		_o, _ := o.Value.(*QueueFormat_MulticastID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_MulticastID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint8(8):
		_o, _ := o.Value.(*QueueFormat_DirectSubqueueID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&QueueFormat_DirectSubqueueID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *QueueFormat_QueueFormat) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint8) error {
	if err := w.ReadSwitch((*uint8)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint8(0):
		o.Value = &QueueFormat_0{}
	case uint8(1):
		o.Value = &QueueFormat_PublicID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(2):
		o.Value = &QueueFormat_PrivateID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(3):
		o.Value = &QueueFormat_DirectID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(4):
		o.Value = &QueueFormat_MachineID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(5):
		o.Value = &QueueFormat_ConnectorID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(6):
		o.Value = &QueueFormat_DLID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(7):
		o.Value = &QueueFormat_MulticastID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint8(8):
		o.Value = &QueueFormat_DirectSubqueueID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// QueueFormat_0 structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 0
type QueueFormat_0 struct {
}

func (*QueueFormat_0) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *QueueFormat_0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// QueueFormat_PublicID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 1
type QueueFormat_PublicID struct {
	// m_gPublicID:  A GUID (as specified in [MS-DTYP] section 2.3.4) of a public queue.
	// Selected when m_qft is set to 0x01.
	PublicID *dtyp.GUID `idl:"name:m_gPublicID" json:"public_id"`
}

func (*QueueFormat_PublicID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_PublicID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PublicID != nil {
		if err := o.PublicID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_PublicID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PublicID == nil {
		o.PublicID = &dtyp.GUID{}
	}
	if err := o.PublicID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_PrivateID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 2
type QueueFormat_PrivateID struct {
	// m_oPrivateID:  An OBJECTID of a private queue; members MUST be used as specified
	// in OBJECTID. Selected when m_qft is set to 0x02.
	PrivateID *ObjectID `idl:"name:m_oPrivateID" json:"private_id"`
}

func (*QueueFormat_PrivateID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_PrivateID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.PrivateID != nil {
		if err := o.PrivateID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ObjectID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_PrivateID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.PrivateID == nil {
		o.PrivateID = &ObjectID{}
	}
	if err := o.PrivateID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_DirectID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 3
type QueueFormat_DirectID struct {
	// m_pDirectID:  A direct format name (as specified in section 2.1.2) with the "DIRECT="
	// prefix removed. It is selected when m_qft is set to 0x03.
	DirectID string `idl:"name:m_pDirectID;string" json:"direct_id"`
}

func (*QueueFormat_DirectID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_DirectID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DirectID != "" {
		_ptr_m_pDirectID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DirectID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DirectID, _ptr_m_pDirectID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_DirectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_m_pDirectID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DirectID); err != nil {
			return err
		}
		return nil
	})
	_s_m_pDirectID := func(ptr interface{}) { o.DirectID = *ptr.(*string) }
	if err := w.ReadPointer(&o.DirectID, _s_m_pDirectID, _ptr_m_pDirectID); err != nil {
		return err
	}
	return nil
}

// QueueFormat_MachineID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 4
type QueueFormat_MachineID struct {
	// m_gMachineID:  The GUID (as specified in [MS-DTYP] section 2.3.4) of a machine. It
	// is selected when m_qft is set to 0x04.
	MachineID *dtyp.GUID `idl:"name:m_gMachineID" json:"machine_id"`
}

func (*QueueFormat_MachineID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_MachineID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MachineID != nil {
		if err := o.MachineID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_MachineID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MachineID == nil {
		o.MachineID = &dtyp.GUID{}
	}
	if err := o.MachineID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_ConnectorID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 5
type QueueFormat_ConnectorID struct {
	// m_GConnectorID:  The GUID (as specified in [MS-DTYP] section 2.3.4) of a connector
	// queue. It is selected when m_qft is set to 0x05.
	ConnectorID *dtyp.GUID `idl:"name:m_GConnectorID" json:"connector_id"`
}

func (*QueueFormat_ConnectorID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_ConnectorID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ConnectorID != nil {
		if err := o.ConnectorID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_ConnectorID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.ConnectorID == nil {
		o.ConnectorID = &dtyp.GUID{}
	}
	if err := o.ConnectorID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_DLID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 6
type QueueFormat_DLID struct {
	// m_DlID:  The identifier of a distribution list. It is selected when m_qft is set
	// to 0x06.
	DLID *DLID `idl:"name:m_DlID" json:"dl_id"`
}

func (*QueueFormat_DLID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_DLID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DLID != nil {
		if err := o.DLID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DLID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_DLID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.DLID == nil {
		o.DLID = &DLID{}
	}
	if err := o.DLID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_MulticastID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 7
type QueueFormat_MulticastID struct {
	// m_MulticastID:  A MULTICAST_ID (section 2.2.10) which specifies a multicast address
	// and port. It is selected when m_qft is set to 0x07.
	MulticastID *MulticastID `idl:"name:m_MulticastID" json:"multicast_id"`
}

func (*QueueFormat_MulticastID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_MulticastID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.MulticastID != nil {
		if err := o.MulticastID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&MulticastID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_MulticastID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.MulticastID == nil {
		o.MulticastID = &MulticastID{}
	}
	if err := o.MulticastID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// QueueFormat_DirectSubqueueID structure represents QueueFormat_QueueFormat RPC union arm.
//
// It has following labels: 8
type QueueFormat_DirectSubqueueID struct {
	// m_pDirectSubqueueID:  The identifier of a subqueue. Selected when m_qft is set to
	// 0x08.
	//
	// The value MUST conform to the ABNF for DirectName and contain the optional <Subqueue>
	// element, as specified in section 2.1.
	DirectSubqueueID string `idl:"name:m_pDirectSubqueueID;string" json:"direct_subqueue_id"`
}

func (*QueueFormat_DirectSubqueueID) is_QueueFormat_QueueFormat() {}

func (o *QueueFormat_DirectSubqueueID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DirectSubqueueID != "" {
		_ptr_m_pDirectSubqueueID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DirectSubqueueID); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DirectSubqueueID, _ptr_m_pDirectSubqueueID); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueueFormat_DirectSubqueueID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_m_pDirectSubqueueID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DirectSubqueueID); err != nil {
			return err
		}
		return nil
	})
	_s_m_pDirectSubqueueID := func(ptr interface{}) { o.DirectSubqueueID = *ptr.(*string) }
	if err := w.ReadPointer(&o.DirectSubqueueID, _s_m_pDirectSubqueueID, _ptr_m_pDirectSubqueueID); err != nil {
		return err
	}
	return nil
}
