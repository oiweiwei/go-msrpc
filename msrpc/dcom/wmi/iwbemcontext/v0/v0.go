package iwbemcontext

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemContext interface identifier 44aca674-e8fc-11d0-a07c-00c04fb68820
	ContextIID = &dcom.IID{Data1: 0x44aca674, Data2: 0xe8fc, Data3: 0x11d0, Data4: []byte{0xa0, 0x7c, 0x00, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax UUID
	ContextSyntaxUUID = &uuid.UUID{TimeLow: 0x44aca674, TimeMid: 0xe8fc, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x7c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax ID
	ContextSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ContextSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemContext interface.
type ContextClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ContextClient
}

type xxx_DefaultContextClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultContextClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultContextClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultContextClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultContextClient) IPID(ctx context.Context, ipid *dcom.IPID) ContextClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultContextClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewContextClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ContextClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ContextSyntaxV0_0))...)
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
	return &xxx_DefaultContextClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}
