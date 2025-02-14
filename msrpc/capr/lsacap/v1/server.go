package lsacap

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

// lsacap server interface.
type LsacapServer interface {

	// This method returns a list of the CAPIDs of all the central access policies available
	// on the specified remote machine. These identifiers are equivalent to the SIDs of
	// the central access policy objects as they are stored in Active Directory.
	//
	// Return Values:
	//
	// If the method succeeds, the function MUST return 0x00000000 (ERROR_SUCCESS).
	GetAvailableCapIDs(context.Context, *GetAvailableCapIDsRequest) (*GetAvailableCapIDsResponse, error)
}

func RegisterLsacapServer(conn dcerpc.Conn, o LsacapServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLsacapServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LsacapSyntaxV1_0))...)
}

func NewLsacapServerHandle(o LsacapServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LsacapServerHandle(ctx, o, opNum, r)
	}
}

func LsacapServerHandle(ctx context.Context, o LsacapServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // LsarGetAvailableCAPIDs
		op := &xxx_GetAvailableCapIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAvailableCapIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAvailableCapIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented lsacap
type UnimplementedLsacapServer struct {
}

func (UnimplementedLsacapServer) GetAvailableCapIDs(context.Context, *GetAvailableCapIDsRequest) (*GetAvailableCapIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ LsacapServer = (*UnimplementedLsacapServer)(nil)
