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
		in := &FreeContextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FreeContext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // AuthzrInitializeContextFromSid
		in := &InitializeContextFromSIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitializeContextFromSID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // AuthzrInitializeCompoundContext
		in := &InitializeCompoundContextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitializeCompoundContext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // AuthzrAccessCheck
		in := &AccessCheckRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AccessCheck(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // AuthzGetInformationFromContext
		in := &GetInformationFromContextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInformationFromContext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AuthzrModifyClaims
		in := &ModifyClaimsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyClaims(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // AuthzrModifySids
		in := &ModifySIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifySIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
