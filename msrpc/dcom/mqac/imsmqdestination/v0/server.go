package imsmqdestination

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

// IMSMQDestination server interface.
type ImsmqDestinationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Open operation.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error)

	// IADs operation.
	GetIaDS(context.Context, *GetIaDSRequest) (*GetIaDSResponse, error)

	// IADs operation.
	SetByRefIaDS(context.Context, *SetByRefIaDSRequest) (*SetByRefIaDSResponse, error)

	// ADsPath operation.
	GetADSPath(context.Context, *GetADSPathRequest) (*GetADSPathResponse, error)

	// ADsPath operation.
	SetADSPath(context.Context, *SetADSPathRequest) (*SetADSPathResponse, error)

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

func RegisterImsmqDestinationServer(conn dcerpc.Conn, o ImsmqDestinationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqDestinationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqDestinationSyntaxV0_0))...)
}

func NewImsmqDestinationServerHandle(o ImsmqDestinationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqDestinationServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqDestinationServerHandle(ctx context.Context, o ImsmqDestinationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
		op := &xxx_GetIaDSOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIaDSRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIaDS(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IADs
		op := &xxx_SetByRefIaDSOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetByRefIaDSRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetByRefIaDS(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ADsPath
		op := &xxx_GetADSPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetADSPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetADSPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ADsPath
		op := &xxx_SetADSPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetADSPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetADSPath(ctx, req)
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
type UnimplementedImsmqDestinationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImsmqDestinationServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetIsOpen(context.Context, *GetIsOpenRequest) (*GetIsOpenResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetIaDS(context.Context, *GetIaDSRequest) (*GetIaDSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) SetByRefIaDS(context.Context, *SetByRefIaDSRequest) (*SetByRefIaDSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetADSPath(context.Context, *GetADSPathRequest) (*GetADSPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) SetADSPath(context.Context, *SetADSPathRequest) (*SetADSPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetPathName(context.Context, *GetPathNameRequest) (*GetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) SetPathName(context.Context, *SetPathNameRequest) (*SetPathNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) SetFormatName(context.Context, *SetFormatNameRequest) (*SetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) SetByRefDestinations(context.Context, *SetByRefDestinationsRequest) (*SetByRefDestinationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqDestinationServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqDestinationServer = (*UnimplementedImsmqDestinationServer)(nil)
