package iwbemfetchsmartenum

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

// IWbemFetchSmartEnum server interface.
type FetchSmartEnumServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemFetchSmartEnum::GetSmartEnum method retrieves an IWbemWCOSmartEnum (section
	// 3.1.4.7) interface, which is a network-optimized enumerator interface.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	GetSmartEnum(context.Context, *GetSmartEnumRequest) (*GetSmartEnumResponse, error)
}

func RegisterFetchSmartEnumServer(conn dcerpc.Conn, o FetchSmartEnumServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFetchSmartEnumServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FetchSmartEnumSyntaxV0_0))...)
}

func NewFetchSmartEnumServerHandle(o FetchSmartEnumServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FetchSmartEnumServerHandle(ctx, o, opNum, r)
	}
}

func FetchSmartEnumServerHandle(ctx context.Context, o FetchSmartEnumServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSmartEnum
		in := &GetSmartEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSmartEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
