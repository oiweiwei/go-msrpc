package iapphostpropertyschema

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

// IAppHostPropertySchema server interface.
type AppHostPropertySchemaServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// DefaultValue operation.
	GetDefaultValue(context.Context, *GetDefaultValueRequest) (*GetDefaultValueResponse, error)

	// IsRequired operation.
	GetIsRequired(context.Context, *GetIsRequiredRequest) (*GetIsRequiredResponse, error)

	// IsUniqueKey operation.
	GetIsUniqueKey(context.Context, *GetIsUniqueKeyRequest) (*GetIsUniqueKeyResponse, error)

	// IsCombinedKey operation.
	GetIsCombinedKey(context.Context, *GetIsCombinedKeyRequest) (*GetIsCombinedKeyResponse, error)

	// IsExpanded operation.
	GetIsExpanded(context.Context, *GetIsExpandedRequest) (*GetIsExpandedResponse, error)

	// ValidationType operation.
	GetValidationType(context.Context, *GetValidationTypeRequest) (*GetValidationTypeResponse, error)

	// ValidationParameter operation.
	GetValidationParameter(context.Context, *GetValidationParameterRequest) (*GetValidationParameterResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// IsCaseSensitive operation.
	GetIsCaseSensitive(context.Context, *GetIsCaseSensitiveRequest) (*GetIsCaseSensitiveResponse, error)

	// PossibleValues operation.
	GetPossibleValues(context.Context, *GetPossibleValuesRequest) (*GetPossibleValuesResponse, error)

	// DoesAllowInfinite operation.
	GetDoesAllowInfinite(context.Context, *GetDoesAllowInfiniteRequest) (*GetDoesAllowInfiniteResponse, error)

	// IsEncrypted operation.
	GetIsEncrypted(context.Context, *GetIsEncryptedRequest) (*GetIsEncryptedResponse, error)

	// TimeSpanFormat operation.
	GetTimeSpanFormat(context.Context, *GetTimeSpanFormatRequest) (*GetTimeSpanFormatResponse, error)
}

func RegisterAppHostPropertySchemaServer(conn dcerpc.Conn, o AppHostPropertySchemaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostPropertySchemaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostPropertySchemaSyntaxV0_0))...)
}

func NewAppHostPropertySchemaServerHandle(o AppHostPropertySchemaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostPropertySchemaServerHandle(ctx, o, opNum, r)
	}
}

func AppHostPropertySchemaServerHandle(ctx context.Context, o AppHostPropertySchemaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // DefaultValue
		op := &xxx_GetDefaultValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDefaultValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDefaultValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // IsRequired
		op := &xxx_GetIsRequiredOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsRequiredRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsRequired(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // IsUniqueKey
		op := &xxx_GetIsUniqueKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsUniqueKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsUniqueKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // IsCombinedKey
		op := &xxx_GetIsCombinedKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsCombinedKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsCombinedKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsExpanded
		op := &xxx_GetIsExpandedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsExpandedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsExpanded(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ValidationType
		op := &xxx_GetValidationTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValidationTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValidationType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ValidationParameter
		op := &xxx_GetValidationParameterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValidationParameterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValidationParameter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetMetadata
		op := &xxx_GetMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IsCaseSensitive
		op := &xxx_GetIsCaseSensitiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsCaseSensitiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsCaseSensitive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // PossibleValues
		op := &xxx_GetPossibleValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPossibleValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPossibleValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // DoesAllowInfinite
		op := &xxx_GetDoesAllowInfiniteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDoesAllowInfiniteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDoesAllowInfinite(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // IsEncrypted
		op := &xxx_GetIsEncryptedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsEncryptedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsEncrypted(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // TimeSpanFormat
		op := &xxx_GetTimeSpanFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTimeSpanFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTimeSpanFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostPropertySchema
type UnimplementedAppHostPropertySchemaServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostPropertySchemaServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetDefaultValue(context.Context, *GetDefaultValueRequest) (*GetDefaultValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsRequired(context.Context, *GetIsRequiredRequest) (*GetIsRequiredResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsUniqueKey(context.Context, *GetIsUniqueKeyRequest) (*GetIsUniqueKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsCombinedKey(context.Context, *GetIsCombinedKeyRequest) (*GetIsCombinedKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsExpanded(context.Context, *GetIsExpandedRequest) (*GetIsExpandedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetValidationType(context.Context, *GetValidationTypeRequest) (*GetValidationTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetValidationParameter(context.Context, *GetValidationParameterRequest) (*GetValidationParameterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsCaseSensitive(context.Context, *GetIsCaseSensitiveRequest) (*GetIsCaseSensitiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetPossibleValues(context.Context, *GetPossibleValuesRequest) (*GetPossibleValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetDoesAllowInfinite(context.Context, *GetDoesAllowInfiniteRequest) (*GetDoesAllowInfiniteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetIsEncrypted(context.Context, *GetIsEncryptedRequest) (*GetIsEncryptedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertySchemaServer) GetTimeSpanFormat(context.Context, *GetTimeSpanFormatRequest) (*GetTimeSpanFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostPropertySchemaServer = (*UnimplementedAppHostPropertySchemaServer)(nil)
