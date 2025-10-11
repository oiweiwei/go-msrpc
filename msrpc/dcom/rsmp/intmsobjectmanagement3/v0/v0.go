package intmsobjectmanagement3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	intmsobjectmanagement2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp/intmsobjectmanagement2/v0"
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
	_ = dcom.GoPackage
	_ = intmsobjectmanagement2.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// INtmsObjectManagement3 interface identifier 3bbed8d9-2c9a-4b21-8936-acb2f995be6c
	ObjectManagement3IID = &dcom.IID{Data1: 0x3bbed8d9, Data2: 0x2c9a, Data3: 0x4b21, Data4: []byte{0x89, 0x36, 0xac, 0xb2, 0xf9, 0x95, 0xbe, 0x6c}}
	// Syntax UUID
	ObjectManagement3SyntaxUUID = &uuid.UUID{TimeLow: 0x3bbed8d9, TimeMid: 0x2c9a, TimeHiAndVersion: 0x4b21, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x36, Node: [6]uint8{0xac, 0xb2, 0xf9, 0x95, 0xbe, 0x6c}}
	// Syntax ID
	ObjectManagement3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectManagement3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// INtmsObjectManagement3 interface.
type ObjectManagement3Client interface {

	// INtmsObjectManagement2 retrieval method.
	ObjectManagement2() intmsobjectmanagement2.ObjectManagement2Client

	// The GetNtmsObjectAttributeAR method retrieves private data from an object, with strings
	// encoded using ASCII.
	GetNTMSObjectAttributeAR(context.Context, *GetNTMSObjectAttributeARRequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeARResponse, error)

	// The GetNtmsObjectAttributeWR method retrieves private data from an object, with strings
	// encoded using Unicode.
	GetNTMSObjectAttributeWR(context.Context, *GetNTMSObjectAttributeWRRequest, ...dcerpc.CallOption) (*GetNTMSObjectAttributeWRResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ObjectManagement3Client
}

type xxx_DefaultObjectManagement3Client struct {
	intmsobjectmanagement2.ObjectManagement2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultObjectManagement3Client) ObjectManagement2() intmsobjectmanagement2.ObjectManagement2Client {
	return o.ObjectManagement2Client
}

func (o *xxx_DefaultObjectManagement3Client) GetNTMSObjectAttributeAR(ctx context.Context, in *GetNTMSObjectAttributeARRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeARResponse, error) {
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
	out := &GetNTMSObjectAttributeARResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement3Client) GetNTMSObjectAttributeWR(ctx context.Context, in *GetNTMSObjectAttributeWRRequest, opts ...dcerpc.CallOption) (*GetNTMSObjectAttributeWRResponse, error) {
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
	out := &GetNTMSObjectAttributeWRResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectManagement3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectManagement3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultObjectManagement3Client) IPID(ctx context.Context, ipid *dcom.IPID) ObjectManagement3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultObjectManagement3Client{
		ObjectManagement2Client: o.ObjectManagement2Client.IPID(ctx, ipid),
		cc:                      o.cc,
		ipid:                    ipid,
	}
}

func NewObjectManagement3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectManagement3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectManagement3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := intmsobjectmanagement2.NewObjectManagement2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultObjectManagement3Client{
		ObjectManagement2Client: base,
		cc:                      cc,
		ipid:                    ipid,
	}, nil
}

// xxx_GetNTMSObjectAttributeAROperation structure represents the GetNtmsObjectAttributeAR operation
type xxx_GetNTMSObjectAttributeAROperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       uint8          `idl:"name:lpAttributeName" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeAROperation) OpNum() int { return 14 }

func (o *xxx_GetNTMSObjectAttributeAROperation) OpName() string {
	return "/INtmsObjectManagement3/v0/GetNtmsObjectAttributeAR"
}

func (o *xxx_GetNTMSObjectAttributeAROperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAROperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.WriteData(o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAROperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{pointer=ref}*(1)(char))
	{
		if err := w.ReadData(&o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAROperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeAROperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualAttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeAROperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualAttributeSize); err != nil {
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

// GetNTMSObjectAttributeARRequest structure represents the GetNtmsObjectAttributeAR operation request
type GetNTMSObjectAttributeARRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	// dwType: The value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of ASCII characters specifying the name
	// of the extended attribute to retrieve.
	AttributeName uint8 `idl:"name:lpAttributeName" json:"attribute_name"`
	// lpdwAttributeBufferSize: A pointer to the size of the lpAttributeData buffer.
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeARRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAROperation) *xxx_GetNTMSObjectAttributeAROperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeAROperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeARRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAROperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeARRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeARRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeAROperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeARResponse structure represents the GetNtmsObjectAttributeAR operation response
type GetNTMSObjectAttributeARResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpAttributeData: A buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	// lpAttributeSize: A pointer to the size of the attribute returned in the lpAttributeData
	// buffer. This will point to 0 when the function returns with an insufficient input
	// buffer error.
	AttributeSize uint32 `idl:"name:lpAttributeSize" json:"attribute_size"`
	// lpActualAttributeSize: A pointer to the actual size of the attribute.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to an object was denied.                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | A parameter is not valid.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The specified buffer size is not large enough.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700E8 ERROR_NO_DATA             | The specified attribute is greater than or equal to the NTMS_MAXATTR_LENGTH      |
	//	|                                      | value, specified in the Platform SDK file NTMSApi.h.                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED       | Unable to connect to the server.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND    | The object was not found.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE    | The database query or update failed.                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// After the server receives this message, it MUST:
	//
	// * Verify that ObjectId is not NULL.
	//
	// * Verify that lpAttributeData is not NULL.
	//
	// * Verify that lpAttributeName is not NULL.
	//
	// * Verify that lpAttributeSize is not NULL.
	//
	// * Verify that dwType is a valid object type.
	//
	// If parameter validation fails, the server MUST immediately fail the operation and
	// return ERROR_INVALID_PARAMETER (0x80070057).
	//
	// If parameter validation succeeds, the server MUST verify that the user has the required
	// access rights, get the value of the extended attribute that is specified by lpAttributeName,
	// return it to the user in the buffer that is pointed to by lpAttributeData, and set
	// the size of the data that is copied in the lpAttributeData in lpAttributeSize. If
	// the client does not have the required access rights, the server MUST return ERROR_ACCESS_DENIED
	// (0x80070005).
	//
	// If the buffer size that is specified by lpdwAttributeBufferSize is too small, the
	// server MUST return ERROR_INSUFFICIENT_BUFFER (0x8007007A) with lpActualAttributeSize
	// set to the required size and lpAttributeSize set to zero.
	//
	// Strings that are sent to this method as parameters MUST be ASCII-encoded.
	ActualAttributeSize uint32 `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	// Return: The GetNtmsObjectAttributeAR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeARResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAROperation) *xxx_GetNTMSObjectAttributeAROperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeAROperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.ActualAttributeSize = o.ActualAttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeARResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeAROperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.ActualAttributeSize = op.ActualAttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeARResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeARResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeAROperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNTMSObjectAttributeWROperation structure represents the GetNtmsObjectAttributeWR operation
type xxx_GetNTMSObjectAttributeWROperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	ObjectID            *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	Type                uint32         `idl:"name:dwType" json:"type"`
	AttributeName       string         `idl:"name:lpAttributeName;string" json:"attribute_name"`
	AttributeData       []byte         `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	AttributeBufferSize uint32         `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
	AttributeSize       uint32         `idl:"name:lpAttributeSize" json:"attribute_size"`
	ActualAttributeSize uint32         `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNTMSObjectAttributeWROperation) OpNum() int { return 15 }

func (o *xxx_GetNTMSObjectAttributeWROperation) OpName() string {
	return "/INtmsObjectManagement3/v0/GetNtmsObjectAttributeWR"
}

func (o *xxx_GetNTMSObjectAttributeWROperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWROperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID != nil {
			if err := o.ObjectID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWROperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lpObjectId {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ObjectID == nil {
			o.ObjectID = &dtyp.GUID{}
		}
		if err := o.ObjectID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpAttributeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.AttributeName); err != nil {
			return err
		}
	}
	// lpdwAttributeBufferSize {in} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWROperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.AttributeData != nil && o.AttributeSize == 0 {
		o.AttributeSize = uint32(len(o.AttributeData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNTMSObjectAttributeWROperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
		dimSize1 := uint64(o.AttributeBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.AttributeSize)
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
		for i1 := range o.AttributeData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AttributeData[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AttributeData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualAttributeSize); err != nil {
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

func (o *xxx_GetNTMSObjectAttributeWROperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// lpAttributeData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpdwAttributeBufferSize,length_is=lpAttributeSize](uint8))
	{
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AttributeData", sizeInfo[0])
		}
		o.AttributeData = make([]byte, sizeInfo[0])
		for i1 := range o.AttributeData {
			i1 := i1
			if err := w.ReadData(&o.AttributeData[i1]); err != nil {
				return err
			}
		}
	}
	// lpAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AttributeSize); err != nil {
			return err
		}
	}
	// lpActualAttributeSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualAttributeSize); err != nil {
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

// GetNTMSObjectAttributeWRRequest structure represents the GetNtmsObjectAttributeWR operation request
type GetNTMSObjectAttributeWRRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	ObjectID *dtyp.GUID     `idl:"name:lpObjectId" json:"object_id"`
	// dwType: The value from the NtmsObjectsTypes (section 2.2.1.6) enumeration specifying
	// the type of the object.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpAttributeName: A null-terminated sequence of Unicode characters specifying the
	// name of the extended attribute to retrieve.
	AttributeName string `idl:"name:lpAttributeName;string" json:"attribute_name"`
	// lpdwAttributeBufferSize: A pointer to the size of the lpAttributeData buffer.
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`
}

func (o *GetNTMSObjectAttributeWRRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWROperation) *xxx_GetNTMSObjectAttributeWROperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWROperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ObjectID = o.ObjectID
	op.Type = o.Type
	op.AttributeName = o.AttributeName
	op.AttributeBufferSize = o.AttributeBufferSize
	return op
}

func (o *GetNTMSObjectAttributeWRRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWROperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ObjectID = op.ObjectID
	o.Type = op.Type
	o.AttributeName = op.AttributeName
	o.AttributeBufferSize = op.AttributeBufferSize
}
func (o *GetNTMSObjectAttributeWRRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNTMSObjectAttributeWRRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWROperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNTMSObjectAttributeWRResponse structure represents the GetNtmsObjectAttributeWR operation response
type GetNTMSObjectAttributeWRResponse struct {
	// XXX: lpdwAttributeBufferSize is an implicit input depedency for output parameters
	AttributeBufferSize uint32 `idl:"name:lpdwAttributeBufferSize" json:"attribute_buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lpAttributeData: A buffer containing the attribute.
	AttributeData []byte `idl:"name:lpAttributeData;size_is:(lpdwAttributeBufferSize);length_is:(lpAttributeSize)" json:"attribute_data"`
	// lpAttributeSize: A pointer to the size of the attribute returned in the lpAttributeData
	// buffer. This will point to zero when the function returns with an insufficient input
	// buffer error.
	AttributeSize uint32 `idl:"name:lpAttributeSize" json:"attribute_size"`
	// lpActualAttributeSize: A pointer to the size of the attribute in the lpAttributeData
	// buffer.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 ERROR_ACCESS_DENIED       | Access to an object was denied.                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER   | A parameter is not valid.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A ERROR_INSUFFICIENT_BUFFER | The specified buffer size is not large enough.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800700E8 ERROR_NO_DATA             | The specified attribute is greater than or equal to the NTMS_MAXATTR_LENGTH      |
	//	|                                      | value, specified in the Platform SDK file NTMSApi.h.                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800708CA ERROR_NOT_CONNECTED       | Unable to connect to the server.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D8 ERROR_OBJECT_NOT_FOUND    | The object was not found.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D9 ERROR_DATABASE_FAILURE    | The database query or update failed.                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	ActualAttributeSize uint32 `idl:"name:lpActualAttributeSize" json:"actual_attribute_size"`
	// Return: The GetNtmsObjectAttributeWR return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNTMSObjectAttributeWRResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWROperation) *xxx_GetNTMSObjectAttributeWROperation {
	if op == nil {
		op = &xxx_GetNTMSObjectAttributeWROperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.AttributeBufferSize == uint32(0) {
		op.AttributeBufferSize = o.AttributeBufferSize
	}

	op.That = o.That
	op.AttributeData = o.AttributeData
	op.AttributeSize = o.AttributeSize
	op.ActualAttributeSize = o.ActualAttributeSize
	op.Return = o.Return
	return op
}

func (o *GetNTMSObjectAttributeWRResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNTMSObjectAttributeWROperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.AttributeBufferSize = op.AttributeBufferSize

	o.That = op.That
	o.AttributeData = op.AttributeData
	o.AttributeSize = op.AttributeSize
	o.ActualAttributeSize = op.ActualAttributeSize
	o.Return = op.Return
}
func (o *GetNTMSObjectAttributeWRResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNTMSObjectAttributeWRResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNTMSObjectAttributeWROperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
