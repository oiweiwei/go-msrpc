package ivdsasync

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

// IVdsAsync server interface.
type AsyncServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Cancel method cancels the asynchronous operation.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)

	// The Wait method blocks and returns when the asynchronous operation has either finished
	// successfully or failed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)

	// The QueryStatus method retrieves the status of the asynchronous operation.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryStatus(context.Context, *QueryStatusRequest) (*QueryStatusResponse, error)
}

func RegisterAsyncServer(conn dcerpc.Conn, o AsyncServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAsyncServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AsyncSyntaxV0_0))...)
}

func NewAsyncServerHandle(o AsyncServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AsyncServerHandle(ctx, o, opNum, r)
	}
}

func AsyncServerHandle(ctx context.Context, o AsyncServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Cancel
		in := &CancelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Cancel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Wait
		in := &WaitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Wait(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // QueryStatus
		in := &QueryStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
