package perflibv2

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

// PerflibV2 server interface.
type PerflibV2Server interface {

	// The PerflibV2EnumerateCounterSet method allows a client to enumerate the available
	// countersets on a server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The return value indicates success.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The server returns this value to the client if the authentication level of the   |
	//	|                                    | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | This return value is used to indicate when the size of the client-provided       |
	//	|                                    | buffer is not large enough to accommodate all of the GUID values that are being  |
	//	|                                    | returned by the server.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | This return value is used to indicate that the server, while attempting to       |
	//	|                                    | return all of the appropriate GUIDs to the client, could not allocate memory.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	EnumerateCounterSet(context.Context, *EnumerateCounterSetRequest) (*EnumerateCounterSetResponse, error)

	// The PerflibV2QueryCounterSetRegistrationInfo method allows a client to enumerate
	// metadata about a counterset or performance counter on a server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The return value indicates success.                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED        | The server returns this value to the client if the authentication level of the   |
	//	|                                       | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | This return value indicates that there was a problem with the parameter that     |
	//	|                                       | was passed by the client to the server. The server MUST return this value        |
	//	|                                       | when: RequestCode (the RequestCode is not between 0x00000001 and 0x0000000A      |
	//	|                                       | inclusive).                                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001068 ERROR_WMI_GUID_NOT_FOUND   | The server returns this value if it does not have a counterset with the same     |
	//	|                                       | GUID as the one passed by the client through the CounterSetGuid parameter of the |
	//	|                                       | method. The server will also return this value if it cannot find the GUID of the |
	//	|                                       | provider to which the counterset belongs.                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY    | The server will return this value to the client if the RequestCode parameter is  |
	//	|                                       | valid, but the buffer pointed to by lpData is not of sufficient size.            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000106A ERROR_WMI_ITEMID_NOT_FOUND | The server returns this error code when the value of RequestCode is 0x02 and a   |
	//	|                                       | counterset with the GUID provided through the CounterSetGuid parameter exists,   |
	//	|                                       | but the counter identifier is not found in the counterset.                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// The data that this method returns depends on the type of information that is requested,
	// as denoted by the RequestCode parameter.
	//
	// * If the value of RequestCode is 0x00000003, 0x00000004, 0x00000005, or 0x00000006,
	// and the language specified by *RequestLCID* is not installed on the server, an error
	// MUST be returned.
	//
	// * If RequestCode = 0x00000001, the server returns information about the counterset.
	// The server MUST return a _PERF_COUNTERSET_REG_INFO ( a2745a4d-3428-4804-868d-b0208e469fa2
	// ) structure that is followed by a set of _PERF_COUNTER_REG_INFO ( a89712bd-4ba1-4dab-8dbb-dea712c237ba
	// ) structures. The number of _PERF_COUNTER_REG_INFO structures MUST be equal to the
	// *NumCounters* field of the PERF_COUNTERSET_REG_INFO structure.
	//
	// [PerflibV2QueryCounterSetRegistrationInfo return if RequestCode = 0x00000001](ms-pcq_files/image001.png)
	QueryCounterSetRegistrationInfo(context.Context, *QueryCounterSetRegistrationInfoRequest) (*QueryCounterSetRegistrationInfoResponse, error)

	// The PerflibV2EnumerateCounterSetInstances method retrieves all active instances of
	// the client-specified counterset on the server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	EnumerateCounterSetInstances(context.Context, *EnumerateCounterSetInstancesRequest) (*EnumerateCounterSetInstancesResponse, error)

	// The PerflibV2OpenQueryHandle method returns a handle to the client that the client
	// then uses to add, remove, and collect performance counters from the server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The return value indicates success.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | The server returns this value to the client if the authentication level of the   |
	//	|                                      | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY         | The server returns this value to the client if for any reason memory allocation  |
	//	|                                      | fails as it tries to allocate memory to begin storing state about the client     |
	//	|                                      | request.                                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | The server returns this value if it cannot allocate other system resource to     |
	//	|                                      | process the client request. This is not specifically memory about the client     |
	//	|                                      | request or handle.                                                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	OpenQueryHandle(context.Context, *OpenQueryHandleRequest) (*OpenQueryHandleResponse, error)

	// The PerflibV2CloseQueryHandle method closes the handle that is returned from the
	// PerflibV2OpenQueryHandle method.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The return value indicates success.                                              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The server returns this value to the client if the authentication level of the   |
	//	|                                | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY. The opened handle, phQuery,   |
	//	|                                | remains in that state until the client calls PerflibV2CloseQueryHandle with      |
	//	|                                | authentication level RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	CloseQueryHandle(context.Context, *CloseQueryHandleRequest) (*CloseQueryHandleResponse, error)

	// The PerflibV2QueryCounterInfo method returns information on the performance counters
	// that belong to the performance counter query associated with the RPC_HQUERY; these
	// performance counters are associated with RPC_HQUERY by calling the PerflibV2ValidateCounters
	// method. The server MUST return performance counter metadata information, stored in
	// a _PERF_COUNTER_IDENTIFIER structure for each performance counter, for the performance
	// counters that are associated with the RPC_HQUERY handle.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The return value indicates success.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The server returns this value to the client if the authentication level of the   |
	//	|                                    | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | The server will return this value if the buffer pointed to by lpData is not of   |
	//	|                                    | sufficient size to return the requested information back to the client.          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	QueryCounterInfo(context.Context, *QueryCounterInfoRequest) (*QueryCounterInfoResponse, error)

	// The PerflibV2QueryCounterData method retrieves data for the performance counters
	// associated with the query. Performance counters can be added or removed from queries
	// by calling PerflibV2ValidateCounters.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows error codes, as specified in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The return value indicates success.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The server returns this value to the client if the authentication level of the   |
	//	|                                    | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | The server will return this value to the client if the size of the buffer        |
	//	|                                    | pointed to by lpData is not of sufficient size to return the performance counter |
	//	|                                    | values to the client.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	QueryCounterData(context.Context, *QueryCounterDataRequest) (*QueryCounterDataResponse, error)

	// This PerflibV2ValidateCounters method either adds or removes performance counters
	// from the query.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows error codes, as specified in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The return value indicates success.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The server returns this value to the client if the authentication level of the   |
	//	|                                    | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The server returns this value to the client for any of the following reasons:    |
	//	|                                    | dwSize is less than the size of the _PERF_COUNTER_IDENTIFIER structure (this     |
	//	|                                    | condition would prevent the server from returning information about one          |
	//	|                                    | counter). The size of a single _PERF_COUNTER_IDENTIFIER structure that is        |
	//	|                                    | passed into the buffer by the client is smaller than the expected size of a      |
	//	|                                    | _PERF_COUNTER_IDENTIFIER structure.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | The server will return this value to the client if, in the process of completing |
	//	|                                    | the client's request of adding or removing performance counters from the query,  |
	//	|                                    | a memory allocation fails.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Errors are returned to the client by the server in one of two ways: the first is
	// if the performance counter infrastructure on the server could not add or remove performance
	// counters from the query; the second is if the provider that is exposing the performance
	// counter returns an error, in which case the performance counter infrastructure passes
	// the error back to the client.
	//
	// When the PerflibV2ValidateCounters method returns, the Status field of each _PERF_COUNTER_IDENTIFIER
	// sent to the server will have the result of whether or not the server was able to
	// successfully add or remove that particular performance counter from the query that
	// is identified by the handle hQuery.
	//
	// If the performance counter infrastructure is setting the Status field to an error
	// value, then it MUST be one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_STATUS                | The return value indicates success. The counter was either successfully added or |
	//	|                                        | removed from the query.                                                          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001068 ERROR_WMI_GUID_NOT_FOUND    | The server cannot find the GUID that was passed by the client in the             |
	//	|                                        | CounterSetGuid field of the _PERF_COUNTER_IDENTIFIER structure.                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000106A ERROR_WMI_ITEMID_NOT_ FOUND | The server cannot find the counter whose numeric identifier is in the CounterId  |
	//	|                                        | field of the _PERF_COUNTER_IDENTIFIER structure.                                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 ERROR_PATH_NOT_FOUND        | The server cannot find an active instance with the name that was placed after    |
	//	|                                        | the _PERF_COUNTER_IDENTIFIER structure.                                          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS        | The client tried to add a performance counter that has already been added in a   |
	//	|                                        | previous call to PerflibV2ValidateCounters.                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The server will return this value in the Status field of the                     |
	//	|                                        | _PERF_COUNTER_IDENTIFIER either when the _PERF_COUNTER_IDENTIFIER is corrupt, or |
	//	|                                        | if the server cannot find the counter to delete from the query that is specified |
	//	|                                        | by the structure.                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY           | The server will return this value to the client if, either in the process of     |
	//	|                                        | adding or removing a counter from a query, a memory allocation failure occurred. |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	ValidateCounters(context.Context, *ValidateCountersRequest) (*ValidateCountersResponse, error)
}

func RegisterPerflibV2Server(conn dcerpc.Conn, o PerflibV2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPerflibV2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PerflibV2SyntaxV1_0))...)
}

func NewPerflibV2ServerHandle(o PerflibV2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PerflibV2ServerHandle(ctx, o, opNum, r)
	}
}

func PerflibV2ServerHandle(ctx context.Context, o PerflibV2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // PerflibV2EnumerateCounterSet
		op := &xxx_EnumerateCounterSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateCounterSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateCounterSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // PerflibV2QueryCounterSetRegistrationInfo
		op := &xxx_QueryCounterSetRegistrationInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryCounterSetRegistrationInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryCounterSetRegistrationInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // PerflibV2EnumerateCounterSetInstances
		op := &xxx_EnumerateCounterSetInstancesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateCounterSetInstancesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateCounterSetInstances(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // PerflibV2OpenQueryHandle
		op := &xxx_OpenQueryHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenQueryHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenQueryHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // PerflibV2CloseQueryHandle
		op := &xxx_CloseQueryHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseQueryHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseQueryHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // PerflibV2QueryCounterInfo
		op := &xxx_QueryCounterInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryCounterInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryCounterInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // PerflibV2QueryCounterData
		op := &xxx_QueryCounterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryCounterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryCounterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // PerflibV2ValidateCounters
		op := &xxx_ValidateCountersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ValidateCountersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ValidateCounters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented PerflibV2
type UnimplementedPerflibV2Server struct {
}

func (UnimplementedPerflibV2Server) EnumerateCounterSet(context.Context, *EnumerateCounterSetRequest) (*EnumerateCounterSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) QueryCounterSetRegistrationInfo(context.Context, *QueryCounterSetRegistrationInfoRequest) (*QueryCounterSetRegistrationInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) EnumerateCounterSetInstances(context.Context, *EnumerateCounterSetInstancesRequest) (*EnumerateCounterSetInstancesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) OpenQueryHandle(context.Context, *OpenQueryHandleRequest) (*OpenQueryHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) CloseQueryHandle(context.Context, *CloseQueryHandleRequest) (*CloseQueryHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) QueryCounterInfo(context.Context, *QueryCounterInfoRequest) (*QueryCounterInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) QueryCounterData(context.Context, *QueryCounterDataRequest) (*QueryCounterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPerflibV2Server) ValidateCounters(context.Context, *ValidateCountersRequest) (*ValidateCountersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PerflibV2Server = (*UnimplementedPerflibV2Server)(nil)
