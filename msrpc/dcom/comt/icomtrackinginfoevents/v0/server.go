package icomtrackinginfoevents

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

// IComTrackingInfoEvents server interface.
type COMTrackingInfoEventsServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The OnNewTrackingInfo method handles a tracker event from the server.
	//
	// Return Values: The OnNewTrackingInfo method MUST return S_OK (0x00000000) on success
	// and a failure result (as specified in [MS-ERREF] section 2.1) on failure.
	OnNewTrackingInfo(context.Context, *OnNewTrackingInfoRequest) (*OnNewTrackingInfoResponse, error)
}

func RegisterCOMTrackingInfoEventsServer(conn dcerpc.Conn, o COMTrackingInfoEventsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCOMTrackingInfoEventsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(COMTrackingInfoEventsSyntaxV0_0))...)
}

func NewCOMTrackingInfoEventsServerHandle(o COMTrackingInfoEventsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return COMTrackingInfoEventsServerHandle(ctx, o, opNum, r)
	}
}

func COMTrackingInfoEventsServerHandle(ctx context.Context, o COMTrackingInfoEventsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // OnNewTrackingInfo
		in := &OnNewTrackingInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OnNewTrackingInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
