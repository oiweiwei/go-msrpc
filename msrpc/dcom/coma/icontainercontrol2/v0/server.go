package icontainercontrol2

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

// IContainerControl2 server interface.
type ContainerControl2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to shut down an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ShutdownContainer(context.Context, *ShutdownContainerRequest) (*ShutdownContainerResponse, error)

	// This method is called by a client to pause an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	PauseContainer(context.Context, *PauseContainerRequest) (*PauseContainerResponse, error)

	// This method is called by a client to resume a paused instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	ResumeContainer(context.Context, *ResumeContainerRequest) (*ResumeContainerResponse, error)

	// This method is called by a client to determine if an instance container is paused.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	IsContainerPaused(context.Context, *IsContainerPausedRequest) (*IsContainerPausedResponse, error)

	// This method is called by a client to get a list of instance containers for a conglomeration,
	// or to get a list of all running instance containers.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetRunningContainers(context.Context, *GetRunningContainersRequest) (*GetRunningContainersResponse, error)

	// This method is called by a client to find the instance container for a process ID.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetContainerIDFromProcessID(context.Context, *GetContainerIDFromProcessIDRequest) (*GetContainerIDFromProcessIDResponse, error)

	// This method is called by a client to recycle an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RecycleContainer(context.Context, *RecycleContainerRequest) (*RecycleContainerResponse, error)

	// This method is called by a client to find the container identifier for the single
	// instance container for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetContainerIDFromConglomerationID(context.Context, *GetContainerIDFromConglomerationIDRequest) (*GetContainerIDFromConglomerationIDResponse, error)
}

func RegisterContainerControl2Server(conn dcerpc.Conn, o ContainerControl2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewContainerControl2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ContainerControl2SyntaxV0_0))...)
}

func NewContainerControl2ServerHandle(o ContainerControl2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ContainerControl2ServerHandle(ctx, o, opNum, r)
	}
}

func ContainerControl2ServerHandle(ctx context.Context, o ContainerControl2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ShutdownContainer
		op := &xxx_ShutdownContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ShutdownContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ShutdownContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // PauseContainer
		op := &xxx_PauseContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PauseContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ResumeContainer
		op := &xxx_ResumeContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResumeContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResumeContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // IsContainerPaused
		op := &xxx_IsContainerPausedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsContainerPausedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsContainerPaused(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetRunningContainers
		op := &xxx_GetRunningContainersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRunningContainersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRunningContainers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetContainerIDFromProcessID
		op := &xxx_GetContainerIDFromProcessIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContainerIDFromProcessIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContainerIDFromProcessID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RecycleContainer
		op := &xxx_RecycleContainerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RecycleContainerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RecycleContainer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetContainerIDFromConglomerationID
		op := &xxx_GetContainerIDFromConglomerationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContainerIDFromConglomerationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContainerIDFromConglomerationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IContainerControl2
type UnimplementedContainerControl2Server struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedContainerControl2Server) ShutdownContainer(context.Context, *ShutdownContainerRequest) (*ShutdownContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) PauseContainer(context.Context, *PauseContainerRequest) (*PauseContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) ResumeContainer(context.Context, *ResumeContainerRequest) (*ResumeContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) IsContainerPaused(context.Context, *IsContainerPausedRequest) (*IsContainerPausedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) GetRunningContainers(context.Context, *GetRunningContainersRequest) (*GetRunningContainersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) GetContainerIDFromProcessID(context.Context, *GetContainerIDFromProcessIDRequest) (*GetContainerIDFromProcessIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) RecycleContainer(context.Context, *RecycleContainerRequest) (*RecycleContainerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedContainerControl2Server) GetContainerIDFromConglomerationID(context.Context, *GetContainerIDFromConglomerationIDRequest) (*GetContainerIDFromConglomerationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ContainerControl2Server = (*UnimplementedContainerControl2Server)(nil)
