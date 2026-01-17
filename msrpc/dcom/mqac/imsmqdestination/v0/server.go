package imsmqdestination

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

// IMSMQDestination server interface.
type DestinationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Open operation.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error)

	// IADs operation.
	GetIADs(context.Context, *GetIADsRequest) (*GetIADsResponse, error)

	// IADs operation.
	SetByRefIADs(context.Context, *SetByRefIADsRequest) (*SetByRefIADsResponse, error)

	// ADsPath operation.
	GetADsPath(context.Context, *GetADsPathRequest) (*GetADsPathResponse, error)

	// ADsPath operation.
	SetADsPath(context.Context, *SetADsPathRequest) (*SetADsPathResponse, error)

	// PathName operation.
	GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error)

	// PathName operation.
	SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error)

	// FormatName operation.
	SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error)

	// Destinations operation.
	GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponse, error)

	// Destinations operation.
	SetByRefDestinations(context.Context, *SetByRefDestinationsRequest) (*SetByRefDestinationsResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterDestinationServer(conn dcerpc.Conn, o DestinationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDestinationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DestinationSyntaxV0_0))...)
}

func NewDestinationServerHandle(o DestinationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DestinationServerHandle(ctx, o, opNum, r)
	}
}

func DestinationServerHandle(ctx context.Context, o DestinationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Open
		op := &xxx_OpenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Open(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Close
		op := &xxx_CloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Close(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsOpen
		op := &xxx_GetIsOpenOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsOpenRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsOpen(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // IADs
		op := &xxx_GetIADsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIADsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIADs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IADs
		op := &xxx_SetByRefIADsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefIADsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefIADs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ADsPath
		op := &xxx_GetADsPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetADsPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetADsPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ADsPath
		op := &xxx_SetADsPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetADsPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetADsPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // PathName
		op := &xxx_GetPathNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPathNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPathName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // PathName
		op := &xxx_SetPathNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPathNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPathName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // FormatName
		op := &xxx_GetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // FormatName
		op := &xxx_SetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // Destinations
		op := &xxx_GetDestinationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDestinationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDestinations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Destinations
		op := &xxx_SetByRefDestinationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefDestinationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefDestinations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQDestination
type UnimplementedDestinationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDestinationServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetIADs(context.Context, *GetIADsRequest) (*GetIADsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) SetByRefIADs(context.Context, *SetByRefIADsRequest) (*SetByRefIADsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetADsPath(context.Context, *GetADsPathRequest) (*GetADsPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) SetADsPath(context.Context, *SetADsPathRequest) (*SetADsPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) SetByRefDestinations(context.Context, *SetByRefDestinationsRequest) (*SetByRefDestinationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDestinationServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DestinationServer = (*UnimplementedDestinationServer)(nil)
