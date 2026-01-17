package icatalogsession

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// ICatalogSession server interface.
type CatalogSessionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// This method is called by a client to perform Catalog Version Negotiation (section
	// 3.1.4.1).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1 on failure. All failure results MUST
	// be treated identically.
	InitializeSession(context.Context, *InitializeSessionRequest) (*InitializeSessionResponse, error)

	// This method is called by a client to perform capability negotiation for the Multiple-partition
	// Support Capability Negotiation (section 3.1.4.3).
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	// A server that does not support catalog version 4.00 or catalog version 5.00 SHOULD
	// immediately return E_NOTIMPL (0x80004001) instead of implementing this method.
	//
	// Otherwise, the server MUST attempt to set the value referenced by plMultiplePartitionSupport
	// to the previously specified value that indicates its support of multiple partitions,
	// and fail the call if it cannot set the value.
	GetServerInformation(context.Context, *GetServerInformationRequest) (*GetServerInformationResponse, error)
}

func RegisterCatalogSessionServer(conn dcerpc.Conn, o CatalogSessionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogSessionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogSessionSyntaxV0_0))...)
}

func NewCatalogSessionServerHandle(o CatalogSessionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogSessionServerHandle(ctx, o, opNum, r)
	}
}

func CatalogSessionServerHandle(ctx context.Context, o CatalogSessionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // InitializeSession
		op := &xxx_InitializeSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetServerInformation
		op := &xxx_GetServerInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICatalogSession
type UnimplementedCatalogSessionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCatalogSessionServer) InitializeSession(context.Context, *InitializeSessionRequest) (*InitializeSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCatalogSessionServer) GetServerInformation(context.Context, *GetServerInformationRequest) (*GetServerInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CatalogSessionServer = (*UnimplementedCatalogSessionServer)(nil)
