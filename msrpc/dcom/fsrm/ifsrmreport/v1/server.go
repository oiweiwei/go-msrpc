package ifsrmreport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IFsrmReport server interface.
type ReportServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Type operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error)

	// The LastGeneratedFileNamePrefix (get) method retrieves the last generated file name
	// prefix of the report for the most recently generated report and returns S_OK upon
	// successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-------------------------------+
	//	|         RETURN          |                               |
	//	|       VALUE/CODE        |          DESCRIPTION          |
	//	|                         |                               |
	//	+-------------------------+-------------------------------+
	//	+-------------------------+-------------------------------+
	//	| 0x80070057 E_INVALIDARG | The prefix parameter is NULL. |
	//	+-------------------------+-------------------------------+
	GetLastGeneratedFileNamePrefix(context.Context, *GetLastGeneratedFileNamePrefixRequest) (*GetLastGeneratedFileNamePrefixResponse, error)

	// The GetFilter method returns the current value of the specified report filter for
	// the report object.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The filter parameter is not a   |
	//	|                         | valid FsrmReportFilter (section 2.2.1.2.16) value. The filterValue parameter is  |
	//	|                         | NULL.                                                                            |
	//	+-------------------------+----------------------------------------------------------------------------------+
	GetFilter(context.Context, *GetFilterRequest) (*GetFilterResponse, error)

	// The SetFilter method sets the value of the specified report filter for the report
	// object. The filter value will override the default value set by using the IFsrmReportManager::SetDefaultFilter
	// (section 3.2.4.2.33.8) method.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The filter parameter is not a   |
	//	|                         | valid FsrmReportFilter (section 2.2.1.2.16) value. An attempt was made to set    |
	//	|                         | the FsrmReportFilter_Property filter value with a supplied value that is not in  |
	//	|                         | a valid property name format, or the property does not exist. The variant does   |
	//	|                         | not have the correct member set for the filter. The string filter values are not |
	//	|                         | valid characters.                                                                |
	//	+-------------------------+----------------------------------------------------------------------------------+
	SetFilter(context.Context, *SetFilterRequest) (*SetFilterResponse, error)

	// Delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterReportServer(conn dcerpc.Conn, o ReportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewReportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ReportSyntaxV1_0))...)
}

func NewReportServerHandle(o ReportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ReportServerHandle(ctx, o, opNum, r)
	}
}

func ReportServerHandle(ctx context.Context, o ReportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Type
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Description
		op := &xxx_GetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Description
		op := &xxx_SetDescriptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDescriptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDescription(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // LastGeneratedFileNamePrefix
		op := &xxx_GetLastGeneratedFileNamePrefixOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLastGeneratedFileNamePrefixRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLastGeneratedFileNamePrefix(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetFilter
		op := &xxx_GetFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SetFilter
		op := &xxx_SetFilterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFilterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFilter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmReport
type UnimplementedReportServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedReportServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) GetDescription(context.Context, *GetDescriptionRequest) (*GetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) SetDescription(context.Context, *SetDescriptionRequest) (*SetDescriptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) GetLastGeneratedFileNamePrefix(context.Context, *GetLastGeneratedFileNamePrefixRequest) (*GetLastGeneratedFileNamePrefixResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) GetFilter(context.Context, *GetFilterRequest) (*GetFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) SetFilter(context.Context, *SetFilterRequest) (*SetFilterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedReportServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ReportServer = (*UnimplementedReportServer)(nil)
