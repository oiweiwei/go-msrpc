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

	// The DefaultValue method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the system-wide default value for the specified property, as defined
	// by the administration system.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
	//
	//	+------------------------------------+-------------------------------------------------------------------+
	//	|               RETURN               |                                                                   |
	//	|             VALUE/CODE             |                            DESCRIPTION                            |
	//	|                                    |                                                                   |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                             |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                     |
	//	+------------------------------------+-------------------------------------------------------------------+
	//	| 0X80070032 ERROR_NOT_SUPPORTED     | The default value has a type that is not supported by the schema. |
	//	+------------------------------------+-------------------------------------------------------------------+
	GetDefaultValue(context.Context, *GetDefaultValueRequest) (*GetDefaultValueResponse, error)

	// The IsRequired method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns whether the specified property is required to be set on the server
	// when the parent IAppHostElement exists.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsRequired(context.Context, *GetIsRequiredRequest) (*GetIsRequiredResponse, error)

	// The IsUniqueKey method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns whether the specified property must be unique compared to all
	// other properties of the peer collection of IAppHostElement objects. In other words,
	// it applies only to properties that are members of the collection of IAppHostElement
	// objects.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsUniqueKey(context.Context, *GetIsUniqueKeyRequest) (*GetIsUniqueKeyResponse, error)

	// The IsCombinedKey method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns whether the specified property is part of a group of properties
	// that combine to be unique compared to all other properties of peer collection IAppHostElement
	// objects. In other words, it applies only to properties that are members of collection
	// IAppHostElement objects.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsCombinedKey(context.Context, *GetIsCombinedKeyRequest) (*GetIsCombinedKeyResponse, error)

	// The IsExpanded method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns whether the specified property supports being expanded on the
	// server side to expand any embedded system environment variables.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsExpanded(context.Context, *GetIsExpandedRequest) (*GetIsExpandedResponse, error)

	// The ValidationType method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a string representing additional custom validation done
	// when processing the corresponding property. The details of the validation are an
	// implementation detail of the administration system.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// received from the client. In this case, *pbstrValidationType is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
	GetValidationType(context.Context, *GetValidationTypeRequest) (*GetValidationTypeResponse, error)

	// The ValidationParameter method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns any parameter that applies to the ValidationType
	// of the specified property. Again, this is implementation-specific.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, pbstrValidationParameter is not NULL.
	// If processing fails, the server MUST return a nonzero HRESULT code as defined in
	// [MS-ERREF]. The following table describes the error conditions that MUST be handled
	// and the corresponding error codes. A server MAY return additional implementation-specific
	// error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
	GetValidationParameter(context.Context, *GetValidationParameterRequest) (*GetValidationParameterResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest) (*GetMetadataResponse, error)

	// The IsCaseSensitive method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns whether the corresponding property is compared to others
	// in a case-sensitive manner, when determining equality for key (combined/unique) evaluation.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsCaseSensitive(context.Context, *GetIsCaseSensitiveRequest) (*GetIsCaseSensitiveResponse, error)

	// The PossibleValues method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a collection of the possible constant values for the
	// specified property, if applicable. The administration system determines the applicability.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppValues is not NULL. If processing
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
	GetPossibleValues(context.Context, *GetPossibleValuesRequest) (*GetPossibleValuesResponse, error)

	// The DoesAllowInfinite method is received by the server in an RPC_REQUEST packet.
	// In response, the server returns whether the property supports having an infinite
	// value set.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetDoesAllowInfinite(context.Context, *GetDoesAllowInfiniteRequest) (*GetDoesAllowInfiniteResponse, error)

	// The IsEncrypted method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns whether the corresponding IAppHostProperty will be encrypted when
	// it is persisted in the administration system.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. If processing fails, the server MUST return a nonzero
	// HRESULT code as defined in [MS-ERREF]. The following table describes the error conditions
	// that MUST be handled and the corresponding error codes. A server MAY return additional
	// implementation-specific error codes.
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
	GetIsEncrypted(context.Context, *GetIsEncryptedRequest) (*GetIsEncryptedResponse, error)

	// The TimeSpanFormat method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns a format string that describes how the corresponding
	// property is supposed to be formatted if the property represents a time span.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrTimeSpanFormat is not NULL.
	// If processing fails, the server MUST return a nonzero HRESULT code as defined in
	// [MS-ERREF]. The following table describes the error conditions that MUST be handled
	// and the corresponding error codes. A server MAY return additional implementation-specific
	// error codes.
	//
	//	+------------------------------------+------------------------------------------------------------------------+
	//	|               RETURN               |                                                                        |
	//	|             VALUE/CODE             |                              DESCRIPTION                               |
	//	|                                    |                                                                        |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                  |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                          |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                |
	//	+------------------------------------+------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted. |
	//	+------------------------------------+------------------------------------------------------------------------+
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
