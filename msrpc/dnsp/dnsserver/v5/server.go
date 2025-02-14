package dnsserver

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

// DnsServer server interface.
type DNSServerServer interface {

	// The R_DnssrvOperation method is used to invoke a set of server functions specified
	// by pszOperation.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. See [MS-ERREF] section 2.2 or
	// section 2.2.1.1.5. All error values MUST be treated the same.
	Operation(context.Context, *OperationRequest) (*OperationResponse, error)

	// The R_DnssrvQuery method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone and pszOperation parameters.
	// For the purpose of selecting an output structure type the server MUST consider the
	// value of dwClientVersion (section 2.2.1.2.1) to be 0x00000000 when responding to
	// this method.
	//
	// Return Values: A Win32 error code indicating whether the operation completed successfully
	// (0x00000000) or failed (any other value). See [MS-ERREF] section 2.2 or section 2.2.1.1.5.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)

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
	ComplexOperation(context.Context, *ComplexOperationRequest) (*ComplexOperationResponse, error)

	// The R_DnssrvEnumRecords method enumerates DNS records on the server.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code if an error occurred. See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same, except that if the return code
	// is ERROR_MORE_DATA (0x000000EA) then the enumeration contains more results than can
	// fit into a single RPC buffer. In this case the client application can call this method
	// again passing the last retrieved child as the pszStartChild argument to retrieve
	// the next set of results.
	EnumRecords(context.Context, *EnumRecordsRequest) (*EnumRecordsResponse, error)

	// The R_DnssrvUpdateRecord method is used to add a new DNS record or modify/delete
	// an existing DNS record at the server. This operation SHOULD<277> be supported.
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// non-zero Win32 error code if an error occurred. See [MS-ERREF] section 2.2 or section
	// 2.2.1.1.5. All error values MUST be treated the same.
	UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error)

	// The R_DnssrvOperation2 method is used to invoke a set of server functions specified
	// by the caller. The DNS server SHOULD implement R_DnssrvOperation2.
	//
	// All parameters are as specified by the R_DnssrvOperation method (section 3.1.4.1)
	// with the following exceptions:
	//
	// Return Values: The method MUST return ERROR_SUCCESS (0x00000000) on success or a
	// nonzero Win32 error code value if an error occurred. See [MS-ERREF] section 2.2 or
	// section 2.2.1.1.5. All error values MUST be treated the same.
	Operation2(context.Context, *Operation2Request) (*Operation2Response, error)

	// The R_DnssrvQuery2 method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone and pszOperation parameters.
	// The DNS server SHOULD implement R_ DnssrvQuery2 <280>.
	//
	// All parameters are as specified by the R_DnssrvQuery method (section 3.1.4.2) with
	// the following exceptions:
	//
	// Return Values: Return values behaviors and interpretations are same as they are for
	// R_DnssrvQuery method (section 3.1.4.2).
	Query2(context.Context, *Query2Request) (*Query2Response, error)

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
	ComplexOperation2(context.Context, *ComplexOperation2Request) (*ComplexOperation2Response, error)

	// The R_DnssrvEnumRecords2 method enumerates DNS records on the server. The DNS server
	// SHOULD implement R_DnssrvEnumRecords2 <281>.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords method (section 3.1.4.4)
	// with the following exceptions:
	//
	// Return Values: Return values behaviors and interpretations are same as they are for
	// R_DnssrvEnumRecords method (section 3.1.4.4).
	EnumRecords2(context.Context, *EnumRecords2Request) (*EnumRecords2Response, error)

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
	UpdateRecord2(context.Context, *UpdateRecord2Request) (*UpdateRecord2Response, error)

	// The R_DnssrvUpdateRecord3 method is used to add a new DNS record or modify or delete
	// an existing DNS record in a zone or in a zone scope, or in a cache zone or cache
	// scope, if specified. The DNS server SHOULD<283> implement R_DnssrvUpdateRecord3.
	//
	// All parameters are as specified by the methods R_DnssrvUpdateRecord (section 3.1.4.5)
	// and R_DnssrvUpdateRecord2 (section 3.1.4.10) with the following exceptions:
	UpdateRecord3(context.Context, *UpdateRecord3Request) (*UpdateRecord3Response, error)

	// The R_DnssrvEnumRecords3 method enumerates DNS records on a zone or a zone scope,
	// or cache zone or a cache scope, if specified. The DNS server SHOULD<284> implement
	// R_DnssrvEnumRecords3.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords method (section 3.1.4.4)
	// and implement the R_DnssrvEnumRecords2 method (section 3.1.4.9) with the following
	// exceptions:
	EnumRecords3(context.Context, *EnumRecords3Request) (*EnumRecords3Response, error)

	// The R_DnssrvOperation3 method is used to invoke a set of server functions specified
	// by the caller on the zone scope or cache scope if specified. The DNS server SHOULD<285>
	// implement R_DnssrvOperation3.
	//
	// All parameters are as specified by the methods R_DnssrvOperation (section 3.1.4.1)
	// and R_DnssrvOperation2 (section 3.1.4.6) with the following exceptions.
	Operation3(context.Context, *Operation3Request) (*Operation3Response, error)

	// The R_DnssrvQuery3 method queries the DNS server for information. The type of information
	// queried for is specified by the client using the pszZone, pwszZoneScopeName, and
	// pszOperation parameters. The DNS server SHOULD<287> implement R_DnssrvQuery3.
	//
	// All the parameters are as specified by the R_DnssrvQuery2 method with the following
	// exceptions:
	Query3(context.Context, *Query3Request) (*Query3Response, error)

	// The R_DnssrvComplexOperation3 method is used to invoke a set of server functions
	// specified by the caller. These functions generally return more complex structures
	// than simple 32-bit integer values, unlike the operations accessible through R_DnssrvOperation
	// (section 3.1.4.1). The DNS server SHOULD<289> implement R_DnssrvComplexOperation2
	// (section 3.1.4.8).
	//
	// All parameters are as specified by the R_DnssrvComplexOperation method (section 3.1.4.3)
	// with the following exceptions:
	ComplexOperation3(context.Context, *ComplexOperation3Request) (*ComplexOperation3Response, error)

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
	Operation4(context.Context, *Operation4Request) (*Operation4Response, error)

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
	Query4(context.Context, *Query4Request) (*Query4Response, error)

	// The R_DnssrvUpdateRecord4 method is used to add a new DNS record or to modify or
	// delete an existing DNS record in a zone or in a zone scope under a virtualization
	// instance, if specified. The DNS server SHOULD<292> implement R_DnssrvUpdateRecord4.
	//
	// All parameters are as specified by the R_DnssrvUpdateRecord3 (section 3.1.4.11) method
	// with the following exceptions:
	UpdateRecord4(context.Context, *UpdateRecord4Request) (*UpdateRecord4Response, error)

	// The R_DnssrvEnumRecords4 method enumerates DNS records on a zone or a zone scope
	// in a virtualization instance, if specified. The DNS server SHOULD<293> implement
	// R_DnssrvEnumRecords4.
	//
	// All parameters are as specified by the R_DnssrvEnumRecords3 method (section 3.1.4.12)
	// with the following exceptions:
	EnumRecords4(context.Context, *EnumRecords4Request) (*EnumRecords4Response, error)
}

func RegisterDNSServerServer(conn dcerpc.Conn, o DNSServerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDNSServerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DNSServerSyntaxV5_0))...)
}

func NewDNSServerServerHandle(o DNSServerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DNSServerServerHandle(ctx, o, opNum, r)
	}
}

func DNSServerServerHandle(ctx context.Context, o DNSServerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_DnssrvOperation
		op := &xxx_OperationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OperationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Operation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // R_DnssrvQuery
		op := &xxx_QueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // R_DnssrvComplexOperation
		op := &xxx_ComplexOperationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ComplexOperationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ComplexOperation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // R_DnssrvEnumRecords
		op := &xxx_EnumRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // R_DnssrvUpdateRecord
		op := &xxx_UpdateRecordOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateRecordRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateRecord(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // R_DnssrvOperation2
		op := &xxx_Operation2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Operation2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Operation2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // R_DnssrvQuery2
		op := &xxx_Query2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Query2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // R_DnssrvComplexOperation2
		op := &xxx_ComplexOperation2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ComplexOperation2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ComplexOperation2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // R_DnssrvEnumRecords2
		op := &xxx_EnumRecords2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRecords2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRecords2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // R_DnssrvUpdateRecord2
		op := &xxx_UpdateRecord2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateRecord2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateRecord2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // R_DnssrvUpdateRecord3
		op := &xxx_UpdateRecord3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateRecord3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateRecord3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // R_DnssrvEnumRecords3
		op := &xxx_EnumRecords3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRecords3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRecords3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // R_DnssrvOperation3
		op := &xxx_Operation3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Operation3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Operation3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // R_DnssrvQuery3
		op := &xxx_Query3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Query3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // R_DnssrvComplexOperation3
		op := &xxx_ComplexOperation3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ComplexOperation3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ComplexOperation3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // R_DnssrvOperation4
		op := &xxx_Operation4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Operation4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Operation4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // R_DnssrvQuery4
		op := &xxx_Query4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Query4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Query4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // R_DnssrvUpdateRecord4
		op := &xxx_UpdateRecord4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateRecord4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateRecord4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // R_DnssrvEnumRecords4
		op := &xxx_EnumRecords4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRecords4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRecords4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented DnsServer
type UnimplementedDNSServerServer struct {
}

func (UnimplementedDNSServerServer) Operation(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) ComplexOperation(context.Context, *ComplexOperationRequest) (*ComplexOperationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) EnumRecords(context.Context, *EnumRecordsRequest) (*EnumRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Operation2(context.Context, *Operation2Request) (*Operation2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Query2(context.Context, *Query2Request) (*Query2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) ComplexOperation2(context.Context, *ComplexOperation2Request) (*ComplexOperation2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) EnumRecords2(context.Context, *EnumRecords2Request) (*EnumRecords2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) UpdateRecord2(context.Context, *UpdateRecord2Request) (*UpdateRecord2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) UpdateRecord3(context.Context, *UpdateRecord3Request) (*UpdateRecord3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) EnumRecords3(context.Context, *EnumRecords3Request) (*EnumRecords3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Operation3(context.Context, *Operation3Request) (*Operation3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Query3(context.Context, *Query3Request) (*Query3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) ComplexOperation3(context.Context, *ComplexOperation3Request) (*ComplexOperation3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Operation4(context.Context, *Operation4Request) (*Operation4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) Query4(context.Context, *Query4Request) (*Query4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) UpdateRecord4(context.Context, *UpdateRecord4Request) (*UpdateRecord4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDNSServerServer) EnumRecords4(context.Context, *EnumRecords4Request) (*EnumRecords4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DNSServerServer = (*UnimplementedDNSServerServer)(nil)
