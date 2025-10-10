package imsmqevent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQEvent interface identifier d7d6e077-dccd-11d0-aa4b-0060970debae
	ImsmqEventIID = &dcom.IID{Data1: 0xd7d6e077, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	ImsmqEventSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e077, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	ImsmqEventSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqEventSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQEvent interface.
type ImsmqEventClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqEventClient
}

type xxx_DefaultImsmqEventClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqEventClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultImsmqEventClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqEventClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqEventClient) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqEventClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqEventClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewImsmqEventClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqEventClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqEventSyntaxV0_0))...)
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
	return &xxx_DefaultImsmqEventClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}
