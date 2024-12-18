package iwbemclassobject

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
	GoPackage = "dcom/wmi"
)

var (
	// IWbemClassObject interface identifier dc12a681-737f-11cf-884d-00aa004b2e24
	ClassObjectIID = &dcom.IID{Data1: 0xdc12a681, Data2: 0x737f, Data3: 0x11cf, Data4: []byte{0x88, 0x4d, 0x00, 0xaa, 0x00, 0x4b, 0x2e, 0x24}}
	// Syntax UUID
	ClassObjectSyntaxUUID = &uuid.UUID{TimeLow: 0xdc12a681, TimeMid: 0x737f, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x4d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4b, 0x2e, 0x24}}
	// Syntax ID
	ClassObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClassObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemClassObject interface.
type ClassObjectClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClassObjectClient
}

type xxx_DefaultClassObjectClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClassObjectClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClassObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClassObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClassObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) ClassObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClassObjectClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClassObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClassObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClassObjectSyntaxV0_0))...)
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
	return &xxx_DefaultClassObjectClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}
