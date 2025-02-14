package perflibv2

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
	GoPackage = "pcq"
)

var (
	// Syntax UUID
	PerflibV2SyntaxUUID = &uuid.UUID{TimeLow: 0xda5a86c5, TimeMid: 0x12c2, TimeHiAndVersion: 0x4943, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x30, Node: [6]uint8{0x7f, 0x74, 0xa8, 0x13, 0xd8, 0x53}}
	// Syntax ID
	PerflibV2SyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: PerflibV2SyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// PerflibV2 interface.
type PerflibV2Client interface {

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
	EnumerateCounterSet(context.Context, *EnumerateCounterSetRequest, ...dcerpc.CallOption) (*EnumerateCounterSetResponse, error)

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
	QueryCounterSetRegistrationInfo(context.Context, *QueryCounterSetRegistrationInfoRequest, ...dcerpc.CallOption) (*QueryCounterSetRegistrationInfoResponse, error)

	// The PerflibV2EnumerateCounterSetInstances method retrieves all active instances of
	// the client-specified counterset on the server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the standard Windows errors, as specified in [MS-ERREF] section
	// 2.2.
	EnumerateCounterSetInstances(context.Context, *EnumerateCounterSetInstancesRequest, ...dcerpc.CallOption) (*EnumerateCounterSetInstancesResponse, error)

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
	OpenQueryHandle(context.Context, *OpenQueryHandleRequest, ...dcerpc.CallOption) (*OpenQueryHandleResponse, error)

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
	CloseQueryHandle(context.Context, *CloseQueryHandleRequest, ...dcerpc.CallOption) (*CloseQueryHandleResponse, error)

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
	QueryCounterInfo(context.Context, *QueryCounterInfoRequest, ...dcerpc.CallOption) (*QueryCounterInfoResponse, error)

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
	QueryCounterData(context.Context, *QueryCounterDataRequest, ...dcerpc.CallOption) (*QueryCounterDataResponse, error)

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
	ValidateCounters(context.Context, *ValidateCountersRequest, ...dcerpc.CallOption) (*ValidateCountersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Query structure represents RPC_HQUERY RPC structure.
type Query dcetypes.ContextHandle

func (o *Query) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Query) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Query) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Query) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultPerflibV2Client struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultPerflibV2Client) EnumerateCounterSet(ctx context.Context, in *EnumerateCounterSetRequest, opts ...dcerpc.CallOption) (*EnumerateCounterSetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateCounterSetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) QueryCounterSetRegistrationInfo(ctx context.Context, in *QueryCounterSetRegistrationInfoRequest, opts ...dcerpc.CallOption) (*QueryCounterSetRegistrationInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryCounterSetRegistrationInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) EnumerateCounterSetInstances(ctx context.Context, in *EnumerateCounterSetInstancesRequest, opts ...dcerpc.CallOption) (*EnumerateCounterSetInstancesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateCounterSetInstancesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) OpenQueryHandle(ctx context.Context, in *OpenQueryHandleRequest, opts ...dcerpc.CallOption) (*OpenQueryHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenQueryHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) CloseQueryHandle(ctx context.Context, in *CloseQueryHandleRequest, opts ...dcerpc.CallOption) (*CloseQueryHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseQueryHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) QueryCounterInfo(ctx context.Context, in *QueryCounterInfoRequest, opts ...dcerpc.CallOption) (*QueryCounterInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryCounterInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) QueryCounterData(ctx context.Context, in *QueryCounterDataRequest, opts ...dcerpc.CallOption) (*QueryCounterDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryCounterDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) ValidateCounters(ctx context.Context, in *ValidateCountersRequest, opts ...dcerpc.CallOption) (*ValidateCountersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ValidateCountersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPerflibV2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPerflibV2Client) Conn() dcerpc.Conn {
	return o.cc
}

func NewPerflibV2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PerflibV2Client, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PerflibV2SyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultPerflibV2Client{cc: cc}, nil
}

// xxx_EnumerateCounterSetOperation structure represents the PerflibV2EnumerateCounterSet operation
type xxx_EnumerateCounterSetOperation struct {
	Machine    string       `idl:"name:szMachine;string" json:"machine"`
	InSize     uint32       `idl:"name:dwInSize" json:"in_size"`
	OutSize    uint32       `idl:"name:pdwOutSize" json:"out_size"`
	ReturnSize uint32       `idl:"name:pdwRtnSize" json:"return_size"`
	Data       []*dtyp.GUID `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	Return     uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateCounterSetOperation) OpNum() int { return 0 }

func (o *xxx_EnumerateCounterSetOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2EnumerateCounterSet"
}

func (o *xxx_EnumerateCounterSetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InSize > uint32(256) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Machine); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Machine); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,256), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwInSize,length_is=pdwOutSize](struct))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Data[i1] != nil {
				if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_EnumerateCounterSetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1))(2:{alias=GUID}[dim:0,size_is=dwInSize,length_is=pdwOutSize](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]*dtyp.GUID, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if o.Data[i1] == nil {
				o.Data[i1] = &dtyp.GUID{}
			}
			if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
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

// EnumerateCounterSetRequest structure represents the PerflibV2EnumerateCounterSet operation request
type EnumerateCounterSetRequest struct {
	// szMachine: A Unicode string specifying a server name, which is passed directly to
	// the counter provider. Counter providers can ignore the server name provided by szMachine.
	Machine string `idl:"name:szMachine;string" json:"machine"`
	// dwInSize: The size of the buffer, in number of GUIDs.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
}

func (o *EnumerateCounterSetRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumerateCounterSetOperation) *xxx_EnumerateCounterSetOperation {
	if op == nil {
		op = &xxx_EnumerateCounterSetOperation{}
	}
	if o == nil {
		return op
	}
	o.Machine = op.Machine
	o.InSize = op.InSize
	return op
}

func (o *EnumerateCounterSetRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateCounterSetOperation) {
	if o == nil {
		return
	}
	o.Machine = op.Machine
	o.InSize = op.InSize
}
func (o *EnumerateCounterSetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateCounterSetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateCounterSetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateCounterSetResponse structure represents the PerflibV2EnumerateCounterSet operation response
type EnumerateCounterSetResponse struct {
	// pdwOutSize: On output, the number of GUIDs that are returned in the array. The server
	// MUST set this value to zero if the value of dwInSize is less than the total number
	// of GUIDs on the server.
	OutSize uint32 `idl:"name:pdwOutSize" json:"out_size"`
	// pdwRtnSize: On output, the total number of GUIDs on the server.
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	// lpData: The buffer that returns an array of GUIDs.
	Data []*dtyp.GUID `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	// Return: The PerflibV2EnumerateCounterSet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateCounterSetResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumerateCounterSetOperation) *xxx_EnumerateCounterSetOperation {
	if op == nil {
		op = &xxx_EnumerateCounterSetOperation{}
	}
	if o == nil {
		return op
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *EnumerateCounterSetResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateCounterSetOperation) {
	if o == nil {
		return
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
}
func (o *EnumerateCounterSetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateCounterSetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateCounterSetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryCounterSetRegistrationInfoOperation structure represents the PerflibV2QueryCounterSetRegistrationInfo operation
type xxx_QueryCounterSetRegistrationInfoOperation struct {
	Machine         string     `idl:"name:szMachine;string" json:"machine"`
	CounterSetGUID  *dtyp.GUID `idl:"name:CounterSetGuid" json:"counter_set_guid"`
	RequestCode     uint32     `idl:"name:RequestCode" json:"request_code"`
	RequestLocaleID uint32     `idl:"name:RequestLCID" json:"request_locale_id"`
	InSize          uint32     `idl:"name:dwInSize" json:"in_size"`
	OutSize         uint32     `idl:"name:pdwOutSize" json:"out_size"`
	ReturnSize      uint32     `idl:"name:pdwRtnSize" json:"return_size"`
	Data            []byte     `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	Return          uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) OpNum() int { return 1 }

func (o *xxx_QueryCounterSetRegistrationInfoOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2QueryCounterSetRegistrationInfo"
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InSize > uint32(134217728) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Machine); err != nil {
			return err
		}
	}
	// CounterSetGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CounterSetGUID != nil {
			if err := o.CounterSetGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RequestCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestCode); err != nil {
			return err
		}
	}
	// RequestLCID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequestLocaleID); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,134217728), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Machine); err != nil {
			return err
		}
	}
	// CounterSetGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CounterSetGUID == nil {
			o.CounterSetGUID = &dtyp.GUID{}
		}
		if err := o.CounterSetGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RequestCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestCode); err != nil {
			return err
		}
	}
	// RequestLCID {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequestLocaleID); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,134217728), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterSetRegistrationInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
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

func (o *xxx_QueryCounterSetRegistrationInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
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

// QueryCounterSetRegistrationInfoRequest structure represents the PerflibV2QueryCounterSetRegistrationInfo operation request
type QueryCounterSetRegistrationInfoRequest struct {
	// szMachine: A Unicode string specifying a server name, which is passed directly to
	// the counter providers. Counter providers can ignore the server name provided by szMachine.
	Machine string `idl:"name:szMachine;string" json:"machine"`
	// CounterSetGuid: The GUID of the counterset whose information needs to be retrieved;
	// this can also be the GUID of the counterset to which the performance counters whose
	// information is being queried belong.
	CounterSetGUID *dtyp.GUID `idl:"name:CounterSetGuid" json:"counter_set_guid"`
	// RequestCode: The type of information on the counterset to retrieve. The value MUST
	// be one of the following.
	//
	//	+------------+----------------------------------------------------------------+
	//	|            |                                                                |
	//	|   VALUE    |                            MEANING                             |
	//	|            |                                                                |
	//	+------------+----------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000001 | Return information about the counterset.                       |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000002 | Return information about a performance counter.                |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000003 | Return the name of the counterset.                             |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000004 | Return the description of the counterset.                      |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000005 | Return the names of the performance counters.                  |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000006 | Return the descriptions of the performance counters.           |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000007 | Return the name of the provider.                               |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000008 | Return the GUID of the provider.                               |
	//	+------------+----------------------------------------------------------------+
	//	| 0x00000009 | Return the English-language name of the counterset.            |
	//	+------------+----------------------------------------------------------------+
	//	| 0x0000000A | Return the English-language names of the performance counters. |
	//	+------------+----------------------------------------------------------------+
	RequestCode uint32 `idl:"name:RequestCode" json:"request_code"`
	// RequestLCID: When the value of RequestCode is 0x00000003, 0x00000004, 0x00000005,
	// or 0x00000006, RequestLCID specifies the locale ID (as specified in [MS-LCID]), or
	// is set to 0 to instruct the server to use its default language.
	//
	// When the value of RequestCode is 0x00000002, RequestLCID specifies the counter ID.
	RequestLocaleID uint32 `idl:"name:RequestLCID" json:"request_locale_id"`
	// dwInSize: The size, in bytes, of the buffer.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
}

func (o *QueryCounterSetRegistrationInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterSetRegistrationInfoOperation) *xxx_QueryCounterSetRegistrationInfoOperation {
	if op == nil {
		op = &xxx_QueryCounterSetRegistrationInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Machine = op.Machine
	o.CounterSetGUID = op.CounterSetGUID
	o.RequestCode = op.RequestCode
	o.RequestLocaleID = op.RequestLocaleID
	o.InSize = op.InSize
	return op
}

func (o *QueryCounterSetRegistrationInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterSetRegistrationInfoOperation) {
	if o == nil {
		return
	}
	o.Machine = op.Machine
	o.CounterSetGUID = op.CounterSetGUID
	o.RequestCode = op.RequestCode
	o.RequestLocaleID = op.RequestLocaleID
	o.InSize = op.InSize
}
func (o *QueryCounterSetRegistrationInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryCounterSetRegistrationInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterSetRegistrationInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryCounterSetRegistrationInfoResponse structure represents the PerflibV2QueryCounterSetRegistrationInfo operation response
type QueryCounterSetRegistrationInfoResponse struct {
	// pdwOutSize: The size, in bytes, of the data in the buffer pointed to by lpData.
	OutSize uint32 `idl:"name:pdwOutSize" json:"out_size"`
	// pdwRtnSize: The necessary size, in bytes, to retrieve all the requested data.
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	// lpData: The buffer that returns the requested data.
	Data []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	// Return: The PerflibV2QueryCounterSetRegistrationInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryCounterSetRegistrationInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterSetRegistrationInfoOperation) *xxx_QueryCounterSetRegistrationInfoOperation {
	if op == nil {
		op = &xxx_QueryCounterSetRegistrationInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *QueryCounterSetRegistrationInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterSetRegistrationInfoOperation) {
	if o == nil {
		return
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
}
func (o *QueryCounterSetRegistrationInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryCounterSetRegistrationInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterSetRegistrationInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateCounterSetInstancesOperation structure represents the PerflibV2EnumerateCounterSetInstances operation
type xxx_EnumerateCounterSetInstancesOperation struct {
	Machine        string     `idl:"name:szMachine;string" json:"machine"`
	CounterSetGUID *dtyp.GUID `idl:"name:CounterSetGuid" json:"counter_set_guid"`
	InSize         uint32     `idl:"name:dwInSize" json:"in_size"`
	OutSize        uint32     `idl:"name:pdwOutSize" json:"out_size"`
	ReturnSize     uint32     `idl:"name:pdwRtnSize" json:"return_size"`
	Data           []byte     `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	Return         uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateCounterSetInstancesOperation) OpNum() int { return 2 }

func (o *xxx_EnumerateCounterSetInstancesOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2EnumerateCounterSetInstances"
}

func (o *xxx_EnumerateCounterSetInstancesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InSize > uint32(67108864) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetInstancesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Machine); err != nil {
			return err
		}
	}
	// CounterSetGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CounterSetGUID != nil {
			if err := o.CounterSetGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetInstancesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Machine); err != nil {
			return err
		}
	}
	// CounterSetGuid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CounterSetGUID == nil {
			o.CounterSetGUID = &dtyp.GUID{}
		}
		if err := o.CounterSetGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetInstancesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateCounterSetInstancesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
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

func (o *xxx_EnumerateCounterSetInstancesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
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

// EnumerateCounterSetInstancesRequest structure represents the PerflibV2EnumerateCounterSetInstances operation request
type EnumerateCounterSetInstancesRequest struct {
	// szMachine: A Unicode string specifying a server name, which is passed directly to
	// the counter providers. Counter providers can ignore the server name provided by szMachine.
	Machine string `idl:"name:szMachine;string" json:"machine"`
	// CounterSetGuid: The GUID of the counterset whose instances are to be enumerated.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                | The return value indicates success.                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED          | The server returns this value to the client if the authentication level of the   |
	//	|                                         | client is less than RPC_C_AUTHN_LEVEL_PKT_PRIVACY.                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001068 ERROR_WMI_GUID_NOT_FOUND     | The server returns this value when it cannot find a counterset with the GUID     |
	//	|                                         | that was specified by the client in the CounterSetGuid parameter.                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY      | The server returns this value to the client when the buffer the client has       |
	//	|                                         | provided is not large enough to accommodate the instance information.            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001069 ERROR_WMI_INSTANCE_NOT_FOUND | The server returns this value to the client when there are no active instances   |
	//	|                                         | of the counterset whose information can be returned.                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001073 ERROR_WMI_INVALID_REGINFO    | The server returns this to the client if, for any reason when trying to          |
	//	|                                         | enumerate counterset instances, the information that the server expected was     |
	//	|                                         | different than what the applications exposing performance counters returned.     |
	//	|                                         | For example, the server (through some standard repository), expected information |
	//	|                                         | about one instance of a counterset to be returned (because it was specified      |
	//	|                                         | as a single-instance counterset), but the application actually maintaining       |
	//	|                                         | the information returned instance information about multiple instances of the    |
	//	|                                         | counterset.                                                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY            | The server returns this value to the client if, for any reason as it tries       |
	//	|                                         | to return the instance information of the specified counterset, it fails to      |
	//	|                                         | allocate memory.                                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	CounterSetGUID *dtyp.GUID `idl:"name:CounterSetGuid" json:"counter_set_guid"`
	// dwInSize: The size, in bytes, of the buffer.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
}

func (o *EnumerateCounterSetInstancesRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumerateCounterSetInstancesOperation) *xxx_EnumerateCounterSetInstancesOperation {
	if op == nil {
		op = &xxx_EnumerateCounterSetInstancesOperation{}
	}
	if o == nil {
		return op
	}
	o.Machine = op.Machine
	o.CounterSetGUID = op.CounterSetGUID
	o.InSize = op.InSize
	return op
}

func (o *EnumerateCounterSetInstancesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateCounterSetInstancesOperation) {
	if o == nil {
		return
	}
	o.Machine = op.Machine
	o.CounterSetGUID = op.CounterSetGUID
	o.InSize = op.InSize
}
func (o *EnumerateCounterSetInstancesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateCounterSetInstancesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateCounterSetInstancesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateCounterSetInstancesResponse structure represents the PerflibV2EnumerateCounterSetInstances operation response
type EnumerateCounterSetInstancesResponse struct {
	// pdwOutSize: The total size, in bytes, of the data that is returned and written to
	// the buffer.
	OutSize uint32 `idl:"name:pdwOutSize" json:"out_size"`
	// pdwRtnSize: The necessary size, in bytes, to retrieve all the requested data.
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	// lpData: The buffer that contains the instances information for the counterset.
	Data []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	// Return: The PerflibV2EnumerateCounterSetInstances return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateCounterSetInstancesResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumerateCounterSetInstancesOperation) *xxx_EnumerateCounterSetInstancesOperation {
	if op == nil {
		op = &xxx_EnumerateCounterSetInstancesOperation{}
	}
	if o == nil {
		return op
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *EnumerateCounterSetInstancesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateCounterSetInstancesOperation) {
	if o == nil {
		return
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
}
func (o *EnumerateCounterSetInstancesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateCounterSetInstancesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateCounterSetInstancesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenQueryHandleOperation structure represents the PerflibV2OpenQueryHandle operation
type xxx_OpenQueryHandleOperation struct {
	Machine string `idl:"name:szMachine;string" json:"machine"`
	Query   *Query `idl:"name:phQuery" json:"query"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenQueryHandleOperation) OpNum() int { return 3 }

func (o *xxx_OpenQueryHandleOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2OpenQueryHandle"
}

func (o *xxx_OpenQueryHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueryHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Machine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueryHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Machine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueryHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenQueryHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phQuery {out} (1:{alias=PRPC_HQUERY}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenQueryHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phQuery {out} (1:{alias=PRPC_HQUERY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
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

// OpenQueryHandleRequest structure represents the PerflibV2OpenQueryHandle operation request
type OpenQueryHandleRequest struct {
	// szMachine: A Unicode string specifying a server name, which is passed directly to
	// the counter providers. Counter providers can ignore the server name provided by szMachine.
	Machine string `idl:"name:szMachine;string" json:"machine"`
}

func (o *OpenQueryHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenQueryHandleOperation) *xxx_OpenQueryHandleOperation {
	if op == nil {
		op = &xxx_OpenQueryHandleOperation{}
	}
	if o == nil {
		return op
	}
	o.Machine = op.Machine
	return op
}

func (o *OpenQueryHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenQueryHandleOperation) {
	if o == nil {
		return
	}
	o.Machine = op.Machine
}
func (o *OpenQueryHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenQueryHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueryHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenQueryHandleResponse structure represents the PerflibV2OpenQueryHandle operation response
type OpenQueryHandleResponse struct {
	// phQuery: A handle used by other methods to add, remove, and collect performance counters.
	Query *Query `idl:"name:phQuery" json:"query"`
	// Return: The PerflibV2OpenQueryHandle return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenQueryHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenQueryHandleOperation) *xxx_OpenQueryHandleOperation {
	if op == nil {
		op = &xxx_OpenQueryHandleOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	o.Return = op.Return
	return op
}

func (o *OpenQueryHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenQueryHandleOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
	o.Return = op.Return
}
func (o *OpenQueryHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenQueryHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenQueryHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseQueryHandleOperation structure represents the PerflibV2CloseQueryHandle operation
type xxx_CloseQueryHandleOperation struct {
	Query  *Query `idl:"name:phQuery" json:"query"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseQueryHandleOperation) OpNum() int { return 4 }

func (o *xxx_CloseQueryHandleOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2CloseQueryHandle"
}

func (o *xxx_CloseQueryHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueryHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phQuery {in, out} (1:{alias=PRPC_HQUERY}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseQueryHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phQuery {in, out} (1:{alias=PRPC_HQUERY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueryHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseQueryHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phQuery {in, out} (1:{alias=PRPC_HQUERY}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseQueryHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phQuery {in, out} (1:{alias=PRPC_HQUERY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
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

// CloseQueryHandleRequest structure represents the PerflibV2CloseQueryHandle operation request
type CloseQueryHandleRequest struct {
	// phQuery: A handle that is created by the PerflibV2OpenQueryHandle method. An exception
	// is thrown or an error is returned by RPC if the handle did not originate from the
	// PerflibV2OpenQueryHandle method. On method return, phQuery MUST be set to NULL.
	Query *Query `idl:"name:phQuery" json:"query"`
}

func (o *CloseQueryHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseQueryHandleOperation) *xxx_CloseQueryHandleOperation {
	if op == nil {
		op = &xxx_CloseQueryHandleOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	return op
}

func (o *CloseQueryHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseQueryHandleOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
}
func (o *CloseQueryHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseQueryHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseQueryHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseQueryHandleResponse structure represents the PerflibV2CloseQueryHandle operation response
type CloseQueryHandleResponse struct {
	// phQuery: A handle that is created by the PerflibV2OpenQueryHandle method. An exception
	// is thrown or an error is returned by RPC if the handle did not originate from the
	// PerflibV2OpenQueryHandle method. On method return, phQuery MUST be set to NULL.
	Query *Query `idl:"name:phQuery" json:"query"`
	// Return: The PerflibV2CloseQueryHandle return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseQueryHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseQueryHandleOperation) *xxx_CloseQueryHandleOperation {
	if op == nil {
		op = &xxx_CloseQueryHandleOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	o.Return = op.Return
	return op
}

func (o *CloseQueryHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseQueryHandleOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
	o.Return = op.Return
}
func (o *CloseQueryHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseQueryHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseQueryHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryCounterInfoOperation structure represents the PerflibV2QueryCounterInfo operation
type xxx_QueryCounterInfoOperation struct {
	Query      *Query `idl:"name:hQuery" json:"query"`
	InSize     uint32 `idl:"name:dwInSize" json:"in_size"`
	OutSize    uint32 `idl:"name:pdwOutSize" json:"out_size"`
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	Data       []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryCounterInfoOperation) OpNum() int { return 5 }

func (o *xxx_QueryCounterInfoOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2QueryCounterInfo"
}

func (o *xxx_QueryCounterInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InSize > uint32(67108864) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
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

func (o *xxx_QueryCounterInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
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

// QueryCounterInfoRequest structure represents the PerflibV2QueryCounterInfo operation request
type QueryCounterInfoRequest struct {
	// hQuery: The handle returned by the PerflibV2OpenQueryHandle method; an exception
	// is thrown or an error is returned by RPC if the handle did not originate from the
	// PerflibV2OpenQueryHandle method.
	Query *Query `idl:"name:hQuery" json:"query"`
	// dwInSize: The size, in bytes, of the buffer.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
}

func (o *QueryCounterInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterInfoOperation) *xxx_QueryCounterInfoOperation {
	if op == nil {
		op = &xxx_QueryCounterInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	o.InSize = op.InSize
	return op
}

func (o *QueryCounterInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterInfoOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
	o.InSize = op.InSize
}
func (o *QueryCounterInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryCounterInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryCounterInfoResponse structure represents the PerflibV2QueryCounterInfo operation response
type QueryCounterInfoResponse struct {
	// pdwOutSize: The size, in bytes, of the data that is written to the buffer.
	OutSize uint32 `idl:"name:pdwOutSize" json:"out_size"`
	// pdwRtnSize: The necessary size, in bytes, to retrieve all the requested data.
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	// lpData: The buffer that contains the requested counter information.
	Data []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	// Return: The PerflibV2QueryCounterInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryCounterInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterInfoOperation) *xxx_QueryCounterInfoOperation {
	if op == nil {
		op = &xxx_QueryCounterInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *QueryCounterInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterInfoOperation) {
	if o == nil {
		return
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
}
func (o *QueryCounterInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryCounterInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryCounterDataOperation structure represents the PerflibV2QueryCounterData operation
type xxx_QueryCounterDataOperation struct {
	Query      *Query `idl:"name:hQuery" json:"query"`
	InSize     uint32 `idl:"name:dwInSize" json:"in_size"`
	OutSize    uint32 `idl:"name:pdwOutSize" json:"out_size"`
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	Data       []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryCounterDataOperation) OpNum() int { return 6 }

func (o *xxx_QueryCounterDataOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2QueryCounterData"
}

func (o *xxx_QueryCounterDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InSize > uint32(1073741824) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwInSize {in} (1:{range=(0,1073741824), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,1073741824), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.OutSize == 0 {
		o.OutSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryCounterDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.OutSize)
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
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
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

func (o *xxx_QueryCounterDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwOutSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutSize); err != nil {
			return err
		}
	}
	// pdwRtnSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnSize); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize,length_is=pdwOutSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
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

// QueryCounterDataRequest structure represents the PerflibV2QueryCounterData operation request
type QueryCounterDataRequest struct {
	// hQuery: The handle returned by the PerflibV2OpenQueryHandle method; an exception
	// is thrown or an error is returned by RPC if the handle did not originate from the
	// PerflibV2OpenQueryHandle method.
	Query *Query `idl:"name:hQuery" json:"query"`
	// dwInSize: The size, in bytes, of the buffer.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
}

func (o *QueryCounterDataRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterDataOperation) *xxx_QueryCounterDataOperation {
	if op == nil {
		op = &xxx_QueryCounterDataOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	o.InSize = op.InSize
	return op
}

func (o *QueryCounterDataRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterDataOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
	o.InSize = op.InSize
}
func (o *QueryCounterDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryCounterDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryCounterDataResponse structure represents the PerflibV2QueryCounterData operation response
type QueryCounterDataResponse struct {
	// pdwOutSize: The size, in bytes, of the data that is returned and written to the buffer.
	OutSize uint32 `idl:"name:pdwOutSize" json:"out_size"`
	// pdwRtnSize: The necessary size, in bytes, to retrieve all the requested data.
	ReturnSize uint32 `idl:"name:pdwRtnSize" json:"return_size"`
	// lpData: The buffer that contains the requested counter information.
	Data []byte `idl:"name:lpData;size_is:(dwInSize);length_is:(pdwOutSize)" json:"data"`
	// Return: The PerflibV2QueryCounterData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryCounterDataResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryCounterDataOperation) *xxx_QueryCounterDataOperation {
	if op == nil {
		op = &xxx_QueryCounterDataOperation{}
	}
	if o == nil {
		return op
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *QueryCounterDataResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryCounterDataOperation) {
	if o == nil {
		return
	}
	o.OutSize = op.OutSize
	o.ReturnSize = op.ReturnSize
	o.Data = op.Data
	o.Return = op.Return
}
func (o *QueryCounterDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryCounterDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryCounterDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ValidateCountersOperation structure represents the PerflibV2ValidateCounters operation
type xxx_ValidateCountersOperation struct {
	Query  *Query `idl:"name:hQuery" json:"query"`
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
	Data   []byte `idl:"name:lpData;size_is:(dwInSize)" json:"data"`
	Add    uint32 `idl:"name:dwAdd" json:"add"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ValidateCountersOperation) OpNum() int { return 7 }

func (o *xxx_ValidateCountersOperation) OpName() string {
	return "/PerflibV2/v1/PerflibV2ValidateCounters"
}

func (o *xxx_ValidateCountersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.InSize == 0 {
		o.InSize = uint32(len(o.Data))
	}
	if o.InSize > uint32(67108864) {
		return fmt.Errorf("InSize is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateCountersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query != nil {
			if err := o.Query.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Query{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InSize); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwAdd {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Add); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateCountersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hQuery {in} (1:{context_handle, alias=RPC_HQUERY, names=ndr_context_handle}(struct))
	{
		if o.Query == nil {
			o.Query = &Query{}
		}
		if err := o.Query.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwInSize {in} (1:{range=(0,67108864), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InSize); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// dwAdd {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Add); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateCountersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateCountersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpData {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize](uchar))
	{
		dimSize1 := uint64(o.InSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
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

func (o *xxx_ValidateCountersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpData {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=dwInSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
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

// ValidateCountersRequest structure represents the PerflibV2ValidateCounters operation request
type ValidateCountersRequest struct {
	// hQuery: The handle that is created by the PerflibV2OpenQueryHandle method; an exception
	// is thrown or an error is returned by RPC if the handle did not originate from the
	// PerflibV2OpenQueryHandle method.
	Query *Query `idl:"name:hQuery" json:"query"`
	// dwInSize: The size, in bytes, of the buffer.
	InSize uint32 `idl:"name:dwInSize" json:"in_size"`
	// lpData: The buffer that contains the counter information to add to, or remove from,
	// the query. The server will return this buffer after it has attempted to add or remove
	// the specified counters; the Status field of each _PERF_COUNTER_IDENTIFIER structure
	// will contain information about whether or not the server was successful.
	Data []byte `idl:"name:lpData;size_is:(dwInSize)" json:"data"`
	// dwAdd: A Boolean value that indicates if counters are being added to, or removed
	// from, the query. If counters are being added, this MUST be set to TRUE; otherwise,
	// it MUST be set to FALSE.
	Add uint32 `idl:"name:dwAdd" json:"add"`
}

func (o *ValidateCountersRequest) xxx_ToOp(ctx context.Context, op *xxx_ValidateCountersOperation) *xxx_ValidateCountersOperation {
	if op == nil {
		op = &xxx_ValidateCountersOperation{}
	}
	if o == nil {
		return op
	}
	o.Query = op.Query
	o.InSize = op.InSize
	o.Data = op.Data
	o.Add = op.Add
	return op
}

func (o *ValidateCountersRequest) xxx_FromOp(ctx context.Context, op *xxx_ValidateCountersOperation) {
	if o == nil {
		return
	}
	o.Query = op.Query
	o.InSize = op.InSize
	o.Data = op.Data
	o.Add = op.Add
}
func (o *ValidateCountersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ValidateCountersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateCountersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ValidateCountersResponse structure represents the PerflibV2ValidateCounters operation response
type ValidateCountersResponse struct {
	// lpData: The buffer that contains the counter information to add to, or remove from,
	// the query. The server will return this buffer after it has attempted to add or remove
	// the specified counters; the Status field of each _PERF_COUNTER_IDENTIFIER structure
	// will contain information about whether or not the server was successful.
	Data []byte `idl:"name:lpData;size_is:(dwInSize)" json:"data"`
	// Return: The PerflibV2ValidateCounters return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ValidateCountersResponse) xxx_ToOp(ctx context.Context, op *xxx_ValidateCountersOperation) *xxx_ValidateCountersOperation {
	if op == nil {
		op = &xxx_ValidateCountersOperation{}
	}
	if o == nil {
		return op
	}
	o.Data = op.Data
	o.Return = op.Return
	return op
}

func (o *ValidateCountersResponse) xxx_FromOp(ctx context.Context, op *xxx_ValidateCountersOperation) {
	if o == nil {
		return
	}
	o.Data = op.Data
	o.Return = op.Return
}
func (o *ValidateCountersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ValidateCountersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateCountersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
