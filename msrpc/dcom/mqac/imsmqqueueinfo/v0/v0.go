package imsmqqueueinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQQueueInfo interface identifier d7d6e07b-dccd-11d0-aa4b-0060970debae
	QueueInfoIID = &dcom.IID{Data1: 0xd7d6e07b, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	QueueInfoSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e07b, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	QueueInfoSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QueueInfoSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQueueInfo interface.
type QueueInfoClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The QueueGuid method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the Queue.Identifier that uniquely identifies the referenced queue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetQueueGUID(context.Context, *GetQueueGUIDRequest, ...dcerpc.CallOption) (*GetQueueGUIDResponse, error)

	// ServiceTypeGuid operation.
	GetServiceTypeGUID(context.Context, *GetServiceTypeGUIDRequest, ...dcerpc.CallOption) (*GetServiceTypeGUIDResponse, error)

	// ServiceTypeGuid operation.
	SetServiceTypeGUID(context.Context, *SetServiceTypeGUIDRequest, ...dcerpc.CallOption) (*SetServiceTypeGUIDResponse, error)

	// Label operation.
	GetLabel(context.Context, *GetLabelRequest, ...dcerpc.CallOption) (*GetLabelResponse, error)

	// Label operation.
	SetLabel(context.Context, *SetLabelRequest, ...dcerpc.CallOption) (*SetLabelResponse, error)

	// PathName operation.
	GetPathName(context.Context, *GetPathNameRequest, ...dcerpc.CallOption) (*GetPathNameResponse, error)

	// PathName operation.
	SetPathName(context.Context, *SetPathNameRequest, ...dcerpc.CallOption) (*SetPathNameResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest, ...dcerpc.CallOption) (*GetFormatNameResponse, error)

	// FormatName operation.
	SetFormatName(context.Context, *SetFormatNameRequest, ...dcerpc.CallOption) (*SetFormatNameResponse, error)

	// The IsTransactional method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a value that indicates whether the referenced queue
	// is transactional or nontransactional.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST abide by the following contract:
	//
	// * If IsRefreshed is FALSE call Refresh ( e4072abd-b433-4b22-a4ed-8f654d182705 ) (section
	// 3.10.4.1.29).
	//
	// * If Refresh returns a value other than S_OK (0x00000000), return the HRESULT returned
	// by Refresh and take no further action.
	//
	// * The Queue ( ../ms-mqdmpr/2e026a09-999e-478f-8c4c-5344b661e579 ).Transactional instance
	// variable is True, and ( IsRefreshed is True or IsQueueCreated is True):
	//
	// * Set the pisTransactional output parameter to MQ_TRANSACTIONAL (0x0001).
	//
	// * Else:
	//
	// * Set the pisTransactional output parameter to MQ_TRANSACTIONAL_NONE (0x0000).
	//
	// * Return S_OK (0x00000000), and take no further action.
	GetIsTransactional(context.Context, *GetIsTransactionalRequest, ...dcerpc.CallOption) (*GetIsTransactionalResponse, error)

	// PrivLevel operation.
	GetPrivacyLevel(context.Context, *GetPrivacyLevelRequest, ...dcerpc.CallOption) (*GetPrivacyLevelResponse, error)

	// PrivLevel operation.
	SetPrivacyLevel(context.Context, *SetPrivacyLevelRequest, ...dcerpc.CallOption) (*SetPrivacyLevelResponse, error)

	// Journal operation.
	GetJournal(context.Context, *GetJournalRequest, ...dcerpc.CallOption) (*GetJournalResponse, error)

	// Journal operation.
	SetJournal(context.Context, *SetJournalRequest, ...dcerpc.CallOption) (*SetJournalResponse, error)

	// Quota operation.
	GetQuota(context.Context, *GetQuotaRequest, ...dcerpc.CallOption) (*GetQuotaResponse, error)

	// Quota operation.
	SetQuota(context.Context, *SetQuotaRequest, ...dcerpc.CallOption) (*SetQuotaResponse, error)

	// BasePriority operation.
	GetBasePriority(context.Context, *GetBasePriorityRequest, ...dcerpc.CallOption) (*GetBasePriorityResponse, error)

	// BasePriority operation.
	SetBasePriority(context.Context, *SetBasePriorityRequest, ...dcerpc.CallOption) (*SetBasePriorityResponse, error)

	// The CreateTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the date and time when the referenced queue was created.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetCreateTime(context.Context, *GetCreateTimeRequest, ...dcerpc.CallOption) (*GetCreateTimeResponse, error)

	// The ModifyTime method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the latest date and time when one of the properties of the referenced
	// queue was updated.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetModifyTime(context.Context, *GetModifyTimeRequest, ...dcerpc.CallOption) (*GetModifyTimeResponse, error)

	// Authenticate operation.
	GetAuthenticate(context.Context, *GetAuthenticateRequest, ...dcerpc.CallOption) (*GetAuthenticateResponse, error)

	// Authenticate operation.
	SetAuthenticate(context.Context, *SetAuthenticateRequest, ...dcerpc.CallOption) (*SetAuthenticateResponse, error)

	// JournalQuota operation.
	GetJournalQuota(context.Context, *GetJournalQuotaRequest, ...dcerpc.CallOption) (*GetJournalQuotaResponse, error)

	// JournalQuota operation.
	SetJournalQuota(context.Context, *SetJournalQuotaRequest, ...dcerpc.CallOption) (*SetJournalQuotaResponse, error)

	// The IsWorldReadable method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a BOOLEAN that indicates whether the referenced queue
	// is accessible to everyone, or only to the owner and the system administrators. This
	// can be computed through the security descriptor in the Queue.Security attribute.
	// The owner is the security principal that has MQSEC_TAKE_QUEUE_OWNERSHIP permissions
	// for the Queue, as defined by the security descriptor in the refQueue.Security attribute.
	//
	// Return Values: The method MUST return S_OK (0x00000000).
	GetIsWorldReadable(context.Context, *GetIsWorldReadableRequest, ...dcerpc.CallOption) (*GetIsWorldReadableResponse, error)

	// The Create method is received by the server in an RPC_REQUEST packet. In response,
	// the server creates a new public or private ApplicationQueue.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Create(context.Context, *CreateRequest, ...dcerpc.CallOption) (*CreateResponse, error)

	// The Delete method is received by the server in an RPC_REQUEST packet. In response,
	// the server deletes the referenced queue.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// Open operation.
	Open(context.Context, *OpenRequest, ...dcerpc.CallOption) (*OpenResponse, error)

	// The Refresh method is received by the server in an RPC_REQUEST packet. In response,
	// the server refreshes the properties of the MSMQQueueInfo object with the values stored
	// in the directory (for public queues) or in the local QueueManager (for private queues).
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Refresh(context.Context, *RefreshRequest, ...dcerpc.CallOption) (*RefreshResponse, error)

	// The Update method is received by the server in an RPC_REQUEST packet. In response,
	// the server updates the directory or the local QueueManager with the current values
	// of the MSMQQueueInfo object's properties.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success or an implementation-specific
	// error HRESULT on failure.
	Update(context.Context, *UpdateRequest, ...dcerpc.CallOption) (*UpdateResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QueueInfoClient
}

type xxx_DefaultQueueInfoClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQueueInfoClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQueueInfoClient) GetQueueGUID(ctx context.Context, in *GetQueueGUIDRequest, opts ...dcerpc.CallOption) (*GetQueueGUIDResponse, error) {
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
	out := &GetQueueGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetServiceTypeGUID(ctx context.Context, in *GetServiceTypeGUIDRequest, opts ...dcerpc.CallOption) (*GetServiceTypeGUIDResponse, error) {
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
	out := &GetServiceTypeGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetServiceTypeGUID(ctx context.Context, in *SetServiceTypeGUIDRequest, opts ...dcerpc.CallOption) (*SetServiceTypeGUIDResponse, error) {
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
	out := &SetServiceTypeGUIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...dcerpc.CallOption) (*GetLabelResponse, error) {
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
	out := &GetLabelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetLabel(ctx context.Context, in *SetLabelRequest, opts ...dcerpc.CallOption) (*SetLabelResponse, error) {
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
	out := &SetLabelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetPathName(ctx context.Context, in *GetPathNameRequest, opts ...dcerpc.CallOption) (*GetPathNameResponse, error) {
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
	out := &GetPathNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetPathName(ctx context.Context, in *SetPathNameRequest, opts ...dcerpc.CallOption) (*SetPathNameResponse, error) {
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
	out := &SetPathNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetFormatName(ctx context.Context, in *GetFormatNameRequest, opts ...dcerpc.CallOption) (*GetFormatNameResponse, error) {
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

func (o *xxx_DefaultQueueInfoClient) SetFormatName(ctx context.Context, in *SetFormatNameRequest, opts ...dcerpc.CallOption) (*SetFormatNameResponse, error) {
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
	out := &SetFormatNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetIsTransactional(ctx context.Context, in *GetIsTransactionalRequest, opts ...dcerpc.CallOption) (*GetIsTransactionalResponse, error) {
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
	out := &GetIsTransactionalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetPrivacyLevel(ctx context.Context, in *GetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*GetPrivacyLevelResponse, error) {
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
	out := &GetPrivacyLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetPrivacyLevel(ctx context.Context, in *SetPrivacyLevelRequest, opts ...dcerpc.CallOption) (*SetPrivacyLevelResponse, error) {
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
	out := &SetPrivacyLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetJournal(ctx context.Context, in *GetJournalRequest, opts ...dcerpc.CallOption) (*GetJournalResponse, error) {
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
	out := &GetJournalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetJournal(ctx context.Context, in *SetJournalRequest, opts ...dcerpc.CallOption) (*SetJournalResponse, error) {
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
	out := &SetJournalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetQuota(ctx context.Context, in *GetQuotaRequest, opts ...dcerpc.CallOption) (*GetQuotaResponse, error) {
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
	out := &GetQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetQuota(ctx context.Context, in *SetQuotaRequest, opts ...dcerpc.CallOption) (*SetQuotaResponse, error) {
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
	out := &SetQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetBasePriority(ctx context.Context, in *GetBasePriorityRequest, opts ...dcerpc.CallOption) (*GetBasePriorityResponse, error) {
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
	out := &GetBasePriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetBasePriority(ctx context.Context, in *SetBasePriorityRequest, opts ...dcerpc.CallOption) (*SetBasePriorityResponse, error) {
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
	out := &SetBasePriorityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetCreateTime(ctx context.Context, in *GetCreateTimeRequest, opts ...dcerpc.CallOption) (*GetCreateTimeResponse, error) {
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
	out := &GetCreateTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetModifyTime(ctx context.Context, in *GetModifyTimeRequest, opts ...dcerpc.CallOption) (*GetModifyTimeResponse, error) {
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
	out := &GetModifyTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetAuthenticate(ctx context.Context, in *GetAuthenticateRequest, opts ...dcerpc.CallOption) (*GetAuthenticateResponse, error) {
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
	out := &GetAuthenticateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetAuthenticate(ctx context.Context, in *SetAuthenticateRequest, opts ...dcerpc.CallOption) (*SetAuthenticateResponse, error) {
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
	out := &SetAuthenticateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetJournalQuota(ctx context.Context, in *GetJournalQuotaRequest, opts ...dcerpc.CallOption) (*GetJournalQuotaResponse, error) {
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
	out := &GetJournalQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) SetJournalQuota(ctx context.Context, in *SetJournalQuotaRequest, opts ...dcerpc.CallOption) (*SetJournalQuotaResponse, error) {
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
	out := &SetJournalQuotaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) GetIsWorldReadable(ctx context.Context, in *GetIsWorldReadableRequest, opts ...dcerpc.CallOption) (*GetIsWorldReadableResponse, error) {
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
	out := &GetIsWorldReadableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) Create(ctx context.Context, in *CreateRequest, opts ...dcerpc.CallOption) (*CreateResponse, error) {
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
	out := &CreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
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
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) Open(ctx context.Context, in *OpenRequest, opts ...dcerpc.CallOption) (*OpenResponse, error) {
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
	out := &OpenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...dcerpc.CallOption) (*RefreshResponse, error) {
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
	out := &RefreshResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) Update(ctx context.Context, in *UpdateRequest, opts ...dcerpc.CallOption) (*UpdateResponse, error) {
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
	out := &UpdateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueInfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQueueInfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQueueInfoClient) IPID(ctx context.Context, ipid *dcom.IPID) QueueInfoClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQueueInfoClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQueueInfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QueueInfoClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QueueInfoSyntaxV0_0))...)
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
	return &xxx_DefaultQueueInfoClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetQueueGUIDOperation structure represents the QueueGuid operation
type xxx_GetQueueGUIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Queue  *oaut.String   `idl:"name:pbstrGuidQueue" json:"queue"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueueGUIDOperation) OpNum() int { return 7 }

func (o *xxx_GetQueueGUIDOperation) OpName() string { return "/IMSMQQueueInfo/v0/QueueGuid" }

func (o *xxx_GetQueueGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQueueGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQueueGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrGuidQueue {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Queue != nil {
			_ptr_pbstrGuidQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queue != nil {
					if err := o.Queue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queue, _ptr_pbstrGuidQueue); err != nil {
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

func (o *xxx_GetQueueGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrGuidQueue {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrGuidQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queue == nil {
				o.Queue = &oaut.String{}
			}
			if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrGuidQueue := func(ptr interface{}) { o.Queue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Queue, _s_pbstrGuidQueue, _ptr_pbstrGuidQueue); err != nil {
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

// GetQueueGUIDRequest structure represents the QueueGuid operation request
type GetQueueGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQueueGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueueGUIDOperation) *xxx_GetQueueGUIDOperation {
	if op == nil {
		op = &xxx_GetQueueGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQueueGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueueGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQueueGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueueGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueueGUIDResponse structure represents the QueueGuid operation response
type GetQueueGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrGuidQueue: A pointer to a BSTR that represents a GUID.
	Queue *oaut.String `idl:"name:pbstrGuidQueue" json:"queue"`
	// Return: The QueueGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQueueGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueueGUIDOperation) *xxx_GetQueueGUIDOperation {
	if op == nil {
		op = &xxx_GetQueueGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Queue = o.Queue
	op.Return = o.Return
	return op
}

func (o *GetQueueGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueueGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Queue = op.Queue
	o.Return = op.Return
}
func (o *GetQueueGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueueGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServiceTypeGUIDOperation structure represents the ServiceTypeGuid operation
type xxx_GetServiceTypeGUIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceType *oaut.String   `idl:"name:pbstrGuidServiceType" json:"service_type"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServiceTypeGUIDOperation) OpNum() int { return 8 }

func (o *xxx_GetServiceTypeGUIDOperation) OpName() string {
	return "/IMSMQQueueInfo/v0/ServiceTypeGuid"
}

func (o *xxx_GetServiceTypeGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceTypeGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetServiceTypeGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetServiceTypeGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceTypeGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrGuidServiceType {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceType != nil {
			_ptr_pbstrGuidServiceType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceType != nil {
					if err := o.ServiceType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceType, _ptr_pbstrGuidServiceType); err != nil {
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

func (o *xxx_GetServiceTypeGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrGuidServiceType {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrGuidServiceType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceType == nil {
				o.ServiceType = &oaut.String{}
			}
			if err := o.ServiceType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrGuidServiceType := func(ptr interface{}) { o.ServiceType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceType, _s_pbstrGuidServiceType, _ptr_pbstrGuidServiceType); err != nil {
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

// GetServiceTypeGUIDRequest structure represents the ServiceTypeGuid operation request
type GetServiceTypeGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServiceTypeGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServiceTypeGUIDOperation) *xxx_GetServiceTypeGUIDOperation {
	if op == nil {
		op = &xxx_GetServiceTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServiceTypeGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServiceTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServiceTypeGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServiceTypeGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceTypeGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServiceTypeGUIDResponse structure represents the ServiceTypeGuid operation response
type GetServiceTypeGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceType *oaut.String   `idl:"name:pbstrGuidServiceType" json:"service_type"`
	// Return: The ServiceTypeGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServiceTypeGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServiceTypeGUIDOperation) *xxx_GetServiceTypeGUIDOperation {
	if op == nil {
		op = &xxx_GetServiceTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ServiceType = o.ServiceType
	op.Return = o.Return
	return op
}

func (o *GetServiceTypeGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServiceTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServiceType = op.ServiceType
	o.Return = op.Return
}
func (o *GetServiceTypeGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServiceTypeGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceTypeGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetServiceTypeGUIDOperation structure represents the ServiceTypeGuid operation
type xxx_SetServiceTypeGUIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ServiceType *oaut.String   `idl:"name:bstrGuidServiceType" json:"service_type"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetServiceTypeGUIDOperation) OpNum() int { return 9 }

func (o *xxx_SetServiceTypeGUIDOperation) OpName() string {
	return "/IMSMQQueueInfo/v0/ServiceTypeGuid"
}

func (o *xxx_SetServiceTypeGUIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceTypeGUIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrGuidServiceType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ServiceType != nil {
			_ptr_bstrGuidServiceType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ServiceType != nil {
					if err := o.ServiceType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceType, _ptr_bstrGuidServiceType); err != nil {
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

func (o *xxx_SetServiceTypeGUIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrGuidServiceType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrGuidServiceType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ServiceType == nil {
				o.ServiceType = &oaut.String{}
			}
			if err := o.ServiceType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrGuidServiceType := func(ptr interface{}) { o.ServiceType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ServiceType, _s_bstrGuidServiceType, _ptr_bstrGuidServiceType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceTypeGUIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceTypeGUIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetServiceTypeGUIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetServiceTypeGUIDRequest structure represents the ServiceTypeGuid operation request
type SetServiceTypeGUIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ServiceType *oaut.String   `idl:"name:bstrGuidServiceType" json:"service_type"`
}

func (o *SetServiceTypeGUIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetServiceTypeGUIDOperation) *xxx_SetServiceTypeGUIDOperation {
	if op == nil {
		op = &xxx_SetServiceTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ServiceType = o.ServiceType
	return op
}

func (o *SetServiceTypeGUIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetServiceTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ServiceType = op.ServiceType
}
func (o *SetServiceTypeGUIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetServiceTypeGUIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceTypeGUIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetServiceTypeGUIDResponse structure represents the ServiceTypeGuid operation response
type SetServiceTypeGUIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ServiceTypeGuid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetServiceTypeGUIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetServiceTypeGUIDOperation) *xxx_SetServiceTypeGUIDOperation {
	if op == nil {
		op = &xxx_SetServiceTypeGUIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetServiceTypeGUIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetServiceTypeGUIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetServiceTypeGUIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetServiceTypeGUIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceTypeGUIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLabelOperation structure represents the Label operation
type xxx_GetLabelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label  *oaut.String   `idl:"name:pbstrLabel" json:"label"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLabelOperation) OpNum() int { return 10 }

func (o *xxx_GetLabelOperation) OpName() string { return "/IMSMQQueueInfo/v0/Label" }

func (o *xxx_GetLabelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLabelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLabelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLabelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLabelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrLabel {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Label != nil {
			_ptr_pbstrLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_pbstrLabel); err != nil {
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

func (o *xxx_GetLabelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrLabel {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.String{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrLabel := func(ptr interface{}) { o.Label = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Label, _s_pbstrLabel, _ptr_pbstrLabel); err != nil {
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

// GetLabelRequest structure represents the Label operation request
type GetLabelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLabelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLabelOperation) *xxx_GetLabelOperation {
	if op == nil {
		op = &xxx_GetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLabelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLabelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLabelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLabelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLabelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLabelResponse structure represents the Label operation response
type GetLabelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label *oaut.String   `idl:"name:pbstrLabel" json:"label"`
	// Return: The Label return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLabelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLabelOperation) *xxx_GetLabelOperation {
	if op == nil {
		op = &xxx_GetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Label = o.Label
	op.Return = o.Return
	return op
}

func (o *GetLabelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLabelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Label = op.Label
	o.Return = op.Return
}
func (o *GetLabelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLabelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLabelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLabelOperation structure represents the Label operation
type xxx_SetLabelOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Label  *oaut.String   `idl:"name:bstrLabel" json:"label"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLabelOperation) OpNum() int { return 11 }

func (o *xxx_SetLabelOperation) OpName() string { return "/IMSMQQueueInfo/v0/Label" }

func (o *xxx_SetLabelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrLabel {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Label != nil {
			_ptr_bstrLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Label != nil {
					if err := o.Label.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_bstrLabel); err != nil {
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

func (o *xxx_SetLabelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrLabel {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Label == nil {
				o.Label = &oaut.String{}
			}
			if err := o.Label.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrLabel := func(ptr interface{}) { o.Label = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Label, _s_bstrLabel, _ptr_bstrLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLabelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetLabelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetLabelRequest structure represents the Label operation request
type SetLabelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Label *oaut.String   `idl:"name:bstrLabel" json:"label"`
}

func (o *SetLabelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetLabelOperation) *xxx_SetLabelOperation {
	if op == nil {
		op = &xxx_SetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Label = o.Label
	return op
}

func (o *SetLabelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLabelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Label = op.Label
}
func (o *SetLabelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetLabelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLabelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLabelResponse structure represents the Label operation response
type SetLabelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Label return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLabelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetLabelOperation) *xxx_SetLabelOperation {
	if op == nil {
		op = &xxx_SetLabelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetLabelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLabelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLabelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetLabelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLabelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPathNameOperation structure represents the PathName operation
type xxx_GetPathNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathName *oaut.String   `idl:"name:pbstrPathName" json:"path_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPathNameOperation) OpNum() int { return 12 }

func (o *xxx_GetPathNameOperation) OpName() string { return "/IMSMQQueueInfo/v0/PathName" }

func (o *xxx_GetPathNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPathNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPathNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPathName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PathName != nil {
			_ptr_pbstrPathName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PathName != nil {
					if err := o.PathName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PathName, _ptr_pbstrPathName); err != nil {
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

func (o *xxx_GetPathNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPathName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPathName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PathName == nil {
				o.PathName = &oaut.String{}
			}
			if err := o.PathName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPathName := func(ptr interface{}) { o.PathName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PathName, _s_pbstrPathName, _ptr_pbstrPathName); err != nil {
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

// GetPathNameRequest structure represents the PathName operation request
type GetPathNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPathNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPathNameOperation) *xxx_GetPathNameOperation {
	if op == nil {
		op = &xxx_GetPathNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPathNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPathNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPathNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPathNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPathNameResponse structure represents the PathName operation response
type GetPathNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathName *oaut.String   `idl:"name:pbstrPathName" json:"path_name"`
	// Return: The PathName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPathNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPathNameOperation) *xxx_GetPathNameOperation {
	if op == nil {
		op = &xxx_GetPathNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PathName = o.PathName
	op.Return = o.Return
	return op
}

func (o *GetPathNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPathNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PathName = op.PathName
	o.Return = op.Return
}
func (o *GetPathNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPathNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPathNameOperation structure represents the PathName operation
type xxx_SetPathNameOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathName *oaut.String   `idl:"name:bstrPathName" json:"path_name"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPathNameOperation) OpNum() int { return 13 }

func (o *xxx_SetPathNameOperation) OpName() string { return "/IMSMQQueueInfo/v0/PathName" }

func (o *xxx_SetPathNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPathNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPathName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PathName != nil {
			_ptr_bstrPathName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PathName != nil {
					if err := o.PathName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PathName, _ptr_bstrPathName); err != nil {
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

func (o *xxx_SetPathNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPathName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPathName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PathName == nil {
				o.PathName = &oaut.String{}
			}
			if err := o.PathName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPathName := func(ptr interface{}) { o.PathName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PathName, _s_bstrPathName, _ptr_bstrPathName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPathNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPathNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPathNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPathNameRequest structure represents the PathName operation request
type SetPathNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	PathName *oaut.String   `idl:"name:bstrPathName" json:"path_name"`
}

func (o *SetPathNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPathNameOperation) *xxx_SetPathNameOperation {
	if op == nil {
		op = &xxx_SetPathNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PathName = o.PathName
	return op
}

func (o *SetPathNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPathNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PathName = op.PathName
}
func (o *SetPathNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPathNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPathNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPathNameResponse structure represents the PathName operation response
type SetPathNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PathName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPathNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPathNameOperation) *xxx_SetPathNameOperation {
	if op == nil {
		op = &xxx_SetPathNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPathNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPathNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPathNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPathNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPathNameOperation{}
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

func (o *xxx_GetFormatNameOperation) OpNum() int { return 14 }

func (o *xxx_GetFormatNameOperation) OpName() string { return "/IMSMQQueueInfo/v0/FormatName" }

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

// xxx_SetFormatNameOperation structure represents the FormatName operation
type xxx_SetFormatNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	FormatName *oaut.String   `idl:"name:bstrFormatName" json:"format_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFormatNameOperation) OpNum() int { return 15 }

func (o *xxx_SetFormatNameOperation) OpName() string { return "/IMSMQQueueInfo/v0/FormatName" }

func (o *xxx_SetFormatNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrFormatName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FormatName != nil {
			_ptr_bstrFormatName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.FormatName, _ptr_bstrFormatName); err != nil {
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

func (o *xxx_SetFormatNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrFormatName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrFormatName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FormatName == nil {
				o.FormatName = &oaut.String{}
			}
			if err := o.FormatName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrFormatName := func(ptr interface{}) { o.FormatName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FormatName, _s_bstrFormatName, _ptr_bstrFormatName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFormatNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFormatNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFormatNameRequest structure represents the FormatName operation request
type SetFormatNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	FormatName *oaut.String   `idl:"name:bstrFormatName" json:"format_name"`
}

func (o *SetFormatNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFormatNameOperation) *xxx_SetFormatNameOperation {
	if op == nil {
		op = &xxx_SetFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FormatName = o.FormatName
	return op
}

func (o *SetFormatNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFormatNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FormatName = op.FormatName
}
func (o *SetFormatNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFormatNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFormatNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFormatNameResponse structure represents the FormatName operation response
type SetFormatNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FormatName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFormatNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFormatNameOperation) *xxx_SetFormatNameOperation {
	if op == nil {
		op = &xxx_SetFormatNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFormatNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFormatNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFormatNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFormatNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFormatNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsTransactionalOperation structure represents the IsTransactional operation
type xxx_GetIsTransactionalOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsTransactional int16          `idl:"name:pisTransactional" json:"is_transactional"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsTransactionalOperation) OpNum() int { return 16 }

func (o *xxx_GetIsTransactionalOperation) OpName() string {
	return "/IMSMQQueueInfo/v0/IsTransactional"
}

func (o *xxx_GetIsTransactionalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsTransactionalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsTransactionalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsTransactionalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsTransactionalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisTransactional {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsTransactional); err != nil {
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

func (o *xxx_GetIsTransactionalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisTransactional {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsTransactional); err != nil {
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

// GetIsTransactionalRequest structure represents the IsTransactional operation request
type GetIsTransactionalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsTransactionalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsTransactionalOperation) *xxx_GetIsTransactionalOperation {
	if op == nil {
		op = &xxx_GetIsTransactionalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsTransactionalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsTransactionalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsTransactionalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsTransactionalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsTransactionalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsTransactionalResponse structure represents the IsTransactional operation response
type GetIsTransactionalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisTransactional: A pointer to a short that corresponds to one of the MQTRANSACTIONAL
	// enumeration values.
	IsTransactional int16 `idl:"name:pisTransactional" json:"is_transactional"`
	// Return: The IsTransactional return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsTransactionalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsTransactionalOperation) *xxx_GetIsTransactionalOperation {
	if op == nil {
		op = &xxx_GetIsTransactionalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsTransactional = o.IsTransactional
	op.Return = o.Return
	return op
}

func (o *GetIsTransactionalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsTransactionalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsTransactional = op.IsTransactional
	o.Return = op.Return
}
func (o *GetIsTransactionalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsTransactionalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsTransactionalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPrivacyLevelOperation structure represents the PrivLevel operation
type xxx_GetPrivacyLevelOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:plPrivLevel" json:"privacy_level"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPrivacyLevelOperation) OpNum() int { return 17 }

func (o *xxx_GetPrivacyLevelOperation) OpName() string { return "/IMSMQQueueInfo/v0/PrivLevel" }

func (o *xxx_GetPrivacyLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivacyLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPrivacyLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPrivacyLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPrivacyLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plPrivLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.PrivacyLevel); err != nil {
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

func (o *xxx_GetPrivacyLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plPrivLevel {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.PrivacyLevel); err != nil {
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

// GetPrivacyLevelRequest structure represents the PrivLevel operation request
type GetPrivacyLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPrivacyLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) *xxx_GetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_GetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPrivacyLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPrivacyLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPrivacyLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPrivacyLevelResponse structure represents the PrivLevel operation response
type GetPrivacyLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:plPrivLevel" json:"privacy_level"`
	// Return: The PrivLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPrivacyLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) *xxx_GetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_GetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PrivacyLevel = o.PrivacyLevel
	op.Return = o.Return
	return op
}

func (o *GetPrivacyLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PrivacyLevel = op.PrivacyLevel
	o.Return = op.Return
}
func (o *GetPrivacyLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPrivacyLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPrivacyLevelOperation structure represents the PrivLevel operation
type xxx_SetPrivacyLevelOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	PrivacyLevel int32          `idl:"name:lPrivLevel" json:"privacy_level"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPrivacyLevelOperation) OpNum() int { return 18 }

func (o *xxx_SetPrivacyLevelOperation) OpName() string { return "/IMSMQQueueInfo/v0/PrivLevel" }

func (o *xxx_SetPrivacyLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lPrivLevel {in} (1:(int32))
	{
		if err := w.WriteData(o.PrivacyLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lPrivLevel {in} (1:(int32))
	{
		if err := w.ReadData(&o.PrivacyLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrivacyLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPrivacyLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPrivacyLevelRequest structure represents the PrivLevel operation request
type SetPrivacyLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	PrivacyLevel int32          `idl:"name:lPrivLevel" json:"privacy_level"`
}

func (o *SetPrivacyLevelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) *xxx_SetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_SetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PrivacyLevel = o.PrivacyLevel
	return op
}

func (o *SetPrivacyLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PrivacyLevel = op.PrivacyLevel
}
func (o *SetPrivacyLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPrivacyLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPrivacyLevelResponse structure represents the PrivLevel operation response
type SetPrivacyLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PrivLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPrivacyLevelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) *xxx_SetPrivacyLevelOperation {
	if op == nil {
		op = &xxx_SetPrivacyLevelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPrivacyLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPrivacyLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPrivacyLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPrivacyLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrivacyLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJournalOperation structure represents the Journal operation
type xxx_GetJournalOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:plJournal" json:"journal"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJournalOperation) OpNum() int { return 19 }

func (o *xxx_GetJournalOperation) OpName() string { return "/IMSMQQueueInfo/v0/Journal" }

func (o *xxx_GetJournalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetJournalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetJournalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plJournal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Journal); err != nil {
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

func (o *xxx_GetJournalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plJournal {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Journal); err != nil {
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

// GetJournalRequest structure represents the Journal operation request
type GetJournalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetJournalRequest) xxx_ToOp(ctx context.Context, op *xxx_GetJournalOperation) *xxx_GetJournalOperation {
	if op == nil {
		op = &xxx_GetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetJournalRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJournalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetJournalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetJournalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJournalResponse structure represents the Journal operation response
type GetJournalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:plJournal" json:"journal"`
	// Return: The Journal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetJournalResponse) xxx_ToOp(ctx context.Context, op *xxx_GetJournalOperation) *xxx_GetJournalOperation {
	if op == nil {
		op = &xxx_GetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Journal = o.Journal
	op.Return = o.Return
	return op
}

func (o *GetJournalResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJournalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Journal = op.Journal
	o.Return = op.Return
}
func (o *GetJournalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetJournalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetJournalOperation structure represents the Journal operation
type xxx_SetJournalOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Journal int32          `idl:"name:lJournal" json:"journal"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetJournalOperation) OpNum() int { return 20 }

func (o *xxx_SetJournalOperation) OpName() string { return "/IMSMQQueueInfo/v0/Journal" }

func (o *xxx_SetJournalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lJournal {in} (1:(int32))
	{
		if err := w.WriteData(o.Journal); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lJournal {in} (1:(int32))
	{
		if err := w.ReadData(&o.Journal); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetJournalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetJournalRequest structure represents the Journal operation request
type SetJournalRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Journal int32          `idl:"name:lJournal" json:"journal"`
}

func (o *SetJournalRequest) xxx_ToOp(ctx context.Context, op *xxx_SetJournalOperation) *xxx_SetJournalOperation {
	if op == nil {
		op = &xxx_SetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Journal = o.Journal
	return op
}

func (o *SetJournalRequest) xxx_FromOp(ctx context.Context, op *xxx_SetJournalOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Journal = op.Journal
}
func (o *SetJournalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetJournalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetJournalResponse structure represents the Journal operation response
type SetJournalResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Journal return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetJournalResponse) xxx_ToOp(ctx context.Context, op *xxx_SetJournalOperation) *xxx_SetJournalOperation {
	if op == nil {
		op = &xxx_SetJournalOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetJournalResponse) xxx_FromOp(ctx context.Context, op *xxx_SetJournalOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetJournalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetJournalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuotaOperation structure represents the Quota operation
type xxx_GetQuotaOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Quota  int32          `idl:"name:plQuota" json:"quota"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuotaOperation) OpNum() int { return 21 }

func (o *xxx_GetQuotaOperation) OpName() string { return "/IMSMQQueueInfo/v0/Quota" }

func (o *xxx_GetQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plQuota {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Quota); err != nil {
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

func (o *xxx_GetQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plQuota {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Quota); err != nil {
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

// GetQuotaRequest structure represents the Quota operation request
type GetQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaOperation) *xxx_GetQuotaOperation {
	if op == nil {
		op = &xxx_GetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuotaResponse structure represents the Quota operation response
type GetQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Quota int32          `idl:"name:plQuota" json:"quota"`
	// Return: The Quota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQuotaOperation) *xxx_GetQuotaOperation {
	if op == nil {
		op = &xxx_GetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Quota = o.Quota
	op.Return = o.Return
	return op
}

func (o *GetQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Quota = op.Quota
	o.Return = op.Return
}
func (o *GetQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetQuotaOperation structure represents the Quota operation
type xxx_SetQuotaOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Quota  int32          `idl:"name:lQuota" json:"quota"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetQuotaOperation) OpNum() int { return 22 }

func (o *xxx_SetQuotaOperation) OpName() string { return "/IMSMQQueueInfo/v0/Quota" }

func (o *xxx_SetQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lQuota {in} (1:(int32))
	{
		if err := w.WriteData(o.Quota); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lQuota {in} (1:(int32))
	{
		if err := w.ReadData(&o.Quota); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetQuotaRequest structure represents the Quota operation request
type SetQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Quota int32          `idl:"name:lQuota" json:"quota"`
}

func (o *SetQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaOperation) *xxx_SetQuotaOperation {
	if op == nil {
		op = &xxx_SetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Quota = o.Quota
	return op
}

func (o *SetQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Quota = op.Quota
}
func (o *SetQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetQuotaResponse structure represents the Quota operation response
type SetQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Quota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_SetQuotaOperation) *xxx_SetQuotaOperation {
	if op == nil {
		op = &xxx_SetQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_SetQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBasePriorityOperation structure represents the BasePriority operation
type xxx_GetBasePriorityOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	BasePriority int32          `idl:"name:plBasePriority" json:"base_priority"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBasePriorityOperation) OpNum() int { return 23 }

func (o *xxx_GetBasePriorityOperation) OpName() string { return "/IMSMQQueueInfo/v0/BasePriority" }

func (o *xxx_GetBasePriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBasePriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBasePriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBasePriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBasePriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plBasePriority {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.BasePriority); err != nil {
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

func (o *xxx_GetBasePriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plBasePriority {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.BasePriority); err != nil {
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

// GetBasePriorityRequest structure represents the BasePriority operation request
type GetBasePriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBasePriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBasePriorityOperation) *xxx_GetBasePriorityOperation {
	if op == nil {
		op = &xxx_GetBasePriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBasePriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBasePriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBasePriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBasePriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBasePriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBasePriorityResponse structure represents the BasePriority operation response
type GetBasePriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	BasePriority int32          `idl:"name:plBasePriority" json:"base_priority"`
	// Return: The BasePriority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBasePriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBasePriorityOperation) *xxx_GetBasePriorityOperation {
	if op == nil {
		op = &xxx_GetBasePriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.BasePriority = o.BasePriority
	op.Return = o.Return
	return op
}

func (o *GetBasePriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBasePriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.BasePriority = op.BasePriority
	o.Return = op.Return
}
func (o *GetBasePriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBasePriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBasePriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetBasePriorityOperation structure represents the BasePriority operation
type xxx_SetBasePriorityOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	BasePriority int32          `idl:"name:lBasePriority" json:"base_priority"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetBasePriorityOperation) OpNum() int { return 24 }

func (o *xxx_SetBasePriorityOperation) OpName() string { return "/IMSMQQueueInfo/v0/BasePriority" }

func (o *xxx_SetBasePriorityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBasePriorityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lBasePriority {in} (1:(int32))
	{
		if err := w.WriteData(o.BasePriority); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBasePriorityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lBasePriority {in} (1:(int32))
	{
		if err := w.ReadData(&o.BasePriority); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBasePriorityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetBasePriorityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetBasePriorityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetBasePriorityRequest structure represents the BasePriority operation request
type SetBasePriorityRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	BasePriority int32          `idl:"name:lBasePriority" json:"base_priority"`
}

func (o *SetBasePriorityRequest) xxx_ToOp(ctx context.Context, op *xxx_SetBasePriorityOperation) *xxx_SetBasePriorityOperation {
	if op == nil {
		op = &xxx_SetBasePriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BasePriority = o.BasePriority
	return op
}

func (o *SetBasePriorityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetBasePriorityOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BasePriority = op.BasePriority
}
func (o *SetBasePriorityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetBasePriorityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBasePriorityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetBasePriorityResponse structure represents the BasePriority operation response
type SetBasePriorityResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The BasePriority return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetBasePriorityResponse) xxx_ToOp(ctx context.Context, op *xxx_SetBasePriorityOperation) *xxx_SetBasePriorityOperation {
	if op == nil {
		op = &xxx_SetBasePriorityOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetBasePriorityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetBasePriorityOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetBasePriorityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetBasePriorityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetBasePriorityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCreateTimeOperation structure represents the CreateTime operation
type xxx_GetCreateTimeOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CreateTime *oaut.Variant  `idl:"name:pvarCreateTime" json:"create_time"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCreateTimeOperation) OpNum() int { return 25 }

func (o *xxx_GetCreateTimeOperation) OpName() string { return "/IMSMQQueueInfo/v0/CreateTime" }

func (o *xxx_GetCreateTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCreateTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCreateTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCreateTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCreateTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarCreateTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.CreateTime != nil {
			_ptr_pvarCreateTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CreateTime != nil {
					if err := o.CreateTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CreateTime, _ptr_pvarCreateTime); err != nil {
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

func (o *xxx_GetCreateTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarCreateTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarCreateTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CreateTime == nil {
				o.CreateTime = &oaut.Variant{}
			}
			if err := o.CreateTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarCreateTime := func(ptr interface{}) { o.CreateTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.CreateTime, _s_pvarCreateTime, _ptr_pvarCreateTime); err != nil {
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

// GetCreateTimeRequest structure represents the CreateTime operation request
type GetCreateTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCreateTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCreateTimeOperation) *xxx_GetCreateTimeOperation {
	if op == nil {
		op = &xxx_GetCreateTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCreateTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCreateTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCreateTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCreateTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCreateTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCreateTimeResponse structure represents the CreateTime operation response
type GetCreateTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarCreateTime: A pointer to a VARIANT that contains a UTC date/time (VT_DATE) that
	// specifies the date and time when the referenced queue was created.
	CreateTime *oaut.Variant `idl:"name:pvarCreateTime" json:"create_time"`
	// Return: The CreateTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCreateTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCreateTimeOperation) *xxx_GetCreateTimeOperation {
	if op == nil {
		op = &xxx_GetCreateTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.CreateTime = o.CreateTime
	op.Return = o.Return
	return op
}

func (o *GetCreateTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCreateTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CreateTime = op.CreateTime
	o.Return = op.Return
}
func (o *GetCreateTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCreateTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCreateTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetModifyTimeOperation structure represents the ModifyTime operation
type xxx_GetModifyTimeOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ModifyTime *oaut.Variant  `idl:"name:pvarModifyTime" json:"modify_time"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetModifyTimeOperation) OpNum() int { return 26 }

func (o *xxx_GetModifyTimeOperation) OpName() string { return "/IMSMQQueueInfo/v0/ModifyTime" }

func (o *xxx_GetModifyTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModifyTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetModifyTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetModifyTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModifyTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarModifyTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ModifyTime != nil {
			_ptr_pvarModifyTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModifyTime != nil {
					if err := o.ModifyTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModifyTime, _ptr_pvarModifyTime); err != nil {
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

func (o *xxx_GetModifyTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarModifyTime {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarModifyTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModifyTime == nil {
				o.ModifyTime = &oaut.Variant{}
			}
			if err := o.ModifyTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarModifyTime := func(ptr interface{}) { o.ModifyTime = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ModifyTime, _s_pvarModifyTime, _ptr_pvarModifyTime); err != nil {
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

// GetModifyTimeRequest structure represents the ModifyTime operation request
type GetModifyTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetModifyTimeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetModifyTimeOperation) *xxx_GetModifyTimeOperation {
	if op == nil {
		op = &xxx_GetModifyTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetModifyTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetModifyTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetModifyTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetModifyTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModifyTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetModifyTimeResponse structure represents the ModifyTime operation response
type GetModifyTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarModifyTime: A pointer to a VARIANT that contains a UTC date/time (VT_DATE) that
	// specifies the latest date and time when one of properties of the referenced queue
	// was updated.
	ModifyTime *oaut.Variant `idl:"name:pvarModifyTime" json:"modify_time"`
	// Return: The ModifyTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetModifyTimeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetModifyTimeOperation) *xxx_GetModifyTimeOperation {
	if op == nil {
		op = &xxx_GetModifyTimeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ModifyTime = o.ModifyTime
	op.Return = o.Return
	return op
}

func (o *GetModifyTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetModifyTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModifyTime = op.ModifyTime
	o.Return = op.Return
}
func (o *GetModifyTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetModifyTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModifyTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAuthenticateOperation structure represents the Authenticate operation
type xxx_GetAuthenticateOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authenticate int32          `idl:"name:plAuthenticate" json:"authenticate"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAuthenticateOperation) OpNum() int { return 27 }

func (o *xxx_GetAuthenticateOperation) OpName() string { return "/IMSMQQueueInfo/v0/Authenticate" }

func (o *xxx_GetAuthenticateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAuthenticateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAuthenticateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAuthenticateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAuthenticate {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Authenticate); err != nil {
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

func (o *xxx_GetAuthenticateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAuthenticate {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Authenticate); err != nil {
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

// GetAuthenticateRequest structure represents the Authenticate operation request
type GetAuthenticateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAuthenticateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticateOperation) *xxx_GetAuthenticateOperation {
	if op == nil {
		op = &xxx_GetAuthenticateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAuthenticateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAuthenticateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAuthenticateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAuthenticateResponse structure represents the Authenticate operation response
type GetAuthenticateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authenticate int32          `idl:"name:plAuthenticate" json:"authenticate"`
	// Return: The Authenticate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAuthenticateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAuthenticateOperation) *xxx_GetAuthenticateOperation {
	if op == nil {
		op = &xxx_GetAuthenticateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Authenticate = o.Authenticate
	op.Return = o.Return
	return op
}

func (o *GetAuthenticateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAuthenticateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Authenticate = op.Authenticate
	o.Return = op.Return
}
func (o *GetAuthenticateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAuthenticateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAuthenticateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAuthenticateOperation structure represents the Authenticate operation
type xxx_SetAuthenticateOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	Authenticate int32          `idl:"name:lAuthenticate" json:"authenticate"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAuthenticateOperation) OpNum() int { return 28 }

func (o *xxx_SetAuthenticateOperation) OpName() string { return "/IMSMQQueueInfo/v0/Authenticate" }

func (o *xxx_SetAuthenticateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lAuthenticate {in} (1:(int32))
	{
		if err := w.WriteData(o.Authenticate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lAuthenticate {in} (1:(int32))
	{
		if err := w.ReadData(&o.Authenticate); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAuthenticateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAuthenticateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAuthenticateRequest structure represents the Authenticate operation request
type SetAuthenticateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	Authenticate int32          `idl:"name:lAuthenticate" json:"authenticate"`
}

func (o *SetAuthenticateRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticateOperation) *xxx_SetAuthenticateOperation {
	if op == nil {
		op = &xxx_SetAuthenticateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Authenticate = o.Authenticate
	return op
}

func (o *SetAuthenticateRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Authenticate = op.Authenticate
}
func (o *SetAuthenticateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAuthenticateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAuthenticateResponse structure represents the Authenticate operation response
type SetAuthenticateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Authenticate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAuthenticateResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAuthenticateOperation) *xxx_SetAuthenticateOperation {
	if op == nil {
		op = &xxx_SetAuthenticateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAuthenticateResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAuthenticateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAuthenticateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAuthenticateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAuthenticateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJournalQuotaOperation structure represents the JournalQuota operation
type xxx_GetJournalQuotaOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	JournalQuota int32          `idl:"name:plJournalQuota" json:"journal_quota"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJournalQuotaOperation) OpNum() int { return 29 }

func (o *xxx_GetJournalQuotaOperation) OpName() string { return "/IMSMQQueueInfo/v0/JournalQuota" }

func (o *xxx_GetJournalQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetJournalQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetJournalQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJournalQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plJournalQuota {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.JournalQuota); err != nil {
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

func (o *xxx_GetJournalQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plJournalQuota {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.JournalQuota); err != nil {
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

// GetJournalQuotaRequest structure represents the JournalQuota operation request
type GetJournalQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetJournalQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_GetJournalQuotaOperation) *xxx_GetJournalQuotaOperation {
	if op == nil {
		op = &xxx_GetJournalQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetJournalQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJournalQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetJournalQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetJournalQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJournalQuotaResponse structure represents the JournalQuota operation response
type GetJournalQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	JournalQuota int32          `idl:"name:plJournalQuota" json:"journal_quota"`
	// Return: The JournalQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetJournalQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_GetJournalQuotaOperation) *xxx_GetJournalQuotaOperation {
	if op == nil {
		op = &xxx_GetJournalQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.JournalQuota = o.JournalQuota
	op.Return = o.Return
	return op
}

func (o *GetJournalQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJournalQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.JournalQuota = op.JournalQuota
	o.Return = op.Return
}
func (o *GetJournalQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetJournalQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJournalQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetJournalQuotaOperation structure represents the JournalQuota operation
type xxx_SetJournalQuotaOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	JournalQuota int32          `idl:"name:lJournalQuota" json:"journal_quota"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetJournalQuotaOperation) OpNum() int { return 30 }

func (o *xxx_SetJournalQuotaOperation) OpName() string { return "/IMSMQQueueInfo/v0/JournalQuota" }

func (o *xxx_SetJournalQuotaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalQuotaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lJournalQuota {in} (1:(int32))
	{
		if err := w.WriteData(o.JournalQuota); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalQuotaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lJournalQuota {in} (1:(int32))
	{
		if err := w.ReadData(&o.JournalQuota); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalQuotaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJournalQuotaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetJournalQuotaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetJournalQuotaRequest structure represents the JournalQuota operation request
type SetJournalQuotaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	JournalQuota int32          `idl:"name:lJournalQuota" json:"journal_quota"`
}

func (o *SetJournalQuotaRequest) xxx_ToOp(ctx context.Context, op *xxx_SetJournalQuotaOperation) *xxx_SetJournalQuotaOperation {
	if op == nil {
		op = &xxx_SetJournalQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.JournalQuota = o.JournalQuota
	return op
}

func (o *SetJournalQuotaRequest) xxx_FromOp(ctx context.Context, op *xxx_SetJournalQuotaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.JournalQuota = op.JournalQuota
}
func (o *SetJournalQuotaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetJournalQuotaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalQuotaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetJournalQuotaResponse structure represents the JournalQuota operation response
type SetJournalQuotaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The JournalQuota return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetJournalQuotaResponse) xxx_ToOp(ctx context.Context, op *xxx_SetJournalQuotaOperation) *xxx_SetJournalQuotaOperation {
	if op == nil {
		op = &xxx_SetJournalQuotaOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetJournalQuotaResponse) xxx_FromOp(ctx context.Context, op *xxx_SetJournalQuotaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetJournalQuotaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetJournalQuotaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJournalQuotaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsWorldReadableOperation structure represents the IsWorldReadable operation
type xxx_GetIsWorldReadableOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsWorldReadable int16          `idl:"name:pisWorldReadable" json:"is_world_readable"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsWorldReadableOperation) OpNum() int { return 31 }

func (o *xxx_GetIsWorldReadableOperation) OpName() string {
	return "/IMSMQQueueInfo/v0/IsWorldReadable"
}

func (o *xxx_GetIsWorldReadableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsWorldReadableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsWorldReadableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsWorldReadableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsWorldReadableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisWorldReadable {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsWorldReadable); err != nil {
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

func (o *xxx_GetIsWorldReadableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisWorldReadable {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsWorldReadable); err != nil {
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

// GetIsWorldReadableRequest structure represents the IsWorldReadable operation request
type GetIsWorldReadableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsWorldReadableRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsWorldReadableOperation) *xxx_GetIsWorldReadableOperation {
	if op == nil {
		op = &xxx_GetIsWorldReadableOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsWorldReadableRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsWorldReadableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsWorldReadableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsWorldReadableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsWorldReadableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsWorldReadableResponse structure represents the IsWorldReadable operation response
type GetIsWorldReadableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pisWorldReadable: A pointer to a short that indicates whether the referenced queue
	// is accessible to everyone or only to the owner and the system administrators.
	IsWorldReadable int16 `idl:"name:pisWorldReadable" json:"is_world_readable"`
	// Return: The IsWorldReadable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsWorldReadableResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsWorldReadableOperation) *xxx_GetIsWorldReadableOperation {
	if op == nil {
		op = &xxx_GetIsWorldReadableOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsWorldReadable = o.IsWorldReadable
	op.Return = o.Return
	return op
}

func (o *GetIsWorldReadableResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsWorldReadableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsWorldReadable = op.IsWorldReadable
	o.Return = op.Return
}
func (o *GetIsWorldReadableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsWorldReadableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsWorldReadableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateOperation structure represents the Create operation
type xxx_CreateOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsTransactional *oaut.Variant  `idl:"name:IsTransactional" json:"is_transactional"`
	IsWorldReadable *oaut.Variant  `idl:"name:IsWorldReadable" json:"is_world_readable"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateOperation) OpNum() int { return 32 }

func (o *xxx_CreateOperation) OpName() string { return "/IMSMQQueueInfo/v0/Create" }

func (o *xxx_CreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// IsTransactional {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.IsTransactional != nil {
			_ptr_IsTransactional := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IsTransactional != nil {
					if err := o.IsTransactional.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IsTransactional, _ptr_IsTransactional); err != nil {
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
	// IsWorldReadable {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.IsWorldReadable != nil {
			_ptr_IsWorldReadable := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.IsWorldReadable != nil {
					if err := o.IsWorldReadable.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IsWorldReadable, _ptr_IsWorldReadable); err != nil {
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

func (o *xxx_CreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// IsTransactional {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_IsTransactional := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IsTransactional == nil {
				o.IsTransactional = &oaut.Variant{}
			}
			if err := o.IsTransactional.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_IsTransactional := func(ptr interface{}) { o.IsTransactional = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.IsTransactional, _s_IsTransactional, _ptr_IsTransactional); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// IsWorldReadable {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_IsWorldReadable := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IsWorldReadable == nil {
				o.IsWorldReadable = &oaut.Variant{}
			}
			if err := o.IsWorldReadable.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_IsWorldReadable := func(ptr interface{}) { o.IsWorldReadable = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.IsWorldReadable, _s_IsWorldReadable, _ptr_IsWorldReadable); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateRequest structure represents the Create operation request
type CreateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// IsTransactional: A VARIANT pointer to a BOOLEAN value (VT_BOOL) that specifies whether
	// the queue is transactional. If the value is TRUE (0x00000001), the queue is transactional.
	// If the value is FALSE (0x00000000), the queue is not transactional. If the value
	// is unspecified, the server MUST assume that this value is FALSE.
	IsTransactional *oaut.Variant `idl:"name:IsTransactional" json:"is_transactional"`
	// IsWorldReadable: A VARIANT pointer to a BOOLEAN value (VT_BOOL) that, if set to TRUE
	// (0x00000001), specifies that the queue is accessible to everyone. If the value is
	// not specified, the server MUST use FALSE (0x00000000), which specifies that the queue
	// will be accessible only to the owner and system administrators.
	IsWorldReadable *oaut.Variant `idl:"name:IsWorldReadable" json:"is_world_readable"`
}

func (o *CreateRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateOperation) *xxx_CreateOperation {
	if op == nil {
		op = &xxx_CreateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.IsTransactional = o.IsTransactional
	op.IsWorldReadable = o.IsWorldReadable
	return op
}

func (o *CreateRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IsTransactional = op.IsTransactional
	o.IsWorldReadable = op.IsWorldReadable
}
func (o *CreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResponse structure represents the Create operation response
type CreateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Create return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateOperation) *xxx_CreateOperation {
	if op == nil {
		op = &xxx_CreateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the Delete operation
type xxx_DeleteOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 33 }

func (o *xxx_DeleteOperation) OpName() string { return "/IMSMQQueueInfo/v0/Delete" }

func (o *xxx_DeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteRequest structure represents the Delete operation request
type DeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the Delete operation response
type DeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Delete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenOperation structure represents the Open operation
type xxx_OpenOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Access    int32          `idl:"name:Access" json:"access"`
	ShareMode int32          `idl:"name:ShareMode" json:"share_mode"`
	Queue     *mqac.Queue    `idl:"name:ppq" json:"queue"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenOperation) OpNum() int { return 34 }

func (o *xxx_OpenOperation) OpName() string { return "/IMSMQQueueInfo/v0/Open" }

func (o *xxx_OpenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Access {in} (1:(int32))
	{
		if err := w.WriteData(o.Access); err != nil {
			return err
		}
	}
	// ShareMode {in} (1:(int32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Access {in} (1:(int32))
	{
		if err := w.ReadData(&o.Access); err != nil {
			return err
		}
	}
	// ShareMode {in} (1:(int32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppq {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueue}(interface))
	{
		if o.Queue != nil {
			_ptr_ppq := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Queue != nil {
					if err := o.Queue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Queue{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Queue, _ptr_ppq); err != nil {
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

func (o *xxx_OpenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppq {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueue}(interface))
	{
		_ptr_ppq := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Queue == nil {
				o.Queue = &mqac.Queue{}
			}
			if err := o.Queue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppq := func(ptr interface{}) { o.Queue = *ptr.(**mqac.Queue) }
		if err := w.ReadPointer(&o.Queue, _s_ppq, _ptr_ppq); err != nil {
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

// OpenRequest structure represents the Open operation request
type OpenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Access    int32          `idl:"name:Access" json:"access"`
	ShareMode int32          `idl:"name:ShareMode" json:"share_mode"`
}

func (o *OpenRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenOperation) *xxx_OpenOperation {
	if op == nil {
		op = &xxx_OpenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Access = o.Access
	op.ShareMode = o.ShareMode
	return op
}

func (o *OpenRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Access = op.Access
	o.ShareMode = op.ShareMode
}
func (o *OpenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenResponse structure represents the Open operation response
type OpenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Queue *mqac.Queue    `idl:"name:ppq" json:"queue"`
	// Return: The Open return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenOperation) *xxx_OpenOperation {
	if op == nil {
		op = &xxx_OpenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Queue = o.Queue
	op.Return = o.Return
	return op
}

func (o *OpenResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Queue = op.Queue
	o.Return = op.Return
}
func (o *OpenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RefreshOperation structure represents the Refresh operation
type xxx_RefreshOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshOperation) OpNum() int { return 35 }

func (o *xxx_RefreshOperation) OpName() string { return "/IMSMQQueueInfo/v0/Refresh" }

func (o *xxx_RefreshOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RefreshOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RefreshRequest structure represents the Refresh operation request
type RefreshRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshRequest) xxx_ToOp(ctx context.Context, op *xxx_RefreshOperation) *xxx_RefreshOperation {
	if op == nil {
		op = &xxx_RefreshOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *RefreshRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RefreshRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshResponse structure represents the Refresh operation response
type RefreshResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Refresh return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshResponse) xxx_ToOp(ctx context.Context, op *xxx_RefreshOperation) *xxx_RefreshOperation {
	if op == nil {
		op = &xxx_RefreshOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RefreshResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RefreshResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateOperation structure represents the Update operation
type xxx_UpdateOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateOperation) OpNum() int { return 36 }

func (o *xxx_UpdateOperation) OpName() string { return "/IMSMQQueueInfo/v0/Update" }

func (o *xxx_UpdateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpdateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_UpdateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpdateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// UpdateRequest structure represents the Update operation request
type UpdateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *UpdateRequest) xxx_ToOp(ctx context.Context, op *xxx_UpdateOperation) *xxx_UpdateOperation {
	if op == nil {
		op = &xxx_UpdateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *UpdateRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *UpdateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpdateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateResponse structure represents the Update operation response
type UpdateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Update return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateResponse) xxx_ToOp(ctx context.Context, op *xxx_UpdateOperation) *xxx_UpdateOperation {
	if op == nil {
		op = &xxx_UpdateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *UpdateResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *UpdateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpdateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
