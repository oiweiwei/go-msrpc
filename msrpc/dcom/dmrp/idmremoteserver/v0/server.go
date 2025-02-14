package idmremoteserver

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

// IDMRemoteServer server interface.
type IDMRemoteServerServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CreateRemoteObject method creates a disk management server, on the remote machine
	// specified by RemoteComputerName, by invoking DCOM with the class GUID of Disk Management
	// server and the name of the remote machine, which starts the disk management server
	// on the remote machine. The method negotiates for the interface as described in section
	// 3.1.3, and as illustrated in section 4. The client holds a reference to the IDMRemoteServer
	// interface binding on the server, until the client has received an IVolumeClient,
	// or IVolumeClient3 interface binding to the remote server. The client MAY then release
	// the IDMRemoteServer interface on the server.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	CreateRemoteObject(context.Context, *CreateRemoteObjectRequest) (*CreateRemoteObjectResponse, error)
}

func RegisterIDMRemoteServerServer(conn dcerpc.Conn, o IDMRemoteServerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIDMRemoteServerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IDMRemoteServerSyntaxV0_0))...)
}

func NewIDMRemoteServerServerHandle(o IDMRemoteServerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IDMRemoteServerServerHandle(ctx, o, opNum, r)
	}
}

func IDMRemoteServerServerHandle(ctx context.Context, o IDMRemoteServerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateRemoteObject
		op := &xxx_CreateRemoteObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRemoteObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateRemoteObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IDMRemoteServer
type UnimplementedIDMRemoteServerServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedIDMRemoteServerServer) CreateRemoteObject(context.Context, *CreateRemoteObjectRequest) (*CreateRemoteObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IDMRemoteServerServer = (*UnimplementedIDMRemoteServerServer)(nil)
