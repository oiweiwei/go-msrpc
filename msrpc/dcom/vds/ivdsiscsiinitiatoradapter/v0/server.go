package ivdsiscsiinitiatoradapter

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

// IVdsIscsiInitiatorAdapter server interface.
type ISCSIInitiatorAdapterServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The QueryInitiatorPortals method returns an object that enumerates the iSCSI initiator
	// portals of the initiator adapter.
	//
	// Return Values: The method MUST return zero or a nonerror HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryInitiatorPortals(context.Context, *QueryInitiatorPortalsRequest) (*QueryInitiatorPortalsResponse, error)

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire

	// Opnum06NotUsedOnWire operation.
	// Opnum06NotUsedOnWire
}

func RegisterISCSIInitiatorAdapterServer(conn dcerpc.Conn, o ISCSIInitiatorAdapterServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewISCSIInitiatorAdapterServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ISCSIInitiatorAdapterSyntaxV0_0))...)
}

func NewISCSIInitiatorAdapterServerHandle(o ISCSIInitiatorAdapterServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ISCSIInitiatorAdapterServerHandle(ctx, o, opNum, r)
	}
}

func ISCSIInitiatorAdapterServerHandle(ctx context.Context, o ISCSIInitiatorAdapterServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 4: // QueryInitiatorPortals
		in := &QueryInitiatorPortalsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInitiatorPortals(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // Opnum05NotUsedOnWire
		// Opnum05NotUsedOnWire
		return nil, nil
	case 6: // Opnum06NotUsedOnWire
		// Opnum06NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
