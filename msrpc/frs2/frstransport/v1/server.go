package frstransport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// FrsTransport server interface.
type TransportServer interface {

	// The CheckConnectivity method determines whether a server can establish an outbound
	// connection (see the EstablishConnection method specified in section 3.2.4.1.2).
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// All nonzero values MUST be treated as equivalent failures unless otherwise specified.
	//
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN          |                                                                                  |
	//	|        VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                          |                                                                                  |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. The server is ready to establish the          |
	//	|                          | specified outbound connection.                                                   |
	//	+--------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error)

	// The EstablishConnection method establishes an outbound connection, uniquely identified
	// by a replication group ID/connection ID pair, from a client to a server. An outbound
	// connection to the server is required before most other operations can be performed.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000235A FRS_ERROR_INCOMPATIBLE_VERSION | The client's DFS-R protocol version is not compatible with the server's DFS-R    |
	//	|                                           | protocol version.                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00002342 FRS_ERROR_CONNECTION_INVALID   | The connection is invalid.                                                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error)

	// The EstablishSession method is used to establish a replicated folder session between
	// the client and server after the client has successfully established an outbound connection
	// with the server via a call to the EstablishConnection method. A replicated folder
	// session with the server is required before most other operations associated with
	// the specified replicated folder can be performed.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+-------------------------------------+
	//	|                  RETURN                   |                                     |
	//	|                VALUE/CODE                 |             DESCRIPTION             |
	//	|                                           |                                     |
	//	+-------------------------------------------+-------------------------------------+
	//	+-------------------------------------------+-------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully.  |
	//	+-------------------------------------------+-------------------------------------+
	//	| 0x00002342 FRS_ERROR_CONNECTION_INVALID   | The connection is invalid.          |
	//	+-------------------------------------------+-------------------------------------+
	//	| 0x00002375 FRS_ERROR_CONTENTSET_READ_ONLY | The replicated folder is read-only. |
	//	+-------------------------------------------+-------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error)

	// The RequestUpdates method is used to obtain a specified set of updates (replicated
	// file metadata) from a server.
	//
	// Return Values: The method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error)

	// The RequestVersionVector method is used to obtain the version chain vector persisted
	// on a server or to request notification when the server's version chain vector changes.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|         *RETURN          |                                    |
	//	|       VALUE/CODE*        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error)

	// The AsyncPoll method is used to register an asynchronous callback, associated with
	// an outbound connection, which the server uses to provide version chain vectors and
	// notifications of version chain vector changes to the client.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error)

	// The RequestRecords method is used to request all (UID, GVSN) pairs that correspond
	// to live (non-tombstone) records on the server for a specified replicated folder during
	// slow sync (see section 3.3.1.3).
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.                                                   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000024FE FRS_ERROR_CSMAN_OFFLINE        | The server is not currently participating in the replication of the specified    |
	//	|                                           | replicated folder.                                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error)

	// Upon successful completion, the client has notified the server that an update received
	// from the server could not be processed by the client.
	UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error)

	// The RawGetFileData method is used to transfer successive segments of compressed marshaled
	// data for a file from the server to the client. This method does not use the Remote
	// Differential Compression Algorithm (as specified in [MS-RDC]) to transfer data.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error)

	// The RdcGetSignatures method is used to obtain a file's RDC signature data from the
	// server.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_RDC_GENERIC          | Unknown error in RDC.              |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method uses the Remote Differential Compression Algorithm, as specified in [MS-RDC],
	// when replicating a changed file.
	GetSignatures(context.Context, *GetSignaturesRequest) (*GetSignaturesResponse, error)

	// The RdcPushSourceNeeds method is used to register requests for file ranges on a server.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	PushSourceNeeds(context.Context, *PushSourceNeedsRequest) (*PushSourceNeedsResponse, error)

	// The RdcGetFileData method is used to obtain file ranges whose requests have previously
	// been queued on a server by calling the RdcPushSourceNeeds method.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+------------------------------------------+------------------------------------+
	//	|                  RETURN                  |                                    |
	//	|                VALUE/CODE                |            DESCRIPTION             |
	//	|                                          |                                    |
	//	+------------------------------------------+------------------------------------+
	//	+------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The method completed successfully. |
	//	+------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER       | The context is invalid.            |
	//	+------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_RDC_GENERIC         | Unknown error in RDC.              |
	//	+------------------------------------------+------------------------------------+
	//	| 0x00002358 FRS_ERROR_XPRESS_INVALID_DATA | The compressed data is invalid.    |
	//	+------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetFileData(context.Context, *GetFileDataRequest) (*GetFileDataResponse, error)

	// The RdcClose method informs the server that the server context information can be
	// released.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+------------------------------------+------------------------------------+
	//	|               RETURN               |                                    |
	//	|             VALUE/CODE             |            DESCRIPTION             |
	//	|                                    |                                    |
	//	+------------------------------------+------------------------------------+
	//	+------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The method completed successfully. |
	//	+------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The context is invalid.            |
	//	+------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// The InitializeFileTransferAsync method is used by a client to start a file download.
	// The client supplies an update to specify which file to download. The server provides
	// its latest version of the update and initial file contents. The server returns information
	// about the file currently being replicated and the first buffer of data from that
	// file (if any).
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002342 FRS_ERROR_CONNECTION_INVALID   | The connection is invalid.         |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_RDC_GENERIC          | Unknown error in RDC.              |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002358 FRS_ERROR_XPRESS_INVALID_DATA  | The compressed data is invalid.    |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// The RawGetFileDataAsync method is used instead of calling RawGetFileData multiple
	// times to obtain file data. As specified in [MS-RPCE], the specification for asynchronous
	// RPC, an RPC client pulls file data from the byte pipe until receiving an end-of-file
	// notification from the pipe.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error)

	// The RdcGetFileDataAsync method is used instead of calling RdcGetFileData multiple
	// times to obtain file data. As specified in [MS-RPCE], the specification for asynchronous
	// RPC, an RPC client pulls file data from the byte pipe until receiving an end-of-file
	// notification from the pipe.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_RDC_GENERIC          | Unknown error in RDC.              |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002358 FRS_ERROR_XPRESS_INVALID_DATA  | The compressed data is invalid.    |
	//	+-------------------------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetFileDataAsync(context.Context, *GetFileDataAsyncRequest) (*GetFileDataAsyncResponse, error)

	// The RdcFileDataTransferKeepAlive method is used to keep the server context alive.
	//
	// Return Values: This method MUST return 0 on success or a nonzero error code on failure.
	// For protocol purposes all nonzero values MUST be treated as equivalent failures unless
	// otherwise specified.
	//
	//	+-------------------------------------------+------------------------------------+
	//	|                  RETURN                   |                                    |
	//	|                VALUE/CODE                 |            DESCRIPTION             |
	//	|                                           |                                    |
	//	+-------------------------------------------+------------------------------------+
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The method completed successfully. |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER        | The context is invalid.            |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x00002344 FRS_ERROR_CONTENTSET_NOT_FOUND | The content set was not found.     |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_RDC_GENERIC          | Unknown error in RDC.              |
	//	+-------------------------------------------+------------------------------------+
	//	| 0x0000234B FRS_ERROR_IN_BACKUP_RESTORE    | Paused for backup or restore.      |
	//	+-------------------------------------------+------------------------------------+
	FileDataTransferKeepAlive(context.Context, *FileDataTransferKeepAliveRequest) (*FileDataTransferKeepAliveResponse, error)
}

func RegisterTransportServer(conn dcerpc.Conn, o TransportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTransportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TransportSyntaxV1_0))...)
}

func NewTransportServerHandle(o TransportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TransportServerHandle(ctx, o, opNum, r)
	}
}

func TransportServerHandle(ctx context.Context, o TransportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // CheckConnectivity
		op := &xxx_CheckConnectivityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckConnectivityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckConnectivity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // EstablishConnection
		op := &xxx_EstablishConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EstablishConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EstablishConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // EstablishSession
		op := &xxx_EstablishSessionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EstablishSessionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EstablishSession(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RequestUpdates
		op := &xxx_RequestUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RequestVersionVector
		op := &xxx_RequestVersionVectorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestVersionVectorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestVersionVector(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // AsyncPoll
		op := &xxx_AsyncPollOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AsyncPollRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AsyncPoll(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RequestRecords
		op := &xxx_RequestRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RequestRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RequestRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // UpdateCancel
		op := &xxx_UpdateCancelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateCancelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateCancel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RawGetFileData
		op := &xxx_RawGetFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawGetFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawGetFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RdcGetSignatures
		op := &xxx_GetSignaturesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSignaturesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSignatures(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RdcPushSourceNeeds
		op := &xxx_PushSourceNeedsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PushSourceNeedsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PushSourceNeeds(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RdcGetFileData
		op := &xxx_GetFileDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RdcClose
		op := &xxx_CloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Close(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // InitializeFileTransferAsync
		op := &xxx_InitializeFileTransferAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeFileTransferAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeFileTransferAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 15: // RawGetFileDataAsync
		op := &xxx_RawGetFileDataAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RawGetFileDataAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RawGetFileDataAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RdcGetFileDataAsync
		op := &xxx_GetFileDataAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileDataAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileDataAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RdcFileDataTransferKeepAlive
		op := &xxx_FileDataTransferKeepAliveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FileDataTransferKeepAliveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FileDataTransferKeepAlive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented FrsTransport
type UnimplementedTransportServer struct {
}

func (UnimplementedTransportServer) CheckConnectivity(context.Context, *CheckConnectivityRequest) (*CheckConnectivityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) EstablishConnection(context.Context, *EstablishConnectionRequest) (*EstablishConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) EstablishSession(context.Context, *EstablishSessionRequest) (*EstablishSessionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestUpdates(context.Context, *RequestUpdatesRequest) (*RequestUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestVersionVector(context.Context, *RequestVersionVectorRequest) (*RequestVersionVectorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) AsyncPoll(context.Context, *AsyncPollRequest) (*AsyncPollResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RequestRecords(context.Context, *RequestRecordsRequest) (*RequestRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) UpdateCancel(context.Context, *UpdateCancelRequest) (*UpdateCancelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RawGetFileData(context.Context, *RawGetFileDataRequest) (*RawGetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetSignatures(context.Context, *GetSignaturesRequest) (*GetSignaturesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) PushSourceNeeds(context.Context, *PushSourceNeedsRequest) (*PushSourceNeedsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetFileData(context.Context, *GetFileDataRequest) (*GetFileDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest) (*InitializeFileTransferAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest) (*RawGetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) GetFileDataAsync(context.Context, *GetFileDataAsyncRequest) (*GetFileDataAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTransportServer) FileDataTransferKeepAlive(context.Context, *FileDataTransferKeepAliveRequest) (*FileDataTransferKeepAliveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TransportServer = (*UnimplementedTransportServer)(nil)
