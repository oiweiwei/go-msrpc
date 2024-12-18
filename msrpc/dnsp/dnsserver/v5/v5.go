package dnsserver

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dnsp "github.com/oiweiwei/go-msrpc/msrpc/dnsp"
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
	_ = dnsp.GoPackage
)

var (
	// import guard
	GoPackage = "dnsp"
)

var (
	// Syntax UUID
	DNSServerSyntaxUUID = &uuid.UUID{TimeLow: 0x50abc2a4, TimeMid: 0x574d, TimeHiAndVersion: 0x40b3, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x66, Node: [6]uint8{0xee, 0x4f, 0xd5, 0xfb, 0xa0, 0x76}}
	// Syntax ID
	DNSServerSyntaxV5_0 = &dcerpc.SyntaxID{IfUUID: DNSServerSyntaxUUID, IfVersionMajor: 5, IfVersionMinor: 0}
)

// DnsServer interface.
type DNSServerClient interface {

	// The R_DnssrvOperation method is used to invoke a set of server functions specified
	// by pszOperation.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. See [MS-ERREF] section 2.2 or
	// section 2.2.1.1.5. All error values MUST be treated the same.
	Operation(context.Context, *OperationRequest, ...dcerpc.CallOption) (*OperationResponse, error)

	// The R_DnssrvQuery method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone and pszOperation parameters.
	// For the purpose of selecting an output structure type the server MUST consider the
	// value of dwClientVersion (section 2.2.1.2.1) to be 0x00000000 when responding to
	// this method.
	//
	// Return Values: A Win32 error code indicating whether the operation completed successfully
	// (0x00000000) or failed (any other value). See [MS-ERREF] section 2.2 or section 2.2.1.1.5.
	Query(context.Context, *QueryRequest, ...dcerpc.CallOption) (*QueryResponse, error)

	// The R_DnssrvComplexOperation method is used to invoke a set of server functions specified
	// by the caller. These functions generally return more complex structures than simple
	// 32-bit integer values, unlike the operations accessible through R_DnssrvOperation
	// (section 3.1.4.1). For the purpose of selecting an output structure type the server
	// MUST consider the value of dwClientVersion (section 2.2.1.2.1) to be 0x00000000 when
	// responding to this method.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. See [MS-ERREF] section 2.2 or
	// section 2.2.1.1.5. All error values MUST be treated the same.
	ComplexOperation(context.Context, *ComplexOperationRequest, ...dcerpc.CallOption) (*ComplexOperationResponse, error)

	// The R_DnssrvEnumRecords method enumerates DNS records on the server.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code if an error occurred. See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same, except that if the return code
	// is ERROR_MORE_DATA (0x000000EA) then the enumeration contains more results than can
	// fit into a single RPC buffer. In this case the client application can call this method
	// again passing the last retrieved child as the pszStartChild argument to retrieve
	// the next set of results.
	EnumRecords(context.Context, *EnumRecordsRequest, ...dcerpc.CallOption) (*EnumRecordsResponse, error)

	// The R_DnssrvUpdateRecord method is used to add a new DNS record or modify/delete
	// an existing DNS record at the server. This operation SHOULD<277> be supported.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code if an error occurred. See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same.
	UpdateRecord(context.Context, *UpdateRecordRequest, ...dcerpc.CallOption) (*UpdateRecordResponse, error)

	// The R_DnssrvOperation2 method is used to invoke a set of server functions specified
	// by the caller. The DNS server SHOULD implement R_DnssrvOperation2.
	//
	// All parameters are as specified by the R_DnssrvOperation method (section 3.1.4.1)
	// with the following exceptions:
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. See [MS-ERREF] section 2.2 or
	// section 2.2.1.1.5. All error values MUST be treated the same.
	Operation2(context.Context, *Operation2Request, ...dcerpc.CallOption) (*Operation2Response, error)

	// The R_DnssrvQuery2 method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone and pszOperation parameters.
	// The DNS server SHOULD implement R_ DnssrvQuery2 <280>.
	//
	// All parameters are as specified by the R_DnssrvQuery method (section 3.1.4.2) with
	// the following exceptions:
	//
	// Return Values: Return values behaviors and interpretations are same as they are for
	// R_DnssrvQuery method (section 3.1.4.2).
	Query2(context.Context, *Query2Request, ...dcerpc.CallOption) (*Query2Response, error)

	// The R_DnssrvComplexOperation2 method is used to invoke a set of server functions
	// specified by the caller. These functions generally return more complex structures
	// than simple 32-bit integer values, unlike the operations accessible through R_DnssrvOperation
	// (section 3.1.4.1). The DNS server SHOULD implement R_DnssrvComplexOperation2.
	//
	// All parameters are as specified by the R_DnssrvComplexOperation method (section 3.1.4.3)
	// with the following exceptions:
	//
	// Return Values: Return values and interpretations are the same as for R_DnssrvComplexOperation
	// (section 3.1.4.3).
	ComplexOperation2(context.Context, *ComplexOperation2Request, ...dcerpc.CallOption) (*ComplexOperation2Response, error)

	// The R_DnssrvEnumRecords2 method enumerates DNS records on the server. The DNS server
	// SHOULD implement R_DnssrvEnumRecords2 <281>.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords method (section 3.1.4.4)
	// with the following exceptions:
	//
	// Return Values: Return values behaviors and interpretations are same as they are for
	// R_DnssrvEnumRecords method (section 3.1.4.4).
	EnumRecords2(context.Context, *EnumRecords2Request, ...dcerpc.CallOption) (*EnumRecords2Response, error)

	// The R_DnssrvUpdateRecord2 method is used to add a new DNS record or modify/delete
	// an existing DNS record at the server. The DNS server SHOULD implement R_ DnssrvEnumRecords2.<282>
	//
	// All parameters are as specified by the R_DnssrvUpdateRecord method (section 3.1.4.5)
	// with the following exceptions:
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code if an error occurred. See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same. All record types SHOULD be
	// supported, but if an operation is attempted on an unsupported record type, the method
	// MUST return a nonzero Win32 error code.
	UpdateRecord2(context.Context, *UpdateRecord2Request, ...dcerpc.CallOption) (*UpdateRecord2Response, error)

	// The R_DnssrvUpdateRecord3 method is used to add a new DNS record or modify or delete
	// an existing DNS record in a zone or in a zone scope, or in a cache zone or cache
	// scope, if specified. The DNS server SHOULD<283> implement R_DnssrvUpdateRecord3.
	//
	// All parameters are as specified by the methods R_DnssrvUpdateRecord (section 3.1.4.5)
	// and R_DnssrvUpdateRecord2 (section 3.1.4.10) with the following exceptions:
	UpdateRecord3(context.Context, *UpdateRecord3Request, ...dcerpc.CallOption) (*UpdateRecord3Response, error)

	// The R_DnssrvEnumRecords3 method enumerates DNS records on a zone or a zone scope,
	// or cache zone or a cache scope, if specified. The DNS server SHOULD<284> implement
	// R_DnssrvEnumRecords3.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords method (section 3.1.4.4)
	// and implement the R_DnssrvEnumRecords2 method (section 3.1.4.9) with the following
	// exceptions:
	EnumRecords3(context.Context, *EnumRecords3Request, ...dcerpc.CallOption) (*EnumRecords3Response, error)

	// The R_DnssrvOperation3 method is used to invoke a set of server functions specified
	// by the caller on the zone scope or cache scope if specified. The DNS server SHOULD<285>
	// implement R_DnssrvOperation3.
	//
	// All parameters are as specified by the methods R_DnssrvOperation (section 3.1.4.1)
	// and R_DnssrvOperation2 (section 3.1.4.6) with the following exceptions.
	Operation3(context.Context, *Operation3Request, ...dcerpc.CallOption) (*Operation3Response, error)

	// The R_DnssrvQuery3 method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone, pwszZoneScopeName, and
	// pszOperation parameters. The DNS server SHOULD<287> implement R_DnssrvQuery3.
	//
	// All the parameters are as specified by the R_DnssrvQuery2 method with the following
	// exceptions:
	Query3(context.Context, *Query3Request, ...dcerpc.CallOption) (*Query3Response, error)

	// The R_DnssrvComplexOperation3 method is used to invoke a set of server functions
	// specified by the caller. These functions generally return more complex structures
	// than simple 32-bit integer values, unlike the operations accessible through R_DnssrvOperation
	// (section 3.1.4.1). The DNS server SHOULD<289> implement R_DnssrvComplexOperation2
	// (section 3.1.4.8).
	//
	// All parameters are as specified by the R_DnssrvComplexOperation method (section 3.1.4.3)
	// with the following exceptions:
	ComplexOperation3(context.Context, *ComplexOperation3Request, ...dcerpc.CallOption) (*ComplexOperation3Response, error)

	// The R_DnssrvOperation4 method is used to invoke a set of server functions specified
	// by the caller on the virtualization instance, if specified. The DNS server SHOULD<290>
	// implement R_DnssrvOperation4.
	//
	// All parameters are as specified by the R_DnssrvOperation3 (section 3.1.4.13) method
	// with the following exceptions.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. If unable to perform the operation,
	// returns error EPT_S_CANT_PERFORM_OP (0x000006D8). See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same.
	Operation4(context.Context, *Operation4Request, ...dcerpc.CallOption) (*Operation4Response, error)

	// The R_DnssrvQuery4 method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pwszVirtualizationInstanceID, and
	// pszOperation parameters. The DNS server SHOULD<291> implement R_DnssrvQuery4.
	//
	// All the parameters are as specified by the R_DnssrvQuery2 method (section 3.1.4.7)
	// method with the following exceptions:
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. If unable to perform the operation,
	// returns error EPT_S_CANT_PERFORM_OP (0x000006D8). See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same.
	Query4(context.Context, *Query4Request, ...dcerpc.CallOption) (*Query4Response, error)

	// The R_DnssrvUpdateRecord4 method is used to add a new DNS record or to modify or
	// delete an existing DNS record in a zone or in a zone scope under a virtualization
	// instance, if specified. The DNS server SHOULD<292> implement R_DnssrvUpdateRecord4.
	//
	// All parameters are as specified by the R_DnssrvUpdateRecord3 (section 3.1.4.11) method
	// with the following exceptions:
	UpdateRecord4(context.Context, *UpdateRecord4Request, ...dcerpc.CallOption) (*UpdateRecord4Response, error)

	// The R_DnssrvEnumRecords4 method enumerates DNS records on a zone or a zone scope
	// in a virtualization instance, if specified. The DNS server SHOULD<293> implement
	// R_DnssrvEnumRecords4.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords3 method (section 3.1.4.12)
	// with the following exceptions:
	EnumRecords4(context.Context, *EnumRecords4Request, ...dcerpc.CallOption) (*EnumRecords4Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultDNSServerClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDNSServerClient) Operation(ctx context.Context, in *OperationRequest, opts ...dcerpc.CallOption) (*OperationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OperationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Query(ctx context.Context, in *QueryRequest, opts ...dcerpc.CallOption) (*QueryResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) ComplexOperation(ctx context.Context, in *ComplexOperationRequest, opts ...dcerpc.CallOption) (*ComplexOperationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ComplexOperationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) EnumRecords(ctx context.Context, in *EnumRecordsRequest, opts ...dcerpc.CallOption) (*EnumRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...dcerpc.CallOption) (*UpdateRecordResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateRecordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Operation2(ctx context.Context, in *Operation2Request, opts ...dcerpc.CallOption) (*Operation2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Operation2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Query2(ctx context.Context, in *Query2Request, opts ...dcerpc.CallOption) (*Query2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Query2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) ComplexOperation2(ctx context.Context, in *ComplexOperation2Request, opts ...dcerpc.CallOption) (*ComplexOperation2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ComplexOperation2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) EnumRecords2(ctx context.Context, in *EnumRecords2Request, opts ...dcerpc.CallOption) (*EnumRecords2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRecords2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) UpdateRecord2(ctx context.Context, in *UpdateRecord2Request, opts ...dcerpc.CallOption) (*UpdateRecord2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateRecord2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) UpdateRecord3(ctx context.Context, in *UpdateRecord3Request, opts ...dcerpc.CallOption) (*UpdateRecord3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateRecord3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) EnumRecords3(ctx context.Context, in *EnumRecords3Request, opts ...dcerpc.CallOption) (*EnumRecords3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRecords3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Operation3(ctx context.Context, in *Operation3Request, opts ...dcerpc.CallOption) (*Operation3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Operation3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Query3(ctx context.Context, in *Query3Request, opts ...dcerpc.CallOption) (*Query3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Query3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) ComplexOperation3(ctx context.Context, in *ComplexOperation3Request, opts ...dcerpc.CallOption) (*ComplexOperation3Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ComplexOperation3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Operation4(ctx context.Context, in *Operation4Request, opts ...dcerpc.CallOption) (*Operation4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Operation4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) Query4(ctx context.Context, in *Query4Request, opts ...dcerpc.CallOption) (*Query4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Query4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) UpdateRecord4(ctx context.Context, in *UpdateRecord4Request, opts ...dcerpc.CallOption) (*UpdateRecord4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateRecord4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) EnumRecords4(ctx context.Context, in *EnumRecords4Request, opts ...dcerpc.CallOption) (*EnumRecords4Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRecords4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDNSServerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDNSServerClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDNSServerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DNSServerClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DNSServerSyntaxV5_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDNSServerClient{cc: cc}, nil
}

// xxx_OperationOperation structure represents the R_DnssrvOperation operation
type xxx_OperationOperation struct {
	ServerName string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone       string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Context    uint32      `idl:"name:dwContext" json:"context"`
	Operation  string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID     uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data       *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
	Return     int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_OperationOperation) OpNum() int { return 0 }

func (o *xxx_OperationOperation) OpName() string { return "/DnsServer/v5/R_DnssrvOperation" }

func (o *xxx_OperationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OperationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OperationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OperationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OperationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OperationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OperationRequest structure represents the R_DnssrvOperation operation request
type OperationRequest struct {
	// pwszServerName: The client SHOULD pass a pointer to the FQDN of the target server
	// as a null-terminated UTF-16LE character string. The server MUST ignore this value.
	ServerName string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pszZone: A pointer to a null-terminated character string that contains the name of
	// the zone to be queried. For operations specific to a particular zone, this string
	// MUST contain the name of the zone in UTF-8 format or a multizone operation string
	// (given in the table that follows) that indicates that the operation is performed
	// on multiple zones, but only if dwContext is zero. If dwContext is not zero, then
	// the value of pszZone MUST be ignored. For all other operations this value MUST be
	// set to NULL. When pszZone is NULL, the valid operations are in the first table under
	// the pszOperation section that follows, or are a property name listed in section 3.1.1.1.2,
	// 3.1.1.1.3, or 3.1.1.1.4. If this value is not NULL, then this value will be used
	// by certain operations as specified in the second table for pszOperation that follows.
	//
	// The following table shows what values are used to request that the operation be performed
	// on multiple zones, using ZONE_REQUEST_FILTERS values (section 2.2.5.1.4).
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllZones"                 | ZONE_REQUEST_PRIMARY | ZONE_REQUEST_SECONDARY | ZONE_REQUEST_AUTO |              |
	//	|                              | ZONE_REQUEST_FORWARD | ZONE_REQUEST_REVERSE | ZONE_REQUEST_FORWARDER             |
	//	|                              | | ZONE_REQUEST_STUB | ZONE_REQUEST_DS | ZONE_REQUEST_NON_DS |                    |
	//	|                              | ZONE_REQUEST_DOMAIN_DP | ZONE_REQUEST_FOREST_DP | ZONE_REQUEST_CUSTOM_DP |       |
	//	|                              | ZONE_REQUEST_LEGACY_DP                                                           |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllZonesAndCache"         | ZONE_REQUEST_PRIMARY | ZONE_REQUEST_SECONDARY | ZONE_REQUEST_CACHE               |
	//	|                              | | ZONE_REQUEST_AUTO | ZONE_REQUEST_FORWARD | ZONE_REQUEST_REVERSE                |
	//	|                              | | ZONE_REQUEST_FORWARDER | ZONE_REQUEST_STUB | ZONE_REQUEST_DS |                 |
	//	|                              | ZONE_REQUEST_NON_DS | ZONE_REQUEST_DOMAIN_DP | ZONE_REQUEST_FOREST_DP |          |
	//	|                              | ZONE_REQUEST_CUSTOM_DP | ZONE_REQUEST_LEGACY_DP                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllPrimaryZones"          | ZONE_REQUEST_PRIMARY                                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllSecondaryZones"        | ZONE_REQUEST_SECONDARY                                                           |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllForwardZones"          | ZONE_REQUEST_FORWARD                                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllReverseZones"          | ZONE_REQUEST_REVERSE                                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllDsZones"               | ZONE_REQUEST_DS                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllNonDsZones"            | ZONE_REQUEST_NON_DS                                                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllPrimaryReverseZones"   | ZONE_REQUEST_REVERSE | ZONE_REQUEST_PRIMARY                                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllPrimaryForwardZones"   | ZONE_REQUEST_FORWARD | ZONE_REQUEST_PRIMARY                                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllSecondaryReverseZones" | ZONE_REQUEST_REVERSE | ZONE_REQUEST_SECONDARY                                    |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| "..AllSecondaryForwardZones" | ZONE_REQUEST_FORWARD | ZONE_REQUEST_SECONDARY                                    |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Zone string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// dwContext: A value used to specify multizone operations in ZONE_REQUEST_FILTERS (section
	// 2.2.5.1.4) format or zero if the operation is not meant to apply to multiple zones.
	// If pszZone is not NULL and matches the name of a zone hosted by the DNS server then
	// the value of dwContext MUST be ignored.
	Context uint32 `idl:"name:dwContext" json:"context"`
	// pszOperation: A pointer to a null-terminated ASCII character string that contains
	// the name of operation to be performed on the server. These are two sets of allowed
	// values for pszOperation:
	//
	// If pszZone is set to NULL, pszOperation MUST be either a property name listed in
	// section 3.1.1.1.2, 3.1.1.1.3 or 3.1.1.1.4, or SHOULD<215> be one of the following.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ResetDwordProperty               | Update the value of a (name, value) pair in the DNS server configuration. On     |
	//	|                                  | input, dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST      |
	//	|                                  | point to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5) that     |
	//	|                                  | specifies the name of a property listed in section 3.1.1.1.1 and a new value for |
	//	|                                  | that property.                                                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| Restart                          | The server SHOULD restart the DNS server process. dwTypeId and pData MUST be     |
	//	|                                  | ignored by the server.<216>                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ClearDebugLog                    | Clear the debug log. dwTypeId and pData MUST be ignored by the server.           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ClearCache                       | Delete all cached records from the cache zone or cache scope memory. dwTypeId    |
	//	|                                  | and pData MUST be ignored by the server.                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| WriteDirtyZones                  | Write all zones that are stored in local persistent storage to local persistent  |
	//	|                                  | storage if the zone's Dirty Flag (section 3.1.1) is set to TRUE. dwTypeId and    |
	//	|                                  | pData MUST be ignored by the server.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ZoneCreate                       | Create a zone. On input, dwTypeId SHOULD<217> be set to                          |
	//	|                                  | DNSSRV_TYPEID_ZONE_CREATE. pData MUST point to a structure of one of the types   |
	//	|                                  | specified in DNS DNS_RPC_ZONE_CREATE_INFO (section 2.2.5.2.7) that contains all  |
	//	|                                  | parameters of a new zone to be created by the DNS server, and pData MUST conform |
	//	|                                  | to the description corresponding to the value of dwTypeId (section 2.2.1.1.1) If |
	//	|                                  | pData points to a DNS_ZONE_TYPE_CACHE or DNS_ZONE_TYPE_SECONDARY_CACHE record,   |
	//	|                                  | the server MUST return a nonzero error. If pData points to a DNS_ZONE_TYPE_STUB, |
	//	|                                  | DNS_ZONE_TYPE_SECONDARY, or DNS_ZONE_TYPE_FORWARDER record, the server MAY       |
	//	|                                  | return a nonzero error, but SHOULD return success.<218>                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ClearStatistics                  | Clears server statistics data on the DNS server. dwTypeId and pData MUST be      |
	//	|                                  | ignored by the server.                                                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnlistDirectoryPartition         | On input dwTypeId MUST be set to DNSSRV_TYPEID_ENLIST_DP, and the pData MUST     |
	//	|                                  | point to a DNS_RPC_ENLIST_DP (section 2.2.7.2.5) structure. This operation       |
	//	|                                  | allows application directory partitions to be added to or deleted from the       |
	//	|                                  | Application Directory Partition Table, and also allows the DNS server to be      |
	//	|                                  | directed to add or remove itself from the replication scope of an existing       |
	//	|                                  | application directory partition.                                                 |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| StartScavenging                  | Initiate a resource record scavenging cycle on the DNS server. dwTypeId, and     |
	//	|                                  | pData MUST be ignored by the server.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| AbortScavenging                  | Terminate a resource record scavenging cycle on the DNS server. dwTypeId and     |
	//	|                                  | pData MUST be ignored by the server.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| AutoConfigure                    | On input, dwTypeId SHOULD be set to DNSSRV_TYPEID_AUTOCONFIGURE, in which        |
	//	|                                  | case pData MUST point to a structure of type DNS_RPC_AUTOCONFIGURE (section      |
	//	|                                  | 2.2.8.2.1)<219>; dwTypeId MAY instead be set to DNSSRV_TYPEID_DWORD in which     |
	//	|                                  | case pData MUST point to a DWORD in DNS_RPC_AUTOCONFIG (section 2.2.8.1.1)       |
	//	|                                  | format.                                                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ExportSettings                   | Export DNS settings on the DNS server to a file on the DNS server. dwTypeId      |
	//	|                                  | SHOULD be set to DNSSRV_TYPEID_LPWSTR, and pData MUST be ignored by the server.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| PrepareForDemotion               | Prepares for demotion by removing references to this DNS server from all zones   |
	//	|                                  | stored in the directory server. dwTypeId and pData MUST be ignored by the        |
	//	|                                  | server.                                                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| PrepareForUninstall              | This operation does nothing on the DNS server. dwTypeId and pData MUST be        |
	//	|                                  | ignored by the server.                                                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteNode                       | On input dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST    |
	//	|                                  | point to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5) that     |
	//	|                                  | contains the FQDN of the node pointed to by pszNodeName in the DNS server's      |
	//	|                                  | cache to be deleted and a Boolean flag in dwParam to indicate if the node        |
	//	|                                  | subtree is to be deleted.                                                        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteRecordSet                  | On input dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST    |
	//	|                                  | point to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5). That    |
	//	|                                  | structure contains the FQDN of the node to be deleted, which is cached on the    |
	//	|                                  | DNS server, and the type of record set in the dwParam member, which indicates    |
	//	|                                  | whether the entire set of this type is to be deleted. The type MUST be a         |
	//	|                                  | DNS_RECORD_TYPE value (section 2.2.2.1.1) or 0x00FF, which specifies all types.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| WriteBackFile                    | Write all information for root hints back to persistent storage. dwTypeId and    |
	//	|                                  | pData MUST be ignored by the server.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ListenAddresses                  | On input, dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                       |
	//	|                                  | DNSSRV_TYPEID_ADDRARRAY and pData MUST point to a structure of type IP4_ARRAY    |
	//	|                                  | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, which    |
	//	|                                  | contains a list of new IP addresses on which the DNS server can listen. The      |
	//	|                                  | server SHOULD accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and MAY accept  |
	//	|                                  | DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.<220>                                        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| Forwarders                       | On input dwTypeId SHOULD be set to DNSSRV_TYPEID_FORWARDERS<221>, and pData      |
	//	|                                  | MUST point to a structure of one of the types specified in DNS_RPC_FORWARDERS    |
	//	|                                  | (section 2.2.5.2.10), which contains information about new IP addresses to which |
	//	|                                  | the DNS server can forward queries.                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| LogFilePath                      | On input dwTypeId MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to   |
	//	|                                  | a Unicode string that contains an absolute or relative pathname or filename for  |
	//	|                                  | the debug log file on the DNS server.                                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| LogIPFilterList                  | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                                  | DNSSRV_TYPEID_ADDRARRAY, and pData MUST point to a structure of type IP4_ARRAY   |
	//	|                                  | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, which    |
	//	|                                  | contains a list of new IP addresses used for debug log filter. The DNS server    |
	//	|                                  | will write to the debug log only for traffic to/from these IP addresses. The     |
	//	|                                  | server SHOULD accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and MAY accept  |
	//	|                                  | DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.<222>                                        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ForestDirectoryPartitionBaseName | The DNS server MUST return an error.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DomainDirectoryPartitionBaseName | The DNS server MUST return an error.                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| GlobalQueryBlockList             | Update the list of single-label names for which queries are blocked. Query       |
	//	|                                  | names that match this list, in any primary zone, will be blocked. On input       |
	//	|                                  | dwTypeId MUST be set to DNSSRV_TYPEID_UTF8_STRING_LIST, and pData MUST point to  |
	//	|                                  | a structure of type DNS_RPC_UTF8_STRING_LIST (section 2.2.1.2.3).                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| BreakOnReceiveFrom               | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                                  | DNSSRV_TYPEID_ADDRARRAY and pData MUST point to a structure of type IP4_ARRAY    |
	//	|                                  | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, that     |
	//	|                                  | contains a list of new IP addresses for which the DNS server will execute        |
	//	|                                  | a breakpoint if a packet is received from these IP addresses. The server         |
	//	|                                  | SHOULD accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and MAY accept         |
	//	|                                  | DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.<223>                                        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| BreakOnUpdateFrom                | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                                  | DNSSRV_TYPEID_ADDRARRAY, and pData MUST point to a structure of type IP4_ARRAY   |
	//	|                                  | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, that     |
	//	|                                  | contains a list of new IP addresses for which the DNS server will execute        |
	//	|                                  | a breakpoint if an update is received from these IP addresses. The server        |
	//	|                                  | SHOULD accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and MAY accept         |
	//	|                                  | DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.<224>                                        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ServerLevelPluginDll             | On input dwTypeId MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to a |
	//	|                                  | Unicode string that contains an absolute pathname for server side plug-in binary |
	//	|                                  | on the DNS server or an empty Unicode string.                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ActiveRefreshAllTrustPoints      | Schedules an immediate RFC 5011 active refresh for all trust points, regardless  |
	//	|                                  | of the time of the last active refresh. The dwTypeId and pData parameters MUST   |
	//	|                                  | be set to zero/NULL by the client and MUST be ignored by the server.<225>        |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| CreateServerScope                | Creates a server scope on the DNS server. The dwTypeId parameter MUST be set     |
	//	|                                  | to DNSSRV_TYPEID_LPWSTR. pData MUST point to a Unicode string that contains the  |
	//	|                                  | name of the server scope to be created.<226>                                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteServerScope                | Deletes a server scope on the DNS server. The dwTypeId parameter MUST be set     |
	//	|                                  | to DNSSRV_TYPEID_LPWSTR. pData MUST point to a Unicode string that contains the  |
	//	|                                  | name of the server scope to be deleted.<227>                                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| CreateClientSubnetRecord         | Creates a client subnet record on a DNS server. The dwTypeId parameter MUST be   |
	//	|                                  | set to DNSSRV_TYPEID_CLIENT_SUBNET_RECORD, and pData MUST point to a structure   |
	//	|                                  | of type DNS_RPC_CLIENT_SUBNET_RECORD.                                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteClientSubnetRecord         | Deletes a client subnet record on a DNS server. The dwTypeId parameter MUST be   |
	//	|                                  | set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to a NULL-terminated Unicode   |
	//	|                                  | string containing the name of the client subnet record to be deleted.            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteSubnetsInRecord            | Updates a client subnet record on a DNS server by deleting the IP/IPv6           |
	//	|                                  | Subnets from the Client Subnet Record. The dwTypeId parameter MUST be set to     |
	//	|                                  | DNSSRV_TYPEID_CLIENT_SUBNET_RECORD, and pData MUST point to a structure of type  |
	//	|                                  | DNS_RPC_CLIENT_SUBNET_RECORD.                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| AddSubnetsInRecord               | Updates a client subnet record on a DNS server by adding the IP/IPv6             |
	//	|                                  | Subnets to the Client Subnet Record. The dwTypeId parameter MUST be set to       |
	//	|                                  | DNSSRV_TYPEID_CLIENT_SUBNET_RECORD, and pData MUST point to a structure of type  |
	//	|                                  | DNS_RPC_CLIENT_SUBNET_RECORD.                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ResetClientSubnetRecord          | Resets a client subnet record on a DNS server by deleting the existing IP/IPv6   |
	//	|                                  | Subnets and adding the IP/IPv6 Subnets specific to the client subnet record. The |
	//	|                                  | dwTypeId parameter MUST be set to DNSSRV_TYPEID_CLIENT_SUBNET_RECORD, and pData  |
	//	|                                  | MUST point to a structure of type DNS_RPC_CLIENT_SUBNET_RECORD.                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| CreatePolicy                     | Creates a DNS Policy at the server level on a DNS server. The dwTypeId parameter |
	//	|                                  | MUST be set to DNSSRV_TYPEID_POLICY, and pData MUST point to a structure of type |
	//	|                                  | DNS_RPC_POLICY.                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeletePolicy                     | Deletes a DNS Policy at the server level on a DNS server. The dwTypeId parameter |
	//	|                                  | MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to NULL-terminated     |
	//	|                                  | Unicode string containing the name of the DNS Policy.                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| UpdatePolicy                     | Updates a DNS Policy at the server level on a DNS server. The dwTypeId parameter |
	//	|                                  | MUST be set to DNSSRV_TYPEID_POLICY, and pData MUST point to a structure of type |
	//	|                                  | DNS_RPC_POLICY.                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| SetRRL                           | Sets Response Rate Limiting parameters at the server level on a DNS server. The  |
	//	|                                  | dwTypeId parameter MUST be set to DNSSRV_TYPEID_RRL, and pData MUST point to a   |
	//	|                                  | structure of type DNS_RPC_RRL_PARAMS.                                            |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| CreateVirtualizationInstance     | Creates a virtualization instance on the DNS server, under which                 |
	//	|                                  | zones can be created later. The dwTypeId parameter MUST be set to                |
	//	|                                  | DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE (section 2.2.1.2.6), and the pData         |
	//	|                                  | parameter MUST point to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE      |
	//	|                                  | (section 2.2.1.2.6).                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteVirtualizationInstance     | Removes a virtualization instance on the DNS server. This also removes           |
	//	|                                  | the zones and zone scopes under the virtualization instance as well. The         |
	//	|                                  | dwTypeId parameter MUST be set to DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE,         |
	//	|                                  | and pData MUST point to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE.     |
	//	|                                  | The value of dwFlags in DNS_RPC_VIRTUALIZATION_INSTANCE if set to                |
	//	|                                  | DNS_RPC_FLAG_PRESERVE_ZONE_FILE, the DNS server deletes the zones under the      |
	//	|                                  | virtualization instance but keeps the zone files. By default a DNS server        |
	//	|                                  | removes the zone files of zones under a virtualization instance when the         |
	//	|                                  | virtualization instance is removed.                                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| UpdateVirtualizationInstance     | Modifies the members of the virtualization instance on the DNS server. The       |
	//	|                                  | dwTypeId parameter MUST be set to DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE, and     |
	//	|                                  | pData MUST point to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE. The     |
	//	|                                  | members of DNS_RPC_VIRTUALIZATION_INSTANCE to be modified is determined by bits  |
	//	|                                  | in member dwFlags, which can be DNS_RPC_FLAG_FRIENDLY_NAME or DNS_RPC_FLAG_DESC  |
	//	|                                  | (section 3.1.4.1)                                                                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//
	// If pszZone is not NULL, and pszOperation does not match a property name listed in
	// sections 3.1.1.2.2 or 3.1.1.2.3, then pszOperation SHOULD<228> be one of the following:
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|            VALUE             |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ResetDwordProperty           | Update the value of a DNS Zone integer property. On input dwTypeId MUST be       |
	//	|                              | set to DNSSRV_TYPEID_NAME_AND_PARAM and pData MUST point to a structure of       |
	//	|                              | type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5), which contains the name of a    |
	//	|                              | property listed in section 3.1.1.2.1 for the zone pointed to by pszZone and a    |
	//	|                              | new value for that property.<229>                                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ZoneTypeReset                | Change the zone's type, for example to convert a secondary zone into a           |
	//	|                              | primary zone. On input dwTypeId SHOULD be set to DNSSRV_TYPEID_ZONE_CREATE,      |
	//	|                              | and pData SHOULD point to a structure of one of the types specified              |
	//	|                              | in DNS_RPC_ZONE_CREATE_INFO (section 2.2.5.2.7), which contains the              |
	//	|                              | new configuration information for the zone. dwTypeId MAY be set to               |
	//	|                              | DNSSRV_TYPEID_ZONE_CREATE_W2K or DNSSRV_TYPEID_ZONE_CREATE_DOTNET.<230> The      |
	//	|                              | server MUST return a nonzero error if the conversion is not implemented.         |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| PauseZone                    | Pause activities for the zone pointed to by pszZone on the DNS server, and       |
	//	|                              | do not use this zone to answer queries or take updates until it is resumed.      |
	//	|                              | dwTypeId, and pData MUST be ignored by the server.                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ResumeZone                   | Resume activities for the zone pointed to by pszZone on the DNS server; the zone |
	//	|                              | thus becomes available to answer queries and take updates. dwTypeId and pData    |
	//	|                              | MUST be ignored by the server.                                                   |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteZone                   | Delete the zone pointed to by pszZone on the DNS server. dwTypeId and pData MUST |
	//	|                              | be ignored by the server.                                                        |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ReloadZone                   | Reloads data for the zone pointed to by pszZone on the DNS server from           |
	//	|                              | persistent storage. dwTypeId, and pData MUST be ignored by the server.           |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| RefreshZone                  | Force a refresh of the secondary zone pointed to by pszZone on the DNS server,   |
	//	|                              | from primary zone server. For this operation pszZone MUST point to a secondary   |
	//	|                              | zone only. dwTypeId and pData MUST be ignored by the server.                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ExpireZone                   | Force a refresh of the secondary zone pointed to by pszZone on the DNS server,   |
	//	|                              | from primary zone server. For this operation pszZone MUST point to a secondary   |
	//	|                              | zone only. dwTypeId and pData MUST be ignored by the server.                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| IncrementVersion             | Same as "WriteBackFile".                                                         |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| WriteBackFile                | If the zone has uncommitted changes, write back all information for the zone     |
	//	|                              | pointed to by pszZone to persistent storage, and notify any secondary DNS        |
	//	|                              | servers. dwTypeId and pData MUST be ignored by the server.                       |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteZoneFromDs             | Delete the zone pointed to by pszZone from the directory server. dwTypeId, and   |
	//	|                              | pData MUST be ignored by the server.                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| UpdateZoneFromDs             | Refresh data for the zone pointed to by pszZone from the directory server.       |
	//	|                              | dwTypeId, and pData MUST be ignored by the server.                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ZoneExport                   | Export zone data to a given file on the DNS server. On input dwTypeId MUST be    |
	//	|                              | set to DNSSRV_TYPEID_ZONE_EXPORT, and pData MUST point to a structure of type    |
	//	|                              | DNS_RPC_ZONE_EXPORT_INFO (section 2.2.5.2.8) that contains a file name pointed   |
	//	|                              | to by pszZoneExportFile.                                                         |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ZoneChangeDirectoryPartition | Move a zone to a given application directory partition. On input dwTypeId        |
	//	|                              | MUST be set to DNSSRV_TYPEID_ZONE_CHANGE_DP, and pData MUST point to structure   |
	//	|                              | of type DNS_RPC_ZONE_CHANGE_DP (section 2.2.7.2.6), which contains the new       |
	//	|                              | application directory partition name pointed to by pszDestPartition.             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteNode                   | Delete a node. On input dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM,    |
	//	|                              | and pData MUST point to a structure of type DNS_RPC_NAME_AND_PARAM (section      |
	//	|                              | 2.2.1.2.5), which contains the FQDN of the node pointed to by pszNodeName        |
	//	|                              | present in the zone pointed to by pszZone on the DNS server to be deleted and a  |
	//	|                              | Boolean flag in dwParam to indicate if the node's subtree is to be deleted.      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteRecordSet              | Delete all the DNS records of a particular type at a particular                  |
	//	|                              | node from the DNS server's cache. On input dwTypeId MUST be set to               |
	//	|                              | DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST point to a structure of type        |
	//	|                              | DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5). That structure contains the FQDN of  |
	//	|                              | the node to be deleted and the DNS record type in the dwParam member. The type   |
	//	|                              | MUST be a DNS_RECORD_TYPE value (section 2.2.2.1.1) or 0x00FF, which specifies   |
	//	|                              | all types.                                                                       |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ForceAgingOnNode             | On input dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST    |
	//	|                              | point to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5), which   |
	//	|                              | contains a node name in pszNodeName, and a Boolean flag in dwParam to indicate   |
	//	|                              | whether aging is performed on all nodes in the subtree. All DNS records at the   |
	//	|                              | specified node in the zone named by pszZone will have their aging time stamp     |
	//	|                              | set to the current time. If subtree aging is specified by dwParam than all DNS   |
	//	|                              | records at all nodes that are children of this node will also have their aging   |
	//	|                              | time stamps set to the current time.                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DatabaseFile                 | On input dwTypeId SHOULD be set to DNSSRV_TYPEID_ZONE_DATABASE<231>, and pData   |
	//	|                              | MUST point to a structure of one of the types specified in DNS_RPC_ZONE_DATABASE |
	//	|                              | (section 2.2.5.2.6), which specifies whether the zone is directory server        |
	//	|                              | integrated by setting fDsIntegrated to TRUE, and if it is not then pszFileName   |
	//	|                              | MUST point to a Unicode string containing the absolute pathname of a file on the |
	//	|                              | DNS server to which the zone database is stored.                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| MasterServers                | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                              | DNSSRV_TYPEID_ADDRARRAY, and pData MUST point to a structure of type IP4_ARRAY   |
	//	|                              | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, which    |
	//	|                              | contains a list of IP addresses of new primary DNS servers for the zone pointed  |
	//	|                              | to by pszZone. This operation is valid only for secondary zones present on the   |
	//	|                              | server. The server SHOULD accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and |
	//	|                              | SHOULD accept DNSSRV_TYPEID_IPARRAY and IP4_ARRAY. If the input data of either   |
	//	|                              | type is accepted and the DNS server is directory-server integrated, the value of |
	//	|                              | pData SHOULD be written to the directory server.<232>                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| LocalMasterServers           | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                              | DNSSRV_TYPEID_ADDRARRAY, and pData MUST point to a structure of type IP4_ARRAY   |
	//	|                              | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively,          |
	//	|                              | which contains a list of IP addresses of new local primary DNS servers           |
	//	|                              | for the zone pointed to by pszZone. This operation is valid only for stub        |
	//	|                              | zones present on the server, and if configured, this value overrides any         |
	//	|                              | primary DNS server configured in the directory server. The server SHOULD         |
	//	|                              | accept DNSSRV_TYPEID_ADDRARRAY and DNS_ADDR_ARRAY, and SHOULD accept             |
	//	|                              | DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.<233>                                        |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| SecondaryServers             | On input dwTypeId SHOULD<234> be set to DNSSRV_TYPEID_ZONE_SECONDARIES,          |
	//	|                              | and pData MUST point to a structure of one of the types specified in             |
	//	|                              | DNS_RPC_ZONE_SECONDARIES (section 2.2.5.2.5), which contains information about   |
	//	|                              | secondary DNS servers for the zone pointed to by pszZone.                        |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ScavengeServers              | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                              | DNSSRV_TYPEID_ADDRARRAY, and pData MUST point to a structure of type IP4_ARRAY   |
	//	|                              | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, which    |
	//	|                              | contains a list of IP addresses of new servers that can run scavenging on        |
	//	|                              | the zone pointed to by pszZone. This operation is valid only for directory       |
	//	|                              | server integrated zones. The server SHOULD accept DNSSRV_TYPEID_ADDRARRAY,       |
	//	|                              | and DNS_ADDR_ARRAY, and SHOULD accept DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.       |
	//	|                              | If the input data of either type is accepted and the DNS server is directory     |
	//	|                              | server-integrated, the value of pData SHOULD be written to the directory         |
	//	|                              | server.<235>                                                                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| AllowNSRecordsAutoCreation   | On input dwTypeId MUST be set to DNSSRV_TYPEID_IPARRAY or                        |
	//	|                              | DNSSRV_TYPEID_ADDRARRAY and pData MUST point to a structure of type IP4_ARRAY    |
	//	|                              | (section 2.2.3.2.1) or DNS_ADDR_ARRAY (section 2.2.3.2.3) respectively, which    |
	//	|                              | contains a list of IP addresses of new servers that can auto-create NS records   |
	//	|                              | for the zone pointed to by pszZone. This operation is valid only for directory   |
	//	|                              | server integrated zones. The server SHOULD accept DNSSRV_TYPEID_ADDRARRAY        |
	//	|                              | and DNS_ADDR_ARRAY, and SHOULD accept DNSSRV_TYPEID_IPARRAY and IP4_ARRAY.       |
	//	|                              | If the input data of either type is accepted and the DNS server is directory     |
	//	|                              | server-integrated, the value of pData SHOULD be written to the directory         |
	//	|                              | server.<236>                                                                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| BreakOnNameUpdate            | On input dwTypeId MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to   |
	//	|                              | a Unicode string that contains the FQDN of the node for which if an update is    |
	//	|                              | received the DNS server will execute a breakpoint.                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| SignZone                     | Sign a zone using DNSSEC, thus making the zone online-signed. The dwTypeId and   |
	//	|                              | pData parameters MUST be set to zero/NULL by the client and MUST be ignored by   |
	//	|                              | the server.                                                                      |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| UnsignZone                   | Unsign a zone signed via online signing and remove all DNSSEC data from the      |
	//	|                              | zone. The dwTypeId and pData parameters MUST be set to zero/NULL by the client   |
	//	|                              | and MUST be ignored by the server.                                               |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| ResignZone                   | Refreshes all DNSSEC data in an online-signed zone. The dwTypeId and pData       |
	//	|                              | parameters MUST be set to zero/NULL by the client and MUST be ignored by the     |
	//	|                              | server.                                                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| PerformZoneKeyRollover       | Queues a signing key descriptor for key rollover. On input, dwTypeId MUST be set |
	//	|                              | to DNSSRV_TYPEID_LPWSTR, and pData MUST point to a Unicode string representation |
	//	|                              | of the GUID of the signing key descriptor to be queued for rollover.             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| PokeZoneKeyRollover          | Instructs the DNS server to stop waiting for DS records [RFC4034] in the parent  |
	//	|                              | zone to be updated and to proceed with key rollover as specified by [RFC4641].   |
	//	|                              | On input, dwTypeId MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to  |
	//	|                              | a Unicode string representation of the GUID of the signing key descriptor to be  |
	//	|                              | queued for rollover.                                                             |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| RetrieveRootTrustAnchors     | Retrieves the root trust anchors from the XML file specified by the              |
	//	|                              | RootTrustAnchorsURL server property (section 3.1.1.1.3) and adds any valid DS    |
	//	|                              | records to the root trust anchors. The dwTypeId and pData parameters MUST be     |
	//	|                              | set to zero/NULL by the client and MUST be ignored by the server. The pszZone    |
	//	|                              | parameter MUST be set to the string "TrustAnchors" to indicate the name of the   |
	//	|                              | zone.                                                                            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| TransferKeymasterRole        | Transfers the key master role to the current server. The dwTypeId parameter MUST |
	//	|                              | be set to DNSSRV_TYPEID_DWORD, and pData MUST point to one of the values defined |
	//	|                              | in the following paragraphs.                                                     |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| CreateZoneScope              | Creates a zone scope in the specified zone or a cache scope in                   |
	//	|                              | the specified cache zone. The dwTypeId parameter MUST be set to                  |
	//	|                              | DNSSRV_TYPEID_ZONE_SCOPE_CREATE. pData MUST point to a structure of the type     |
	//	|                              | DNS_RPC_ZONE_SCOPE_CREATE_INFO_V1 (section 2.2.13.1.2.1) that contains all the   |
	//	|                              | parameters needed to create the zone scope or cache scope. pszZone MUST be the   |
	//	|                              | name of the zone in which the zone scope is to be created or be specified as     |
	//	|                              | "..cache" for a cache scope.<237>                                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeleteZoneScope              | Deletes a zone scope from the specified zone or a cache scope from a specified   |
	//	|                              | cache zone. The dwTypeId MUST be set to DNSSRV_TYPEID_LPWSTR. pData MUST point   |
	//	|                              | to the name of the zone scope or cache scope that is to be deleted. pszZone MUST |
	//	|                              | be the name of the zone from which the zone scope is to be deleted or set to     |
	//	|                              | "..cache" for a cache scope.<238>                                                |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| CreatePolicy                 | Creates a DNS Policy for the specified zone or a cache zone on a DNS server. The |
	//	|                              | dwTypeId parameter MUST be set to DNSSRV_TYPEID_POLICY, and pData MUST point to  |
	//	|                              | a structure of type DNS_RPC_POLICY.                                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| DeletePolicy                 | Deletes a DNS Policy for the specified zone or a cache zone on a DNS server. The |
	//	|                              | dwTypeId parameter MUST be set to DNSSRV_TYPEID_LPWSTR, and pData MUST point to  |
	//	|                              | NULL-terminated Unicode string containing the name of the DNS Policy.            |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| UpdatePolicy                 | Updates a DNS Policy for the specified zone or a cache zone on a DNS server. The |
	//	|                              | dwTypeId parameter MUST be set to DNSSRV_TYPEID_POLICY, and pData MUST point to  |
	//	|                              | a structure of type DNS_RPC_POLICY.                                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	Operation string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	// dwTypeId: A DNS_RPC_TYPEID (section 2.2.1.1.1) value that indicates the type of input
	// data pointed to by pData.
	TypeID uint32 `idl:"name:dwTypeId" json:"type_id"`
	// pData: Input data of type DNSSRV_RPC_UNION (section 2.2.1.2.6), which contains a
	// data structure as specified by dwTypeId.
	Data *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
}

func (o *OperationRequest) xxx_ToOp(ctx context.Context) *xxx_OperationOperation {
	if o == nil {
		return &xxx_OperationOperation{}
	}
	return &xxx_OperationOperation{
		ServerName: o.ServerName,
		Zone:       o.Zone,
		Context:    o.Context,
		Operation:  o.Operation,
		TypeID:     o.TypeID,
		Data:       o.Data,
	}
}

func (o *OperationRequest) xxx_FromOp(ctx context.Context, op *xxx_OperationOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Context = op.Context
	o.Operation = op.Operation
	o.TypeID = op.TypeID
	o.Data = op.Data
}
func (o *OperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OperationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OperationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OperationResponse structure represents the R_DnssrvOperation operation response
type OperationResponse struct {
	// Return: The R_DnssrvOperation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OperationResponse) xxx_ToOp(ctx context.Context) *xxx_OperationOperation {
	if o == nil {
		return &xxx_OperationOperation{}
	}
	return &xxx_OperationOperation{
		Return: o.Return,
	}
}

func (o *OperationResponse) xxx_FromOp(ctx context.Context, op *xxx_OperationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OperationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OperationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OperationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryOperation structure represents the R_DnssrvQuery operation
type xxx_QueryOperation struct {
	ServerName string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone       string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation  string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID     uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data       *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	Return     int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryOperation) OpNum() int { return 1 }

func (o *xxx_QueryOperation) OpName() string { return "/DnsServer/v5/R_DnssrvQuery" }

func (o *xxx_QueryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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

func (o *xxx_QueryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeId {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeId {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryRequest structure represents the R_DnssrvQuery operation request
type QueryRequest struct {
	// pwszServerName: The client can pass a pointer to the FQDN of the target server as
	// a null-terminated UTF-16LE character string. The server MUST ignore this value.
	ServerName string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pszZone: A pointer to a null-terminated character string that contains name of the
	// zone to be queried. For operations specific to a particular zone, this field MUST
	// contain the name of the zone in UTF-8 format. For all other operations, this field
	// MUST be NULL.
	Zone string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pszOperation: A pointer to a null-terminated character string that contains the name
	// of the operation to be performed on the server. These are two sets of allowed values
	// for pszOperation:
	//
	// If pszZone is set to NULL, pszOperation MUST be either a property name listed in
	// section 3.1.1.1, or it SHOULD<265> be the following.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|                        |                                                                                  |
	//	|         VALUE          |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| "ServerInfo"           | On output pdwTypeId SHOULD be set according to the value of the dwClientVersion  |
	//	|                        | field (section 2.2.1.2.1). If dwClientVersion is 0x00000000, then pdwTypeId      |
	//	|                        | SHOULD be set to DNSSRV_TYPEID_SERVER_INFO_W2K. If dwClientVersion is            |
	//	|                        | 0x00060000, then pdwTypeId SHOULD be set to DNSSRV_TYPEID_SERVER_INFO_DOTNET.    |
	//	|                        | If dwClientVersion is 0x00070000, then pdwTypeId SHOULD be set to                |
	//	|                        | DNSSRV_TYPEID_SERVER_INFO<266> ppData MUST point to a structure of one of the    |
	//	|                        | types specified in DNS_RPC_SERVER_INFO (section 2.2.4.2.2), which SHOULD contain |
	//	|                        | the configuration information for the DNS server, but MAY have some fields set   |
	//	|                        | to zero even when the related configuration value is nonzero.<267>               |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| VirtualizationInstance | Gets the details of a virtualization instance in the DNS server.                 |
	//	|                        | On success, the DNS server MUST set the dwTypeId parameter to                    |
	//	|                        | DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE (section 2.2.1.2.6), and MUST set the      |
	//	|                        | ppData point to a structure of type DNS_RPC_VIRTUALIZATION_INSTANCE (section     |
	//	|                        | 2.2.17.1.1).                                                                     |
	//	+------------------------+----------------------------------------------------------------------------------+
	//
	// If pszZone is not NULL, pszOperation MUST be either a property name listed in section
	// 3.1.1.2, or one of the following.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| "Zone"     | On output the value pointed to by pdwTypeId SHOULD<268> be set to                |
	//	|            | DNSSRV_TYPEID_ZONE and ppData MUST point to a structure of one of the types      |
	//	|            | specified in DNS_RPC_ZONE (section 2.2.5.2.1), which contains abbreviated        |
	//	|            | information about the zone pointed to by pszZone.                                |
	//	+------------+----------------------------------------------------------------------------------+
	//	| "ZoneInfo" | On output the value pointed to by pdwTypeId SHOULD<269> be set to                |
	//	|            | DNSSRV_TYPEID_ZONE_INFOand ppData MUST point to a structure of one of the        |
	//	|            | types specified in DNS_RPC_ZONE_INFO (section 2.2.5.2.4), which contains full    |
	//	|            | information about the zone pointed to by pszZone.                                |
	//	+------------+----------------------------------------------------------------------------------+
	Operation string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
}

func (o *QueryRequest) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		ServerName: o.ServerName,
		Zone:       o.Zone,
		Operation:  o.Operation,
	}
}

func (o *QueryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Operation = op.Operation
}
func (o *QueryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryResponse structure represents the R_DnssrvQuery operation response
type QueryResponse struct {
	// pdwTypeId: A pointer to an integer that on success contains a value of type DNS_RPC_TYPEID
	// (section 2.2.1.1.1) that indicates the type of data pointed to by ppData.
	TypeID uint32 `idl:"name:pdwTypeId" json:"type_id"`
	// ppData: A DNSSRV_RPC_UNION(section 2.2.1.2.6) that contains a data-structure as indicated
	// by dwTypeId.
	Data *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	// Return: The R_DnssrvQuery return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryResponse) xxx_ToOp(ctx context.Context) *xxx_QueryOperation {
	if o == nil {
		return &xxx_QueryOperation{}
	}
	return &xxx_QueryOperation{
		TypeID: o.TypeID,
		Data:   o.Data,
		Return: o.Return,
	}
}

func (o *QueryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryOperation) {
	if o == nil {
		return
	}
	o.TypeID = op.TypeID
	o.Data = op.Data
	o.Return = op.Return
}
func (o *QueryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComplexOperationOperation structure represents the R_DnssrvComplexOperation operation
type xxx_ComplexOperationOperation struct {
	ServerName string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone       string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation  string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeIn     uint32      `idl:"name:dwTypeIn" json:"type_in"`
	DataIn     *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
	TypeOut    uint32      `idl:"name:pdwTypeOut" json:"type_out"`
	DataOut    *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	Return     int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_ComplexOperationOperation) OpNum() int { return 2 }

func (o *xxx_ComplexOperationOperation) OpName() string {
	return "/DnsServer/v5/R_DnssrvComplexOperation"
}

func (o *xxx_ComplexOperationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataIn := uint32(o.TypeIn)
		if o.DataIn != nil {
			if err := o.DataIn.MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataIn == nil {
			o.DataIn = &dnsp.Union{}
		}
		_swDataIn := uint32(o.TypeIn)
		if err := o.DataIn.UnmarshalUnionNDR(ctx, w, _swDataIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeOut {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataOut := uint32(o.TypeOut)
		if o.DataOut != nil {
			if err := o.DataOut.MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeOut {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataOut == nil {
			o.DataOut = &dnsp.Union{}
		}
		_swDataOut := uint32(o.TypeOut)
		if err := o.DataOut.UnmarshalUnionNDR(ctx, w, _swDataOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ComplexOperationRequest structure represents the R_DnssrvComplexOperation operation request
type ComplexOperationRequest struct {
	// pwszServerName: The client SHOULD pass a pointer to the FQDN of the target server
	// as a null-terminated UTF-16LE character string. The server MUST ignore this value.
	ServerName string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pszZone: The name of the zone that is being operated on. This MUST be set to NULL
	// unless pszOperation is set to QueryDwordProperty, in which case this value MUST be
	// set either to NULL (to indicate that DNS server Configuration information is being
	// requested) or to the name of the zone to be queried in UTF-8 format (to indicate
	// that a DNS Zone integer property is being requested). This value will be used by
	// certain operations as specified in the table below.
	Zone string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pszOperation: The operation to perform. The value of pszOperation SHOULD<275> be
	// one of the following:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|                                  |                                                                                  |
	//	|              VALUE               |                                     MEANING                                      |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumZones                        | Enumerate zones present on the DNS server qualifying for a specified simple      |
	//	|                                  | zone filter value. On input, dwTypeIn MUST be set to DNSSRV_TYPEID_DWORD         |
	//	|                                  | and pDataIn MUST point to any combination of ZONE_REQUEST_FILTERS (section       |
	//	|                                  | 2.2.5.1.4) values. Unless an error is returned, on output the value pointed to   |
	//	|                                  | by pdwTypeOut MUST be set to DNSSRV_TYPEID_ZONE_LIST and ppDataOut MUST point    |
	//	|                                  | to a structure of one of the types specified in DNS_RPC_ZONE_LIST (section       |
	//	|                                  | 2.2.5.2.3).                                                                      |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumZones2                       | Enumerate zones present on the DNS server qualifying for a specified             |
	//	|                                  | complex zone filter value. On input, dwTypeIn MUST be set to                     |
	//	|                                  | DNSSRV_TYPEID_ENUM_ZONES_FILTER and pDataIn MUST point to a structure of type    |
	//	|                                  | DNS_RPC_ENUM_ZONES_FILTER (section 2.2.5.2.9). Unless an error is returned, on   |
	//	|                                  | output the value pointed to by pdwTypeOut MUST be set to DNSSRV_TYPEID_ZONE_LIST |
	//	|                                  | and MUST ppDataOut point to a structure of one of the types specified in         |
	//	|                                  | DNS_RPC_ZONE_LIST.                                                               |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumDirectoryPartitions          | Enumerate the Application Directory Partition Table known to the DNS server. On  |
	//	|                                  | input, dwTypeIn MUST be set to DNSSRV_TYPEID_DWORD and pDataIn MUST be set to    |
	//	|                                  | zero if all application directory partitions are enumerated or to 0x000000001    |
	//	|                                  | if the DNS domain partition and DNS forest partition are excluded from results.  |
	//	|                                  | Unless an error is returned, on output the value pointed to by pdwTypeOut MUST   |
	//	|                                  | be set to DNSSRV_TYPEID_DP_LIST and ppDataOut MUST point to a structure of type  |
	//	|                                  | DNS_RPC_DP_LIST (section 2.2.7.2.4).                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| DirectoryPartitionInfo           | Retrieve detailed information about a specified application directory partition. |
	//	|                                  | On input, dwTypeIn MUST be set to DNSSRV_TYPEID_LPSTR and pDataIn MUST point     |
	//	|                                  | to a null-terminated UTF-8 string specifying the distinguished name of an        |
	//	|                                  | application directory partition. Unless an error is returned, on output the      |
	//	|                                  | value pointed to by pdwTypeOut MUST be DNSSRV_TYPEID_DP_INFO and ppDataOut MUST  |
	//	|                                  | point to a structure of type DNS_RPC_DP_INFO (section 2.2.7.2.1).                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| Statistics                       | Retrieve statistics. On input dwTypeIn MUST be set to DNSSRV_TYPEID_DWORD        |
	//	|                                  | and pDataIn MUST point to any combination of the DNSSRV_STATID_TYPES (section    |
	//	|                                  | 2.2.10.1.1) values. Unless an error is returned, on output the value pointed to  |
	//	|                                  | by pdwTypeOut MUST be set to DNSSRV_TYPEID_BUFFER and ppDataOut MUST point to    |
	//	|                                  | a DNS_RPC_BUFFER structure (section 2.2.1.2.2) that contains a list of variable  |
	//	|                                  | sized DNSSRV_STATS structures (section 2.2.10.2.2).                              |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| QueryDwordProperty               | Retrieve the value of a 32-bit integer property. On input, dwTypeIn MUST be      |
	//	|                                  | set to DNSSRV_TYPEID _LPSTR and pDataIn MUST point to a null-terminated UTF-8    |
	//	|                                  | string specifying a zone property name listed in section 3.1.1.2.1 (if pszZone   |
	//	|                                  | is non-NULL) or server property name listed in section 3.1.1.1.1 (if pszZone is  |
	//	|                                  | NULL). Unless an error is returned, on output the value pointed to by pdwTypeOut |
	//	|                                  | MUST be set to DNSSRV_TYPEID_DWORD and ppDataOut MUST point to a DWORD value.    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| IpValidate                       | Validate a list of IP addresses. On input, dwTypeIn MUST be set to               |
	//	|                                  | DNSSRV_TYPEID_IP_VALIDATE and pDataIn MUST point to a DNS_RPC_IP_VALIDATE        |
	//	|                                  | structure (section 2.2.3.2.4) containing a list of IP addresses to be validated  |
	//	|                                  | and the context information for validation as specified in section 2.2.3.2.4.    |
	//	|                                  | Unless an error is returned, on output the value pointed to by pdwTypeOut MUST   |
	//	|                                  | be set to DNSSRV_TYPEID_ADDRARRAY and ppDataOut MUST point to a structure of     |
	//	|                                  | type DNS_ADDR_ARRAY (section 2.2.3.2.3) that contains IP validation results      |
	//	|                                  | (section 2.2.3.2.1).                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ModifyZoneSigningKeyDescriptor   | Creates, deletes, or modifies a signing key descriptor (SKD) for the specified   |
	//	|                                  | zone. On input, dwTypeIn MUST be set to DNSSRV_TYPEID_SKD and pDataIn MUST       |
	//	|                                  | point to a structure of type DNS_RPC_SKD (section 2.2.6.2.1). If GUID inside     |
	//	|                                  | DNS_RPC_SKD is set to zero, the server MUST create a new signing key descriptor. |
	//	|                                  | If GUID inside DNS_RPC_SKD is set to a nonzero value and if all other fields     |
	//	|                                  | in the structure are NULL, the server MUST delete the signing key descriptor     |
	//	|                                  | from the zone. Otherwise, the server MUST modify the signing key descriptor for  |
	//	|                                  | the specified zone. Unless an error is returned, on output the value pointed     |
	//	|                                  | to by pdwTypeOut MUST be set to DNSSRV_TYPEID_SKD and ppDataOut MUST point to a  |
	//	|                                  | structure of type DNS_RPC_SKD (section 2.2.6.2.1).                               |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumZoneSigningKeyDescriptors    | Retrieves the signing key descriptor found in the zone's signing key descriptor  |
	//	|                                  | list for the specified zone. Input parameters (dwTypeIn and pDataIn) are         |
	//	|                                  | ignored. Unless an error is returned, on output the value pointed to by          |
	//	|                                  | pdwTypeOut MUST be set to DNSSRV_TYPEID_SKD_LIST and ppDataOut MUST point to a   |
	//	|                                  | structure of type DNS_RPC_SKD_LIST (section 2.2.6.2.2).                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| GetZoneSigningKeyDescriptorState | Retrieves the SKD state (section 2.2.6.2.3) for the specified zone               |
	//	|                                  | and the signing key descriptor GUID. On input, dwTypeIn MUST be set to           |
	//	|                                  | DNSSRV_TYPEID_LPWSTR and pDataIn MUST point to a string containing the GUID of   |
	//	|                                  | the signing key descriptor. Unless an error is returned, on output the value     |
	//	|                                  | pointed to by pdwTypeOut MUST be set to DNSSRV_TYPEID_SKD_STATE and ppDataOut    |
	//	|                                  | MUST point to a structure of type DNS_RPC_SKD_STATE (section 2.2.6.2.3).         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| SetZoneSigningKeyDescriptorState | Modifies the SKD state (section 2.2.6.2.3) for the specified zone and            |
	//	|                                  | the signing key descriptor GUID. On input, dwTypeIn MUST be set to               |
	//	|                                  | DNSSRV_TYPEID_SKD_STATE and pDataIn MUST point to a structure of type            |
	//	|                                  | DNS_RPC_SKD_STATE (section 2.2.6.2.3). Note that only one key pointer string     |
	//	|                                  | inside DNS_RPC_SKD_STATE will be set per a specific operation as described       |
	//	|                                  | below. Unless an error is returned, on output the value pointed to by pdwTypeOut |
	//	|                                  | MUST be set to DNSSRV_TYPEID_SKD_STATE and ppDataOut MUST point to a structure   |
	//	|                                  | of type DNS_RPC_SKD_STATE (section 2.2.6.2.3) containing the modified SKD state. |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ValidateZoneSigningParameters    | Validates the zone signing parameters and returns a structure describing the     |
	//	|                                  | invalid signing parameters. Input parameters (dwTypeIn and pDataIn) are ignored. |
	//	|                                  | Unless success is returned, on output the value pointed to by pdwTypeOut MUST    |
	//	|                                  | be set to DNSSRV_TYPEID_SIGNING_VALIDATION_ERROR and ppDataOut MUST point        |
	//	|                                  | to a structure of type DNS_RPC_SIGNING_VALIDATION_ERROR (section 2.2.6.2.8)      |
	//	|                                  | containing invalid elements of the zone configuration.                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumerateKeyStorageProviders     | Enumerates key storage providers installed on the DNS server. On input, dwTypeIn |
	//	|                                  | and pDataIn are ignored. Unless an error is returned, on output the value        |
	//	|                                  | pointed to by pdwTypeOut MUST be set to DNSSRV_TYPEID_UNICODE_STRING_LIST and    |
	//	|                                  | ppDataOut MUST point to a structure of type DNS_RPC_UNICODE_STRING_LIST (section |
	//	|                                  | 2.2.1.2.4) that contains a list of storage providers installed on the DNS        |
	//	|                                  | server.                                                                          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumerateTrustPoints             | Retrieve a list of trust points, containing either all trust points or           |
	//	|                                  | only those at or below a given FQDN. On input, dwTypeIn MUST be set to           |
	//	|                                  | DNSSRV_TYPEID_LPSTR and pDataIn MUST point either to a null pointer or to        |
	//	|                                  | a null-terminated UTF-8 string specifying an FQDN. If pDataIn points to a        |
	//	|                                  | null pointer, the server MUST return all of the trust points. If pDataIn is      |
	//	|                                  | an FQDN and there is a trust point or parent of a trust point at the FQDN,       |
	//	|                                  | the server MUST return the trust point at the FQDN (or an empty trust-point      |
	//	|                                  | structure for the FQDN if the FQDN is not a trust point) followed by empty       |
	//	|                                  | trust-point structures for each immediate child of the FQDN, if any. An          |
	//	|                                  | empty trust-point structure is a structure in which eTrustPointState is          |
	//	|                                  | TRUST_POINT_STATE_INITIALIZED and all elements other than pszTrustPointName and  |
	//	|                                  | dwRpcStructureVersion are zero. If pDataIn is an FQDN and there is neither a     |
	//	|                                  | trust point nor the parent of a trust point at the FQDN, the server MUST return  |
	//	|                                  | a nonzero error. Unless an error is returned, on output the value pointed to     |
	//	|                                  | by pdwTypeOut MUST be set to DNSSRV_TYPEID_TRUST_POINT_LIST and ppDataOut MUST   |
	//	|                                  | point to a structure of type DNS_RPC_TRUST_POINT_LIST (section 2.2.6.2.5).       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumerateTrustAnchors            | Retrieve a list of the trust anchors at a given FQDN. On input, dwTypeIn MUST    |
	//	|                                  | be set to DNSSRV_TYPEID_LPSTR and pDataIn MUST point to a null-terminated UTF-8  |
	//	|                                  | string specifying an FQDN. If the FQDN specified is not a trust point, the       |
	//	|                                  | server MUST return a nonzero error. Unless an error is returned, on output the   |
	//	|                                  | value pointed to by pdwTypeOut MUST be set to DNSSRV_TYPEID_TRUST_ANCHOR_LIST    |
	//	|                                  | and ppDataOut MUST point to a structure of type DNS_RPC_TRUST_ANCHOR_LIST        |
	//	|                                  | (section 2.2.6.2.7).                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ExportZoneSigningSettings        | Exports all the Dnssec settings of a file-backed primary zone from a server.     |
	//	|                                  | On input, dwTypeIn MUST be set to DNSSRV_TYPEID_DWORD and pDataIn SHOULD be      |
	//	|                                  | 1 to get KSK details in the exported DNS_RPC_ZONE_DNSSEC_SETTINGS structure;     |
	//	|                                  | otherwise, pDataIn SHOULD be zero. Unless an error is returned, on output        |
	//	|                                  | pdwTypeOut is set to DNSSRV_TYPEID_ZONE_SIGNING_SETTINGS and ppDataOut points to |
	//	|                                  | a structure of type PDNS_RPC_ZONE_DNSSEC_SETTINGS.                               |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ImportZoneSigningSettings        | Imports the Dnssec settings to a primary file-backed zone on a server and        |
	//	|                                  | takes appropriate action based on the signing metadata imported. On input,       |
	//	|                                  | dwTypeIn MUST be set to DNSSRV_TYPEID_ZONE_SIGNING_SETTINGS and pDataIn SHOULD   |
	//	|                                  | be a structure of type PDNS_RPC_ZONE_DNSSEC_SETTINGS. If this operation is       |
	//	|                                  | invoked on a server that hosts a primary unsigned copy of a file-backed zone     |
	//	|                                  | and the fIsSigned Property of PDNS_RPC_ZONE_DNSSEC_SETTINGS is 1, then the       |
	//	|                                  | server becomes a nonkey master primary server for that zone. Unless an error     |
	//	|                                  | is returned, on output pdwTypeOut is set to DNSSRV_TYPEID_DWORD and ppDataOut    |
	//	|                                  | points to a structure of type ImportOpResult.                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumZoneScopes                   | Enumerates all the zone scopes in a zone or cache scopes in a cache zone. On     |
	//	|                                  | input, dwTypeIn MUST be set to DNSSRV_TYPEID_NULL and pDataIn SHOULD be NULL.    |
	//	|                                  | The pszZone MUST be the zone name for which zone scopes are to be enumerated or  |
	//	|                                  | it MUST be "..cache". Unless an error is returned, on output pdwTypeOut is set   |
	//	|                                  | to DNSSRV_TYPEID_ZONE_SCOPE_ENUM and ppDataOut points to a structure of type     |
	//	|                                  | PDNS_RPC_ENUM_ZONE_SCOPE_LIST.                                                   |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| ZoneStatistics                   | Gets the zone statistics from the server. On input, dwTypeIn MUST be set to      |
	//	|                                  | DNSSRV_TYPEID_DWORD and pDataIn SHOULD be either DNS_RPC_ZONE_STATS_GET, which   |
	//	|                                  | gets the current zone statistics, or DNS_RPC_ZONE_STATS_CLEAR, which clears      |
	//	|                                  | the zone statistics after getting them. The pszZone MUST point to the zone       |
	//	|                                  | information for which statistics are required. Unless an error is returned, on   |
	//	|                                  | output pdwTypeOut SHOULD be set to DNSSRV_TYPEID_ZONE_STATS and ppDataOut SHOULD |
	//	|                                  | point to a structure of type DNS_RPC_ZONE_STATS_V1 (section 2.2.12.2.5).         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumServerScopes                 | Enumerates all the server scopes in a DNS server. On input, dwTypeIn MUST be set |
	//	|                                  | to DNSSRV_TYPEID_NULL and pDataIn SHOULD be NULL. Unless an error is returned,   |
	//	|                                  | on output, pdwTypeOut SHOULD be set to DNSSRV_TYPEID_SCOPE_ENUM and ppDataOut    |
	//	|                                  | points to a structure of type PDNS_RPC_ENUM_SCOPE_LIST.<276>                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumerateClientSubnetRecord      | Enumerates the names all the Client Subnet Records on the DNS server. On         |
	//	|                                  | input, dwTypeIn MUST be set to DNSSRV_TYPEID_NULL and pDataIn SHOULD be          |
	//	|                                  | NULL. Unless an error is returned, on output pdwTypeOut SHOULD be set to         |
	//	|                                  | DNSSRV_TYPEID_UNICODE_STRING_LIST and ppDataOut points to a structure of type    |
	//	|                                  | PDNS_RPC_UNICODE_STRING_LIST.                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| GetClientSubnetRecord            | Gets the details of the Client Subnet Record on the DNS server. On input,        |
	//	|                                  | dwTypeIn MUST be set to DNSSRV_TYPEID_LPWSTR and pDataIn SHOULD be name of the   |
	//	|                                  | Client Subnet Record. Unless an error is returned, on output pdwTypeOut SHOULD   |
	//	|                                  | be set to DNSSRV_TYPEID_CLIENT_SUBNET_RECORD and ppDataOut points to a structure |
	//	|                                  | of type PDNS_RPC_CLIENT_SUBNET_RECORD.                                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumeratePolicy                  | Enumerates the policies configured on the server level or zone level             |
	//	|                                  | for a specified zone on a DNS server. On input, dwTypeIn MUST be set to          |
	//	|                                  | DNSSRV_TYPEID_NULL and pDataIn SHOULD be NULL. Unless an error is returned,      |
	//	|                                  | on output pdwTypeOut SHOULD be set to DNSSRV_TYPEID_POLICY_ENUM and ppDataOut    |
	//	|                                  | points to a structure of type PDNS_RPC_ENUMERATE_POLICY_LIST.                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| GetPolicy                        | Gets the details of a DNS Policy configured on the server level or on the zone   |
	//	|                                  | level for a specified zone on the DNS server. On input, dwTypeIn MUST be set     |
	//	|                                  | to DNSSRV_TYPEID_LPWSTR and pDataIn SHOULD be name of the DNS Policy. Unless an  |
	//	|                                  | error is returned, on output pdwTypeOut SHOULD be set to DNSSRV_TYPEID_POLICY    |
	//	|                                  | and ppDataOut points to a structure of type PDNS_RPC_POLICY.                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| GetRRLInfo                       | Gets the details of Response Rate Limiting parameters configured on the server   |
	//	|                                  | level on the DNS server. On input, dwTypeIn MUST be set to DNSSRV_TYPEID_NULL    |
	//	|                                  | and pDataIn SHOULD be NULL. Unless an error is returned, on output pdwTypeOut    |
	//	|                                  | SHOULD be set to DNSSRV_TYPEID_RRL and ppDataOut points to a structure of type   |
	//	|                                  | PDNS_RPC_RRL_PARAMS.                                                             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| EnumVirtualizationInstances      | Enumerates the virtualization instance present in DNS server. The                |
	//	|                                  | dwTypeIn parameter MUST be set to DNSSRV_TYPEID_NULL and pDataIn                 |
	//	|                                  | MUST be set to NULL. On successful enumeration pdwTypeOut is set to              |
	//	|                                  | DNSSRV_TYPEID_VIRTUALIZATION_INSTANCE_ENUM, and ppDataOut MUST point to a        |
	//	|                                  | structure of type DNS_RPC_ENUM_VIRTUALIZATION_INSTANCE_LIST.                     |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	Operation string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	// dwTypeIn: A DNS_RPC_TYPEID (section 2.2.1.1.1) value indicating the type of input
	// data pointed to by pDataIn.
	TypeIn uint32 `idl:"name:dwTypeIn" json:"type_in"`
	// pDataIn: Input data of type DNSSRV_RPC_UNION (section 2.2.1.2.6), which contains
	// a data structure of the type indicated by dwTypeIn.
	DataIn *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
}

func (o *ComplexOperationRequest) xxx_ToOp(ctx context.Context) *xxx_ComplexOperationOperation {
	if o == nil {
		return &xxx_ComplexOperationOperation{}
	}
	return &xxx_ComplexOperationOperation{
		ServerName: o.ServerName,
		Zone:       o.Zone,
		Operation:  o.Operation,
		TypeIn:     o.TypeIn,
		DataIn:     o.DataIn,
	}
}

func (o *ComplexOperationRequest) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperationOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Operation = op.Operation
	o.TypeIn = op.TypeIn
	o.DataIn = op.DataIn
}
func (o *ComplexOperationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ComplexOperationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComplexOperationResponse structure represents the R_DnssrvComplexOperation operation response
type ComplexOperationResponse struct {
	// pdwTypeOut: A pointer to a DWORD that on success returns a DNS_RPC_TYPEID (section
	// 2.2.1.1.1) value indicating the type of output data pointed to by ppDataOut.
	TypeOut uint32 `idl:"name:pdwTypeOut" json:"type_out"`
	// ppDataOut: A pointer to output data of type DNSSRV_RPC_UNION, which on success contains
	// a data structure of the type indicated by pdwTypeOut.
	DataOut *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	// Return: The R_DnssrvComplexOperation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ComplexOperationResponse) xxx_ToOp(ctx context.Context) *xxx_ComplexOperationOperation {
	if o == nil {
		return &xxx_ComplexOperationOperation{}
	}
	return &xxx_ComplexOperationOperation{
		TypeOut: o.TypeOut,
		DataOut: o.DataOut,
		Return:  o.Return,
	}
}

func (o *ComplexOperationResponse) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperationOperation) {
	if o == nil {
		return
	}
	o.TypeOut = op.TypeOut
	o.DataOut = op.DataOut
	o.Return = op.Return
}
func (o *ComplexOperationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ComplexOperationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRecordsOperation structure represents the R_DnssrvEnumRecords operation
type xxx_EnumRecordsOperation struct {
	ServerName   string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName     string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild   string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType   uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag   uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart  string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop   string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
	BufferLength uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer       []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	Return       int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRecordsOperation) OpNum() int { return 3 }

func (o *xxx_EnumRecordsOperation) OpName() string { return "/DnsServer/v5/R_DnssrvEnumRecords" }

func (o *xxx_EnumRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.NodeName != "" {
			_ptr_pszNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_pszNodeName); err != nil {
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
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.StartChild != "" {
			_ptr_pszStartChild := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.StartChild); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.StartChild, _ptr_pszStartChild); err != nil {
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
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStart != "" {
			_ptr_pszFilterStart := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStart); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStart, _ptr_pszFilterStart); err != nil {
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
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStop != "" {
			_ptr_pszFilterStop := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStop); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStop, _ptr_pszFilterStop); err != nil {
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

func (o *xxx_EnumRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszNodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_pszNodeName, _ptr_pszNodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszStartChild := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.StartChild); err != nil {
				return err
			}
			return nil
		})
		_s_pszStartChild := func(ptr interface{}) { o.StartChild = *ptr.(*string) }
		if err := w.ReadPointer(&o.StartChild, _s_pszStartChild, _ptr_pszStartChild); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStart := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStart); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStart := func(ptr interface{}) { o.FilterStart = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStart, _s_pszFilterStart, _ptr_pszFilterStart); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStop := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStop); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStop := func(ptr interface{}) { o.FilterStop = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStop, _s_pszFilterStop, _ptr_pszFilterStop); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwBufferLength {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		if o.Buffer != nil || o.BufferLength > 0 {
			_ptr_ppBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
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
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuffer); err != nil {
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
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwBufferLength {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		_ptr_ppBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
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
			return nil
		})
		_s_ppBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuffer, _ptr_ppBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRecordsRequest structure represents the R_DnssrvEnumRecords operation request
type EnumRecordsRequest struct {
	// pwszServerName: The client SHOULD pass a pointer to the FQDN of the target server
	// as a null-terminated UTF-16LE character string. The server MUST ignore this value.
	ServerName string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pszZone: A pointer to a null-terminated character string that contains the name of
	// the zone to be queried. For operations specific to a particular zone, this field
	// MUST contain the name of the zone in UTF-8 format. For all other operations, this
	// field MUST be NULL.
	Zone string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pszNodeName: A pointer to a null-terminated character string that contains the node
	// name at which to modify a record. A string that is not dot-terminated indicates a
	// name relative to the zone root. A value of "@" indicates the zone root itself. A
	// dot-terminated string indicates the name is an FQDN.
	NodeName string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	// pszStartChild: A pointer to a null-terminated character string that contains the
	// name of the child node at which to start enumeration. A NULL value indicates to start
	// a new record enumeration. The client application can pass the last retrieved child
	// node of pszNodeName to continue a previous enumeration.
	StartChild string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	// wRecordType: An integer value that indicates the type of record to enumerate. Any
	// value can be used, as specified in DNS_RECORD_TYPE (section 2.2.2.1.1). The query-only
	// value DNS_TYPE_ALL indicates all types of records.
	RecordType uint16 `idl:"name:wRecordType" json:"record_type"`
	// fSelectFlag: An integer value that specifies what records are included in the response.
	// Any combination of the values below MUST be supported. Values not listed below MUST
	// be ignored.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_AUTHORITY_DATA 0x00000001  | Include records from authoritative zones.                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_CACHE_DATA 0x00000002      | Include records from the DNS server's cache.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_GLUE_DATA 0x00000004       | Include glue records.                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_ROOT_HINT_DATA 0x00000008  | Include root hint records.                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_ADDITIONAL_DATA 0x00000010 | Include additional records.                                                      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_NO_CHILDREN 0x00010000     | Do not include any records from child nodes.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_RPC_VIEW_ONLY_CHILDREN 0x00020000   | Include only children nodes of the specified node in the results. For example:   |
	//	|                                         | if a zone, "example.com", has child nodes, "a.example.com" and "b.example.com",  |
	//	|                                         | calling R_DnssrcEnumRecords(…,"example.com", "example.com", NULL, DNS_TYPE_ALL,  |
	//	|                                         | DNS_RPC_VIEW_ONLY_CHILDREN, …, …, …, …) will return DNS_RPC_NODES for "a" and    |
	//	|                                         | "b".                                                                             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	SelectFlag uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	// pszFilterStart: Reserved for future use only. This MUST be set to NULL by clients
	// and ignored by servers.
	FilterStart string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	// pszFilterStop: Reserved for future use only. This MUST be set to NULL by clients
	// and ignored by servers.
	FilterStop string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
}

func (o *EnumRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumRecordsOperation {
	if o == nil {
		return &xxx_EnumRecordsOperation{}
	}
	return &xxx_EnumRecordsOperation{
		ServerName:  o.ServerName,
		Zone:        o.Zone,
		NodeName:    o.NodeName,
		StartChild:  o.StartChild,
		RecordType:  o.RecordType,
		SelectFlag:  o.SelectFlag,
		FilterStart: o.FilterStart,
		FilterStop:  o.FilterStop,
	}
}

func (o *EnumRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumRecordsOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.NodeName = op.NodeName
	o.StartChild = op.StartChild
	o.RecordType = op.RecordType
	o.SelectFlag = op.SelectFlag
	o.FilterStart = op.FilterStart
	o.FilterStop = op.FilterStop
}
func (o *EnumRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRecordsResponse structure represents the R_DnssrvEnumRecords operation response
type EnumRecordsResponse struct {
	// pdwBufferLength: A pointer to an integer that on success contains the length of the
	// buffer pointed to by ppBuffer.
	BufferLength uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	// ppBuffer: A pointer to a pointer that on success points to a buffer containing the
	// enumerated records. The buffer is a series of structures beginning with a DNS_RPC_NODE
	// structure (section 2.2.2.2.3). The records for the node will be represented by a
	// series of DNS_RPC_RECORD (section 2.2.2.2.5) structures. The number of DNS_RPC_RECORD
	// structures following a DNS_RPC_NODE structure is given by the wRecordCount member
	// of DNS_RPC_NODE.
	Buffer []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	// Return: The R_DnssrvEnumRecords return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumRecordsOperation {
	if o == nil {
		return &xxx_EnumRecordsOperation{}
	}
	return &xxx_EnumRecordsOperation{
		BufferLength: o.BufferLength,
		Buffer:       o.Buffer,
		Return:       o.Return,
	}
}

func (o *EnumRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumRecordsOperation) {
	if o == nil {
		return
	}
	o.BufferLength = op.BufferLength
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *EnumRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateRecordOperation structure represents the R_DnssrvUpdateRecord operation
type xxx_UpdateRecordOperation struct {
	ServerName   string       `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName     string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord    *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
	Return       int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateRecordOperation) OpNum() int { return 4 }

func (o *xxx_UpdateRecordOperation) OpName() string { return "/DnsServer/v5/R_DnssrvUpdateRecord" }

func (o *xxx_UpdateRecordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszNodeName {in} (1:{string, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.AddRecord != nil {
			_ptr_pAddRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AddRecord != nil {
					if err := o.AddRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AddRecord, _ptr_pAddRecord); err != nil {
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
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.DeleteRecord != nil {
			_ptr_pDeleteRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DeleteRecord != nil {
					if err := o.DeleteRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeleteRecord, _ptr_pDeleteRecord); err != nil {
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

func (o *xxx_UpdateRecordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, alias=LPCSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pAddRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AddRecord == nil {
				o.AddRecord = &dnsp.Record{}
			}
			if err := o.AddRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pAddRecord := func(ptr interface{}) { o.AddRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.AddRecord, _s_pAddRecord, _ptr_pAddRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pDeleteRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DeleteRecord == nil {
				o.DeleteRecord = &dnsp.Record{}
			}
			if err := o.DeleteRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pDeleteRecord := func(ptr interface{}) { o.DeleteRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.DeleteRecord, _s_pDeleteRecord, _ptr_pDeleteRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateRecordRequest structure represents the R_DnssrvUpdateRecord operation request
type UpdateRecordRequest struct {
	// pwszServerName: The client SHOULD pass a pointer to the FQDN of the target server
	// as a null-terminated UTF-16LE character string. The server MUST ignore this value.
	ServerName string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pszZone: A pointer to a null-terminated character string that contains the name of
	// the zone to be queried. For operations specific to a particular zone, this field
	// MUST contain the name of the zone in UTF-8 format. For all other operations, this
	// field MUST be NULL.
	Zone string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pszNodeName: A pointer to a null-terminated character string that contains the node
	// name at which to modify a record. A string that is not dot-terminated indicates a
	// name relative to the zone root. A value of "@" indicates the zone root itself. A
	// dot-terminated string indicates the name is an FQDN.
	NodeName string `idl:"name:pszNodeName;string" json:"node_name"`
	// pAddRecord: A pointer to a structure of type DNS_RPC_RECORD (section 2.2.2.2.5) that
	// contains information based on the operation being performed as specified below.
	AddRecord *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	// pDeleteRecord: A pointer to a structure of type DNS_RPC_RECORD (section 2.2.2.2.5)
	// that contains information based on the operation being performed as specified below.
	//
	// * pAddRecord: The DNS RR data of the new record.
	//
	// * pDeleteRecord: MUST be set to NULL.
	//
	// * pAddRecord: MUST be set to NULL.
	//
	// * pDeleteRecord: Individual DNS RR data of the record to be deleted.
	//
	// * pAddRecord: New record data.
	//
	// * pDeleteRecord: Old record data.
	//
	// * pAddRecord: MUST be set to NULL.
	//
	// * pDeleteRecord: MUST be set to NULL.
	DeleteRecord *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
}

func (o *UpdateRecordRequest) xxx_ToOp(ctx context.Context) *xxx_UpdateRecordOperation {
	if o == nil {
		return &xxx_UpdateRecordOperation{}
	}
	return &xxx_UpdateRecordOperation{
		ServerName:   o.ServerName,
		Zone:         o.Zone,
		NodeName:     o.NodeName,
		AddRecord:    o.AddRecord,
		DeleteRecord: o.DeleteRecord,
	}
}

func (o *UpdateRecordRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecordOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.NodeName = op.NodeName
	o.AddRecord = op.AddRecord
	o.DeleteRecord = op.DeleteRecord
}
func (o *UpdateRecordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UpdateRecordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateRecordResponse structure represents the R_DnssrvUpdateRecord operation response
type UpdateRecordResponse struct {
	// Return: The R_DnssrvUpdateRecord return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateRecordResponse) xxx_ToOp(ctx context.Context) *xxx_UpdateRecordOperation {
	if o == nil {
		return &xxx_UpdateRecordOperation{}
	}
	return &xxx_UpdateRecordOperation{
		Return: o.Return,
	}
}

func (o *UpdateRecordResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecordOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateRecordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UpdateRecordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Operation2Operation structure represents the R_DnssrvOperation2 operation
type xxx_Operation2Operation struct {
	ClientVersion uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Context       uint32      `idl:"name:dwContext" json:"context"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID        uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data          *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
	Return        int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Operation2Operation) OpNum() int { return 5 }

func (o *xxx_Operation2Operation) OpName() string { return "/DnsServer/v5/R_DnssrvOperation2" }

func (o *xxx_Operation2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Operation2Request structure represents the R_DnssrvOperation2 operation request
type Operation2Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use. MUST be set to zero by clients and MUST
	// be ignored by servers.
	SettingFlags uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Context      uint32      `idl:"name:dwContext" json:"context"`
	Operation    string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID       uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data         *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
}

func (o *Operation2Request) xxx_ToOp(ctx context.Context) *xxx_Operation2Operation {
	if o == nil {
		return &xxx_Operation2Operation{}
	}
	return &xxx_Operation2Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		Context:       o.Context,
		Operation:     o.Operation,
		TypeID:        o.TypeID,
		Data:          o.Data,
	}
}

func (o *Operation2Request) xxx_FromOp(ctx context.Context, op *xxx_Operation2Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Context = op.Context
	o.Operation = op.Operation
	o.TypeID = op.TypeID
	o.Data = op.Data
}
func (o *Operation2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Operation2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Operation2Response structure represents the R_DnssrvOperation2 operation response
type Operation2Response struct {
	// Return: The R_DnssrvOperation2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Operation2Response) xxx_ToOp(ctx context.Context) *xxx_Operation2Operation {
	if o == nil {
		return &xxx_Operation2Operation{}
	}
	return &xxx_Operation2Operation{
		Return: o.Return,
	}
}

func (o *Operation2Response) xxx_FromOp(ctx context.Context, op *xxx_Operation2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *Operation2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Operation2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Query2Operation structure represents the R_DnssrvQuery2 operation
type xxx_Query2Operation struct {
	ClientVersion uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID        uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data          *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	Return        int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Query2Operation) OpNum() int { return 6 }

func (o *xxx_Query2Operation) OpName() string { return "/DnsServer/v5/R_DnssrvQuery2" }

func (o *xxx_Query2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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

func (o *xxx_Query2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeId {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeId {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Query2Request structure represents the R_DnssrvQuery2 operation request
type Query2Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use only. This field MUST be set to zero by clients
	// and ignored by servers.
	SettingFlags uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation    string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
}

func (o *Query2Request) xxx_ToOp(ctx context.Context) *xxx_Query2Operation {
	if o == nil {
		return &xxx_Query2Operation{}
	}
	return &xxx_Query2Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		Operation:     o.Operation,
	}
}

func (o *Query2Request) xxx_FromOp(ctx context.Context, op *xxx_Query2Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Operation = op.Operation
}
func (o *Query2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Query2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Query2Response structure represents the R_DnssrvQuery2 operation response
type Query2Response struct {
	TypeID uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data   *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	// Return: The R_DnssrvQuery2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Query2Response) xxx_ToOp(ctx context.Context) *xxx_Query2Operation {
	if o == nil {
		return &xxx_Query2Operation{}
	}
	return &xxx_Query2Operation{
		TypeID: o.TypeID,
		Data:   o.Data,
		Return: o.Return,
	}
}

func (o *Query2Response) xxx_FromOp(ctx context.Context, op *xxx_Query2Operation) {
	if o == nil {
		return
	}
	o.TypeID = op.TypeID
	o.Data = op.Data
	o.Return = op.Return
}
func (o *Query2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Query2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComplexOperation2Operation structure represents the R_DnssrvComplexOperation2 operation
type xxx_ComplexOperation2Operation struct {
	ClientVersion uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeIn        uint32      `idl:"name:dwTypeIn" json:"type_in"`
	DataIn        *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
	TypeOut       uint32      `idl:"name:pdwTypeOut" json:"type_out"`
	DataOut       *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	Return        int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_ComplexOperation2Operation) OpNum() int { return 7 }

func (o *xxx_ComplexOperation2Operation) OpName() string {
	return "/DnsServer/v5/R_DnssrvComplexOperation2"
}

func (o *xxx_ComplexOperation2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataIn := uint32(o.TypeIn)
		if o.DataIn != nil {
			if err := o.DataIn.MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataIn == nil {
			o.DataIn = &dnsp.Union{}
		}
		_swDataIn := uint32(o.TypeIn)
		if err := o.DataIn.UnmarshalUnionNDR(ctx, w, _swDataIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeOut {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataOut := uint32(o.TypeOut)
		if o.DataOut != nil {
			if err := o.DataOut.MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeOut {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataOut == nil {
			o.DataOut = &dnsp.Union{}
		}
		_swDataOut := uint32(o.TypeOut)
		if err := o.DataOut.UnmarshalUnionNDR(ctx, w, _swDataOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ComplexOperation2Request structure represents the R_DnssrvComplexOperation2 operation request
type ComplexOperation2Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use only. This field MUST be set to zero by clients
	// and ignored by servers.
	SettingFlags uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation    string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeIn       uint32      `idl:"name:dwTypeIn" json:"type_in"`
	DataIn       *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
}

func (o *ComplexOperation2Request) xxx_ToOp(ctx context.Context) *xxx_ComplexOperation2Operation {
	if o == nil {
		return &xxx_ComplexOperation2Operation{}
	}
	return &xxx_ComplexOperation2Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		Operation:     o.Operation,
		TypeIn:        o.TypeIn,
		DataIn:        o.DataIn,
	}
}

func (o *ComplexOperation2Request) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperation2Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.Operation = op.Operation
	o.TypeIn = op.TypeIn
	o.DataIn = op.DataIn
}
func (o *ComplexOperation2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ComplexOperation2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperation2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComplexOperation2Response structure represents the R_DnssrvComplexOperation2 operation response
type ComplexOperation2Response struct {
	TypeOut uint32      `idl:"name:pdwTypeOut" json:"type_out"`
	DataOut *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	// Return: The R_DnssrvComplexOperation2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ComplexOperation2Response) xxx_ToOp(ctx context.Context) *xxx_ComplexOperation2Operation {
	if o == nil {
		return &xxx_ComplexOperation2Operation{}
	}
	return &xxx_ComplexOperation2Operation{
		TypeOut: o.TypeOut,
		DataOut: o.DataOut,
		Return:  o.Return,
	}
}

func (o *ComplexOperation2Response) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperation2Operation) {
	if o == nil {
		return
	}
	o.TypeOut = op.TypeOut
	o.DataOut = op.DataOut
	o.Return = op.Return
}
func (o *ComplexOperation2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ComplexOperation2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperation2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRecords2Operation structure represents the R_DnssrvEnumRecords2 operation
type xxx_EnumRecords2Operation struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName      string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild    string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType    uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag    uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart   string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop    string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
	BufferLength  uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer        []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	Return        int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRecords2Operation) OpNum() int { return 8 }

func (o *xxx_EnumRecords2Operation) OpName() string { return "/DnsServer/v5/R_DnssrvEnumRecords2" }

func (o *xxx_EnumRecords2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.NodeName != "" {
			_ptr_pszNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_pszNodeName); err != nil {
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
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.StartChild != "" {
			_ptr_pszStartChild := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.StartChild); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.StartChild, _ptr_pszStartChild); err != nil {
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
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStart != "" {
			_ptr_pszFilterStart := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStart); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStart, _ptr_pszFilterStart); err != nil {
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
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStop != "" {
			_ptr_pszFilterStop := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStop); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStop, _ptr_pszFilterStop); err != nil {
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

func (o *xxx_EnumRecords2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszNodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_pszNodeName, _ptr_pszNodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszStartChild := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.StartChild); err != nil {
				return err
			}
			return nil
		})
		_s_pszStartChild := func(ptr interface{}) { o.StartChild = *ptr.(*string) }
		if err := w.ReadPointer(&o.StartChild, _s_pszStartChild, _ptr_pszStartChild); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStart := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStart); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStart := func(ptr interface{}) { o.FilterStart = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStart, _s_pszFilterStart, _ptr_pszFilterStart); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStop := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStop); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStop := func(ptr interface{}) { o.FilterStop = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStop, _s_pszFilterStop, _ptr_pszFilterStop); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwBufferLength {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		if o.Buffer != nil || o.BufferLength > 0 {
			_ptr_ppBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
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
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuffer); err != nil {
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
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwBufferLength {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		_ptr_ppBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
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
			return nil
		})
		_s_ppBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuffer, _ptr_ppBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRecords2Request structure represents the R_DnssrvEnumRecords2 operation request
type EnumRecords2Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use only. This field MUST be set to zero by clients
	// and ignored by servers.
	SettingFlags uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName     string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild   string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType   uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag   uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart  string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop   string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
}

func (o *EnumRecords2Request) xxx_ToOp(ctx context.Context) *xxx_EnumRecords2Operation {
	if o == nil {
		return &xxx_EnumRecords2Operation{}
	}
	return &xxx_EnumRecords2Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		NodeName:      o.NodeName,
		StartChild:    o.StartChild,
		RecordType:    o.RecordType,
		SelectFlag:    o.SelectFlag,
		FilterStart:   o.FilterStart,
		FilterStop:    o.FilterStop,
	}
}

func (o *EnumRecords2Request) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords2Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.NodeName = op.NodeName
	o.StartChild = op.StartChild
	o.RecordType = op.RecordType
	o.SelectFlag = op.SelectFlag
	o.FilterStart = op.FilterStart
	o.FilterStop = op.FilterStop
}
func (o *EnumRecords2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumRecords2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRecords2Response structure represents the R_DnssrvEnumRecords2 operation response
type EnumRecords2Response struct {
	BufferLength uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer       []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	// Return: The R_DnssrvEnumRecords2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRecords2Response) xxx_ToOp(ctx context.Context) *xxx_EnumRecords2Operation {
	if o == nil {
		return &xxx_EnumRecords2Operation{}
	}
	return &xxx_EnumRecords2Operation{
		BufferLength: o.BufferLength,
		Buffer:       o.Buffer,
		Return:       o.Return,
	}
}

func (o *EnumRecords2Response) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords2Operation) {
	if o == nil {
		return
	}
	o.BufferLength = op.BufferLength
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *EnumRecords2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumRecords2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateRecord2Operation structure represents the R_DnssrvUpdateRecord2 operation
type xxx_UpdateRecord2Operation struct {
	ClientVersion uint32       `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32       `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string       `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName      string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord     *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord  *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
	Return        int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateRecord2Operation) OpNum() int { return 9 }

func (o *xxx_UpdateRecord2Operation) OpName() string { return "/DnsServer/v5/R_DnssrvUpdateRecord2" }

func (o *xxx_UpdateRecord2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszNodeName {in} (1:{string, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.AddRecord != nil {
			_ptr_pAddRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AddRecord != nil {
					if err := o.AddRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AddRecord, _ptr_pAddRecord); err != nil {
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
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.DeleteRecord != nil {
			_ptr_pDeleteRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DeleteRecord != nil {
					if err := o.DeleteRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeleteRecord, _ptr_pDeleteRecord); err != nil {
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

func (o *xxx_UpdateRecord2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, alias=LPCSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pAddRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AddRecord == nil {
				o.AddRecord = &dnsp.Record{}
			}
			if err := o.AddRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pAddRecord := func(ptr interface{}) { o.AddRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.AddRecord, _s_pAddRecord, _ptr_pAddRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pDeleteRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DeleteRecord == nil {
				o.DeleteRecord = &dnsp.Record{}
			}
			if err := o.DeleteRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pDeleteRecord := func(ptr interface{}) { o.DeleteRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.DeleteRecord, _s_pDeleteRecord, _ptr_pDeleteRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateRecord2Request structure represents the R_DnssrvUpdateRecord2 operation request
type UpdateRecord2Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use only. This field MUST be set to zero by clients
	// and ignored by servers.
	SettingFlags uint32       `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string       `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone         string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	NodeName     string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord    *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
}

func (o *UpdateRecord2Request) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord2Operation {
	if o == nil {
		return &xxx_UpdateRecord2Operation{}
	}
	return &xxx_UpdateRecord2Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		NodeName:      o.NodeName,
		AddRecord:     o.AddRecord,
		DeleteRecord:  o.DeleteRecord,
	}
}

func (o *UpdateRecord2Request) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord2Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.NodeName = op.NodeName
	o.AddRecord = op.AddRecord
	o.DeleteRecord = op.DeleteRecord
}
func (o *UpdateRecord2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UpdateRecord2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateRecord2Response structure represents the R_DnssrvUpdateRecord2 operation response
type UpdateRecord2Response struct {
	// Return: The R_DnssrvUpdateRecord2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateRecord2Response) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord2Operation {
	if o == nil {
		return &xxx_UpdateRecord2Operation{}
	}
	return &xxx_UpdateRecord2Operation{
		Return: o.Return,
	}
}

func (o *UpdateRecord2Response) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateRecord2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UpdateRecord2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateRecord3Operation structure represents the R_DnssrvUpdateRecord3 operation
type xxx_UpdateRecord3Operation struct {
	ClientVersion uint32       `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32       `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string       `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope     string       `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName      string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord     *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord  *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
	Return        int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateRecord3Operation) OpNum() int { return 10 }

func (o *xxx_UpdateRecord3Operation) OpName() string { return "/DnsServer/v5/R_DnssrvUpdateRecord3" }

func (o *xxx_UpdateRecord3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScope != "" {
			_ptr_pwszZoneScope := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScope); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScope, _ptr_pwszZoneScope); err != nil {
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
	// pszNodeName {in} (1:{string, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.AddRecord != nil {
			_ptr_pAddRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AddRecord != nil {
					if err := o.AddRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AddRecord, _ptr_pAddRecord); err != nil {
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
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.DeleteRecord != nil {
			_ptr_pDeleteRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DeleteRecord != nil {
					if err := o.DeleteRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeleteRecord, _ptr_pDeleteRecord); err != nil {
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

func (o *xxx_UpdateRecord3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScope := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScope); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScope := func(ptr interface{}) { o.ZoneScope = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScope, _s_pwszZoneScope, _ptr_pwszZoneScope); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, alias=LPCSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pAddRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AddRecord == nil {
				o.AddRecord = &dnsp.Record{}
			}
			if err := o.AddRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pAddRecord := func(ptr interface{}) { o.AddRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.AddRecord, _s_pAddRecord, _ptr_pAddRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pDeleteRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DeleteRecord == nil {
				o.DeleteRecord = &dnsp.Record{}
			}
			if err := o.DeleteRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pDeleteRecord := func(ptr interface{}) { o.DeleteRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.DeleteRecord, _s_pDeleteRecord, _ptr_pDeleteRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateRecord3Request structure represents the R_DnssrvUpdateRecord3 operation request
type UpdateRecord3Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pwszZoneScope: A pointer to a null-terminated character string that contains the
	// name of the zone scope or cache scope inside the zone to be queried. For operations
	// specific to a particular zone scope or cache scope, this field MUST contain the name
	// of the zone scope or cache scope. If the value is NULL, the API gives the same behavior
	// as R_DnssrvUpdateRecord2.
	ZoneScope    string       `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName     string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord    *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
}

func (o *UpdateRecord3Request) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord3Operation {
	if o == nil {
		return &xxx_UpdateRecord3Operation{}
	}
	return &xxx_UpdateRecord3Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		ZoneScope:     o.ZoneScope,
		NodeName:      o.NodeName,
		AddRecord:     o.AddRecord,
		DeleteRecord:  o.DeleteRecord,
	}
}

func (o *UpdateRecord3Request) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord3Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.ZoneScope = op.ZoneScope
	o.NodeName = op.NodeName
	o.AddRecord = op.AddRecord
	o.DeleteRecord = op.DeleteRecord
}
func (o *UpdateRecord3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UpdateRecord3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateRecord3Response structure represents the R_DnssrvUpdateRecord3 operation response
type UpdateRecord3Response struct {
	// Return: The R_DnssrvUpdateRecord3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateRecord3Response) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord3Operation {
	if o == nil {
		return &xxx_UpdateRecord3Operation{}
	}
	return &xxx_UpdateRecord3Operation{
		Return: o.Return,
	}
}

func (o *UpdateRecord3Response) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord3Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateRecord3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UpdateRecord3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRecords3Operation structure represents the R_DnssrvEnumRecords3 operation
type xxx_EnumRecords3Operation struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope     string `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName      string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild    string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType    uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag    uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart   string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop    string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
	BufferLength  uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer        []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	Return        int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRecords3Operation) OpNum() int { return 11 }

func (o *xxx_EnumRecords3Operation) OpName() string { return "/DnsServer/v5/R_DnssrvEnumRecords3" }

func (o *xxx_EnumRecords3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScope != "" {
			_ptr_pwszZoneScope := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScope); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScope, _ptr_pwszZoneScope); err != nil {
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
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.NodeName != "" {
			_ptr_pszNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_pszNodeName); err != nil {
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
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.StartChild != "" {
			_ptr_pszStartChild := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.StartChild); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.StartChild, _ptr_pszStartChild); err != nil {
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
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStart != "" {
			_ptr_pszFilterStart := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStart); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStart, _ptr_pszFilterStart); err != nil {
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
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStop != "" {
			_ptr_pszFilterStop := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStop); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStop, _ptr_pszFilterStop); err != nil {
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

func (o *xxx_EnumRecords3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScope := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScope); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScope := func(ptr interface{}) { o.ZoneScope = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScope, _s_pwszZoneScope, _ptr_pwszZoneScope); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszNodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_pszNodeName, _ptr_pszNodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszStartChild := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.StartChild); err != nil {
				return err
			}
			return nil
		})
		_s_pszStartChild := func(ptr interface{}) { o.StartChild = *ptr.(*string) }
		if err := w.ReadPointer(&o.StartChild, _s_pszStartChild, _ptr_pszStartChild); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStart := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStart); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStart := func(ptr interface{}) { o.FilterStart = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStart, _s_pszFilterStart, _ptr_pszFilterStart); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStop := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStop); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStop := func(ptr interface{}) { o.FilterStop = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStop, _s_pszFilterStop, _ptr_pszFilterStop); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwBufferLength {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		if o.Buffer != nil || o.BufferLength > 0 {
			_ptr_ppBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
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
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuffer); err != nil {
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
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwBufferLength {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		_ptr_ppBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
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
			return nil
		})
		_s_ppBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuffer, _ptr_ppBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRecords3Request structure represents the R_DnssrvEnumRecords3 operation request
type EnumRecords3Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pwszZoneScope: A pointer to a null-terminated character string that contains the
	// name of the zone scope inside the zone or cache scope inside the cache zone that
	// is to be queried. For operations specific to a particular zone scope or cache scope,
	// this field MUST contain the name of the zone scope or cache scope. If the value is
	// NULL, the API gives the same behavior as the R_DnssrvEnumRecords2 method.
	ZoneScope   string `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName    string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild  string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType  uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag  uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop  string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
}

func (o *EnumRecords3Request) xxx_ToOp(ctx context.Context) *xxx_EnumRecords3Operation {
	if o == nil {
		return &xxx_EnumRecords3Operation{}
	}
	return &xxx_EnumRecords3Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		ZoneScope:     o.ZoneScope,
		NodeName:      o.NodeName,
		StartChild:    o.StartChild,
		RecordType:    o.RecordType,
		SelectFlag:    o.SelectFlag,
		FilterStart:   o.FilterStart,
		FilterStop:    o.FilterStop,
	}
}

func (o *EnumRecords3Request) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords3Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.ZoneScope = op.ZoneScope
	o.NodeName = op.NodeName
	o.StartChild = op.StartChild
	o.RecordType = op.RecordType
	o.SelectFlag = op.SelectFlag
	o.FilterStart = op.FilterStart
	o.FilterStop = op.FilterStop
}
func (o *EnumRecords3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumRecords3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRecords3Response structure represents the R_DnssrvEnumRecords3 operation response
type EnumRecords3Response struct {
	BufferLength uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer       []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	// Return: The R_DnssrvEnumRecords3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRecords3Response) xxx_ToOp(ctx context.Context) *xxx_EnumRecords3Operation {
	if o == nil {
		return &xxx_EnumRecords3Operation{}
	}
	return &xxx_EnumRecords3Operation{
		BufferLength: o.BufferLength,
		Buffer:       o.Buffer,
		Return:       o.Return,
	}
}

func (o *EnumRecords3Response) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords3Operation) {
	if o == nil {
		return
	}
	o.BufferLength = op.BufferLength
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *EnumRecords3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumRecords3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Operation3Operation structure represents the R_DnssrvOperation3 operation
type xxx_Operation3Operation struct {
	ClientVersion uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName string      `idl:"name:pwszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Context       uint32      `idl:"name:dwContext" json:"context"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID        uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data          *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
	Return        int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Operation3Operation) OpNum() int { return 12 }

func (o *xxx_Operation3Operation) OpName() string { return "/DnsServer/v5/R_DnssrvOperation3" }

func (o *xxx_Operation3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScopeName != "" {
			_ptr_pwszZoneScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScopeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScopeName, _ptr_pwszZoneScopeName); err != nil {
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
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScopeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScopeName := func(ptr interface{}) { o.ZoneScopeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScopeName, _s_pwszZoneScopeName, _ptr_pwszZoneScopeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Operation3Request structure represents the R_DnssrvOperation3 operation request
type Operation3Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	// pwszZoneScopeName: A pointer to a null-terminated character string that contains
	// the name of the zone scope in the zone or cache scope in a cache zone, or a server
	// scope configured on the DNS server in which the operation is to be performed. For
	// operations specific to a particular zone scope, this field MUST contain the name
	// of the zone scope. If the value is NULL then the API gives the same behavior as R_DnssrvOperation2
	// (section 3.1.4.6). If the value is not NULL then pszZone MUST point to a null-terminated
	// character string that contains the name of the zone in UTF-8 format. In this case
	// the type of the zone pointed to by pszZone MUST be a primary zone. It MUST also be
	// a non-autocreated zone, a non-reverse lookup zone, and a non-AD integrated zone.
	// For operations specific to server scopes, this field MUST contain the name of the
	// server scope and the pszZone field MUST be NULL. For operations specific to cache
	// scopes, this field MUST contain the name of the cache scope, and the pszZone field
	// MUST be "..cache."
	//
	// If pszZone is not NULL and pwszZoneScopeName is not NULL, pszOperation MUST be set
	// to one of the following values:
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|                   |                                                                                  |
	//	|       VALUE       |                                     MEANING                                      |
	//	|                   |                                                                                  |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| "DeleteNode"      | Deletes a node from the given zone scope of a specified zone or a cache          |
	//	|                   | scope in a specified cache zone. On input, dwTypeId MUST be set to               |
	//	|                   | DNSSRV_TYPEID_NAME_AND_PARAM (section 2.2.2.1.1), and pData MUST point to a      |
	//	|                   | structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5). That structure     |
	//	|                   | contains the name of the node to be deleted as pointed by pszNodeName, and a     |
	//	|                   | Boolean flag in dwParam that indicates if the node's subtree is to be deleted.   |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| "DeleteRecordSet" | Deletes all the DNS records of a particular type at a particular node            |
	//	|                   | that is present in the given zone scope of a specified zone or a cache           |
	//	|                   | scope in a specified cache zone. On input, dwTypeId MUST be set to               |
	//	|                   | DNSSRV_TYPEID_NAME_AND_PARAM (section 2.2.2.1.1). pData MUST point to a          |
	//	|                   | structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5). That structure     |
	//	|                   | contains the name of the node to be deleted and the DNS record type in the       |
	//	|                   | dwParam member. The type MUST be of value of either DNS_RECORD_TYPE (section     |
	//	|                   | 2.2.2.1.1) or 0x00FF, which specifies all types.                                 |
	//	+-------------------+----------------------------------------------------------------------------------+
	//
	// If pszZone is NULL and pwszZoneScopeName is not NULL, pszOperation MUST be set to
	// one of the following values.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| "Forwarders"         | On input, dwTypeId SHOULD<286> be set to DNSSRV_TYPEID_FORWARDERS, and pData     |
	//	|                      | MUST point to a structure of one of the types specified in DNS_RPC_FORWARDERS    |
	//	|                      | (section 2.2.5.2.10), which contains information about new IP addresses to which |
	//	|                      | the DNS server SHOULD forward queries. These IP addresses are part of the server |
	//	|                      | scope as specified in the pwszZoneScopeName.                                     |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| "ResetDwordProperty" | Update the value of a (name, value) pair in the DNS Server Configuration. On     |
	//	|                      | input, dwTypeId MUST be set to DNSSRV_TYPEID_NAME_AND_PARAM, and pData MUST      |
	//	|                      | point to a structure of type DNS_RPC_NAME_AND_PARAM (section 2.2.1.2.5) that     |
	//	|                      | specifies the name of a property listed in section 3.1.1.3.1 and a new value for |
	//	|                      | that property.                                                                   |
	//	+----------------------+----------------------------------------------------------------------------------+
	ZoneScopeName string      `idl:"name:pwszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Context       uint32      `idl:"name:dwContext" json:"context"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID        uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data          *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
}

func (o *Operation3Request) xxx_ToOp(ctx context.Context) *xxx_Operation3Operation {
	if o == nil {
		return &xxx_Operation3Operation{}
	}
	return &xxx_Operation3Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		ZoneScopeName: o.ZoneScopeName,
		Context:       o.Context,
		Operation:     o.Operation,
		TypeID:        o.TypeID,
		Data:          o.Data,
	}
}

func (o *Operation3Request) xxx_FromOp(ctx context.Context, op *xxx_Operation3Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.ZoneScopeName = op.ZoneScopeName
	o.Context = op.Context
	o.Operation = op.Operation
	o.TypeID = op.TypeID
	o.Data = op.Data
}
func (o *Operation3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Operation3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Operation3Response structure represents the R_DnssrvOperation3 operation response
type Operation3Response struct {
	// Return: The R_DnssrvOperation3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Operation3Response) xxx_ToOp(ctx context.Context) *xxx_Operation3Operation {
	if o == nil {
		return &xxx_Operation3Operation{}
	}
	return &xxx_Operation3Operation{
		Return: o.Return,
	}
}

func (o *Operation3Response) xxx_FromOp(ctx context.Context, op *xxx_Operation3Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *Operation3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Operation3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Query3Operation structure represents the R_DnssrvQuery3 operation
type xxx_Query3Operation struct {
	ClientVersion uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName string      `idl:"name:pszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Operation     string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID        uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data          *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	Return        int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Query3Operation) OpNum() int { return 13 }

func (o *xxx_Query3Operation) OpName() string { return "/DnsServer/v5/R_DnssrvQuery3" }

func (o *xxx_Query3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScopeName != "" {
			_ptr_pszZoneScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScopeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScopeName, _ptr_pszZoneScopeName); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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

func (o *xxx_Query3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszZoneScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScopeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszZoneScopeName := func(ptr interface{}) { o.ZoneScopeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScopeName, _s_pszZoneScopeName, _ptr_pszZoneScopeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeId {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeId {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Query3Request structure represents the R_DnssrvQuery3 operation request
type Query3Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	Zone          string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName string `idl:"name:pszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Operation     string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
}

func (o *Query3Request) xxx_ToOp(ctx context.Context) *xxx_Query3Operation {
	if o == nil {
		return &xxx_Query3Operation{}
	}
	return &xxx_Query3Operation{
		ClientVersion: o.ClientVersion,
		SettingFlags:  o.SettingFlags,
		ServerName:    o.ServerName,
		Zone:          o.Zone,
		ZoneScopeName: o.ZoneScopeName,
		Operation:     o.Operation,
	}
}

func (o *Query3Request) xxx_FromOp(ctx context.Context, op *xxx_Query3Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.Zone = op.Zone
	o.ZoneScopeName = op.ZoneScopeName
	o.Operation = op.Operation
}
func (o *Query3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Query3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Query3Response structure represents the R_DnssrvQuery3 operation response
type Query3Response struct {
	TypeID uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data   *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	// Return: The R_DnssrvQuery3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Query3Response) xxx_ToOp(ctx context.Context) *xxx_Query3Operation {
	if o == nil {
		return &xxx_Query3Operation{}
	}
	return &xxx_Query3Operation{
		TypeID: o.TypeID,
		Data:   o.Data,
		Return: o.Return,
	}
}

func (o *Query3Response) xxx_FromOp(ctx context.Context, op *xxx_Query3Operation) {
	if o == nil {
		return
	}
	o.TypeID = op.TypeID
	o.Data = op.Data
	o.Return = op.Return
}
func (o *Query3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Query3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComplexOperation3Operation structure represents the R_DnssrvComplexOperation3 operation
type xxx_ComplexOperation3Operation struct {
	ClientVersion            uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags             uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName               string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	VirtualizationInstanceID string      `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation                string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeIn                   uint32      `idl:"name:dwTypeIn" json:"type_in"`
	DataIn                   *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
	TypeOut                  uint32      `idl:"name:pdwTypeOut" json:"type_out"`
	DataOut                  *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	Return                   int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_ComplexOperation3Operation) OpNum() int { return 14 }

func (o *xxx_ComplexOperation3Operation) OpName() string {
	return "/DnsServer/v5/R_DnssrvComplexOperation3"
}

func (o *xxx_ComplexOperation3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VirtualizationInstanceID != "" {
			_ptr_pwszVirtualizationInstanceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VirtualizationInstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataIn := uint32(o.TypeIn)
		if o.DataIn != nil {
			if err := o.DataIn.MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataIn); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszVirtualizationInstanceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VirtualizationInstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVirtualizationInstanceID := func(ptr interface{}) { o.VirtualizationInstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VirtualizationInstanceID, _s_pwszVirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeIn); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataIn == nil {
			o.DataIn = &dnsp.Union{}
		}
		_swDataIn := uint32(o.TypeIn)
		if err := o.DataIn.UnmarshalUnionNDR(ctx, w, _swDataIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeOut {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swDataOut := uint32(o.TypeOut)
		if o.DataOut != nil {
			if err := o.DataOut.MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swDataOut); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexOperation3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeOut {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeOut); err != nil {
			return err
		}
	}
	// ppDataOut {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.DataOut == nil {
			o.DataOut = &dnsp.Union{}
		}
		_swDataOut := uint32(o.TypeOut)
		if err := o.DataOut.UnmarshalUnionNDR(ctx, w, _swDataOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ComplexOperation3Request structure represents the R_DnssrvComplexOperation3 operation request
type ComplexOperation3Request struct {
	// dwClientVersion: The client version in DNS_RPC_CURRENT_CLIENT_VER (section 2.2.1.2.1)
	// format.
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	// dwSettingFlags: Reserved for future use only. This field MUST be set to zero by clients
	// and ignored by servers.
	SettingFlags uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName   string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pwszVirtualizationInstanceID: A pointer to a null-terminated Unicode string that
	// contains the name of the virtualization instance configured in the DNS server. For
	// operations specific to a virtualization instance, this field MUST contain the name
	// of the virtualization instance. If the value is NULL, then the API is as specified
	// in R_DnssrvComplexOperation2 (section 3.1.4.8). Apart from the EnumVirtualizationInstances
	// operation (section 3.1.4.3), R_DnssrvComplexOperation3 changes the behavior of the
	// following operations: EnumZoneScopes, EnumZones2, and EnumZones (section 3.1.4.3),
	// if these operation are called with R_DnssrvComplexOperation3 and a non-NULL pwszVirtualizationInstanceID,
	// they are performed under the given virtualization instance.
	VirtualizationInstanceID string      `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	Operation                string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeIn                   uint32      `idl:"name:dwTypeIn" json:"type_in"`
	DataIn                   *dnsp.Union `idl:"name:pDataIn;switch_is:dwTypeIn" json:"data_in"`
}

func (o *ComplexOperation3Request) xxx_ToOp(ctx context.Context) *xxx_ComplexOperation3Operation {
	if o == nil {
		return &xxx_ComplexOperation3Operation{}
	}
	return &xxx_ComplexOperation3Operation{
		ClientVersion:            o.ClientVersion,
		SettingFlags:             o.SettingFlags,
		ServerName:               o.ServerName,
		VirtualizationInstanceID: o.VirtualizationInstanceID,
		Zone:                     o.Zone,
		Operation:                o.Operation,
		TypeIn:                   o.TypeIn,
		DataIn:                   o.DataIn,
	}
}

func (o *ComplexOperation3Request) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperation3Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.VirtualizationInstanceID = op.VirtualizationInstanceID
	o.Zone = op.Zone
	o.Operation = op.Operation
	o.TypeIn = op.TypeIn
	o.DataIn = op.DataIn
}
func (o *ComplexOperation3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ComplexOperation3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperation3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComplexOperation3Response structure represents the R_DnssrvComplexOperation3 operation response
type ComplexOperation3Response struct {
	TypeOut uint32      `idl:"name:pdwTypeOut" json:"type_out"`
	DataOut *dnsp.Union `idl:"name:ppDataOut;switch_is:*pdwTypeOut" json:"data_out"`
	// Return: The R_DnssrvComplexOperation3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ComplexOperation3Response) xxx_ToOp(ctx context.Context) *xxx_ComplexOperation3Operation {
	if o == nil {
		return &xxx_ComplexOperation3Operation{}
	}
	return &xxx_ComplexOperation3Operation{
		TypeOut: o.TypeOut,
		DataOut: o.DataOut,
		Return:  o.Return,
	}
}

func (o *ComplexOperation3Response) xxx_FromOp(ctx context.Context, op *xxx_ComplexOperation3Operation) {
	if o == nil {
		return
	}
	o.TypeOut = op.TypeOut
	o.DataOut = op.DataOut
	o.Return = op.Return
}
func (o *ComplexOperation3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ComplexOperation3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexOperation3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Operation4Operation structure represents the R_DnssrvOperation4 operation
type xxx_Operation4Operation struct {
	ClientVersion            uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags             uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName               string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	VirtualizationInstanceID string      `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName            string      `idl:"name:pwszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Context                  uint32      `idl:"name:dwContext" json:"context"`
	Operation                string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID                   uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data                     *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
	Return                   int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Operation4Operation) OpNum() int { return 15 }

func (o *xxx_Operation4Operation) OpName() string { return "/DnsServer/v5/R_DnssrvOperation4" }

func (o *xxx_Operation4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VirtualizationInstanceID != "" {
			_ptr_pwszVirtualizationInstanceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VirtualizationInstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScopeName != "" {
			_ptr_pwszZoneScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScopeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScopeName, _ptr_pwszZoneScopeName); err != nil {
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
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszVirtualizationInstanceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VirtualizationInstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVirtualizationInstanceID := func(ptr interface{}) { o.VirtualizationInstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VirtualizationInstanceID, _s_pwszVirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScopeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScopeName := func(ptr interface{}) { o.ZoneScopeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScopeName, _s_pwszZoneScopeName, _ptr_pwszZoneScopeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwContext {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTypeId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// pData {in} (1:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Operation4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Operation4Request structure represents the R_DnssrvOperation4 operation request
type Operation4Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pwszVirtualizationInstanceID: A pointer to a null-terminated Unicode string that
	// contains the name of the virtualization instance configured on the DNS server in
	// which the operation is to be performed. For operations specific to a particular virtualization
	// instance, this field MUST contain the name of the virtualization instance. If the
	// value is NULL then the API gives the same behavior as R_DnssrvOperation3. If the
	// value is not NULL then pszZone MUST point to a null-terminated character string that
	// contains the name of the zone in UTF-8 format. In this case the type of the zone
	// pointed to by pszZone MUST be a primary zone. It MUST also be a non-autocreated zone,
	// and a non-AD integrated zone. Apart from the CreateVirtualizationInstance, DeleteVirtualizationInstance,
	// and UpdateVirtualizationInstance operations (section 3.1.4.1), R_DnssrvComplexOperation4
	// changes the behavior of the following operations: WriteDirtyZones, ZoneCreate, DeleteNode,
	// DeleteRecordSet, WriteBackFile, PauseZone, ResumeZone, DeleteZone, ReloadZone, RefreshZone,
	// CreateZoneScope, and DeleteZoneScope (section 3.1.4.1). If these operations are called
	// with R_DnssrvOperation4 and a non-NULL pwszVirtualizationInstanceID, they are performed
	// under the given virtualization instance.
	VirtualizationInstanceID string      `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName            string      `idl:"name:pwszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Context                  uint32      `idl:"name:dwContext" json:"context"`
	Operation                string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID                   uint32      `idl:"name:dwTypeId" json:"type_id"`
	Data                     *dnsp.Union `idl:"name:pData;switch_is:dwTypeId" json:"data"`
}

func (o *Operation4Request) xxx_ToOp(ctx context.Context) *xxx_Operation4Operation {
	if o == nil {
		return &xxx_Operation4Operation{}
	}
	return &xxx_Operation4Operation{
		ClientVersion:            o.ClientVersion,
		SettingFlags:             o.SettingFlags,
		ServerName:               o.ServerName,
		VirtualizationInstanceID: o.VirtualizationInstanceID,
		Zone:                     o.Zone,
		ZoneScopeName:            o.ZoneScopeName,
		Context:                  o.Context,
		Operation:                o.Operation,
		TypeID:                   o.TypeID,
		Data:                     o.Data,
	}
}

func (o *Operation4Request) xxx_FromOp(ctx context.Context, op *xxx_Operation4Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.VirtualizationInstanceID = op.VirtualizationInstanceID
	o.Zone = op.Zone
	o.ZoneScopeName = op.ZoneScopeName
	o.Context = op.Context
	o.Operation = op.Operation
	o.TypeID = op.TypeID
	o.Data = op.Data
}
func (o *Operation4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Operation4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Operation4Response structure represents the R_DnssrvOperation4 operation response
type Operation4Response struct {
	// Return: The R_DnssrvOperation4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Operation4Response) xxx_ToOp(ctx context.Context) *xxx_Operation4Operation {
	if o == nil {
		return &xxx_Operation4Operation{}
	}
	return &xxx_Operation4Operation{
		Return: o.Return,
	}
}

func (o *Operation4Response) xxx_FromOp(ctx context.Context, op *xxx_Operation4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *Operation4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Operation4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Operation4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Query4Operation structure represents the R_DnssrvQuery4 operation
type xxx_Query4Operation struct {
	ClientVersion            uint32      `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags             uint32      `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName               string      `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	VirtualizationInstanceID string      `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string      `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName            string      `idl:"name:pszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Operation                string      `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
	TypeID                   uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data                     *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	Return                   int32       `idl:"name:Return" json:"return"`
}

func (o *xxx_Query4Operation) OpNum() int { return 16 }

func (o *xxx_Query4Operation) OpName() string { return "/DnsServer/v5/R_DnssrvQuery4" }

func (o *xxx_Query4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VirtualizationInstanceID != "" {
			_ptr_pwszVirtualizationInstanceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VirtualizationInstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScopeName != "" {
			_ptr_pszZoneScopeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScopeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScopeName, _ptr_pszZoneScopeName); err != nil {
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
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Operation != "" {
			_ptr_pszOperation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Operation); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Operation, _ptr_pszOperation); err != nil {
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

func (o *xxx_Query4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszVirtualizationInstanceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VirtualizationInstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVirtualizationInstanceID := func(ptr interface{}) { o.VirtualizationInstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VirtualizationInstanceID, _s_pwszVirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZoneScopeName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszZoneScopeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScopeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszZoneScopeName := func(ptr interface{}) { o.ZoneScopeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScopeName, _s_pszZoneScopeName, _ptr_pszZoneScopeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszOperation {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszOperation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Operation); err != nil {
				return err
			}
			return nil
		})
		_s_pszOperation := func(ptr interface{}) { o.Operation = *ptr.(*string) }
		if err := w.ReadPointer(&o.Operation, _s_pszOperation, _ptr_pszOperation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwTypeId {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		_swData := uint32(o.TypeID)
		if o.Data != nil {
			if err := o.Data.MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		} else {
			if err := (&dnsp.Union{}).MarshalUnionNDR(ctx, w, _swData); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Query4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwTypeId {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeID); err != nil {
			return err
		}
	}
	// ppData {out} (1:{pointer=ref}*(1))(2:{switch_type={alias=DWORD}(uint32), alias=DNSSRV_RPC_UNION}(union))
	{
		if o.Data == nil {
			o.Data = &dnsp.Union{}
		}
		_swData := uint32(o.TypeID)
		if err := o.Data.UnmarshalUnionNDR(ctx, w, _swData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Query4Request structure represents the R_DnssrvQuery4 operation request
type Query4Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pwszVirtualizationInstanceID: A pointer to a null-terminated Unicode string that
	// contains the name of the virtualization instance configured on the DNS server. For
	// operations specific to a particular virtualization instance, this field MUST contain
	// the name of the virtualization instance. If the value is NULL, then the API is as
	// specified in R_DnssrvQuery3 (section 3.1.4.14). Apart from the VirtualizationInstances
	// operation (section 3.1.4.2), R_DnssrvQuery3 (section 3.1.4.14) changes the behavior
	// of the following operations: Zone, ZoneInfo (section 3.1.4.2) and ScopeInfo (section
	// 3.1.4.14) operations. If these operations are called with R_DnssrvQuery4 and a non-NULL
	// pwszVirtualizationInstanceID parameter, they are performed under the given virtualization
	// instance. The ScopeInfo operation is defined for R_DnssrvQuery4 with a non-NULL virtualization
	// instance only if pszZone is not NULL.
	VirtualizationInstanceID string `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScopeName            string `idl:"name:pszZoneScopeName;string;pointer:unique" json:"zone_scope_name"`
	Operation                string `idl:"name:pszOperation;string;pointer:unique" json:"operation"`
}

func (o *Query4Request) xxx_ToOp(ctx context.Context) *xxx_Query4Operation {
	if o == nil {
		return &xxx_Query4Operation{}
	}
	return &xxx_Query4Operation{
		ClientVersion:            o.ClientVersion,
		SettingFlags:             o.SettingFlags,
		ServerName:               o.ServerName,
		VirtualizationInstanceID: o.VirtualizationInstanceID,
		Zone:                     o.Zone,
		ZoneScopeName:            o.ZoneScopeName,
		Operation:                o.Operation,
	}
}

func (o *Query4Request) xxx_FromOp(ctx context.Context, op *xxx_Query4Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.VirtualizationInstanceID = op.VirtualizationInstanceID
	o.Zone = op.Zone
	o.ZoneScopeName = op.ZoneScopeName
	o.Operation = op.Operation
}
func (o *Query4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *Query4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Query4Response structure represents the R_DnssrvQuery4 operation response
type Query4Response struct {
	TypeID uint32      `idl:"name:pdwTypeId" json:"type_id"`
	Data   *dnsp.Union `idl:"name:ppData;switch_is:*pdwTypeId" json:"data"`
	// Return: The R_DnssrvQuery4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *Query4Response) xxx_ToOp(ctx context.Context) *xxx_Query4Operation {
	if o == nil {
		return &xxx_Query4Operation{}
	}
	return &xxx_Query4Operation{
		TypeID: o.TypeID,
		Data:   o.Data,
		Return: o.Return,
	}
}

func (o *Query4Response) xxx_FromOp(ctx context.Context, op *xxx_Query4Operation) {
	if o == nil {
		return
	}
	o.TypeID = op.TypeID
	o.Data = op.Data
	o.Return = op.Return
}
func (o *Query4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *Query4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Query4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateRecord4Operation structure represents the R_DnssrvUpdateRecord4 operation
type xxx_UpdateRecord4Operation struct {
	ClientVersion            uint32       `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags             uint32       `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName               string       `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	VirtualizationInstanceID string       `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope                string       `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName                 string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord                *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord             *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
	Return                   int32        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateRecord4Operation) OpNum() int { return 17 }

func (o *xxx_UpdateRecord4Operation) OpName() string { return "/DnsServer/v5/R_DnssrvUpdateRecord4" }

func (o *xxx_UpdateRecord4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VirtualizationInstanceID != "" {
			_ptr_pwszVirtualizationInstanceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VirtualizationInstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScope != "" {
			_ptr_pwszZoneScope := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScope); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScope, _ptr_pwszZoneScope); err != nil {
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
	// pszNodeName {in} (1:{string, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.AddRecord != nil {
			_ptr_pAddRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AddRecord != nil {
					if err := o.AddRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AddRecord, _ptr_pAddRecord); err != nil {
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
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		if o.DeleteRecord != nil {
			_ptr_pDeleteRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DeleteRecord != nil {
					if err := o.DeleteRecord.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dnsp.Record{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeleteRecord, _ptr_pDeleteRecord); err != nil {
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

func (o *xxx_UpdateRecord4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszVirtualizationInstanceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VirtualizationInstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVirtualizationInstanceID := func(ptr interface{}) { o.VirtualizationInstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VirtualizationInstanceID, _s_pwszVirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScope := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScope); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScope := func(ptr interface{}) { o.ZoneScope = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScope, _s_pwszZoneScope, _ptr_pwszZoneScope); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, alias=LPCSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	// pAddRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pAddRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AddRecord == nil {
				o.AddRecord = &dnsp.Record{}
			}
			if err := o.AddRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pAddRecord := func(ptr interface{}) { o.AddRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.AddRecord, _s_pAddRecord, _ptr_pAddRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDeleteRecord {in} (1:{pointer=unique, alias=PDNS_RPC_RECORD}*(1))(2:{alias=DNS_RPC_RECORD}(struct))
	{
		_ptr_pDeleteRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DeleteRecord == nil {
				o.DeleteRecord = &dnsp.Record{}
			}
			if err := o.DeleteRecord.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pDeleteRecord := func(ptr interface{}) { o.DeleteRecord = *ptr.(**dnsp.Record) }
		if err := w.ReadPointer(&o.DeleteRecord, _s_pDeleteRecord, _ptr_pDeleteRecord); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateRecord4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateRecord4Request structure represents the R_DnssrvUpdateRecord4 operation request
type UpdateRecord4Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pwszVirtualizationInstanceID: A pointer to a null-terminated Unicode string that
	// contains the name of the virtualization instance configured in the DNS server. For
	// operations specific to a particular zone or zone scope, details must be given in
	// pszZone and pwszZonescope as specified in section 3.1.4.11. If the value pwszVirtualizationInstanceID
	// is NULL, the API gives the same behavior as R_DnssrvUpdateRecord3.
	VirtualizationInstanceID string       `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string       `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope                string       `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName                 string       `idl:"name:pszNodeName;string" json:"node_name"`
	AddRecord                *dnsp.Record `idl:"name:pAddRecord;pointer:unique" json:"add_record"`
	DeleteRecord             *dnsp.Record `idl:"name:pDeleteRecord;pointer:unique" json:"delete_record"`
}

func (o *UpdateRecord4Request) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord4Operation {
	if o == nil {
		return &xxx_UpdateRecord4Operation{}
	}
	return &xxx_UpdateRecord4Operation{
		ClientVersion:            o.ClientVersion,
		SettingFlags:             o.SettingFlags,
		ServerName:               o.ServerName,
		VirtualizationInstanceID: o.VirtualizationInstanceID,
		Zone:                     o.Zone,
		ZoneScope:                o.ZoneScope,
		NodeName:                 o.NodeName,
		AddRecord:                o.AddRecord,
		DeleteRecord:             o.DeleteRecord,
	}
}

func (o *UpdateRecord4Request) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord4Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.VirtualizationInstanceID = op.VirtualizationInstanceID
	o.Zone = op.Zone
	o.ZoneScope = op.ZoneScope
	o.NodeName = op.NodeName
	o.AddRecord = op.AddRecord
	o.DeleteRecord = op.DeleteRecord
}
func (o *UpdateRecord4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UpdateRecord4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateRecord4Response structure represents the R_DnssrvUpdateRecord4 operation response
type UpdateRecord4Response struct {
	// Return: The R_DnssrvUpdateRecord4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UpdateRecord4Response) xxx_ToOp(ctx context.Context) *xxx_UpdateRecord4Operation {
	if o == nil {
		return &xxx_UpdateRecord4Operation{}
	}
	return &xxx_UpdateRecord4Operation{
		Return: o.Return,
	}
}

func (o *UpdateRecord4Response) xxx_FromOp(ctx context.Context, op *xxx_UpdateRecord4Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateRecord4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UpdateRecord4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateRecord4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRecords4Operation structure represents the R_DnssrvEnumRecords4 operation
type xxx_EnumRecords4Operation struct {
	ClientVersion            uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags             uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName               string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	VirtualizationInstanceID string `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope                string `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName                 string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild               string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType               uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag               uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart              string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop               string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
	BufferLength             uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer                   []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	Return                   int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRecords4Operation) OpNum() int { return 18 }

func (o *xxx_EnumRecords4Operation) OpName() string { return "/DnsServer/v5/R_DnssrvEnumRecords4" }

func (o *xxx_EnumRecords4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_pwszServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_pwszServerName); err != nil {
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
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VirtualizationInstanceID != "" {
			_ptr_pwszVirtualizationInstanceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VirtualizationInstanceID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
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
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.Zone != "" {
			_ptr_pszZone := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.Zone); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Zone, _ptr_pszZone); err != nil {
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
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ZoneScope != "" {
			_ptr_pwszZoneScope := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ZoneScope); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ZoneScope, _ptr_pwszZoneScope); err != nil {
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
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.NodeName != "" {
			_ptr_pszNodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_pszNodeName); err != nil {
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
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.StartChild != "" {
			_ptr_pszStartChild := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.StartChild); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.StartChild, _ptr_pszStartChild); err != nil {
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
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStart != "" {
			_ptr_pszFilterStart := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStart); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStart, _ptr_pszFilterStart); err != nil {
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
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		if o.FilterStop != "" {
			_ptr_pszFilterStop := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.FilterStop); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterStop, _ptr_pszFilterStop); err != nil {
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

func (o *xxx_EnumRecords4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwClientVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientVersion); err != nil {
			return err
		}
	}
	// dwSettingFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SettingFlags); err != nil {
			return err
		}
	}
	// pwszServerName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_pwszServerName, _ptr_pwszServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszVirtualizationInstanceID {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszVirtualizationInstanceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VirtualizationInstanceID); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVirtualizationInstanceID := func(ptr interface{}) { o.VirtualizationInstanceID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VirtualizationInstanceID, _s_pwszVirtualizationInstanceID, _ptr_pwszVirtualizationInstanceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszZone {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszZone := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.Zone); err != nil {
				return err
			}
			return nil
		})
		_s_pszZone := func(ptr interface{}) { o.Zone = *ptr.(*string) }
		if err := w.ReadPointer(&o.Zone, _s_pszZone, _ptr_pszZone); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwszZoneScope {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszZoneScope := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ZoneScope); err != nil {
				return err
			}
			return nil
		})
		_s_pwszZoneScope := func(ptr interface{}) { o.ZoneScope = *ptr.(*string) }
		if err := w.ReadPointer(&o.ZoneScope, _s_pwszZoneScope, _ptr_pwszZoneScope); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszNodeName {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszNodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_pszNodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_pszNodeName, _ptr_pszNodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszStartChild {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszStartChild := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.StartChild); err != nil {
				return err
			}
			return nil
		})
		_s_pszStartChild := func(ptr interface{}) { o.StartChild = *ptr.(*string) }
		if err := w.ReadPointer(&o.StartChild, _s_pszStartChild, _ptr_pszStartChild); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wRecordType {in} (1:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.RecordType); err != nil {
			return err
		}
	}
	// fSelectFlag {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SelectFlag); err != nil {
			return err
		}
	}
	// pszFilterStart {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStart := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStart); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStart := func(ptr interface{}) { o.FilterStart = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStart, _s_pszFilterStart, _ptr_pszFilterStart); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pszFilterStop {in} (1:{string, pointer=unique, alias=LPCSTR}*(1)[dim:0,string,null](char))
	{
		_ptr_pszFilterStop := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.FilterStop); err != nil {
				return err
			}
			return nil
		})
		_s_pszFilterStop := func(ptr interface{}) { o.FilterStop = *ptr.(*string) }
		if err := w.ReadPointer(&o.FilterStop, _s_pszFilterStop, _ptr_pszFilterStop); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwBufferLength {out} (1:{alias=PDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		if o.Buffer != nil || o.BufferLength > 0 {
			_ptr_ppBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
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
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_ppBuffer); err != nil {
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
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRecords4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwBufferLength {out} (1:{alias=PDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferLength); err != nil {
			return err
		}
	}
	// ppBuffer {out} (1:{pointer=ref}*(2))(2:{alias=PBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=pdwBufferLength](uchar))
	{
		_ptr_ppBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
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
			return nil
		})
		_s_ppBuffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_ppBuffer, _ptr_ppBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRecords4Request structure represents the R_DnssrvEnumRecords4 operation request
type EnumRecords4Request struct {
	ClientVersion uint32 `idl:"name:dwClientVersion" json:"client_version"`
	SettingFlags  uint32 `idl:"name:dwSettingFlags" json:"setting_flags"`
	ServerName    string `idl:"name:pwszServerName;string;pointer:unique" json:"server_name"`
	// pwszVirtualizationInstanceID: A pointer to a null-terminated character string that
	// contains the name of the virtualization instance under which zone and zone scope
	// records are to be enumerated. For operations specific to a particular zone or zone
	// scope, details must be given in pszZone and pwszZoneScope as specified in section
	// 3.1.4.12. If the value pwszVirtualizationInstanceID is NULL, the API gives the same
	// behavior as R_DnssrvEnumRecord3.
	VirtualizationInstanceID string `idl:"name:pwszVirtualizationInstanceID;string;pointer:unique" json:"virtualization_instance_id"`
	Zone                     string `idl:"name:pszZone;string;pointer:unique" json:"zone"`
	ZoneScope                string `idl:"name:pwszZoneScope;string;pointer:unique" json:"zone_scope"`
	NodeName                 string `idl:"name:pszNodeName;string;pointer:unique" json:"node_name"`
	StartChild               string `idl:"name:pszStartChild;string;pointer:unique" json:"start_child"`
	RecordType               uint16 `idl:"name:wRecordType" json:"record_type"`
	SelectFlag               uint32 `idl:"name:fSelectFlag" json:"select_flag"`
	FilterStart              string `idl:"name:pszFilterStart;string;pointer:unique" json:"filter_start"`
	FilterStop               string `idl:"name:pszFilterStop;string;pointer:unique" json:"filter_stop"`
}

func (o *EnumRecords4Request) xxx_ToOp(ctx context.Context) *xxx_EnumRecords4Operation {
	if o == nil {
		return &xxx_EnumRecords4Operation{}
	}
	return &xxx_EnumRecords4Operation{
		ClientVersion:            o.ClientVersion,
		SettingFlags:             o.SettingFlags,
		ServerName:               o.ServerName,
		VirtualizationInstanceID: o.VirtualizationInstanceID,
		Zone:                     o.Zone,
		ZoneScope:                o.ZoneScope,
		NodeName:                 o.NodeName,
		StartChild:               o.StartChild,
		RecordType:               o.RecordType,
		SelectFlag:               o.SelectFlag,
		FilterStart:              o.FilterStart,
		FilterStop:               o.FilterStop,
	}
}

func (o *EnumRecords4Request) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords4Operation) {
	if o == nil {
		return
	}
	o.ClientVersion = op.ClientVersion
	o.SettingFlags = op.SettingFlags
	o.ServerName = op.ServerName
	o.VirtualizationInstanceID = op.VirtualizationInstanceID
	o.Zone = op.Zone
	o.ZoneScope = op.ZoneScope
	o.NodeName = op.NodeName
	o.StartChild = op.StartChild
	o.RecordType = op.RecordType
	o.SelectFlag = op.SelectFlag
	o.FilterStart = op.FilterStart
	o.FilterStop = op.FilterStop
}
func (o *EnumRecords4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumRecords4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRecords4Response structure represents the R_DnssrvEnumRecords4 operation response
type EnumRecords4Response struct {
	BufferLength uint32 `idl:"name:pdwBufferLength" json:"buffer_length"`
	Buffer       []byte `idl:"name:ppBuffer;size_is:(, pdwBufferLength)" json:"buffer"`
	// Return: The R_DnssrvEnumRecords4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRecords4Response) xxx_ToOp(ctx context.Context) *xxx_EnumRecords4Operation {
	if o == nil {
		return &xxx_EnumRecords4Operation{}
	}
	return &xxx_EnumRecords4Operation{
		BufferLength: o.BufferLength,
		Buffer:       o.Buffer,
		Return:       o.Return,
	}
}

func (o *EnumRecords4Response) xxx_FromOp(ctx context.Context, op *xxx_EnumRecords4Operation) {
	if o == nil {
		return
	}
	o.BufferLength = op.BufferLength
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *EnumRecords4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumRecords4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRecords4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
