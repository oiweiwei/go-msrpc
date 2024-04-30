package ivdspack2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
	_ = iunknown.GoPackage
)

// IVdsPack2 server interface.
type Pack2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CreateVolume2 method creates a volume in a disk pack with an optional alignment
	// parameter.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// IVdsPack2::CreateVolume2 has the same sequencing rules as IVdsPack::CreateVolume
	// (Opnum 7), as specified in section 3.4.5.2.19.5.
	CreateVolume2(context.Context, *CreateVolume2Request) (*CreateVolume2Response, error)
}

func RegisterPack2Server(conn dcerpc.Conn, o Pack2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPack2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Pack2SyntaxV0_0))...)
}

func NewPack2ServerHandle(o Pack2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Pack2ServerHandle(ctx, o, opNum, r)
	}
}

func Pack2ServerHandle(ctx context.Context, o Pack2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateVolume2
		in := &CreateVolume2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVolume2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
