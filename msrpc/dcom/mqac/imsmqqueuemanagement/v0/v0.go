package imsmqqueuemanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = imsmqmanagement.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQQueueManagement interface identifier 7fbe7759-5760-444d-b8a5-5e7ab9a84cce
	QueueManagementIID = &dcom.IID{Data1: 0x7fbe7759, Data2: 0x5760, Data3: 0x444d, Data4: []byte{0xb8, 0xa5, 0x5e, 0x7a, 0xb9, 0xa8, 0x4c, 0xce}}
	// Syntax UUID
	QueueManagementSyntaxUUID = &uuid.UUID{TimeLow: 0x7fbe7759, TimeMid: 0x5760, TimeHiAndVersion: 0x444d, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0xa5, Node: [6]uint8{0x5e, 0x7a, 0xb9, 0xa8, 0x4c, 0xce}}
	// Syntax ID
	QueueManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QueueManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQueueManagement interface.
type QueueManagementClient interface {

	// IMSMQManagement retrieval method.
	Management() imsmqmanagement.ManagementClient

	// The JournalMessageCount method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the number of messages in the Queue that is associated
	// with the represented Queue.JournalQueueReference.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_QUEUE_JOURNAL_MESSAGE_COUNT
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * The plJournalMessageCount output variable MUST be set to the value of the returned
	// *rPropVar*.
	GetJournalMessageCount(context.Context, *GetJournalMessageCountRequest, ...dcerpc.CallOption) (*GetJournalMessageCountResponse, error)

	// The BytesInJournal method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the Queue.TotalBytes of the Queue that is associated
	// with the represented Queue.JournalQueueReference.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_QUEUE_BYTES_IN_JOURNAL
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * The plJournalMessageCount output variable MUST be set to the value of the returned
	// *rPropVar*.
	GetBytesInJournal(context.Context, *GetBytesInJournalRequest, ...dcerpc.CallOption) (*GetBytesInJournalResponse, error)

	// The EodGetReceiveInfo method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Queue.IncomingTransactionalTransferInfoCollection.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094), and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_EOD_SOURCE_INFO
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * The pvCollection output variable MUST be set to the returned *rPropVar*.
	EODGetReceiveInfo(context.Context, *EODGetReceiveInfoRequest, ...dcerpc.CallOption) (*EODGetReceiveInfoResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QueueManagementClient
}

type xxx_DefaultQueueManagementClient struct {
	imsmqmanagement.ManagementClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQueueManagementClient) Management() imsmqmanagement.ManagementClient {
	return o.ManagementClient
}

func (o *xxx_DefaultQueueManagementClient) GetJournalMessageCount(ctx context.Context, in *GetJournalMessageCountRequest, opts ...dcerpc.CallOption) (*GetJournalMessageCountResponse, error) {
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
	out := &GetJournalMessageCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueManagementClient) GetBytesInJournal(ctx context.Context, in *GetBytesInJournalRequest, opts ...dcerpc.CallOption) (*GetBytesInJournalResponse, error) {
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
	out := &GetBytesInJournalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueManagementClient) EODGetReceiveInfo(ctx context.Context, in *EODGetReceiveInfoRequest, opts ...dcerpc.CallOption) (*EODGetReceiveInfoResponse, error) {
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
	out := &EODGetReceiveInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQueueManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQueueManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) QueueManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQueueManagementClient{
		ManagementClient: o.ManagementClient.IPID(ctx, ipid),
		cc:               o.cc,
		ipid:             ipid,
	}
}

func NewQueueManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QueueManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QueueManagementSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqmanagement.NewManagementClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQueueManagementClient{
		ManagementClient: base,
		cc:               cc,
		ipid:             ipid,
	}, nil
}

// xxx_GetJournalMessageCountOperation structure represents the JournalMessageCount operation
type xxx_GetJournalMessageCountOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	JournalMessageCount int32          `idl:"name:plJournalMessageCount" json:"journal_message_count"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJournalMessageCountOperation) OpNum() int { return 16 }

func (o *xxx_GetJournalMessageCountOperation) OpName() string {
	return "/IMSMQQueueManagement/v0/JournalMessageCount"
}

func (o *xxx_GetJournalMessageCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalMessageCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetJournalMessageCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetJournalMessageCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalMessageCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plJournalMessageCount {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.JournalMessageCount); err != nil {
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

func (o *xxx_GetJournalMessageCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plJournalMessageCount {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.JournalMessageCount); err != nil {
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

// GetJournalMessageCountRequest structure represents the JournalMessageCount operation request
type GetJournalMessageCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetJournalMessageCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetJournalMessageCountOperation) *xxx_GetJournalMessageCountOperation {
	if op == nil {
		op = &xxx_GetJournalMessageCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetJournalMessageCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJournalMessageCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetJournalMessageCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetJournalMessageCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalMessageCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJournalMessageCountResponse structure represents the JournalMessageCount operation response
type GetJournalMessageCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plJournalMessageCount: A pointer to a long that, when successfully completed, contains
	// the number of messages in the Queue that is associated with the represented Queue.JournalQueueReference.
	JournalMessageCount int32 `idl:"name:plJournalMessageCount" json:"journal_message_count"`
	// Return: The JournalMessageCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetJournalMessageCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetJournalMessageCountOperation) *xxx_GetJournalMessageCountOperation {
	if op == nil {
		op = &xxx_GetJournalMessageCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.JournalMessageCount = o.JournalMessageCount
	op.Return = o.Return
	return op
}

func (o *GetJournalMessageCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJournalMessageCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.JournalMessageCount = op.JournalMessageCount
	o.Return = op.Return
}
func (o *GetJournalMessageCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetJournalMessageCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalMessageCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBytesInJournalOperation structure represents the BytesInJournal operation
type xxx_GetBytesInJournalOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	BytesInJournal *oaut.Variant  `idl:"name:pvBytesInJournal" json:"bytes_in_journal"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBytesInJournalOperation) OpNum() int { return 17 }

func (o *xxx_GetBytesInJournalOperation) OpName() string {
	return "/IMSMQQueueManagement/v0/BytesInJournal"
}

func (o *xxx_GetBytesInJournalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInJournalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetBytesInJournalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetBytesInJournalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInJournalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvBytesInJournal {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.BytesInJournal != nil {
			_ptr_pvBytesInJournal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BytesInJournal != nil {
					if err := o.BytesInJournal.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BytesInJournal, _ptr_pvBytesInJournal); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
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

func (o *xxx_GetBytesInJournalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvBytesInJournal {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvBytesInJournal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BytesInJournal == nil {
				o.BytesInJournal = &oaut.Variant{}
			}
			if err := o.BytesInJournal.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvBytesInJournal := func(ptr interface{}) { o.BytesInJournal = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.BytesInJournal, _s_pvBytesInJournal, _ptr_pvBytesInJournal); err != nil {
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

// GetBytesInJournalRequest structure represents the BytesInJournal operation request
type GetBytesInJournalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBytesInJournalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInJournalOperation) *xxx_GetBytesInJournalOperation {
	if op == nil {
		op = &xxx_GetBytesInJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBytesInJournalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInJournalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBytesInJournalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBytesInJournalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInJournalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBytesInJournalResponse structure represents the BytesInJournal operation response
type GetBytesInJournalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvBytesInJournal: A pointer to a VARIANT that, when successfully completed, contains
	// an unsigned 64-bit integer (VT_UI8) that specifies the number of message bytes in
	// the Queue that is associated with the represented Queue.JournalQueueReference.
	BytesInJournal *oaut.Variant `idl:"name:pvBytesInJournal" json:"bytes_in_journal"`
	// Return: The BytesInJournal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBytesInJournalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInJournalOperation) *xxx_GetBytesInJournalOperation {
	if op == nil {
		op = &xxx_GetBytesInJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BytesInJournal = o.BytesInJournal
	op.Return = o.Return
	return op
}

func (o *GetBytesInJournalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInJournalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BytesInJournal = op.BytesInJournal
	o.Return = op.Return
}
func (o *GetBytesInJournalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBytesInJournalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInJournalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EODGetReceiveInfoOperation structure represents the EodGetReceiveInfo operation
type xxx_EODGetReceiveInfoOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Collection *oaut.Variant  `idl:"name:pvCollection" json:"collection"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EODGetReceiveInfoOperation) OpNum() int { return 18 }

func (o *xxx_EODGetReceiveInfoOperation) OpName() string {
	return "/IMSMQQueueManagement/v0/EodGetReceiveInfo"
}

func (o *xxx_EODGetReceiveInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetReceiveInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_EODGetReceiveInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_EODGetReceiveInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EODGetReceiveInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvCollection {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Collection != nil {
			_ptr_pvCollection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_pvCollection); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
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

func (o *xxx_EODGetReceiveInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvCollection {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvCollection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &oaut.Variant{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvCollection := func(ptr interface{}) { o.Collection = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Collection, _s_pvCollection, _ptr_pvCollection); err != nil {
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

// EODGetReceiveInfoRequest structure represents the EodGetReceiveInfo operation request
type EODGetReceiveInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EODGetReceiveInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_EODGetReceiveInfoOperation) *xxx_EODGetReceiveInfoOperation {
	if op == nil {
		op = &xxx_EODGetReceiveInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *EODGetReceiveInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_EODGetReceiveInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EODGetReceiveInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EODGetReceiveInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODGetReceiveInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EODGetReceiveInfoResponse structure represents the EodGetReceiveInfo operation response
type EODGetReceiveInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvCollection: A pointer to a VARIANT that contains an array of VARIANTs (VT_ARRAY | VT_VARIANT) in which each VARIANT contains an EODTargetInfo (section 2.2.4.1) collection.
	Collection *oaut.Variant `idl:"name:pvCollection" json:"collection"`
	// Return: The EodGetReceiveInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EODGetReceiveInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_EODGetReceiveInfoOperation) *xxx_EODGetReceiveInfoOperation {
	if op == nil {
		op = &xxx_EODGetReceiveInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *EODGetReceiveInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_EODGetReceiveInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *EODGetReceiveInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EODGetReceiveInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EODGetReceiveInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
