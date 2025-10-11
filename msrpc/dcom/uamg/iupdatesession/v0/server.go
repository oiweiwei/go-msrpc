package iupdatesession

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

// IUpdateSession server interface.
type UpdateSessionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
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
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
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
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	// The IUpdateSession::ReadOnly (opnum 10) method returns whether the session is read-only.
	//
	// The IUpdateCollection::ReadOnly (opnum 12) method returns whether the collection
	// is read-only.
	//
	// The IStringCollection::ReadOnly (opnum 12) method retrieves whether the collection
	// is read-only.
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
	// This method SHOULD return the value of the ReadOnly ADM element.
	GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error)

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// The IUpdateSession::CreateUpdateSearcher (opnum 13) method retrieves an instance
	// of the IUpdateSearcher interface.
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
	// This method SHOULD return a newly created object implementing the IUpdateSearcher
	// interface.
	//
	// The returned IUpdateSearcher instance SHOULD use the current value of the server's
	// ClientApplicationID ADM element as the client application ID for operations that
	// it performs. The IUpdateSearcher instance SHOULD use the current value of the server's
	// UserLocale ADM element to determine the language in which to generate localized results
	// for operations that it performs.
	CreateUpdateSearcher(context.Context, *CreateUpdateSearcherRequest) (*CreateUpdateSearcherResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire
}

func RegisterUpdateSessionServer(conn dcerpc.Conn, o UpdateSessionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSessionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSessionSyntaxV0_0))...)
}

func NewUpdateSessionServerHandle(o UpdateSessionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSessionServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSessionServerHandle(ctx context.Context, o UpdateSessionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ReadOnly
		op := &xxx_GetReadOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReadOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReadOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 11: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	case 12: // CreateUpdateSearcher
		op := &xxx_CreateUpdateSearcherOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateUpdateSearcherRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateUpdateSearcher(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 14: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IUpdateSession
type UnimplementedUpdateSessionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateSessionServer) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) CreateUpdateSearcher(context.Context, *CreateUpdateSearcherRequest) (*CreateUpdateSearcherResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSessionServer = (*UnimplementedUpdateSessionServer)(nil)
