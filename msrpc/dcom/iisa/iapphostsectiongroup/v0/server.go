package iapphostsectiongroup

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

// IAppHostSectionGroup server interface.
type AppHostSectionGroupServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// The Sections method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns a collection of section definitions in the specified section group.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppSections is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+-----------------------------------------------+
	//	|               RETURN               |                                               |
	//	|             VALUE/CODE             |                  DESCRIPTION                  |
	//	|                                    |                                               |
	//	+------------------------------------+-----------------------------------------------+
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.         |
	//	+------------------------------------+-----------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null. |
	//	+------------------------------------+-----------------------------------------------+
	GetSections(context.Context, *GetSectionsRequest) (*GetSectionsResponse, error)

	// The AddSectionGroup method is received by the server in an RPC_REQUEST packet. In
	// response, the server adds a new section group to the specified section group.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppSectionGroup is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------+
	//	|               RETURN               |                                                                |
	//	|             VALUE/CODE             |                          DESCRIPTION                           |
	//	|                                    |                                                                |
	//	+------------------------------------+----------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                          |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                  |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.        |
	//	+------------------------------------+----------------------------------------------------------------+
	//	| 0X800700B7 ERROR_ALREADY_EXISTS    | A section group with name bstrSectionGroupName already exists. |
	//	+------------------------------------+----------------------------------------------------------------+
	AddSectionGroup(context.Context, *AddSectionGroupRequest) (*AddSectionGroupResponse, error)

	// The DeleteSectionGroup method is received by the server in an RPC_REQUEST packet.
	// In response, the server deletes the specified section group.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR             | The operation completed successfully.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070585 ERROR_INVALID_INDEX  | The integer index specified by varIndex is invalid, or the section group with    |
	//	|                                 | name specified by cIndex could not be found.                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0X80070021 ERROR_LOCK_VIOLATION | The instance is not editable.                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	DeleteSectionGroup(context.Context, *DeleteSectionGroupRequest) (*DeleteSectionGroupResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Type operation.
	SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error)
}

func RegisterAppHostSectionGroupServer(conn dcerpc.Conn, o AppHostSectionGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostSectionGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostSectionGroupSyntaxV0_0))...)
}

func NewAppHostSectionGroupServerHandle(o AppHostSectionGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostSectionGroupServerHandle(ctx, o, opNum, r)
	}
}

func AppHostSectionGroupServerHandle(ctx context.Context, o AppHostSectionGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Count
		op := &xxx_GetCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Item
		op := &xxx_GetItemOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetItemRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetItem(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // Sections
		op := &xxx_GetSectionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSectionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSections(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // AddSectionGroup
		op := &xxx_AddSectionGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddSectionGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddSectionGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // DeleteSectionGroup
		op := &xxx_DeleteSectionGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteSectionGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteSectionGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Type
		op := &xxx_SetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostSectionGroup
type UnimplementedAppHostSectionGroupServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostSectionGroupServer) GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) GetSections(context.Context, *GetSectionsRequest) (*GetSectionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) AddSectionGroup(context.Context, *AddSectionGroupRequest) (*AddSectionGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) DeleteSectionGroup(context.Context, *DeleteSectionGroupRequest) (*DeleteSectionGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostSectionGroupServer) SetType(context.Context, *SetTypeRequest) (*SetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostSectionGroupServer = (*UnimplementedAppHostSectionGroupServer)(nil)
