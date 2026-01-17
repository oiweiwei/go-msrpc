package qmmgmt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	mqmq "github.com/oiweiwei/go-msrpc/msrpc/mqmq"
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
	_ = mqmq.GoPackage
)

var (
	// import guard
	GoPackage = "mqmr"
)

var (
	// Syntax UUID
	QmmgmtSyntaxUUID = &uuid.UUID{TimeLow: 0x41208ee0, TimeMid: 0xe970, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x9e, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x6, 0x4c, 0x39}}
	// Syntax ID
	QmmgmtSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: QmmgmtSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// qmmgmt interface.
type QmmgmtClient interface {

	// The R_QMMgmtGetInfo method requests information on an MSMQ installation on a server
	// or on a specific queue.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	//	+---------------------+----------------------------------------------------------------------------------+
	//	|       RETURN        |                                                                                  |
	//	|     VALUE/CODE      |                                   DESCRIPTION                                    |
	//	|                     |                                                                                  |
	//	+---------------------+----------------------------------------------------------------------------------+
	//	+---------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 MQ_OK    |                                                                                  |
	//	+---------------------+----------------------------------------------------------------------------------+
	//	| 0xC00E0001 MQ_ERROR | Generic error code. This error code is also the first of several error codes     |
	//	|                     | beginning with the string "MQ_ERR". A list of the errors prefaced with "MQ-ERR"  |
	//	|                     | is specified in 2.4.                                                             |
	//	+---------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// If an error occurs, the server MUST return a failure HRESULT and MUST NOT set any
	// [out] parameter values.
	//
	// The opnum field value for this method MUST be 0 and is received at a dynamically
	// assigned endpoint supplied by the RPC endpoint mapper, as specified in [MS-RPCE].
	//
	// If the pObjectFormat parameter specifies an MgmtObjectType of MGMT_MACHINE, the server
	// MUST return only those properties that pertain to the MSMQ installation. If pObjectFormat
	// specifies an MgmtObjectType of MGMT_QUEUE, the server MUST return only those properties
	// that pertain to a queue. If pObjectFormat specifies an MgmtObjectType of MGMT_SESSION,
	// the call MUST fail, and the error message MAY be MQ_ERROR_INVALID_PARAMETER (0xC00E0006).<15>
	//
	// If the pObjectFormat parameter specifies a computer, and one or more of the properties
	// specified in aProp are different than those specified in section 2.2.3.1, the call
	// MAY fail with MQ_ERROR_ILLEGAL_PROPID (0xC00E0039). If the pObjectFormat parameter
	// specifies a queue, and one or more of the properties specified in aProp are different
	// than those specified in section 2.2.3.2, the call MAY fail with MQ_ERROR_ILLEGAL_PROPID
	// (0xC00E0039).<16>
	//
	// MSMQ properties are specified in [MS-MQMQ] section 2.
	//
	// For MSMQ error codes, see [MSDN-MQEIC]. The structure and sequence of data on the
	// wire are specified in [C706] Transfer Syntax NDR.
	ManagementGetInfo(context.Context, *ManagementGetInfoRequest, ...dcerpc.CallOption) (*ManagementGetInfoResponse, error)

	// The R_QMMgmtAction method requests the server to perform a management function on
	// a specific queue or MSMQ installation.
	//
	// Return Values: On success, this method MUST return MQ_OK (0x00000000).
	//
	// MQ_OK (0x00000000)
	//
	// MQ_ERROR (0xC00E0001)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// If pObjectFormat specifies an MgmtObjectType of MGMT_SESSION or an lpwszAction has
	// different value than those in the table above, the call MUST fail and the error message
	// MAY be MQ_ERROR_INVALID_PARAMETER (0xC00E0006).<17>
	//
	// If an error occurs, the server MUST return a failure HRESULT.
	//
	// The opnum field value for this method MUST be 1 and is received at a dynamically
	// assigned endpoint supplied by the RPC endpoint mapper, as specified in [MS-RPCE].
	//
	// For MSMQ error codes, see [MSDN-MQEIC]. The structure and sequence of data on the
	// wire are specified in the Transfer Syntax NDR section in [C706].
	ManagementAction(context.Context, *ManagementActionRequest, ...dcerpc.CallOption) (*ManagementActionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ManagementObjectType type represents MgmtObjectType RPC enumeration.
//
// The MgmtObjectType enumeration identifies the type of management object (as specified
// in section 2.2.1.2) being used.
type ManagementObjectType uint16

var (
	// MGMT_MACHINE:  A machine management object.
	ManagementObjectTypeMachine ManagementObjectType = 1
	// MGMT_QUEUE:  A queue management object.
	ManagementObjectTypeQueue ManagementObjectType = 2
	// MGMT_SESSION:   A session management object.
	ManagementObjectTypeSession ManagementObjectType = 3
)

func (o ManagementObjectType) String() string {
	switch o {
	case ManagementObjectTypeMachine:
		return "ManagementObjectTypeMachine"
	case ManagementObjectTypeQueue:
		return "ManagementObjectTypeQueue"
	case ManagementObjectTypeSession:
		return "ManagementObjectTypeSession"
	}
	return "Invalid"
}

// ManagementObject structure represents MGMT_OBJECT RPC structure.
//
// The MGMT_OBJECT structure defines information on a queue, a computer, or a session.
// The structure includes an embedded discriminated union.
type ManagementObject struct {
	// type:  An integer discriminator for the embedded discriminated union. The value of
	// this field MUST be 1, 2, or 3, as specified in section 2.2.2.1.
	Type             ManagementObjectType               `idl:"name:type" json:"type"`
	ManagementObject *ManagementObject_ManagementObject `idl:"name:ManagementObject;switch_is:type" json:"management_object"`
}

func (o *ManagementObject) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ManagementObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Type)); err != nil {
		return err
	}
	_swManagementObject := uint16(o.Type)
	if o.ManagementObject != nil {
		if err := o.ManagementObject.MarshalUnionNDR(ctx, w, _swManagementObject); err != nil {
			return err
		}
	} else {
		if err := (&ManagementObject_ManagementObject{}).MarshalUnionNDR(ctx, w, _swManagementObject); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
		return err
	}
	if o.ManagementObject == nil {
		o.ManagementObject = &ManagementObject_ManagementObject{}
	}
	_swManagementObject := uint16(o.Type)
	if err := o.ManagementObject.UnmarshalUnionNDR(ctx, w, _swManagementObject); err != nil {
		return err
	}
	return nil
}

// ManagementObject_ManagementObject structure represents MGMT_OBJECT union anonymous member.
//
// The MGMT_OBJECT structure defines information on a queue, a computer, or a session.
// The structure includes an embedded discriminated union.
type ManagementObject_ManagementObject struct {
	// Types that are assignable to Value
	//
	// *ManagementObject_QueueFormat
	// *ManagementObject_Reserved1
	// *ManagementObject_Reserved2
	Value is_ManagementObject_ManagementObject `json:"value"`
}

func (o *ManagementObject_ManagementObject) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ManagementObject_QueueFormat:
		if value != nil {
			return value.QueueFormat
		}
	}
	return nil
}

type is_ManagementObject_ManagementObject interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ManagementObject_ManagementObject()
}

func (o *ManagementObject_ManagementObject) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ManagementObject_QueueFormat:
		return uint16(2)
	case *ManagementObject_Reserved1:
		return uint16(1)
	case *ManagementObject_Reserved2:
		return uint16(3)
	}
	return uint16(0)
}

func (o *ManagementObject_ManagementObject) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint16(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(2):
		_o, _ := o.Value.(*ManagementObject_QueueFormat)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObject_QueueFormat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(1):
		_o, _ := o.Value.(*ManagementObject_Reserved1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObject_Reserved1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*ManagementObject_Reserved2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObject_Reserved2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ManagementObject_ManagementObject) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint16)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(2):
		o.Value = &ManagementObject_QueueFormat{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(1):
		o.Value = &ManagementObject_Reserved1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &ManagementObject_Reserved2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ManagementObject_QueueFormat structure represents ManagementObject_ManagementObject RPC union arm.
//
// It has following labels: 2
type ManagementObject_QueueFormat struct {
	// pQueueFormat:  A pointer to a QUEUE_FORMAT ([MS-MQMQ] section 2.2.7) structure that
	// describes the type of the queue.
	QueueFormat *mqmq.QueueFormat `idl:"name:pQueueFormat" json:"queue_format"`
}

func (*ManagementObject_QueueFormat) is_ManagementObject_ManagementObject() {}

func (o *ManagementObject_QueueFormat) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.QueueFormat != nil {
		_ptr_pQueueFormat := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.QueueFormat != nil {
				if err := o.QueueFormat.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.QueueFormat{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.QueueFormat, _ptr_pQueueFormat); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ManagementObject_QueueFormat) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pQueueFormat := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.QueueFormat == nil {
			o.QueueFormat = &mqmq.QueueFormat{}
		}
		if err := o.QueueFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pQueueFormat := func(ptr interface{}) { o.QueueFormat = *ptr.(**mqmq.QueueFormat) }
	if err := w.ReadPointer(&o.QueueFormat, _s_pQueueFormat, _ptr_pQueueFormat); err != nil {
		return err
	}
	return nil
}

// ManagementObject_Reserved1 structure represents ManagementObject_ManagementObject RPC union arm.
//
// It has following labels: 1
type ManagementObject_Reserved1 struct {
	// Reserved1:  A 32-bit unsigned integer.<5>
	_ uint32 `idl:"name:Reserved1"`
}

func (*ManagementObject_Reserved1) is_ManagementObject_ManagementObject() {}

func (o *ManagementObject_Reserved1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	// reserved Reserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *ManagementObject_Reserved1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// reserved Reserved1
	var _Reserved1 uint32
	if err := w.ReadData(&_Reserved1); err != nil {
		return err
	}
	return nil
}

// ManagementObject_Reserved2 structure represents ManagementObject_ManagementObject RPC union arm.
//
// It has following labels: 3
type ManagementObject_Reserved2 struct {
	// Reserved2:  A 32-bit unsigned integer.<6>
	_ uint32 `idl:"name:Reserved2"`
}

func (*ManagementObject_Reserved2) is_ManagementObject_ManagementObject() {}

func (o *ManagementObject_Reserved2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	// reserved Reserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *ManagementObject_Reserved2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	// reserved Reserved2
	var _Reserved2 uint32
	if err := w.ReadData(&_Reserved2); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultQmmgmtClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultQmmgmtClient) ManagementGetInfo(ctx context.Context, in *ManagementGetInfoRequest, opts ...dcerpc.CallOption) (*ManagementGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ManagementGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmmgmtClient) ManagementAction(ctx context.Context, in *ManagementActionRequest, opts ...dcerpc.CallOption) (*ManagementActionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ManagementActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQmmgmtClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQmmgmtClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewQmmgmtClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QmmgmtClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QmmgmtSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultQmmgmtClient{cc: cc}, nil
}

// xxx_ManagementGetInfoOperation structure represents the R_QMMgmtGetInfo operation
type xxx_ManagementGetInfoOperation struct {
	ObjectFormat    *ManagementObject       `idl:"name:pObjectFormat" json:"object_format"`
	CreatePartition uint32                  `idl:"name:cp" json:"create_partition"`
	Property        []uint32                `idl:"name:aProp;size_is:(cp)" json:"property"`
	Var             []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	Return          int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ManagementGetInfoOperation) OpNum() int { return 0 }

func (o *xxx_ManagementGetInfoOperation) OpName() string { return "/qmmgmt/v1/R_QMMgmtGetInfo" }

func (o *xxx_ManagementGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Property != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Property))
	}
	if o.Var != nil && o.CreatePartition == 0 {
		o.CreatePartition = uint32(len(o.Var))
	}
	if o.CreatePartition < uint32(1) || o.CreatePartition > uint32(128) {
		return fmt.Errorf("CreatePartition is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1))(2:{alias=MGMT_OBJECT}(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=ULONG}(uint32))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Property {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Property[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Property); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint32(0)); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1))(2:{alias=MGMT_OBJECT}(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ManagementObject{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cp {in} (1:{range=(1,128), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreatePartition); err != nil {
			return err
		}
	}
	// aProp {in} (1:[dim:0,size_is=cp])(2:{alias=ULONG}(uint32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Property", sizeInfo[0])
		}
		o.Property = make([]uint32, sizeInfo[0])
		for i1 := range o.Property {
			i1 := i1
			if err := w.ReadData(&o.Property[i1]); err != nil {
				return err
			}
		}
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
	{
		dimSize1 := uint64(o.CreatePartition)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Var {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Var[i1] != nil {
				if err := o.Var[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Var); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&mqmq.PropertyVariant{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ManagementGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// apVar {in, out} (1:[dim:0,size_is=cp])(2:{alias=PROPVARIANT}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Var", sizeInfo[0])
		}
		o.Var = make([]*mqmq.PropertyVariant, sizeInfo[0])
		for i1 := range o.Var {
			i1 := i1
			if o.Var[i1] == nil {
				o.Var[i1] = &mqmq.PropertyVariant{}
			}
			if err := o.Var[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
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

// ManagementGetInfoRequest structure represents the R_QMMgmtGetInfo operation request
type ManagementGetInfoRequest struct {
	// pObjectFormat: A pointer to an MGMT_OBJECT structure that defines the queue or computer
	// on which to return information.
	ObjectFormat *ManagementObject `idl:"name:pObjectFormat" json:"object_format"`
	// cp: The length (in elements) of the arrays aProp and apVar MUST be at least 1, and
	// MUST be at most 128.
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`
	// aProp: Points to an array of property identifiers associated with the array of property
	// values. This array MUST contain at least one element. Each element MUST specify a
	// value from the property identifiers table, as specified in section 2.2.3. Each element
	// MUST specify the property identifier for the corresponding property value at the
	// same element index in apVar. This array and the array to which apVar points MUST
	// be of the same length.
	Property []uint32 `idl:"name:aProp;size_is:(cp)" json:"property"`
	// apVar: Points to an array that specifies the property values associated with the
	// array of property identifiers. Each element in this array specifies the property
	// value for the corresponding property identifier at the same element index in the
	// array to which aProp points. This array MUST contain at least one element. The property
	// value in each element MUST correspond accordingly to the property identifier from
	// aProp, as specified in section 2.2.3, and MUST be set to VT_NULL<14> (as specified
	// in [MS-MQMQ] section 2.2.12) before each call to R_QMMgmtGetInfo. This array and
	// the array to which aProp points MUST be of the same length.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
}

func (o *ManagementGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ManagementGetInfoOperation) *xxx_ManagementGetInfoOperation {
	if op == nil {
		op = &xxx_ManagementGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.CreatePartition = o.CreatePartition
	op.Property = o.Property
	op.Var = o.Var
	return op
}

func (o *ManagementGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ManagementGetInfoOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.CreatePartition = op.CreatePartition
	o.Property = op.Property
	o.Var = op.Var
}
func (o *ManagementGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ManagementGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ManagementGetInfoResponse structure represents the R_QMMgmtGetInfo operation response
type ManagementGetInfoResponse struct {
	// XXX: cp is an implicit input depedency for output parameters
	CreatePartition uint32 `idl:"name:cp" json:"create_partition"`

	// apVar: Points to an array that specifies the property values associated with the
	// array of property identifiers. Each element in this array specifies the property
	// value for the corresponding property identifier at the same element index in the
	// array to which aProp points. This array MUST contain at least one element. The property
	// value in each element MUST correspond accordingly to the property identifier from
	// aProp, as specified in section 2.2.3, and MUST be set to VT_NULL<14> (as specified
	// in [MS-MQMQ] section 2.2.12) before each call to R_QMMgmtGetInfo. This array and
	// the array to which aProp points MUST be of the same length.
	Var []*mqmq.PropertyVariant `idl:"name:apVar;size_is:(cp)" json:"var"`
	// Return: The R_QMMgmtGetInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ManagementGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ManagementGetInfoOperation) *xxx_ManagementGetInfoOperation {
	if op == nil {
		op = &xxx_ManagementGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.CreatePartition == uint32(0) {
		op.CreatePartition = o.CreatePartition
	}

	op.Var = o.Var
	op.Return = o.Return
	return op
}

func (o *ManagementGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ManagementGetInfoOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.CreatePartition = op.CreatePartition

	o.Var = op.Var
	o.Return = op.Return
}
func (o *ManagementGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ManagementGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ManagementActionOperation structure represents the R_QMMgmtAction operation
type xxx_ManagementActionOperation struct {
	ObjectFormat *ManagementObject `idl:"name:pObjectFormat" json:"object_format"`
	Action       string            `idl:"name:lpwszAction;string" json:"action"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ManagementActionOperation) OpNum() int { return 1 }

func (o *xxx_ManagementActionOperation) OpName() string { return "/qmmgmt/v1/R_QMMgmtAction" }

func (o *xxx_ManagementActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pObjectFormat {in} (1:{pointer=ref}*(1))(2:{alias=MGMT_OBJECT}(struct))
	{
		if o.ObjectFormat != nil {
			if err := o.ObjectFormat.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ManagementObject{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpwszAction {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pObjectFormat {in} (1:{pointer=ref}*(1))(2:{alias=MGMT_OBJECT}(struct))
	{
		if o.ObjectFormat == nil {
			o.ObjectFormat = &ManagementObject{}
		}
		if err := o.ObjectFormat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwszAction {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Action); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagementActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ManagementActionRequest structure represents the R_QMMgmtAction operation request
type ManagementActionRequest struct {
	// pObjectFormat: A pointer to a MGMT_OBJECT structure that specifies the queue or computer
	// to which the action is being applied.
	ObjectFormat *ManagementObject `idl:"name:pObjectFormat" json:"object_format"`
	// lpwszAction:  A pointer to a null-terminated Unicode string that specifies the action
	// to perform on the computer. The lpwszAction value MUST be one of the following (the
	// value is not case-sensitive).
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|              |                                                                                  |
	//	|    VALUE     |                                     MEANING                                      |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "CONNECT"    | A machine action. Connects the computer to the network and the MSMQ Directory    |
	//	|              | Service server.                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "DISCONNECT" | A machine action. Disconnects the computer from the network and the MSMQ         |
	//	|              | Directory Service server.                                                        |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "TIDY"       | A machine action. Cleans up empty message files. MSMQ does this every 6 hours.   |
	//	|              | It is helpful when a large number of messages are deleted (purged or received by |
	//	|              | an application), and the application needs the disk space immediately.           |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "PAUSE"      | A queue action. Valid for outgoing queues only. Stops the sending of messages    |
	//	|              | from the computer. The queue manager will not send messages to the applicable    |
	//	|              | destination queue until a RESUME action is initiated.                            |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "RESUME"     | A queue action. Valid for outgoing queues only. Restarts the sending of messages |
	//	|              | after a PAUSE action is initiated.                                               |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| "EOD_RESEND" | A queue action. Resends the pending transaction sequence.                        |
	//	+--------------+----------------------------------------------------------------------------------+
	Action string `idl:"name:lpwszAction;string" json:"action"`
}

func (o *ManagementActionRequest) xxx_ToOp(ctx context.Context, op *xxx_ManagementActionOperation) *xxx_ManagementActionOperation {
	if op == nil {
		op = &xxx_ManagementActionOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectFormat = o.ObjectFormat
	op.Action = o.Action
	return op
}

func (o *ManagementActionRequest) xxx_FromOp(ctx context.Context, op *xxx_ManagementActionOperation) {
	if o == nil {
		return
	}
	o.ObjectFormat = op.ObjectFormat
	o.Action = op.Action
}
func (o *ManagementActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ManagementActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ManagementActionResponse structure represents the R_QMMgmtAction operation response
type ManagementActionResponse struct {
	// Return: The R_QMMgmtAction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ManagementActionResponse) xxx_ToOp(ctx context.Context, op *xxx_ManagementActionOperation) *xxx_ManagementActionOperation {
	if op == nil {
		op = &xxx_ManagementActionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ManagementActionResponse) xxx_FromOp(ctx context.Context, op *xxx_ManagementActionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ManagementActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ManagementActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagementActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
