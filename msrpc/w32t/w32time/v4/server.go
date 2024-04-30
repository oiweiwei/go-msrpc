package w32time

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

// W32Time server interface.
type W32TimeServer interface {

	// The W32TimeSync method is invoked to request that the time service immediately initiate
	// an attempt to synchronize its time. The MIDL syntax of this method is specified as
	// follows.
	//
	// Return Values: If the TimeSyncFlag_ReturnResult flag is specified, the return value
	// MUST be one of the following specific TimeSync_ReturnResult codes. Otherwise, this
	// method MUST return zero on success or an implementation-specific nonzero error code
	// on failure.<36>
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ResyncResult_Success      | Synchronization between the time provider and the caller was successful. For     |
	//	|                                      | asynchronous calls, this result does not guarantee that the server has acquired  |
	//	|                                      | a new time sample. It merely states that the synchronization attempt has been    |
	//	|                                      | successfully initiated.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 ResyncResult_NoData       | The time service could not obtain a new time sample from the time provider.      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ResyncResult_StaleData    | The time service received data that was time stamped earlier than the last good  |
	//	|                                      | sample.                                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 ResyncResult_ChangeTooBig | The time service received data in which the time difference from the local clock |
	//	|                                      | was too large to trust.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 ResyncResult_Shutdown     | The time service was shutting down.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)

	// The W32TimeGetNetlogonServiceBits method returns information about the functionality
	// that the time service provides. The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method returns an unsigned 32-bit field that contains information
	// about the functionality that the time service provides. Multiple bits can be set
	// in the return value. Any bits not defined as follows MUST be set to zero by servers
	// and ignored by clients.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000040 DS_TIMESERV_FLAG      | The time service provides a time source with which clients can synchronize using |
	//	|                                  | NTP, as specified in [RFC1305].                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000200 DS_GOOD_TIMESERV_FLAG | The time service provides a reliable time source with which clients can          |
	//	|                                  | synchronize using NTP, as specified in [RFC1305].                                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetNetlogonServiceBits(context.Context, *GetNetlogonServiceBitsRequest) (*GetNetlogonServiceBitsResponse, error)

	// The W32TimeQueryProviderStatus method returns operational information for a specified
	// time provider (either an NTP or a hardware time provider) within the time service's
	// list of time providers.<39> The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return zero on success; on failure, it returns a
	// nonzero error code. The values transmitted in this field are implementation specific.
	// All nonzero values MUST be treated as equivalent for protocol purposes.<41>
	QueryProviderStatus(context.Context, *QueryProviderStatusRequest) (*QueryProviderStatusResponse, error)

	// The W32TimeQuerySource method returns the current time source of the time service.
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<42>
	QuerySource(context.Context, *QuerySourceRequest) (*QuerySourceResponse, error)

	// The W32TimeQueryProviderConfiguration method returns configuration data for a specific
	// time provider within the time service's list of time providers.<44> The MIDL syntax
	// of this method is specified as follows.
	//
	// Return Values: This method MUST return zero on success; on failure, it returns a
	// nonzero error code. The values transmitted in this field are implementation specific.
	// All nonzero values MUST be treated as equivalent for the purposes of this protocol.<46>
	QueryProviderConfiguration(context.Context, *QueryProviderConfigurationRequest) (*QueryProviderConfigurationResponse, error)

	// The W32TimeQueryConfiguration method returns the configuration data of the time service.<47>
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<48>
	QueryConfiguration(context.Context, *QueryConfigurationRequest) (*QueryConfigurationResponse, error)

	// The W32TimeQueryStatus method returns the service status data of the time service.<49>
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<50>
	QueryStatus(context.Context, *QueryStatusRequest) (*QueryStatusResponse, error)

	// The W32TimeLog method is invoked to request that the time service update its logging
	// configuration.<51> The logging of the time service is implementation specific.<52>
	//
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for protocol purposes.<53>
	Log(context.Context, *LogRequest) (*LogResponse, error)
}

func RegisterW32TimeServer(conn dcerpc.Conn, o W32TimeServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewW32TimeServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(W32TimeSyntaxV4_1))...)
}

func NewW32TimeServerHandle(o W32TimeServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return W32TimeServerHandle(ctx, o, opNum, r)
	}
}

func W32TimeServerHandle(ctx context.Context, o W32TimeServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // W32TimeSync
		in := &SyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Sync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // W32TimeGetNetlogonServiceBits
		in := &GetNetlogonServiceBitsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetlogonServiceBits(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // W32TimeQueryProviderStatus
		in := &QueryProviderStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryProviderStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // W32TimeQuerySource
		in := &QuerySourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QuerySource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // W32TimeQueryProviderConfiguration
		in := &QueryProviderConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryProviderConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // W32TimeQueryConfiguration
		in := &QueryConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // W32TimeQueryStatus
		in := &QueryStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // W32TimeLog
		in := &LogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Log(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
