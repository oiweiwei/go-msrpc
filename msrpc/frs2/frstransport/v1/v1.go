package frstransport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	frs2 "github.com/oiweiwei/go-msrpc/msrpc/frs2"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
	_ = frs2.GoPackage
)

var (
	// import guard
	GoPackage = "frs2"
)

var (
	// Syntax UUID
	TransportSyntaxUUID = &uuid.UUID{TimeLow: 0x897e2e5f, TimeMid: 0x93f3, TimeHiAndVersion: 0x4376, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x9c, Node: [6]uint8{0xfd, 0x22, 0x77, 0x49, 0x5c, 0x27}}
	// Syntax ID
	TransportSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TransportSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// FrsTransport interface.
type TransportClient interface {

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
	CheckConnectivity(context.Context, *CheckConnectivityRequest, ...dcerpc.CallOption) (*CheckConnectivityResponse, error)

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
	EstablishConnection(context.Context, *EstablishConnectionRequest, ...dcerpc.CallOption) (*EstablishConnectionResponse, error)

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
	EstablishSession(context.Context, *EstablishSessionRequest, ...dcerpc.CallOption) (*EstablishSessionResponse, error)

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
	RequestUpdates(context.Context, *RequestUpdatesRequest, ...dcerpc.CallOption) (*RequestUpdatesResponse, error)

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
	RequestVersionVector(context.Context, *RequestVersionVectorRequest, ...dcerpc.CallOption) (*RequestVersionVectorResponse, error)

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
	AsyncPoll(context.Context, *AsyncPollRequest, ...dcerpc.CallOption) (*AsyncPollResponse, error)

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
	RequestRecords(context.Context, *RequestRecordsRequest, ...dcerpc.CallOption) (*RequestRecordsResponse, error)

	// Upon successful completion, the client has notified the server that an update received
	// from the server could not be processed by the client.
	UpdateCancel(context.Context, *UpdateCancelRequest, ...dcerpc.CallOption) (*UpdateCancelResponse, error)

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
	RawGetFileData(context.Context, *RawGetFileDataRequest, ...dcerpc.CallOption) (*RawGetFileDataResponse, error)

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
	GetSignatures(context.Context, *GetSignaturesRequest, ...dcerpc.CallOption) (*GetSignaturesResponse, error)

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
	PushSourceNeeds(context.Context, *PushSourceNeedsRequest, ...dcerpc.CallOption) (*PushSourceNeedsResponse, error)

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
	GetFileData(context.Context, *GetFileDataRequest, ...dcerpc.CallOption) (*GetFileDataResponse, error)

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
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

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
	InitializeFileTransferAsync(context.Context, *InitializeFileTransferAsyncRequest, ...dcerpc.CallOption) (*InitializeFileTransferAsyncResponse, error)

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
	RawGetFileDataAsync(context.Context, *RawGetFileDataAsyncRequest, ...dcerpc.CallOption) (*RawGetFileDataAsyncResponse, error)

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
	GetFileDataAsync(context.Context, *GetFileDataAsyncRequest, ...dcerpc.CallOption) (*GetFileDataAsyncResponse, error)

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
	FileDataTransferKeepAlive(context.Context, *FileDataTransferKeepAliveRequest, ...dcerpc.CallOption) (*FileDataTransferKeepAliveResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// ServerContext structure represents PFRS_SERVER_CONTEXT RPC structure.
type ServerContext dcetypes.ContextHandle

func (o *ServerContext) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *ServerContext) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ServerContext) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ServerContext) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTransportClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTransportClient) CheckConnectivity(ctx context.Context, in *CheckConnectivityRequest, opts ...dcerpc.CallOption) (*CheckConnectivityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CheckConnectivityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) EstablishConnection(ctx context.Context, in *EstablishConnectionRequest, opts ...dcerpc.CallOption) (*EstablishConnectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EstablishConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) EstablishSession(ctx context.Context, in *EstablishSessionRequest, opts ...dcerpc.CallOption) (*EstablishSessionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EstablishSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) RequestUpdates(ctx context.Context, in *RequestUpdatesRequest, opts ...dcerpc.CallOption) (*RequestUpdatesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RequestUpdatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) RequestVersionVector(ctx context.Context, in *RequestVersionVectorRequest, opts ...dcerpc.CallOption) (*RequestVersionVectorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RequestVersionVectorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) AsyncPoll(ctx context.Context, in *AsyncPollRequest, opts ...dcerpc.CallOption) (*AsyncPollResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AsyncPollResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) RequestRecords(ctx context.Context, in *RequestRecordsRequest, opts ...dcerpc.CallOption) (*RequestRecordsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RequestRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) UpdateCancel(ctx context.Context, in *UpdateCancelRequest, opts ...dcerpc.CallOption) (*UpdateCancelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateCancelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) RawGetFileData(ctx context.Context, in *RawGetFileDataRequest, opts ...dcerpc.CallOption) (*RawGetFileDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RawGetFileDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) GetSignatures(ctx context.Context, in *GetSignaturesRequest, opts ...dcerpc.CallOption) (*GetSignaturesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSignaturesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) PushSourceNeeds(ctx context.Context, in *PushSourceNeedsRequest, opts ...dcerpc.CallOption) (*PushSourceNeedsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PushSourceNeedsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) GetFileData(ctx context.Context, in *GetFileDataRequest, opts ...dcerpc.CallOption) (*GetFileDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFileDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) InitializeFileTransferAsync(ctx context.Context, in *InitializeFileTransferAsyncRequest, opts ...dcerpc.CallOption) (*InitializeFileTransferAsyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitializeFileTransferAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) RawGetFileDataAsync(ctx context.Context, in *RawGetFileDataAsyncRequest, opts ...dcerpc.CallOption) (*RawGetFileDataAsyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RawGetFileDataAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) GetFileDataAsync(ctx context.Context, in *GetFileDataAsyncRequest, opts ...dcerpc.CallOption) (*GetFileDataAsyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFileDataAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) FileDataTransferKeepAlive(ctx context.Context, in *FileDataTransferKeepAliveRequest, opts ...dcerpc.CallOption) (*FileDataTransferKeepAliveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FileDataTransferKeepAliveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTransportClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTransportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TransportClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TransportSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTransportClient{cc: cc}, nil
}

// xxx_CheckConnectivityOperation structure represents the CheckConnectivity operation
type xxx_CheckConnectivityOperation struct {
	SetIDReplica *frs2.ReplicaSetID `idl:"name:replicaSetId" json:"set_id_replica"`
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_CheckConnectivityOperation) OpNum() int { return 0 }

func (o *xxx_CheckConnectivityOperation) OpName() string { return "/FrsTransport/v1/CheckConnectivity" }

func (o *xxx_CheckConnectivityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckConnectivityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// replicaSetId {in} (1:{alias=FRS_REPLICA_SET_ID, names=GUID}(struct))
	{
		if o.SetIDReplica != nil {
			if err := o.SetIDReplica.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ReplicaSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CheckConnectivityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// replicaSetId {in} (1:{alias=FRS_REPLICA_SET_ID, names=GUID}(struct))
	{
		if o.SetIDReplica == nil {
			o.SetIDReplica = &frs2.ReplicaSetID{}
		}
		if err := o.SetIDReplica.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckConnectivityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckConnectivityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CheckConnectivityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CheckConnectivityRequest structure represents the CheckConnectivity operation request
type CheckConnectivityRequest struct {
	// replicaSetId: The GUID of the outbound connection’s replication group (see the
	// objectGUID attribute specified in section 2.3.5).
	SetIDReplica *frs2.ReplicaSetID `idl:"name:replicaSetId" json:"set_id_replica"`
	// connectionId: The GUID of the outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) in the specified replication group.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
}

func (o *CheckConnectivityRequest) xxx_ToOp(ctx context.Context, op *xxx_CheckConnectivityOperation) *xxx_CheckConnectivityOperation {
	if op == nil {
		op = &xxx_CheckConnectivityOperation{}
	}
	if o == nil {
		return op
	}
	op.SetIDReplica = o.SetIDReplica
	op.ConnectionID = o.ConnectionID
	return op
}

func (o *CheckConnectivityRequest) xxx_FromOp(ctx context.Context, op *xxx_CheckConnectivityOperation) {
	if o == nil {
		return
	}
	o.SetIDReplica = op.SetIDReplica
	o.ConnectionID = op.ConnectionID
}
func (o *CheckConnectivityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CheckConnectivityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckConnectivityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CheckConnectivityResponse structure represents the CheckConnectivity operation response
type CheckConnectivityResponse struct {
	// Return: The CheckConnectivity return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CheckConnectivityResponse) xxx_ToOp(ctx context.Context, op *xxx_CheckConnectivityOperation) *xxx_CheckConnectivityOperation {
	if op == nil {
		op = &xxx_CheckConnectivityOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CheckConnectivityResponse) xxx_FromOp(ctx context.Context, op *xxx_CheckConnectivityOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CheckConnectivityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CheckConnectivityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckConnectivityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EstablishConnectionOperation structure represents the EstablishConnection operation
type xxx_EstablishConnectionOperation struct {
	SetIDReplica              *frs2.ReplicaSetID `idl:"name:replicaSetId" json:"set_id_replica"`
	ConnectionID              *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	DownstreamProtocolVersion uint32             `idl:"name:downstreamProtocolVersion" json:"downstream_protocol_version"`
	DownstreamFlags           uint32             `idl:"name:downstreamFlags" json:"downstream_flags"`
	UpstreamProtocolVersion   uint32             `idl:"name:upstreamProtocolVersion" json:"upstream_protocol_version"`
	UpstreamFlags             uint32             `idl:"name:upstreamFlags" json:"upstream_flags"`
	Return                    uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_EstablishConnectionOperation) OpNum() int { return 1 }

func (o *xxx_EstablishConnectionOperation) OpName() string {
	return "/FrsTransport/v1/EstablishConnection"
}

func (o *xxx_EstablishConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// replicaSetId {in} (1:{alias=FRS_REPLICA_SET_ID, names=GUID}(struct))
	{
		if o.SetIDReplica != nil {
			if err := o.SetIDReplica.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ReplicaSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// downstreamProtocolVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DownstreamProtocolVersion); err != nil {
			return err
		}
	}
	// downstreamFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DownstreamFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// replicaSetId {in} (1:{alias=FRS_REPLICA_SET_ID, names=GUID}(struct))
	{
		if o.SetIDReplica == nil {
			o.SetIDReplica = &frs2.ReplicaSetID{}
		}
		if err := o.SetIDReplica.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// downstreamProtocolVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DownstreamProtocolVersion); err != nil {
			return err
		}
	}
	// downstreamFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DownstreamFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// upstreamProtocolVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UpstreamProtocolVersion); err != nil {
			return err
		}
	}
	// upstreamFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UpstreamFlags); err != nil {
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

func (o *xxx_EstablishConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// upstreamProtocolVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UpstreamProtocolVersion); err != nil {
			return err
		}
	}
	// upstreamFlags {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UpstreamFlags); err != nil {
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

// EstablishConnectionRequest structure represents the EstablishConnection operation request
type EstablishConnectionRequest struct {
	// replicaSetId: The GUID of the outbound connection's replication group (see the objectGUID
	// attribute specified in section 2.3.5).
	SetIDReplica *frs2.ReplicaSetID `idl:"name:replicaSetId" json:"set_id_replica"`
	// connectionId: The GUID of the outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) in the specified replication group.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// downstreamProtocolVersion: Identifies the version of the DFS-R protocol implemented
	// by the client. Currently implemented protocol versions are specified in section 2.2.1.1.1.
	DownstreamProtocolVersion uint32 `idl:"name:downstreamProtocolVersion" json:"downstream_protocol_version"`
	// downstreamFlags: This parameter is unused and SHOULD be set to 0 by the client.
	// <22>
	DownstreamFlags uint32 `idl:"name:downstreamFlags" json:"downstream_flags"`
}

func (o *EstablishConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_EstablishConnectionOperation) *xxx_EstablishConnectionOperation {
	if op == nil {
		op = &xxx_EstablishConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.SetIDReplica = o.SetIDReplica
	op.ConnectionID = o.ConnectionID
	op.DownstreamProtocolVersion = o.DownstreamProtocolVersion
	op.DownstreamFlags = o.DownstreamFlags
	return op
}

func (o *EstablishConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_EstablishConnectionOperation) {
	if o == nil {
		return
	}
	o.SetIDReplica = op.SetIDReplica
	o.ConnectionID = op.ConnectionID
	o.DownstreamProtocolVersion = op.DownstreamProtocolVersion
	o.DownstreamFlags = op.DownstreamFlags
}
func (o *EstablishConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EstablishConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EstablishConnectionResponse structure represents the EstablishConnection operation response
type EstablishConnectionResponse struct {
	// upstreamProtocolVersion: Receives the version of the DFS-R protocol implemented by
	// the server. Currently implemented protocol versions are specified in section 2.2.1.1.1.
	UpstreamProtocolVersion uint32 `idl:"name:upstreamProtocolVersion" json:"upstream_protocol_version"`
	// upstreamFlags: A flags bitmask. The server MUST set the TRANSPORT_SUPPORTS_RDC_SIMILARITY
	// bit flag to 1 if the server supports RDC similarity (as specified in [MS-RDC] section
	// 3.1.5.4). Otherwise, the server MUST clear this bitmask (set all bits to 0). The
	// client MUST ignore any bit flags other than TRANSPORT_SUPPORTS_RDC_SIMILARITY.
	UpstreamFlags uint32 `idl:"name:upstreamFlags" json:"upstream_flags"`
	// Return: The EstablishConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EstablishConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_EstablishConnectionOperation) *xxx_EstablishConnectionOperation {
	if op == nil {
		op = &xxx_EstablishConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.UpstreamProtocolVersion = o.UpstreamProtocolVersion
	op.UpstreamFlags = o.UpstreamFlags
	op.Return = o.Return
	return op
}

func (o *EstablishConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_EstablishConnectionOperation) {
	if o == nil {
		return
	}
	o.UpstreamProtocolVersion = op.UpstreamProtocolVersion
	o.UpstreamFlags = op.UpstreamFlags
	o.Return = op.Return
}
func (o *EstablishConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EstablishConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EstablishSessionOperation structure represents the EstablishSession operation
type xxx_EstablishSessionOperation struct {
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	ContentSetID *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_EstablishSessionOperation) OpNum() int { return 2 }

func (o *xxx_EstablishSessionOperation) OpName() string { return "/FrsTransport/v1/EstablishSession" }

func (o *xxx_EstablishSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID != nil {
			if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ContentSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EstablishSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID == nil {
			o.ContentSetID = &frs2.ContentSetID{}
		}
		if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EstablishSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EstablishSessionRequest structure represents the EstablishSession operation request
type EstablishSessionRequest struct {
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// contentSetId: The GUID of the replicated folder (see the objectGUID specified in
	// section 2.3.7) in the specified connection's replication group.
	ContentSetID *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
}

func (o *EstablishSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_EstablishSessionOperation) *xxx_EstablishSessionOperation {
	if op == nil {
		op = &xxx_EstablishSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	op.ContentSetID = o.ContentSetID
	return op
}

func (o *EstablishSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_EstablishSessionOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
	o.ContentSetID = op.ContentSetID
}
func (o *EstablishSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EstablishSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EstablishSessionResponse structure represents the EstablishSession operation response
type EstablishSessionResponse struct {
	// Return: The EstablishSession return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EstablishSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_EstablishSessionOperation) *xxx_EstablishSessionOperation {
	if op == nil {
		op = &xxx_EstablishSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *EstablishSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_EstablishSessionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EstablishSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EstablishSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RequestUpdatesOperation structure represents the RequestUpdates operation
type xxx_RequestUpdatesOperation struct {
	ConnectionID           *frs2.ConnectionID     `idl:"name:connectionId" json:"connection_id"`
	ContentSetID           *frs2.ContentSetID     `idl:"name:contentSetId" json:"content_set_id"`
	CreditsAvailable       uint32                 `idl:"name:creditsAvailable" json:"credits_available"`
	HashRequested          int32                  `idl:"name:hashRequested" json:"hash_requested"`
	UpdateRequestType      frs2.UpdateRequestType `idl:"name:updateRequestType" json:"update_request_type"`
	VersionVectorDiffCount uint32                 `idl:"name:versionVectorDiffCount" json:"version_vector_diff_count"`
	VersionVectorDiff      []*frs2.VersionVector  `idl:"name:versionVectorDiff;size_is:(versionVectorDiffCount)" json:"version_vector_diff"`
	Update                 []*frs2.Update         `idl:"name:frsUpdate;size_is:(creditsAvailable);length_is:(updateCount)" json:"update"`
	UpdateCount            uint32                 `idl:"name:updateCount" json:"update_count"`
	UpdateStatus           frs2.UpdateStatus      `idl:"name:updateStatus" json:"update_status"`
	GVSNDBGUID             *dtyp.GUID             `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	GVSNVersion            uint64                 `idl:"name:gvsnVersion" json:"gvsn_version"`
	Return                 uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_RequestUpdatesOperation) OpNum() int { return 3 }

func (o *xxx_RequestUpdatesOperation) OpName() string { return "/FrsTransport/v1/RequestUpdates" }

func (o *xxx_RequestUpdatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.VersionVectorDiff != nil && o.VersionVectorDiffCount == 0 {
		o.VersionVectorDiffCount = uint32(len(o.VersionVectorDiff))
	}
	if o.CreditsAvailable > uint32(256) {
		return fmt.Errorf("CreditsAvailable is out of range")
	}
	if o.HashRequested < int32(0) || o.HashRequested > int32(1) {
		return fmt.Errorf("HashRequested is out of range")
	}
	if o.UpdateRequestType > frs2.UpdateRequestType(2) {
		return fmt.Errorf("UpdateRequestType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestUpdatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID != nil {
			if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ContentSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// creditsAvailable {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CreditsAvailable); err != nil {
			return err
		}
	}
	// hashRequested {in} (1:{range=(0,1)}(int32))
	{
		if err := w.WriteData(o.HashRequested); err != nil {
			return err
		}
	}
	// updateRequestType {in} (1:{range=(0,2), alias=UPDATE_REQUEST_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.UpdateRequestType)); err != nil {
			return err
		}
	}
	// versionVectorDiffCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.VersionVectorDiffCount); err != nil {
			return err
		}
	}
	// versionVectorDiff {in} (1:{pointer=ref}*(1))(2:{alias=FRS_VERSION_VECTOR}[dim:0,size_is=versionVectorDiffCount](struct))
	{
		dimSize1 := uint64(o.VersionVectorDiffCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.VersionVectorDiff {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.VersionVectorDiff[i1] != nil {
				if err := o.VersionVectorDiff[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&frs2.VersionVector{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.VersionVectorDiff); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&frs2.VersionVector{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RequestUpdatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID == nil {
			o.ContentSetID = &frs2.ContentSetID{}
		}
		if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// creditsAvailable {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CreditsAvailable); err != nil {
			return err
		}
	}
	// hashRequested {in} (1:{range=(0,1)}(int32))
	{
		if err := w.ReadData(&o.HashRequested); err != nil {
			return err
		}
	}
	// updateRequestType {in} (1:{range=(0,2), alias=UPDATE_REQUEST_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UpdateRequestType)); err != nil {
			return err
		}
	}
	// versionVectorDiffCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.VersionVectorDiffCount); err != nil {
			return err
		}
	}
	// versionVectorDiff {in} (1:{pointer=ref}*(1))(2:{alias=FRS_VERSION_VECTOR}[dim:0,size_is=versionVectorDiffCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.VersionVectorDiff", sizeInfo[0])
		}
		o.VersionVectorDiff = make([]*frs2.VersionVector, sizeInfo[0])
		for i1 := range o.VersionVectorDiff {
			i1 := i1
			if o.VersionVectorDiff[i1] == nil {
				o.VersionVectorDiff[i1] = &frs2.VersionVector{}
			}
			if err := o.VersionVectorDiff[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RequestUpdatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Update != nil && o.UpdateCount == 0 {
		o.UpdateCount = uint32(len(o.Update))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestUpdatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// frsUpdate {out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}[dim:0,size_is=creditsAvailable,length_is=updateCount](struct))
	{
		dimSize1 := uint64(o.CreditsAvailable)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.UpdateCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Update {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Update[i1] != nil {
				if err := o.Update[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&frs2.Update{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Update); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&frs2.Update{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// updateCount {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UpdateCount); err != nil {
			return err
		}
	}
	// updateStatus {out} (1:{pointer=ref}*(1))(2:{alias=UPDATE_STATUS}(enum))
	{
		if err := w.WriteEnum(uint16(o.UpdateStatus)); err != nil {
			return err
		}
	}
	// gvsnDbGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GVSNDBGUID != nil {
			if err := o.GVSNDBGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// gvsnVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.GVSNVersion); err != nil {
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

func (o *xxx_RequestUpdatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// frsUpdate {out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}[dim:0,size_is=creditsAvailable,length_is=updateCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Update", sizeInfo[0])
		}
		o.Update = make([]*frs2.Update, sizeInfo[0])
		for i1 := range o.Update {
			i1 := i1
			if o.Update[i1] == nil {
				o.Update[i1] = &frs2.Update{}
			}
			if err := o.Update[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// updateCount {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UpdateCount); err != nil {
			return err
		}
	}
	// updateStatus {out} (1:{pointer=ref}*(1))(2:{alias=UPDATE_STATUS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.UpdateStatus)); err != nil {
			return err
		}
	}
	// gvsnDbGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GVSNDBGUID == nil {
			o.GVSNDBGUID = &dtyp.GUID{}
		}
		if err := o.GVSNDBGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// gvsnVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.GVSNVersion); err != nil {
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

// RequestUpdatesRequest structure represents the RequestUpdates operation request
type RequestUpdatesRequest struct {
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// contentSetId: The GUID of the replicated folder (see the objectGUID attribute specified
	// in section 2.3.7) in the specified connection's replication group.
	ContentSetID *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	// creditsAvailable: The maximum number of updates that the client can receive in the
	// frsUpdate buffer.
	CreditsAvailable uint32 `idl:"name:creditsAvailable" json:"credits_available"`
	// hashRequested: The client sets the hashRequested parameter to TRUE to request that
	// the server compute the hash (see the hash field of the FRS_UPDATE structure specified
	// in section 2.2.1.4.4) for each update that it sends, or FALSE if the hashes are not
	// desired. The server SHOULD compute hashes when hashes are requested, although it
	// is not required to do so. Computing a file's hash requires DFS-R to read the file's
	// data. It is possible that another process has already opened the file for exclusive
	// access, which prevents DFS-R from computing the file hash. In this scenario, the
	// DFS-R server does not compute the hash even if the client requested that it does.
	HashRequested int32 `idl:"name:hashRequested" json:"hash_requested"`
	// updateRequestType: The value from the UPDATE_REQUEST_TYPE enumeration that indicates
	// the type of replication updates requested.
	UpdateRequestType frs2.UpdateRequestType `idl:"name:updateRequestType" json:"update_request_type"`
	// versionVectorDiffCount: The number of items specified in the versionVectorDiff parameter.
	VersionVectorDiffCount uint32 `idl:"name:versionVectorDiffCount" json:"version_vector_diff_count"`
	// versionVectorDiff: The set of FRS_VERSION_VECTOR structures that specifies what updates
	// the client requires from the server. This parameter specifies the difference between
	// the client's version vector and the client's most recent copy of the server's version
	// vector obtained from a previous call to the RequestVersionVector method.
	VersionVectorDiff []*frs2.VersionVector `idl:"name:versionVectorDiff;size_is:(versionVectorDiffCount)" json:"version_vector_diff"`
}

func (o *RequestUpdatesRequest) xxx_ToOp(ctx context.Context, op *xxx_RequestUpdatesOperation) *xxx_RequestUpdatesOperation {
	if op == nil {
		op = &xxx_RequestUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	op.ContentSetID = o.ContentSetID
	op.CreditsAvailable = o.CreditsAvailable
	op.HashRequested = o.HashRequested
	op.UpdateRequestType = o.UpdateRequestType
	op.VersionVectorDiffCount = o.VersionVectorDiffCount
	op.VersionVectorDiff = o.VersionVectorDiff
	return op
}

func (o *RequestUpdatesRequest) xxx_FromOp(ctx context.Context, op *xxx_RequestUpdatesOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
	o.ContentSetID = op.ContentSetID
	o.CreditsAvailable = op.CreditsAvailable
	o.HashRequested = op.HashRequested
	o.UpdateRequestType = op.UpdateRequestType
	o.VersionVectorDiffCount = op.VersionVectorDiffCount
	o.VersionVectorDiff = op.VersionVectorDiff
}
func (o *RequestUpdatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RequestUpdatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestUpdatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RequestUpdatesResponse structure represents the RequestUpdates operation response
type RequestUpdatesResponse struct {
	// XXX: creditsAvailable is an implicit input depedency for output parameters
	CreditsAvailable uint32 `idl:"name:creditsAvailable" json:"credits_available"`

	// frsUpdate: The set of FRS_UPDATE structures that describes the update that occurred
	// to each of the files to be replicated.
	Update []*frs2.Update `idl:"name:frsUpdate;size_is:(creditsAvailable);length_is:(updateCount)" json:"update"`
	// updateCount: The number of updates that the server wrote into the frsUpdate buffer.
	UpdateCount uint32 `idl:"name:updateCount" json:"update_count"`
	// updateStatus: The value from the UPDATE_STATUS enumeration that specifies if all
	// of the requested updates have been sent by the server.
	UpdateStatus frs2.UpdateStatus `idl:"name:updateStatus" json:"update_status"`
	// gvsnDbGuid: The GVSN GUID (as specified in [MS-DTYP] section 2.3.4) for the last
	// field in the versionVectorDiff that was processed.
	GVSNDBGUID *dtyp.GUID `idl:"name:gvsnDbGuid" json:"gvsn_db_guid"`
	// gvsnVersion: The version of the gvsnDbGuid.
	GVSNVersion uint64 `idl:"name:gvsnVersion" json:"gvsn_version"`
	// Return: The RequestUpdates return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RequestUpdatesResponse) xxx_ToOp(ctx context.Context, op *xxx_RequestUpdatesOperation) *xxx_RequestUpdatesOperation {
	if op == nil {
		op = &xxx_RequestUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.CreditsAvailable == uint32(0) {
		op.CreditsAvailable = o.CreditsAvailable
	}

	op.Update = o.Update
	op.UpdateCount = o.UpdateCount
	op.UpdateStatus = o.UpdateStatus
	op.GVSNDBGUID = o.GVSNDBGUID
	op.GVSNVersion = o.GVSNVersion
	op.Return = o.Return
	return op
}

func (o *RequestUpdatesResponse) xxx_FromOp(ctx context.Context, op *xxx_RequestUpdatesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.CreditsAvailable = op.CreditsAvailable

	o.Update = op.Update
	o.UpdateCount = op.UpdateCount
	o.UpdateStatus = op.UpdateStatus
	o.GVSNDBGUID = op.GVSNDBGUID
	o.GVSNVersion = op.GVSNVersion
	o.Return = op.Return
}
func (o *RequestUpdatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RequestUpdatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestUpdatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RequestVersionVectorOperation structure represents the RequestVersionVector operation
type xxx_RequestVersionVectorOperation struct {
	SequenceNumber uint32                  `idl:"name:sequenceNumber" json:"sequence_number"`
	ConnectionID   *frs2.ConnectionID      `idl:"name:connectionId" json:"connection_id"`
	ContentSetID   *frs2.ContentSetID      `idl:"name:contentSetId" json:"content_set_id"`
	RequestType    frs2.VersionRequestType `idl:"name:requestType" json:"request_type"`
	ChangeType     frs2.VersionChangeType  `idl:"name:changeType" json:"change_type"`
	Generation     uint64                  `idl:"name:vvGeneration" json:"generation"`
	Return         uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RequestVersionVectorOperation) OpNum() int { return 4 }

func (o *xxx_RequestVersionVectorOperation) OpName() string {
	return "/FrsTransport/v1/RequestVersionVector"
}

func (o *xxx_RequestVersionVectorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestType > frs2.VersionRequestType(2) {
		return fmt.Errorf("RequestType is out of range")
	}
	if o.ChangeType > frs2.VersionChangeType(2) {
		return fmt.Errorf("ChangeType is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestVersionVectorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sequenceNumber {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SequenceNumber); err != nil {
			return err
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID != nil {
			if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ContentSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// requestType {in} (1:{range=(0,2), alias=VERSION_REQUEST_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.RequestType)); err != nil {
			return err
		}
	}
	// changeType {in} (1:{range=(0,2), alias=VERSION_CHANGE_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.ChangeType)); err != nil {
			return err
		}
	}
	// vvGeneration {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Generation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestVersionVectorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sequenceNumber {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SequenceNumber); err != nil {
			return err
		}
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID == nil {
			o.ContentSetID = &frs2.ContentSetID{}
		}
		if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// requestType {in} (1:{range=(0,2), alias=VERSION_REQUEST_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.RequestType)); err != nil {
			return err
		}
	}
	// changeType {in} (1:{range=(0,2), alias=VERSION_CHANGE_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ChangeType)); err != nil {
			return err
		}
	}
	// vvGeneration {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Generation); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestVersionVectorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestVersionVectorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RequestVersionVectorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RequestVersionVectorRequest structure represents the RequestVersionVector operation request
type RequestVersionVectorRequest struct {
	// sequenceNumber: The sequence number for this request. The sequence number is used
	// to pair the version vector request with the asynchronous response in AsyncPoll. During
	// a given session, the client SHOULD supply a unique sequence number for each call
	// to this function or else they will not be able to match server responses via the
	// AsyncPoll method to the original version vector request.
	SequenceNumber uint32 `idl:"name:sequenceNumber" json:"sequence_number"`
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// contentSetId: The GUID of the replicated folder (see the objectGUID attribute specified
	// in section 2.3.7) in the specified connection's replication group.
	ContentSetID *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	// requestType: The value from the VERSION_REQUEST_TYPE enumeration that describes the
	// type of replication sync to perform.
	RequestType frs2.VersionRequestType `idl:"name:requestType" json:"request_type"`
	// changeType: The value from the VERSION_CHANGE_TYPE enumeration that indicates whether
	// to notify change only or send the entire version chain vector.
	ChangeType frs2.VersionChangeType `idl:"name:changeType" json:"change_type"`
	// vvGeneration: The vvGeneration parameter is used to calibrate what incarnation of
	// the server's version chain vector is known to the client. The client supplies the
	// last generation number that it received from the server when the requestType parameter
	// is set to REQUEST_NORMAL_SYNC. Otherwise the client MUST supply zero.
	Generation uint64 `idl:"name:vvGeneration" json:"generation"`
}

func (o *RequestVersionVectorRequest) xxx_ToOp(ctx context.Context, op *xxx_RequestVersionVectorOperation) *xxx_RequestVersionVectorOperation {
	if op == nil {
		op = &xxx_RequestVersionVectorOperation{}
	}
	if o == nil {
		return op
	}
	op.SequenceNumber = o.SequenceNumber
	op.ConnectionID = o.ConnectionID
	op.ContentSetID = o.ContentSetID
	op.RequestType = o.RequestType
	op.ChangeType = o.ChangeType
	op.Generation = o.Generation
	return op
}

func (o *RequestVersionVectorRequest) xxx_FromOp(ctx context.Context, op *xxx_RequestVersionVectorOperation) {
	if o == nil {
		return
	}
	o.SequenceNumber = op.SequenceNumber
	o.ConnectionID = op.ConnectionID
	o.ContentSetID = op.ContentSetID
	o.RequestType = op.RequestType
	o.ChangeType = op.ChangeType
	o.Generation = op.Generation
}
func (o *RequestVersionVectorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RequestVersionVectorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestVersionVectorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RequestVersionVectorResponse structure represents the RequestVersionVector operation response
type RequestVersionVectorResponse struct {
	// Return: The RequestVersionVector return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RequestVersionVectorResponse) xxx_ToOp(ctx context.Context, op *xxx_RequestVersionVectorOperation) *xxx_RequestVersionVectorOperation {
	if op == nil {
		op = &xxx_RequestVersionVectorOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RequestVersionVectorResponse) xxx_FromOp(ctx context.Context, op *xxx_RequestVersionVectorOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RequestVersionVectorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RequestVersionVectorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestVersionVectorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AsyncPollOperation structure represents the AsyncPoll operation
type xxx_AsyncPollOperation struct {
	ConnectionID *frs2.ConnectionID         `idl:"name:connectionId" json:"connection_id"`
	Response     *frs2.AsyncResponseContext `idl:"name:response" json:"response"`
	Return       uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_AsyncPollOperation) OpNum() int { return 5 }

func (o *xxx_AsyncPollOperation) OpName() string { return "/FrsTransport/v1/AsyncPoll" }

func (o *xxx_AsyncPollOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncPollOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AsyncPollOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncPollOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncPollOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// response {out} (1:{pointer=ref}*(1))(2:{alias=FRS_ASYNC_RESPONSE_CONTEXT}(struct))
	{
		if o.Response != nil {
			if err := o.Response.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.AsyncResponseContext{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AsyncPollOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// response {out} (1:{pointer=ref}*(1))(2:{alias=FRS_ASYNC_RESPONSE_CONTEXT}(struct))
	{
		if o.Response == nil {
			o.Response = &frs2.AsyncResponseContext{}
		}
		if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
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

// AsyncPollRequest structure represents the AsyncPoll operation request
type AsyncPollRequest struct {
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
}

func (o *AsyncPollRequest) xxx_ToOp(ctx context.Context, op *xxx_AsyncPollOperation) *xxx_AsyncPollOperation {
	if op == nil {
		op = &xxx_AsyncPollOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	return op
}

func (o *AsyncPollRequest) xxx_FromOp(ctx context.Context, op *xxx_AsyncPollOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
}
func (o *AsyncPollRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AsyncPollRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AsyncPollOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AsyncPollResponse structure represents the AsyncPoll operation response
type AsyncPollResponse struct {
	// response: The FRS_ASYNC_RESPONSE_CONTEXT structure that contains the context for
	// the requested poll.
	Response *frs2.AsyncResponseContext `idl:"name:response" json:"response"`
	// Return: The AsyncPoll return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AsyncPollResponse) xxx_ToOp(ctx context.Context, op *xxx_AsyncPollOperation) *xxx_AsyncPollOperation {
	if op == nil {
		op = &xxx_AsyncPollOperation{}
	}
	if o == nil {
		return op
	}
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *AsyncPollResponse) xxx_FromOp(ctx context.Context, op *xxx_AsyncPollOperation) {
	if o == nil {
		return
	}
	o.Response = op.Response
	o.Return = op.Return
}
func (o *AsyncPollResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AsyncPollResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AsyncPollOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RequestRecordsOperation structure represents the RequestRecords operation
type xxx_RequestRecordsOperation struct {
	ConnectionID      *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	ContentSetID      *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	UIDDBGUID         *frs2.DatabaseID   `idl:"name:uidDbGuid" json:"uid_db_guid"`
	UIDVersion        uint64             `idl:"name:uidVersion" json:"uid_version"`
	MaxRecords        uint32             `idl:"name:maxRecords" json:"max_records"`
	RecordsLength     uint32             `idl:"name:numRecords" json:"records_length"`
	BytesLength       uint32             `idl:"name:numBytes" json:"bytes_length"`
	CompressedRecords []byte             `idl:"name:compressedRecords;size_is:(, numBytes)" json:"compressed_records"`
	RecordsStatus     frs2.RecordsStatus `idl:"name:recordsStatus" json:"records_status"`
	Return            uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_RequestRecordsOperation) OpNum() int { return 6 }

func (o *xxx_RequestRecordsOperation) OpName() string { return "/FrsTransport/v1/RequestRecords" }

func (o *xxx_RequestRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID != nil {
			if err := o.ContentSetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ContentSetID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// uidDbGuid {in} (1:{alias=FRS_DATABASE_ID, names=GUID}(struct))
	{
		if o.UIDDBGUID != nil {
			if err := o.UIDDBGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.DatabaseID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// uidVersion {in} (1:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.UIDVersion); err != nil {
			return err
		}
	}
	// maxRecords {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxRecords); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// contentSetId {in} (1:{alias=FRS_CONTENT_SET_ID, names=GUID}(struct))
	{
		if o.ContentSetID == nil {
			o.ContentSetID = &frs2.ContentSetID{}
		}
		if err := o.ContentSetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// uidDbGuid {in} (1:{alias=FRS_DATABASE_ID, names=GUID}(struct))
	{
		if o.UIDDBGUID == nil {
			o.UIDDBGUID = &frs2.DatabaseID{}
		}
		if err := o.UIDDBGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// uidVersion {in} (1:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.UIDVersion); err != nil {
			return err
		}
	}
	// maxRecords {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxRecords); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.CompressedRecords != nil && o.BytesLength == 0 {
		o.BytesLength = uint32(len(o.CompressedRecords))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// maxRecords {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxRecords); err != nil {
			return err
		}
	}
	// numRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RecordsLength); err != nil {
			return err
		}
	}
	// numBytes {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesLength); err != nil {
			return err
		}
	}
	// compressedRecords {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=numBytes](uint8))
	{
		if o.CompressedRecords != nil || o.BytesLength > 0 {
			_ptr_compressedRecords := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BytesLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.CompressedRecords {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.CompressedRecords[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.CompressedRecords); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CompressedRecords, _ptr_compressedRecords); err != nil {
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
	// recordsStatus {out} (1:{pointer=ref}*(1))(2:{alias=RECORDS_STATUS}(enum))
	{
		if err := w.WriteEnum(uint16(o.RecordsStatus)); err != nil {
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

func (o *xxx_RequestRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// maxRecords {in, out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxRecords); err != nil {
			return err
		}
	}
	// numRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RecordsLength); err != nil {
			return err
		}
	}
	// numBytes {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesLength); err != nil {
			return err
		}
	}
	// compressedRecords {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=numBytes](uint8))
	{
		_ptr_compressedRecords := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.CompressedRecords", sizeInfo[0])
			}
			o.CompressedRecords = make([]byte, sizeInfo[0])
			for i1 := range o.CompressedRecords {
				i1 := i1
				if err := w.ReadData(&o.CompressedRecords[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_compressedRecords := func(ptr interface{}) { o.CompressedRecords = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.CompressedRecords, _s_compressedRecords, _ptr_compressedRecords); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// recordsStatus {out} (1:{pointer=ref}*(1))(2:{alias=RECORDS_STATUS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.RecordsStatus)); err != nil {
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

// RequestRecordsRequest structure represents the RequestRecords operation request
type RequestRecordsRequest struct {
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// contentSetId: The GUID of the replicated folder (see the objectGUID attribute specified
	// in section 2.3.7) in the specified connection's replication group.
	ContentSetID *frs2.ContentSetID `idl:"name:contentSetId" json:"content_set_id"`
	// uidDbGuid: A UID database GUID. This parameter, along with the uidVersion parameter,
	// specifies an iterator into the server's records. A value of zero specifies a request
	// for all of a replicated folder's records from the server.
	UIDDBGUID *frs2.DatabaseID `idl:"name:uidDbGuid" json:"uid_db_guid"`
	// uidVersion: A UID version. The parameter, along with uidDbGuid parameter, specifies
	// an iterator into the server's records. A value of zero specifies a request for all
	// of a replicated folder's records from the server.
	UIDVersion uint64 `idl:"name:uidVersion" json:"uid_version"`
	// maxRecords: The maximum number of records that the server can send to the client.
	// The server returns the lesser of the client-specified value and the maximum number
	// of records that the server is capable of sending.<26>
	MaxRecords uint32 `idl:"name:maxRecords" json:"max_records"`
}

func (o *RequestRecordsRequest) xxx_ToOp(ctx context.Context, op *xxx_RequestRecordsOperation) *xxx_RequestRecordsOperation {
	if op == nil {
		op = &xxx_RequestRecordsOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	op.ContentSetID = o.ContentSetID
	op.UIDDBGUID = o.UIDDBGUID
	op.UIDVersion = o.UIDVersion
	op.MaxRecords = o.MaxRecords
	return op
}

func (o *RequestRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_RequestRecordsOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
	o.ContentSetID = op.ContentSetID
	o.UIDDBGUID = op.UIDDBGUID
	o.UIDVersion = op.UIDVersion
	o.MaxRecords = op.MaxRecords
}
func (o *RequestRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RequestRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RequestRecordsResponse structure represents the RequestRecords operation response
type RequestRecordsResponse struct {
	// maxRecords: The maximum number of records that the server can send to the client.
	// The server returns the lesser of the client-specified value and the maximum number
	// of records that the server is capable of sending.<26>
	MaxRecords uint32 `idl:"name:maxRecords" json:"max_records"`
	// numRecords: The number of records written into the compressedRecords buffer by the
	// server.
	RecordsLength uint32 `idl:"name:numRecords" json:"records_length"`
	// numBytes: The size, in bytes, of the compressedRecords buffer.
	BytesLength uint32 `idl:"name:numBytes" json:"bytes_length"`
	// compressedRecords: The data records, compressed using the algorithm specified in
	// section 3.1.1.1.
	CompressedRecords []byte `idl:"name:compressedRecords;size_is:(, numBytes)" json:"compressed_records"`
	// recordsStatus: The value from the RECORDS_STATUS enumeration that indicates whether
	// more update records are available.
	RecordsStatus frs2.RecordsStatus `idl:"name:recordsStatus" json:"records_status"`
	// Return: The RequestRecords return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RequestRecordsResponse) xxx_ToOp(ctx context.Context, op *xxx_RequestRecordsOperation) *xxx_RequestRecordsOperation {
	if op == nil {
		op = &xxx_RequestRecordsOperation{}
	}
	if o == nil {
		return op
	}
	op.MaxRecords = o.MaxRecords
	op.RecordsLength = o.RecordsLength
	op.BytesLength = o.BytesLength
	op.CompressedRecords = o.CompressedRecords
	op.RecordsStatus = o.RecordsStatus
	op.Return = o.Return
	return op
}

func (o *RequestRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_RequestRecordsOperation) {
	if o == nil {
		return
	}
	o.MaxRecords = op.MaxRecords
	o.RecordsLength = op.RecordsLength
	o.BytesLength = op.BytesLength
	o.CompressedRecords = op.CompressedRecords
	o.RecordsStatus = op.RecordsStatus
	o.Return = op.Return
}
func (o *RequestRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RequestRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateCancelOperation structure represents the UpdateCancel operation
type xxx_UpdateCancelOperation struct {
	ConnectionID *frs2.ConnectionID     `idl:"name:connectionId" json:"connection_id"`
	CancelData   *frs2.UpdateCancelData `idl:"name:cancelData" json:"cancel_data"`
	Return       uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateCancelOperation) OpNum() int { return 7 }

func (o *xxx_UpdateCancelOperation) OpName() string { return "/FrsTransport/v1/UpdateCancel" }

func (o *xxx_UpdateCancelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateCancelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cancelData {in} (1:{alias=FRS_UPDATE_CANCEL_DATA}(struct))
	{
		if o.CancelData != nil {
			if err := o.CancelData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.UpdateCancelData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UpdateCancelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cancelData {in} (1:{alias=FRS_UPDATE_CANCEL_DATA}(struct))
	{
		if o.CancelData == nil {
			o.CancelData = &frs2.UpdateCancelData{}
		}
		if err := o.CancelData.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateCancelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateCancelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_UpdateCancelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateCancelRequest structure represents the UpdateCancel operation request
type UpdateCancelRequest struct {
	ConnectionID *frs2.ConnectionID     `idl:"name:connectionId" json:"connection_id"`
	CancelData   *frs2.UpdateCancelData `idl:"name:cancelData" json:"cancel_data"`
}

func (o *UpdateCancelRequest) xxx_ToOp(ctx context.Context, op *xxx_UpdateCancelOperation) *xxx_UpdateCancelOperation {
	if op == nil {
		op = &xxx_UpdateCancelOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	op.CancelData = o.CancelData
	return op
}

func (o *UpdateCancelRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateCancelOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
	o.CancelData = op.CancelData
}
func (o *UpdateCancelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpdateCancelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateCancelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateCancelResponse structure represents the UpdateCancel operation response
type UpdateCancelResponse struct {
	// Return: The UpdateCancel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UpdateCancelResponse) xxx_ToOp(ctx context.Context, op *xxx_UpdateCancelOperation) *xxx_UpdateCancelOperation {
	if op == nil {
		op = &xxx_UpdateCancelOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *UpdateCancelResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateCancelOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateCancelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpdateCancelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateCancelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RawGetFileDataOperation structure represents the RawGetFileData operation
type xxx_RawGetFileDataOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	DataBuffer    []byte         `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeRead)" json:"data_buffer"`
	BufferSize    uint32         `idl:"name:bufferSize" json:"buffer_size"`
	SizeRead      uint32         `idl:"name:sizeRead" json:"size_read"`
	IsEndOfFile   int32          `idl:"name:isEndOfFile" json:"is_end_of_file"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_RawGetFileDataOperation) OpNum() int { return 8 }

func (o *xxx_RawGetFileDataOperation) OpName() string { return "/FrsTransport/v1/RawGetFileData" }

func (o *xxx_RawGetFileDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BufferSize > uint32(262144) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DataBuffer != nil && o.SizeRead == 0 {
		o.SizeRead = uint32(len(o.DataBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeRead](uint8))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SizeRead)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.DataBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeRead); err != nil {
			return err
		}
	}
	// isEndOfFile {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.IsEndOfFile); err != nil {
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

func (o *xxx_RawGetFileDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeRead](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DataBuffer", sizeInfo[0])
		}
		o.DataBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.DataBuffer {
			i1 := i1
			if err := w.ReadData(&o.DataBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeRead); err != nil {
			return err
		}
	}
	// isEndOfFile {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.IsEndOfFile); err != nil {
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

// RawGetFileDataRequest structure represents the RawGetFileData operation request
type RawGetFileDataRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a pointer to a server context that was retrieved
	// by a previously successful call to the InitializeFileTransferAsync method. The server
	// MUST NOT change the value of serverContext and then return the same serverContext
	// that was passed in.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// bufferSize: The size, in bytes, of dataBuffer.
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`
}

func (o *RawGetFileDataRequest) xxx_ToOp(ctx context.Context, op *xxx_RawGetFileDataOperation) *xxx_RawGetFileDataOperation {
	if op == nil {
		op = &xxx_RawGetFileDataOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	op.BufferSize = o.BufferSize
	return op
}

func (o *RawGetFileDataRequest) xxx_FromOp(ctx context.Context, op *xxx_RawGetFileDataOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
	o.BufferSize = op.BufferSize
}
func (o *RawGetFileDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RawGetFileDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawGetFileDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RawGetFileDataResponse structure represents the RawGetFileData operation response
type RawGetFileDataResponse struct {
	// XXX: bufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`

	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a pointer to a server context that was retrieved
	// by a previously successful call to the InitializeFileTransferAsync method. The server
	// MUST NOT change the value of serverContext and then return the same serverContext
	// that was passed in.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// dataBuffer: The file data received from the server.
	DataBuffer []byte `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeRead)" json:"data_buffer"`
	// sizeRead: The size, in bytes, of the file data returned in dataBuffer.
	SizeRead uint32 `idl:"name:sizeRead" json:"size_read"`
	// isEndOfFile: The value is TRUE if the end of the specified file has been reached
	// and there is no more file data to replicate to the client; otherwise, the value is
	// FALSE.
	IsEndOfFile int32 `idl:"name:isEndOfFile" json:"is_end_of_file"`
	// Return: The RawGetFileData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RawGetFileDataResponse) xxx_ToOp(ctx context.Context, op *xxx_RawGetFileDataOperation) *xxx_RawGetFileDataOperation {
	if op == nil {
		op = &xxx_RawGetFileDataOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.ServerContext = o.ServerContext
	op.DataBuffer = o.DataBuffer
	op.SizeRead = o.SizeRead
	op.IsEndOfFile = o.IsEndOfFile
	op.Return = o.Return
	return op
}

func (o *RawGetFileDataResponse) xxx_FromOp(ctx context.Context, op *xxx_RawGetFileDataOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.ServerContext = op.ServerContext
	o.DataBuffer = op.DataBuffer
	o.SizeRead = op.SizeRead
	o.IsEndOfFile = op.IsEndOfFile
	o.Return = op.Return
}
func (o *RawGetFileDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RawGetFileDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawGetFileDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSignaturesOperation structure represents the RdcGetSignatures operation
type xxx_GetSignaturesOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	Level         uint8          `idl:"name:level" json:"level"`
	Offset        uint64         `idl:"name:offset" json:"offset"`
	Buffer        []byte         `idl:"name:buffer;size_is:(length);length_is:(sizeRead)" json:"buffer"`
	Length        uint32         `idl:"name:length" json:"length"`
	SizeRead      uint32         `idl:"name:sizeRead" json:"size_read"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSignaturesOperation) OpNum() int { return 9 }

func (o *xxx_GetSignaturesOperation) OpName() string { return "/FrsTransport/v1/RdcGetSignatures" }

func (o *xxx_GetSignaturesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Level < uint8(1) || o.Level > uint8(8) {
		return fmt.Errorf("Level is out of range")
	}
	if o.Length < uint32(1) || o.Length > uint32(65536) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignaturesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// level {in} (1:{range=(1,8)}(uint8))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// offset {in} (1:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	// length {in} (1:{range=(1,65536), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignaturesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// level {in} (1:{range=(1,8)}(uint8))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// offset {in} (1:{alias=DWORDLONG, names=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	// length {in} (1:{range=(1,65536), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Length); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignaturesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.SizeRead == 0 {
		o.SizeRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSignaturesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=length,length_is=sizeRead](uint8))
	{
		dimSize1 := uint64(o.Length)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SizeRead)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeRead); err != nil {
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

func (o *xxx_GetSignaturesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// buffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=length,length_is=sizeRead](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeRead); err != nil {
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

// GetSignaturesRequest structure represents the RdcGetSignatures operation request
type GetSignaturesRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to InitializeFileTransferAsync method in which the client set the
	// rdcDesired parameter to TRUE.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// level: The RDC recursion level being requested. A client MUST specify a number in
	// the range of 1 to x, where x is the value of the rdcSignatureLevels field of the
	// rdcInfo structure that was returned by the InitializeFileTransferAsync method call
	// associated with the specified server context.
	Level uint8 `idl:"name:level" json:"level"`
	// offset: The zero-based offset, in bytes, at which to retrieve data from the file.
	Offset uint64 `idl:"name:offset" json:"offset"`
	// length: The size, in bytes, of buffer.
	Length uint32 `idl:"name:length" json:"length"`
}

func (o *GetSignaturesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSignaturesOperation) *xxx_GetSignaturesOperation {
	if op == nil {
		op = &xxx_GetSignaturesOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	op.Level = o.Level
	op.Offset = o.Offset
	op.Length = o.Length
	return op
}

func (o *GetSignaturesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSignaturesOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
	o.Level = op.Level
	o.Offset = op.Offset
	o.Length = op.Length
}
func (o *GetSignaturesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSignaturesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSignaturesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSignaturesResponse structure represents the RdcGetSignatures operation response
type GetSignaturesResponse struct {
	// XXX: length is an implicit input depedency for output parameters
	Length uint32 `idl:"name:length" json:"length"`

	// buffer: The file signature data received from the server.
	Buffer []byte `idl:"name:buffer;size_is:(length);length_is:(sizeRead)" json:"buffer"`
	// sizeRead: The size, in bytes, of the file data returned in buffer.
	SizeRead uint32 `idl:"name:sizeRead" json:"size_read"`
	// Return: The RdcGetSignatures return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSignaturesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSignaturesOperation) *xxx_GetSignaturesOperation {
	if op == nil {
		op = &xxx_GetSignaturesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Length == uint32(0) {
		op.Length = o.Length
	}

	op.Buffer = o.Buffer
	op.SizeRead = o.SizeRead
	op.Return = o.Return
	return op
}

func (o *GetSignaturesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSignaturesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Length = op.Length

	o.Buffer = op.Buffer
	o.SizeRead = op.SizeRead
	o.Return = op.Return
}
func (o *GetSignaturesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSignaturesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSignaturesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PushSourceNeedsOperation structure represents the RdcPushSourceNeeds operation
type xxx_PushSourceNeedsOperation struct {
	ServerContext *ServerContext     `idl:"name:serverContext" json:"server_context"`
	SourceNeeds   []*frs2.SourceNeed `idl:"name:sourceNeeds;size_is:(needCount)" json:"source_needs"`
	NeedCount     uint32             `idl:"name:needCount" json:"need_count"`
	Return        uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_PushSourceNeedsOperation) OpNum() int { return 10 }

func (o *xxx_PushSourceNeedsOperation) OpName() string { return "/FrsTransport/v1/RdcPushSourceNeeds" }

func (o *xxx_PushSourceNeedsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SourceNeeds != nil && o.NeedCount == 0 {
		o.NeedCount = uint32(len(o.SourceNeeds))
	}
	if o.NeedCount > uint32(20) {
		return fmt.Errorf("NeedCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PushSourceNeedsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// sourceNeeds {in} (1:{pointer=ref}*(1))(2:{alias=FRS_RDC_SOURCE_NEED}[dim:0,size_is=needCount](struct))
	{
		dimSize1 := uint64(o.NeedCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SourceNeeds {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.SourceNeeds[i1] != nil {
				if err := o.SourceNeeds[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&frs2.SourceNeed{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.SourceNeeds); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&frs2.SourceNeed{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// needCount {in} (1:{range=(0,20), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NeedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PushSourceNeedsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// sourceNeeds {in} (1:{pointer=ref}*(1))(2:{alias=FRS_RDC_SOURCE_NEED}[dim:0,size_is=needCount](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SourceNeeds", sizeInfo[0])
		}
		o.SourceNeeds = make([]*frs2.SourceNeed, sizeInfo[0])
		for i1 := range o.SourceNeeds {
			i1 := i1
			if o.SourceNeeds[i1] == nil {
				o.SourceNeeds[i1] = &frs2.SourceNeed{}
			}
			if err := o.SourceNeeds[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// needCount {in} (1:{range=(0,20), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NeedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PushSourceNeedsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PushSourceNeedsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_PushSourceNeedsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PushSourceNeedsRequest structure represents the RdcPushSourceNeeds operation request
type PushSourceNeedsRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method in which the client set
	// the rdcDesired parameter to TRUE.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// sourceNeeds: The pointer to a set of FRS_RDC_SOURCE_NEED structures that indicate
	// the offsets and lengths of file data that is sent from the server to the client.
	SourceNeeds []*frs2.SourceNeed `idl:"name:sourceNeeds;size_is:(needCount)" json:"source_needs"`
	// needCount: The number of FRS_RDC_SOURCE_NEED structures pointed to by sourceNeeds.
	NeedCount uint32 `idl:"name:needCount" json:"need_count"`
}

func (o *PushSourceNeedsRequest) xxx_ToOp(ctx context.Context, op *xxx_PushSourceNeedsOperation) *xxx_PushSourceNeedsOperation {
	if op == nil {
		op = &xxx_PushSourceNeedsOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	op.SourceNeeds = o.SourceNeeds
	op.NeedCount = o.NeedCount
	return op
}

func (o *PushSourceNeedsRequest) xxx_FromOp(ctx context.Context, op *xxx_PushSourceNeedsOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
	o.SourceNeeds = op.SourceNeeds
	o.NeedCount = op.NeedCount
}
func (o *PushSourceNeedsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PushSourceNeedsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PushSourceNeedsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PushSourceNeedsResponse structure represents the RdcPushSourceNeeds operation response
type PushSourceNeedsResponse struct {
	// Return: The RdcPushSourceNeeds return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PushSourceNeedsResponse) xxx_ToOp(ctx context.Context, op *xxx_PushSourceNeedsOperation) *xxx_PushSourceNeedsOperation {
	if op == nil {
		op = &xxx_PushSourceNeedsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PushSourceNeedsResponse) xxx_FromOp(ctx context.Context, op *xxx_PushSourceNeedsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PushSourceNeedsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PushSourceNeedsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PushSourceNeedsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileDataOperation structure represents the RdcGetFileData operation
type xxx_GetFileDataOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	DataBuffer    []byte         `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeReturned)" json:"data_buffer"`
	BufferSize    uint32         `idl:"name:bufferSize" json:"buffer_size"`
	SizeReturned  uint32         `idl:"name:sizeReturned" json:"size_returned"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileDataOperation) OpNum() int { return 11 }

func (o *xxx_GetFileDataOperation) OpName() string { return "/FrsTransport/v1/RdcGetFileData" }

func (o *xxx_GetFileDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.BufferSize > uint32(262144) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DataBuffer != nil && o.SizeReturned == 0 {
		o.SizeReturned = uint32(len(o.DataBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeReturned](uint8))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SizeReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.DataBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// sizeReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeReturned); err != nil {
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

func (o *xxx_GetFileDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeReturned](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DataBuffer", sizeInfo[0])
		}
		o.DataBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.DataBuffer {
			i1 := i1
			if err := w.ReadData(&o.DataBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// sizeReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeReturned); err != nil {
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

// GetFileDataRequest structure represents the RdcGetFileData operation request
type GetFileDataRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method in which the client set
	// the rdcDesired parameter to TRUE.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// bufferSize: The size, in bytes, of dataBuffer
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`
}

func (o *GetFileDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileDataOperation) *xxx_GetFileDataOperation {
	if op == nil {
		op = &xxx_GetFileDataOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	op.BufferSize = o.BufferSize
	return op
}

func (o *GetFileDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileDataOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
	o.BufferSize = op.BufferSize
}
func (o *GetFileDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileDataResponse structure represents the RdcGetFileData operation response
type GetFileDataResponse struct {
	// XXX: bufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`

	// dataBuffer: The file data received from the server.
	DataBuffer []byte `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeReturned)" json:"data_buffer"`
	// sizeReturned: The size, in bytes, of the file data returned in dataBuffer.
	SizeReturned uint32 `idl:"name:sizeReturned" json:"size_returned"`
	// Return: The RdcGetFileData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetFileDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileDataOperation) *xxx_GetFileDataOperation {
	if op == nil {
		op = &xxx_GetFileDataOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.DataBuffer = o.DataBuffer
	op.SizeReturned = o.SizeReturned
	op.Return = o.Return
	return op
}

func (o *GetFileDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileDataOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.DataBuffer = op.DataBuffer
	o.SizeReturned = op.SizeReturned
	o.Return = op.Return
}
func (o *GetFileDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseOperation structure represents the RdcClose operation
type xxx_CloseOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 12 }

func (o *xxx_CloseOperation) OpName() string { return "/FrsTransport/v1/RdcClose" }

func (o *xxx_CloseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// serverContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
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

// CloseRequest structure represents the RdcClose operation request
type CloseRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	return op
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
}
func (o *CloseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResponse structure represents the RdcClose operation response
type CloseResponse struct {
	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// Return: The RdcClose return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	op.Return = o.Return
	return op
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
	o.Return = op.Return
}
func (o *CloseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeFileTransferAsyncOperation structure represents the InitializeFileTransferAsync operation
type xxx_InitializeFileTransferAsyncOperation struct {
	ConnectionID  *frs2.ConnectionID          `idl:"name:connectionId" json:"connection_id"`
	Update        *frs2.Update                `idl:"name:frsUpdate" json:"update"`
	Desired       int32                       `idl:"name:rdcDesired" json:"desired"`
	StagingPolicy frs2.RequestedStagingPolicy `idl:"name:stagingPolicy" json:"staging_policy"`
	ServerContext *ServerContext              `idl:"name:serverContext" json:"server_context"`
	FileInfo      *frs2.FileInfo              `idl:"name:rdcFileInfo" json:"file_info"`
	DataBuffer    []byte                      `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeRead)" json:"data_buffer"`
	BufferSize    uint32                      `idl:"name:bufferSize" json:"buffer_size"`
	SizeRead      uint32                      `idl:"name:sizeRead" json:"size_read"`
	IsEndOfFile   int32                       `idl:"name:isEndOfFile" json:"is_end_of_file"`
	Return        uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeFileTransferAsyncOperation) OpNum() int { return 13 }

func (o *xxx_InitializeFileTransferAsyncOperation) OpName() string {
	return "/FrsTransport/v1/InitializeFileTransferAsync"
}

func (o *xxx_InitializeFileTransferAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Desired < int32(0) || o.Desired > int32(1) {
		return fmt.Errorf("Desired is out of range")
	}
	if o.BufferSize > uint32(262144) {
		return fmt.Errorf("BufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeFileTransferAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID != nil {
			if err := o.ConnectionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.ConnectionID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// frsUpdate {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}(struct))
	{
		if o.Update != nil {
			if err := o.Update.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.Update{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// rdcDesired {in} (1:{range=(0,1)}(int32))
	{
		if err := w.WriteData(o.Desired); err != nil {
			return err
		}
	}
	// stagingPolicy {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_REQUESTED_STAGING_POLICY}(enum))
	{
		if err := w.WriteEnum(uint16(o.StagingPolicy)); err != nil {
			return err
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeFileTransferAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// connectionId {in} (1:{alias=FRS_CONNECTION_ID, names=GUID}(struct))
	{
		if o.ConnectionID == nil {
			o.ConnectionID = &frs2.ConnectionID{}
		}
		if err := o.ConnectionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// frsUpdate {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}(struct))
	{
		if o.Update == nil {
			o.Update = &frs2.Update{}
		}
		if err := o.Update.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// rdcDesired {in} (1:{range=(0,1)}(int32))
	{
		if err := w.ReadData(&o.Desired); err != nil {
			return err
		}
	}
	// stagingPolicy {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_REQUESTED_STAGING_POLICY}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.StagingPolicy)); err != nil {
			return err
		}
	}
	// bufferSize {in} (1:{range=(0,262144), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeFileTransferAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DataBuffer != nil && o.SizeRead == 0 {
		o.SizeRead = uint32(len(o.DataBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeFileTransferAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// frsUpdate {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}(struct))
	{
		if o.Update != nil {
			if err := o.Update.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&frs2.Update{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// stagingPolicy {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_REQUESTED_STAGING_POLICY}(enum))
	{
		if err := w.WriteEnum(uint16(o.StagingPolicy)); err != nil {
			return err
		}
	}
	// serverContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// rdcFileInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=FRS_RDC_FILEINFO}(struct))
	{
		if o.FileInfo != nil {
			_ptr_rdcFileInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileInfo != nil {
					if err := o.FileInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&frs2.FileInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileInfo, _ptr_rdcFileInfo); err != nil {
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
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeRead](uint8))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SizeRead)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.DataBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeRead); err != nil {
			return err
		}
	}
	// isEndOfFile {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.IsEndOfFile); err != nil {
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

func (o *xxx_InitializeFileTransferAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// frsUpdate {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_UPDATE}(struct))
	{
		if o.Update == nil {
			o.Update = &frs2.Update{}
		}
		if err := o.Update.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// stagingPolicy {in, out} (1:{pointer=ref}*(1))(2:{alias=FRS_REQUESTED_STAGING_POLICY}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.StagingPolicy)); err != nil {
			return err
		}
	}
	// serverContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// rdcFileInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=FRS_RDC_FILEINFO}(struct))
	{
		_ptr_rdcFileInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileInfo == nil {
				o.FileInfo = &frs2.FileInfo{}
			}
			if err := o.FileInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_rdcFileInfo := func(ptr interface{}) { o.FileInfo = *ptr.(**frs2.FileInfo) }
		if err := w.ReadPointer(&o.FileInfo, _s_rdcFileInfo, _ptr_rdcFileInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dataBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=bufferSize,length_is=sizeRead](uint8))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DataBuffer", sizeInfo[0])
		}
		o.DataBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.DataBuffer {
			i1 := i1
			if err := w.ReadData(&o.DataBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// sizeRead {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeRead); err != nil {
			return err
		}
	}
	// isEndOfFile {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.IsEndOfFile); err != nil {
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

// InitializeFileTransferAsyncRequest structure represents the InitializeFileTransferAsync operation request
type InitializeFileTransferAsyncRequest struct {
	// connectionId: The GUID of an outbound connection (see the objectGUID attribute specified
	// in section 2.3.11) that the client established by a previous call to the EstablishConnection
	// method.
	ConnectionID *frs2.ConnectionID `idl:"name:connectionId" json:"connection_id"`
	// frsUpdate:  The FRS_UPDATE structure that contains information about the file being
	// replicated. The fields for the UID in frsUpdate MUST be set to the UID of the file
	// to be downloaded. All other fields are cleared (zeroed out) or can have the values
	// provided by the server in the response to a RequestUpdates call. On return, all fields
	// of frsUpdate MUST contain the values that are held by the server.
	Update *frs2.Update `idl:"name:frsUpdate" json:"update"`
	// rdcDesired: The value is TRUE if RDC has to be used when replicating this file; otherwise,
	// the value is FALSE.
	Desired int32 `idl:"name:rdcDesired" json:"desired"`
	// stagingPolicy: The FRS_REQUESTED_STAGING_POLICY enumeration value that indicates
	// the type of staging requested.  If the client-supplied value of rdcDesired is TRUE
	// and the client-supplied value of stagingPolicy is SERVER_DEFAULT, then the server
	// MUST set stagingPolicy to STAGING_REQUIRED. If the client-supplied value of rdcDesired
	// is FALSE and the client-supplied value of stagingPolicy is STAGING_REQUIRED, then
	// the server MUST set stagingPolicy to STAGING_REQUIRED. If the client-supplied value
	// of rdcDesired is FALSE and the client-supplied value of stagingPolicy is RESTAGING_REQUIRED,
	// then the server MUST set stagingPolicy to RESTAGING_REQUIRED.
	StagingPolicy frs2.RequestedStagingPolicy `idl:"name:stagingPolicy" json:"staging_policy"`
	// bufferSize: The size, in bytes, of dataBuffer. CONFIG_TRANSPORT_MAX_BUFFER_SIZE is
	// 262,144.
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`
}

func (o *InitializeFileTransferAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_InitializeFileTransferAsyncOperation) *xxx_InitializeFileTransferAsyncOperation {
	if op == nil {
		op = &xxx_InitializeFileTransferAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.ConnectionID = o.ConnectionID
	op.Update = o.Update
	op.Desired = o.Desired
	op.StagingPolicy = o.StagingPolicy
	op.BufferSize = o.BufferSize
	return op
}

func (o *InitializeFileTransferAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeFileTransferAsyncOperation) {
	if o == nil {
		return
	}
	o.ConnectionID = op.ConnectionID
	o.Update = op.Update
	o.Desired = op.Desired
	o.StagingPolicy = op.StagingPolicy
	o.BufferSize = op.BufferSize
}
func (o *InitializeFileTransferAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitializeFileTransferAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeFileTransferAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeFileTransferAsyncResponse structure represents the InitializeFileTransferAsync operation response
type InitializeFileTransferAsyncResponse struct {
	// XXX: bufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:bufferSize" json:"buffer_size"`

	// frsUpdate:  The FRS_UPDATE structure that contains information about the file being
	// replicated. The fields for the UID in frsUpdate MUST be set to the UID of the file
	// to be downloaded. All other fields are cleared (zeroed out) or can have the values
	// provided by the server in the response to a RequestUpdates call. On return, all fields
	// of frsUpdate MUST contain the values that are held by the server.
	Update *frs2.Update `idl:"name:frsUpdate" json:"update"`
	// stagingPolicy: The FRS_REQUESTED_STAGING_POLICY enumeration value that indicates
	// the type of staging requested.  If the client-supplied value of rdcDesired is TRUE
	// and the client-supplied value of stagingPolicy is SERVER_DEFAULT, then the server
	// MUST set stagingPolicy to STAGING_REQUIRED. If the client-supplied value of rdcDesired
	// is FALSE and the client-supplied value of stagingPolicy is STAGING_REQUIRED, then
	// the server MUST set stagingPolicy to STAGING_REQUIRED. If the client-supplied value
	// of rdcDesired is FALSE and the client-supplied value of stagingPolicy is RESTAGING_REQUIRED,
	// then the server MUST set stagingPolicy to RESTAGING_REQUIRED.
	StagingPolicy frs2.RequestedStagingPolicy `idl:"name:stagingPolicy" json:"staging_policy"`
	// serverContext: The context handle that represents the requested file replication
	// operation.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	// rdcFileInfo: The FRS_RDC_FILEINFO structure that describes the file whose replication
	// is in progress.
	FileInfo *frs2.FileInfo `idl:"name:rdcFileInfo" json:"file_info"`
	// dataBuffer: The file data received from the server.
	DataBuffer []byte `idl:"name:dataBuffer;size_is:(bufferSize);length_is:(sizeRead)" json:"data_buffer"`
	// sizeRead: The size, in bytes, of the file data returned in dataBuffer.
	SizeRead uint32 `idl:"name:sizeRead" json:"size_read"`
	// isEndOfFile: The value is TRUE if the end of the specified file has been reached
	// and there is no more file data to replicate to the client; otherwise, the value is
	// FALSE.
	IsEndOfFile int32 `idl:"name:isEndOfFile" json:"is_end_of_file"`
	// Return: The InitializeFileTransferAsync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *InitializeFileTransferAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_InitializeFileTransferAsyncOperation) *xxx_InitializeFileTransferAsyncOperation {
	if op == nil {
		op = &xxx_InitializeFileTransferAsyncOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.Update = o.Update
	op.StagingPolicy = o.StagingPolicy
	op.ServerContext = o.ServerContext
	op.FileInfo = o.FileInfo
	op.DataBuffer = o.DataBuffer
	op.SizeRead = o.SizeRead
	op.IsEndOfFile = o.IsEndOfFile
	op.Return = o.Return
	return op
}

func (o *InitializeFileTransferAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeFileTransferAsyncOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.Update = op.Update
	o.StagingPolicy = op.StagingPolicy
	o.ServerContext = op.ServerContext
	o.FileInfo = op.FileInfo
	o.DataBuffer = op.DataBuffer
	o.SizeRead = op.SizeRead
	o.IsEndOfFile = op.IsEndOfFile
	o.Return = op.Return
}
func (o *InitializeFileTransferAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitializeFileTransferAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeFileTransferAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RawGetFileDataAsyncOperation structure represents the RawGetFileDataAsync operation
type xxx_RawGetFileDataAsyncOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	BytePipe      frs2.BytePipe  `idl:"name:bytePipe" json:"byte_pipe"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_RawGetFileDataAsyncOperation) OpNum() int { return 15 }

func (o *xxx_RawGetFileDataAsyncOperation) OpName() string {
	return "/FrsTransport/v1/RawGetFileDataAsync"
}

func (o *xxx_RawGetFileDataAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RawGetFileDataAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// bytePipe {out} (1:{pointer=ref}*(1))(2:{alias=BYTE_PIPE}(pipe)(uint8))
	{
		// marshal pipe bytePipe
		if o.BytePipe != nil {
			for _chunk := range o.BytePipe.Pipe() {
				dimSize1 := uint64(len(_chunk))
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range _chunk {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_chunk[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_chunk); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
		}
		if err := w.WriteSize(0); err != nil {
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

func (o *xxx_RawGetFileDataAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// bytePipe {out} (1:{pointer=ref}*(1))(2:{alias=BYTE_PIPE}(pipe)(uint8))
	{
		if o.BytePipe == nil {
			o.BytePipe = frs2.NewBytePipe()
		}
		for {
			var _chunk []byte
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _chunk", sizeInfo[0])
			}
			_chunk = make([]byte, sizeInfo[0])
			for i1 := range _chunk {
				i1 := i1
				if err := w.ReadData(&_chunk[i1]); err != nil {
					return err
				}
			}
			if len(_chunk) == 0 /* end */ {
				break
			}
			o.BytePipe.Append(_chunk)
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

// RawGetFileDataAsyncRequest structure represents the RawGetFileDataAsync operation request
type RawGetFileDataAsyncRequest struct {
	// XXX: bytePipe is an implicit input depedency for output parameters
	BytePipe frs2.BytePipe `idl:"name:bytePipe" json:"byte_pipe"`

	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
}

func (o *RawGetFileDataAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_RawGetFileDataAsyncOperation) *xxx_RawGetFileDataAsyncOperation {
	if op == nil {
		op = &xxx_RawGetFileDataAsyncOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BytePipe == nil {
		op.BytePipe = o.BytePipe
	}

	op.ServerContext = o.ServerContext
	return op
}

func (o *RawGetFileDataAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_RawGetFileDataAsyncOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BytePipe = op.BytePipe

	o.ServerContext = op.ServerContext
}
func (o *RawGetFileDataAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RawGetFileDataAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawGetFileDataAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RawGetFileDataAsyncResponse structure represents the RawGetFileDataAsync operation response
type RawGetFileDataAsyncResponse struct {
	// bytePipe: The asynchronous RPC byte pipe that contains returned file data.
	BytePipe frs2.BytePipe `idl:"name:bytePipe" json:"byte_pipe"`
	// Return: The RawGetFileDataAsync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RawGetFileDataAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_RawGetFileDataAsyncOperation) *xxx_RawGetFileDataAsyncOperation {
	if op == nil {
		op = &xxx_RawGetFileDataAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.BytePipe = o.BytePipe
	op.Return = o.Return
	return op
}

func (o *RawGetFileDataAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_RawGetFileDataAsyncOperation) {
	if o == nil {
		return
	}
	o.BytePipe = op.BytePipe
	o.Return = op.Return
}
func (o *RawGetFileDataAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RawGetFileDataAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RawGetFileDataAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileDataAsyncOperation structure represents the RdcGetFileDataAsync operation
type xxx_GetFileDataAsyncOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	BytePipe      frs2.BytePipe  `idl:"name:bytePipe" json:"byte_pipe"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileDataAsyncOperation) OpNum() int { return 16 }

func (o *xxx_GetFileDataAsyncOperation) OpName() string {
	return "/FrsTransport/v1/RdcGetFileDataAsync"
}

func (o *xxx_GetFileDataAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetFileDataAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileDataAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// bytePipe {out} (1:{pointer=ref}*(1))(2:{alias=BYTE_PIPE}(pipe)(uint8))
	{
		// marshal pipe bytePipe
		if o.BytePipe != nil {
			for _chunk := range o.BytePipe.Pipe() {
				dimSize1 := uint64(len(_chunk))
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range _chunk {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_chunk[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_chunk); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
			}
		}
		if err := w.WriteSize(0); err != nil {
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

func (o *xxx_GetFileDataAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// bytePipe {out} (1:{pointer=ref}*(1))(2:{alias=BYTE_PIPE}(pipe)(uint8))
	{
		if o.BytePipe == nil {
			o.BytePipe = frs2.NewBytePipe()
		}
		for {
			var _chunk []byte
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _chunk", sizeInfo[0])
			}
			_chunk = make([]byte, sizeInfo[0])
			for i1 := range _chunk {
				i1 := i1
				if err := w.ReadData(&_chunk[i1]); err != nil {
					return err
				}
			}
			if len(_chunk) == 0 /* end */ {
				break
			}
			o.BytePipe.Append(_chunk)
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

// GetFileDataAsyncRequest structure represents the RdcGetFileDataAsync operation request
type GetFileDataAsyncRequest struct {
	// XXX: bytePipe is an implicit input depedency for output parameters
	BytePipe frs2.BytePipe `idl:"name:bytePipe" json:"byte_pipe"`

	// serverContext: The context handle that represents the requested file replication
	// operation. The client MUST specify a server context that was retrieved by a previously
	// successful call to the InitializeFileTransferAsync method in which the client set
	// the rdcDesired parameter to TRUE.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
}

func (o *GetFileDataAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileDataAsyncOperation) *xxx_GetFileDataAsyncOperation {
	if op == nil {
		op = &xxx_GetFileDataAsyncOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BytePipe == nil {
		op.BytePipe = o.BytePipe
	}

	op.ServerContext = o.ServerContext
	return op
}

func (o *GetFileDataAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileDataAsyncOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BytePipe = op.BytePipe

	o.ServerContext = op.ServerContext
}
func (o *GetFileDataAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileDataAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileDataAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileDataAsyncResponse structure represents the RdcGetFileDataAsync operation response
type GetFileDataAsyncResponse struct {
	// bytePipe: The asynchronous RPC byte pipe that contains returned file data.
	BytePipe frs2.BytePipe `idl:"name:bytePipe" json:"byte_pipe"`
	// Return: The RdcGetFileDataAsync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetFileDataAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileDataAsyncOperation) *xxx_GetFileDataAsyncOperation {
	if op == nil {
		op = &xxx_GetFileDataAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.BytePipe = o.BytePipe
	op.Return = o.Return
	return op
}

func (o *GetFileDataAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileDataAsyncOperation) {
	if o == nil {
		return
	}
	o.BytePipe = op.BytePipe
	o.Return = op.Return
}
func (o *GetFileDataAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileDataAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileDataAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FileDataTransferKeepAliveOperation structure represents the RdcFileDataTransferKeepAlive operation
type xxx_FileDataTransferKeepAliveOperation struct {
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_FileDataTransferKeepAliveOperation) OpNum() int { return 17 }

func (o *xxx_FileDataTransferKeepAliveOperation) OpName() string {
	return "/FrsTransport/v1/RdcFileDataTransferKeepAlive"
}

func (o *xxx_FileDataTransferKeepAliveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileDataTransferKeepAliveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext != nil {
			if err := o.ServerContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ServerContext{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FileDataTransferKeepAliveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// serverContext {in} (1:{context_handle, alias=PFRS_SERVER_CONTEXT, names=ndr_context_handle}(struct))
	{
		if o.ServerContext == nil {
			o.ServerContext = &ServerContext{}
		}
		if err := o.ServerContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileDataTransferKeepAliveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FileDataTransferKeepAliveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FileDataTransferKeepAliveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FileDataTransferKeepAliveRequest structure represents the RdcFileDataTransferKeepAlive operation request
type FileDataTransferKeepAliveRequest struct {
	// serverContext: The context handle that represents the requested file replication
	// operation that was retrieved by a previously successful call to the InitializeFileTransferAsync
	// method.
	ServerContext *ServerContext `idl:"name:serverContext" json:"server_context"`
}

func (o *FileDataTransferKeepAliveRequest) xxx_ToOp(ctx context.Context, op *xxx_FileDataTransferKeepAliveOperation) *xxx_FileDataTransferKeepAliveOperation {
	if op == nil {
		op = &xxx_FileDataTransferKeepAliveOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerContext = o.ServerContext
	return op
}

func (o *FileDataTransferKeepAliveRequest) xxx_FromOp(ctx context.Context, op *xxx_FileDataTransferKeepAliveOperation) {
	if o == nil {
		return
	}
	o.ServerContext = op.ServerContext
}
func (o *FileDataTransferKeepAliveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FileDataTransferKeepAliveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileDataTransferKeepAliveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FileDataTransferKeepAliveResponse structure represents the RdcFileDataTransferKeepAlive operation response
type FileDataTransferKeepAliveResponse struct {
	// Return: The RdcFileDataTransferKeepAlive return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FileDataTransferKeepAliveResponse) xxx_ToOp(ctx context.Context, op *xxx_FileDataTransferKeepAliveOperation) *xxx_FileDataTransferKeepAliveOperation {
	if op == nil {
		op = &xxx_FileDataTransferKeepAliveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FileDataTransferKeepAliveResponse) xxx_FromOp(ctx context.Context, op *xxx_FileDataTransferKeepAliveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FileDataTransferKeepAliveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FileDataTransferKeepAliveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FileDataTransferKeepAliveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
