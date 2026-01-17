package iinstallationbehavior

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IInstallationBehavior server interface.
type InstallationBehaviorServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IInstallationBehavior::CanRequestUserInput (opnum 8) method retrieves whether
	// the operation can request user input.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the CanRequestUserInput ADM element.
	GetCanRequestUserInput(context.Context, *GetCanRequestUserInputRequest) (*GetCanRequestUserInputResponse, error)

	// The IInstallationBehavior::Impact (opnum 9) method retrieves an enumeration value
	// that describes the impact of the installation or uninstallation operation on the
	// computer.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the Impact ADM element.
	GetImpact(context.Context, *GetImpactRequest) (*GetImpactResponse, error)

	// The IInstallationBehavior::RebootBehavior (opnum 10) method retrieves an enumeration
	// value that describes the likelihood that a reboot is needed for this operation.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the RebootBehavior ADM element.
	GetRebootBehavior(context.Context, *GetRebootBehaviorRequest) (*GetRebootBehaviorResponse, error)

	// The IInstallationBehavior::RequiresNetworkConnectivity (opnum 11) method retrieves
	// whether the operation can require network connectivity.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the RequiresNetworkConnectivity ADM element.
	GetRequiresNetworkConnectivity(context.Context, *GetRequiresNetworkConnectivityRequest) (*GetRequiresNetworkConnectivityResponse, error)
}

func RegisterInstallationBehaviorServer(conn dcerpc.Conn, o InstallationBehaviorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewInstallationBehaviorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(InstallationBehaviorSyntaxV0_0))...)
}

func NewInstallationBehaviorServerHandle(o InstallationBehaviorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return InstallationBehaviorServerHandle(ctx, o, opNum, r)
	}
}

func InstallationBehaviorServerHandle(ctx context.Context, o InstallationBehaviorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CanRequestUserInput
		op := &xxx_GetCanRequestUserInputOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanRequestUserInputRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanRequestUserInput(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Impact
		op := &xxx_GetImpactOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImpactRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImpact(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RebootBehavior
		op := &xxx_GetRebootBehaviorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRebootBehaviorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRebootBehavior(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RequiresNetworkConnectivity
		op := &xxx_GetRequiresNetworkConnectivityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRequiresNetworkConnectivityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRequiresNetworkConnectivity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IInstallationBehavior
type UnimplementedInstallationBehaviorServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedInstallationBehaviorServer) GetCanRequestUserInput(context.Context, *GetCanRequestUserInputRequest) (*GetCanRequestUserInputResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInstallationBehaviorServer) GetImpact(context.Context, *GetImpactRequest) (*GetImpactResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInstallationBehaviorServer) GetRebootBehavior(context.Context, *GetRebootBehaviorRequest) (*GetRebootBehaviorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedInstallationBehaviorServer) GetRequiresNetworkConnectivity(context.Context, *GetRequiresNetworkConnectivityRequest) (*GetRequiresNetworkConnectivityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ InstallationBehaviorServer = (*UnimplementedInstallationBehaviorServer)(nil)
