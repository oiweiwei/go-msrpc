package imsmqqueue3

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
	_ = mqac.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQQueue3 interface identifier eba96b1b-2168-11d3-898c-00e02c074f6b
	Queue3IID = &dcom.IID{Data1: 0xeba96b1b, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	Queue3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b1b, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	Queue3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Queue3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQueue3 interface.
type Queue3Client interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Access operation.
	GetAccess(context.Context, *GetAccessRequest, ...dcerpc.CallOption) (*GetAccessResponse, error)

	// ShareMode operation.
	GetShareMode(context.Context, *GetShareModeRequest, ...dcerpc.CallOption) (*GetShareModeResponse, error)

	// QueueInfo operation.
	GetQueueInfo(context.Context, *GetQueueInfoRequest, ...dcerpc.CallOption) (*GetQueueInfoResponse, error)

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest, ...dcerpc.CallOption) (*GetHandleResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest, ...dcerpc.CallOption) (*GetIsOpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

	// Receive_v1 operation.
	ReceiveV1(context.Context, *ReceiveV1Request, ...dcerpc.CallOption) (*ReceiveV1Response, error)

	// Peek_v1 operation.
	PeekV1(context.Context, *PeekV1Request, ...dcerpc.CallOption) (*PeekV1Response, error)

	// EnableNotification operation.
	EnableNotification(context.Context, *EnableNotificationRequest, ...dcerpc.CallOption) (*EnableNotificationResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest, ...dcerpc.CallOption) (*ResetResponse, error)

	// ReceiveCurrent_v1 operation.
	ReceiveCurrentV1(context.Context, *ReceiveCurrentV1Request, ...dcerpc.CallOption) (*ReceiveCurrentV1Response, error)

	// PeekNext_v1 operation.
	PeekNextV1(context.Context, *PeekNextV1Request, ...dcerpc.CallOption) (*PeekNextV1Response, error)

	// PeekCurrent_v1 operation.
	PeekCurrentV1(context.Context, *PeekCurrentV1Request, ...dcerpc.CallOption) (*PeekCurrentV1Response, error)

	// Receive operation.
	Receive(context.Context, *ReceiveRequest, ...dcerpc.CallOption) (*ReceiveResponse, error)

	// Peek operation.
	Peek(context.Context, *PeekRequest, ...dcerpc.CallOption) (*PeekResponse, error)

	// ReceiveCurrent operation.
	ReceiveCurrent(context.Context, *ReceiveCurrentRequest, ...dcerpc.CallOption) (*ReceiveCurrentResponse, error)

	// PeekNext operation.
	PeekNext(context.Context, *PeekNextRequest, ...dcerpc.CallOption) (*PeekNextResponse, error)

	// PeekCurrent operation.
	PeekCurrent(context.Context, *PeekCurrentRequest, ...dcerpc.CallOption) (*PeekCurrentResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// Handle2 operation.
	GetHandle2(context.Context, *GetHandle2Request, ...dcerpc.CallOption) (*GetHandle2Response, error)

	// ReceiveByLookupId operation.
	ReceiveByLookupID(context.Context, *ReceiveByLookupIDRequest, ...dcerpc.CallOption) (*ReceiveByLookupIDResponse, error)

	// ReceiveNextByLookupId operation.
	ReceiveNextByLookupID(context.Context, *ReceiveNextByLookupIDRequest, ...dcerpc.CallOption) (*ReceiveNextByLookupIDResponse, error)

	// ReceivePreviousByLookupId operation.
	ReceivePreviousByLookupID(context.Context, *ReceivePreviousByLookupIDRequest, ...dcerpc.CallOption) (*ReceivePreviousByLookupIDResponse, error)

	// ReceiveFirstByLookupId operation.
	ReceiveFirstByLookupID(context.Context, *ReceiveFirstByLookupIDRequest, ...dcerpc.CallOption) (*ReceiveFirstByLookupIDResponse, error)

	// ReceiveLastByLookupId operation.
	ReceiveLastByLookupID(context.Context, *ReceiveLastByLookupIDRequest, ...dcerpc.CallOption) (*ReceiveLastByLookupIDResponse, error)

	// PeekByLookupId operation.
	PeekByLookupID(context.Context, *PeekByLookupIDRequest, ...dcerpc.CallOption) (*PeekByLookupIDResponse, error)

	// PeekNextByLookupId operation.
	PeekNextByLookupID(context.Context, *PeekNextByLookupIDRequest, ...dcerpc.CallOption) (*PeekNextByLookupIDResponse, error)

	// PeekPreviousByLookupId operation.
	PeekPreviousByLookupID(context.Context, *PeekPreviousByLookupIDRequest, ...dcerpc.CallOption) (*PeekPreviousByLookupIDResponse, error)

	// PeekFirstByLookupId operation.
	PeekFirstByLookupID(context.Context, *PeekFirstByLookupIDRequest, ...dcerpc.CallOption) (*PeekFirstByLookupIDResponse, error)

	// PeekLastByLookupId operation.
	PeekLastByLookupID(context.Context, *PeekLastByLookupIDRequest, ...dcerpc.CallOption) (*PeekLastByLookupIDResponse, error)

	// Purge operation.
	Purge(context.Context, *PurgeRequest, ...dcerpc.CallOption) (*PurgeResponse, error)

	// IsOpen2 operation.
	GetIsOpen2(context.Context, *GetIsOpen2Request, ...dcerpc.CallOption) (*GetIsOpen2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Queue3Client
}

type xxx_DefaultQueue3Client struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQueue3Client) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQueue3Client) GetAccess(ctx context.Context, in *GetAccessRequest, opts ...dcerpc.CallOption) (*GetAccessResponse, error) {
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
	out := &GetAccessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetShareMode(ctx context.Context, in *GetShareModeRequest, opts ...dcerpc.CallOption) (*GetShareModeResponse, error) {
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
	out := &GetShareModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetQueueInfo(ctx context.Context, in *GetQueueInfoRequest, opts ...dcerpc.CallOption) (*GetQueueInfoResponse, error) {
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
	out := &GetQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetHandle(ctx context.Context, in *GetHandleRequest, opts ...dcerpc.CallOption) (*GetHandleResponse, error) {
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
	out := &GetHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetIsOpen(ctx context.Context, in *GetIsOpenRequest, opts ...dcerpc.CallOption) (*GetIsOpenResponse, error) {
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
	out := &GetIsOpenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
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
	out := &CloseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveV1(ctx context.Context, in *ReceiveV1Request, opts ...dcerpc.CallOption) (*ReceiveV1Response, error) {
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
	out := &ReceiveV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekV1(ctx context.Context, in *PeekV1Request, opts ...dcerpc.CallOption) (*PeekV1Response, error) {
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
	out := &PeekV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) EnableNotification(ctx context.Context, in *EnableNotificationRequest, opts ...dcerpc.CallOption) (*EnableNotificationResponse, error) {
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
	out := &EnableNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) Reset(ctx context.Context, in *ResetRequest, opts ...dcerpc.CallOption) (*ResetResponse, error) {
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
	out := &ResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveCurrentV1(ctx context.Context, in *ReceiveCurrentV1Request, opts ...dcerpc.CallOption) (*ReceiveCurrentV1Response, error) {
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
	out := &ReceiveCurrentV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekNextV1(ctx context.Context, in *PeekNextV1Request, opts ...dcerpc.CallOption) (*PeekNextV1Response, error) {
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
	out := &PeekNextV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekCurrentV1(ctx context.Context, in *PeekCurrentV1Request, opts ...dcerpc.CallOption) (*PeekCurrentV1Response, error) {
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
	out := &PeekCurrentV1Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) Receive(ctx context.Context, in *ReceiveRequest, opts ...dcerpc.CallOption) (*ReceiveResponse, error) {
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
	out := &ReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) Peek(ctx context.Context, in *PeekRequest, opts ...dcerpc.CallOption) (*PeekResponse, error) {
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
	out := &PeekResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveCurrent(ctx context.Context, in *ReceiveCurrentRequest, opts ...dcerpc.CallOption) (*ReceiveCurrentResponse, error) {
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
	out := &ReceiveCurrentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekNext(ctx context.Context, in *PeekNextRequest, opts ...dcerpc.CallOption) (*PeekNextResponse, error) {
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
	out := &PeekNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekCurrent(ctx context.Context, in *PeekCurrentRequest, opts ...dcerpc.CallOption) (*PeekCurrentResponse, error) {
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
	out := &PeekCurrentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetHandle2(ctx context.Context, in *GetHandle2Request, opts ...dcerpc.CallOption) (*GetHandle2Response, error) {
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
	out := &GetHandle2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveByLookupID(ctx context.Context, in *ReceiveByLookupIDRequest, opts ...dcerpc.CallOption) (*ReceiveByLookupIDResponse, error) {
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
	out := &ReceiveByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveNextByLookupID(ctx context.Context, in *ReceiveNextByLookupIDRequest, opts ...dcerpc.CallOption) (*ReceiveNextByLookupIDResponse, error) {
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
	out := &ReceiveNextByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceivePreviousByLookupID(ctx context.Context, in *ReceivePreviousByLookupIDRequest, opts ...dcerpc.CallOption) (*ReceivePreviousByLookupIDResponse, error) {
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
	out := &ReceivePreviousByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveFirstByLookupID(ctx context.Context, in *ReceiveFirstByLookupIDRequest, opts ...dcerpc.CallOption) (*ReceiveFirstByLookupIDResponse, error) {
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
	out := &ReceiveFirstByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) ReceiveLastByLookupID(ctx context.Context, in *ReceiveLastByLookupIDRequest, opts ...dcerpc.CallOption) (*ReceiveLastByLookupIDResponse, error) {
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
	out := &ReceiveLastByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekByLookupID(ctx context.Context, in *PeekByLookupIDRequest, opts ...dcerpc.CallOption) (*PeekByLookupIDResponse, error) {
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
	out := &PeekByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekNextByLookupID(ctx context.Context, in *PeekNextByLookupIDRequest, opts ...dcerpc.CallOption) (*PeekNextByLookupIDResponse, error) {
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
	out := &PeekNextByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekPreviousByLookupID(ctx context.Context, in *PeekPreviousByLookupIDRequest, opts ...dcerpc.CallOption) (*PeekPreviousByLookupIDResponse, error) {
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
	out := &PeekPreviousByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekFirstByLookupID(ctx context.Context, in *PeekFirstByLookupIDRequest, opts ...dcerpc.CallOption) (*PeekFirstByLookupIDResponse, error) {
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
	out := &PeekFirstByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) PeekLastByLookupID(ctx context.Context, in *PeekLastByLookupIDRequest, opts ...dcerpc.CallOption) (*PeekLastByLookupIDResponse, error) {
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
	out := &PeekLastByLookupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) Purge(ctx context.Context, in *PurgeRequest, opts ...dcerpc.CallOption) (*PurgeResponse, error) {
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
	out := &PurgeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) GetIsOpen2(ctx context.Context, in *GetIsOpen2Request, opts ...dcerpc.CallOption) (*GetIsOpen2Response, error) {
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
	out := &GetIsOpen2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueue3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQueue3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQueue3Client) IPID(ctx context.Context, ipid *dcom.IPID) Queue3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQueue3Client{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQueue3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Queue3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Queue3SyntaxV0_0))...)
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
	return &xxx_DefaultQueue3Client{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetAccessOperation structure represents the Access operation
type xxx_GetAccessOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Access int32          `idl:"name:plAccess" json:"access"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccessOperation) OpNum() int { return 7 }

func (o *xxx_GetAccessOperation) OpName() string { return "/IMSMQQueue3/v0/Access" }

func (o *xxx_GetAccessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Access); err != nil {
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

func (o *xxx_GetAccessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Access); err != nil {
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

// GetAccessRequest structure represents the Access operation request
type GetAccessRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccessRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccessOperation) *xxx_GetAccessOperation {
	if op == nil {
		op = &xxx_GetAccessOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAccessRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccessOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccessResponse structure represents the Access operation response
type GetAccessResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Access int32          `idl:"name:plAccess" json:"access"`
	// Return: The Access return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccessResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccessOperation) *xxx_GetAccessOperation {
	if op == nil {
		op = &xxx_GetAccessOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Access = o.Access
	op.Return = o.Return
	return op
}

func (o *GetAccessResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccessOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Access = op.Access
	o.Return = op.Return
}
func (o *GetAccessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetShareModeOperation structure represents the ShareMode operation
type xxx_GetShareModeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareMode int32          `idl:"name:plShareMode" json:"share_mode"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetShareModeOperation) OpNum() int { return 8 }

func (o *xxx_GetShareModeOperation) OpName() string { return "/IMSMQQueue3/v0/ShareMode" }

func (o *xxx_GetShareModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetShareModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetShareModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plShareMode {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
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

func (o *xxx_GetShareModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plShareMode {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
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

// GetShareModeRequest structure represents the ShareMode operation request
type GetShareModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetShareModeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetShareModeOperation) *xxx_GetShareModeOperation {
	if op == nil {
		op = &xxx_GetShareModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetShareModeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetShareModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetShareModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetShareModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetShareModeResponse structure represents the ShareMode operation response
type GetShareModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareMode int32          `idl:"name:plShareMode" json:"share_mode"`
	// Return: The ShareMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetShareModeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetShareModeOperation) *xxx_GetShareModeOperation {
	if op == nil {
		op = &xxx_GetShareModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ShareMode = o.ShareMode
	op.Return = o.Return
	return op
}

func (o *GetShareModeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetShareModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ShareMode = op.ShareMode
	o.Return = op.Return
}
func (o *GetShareModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetShareModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQueueInfoOperation structure represents the QueueInfo operation
type xxx_GetQueueInfoOperation struct {
	This      *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfo *mqac.QueueInfo3 `idl:"name:ppqinfo" json:"queue_info"`
	Return    int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueueInfoOperation) OpNum() int { return 9 }

func (o *xxx_GetQueueInfoOperation) OpName() string { return "/IMSMQQueue3/v0/QueueInfo" }

func (o *xxx_GetQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfo {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		if o.QueueInfo != nil {
			_ptr_ppqinfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfo != nil {
					if err := o.QueueInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueInfo, _ptr_ppqinfo); err != nil {
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

func (o *xxx_GetQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfo {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo3}(interface))
	{
		_ptr_ppqinfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfo == nil {
				o.QueueInfo = &mqac.QueueInfo3{}
			}
			if err := o.QueueInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfo := func(ptr interface{}) { o.QueueInfo = *ptr.(**mqac.QueueInfo3) }
		if err := w.ReadPointer(&o.QueueInfo, _s_ppqinfo, _ptr_ppqinfo); err != nil {
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

// GetQueueInfoRequest structure represents the QueueInfo operation request
type GetQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueueInfoOperation) *xxx_GetQueueInfoOperation {
	if op == nil {
		op = &xxx_GetQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueueInfoResponse structure represents the QueueInfo operation response
type GetQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat   `idl:"name:That" json:"that"`
	QueueInfo *mqac.QueueInfo3 `idl:"name:ppqinfo" json:"queue_info"`
	// Return: The QueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueueInfoOperation) *xxx_GetQueueInfoOperation {
	if op == nil {
		op = &xxx_GetQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QueueInfo = o.QueueInfo
	op.Return = o.Return
	return op
}

func (o *GetQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QueueInfo = op.QueueInfo
	o.Return = op.Return
}
func (o *GetQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHandleOperation structure represents the Handle operation
type xxx_GetHandleOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle int32          `idl:"name:plHandle" json:"handle"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHandleOperation) OpNum() int { return 10 }

func (o *xxx_GetHandleOperation) OpName() string { return "/IMSMQQueue3/v0/Handle" }

func (o *xxx_GetHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plHandle {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Handle); err != nil {
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

func (o *xxx_GetHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plHandle {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
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

// GetHandleRequest structure represents the Handle operation request
type GetHandleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHandleResponse structure represents the Handle operation response
type GetHandleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle int32          `idl:"name:plHandle" json:"handle"`
	// Return: The Handle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *GetHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *GetHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsOpenOperation structure represents the IsOpen operation
type xxx_GetIsOpenOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsOpenOperation) OpNum() int { return 11 }

func (o *xxx_GetIsOpenOperation) OpName() string { return "/IMSMQQueue3/v0/IsOpen" }

func (o *xxx_GetIsOpenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsOpenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsOpenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsOpen); err != nil {
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

func (o *xxx_GetIsOpenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsOpen); err != nil {
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

// GetIsOpenRequest structure represents the IsOpen operation request
type GetIsOpenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsOpenRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpenOperation) *xxx_GetIsOpenOperation {
	if op == nil {
		op = &xxx_GetIsOpenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsOpenRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsOpenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsOpenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsOpenResponse structure represents the IsOpen operation response
type GetIsOpenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	// Return: The IsOpen return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsOpenResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpenOperation) *xxx_GetIsOpenOperation {
	if op == nil {
		op = &xxx_GetIsOpenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsOpen = o.IsOpen
	op.Return = o.Return
	return op
}

func (o *GetIsOpenResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsOpen = op.IsOpen
	o.Return = op.Return
}
func (o *GetIsOpenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsOpenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseOperation structure represents the Close operation
type xxx_CloseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 12 }

func (o *xxx_CloseOperation) OpName() string { return "/IMSMQQueue3/v0/Close" }

func (o *xxx_CloseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CloseRequest structure represents the Close operation request
type CloseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResponse structure represents the Close operation response
type CloseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Close return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CloseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveV1Operation structure represents the Receive_v1 operation
type xxx_ReceiveV1Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveV1Operation) OpNum() int { return 13 }

func (o *xxx_ReceiveV1Operation) OpName() string { return "/IMSMQQueue3/v0/Receive_v1" }

func (o *xxx_ReceiveV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_ReceiveV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveV1Request structure represents the Receive_v1 operation request
type ReceiveV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *ReceiveV1Request) xxx_ToOp(ctx context.Context, op *xxx_ReceiveV1Operation) *xxx_ReceiveV1Operation {
	if op == nil {
		op = &xxx_ReceiveV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *ReceiveV1Request) xxx_FromOp(ctx context.Context, op *xxx_ReceiveV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *ReceiveV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveV1Response structure represents the Receive_v1 operation response
type ReceiveV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The Receive_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveV1Response) xxx_ToOp(ctx context.Context, op *xxx_ReceiveV1Operation) *xxx_ReceiveV1Operation {
	if op == nil {
		op = &xxx_ReceiveV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveV1Response) xxx_FromOp(ctx context.Context, op *xxx_ReceiveV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekV1Operation structure represents the Peek_v1 operation
type xxx_PeekV1Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekV1Operation) OpNum() int { return 14 }

func (o *xxx_PeekV1Operation) OpName() string { return "/IMSMQQueue3/v0/Peek_v1" }

func (o *xxx_PeekV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_PeekV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekV1Request structure represents the Peek_v1 operation request
type PeekV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekV1Request) xxx_ToOp(ctx context.Context, op *xxx_PeekV1Operation) *xxx_PeekV1Operation {
	if op == nil {
		op = &xxx_PeekV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekV1Request) xxx_FromOp(ctx context.Context, op *xxx_PeekV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekV1Response structure represents the Peek_v1 operation response
type PeekV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The Peek_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekV1Response) xxx_ToOp(ctx context.Context, op *xxx_PeekV1Operation) *xxx_PeekV1Operation {
	if op == nil {
		op = &xxx_PeekV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekV1Response) xxx_FromOp(ctx context.Context, op *xxx_PeekV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableNotificationOperation structure represents the EnableNotification operation
type xxx_EnableNotificationOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Event          *mqac.Event3   `idl:"name:Event" json:"event"`
	Cursor         *oaut.Variant  `idl:"name:Cursor" json:"cursor"`
	ReceiveTimeout *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableNotificationOperation) OpNum() int { return 15 }

func (o *xxx_EnableNotificationOperation) OpName() string {
	return "/IMSMQQueue3/v0/EnableNotification"
}

func (o *xxx_EnableNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Event {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQEvent3}(interface))
	{
		if o.Event != nil {
			_ptr_Event := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Event != nil {
					if err := o.Event.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Event3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Event, _ptr_Event); err != nil {
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
	// Cursor {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Cursor != nil {
			_ptr_Cursor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Cursor != nil {
					if err := o.Cursor.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Cursor, _ptr_Cursor); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_EnableNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Event {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQEvent3}(interface))
	{
		_ptr_Event := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Event == nil {
				o.Event = &mqac.Event3{}
			}
			if err := o.Event.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Event := func(ptr interface{}) { o.Event = *ptr.(**mqac.Event3) }
		if err := w.ReadPointer(&o.Event, _s_Event, _ptr_Event); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Cursor {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Cursor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Cursor == nil {
				o.Cursor = &oaut.Variant{}
			}
			if err := o.Cursor.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Cursor := func(ptr interface{}) { o.Cursor = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Cursor, _s_Cursor, _ptr_Cursor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnableNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnableNotificationRequest structure represents the EnableNotification operation request
type EnableNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	Event          *mqac.Event3   `idl:"name:Event" json:"event"`
	Cursor         *oaut.Variant  `idl:"name:Cursor" json:"cursor"`
	ReceiveTimeout *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *EnableNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableNotificationOperation) *xxx_EnableNotificationOperation {
	if op == nil {
		op = &xxx_EnableNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Event = o.Event
	op.Cursor = o.Cursor
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *EnableNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Event = op.Event
	o.Cursor = op.Cursor
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *EnableNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableNotificationResponse structure represents the EnableNotification operation response
type EnableNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EnableNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnableNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableNotificationOperation) *xxx_EnableNotificationOperation {
	if op == nil {
		op = &xxx_EnableNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EnableNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EnableNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResetOperation structure represents the Reset operation
type xxx_ResetOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetOperation) OpNum() int { return 16 }

func (o *xxx_ResetOperation) OpName() string { return "/IMSMQQueue3/v0/Reset" }

func (o *xxx_ResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResetRequest structure represents the Reset operation request
type ResetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResetRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetResponse structure represents the Reset operation response
type ResetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reset return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResetResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveCurrentV1Operation structure represents the ReceiveCurrent_v1 operation
type xxx_ReceiveCurrentV1Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveCurrentV1Operation) OpNum() int { return 17 }

func (o *xxx_ReceiveCurrentV1Operation) OpName() string { return "/IMSMQQueue3/v0/ReceiveCurrent_v1" }

func (o *xxx_ReceiveCurrentV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_ReceiveCurrentV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveCurrentV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveCurrentV1Request structure represents the ReceiveCurrent_v1 operation request
type ReceiveCurrentV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *ReceiveCurrentV1Request) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentV1Operation) *xxx_ReceiveCurrentV1Operation {
	if op == nil {
		op = &xxx_ReceiveCurrentV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *ReceiveCurrentV1Request) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *ReceiveCurrentV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveCurrentV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveCurrentV1Response structure represents the ReceiveCurrent_v1 operation response
type ReceiveCurrentV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveCurrent_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveCurrentV1Response) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentV1Operation) *xxx_ReceiveCurrentV1Operation {
	if op == nil {
		op = &xxx_ReceiveCurrentV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveCurrentV1Response) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveCurrentV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveCurrentV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekNextV1Operation structure represents the PeekNext_v1 operation
type xxx_PeekNextV1Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekNextV1Operation) OpNum() int { return 18 }

func (o *xxx_PeekNextV1Operation) OpName() string { return "/IMSMQQueue3/v0/PeekNext_v1" }

func (o *xxx_PeekNextV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_PeekNextV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekNextV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekNextV1Request structure represents the PeekNext_v1 operation request
type PeekNextV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekNextV1Request) xxx_ToOp(ctx context.Context, op *xxx_PeekNextV1Operation) *xxx_PeekNextV1Operation {
	if op == nil {
		op = &xxx_PeekNextV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekNextV1Request) xxx_FromOp(ctx context.Context, op *xxx_PeekNextV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekNextV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekNextV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekNextV1Response structure represents the PeekNext_v1 operation response
type PeekNextV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The PeekNext_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekNextV1Response) xxx_ToOp(ctx context.Context, op *xxx_PeekNextV1Operation) *xxx_PeekNextV1Operation {
	if op == nil {
		op = &xxx_PeekNextV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekNextV1Response) xxx_FromOp(ctx context.Context, op *xxx_PeekNextV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekNextV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekNextV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekCurrentV1Operation structure represents the PeekCurrent_v1 operation
type xxx_PeekCurrentV1Operation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekCurrentV1Operation) OpNum() int { return 19 }

func (o *xxx_PeekCurrentV1Operation) OpName() string { return "/IMSMQQueue3/v0/PeekCurrent_v1" }

func (o *xxx_PeekCurrentV1Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentV1Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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

func (o *xxx_PeekCurrentV1Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentV1Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentV1Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekCurrentV1Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekCurrentV1Request structure represents the PeekCurrent_v1 operation request
type PeekCurrentV1Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekCurrentV1Request) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentV1Operation) *xxx_PeekCurrentV1Operation {
	if op == nil {
		op = &xxx_PeekCurrentV1Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekCurrentV1Request) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentV1Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekCurrentV1Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekCurrentV1Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentV1Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekCurrentV1Response structure represents the PeekCurrent_v1 operation response
type PeekCurrentV1Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The PeekCurrent_v1 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekCurrentV1Response) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentV1Operation) *xxx_PeekCurrentV1Operation {
	if op == nil {
		op = &xxx_PeekCurrentV1Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekCurrentV1Response) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentV1Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekCurrentV1Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekCurrentV1Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentV1Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveOperation structure represents the Receive operation
type xxx_ReceiveOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveOperation) OpNum() int { return 20 }

func (o *xxx_ReceiveOperation) OpName() string { return "/IMSMQQueue3/v0/Receive" }

func (o *xxx_ReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveRequest structure represents the Receive operation request
type ReceiveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveOperation) *xxx_ReceiveOperation {
	if op == nil {
		op = &xxx_ReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveResponse structure represents the Receive operation response
type ReceiveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The Receive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveOperation) *xxx_ReceiveOperation {
	if op == nil {
		op = &xxx_ReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekOperation structure represents the Peek operation
type xxx_PeekOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekOperation) OpNum() int { return 21 }

func (o *xxx_PeekOperation) OpName() string { return "/IMSMQQueue3/v0/Peek" }

func (o *xxx_PeekOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekRequest structure represents the Peek operation request
type PeekRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekOperation) *xxx_PeekOperation {
	if op == nil {
		op = &xxx_PeekOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekResponse structure represents the Peek operation response
type PeekResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The Peek return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekOperation) *xxx_PeekOperation {
	if op == nil {
		op = &xxx_PeekOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveCurrentOperation structure represents the ReceiveCurrent operation
type xxx_ReceiveCurrentOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveCurrentOperation) OpNum() int { return 22 }

func (o *xxx_ReceiveCurrentOperation) OpName() string { return "/IMSMQQueue3/v0/ReceiveCurrent" }

func (o *xxx_ReceiveCurrentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveCurrentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveCurrentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveCurrentRequest structure represents the ReceiveCurrent operation request
type ReceiveCurrentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveCurrentRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) *xxx_ReceiveCurrentOperation {
	if op == nil {
		op = &xxx_ReceiveCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveCurrentRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveCurrentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveCurrentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveCurrentResponse structure represents the ReceiveCurrent operation response
type ReceiveCurrentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveCurrent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveCurrentResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) *xxx_ReceiveCurrentOperation {
	if op == nil {
		op = &xxx_ReceiveCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveCurrentResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveCurrentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveCurrentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekNextOperation structure represents the PeekNext operation
type xxx_PeekNextOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekNextOperation) OpNum() int { return 23 }

func (o *xxx_PeekNextOperation) OpName() string { return "/IMSMQQueue3/v0/PeekNext" }

func (o *xxx_PeekNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekNextRequest structure represents the PeekNext operation request
type PeekNextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekNextRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekNextOperation) *xxx_PeekNextOperation {
	if op == nil {
		op = &xxx_PeekNextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekNextRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekNextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekNextResponse structure represents the PeekNext operation response
type PeekNextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekNext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekNextResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekNextOperation) *xxx_PeekNextOperation {
	if op == nil {
		op = &xxx_PeekNextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekNextResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekNextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekCurrentOperation structure represents the PeekCurrent operation
type xxx_PeekCurrentOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekCurrentOperation) OpNum() int { return 24 }

func (o *xxx_PeekCurrentOperation) OpName() string { return "/IMSMQQueue3/v0/PeekCurrent" }

func (o *xxx_PeekCurrentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekCurrentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekCurrentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekCurrentRequest structure represents the PeekCurrent operation request
type PeekCurrentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekCurrentRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentOperation) *xxx_PeekCurrentOperation {
	if op == nil {
		op = &xxx_PeekCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekCurrentRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekCurrentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekCurrentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekCurrentResponse structure represents the PeekCurrent operation response
type PeekCurrentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekCurrent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekCurrentResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentOperation) *xxx_PeekCurrentOperation {
	if op == nil {
		op = &xxx_PeekCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekCurrentResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekCurrentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekCurrentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the Properties operation
type xxx_GetPropertiesOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 25 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IMSMQQueue3/v0/Properties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.Properties != nil {
			_ptr_ppcolProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Properties != nil {
					if err := o.Properties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Properties, _ptr_ppcolProperties); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppcolProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Properties == nil {
				o.Properties = &oaut.Dispatch{}
			}
			if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppcolProperties := func(ptr interface{}) { o.Properties = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.Properties, _s_ppcolProperties, _ptr_ppcolProperties); err != nil {
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

// GetPropertiesRequest structure represents the Properties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the Properties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	Properties *oaut.Dispatch `idl:"name:ppcolProperties" json:"properties"`
	// Return: The Properties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Properties = o.Properties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Properties = op.Properties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHandle2Operation structure represents the Handle2 operation
type xxx_GetHandle2Operation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle *oaut.Variant  `idl:"name:pvarHandle" json:"handle"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHandle2Operation) OpNum() int { return 26 }

func (o *xxx_GetHandle2Operation) OpName() string { return "/IMSMQQueue3/v0/Handle2" }

func (o *xxx_GetHandle2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandle2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHandle2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHandle2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandle2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarHandle {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Handle != nil {
			_ptr_pvarHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Handle != nil {
					if err := o.Handle.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Handle, _ptr_pvarHandle); err != nil {
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

func (o *xxx_GetHandle2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarHandle {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Handle == nil {
				o.Handle = &oaut.Variant{}
			}
			if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarHandle := func(ptr interface{}) { o.Handle = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Handle, _s_pvarHandle, _ptr_pvarHandle); err != nil {
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

// GetHandle2Request structure represents the Handle2 operation request
type GetHandle2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHandle2Request) xxx_ToOp(ctx context.Context, op *xxx_GetHandle2Operation) *xxx_GetHandle2Operation {
	if op == nil {
		op = &xxx_GetHandle2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHandle2Request) xxx_FromOp(ctx context.Context, op *xxx_GetHandle2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHandle2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHandle2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandle2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHandle2Response structure represents the Handle2 operation response
type GetHandle2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle *oaut.Variant  `idl:"name:pvarHandle" json:"handle"`
	// Return: The Handle2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHandle2Response) xxx_ToOp(ctx context.Context, op *xxx_GetHandle2Operation) *xxx_GetHandle2Operation {
	if op == nil {
		op = &xxx_GetHandle2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *GetHandle2Response) xxx_FromOp(ctx context.Context, op *xxx_GetHandle2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *GetHandle2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHandle2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandle2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveByLookupIDOperation structure represents the ReceiveByLookupId operation
type xxx_ReceiveByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveByLookupIDOperation) OpNum() int { return 27 }

func (o *xxx_ReceiveByLookupIDOperation) OpName() string { return "/IMSMQQueue3/v0/ReceiveByLookupId" }

func (o *xxx_ReceiveByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveByLookupIDRequest structure represents the ReceiveByLookupId operation request
type ReceiveByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveByLookupIDOperation) *xxx_ReceiveByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveByLookupIDResponse structure represents the ReceiveByLookupId operation response
type ReceiveByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveByLookupIDOperation) *xxx_ReceiveByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveNextByLookupIDOperation structure represents the ReceiveNextByLookupId operation
type xxx_ReceiveNextByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveNextByLookupIDOperation) OpNum() int { return 28 }

func (o *xxx_ReceiveNextByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/ReceiveNextByLookupId"
}

func (o *xxx_ReceiveNextByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveNextByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveNextByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveNextByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveNextByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveNextByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveNextByLookupIDRequest structure represents the ReceiveNextByLookupId operation request
type ReceiveNextByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveNextByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveNextByLookupIDOperation) *xxx_ReceiveNextByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveNextByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveNextByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveNextByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveNextByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveNextByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveNextByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveNextByLookupIDResponse structure represents the ReceiveNextByLookupId operation response
type ReceiveNextByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveNextByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveNextByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveNextByLookupIDOperation) *xxx_ReceiveNextByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveNextByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveNextByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveNextByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveNextByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveNextByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveNextByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceivePreviousByLookupIDOperation structure represents the ReceivePreviousByLookupId operation
type xxx_ReceivePreviousByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceivePreviousByLookupIDOperation) OpNum() int { return 29 }

func (o *xxx_ReceivePreviousByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/ReceivePreviousByLookupId"
}

func (o *xxx_ReceivePreviousByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceivePreviousByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceivePreviousByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceivePreviousByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceivePreviousByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceivePreviousByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceivePreviousByLookupIDRequest structure represents the ReceivePreviousByLookupId operation request
type ReceivePreviousByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceivePreviousByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceivePreviousByLookupIDOperation) *xxx_ReceivePreviousByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceivePreviousByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceivePreviousByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceivePreviousByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceivePreviousByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceivePreviousByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceivePreviousByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceivePreviousByLookupIDResponse structure represents the ReceivePreviousByLookupId operation response
type ReceivePreviousByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceivePreviousByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceivePreviousByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceivePreviousByLookupIDOperation) *xxx_ReceivePreviousByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceivePreviousByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceivePreviousByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceivePreviousByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceivePreviousByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceivePreviousByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceivePreviousByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveFirstByLookupIDOperation structure represents the ReceiveFirstByLookupId operation
type xxx_ReceiveFirstByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveFirstByLookupIDOperation) OpNum() int { return 30 }

func (o *xxx_ReceiveFirstByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/ReceiveFirstByLookupId"
}

func (o *xxx_ReceiveFirstByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveFirstByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveFirstByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveFirstByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveFirstByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveFirstByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveFirstByLookupIDRequest structure represents the ReceiveFirstByLookupId operation request
type ReceiveFirstByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveFirstByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveFirstByLookupIDOperation) *xxx_ReceiveFirstByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveFirstByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveFirstByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveFirstByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveFirstByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveFirstByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveFirstByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveFirstByLookupIDResponse structure represents the ReceiveFirstByLookupId operation response
type ReceiveFirstByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveFirstByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveFirstByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveFirstByLookupIDOperation) *xxx_ReceiveFirstByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveFirstByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveFirstByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveFirstByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveFirstByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveFirstByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveFirstByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveLastByLookupIDOperation structure represents the ReceiveLastByLookupId operation
type xxx_ReceiveLastByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveLastByLookupIDOperation) OpNum() int { return 31 }

func (o *xxx_ReceiveLastByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/ReceiveLastByLookupId"
}

func (o *xxx_ReceiveLastByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveLastByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_ReceiveLastByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveLastByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveLastByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveLastByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveLastByLookupIDRequest structure represents the ReceiveLastByLookupId operation request
type ReceiveLastByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *ReceiveLastByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveLastByLookupIDOperation) *xxx_ReceiveLastByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveLastByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *ReceiveLastByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveLastByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *ReceiveLastByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveLastByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveLastByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveLastByLookupIDResponse structure represents the ReceiveLastByLookupId operation response
type ReceiveLastByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveLastByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveLastByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveLastByLookupIDOperation) *xxx_ReceiveLastByLookupIDOperation {
	if op == nil {
		op = &xxx_ReceiveLastByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveLastByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveLastByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveLastByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveLastByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveLastByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekByLookupIDOperation structure represents the PeekByLookupId operation
type xxx_PeekByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekByLookupIDOperation) OpNum() int { return 32 }

func (o *xxx_PeekByLookupIDOperation) OpName() string { return "/IMSMQQueue3/v0/PeekByLookupId" }

func (o *xxx_PeekByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekByLookupIDRequest structure represents the PeekByLookupId operation request
type PeekByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekByLookupIDOperation) *xxx_PeekByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekByLookupIDResponse structure represents the PeekByLookupId operation response
type PeekByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekByLookupIDOperation) *xxx_PeekByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekNextByLookupIDOperation structure represents the PeekNextByLookupId operation
type xxx_PeekNextByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekNextByLookupIDOperation) OpNum() int { return 33 }

func (o *xxx_PeekNextByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/PeekNextByLookupId"
}

func (o *xxx_PeekNextByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekNextByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekNextByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekNextByLookupIDRequest structure represents the PeekNextByLookupId operation request
type PeekNextByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekNextByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekNextByLookupIDOperation) *xxx_PeekNextByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekNextByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekNextByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekNextByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekNextByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekNextByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekNextByLookupIDResponse structure represents the PeekNextByLookupId operation response
type PeekNextByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekNextByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekNextByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekNextByLookupIDOperation) *xxx_PeekNextByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekNextByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekNextByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekNextByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekNextByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekNextByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekPreviousByLookupIDOperation structure represents the PeekPreviousByLookupId operation
type xxx_PeekPreviousByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekPreviousByLookupIDOperation) OpNum() int { return 34 }

func (o *xxx_PeekPreviousByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/PeekPreviousByLookupId"
}

func (o *xxx_PeekPreviousByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekPreviousByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// LookupId {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID != nil {
			if err := o.LookupID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekPreviousByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// LookupId {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.LookupID == nil {
			o.LookupID = &oaut.Variant{}
		}
		if err := o.LookupID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekPreviousByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekPreviousByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekPreviousByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekPreviousByLookupIDRequest structure represents the PeekPreviousByLookupId operation request
type PeekPreviousByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	LookupID             *oaut.Variant  `idl:"name:LookupId" json:"lookup_id"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekPreviousByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekPreviousByLookupIDOperation) *xxx_PeekPreviousByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekPreviousByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LookupID = o.LookupID
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekPreviousByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekPreviousByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LookupID = op.LookupID
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekPreviousByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekPreviousByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekPreviousByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekPreviousByLookupIDResponse structure represents the PeekPreviousByLookupId operation response
type PeekPreviousByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekPreviousByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekPreviousByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekPreviousByLookupIDOperation) *xxx_PeekPreviousByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekPreviousByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekPreviousByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekPreviousByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekPreviousByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekPreviousByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekPreviousByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekFirstByLookupIDOperation structure represents the PeekFirstByLookupId operation
type xxx_PeekFirstByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekFirstByLookupIDOperation) OpNum() int { return 35 }

func (o *xxx_PeekFirstByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/PeekFirstByLookupId"
}

func (o *xxx_PeekFirstByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekFirstByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekFirstByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekFirstByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekFirstByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekFirstByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekFirstByLookupIDRequest structure represents the PeekFirstByLookupId operation request
type PeekFirstByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekFirstByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekFirstByLookupIDOperation) *xxx_PeekFirstByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekFirstByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekFirstByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekFirstByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekFirstByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekFirstByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekFirstByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekFirstByLookupIDResponse structure represents the PeekFirstByLookupId operation response
type PeekFirstByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekFirstByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekFirstByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekFirstByLookupIDOperation) *xxx_PeekFirstByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekFirstByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekFirstByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekFirstByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekFirstByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekFirstByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekFirstByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekLastByLookupIDOperation structure represents the PeekLastByLookupId operation
type xxx_PeekLastByLookupIDOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
	Message              *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekLastByLookupIDOperation) OpNum() int { return 36 }

func (o *xxx_PeekLastByLookupIDOperation) OpName() string {
	return "/IMSMQQueue3/v0/PeekLastByLookupId"
}

func (o *xxx_PeekLastByLookupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekLastByLookupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantConnectorType != nil {
			_ptr_WantConnectorType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantConnectorType != nil {
					if err := o.WantConnectorType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantConnectorType, _ptr_WantConnectorType); err != nil {
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

func (o *xxx_PeekLastByLookupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantConnectorType {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantConnectorType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantConnectorType == nil {
				o.WantConnectorType = &oaut.Variant{}
			}
			if err := o.WantConnectorType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantConnectorType := func(ptr interface{}) { o.WantConnectorType = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantConnectorType, _s_WantConnectorType, _ptr_WantConnectorType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekLastByLookupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekLastByLookupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekLastByLookupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage3}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message3{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message3) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekLastByLookupIDRequest structure represents the PeekLastByLookupId operation request
type PeekLastByLookupIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	WantConnectorType    *oaut.Variant  `idl:"name:WantConnectorType" json:"want_connector_type"`
}

func (o *PeekLastByLookupIDRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekLastByLookupIDOperation) *xxx_PeekLastByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekLastByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.WantConnectorType = o.WantConnectorType
	return op
}

func (o *PeekLastByLookupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekLastByLookupIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.WantConnectorType = op.WantConnectorType
}
func (o *PeekLastByLookupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekLastByLookupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekLastByLookupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekLastByLookupIDResponse structure represents the PeekLastByLookupId operation response
type PeekLastByLookupIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message3 `idl:"name:ppmsg" json:"message"`
	// Return: The PeekLastByLookupId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekLastByLookupIDResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekLastByLookupIDOperation) *xxx_PeekLastByLookupIDOperation {
	if op == nil {
		op = &xxx_PeekLastByLookupIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekLastByLookupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekLastByLookupIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekLastByLookupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekLastByLookupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekLastByLookupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PurgeOperation structure represents the Purge operation
type xxx_PurgeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PurgeOperation) OpNum() int { return 37 }

func (o *xxx_PurgeOperation) OpName() string { return "/IMSMQQueue3/v0/Purge" }

func (o *xxx_PurgeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PurgeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_PurgeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PurgeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PurgeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// PurgeRequest structure represents the Purge operation request
type PurgeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *PurgeRequest) xxx_ToOp(ctx context.Context, op *xxx_PurgeOperation) *xxx_PurgeOperation {
	if op == nil {
		op = &xxx_PurgeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *PurgeRequest) xxx_FromOp(ctx context.Context, op *xxx_PurgeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *PurgeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PurgeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PurgeResponse structure represents the Purge operation response
type PurgeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Purge return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PurgeResponse) xxx_ToOp(ctx context.Context, op *xxx_PurgeOperation) *xxx_PurgeOperation {
	if op == nil {
		op = &xxx_PurgeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *PurgeResponse) xxx_FromOp(ctx context.Context, op *xxx_PurgeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PurgeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PurgeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PurgeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsOpen2Operation structure represents the IsOpen2 operation
type xxx_GetIsOpen2Operation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsOpen2Operation) OpNum() int { return 38 }

func (o *xxx_GetIsOpen2Operation) OpName() string { return "/IMSMQQueue3/v0/IsOpen2" }

func (o *xxx_GetIsOpen2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpen2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsOpen2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsOpen2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpen2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.IsOpen); err != nil {
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

func (o *xxx_GetIsOpen2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.IsOpen); err != nil {
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

// GetIsOpen2Request structure represents the IsOpen2 operation request
type GetIsOpen2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsOpen2Request) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpen2Operation) *xxx_GetIsOpen2Operation {
	if op == nil {
		op = &xxx_GetIsOpen2Operation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsOpen2Request) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpen2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsOpen2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsOpen2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpen2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsOpen2Response structure represents the IsOpen2 operation response
type GetIsOpen2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	// Return: The IsOpen2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsOpen2Response) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpen2Operation) *xxx_GetIsOpen2Operation {
	if op == nil {
		op = &xxx_GetIsOpen2Operation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsOpen = o.IsOpen
	op.Return = o.Return
	return op
}

func (o *GetIsOpen2Response) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpen2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsOpen = op.IsOpen
	o.Return = op.Return
}
func (o *GetIsOpen2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsOpen2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpen2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
