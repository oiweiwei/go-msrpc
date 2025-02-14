package iwbemlevel1login

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

// IWbemLevel1Login server interface.
type Level1LoginServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The IWbemLevel1Login::EstablishPosition method does not perform any action. The return
	// value and output parameter are used in locale negotiation as specified in section
	// 3.2.3.
	//
	// Return Values: The server MUST return one of the following values, based on server
	// behavior for the wszPreferredLocale parameter in IWbemLevel1Login::NTLMLogin.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|        RETURN        |                                                                                  |
	//	|      VALUE/CODE      |                                   DESCRIPTION                                    |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| 0x00 WBEM_S_NO_ERROR | The connection was established and no error occurred.<27>                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| 0x80004001 E_NOTIMPL | The attempted operation is not implemented. The value of this element is as      |
	//	|                      | specified in [MS-ERREF] section 2.1.<28>                                         |
	//	+----------------------+----------------------------------------------------------------------------------+
	EstablishPosition(context.Context, *EstablishPositionRequest) (*EstablishPositionResponse, error)

	// This method does not perform any action.
	RequestChallenge(context.Context, *RequestChallengeRequest) (*RequestChallengeResponse, error)

	// This method does not perform any action.
	WBEMLogin(context.Context, *WBEMLoginRequest) (*WBEMLoginResponse, error)

	// The IWbemLevel1Login::NTLMLogin method MUST connect a user to the management services
	// interface in a specified namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	NTLMLogin(context.Context, *NTLMLoginRequest) (*NTLMLoginResponse, error)
}

func RegisterLevel1LoginServer(conn dcerpc.Conn, o Level1LoginServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLevel1LoginServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Level1LoginSyntaxV0_0))...)
}

func NewLevel1LoginServerHandle(o Level1LoginServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Level1LoginServerHandle(ctx, o, opNum, r)
	}
}

func Level1LoginServerHandle(ctx context.Context, o Level1LoginServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // EstablishPosition
		op := &xxx_EstablishPositionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EstablishPositionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EstablishPosition(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RequestChallenge
		op := &xxx_RequestChallengeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestChallengeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestChallenge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // WBEMLogin
		op := &xxx_WBEMLoginOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WBEMLoginRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WBEMLogin(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // NTLMLogin
		op := &xxx_NTLMLoginOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NTLMLoginRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NTLMLogin(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWbemLevel1Login
type UnimplementedLevel1LoginServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedLevel1LoginServer) EstablishPosition(context.Context, *EstablishPositionRequest) (*EstablishPositionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLevel1LoginServer) RequestChallenge(context.Context, *RequestChallengeRequest) (*RequestChallengeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLevel1LoginServer) WBEMLogin(context.Context, *WBEMLoginRequest) (*WBEMLoginResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLevel1LoginServer) NTLMLogin(context.Context, *NTLMLoginRequest) (*NTLMLoginResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Level1LoginServer = (*UnimplementedLevel1LoginServer)(nil)
