package idmnotify

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/dmrp"
)

var (
	// IDMNotify interface identifier d2d79df7-3400-11d0-b40b-00aa005ff586
	IDMNotifyIID = &dcom.IID{Data1: 0xd2d79df7, Data2: 0x3400, Data3: 0x11d0, Data4: []byte{0xb4, 0x0b, 0x00, 0xaa, 0x00, 0x5f, 0xf5, 0x86}}
	// Syntax UUID
	IDMNotifySyntaxUUID = &uuid.UUID{TimeLow: 0xd2d79df7, TimeMid: 0x3400, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x5f, 0xf5, 0x86}}
	// Syntax ID
	IDMNotifySyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IDMNotifySyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IDMNotify interface.
type IDMNotifyClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The ObjectsChanged method notifies the client of object changes.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF] section 2.1).
	ObjectsChanged(context.Context, *ObjectsChangedRequest, ...dcerpc.CallOption) (*ObjectsChangedResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IDMNotifyClient
}

type xxx_DefaultIDMNotifyClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIDMNotifyClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultIDMNotifyClient) ObjectsChanged(ctx context.Context, in *ObjectsChangedRequest, opts ...dcerpc.CallOption) (*ObjectsChangedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ObjectsChangedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIDMNotifyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIDMNotifyClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIDMNotifyClient) IPID(ctx context.Context, ipid *dcom.IPID) IDMNotifyClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIDMNotifyClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewIDMNotifyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IDMNotifyClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IDMNotifySyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultIDMNotifyClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ObjectsChangedOperation structure represents the ObjectsChanged operation
type xxx_ObjectsChangedOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ByteCount  uint32         `idl:"name:ByteCount" json:"byte_count"`
	ByteStream []byte         `idl:"name:ByteStream;size_is:(ByteCount)" json:"byte_stream"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ObjectsChangedOperation) OpNum() int { return 3 }

func (o *xxx_ObjectsChangedOperation) OpName() string { return "/IDMNotify/v0/ObjectsChanged" }

func (o *xxx_ObjectsChangedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ByteStream != nil && o.ByteCount == 0 {
		o.ByteCount = uint32(len(o.ByteStream))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectsChangedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ByteCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ByteCount); err != nil {
			return err
		}
	}
	// ByteStream {in} (1:{pointer=ref}*(1)[dim:0,size_is=ByteCount](uint8))
	{
		dimSize1 := uint64(o.ByteCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ByteStream {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.ByteStream[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.ByteStream); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ObjectsChangedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ByteCount {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ByteCount); err != nil {
			return err
		}
	}
	// ByteStream {in} (1:{pointer=ref}*(1)[dim:0,size_is=ByteCount](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ByteStream", sizeInfo[0])
		}
		o.ByteStream = make([]byte, sizeInfo[0])
		for i1 := range o.ByteStream {
			i1 := i1
			if err := w.ReadData(&o.ByteStream[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ObjectsChangedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectsChangedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ObjectsChangedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ObjectsChangedRequest structure represents the ObjectsChanged operation request
type ObjectsChangedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ByteCount: Length of ByteStream in bytes.
	ByteCount uint32 `idl:"name:ByteCount" json:"byte_count"`
	// ByteStream: Array of bytes that compose any number of variable-length change notification
	// structures. Memory for the array is allocated and freed by the caller (that is, the
	// server).
	//
	// Any variable-length change notification structure in the array starts with a fixed
	// header that contains the fields shown in the following table.
	//
	//	+------------+--------------------+-----------------------------------------------+
	//	|   FIELD    |        DATA        |                                               |
	//	|    NAME    |        TYPE        |                  DESCRIPTION                  |
	//	|            |                    |                                               |
	//	+------------+--------------------+-----------------------------------------------+
	//	+------------+--------------------+-----------------------------------------------+
	//	| size       | ULONG              | The total size of the structure in bytes.     |
	//	+------------+--------------------+-----------------------------------------------+
	//	| type       | DMNOTIFY_INFO_TYPE | The type of object that changed.              |
	//	+------------+--------------------+-----------------------------------------------+
	//	| action     | LDMACTION          | The type of change that the object underwent. |
	//	+------------+--------------------+-----------------------------------------------+
	//
	// Depending on the value of type, the fixed header of the notification structure is
	// followed by one of the following items.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                          STRUCTURE FOLLOWING THE FIXED                           |
	//	|         TYPE         |                                      HEADER                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_VOLUME_INFO | VOLUME_INFO                                                                      |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_TASK_INFO   | TASK_INFO                                                                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_DL_INFO     | DRIVE_LETTER_INFO                                                                |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_FS_INFO     | FILE_SYSTEM_INFO                                                                 |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_SYSTEM_INFO | ULONG                                                                            |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_DISK_INFO   | If client called Initialize on IVolumeClient interface, then DISK_INFO. If       |
	//	|                      | client called Initialize on IVolumeClient3 interface, then DISK_INFO_EX.         |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| DMNOTIFY_REGION_INFO | If client called Initialize on IVolumeClient interface, then REGION_INFO. If     |
	//	|                      | client called Initialize on IVolumeClient3 interface, then REGION_INFO_EX.       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//
	// Note  The structures that are transmitted within ByteStream are not marshaled in
	// RPC Network Data Representation (NDR) format. They are C structures, and the memory
	// layout and field types are those found on the Windows/Intel 32-bit and 64-bit architectures,
	// and, Windows/AMD 64-bit architecture. These structures are not packed, and padding
	// bytes can exist between successive structure fields to ensure that the field of a
	// given data type begins at a byte offset that is an integer multiple of the type's
	// size with respect to the beginning of the structure. The structures transmitted within
	// ByteStream also appear in other interfaces as RPC-marshaled structures. In these
	// interfaces, the structure fields will be marshaled in NDR format.
	//
	// *
	//
	// The structure is copied one byte at a time from memory into ByteStream beginning
	// at first byte after action field. If the structure contains character pointer fields,
	// those fields are omitted.
	//
	// *
	//
	// The character strings of the character pointer fields are copied into ByteStream
	// following the structure in the order in which they appear in the structure. All strings
	// are null-terminated. There is no padding between the end of the structure and the
	// first string, or between successive strings.
	//
	// At the client, the following technique is used to parse the byte stream back into
	// the appropriate structures:
	//
	// *
	//
	// The notification size, type, and action are parsed from the byte stream.
	//
	// *
	//
	// The notification object structure, up through the first string field, is copied out
	// of the byte stream and into the appropriate structure. For the IVolumeClient interface,
	// the disk ( 3c91641b-e5cd-409e-a445-2e4d61ae33c6#gt_c4133b2a-a990-4042-ba44-7fda3090f118
	// ) and region structures are DISK_INFO and REGION_INFO; for the IVolumeClient3 interface,
	// the structures are DISK_INFO_EX and REGION_INFO_EX. The client's ObjectsChanged implementation
	// MUST switch based on which version of the IVolumeClient interface is being used.
	// The client MUST also determine the type of processor architecture for both the server
	// and client. If the architectures are the same, the padding in the client-defined
	// structures will match that used in the server's byte stream. If the architectures
	// are not the same, the client MUST use the proper method for parsing the byte stream,
	// taking into account padding that MAY have been added for alignment purposes on either
	// the client or on the server. For more information, see section 8 ( 562d58df-633f-4afa-89ec-80917a2852b3
	// ).
	//
	// Allocations are done on the client to hold the character strings of the character
	// pointer fields. These fields are copied from ByteStream to the client-allocated buffers,
	// and appropriate structure fields are set to point to the client-allocated buffers.
	// All strings are null-terminated.
	ByteStream []byte `idl:"name:ByteStream;size_is:(ByteCount)" json:"byte_stream"`
}

func (o *ObjectsChangedRequest) xxx_ToOp(ctx context.Context, op *xxx_ObjectsChangedOperation) *xxx_ObjectsChangedOperation {
	if op == nil {
		op = &xxx_ObjectsChangedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ByteCount = o.ByteCount
	op.ByteStream = o.ByteStream
	return op
}

func (o *ObjectsChangedRequest) xxx_FromOp(ctx context.Context, op *xxx_ObjectsChangedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ByteCount = op.ByteCount
	o.ByteStream = op.ByteStream
}
func (o *ObjectsChangedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ObjectsChangedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ObjectsChangedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ObjectsChangedResponse structure represents the ObjectsChanged operation response
type ObjectsChangedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ObjectsChanged return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ObjectsChangedResponse) xxx_ToOp(ctx context.Context, op *xxx_ObjectsChangedOperation) *xxx_ObjectsChangedOperation {
	if op == nil {
		op = &xxx_ObjectsChangedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ObjectsChangedResponse) xxx_FromOp(ctx context.Context, op *xxx_ObjectsChangedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ObjectsChangedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ObjectsChangedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ObjectsChangedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
