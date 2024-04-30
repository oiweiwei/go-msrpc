package iactivation

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// IActivation server interface.
type ActivationServer interface {

	// The RemoteActivation (Opnum 0) method is used by clients to request the activation
	// of an object. It returns the bindings, the IPID for the Remote Unknown, and the COMVERSION
	// of the object exporter that hosts the object.
	RemoteActivation(context.Context, *RemoteActivationRequest) (*RemoteActivationResponse, error)
}

func RegisterActivationServer(conn dcerpc.Conn, o ActivationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActivationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActivationSyntaxV0_0))...)
}

func NewActivationServerHandle(o ActivationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActivationServerHandle(ctx, o, opNum, r)
	}
}

func ActivationServerHandle(ctx context.Context, o ActivationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RemoteActivation
		in := &RemoteActivationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteActivation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
