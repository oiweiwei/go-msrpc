package icontainercontrol

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

// IContainerControl server interface.
type ContainerControlServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to create an instance container for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)

	// This method is called by a client to shut down all instance containers for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ShutdownContainers(context.Context, *ShutdownContainersRequest) (*ShutdownContainersResponse, error)

	// This method is called by a client to perform implementation-specific repairs on the
	// server's catalog.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RefreshComponents(context.Context, *RefreshComponentsRequest) (*RefreshComponentsResponse, error)
}

func RegisterContainerControlServer(conn dcerpc.Conn, o ContainerControlServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewContainerControlServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ContainerControlSyntaxV0_0))...)
}

func NewContainerControlServerHandle(o ContainerControlServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ContainerControlServerHandle(ctx, o, opNum, r)
	}
}

func ContainerControlServerHandle(ctx context.Context, o ContainerControlServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateContainer
		op := &xxx_CreateContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ShutdownContainers
		op := &xxx_ShutdownContainersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ShutdownContainersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ShutdownContainers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RefreshComponents
		op := &xxx_RefreshComponentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshComponentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshComponents(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IContainerControl
type UnimplementedContainerControlServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedContainerControlServer) CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControlServer) ShutdownContainers(context.Context, *ShutdownContainersRequest) (*ShutdownContainersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControlServer) RefreshComponents(context.Context, *RefreshComponentsRequest) (*RefreshComponentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ContainerControlServer = (*UnimplementedContainerControlServer)(nil)
