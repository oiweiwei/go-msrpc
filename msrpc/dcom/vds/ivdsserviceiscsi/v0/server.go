package ivdsserviceiscsi

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

// IVdsServiceIscsi server interface.
type ServiceISCSIServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetInitiatorName method returns the iSCSI name of the initiator service.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetInitiatorName(context.Context, *GetInitiatorNameRequest) (*GetInitiatorNameResponse, error)

	// The QueryInitiatorAdapters method returns an object that enumerates the iSCSI initiator
	// adapters of the initiator.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryInitiatorAdapters(context.Context, *QueryInitiatorAdaptersRequest) (*QueryInitiatorAdaptersResponse, error)

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire

	// Opnum06NotUsedOnWire operation.
	// Opnum06NotUsedOnWire

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire

	// The SetInitiatorSharedSecret method sets the initiator CHAP shared secret that is
	// used for mutual CHAP authentication when the initiator authenticates the target.
	// For more information on CHAP, see [MS-CHAP].<74>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetInitiatorSharedSecret(context.Context, *SetInitiatorSharedSecretRequest) (*SetInitiatorSharedSecretResponse, error)

	// Opnum09NotUsedOnWire operation.
	// Opnum09NotUsedOnWire
}

func RegisterServiceISCSIServer(conn dcerpc.Conn, o ServiceISCSIServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceISCSIServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceISCSISyntaxV0_0))...)
}

func NewServiceISCSIServerHandle(o ServiceISCSIServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceISCSIServerHandle(ctx, o, opNum, r)
	}
}

func ServiceISCSIServerHandle(ctx context.Context, o ServiceISCSIServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetInitiatorName
		op := &xxx_GetInitiatorNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInitiatorNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInitiatorName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // QueryInitiatorAdapters
		op := &xxx_QueryInitiatorAdaptersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInitiatorAdaptersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInitiatorAdapters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Opnum05NotUsedOnWire
		// Opnum05NotUsedOnWire
		return nil, nil
	case 6: // Opnum06NotUsedOnWire
		// Opnum06NotUsedOnWire
		return nil, nil
	case 7: // Opnum07NotUsedOnWire
		// Opnum07NotUsedOnWire
		return nil, nil
	case 8: // SetInitiatorSharedSecret
		op := &xxx_SetInitiatorSharedSecretOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInitiatorSharedSecretRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInitiatorSharedSecret(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Opnum09NotUsedOnWire
		// Opnum09NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IVdsServiceIscsi
type UnimplementedServiceISCSIServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServiceISCSIServer) GetInitiatorName(context.Context, *GetInitiatorNameRequest) (*GetInitiatorNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceISCSIServer) QueryInitiatorAdapters(context.Context, *QueryInitiatorAdaptersRequest) (*QueryInitiatorAdaptersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServiceISCSIServer) SetInitiatorSharedSecret(context.Context, *SetInitiatorSharedSecretRequest) (*SetInitiatorSharedSecretResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServiceISCSIServer = (*UnimplementedServiceISCSIServer)(nil)
