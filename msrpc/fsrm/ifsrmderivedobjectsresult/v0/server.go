package ifsrmderivedobjectsresult

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = idispatch.GoPackage
)

// IFsrmDerivedObjectsResult server interface.
type DerivedObjectsResultServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The DerivedObjects (get) method returns the collection of derived objects for the
	// calling template.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+---------------------------------------+
	//	|         RETURN          |                                       |
	//	|       VALUE/CODE        |              DESCRIPTION              |
	//	|                         |                                       |
	//	+-------------------------+---------------------------------------+
	//	+-------------------------+---------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The derivedObjects parameter is NULL. |
	//	+-------------------------+---------------------------------------+
	GetDerivedObjects(context.Context, *GetDerivedObjectsRequest) (*GetDerivedObjectsResponse, error)

	// The Results (get) method returns the collection HRESULTS received when committing
	// derived objects that were updated as a result of the source template's call to CommitAndUpdateDerived.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------+
	//	|         RETURN          |                                |
	//	|       VALUE/CODE        |          DESCRIPTION           |
	//	|                         |                                |
	//	+-------------------------+--------------------------------+
	//	+-------------------------+--------------------------------+
	//	| 0x80070057 E_INVALIDARG | The results parameter is NULL. |
	//	+-------------------------+--------------------------------+
	GetResults(context.Context, *GetResultsRequest) (*GetResultsResponse, error)
}

func RegisterDerivedObjectsResultServer(conn dcerpc.Conn, o DerivedObjectsResultServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDerivedObjectsResultServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DerivedObjectsResultSyntaxV0_0))...)
}

func NewDerivedObjectsResultServerHandle(o DerivedObjectsResultServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DerivedObjectsResultServerHandle(ctx, o, opNum, r)
	}
}

func DerivedObjectsResultServerHandle(ctx context.Context, o DerivedObjectsResultServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DerivedObjects
		op := &xxx_GetDerivedObjectsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDerivedObjectsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDerivedObjects(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Results
		op := &xxx_GetResultsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResultsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResults(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmDerivedObjectsResult
type UnimplementedDerivedObjectsResultServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDerivedObjectsResultServer) GetDerivedObjects(context.Context, *GetDerivedObjectsRequest) (*GetDerivedObjectsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDerivedObjectsResultServer) GetResults(context.Context, *GetResultsRequest) (*GetResultsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DerivedObjectsResultServer = (*UnimplementedDerivedObjectsResultServer)(nil)
