package fileservervssagent

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "fsrvp"
)

var (
	// Syntax UUID
	FileServerVSSAgentSyntaxUUID = &uuid.UUID{TimeLow: 0xa8e0653c, TimeMid: 0x2744, TimeHiAndVersion: 0x4389, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x1d, Node: [6]uint8{0x73, 0x73, 0xdf, 0x8b, 0x22, 0x92}}
	// Syntax ID
	FileServerVSSAgentSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: FileServerVSSAgentSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// FileServerVssAgent interface.
type FileServerVSSAgentClient interface {

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
	GetSupportedVersion(context.Context, *GetSupportedVersionRequest, ...dcerpc.CallOption) (*GetSupportedVersionResponse, error)

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
	SetContext(context.Context, *SetContextRequest, ...dcerpc.CallOption) (*SetContextResponse, error)

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
	StartShadowCopySet(context.Context, *StartShadowCopySetRequest, ...dcerpc.CallOption) (*StartShadowCopySetResponse, error)

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
	AddToShadowCopySet(context.Context, *AddToShadowCopySetRequest, ...dcerpc.CallOption) (*AddToShadowCopySetResponse, error)

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
	CommitShadowCopySet(context.Context, *CommitShadowCopySetRequest, ...dcerpc.CallOption) (*CommitShadowCopySetResponse, error)

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
	ExposeShadowCopySet(context.Context, *ExposeShadowCopySetRequest, ...dcerpc.CallOption) (*ExposeShadowCopySetResponse, error)

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
	RecoveryCompleteShadowCopySet(context.Context, *RecoveryCompleteShadowCopySetRequest, ...dcerpc.CallOption) (*RecoveryCompleteShadowCopySetResponse, error)

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
	AbortShadowCopySet(context.Context, *AbortShadowCopySetRequest, ...dcerpc.CallOption) (*AbortShadowCopySetResponse, error)

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
	IsPathSupported(context.Context, *IsPathSupportedRequest, ...dcerpc.CallOption) (*IsPathSupportedResponse, error)

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
	IsPathShadowCopied(context.Context, *IsPathShadowCopiedRequest, ...dcerpc.CallOption) (*IsPathShadowCopiedResponse, error)

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
	GetShareMapping(context.Context, *GetShareMappingRequest, ...dcerpc.CallOption) (*GetShareMappingResponse, error)

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
	DeleteShareMapping(context.Context, *DeleteShareMappingRequest, ...dcerpc.CallOption) (*DeleteShareMappingResponse, error)

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
	PrepareShadowCopySet(context.Context, *PrepareShadowCopySetRequest, ...dcerpc.CallOption) (*PrepareShadowCopySetResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ShareMapping1 structure represents FSSAGENT_SHARE_MAPPING_1 RPC structure.
//
// This structure contains the mapping information for a file share to its shadow copy.
type ShareMapping1 struct {
	// ShadowCopySetId:  The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// ShadowCopyId:  The GUID of the shadow copy.
	ShadowCopyID *dtyp.GUID `idl:"name:ShadowCopyId" json:"shadow_copy_id"`
	// ShareNameUNC:  The name of the share in UNC format.
	ShareNameUNC string `idl:"name:ShareNameUNC;string" json:"share_name_unc"`
	// ShadowCopyShareName:  The name of the share exposing the shadow copy of the base
	// share identified by ShareNameUNC, in UNC format.
	ShadowCopyShareName string `idl:"name:ShadowCopyShareName;string" json:"shadow_copy_share_name"`
	// CreationTimestamp:  The time at which the shadow copy of the share is created. This
	// MUST be a 64-bit integer value containing the number of 100-nanosecond intervals
	// since January 1, 1601 (UTC).
	CreationTimestamp int64 `idl:"name:CreationTimestamp" json:"creation_timestamp"`
}

func (o *ShareMapping1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ShareMapping1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.ShadowCopySetID != nil {
		if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ShadowCopyID != nil {
		if err := o.ShadowCopyID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ShareNameUNC != "" {
		_ptr_ShareNameUNC := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ShareNameUNC); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ShareNameUNC, _ptr_ShareNameUNC); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ShadowCopyShareName != "" {
		_ptr_ShadowCopyShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ShadowCopyShareName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ShadowCopyShareName, _ptr_ShadowCopyShareName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CreationTimestamp); err != nil {
		return err
	}
	return nil
}
func (o *ShareMapping1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.ShadowCopySetID == nil {
		o.ShadowCopySetID = &dtyp.GUID{}
	}
	if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ShadowCopyID == nil {
		o.ShadowCopyID = &dtyp.GUID{}
	}
	if err := o.ShadowCopyID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ShareNameUNC := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareNameUNC); err != nil {
			return err
		}
		return nil
	})
	_s_ShareNameUNC := func(ptr interface{}) { o.ShareNameUNC = *ptr.(*string) }
	if err := w.ReadPointer(&o.ShareNameUNC, _s_ShareNameUNC, _ptr_ShareNameUNC); err != nil {
		return err
	}
	_ptr_ShadowCopyShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShadowCopyShareName); err != nil {
			return err
		}
		return nil
	})
	_s_ShadowCopyShareName := func(ptr interface{}) { o.ShadowCopyShareName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ShadowCopyShareName, _s_ShadowCopyShareName, _ptr_ShadowCopyShareName); err != nil {
		return err
	}
	if err := w.ReadData(&o.CreationTimestamp); err != nil {
		return err
	}
	return nil
}

// ShareMapping structure represents FSSAGENT_SHARE_MAPPING RPC union.
//
// The FSSAGENT_SHARE_MAPPING union contains mapping information for a share to its
// shadow copy based on the level value.
type ShareMapping struct {
	// Types that are assignable to Value
	//
	// *ShareMapping_1
	Value is_ShareMapping `json:"value"`
}

func (o *ShareMapping) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ShareMapping_1:
		if value != nil {
			return value.ShareMapping1
		}
	}
	return nil
}

type is_ShareMapping interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ShareMapping()
}

func (o *ShareMapping) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ShareMapping_1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *ShareMapping) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*ShareMapping_1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ShareMapping_1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *ShareMapping) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &ShareMapping_1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// ShareMapping_1 structure represents FSSAGENT_SHARE_MAPPING RPC union arm.
//
// It has following labels: 1
type ShareMapping_1 struct {
	// ShareMapping1:  A pointer to an FSSAGENT_SHARE_MAPPING_1 structure, as specified
	// in section 2.2.1.1.
	ShareMapping1 *ShareMapping1 `idl:"name:ShareMapping1" json:"share_mapping1"`
}

func (*ShareMapping_1) is_ShareMapping() {}

func (o *ShareMapping_1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ShareMapping1 != nil {
		_ptr_ShareMapping1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ShareMapping1 != nil {
				if err := o.ShareMapping1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ShareMapping1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ShareMapping1, _ptr_ShareMapping1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ShareMapping_1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ShareMapping1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ShareMapping1 == nil {
			o.ShareMapping1 = &ShareMapping1{}
		}
		if err := o.ShareMapping1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ShareMapping1 := func(ptr interface{}) { o.ShareMapping1 = *ptr.(**ShareMapping1) }
	if err := w.ReadPointer(&o.ShareMapping1, _s_ShareMapping1, _ptr_ShareMapping1); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultFileServerVSSAgentClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultFileServerVSSAgentClient) GetSupportedVersion(ctx context.Context, in *GetSupportedVersionRequest, opts ...dcerpc.CallOption) (*GetSupportedVersionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSupportedVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) SetContext(ctx context.Context, in *SetContextRequest, opts ...dcerpc.CallOption) (*SetContextResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) StartShadowCopySet(ctx context.Context, in *StartShadowCopySetRequest, opts ...dcerpc.CallOption) (*StartShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) AddToShadowCopySet(ctx context.Context, in *AddToShadowCopySetRequest, opts ...dcerpc.CallOption) (*AddToShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddToShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) CommitShadowCopySet(ctx context.Context, in *CommitShadowCopySetRequest, opts ...dcerpc.CallOption) (*CommitShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CommitShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) ExposeShadowCopySet(ctx context.Context, in *ExposeShadowCopySetRequest, opts ...dcerpc.CallOption) (*ExposeShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ExposeShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) RecoveryCompleteShadowCopySet(ctx context.Context, in *RecoveryCompleteShadowCopySetRequest, opts ...dcerpc.CallOption) (*RecoveryCompleteShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RecoveryCompleteShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) AbortShadowCopySet(ctx context.Context, in *AbortShadowCopySetRequest, opts ...dcerpc.CallOption) (*AbortShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) IsPathSupported(ctx context.Context, in *IsPathSupportedRequest, opts ...dcerpc.CallOption) (*IsPathSupportedResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsPathSupportedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) IsPathShadowCopied(ctx context.Context, in *IsPathShadowCopiedRequest, opts ...dcerpc.CallOption) (*IsPathShadowCopiedResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsPathShadowCopiedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) GetShareMapping(ctx context.Context, in *GetShareMappingRequest, opts ...dcerpc.CallOption) (*GetShareMappingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetShareMappingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) DeleteShareMapping(ctx context.Context, in *DeleteShareMappingRequest, opts ...dcerpc.CallOption) (*DeleteShareMappingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteShareMappingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) PrepareShadowCopySet(ctx context.Context, in *PrepareShadowCopySetRequest, opts ...dcerpc.CallOption) (*PrepareShadowCopySetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PrepareShadowCopySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileServerVSSAgentClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileServerVSSAgentClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewFileServerVSSAgentClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileServerVSSAgentClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileServerVSSAgentSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultFileServerVSSAgentClient{cc: cc}, nil
}

// xxx_GetSupportedVersionOperation structure represents the GetSupportedVersion operation
type xxx_GetSupportedVersionOperation struct {
	MinVersion uint32 `idl:"name:MinVersion" json:"min_version"`
	MaxVersion uint32 `idl:"name:MaxVersion" json:"max_version"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSupportedVersionOperation) OpNum() int { return 0 }

func (o *xxx_GetSupportedVersionOperation) OpName() string {
	return "/FileServerVssAgent/v1/GetSupportedVersion"
}

func (o *xxx_GetSupportedVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetSupportedVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetSupportedVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// MinVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MinVersion); err != nil {
			return err
		}
	}
	// MaxVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// MinVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MinVersion); err != nil {
			return err
		}
	}
	// MaxVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSupportedVersionRequest structure represents the GetSupportedVersion operation request
type GetSupportedVersionRequest struct {
}

func (o *GetSupportedVersionRequest) xxx_ToOp(ctx context.Context) *xxx_GetSupportedVersionOperation {
	if o == nil {
		return &xxx_GetSupportedVersionOperation{}
	}
	return &xxx_GetSupportedVersionOperation{}
}

func (o *GetSupportedVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedVersionOperation) {
	if o == nil {
		return
	}
}
func (o *GetSupportedVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSupportedVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSupportedVersionResponse structure represents the GetSupportedVersion operation response
type GetSupportedVersionResponse struct {
	// MinVersion:  The minimum version of the protocol that the server supports.
	MinVersion uint32 `idl:"name:MinVersion" json:"min_version"`
	// MaxVersion:  The maximum version of the protocol that the server supports.
	MaxVersion uint32 `idl:"name:MaxVersion" json:"max_version"`
	// Return: The GetSupportedVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSupportedVersionResponse) xxx_ToOp(ctx context.Context) *xxx_GetSupportedVersionOperation {
	if o == nil {
		return &xxx_GetSupportedVersionOperation{}
	}
	return &xxx_GetSupportedVersionOperation{
		MinVersion: o.MinVersion,
		MaxVersion: o.MaxVersion,
		Return:     o.Return,
	}
}

func (o *GetSupportedVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedVersionOperation) {
	if o == nil {
		return
	}
	o.MinVersion = op.MinVersion
	o.MaxVersion = op.MaxVersion
	o.Return = op.Return
}
func (o *GetSupportedVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSupportedVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetContextOperation structure represents the SetContext operation
type xxx_SetContextOperation struct {
	Context uint32 `idl:"name:Context" json:"context"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetContextOperation) OpNum() int { return 1 }

func (o *xxx_SetContextOperation) OpName() string { return "/FileServerVssAgent/v1/SetContext" }

func (o *xxx_SetContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Context {in} (1:(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Context {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetContextRequest structure represents the SetContext operation request
type SetContextRequest struct {
	// Context: The context to be used for the shadow copy operations. It MUST be set to
	// one of the CONTEXT_VALUES specified in section 2.2.2.2.
	Context uint32 `idl:"name:Context" json:"context"`
}

func (o *SetContextRequest) xxx_ToOp(ctx context.Context) *xxx_SetContextOperation {
	if o == nil {
		return &xxx_SetContextOperation{}
	}
	return &xxx_SetContextOperation{
		Context: o.Context,
	}
}

func (o *SetContextRequest) xxx_FromOp(ctx context.Context, op *xxx_SetContextOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *SetContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetContextResponse structure represents the SetContext operation response
type SetContextResponse struct {
	// Return: The SetContext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetContextResponse) xxx_ToOp(ctx context.Context) *xxx_SetContextOperation {
	if o == nil {
		return &xxx_SetContextOperation{}
	}
	return &xxx_SetContextOperation{
		Return: o.Return,
	}
}

func (o *SetContextResponse) xxx_FromOp(ctx context.Context, op *xxx_SetContextOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartShadowCopySetOperation structure represents the StartShadowCopySet operation
type xxx_StartShadowCopySetOperation struct {
	ClientShadowCopySetID *dtyp.GUID `idl:"name:ClientShadowCopySetId" json:"client_shadow_copy_set_id"`
	ShadowCopySetID       *dtyp.GUID `idl:"name:pShadowCopySetId" json:"shadow_copy_set_id"`
	Return                uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_StartShadowCopySetOperation) OpNum() int { return 2 }

func (o *xxx_StartShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/StartShadowCopySet"
}

func (o *xxx_StartShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ClientShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ClientShadowCopySetID != nil {
			if err := o.ClientShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_StartShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ClientShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ClientShadowCopySetID == nil {
			o.ClientShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ClientShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pShadowCopySetId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pShadowCopySetId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartShadowCopySetRequest structure represents the StartShadowCopySet operation request
type StartShadowCopySetRequest struct {
	// ClientShadowCopySetId: The GUID assigned by the client for the shadow copy set.<7>
	ClientShadowCopySetID *dtyp.GUID `idl:"name:ClientShadowCopySetId" json:"client_shadow_copy_set_id"`
}

func (o *StartShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_StartShadowCopySetOperation {
	if o == nil {
		return &xxx_StartShadowCopySetOperation{}
	}
	return &xxx_StartShadowCopySetOperation{
		ClientShadowCopySetID: o.ClientShadowCopySetID,
	}
}

func (o *StartShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_StartShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ClientShadowCopySetID = op.ClientShadowCopySetID
}
func (o *StartShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StartShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartShadowCopySetResponse structure represents the StartShadowCopySet operation response
type StartShadowCopySetResponse struct {
	// pShadowCopySetId: The GUID of the shadow copy set, assigned by the server.
	ShadowCopySetID *dtyp.GUID `idl:"name:pShadowCopySetId" json:"shadow_copy_set_id"`
	// Return: The StartShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *StartShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_StartShadowCopySetOperation {
	if o == nil {
		return &xxx_StartShadowCopySetOperation{}
	}
	return &xxx_StartShadowCopySetOperation{
		ShadowCopySetID: o.ShadowCopySetID,
		Return:          o.Return,
	}
}

func (o *StartShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_StartShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
	o.Return = op.Return
}
func (o *StartShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StartShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddToShadowCopySetOperation structure represents the AddToShadowCopySet operation
type xxx_AddToShadowCopySetOperation struct {
	ClientShadowCopyID *dtyp.GUID `idl:"name:ClientShadowCopyId" json:"client_shadow_copy_id"`
	ShadowCopySetID    *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	ShareName          string     `idl:"name:ShareName;string" json:"share_name"`
	ShadowCopyID       *dtyp.GUID `idl:"name:pShadowCopyId" json:"shadow_copy_id"`
	Return             uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_AddToShadowCopySetOperation) OpNum() int { return 3 }

func (o *xxx_AddToShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/AddToShadowCopySet"
}

func (o *xxx_AddToShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddToShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ClientShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ClientShadowCopyID != nil {
			if err := o.ClientShadowCopyID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddToShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ClientShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ClientShadowCopyID == nil {
			o.ClientShadowCopyID = &dtyp.GUID{}
		}
		if err := o.ClientShadowCopyID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddToShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddToShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pShadowCopyId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ShadowCopyID != nil {
			if err := o.ShadowCopyID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddToShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pShadowCopyId {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ShadowCopyID == nil {
			o.ShadowCopyID = &dtyp.GUID{}
		}
		if err := o.ShadowCopyID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddToShadowCopySetRequest structure represents the AddToShadowCopySet operation request
type AddToShadowCopySetRequest struct {
	// ClientShadowCopyId: The GUID for the shadow copy, assigned by the client.<8>
	ClientShadowCopyID *dtyp.GUID `idl:"name:ClientShadowCopyId" json:"client_shadow_copy_id"`
	// ShadowCopySetId: The GUID of the shadow copy set to which ShareName is to be added.
	// This GUID is assigned by the server.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// ShareName: The name of the share, in UNC format, for which a shadow copy is required.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
}

func (o *AddToShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_AddToShadowCopySetOperation {
	if o == nil {
		return &xxx_AddToShadowCopySetOperation{}
	}
	return &xxx_AddToShadowCopySetOperation{
		ClientShadowCopyID: o.ClientShadowCopyID,
		ShadowCopySetID:    o.ShadowCopySetID,
		ShareName:          o.ShareName,
	}
}

func (o *AddToShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_AddToShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ClientShadowCopyID = op.ClientShadowCopyID
	o.ShadowCopySetID = op.ShadowCopySetID
	o.ShareName = op.ShareName
}
func (o *AddToShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddToShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddToShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddToShadowCopySetResponse structure represents the AddToShadowCopySet operation response
type AddToShadowCopySetResponse struct {
	// pShadowCopyId: The GUID of the shadow copy associated with the share.
	ShadowCopyID *dtyp.GUID `idl:"name:pShadowCopyId" json:"shadow_copy_id"`
	// Return: The AddToShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddToShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_AddToShadowCopySetOperation {
	if o == nil {
		return &xxx_AddToShadowCopySetOperation{}
	}
	return &xxx_AddToShadowCopySetOperation{
		ShadowCopyID: o.ShadowCopyID,
		Return:       o.Return,
	}
}

func (o *AddToShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_AddToShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopyID = op.ShadowCopyID
	o.Return = op.Return
}
func (o *AddToShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddToShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddToShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommitShadowCopySetOperation structure represents the CommitShadowCopySet operation
type xxx_CommitShadowCopySetOperation struct {
	ShadowCopySetID       *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	TimeoutInMilliseconds uint32     `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
	Return                uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitShadowCopySetOperation) OpNum() int { return 4 }

func (o *xxx_CommitShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/CommitShadowCopySet"
}

func (o *xxx_CommitShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CommitShadowCopySetRequest structure represents the CommitShadowCopySet operation request
type CommitShadowCopySetRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set, assigned by the server.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// TimeOutInMilliseconds: The time in milliseconds that the server MUST wait for the
	// shadow copy commit process.
	TimeoutInMilliseconds uint32 `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
}

func (o *CommitShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_CommitShadowCopySetOperation {
	if o == nil {
		return &xxx_CommitShadowCopySetOperation{}
	}
	return &xxx_CommitShadowCopySetOperation{
		ShadowCopySetID:       o.ShadowCopySetID,
		TimeoutInMilliseconds: o.TimeoutInMilliseconds,
	}
}

func (o *CommitShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
	o.TimeoutInMilliseconds = op.TimeoutInMilliseconds
}
func (o *CommitShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CommitShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitShadowCopySetResponse structure represents the CommitShadowCopySet operation response
type CommitShadowCopySetResponse struct {
	// Return: The CommitShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CommitShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_CommitShadowCopySetOperation {
	if o == nil {
		return &xxx_CommitShadowCopySetOperation{}
	}
	return &xxx_CommitShadowCopySetOperation{
		Return: o.Return,
	}
}

func (o *CommitShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CommitShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CommitShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExposeShadowCopySetOperation structure represents the ExposeShadowCopySet operation
type xxx_ExposeShadowCopySetOperation struct {
	ShadowCopySetID       *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	TimeoutInMilliseconds uint32     `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
	Return                uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_ExposeShadowCopySetOperation) OpNum() int { return 5 }

func (o *xxx_ExposeShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/ExposeShadowCopySet"
}

func (o *xxx_ExposeShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExposeShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExposeShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExposeShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExposeShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExposeShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExposeShadowCopySetRequest structure represents the ExposeShadowCopySet operation request
type ExposeShadowCopySetRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// TimeOutInMilliseconds:  The maximum time, in milliseconds, for which the server
	// MUST wait for completion of the expose operation.
	TimeoutInMilliseconds uint32 `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
}

func (o *ExposeShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_ExposeShadowCopySetOperation {
	if o == nil {
		return &xxx_ExposeShadowCopySetOperation{}
	}
	return &xxx_ExposeShadowCopySetOperation{
		ShadowCopySetID:       o.ShadowCopySetID,
		TimeoutInMilliseconds: o.TimeoutInMilliseconds,
	}
}

func (o *ExposeShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_ExposeShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
	o.TimeoutInMilliseconds = op.TimeoutInMilliseconds
}
func (o *ExposeShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExposeShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExposeShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExposeShadowCopySetResponse structure represents the ExposeShadowCopySet operation response
type ExposeShadowCopySetResponse struct {
	// Return: The ExposeShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ExposeShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_ExposeShadowCopySetOperation {
	if o == nil {
		return &xxx_ExposeShadowCopySetOperation{}
	}
	return &xxx_ExposeShadowCopySetOperation{
		Return: o.Return,
	}
}

func (o *ExposeShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_ExposeShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ExposeShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExposeShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExposeShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RecoveryCompleteShadowCopySetOperation structure represents the RecoveryCompleteShadowCopySet operation
type xxx_RecoveryCompleteShadowCopySetOperation struct {
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	Return          uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) OpNum() int { return 6 }

func (o *xxx_RecoveryCompleteShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/RecoveryCompleteShadowCopySet"
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoveryCompleteShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RecoveryCompleteShadowCopySetRequest structure represents the RecoveryCompleteShadowCopySet operation request
type RecoveryCompleteShadowCopySetRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
}

func (o *RecoveryCompleteShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_RecoveryCompleteShadowCopySetOperation {
	if o == nil {
		return &xxx_RecoveryCompleteShadowCopySetOperation{}
	}
	return &xxx_RecoveryCompleteShadowCopySetOperation{
		ShadowCopySetID: o.ShadowCopySetID,
	}
}

func (o *RecoveryCompleteShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_RecoveryCompleteShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
}
func (o *RecoveryCompleteShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RecoveryCompleteShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecoveryCompleteShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecoveryCompleteShadowCopySetResponse structure represents the RecoveryCompleteShadowCopySet operation response
type RecoveryCompleteShadowCopySetResponse struct {
	// Return: The RecoveryCompleteShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RecoveryCompleteShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_RecoveryCompleteShadowCopySetOperation {
	if o == nil {
		return &xxx_RecoveryCompleteShadowCopySetOperation{}
	}
	return &xxx_RecoveryCompleteShadowCopySetOperation{
		Return: o.Return,
	}
}

func (o *RecoveryCompleteShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_RecoveryCompleteShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RecoveryCompleteShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RecoveryCompleteShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecoveryCompleteShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortShadowCopySetOperation structure represents the AbortShadowCopySet operation
type xxx_AbortShadowCopySetOperation struct {
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	Return          uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortShadowCopySetOperation) OpNum() int { return 7 }

func (o *xxx_AbortShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/AbortShadowCopySet"
}

func (o *xxx_AbortShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AbortShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AbortShadowCopySetRequest structure represents the AbortShadowCopySet operation request
type AbortShadowCopySetRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
}

func (o *AbortShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_AbortShadowCopySetOperation {
	if o == nil {
		return &xxx_AbortShadowCopySetOperation{}
	}
	return &xxx_AbortShadowCopySetOperation{
		ShadowCopySetID: o.ShadowCopySetID,
	}
}

func (o *AbortShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
}
func (o *AbortShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AbortShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortShadowCopySetResponse structure represents the AbortShadowCopySet operation response
type AbortShadowCopySetResponse struct {
	// Return: The AbortShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AbortShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_AbortShadowCopySetOperation {
	if o == nil {
		return &xxx_AbortShadowCopySetOperation{}
	}
	return &xxx_AbortShadowCopySetOperation{
		Return: o.Return,
	}
}

func (o *AbortShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AbortShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AbortShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsPathSupportedOperation structure represents the IsPathSupported operation
type xxx_IsPathSupportedOperation struct {
	ShareName               string `idl:"name:ShareName;string" json:"share_name"`
	SupportedByThisProvider bool   `idl:"name:SupportedByThisProvider" json:"supported_by_this_provider"`
	OwnerMachineName        string `idl:"name:OwnerMachineName;string" json:"owner_machine_name"`
	Return                  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_IsPathSupportedOperation) OpNum() int { return 8 }

func (o *xxx_IsPathSupportedOperation) OpName() string {
	return "/FileServerVssAgent/v1/IsPathSupported"
}

func (o *xxx_IsPathSupportedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathSupportedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShareName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathSupportedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShareName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathSupportedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathSupportedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SupportedByThisProvider {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.SupportedByThisProvider {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// OwnerMachineName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.OwnerMachineName != "" {
			_ptr_OwnerMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.OwnerMachineName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.OwnerMachineName, _ptr_OwnerMachineName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathSupportedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SupportedByThisProvider {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bSupportedByThisProvider int32
		if err := w.ReadData(&_bSupportedByThisProvider); err != nil {
			return err
		}
		o.SupportedByThisProvider = _bSupportedByThisProvider != 0
	}
	// OwnerMachineName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_OwnerMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.OwnerMachineName); err != nil {
				return err
			}
			return nil
		})
		_s_OwnerMachineName := func(ptr interface{}) { o.OwnerMachineName = *ptr.(*string) }
		if err := w.ReadPointer(&o.OwnerMachineName, _s_OwnerMachineName, _ptr_OwnerMachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsPathSupportedRequest structure represents the IsPathSupported operation request
type IsPathSupportedRequest struct {
	// ShareName: The full path of the share in UNC format.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
}

func (o *IsPathSupportedRequest) xxx_ToOp(ctx context.Context) *xxx_IsPathSupportedOperation {
	if o == nil {
		return &xxx_IsPathSupportedOperation{}
	}
	return &xxx_IsPathSupportedOperation{
		ShareName: o.ShareName,
	}
}

func (o *IsPathSupportedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsPathSupportedOperation) {
	if o == nil {
		return
	}
	o.ShareName = op.ShareName
}
func (o *IsPathSupportedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsPathSupportedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathSupportedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsPathSupportedResponse structure represents the IsPathSupported operation response
type IsPathSupportedResponse struct {
	// SupportedByThisProvider:  A Boolean, when set to TRUE, that indicates that shadow
	// copies of this share are supported by the server.
	SupportedByThisProvider bool `idl:"name:SupportedByThisProvider" json:"supported_by_this_provider"`
	// OwnerMachineName:  The name of the server machine to which the client MUST connect
	// to create shadow copies of the specified ShareName.
	OwnerMachineName string `idl:"name:OwnerMachineName;string" json:"owner_machine_name"`
	// Return: The IsPathSupported return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *IsPathSupportedResponse) xxx_ToOp(ctx context.Context) *xxx_IsPathSupportedOperation {
	if o == nil {
		return &xxx_IsPathSupportedOperation{}
	}
	return &xxx_IsPathSupportedOperation{
		SupportedByThisProvider: o.SupportedByThisProvider,
		OwnerMachineName:        o.OwnerMachineName,
		Return:                  o.Return,
	}
}

func (o *IsPathSupportedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsPathSupportedOperation) {
	if o == nil {
		return
	}
	o.SupportedByThisProvider = op.SupportedByThisProvider
	o.OwnerMachineName = op.OwnerMachineName
	o.Return = op.Return
}
func (o *IsPathSupportedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsPathSupportedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathSupportedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsPathShadowCopiedOperation structure represents the IsPathShadowCopied operation
type xxx_IsPathShadowCopiedOperation struct {
	ShareName               string `idl:"name:ShareName;string" json:"share_name"`
	ShadowCopyPresent       bool   `idl:"name:ShadowCopyPresent" json:"shadow_copy_present"`
	ShadowCopyCompatibility int32  `idl:"name:ShadowCopyCompatibility" json:"shadow_copy_compatibility"`
	Return                  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_IsPathShadowCopiedOperation) OpNum() int { return 9 }

func (o *xxx_IsPathShadowCopiedOperation) OpName() string {
	return "/FileServerVssAgent/v1/IsPathShadowCopied"
}

func (o *xxx_IsPathShadowCopiedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathShadowCopiedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShareName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathShadowCopiedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShareName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathShadowCopiedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathShadowCopiedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ShadowCopyPresent {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.ShadowCopyPresent {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// ShadowCopyCompatibility {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ShadowCopyCompatibility); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPathShadowCopiedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ShadowCopyPresent {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bShadowCopyPresent int32
		if err := w.ReadData(&_bShadowCopyPresent); err != nil {
			return err
		}
		o.ShadowCopyPresent = _bShadowCopyPresent != 0
	}
	// ShadowCopyCompatibility {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ShadowCopyCompatibility); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsPathShadowCopiedRequest structure represents the IsPathShadowCopied operation request
type IsPathShadowCopiedRequest struct {
	// ShareName: The full path of the share in UNC format.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
}

func (o *IsPathShadowCopiedRequest) xxx_ToOp(ctx context.Context) *xxx_IsPathShadowCopiedOperation {
	if o == nil {
		return &xxx_IsPathShadowCopiedOperation{}
	}
	return &xxx_IsPathShadowCopiedOperation{
		ShareName: o.ShareName,
	}
}

func (o *IsPathShadowCopiedRequest) xxx_FromOp(ctx context.Context, op *xxx_IsPathShadowCopiedOperation) {
	if o == nil {
		return
	}
	o.ShareName = op.ShareName
}
func (o *IsPathShadowCopiedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsPathShadowCopiedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathShadowCopiedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsPathShadowCopiedResponse structure represents the IsPathShadowCopied operation response
type IsPathShadowCopiedResponse struct {
	// ShadowCopyPresent: This value is set to TRUE if the ShareName specified has a shadow
	// copy; otherwise set to FALSE.
	ShadowCopyPresent bool `idl:"name:ShadowCopyPresent" json:"shadow_copy_present"`
	// ShadowCopyCompatibility: This value indicates whether certain I/O operations on the
	// file store containing the shadow copy are disabled. This MUST be zero or a combination
	// of the values as specified in section 2.2.2.3.
	ShadowCopyCompatibility int32 `idl:"name:ShadowCopyCompatibility" json:"shadow_copy_compatibility"`
	// Return: The IsPathShadowCopied return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *IsPathShadowCopiedResponse) xxx_ToOp(ctx context.Context) *xxx_IsPathShadowCopiedOperation {
	if o == nil {
		return &xxx_IsPathShadowCopiedOperation{}
	}
	return &xxx_IsPathShadowCopiedOperation{
		ShadowCopyPresent:       o.ShadowCopyPresent,
		ShadowCopyCompatibility: o.ShadowCopyCompatibility,
		Return:                  o.Return,
	}
}

func (o *IsPathShadowCopiedResponse) xxx_FromOp(ctx context.Context, op *xxx_IsPathShadowCopiedOperation) {
	if o == nil {
		return
	}
	o.ShadowCopyPresent = op.ShadowCopyPresent
	o.ShadowCopyCompatibility = op.ShadowCopyCompatibility
	o.Return = op.Return
}
func (o *IsPathShadowCopiedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsPathShadowCopiedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPathShadowCopiedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetShareMappingOperation structure represents the GetShareMapping operation
type xxx_GetShareMappingOperation struct {
	ShadowCopyID    *dtyp.GUID    `idl:"name:ShadowCopyId" json:"shadow_copy_id"`
	ShadowCopySetID *dtyp.GUID    `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	ShareName       string        `idl:"name:ShareName;string" json:"share_name"`
	Level           uint32        `idl:"name:Level" json:"level"`
	ShareMapping    *ShareMapping `idl:"name:ShareMapping;switch_is:Level" json:"share_mapping"`
	Return          uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetShareMappingOperation) OpNum() int { return 10 }

func (o *xxx_GetShareMappingOperation) OpName() string {
	return "/FileServerVssAgent/v1/GetShareMapping"
}

func (o *xxx_GetShareMappingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareMappingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopyID != nil {
			if err := o.ShadowCopyID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareMappingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopyID == nil {
			o.ShadowCopyID = &dtyp.GUID{}
		}
		if err := o.ShadowCopyID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareMappingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareMappingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ShareMapping {out} (1:{switch_type={}(uint32), alias=PFSSAGENT_SHARE_MAPPING}*(1))(2:{switch_type={}(uint32), alias=FSSAGENT_SHARE_MAPPING}(union))
	{
		_swShareMapping := uint32(o.Level)
		if o.ShareMapping != nil {
			if err := o.ShareMapping.MarshalUnionNDR(ctx, w, _swShareMapping); err != nil {
				return err
			}
		} else {
			if err := (&ShareMapping{}).MarshalUnionNDR(ctx, w, _swShareMapping); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareMappingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ShareMapping {out} (1:{switch_type={}(uint32), alias=PFSSAGENT_SHARE_MAPPING,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=FSSAGENT_SHARE_MAPPING}(union))
	{
		if o.ShareMapping == nil {
			o.ShareMapping = &ShareMapping{}
		}
		_swShareMapping := uint32(o.Level)
		if err := o.ShareMapping.UnmarshalUnionNDR(ctx, w, _swShareMapping); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetShareMappingRequest structure represents the GetShareMapping operation request
type GetShareMappingRequest struct {
	// ShadowCopyId: The GUID of the shadow copy associated with the share.
	ShadowCopyID *dtyp.GUID `idl:"name:ShadowCopyId" json:"shadow_copy_id"`
	// ShadowCopySetId: The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// ShareName: The name of the share in UNC format.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
	// Level: The information level of the share mapping data. This parameter MUST be one
	// of the following values.
	//
	//	+-------+--------------------------+
	//	|       |                          |
	//	| VALUE |         MEANING          |
	//	|       |                          |
	//	+-------+--------------------------+
	//	+-------+--------------------------+
	//	|     1 | FSSAGENT_SHARE_MAPPING_1 |
	//	+-------+--------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetShareMappingRequest) xxx_ToOp(ctx context.Context) *xxx_GetShareMappingOperation {
	if o == nil {
		return &xxx_GetShareMappingOperation{}
	}
	return &xxx_GetShareMappingOperation{
		ShadowCopyID:    o.ShadowCopyID,
		ShadowCopySetID: o.ShadowCopySetID,
		ShareName:       o.ShareName,
		Level:           o.Level,
	}
}

func (o *GetShareMappingRequest) xxx_FromOp(ctx context.Context, op *xxx_GetShareMappingOperation) {
	if o == nil {
		return
	}
	o.ShadowCopyID = op.ShadowCopyID
	o.ShadowCopySetID = op.ShadowCopySetID
	o.ShareName = op.ShareName
	o.Level = op.Level
}
func (o *GetShareMappingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetShareMappingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareMappingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetShareMappingResponse structure represents the GetShareMapping operation response
type GetShareMappingResponse struct {
	// ShareMapping: A pointer to an FSSAGENT_SHARE_MAPPING structure, as specified in section
	// 2.2.3.1.
	ShareMapping *ShareMapping `idl:"name:ShareMapping;switch_is:Level" json:"share_mapping"`
	// Return: The GetShareMapping return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetShareMappingResponse) xxx_ToOp(ctx context.Context) *xxx_GetShareMappingOperation {
	if o == nil {
		return &xxx_GetShareMappingOperation{}
	}
	return &xxx_GetShareMappingOperation{
		ShareMapping: o.ShareMapping,
		Return:       o.Return,
	}
}

func (o *GetShareMappingResponse) xxx_FromOp(ctx context.Context, op *xxx_GetShareMappingOperation) {
	if o == nil {
		return
	}
	o.ShareMapping = op.ShareMapping
	o.Return = op.Return
}
func (o *GetShareMappingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetShareMappingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareMappingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteShareMappingOperation structure represents the DeleteShareMapping operation
type xxx_DeleteShareMappingOperation struct {
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	ShadowCopyID    *dtyp.GUID `idl:"name:ShadowCopyId" json:"shadow_copy_id"`
	ShareName       string     `idl:"name:ShareName;string" json:"share_name"`
	Return          uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteShareMappingOperation) OpNum() int { return 11 }

func (o *xxx_DeleteShareMappingOperation) OpName() string {
	return "/FileServerVssAgent/v1/DeleteShareMapping"
}

func (o *xxx_DeleteShareMappingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteShareMappingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopyID != nil {
			if err := o.ShadowCopyID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteShareMappingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShadowCopyId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopyID == nil {
			o.ShadowCopyID = &dtyp.GUID{}
		}
		if err := o.ShadowCopyID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteShareMappingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteShareMappingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteShareMappingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteShareMappingRequest structure represents the DeleteShareMapping operation request
type DeleteShareMappingRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// ShadowCopyId: The GUID of the shadow copy.
	ShadowCopyID *dtyp.GUID `idl:"name:ShadowCopyId" json:"shadow_copy_id"`
	// ShareName: The name of the share for which the share mapping is to be deleted.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
}

func (o *DeleteShareMappingRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteShareMappingOperation {
	if o == nil {
		return &xxx_DeleteShareMappingOperation{}
	}
	return &xxx_DeleteShareMappingOperation{
		ShadowCopySetID: o.ShadowCopySetID,
		ShadowCopyID:    o.ShadowCopyID,
		ShareName:       o.ShareName,
	}
}

func (o *DeleteShareMappingRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteShareMappingOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
	o.ShadowCopyID = op.ShadowCopyID
	o.ShareName = op.ShareName
}
func (o *DeleteShareMappingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteShareMappingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteShareMappingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteShareMappingResponse structure represents the DeleteShareMapping operation response
type DeleteShareMappingResponse struct {
	// Return: The DeleteShareMapping return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteShareMappingResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteShareMappingOperation {
	if o == nil {
		return &xxx_DeleteShareMappingOperation{}
	}
	return &xxx_DeleteShareMappingOperation{
		Return: o.Return,
	}
}

func (o *DeleteShareMappingResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteShareMappingOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteShareMappingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteShareMappingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteShareMappingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PrepareShadowCopySetOperation structure represents the PrepareShadowCopySet operation
type xxx_PrepareShadowCopySetOperation struct {
	ShadowCopySetID       *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	TimeoutInMilliseconds uint32     `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
	Return                uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_PrepareShadowCopySetOperation) OpNum() int { return 12 }

func (o *xxx_PrepareShadowCopySetOperation) OpName() string {
	return "/FileServerVssAgent/v1/PrepareShadowCopySet"
}

func (o *xxx_PrepareShadowCopySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareShadowCopySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID != nil {
			if err := o.ShadowCopySetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.WriteData(o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareShadowCopySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ShadowCopySetId {in} (1:{alias=GUID}(struct))
	{
		if o.ShadowCopySetID == nil {
			o.ShadowCopySetID = &dtyp.GUID{}
		}
		if err := o.ShadowCopySetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TimeOutInMilliseconds {in} (1:(uint32))
	{
		if err := w.ReadData(&o.TimeoutInMilliseconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareShadowCopySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareShadowCopySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PrepareShadowCopySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PrepareShadowCopySetRequest structure represents the PrepareShadowCopySet operation request
type PrepareShadowCopySetRequest struct {
	// ShadowCopySetId: The GUID of the shadow copy set, assigned by the server.
	ShadowCopySetID *dtyp.GUID `idl:"name:ShadowCopySetId" json:"shadow_copy_set_id"`
	// TimeOutInMilliseconds: The time in milliseconds for which the server MUST wait for
	// the shadow copy preparation process to complete.
	TimeoutInMilliseconds uint32 `idl:"name:TimeOutInMilliseconds" json:"timeout_in_milliseconds"`
}

func (o *PrepareShadowCopySetRequest) xxx_ToOp(ctx context.Context) *xxx_PrepareShadowCopySetOperation {
	if o == nil {
		return &xxx_PrepareShadowCopySetOperation{}
	}
	return &xxx_PrepareShadowCopySetOperation{
		ShadowCopySetID:       o.ShadowCopySetID,
		TimeoutInMilliseconds: o.TimeoutInMilliseconds,
	}
}

func (o *PrepareShadowCopySetRequest) xxx_FromOp(ctx context.Context, op *xxx_PrepareShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.ShadowCopySetID = op.ShadowCopySetID
	o.TimeoutInMilliseconds = op.TimeoutInMilliseconds
}
func (o *PrepareShadowCopySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PrepareShadowCopySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareShadowCopySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PrepareShadowCopySetResponse structure represents the PrepareShadowCopySet operation response
type PrepareShadowCopySetResponse struct {
	// Return: The PrepareShadowCopySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PrepareShadowCopySetResponse) xxx_ToOp(ctx context.Context) *xxx_PrepareShadowCopySetOperation {
	if o == nil {
		return &xxx_PrepareShadowCopySetOperation{}
	}
	return &xxx_PrepareShadowCopySetOperation{
		Return: o.Return,
	}
}

func (o *PrepareShadowCopySetResponse) xxx_FromOp(ctx context.Context, op *xxx_PrepareShadowCopySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PrepareShadowCopySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PrepareShadowCopySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PrepareShadowCopySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
