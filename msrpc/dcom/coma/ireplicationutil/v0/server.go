package ireplicationutil

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

// IReplicationUtil server interface.
type ReplicationUtilServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a replication client application to create a Common Internet
	// File System (CIFS) [MS-CIFS] file share for copying installer package files.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateShare(context.Context, *CreateShareRequest) (*CreateShareResponse, error)

	// This method is called by a replication client application to create an empty directory
	// to back up a replication file share.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateEmptyDir(context.Context, *CreateEmptyDirRequest) (*CreateEmptyDirResponse, error)

	// This method is called by a replication client application to remove a file share
	// that was used during replication and is no longer needed.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	RemoveShare(context.Context, *RemoveShareRequest) (*RemoveShareResponse, error)

	// This method is called by a replication client application to request that a server
	// perform the actions necessary to begin a replication operation in which the server
	// is a replication target. This typically happens after the application has copied
	// import package files for the conglomerations to be copied to a replication target
	// file share on the server. As part of the handling of the method, the server sets
	// up a replication working directory and file share. The server's handling of the method
	// might also include management of replication history and backup state.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	BeginReplicationAsTarget(context.Context, *BeginReplicationAsTargetRequest) (*BeginReplicationAsTargetResponse, error)

	// This method is called by a replication client application to obtain the value of
	// the Password property of a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	QueryConglomerationPassword(context.Context, *QueryConglomerationPasswordRequest) (*QueryConglomerationPasswordResponse, error)

	// This method is called by a replication client application to ensure that the server's
	// base replication directory exists and to get its path.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CreateReplicationDir(context.Context, *CreateReplicationDirRequest) (*CreateReplicationDirResponse, error)
}

func RegisterReplicationUtilServer(conn dcerpc.Conn, o ReplicationUtilServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewReplicationUtilServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ReplicationUtilSyntaxV0_0))...)
}

func NewReplicationUtilServerHandle(o ReplicationUtilServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ReplicationUtilServerHandle(ctx, o, opNum, r)
	}
}

func ReplicationUtilServerHandle(ctx context.Context, o ReplicationUtilServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateShare
		in := &CreateShareRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateShare(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // CreateEmptyDir
		in := &CreateEmptyDirRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateEmptyDir(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RemoveShare
		in := &RemoveShareRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveShare(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // BeginReplicationAsTarget
		in := &BeginReplicationAsTargetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BeginReplicationAsTarget(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // QueryConglomerationPassword
		in := &QueryConglomerationPasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryConglomerationPassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // CreateReplicationDir
		in := &CreateReplicationDirRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateReplicationDir(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
