package imsmqmanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQManagement interface identifier be5f0241-e489-4957-8cc4-a452fcf3e23e
	ManagementIID = &dcom.IID{Data1: 0xbe5f0241, Data2: 0xe489, Data3: 0x4957, Data4: []byte{0x8c, 0xc4, 0xa4, 0x52, 0xfc, 0xf3, 0xe2, 0x3e}}
	// Syntax UUID
	ManagementSyntaxUUID = &uuid.UUID{TimeLow: 0xbe5f0241, TimeMid: 0xe489, TimeHiAndVersion: 0x4957, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xc4, Node: [6]uint8{0xa4, 0x52, 0xfc, 0xf3, 0xe2, 0x3e}}
	// Syntax ID
	ManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQManagement interface.
type ManagementClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The Init method is received by the server in an RPC_REQUEST packet. In response,
	// the server initializes the object to represent the state of a Queue. The represented
	// Queue MUST have Queue.Active set to True. If the represented Queue.Active is False,
	// or the Pathname and FormatName input parameters cannot be resolved, this method MUST
	// return an error without setting the ObjectIsInitialized instance variable to True.
	// This method MUST be called prior to calling any other operation on MSMQManagement.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	Init(context.Context, *InitRequest, ...dcerpc.CallOption) (*InitResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest, ...dcerpc.CallOption) (*GetFormatNameResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest, ...dcerpc.CallOption) (*GetMachineResponse, error)

	// The MessageCount method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the number of messages in the represented Queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetMessageCount(context.Context, *GetMessageCountRequest, ...dcerpc.CallOption) (*GetMessageCountResponse, error)

	// The ForeignStatus method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return an enumerated value to indicate whether a Queue is a foreign
	// queue or an OutgoingQueue that transfers messages to a foreign queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094) and MUST take
	// no further action.
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_QUEUE_FOREIGN
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * If the value of the returned *rPropVar* was "Yes":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_FOREIGN, and the
	// server MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "No":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_NOT_FOREIGN, and
	// the server MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "Unknown":
	//
	// * The *pIForeignStatus* output variable MUST be set to MQ_STATUS_UNKNOWN, and the
	// server MUST take no further action.
	GetForeignStatus(context.Context, *GetForeignStatusRequest, ...dcerpc.CallOption) (*GetForeignStatusResponse, error)

	// The QueueType method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the type of the referenced Queue.
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
	// * Else:
	//
	// * The server MUST generate a *QMMgmt Get Info* event with the following inputs:
	//
	// * iPropID = PROPID_MGMT_QUEUE_TYPE
	//
	// * If the rStatus return value is not equal to MQ_OK (0x00000000), the server MUST
	// return rStatus and MUST take no further action.
	//
	// * If the value of the returned *rPropVar* was "PUBLIC":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_PUBLIC, and the server MUST
	// return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "PRIVATE":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_PRIVATE, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "MACHINE":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_MACHINE, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "CONNECTOR":
	//
	// * The plQueueType output variable MUST be set to MGMT_TYPE_CONNECTOR, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	//
	// * If the value of the returned was "MULTICAST":
	//
	// * The plQueueType output variable MUST be set to MQ_TYPE_MULTICAST, and the server
	// MUST return S_OK (0x00000000) and MUST take no further action.
	GetQueueType(context.Context, *GetQueueTypeRequest, ...dcerpc.CallOption) (*GetQueueTypeResponse, error)

	// The IsLocal method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a BOOLEAN value that indicates whether the represented Queue
	// is an OutgoingQueue (False) or not (True).
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ObjectIsInitialized* instance variable is False:
	//
	// * The server MUST return MQ_ERROR_UNINITIALIZED_OBJECT (0xC00E0094) and MUST take
	// no further action.
	//
	// * The server MUST generate a QMMgmt Get Info ( a0b1e28b-0f93-415d-8753-e3e1133678db
	// ) event with the following inputs:
	//
	// * *iPropID* = PROPID_MGMT_QUEUE_LOCATION
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * If the value of the returned *rPropVar* was "LOCAL":
	//
	// * The pfIsLocal output variable MUST be set to True, and the server MUST take no
	// further action.
	//
	// * If the value of the returned *rPropVar* was "REMOTE":
	//
	// * The pfIsLocal output variable MUST be set to False, and the server MUST take no
	// further action.
	//
	// * If the value of the returned *rPropVar* was "LOCAL":
	//
	// * The pfIsLocal output variable MUST be set to True, and the server MUST take no
	// further action.
	GetIsLocal(context.Context, *GetIsLocalRequest, ...dcerpc.CallOption) (*GetIsLocalResponse, error)

	// The TransactionalStatus method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return the represented Queue.Transactional value.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetTransactionalStatus(context.Context, *GetTransactionalStatusRequest, ...dcerpc.CallOption) (*GetTransactionalStatusResponse, error)

	// The BytesInQueue method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return the value of the represented Queue.TotalBytes property.
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
	// * *iPropID* = PROPID_MGMT_QUEUE_BYTES_IN_ QUEUE
	//
	// * If the *rStatus* return value is not equal to MQ_OK (0x00000000), the server MUST
	// return *rStatus* and MUST take no further action.
	//
	// * Else:
	//
	// * The pvBytesInQueue output variable MUST be set to the value of the returned *rPropVar*.
	GetBytesInQueue(context.Context, *GetBytesInQueueRequest, ...dcerpc.CallOption) (*GetBytesInQueueResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ManagementClient
}

type xxx_DefaultManagementClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultManagementClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultManagementClient) Init(ctx context.Context, in *InitRequest, opts ...dcerpc.CallOption) (*InitResponse, error) {
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
	out := &InitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetFormatName(ctx context.Context, in *GetFormatNameRequest, opts ...dcerpc.CallOption) (*GetFormatNameResponse, error) {
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
	out := &GetFormatNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetMachine(ctx context.Context, in *GetMachineRequest, opts ...dcerpc.CallOption) (*GetMachineResponse, error) {
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
	out := &GetMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetMessageCount(ctx context.Context, in *GetMessageCountRequest, opts ...dcerpc.CallOption) (*GetMessageCountResponse, error) {
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
	out := &GetMessageCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetForeignStatus(ctx context.Context, in *GetForeignStatusRequest, opts ...dcerpc.CallOption) (*GetForeignStatusResponse, error) {
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
	out := &GetForeignStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetQueueType(ctx context.Context, in *GetQueueTypeRequest, opts ...dcerpc.CallOption) (*GetQueueTypeResponse, error) {
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
	out := &GetQueueTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetIsLocal(ctx context.Context, in *GetIsLocalRequest, opts ...dcerpc.CallOption) (*GetIsLocalResponse, error) {
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
	out := &GetIsLocalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetTransactionalStatus(ctx context.Context, in *GetTransactionalStatusRequest, opts ...dcerpc.CallOption) (*GetTransactionalStatusResponse, error) {
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
	out := &GetTransactionalStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) GetBytesInQueue(ctx context.Context, in *GetBytesInQueueRequest, opts ...dcerpc.CallOption) (*GetBytesInQueueResponse, error) {
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
	out := &GetBytesInQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) ManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultManagementClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ManagementSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultManagementClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_InitOperation structure represents the Init operation
type xxx_InitOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine    *oaut.Variant  `idl:"name:Machine" json:"machine"`
	PathName   *oaut.Variant  `idl:"name:Pathname" json:"path_name"`
	FormatName *oaut.Variant  `idl:"name:FormatName" json:"format_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_InitOperation) OpNum() int { return 7 }

func (o *xxx_InitOperation) OpName() string { return "/IMSMQManagement/v0/Init" }

func (o *xxx_InitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Machine {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Machine != nil {
			_ptr_Machine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Machine != nil {
					if err := o.Machine.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Machine, _ptr_Machine); err != nil {
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
	// Pathname {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.PathName != nil {
			_ptr_Pathname := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PathName != nil {
					if err := o.PathName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PathName, _ptr_Pathname); err != nil {
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
	// FormatName {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.FormatName != nil {
			_ptr_FormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FormatName != nil {
					if err := o.FormatName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FormatName, _ptr_FormatName); err != nil {
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
	return nil
}

func (o *xxx_InitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Machine {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Machine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Machine == nil {
				o.Machine = &oaut.Variant{}
			}
			if err := o.Machine.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Machine := func(ptr interface{}) { o.Machine = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Machine, _s_Machine, _ptr_Machine); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Pathname {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Pathname := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PathName == nil {
				o.PathName = &oaut.Variant{}
			}
			if err := o.PathName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Pathname := func(ptr interface{}) { o.PathName = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.PathName, _s_Pathname, _ptr_Pathname); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FormatName {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_FormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FormatName == nil {
				o.FormatName = &oaut.Variant{}
			}
			if err := o.FormatName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_FormatName := func(ptr interface{}) { o.FormatName = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.FormatName, _s_FormatName, _ptr_FormatName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_InitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// InitRequest structure represents the Init operation request
type InitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Machine: A pointer to a VARIANT that contains a BSTR that contains a string representation
	// of a computer name in the DNS or NetBIOS format. If this parameter is not specified
	// or is NULL, the server MUST ignore this parameter.
	Machine *oaut.Variant `idl:"name:Machine" json:"machine"`
	// Pathname: A pointer to a VARIANT that contains a BSTR that contains a string representation
	// of the path name describing a Queue.
	PathName *oaut.Variant `idl:"name:Pathname" json:"path_name"`
	// FormatName: A pointer to a VARIANT that contains a BSTR that contains a string representation
	// of the format name describing a Queue.
	FormatName *oaut.Variant `idl:"name:FormatName" json:"format_name"`
}

func (o *InitRequest) xxx_ToOp(ctx context.Context, op *xxx_InitOperation) *xxx_InitOperation {
	if op == nil {
		op = &xxx_InitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Machine = o.Machine
	op.PathName = o.PathName
	op.FormatName = o.FormatName
	return op
}

func (o *InitRequest) xxx_FromOp(ctx context.Context, op *xxx_InitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Machine = op.Machine
	o.PathName = op.PathName
	o.FormatName = op.FormatName
}
func (o *InitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitResponse structure represents the Init operation response
type InitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Init return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *InitResponse) xxx_ToOp(ctx context.Context, op *xxx_InitOperation) *xxx_InitOperation {
	if op == nil {
		op = &xxx_InitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *InitResponse) xxx_FromOp(ctx context.Context, op *xxx_InitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *InitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFormatNameOperation structure represents the FormatName operation
type xxx_GetFormatNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	FormatName *oaut.String   `idl:"name:pbstrFormatName" json:"format_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFormatNameOperation) OpNum() int { return 8 }

func (o *xxx_GetFormatNameOperation) OpName() string { return "/IMSMQManagement/v0/FormatName" }

func (o *xxx_GetFormatNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFormatNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFormatNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFormatNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFormatNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFormatName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FormatName != nil {
			_ptr_pbstrFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FormatName != nil {
					if err := o.FormatName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FormatName, _ptr_pbstrFormatName); err != nil {
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

func (o *xxx_GetFormatNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFormatName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FormatName == nil {
				o.FormatName = &oaut.String{}
			}
			if err := o.FormatName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFormatName := func(ptr interface{}) { o.FormatName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FormatName, _s_pbstrFormatName, _ptr_pbstrFormatName); err != nil {
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

// GetFormatNameRequest structure represents the FormatName operation request
type GetFormatNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFormatNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFormatNameOperation) *xxx_GetFormatNameOperation {
	if op == nil {
		op = &xxx_GetFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFormatNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFormatNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFormatNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFormatNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFormatNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFormatNameResponse structure represents the FormatName operation response
type GetFormatNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	FormatName *oaut.String   `idl:"name:pbstrFormatName" json:"format_name"`
	// Return: The FormatName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFormatNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFormatNameOperation) *xxx_GetFormatNameOperation {
	if op == nil {
		op = &xxx_GetFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FormatName = o.FormatName
	op.Return = o.Return
	return op
}

func (o *GetFormatNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFormatNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FormatName = op.FormatName
	o.Return = op.Return
}
func (o *GetFormatNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFormatNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFormatNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMachineOperation structure represents the Machine operation
type xxx_GetMachineOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine *oaut.String   `idl:"name:pbstrMachine" json:"machine"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMachineOperation) OpNum() int { return 9 }

func (o *xxx_GetMachineOperation) OpName() string { return "/IMSMQManagement/v0/Machine" }

func (o *xxx_GetMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Machine != nil {
			_ptr_pbstrMachine := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Machine != nil {
					if err := o.Machine.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Machine, _ptr_pbstrMachine); err != nil {
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

func (o *xxx_GetMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMachine {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachine := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Machine == nil {
				o.Machine = &oaut.String{}
			}
			if err := o.Machine.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachine := func(ptr interface{}) { o.Machine = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Machine, _s_pbstrMachine, _ptr_pbstrMachine); err != nil {
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

// GetMachineRequest structure represents the Machine operation request
type GetMachineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMachineOperation) *xxx_GetMachineOperation {
	if op == nil {
		op = &xxx_GetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMachineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMachineResponse structure represents the Machine operation response
type GetMachineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Machine *oaut.String   `idl:"name:pbstrMachine" json:"machine"`
	// Return: The Machine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMachineOperation) *xxx_GetMachineOperation {
	if op == nil {
		op = &xxx_GetMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Machine = o.Machine
	op.Return = o.Return
	return op
}

func (o *GetMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMachineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Machine = op.Machine
	o.Return = op.Return
}
func (o *GetMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMessageCountOperation structure represents the MessageCount operation
type xxx_GetMessageCountOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageCount int32          `idl:"name:plMessageCount" json:"message_count"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMessageCountOperation) OpNum() int { return 10 }

func (o *xxx_GetMessageCountOperation) OpName() string { return "/IMSMQManagement/v0/MessageCount" }

func (o *xxx_GetMessageCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMessageCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMessageCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plMessageCount {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.MessageCount); err != nil {
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

func (o *xxx_GetMessageCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plMessageCount {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.MessageCount); err != nil {
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

// GetMessageCountRequest structure represents the MessageCount operation request
type GetMessageCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMessageCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMessageCountOperation) *xxx_GetMessageCountOperation {
	if op == nil {
		op = &xxx_GetMessageCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMessageCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMessageCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMessageCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMessageCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMessageCountResponse structure represents the MessageCount operation response
type GetMessageCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plMessageCount: A pointer to a long that, when successfully completed, contains the
	// number of messages in the represented Queue.
	MessageCount int32 `idl:"name:plMessageCount" json:"message_count"`
	// Return: The MessageCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMessageCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMessageCountOperation) *xxx_GetMessageCountOperation {
	if op == nil {
		op = &xxx_GetMessageCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MessageCount = o.MessageCount
	op.Return = o.Return
	return op
}

func (o *GetMessageCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMessageCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageCount = op.MessageCount
	o.Return = op.Return
}
func (o *GetMessageCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMessageCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetForeignStatusOperation structure represents the ForeignStatus operation
type xxx_GetForeignStatusOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ForeignStatus int32          `idl:"name:plForeignStatus" json:"foreign_status"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetForeignStatusOperation) OpNum() int { return 11 }

func (o *xxx_GetForeignStatusOperation) OpName() string { return "/IMSMQManagement/v0/ForeignStatus" }

func (o *xxx_GetForeignStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetForeignStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetForeignStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetForeignStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetForeignStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plForeignStatus {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ForeignStatus); err != nil {
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

func (o *xxx_GetForeignStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plForeignStatus {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ForeignStatus); err != nil {
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

// GetForeignStatusRequest structure represents the ForeignStatus operation request
type GetForeignStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetForeignStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetForeignStatusOperation) *xxx_GetForeignStatusOperation {
	if op == nil {
		op = &xxx_GetForeignStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetForeignStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetForeignStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetForeignStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetForeignStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetForeignStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetForeignStatusResponse structure represents the ForeignStatus operation response
type GetForeignStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plForeignStatus: A pointer to a long that corresponds to the FOREIGN_STATUS enumeration
	// as defined in the following table.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_STATUS_FOREIGN 0x00000000     | The represented Queue is a foreign queue, or an OutgoingQueue that transfers to  |
	//	|                                  | a foreign queue.                                                                 |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_STATUS_NOT_FOREIGN 0x00000001 | The represented Queue is not a foreign queue, or an OutgoingQueue that transfers |
	//	|                                  | to a foreign queue.                                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| STATUS_UNKNOWN 0x00000002        | The message queuing system is unable to determine whether the represented Queue  |
	//	|                                  | is a foreign queue.                                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	ForeignStatus int32 `idl:"name:plForeignStatus" json:"foreign_status"`
	// Return: The ForeignStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetForeignStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetForeignStatusOperation) *xxx_GetForeignStatusOperation {
	if op == nil {
		op = &xxx_GetForeignStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ForeignStatus = o.ForeignStatus
	op.Return = o.Return
	return op
}

func (o *GetForeignStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetForeignStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ForeignStatus = op.ForeignStatus
	o.Return = op.Return
}
func (o *GetForeignStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetForeignStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetForeignStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQueueTypeOperation structure represents the QueueType operation
type xxx_GetQueueTypeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	QueueType int32          `idl:"name:plQueueType" json:"queue_type"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueueTypeOperation) OpNum() int { return 12 }

func (o *xxx_GetQueueTypeOperation) OpName() string { return "/IMSMQManagement/v0/QueueType" }

func (o *xxx_GetQueueTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQueueTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQueueTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plQueueType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.QueueType); err != nil {
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

func (o *xxx_GetQueueTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plQueueType {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.QueueType); err != nil {
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

// GetQueueTypeRequest structure represents the QueueType operation request
type GetQueueTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQueueTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueueTypeOperation) *xxx_GetQueueTypeOperation {
	if op == nil {
		op = &xxx_GetQueueTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQueueTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueueTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQueueTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueueTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueueTypeResponse structure represents the QueueType operation response
type GetQueueTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plQueueType: A pointer to a long that corresponds to the QUEUE_TYPE enumeration as
	// defined in the following table.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_TYPE_PUBLIC 0x00000000    | The represented Queue.QueueType equals Public.                                   |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_TYPE_PRIVATE 0x00000001   | The represented Queue.QueueType equals Private.                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_TYPE_MACHINE 0x00000002   | The represented Queue.QueueType equals System.                                   |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_TYPE_CONNECTOR 0x00000003 | The represented Queue.QueueType equals Connector.                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_TYPE_MULTICAST 0x00000004 | The represented Queue is an OutgoingQueue with the                               |
	//	|                              | OutgoingQueue.DestinationFormatName specifying a multicast address.              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	QueueType int32 `idl:"name:plQueueType" json:"queue_type"`
	// Return: The QueueType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQueueTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueueTypeOperation) *xxx_GetQueueTypeOperation {
	if op == nil {
		op = &xxx_GetQueueTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QueueType = o.QueueType
	op.Return = o.Return
	return op
}

func (o *GetQueueTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueueTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QueueType = op.QueueType
	o.Return = op.Return
}
func (o *GetQueueTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueueTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsLocalOperation structure represents the IsLocal operation
type xxx_GetIsLocalOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsLocal int16          `idl:"name:pfIsLocal" json:"is_local"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsLocalOperation) OpNum() int { return 13 }

func (o *xxx_GetIsLocalOperation) OpName() string { return "/IMSMQManagement/v0/IsLocal" }

func (o *xxx_GetIsLocalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLocalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsLocalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsLocalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfIsLocal {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsLocal); err != nil {
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

func (o *xxx_GetIsLocalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfIsLocal {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsLocal); err != nil {
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

// GetIsLocalRequest structure represents the IsLocal operation request
type GetIsLocalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsLocalOperation) *xxx_GetIsLocalOperation {
	if op == nil {
		op = &xxx_GetIsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsLocalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsLocalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsLocalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLocalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsLocalResponse structure represents the IsLocal operation response
type GetIsLocalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pfIsLocal: A pointer to a VARIANT_BOOL that, when successfully completed, contains
	// VARIANT_FALSE (0x0000) or VARIANT_TRUE (0xffff) values, depending on whether the
	// represented Queue is an OutgoingQueue.
	IsLocal int16 `idl:"name:pfIsLocal" json:"is_local"`
	// Return: The IsLocal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsLocalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsLocalOperation) *xxx_GetIsLocalOperation {
	if op == nil {
		op = &xxx_GetIsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsLocal = o.IsLocal
	op.Return = o.Return
	return op
}

func (o *GetIsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsLocalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsLocal = op.IsLocal
	o.Return = op.Return
}
func (o *GetIsLocalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsLocalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsLocalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTransactionalStatusOperation structure represents the TransactionalStatus operation
type xxx_GetTransactionalStatusOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	TransactionalStatus int32          `idl:"name:plTransactionalStatus" json:"transactional_status"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTransactionalStatusOperation) OpNum() int { return 14 }

func (o *xxx_GetTransactionalStatusOperation) OpName() string {
	return "/IMSMQManagement/v0/TransactionalStatus"
}

func (o *xxx_GetTransactionalStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionalStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTransactionalStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTransactionalStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTransactionalStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plTransactionalStatus {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.TransactionalStatus); err != nil {
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

func (o *xxx_GetTransactionalStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plTransactionalStatus {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.TransactionalStatus); err != nil {
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

// GetTransactionalStatusRequest structure represents the TransactionalStatus operation request
type GetTransactionalStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTransactionalStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionalStatusOperation) *xxx_GetTransactionalStatusOperation {
	if op == nil {
		op = &xxx_GetTransactionalStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTransactionalStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionalStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTransactionalStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTransactionalStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionalStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTransactionalStatusResponse structure represents the TransactionalStatus operation response
type GetTransactionalStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plTransactionalStatus: A pointer to a long that corresponds to the XACT_STATUS enumeration
	// as defined in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_XACT_STATUS_XACT 0x00000000     | The represented Queue.Transactional equals True.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_XACT_STATUS_NOT_XACT 0x00000001 | The represented Queue.Transactional equals False.                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| MQ_XACT_STATUS_UNKNOWN 0x00000002  | The represented Queue.QueueType equals Public, and the                           |
	//	|                                    | QueueManager.DirectoryIntegrated of the QueueManager that owns the represented   |
	//	|                                    | Queue equals False.                                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	TransactionalStatus int32 `idl:"name:plTransactionalStatus" json:"transactional_status"`
	// Return: The TransactionalStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTransactionalStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTransactionalStatusOperation) *xxx_GetTransactionalStatusOperation {
	if op == nil {
		op = &xxx_GetTransactionalStatusOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TransactionalStatus = o.TransactionalStatus
	op.Return = o.Return
	return op
}

func (o *GetTransactionalStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTransactionalStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TransactionalStatus = op.TransactionalStatus
	o.Return = op.Return
}
func (o *GetTransactionalStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTransactionalStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTransactionalStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBytesInQueueOperation structure represents the BytesInQueue operation
type xxx_GetBytesInQueueOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	BytesInQueue *oaut.Variant  `idl:"name:pvBytesInQueue" json:"bytes_in_queue"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBytesInQueueOperation) OpNum() int { return 15 }

func (o *xxx_GetBytesInQueueOperation) OpName() string { return "/IMSMQManagement/v0/BytesInQueue" }

func (o *xxx_GetBytesInQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBytesInQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBytesInQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBytesInQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvBytesInQueue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.BytesInQueue != nil {
			_ptr_pvBytesInQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.BytesInQueue != nil {
					if err := o.BytesInQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.BytesInQueue, _ptr_pvBytesInQueue); err != nil {
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

func (o *xxx_GetBytesInQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvBytesInQueue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvBytesInQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BytesInQueue == nil {
				o.BytesInQueue = &oaut.Variant{}
			}
			if err := o.BytesInQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvBytesInQueue := func(ptr interface{}) { o.BytesInQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.BytesInQueue, _s_pvBytesInQueue, _ptr_pvBytesInQueue); err != nil {
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

// GetBytesInQueueRequest structure represents the BytesInQueue operation request
type GetBytesInQueueRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBytesInQueueRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInQueueOperation) *xxx_GetBytesInQueueOperation {
	if op == nil {
		op = &xxx_GetBytesInQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBytesInQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInQueueOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBytesInQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBytesInQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBytesInQueueResponse structure represents the BytesInQueue operation response
type GetBytesInQueueResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvBytesInQueue: A pointer to a VARIANT that, when successfully completed, contains
	// an unsigned 64-bit integer (VT_UI8) that specifies, in bytes, the amount of data
	// in the represented Queue.
	BytesInQueue *oaut.Variant `idl:"name:pvBytesInQueue" json:"bytes_in_queue"`
	// Return: The BytesInQueue return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBytesInQueueResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBytesInQueueOperation) *xxx_GetBytesInQueueOperation {
	if op == nil {
		op = &xxx_GetBytesInQueueOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BytesInQueue = o.BytesInQueue
	op.Return = o.Return
	return op
}

func (o *GetBytesInQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBytesInQueueOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BytesInQueue = op.BytesInQueue
	o.Return = op.Return
}
func (o *GetBytesInQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBytesInQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBytesInQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
