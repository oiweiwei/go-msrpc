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
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Type
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // DefaultValue
		in := &GetDefaultValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDefaultValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // IsRequired
		in := &GetIsRequiredRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsRequired(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // IsUniqueKey
		in := &GetIsUniqueKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsUniqueKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // IsCombinedKey
		in := &GetIsCombinedKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsCombinedKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // IsExpanded
		in := &GetIsExpandedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsExpanded(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // ValidationType
		in := &GetValidationTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValidationType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ValidationParameter
		in := &GetValidationParameterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValidationParameter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // GetMetadata
		in := &GetMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // IsCaseSensitive
		in := &GetIsCaseSensitiveRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsCaseSensitive(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // PossibleValues
		in := &GetPossibleValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPossibleValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // DoesAllowInfinite
		in := &GetDoesAllowInfiniteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDoesAllowInfinite(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // IsEncrypted
		in := &GetIsEncryptedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIsEncrypted(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // TimeSpanFormat
		in := &GetTimeSpanFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTimeSpanFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
