package ocspa

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/ocspa"
)

// CertTransportBlob structure represents CERTTRANSBLOB RPC structure.
type CertTransportBlob struct {
	Length uint32 `idl:"name:cb" json:"length"`
	Buffer []byte `idl:"name:pb;size_is:(cb);pointer:unique" json:"buffer"`
}

func (o *CertTransportBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransportBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Buffer != nil || o.Length > 0 {
		_ptr_pb := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_pb); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransportBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pb := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pb := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_pb, _ptr_pb); err != nil {
		return err
	}
	return nil
}

// OCSPAdminD structure represents IOCSPAdminD RPC structure.
//
// The IOCSPAdminD interface provides an RPC interface for a client to manage an Online
// Responder.
//
// IOCSPAdminD interface inherits the IUnknown interface.
//
// The version number for IUnknown is 1.0. The UUID for the IOCSPAdminD interface is:
// "784b693d-95f3-420b-8126-365c098659f2". Method opnum field values start with 3; opnum
// values 0 through 2 represent the IUnknown methods: QueryInterface, AddRef, and Release
// methods inherited from [MS-DCOM].
//
// Each method MUST NOT throw exceptions.
//
// Methods in RPC Opnum Order
//
//	+------------------------+-------------+
//	|                        |             |
//	|         METHOD         | DESCRIPTION |
//	|                        |             |
//	+------------------------+-------------+
//	+------------------------+-------------+
//	| GetOCSPProperty        | Opnum: 3    |
//	+------------------------+-------------+
//	| SetOCSPProperty        | Opnum: 4    |
//	+------------------------+-------------+
//	| GetCAConfigInformation | Opnum: 5    |
//	+------------------------+-------------+
//	| SetCAConfigInformation | Opnum: 6    |
//	+------------------------+-------------+
//	| GetSecurity            | Opnum: 7    |
//	+------------------------+-------------+
//	| SetSecurity            | Opnum: 8    |
//	+------------------------+-------------+
//	| GetSigningCertificates | Opnum: 9    |
//	+------------------------+-------------+
//	| GetHashAlgorithms      | Opnum: 10   |
//	+------------------------+-------------+
//	| GetMyRoles             | Opnum: 11   |
//	+------------------------+-------------+
//	| Ping                   | Opnum: 12   |
//	+------------------------+-------------+
type OCSPAdminD dcom.InterfacePointer

func (o *OCSPAdminD) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *OCSPAdminD) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OCSPAdminD) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *OCSPAdminD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OCSPAdminD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
