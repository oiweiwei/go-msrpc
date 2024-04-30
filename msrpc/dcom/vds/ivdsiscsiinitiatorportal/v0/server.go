package ivdsiscsiinitiatorportal

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

// IVdsIscsiInitiatorPortal server interface.
type ISCSIInitiatorPortalServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The GetInitiatorAdapter method returns the initiator adapter to the initiator portal
	// it belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetInitiatorAdapter(context.Context, *GetInitiatorAdapterRequest) (*GetInitiatorAdapterResponse, error)

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire

	// Opnum06NotUsedOnWire operation.
	// Opnum06NotUsedOnWire

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire
}

func RegisterISCSIInitiatorPortalServer(conn dcerpc.Conn, o ISCSIInitiatorPortalServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewISCSIInitiatorPortalServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ISCSIInitiatorPortalSyntaxV0_0))...)
}

func NewISCSIInitiatorPortalServerHandle(o ISCSIInitiatorPortalServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ISCSIInitiatorPortalServerHandle(ctx, o, opNum, r)
	}
}

func ISCSIInitiatorPortalServerHandle(ctx context.Context, o ISCSIInitiatorPortalServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetProperties
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetInitiatorAdapter
		in := &GetInitiatorAdapterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInitiatorAdapter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum05NotUsedOnWire
		// Opnum05NotUsedOnWire
		return nil, nil
	case 6: // Opnum06NotUsedOnWire
		// Opnum06NotUsedOnWire
		return nil, nil
	case 7: // Opnum07NotUsedOnWire
		// Opnum07NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
