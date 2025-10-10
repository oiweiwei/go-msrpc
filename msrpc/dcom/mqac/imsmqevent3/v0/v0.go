package imsmqevent3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsmqevent2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqevent2/v0"
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
	_ = imsmqevent2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQEvent3 interface identifier eba96b1c-2168-11d3-898c-00e02c074f6b
	ImsmqEvent3IID = &dcom.IID{Data1: 0xeba96b1c, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	ImsmqEvent3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b1c, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	ImsmqEvent3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqEvent3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQEvent3 interface.
type ImsmqEvent3Client interface {

	// IMSMQEvent2 retrieval method.
	ImsmqEvent2() imsmqevent2.ImsmqEvent2Client

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqEvent3Client
}

type xxx_DefaultImsmqEvent3Client struct {
	imsmqevent2.ImsmqEvent2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqEvent3Client) ImsmqEvent2() imsmqevent2.ImsmqEvent2Client {
	return o.ImsmqEvent2Client
}

func (o *xxx_DefaultImsmqEvent3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqEvent3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqEvent3Client) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqEvent3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqEvent3Client{
		ImsmqEvent2Client: o.ImsmqEvent2Client.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewImsmqEvent3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqEvent3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqEvent3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqevent2.NewImsmqEvent2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultImsmqEvent3Client{
		ImsmqEvent2Client: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}
