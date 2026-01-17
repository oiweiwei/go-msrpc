package icatalog64bitsupport

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

// ICatalog64BitSupport server interface.
type Catalog64BitSupportServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to perform capability negotiation for the Multiple-Bitness
	// Capability (section 3.1.4.4).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	SupportsMultipleBitness(context.Context, *SupportsMultipleBitnessRequest) (*SupportsMultipleBitnessResponse, error)

	// This method is called by a client to perform capability negotiation for the 64-bit
	// QueryCell Marshaling Format Capability (section 3.1.4.2).
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	Initialize64BitQueryCellSupport(context.Context, *Initialize64BitQueryCellSupportRequest) (*Initialize64BitQueryCellSupportResponse, error)
}

func RegisterCatalog64BitSupportServer(conn dcerpc.Conn, o Catalog64BitSupportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalog64BitSupportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Catalog64BitSupportSyntaxV0_0))...)
}

func NewCatalog64BitSupportServerHandle(o Catalog64BitSupportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Catalog64BitSupportServerHandle(ctx, o, opNum, r)
	}
}

func Catalog64BitSupportServerHandle(ctx context.Context, o Catalog64BitSupportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SupportsMultipleBitness
		op := &xxx_SupportsMultipleBitnessOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SupportsMultipleBitnessRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SupportsMultipleBitness(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Initialize64BitQueryCellSupport
		op := &xxx_Initialize64BitQueryCellSupportOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Initialize64BitQueryCellSupportRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Initialize64BitQueryCellSupport(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICatalog64BitSupport
type UnimplementedCatalog64BitSupportServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedCatalog64BitSupportServer) SupportsMultipleBitness(context.Context, *SupportsMultipleBitnessRequest) (*SupportsMultipleBitnessResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCatalog64BitSupportServer) Initialize64BitQueryCellSupport(context.Context, *Initialize64BitQueryCellSupportRequest) (*Initialize64BitQueryCellSupportResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Catalog64BitSupportServer = (*UnimplementedCatalog64BitSupportServer)(nil)
