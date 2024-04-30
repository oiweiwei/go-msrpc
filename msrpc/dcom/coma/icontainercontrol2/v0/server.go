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
		in := &ShutdownContainerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShutdownContainer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // PauseContainer
		in := &PauseContainerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PauseContainer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ResumeContainer
		in := &ResumeContainerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResumeContainer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // IsContainerPaused
		in := &IsContainerPausedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsContainerPaused(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetRunningContainers
		in := &GetRunningContainersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRunningContainers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetContainerIDFromProcessID
		in := &GetContainerIDFromProcessIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetContainerIDFromProcessID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // RecycleContainer
		in := &RecycleContainerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RecycleContainer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // GetContainerIDFromConglomerationID
		in := &GetContainerIDFromConglomerationIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetContainerIDFromConglomerationID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
