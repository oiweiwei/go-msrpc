package ieventservice

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
)

var (
	// import guard
	GoPackage = "even6"
)

var (
	// Syntax UUID
	EventServiceSyntaxUUID = &uuid.UUID{TimeLow: 0xf6beaff7, TimeMid: 0x1e19, TimeHiAndVersion: 0x4fbb, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x8f, Node: [6]uint8{0xb8, 0x9e, 0x20, 0x18, 0x33, 0x7c}}
	// Syntax ID
	EventServiceSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: EventServiceSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// IEventService interface.
type EventServiceClient interface {

	// The EvtRpcRegisterRemoteSubscription (Opnum 0) method is used by a client to create
	// either a push or a pull subscription. In push subscriptions, the server calls the
	// client when new events are ready. In pull subscriptions, the client polls the server
	// for new events. Subscriptions can be to either a single channel and its associated
	// log, or to multiple channels and their logs.
	//
	// A client can use bookmarks to ensure a reliable subscription even if the client is
	// not continuously connected. A client can create a bookmark locally based on the contents
	// of an event that the client has processed. If the client disconnects and later reconnects,
	// it can use the bookmark to pick up where it left off. For information on bookmarks,
	// see section 2.2.14.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RegisterRemoteSubscription(context.Context, *RegisterRemoteSubscriptionRequest, ...dcerpc.CallOption) (*RegisterRemoteSubscriptionResponse, error)

	// The EvtRpcRemoteSubscriptionNextAsync (Opnum 1) method is used by a client to request
	// asynchronous delivery of events that are delivered to a subscription.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RemoteSubscriptionNextAsync(context.Context, *RemoteSubscriptionNextAsyncRequest, ...dcerpc.CallOption) (*RemoteSubscriptionNextAsyncResponse, error)

	// This EvtRpcRemoteSubscriptionNext (Opnum 2) method is a synchronous request for events
	// that have been delivered to a subscription. This method is only used for pull subscriptions
	// in which the client polls for events. The EvtRpcRemoteSubscriptionWaitAsync (section
	// 3.1.4.11) method can be used along with this method to minimize the frequency of
	// polling.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_TIMEOUT (0x000005b4) if fewer than numRequestedRecords records
	// are found within the time-out period. Otherwise, it MUST return a different implementation-specific
	// nonzero value as specified in [MS-ERREF].
	RemoteSubscriptionNext(context.Context, *RemoteSubscriptionNextRequest, ...dcerpc.CallOption) (*RemoteSubscriptionNextResponse, error)

	// Pull subscriptions are subscriptions in which the client requests records. The requests
	// can be done by using a polling mechanism. The EvtRpcRemoteSubscriptionWaitAsync (Opnum
	// 3) method can be used to enable the client to only poll when results are likely,
	// and is typically used in conjunction with the EvtRpcRemoteSubscriptionNext (Opnum
	// 2) (section 3.1.4.10) method, which is a blocking call; so this asynchronous method
	// is used to provide a way for the caller to not have to block or continuously poll
	// the server.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RemoteSubscriptionWaitAsync(context.Context, *RemoteSubscriptionWaitAsyncRequest, ...dcerpc.CallOption) (*RemoteSubscriptionWaitAsyncResponse, error)

	// The EvtRpcRegisterControllableOperation (Opnum 4) method obtains a CONTEXT_HANDLE_OPERATION_CONTROL
	// handle that can be used to cancel other operations.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RegisterControllableOperation(context.Context, *RegisterControllableOperationRequest, ...dcerpc.CallOption) (*RegisterControllableOperationResponse, error)

	// The EvtRpcRegisterLogQuery (Opnum 5) method is used to query one or more channels.
	// It can also be used to query a specific file. Actual retrieval of events is done
	// by subsequent calls to the EvtRpcQueryNext (section 3.1.4.13) method.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RegisterLogQuery(context.Context, *RegisterLogQueryRequest, ...dcerpc.CallOption) (*RegisterLogQueryResponse, error)

	// The EvtRpcClearLog (Opnum 6) method instructs the server to clear all the events
	// in a live channel, and optionally, to create a backup event log before the clear
	// takes place.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) on success; otherwise, it MUST
	// return an implementation-specific nonzero value as specified in [MS-ERREF].
	//
	// The server does not validate the control handle passed to EvtRpcClearLog and it SHOULD
	// assume that this parameter is always valid when the method is invoked.
	ClearLog(context.Context, *ClearLogRequest, ...dcerpc.CallOption) (*ClearLogResponse, error)

	// The EvtRpcExportLog (Opnum 7) method instructs the server to create a backup event
	// log at a specified file name.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	//
	// The server does not validate the control handle passed to EvtRpcExportLog, and it
	// SHOULD assume that this parameter is always valid when the method is invoked.
	ExportLog(context.Context, *ExportLogRequest, ...dcerpc.CallOption) (*ExportLogResponse, error)

	// The EvtRpcLocalizeExportLog (Opnum 8) method is used by a client to add localized
	// information to a previously created backup event log, because the backup event log
	// does not contain the localized strings for event descriptions. An example of how
	// this can be useful is if a backup event log needs to be copied to other computers
	// so that support personnel on those other computers can access it; if this method
	// has been called, such support personnel can access or view the localized backup event
	// log, which will then contain events with localized strings. Note that a discussion
	// of tools by which administrators or support personnel can work with localized backup
	// event log files in scenarios such as this is out of scope with respect to this protocol
	// specification.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an error value as specified in the processing rules in this section.<33>
	// Callers SHOULD treat all return values other than ERROR_SUCCESS equally and not alter
	// their behavior based on any specific error values.
	//
	// The server does not validate the control handle passed to EvtRpcLocalizeExportLog,
	// and it SHOULD assume that this parameter is always valid when the method is invoked.
	LocalizeExportLog(context.Context, *LocalizeExportLogRequest, ...dcerpc.CallOption) (*LocalizeExportLogResponse, error)

	// The EvtRpcMessageRender (Opnum 9) method is used by a client to get localized descriptive
	// strings for an event.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_INSUFFICIENT_BUFFER (0x0000007A) if maxSizeString is too
	// small to hold the result string. In that case, neededSizeString MUST be set to the
	// necessary size. Otherwise, the method MUST return a different implementation-specific
	// nonzero value as specified in [MS-ERREF].
	MessageRender(context.Context, *MessageRenderRequest, ...dcerpc.CallOption) (*MessageRenderResponse, error)

	// The EvtRpcMessageRenderDefault (Opnum 10) method is used by a client to get localized
	// strings for common values of opcodes, tasks, or keywords, as specified in section
	// 3.1.4.31.
	//
	// Return Values:  The method MUST return the following value on success.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// The method MUST return ERROR_INSUFFICIENT_BUFFER (0x0000007A) if maxSizeString is
	// too small to hold the result string. In that case, neededSizeString MUST be set to
	// the necessary size.
	//
	// Otherwise, the method MUST return a different implementation-specific nonzero value
	// as specified in [MS-ERREF].
	//
	// This method is the same as the EvtRpcMessageRender (section 3.1.4.31) method, except
	// that this method always uses the server's default strings (default strings come from
	// the server's default publisher, so a publisher handle is not required), whereas the
	// EvtRpcMessageRender (section 3.1.4.31) method uses only the default strings in the
	// case of level, task, opcode, and keyword values that fall in certain ranges. Therefore
	// it takes only 6 possible format flags. The server MUST fail the method with ERROR_INVALID_PARAMETER
	// (0x00000057) for any other flags than the 6 values given in the flags table.
	MessageRenderDefault(context.Context, *MessageRenderDefaultRequest, ...dcerpc.CallOption) (*MessageRenderDefaultResponse, error)

	// The EvtRpcQueryNext (Opnum 11) method is used by a client to get the next batch of
	// records from a query result set.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_TIMEOUT (0x000005bf) if no records are found within the
	// time-out period. The method MUST return ERROR_NO_MORE_ITEMS (0x00000103) once the
	// query has finished going through all the log(s); otherwise, it MUST return a different
	// implementation-specific nonzero value as specified in [MS-ERREF].
	QueryNext(context.Context, *QueryNextRequest, ...dcerpc.CallOption) (*QueryNextResponse, error)

	// The EvtRpcQuerySeek (Opnum 12) method is used by a client to move a query cursor
	// within a result set.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	QuerySeek(context.Context, *QuerySeekRequest, ...dcerpc.CallOption) (*QuerySeekResponse, error)

	// The EvtRpcClose (Opnum 13) method is used by a client to close context handles that
	// are opened by other methods in this protocol.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

	// The EvtRpcCancel (Opnum 14) method is used by a client to cancel another method.
	// This can be used to terminate long-running methods gracefully. Methods that can be
	// canceled include the subscription and query functions, and other functions that take
	// a CONTEXT_HANDLE_OPERATION_CONTROL argument.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	Cancel(context.Context, *CancelRequest, ...dcerpc.CallOption) (*CancelResponse, error)

	// The EvtRpcAssertConfig (Opnum 15) method indicates to the server that the publisher
	// or channel configuration has been updated.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	AssertConfig(context.Context, *AssertConfigRequest, ...dcerpc.CallOption) (*AssertConfigResponse, error)

	// The EvtRpcRetractConfig (Opnum 16) method indicates to the server that the publisher
	// or channel is to be removed.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RetractConfig(context.Context, *RetractConfigRequest, ...dcerpc.CallOption) (*RetractConfigResponse, error)

	// The EvtRpcOpenLogHandle (Opnum 17) method is used by a client to get information
	// about a channel or a backup event log.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	OpenLog(context.Context, *OpenLogRequest, ...dcerpc.CallOption) (*OpenLogResponse, error)

	// The EvtRpcGetLogFileInfo (Opnum 18) method is used by a client to get information
	// about a live channel or a backup event log.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_INSUFFICIENT_BUFFER (0x0000007A) if the buffer is too small;
	// otherwise, it MUST return a different implementation-specific nonzero value as specified
	// in [MS-ERREF].
	GetLogFileInfo(context.Context, *GetLogFileInfoRequest, ...dcerpc.CallOption) (*GetLogFileInfoResponse, error)

	// The EvtRpcGetChannelList (Opnum 19) method is used to enumerate the set of available
	// channels.
	//
	// Return Values:  The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetChannelList(context.Context, *GetChannelListRequest, ...dcerpc.CallOption) (*GetChannelListResponse, error)

	// The EvtRpcGetChannelConfig (opnum 20) method is used by a client to get the configuration
	// for a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetChannelConfig(context.Context, *GetChannelConfigRequest, ...dcerpc.CallOption) (*GetChannelConfigResponse, error)

	// The EvtRpcPutChannelConfig (Opnum 21) method is used by a client to update the configuration
	// for a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].<42>
	PutChannelConfig(context.Context, *PutChannelConfigRequest, ...dcerpc.CallOption) (*PutChannelConfigResponse, error)

	// The EvtRpcGetPublisherList (Opnum 22) method is used by a client to get the list
	// of publishers.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherList(context.Context, *GetPublisherListRequest, ...dcerpc.CallOption) (*GetPublisherListResponse, error)

	// The EvtRpcGetPublisherListForChannel (Opnum 23) method is used by a client to get
	// the list of publishers that write events to a particular channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherListForChannel(context.Context, *GetPublisherListForChannelRequest, ...dcerpc.CallOption) (*GetPublisherListForChannelResponse, error)

	// The EvtRpcGetPublisherMetadata (Opnum 24) method is used by a client to open a handle
	// to publisher metadata. It also gets some initial information from the metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherMetadata(context.Context, *GetPublisherMetadataRequest, ...dcerpc.CallOption) (*GetPublisherMetadataResponse, error)

	// The EvtRpcGetPublisherResourceMetadata (Opnum 25) method obtains information from
	// the publisher metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherResourceMetadata(context.Context, *GetPublisherResourceMetadataRequest, ...dcerpc.CallOption) (*GetPublisherResourceMetadataResponse, error)

	// The EvtRpcGetEventMetadataEnum (Opnum 26) method obtains a handle for enumerating
	// a publisher's event metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetEventMetadataEnum(context.Context, *GetEventMetadataEnumRequest, ...dcerpc.CallOption) (*GetEventMetadataEnumResponse, error)

	// The EvtRpcGetNextEventMetadata (Opnum 27) method gets details about a possible event
	// and also returns the next event metadata in the enumeration. It is used to enumerate
	// through the event definitions for the publisher associated with the handle. The enumeration
	// is in the forward direction only, and there is no reset functionality.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetNextEventMetadata(context.Context, *GetNextEventMetadataRequest, ...dcerpc.CallOption) (*GetNextEventMetadataResponse, error)

	// The EvtRpcGetClassicLogDisplayName (Opnum 28) method obtains a descriptive name for
	// a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetClassicLogDisplayName(context.Context, *GetClassicLogDisplayNameRequest, ...dcerpc.CallOption) (*GetClassicLogDisplayNameResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// SubscribePull represents the EvtRpcSubscribePull RPC constant
var SubscribePull = 268435456

// VarFlagsModified represents the EvtRpcVarFlagsModified RPC constant
var VarFlagsModified = 1

// RemoteSubscription structure represents PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION RPC structure.
type RemoteSubscription dcetypes.ContextHandle

func (o *RemoteSubscription) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *RemoteSubscription) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteSubscription) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteSubscription) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// LogQuery structure represents PCONTEXT_HANDLE_LOG_QUERY RPC structure.
type LogQuery dcetypes.ContextHandle

func (o *LogQuery) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *LogQuery) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *LogQuery) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *LogQuery) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Log structure represents PCONTEXT_HANDLE_LOG_HANDLE RPC structure.
type Log dcetypes.ContextHandle

func (o *Log) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Log) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Log) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Log) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// OperationControl structure represents PCONTEXT_HANDLE_OPERATION_CONTROL RPC structure.
type OperationControl dcetypes.ContextHandle

func (o *OperationControl) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *OperationControl) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *OperationControl) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *OperationControl) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PublisherMetadata structure represents PCONTEXT_HANDLE_PUBLISHER_METADATA RPC structure.
type PublisherMetadata dcetypes.ContextHandle

func (o *PublisherMetadata) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *PublisherMetadata) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PublisherMetadata) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PublisherMetadata) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventMetadataEnum structure represents PCONTEXT_HANDLE_EVENT_METADATA_ENUM RPC structure.
type EventMetadataEnum dcetypes.ContextHandle

func (o *EventMetadataEnum) ContextHandle() *dcetypes.ContextHandle {
	return (*dcetypes.ContextHandle)(o)
}

func (o *EventMetadataEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventMetadataEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventMetadataEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Info structure represents RpcInfo RPC structure.
//
// The RpcInfo structure is used for certain methods that return additional information
// about errors.
type Info struct {
	// m_error:  A Win32 error code that contains a general operation success or failure
	// status. A value of 0x00000000 indicates success; any other value indicates failure.
	// Unless noted otherwise, all failure values MUST be treated equally.
	Error uint32 `idl:"name:m_error" json:"error"`
	// m_subErr:  MUST be zero unless specified otherwise in the method using this structure.
	// Unless noted otherwise, all nonzero values MUST be treated equally.
	SubError uint32 `idl:"name:m_subErr" json:"sub_error"`
	// m_subErrParam:   MUST be zero unless specified otherwise in the method using this
	// structure. Unless noted otherwise, all nonzero values MUST be treated equally.
	SubErrorParam uint32 `idl:"name:m_subErrParam" json:"sub_error_param"`
}

func (o *Info) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	if err := w.WriteData(o.SubError); err != nil {
		return err
	}
	if err := w.WriteData(o.SubErrorParam); err != nil {
		return err
	}
	return nil
}
func (o *Info) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubError); err != nil {
		return err
	}
	if err := w.ReadData(&o.SubErrorParam); err != nil {
		return err
	}
	return nil
}

// BooleanArray structure represents BooleanArray RPC structure.
//
// The BooleanArray structure is defined as follows.
type BooleanArray struct {
	// count:  A 32-bit unsigned integer that contains the number of BOOLEAN values pointed
	// to by ptr.
	Count uint32 `idl:"name:count" json:"count"`
	// ptr:  A pointer to an array of BOOLEAN values.
	Array []bool `idl:"name:ptr;size_is:(count)" json:"array"`
}

func (o *BooleanArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Array != nil && o.Count == 0 {
		o.Count = uint32(len(o.Array))
	}
	if o.Count > uint32(524288) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BooleanArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Array != nil || o.Count > 0 {
		_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Array {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Array[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Array); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(false); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Array, _ptr_ptr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *BooleanArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Array", sizeInfo[0])
		}
		o.Array = make([]bool, sizeInfo[0])
		for i1 := range o.Array {
			i1 := i1
			if err := w.ReadData(&o.Array[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ptr := func(ptr interface{}) { o.Array = *ptr.(*[]bool) }
	if err := w.ReadPointer(&o.Array, _s_ptr, _ptr_ptr); err != nil {
		return err
	}
	return nil
}

// Uint32Array structure represents UInt32Array RPC structure.
//
// The UInt32Array structure is defined as follows.
type Uint32Array struct {
	// count:  An unsigned 32-bit integer that contains the number of unsigned 32-bit integers
	// pointed to by ptr.
	Count uint32 `idl:"name:count" json:"count"`
	// ptr:  A pointer to an array of unsigned 32-bit integers.
	Array []uint32 `idl:"name:ptr;size_is:(count)" json:"array"`
}

func (o *Uint32Array) xxx_PreparePayload(ctx context.Context) error {
	if o.Array != nil && o.Count == 0 {
		o.Count = uint32(len(o.Array))
	}
	if o.Count > uint32(524288) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Array != nil || o.Count > 0 {
		_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Array {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Array[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Array); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Array, _ptr_ptr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Array", sizeInfo[0])
		}
		o.Array = make([]uint32, sizeInfo[0])
		for i1 := range o.Array {
			i1 := i1
			if err := w.ReadData(&o.Array[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ptr := func(ptr interface{}) { o.Array = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Array, _s_ptr, _ptr_ptr); err != nil {
		return err
	}
	return nil
}

// Uint64Array structure represents UInt64Array RPC structure.
//
// The UInt64Array structure is defined as follows.
type Uint64Array struct {
	// count:  A 32-bit unsigned integer that contains the number of 64-bit integers pointed
	// to by ptr.
	Count uint32 `idl:"name:count" json:"count"`
	// ptr:  A pointer to an array of unsigned 64-bit integers.
	Array []uint64 `idl:"name:ptr;size_is:(count)" json:"array"`
}

func (o *Uint64Array) xxx_PreparePayload(ctx context.Context) error {
	if o.Array != nil && o.Count == 0 {
		o.Count = uint32(len(o.Array))
	}
	if o.Count > uint32(262144) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint64Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Array != nil || o.Count > 0 {
		_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Array {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Array[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Array); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint64(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Array, _ptr_ptr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Uint64Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Array", sizeInfo[0])
		}
		o.Array = make([]uint64, sizeInfo[0])
		for i1 := range o.Array {
			i1 := i1
			if err := w.ReadData(&o.Array[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ptr := func(ptr interface{}) { o.Array = *ptr.(*[]uint64) }
	if err := w.ReadPointer(&o.Array, _s_ptr, _ptr_ptr); err != nil {
		return err
	}
	return nil
}

// StringArray structure represents StringArray RPC structure.
//
// The StringArray structure is defined as follows.
type StringArray struct {
	// count:  A 32-bit unsigned integer that contains the number of strings pointed to
	// by ptr.
	Count uint32 `idl:"name:count" json:"count"`
	// ptr:  A pointer to an array of null-terminated Unicode (as specified in [UNICODE])
	// strings.
	Array []string `idl:"name:ptr;size_is:(count);string" json:"array"`
}

func (o *StringArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Array != nil && o.Count == 0 {
		o.Count = uint32(len(o.Array))
	}
	if o.Count > uint32(4096) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Array != nil || o.Count > 0 {
		_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Array {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Array[i1] != "" {
					_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteUTF16NString(ctx, w, o.Array[i1]); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Array[i1], _ptr_ptr); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Array); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Array, _ptr_ptr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Array", sizeInfo[0])
		}
		o.Array = make([]string, sizeInfo[0])
		for i1 := range o.Array {
			i1 := i1
			_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.Array[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ptr := func(ptr interface{}) { o.Array[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.Array[i1], _s_ptr, _ptr_ptr); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ptr := func(ptr interface{}) { o.Array = *ptr.(*[]string) }
	if err := w.ReadPointer(&o.Array, _s_ptr, _ptr_ptr); err != nil {
		return err
	}
	return nil
}

// GUIDArray structure represents GuidArray RPC structure.
//
// The GuidArray structure is defined as follows.
type GUIDArray struct {
	// count:  A 32-bit unsigned integer that contains the number of GUIDs pointed to by
	// ptr.
	Count uint32 `idl:"name:count" json:"count"`
	// ptr:   A pointer to an array of GUIDs.
	Array []*dtyp.GUID `idl:"name:ptr;size_is:(count)" json:"array"`
}

func (o *GUIDArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Array != nil && o.Count == 0 {
		o.Count = uint32(len(o.Array))
	}
	if o.Count > uint32(131072) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GUIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Array != nil || o.Count > 0 {
		_ptr_ptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Array {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Array[i1] != nil {
					if err := o.Array[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Array); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Array, _ptr_ptr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *GUIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_ptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Array", sizeInfo[0])
		}
		o.Array = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.Array {
			i1 := i1
			if o.Array[i1] == nil {
				o.Array[i1] = &dtyp.GUID{}
			}
			if err := o.Array[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ptr := func(ptr interface{}) { o.Array = *ptr.(*[]*dtyp.GUID) }
	if err := w.ReadPointer(&o.Array, _s_ptr, _ptr_ptr); err != nil {
		return err
	}
	return nil
}

// VariantType type represents EvtRpcVariantType RPC enumeration.
type VariantType uint32

var (
	VariantTypeNull         VariantType = 0
	VariantTypeBoolean      VariantType = 1
	VariantTypeUint32       VariantType = 2
	VariantTypeUint64       VariantType = 3
	VariantTypeString       VariantType = 4
	VariantTypeGUID         VariantType = 5
	VariantTypeBooleanArray VariantType = 6
	VariantTypeUint32Array  VariantType = 7
	VariantTypeUint64Array  VariantType = 8
	VariantTypeStringArray  VariantType = 9
	VariantTypeGUIDArray    VariantType = 10
)

func (o VariantType) String() string {
	switch o {
	case VariantTypeNull:
		return "VariantTypeNull"
	case VariantTypeBoolean:
		return "VariantTypeBoolean"
	case VariantTypeUint32:
		return "VariantTypeUint32"
	case VariantTypeUint64:
		return "VariantTypeUint64"
	case VariantTypeString:
		return "VariantTypeString"
	case VariantTypeGUID:
		return "VariantTypeGUID"
	case VariantTypeBooleanArray:
		return "VariantTypeBooleanArray"
	case VariantTypeUint32Array:
		return "VariantTypeUint32Array"
	case VariantTypeUint64Array:
		return "VariantTypeUint64Array"
	case VariantTypeStringArray:
		return "VariantTypeStringArray"
	case VariantTypeGUIDArray:
		return "VariantTypeGUIDArray"
	}
	return "Invalid"
}

// AssertConfigFlags type represents EvtRpcAssertConfigFlags RPC enumeration.
//
// The EvtRpcAssertConfigFlags Enumeration members specify how the path and channelPath
// parameters (used by a number of the methods in 3.1.4) are to be interpreted.
type AssertConfigFlags uint32

var (
	// EvtRpcChannelPath:  The associated parameter string contains a path to a channel.
	AssertConfigFlagsChannelPath AssertConfigFlags = 0
	// EvtRpcPublisherName:  The associated parameter string contains a publisher name.
	AssertConfigFlagsPublisherName AssertConfigFlags = 1
)

func (o AssertConfigFlags) String() string {
	switch o {
	case AssertConfigFlagsChannelPath:
		return "AssertConfigFlagsChannelPath"
	case AssertConfigFlagsPublisherName:
		return "AssertConfigFlagsPublisherName"
	}
	return "Invalid"
}

// Variant structure represents EvtRpcVariant RPC structure.
//
// The EvtRpcVariant structure is defined as follows.
type Variant struct {
	// type:  Indicates the actual type of the union.
	Type VariantType `idl:"name:type" json:"type"`
	// flags:   This flag MUST be set to either 0x0000 or 0x0001. If this flag is set to
	// 0x0001, it indicates that an EvtRpcVariant structure has been changed by the client.
	// For an example of how this flag might be set, suppose the client application retrieved
	// an EvtRpcVariantList structure by calling EvtRpcGetChannelConfig, changed one or
	// more EvtRpcVariant structures in the list, and then sent the list back to the server
	// via EvtRpcPutChannelConfig. In this example, the server updates the values corresponding
	// to the EvtRpcVariant structures with this flag set.
	//
	//	+--------+----------------------------------------------------------------------------------+
	//	|        |                                                                                  |
	//	| VALUE  |                                     MEANING                                      |
	//	|        |                                                                                  |
	//	+--------+----------------------------------------------------------------------------------+
	//	+--------+----------------------------------------------------------------------------------+
	//	| 0x0000 | A flag indicating that no instance of an EvtRpcVariant structure was changed by  |
	//	|        | the client.                                                                      |
	//	+--------+----------------------------------------------------------------------------------+
	//	| 0x0001 | A flag indicating that an EvtRpcVariant structure was changed by the client.     |
	//	+--------+----------------------------------------------------------------------------------+
	Flags   uint32           `idl:"name:flags" json:"flags"`
	Variant *Variant_Variant `idl:"name:Variant;switch_is:type" json:"variant"`
}

func (o *Variant) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.Type)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	_swVariant := uint32(o.Type)
	if o.Variant != nil {
		if err := o.Variant.MarshalUnionNDR(ctx, w, _swVariant); err != nil {
			return err
		}
	} else {
		if err := (&Variant_Variant{}).MarshalUnionNDR(ctx, w, _swVariant); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.Type)); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if o.Variant == nil {
		o.Variant = &Variant_Variant{}
	}
	_swVariant := uint32(o.Type)
	if err := o.Variant.UnmarshalUnionNDR(ctx, w, _swVariant); err != nil {
		return err
	}
	return nil
}

// Variant_Variant structure represents EvtRpcVariant union anonymous member.
//
// The EvtRpcVariant structure is defined as follows.
type Variant_Variant struct {
	// Types that are assignable to Value
	//
	// *Variant_NullValue
	// *Variant_BooleanValue
	// *Variant_Uint32Value
	// *Variant_Uint64Value
	// *Variant_StringValue
	// *Variant_GUIDValue
	// *Variant_BooleanArray
	// *Variant_Uint32Array
	// *Variant_Uint64Array
	// *Variant_StringArray
	// *Variant_GUIDArray
	Value is_Variant_Variant `json:"value"`
}

func (o *Variant_Variant) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Variant_NullValue:
		if value != nil {
			return value.NullValue
		}
	case *Variant_BooleanValue:
		if value != nil {
			return value.BooleanValue
		}
	case *Variant_Uint32Value:
		if value != nil {
			return value.Uint32Value
		}
	case *Variant_Uint64Value:
		if value != nil {
			return value.Uint64Value
		}
	case *Variant_StringValue:
		if value != nil {
			return value.StringValue
		}
	case *Variant_GUIDValue:
		if value != nil {
			return value.GUIDValue
		}
	case *Variant_BooleanArray:
		if value != nil {
			return value.BooleanArray
		}
	case *Variant_Uint32Array:
		if value != nil {
			return value.Uint32Array
		}
	case *Variant_Uint64Array:
		if value != nil {
			return value.Uint64Array
		}
	case *Variant_StringArray:
		if value != nil {
			return value.StringArray
		}
	case *Variant_GUIDArray:
		if value != nil {
			return value.GUIDArray
		}
	}
	return nil
}

type is_Variant_Variant interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Variant_Variant()
}

func (o *Variant_Variant) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Variant_NullValue:
		return uint32(0)
	case *Variant_BooleanValue:
		return uint32(1)
	case *Variant_Uint32Value:
		return uint32(2)
	case *Variant_Uint64Value:
		return uint32(3)
	case *Variant_StringValue:
		return uint32(4)
	case *Variant_GUIDValue:
		return uint32(5)
	case *Variant_BooleanArray:
		return uint32(6)
	case *Variant_Uint32Array:
		return uint32(7)
	case *Variant_Uint64Array:
		return uint32(8)
	case *Variant_StringArray:
		return uint32(9)
	case *Variant_GUIDArray:
		return uint32(10)
	}
	return uint32(0)
}

func (o *Variant_Variant) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint32(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*Variant_NullValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_NullValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*Variant_BooleanValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_BooleanValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*Variant_Uint32Value)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_Uint32Value{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*Variant_Uint64Value)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_Uint64Value{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*Variant_StringValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_StringValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(5):
		_o, _ := o.Value.(*Variant_GUIDValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_GUIDValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*Variant_BooleanArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_BooleanArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*Variant_Uint32Array)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_Uint32Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*Variant_Uint64Array)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_Uint64Array{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*Variant_StringArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_StringArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*Variant_GUIDArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_GUIDArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Variant_Variant) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint32)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &Variant_NullValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &Variant_BooleanValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &Variant_Uint32Value{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &Variant_Uint64Value{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &Variant_StringValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(5):
		o.Value = &Variant_GUIDValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &Variant_BooleanArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &Variant_Uint32Array{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &Variant_Uint64Array{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &Variant_StringArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &Variant_GUIDArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Variant_NullValue structure represents Variant_Variant RPC union arm.
//
// It has following labels: 0
type Variant_NullValue struct {
	// nullVal:  MUST be set to 0x00000000.
	NullValue int32 `idl:"name:nullVal" json:"null_value"`
}

func (*Variant_NullValue) is_Variant_Variant() {}

func (o *Variant_NullValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.NullValue); err != nil {
		return err
	}
	return nil
}
func (o *Variant_NullValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.NullValue); err != nil {
		return err
	}
	return nil
}

// Variant_BooleanValue structure represents Variant_Variant RPC union arm.
//
// It has following labels: 1
type Variant_BooleanValue struct {
	// booleanVal:  A BOOLEAN value.
	BooleanValue bool `idl:"name:booleanVal" json:"boolean_value"`
}

func (*Variant_BooleanValue) is_Variant_Variant() {}

func (o *Variant_BooleanValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.BooleanValue); err != nil {
		return err
	}
	return nil
}
func (o *Variant_BooleanValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.BooleanValue); err != nil {
		return err
	}
	return nil
}

// Variant_Uint32Value structure represents Variant_Variant RPC union arm.
//
// It has following labels: 2
type Variant_Uint32Value struct {
	// uint32Val:  A 32-bit unsigned integer.
	Uint32Value uint32 `idl:"name:uint32Val" json:"uint32_value"`
}

func (*Variant_Uint32Value) is_Variant_Variant() {}

func (o *Variant_Uint32Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint32Value); err != nil {
		return err
	}
	return nil
}
func (o *Variant_Uint32Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint32Value); err != nil {
		return err
	}
	return nil
}

// Variant_Uint64Value structure represents Variant_Variant RPC union arm.
//
// It has following labels: 3
type Variant_Uint64Value struct {
	// uint64Val:  A 64-bit unsigned integer.
	Uint64Value uint64 `idl:"name:uint64Val" json:"uint64_value"`
}

func (*Variant_Uint64Value) is_Variant_Variant() {}

func (o *Variant_Uint64Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint64Value); err != nil {
		return err
	}
	return nil
}
func (o *Variant_Uint64Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint64Value); err != nil {
		return err
	}
	return nil
}

// Variant_StringValue structure represents Variant_Variant RPC union arm.
//
// It has following labels: 4
type Variant_StringValue struct {
	// stringVal:  A null-terminated UNICODE string.
	StringValue string `idl:"name:stringVal;string" json:"string_value"`
}

func (*Variant_StringValue) is_Variant_Variant() {}

func (o *Variant_StringValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StringValue != "" {
		_ptr_stringVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.StringValue); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.StringValue, _ptr_stringVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_StringValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_stringVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.StringValue); err != nil {
			return err
		}
		return nil
	})
	_s_stringVal := func(ptr interface{}) { o.StringValue = *ptr.(*string) }
	if err := w.ReadPointer(&o.StringValue, _s_stringVal, _ptr_stringVal); err != nil {
		return err
	}
	return nil
}

// Variant_GUIDValue structure represents Variant_Variant RPC union arm.
//
// It has following labels: 5
type Variant_GUIDValue struct {
	// guidVal:  A GUID.
	GUIDValue *dtyp.GUID `idl:"name:guidVal" json:"guid_value"`
}

func (*Variant_GUIDValue) is_Variant_Variant() {}

func (o *Variant_GUIDValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GUIDValue != nil {
		_ptr_guidVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.GUIDValue != nil {
				if err := o.GUIDValue.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GUIDValue, _ptr_guidVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_GUIDValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_guidVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.GUIDValue == nil {
			o.GUIDValue = &dtyp.GUID{}
		}
		if err := o.GUIDValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_guidVal := func(ptr interface{}) { o.GUIDValue = *ptr.(**dtyp.GUID) }
	if err := w.ReadPointer(&o.GUIDValue, _s_guidVal, _ptr_guidVal); err != nil {
		return err
	}
	return nil
}

// Variant_BooleanArray structure represents Variant_Variant RPC union arm.
//
// It has following labels: 6
type Variant_BooleanArray struct {
	// booleanArray:  An array of BOOLEAN values that are stored as a BooleanArray.
	BooleanArray *BooleanArray `idl:"name:booleanArray" json:"boolean_array"`
}

func (*Variant_BooleanArray) is_Variant_Variant() {}

func (o *Variant_BooleanArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BooleanArray != nil {
		if err := o.BooleanArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&BooleanArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_BooleanArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.BooleanArray == nil {
		o.BooleanArray = &BooleanArray{}
	}
	if err := o.BooleanArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_Uint32Array structure represents Variant_Variant RPC union arm.
//
// It has following labels: 7
type Variant_Uint32Array struct {
	// uint32Array:  An array of 32-bit unsigned integers that are stored as a UInt32Array.
	Uint32Array *Uint32Array `idl:"name:uint32Array" json:"uint32_array"`
}

func (*Variant_Uint32Array) is_Variant_Variant() {}

func (o *Variant_Uint32Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Uint32Array != nil {
		if err := o.Uint32Array.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Uint32Array{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_Uint32Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Uint32Array == nil {
		o.Uint32Array = &Uint32Array{}
	}
	if err := o.Uint32Array.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_Uint64Array structure represents Variant_Variant RPC union arm.
//
// It has following labels: 8
type Variant_Uint64Array struct {
	// uint64Array:  An array of 64-bit unsigned integers that are stored as a UInt64Array.
	Uint64Array *Uint64Array `idl:"name:uint64Array" json:"uint64_array"`
}

func (*Variant_Uint64Array) is_Variant_Variant() {}

func (o *Variant_Uint64Array) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Uint64Array != nil {
		if err := o.Uint64Array.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Uint64Array{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_Uint64Array) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Uint64Array == nil {
		o.Uint64Array = &Uint64Array{}
	}
	if err := o.Uint64Array.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_StringArray structure represents Variant_Variant RPC union arm.
//
// It has following labels: 9
type Variant_StringArray struct {
	// stringArray:  An array of strings that are stored as a StringArray.
	StringArray *StringArray `idl:"name:stringArray" json:"string_array"`
}

func (*Variant_StringArray) is_Variant_Variant() {}

func (o *Variant_StringArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StringArray != nil {
		if err := o.StringArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&StringArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_StringArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.StringArray == nil {
		o.StringArray = &StringArray{}
	}
	if err := o.StringArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_GUIDArray structure represents Variant_Variant RPC union arm.
//
// It has following labels: 10
type Variant_GUIDArray struct {
	// guidArray:  An array of GUIDs that are stored as a GuidArray.
	GUIDArray *GUIDArray `idl:"name:guidArray" json:"guid_array"`
}

func (*Variant_GUIDArray) is_Variant_Variant() {}

func (o *Variant_GUIDArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.GUIDArray != nil {
		if err := o.GUIDArray.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&GUIDArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_GUIDArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.GUIDArray == nil {
		o.GUIDArray = &GUIDArray{}
	}
	if err := o.GUIDArray.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// VariantList structure represents EvtRpcVariantList RPC structure.
//
// The EvtRpcVariantList data type is a wrapper for multiple EvtRpcVariant (section
// 2.2.7) data types.
type VariantList struct {
	// count:  Number of EvtRpcVariant values pointed to by the props field.
	Count uint32 `idl:"name:count" json:"count"`
	// props:  Pointer to an array of EvtRpcVariant values.
	Properties []*Variant `idl:"name:props;size_is:(count)" json:"properties"`
}

func (o *VariantList) xxx_PreparePayload(ctx context.Context) error {
	if o.Properties != nil && o.Count == 0 {
		o.Count = uint32(len(o.Properties))
	}
	if o.Count > uint32(256) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VariantList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	if o.Properties != nil || o.Count > 0 {
		_ptr_props := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Count)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Properties {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Properties[i1] != nil {
					if err := o.Properties[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Properties); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Properties, _ptr_props); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VariantList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	_ptr_props := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Count > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Count)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Properties", sizeInfo[0])
		}
		o.Properties = make([]*Variant, sizeInfo[0])
		for i1 := range o.Properties {
			i1 := i1
			if o.Properties[i1] == nil {
				o.Properties[i1] = &Variant{}
			}
			if err := o.Properties[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_props := func(ptr interface{}) { o.Properties = *ptr.(*[]*Variant) }
	if err := w.ReadPointer(&o.Properties, _s_props, _ptr_props); err != nil {
		return err
	}
	return nil
}

// QueryChannelInfo structure represents EvtRpcQueryChannelInfo RPC structure.
//
// The format of the EvtRpcQueryChannelInfo data type is as follows.
type QueryChannelInfo struct {
	// name:  Name of the channel to which the status applies.
	Name string `idl:"name:name" json:"name"`
	// status:   A Win32 error code that indicates the channel status. A value of 0x00000000
	// indicates success; any other value indicates failure. Unless otherwise noted, all
	// failure values MUST be treated equally.
	Status uint32 `idl:"name:status" json:"status"`
}

func (o *QueryChannelInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *QueryChannelInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16String(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	return nil
}
func (o *QueryChannelInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16String(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultEventServiceClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultEventServiceClient) RegisterRemoteSubscription(ctx context.Context, in *RegisterRemoteSubscriptionRequest, opts ...dcerpc.CallOption) (*RegisterRemoteSubscriptionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterRemoteSubscriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RemoteSubscriptionNextAsync(ctx context.Context, in *RemoteSubscriptionNextAsyncRequest, opts ...dcerpc.CallOption) (*RemoteSubscriptionNextAsyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteSubscriptionNextAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RemoteSubscriptionNext(ctx context.Context, in *RemoteSubscriptionNextRequest, opts ...dcerpc.CallOption) (*RemoteSubscriptionNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteSubscriptionNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RemoteSubscriptionWaitAsync(ctx context.Context, in *RemoteSubscriptionWaitAsyncRequest, opts ...dcerpc.CallOption) (*RemoteSubscriptionWaitAsyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteSubscriptionWaitAsyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RegisterControllableOperation(ctx context.Context, in *RegisterControllableOperationRequest, opts ...dcerpc.CallOption) (*RegisterControllableOperationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterControllableOperationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RegisterLogQuery(ctx context.Context, in *RegisterLogQueryRequest, opts ...dcerpc.CallOption) (*RegisterLogQueryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterLogQueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) ClearLog(ctx context.Context, in *ClearLogRequest, opts ...dcerpc.CallOption) (*ClearLogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) ExportLog(ctx context.Context, in *ExportLogRequest, opts ...dcerpc.CallOption) (*ExportLogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ExportLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) LocalizeExportLog(ctx context.Context, in *LocalizeExportLogRequest, opts ...dcerpc.CallOption) (*LocalizeExportLogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LocalizeExportLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) MessageRender(ctx context.Context, in *MessageRenderRequest, opts ...dcerpc.CallOption) (*MessageRenderResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageRenderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) MessageRenderDefault(ctx context.Context, in *MessageRenderDefaultRequest, opts ...dcerpc.CallOption) (*MessageRenderDefaultResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MessageRenderDefaultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) QueryNext(ctx context.Context, in *QueryNextRequest, opts ...dcerpc.CallOption) (*QueryNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) QuerySeek(ctx context.Context, in *QuerySeekRequest, opts ...dcerpc.CallOption) (*QuerySeekResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySeekResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
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

func (o *xxx_DefaultEventServiceClient) Cancel(ctx context.Context, in *CancelRequest, opts ...dcerpc.CallOption) (*CancelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CancelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) AssertConfig(ctx context.Context, in *AssertConfigRequest, opts ...dcerpc.CallOption) (*AssertConfigResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AssertConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) RetractConfig(ctx context.Context, in *RetractConfigRequest, opts ...dcerpc.CallOption) (*RetractConfigResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetractConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) OpenLog(ctx context.Context, in *OpenLogRequest, opts ...dcerpc.CallOption) (*OpenLogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetLogFileInfo(ctx context.Context, in *GetLogFileInfoRequest, opts ...dcerpc.CallOption) (*GetLogFileInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLogFileInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetChannelList(ctx context.Context, in *GetChannelListRequest, opts ...dcerpc.CallOption) (*GetChannelListResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetChannelListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetChannelConfig(ctx context.Context, in *GetChannelConfigRequest, opts ...dcerpc.CallOption) (*GetChannelConfigResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetChannelConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) PutChannelConfig(ctx context.Context, in *PutChannelConfigRequest, opts ...dcerpc.CallOption) (*PutChannelConfigResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PutChannelConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetPublisherList(ctx context.Context, in *GetPublisherListRequest, opts ...dcerpc.CallOption) (*GetPublisherListResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetPublisherListForChannel(ctx context.Context, in *GetPublisherListForChannelRequest, opts ...dcerpc.CallOption) (*GetPublisherListForChannelResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherListForChannelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetPublisherMetadata(ctx context.Context, in *GetPublisherMetadataRequest, opts ...dcerpc.CallOption) (*GetPublisherMetadataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetPublisherResourceMetadata(ctx context.Context, in *GetPublisherResourceMetadataRequest, opts ...dcerpc.CallOption) (*GetPublisherResourceMetadataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPublisherResourceMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetEventMetadataEnum(ctx context.Context, in *GetEventMetadataEnumRequest, opts ...dcerpc.CallOption) (*GetEventMetadataEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetEventMetadataEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetNextEventMetadata(ctx context.Context, in *GetNextEventMetadataRequest, opts ...dcerpc.CallOption) (*GetNextEventMetadataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNextEventMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) GetClassicLogDisplayName(ctx context.Context, in *GetClassicLogDisplayNameRequest, opts ...dcerpc.CallOption) (*GetClassicLogDisplayNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassicLogDisplayNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventServiceClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventServiceClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewEventServiceClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventServiceClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventServiceSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultEventServiceClient{cc: cc}, nil
}

// xxx_RegisterRemoteSubscriptionOperation structure represents the EvtRpcRegisterRemoteSubscription operation
type xxx_RegisterRemoteSubscriptionOperation struct {
	ChannelPath          string              `idl:"name:channelPath;string;pointer:unique" json:"channel_path"`
	Query                string              `idl:"name:query;string" json:"query"`
	BookmarkXML          string              `idl:"name:bookmarkXml;string;pointer:unique" json:"bookmark_xml"`
	Flags                uint32              `idl:"name:flags" json:"flags"`
	Handle               *RemoteSubscription `idl:"name:handle" json:"handle"`
	Control              *OperationControl   `idl:"name:control" json:"control"`
	QueryChannelInfoSize uint32              `idl:"name:queryChannelInfoSize" json:"query_channel_info_size"`
	QueryChannelInfo     []*QueryChannelInfo `idl:"name:queryChannelInfo;size_is:(, queryChannelInfoSize)" json:"query_channel_info"`
	Error                *Info               `idl:"name:error" json:"error"`
	Return               uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterRemoteSubscriptionOperation) OpNum() int { return 0 }

func (o *xxx_RegisterRemoteSubscriptionOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRegisterRemoteSubscription"
}

func (o *xxx_RegisterRemoteSubscriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ChannelPath) > int(512) {
		return fmt.Errorf("ChannelPath is out of range")
	}
	if len(o.Query) > int(1048576) {
		return fmt.Errorf("Query is out of range")
	}
	if len(o.BookmarkXML) > int(1048576) {
		return fmt.Errorf("BookmarkXML is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterRemoteSubscriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// channelPath {in} (1:{string, pointer=unique, range=(0,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ChannelPath != "" {
			_ptr_channelPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ChannelPath, _ptr_channelPath); err != nil {
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
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Query); err != nil {
			return err
		}
	}
	// bookmarkXml {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BookmarkXML != "" {
			_ptr_bookmarkXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BookmarkXML); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BookmarkXML, _ptr_bookmarkXml); err != nil {
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
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterRemoteSubscriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// channelPath {in} (1:{string, pointer=unique, range=(0,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_channelPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPath); err != nil {
				return err
			}
			return nil
		})
		_s_channelPath := func(ptr interface{}) { o.ChannelPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.ChannelPath, _s_channelPath, _ptr_channelPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Query); err != nil {
			return err
		}
	}
	// bookmarkXml {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_bookmarkXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BookmarkXML); err != nil {
				return err
			}
			return nil
		})
		_s_bookmarkXml := func(ptr interface{}) { o.BookmarkXML = *ptr.(*string) }
		if err := w.ReadPointer(&o.BookmarkXML, _s_bookmarkXml, _ptr_bookmarkXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterRemoteSubscriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.QueryChannelInfo != nil && o.QueryChannelInfoSize == 0 {
		o.QueryChannelInfoSize = uint32(len(o.QueryChannelInfo))
	}
	if len(o.QueryChannelInfo) > int(512) {
		return fmt.Errorf("QueryChannelInfo is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterRemoteSubscriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteSubscription{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// control {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control != nil {
			if err := o.Control.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// queryChannelInfoSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueryChannelInfoSize); err != nil {
			return err
		}
	}
	// queryChannelInfo {out} (1:{pointer=ref, range=(0,512)}*(2)*(1))(2:{alias=EvtRpcQueryChannelInfo}[dim:0,size_is=queryChannelInfoSize](struct))
	{
		if o.QueryChannelInfo != nil || o.QueryChannelInfoSize > 0 {
			_ptr_queryChannelInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryChannelInfoSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryChannelInfo {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.QueryChannelInfo[i1] != nil {
						if err := o.QueryChannelInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&QueryChannelInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.QueryChannelInfo); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&QueryChannelInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryChannelInfo, _ptr_queryChannelInfo); err != nil {
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
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterRemoteSubscriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &RemoteSubscription{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// control {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control == nil {
			o.Control = &OperationControl{}
		}
		if err := o.Control.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// queryChannelInfoSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueryChannelInfoSize); err != nil {
			return err
		}
	}
	// queryChannelInfo {out} (1:{pointer=ref, range=(0,512)}*(2)*(1))(2:{alias=EvtRpcQueryChannelInfo}[dim:0,size_is=queryChannelInfoSize](struct))
	{
		_ptr_queryChannelInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryChannelInfo", sizeInfo[0])
			}
			o.QueryChannelInfo = make([]*QueryChannelInfo, sizeInfo[0])
			for i1 := range o.QueryChannelInfo {
				i1 := i1
				if o.QueryChannelInfo[i1] == nil {
					o.QueryChannelInfo[i1] = &QueryChannelInfo{}
				}
				if err := o.QueryChannelInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_queryChannelInfo := func(ptr interface{}) { o.QueryChannelInfo = *ptr.(*[]*QueryChannelInfo) }
		if err := w.ReadPointer(&o.QueryChannelInfo, _s_queryChannelInfo, _ptr_queryChannelInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterRemoteSubscriptionRequest structure represents the EvtRpcRegisterRemoteSubscription operation request
type RegisterRemoteSubscriptionRequest struct {
	// channelPath: A pointer to a string that contains a channel name or is a null pointer.
	// In the case of a null pointer, the query field indicates the channels to which the
	// subscription applies.
	ChannelPath string `idl:"name:channelPath;string;pointer:unique" json:"channel_path"`
	// query: A pointer to a string that contains a query that specifies events of interest
	// to the application. The pointer MUST be either an XPath filter, as specified in section
	// 2.2.15, or a query as specified in section 2.2.16.
	Query string `idl:"name:query;string" json:"query"`
	// bookmarkXml: Either NULL or a pointer to a string that contains a bookmark indicating
	// the last event that the client processed during a previous subscription. The server
	// MUST ignore the bookmarkXML parameter unless the flags field has the bit 0x00000003
	// set.
	BookmarkXML string `idl:"name:bookmarkXml;string;pointer:unique" json:"bookmark_xml"`
	// flags: Flags that determine the behavior of the query.
	//
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//	|                                            |                                                                    |
	//	|                   VALUE                    |                              MEANING                               |
	//	|                                            |                                                                    |
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//	| EvtSubscribeToFutureEvents 0x00000001      | Get events starting from the present time.                         |
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//	| EvtSubscribeStartAtOldestRecord 0x00000002 | Get all events from the logs, and any future events.               |
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//	| EvtSubscribeStartAfterBookmark 0x00000003  | Get all events starting after the event indicated by the bookmark. |
	//	+--------------------------------------------+--------------------------------------------------------------------+
	//
	// The following bits control other aspects of the subscription. These bits are set
	// independently of the flags defined for the lower two bits, and independently of each
	// other.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSubscribeTolerateQueryErrors 0x00001000 | The server does not fail the function as long as there is one valid channel.     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSubscribeStrict 0x00010000              | Fail if any events are missed for reasons such as log clearing.                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSubscribePull 0x10000000                | Subscription is going to be a pull subscription. A pull subscription requires    |
	//	|                                            | the client to call the EvtRpcRemoteSubscriptionNext (as specified in section     |
	//	|                                            | 3.1.4.10) method to fetch the subscribed events. If this flag is not set, the    |
	//	|                                            | subscription is a push subscription. A push subscription requires the client to  |
	//	|                                            | call the EvtRpcRemoteSubscriptionNextAsync (as specified in section 3.1.4.9) to  |
	//	|                                            | receive notifications from the server when the subscribed events arrive.         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RegisterRemoteSubscriptionRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterRemoteSubscriptionOperation) *xxx_RegisterRemoteSubscriptionOperation {
	if op == nil {
		op = &xxx_RegisterRemoteSubscriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelPath = o.ChannelPath
	op.Query = o.Query
	op.BookmarkXML = o.BookmarkXML
	op.Flags = o.Flags
	return op
}

func (o *RegisterRemoteSubscriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterRemoteSubscriptionOperation) {
	if o == nil {
		return
	}
	o.ChannelPath = op.ChannelPath
	o.Query = op.Query
	o.BookmarkXML = op.BookmarkXML
	o.Flags = op.Flags
}
func (o *RegisterRemoteSubscriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterRemoteSubscriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterRemoteSubscriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterRemoteSubscriptionResponse structure represents the EvtRpcRegisterRemoteSubscription operation response
type RegisterRemoteSubscriptionResponse struct {
	// handle: A context handle for the subscription. This parameter is an RPC context handle,
	// as specified in [C706], Context Handles.
	Handle *RemoteSubscription `idl:"name:handle" json:"handle"`
	// control: A context handle for the subscription. This parameter is an RPC context
	// handle, as specified in [C706], Context Handles.
	Control *OperationControl `idl:"name:control" json:"control"`
	// queryChannelInfoSize: A pointer to a 32-bit unsigned integer that contains the number
	// of EvtRpcQueryChannelInfo structures returned in queryChannelInfo.
	QueryChannelInfoSize uint32 `idl:"name:queryChannelInfoSize" json:"query_channel_info_size"`
	// queryChannelInfo: A pointer to an array of EvtRpcQueryChannelInfo (section 2.2.11)
	// structures that indicate the status of each channel in the subscription.
	QueryChannelInfo []*QueryChannelInfo `idl:"name:queryChannelInfo;size_is:(, queryChannelInfoSize)" json:"query_channel_info"`
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to parsing the query. If the
	// method succeeds, the server MUST set all of the values in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcRegisterRemoteSubscription return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RegisterRemoteSubscriptionResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterRemoteSubscriptionOperation) *xxx_RegisterRemoteSubscriptionOperation {
	if op == nil {
		op = &xxx_RegisterRemoteSubscriptionOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Control = o.Control
	op.QueryChannelInfoSize = o.QueryChannelInfoSize
	op.QueryChannelInfo = o.QueryChannelInfo
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *RegisterRemoteSubscriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterRemoteSubscriptionOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Control = op.Control
	o.QueryChannelInfoSize = op.QueryChannelInfoSize
	o.QueryChannelInfo = op.QueryChannelInfo
	o.Error = op.Error
	o.Return = op.Return
}
func (o *RegisterRemoteSubscriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterRemoteSubscriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterRemoteSubscriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteSubscriptionNextAsyncOperation structure represents the EvtRpcRemoteSubscriptionNextAsync operation
type xxx_RemoteSubscriptionNextAsyncOperation struct {
	Handle                 *RemoteSubscription `idl:"name:handle" json:"handle"`
	RequestedRecordsLength uint32              `idl:"name:numRequestedRecords" json:"requested_records_length"`
	Flags                  uint32              `idl:"name:flags" json:"flags"`
	ActualRecordsLength    uint32              `idl:"name:numActualRecords" json:"actual_records_length"`
	EventDataIndices       []uint32            `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	EventDataSizes         []uint32            `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	ResultBufferSize       uint32              `idl:"name:resultBufferSize" json:"result_buffer_size"`
	ResultBuffer           []byte              `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	Return                 uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) OpNum() int { return 1 }

func (o *xxx_RemoteSubscriptionNextAsyncOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRemoteSubscriptionNextAsync"
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteSubscription{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &RemoteSubscription{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.EventDataIndices != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataIndices))
	}
	if o.EventDataSizes != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataSizes))
	}
	if o.ResultBuffer != nil && o.ResultBufferSize == 0 {
		o.ResultBufferSize = uint32(len(o.ResultBuffer))
	}
	if len(o.EventDataIndices) > int(1024) {
		return fmt.Errorf("EventDataIndices is out of range")
	}
	if len(o.EventDataSizes) > int(1024) {
		return fmt.Errorf("EventDataSizes is out of range")
	}
	if len(o.ResultBuffer) > int(2097152) {
		return fmt.Errorf("ResultBuffer is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataIndices != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataIndices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataIndices {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataIndices[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataIndices); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataIndices, _ptr_eventDataIndices); err != nil {
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
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataSizes != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataSizes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataSizes {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataSizes[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataSizes); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataSizes, _ptr_eventDataSizes); err != nil {
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
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		if o.ResultBuffer != nil || o.ResultBufferSize > 0 {
			_ptr_resultBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultBuffer, _ptr_resultBuffer); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataIndices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataIndices", sizeInfo[0])
			}
			o.EventDataIndices = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataIndices {
				i1 := i1
				if err := w.ReadData(&o.EventDataIndices[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataIndices := func(ptr interface{}) { o.EventDataIndices = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataIndices, _s_eventDataIndices, _ptr_eventDataIndices); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataSizes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataSizes", sizeInfo[0])
			}
			o.EventDataSizes = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataSizes {
				i1 := i1
				if err := w.ReadData(&o.EventDataSizes[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataSizes := func(ptr interface{}) { o.EventDataSizes = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataSizes, _s_eventDataSizes, _ptr_eventDataSizes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		_ptr_resultBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultBuffer", sizeInfo[0])
			}
			o.ResultBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.ResultBuffer {
				i1 := i1
				if err := w.ReadData(&o.ResultBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_resultBuffer := func(ptr interface{}) { o.ResultBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ResultBuffer, _s_resultBuffer, _ptr_resultBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteSubscriptionNextAsyncRequest structure represents the EvtRpcRemoteSubscriptionNextAsync operation request
type RemoteSubscriptionNextAsyncRequest struct {
	// handle: A handle to the subscription. This parameter is an RPC context handle, as
	// specified in [C706], Context Handles.
	Handle *RemoteSubscription `idl:"name:handle" json:"handle"`
	// numRequestedRecords: A 32-bit unsigned integer that contains the number of events
	// to return.
	RequestedRecordsLength uint32 `idl:"name:numRequestedRecords" json:"requested_records_length"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<13>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RemoteSubscriptionNextAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionNextAsyncOperation) *xxx_RemoteSubscriptionNextAsyncOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionNextAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.RequestedRecordsLength = o.RequestedRecordsLength
	op.Flags = o.Flags
	return op
}

func (o *RemoteSubscriptionNextAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionNextAsyncOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.RequestedRecordsLength = op.RequestedRecordsLength
	o.Flags = op.Flags
}
func (o *RemoteSubscriptionNextAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteSubscriptionNextAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionNextAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteSubscriptionNextAsyncResponse structure represents the EvtRpcRemoteSubscriptionNextAsync operation response
type RemoteSubscriptionNextAsyncResponse struct {
	// numActualRecords: A pointer to a 32-bit unsigned integer that contains the value
	// that, on success, MUST be set to the number of events retrieved. This might be used,
	// for example, if the method times out without receiving the full number of events
	// specified in numRequestedRecords.
	ActualRecordsLength uint32 `idl:"name:numActualRecords" json:"actual_records_length"`
	// eventDataIndices: A pointer to an array of 32-bit unsigned integers that contain
	// the offsets for the event. An event's offset is its position relative to the start
	// of resultBuffer.
	EventDataIndices []uint32 `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	// eventDataSizes: A pointer to an array of 32-bit unsigned integers that contain the
	// event sizes in bytes.
	EventDataSizes []uint32 `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	// resultBufferSize: A pointer to a 32-bit unsigned integer that contains the number
	// of bytes of data returned in resultBuffer.
	ResultBufferSize uint32 `idl:"name:resultBufferSize" json:"result_buffer_size"`
	// resultBuffer: A pointer to a byte-array that contains the result set of one or more
	// events. The events MUST be in binary XML format, as specified in section 2.2.17.
	ResultBuffer []byte `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	// Return: The EvtRpcRemoteSubscriptionNextAsync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoteSubscriptionNextAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionNextAsyncOperation) *xxx_RemoteSubscriptionNextAsyncOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionNextAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.ActualRecordsLength = o.ActualRecordsLength
	op.EventDataIndices = o.EventDataIndices
	op.EventDataSizes = o.EventDataSizes
	op.ResultBufferSize = o.ResultBufferSize
	op.ResultBuffer = o.ResultBuffer
	op.Return = o.Return
	return op
}

func (o *RemoteSubscriptionNextAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionNextAsyncOperation) {
	if o == nil {
		return
	}
	o.ActualRecordsLength = op.ActualRecordsLength
	o.EventDataIndices = op.EventDataIndices
	o.EventDataSizes = op.EventDataSizes
	o.ResultBufferSize = op.ResultBufferSize
	o.ResultBuffer = op.ResultBuffer
	o.Return = op.Return
}
func (o *RemoteSubscriptionNextAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteSubscriptionNextAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionNextAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteSubscriptionNextOperation structure represents the EvtRpcRemoteSubscriptionNext operation
type xxx_RemoteSubscriptionNextOperation struct {
	Handle                 *RemoteSubscription `idl:"name:handle" json:"handle"`
	RequestedRecordsLength uint32              `idl:"name:numRequestedRecords" json:"requested_records_length"`
	Timeout                uint32              `idl:"name:timeOut" json:"timeout"`
	Flags                  uint32              `idl:"name:flags" json:"flags"`
	ActualRecordsLength    uint32              `idl:"name:numActualRecords" json:"actual_records_length"`
	EventDataIndices       []uint32            `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	EventDataSizes         []uint32            `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	ResultBufferSize       uint32              `idl:"name:resultBufferSize" json:"result_buffer_size"`
	ResultBuffer           []byte              `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	Return                 uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteSubscriptionNextOperation) OpNum() int { return 2 }

func (o *xxx_RemoteSubscriptionNextOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRemoteSubscriptionNext"
}

func (o *xxx_RemoteSubscriptionNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteSubscription{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// timeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &RemoteSubscription{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// timeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.EventDataIndices != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataIndices))
	}
	if o.EventDataSizes != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataSizes))
	}
	if o.ResultBuffer != nil && o.ResultBufferSize == 0 {
		o.ResultBufferSize = uint32(len(o.ResultBuffer))
	}
	if len(o.EventDataIndices) > int(1024) {
		return fmt.Errorf("EventDataIndices is out of range")
	}
	if len(o.EventDataSizes) > int(1024) {
		return fmt.Errorf("EventDataSizes is out of range")
	}
	if len(o.ResultBuffer) > int(2097152) {
		return fmt.Errorf("ResultBuffer is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataIndices != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataIndices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataIndices {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataIndices[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataIndices); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataIndices, _ptr_eventDataIndices); err != nil {
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
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataSizes != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataSizes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataSizes {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataSizes[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataSizes); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataSizes, _ptr_eventDataSizes); err != nil {
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
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		if o.ResultBuffer != nil || o.ResultBufferSize > 0 {
			_ptr_resultBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultBuffer, _ptr_resultBuffer); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataIndices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataIndices", sizeInfo[0])
			}
			o.EventDataIndices = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataIndices {
				i1 := i1
				if err := w.ReadData(&o.EventDataIndices[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataIndices := func(ptr interface{}) { o.EventDataIndices = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataIndices, _s_eventDataIndices, _ptr_eventDataIndices); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataSizes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataSizes", sizeInfo[0])
			}
			o.EventDataSizes = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataSizes {
				i1 := i1
				if err := w.ReadData(&o.EventDataSizes[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataSizes := func(ptr interface{}) { o.EventDataSizes = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataSizes, _s_eventDataSizes, _ptr_eventDataSizes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		_ptr_resultBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultBuffer", sizeInfo[0])
			}
			o.ResultBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.ResultBuffer {
				i1 := i1
				if err := w.ReadData(&o.ResultBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_resultBuffer := func(ptr interface{}) { o.ResultBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ResultBuffer, _s_resultBuffer, _ptr_resultBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteSubscriptionNextRequest structure represents the EvtRpcRemoteSubscriptionNext operation request
type RemoteSubscriptionNextRequest struct {
	// handle: A handle to a subscription. This parameter is an RPC context handle, as specified
	// in [C706] Context Handles.
	Handle *RemoteSubscription `idl:"name:handle" json:"handle"`
	// numRequestedRecords: A 32-bit unsigned integer that contains the maximum number of
	// events to return.
	RequestedRecordsLength uint32 `idl:"name:numRequestedRecords" json:"requested_records_length"`
	// timeOut: A 32-bit unsigned integer that contains the maximum number of milliseconds
	// to wait before returning.
	Timeout uint32 `idl:"name:timeOut" json:"timeout"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<15>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RemoteSubscriptionNextRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionNextOperation) *xxx_RemoteSubscriptionNextOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionNextOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.RequestedRecordsLength = o.RequestedRecordsLength
	op.Timeout = o.Timeout
	op.Flags = o.Flags
	return op
}

func (o *RemoteSubscriptionNextRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionNextOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.RequestedRecordsLength = op.RequestedRecordsLength
	o.Timeout = op.Timeout
	o.Flags = op.Flags
}
func (o *RemoteSubscriptionNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteSubscriptionNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteSubscriptionNextResponse structure represents the EvtRpcRemoteSubscriptionNext operation response
type RemoteSubscriptionNextResponse struct {
	// numActualRecords: A pointer to a 32-bit unsigned integer that contains the value
	// that, on success, MUST be set to the number of events that are retrieved. This is
	// useful in the case in which the method times out without receiving the full number
	// of events specified in numRequestedRecords. If the method fails, the client MUST
	// NOT use the value.
	ActualRecordsLength uint32 `idl:"name:numActualRecords" json:"actual_records_length"`
	// eventDataIndices: A pointer to an array of 32-bit unsigned integers that contain
	// the offsets for the events. An event offset is its position relative to the start
	// of resultBuffer.
	EventDataIndices []uint32 `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	// eventDataSizes: A pointer to an array of 32-bit unsigned integers that contain the
	// event sizes in bytes.
	EventDataSizes []uint32 `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	// resultBufferSize: A pointer to a 32-bit unsigned integer that contains the number
	// of bytes of data returned in resultBuffer.
	ResultBufferSize uint32 `idl:"name:resultBufferSize" json:"result_buffer_size"`
	// resultBuffer: A pointer to a byte-array that contains the result set of one or more
	// events. The events MUST be in binary XML format, as specified in section 2.2.17.
	ResultBuffer []byte `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	// Return: The EvtRpcRemoteSubscriptionNext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoteSubscriptionNextResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionNextOperation) *xxx_RemoteSubscriptionNextOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionNextOperation{}
	}
	if o == nil {
		return op
	}
	op.ActualRecordsLength = o.ActualRecordsLength
	op.EventDataIndices = o.EventDataIndices
	op.EventDataSizes = o.EventDataSizes
	op.ResultBufferSize = o.ResultBufferSize
	op.ResultBuffer = o.ResultBuffer
	op.Return = o.Return
	return op
}

func (o *RemoteSubscriptionNextResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionNextOperation) {
	if o == nil {
		return
	}
	o.ActualRecordsLength = op.ActualRecordsLength
	o.EventDataIndices = op.EventDataIndices
	o.EventDataSizes = op.EventDataSizes
	o.ResultBufferSize = op.ResultBufferSize
	o.ResultBuffer = op.ResultBuffer
	o.Return = op.Return
}
func (o *RemoteSubscriptionNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteSubscriptionNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteSubscriptionWaitAsyncOperation structure represents the EvtRpcRemoteSubscriptionWaitAsync operation
type xxx_RemoteSubscriptionWaitAsyncOperation struct {
	Handle *RemoteSubscription `idl:"name:handle" json:"handle"`
	Return uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) OpNum() int { return 3 }

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRemoteSubscriptionWaitAsync"
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RemoteSubscription{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_REMOTE_SUBSCRIPTION, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &RemoteSubscription{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteSubscriptionWaitAsyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteSubscriptionWaitAsyncRequest structure represents the EvtRpcRemoteSubscriptionWaitAsync operation request
type RemoteSubscriptionWaitAsyncRequest struct {
	// handle: A handle to a subscription, as obtained from the EvtRpcRegisterRemoteSubscription
	// (section 3.1.4.8) method. This parameter MUST be an RPC context handle, as specified
	// in [C706] Context Handles.
	Handle *RemoteSubscription `idl:"name:handle" json:"handle"`
}

func (o *RemoteSubscriptionWaitAsyncRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionWaitAsyncOperation) *xxx_RemoteSubscriptionWaitAsyncOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionWaitAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	return op
}

func (o *RemoteSubscriptionWaitAsyncRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionWaitAsyncOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
}
func (o *RemoteSubscriptionWaitAsyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteSubscriptionWaitAsyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionWaitAsyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteSubscriptionWaitAsyncResponse structure represents the EvtRpcRemoteSubscriptionWaitAsync operation response
type RemoteSubscriptionWaitAsyncResponse struct {
	// Return: The EvtRpcRemoteSubscriptionWaitAsync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoteSubscriptionWaitAsyncResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteSubscriptionWaitAsyncOperation) *xxx_RemoteSubscriptionWaitAsyncOperation {
	if op == nil {
		op = &xxx_RemoteSubscriptionWaitAsyncOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RemoteSubscriptionWaitAsyncResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteSubscriptionWaitAsyncOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoteSubscriptionWaitAsyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteSubscriptionWaitAsyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteSubscriptionWaitAsyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterControllableOperationOperation structure represents the EvtRpcRegisterControllableOperation operation
type xxx_RegisterControllableOperationOperation struct {
	Handle *OperationControl `idl:"name:handle" json:"handle"`
	Return uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterControllableOperationOperation) OpNum() int { return 4 }

func (o *xxx_RegisterControllableOperationOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRegisterControllableOperation"
}

func (o *xxx_RegisterControllableOperationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterControllableOperationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_RegisterControllableOperationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_RegisterControllableOperationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterControllableOperationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterControllableOperationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &OperationControl{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterControllableOperationRequest structure represents the EvtRpcRegisterControllableOperation operation request
type RegisterControllableOperationRequest struct {
}

func (o *RegisterControllableOperationRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterControllableOperationOperation) *xxx_RegisterControllableOperationOperation {
	if op == nil {
		op = &xxx_RegisterControllableOperationOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *RegisterControllableOperationRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterControllableOperationOperation) {
	if o == nil {
		return
	}
}
func (o *RegisterControllableOperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterControllableOperationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterControllableOperationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterControllableOperationResponse structure represents the EvtRpcRegisterControllableOperation operation response
type RegisterControllableOperationResponse struct {
	// handle: A context handle for a control object. This parameter MUST be an RPC context
	// handle, as specified in [C706], Context Handles. For information on handle security
	// and authentication considerations, see sections 2.2.20 and 5.1.
	Handle *OperationControl `idl:"name:handle" json:"handle"`
	// Return: The EvtRpcRegisterControllableOperation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RegisterControllableOperationResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterControllableOperationOperation) *xxx_RegisterControllableOperationOperation {
	if op == nil {
		op = &xxx_RegisterControllableOperationOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *RegisterControllableOperationResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterControllableOperationOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *RegisterControllableOperationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterControllableOperationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterControllableOperationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterLogQueryOperation structure represents the EvtRpcRegisterLogQuery operation
type xxx_RegisterLogQueryOperation struct {
	Path                 string              `idl:"name:path;string;pointer:unique" json:"path"`
	Query                string              `idl:"name:query;string" json:"query"`
	Flags                uint32              `idl:"name:flags" json:"flags"`
	Handle               *LogQuery           `idl:"name:handle" json:"handle"`
	OperationControl     *OperationControl   `idl:"name:opControl" json:"operation_control"`
	QueryChannelInfoSize uint32              `idl:"name:queryChannelInfoSize" json:"query_channel_info_size"`
	QueryChannelInfo     []*QueryChannelInfo `idl:"name:queryChannelInfo;size_is:(, queryChannelInfoSize)" json:"query_channel_info"`
	Error                *Info               `idl:"name:error" json:"error"`
	Return               uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterLogQueryOperation) OpNum() int { return 5 }

func (o *xxx_RegisterLogQueryOperation) OpName() string {
	return "/IEventService/v1/EvtRpcRegisterLogQuery"
}

func (o *xxx_RegisterLogQueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Path) > int(32768) {
		return fmt.Errorf("Path is out of range")
	}
	if len(o.Query) > int(1048576) {
		return fmt.Errorf("Query is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterLogQueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Path != "" {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Query); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterLogQueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(*string) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Query); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterLogQueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.QueryChannelInfo != nil && o.QueryChannelInfoSize == 0 {
		o.QueryChannelInfoSize = uint32(len(o.QueryChannelInfo))
	}
	if len(o.QueryChannelInfo) > int(512) {
		return fmt.Errorf("QueryChannelInfo is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterLogQueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LogQuery{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// opControl {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.OperationControl != nil {
			if err := o.OperationControl.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// queryChannelInfoSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueryChannelInfoSize); err != nil {
			return err
		}
	}
	// queryChannelInfo {out} (1:{pointer=ref, range=(0,512)}*(2)*(1))(2:{alias=EvtRpcQueryChannelInfo}[dim:0,size_is=queryChannelInfoSize](struct))
	{
		if o.QueryChannelInfo != nil || o.QueryChannelInfoSize > 0 {
			_ptr_queryChannelInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryChannelInfoSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryChannelInfo {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.QueryChannelInfo[i1] != nil {
						if err := o.QueryChannelInfo[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&QueryChannelInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.QueryChannelInfo); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&QueryChannelInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryChannelInfo, _ptr_queryChannelInfo); err != nil {
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
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterLogQueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &LogQuery{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// opControl {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.OperationControl == nil {
			o.OperationControl = &OperationControl{}
		}
		if err := o.OperationControl.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// queryChannelInfoSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueryChannelInfoSize); err != nil {
			return err
		}
	}
	// queryChannelInfo {out} (1:{pointer=ref, range=(0,512)}*(2)*(1))(2:{alias=EvtRpcQueryChannelInfo}[dim:0,size_is=queryChannelInfoSize](struct))
	{
		_ptr_queryChannelInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryChannelInfo", sizeInfo[0])
			}
			o.QueryChannelInfo = make([]*QueryChannelInfo, sizeInfo[0])
			for i1 := range o.QueryChannelInfo {
				i1 := i1
				if o.QueryChannelInfo[i1] == nil {
					o.QueryChannelInfo[i1] = &QueryChannelInfo{}
				}
				if err := o.QueryChannelInfo[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_queryChannelInfo := func(ptr interface{}) { o.QueryChannelInfo = *ptr.(*[]*QueryChannelInfo) }
		if err := w.ReadPointer(&o.QueryChannelInfo, _s_queryChannelInfo, _ptr_queryChannelInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterLogQueryRequest structure represents the EvtRpcRegisterLogQuery operation request
type RegisterLogQueryRequest struct {
	// path: A pointer to a string that contains a channel or file path.
	Path string `idl:"name:path;string;pointer:unique" json:"path"`
	// query: A pointer to a string that contains a query that specifies events of interest
	// to the application. The pointer MUST be either an XPath filter, as specified in section
	// 2.2.15, or a query, as specified in section 2.2.16.
	Query string `idl:"name:query;string" json:"query"`
	// flags: The flags field MUST be set as follows. The first two bits indicate how the
	// path argument MUST be interpreted. Callers MUST specify one and only one value.
	//
	//	+--------------------------------+--------------------------------+
	//	|                                |                                |
	//	|             VALUE              |            MEANING             |
	//	|                                |                                |
	//	+--------------------------------+--------------------------------+
	//	+--------------------------------+--------------------------------+
	//	| EvtQueryChannelPath 0x00000001 | Path specifies a channel name. |
	//	+--------------------------------+--------------------------------+
	//	| EvtQueryFilePath 0x00000002    | Path specifies a file name.    |
	//	+--------------------------------+--------------------------------+
	//
	// These bits control the direction of the query. Callers MUST specify one and only
	// one value.
	//
	//	+------------+----------------------------------------+
	//	|            |                                        |
	//	|   VALUE    |                MEANING                 |
	//	|            |                                        |
	//	+------------+----------------------------------------+
	//	+------------+----------------------------------------+
	//	| 0x00000100 | Events are read from oldest to newest. |
	//	+------------+----------------------------------------+
	//	| 0x00000200 | Events are read from newest to oldest. |
	//	+------------+----------------------------------------+
	//
	// The following bit can be set independently of the previously mentioned bits.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00001000 | Specifies to return the query result set even if one or more errors result from  |
	//	|            | the query. For example, if a structured XML query specifies multiple channels,   |
	//	|            | some channels are valid while others are not. A query that is used on many       |
	//	|            | computers might be sent to a computer that is missing one or more channels in    |
	//	|            | the query. If this bit is not set, the server MUST fail the query. If this bit   |
	//	|            | is set, the query MUST succeed even if all channels are not present.             |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RegisterLogQueryRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterLogQueryOperation) *xxx_RegisterLogQueryOperation {
	if op == nil {
		op = &xxx_RegisterLogQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.Path = o.Path
	op.Query = o.Query
	op.Flags = o.Flags
	return op
}

func (o *RegisterLogQueryRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterLogQueryOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Query = op.Query
	o.Flags = op.Flags
}
func (o *RegisterLogQueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterLogQueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterLogQueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterLogQueryResponse structure represents the EvtRpcRegisterLogQuery operation response
type RegisterLogQueryResponse struct {
	// handle: A pointer to a query handle. This parameter MUST be an RPC context handle,
	// as specified in [C706], Context Handles.
	Handle *LogQuery `idl:"name:handle" json:"handle"`
	// opControl: A pointer to a control handle. This parameter MUST be an RPC context handle,
	// as specified in [C706], Context Handles.
	OperationControl *OperationControl `idl:"name:opControl" json:"operation_control"`
	// queryChannelInfoSize: A pointer to a 32-bit unsigned integer that contains the number
	// of EvtRpcQueryChannelInfo structures returned in queryChannelInfo.
	QueryChannelInfoSize uint32 `idl:"name:queryChannelInfoSize" json:"query_channel_info_size"`
	// queryChannelInfo: A pointer to an array of section EvtRpcQueryChannelInfo structures,
	// as specified in section 2.2.11.
	QueryChannelInfo []*QueryChannelInfo `idl:"name:queryChannelInfo;size_is:(, queryChannelInfoSize)" json:"query_channel_info"`
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to parsing the query; in addition,
	// the server MAY set the structure fields to nonzero values for errors unrelated to
	// query parsing (for example, for an invalid channel name). All nonzero values MUST
	// be treated the same. If the method succeeds, the server MUST set all the fields in
	// the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcRegisterLogQuery return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RegisterLogQueryResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterLogQueryOperation) *xxx_RegisterLogQueryOperation {
	if op == nil {
		op = &xxx_RegisterLogQueryOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.OperationControl = o.OperationControl
	op.QueryChannelInfoSize = o.QueryChannelInfoSize
	op.QueryChannelInfo = o.QueryChannelInfo
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *RegisterLogQueryResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterLogQueryOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.OperationControl = op.OperationControl
	o.QueryChannelInfoSize = op.QueryChannelInfoSize
	o.QueryChannelInfo = op.QueryChannelInfo
	o.Error = op.Error
	o.Return = op.Return
}
func (o *RegisterLogQueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterLogQueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterLogQueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearLogOperation structure represents the EvtRpcClearLog operation
type xxx_ClearLogOperation struct {
	Control     *OperationControl `idl:"name:control" json:"control"`
	ChannelPath string            `idl:"name:channelPath;string" json:"channel_path"`
	BackupPath  string            `idl:"name:backupPath;string;pointer:unique" json:"backup_path"`
	Flags       uint32            `idl:"name:flags" json:"flags"`
	Error       *Info             `idl:"name:error" json:"error"`
	Return      uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearLogOperation) OpNum() int { return 6 }

func (o *xxx_ClearLogOperation) OpName() string { return "/IEventService/v1/EvtRpcClearLog" }

func (o *xxx_ClearLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ChannelPath) > int(512) {
		return fmt.Errorf("ChannelPath is out of range")
	}
	if len(o.BackupPath) > int(32768) {
		return fmt.Errorf("BackupPath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control != nil {
			if err := o.Control.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// channelPath {in} (1:{string, range=(0,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPath); err != nil {
			return err
		}
	}
	// backupPath {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BackupPath != "" {
			_ptr_backupPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BackupPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BackupPath, _ptr_backupPath); err != nil {
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
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control == nil {
			o.Control = &OperationControl{}
		}
		if err := o.Control.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// channelPath {in} (1:{string, range=(0,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPath); err != nil {
			return err
		}
	}
	// backupPath {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_backupPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BackupPath); err != nil {
				return err
			}
			return nil
		})
		_s_backupPath := func(ptr interface{}) { o.BackupPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.BackupPath, _s_backupPath, _ptr_backupPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearLogRequest structure represents the EvtRpcClearLog operation request
type ClearLogRequest struct {
	// control: A handle to an operation control object. This parameter is an RPC context
	// handle, as specified in [C706], Context Handles.
	Control *OperationControl `idl:"name:control" json:"control"`
	// channelPath:  A pointer to a string that contains the path of the channel to be
	// cleared.
	ChannelPath string `idl:"name:channelPath;string" json:"channel_path"`
	// backupPath: A pointer to a string that contains the path of the file in which events
	// are to be saved before the clear is performed. A value of NULL indicates that no
	// backup event log is to be created (the events to be cleared are not to be saved).
	BackupPath string `idl:"name:backupPath;string;pointer:unique" json:"backup_path"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<25>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *ClearLogRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearLogOperation) *xxx_ClearLogOperation {
	if op == nil {
		op = &xxx_ClearLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Control = o.Control
	op.ChannelPath = o.ChannelPath
	op.BackupPath = o.BackupPath
	op.Flags = o.Flags
	return op
}

func (o *ClearLogRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearLogOperation) {
	if o == nil {
		return
	}
	o.Control = op.Control
	o.ChannelPath = op.ChannelPath
	o.BackupPath = op.BackupPath
	o.Flags = op.Flags
}
func (o *ClearLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearLogResponse structure represents the EvtRpcClearLog operation response
type ClearLogResponse struct {
	// error: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific nonzero value as specified in [MS-ERREF].<26>
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcClearLog return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClearLogResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearLogOperation) *xxx_ClearLogOperation {
	if op == nil {
		op = &xxx_ClearLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *ClearLogResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearLogOperation) {
	if o == nil {
		return
	}
	o.Error = op.Error
	o.Return = op.Return
}
func (o *ClearLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportLogOperation structure represents the EvtRpcExportLog operation
type xxx_ExportLogOperation struct {
	Control     *OperationControl `idl:"name:control" json:"control"`
	ChannelPath string            `idl:"name:channelPath;string;pointer:unique" json:"channel_path"`
	Query       string            `idl:"name:query;string" json:"query"`
	BackupPath  string            `idl:"name:backupPath;string" json:"backup_path"`
	Flags       uint32            `idl:"name:flags" json:"flags"`
	Error       *Info             `idl:"name:error" json:"error"`
	Return      uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportLogOperation) OpNum() int { return 7 }

func (o *xxx_ExportLogOperation) OpName() string { return "/IEventService/v1/EvtRpcExportLog" }

func (o *xxx_ExportLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ChannelPath) > int(512) {
		return fmt.Errorf("ChannelPath is out of range")
	}
	if len(o.Query) > int(1048576) {
		return fmt.Errorf("Query is out of range")
	}
	if len(o.BackupPath) > int(32768) {
		return fmt.Errorf("BackupPath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control != nil {
			if err := o.Control.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// channelPath {in} (1:{string, pointer=unique, range=(0,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ChannelPath != "" {
			_ptr_channelPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ChannelPath, _ptr_channelPath); err != nil {
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
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Query); err != nil {
			return err
		}
	}
	// backupPath {in} (1:{string, range=(1,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.BackupPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control == nil {
			o.Control = &OperationControl{}
		}
		if err := o.Control.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// channelPath {in} (1:{string, pointer=unique, range=(0,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_channelPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPath); err != nil {
				return err
			}
			return nil
		})
		_s_channelPath := func(ptr interface{}) { o.ChannelPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.ChannelPath, _s_channelPath, _ptr_channelPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// query {in} (1:{string, range=(1,1048576), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Query); err != nil {
			return err
		}
	}
	// backupPath {in} (1:{string, range=(1,32768), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.BackupPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ExportLogRequest structure represents the EvtRpcExportLog operation request
type ExportLogRequest struct {
	// control: A handle to an operation control object. This parameter is an RPC context
	// handle, as specified in [C706] Context Handles.
	Control *OperationControl `idl:"name:control" json:"control"`
	// channelPath: A pointer to a string that contains the channel name (for a live event
	// log) or file path (for an existing backup event log) to be used to create a backup
	// event log.
	ChannelPath string `idl:"name:channelPath;string;pointer:unique" json:"channel_path"`
	// query: A pointer to a string that contains a query that specifies events to be included
	// in the backup event log.
	Query string `idl:"name:query;string" json:"query"`
	// backupPath: A pointer to a string that contains the path of the file for the backup
	// event logs to be created.
	BackupPath string `idl:"name:backupPath;string" json:"backup_path"`
	// flags: The client MUST set the flags parameter to one of the following values.
	//
	//	+--------------------------------+---------------------------------------------+
	//	|                                |                                             |
	//	|             VALUE              |                   MEANING                   |
	//	|                                |                                             |
	//	+--------------------------------+---------------------------------------------+
	//	+--------------------------------+---------------------------------------------+
	//	| EvtQueryChannelPath 0x00000001 | Channel parameter specifies a channel name. |
	//	+--------------------------------+---------------------------------------------+
	//	| EvtQueryFilePath 0x00000002    | Channel parameter specifies a file name.    |
	//	+--------------------------------+---------------------------------------------+
	//
	// In addition, the client MAY set the following value in the flags parameter:
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtQueryTolerateQueryErrors 0x00001000 | The query MUST succeed even if not all the channels or backup event logs         |
	//	|                                        | that are specified in the query are present, or if the channelPath parameter     |
	//	|                                        | specifies channels that do not exist.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *ExportLogRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportLogOperation) *xxx_ExportLogOperation {
	if op == nil {
		op = &xxx_ExportLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Control = o.Control
	op.ChannelPath = o.ChannelPath
	op.Query = o.Query
	op.BackupPath = o.BackupPath
	op.Flags = o.Flags
	return op
}

func (o *ExportLogRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportLogOperation) {
	if o == nil {
		return
	}
	o.Control = op.Control
	o.ChannelPath = op.ChannelPath
	o.Query = op.Query
	o.BackupPath = op.BackupPath
	o.Flags = op.Flags
}
func (o *ExportLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportLogResponse structure represents the EvtRpcExportLog operation response
type ExportLogResponse struct {
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to a nonzero value if the error is related to parsing the query. In addition,
	// the server MAY set the suberror fields to nonzero values for other types of errors.
	// All nonzero values MUST be treated the same. If the method succeeds, the server MUST
	// set all of the values in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcExportLog return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ExportLogResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportLogOperation) *xxx_ExportLogOperation {
	if op == nil {
		op = &xxx_ExportLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *ExportLogResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportLogOperation) {
	if o == nil {
		return
	}
	o.Error = op.Error
	o.Return = op.Return
}
func (o *ExportLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LocalizeExportLogOperation structure represents the EvtRpcLocalizeExportLog operation
type xxx_LocalizeExportLogOperation struct {
	Control     *OperationControl `idl:"name:control" json:"control"`
	LogFilePath string            `idl:"name:logFilePath;string" json:"log_file_path"`
	Locale      uint32            `idl:"name:locale" json:"locale"`
	Flags       uint32            `idl:"name:flags" json:"flags"`
	Error       *Info             `idl:"name:error" json:"error"`
	Return      uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_LocalizeExportLogOperation) OpNum() int { return 8 }

func (o *xxx_LocalizeExportLogOperation) OpName() string {
	return "/IEventService/v1/EvtRpcLocalizeExportLog"
}

func (o *xxx_LocalizeExportLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.LogFilePath) > int(32768) {
		return fmt.Errorf("LogFilePath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LocalizeExportLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control != nil {
			if err := o.Control.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// logFilePath {in} (1:{string, range=(1,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LogFilePath); err != nil {
			return err
		}
	}
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LocalizeExportLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// control {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Control == nil {
			o.Control = &OperationControl{}
		}
		if err := o.Control.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// logFilePath {in} (1:{string, range=(1,32768), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogFilePath); err != nil {
			return err
		}
	}
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LocalizeExportLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LocalizeExportLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LocalizeExportLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LocalizeExportLogRequest structure represents the EvtRpcLocalizeExportLog operation request
type LocalizeExportLogRequest struct {
	// control: A handle to an operation control object. This parameter MUST be an RPC context
	// handle, as specified in [C706], Context Handles.
	Control *OperationControl `idl:"name:control" json:"control"`
	// logFilePath: A pointer to a string that contains the path of the backup event log
	// to be localized.
	LogFilePath string `idl:"name:logFilePath;string" json:"log_file_path"`
	// locale: Locale, as specified in [MS-GPSI] Appendix A, to be used for localizing the
	// log.
	Locale uint32 `idl:"name:locale" json:"locale"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<31>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *LocalizeExportLogRequest) xxx_ToOp(ctx context.Context, op *xxx_LocalizeExportLogOperation) *xxx_LocalizeExportLogOperation {
	if op == nil {
		op = &xxx_LocalizeExportLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Control = o.Control
	op.LogFilePath = o.LogFilePath
	op.Locale = o.Locale
	op.Flags = o.Flags
	return op
}

func (o *LocalizeExportLogRequest) xxx_FromOp(ctx context.Context, op *xxx_LocalizeExportLogOperation) {
	if o == nil {
		return
	}
	o.Control = op.Control
	o.LogFilePath = op.LogFilePath
	o.Locale = op.Locale
	o.Flags = op.Flags
}
func (o *LocalizeExportLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LocalizeExportLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LocalizeExportLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LocalizeExportLogResponse structure represents the EvtRpcLocalizeExportLog operation response
type LocalizeExportLogResponse struct {
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to loading localization information.
	// All nonzero values MUST be treated the same. If the method succeeds, the server MUST
	// set all of the values in the structure to zero.<32>
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcLocalizeExportLog return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *LocalizeExportLogResponse) xxx_ToOp(ctx context.Context, op *xxx_LocalizeExportLogOperation) *xxx_LocalizeExportLogOperation {
	if op == nil {
		op = &xxx_LocalizeExportLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *LocalizeExportLogResponse) xxx_FromOp(ctx context.Context, op *xxx_LocalizeExportLogOperation) {
	if o == nil {
		return
	}
	o.Error = op.Error
	o.Return = op.Return
}
func (o *LocalizeExportLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LocalizeExportLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LocalizeExportLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MessageRenderOperation structure represents the EvtRpcMessageRender operation
type xxx_MessageRenderOperation struct {
	PublisherConfigObject *PublisherMetadata `idl:"name:pubCfgObj" json:"publisher_config_object"`
	SizeEventID           uint32             `idl:"name:sizeEventId" json:"size_event_id"`
	EventID               []byte             `idl:"name:eventId;size_is:(sizeEventId)" json:"event_id"`
	MessageID             uint32             `idl:"name:messageId" json:"message_id"`
	Values                *VariantList       `idl:"name:values" json:"values"`
	Flags                 uint32             `idl:"name:flags" json:"flags"`
	MaxSizeString         uint32             `idl:"name:maxSizeString" json:"max_size_string"`
	ActualSizeString      uint32             `idl:"name:actualSizeString" json:"actual_size_string"`
	NeededSizeString      uint32             `idl:"name:neededSizeString" json:"needed_size_string"`
	String                []byte             `idl:"name:string;size_is:(, actualSizeString)" json:"string"`
	Error                 *Info              `idl:"name:error" json:"error"`
	Return                uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageRenderOperation) OpNum() int { return 9 }

func (o *xxx_MessageRenderOperation) OpName() string { return "/IEventService/v1/EvtRpcMessageRender" }

func (o *xxx_MessageRenderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.EventID != nil && o.SizeEventID == 0 {
		o.SizeEventID = uint32(len(o.EventID))
	}
	if o.SizeEventID < uint32(1) || o.SizeEventID > uint32(256) {
		return fmt.Errorf("SizeEventID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pubCfgObj {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherConfigObject != nil {
			if err := o.PublisherConfigObject.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PublisherMetadata{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// sizeEventId {in} (1:{range=(1,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeEventID); err != nil {
			return err
		}
	}
	// eventId {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=sizeEventId](uchar))
	{
		dimSize1 := uint64(o.SizeEventID)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.EventID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.EventID[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.EventID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// messageId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MessageID); err != nil {
			return err
		}
	}
	// values {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Values != nil {
			if err := o.Values.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// maxSizeString {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxSizeString); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pubCfgObj {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherConfigObject == nil {
			o.PublisherConfigObject = &PublisherMetadata{}
		}
		if err := o.PublisherConfigObject.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// sizeEventId {in} (1:{range=(1,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeEventID); err != nil {
			return err
		}
	}
	// eventId {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=sizeEventId](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.EventID", sizeInfo[0])
		}
		o.EventID = make([]byte, sizeInfo[0])
		for i1 := range o.EventID {
			i1 := i1
			if err := w.ReadData(&o.EventID[i1]); err != nil {
				return err
			}
		}
	}
	// messageId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MessageID); err != nil {
			return err
		}
	}
	// values {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Values == nil {
			o.Values = &VariantList{}
		}
		if err := o.Values.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// maxSizeString {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxSizeString); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.String != nil && o.ActualSizeString == 0 {
		o.ActualSizeString = uint32(len(o.String))
	}
	if len(o.String) > int(2097152) {
		return fmt.Errorf("String is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// actualSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualSizeString); err != nil {
			return err
		}
	}
	// neededSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NeededSizeString); err != nil {
			return err
		}
	}
	// string {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=actualSizeString](uchar))
	{
		if o.String != nil || o.ActualSizeString > 0 {
			_ptr_string := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualSizeString)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.String {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.String[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.String); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.String, _ptr_string); err != nil {
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
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// actualSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualSizeString); err != nil {
			return err
		}
	}
	// neededSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NeededSizeString); err != nil {
			return err
		}
	}
	// string {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=actualSizeString](uchar))
	{
		_ptr_string := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.String", sizeInfo[0])
			}
			o.String = make([]byte, sizeInfo[0])
			for i1 := range o.String {
				i1 := i1
				if err := w.ReadData(&o.String[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_string := func(ptr interface{}) { o.String = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.String, _s_string, _ptr_string); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageRenderRequest structure represents the EvtRpcMessageRender operation request
type MessageRenderRequest struct {
	// pubCfgObj: A handle to a publisher object. This parameter is an RPC context handle,
	// as specified in [C706], Context Handles. This value comes from the return parameter
	// pubMetadata of the function EvtRpcGetPublisherMetadata (section 3.1.4.25).
	PublisherConfigObject *PublisherMetadata `idl:"name:pubCfgObj" json:"publisher_config_object"`
	// sizeEventId: A 32-bit unsigned integer that contains the size, in bytes, of the data
	// in the eventId field. The server MUST ignore this value if EvtFormatMessageId is
	// specified as the flags parameter. If EvtFormatMessageId is not specified in the flags
	// parameter, the server MUST use the sizeEventId parameter and ignore the messageId
	// parameter.
	SizeEventID uint32 `idl:"name:sizeEventId" json:"size_event_id"`
	// eventId: A pointer to an EVENT_DESCRIPTOR structure, as specified in [MS-DTYP] section
	// 2.3.1.
	EventID []byte `idl:"name:eventId;size_is:(sizeEventId)" json:"event_id"`
	// messageId: A 32-bit unsigned integer that specifies the required message. This is
	// an alternative to using the eventID parameter used by a client application that has
	// obtained the value through some method outside those documented by this protocol.
	// The server MUST ignore this value unless the flags value is set to EvtFormatMessageId;
	// in which case, the server MUST use this value to determine the required message and
	// ignore the eventID parameter.
	MessageID uint32 `idl:"name:messageId" json:"message_id"`
	// values: An array of strings used as substitution values for event description strings.
	// The number of strings submitted is determined by the number of description strings
	// contained in the event message specified by the eventID or messageId parameter.<62>
	Values *VariantList `idl:"name:values" json:"values"`
	// flags: For all options except EvtFormatMessageId, the eventId parameter is used to
	// specify an event descriptor. For the EvtFormatMessageId option, the messageId is
	// used for locating the message. This MUST be set to one of the values in the following
	// table, which indicates the action a server is requested to perform.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageEvent 0x00000001    | Locate the message for the event that corresponds to eventID, and then insert    |
	//	|                                     | the values specified by the values parameter.                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageLevel 0x00000002    | Extract the level field from eventID, and then return the localized name for     |
	//	|                                     | that level.                                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageTask 0x00000003     | Extract the task field from eventID, and then return the localized name for that |
	//	|                                     | task.                                                                            |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageOpcode 0x00000004   | Extract the opcode field from eventID, and then return the localized name for    |
	//	|                                     | that opcode.                                                                     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageKeyword 0x00000005  | Extract the keyword field from eventID, and then return the localized name for   |
	//	|                                     | that keyword.                                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageChannel 0x00000006  | Extract the channel field from eventID, and then return the localized name for   |
	//	|                                     | that channel.                                                                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageProvider 0x00000007 | Return the localized name of the publisher.                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageId 0x00000008       | Locate the message for the event corresponding to the messageId parameter, and   |
	//	|                                     | then insert the values specified by the values parameter.                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// maxSizeString: A 32-bit unsigned integer that contains the size, in bytes, of the
	// string that is provided by the caller.
	MaxSizeString uint32 `idl:"name:maxSizeString" json:"max_size_string"`
}

func (o *MessageRenderRequest) xxx_ToOp(ctx context.Context, op *xxx_MessageRenderOperation) *xxx_MessageRenderOperation {
	if op == nil {
		op = &xxx_MessageRenderOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherConfigObject = o.PublisherConfigObject
	op.SizeEventID = o.SizeEventID
	op.EventID = o.EventID
	op.MessageID = o.MessageID
	op.Values = o.Values
	op.Flags = o.Flags
	op.MaxSizeString = o.MaxSizeString
	return op
}

func (o *MessageRenderRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageRenderOperation) {
	if o == nil {
		return
	}
	o.PublisherConfigObject = op.PublisherConfigObject
	o.SizeEventID = op.SizeEventID
	o.EventID = op.EventID
	o.MessageID = op.MessageID
	o.Values = op.Values
	o.Flags = op.Flags
	o.MaxSizeString = op.MaxSizeString
}
func (o *MessageRenderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MessageRenderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageRenderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageRenderResponse structure represents the EvtRpcMessageRender operation response
type MessageRenderResponse struct {
	// actualSizeString: A pointer to a 32-bit unsigned integer that, on return, contains
	// the actual size, in bytes, of the resulting description (including null termination).
	ActualSizeString uint32 `idl:"name:actualSizeString" json:"actual_size_string"`
	// neededSizeString: A pointer to a 32-bit unsigned integer that, on return, contains
	// the needed size, in bytes (including null termination).
	NeededSizeString uint32 `idl:"name:neededSizeString" json:"needed_size_string"`
	// string: A pointer to a bytearray that, on return, contains a localized string containing
	// the message requested. This can contain a simple string, such as the localized name
	// of a keyword, or a fully rendered message that contains multiple inserts.
	String []byte `idl:"name:string;size_is:(, actualSizeString)" json:"string"`
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to loading the necessary resource.
	// All nonzero values MUST be treated the same. If the method succeeds, the server MUST
	// set all the fields in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcMessageRender return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageRenderResponse) xxx_ToOp(ctx context.Context, op *xxx_MessageRenderOperation) *xxx_MessageRenderOperation {
	if op == nil {
		op = &xxx_MessageRenderOperation{}
	}
	if o == nil {
		return op
	}
	op.ActualSizeString = o.ActualSizeString
	op.NeededSizeString = o.NeededSizeString
	op.String = o.String
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *MessageRenderResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageRenderOperation) {
	if o == nil {
		return
	}
	o.ActualSizeString = op.ActualSizeString
	o.NeededSizeString = op.NeededSizeString
	o.String = op.String
	o.Error = op.Error
	o.Return = op.Return
}
func (o *MessageRenderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MessageRenderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageRenderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MessageRenderDefaultOperation structure represents the EvtRpcMessageRenderDefault operation
type xxx_MessageRenderDefaultOperation struct {
	SizeEventID      uint32       `idl:"name:sizeEventId" json:"size_event_id"`
	EventID          []byte       `idl:"name:eventId;size_is:(sizeEventId)" json:"event_id"`
	MessageID        uint32       `idl:"name:messageId" json:"message_id"`
	Values           *VariantList `idl:"name:values" json:"values"`
	Flags            uint32       `idl:"name:flags" json:"flags"`
	MaxSizeString    uint32       `idl:"name:maxSizeString" json:"max_size_string"`
	ActualSizeString uint32       `idl:"name:actualSizeString" json:"actual_size_string"`
	NeededSizeString uint32       `idl:"name:neededSizeString" json:"needed_size_string"`
	String           []byte       `idl:"name:string;size_is:(, actualSizeString)" json:"string"`
	Error            *Info        `idl:"name:error" json:"error"`
	Return           uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_MessageRenderDefaultOperation) OpNum() int { return 10 }

func (o *xxx_MessageRenderDefaultOperation) OpName() string {
	return "/IEventService/v1/EvtRpcMessageRenderDefault"
}

func (o *xxx_MessageRenderDefaultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.EventID != nil && o.SizeEventID == 0 {
		o.SizeEventID = uint32(len(o.EventID))
	}
	if o.SizeEventID < uint32(1) || o.SizeEventID > uint32(256) {
		return fmt.Errorf("SizeEventID is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderDefaultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// sizeEventId {in} (1:{range=(1,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeEventID); err != nil {
			return err
		}
	}
	// eventId {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=sizeEventId](uchar))
	{
		dimSize1 := uint64(o.SizeEventID)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.EventID {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.EventID[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.EventID); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// messageId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MessageID); err != nil {
			return err
		}
	}
	// values {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Values != nil {
			if err := o.Values.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// maxSizeString {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxSizeString); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderDefaultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// sizeEventId {in} (1:{range=(1,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeEventID); err != nil {
			return err
		}
	}
	// eventId {in} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=sizeEventId](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.EventID", sizeInfo[0])
		}
		o.EventID = make([]byte, sizeInfo[0])
		for i1 := range o.EventID {
			i1 := i1
			if err := w.ReadData(&o.EventID[i1]); err != nil {
				return err
			}
		}
	}
	// messageId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MessageID); err != nil {
			return err
		}
	}
	// values {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Values == nil {
			o.Values = &VariantList{}
		}
		if err := o.Values.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// maxSizeString {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxSizeString); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderDefaultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.String != nil && o.ActualSizeString == 0 {
		o.ActualSizeString = uint32(len(o.String))
	}
	if len(o.String) > int(2097152) {
		return fmt.Errorf("String is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderDefaultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// actualSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualSizeString); err != nil {
			return err
		}
	}
	// neededSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NeededSizeString); err != nil {
			return err
		}
	}
	// string {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=actualSizeString](uchar))
	{
		if o.String != nil || o.ActualSizeString > 0 {
			_ptr_string := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualSizeString)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.String {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.String[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.String); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.String, _ptr_string); err != nil {
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
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MessageRenderDefaultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// actualSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualSizeString); err != nil {
			return err
		}
	}
	// neededSizeString {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NeededSizeString); err != nil {
			return err
		}
	}
	// string {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=actualSizeString](uchar))
	{
		_ptr_string := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.String", sizeInfo[0])
			}
			o.String = make([]byte, sizeInfo[0])
			for i1 := range o.String {
				i1 := i1
				if err := w.ReadData(&o.String[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_string := func(ptr interface{}) { o.String = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.String, _s_string, _ptr_string); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MessageRenderDefaultRequest structure represents the EvtRpcMessageRenderDefault operation request
type MessageRenderDefaultRequest struct {
	// sizeEventId: A 32-bit unsigned integer that contains the size in bytes of the eventId
	// field.
	SizeEventID uint32 `idl:"name:sizeEventId" json:"size_event_id"`
	// eventId: A pointer to an EVENT_DESCRIPTOR structure, as specified in [MS-DTYP] section
	// 2.3.1.
	EventID []byte `idl:"name:eventId;size_is:(sizeEventId)" json:"event_id"`
	// messageId: A 32-bit unsigned integer that specifies the required message. This is
	// an alternative to using the eventID parameter that can be used by a client application
	// that has obtained the value through some method outside those documented by this
	// protocol. The server MUST ignore this value unless the flags value is set to EvtFormatMessageId,
	// in which case the server MUST use this value to determine the required message and
	// ignore the eventID parameter.
	MessageID uint32 `idl:"name:messageId" json:"message_id"`
	// values: An array of strings to be used as substitution values for event description
	// strings. Substitution values MUST be ignored by the server except when the flags
	// are set to either EvtFormatMessageEvent or EvtFormatMessageId.
	Values *VariantList `idl:"name:values" json:"values"`
	// flags: This field MUST be set to a value from the following table, which indicates
	// the action that the server is requested to perform.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageEvent 0x00000001   | Locate the message for the event corresponding to eventId, and then insert the   |
	//	|                                    | values specified by the values parameter.                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageLevel 0x00000002   | Extract the level field from eventId, and then return the localized name for     |
	//	|                                    | that level.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageTask 0x00000003    | Extract the task field from eventId, and then return the localized name for that |
	//	|                                    | task.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageOpcode 0x00000004  | Extract the opcode field from eventId, and then return the localized name for    |
	//	|                                    | that opcode.                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageKeyword 0x00000005 | Extract the keyword field from eventId, and then return the localized name for   |
	//	|                                    | that keyword.                                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtFormatMessageId 0x00000008      | Locate the message for the event corresponding to the messageId parameter, and   |
	//	|                                    | then insert the values specified by the values parameter.                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// maxSizeString: A 32-bit unsigned integer that contains the maximum size in bytes
	// allowed for the string field.
	MaxSizeString uint32 `idl:"name:maxSizeString" json:"max_size_string"`
}

func (o *MessageRenderDefaultRequest) xxx_ToOp(ctx context.Context, op *xxx_MessageRenderDefaultOperation) *xxx_MessageRenderDefaultOperation {
	if op == nil {
		op = &xxx_MessageRenderDefaultOperation{}
	}
	if o == nil {
		return op
	}
	op.SizeEventID = o.SizeEventID
	op.EventID = o.EventID
	op.MessageID = o.MessageID
	op.Values = o.Values
	op.Flags = o.Flags
	op.MaxSizeString = o.MaxSizeString
	return op
}

func (o *MessageRenderDefaultRequest) xxx_FromOp(ctx context.Context, op *xxx_MessageRenderDefaultOperation) {
	if o == nil {
		return
	}
	o.SizeEventID = op.SizeEventID
	o.EventID = op.EventID
	o.MessageID = op.MessageID
	o.Values = op.Values
	o.Flags = op.Flags
	o.MaxSizeString = op.MaxSizeString
}
func (o *MessageRenderDefaultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MessageRenderDefaultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageRenderDefaultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MessageRenderDefaultResponse structure represents the EvtRpcMessageRenderDefault operation response
type MessageRenderDefaultResponse struct {
	// actualSizeString: A pointer to a 32-bit unsigned integer that contains the actual
	// size of the resulting description string returned in the string. It MUST be set to
	// the size in bytes of the string returned in the string parameter, including the NULL
	// ('\0') terminating character. If the description string cannot be retrieved, actualSizeString
	// MUST be set to zero.
	ActualSizeString uint32 `idl:"name:actualSizeString" json:"actual_size_string"`
	// neededSizeString: A pointer to a 32-bit unsigned integer that contains the size in
	// bytes of the fully instantiated description string, even if the length of the description
	// string is greater than maxSizeString. The returned value MUST be zero when the description
	// string cannot be computed by the server.
	NeededSizeString uint32 `idl:"name:neededSizeString" json:"needed_size_string"`
	// string: A buffer in which to return either a null-terminated string or multiple null-terminated
	// strings, terminated by a double NULL in the case of keywords. In the case of failure,
	// the client MUST ignore this value.
	String []byte `idl:"name:string;size_is:(, actualSizeString)" json:"string"`
	// error:  A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to loading the necessary resource.
	// All nonzero values MUST be treated the same. If the method succeeds, the server MUST
	// set all of the values in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcMessageRenderDefault return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MessageRenderDefaultResponse) xxx_ToOp(ctx context.Context, op *xxx_MessageRenderDefaultOperation) *xxx_MessageRenderDefaultOperation {
	if op == nil {
		op = &xxx_MessageRenderDefaultOperation{}
	}
	if o == nil {
		return op
	}
	op.ActualSizeString = o.ActualSizeString
	op.NeededSizeString = o.NeededSizeString
	op.String = o.String
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *MessageRenderDefaultResponse) xxx_FromOp(ctx context.Context, op *xxx_MessageRenderDefaultOperation) {
	if o == nil {
		return
	}
	o.ActualSizeString = op.ActualSizeString
	o.NeededSizeString = op.NeededSizeString
	o.String = op.String
	o.Error = op.Error
	o.Return = op.Return
}
func (o *MessageRenderDefaultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MessageRenderDefaultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MessageRenderDefaultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryNextOperation structure represents the EvtRpcQueryNext operation
type xxx_QueryNextOperation struct {
	LogQuery               *LogQuery `idl:"name:logQuery" json:"log_query"`
	RequestedRecordsLength uint32    `idl:"name:numRequestedRecords" json:"requested_records_length"`
	TimeoutEnd             uint32    `idl:"name:timeOutEnd" json:"timeout_end"`
	Flags                  uint32    `idl:"name:flags" json:"flags"`
	ActualRecordsLength    uint32    `idl:"name:numActualRecords" json:"actual_records_length"`
	EventDataIndices       []uint32  `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	EventDataSizes         []uint32  `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	ResultBufferSize       uint32    `idl:"name:resultBufferSize" json:"result_buffer_size"`
	ResultBuffer           []byte    `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	Return                 uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryNextOperation) OpNum() int { return 11 }

func (o *xxx_QueryNextOperation) OpName() string { return "/IEventService/v1/EvtRpcQueryNext" }

func (o *xxx_QueryNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// logQuery {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.LogQuery != nil {
			if err := o.LogQuery.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LogQuery{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// timeOutEnd {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TimeoutEnd); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// logQuery {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.LogQuery == nil {
			o.LogQuery = &LogQuery{}
		}
		if err := o.LogQuery.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// numRequestedRecords {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedRecordsLength); err != nil {
			return err
		}
	}
	// timeOutEnd {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TimeoutEnd); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.EventDataIndices != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataIndices))
	}
	if o.EventDataSizes != nil && o.ActualRecordsLength == 0 {
		o.ActualRecordsLength = uint32(len(o.EventDataSizes))
	}
	if o.ResultBuffer != nil && o.ResultBufferSize == 0 {
		o.ResultBufferSize = uint32(len(o.ResultBuffer))
	}
	if len(o.EventDataIndices) > int(1024) {
		return fmt.Errorf("EventDataIndices is out of range")
	}
	if len(o.EventDataSizes) > int(1024) {
		return fmt.Errorf("EventDataSizes is out of range")
	}
	if len(o.ResultBuffer) > int(2097152) {
		return fmt.Errorf("ResultBuffer is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataIndices != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataIndices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataIndices {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataIndices[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataIndices); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataIndices, _ptr_eventDataIndices); err != nil {
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
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		if o.EventDataSizes != nil || o.ActualRecordsLength > 0 {
			_ptr_eventDataSizes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ActualRecordsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventDataSizes {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.EventDataSizes[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.EventDataSizes); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventDataSizes, _ptr_eventDataSizes); err != nil {
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
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		if o.ResultBuffer != nil || o.ResultBufferSize > 0 {
			_ptr_resultBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ResultBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ResultBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ResultBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ResultBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResultBuffer, _ptr_resultBuffer); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numActualRecords {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ActualRecordsLength); err != nil {
			return err
		}
	}
	// eventDataIndices {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataIndices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataIndices", sizeInfo[0])
			}
			o.EventDataIndices = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataIndices {
				i1 := i1
				if err := w.ReadData(&o.EventDataIndices[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataIndices := func(ptr interface{}) { o.EventDataIndices = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataIndices, _s_eventDataIndices, _ptr_eventDataIndices); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// eventDataSizes {out} (1:{pointer=ref, range=(0,1024)}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=numActualRecords](uint32))
	{
		_ptr_eventDataSizes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventDataSizes", sizeInfo[0])
			}
			o.EventDataSizes = make([]uint32, sizeInfo[0])
			for i1 := range o.EventDataSizes {
				i1 := i1
				if err := w.ReadData(&o.EventDataSizes[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventDataSizes := func(ptr interface{}) { o.EventDataSizes = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.EventDataSizes, _s_eventDataSizes, _ptr_eventDataSizes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// resultBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ResultBufferSize); err != nil {
			return err
		}
	}
	// resultBuffer {out} (1:{pointer=ref, range=(0,2097152)}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=resultBufferSize](uchar))
	{
		_ptr_resultBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ResultBuffer", sizeInfo[0])
			}
			o.ResultBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.ResultBuffer {
				i1 := i1
				if err := w.ReadData(&o.ResultBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_resultBuffer := func(ptr interface{}) { o.ResultBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ResultBuffer, _s_resultBuffer, _ptr_resultBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryNextRequest structure represents the EvtRpcQueryNext operation request
type QueryNextRequest struct {
	// logQuery: A handle to an event log. This parameter is an RPC context handle, as specified
	// in [C706], Context Handles.
	LogQuery *LogQuery `idl:"name:logQuery" json:"log_query"`
	// numRequestedRecords: A 32-bit unsigned integer that contains the number of events
	// to return.<19>
	RequestedRecordsLength uint32 `idl:"name:numRequestedRecords" json:"requested_records_length"`
	// timeOutEnd: A 32-bit unsigned integer that contains the maximum number of milliseconds
	// to wait before returning, starting from the time the server begins processing the
	// call.
	TimeoutEnd uint32 `idl:"name:timeOutEnd" json:"timeout_end"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<20>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *QueryNextRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryNextOperation) *xxx_QueryNextOperation {
	if op == nil {
		op = &xxx_QueryNextOperation{}
	}
	if o == nil {
		return op
	}
	op.LogQuery = o.LogQuery
	op.RequestedRecordsLength = o.RequestedRecordsLength
	op.TimeoutEnd = o.TimeoutEnd
	op.Flags = o.Flags
	return op
}

func (o *QueryNextRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryNextOperation) {
	if o == nil {
		return
	}
	o.LogQuery = op.LogQuery
	o.RequestedRecordsLength = op.RequestedRecordsLength
	o.TimeoutEnd = op.TimeoutEnd
	o.Flags = op.Flags
}
func (o *QueryNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryNextResponse structure represents the EvtRpcQueryNext operation response
type QueryNextResponse struct {
	// numActualRecords: A pointer to a 32-bit unsigned integer that contains the value
	// that, on success, MUST be set to the number of events that are retrieved. This is
	// useful when the method times out without receiving the full number of events specified
	// in numRequestedRecords. If the method fails, the client MUST NOT use the value.
	ActualRecordsLength uint32 `idl:"name:numActualRecords" json:"actual_records_length"`
	// eventDataIndices: A pointer to an array of 32-bit unsigned integers that contain
	// the offsets (in bytes) within the resultBuffer for the events that are read.
	EventDataIndices []uint32 `idl:"name:eventDataIndices;size_is:(, numActualRecords)" json:"event_data_indices"`
	// eventDataSizes: A pointer to an array of 32-bit unsigned integers that contain the
	// sizes (in bytes) within the resultBuffer for the events that are read.
	EventDataSizes []uint32 `idl:"name:eventDataSizes;size_is:(, numActualRecords)" json:"event_data_sizes"`
	// resultBufferSize: A pointer to a 32-bit unsigned integer that contains the number
	// of bytes of data returned in resultBuffer.
	ResultBufferSize uint32 `idl:"name:resultBufferSize" json:"result_buffer_size"`
	// resultBuffer: A pointer to a byte-array that contains the result set of one or more
	// events. The events MUST be in binary XML format, as specified in section 2.2.17.
	ResultBuffer []byte `idl:"name:resultBuffer;size_is:(, resultBufferSize)" json:"result_buffer"`
	// Return: The EvtRpcQueryNext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryNextResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryNextOperation) *xxx_QueryNextOperation {
	if op == nil {
		op = &xxx_QueryNextOperation{}
	}
	if o == nil {
		return op
	}
	op.ActualRecordsLength = o.ActualRecordsLength
	op.EventDataIndices = o.EventDataIndices
	op.EventDataSizes = o.EventDataSizes
	op.ResultBufferSize = o.ResultBufferSize
	op.ResultBuffer = o.ResultBuffer
	op.Return = o.Return
	return op
}

func (o *QueryNextResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryNextOperation) {
	if o == nil {
		return
	}
	o.ActualRecordsLength = op.ActualRecordsLength
	o.EventDataIndices = op.EventDataIndices
	o.EventDataSizes = op.EventDataSizes
	o.ResultBufferSize = op.ResultBufferSize
	o.ResultBuffer = op.ResultBuffer
	o.Return = op.Return
}
func (o *QueryNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySeekOperation structure represents the EvtRpcQuerySeek operation
type xxx_QuerySeekOperation struct {
	LogQuery    *LogQuery `idl:"name:logQuery" json:"log_query"`
	Pos         int64     `idl:"name:pos" json:"pos"`
	BookmarkXML string    `idl:"name:bookmarkXml;string;pointer:unique" json:"bookmark_xml"`
	Timeout     uint32    `idl:"name:timeOut" json:"timeout"`
	Flags       uint32    `idl:"name:flags" json:"flags"`
	Error       *Info     `idl:"name:error" json:"error"`
	Return      uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySeekOperation) OpNum() int { return 12 }

func (o *xxx_QuerySeekOperation) OpName() string { return "/IEventService/v1/EvtRpcQuerySeek" }

func (o *xxx_QuerySeekOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.BookmarkXML) > int(1048576) {
		return fmt.Errorf("BookmarkXML is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySeekOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// logQuery {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.LogQuery != nil {
			if err := o.LogQuery.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&LogQuery{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pos {in} (1:(int64))
	{
		if err := w.WriteData(o.Pos); err != nil {
			return err
		}
	}
	// bookmarkXml {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.BookmarkXML != "" {
			_ptr_bookmarkXml := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.BookmarkXML); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.BookmarkXML, _ptr_bookmarkXml); err != nil {
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
	// timeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySeekOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// logQuery {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_QUERY, names=ndr_context_handle}(struct))
	{
		if o.LogQuery == nil {
			o.LogQuery = &LogQuery{}
		}
		if err := o.LogQuery.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pos {in} (1:(int64))
	{
		if err := w.ReadData(&o.Pos); err != nil {
			return err
		}
	}
	// bookmarkXml {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_bookmarkXml := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.BookmarkXML); err != nil {
				return err
			}
			return nil
		})
		_s_bookmarkXml := func(ptr interface{}) { o.BookmarkXML = *ptr.(*string) }
		if err := w.ReadPointer(&o.BookmarkXML, _s_bookmarkXml, _ptr_bookmarkXml); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// timeOut {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySeekOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySeekOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySeekOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySeekRequest structure represents the EvtRpcQuerySeek operation request
type QuerySeekRequest struct {
	// logQuery: A handle to an event log. This parameter is an RPC context handle, as specified
	// in [C706], Context Handles.
	LogQuery *LogQuery `idl:"name:logQuery" json:"log_query"`
	// pos: The number of records in the result set to move by. If the number is positive,
	// the movement is the same as the direction of the query that was specified in the
	// EvtRpcRegisterLogQuery (section 3.1.4.12) method call that was used to obtain the
	// handle specified by the logQuery parameter. If the number is negative, the movement
	// is in the opposite direction of the query.
	Pos int64 `idl:"name:pos" json:"pos"`
	// bookmarkXml: A pointer to a string that contains a bookmark.
	BookmarkXML string `idl:"name:bookmarkXml;string;pointer:unique" json:"bookmark_xml"`
	// timeOut: A 32-bit unsigned integer that MUST be set to 0x00000000 when sent and MAY
	// be ignored on receipt.
	Timeout uint32 `idl:"name:timeOut" json:"timeout"`
	// flags: This MUST be set as follows: this 32-bit unsigned integer contains flags that
	// describe the absolute position from which EvtRpcQuerySeek (Opnum 12) starts its seek.
	// The origin flags (the first four flags that follow) are mutually exclusive; however,
	// the last flag can be set independently. The pos parameter specifies the offset used
	// in the definitions of these flags.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSeekRelativeToFirst 0x00000001    | The offset is relative to the first entry in the result set and SHOULD be        |
	//	|                                      | nonnegative. Therefore, if an offset of 0 is specified, the cursor is moved to   |
	//	|                                      | the first entry in the result set.                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSeekRelativeToLast 0x00000002     | The offset is relative to the last entry in the result set and SHOULD be         |
	//	|                                      | nonpositive. Therefore, if an offset of 0 is specified, the cursor is moved to   |
	//	|                                      | the last entry in the result set.                                                |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSeekRelativeToCurrent 0x00000003  | The offset is relative to the current cursor location. If an offset of 0 is      |
	//	|                                      | specified, the cursor is not to be moved. A positive or negative number can be   |
	//	|                                      | used in this case to move the cursor to any other location.                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSeekRelativeToBookmark 0x00000004 | The offset is relative to the bookmark location. If an offset of 0 is specified, |
	//	|                                      | the cursor is positioned at the bookmark. A positive or negative number can be   |
	//	|                                      | used in this case to move the cursor to any other location. The server MUST fail |
	//	|                                      | the operation if the bookmarkXml parameter does not specify a valid position in  |
	//	|                                      | the log. The presence of the EvtSeekStrict flag SHOULD influence the behavior of |
	//	|                                      | this flag, as specified below.                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtSeekStrict 0x00010000             | If this is set, the query fails if the seek cannot go to the record indicated    |
	//	|                                      | by the other flags/parameters. If not set, the seek uses a best effort. For      |
	//	|                                      | example, if 99 records remain in the result set and the pos parameter specifies  |
	//	|                                      | 100 with the EvtSeekRelativeToCurrent flag set, the 99th record is selected.     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *QuerySeekRequest) xxx_ToOp(ctx context.Context, op *xxx_QuerySeekOperation) *xxx_QuerySeekOperation {
	if op == nil {
		op = &xxx_QuerySeekOperation{}
	}
	if o == nil {
		return op
	}
	op.LogQuery = o.LogQuery
	op.Pos = o.Pos
	op.BookmarkXML = o.BookmarkXML
	op.Timeout = o.Timeout
	op.Flags = o.Flags
	return op
}

func (o *QuerySeekRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySeekOperation) {
	if o == nil {
		return
	}
	o.LogQuery = op.LogQuery
	o.Pos = op.Pos
	o.BookmarkXML = op.BookmarkXML
	o.Timeout = op.Timeout
	o.Flags = op.Flags
}
func (o *QuerySeekRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QuerySeekRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySeekOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySeekResponse structure represents the EvtRpcQuerySeek operation response
type QuerySeekResponse struct {
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo structure fields MUST be set to
	// nonzero values if the error is related to parsing the query. In addition, the server
	// MAY set the structure fields to nonzero values for errors unrelated to query parsing.
	// All nonzero values MUST be treated the same by the client.
	//
	// If the method succeeds, the server MUST set all the values in the structure to zero.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcQuerySeek return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QuerySeekResponse) xxx_ToOp(ctx context.Context, op *xxx_QuerySeekOperation) *xxx_QuerySeekOperation {
	if op == nil {
		op = &xxx_QuerySeekOperation{}
	}
	if o == nil {
		return op
	}
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *QuerySeekResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySeekOperation) {
	if o == nil {
		return
	}
	o.Error = op.Error
	o.Return = op.Return
}
func (o *QuerySeekResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QuerySeekResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySeekOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseOperation structure represents the EvtRpcClose operation
type xxx_CloseOperation struct {
	Handle *dcetypes.ContextHandle `idl:"name:handle" json:"handle"`
	Return uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 13 }

func (o *xxx_CloseOperation) OpName() string { return "/IEventService/v1/EvtRpcClose" }

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
	// handle {in, out} (1:{context_handle, pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in, out} (1:{context_handle, pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dcetypes.ContextHandle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
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
	// handle {in, out} (1:{context_handle, pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// handle {in, out} (1:{context_handle, pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &dcetypes.ContextHandle{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseRequest structure represents the EvtRpcClose operation request
type CloseRequest struct {
	// handle: This parameter is an RPC context handle, as specified in [C706], Context
	// Handles.
	Handle *dcetypes.ContextHandle `idl:"name:handle" json:"handle"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	return op
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
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

// CloseResponse structure represents the EvtRpcClose operation response
type CloseResponse struct {
	// handle: This parameter is an RPC context handle, as specified in [C706], Context
	// Handles.
	Handle *dcetypes.ContextHandle `idl:"name:handle" json:"handle"`
	// Return: The EvtRpcClose return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
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

// xxx_CancelOperation structure represents the EvtRpcCancel operation
type xxx_CancelOperation struct {
	Handle *OperationControl `idl:"name:handle" json:"handle"`
	Return uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelOperation) OpNum() int { return 14 }

func (o *xxx_CancelOperation) OpName() string { return "/IEventService/v1/EvtRpcCancel" }

func (o *xxx_CancelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&OperationControl{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CancelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_OPERATION_CONTROL, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &OperationControl{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CancelRequest structure represents the EvtRpcCancel operation request
type CancelRequest struct {
	// handle: A handle obtained by any of the other methods in this interface. This parameter
	// is an RPC context handle, as specified in [C706], Context Handles.
	Handle *OperationControl `idl:"name:handle" json:"handle"`
}

func (o *CancelRequest) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	return op
}

func (o *CancelRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
}
func (o *CancelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CancelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelResponse structure represents the EvtRpcCancel operation response
type CancelResponse struct {
	// Return: The EvtRpcCancel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CancelResponse) xxx_ToOp(ctx context.Context, op *xxx_CancelOperation) *xxx_CancelOperation {
	if op == nil {
		op = &xxx_CancelOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *CancelResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CancelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CancelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AssertConfigOperation structure represents the EvtRpcAssertConfig operation
type xxx_AssertConfigOperation struct {
	Path   string `idl:"name:path;string" json:"path"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AssertConfigOperation) OpNum() int { return 15 }

func (o *xxx_AssertConfigOperation) OpName() string { return "/IEventService/v1/EvtRpcAssertConfig" }

func (o *xxx_AssertConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Path) > int(512) {
		return fmt.Errorf("Path is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssertConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssertConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssertConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssertConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssertConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AssertConfigRequest structure represents the EvtRpcAssertConfig operation request
type AssertConfigRequest struct {
	// path: A pointer to a string that contains a channel or publisher name to be updated.
	Path string `idl:"name:path;string" json:"path"`
	// flags: The client MUST specify exactly one of the following.
	//
	//	+--------------------------------+----------------------------------+
	//	|                                |                                  |
	//	|             VALUE              |             MEANING              |
	//	|                                |                                  |
	//	+--------------------------------+----------------------------------+
	//	+--------------------------------+----------------------------------+
	//	| EvtRpcChannelPath 0x00000000   | Path specifies a channel name.   |
	//	+--------------------------------+----------------------------------+
	//	| EvtRpcPublisherName 0x00000001 | Path specifies a publisher name. |
	//	+--------------------------------+----------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *AssertConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_AssertConfigOperation) *xxx_AssertConfigOperation {
	if op == nil {
		op = &xxx_AssertConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Path = o.Path
	op.Flags = o.Flags
	return op
}

func (o *AssertConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_AssertConfigOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *AssertConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AssertConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssertConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AssertConfigResponse structure represents the EvtRpcAssertConfig operation response
type AssertConfigResponse struct {
	// Return: The EvtRpcAssertConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AssertConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_AssertConfigOperation) *xxx_AssertConfigOperation {
	if op == nil {
		op = &xxx_AssertConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AssertConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_AssertConfigOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AssertConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AssertConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssertConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetractConfigOperation structure represents the EvtRpcRetractConfig operation
type xxx_RetractConfigOperation struct {
	Path   string `idl:"name:path;string" json:"path"`
	Flags  uint32 `idl:"name:flags" json:"flags"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RetractConfigOperation) OpNum() int { return 16 }

func (o *xxx_RetractConfigOperation) OpName() string { return "/IEventService/v1/EvtRpcRetractConfig" }

func (o *xxx_RetractConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Path) > int(512) {
		return fmt.Errorf("Path is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetractConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// path {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetractConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// path {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetractConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetractConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetractConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetractConfigRequest structure represents the EvtRpcRetractConfig operation request
type RetractConfigRequest struct {
	// path: A pointer to a string that contains a channel or publisher name to be removed.
	Path string `idl:"name:path;string" json:"path"`
	// flags: A 32-bit unsigned integer that indicates how the path parameter is to be interpreted.
	// This MUST be set as follows.
	//
	//	+--------------------------------+----------------------------------+
	//	|                                |                                  |
	//	|             VALUE              |             MEANING              |
	//	|                                |                                  |
	//	+--------------------------------+----------------------------------+
	//	+--------------------------------+----------------------------------+
	//	| EvtRpcChannelPath 0x00000000   | Path specifies a channel name.   |
	//	+--------------------------------+----------------------------------+
	//	| EvtRpcPublisherName 0x00000001 | Path specifies a publisher name. |
	//	+--------------------------------+----------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *RetractConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_RetractConfigOperation) *xxx_RetractConfigOperation {
	if op == nil {
		op = &xxx_RetractConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Path = o.Path
	op.Flags = o.Flags
	return op
}

func (o *RetractConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_RetractConfigOperation) {
	if o == nil {
		return
	}
	o.Path = op.Path
	o.Flags = op.Flags
}
func (o *RetractConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RetractConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetractConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetractConfigResponse structure represents the EvtRpcRetractConfig operation response
type RetractConfigResponse struct {
	// Return: The EvtRpcRetractConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RetractConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_RetractConfigOperation) *xxx_RetractConfigOperation {
	if op == nil {
		op = &xxx_RetractConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RetractConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_RetractConfigOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RetractConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RetractConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetractConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenLogOperation structure represents the EvtRpcOpenLogHandle operation
type xxx_OpenLogOperation struct {
	Channel string `idl:"name:channel;string" json:"channel"`
	Flags   uint32 `idl:"name:flags" json:"flags"`
	Handle  *Log   `idl:"name:handle" json:"handle"`
	Error   *Info  `idl:"name:error" json:"error"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenLogOperation) OpNum() int { return 17 }

func (o *xxx_OpenLogOperation) OpName() string { return "/IEventService/v1/EvtRpcOpenLogHandle" }

func (o *xxx_OpenLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.Channel) > int(512) {
		return fmt.Errorf("Channel is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// channel {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Channel); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// channel {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Channel); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_LOG_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Log{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// handle {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_LOG_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &Log{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenLogRequest structure represents the EvtRpcOpenLogHandle operation request
type OpenLogRequest struct {
	// channel: A pointer to a string that contains a channel or a file path.
	Channel string `idl:"name:channel;string" json:"channel"`
	// flags: MUST be one of the following two values.
	//
	//	+------------+---------------------------------------------+
	//	|            |                                             |
	//	|   VALUE    |                   MEANING                   |
	//	|            |                                             |
	//	+------------+---------------------------------------------+
	//	+------------+---------------------------------------------+
	//	| 0x00000001 | Channel parameter specifies a channel name. |
	//	+------------+---------------------------------------------+
	//	| 0x00000002 | Channel parameter specifies a file name.    |
	//	+------------+---------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *OpenLogRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenLogOperation) *xxx_OpenLogOperation {
	if op == nil {
		op = &xxx_OpenLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Channel = o.Channel
	op.Flags = o.Flags
	return op
}

func (o *OpenLogRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenLogOperation) {
	if o == nil {
		return
	}
	o.Channel = op.Channel
	o.Flags = op.Flags
}
func (o *OpenLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenLogResponse structure represents the EvtRpcOpenLogHandle operation response
type OpenLogResponse struct {
	// handle: A pointer to a log handle. This parameter is an RPC context handle, as specified
	// in [C706], Context Handles.
	Handle *Log `idl:"name:handle" json:"handle"`
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The server MAY set the suberror fields to supply
	// more comprehensive error information.<35> If the method succeeds, the server MUST
	// set all of the values in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcOpenLogHandle return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenLogResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenLogOperation) *xxx_OpenLogOperation {
	if op == nil {
		op = &xxx_OpenLogOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *OpenLogResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenLogOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Error = op.Error
	o.Return = op.Return
}
func (o *OpenLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLogFileInfoOperation structure represents the EvtRpcGetLogFileInfo operation
type xxx_GetLogFileInfoOperation struct {
	Log                       *Log   `idl:"name:logHandle" json:"log"`
	PropertyID                uint32 `idl:"name:propertyId" json:"property_id"`
	PropertyValueBufferSize   uint32 `idl:"name:propertyValueBufferSize" json:"property_value_buffer_size"`
	PropertyValueBuffer       []byte `idl:"name:propertyValueBuffer;size_is:(propertyValueBufferSize)" json:"property_value_buffer"`
	PropertyValueBufferLength uint32 `idl:"name:propertyValueBufferLength" json:"property_value_buffer_length"`
	Return                    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLogFileInfoOperation) OpNum() int { return 18 }

func (o *xxx_GetLogFileInfoOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetLogFileInfo"
}

func (o *xxx_GetLogFileInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PropertyValueBufferSize > uint32(2097152) {
		return fmt.Errorf("PropertyValueBufferSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// logHandle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log != nil {
			if err := o.Log.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Log{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// propertyId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyID); err != nil {
			return err
		}
	}
	// propertyValueBufferSize {in} (1:{range=(0,2097152), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyValueBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// logHandle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_LOG_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Log == nil {
			o.Log = &Log{}
		}
		if err := o.Log.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// propertyId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyID); err != nil {
			return err
		}
	}
	// propertyValueBufferSize {in} (1:{range=(0,2097152), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyValueBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// propertyValueBuffer {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=propertyValueBufferSize](uchar))
	{
		dimSize1 := uint64(o.PropertyValueBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.PropertyValueBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.PropertyValueBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.PropertyValueBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// propertyValueBufferLength {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyValueBufferLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLogFileInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// propertyValueBuffer {out} (1:{pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=propertyValueBufferSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.PropertyValueBuffer", sizeInfo[0])
		}
		o.PropertyValueBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.PropertyValueBuffer {
			i1 := i1
			if err := w.ReadData(&o.PropertyValueBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// propertyValueBufferLength {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyValueBufferLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetLogFileInfoRequest structure represents the EvtRpcGetLogFileInfo operation request
type GetLogFileInfoRequest struct {
	// logHandle: A handle to an event log. This parameter is an RPC context handle, as
	// specified in [C706], Context Handles. For more information about the server-side
	// object that maps to this handle, see section 3.1.4.19.
	Log *Log `idl:"name:logHandle" json:"log"`
	// propertyId: A 32-bit unsigned integer that indicates what log file property (as specified
	// in section 3.1.1.6) needs to be retrieved.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogCreationTime 0x00000000       | A FILETIME containing the creation time of the file. This is the creation time   |
	//	|                                     | of a log file associated with the channel or the creation time of the backup     |
	//	|                                     | event log file in the server's file system.                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogLastAccessTime 0x00000001     | A FILETIME containing the last access time of the file. This is the last access  |
	//	|                                     | time of a log file associated with the channel or the last access time of the    |
	//	|                                     | backup event log file in the server's file system.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogLastWriteTime 0x00000002      | A FILETIME containing the last write time of the file. This is the last written  |
	//	|                                     | time of a log file associated with the channel or the last written time of the   |
	//	|                                     | backup event log file in the server's file system.                               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogFileSize 0x00000003           | An unsigned 64-bit integer containing the size of the file. This is the file     |
	//	|                                     | size of a log file associated with the channel or the file size of the backup    |
	//	|                                     | event log file in the server's file system.                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogAttributes 0x00000004         | An unsigned 32-bit integer containing the attributes of the file. The attributes |
	//	|                                     | are implementation-specific, and clients MUST<23> treat all values equally.      |
	//	|                                     | The attributes are tracked by the server's file system and SHOULD be able to be  |
	//	|                                     | retrieved from the file system.                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogNumberOfLogRecords 0x00000005 | An unsigned 64-bit integer containing the number of records in the file. See the |
	//	|                                     | following processing rules for how the server gets this value.                   |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogOldestRecordNumber 0x00000006 | An unsigned 64-bit integer containing the oldest record number in the file. See  |
	//	|                                     | the following processing rules for how the server gets this value.               |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| EvtLogFull 0x00000007               | A BOOLEAN value; MUST be true if the log is full, and MUST be false otherwise.   |
	//	|                                     | See the following processing rules for how the server gets this value.           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	PropertyID uint32 `idl:"name:propertyId" json:"property_id"`
	// propertyValueBufferSize: A 32-bit unsigned integer that contains the length of caller's
	// buffer in bytes.
	PropertyValueBufferSize uint32 `idl:"name:propertyValueBufferSize" json:"property_value_buffer_size"`
}

func (o *GetLogFileInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLogFileInfoOperation) *xxx_GetLogFileInfoOperation {
	if op == nil {
		op = &xxx_GetLogFileInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Log = o.Log
	op.PropertyID = o.PropertyID
	op.PropertyValueBufferSize = o.PropertyValueBufferSize
	return op
}

func (o *GetLogFileInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLogFileInfoOperation) {
	if o == nil {
		return
	}
	o.Log = op.Log
	o.PropertyID = op.PropertyID
	o.PropertyValueBufferSize = op.PropertyValueBufferSize
}
func (o *GetLogFileInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLogFileInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFileInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLogFileInfoResponse structure represents the EvtRpcGetLogFileInfo operation response
type GetLogFileInfoResponse struct {
	// XXX: propertyValueBufferSize is an implicit input depedency for output parameters
	PropertyValueBufferSize uint32 `idl:"name:propertyValueBufferSize" json:"property_value_buffer_size"`

	// propertyValueBuffer: A byte-array that contains the buffer for returned data.
	PropertyValueBuffer []byte `idl:"name:propertyValueBuffer;size_is:(propertyValueBufferSize)" json:"property_value_buffer"`
	// propertyValueBufferLength: A pointer to a 32-bit unsigned integer that contains the
	// size in bytes of the returned data.
	PropertyValueBufferLength uint32 `idl:"name:propertyValueBufferLength" json:"property_value_buffer_length"`
	// Return: The EvtRpcGetLogFileInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetLogFileInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLogFileInfoOperation) *xxx_GetLogFileInfoOperation {
	if op == nil {
		op = &xxx_GetLogFileInfoOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.PropertyValueBufferSize == uint32(0) {
		op.PropertyValueBufferSize = o.PropertyValueBufferSize
	}

	op.PropertyValueBuffer = o.PropertyValueBuffer
	op.PropertyValueBufferLength = o.PropertyValueBufferLength
	op.Return = o.Return
	return op
}

func (o *GetLogFileInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLogFileInfoOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.PropertyValueBufferSize = op.PropertyValueBufferSize

	o.PropertyValueBuffer = op.PropertyValueBuffer
	o.PropertyValueBufferLength = op.PropertyValueBufferLength
	o.Return = op.Return
}
func (o *GetLogFileInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLogFileInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLogFileInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetChannelListOperation structure represents the EvtRpcGetChannelList operation
type xxx_GetChannelListOperation struct {
	Flags              uint32   `idl:"name:flags" json:"flags"`
	ChannelPathsLength uint32   `idl:"name:numChannelPaths" json:"channel_paths_length"`
	ChannelPaths       []string `idl:"name:channelPaths;size_is:(, numChannelPaths);string" json:"channel_paths"`
	Return             uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetChannelListOperation) OpNum() int { return 19 }

func (o *xxx_GetChannelListOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetChannelList"
}

func (o *xxx_GetChannelListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ChannelPaths != nil && o.ChannelPathsLength == 0 {
		o.ChannelPathsLength = uint32(len(o.ChannelPaths))
	}
	if len(o.ChannelPaths) > int(8192) {
		return fmt.Errorf("ChannelPaths is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numChannelPaths {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ChannelPathsLength); err != nil {
			return err
		}
	}
	// channelPaths {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numChannelPaths]*(1)[dim:0,string,null](wchar))
	{
		if o.ChannelPaths != nil || o.ChannelPathsLength > 0 {
			_ptr_channelPaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ChannelPathsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ChannelPaths {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ChannelPaths[i1] != "" {
						_ptr_channelPaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPaths[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ChannelPaths[i1], _ptr_channelPaths); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ChannelPaths); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ChannelPaths, _ptr_channelPaths); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numChannelPaths {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ChannelPathsLength); err != nil {
			return err
		}
	}
	// channelPaths {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numChannelPaths]*(1)[dim:0,string,null](wchar))
	{
		_ptr_channelPaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ChannelPaths", sizeInfo[0])
			}
			o.ChannelPaths = make([]string, sizeInfo[0])
			for i1 := range o.ChannelPaths {
				i1 := i1
				_ptr_channelPaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPaths[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_channelPaths := func(ptr interface{}) { o.ChannelPaths[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ChannelPaths[i1], _s_channelPaths, _ptr_channelPaths); err != nil {
					return err
				}
			}
			return nil
		})
		_s_channelPaths := func(ptr interface{}) { o.ChannelPaths = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ChannelPaths, _s_channelPaths, _ptr_channelPaths); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetChannelListRequest structure represents the EvtRpcGetChannelList operation request
type GetChannelListRequest struct {
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<37>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetChannelListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetChannelListOperation) *xxx_GetChannelListOperation {
	if op == nil {
		op = &xxx_GetChannelListOperation{}
	}
	if o == nil {
		return op
	}
	op.Flags = o.Flags
	return op
}

func (o *GetChannelListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetChannelListOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
}
func (o *GetChannelListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetChannelListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChannelListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetChannelListResponse structure represents the EvtRpcGetChannelList operation response
type GetChannelListResponse struct {
	// numChannelPaths:  A pointer to a 32-bit unsigned integer that contains the number
	// of channel names.
	ChannelPathsLength uint32 `idl:"name:numChannelPaths" json:"channel_paths_length"`
	// channelPaths: A pointer to an array of strings that contain all the channel names.
	ChannelPaths []string `idl:"name:channelPaths;size_is:(, numChannelPaths);string" json:"channel_paths"`
	// Return: The EvtRpcGetChannelList return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetChannelListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetChannelListOperation) *xxx_GetChannelListOperation {
	if op == nil {
		op = &xxx_GetChannelListOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelPathsLength = o.ChannelPathsLength
	op.ChannelPaths = o.ChannelPaths
	op.Return = o.Return
	return op
}

func (o *GetChannelListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetChannelListOperation) {
	if o == nil {
		return
	}
	o.ChannelPathsLength = op.ChannelPathsLength
	o.ChannelPaths = op.ChannelPaths
	o.Return = op.Return
}
func (o *GetChannelListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetChannelListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChannelListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetChannelConfigOperation structure represents the EvtRpcGetChannelConfig operation
type xxx_GetChannelConfigOperation struct {
	ChannelPath string       `idl:"name:channelPath;string" json:"channel_path"`
	Flags       uint32       `idl:"name:flags" json:"flags"`
	Properties  *VariantList `idl:"name:props" json:"properties"`
	Return      uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_GetChannelConfigOperation) OpNum() int { return 20 }

func (o *xxx_GetChannelConfigOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetChannelConfig"
}

func (o *xxx_GetChannelConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ChannelPath) > int(512) {
		return fmt.Errorf("ChannelPath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// channelPath {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// channelPath {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// props {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Properties != nil {
			if err := o.Properties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetChannelConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// props {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Properties == nil {
			o.Properties = &VariantList{}
		}
		if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetChannelConfigRequest structure represents the EvtRpcGetChannelConfig operation request
type GetChannelConfigRequest struct {
	// channelPath: A pointer to a string that contains the name of a channel for which
	// the information is needed.
	ChannelPath string `idl:"name:channelPath;string" json:"channel_path"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<38>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetChannelConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_GetChannelConfigOperation) *xxx_GetChannelConfigOperation {
	if op == nil {
		op = &xxx_GetChannelConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelPath = o.ChannelPath
	op.Flags = o.Flags
	return op
}

func (o *GetChannelConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_GetChannelConfigOperation) {
	if o == nil {
		return
	}
	o.ChannelPath = op.ChannelPath
	o.Flags = op.Flags
}
func (o *GetChannelConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetChannelConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChannelConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetChannelConfigResponse structure represents the EvtRpcGetChannelConfig operation response
type GetChannelConfigResponse struct {
	// props: A pointer to an EvtRpcVariantList structure to be filled with channel properties,
	// as defined in the following table.
	//
	// Note  The index column in the following table is the array index, not the actual
	// field of the EvtRpcVariantList structure. The returned data is an array of EvtRpcVariantList
	// for which the index value is used to identify the elements in the array. For example,
	// index 0 means the first element of the returned array.
	//
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|       |                          |                                                                                  |
	//	| INDEX |           TYPE           |                                     MEANING                                      |
	//	|       |                          |                                                                                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     0 | EvtRpcVarTypeBoolean     | Enabled. If true, the channel can accept new events. If false, any attempts to   |
	//	|       |                          | write events into this channel are automatically dropped.                        |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     1 | EvtRpcVarTypeUInt32      | Channel Isolation. It defines the default access permissions for the channel.    |
	//	|       |                          | Three values are allowed: 0: Application. 1: System. 2: Custom. The default      |
	//	|       |                          | isolation is Application. The default permissions for Application are (shown     |
	//	|       |                          | using SDDL): L"O:BAG:SYD:" L"(A;;0xf0007;;;SY)" // local system (read,           |
	//	|       |                          | write, clear) L"(A;;0x7;;;BA)" // built-in admins (read, write, clear)           |
	//	|       |                          | L"(A;;0x7;;;SO)" // server operators (read, write, clear) L"(A;;0x3;;;IU)"       |
	//	|       |                          | // INTERACTIVE LOGON (read, write) L"(A;;0x3;;;SU)" // SERVICES LOGON (read,     |
	//	|       |                          | write) L"(A;;0x3;;;S-1-5-3)" // BATCH LOGON (read, write) L"(A;;0x3;;;S-1-5-33)" |
	//	|       |                          | // write restricted service (read,write) L"(A;;0x1;;;S-1-5-32-573)"; //          |
	//	|       |                          | event log readers (read) The default permissions for System are (shown           |
	//	|       |                          | using SDDL): L"O:BAG:SYD:" L"(A;;0xf0007;;;SY)" // local system (read,           |
	//	|       |                          | write, clear) L"(A;;0x7;;;BA)" // built-in admins (read, write, clear)           |
	//	|       |                          | L"(A;;0x3;;;BO)" // backup operators (read, write) L"(A;;0x5;;;SO)" //           |
	//	|       |                          | server operators (read, clear) L"(A;;0x1;;;IU)" // INTERACTIVE LOGON (read)      |
	//	|       |                          | L"(A;;0x3;;;SU)" // SERVICES LOGON (read, write) L"(A;;0x1;;;S-1-5-3)" //        |
	//	|       |                          | BATCH LOGON (read) L"(A;;0x2;;;S-1-5-33)" // write restricted service (write)    |
	//	|       |                          | L"(A;;0x1;;;S-1-5-32-573)"; // event log readers (read) When the Custom value is |
	//	|       |                          | used, the Access property will contain the defined SDDL.                         |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     2 | EvtRpcVarTypeUInt32      | Channel type. One of four values: 0: Admin 1: Operational 2: Analytic 3: Debug   |
	//	|       |                          | For more information, see [MSDN-EVTLGCHWINEVTLG].                                |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     3 | EvtRpcVarTypeString      | OwningPublisher. Name of the publisher that defines and registers the channel    |
	//	|       |                          | with the system. For more information on how the server reacts to changes of     |
	//	|       |                          | this property, see section 3.1.4.22.                                             |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     4 | EvtRpcVarTypeBoolean     | ClassicEventlog. If true, the channel represents an event log created according  |
	//	|       |                          | to the EventLog Remoting Protocol, not this protocol (EventLog Remoting Protocol |
	//	|       |                          | Version 6.0). The server maintains two channel tables: one for the EventLog      |
	//	|       |                          | Remoting Protocol Version 6.0 and one for the legacy EventLog Remoting Protocol. |
	//	|       |                          | The table for the legacy EventLog Remoting Protocol is called "log table". For   |
	//	|       |                          | more information on the legacy "log table", see [MS-EVEN]. Any channel coming    |
	//	|       |                          | from the new "channel table" gets the value as false, any channel name that is   |
	//	|       |                          | in the legacy "log table" gets the value as true.                                |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     5 | EvtRpcVarTypeString      | Access. A Security Descriptor Description Language (SDDL) string, as specified   |
	//	|       |                          | in [MS-DTYP], that represents access permissions to the channels. If the         |
	//	|       |                          | isolation attribute is set to Application or System, the access descriptor       |
	//	|       |                          | controls read access to the file (the write permissions are ignored). If the     |
	//	|       |                          | isolation attribute is set to Custom, the access descriptor controls write       |
	//	|       |                          | access to the channel and read access to the file.                               |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     6 | EvtRpcVarTypeBoolean     | Retention. If set to true, events can never be overwritten unless explicitly     |
	//	|       |                          | cleared. If set to false, events are overwritten as needed when the event log is |
	//	|       |                          | full.                                                                            |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     7 | EvtRpcVarTypeBoolean     | AutoBackup. When set to true, the event log file associated with the channel is  |
	//	|       |                          | closed as soon as it reaches the maximum size specified by the MaxSize property, |
	//	|       |                          | and a new file is opened to accept new events. If the new file reaches maximum   |
	//	|       |                          | size, another new file will be generated and the previous new file will be       |
	//	|       |                          | backed up. The events in backed up files cannot be queried from this channel in  |
	//	|       |                          | the server unless the client specifies the backup log file names in a separate   |
	//	|       |                          | query.                                                                           |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     8 | EvtRpcVarTypeUInt64      | MaxSize. The value that indicates at which point the size (in bytes) of the      |
	//	|       |                          | event log file stops increasing. When the size is greater than or equal to this  |
	//	|       |                          | value, the file growth stops.                                                    |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     9 | EvtRpcVarTypeString      | LogFilePath. File path to the event log file for the channel. The path is saved  |
	//	|       |                          | in the channel config and read out by the server when client requests it.        |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    10 | EvtRpcVarTypeUInt32      | Level. Events with a level property less than or equal to this specified value   |
	//	|       |                          | are logged to the channel.                                                       |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    11 | EvtRpcVarTypeUInt64      | Keywords. Events with a keyword bit contained in the Keywords bitmask set are    |
	//	|       |                          | logged to the channel.                                                           |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    12 | EvtRpcVarTypeGuid        | ControlGuid. A GUID value. The client SHOULD ignore this value.                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    13 | EvtRpcVarTypeUInt64      | BufferSize. Size of the events buffer (in kilobytes) used for asynchronous event |
	//	|       |                          | delivery. This property is for providing events. Typically the events generated  |
	//	|       |                          | by a publisher are first written to memory buffers on the server. Once the       |
	//	|       |                          | buffer used is full, that buffer is written to a disk file. The BufferSize is    |
	//	|       |                          | used to specify the size of the buffer. The server allocates buffers according   |
	//	|       |                          | to the BufferSize value. The number of buffers the server can allocate is        |
	//	|       |                          | controlled by the MinBuffers and MaxBuffers properties. The server's specific    |
	//	|       |                          | implementation can allocate any number of buffers between MinBuffers and         |
	//	|       |                          | MaxBuffers.                                                                      |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    14 | EvtRpcVarTypeUInt32      | MinBuffers. The minimum number of buffers used for asynchronous event delivery.  |
	//	|       |                          | For more information, see the preceding BufferSize information.                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    15 | EvtRpcVarTypeUInt32      | MaxBuffers. The maximum number of buffers used for asynchronous event delivery.  |
	//	|       |                          | For more information, see the preceding BufferSize information.                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    16 | EvtRpcVarTypeUInt32      | Latency. The number of seconds of inactivity (if events are delivered            |
	//	|       |                          | asynchronously and no new events are arriving) after which the event buffers     |
	//	|       |                          | MUST be flushed to the server. As specified in the description for BufferSize    |
	//	|       |                          | property, the server keeps a number of buffers when writing events. If the       |
	//	|       |                          | buffers are full, the server writes the buffers to disk file. However, if a      |
	//	|       |                          | certain amount of time elapses and the buffers are still not full, the server    |
	//	|       |                          | SHOULD write the buffers to disk. That certain amount of time is the latency     |
	//	|       |                          | property.                                                                        |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    17 | EvtRpcVarTypeUInt32      | ClockType. One of two values: 0: SystemTime. Use the system time. When set       |
	//	|       |                          | to this value, the server uses the system time type (which is low-resolution     |
	//	|       |                          | on most platforms) for a time stamp field of any event it writes into this       |
	//	|       |                          | channel. 1: Query Performance Counter. The server uses a high-resolution time    |
	//	|       |                          | type for the time stamp field of any event it writes into this channel. Note The |
	//	|       |                          | timestamp is simply written into the event without any special handling. Which   |
	//	|       |                          | is to say, the server behavior does not change if a channel's Clock type is      |
	//	|       |                          | SystemTime or Query Performance Counter.                                         |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    18 | EvtRpcVarTypeUInt32      | SIDType. One of two values: 0: The events written by the server to this channel  |
	//	|       |                          | will not include the publisher's SID. 1: The events written by the server to     |
	//	|       |                          | this channel will include the publisher's SID.                                   |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    19 | EvtRpcVarTypeStringArray | PublisherList. List of publishers that can raise events into the channel. This   |
	//	|       |                          | returns the same list as is returned by the EvtRpcGetPublisherList method, as    |
	//	|       |                          | specified in section 3.1.4.24.                                                   |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    20 | EvtRpcVarTypeUint32      | FileMax. Maximum number of log files associated with an analytic or debug        |
	//	|       |                          | channel. When the number of logs reaches the specified maximum, the system       |
	//	|       |                          | begins to overwrite the logs, beginning with the oldest. A FileMax value of 0 or |
	//	|       |                          | 1 indicates that only one file is associated with this channel. A FileMax of 0   |
	//	|       |                          | is default.<39>                                                                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	Properties *VariantList `idl:"name:props" json:"properties"`
	// Return: The EvtRpcGetChannelConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetChannelConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_GetChannelConfigOperation) *xxx_GetChannelConfigOperation {
	if op == nil {
		op = &xxx_GetChannelConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Properties = o.Properties
	op.Return = o.Return
	return op
}

func (o *GetChannelConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_GetChannelConfigOperation) {
	if o == nil {
		return
	}
	o.Properties = op.Properties
	o.Return = op.Return
}
func (o *GetChannelConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetChannelConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetChannelConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PutChannelConfigOperation structure represents the EvtRpcPutChannelConfig operation
type xxx_PutChannelConfigOperation struct {
	ChannelPath string       `idl:"name:channelPath;string" json:"channel_path"`
	Flags       uint32       `idl:"name:flags" json:"flags"`
	Properties  *VariantList `idl:"name:props" json:"properties"`
	Error       *Info        `idl:"name:error" json:"error"`
	Return      uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_PutChannelConfigOperation) OpNum() int { return 21 }

func (o *xxx_PutChannelConfigOperation) OpName() string {
	return "/IEventService/v1/EvtRpcPutChannelConfig"
}

func (o *xxx_PutChannelConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ChannelPath) > int(512) {
		return fmt.Errorf("ChannelPath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutChannelConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// channelPath {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ChannelPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// props {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Properties != nil {
			if err := o.Properties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutChannelConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// channelPath {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ChannelPath); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// props {in} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.Properties == nil {
			o.Properties = &VariantList{}
		}
		if err := o.Properties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutChannelConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutChannelConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error != nil {
			if err := o.Error.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PutChannelConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// error {out} (1:{pointer=ref}*(1))(2:{alias=RpcInfo}(struct))
	{
		if o.Error == nil {
			o.Error = &Info{}
		}
		if err := o.Error.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PutChannelConfigRequest structure represents the EvtRpcPutChannelConfig operation request
type PutChannelConfigRequest struct {
	// channelPath: A pointer to a string that contains a channel name (this is not a file
	// path as the parameter name might suggest).
	ChannelPath string `idl:"name:channelPath;string" json:"channel_path"`
	// flags: A 32-bit unsigned integer that indicates what to do depending on the existence
	// of the channel. This MUST be set to one of the following, and the server SHOULD return
	// ERROR_INVALID_PARAMETER (0x00000057) if the flag is not one of the following values.<40>
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | The server MUST open the existing channel entry in its channel table or create a |
	//	|            | new entry if the specified channel is not in the table.                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The server MUST open the existing channel entry in its channel table.            |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Always create a new channel entry in the server's channel table and delete the   |
	//	|            | existing entry.                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 | Only create a new channel entry in the server's channel table.                   |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
	// props: A pointer to an EvtRpcVariantList (section 2.2.9) structure containing channel
	// properties, as defined in the following table.
	//
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|       |                          |                                                                                  |
	//	| INDEX |           TYPE           |                                     MEANING                                      |
	//	|       |                          |                                                                                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     0 | EvtRpcVarTypeBoolean     | Enabled. If true, the channel can accept new events. If false, any attempts to   |
	//	|       |                          | write events into this channel are automatically dropped.                        |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     1 | EvtRpcVarTypeUInt32      | Channel Isolation. It defines the default access permissions for the channel.    |
	//	|       |                          | Three values are allowed: 0: Application. 1: System. 2: Custom. The default      |
	//	|       |                          | isolation is Application. The default permissions for Application are (shown     |
	//	|       |                          | using SDDL): L"O:BAG:SYD:" L"(A;;0xf0007;;;SY)" // local system (read,           |
	//	|       |                          | write, clear) L"(A;;0x7;;;BA)" // built-in admins (read, write, clear)           |
	//	|       |                          | L"(A;;0x7;;;SO)" // server operators (read, write, clear) L"(A;;0x3;;;IU)"       |
	//	|       |                          | // INTERACTIVE LOGON (read, write) L"(A;;0x3;;;SU)" // SERVICES LOGON (read,     |
	//	|       |                          | write) L"(A;;0x3;;;S-1-5-3)" // BATCH LOGON (read, write) L"(A;;0x3;;;S-1-5-33)" |
	//	|       |                          | // write restricted service (read,write) L"(A;;0x1;;;S-1-5-32-573)"; //          |
	//	|       |                          | event log readers (read) The default permissions for System are (shown           |
	//	|       |                          | using SDDL): L"O:BAG:SYD:" L"(A;;0xf0007;;;SY)" // local system (read,           |
	//	|       |                          | write, clear) L"(A;;0x7;;;BA)" // built-in admins (read, write, clear)           |
	//	|       |                          | L"(A;;0x3;;;BO)" // backup operators (read, write) L"(A;;0x5;;;SO)" //           |
	//	|       |                          | server operators (read, clear) L"(A;;0x1;;;IU)" // INTERACTIVE LOGON (read)      |
	//	|       |                          | L"(A;;0x3;;;SU)" // SERVICES LOGON (read, write) L"(A;;0x1;;;S-1-5-3)" //        |
	//	|       |                          | BATCH LOGON (read) L"(A;;0x2;;;S-1-5-33)" // write restricted service (write)    |
	//	|       |                          | L"(A;;0x1;;;S-1-5-32-573)"; // event log readers (read) When the Custom value is |
	//	|       |                          | used, the Access property will contain the defined SDDL.                         |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     2 | EvtRpcVarTypeUInt32      | Channel Type. One of four values: 0: Admin 1: Operational. 2: Analytic 3: Debug  |
	//	|       |                          | For more information, see [MSDN-EVTLGCHWINEVTLG].                                |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     3 | EvtRpcVarTypeString      | OwningPublisher. The name of the publisher that defines and registers the        |
	//	|       |                          | channel with the system.                                                         |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     4 | EvtRpcVarTypeBoolean     | ClassicEventlog. If true, the channel represents an event log created according  |
	//	|       |                          | to the EventLog Remoting Protocol, not this protocol (EventLog Remoting Protocol |
	//	|       |                          | Version 6.0).                                                                    |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     5 | EvtRpcVarTypeString      | Access. A Security Descriptor Description Language (SDDL) string, as specified   |
	//	|       |                          | in [MS-DTYP], that represents access permissions to the channels. If the         |
	//	|       |                          | isolation attribute is set to Application or System, the access descriptor       |
	//	|       |                          | controls read access to the file (the write permissions are ignored). If the     |
	//	|       |                          | isolation attribute is set to Custom, the access descriptor controls write       |
	//	|       |                          | access to the channel and read access to the file.                               |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     6 | EvtRpcVarTypeBoolean     | Retention. If set to true, events can never be overwritten unless explicitly     |
	//	|       |                          | cleared. This is the way to configure the logs to be circular. If set to false,  |
	//	|       |                          | events are overwritten as needed when the event log is full.                     |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     7 | EvtRpcVarTypeBoolean     | AutoBackup. When set to true, the event log file associated with the channel is  |
	//	|       |                          | closed as soon as it reaches the maximum size specified by the MaxSize property, |
	//	|       |                          | and a new file is opened to accept new events. If the new file reaches maximum   |
	//	|       |                          | size, another new file will be generated and the previous new file will be       |
	//	|       |                          | backed up. The events in backed up files cannot be queried from this channel in  |
	//	|       |                          | the server unless the client specifies the backup log file names in a separate   |
	//	|       |                          | query.                                                                           |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     8 | EvtRpcVarTypeUInt64      | MaxSize. The value that indicates at which point the size (in bytes) of the      |
	//	|       |                          | event log file stops increasing. When the size is greater than or equal to this  |
	//	|       |                          | value, the file growth stops.                                                    |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|     9 | EvtRpcVarTypeString      | LogFilePath. The server changes the file path to the event log file for the      |
	//	|       |                          | channel.                                                                         |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    10 | EvtRpcVarTypeUInt32      | Level. Events with a level property less than or equal to this specified value   |
	//	|       |                          | are logged to the channel.                                                       |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    11 | EvtRpcVarTypeUInt64      | Keywords. Events with a keyword bit contained in the keywords bitmask set are    |
	//	|       |                          | logged to the channel.                                                           |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    12 | EvtRpcVarTypeGuid        | ControlGuid. A GUID value. The server SHOULD ignore this value.<41>              |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    19 | EvtRpcVarTypeStringArray | PublisherList. A list of publishers that can raise events into the channel.      |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	//	|    20 | EvtRpcVarTypeUInt32      | FileMax. The maximum number of log files associated with an analytic or debug    |
	//	|       |                          | channel. When the number of logs reaches the specified maximum, the system       |
	//	|       |                          | begins to overwrite the logs, beginning with the oldest. A FileMax value of 0 or |
	//	|       |                          | 1 indicates that only one file is associated with this channel. A MaxFile of 0   |
	//	|       |                          | is the default.                                                                  |
	//	+-------+--------------------------+----------------------------------------------------------------------------------+
	Properties *VariantList `idl:"name:props" json:"properties"`
}

func (o *PutChannelConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_PutChannelConfigOperation) *xxx_PutChannelConfigOperation {
	if op == nil {
		op = &xxx_PutChannelConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelPath = o.ChannelPath
	op.Flags = o.Flags
	op.Properties = o.Properties
	return op
}

func (o *PutChannelConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_PutChannelConfigOperation) {
	if o == nil {
		return
	}
	o.ChannelPath = op.ChannelPath
	o.Flags = op.Flags
	o.Properties = op.Properties
}
func (o *PutChannelConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PutChannelConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutChannelConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PutChannelConfigResponse structure represents the EvtRpcPutChannelConfig operation response
type PutChannelConfigResponse struct {
	// error: A pointer to an RpcInfo (section 2.2.1) structure in which to place error
	// information in the case of a failure. The RpcInfo (section 2.2.1) structure fields
	// MUST be set to nonzero values if the error is related to a particular property. All
	// nonzero values MUST be treated the same. If the method succeeds, the server MUST
	// set all of the values in the structure to 0.
	Error *Info `idl:"name:error" json:"error"`
	// Return: The EvtRpcPutChannelConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PutChannelConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_PutChannelConfigOperation) *xxx_PutChannelConfigOperation {
	if op == nil {
		op = &xxx_PutChannelConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Error = o.Error
	op.Return = o.Return
	return op
}

func (o *PutChannelConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_PutChannelConfigOperation) {
	if o == nil {
		return
	}
	o.Error = op.Error
	o.Return = op.Return
}
func (o *PutChannelConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PutChannelConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PutChannelConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherListOperation structure represents the EvtRpcGetPublisherList operation
type xxx_GetPublisherListOperation struct {
	Flags              uint32   `idl:"name:flags" json:"flags"`
	PublisherIDsLength uint32   `idl:"name:numPublisherIds" json:"publisher_ids_length"`
	PublisherIDs       []string `idl:"name:publisherIds;size_is:(, numPublisherIds);string" json:"publisher_ids"`
	Return             uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherListOperation) OpNum() int { return 22 }

func (o *xxx_GetPublisherListOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetPublisherList"
}

func (o *xxx_GetPublisherListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PublisherIDs != nil && o.PublisherIDsLength == 0 {
		o.PublisherIDsLength = uint32(len(o.PublisherIDs))
	}
	if len(o.PublisherIDs) > int(8192) {
		return fmt.Errorf("PublisherIDs is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numPublisherIds {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PublisherIDsLength); err != nil {
			return err
		}
	}
	// publisherIds {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numPublisherIds]*(1)[dim:0,string,null](wchar))
	{
		if o.PublisherIDs != nil || o.PublisherIDsLength > 0 {
			_ptr_publisherIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PublisherIDsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PublisherIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PublisherIDs[i1] != "" {
						_ptr_publisherIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.PublisherIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.PublisherIDs[i1], _ptr_publisherIds); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PublisherIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherIDs, _ptr_publisherIds); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numPublisherIds {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PublisherIDsLength); err != nil {
			return err
		}
	}
	// publisherIds {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numPublisherIds]*(1)[dim:0,string,null](wchar))
	{
		_ptr_publisherIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PublisherIDs", sizeInfo[0])
			}
			o.PublisherIDs = make([]string, sizeInfo[0])
			for i1 := range o.PublisherIDs {
				i1 := i1
				_ptr_publisherIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.PublisherIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_publisherIds := func(ptr interface{}) { o.PublisherIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.PublisherIDs[i1], _s_publisherIds, _ptr_publisherIds); err != nil {
					return err
				}
			}
			return nil
		})
		_s_publisherIds := func(ptr interface{}) { o.PublisherIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.PublisherIDs, _s_publisherIds, _ptr_publisherIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherListRequest structure represents the EvtRpcGetPublisherList operation request
type GetPublisherListRequest struct {
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<45>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetPublisherListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherListOperation) *xxx_GetPublisherListOperation {
	if op == nil {
		op = &xxx_GetPublisherListOperation{}
	}
	if o == nil {
		return op
	}
	op.Flags = o.Flags
	return op
}

func (o *GetPublisherListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherListOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
}
func (o *GetPublisherListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherListResponse structure represents the EvtRpcGetPublisherList operation response
type GetPublisherListResponse struct {
	// numPublisherIds: A pointer to a 32-bit unsigned integer that contains the number
	// of publisher names.
	PublisherIDsLength uint32 `idl:"name:numPublisherIds" json:"publisher_ids_length"`
	// publisherIds: A pointer to an array of strings that contain publisher names.
	PublisherIDs []string `idl:"name:publisherIds;size_is:(, numPublisherIds);string" json:"publisher_ids"`
	// Return: The EvtRpcGetPublisherList return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherListOperation) *xxx_GetPublisherListOperation {
	if op == nil {
		op = &xxx_GetPublisherListOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherIDsLength = o.PublisherIDsLength
	op.PublisherIDs = o.PublisherIDs
	op.Return = o.Return
	return op
}

func (o *GetPublisherListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherListOperation) {
	if o == nil {
		return
	}
	o.PublisherIDsLength = op.PublisherIDsLength
	o.PublisherIDs = op.PublisherIDs
	o.Return = op.Return
}
func (o *GetPublisherListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherListForChannelOperation structure represents the EvtRpcGetPublisherListForChannel operation
type xxx_GetPublisherListForChannelOperation struct {
	ChannelName        string   `idl:"name:channelName" json:"channel_name"`
	Flags              uint32   `idl:"name:flags" json:"flags"`
	PublisherIDsLength uint32   `idl:"name:numPublisherIds" json:"publisher_ids_length"`
	PublisherIDs       []string `idl:"name:publisherIds;size_is:(, numPublisherIds);string" json:"publisher_ids"`
	Return             uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherListForChannelOperation) OpNum() int { return 23 }

func (o *xxx_GetPublisherListForChannelOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetPublisherListForChannel"
}

func (o *xxx_GetPublisherListForChannelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListForChannelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// channelName {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.ChannelName); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListForChannelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// channelName {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.ChannelName); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListForChannelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PublisherIDs != nil && o.PublisherIDsLength == 0 {
		o.PublisherIDsLength = uint32(len(o.PublisherIDs))
	}
	if len(o.PublisherIDs) > int(8192) {
		return fmt.Errorf("PublisherIDs is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListForChannelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numPublisherIds {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PublisherIDsLength); err != nil {
			return err
		}
	}
	// publisherIds {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numPublisherIds]*(1)[dim:0,string,null](wchar))
	{
		if o.PublisherIDs != nil || o.PublisherIDsLength > 0 {
			_ptr_publisherIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.PublisherIDsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PublisherIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PublisherIDs[i1] != "" {
						_ptr_publisherIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.PublisherIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.PublisherIDs[i1], _ptr_publisherIds); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PublisherIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherIDs, _ptr_publisherIds); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherListForChannelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numPublisherIds {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PublisherIDsLength); err != nil {
			return err
		}
	}
	// publisherIds {out} (1:{string, pointer=ref, range=(0,8192)}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=numPublisherIds]*(1)[dim:0,string,null](wchar))
	{
		_ptr_publisherIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PublisherIDs", sizeInfo[0])
			}
			o.PublisherIDs = make([]string, sizeInfo[0])
			for i1 := range o.PublisherIDs {
				i1 := i1
				_ptr_publisherIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.PublisherIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_publisherIds := func(ptr interface{}) { o.PublisherIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.PublisherIDs[i1], _s_publisherIds, _ptr_publisherIds); err != nil {
					return err
				}
			}
			return nil
		})
		_s_publisherIds := func(ptr interface{}) { o.PublisherIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.PublisherIDs, _s_publisherIds, _ptr_publisherIds); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherListForChannelRequest structure represents the EvtRpcGetPublisherListForChannel operation request
type GetPublisherListForChannelRequest struct {
	// channelName:  A pointer to a string that contains the name of the channel for which
	// the publisher list is needed.
	ChannelName string `idl:"name:channelName" json:"channel_name"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<47>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetPublisherListForChannelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherListForChannelOperation) *xxx_GetPublisherListForChannelOperation {
	if op == nil {
		op = &xxx_GetPublisherListForChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.ChannelName = o.ChannelName
	op.Flags = o.Flags
	return op
}

func (o *GetPublisherListForChannelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherListForChannelOperation) {
	if o == nil {
		return
	}
	o.ChannelName = op.ChannelName
	o.Flags = op.Flags
}
func (o *GetPublisherListForChannelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherListForChannelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherListForChannelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherListForChannelResponse structure represents the EvtRpcGetPublisherListForChannel operation response
type GetPublisherListForChannelResponse struct {
	// numPublisherIds: A pointer to a 32-bit unsigned integer that contains the number
	// of publishers that are registered and that can write to the log.
	PublisherIDsLength uint32 `idl:"name:numPublisherIds" json:"publisher_ids_length"`
	// publisherIds: A pointer to an array of strings that contain publisher names.
	PublisherIDs []string `idl:"name:publisherIds;size_is:(, numPublisherIds);string" json:"publisher_ids"`
	// Return: The EvtRpcGetPublisherListForChannel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherListForChannelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherListForChannelOperation) *xxx_GetPublisherListForChannelOperation {
	if op == nil {
		op = &xxx_GetPublisherListForChannelOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherIDsLength = o.PublisherIDsLength
	op.PublisherIDs = o.PublisherIDs
	op.Return = o.Return
	return op
}

func (o *GetPublisherListForChannelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherListForChannelOperation) {
	if o == nil {
		return
	}
	o.PublisherIDsLength = op.PublisherIDsLength
	o.PublisherIDs = op.PublisherIDs
	o.Return = op.Return
}
func (o *GetPublisherListForChannelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherListForChannelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherListForChannelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherMetadataOperation structure represents the EvtRpcGetPublisherMetadata operation
type xxx_GetPublisherMetadataOperation struct {
	PublisherID                 string             `idl:"name:publisherId;string;pointer:unique" json:"publisher_id"`
	LogFilePath                 string             `idl:"name:logFilePath;string;pointer:unique" json:"log_file_path"`
	Locale                      uint32             `idl:"name:locale" json:"locale"`
	Flags                       uint32             `idl:"name:flags" json:"flags"`
	PublisherMetadataProperties *VariantList       `idl:"name:pubMetadataProps" json:"publisher_metadata_properties"`
	PublisherMetadata           *PublisherMetadata `idl:"name:pubMetadata" json:"publisher_metadata"`
	Return                      uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherMetadataOperation) OpNum() int { return 24 }

func (o *xxx_GetPublisherMetadataOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetPublisherMetadata"
}

func (o *xxx_GetPublisherMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.PublisherID) > int(2048) {
		return fmt.Errorf("PublisherID is out of range")
	}
	if len(o.LogFilePath) > int(32768) {
		return fmt.Errorf("LogFilePath is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// publisherId {in} (1:{string, pointer=unique, range=(0,2048), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.PublisherID != "" {
			_ptr_publisherId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PublisherID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherID, _ptr_publisherId); err != nil {
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
	// logFilePath {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.LogFilePath != "" {
			_ptr_logFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.LogFilePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.LogFilePath, _ptr_logFilePath); err != nil {
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
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// publisherId {in} (1:{string, pointer=unique, range=(0,2048), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_publisherId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PublisherID); err != nil {
				return err
			}
			return nil
		})
		_s_publisherId := func(ptr interface{}) { o.PublisherID = *ptr.(*string) }
		if err := w.ReadPointer(&o.PublisherID, _s_publisherId, _ptr_publisherId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// logFilePath {in} (1:{string, pointer=unique, range=(0,32768), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_logFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.LogFilePath); err != nil {
				return err
			}
			return nil
		})
		_s_logFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.LogFilePath, _s_logFilePath, _ptr_logFilePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pubMetadataProps {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.PublisherMetadataProperties != nil {
			if err := o.PublisherMetadataProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pubMetadata {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherMetadata != nil {
			if err := o.PublisherMetadata.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PublisherMetadata{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pubMetadataProps {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.PublisherMetadataProperties == nil {
			o.PublisherMetadataProperties = &VariantList{}
		}
		if err := o.PublisherMetadataProperties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pubMetadata {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherMetadata == nil {
			o.PublisherMetadata = &PublisherMetadata{}
		}
		if err := o.PublisherMetadata.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherMetadataRequest structure represents the EvtRpcGetPublisherMetadata operation request
type GetPublisherMetadataRequest struct {
	// publisherId: A pointer to a string that contains the publisher for which information
	// is needed.
	PublisherID string `idl:"name:publisherId;string;pointer:unique" json:"publisher_id"`
	// logFilePath: A pointer to a null string that MUST be ignored on receipt.
	LogFilePath string `idl:"name:logFilePath;string;pointer:unique" json:"log_file_path"`
	// locale: A Locale value, as specified in [MS-GPSI]. This is used later if the pubMetadata
	// handle is used for rendering.
	Locale uint32 `idl:"name:locale" json:"locale"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<48>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetPublisherMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherMetadataOperation) *xxx_GetPublisherMetadataOperation {
	if op == nil {
		op = &xxx_GetPublisherMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherID = o.PublisherID
	op.LogFilePath = o.LogFilePath
	op.Locale = o.Locale
	op.Flags = o.Flags
	return op
}

func (o *GetPublisherMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherMetadataOperation) {
	if o == nil {
		return
	}
	o.PublisherID = op.PublisherID
	o.LogFilePath = op.LogFilePath
	o.Locale = op.Locale
	o.Flags = op.Flags
}
func (o *GetPublisherMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherMetadataResponse structure represents the EvtRpcGetPublisherMetadata operation response
type GetPublisherMetadataResponse struct {
	// pubMetadataProps: A pointer to an EvtRpcVariantList (section 2.2.9) structure containing
	// publisher properties.
	PublisherMetadataProperties *VariantList `idl:"name:pubMetadataProps" json:"publisher_metadata_properties"`
	// pubMetadata: A pointer to a publisher handle. This parameter is an RPC context handle,
	// as specified in [C706], Context Handles. For information on handle security and authentication
	// considerations, see sections 2.2.20 and 5.1.
	PublisherMetadata *PublisherMetadata `idl:"name:pubMetadata" json:"publisher_metadata"`
	// Return: The EvtRpcGetPublisherMetadata return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherMetadataOperation) *xxx_GetPublisherMetadataOperation {
	if op == nil {
		op = &xxx_GetPublisherMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherMetadataProperties = o.PublisherMetadataProperties
	op.PublisherMetadata = o.PublisherMetadata
	op.Return = o.Return
	return op
}

func (o *GetPublisherMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherMetadataOperation) {
	if o == nil {
		return
	}
	o.PublisherMetadataProperties = op.PublisherMetadataProperties
	o.PublisherMetadata = op.PublisherMetadata
	o.Return = op.Return
}
func (o *GetPublisherMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPublisherResourceMetadataOperation structure represents the EvtRpcGetPublisherResourceMetadata operation
type xxx_GetPublisherResourceMetadataOperation struct {
	Handle                      *PublisherMetadata `idl:"name:handle" json:"handle"`
	PropertyID                  uint32             `idl:"name:propertyId" json:"property_id"`
	Flags                       uint32             `idl:"name:flags" json:"flags"`
	PublisherMetadataProperties *VariantList       `idl:"name:pubMetadataProps" json:"publisher_metadata_properties"`
	Return                      uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherResourceMetadataOperation) OpNum() int { return 25 }

func (o *xxx_GetPublisherResourceMetadataOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetPublisherResourceMetadata"
}

func (o *xxx_GetPublisherResourceMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherResourceMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PublisherMetadata{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// propertyId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PropertyID); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherResourceMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// handle {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &PublisherMetadata{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// propertyId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PropertyID); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherResourceMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherResourceMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pubMetadataProps {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.PublisherMetadataProperties != nil {
			if err := o.PublisherMetadataProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherResourceMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pubMetadataProps {out} (1:{pointer=ref}*(1))(2:{alias=EvtRpcVariantList}(struct))
	{
		if o.PublisherMetadataProperties == nil {
			o.PublisherMetadataProperties = &VariantList{}
		}
		if err := o.PublisherMetadataProperties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPublisherResourceMetadataRequest structure represents the EvtRpcGetPublisherResourceMetadata operation request
type GetPublisherResourceMetadataRequest struct {
	// handle: A handle to an event log. This handle is returned by the EvtRpcGetPublisherMetadata
	// (Opnum 24) method. This parameter is an RPC context handle, as specified in [C706],
	// Context Handles.
	Handle *PublisherMetadata `idl:"name:handle" json:"handle"`
	// propertyId: Type of information as specified in the following table.
	//
	//	+------------+--------------------------+
	//	|            |                          |
	//	|   VALUE    |         MEANING          |
	//	|            |                          |
	//	+------------+--------------------------+
	//	+------------+--------------------------+
	//	| 0x00000004 | Publisher help link.     |
	//	+------------+--------------------------+
	//	| 0x00000005 | Publisher friendly name. |
	//	+------------+--------------------------+
	//	| 0x0000000C | Level information.       |
	//	+------------+--------------------------+
	//	| 0x00000010 | Task information.        |
	//	+------------+--------------------------+
	//	| 0x00000015 | Opcode information.      |
	//	+------------+--------------------------+
	//	| 0x00000019 | Keyword information.     |
	//	+------------+--------------------------+
	PropertyID uint32 `idl:"name:propertyId" json:"property_id"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<50>
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetPublisherResourceMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherResourceMetadataOperation) *xxx_GetPublisherResourceMetadataOperation {
	if op == nil {
		op = &xxx_GetPublisherResourceMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.Handle = o.Handle
	op.PropertyID = o.PropertyID
	op.Flags = o.Flags
	return op
}

func (o *GetPublisherResourceMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherResourceMetadataOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.PropertyID = op.PropertyID
	o.Flags = op.Flags
}
func (o *GetPublisherResourceMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherResourceMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherResourceMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherResourceMetadataResponse structure represents the EvtRpcGetPublisherResourceMetadata operation response
type GetPublisherResourceMetadataResponse struct {
	// pubMetadataProps: Pointer to an EvtRpcVariantList (section 2.2.9) structure. This
	// list MUST contain multiple entries.
	PublisherMetadataProperties *VariantList `idl:"name:pubMetadataProps" json:"publisher_metadata_properties"`
	// Return: The EvtRpcGetPublisherResourceMetadata return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherResourceMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherResourceMetadataOperation) *xxx_GetPublisherResourceMetadataOperation {
	if op == nil {
		op = &xxx_GetPublisherResourceMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherMetadataProperties = o.PublisherMetadataProperties
	op.Return = o.Return
	return op
}

func (o *GetPublisherResourceMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherResourceMetadataOperation) {
	if o == nil {
		return
	}
	o.PublisherMetadataProperties = op.PublisherMetadataProperties
	o.Return = op.Return
}
func (o *GetPublisherResourceMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherResourceMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherResourceMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventMetadataEnumOperation structure represents the EvtRpcGetEventMetadataEnum operation
type xxx_GetEventMetadataEnumOperation struct {
	PublisherMetadata *PublisherMetadata `idl:"name:pubMetadata" json:"publisher_metadata"`
	Flags             uint32             `idl:"name:flags" json:"flags"`
	ReservedForFilter string             `idl:"name:reservedForFilter;string;pointer:unique" json:"reserved_for_filter"`
	EventMetadataEnum *EventMetadataEnum `idl:"name:eventMetaDataEnum" json:"event_metadata_enum"`
	Return            uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventMetadataEnumOperation) OpNum() int { return 26 }

func (o *xxx_GetEventMetadataEnumOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetEventMetadataEnum"
}

func (o *xxx_GetEventMetadataEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.ReservedForFilter) > int(1048576) {
		return fmt.Errorf("ReservedForFilter is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventMetadataEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pubMetadata {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherMetadata != nil {
			if err := o.PublisherMetadata.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PublisherMetadata{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// reservedForFilter {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ReservedForFilter != "" {
			_ptr_reservedForFilter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ReservedForFilter); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ReservedForFilter, _ptr_reservedForFilter); err != nil {
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
	return nil
}

func (o *xxx_GetEventMetadataEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pubMetadata {in} (1:{context_handle, alias=PCONTEXT_HANDLE_PUBLISHER_METADATA, names=ndr_context_handle}(struct))
	{
		if o.PublisherMetadata == nil {
			o.PublisherMetadata = &PublisherMetadata{}
		}
		if err := o.PublisherMetadata.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// reservedForFilter {in} (1:{string, pointer=unique, range=(0,1048576), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_reservedForFilter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ReservedForFilter); err != nil {
				return err
			}
			return nil
		})
		_s_reservedForFilter := func(ptr interface{}) { o.ReservedForFilter = *ptr.(*string) }
		if err := w.ReadPointer(&o.ReservedForFilter, _s_reservedForFilter, _ptr_reservedForFilter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventMetadataEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventMetadataEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// eventMetaDataEnum {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_EVENT_METADATA_ENUM, names=ndr_context_handle}(struct))
	{
		if o.EventMetadataEnum != nil {
			if err := o.EventMetadataEnum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EventMetadataEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventMetadataEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// eventMetaDataEnum {out} (1:{context_handle, pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_EVENT_METADATA_ENUM, names=ndr_context_handle}(struct))
	{
		if o.EventMetadataEnum == nil {
			o.EventMetadataEnum = &EventMetadataEnum{}
		}
		if err := o.EventMetadataEnum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetEventMetadataEnumRequest structure represents the EvtRpcGetEventMetadataEnum operation request
type GetEventMetadataEnumRequest struct {
	// pubMetadata: This parameter is an RPC context handle, as specified in [C706], Context
	// Handles. For information on handle security and authentication considerations, see
	// sections 2.2.20 and 5.1.
	PublisherMetadata *PublisherMetadata `idl:"name:pubMetadata" json:"publisher_metadata"`
	// flags: A 32-bit unsigned integer that MUST be set to zero when sent and MAY be ignored
	// on receipt.<54>
	Flags uint32 `idl:"name:flags" json:"flags"`
	// reservedForFilter: A pointer to a null string that MUST be ignored on receipt.
	ReservedForFilter string `idl:"name:reservedForFilter;string;pointer:unique" json:"reserved_for_filter"`
}

func (o *GetEventMetadataEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEventMetadataEnumOperation) *xxx_GetEventMetadataEnumOperation {
	if op == nil {
		op = &xxx_GetEventMetadataEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.PublisherMetadata = o.PublisherMetadata
	op.Flags = o.Flags
	op.ReservedForFilter = o.ReservedForFilter
	return op
}

func (o *GetEventMetadataEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventMetadataEnumOperation) {
	if o == nil {
		return
	}
	o.PublisherMetadata = op.PublisherMetadata
	o.Flags = op.Flags
	o.ReservedForFilter = op.ReservedForFilter
}
func (o *GetEventMetadataEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEventMetadataEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventMetadataEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventMetadataEnumResponse structure represents the EvtRpcGetEventMetadataEnum operation response
type GetEventMetadataEnumResponse struct {
	// eventMetaDataEnum: A pointer to an event numeration handle. This parameter is an
	// RPC context handle, as specified in [C706], Context Handles.
	EventMetadataEnum *EventMetadataEnum `idl:"name:eventMetaDataEnum" json:"event_metadata_enum"`
	// Return: The EvtRpcGetEventMetadataEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetEventMetadataEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEventMetadataEnumOperation) *xxx_GetEventMetadataEnumOperation {
	if op == nil {
		op = &xxx_GetEventMetadataEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.EventMetadataEnum = o.EventMetadataEnum
	op.Return = o.Return
	return op
}

func (o *GetEventMetadataEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventMetadataEnumOperation) {
	if o == nil {
		return
	}
	o.EventMetadataEnum = op.EventMetadataEnum
	o.Return = op.Return
}
func (o *GetEventMetadataEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEventMetadataEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventMetadataEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNextEventMetadataOperation structure represents the EvtRpcGetNextEventMetadata operation
type xxx_GetNextEventMetadataOperation struct {
	EventMetadataEnum      *EventMetadataEnum `idl:"name:eventMetaDataEnum" json:"event_metadata_enum"`
	Flags                  uint32             `idl:"name:flags" json:"flags"`
	RequestedLength        uint32             `idl:"name:numRequested" json:"requested_length"`
	ReturnedLength         uint32             `idl:"name:numReturned" json:"returned_length"`
	EventMetadataInstances []*VariantList     `idl:"name:eventMetadataInstances;size_is:(, numReturned)" json:"event_metadata_instances"`
	Return                 uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNextEventMetadataOperation) OpNum() int { return 27 }

func (o *xxx_GetNextEventMetadataOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetNextEventMetadata"
}

func (o *xxx_GetNextEventMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextEventMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// eventMetaDataEnum {in} (1:{context_handle, alias=PCONTEXT_HANDLE_EVENT_METADATA_ENUM, names=ndr_context_handle}(struct))
	{
		if o.EventMetadataEnum != nil {
			if err := o.EventMetadataEnum.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EventMetadataEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// numRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestedLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextEventMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// eventMetaDataEnum {in} (1:{context_handle, alias=PCONTEXT_HANDLE_EVENT_METADATA_ENUM, names=ndr_context_handle}(struct))
	{
		if o.EventMetadataEnum == nil {
			o.EventMetadataEnum = &EventMetadataEnum{}
		}
		if err := o.EventMetadataEnum.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// numRequested {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestedLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextEventMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.EventMetadataInstances != nil && o.ReturnedLength == 0 {
		o.ReturnedLength = uint32(len(o.EventMetadataInstances))
	}
	if len(o.EventMetadataInstances) > int(256) {
		return fmt.Errorf("EventMetadataInstances is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextEventMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// numReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnedLength); err != nil {
			return err
		}
	}
	// eventMetadataInstances {out} (1:{pointer=ref, range=(0,256)}*(2)*(1))(2:{alias=EvtRpcVariantList}[dim:0,size_is=numReturned](struct))
	{
		if o.EventMetadataInstances != nil || o.ReturnedLength > 0 {
			_ptr_eventMetadataInstances := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ReturnedLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.EventMetadataInstances {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.EventMetadataInstances[i1] != nil {
						if err := o.EventMetadataInstances[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.EventMetadataInstances); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&VariantList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventMetadataInstances, _ptr_eventMetadataInstances); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNextEventMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// numReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnedLength); err != nil {
			return err
		}
	}
	// eventMetadataInstances {out} (1:{pointer=ref, range=(0,256)}*(2)*(1))(2:{alias=EvtRpcVariantList}[dim:0,size_is=numReturned](struct))
	{
		_ptr_eventMetadataInstances := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.EventMetadataInstances", sizeInfo[0])
			}
			o.EventMetadataInstances = make([]*VariantList, sizeInfo[0])
			for i1 := range o.EventMetadataInstances {
				i1 := i1
				if o.EventMetadataInstances[i1] == nil {
					o.EventMetadataInstances[i1] = &VariantList{}
				}
				if err := o.EventMetadataInstances[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_eventMetadataInstances := func(ptr interface{}) { o.EventMetadataInstances = *ptr.(*[]*VariantList) }
		if err := w.ReadPointer(&o.EventMetadataInstances, _s_eventMetadataInstances, _ptr_eventMetadataInstances); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNextEventMetadataRequest structure represents the EvtRpcGetNextEventMetadata operation request
type GetNextEventMetadataRequest struct {
	// eventMetaDataEnum: A handle to an event metadata enumerator. This parameter is an
	// RPC context handle, as specified in [C706], Context Handles. For information on handle
	// security and authentication considerations, see sections 2.2.20 and 5.1. This is
	// the value which comes from the return parameter eventMetaDataEnum of function EvtRpcGetEventMetadataEnum
	// (as specified in 3.1.4.27).
	EventMetadataEnum *EventMetadataEnum `idl:"name:eventMetaDataEnum" json:"event_metadata_enum"`
	// flags: A 32-bit unsigned integer that MUST be set to 0x00000000 when sent and MAY
	// be ignored on receipt.<56>
	Flags uint32 `idl:"name:flags" json:"flags"`
	// numRequested: A 32-bit unsigned integer that contains the number of events for which
	// the properties are needed.
	RequestedLength uint32 `idl:"name:numRequested" json:"requested_length"`
}

func (o *GetNextEventMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNextEventMetadataOperation) *xxx_GetNextEventMetadataOperation {
	if op == nil {
		op = &xxx_GetNextEventMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.EventMetadataEnum = o.EventMetadataEnum
	op.Flags = o.Flags
	op.RequestedLength = o.RequestedLength
	return op
}

func (o *GetNextEventMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNextEventMetadataOperation) {
	if o == nil {
		return
	}
	o.EventMetadataEnum = op.EventMetadataEnum
	o.Flags = op.Flags
	o.RequestedLength = op.RequestedLength
}
func (o *GetNextEventMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNextEventMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextEventMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNextEventMetadataResponse structure represents the EvtRpcGetNextEventMetadata operation response
type GetNextEventMetadataResponse struct {
	// numReturned: A pointer to a 32-bit unsigned integer that contains the number of events
	// for which the properties are retrieved.
	ReturnedLength uint32 `idl:"name:numReturned" json:"returned_length"`
	// eventMetadataInstances: A pointer to an array of EvtRpcVariantList (section 2.2.9)
	// structures.
	EventMetadataInstances []*VariantList `idl:"name:eventMetadataInstances;size_is:(, numReturned)" json:"event_metadata_instances"`
	// Return: The EvtRpcGetNextEventMetadata return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNextEventMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNextEventMetadataOperation) *xxx_GetNextEventMetadataOperation {
	if op == nil {
		op = &xxx_GetNextEventMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.ReturnedLength = o.ReturnedLength
	op.EventMetadataInstances = o.EventMetadataInstances
	op.Return = o.Return
	return op
}

func (o *GetNextEventMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNextEventMetadataOperation) {
	if o == nil {
		return
	}
	o.ReturnedLength = op.ReturnedLength
	o.EventMetadataInstances = op.EventMetadataInstances
	o.Return = op.Return
}
func (o *GetNextEventMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNextEventMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNextEventMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassicLogDisplayNameOperation structure represents the EvtRpcGetClassicLogDisplayName operation
type xxx_GetClassicLogDisplayNameOperation struct {
	LogName     string `idl:"name:logName;string" json:"log_name"`
	Locale      uint32 `idl:"name:locale" json:"locale"`
	Flags       uint32 `idl:"name:flags" json:"flags"`
	DisplayName string `idl:"name:displayName" json:"display_name"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassicLogDisplayNameOperation) OpNum() int { return 28 }

func (o *xxx_GetClassicLogDisplayNameOperation) OpName() string {
	return "/IEventService/v1/EvtRpcGetClassicLogDisplayName"
}

func (o *xxx_GetClassicLogDisplayNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if len(o.LogName) > int(512) {
		return fmt.Errorf("LogName is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassicLogDisplayNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// logName {in} (1:{string, range=(1,512), alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LogName); err != nil {
			return err
		}
	}
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassicLogDisplayNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// logName {in} (1:{string, range=(1,512), alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogName); err != nil {
			return err
		}
	}
	// locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Locale); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassicLogDisplayNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassicLogDisplayNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// displayName {out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		if o.DisplayName != "" {
			_ptr_displayName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.DisplayName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DisplayName, _ptr_displayName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassicLogDisplayNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// displayName {out} (1:{pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		_ptr_displayName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.DisplayName); err != nil {
				return err
			}
			return nil
		})
		_s_displayName := func(ptr interface{}) { o.DisplayName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DisplayName, _s_displayName, _ptr_displayName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassicLogDisplayNameRequest structure represents the EvtRpcGetClassicLogDisplayName operation request
type GetClassicLogDisplayNameRequest struct {
	// logName: The channel name for which the descriptive name is needed.
	LogName string `idl:"name:logName;string" json:"log_name"`
	// locale: The locale, as specified in [MS-GPSI] Appendix A, to be used for localizing
	// the log.
	Locale uint32 `idl:"name:locale" json:"locale"`
	// flags: A 32-bit unsigned integer that MUST be set to one of the following values:
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                     MEANING                                      |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x0   | If a locale is specified, that locale will be used and no fallback locale will   |
	//	|       | be attempted if the locale is not present.                                       |
	//	+-------+----------------------------------------------------------------------------------+
	//	| 0x100 | If set, instructs the server to pick the best locale, if the locale specified by |
	//	|       | the locale parameter is not present. Please see the following processing rules   |
	//	|       | for more information on how the server picks the best locale.                    |
	//	+-------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:flags" json:"flags"`
}

func (o *GetClassicLogDisplayNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClassicLogDisplayNameOperation) *xxx_GetClassicLogDisplayNameOperation {
	if op == nil {
		op = &xxx_GetClassicLogDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.LogName = o.LogName
	op.Locale = o.Locale
	op.Flags = o.Flags
	return op
}

func (o *GetClassicLogDisplayNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassicLogDisplayNameOperation) {
	if o == nil {
		return
	}
	o.LogName = op.LogName
	o.Locale = op.Locale
	o.Flags = op.Flags
}
func (o *GetClassicLogDisplayNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClassicLogDisplayNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassicLogDisplayNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassicLogDisplayNameResponse structure represents the EvtRpcGetClassicLogDisplayName operation response
type GetClassicLogDisplayNameResponse struct {
	// displayName: Returned display name.
	DisplayName string `idl:"name:displayName" json:"display_name"`
	// Return: The EvtRpcGetClassicLogDisplayName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClassicLogDisplayNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClassicLogDisplayNameOperation) *xxx_GetClassicLogDisplayNameOperation {
	if op == nil {
		op = &xxx_GetClassicLogDisplayNameOperation{}
	}
	if o == nil {
		return op
	}
	op.DisplayName = o.DisplayName
	op.Return = o.Return
	return op
}

func (o *GetClassicLogDisplayNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassicLogDisplayNameOperation) {
	if o == nil {
		return
	}
	o.DisplayName = op.DisplayName
	o.Return = op.Return
}
func (o *GetClassicLogDisplayNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClassicLogDisplayNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassicLogDisplayNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
