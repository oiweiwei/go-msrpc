package ieventservice

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

// IEventService server interface.
type EventServiceServer interface {

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
	RegisterRemoteSubscription(context.Context, *RegisterRemoteSubscriptionRequest) (*RegisterRemoteSubscriptionResponse, error)

	// The EvtRpcRemoteSubscriptionNextAsync (Opnum 1) method is used by a client to request
	// asynchronous delivery of events that are delivered to a subscription.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RemoteSubscriptionNextAsync(context.Context, *RemoteSubscriptionNextAsyncRequest) (*RemoteSubscriptionNextAsyncResponse, error)

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
	RemoteSubscriptionNext(context.Context, *RemoteSubscriptionNextRequest) (*RemoteSubscriptionNextResponse, error)

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
	RemoteSubscriptionWaitAsync(context.Context, *RemoteSubscriptionWaitAsyncRequest) (*RemoteSubscriptionWaitAsyncResponse, error)

	// The EvtRpcRegisterControllableOperation (Opnum 4) method obtains a CONTEXT_HANDLE_OPERATION_CONTROL
	// handle that can be used to cancel other operations.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RegisterControllableOperation(context.Context, *RegisterControllableOperationRequest) (*RegisterControllableOperationResponse, error)

	// The EvtRpcRegisterLogQuery (Opnum 5) method is used to query one or more channels.
	// It can also be used to query a specific file. Actual retrieval of events is done
	// by subsequent calls to the EvtRpcQueryNext (section 3.1.4.13) method.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RegisterLogQuery(context.Context, *RegisterLogQueryRequest) (*RegisterLogQueryResponse, error)

	// The EvtRpcClearLog (Opnum 6) method instructs the server to clear all the events
	// in a live channel, and optionally, to create a backup event log before the clear
	// takes place.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) on success; otherwise, it MUST
	// return an implementation-specific nonzero value as specified in [MS-ERREF].
	//
	// The server does not validate the control handle passed to EvtRpcClearLog and it SHOULD
	// assume that this parameter is always valid when the method is invoked.
	ClearLog(context.Context, *ClearLogRequest) (*ClearLogResponse, error)

	// The EvtRpcExportLog (Opnum 7) method instructs the server to create a backup event
	// log at a specified file name.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	//
	// The server does not validate the control handle passed to EvtRpcExportLog, and it
	// SHOULD assume that this parameter is always valid when the method is invoked.
	ExportLog(context.Context, *ExportLogRequest) (*ExportLogResponse, error)

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
	LocalizeExportLog(context.Context, *LocalizeExportLogRequest) (*LocalizeExportLogResponse, error)

	// The EvtRpcMessageRender (Opnum 9) method is used by a client to get localized descriptive
	// strings for an event.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_INSUFFICIENT_BUFFER (0x0000007A) if maxSizeString is too
	// small to hold the result string. In that case, neededSizeString MUST be set to the
	// necessary size. Otherwise, the method MUST return a different implementation-specific
	// nonzero value as specified in [MS-ERREF].
	MessageRender(context.Context, *MessageRenderRequest) (*MessageRenderResponse, error)

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
	MessageRenderDefault(context.Context, *MessageRenderDefaultRequest) (*MessageRenderDefaultResponse, error)

	// The EvtRpcQueryNext (Opnum 11) method is used by a client to get the next batch of
	// records from a query result set.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_TIMEOUT (0x000005bf) if no records are found within the
	// time-out period. The method MUST return ERROR_NO_MORE_ITEMS (0x00000103) once the
	// query has finished going through all the log(s); otherwise, it MUST return a different
	// implementation-specific nonzero value as specified in [MS-ERREF].
	QueryNext(context.Context, *QueryNextRequest) (*QueryNextResponse, error)

	// The EvtRpcQuerySeek (Opnum 12) method is used by a client to move a query cursor
	// within a result set.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	QuerySeek(context.Context, *QuerySeekRequest) (*QuerySeekResponse, error)

	// The EvtRpcClose (Opnum 13) method is used by a client to close context handles that
	// are opened by other methods in this protocol.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

	// The EvtRpcCancel (Opnum 14) method is used by a client to cancel another method.
	// This can be used to terminate long-running methods gracefully. Methods that can be
	// canceled include the subscription and query functions, and other functions that take
	// a CONTEXT_HANDLE_OPERATION_CONTROL argument.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	Cancel(context.Context, *CancelRequest) (*CancelResponse, error)

	// The EvtRpcAssertConfig (Opnum 15) method indicates to the server that the publisher
	// or channel configuration has been updated.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	AssertConfig(context.Context, *AssertConfigRequest) (*AssertConfigResponse, error)

	// The EvtRpcRetractConfig (Opnum 16) method indicates to the server that the publisher
	// or channel is to be removed.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	RetractConfig(context.Context, *RetractConfigRequest) (*RetractConfigResponse, error)

	// The EvtRpcOpenLogHandle (Opnum 17) method is used by a client to get information
	// about a channel or a backup event log.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	OpenLog(context.Context, *OpenLogRequest) (*OpenLogResponse, error)

	// The EvtRpcGetLogFileInfo (Opnum 18) method is used by a client to get information
	// about a live channel or a backup event log.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success. The
	// method MUST return ERROR_INSUFFICIENT_BUFFER (0x0000007A) if the buffer is too small;
	// otherwise, it MUST return a different implementation-specific nonzero value as specified
	// in [MS-ERREF].
	GetLogFileInfo(context.Context, *GetLogFileInfoRequest) (*GetLogFileInfoResponse, error)

	// The EvtRpcGetChannelList (Opnum 19) method is used to enumerate the set of available
	// channels.
	//
	// Return Values:  The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetChannelList(context.Context, *GetChannelListRequest) (*GetChannelListResponse, error)

	// The EvtRpcGetChannelConfig (opnum 20) method is used by a client to get the configuration
	// for a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetChannelConfig(context.Context, *GetChannelConfigRequest) (*GetChannelConfigResponse, error)

	// The EvtRpcPutChannelConfig (Opnum 21) method is used by a client to update the configuration
	// for a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].<42>
	PutChannelConfig(context.Context, *PutChannelConfigRequest) (*PutChannelConfigResponse, error)

	// The EvtRpcGetPublisherList (Opnum 22) method is used by a client to get the list
	// of publishers.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherList(context.Context, *GetPublisherListRequest) (*GetPublisherListResponse, error)

	// The EvtRpcGetPublisherListForChannel (Opnum 23) method is used by a client to get
	// the list of publishers that write events to a particular channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherListForChannel(context.Context, *GetPublisherListForChannelRequest) (*GetPublisherListForChannelResponse, error)

	// The EvtRpcGetPublisherMetadata (Opnum 24) method is used by a client to open a handle
	// to publisher metadata. It also gets some initial information from the metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherMetadata(context.Context, *GetPublisherMetadataRequest) (*GetPublisherMetadataResponse, error)

	// The EvtRpcGetPublisherResourceMetadata (Opnum 25) method obtains information from
	// the publisher metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetPublisherResourceMetadata(context.Context, *GetPublisherResourceMetadataRequest) (*GetPublisherResourceMetadataResponse, error)

	// The EvtRpcGetEventMetadataEnum (Opnum 26) method obtains a handle for enumerating
	// a publisher's event metadata.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetEventMetadataEnum(context.Context, *GetEventMetadataEnumRequest) (*GetEventMetadataEnumResponse, error)

	// The EvtRpcGetNextEventMetadata (Opnum 27) method gets details about a possible event
	// and also returns the next event metadata in the enumeration. It is used to enumerate
	// through the event definitions for the publisher associated with the handle. The enumeration
	// is in the forward direction only, and there is no reset functionality.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetNextEventMetadata(context.Context, *GetNextEventMetadataRequest) (*GetNextEventMetadataResponse, error)

	// The EvtRpcGetClassicLogDisplayName (Opnum 28) method obtains a descriptive name for
	// a channel.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific nonzero value as specified in [MS-ERREF].
	GetClassicLogDisplayName(context.Context, *GetClassicLogDisplayNameRequest) (*GetClassicLogDisplayNameResponse, error)
}

func RegisterEventServiceServer(conn dcerpc.Conn, o EventServiceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEventServiceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EventServiceSyntaxV1_0))...)
}

func NewEventServiceServerHandle(o EventServiceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EventServiceServerHandle(ctx, o, opNum, r)
	}
}

func EventServiceServerHandle(ctx context.Context, o EventServiceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // EvtRpcRegisterRemoteSubscription
		in := &RegisterRemoteSubscriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterRemoteSubscription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // EvtRpcRemoteSubscriptionNextAsync
		in := &RemoteSubscriptionNextAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteSubscriptionNextAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // EvtRpcRemoteSubscriptionNext
		in := &RemoteSubscriptionNextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteSubscriptionNext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // EvtRpcRemoteSubscriptionWaitAsync
		in := &RemoteSubscriptionWaitAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteSubscriptionWaitAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // EvtRpcRegisterControllableOperation
		in := &RegisterControllableOperationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterControllableOperation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // EvtRpcRegisterLogQuery
		in := &RegisterLogQueryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegisterLogQuery(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // EvtRpcClearLog
		in := &ClearLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // EvtRpcExportLog
		in := &ExportLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExportLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // EvtRpcLocalizeExportLog
		in := &LocalizeExportLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LocalizeExportLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // EvtRpcMessageRender
		in := &MessageRenderRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MessageRender(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // EvtRpcMessageRenderDefault
		in := &MessageRenderDefaultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MessageRenderDefault(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // EvtRpcQueryNext
		in := &QueryNextRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryNext(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // EvtRpcQuerySeek
		in := &QuerySeekRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QuerySeek(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // EvtRpcClose
		in := &CloseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Close(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // EvtRpcCancel
		in := &CancelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Cancel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // EvtRpcAssertConfig
		in := &AssertConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AssertConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // EvtRpcRetractConfig
		in := &RetractConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetractConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // EvtRpcOpenLogHandle
		in := &OpenLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // EvtRpcGetLogFileInfo
		in := &GetLogFileInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogFileInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // EvtRpcGetChannelList
		in := &GetChannelListRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetChannelList(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // EvtRpcGetChannelConfig
		in := &GetChannelConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetChannelConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // EvtRpcPutChannelConfig
		in := &PutChannelConfigRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PutChannelConfig(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // EvtRpcGetPublisherList
		in := &GetPublisherListRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherList(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // EvtRpcGetPublisherListForChannel
		in := &GetPublisherListForChannelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherListForChannel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // EvtRpcGetPublisherMetadata
		in := &GetPublisherMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // EvtRpcGetPublisherResourceMetadata
		in := &GetPublisherResourceMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPublisherResourceMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // EvtRpcGetEventMetadataEnum
		in := &GetEventMetadataEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventMetadataEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // EvtRpcGetNextEventMetadata
		in := &GetNextEventMetadataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNextEventMetadata(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // EvtRpcGetClassicLogDisplayName
		in := &GetClassicLogDisplayNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassicLogDisplayName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
