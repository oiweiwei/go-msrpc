package iunknown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// IUnknown interface identifier 00000000-0000-0000-c000-000000000046
	UnknownIID = &dcom.IID{Data1: 0x00000000, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	UnknownSyntaxUUID = &uuid.UUID{TimeLow: 0x0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	UnknownSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UnknownSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUnknown interface.
type UnknownClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UnknownClient
}

type xxx_DefaultUnknownClient struct {
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUnknownClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUnknownClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUnknownClient) IPID(ctx context.Context, ipid *dcom.IPID) UnknownClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUnknownClient{
		cc:   o.cc,
		ipid: ipid,
	}
}

func NewUnknownClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UnknownClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UnknownSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	ipid, _ := dcom.HasIPID(opts)
	return &xxx_DefaultUnknownClient{
		cc:   cc,
		ipid: ipid,
	}, nil
}
