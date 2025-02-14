package authzr

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

// authzr server interface.
type AuthzrServer interface {

	// The AuthzrFreeContext method (opnum 0) frees all remote structures and memory associated
	// with the client context identified by the ContextHandle parameter.
	//
	// Return Values:
	//
	// If the function succeeds, it MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero 32-bit error code.
	//
	// When a remote authorization server receives this message, it MUST look up the ClientContext
	// structure in the ClientContextTable ADM element and free all structures and memory
	// associated with the ClientContext.
	FreeContext(context.Context, *FreeContextRequest) (*FreeContextResponse, error)

	// The AuthzrInitializeContextFromSid method (opnum 1) creates a client context from
	// a given security identifier (SID). For domain SIDs, token group and claim attributes
	// will be retrieved from Active Directory through Kerberos.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code.
	InitializeContextFromSID(context.Context, *InitializeContextFromSIDRequest) (*InitializeContextFromSIDResponse, error)

	// The AuthzrInitializeCompoundContext method (opnum 2) creates a compound context from
	// two specified context handles.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000. If the function fails,
	// it MUST return a nonzero value.
	InitializeCompoundContext(context.Context, *InitializeCompoundContextRequest) (*InitializeCompoundContextResponse, error)

	// The AuthzrAccessCheck method (opnum 3) determines which access bits can be granted
	// to a client for a given set of security descriptors. The AUTHZR_ACCESS_REPLY structure
	// returns an array of granted access masks and error status.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code.
	AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error)

	// The AuthzGetInformationFromContext method (opnum 4) returns information about the
	// identified client context.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	GetInformationFromContext(context.Context, *GetInformationFromContextRequest) (*GetInformationFromContextResponse, error)

	// The AuthzrModifyClaims method (opnum 5) modifies information about the identified
	// client context.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	ModifyClaims(context.Context, *ModifyClaimsRequest) (*ModifyClaimsResponse, error)

	// The AuthzrModifySids method (opnum 6) modifies the list of SIDs associated with the
	// identified client context.
	//
	// Return Values:
	//
	// If the function succeeds, it MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	ModifySIDs(context.Context, *ModifySIDsRequest) (*ModifySIDsResponse, error)
}

func RegisterAuthzrServer(conn dcerpc.Conn, o AuthzrServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAuthzrServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AuthzrSyntaxV0_0))...)
}

func NewAuthzrServerHandle(o AuthzrServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AuthzrServerHandle(ctx, o, opNum, r)
	}
}

func AuthzrServerHandle(ctx context.Context, o AuthzrServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // AuthzrFreeContext
		op := &xxx_FreeContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FreeContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FreeContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // AuthzrInitializeContextFromSid
		op := &xxx_InitializeContextFromSIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeContextFromSIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeContextFromSID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // AuthzrInitializeCompoundContext
		op := &xxx_InitializeCompoundContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeCompoundContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeCompoundContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // AuthzrAccessCheck
		op := &xxx_AccessCheckOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AccessCheckRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AccessCheck(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // AuthzGetInformationFromContext
		op := &xxx_GetInformationFromContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInformationFromContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInformationFromContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AuthzrModifyClaims
		op := &xxx_ModifyClaimsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyClaimsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyClaims(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // AuthzrModifySids
		op := &xxx_ModifySIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifySIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifySIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented authzr
type UnimplementedAuthzrServer struct {
}

func (UnimplementedAuthzrServer) FreeContext(context.Context, *FreeContextRequest) (*FreeContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) InitializeContextFromSID(context.Context, *InitializeContextFromSIDRequest) (*InitializeContextFromSIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) InitializeCompoundContext(context.Context, *InitializeCompoundContextRequest) (*InitializeCompoundContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) GetInformationFromContext(context.Context, *GetInformationFromContextRequest) (*GetInformationFromContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) ModifyClaims(context.Context, *ModifyClaimsRequest) (*ModifyClaimsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAuthzrServer) ModifySIDs(context.Context, *ModifySIDsRequest) (*ModifySIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AuthzrServer = (*UnimplementedAuthzrServer)(nil)
