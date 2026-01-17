package fileservervssagent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// FileServerVssAgent server interface.
type FileServerVSSAgentServer interface {

	// The GetSupportedVersion method is invoked by the client to get the minimum and maximum
	// versions of the protocol that the server supports.
	//
	// Return Values: The method returns one of the values specified in section 2.2.4. The
	// most common error codes are listed below.
	//
	//	+---------------------------+--------------------------------------------------------------------+
	//	|          RETURN           |                                                                    |
	//	|        VALUE/CODE         |                            DESCRIPTION                             |
	//	|                           |                                                                    |
	//	+---------------------------+--------------------------------------------------------------------+
	//	+---------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | The caller does not have the permissions to perform the operation. |
	//	+---------------------------+--------------------------------------------------------------------+
	GetSupportedVersion(context.Context, *GetSupportedVersionRequest) (*GetSupportedVersionResponse, error)

	// The SetContext method sets the context for the current shadow copy creation process.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed below.
	//
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                    |
	//	|                   VALUE/CODE                   |                            DESCRIPTION                             |
	//	|                                                |                                                                    |
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                      | The caller does not have the permissions to perform the operation. |
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//	| 0x8004231B FSRVP_E_UNSUPPORTED_CONTEXT         | The context value specified is invalid.                            |
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042316 FSRVP_E_SHADOW_COPY_SET_IN_PROGRESS | Creation of another shadow copy set is in progress.                |
	//	+------------------------------------------------+--------------------------------------------------------------------+
	//
	// If the Context parameter contains an invalid value, the server MUST fail the call
	// with FSRVP_E_UNSUPPORTED_CONTEXT.
	SetContext(context.Context, *SetContextRequest) (*SetContextResponse, error)

	// The StartShadowCopySet method is called by the client to initiate a new shadow copy
	// set for shadow copy creation.<6>
	//
	// Return Values: The method returns one of the values specified in section 2.2.4. The
	// most common error codes are listed in the following table:
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                      | The caller does not have the permissions to perform the operation.               |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                        | One or more arguments are invalid.                                               |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042316 FSRVP_E_SHADOW_COPY_SET_IN_PROGRESS | StartShadowCopySet (Opnum 2) was called while the creation of another shadow     |
	//	|                                                | copy set was in progress.                                                        |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If ContextSet is FALSE, the server MUST fail the call with FSRVP_E_BAD_STATE.
	//
	// If there is a ShadowCopySet in the GlobalShadowCopySetTable where ShadowCopySet.Status
	// is not equal to "Recovered", the server MUST fail the call with FSRVP_E_SHADOW_COPY_SET_IN_PROGRESS.
	StartShadowCopySet(context.Context, *StartShadowCopySetRequest) (*StartShadowCopySetResponse, error)

	// The AddToShadowCopySet method adds a share to an existing shadow copy set.
	//
	// Return Values: The method returns one of the values specified in section 2.2.4. The
	// most common error codes are listed in the following table:
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have permission to perform the operation.                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230C FSRVP_E_NOT_SUPPORTED             | The file store that contains the share to be shadow copied is not supported by   |
	//	|                                              | the server.                                                                      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the state of the server.                   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230D FSRVP_E_OBJECT_ALREADY_EXISTS     | The object already exists.                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	AddToShadowCopySet(context.Context, *AddToShadowCopySetRequest) (*AddToShadowCopySetResponse, error)

	// The CommitShadowCopySet method is invoked by the client to commit a given shadow
	// copy set.
	//
	// Return Values: The method returns one of the values specified in section 2.2.4. The
	// most common error codes are listed below.
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the state of the server.     |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042500 FSSAGENT_E_TIMEOUT                | The wait for the shadow copy commit operation has timed out.       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0xFFFFFFFF FSRVP_E_WAIT_FAILED               | The wait for shadow copy commit operation has failed.              |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	CommitShadowCopySet(context.Context, *CommitShadowCopySetRequest) (*CommitShadowCopySetResponse, error)

	// The ExposeShadowCopySet method exposes all the shadow copies in a shadow copy set
	// as file shares on the file server.
	//
	// Return Values: The method returns one of the values specified in section 2.2.4. The
	// most common error codes are listed in the following table.
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the server's state.          |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	ExposeShadowCopySet(context.Context, *ExposeShadowCopySetRequest) (*ExposeShadowCopySetResponse, error)

	// The RecoveryCompleteShadowCopySet method is invoked by the client to indicate to
	// the server that the data associated with the file shares in a shadow copy set have
	// been recovered by the VSS writers.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table:
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x00000000 ZERO                              | The operation completed successfully.                              |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the server's state.          |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	RecoveryCompleteShadowCopySet(context.Context, *RecoveryCompleteShadowCopySetRequest) (*RecoveryCompleteShadowCopySetResponse, error)

	// The AbortShadowCopySet method is invoked by the client to delete a given shadow copy
	// set on the server.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table.
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the server's state.          |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	AbortShadowCopySet(context.Context, *AbortShadowCopySetRequest) (*AbortShadowCopySetResponse, error)

	// The IsPathSupported method is invoked by the client to query if a given share is
	// supported by the server for shadow copy operations.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table:
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED         | The caller does not have the permissions to perform the operation.               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG           | One or more arguments are invalid.                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004230CL FSRVP_E_NOT_SUPPORTED | The file store that contains the share to be shadow copied is not supported by   |
	//	|                                   | the server.                                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	IsPathSupported(context.Context, *IsPathSupportedRequest) (*IsPathSupportedResponse, error)

	// The IsPathShadowCopied method is invoked by the client to query if any shadow copy
	// for a share already exists.
	//
	// Return Values:  The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table.
	//
	//	+---------------------------+--------------------------------------------------------------------+
	//	|          RETURN           |                                                                    |
	//	|        VALUE/CODE         |                            DESCRIPTION                             |
	//	|                           |                                                                    |
	//	+---------------------------+--------------------------------------------------------------------+
	//	+---------------------------+--------------------------------------------------------------------+
	//	| 0x00000000 ZERO           | The operation completed successfully.                              |
	//	+---------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED | The caller does not have the permissions to perform the operation. |
	//	+---------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG   | One or more arguments are invalid.                                 |
	//	+---------------------------+--------------------------------------------------------------------+
	IsPathShadowCopied(context.Context, *IsPathShadowCopiedRequest) (*IsPathShadowCopiedResponse, error)

	// The GetShareMapping method is invoked by the client to get the shadow copy information
	// on a given file share on the server after the shadow copy of the share has been exposed.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table:
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//
	// If the value of Level is invalid, the server MUST fail the call with E_INVALIDARG.
	GetShareMapping(context.Context, *GetShareMappingRequest) (*GetShareMappingResponse, error)

	// The DeleteShareMapping method deletes the mapping of a share's shadow copy from a
	// shadow copy set.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed in the following table:
	//
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                    |
	//	|                  VALUE/CODE                  |                            DESCRIPTION                             |
	//	|                                              |                                                                    |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have the permissions to perform the operation. |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                                 |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042308 FSRVP_E_OBJECT_NOT_FOUND          | The specified object does not exist.                               |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                       |
	//	+----------------------------------------------+--------------------------------------------------------------------+
	DeleteShareMapping(context.Context, *DeleteShareMappingRequest) (*DeleteShareMappingResponse, error)

	// The PrepareShadowCopySet method is invoked by the client to ensure that the server
	// has completed preparation for creating the shadow copy set.
	//
	// Return Values: The method returns one of the values as specified in section 2.2.4.
	// The most common error codes are listed below.
	//
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	|                    RETURN                    |                                                                |
	//	|                  VALUE/CODE                  |                          DESCRIPTION                           |
	//	|                                              |                                                                |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                    | The caller does not have permission to perform the operation.  |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                      | One or more arguments are invalid.                             |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0x80042301 FSRVP_E_BAD_STATE                 | The method call is invalid because of the state of the server. |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0x00000102 FSRVP_E_WAIT_TIMEOUT              | The wait for shadow copy preparation operation has timed out.  |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0xFFFFFFFF FSRVP_E_WAIT_FAILED               | The wait for shadow copy preparation operation has failed.     |
	//	+----------------------------------------------+----------------------------------------------------------------+
	//	| 0x80042501 FSRVP_E_SHADOWCOPYSET_ID_MISMATCH | The provided ShadowCopySetId does not exist.                   |
	//	+----------------------------------------------+----------------------------------------------------------------+
	PrepareShadowCopySet(context.Context, *PrepareShadowCopySetRequest) (*PrepareShadowCopySetResponse, error)
}

func RegisterFileServerVSSAgentServer(conn dcerpc.Conn, o FileServerVSSAgentServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFileServerVSSAgentServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FileServerVSSAgentSyntaxV1_0))...)
}

func NewFileServerVSSAgentServerHandle(o FileServerVSSAgentServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FileServerVSSAgentServerHandle(ctx, o, opNum, r)
	}
}

func FileServerVSSAgentServerHandle(ctx context.Context, o FileServerVSSAgentServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // GetSupportedVersion
		op := &xxx_GetSupportedVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportedVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportedVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // SetContext
		op := &xxx_SetContextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetContextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetContext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // StartShadowCopySet
		op := &xxx_StartShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // AddToShadowCopySet
		op := &xxx_AddToShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddToShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddToShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // CommitShadowCopySet
		op := &xxx_CommitShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CommitShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CommitShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ExposeShadowCopySet
		op := &xxx_ExposeShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExposeShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExposeShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RecoveryCompleteShadowCopySet
		op := &xxx_RecoveryCompleteShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RecoveryCompleteShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RecoveryCompleteShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // AbortShadowCopySet
		op := &xxx_AbortShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AbortShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // IsPathSupported
		op := &xxx_IsPathSupportedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsPathSupportedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsPathSupported(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsPathShadowCopied
		op := &xxx_IsPathShadowCopiedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsPathShadowCopiedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsPathShadowCopied(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetShareMapping
		op := &xxx_GetShareMappingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetShareMappingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetShareMapping(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // DeleteShareMapping
		op := &xxx_DeleteShareMappingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteShareMappingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteShareMapping(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // PrepareShadowCopySet
		op := &xxx_PrepareShadowCopySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PrepareShadowCopySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PrepareShadowCopySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented FileServerVssAgent
type UnimplementedFileServerVSSAgentServer struct {
}

func (UnimplementedFileServerVSSAgentServer) GetSupportedVersion(context.Context, *GetSupportedVersionRequest) (*GetSupportedVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) SetContext(context.Context, *SetContextRequest) (*SetContextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) StartShadowCopySet(context.Context, *StartShadowCopySetRequest) (*StartShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) AddToShadowCopySet(context.Context, *AddToShadowCopySetRequest) (*AddToShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) CommitShadowCopySet(context.Context, *CommitShadowCopySetRequest) (*CommitShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) ExposeShadowCopySet(context.Context, *ExposeShadowCopySetRequest) (*ExposeShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) RecoveryCompleteShadowCopySet(context.Context, *RecoveryCompleteShadowCopySetRequest) (*RecoveryCompleteShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) AbortShadowCopySet(context.Context, *AbortShadowCopySetRequest) (*AbortShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) IsPathSupported(context.Context, *IsPathSupportedRequest) (*IsPathSupportedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) IsPathShadowCopied(context.Context, *IsPathShadowCopiedRequest) (*IsPathShadowCopiedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) GetShareMapping(context.Context, *GetShareMappingRequest) (*GetShareMappingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) DeleteShareMapping(context.Context, *DeleteShareMappingRequest) (*DeleteShareMappingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFileServerVSSAgentServer) PrepareShadowCopySet(context.Context, *PrepareShadowCopySetRequest) (*PrepareShadowCopySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FileServerVSSAgentServer = (*UnimplementedFileServerVSSAgentServer)(nil)
