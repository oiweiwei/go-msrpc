package ivaluemap

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

// IValueMap server interface.
type ValueMapServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// _NewEnum operation.
	Get_NewEnum(context.Context, *Get_NewEnumRequest) (*Get_NewEnumResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// Value operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Value operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

	// The ValueMapType enumeration defines a value map type. A value map defines a named-value
	// pair. A value map can be used in different ways. A value map type defines which way
	// the value map is to be used; each type has a different way of evaluating the "value"
	// of the "value map" based on the "values" of each individual "value map item".
	GetValueMapType(context.Context, *GetValueMapTypeRequest) (*GetValueMapTypeResponse, error)

	// The ValueMapType enumeration defines a value map type. A value map defines a named-value
	// pair. A value map can be used in different ways. A value map type defines which way
	// the value map is to be used; each type has a different way of evaluating the "value"
	// of the "value map" based on the "values" of each individual "value map item".
	SetValueMapType(context.Context, *SetValueMapTypeRequest) (*SetValueMapTypeResponse, error)

	// Add operation.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// Remove operation.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// AddRange operation.
	AddRange(context.Context, *AddRangeRequest) (*AddRangeResponse, error)

	// The CreateValueMapItem method creates an object implementing the IValueMapItem interface.
	// This object can be configured and then passed to IValueMap::Add. This method exists
	// to provide a means for populating ValueMaps.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	CreateValueMapItem(context.Context, *CreateValueMapItemRequest) (*CreateValueMapItemResponse, error)
}

func RegisterValueMapServer(conn dcerpc.Conn, o ValueMapServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewValueMapServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ValueMapSyntaxV0_0))...)
}

func NewValueMapServerHandle(o ValueMapServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ValueMapServerHandle(ctx, o, opNum, r)
	}
}

func ValueMapServerHandle(ctx context.Context, o ValueMapServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // _NewEnum
		in := &Get_NewEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Get_NewEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Value
		in := &GetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Value
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // ValueMapType
		in := &GetValueMapTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValueMapType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ValueMapType
		in := &SetValueMapTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValueMapType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Add
		in := &AddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Add(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Remove
		in := &RemoveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Remove(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // Clear
		in := &ClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // AddRange
		in := &AddRangeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddRange(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // CreateValueMapItem
		in := &CreateValueMapItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateValueMapItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
