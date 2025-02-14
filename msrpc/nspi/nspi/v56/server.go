package nspi

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

// nspi server interface.
type NspiServer interface {

	// The NspiBind method initiates a session between a client and the NSPI server.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Bind(context.Context, *BindRequest) (*BindResponse, error)

	// The NspiUnbind method destroys the context handle. No other action is taken.
	//
	// Return Values: The server returns a DWORD value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Unbind(context.Context, *UnbindRequest) (*UnbindResponse, error)

	// The NspiUpdateStat method updates the STAT block representing position in a table
	// to reflect positioning changes requested by the client.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	UpdateStat(context.Context, *UpdateStatRequest) (*UpdateStatResponse, error)

	// The NspiQueryRows method returns to the client a number of rows from a specified
	// table. The server MUST return no more rows than the number specified in the input
	// parameter Count. Although the protocol places no further boundary or requirements
	// on the minimum number of rows the server returns, implementations SHOULD return as
	// many rows as possible subject to this maximum limit to improve usability of the NSPI
	// server for clients.
	//
	// Return Values:  The server returns a long value specifying the return status of
	// the method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryRows(context.Context, *QueryRowsRequest) (*QueryRowsResponse, error)

	// The NspiSeekEntries method searches for and sets the logical position in a specific
	// table to the first entry greater than or equal to a specified value. Optionally,
	// it might also return information about rows in the table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	SeekEntries(context.Context, *SeekEntriesRequest) (*SeekEntriesResponse, error)

	// The NspiGetMatches method returns an Explicit Table. The rows in the table are chosen
	// based on a two possible criteria: a restriction applied to an address book container
	// or the values of a property on a single object that hold references to other objects.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error)

	// The NspiResortRestriction method applies a sort order to the objects in a restricted
	// address book container.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResortRestriction(context.Context, *ResortRestrictionRequest) (*ResortRestrictionResponse, error)

	// The NspiDNToMId method maps a set of DN to a set of MId.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	DNToMID(context.Context, *DNToMIDRequest) (*DNToMIDResponse, error)

	// The NspiGetPropList method returns a list of all the properties that have values
	// on a specified object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetPropertyList(context.Context, *GetPropertyListRequest) (*GetPropertyListResponse, error)

	// The NspiGetProps method returns an address book row containing a set of the properties
	// and values that exist on an object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The NspiCompareMIds method compares the position in an address book container of
	// two objects identified by MId and returns the value of the comparison.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CompareMIDs(context.Context, *CompareMIDsRequest) (*CompareMIDsResponse, error)

	// The NspiModProps method is used to modify the properties of an object in the address
	// book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyProperties(context.Context, *ModifyPropertiesRequest) (*ModifyPropertiesResponse, error)

	// The NspiGetSpecialTable method returns the rows of a special table to the client.
	// The special table can be an Address Book Hierarchy Table or an Address Creation Table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetSpecialTable(context.Context, *GetSpecialTableRequest) (*GetSpecialTableResponse, error)

	// The NspiGetTemplateInfo method returns information about template objects in the
	// address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetTemplateInfo(context.Context, *GetTemplateInfoRequest) (*GetTemplateInfoResponse, error)

	// The NspiModLinkAtt method modifies the values of a specific property of a specific
	// row in the address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyLinkAttribute(context.Context, *ModifyLinkAttributeRequest) (*ModifyLinkAttributeResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The NspiQueryColumns method returns a list of all the properties the NSPI server
	// is aware of. It returns this list as an array of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryColumns(context.Context, *QueryColumnsRequest) (*QueryColumnsResponse, error)

	// The NspiGetNamesFromIDs method returns a list of property names for a set of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetNamesFromIDs(context.Context, *GetNamesFromIDsRequest) (*GetNamesFromIDsResponse, error)

	// The NspiGetIDsFromNames method returns a list of proptags for a set of property names.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetIDsFromNames(context.Context, *GetIDsFromNamesRequest) (*GetIDsFromNamesResponse, error)

	// The NspiResolveNames method takes a set of string values in an 8-bit character set
	// and performs ANR (as specified in 3.1.1.6) on those strings. The server reports the
	// MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNames(context.Context, *ResolveNamesRequest) (*ResolveNamesResponse, error)

	// The NspiResolveNamesW method takes a set of string values in the Unicode character
	// set and performs ANR (as specified in 3.1.1.6) on those strings. The server reports
	// the MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNamesW(context.Context, *ResolveNamesWRequest) (*ResolveNamesWResponse, error)
}

func RegisterNspiServer(conn dcerpc.Conn, o NspiServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNspiServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NspiSyntaxV56_0))...)
}

func NewNspiServerHandle(o NspiServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NspiServerHandle(ctx, o, opNum, r)
	}
}

func NspiServerHandle(ctx context.Context, o NspiServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NspiBind
		op := &xxx_BindOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BindRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Bind(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // NspiUnbind
		op := &xxx_UnbindOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnbindRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Unbind(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NspiUpdateStat
		op := &xxx_UpdateStatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateStatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateStat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // NspiQueryRows
		op := &xxx_QueryRowsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryRowsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryRows(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // NspiSeekEntries
		op := &xxx_SeekEntriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SeekEntriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SeekEntries(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // NspiGetMatches
		op := &xxx_GetMatchesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMatchesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMatches(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // NspiResortRestriction
		op := &xxx_ResortRestrictionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResortRestrictionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResortRestriction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // NspiDNToMId
		op := &xxx_DNToMIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DNToMIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DNToMID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // NspiGetPropList
		op := &xxx_GetPropertyListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertyListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPropertyList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // NspiGetProps
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // NspiCompareMIds
		op := &xxx_CompareMIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CompareMIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CompareMIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // NspiModProps
		op := &xxx_ModifyPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // NspiGetSpecialTable
		op := &xxx_GetSpecialTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSpecialTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSpecialTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // NspiGetTemplateInfo
		op := &xxx_GetTemplateInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTemplateInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTemplateInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // NspiModLinkAtt
		op := &xxx_ModifyLinkAttributeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyLinkAttributeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyLinkAttribute(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // NspiQueryColumns
		op := &xxx_QueryColumnsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryColumnsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryColumns(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NspiGetNamesFromIDs
		op := &xxx_GetNamesFromIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNamesFromIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNamesFromIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // NspiGetIDsFromNames
		op := &xxx_GetIDsFromNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIDsFromNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIDsFromNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // NspiResolveNames
		op := &xxx_ResolveNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResolveNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResolveNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // NspiResolveNamesW
		op := &xxx_ResolveNamesWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResolveNamesWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResolveNamesW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented nspi
type UnimplementedNspiServer struct {
}

func (UnimplementedNspiServer) Bind(context.Context, *BindRequest) (*BindResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) Unbind(context.Context, *UnbindRequest) (*UnbindResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) UpdateStat(context.Context, *UpdateStatRequest) (*UpdateStatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) QueryRows(context.Context, *QueryRowsRequest) (*QueryRowsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) SeekEntries(context.Context, *SeekEntriesRequest) (*SeekEntriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) ResortRestriction(context.Context, *ResortRestrictionRequest) (*ResortRestrictionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) DNToMID(context.Context, *DNToMIDRequest) (*DNToMIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetPropertyList(context.Context, *GetPropertyListRequest) (*GetPropertyListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) CompareMIDs(context.Context, *CompareMIDsRequest) (*CompareMIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) ModifyProperties(context.Context, *ModifyPropertiesRequest) (*ModifyPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetSpecialTable(context.Context, *GetSpecialTableRequest) (*GetSpecialTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetTemplateInfo(context.Context, *GetTemplateInfoRequest) (*GetTemplateInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) ModifyLinkAttribute(context.Context, *ModifyLinkAttributeRequest) (*ModifyLinkAttributeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) QueryColumns(context.Context, *QueryColumnsRequest) (*QueryColumnsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetNamesFromIDs(context.Context, *GetNamesFromIDsRequest) (*GetNamesFromIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) GetIDsFromNames(context.Context, *GetIDsFromNamesRequest) (*GetIDsFromNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) ResolveNames(context.Context, *ResolveNamesRequest) (*ResolveNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNspiServer) ResolveNamesW(context.Context, *ResolveNamesWRequest) (*ResolveNamesWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NspiServer = (*UnimplementedNspiServer)(nil)
