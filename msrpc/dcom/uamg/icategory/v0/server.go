package icategory

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

// ICategory server interface.
type CategoryServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The ICategory::Name (opnum 8) method retrieves the name of the update category.
	//
	// The IUpdateService::Name (opnum 8) method retrieves the name of the update service.
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
	// This method SHOULD return the value of the Name ADM element.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// The ICategory:: CategoryID (opnum 9) method retrieves the ID of the update category.
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
	// This method SHOULD return the value of the ID ADM element.
	GetCategoryID(context.Context, *GetCategoryIDRequest) (*GetCategoryIDResponse, error)

	// The ICategory::Children (opnum 10) method retrieves an ICategoryCollection interface
	// instance containing the children of the update category.
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
	// This method SHOULD return the value of the Children ADM element.
	GetChildren(context.Context, *GetChildrenRequest) (*GetChildrenResponse, error)

	// The ICategory::Description (opnum 11) method retrieves the description of the update
	// category.
	//
	// The IUpdateHistoryEntry::Description (opnum 14) method retrieves the description
	// of the update for which the operation is performed.
	//
	// The IUpdate::Description (opnum 16) method retrieves the localized description for
	// the update.
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
	// This method SHOULD return the value of the Description ADM element.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// The ICategory::Image (opnum 12) method retrieves an IImageInformation interface instance
	// that contains information on the image associated with the update category.
	//
	// The IUpdate::Image (opnum 21) method retrieves a localized image associated with
	// the update.
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
	// This method SHOULD return the value of the Image ADM element.
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)

	// The ICategory::Order (opnum 13) method retrieves the recommended display order of
	// the update category among its siblings.
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
	// This method SHOULD return the value of the Order ADM element.
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)

	// The ICategory::Parent (opnum 14) method retrieves the parent of the update category.
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
	// This method SHOULD return the value of the Parent ADM element.
	GetParent(context.Context, *GetParentRequest) (*GetParentResponse, error)

	// The ICategory::Type (opnum 15) method retrieves the type of the update category.
	//
	// The IUpdate::Type (opnum 43) method retrieves the type of the update.
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
	// This method SHOULD return the value of the Type ADM element.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// The ICategory::Updates (opnum 16) method retrieves an IUpdateCollection interface
	// containing the updates that belong to the update category.
	//
	// The ISearchResult::Updates (opnum 10) method retrieves the updates found during the
	// search.
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
	// This method SHOULD return the value of the Updates ADM element.
	GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error)
}

func RegisterCategoryServer(conn dcerpc.Conn, o CategoryServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCategoryServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CategorySyntaxV0_0))...)
}

func NewCategoryServerHandle(o CategoryServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CategoryServerHandle(ctx, o, opNum, r)
	}
}

func CategoryServerHandle(ctx context.Context, o CategoryServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CategoryID
		op := &xxx_GetCategoryIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCategoryIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCategoryID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Children
		op := &xxx_GetChildrenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetChildrenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetChildren(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Image
		op := &xxx_GetImageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Order
		op := &xxx_GetOrderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOrderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOrder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Parent
		op := &xxx_GetParentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Updates
		op := &xxx_GetUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICategory
type UnimplementedCategoryServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCategoryServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetCategoryID(context.Context, *GetCategoryIDRequest) (*GetCategoryIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetChildren(context.Context, *GetChildrenRequest) (*GetChildrenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetParent(context.Context, *GetParentRequest) (*GetParentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCategoryServer) GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CategoryServer = (*UnimplementedCategoryServer)(nil)
