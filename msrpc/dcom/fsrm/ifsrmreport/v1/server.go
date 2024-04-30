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
		in := &GetTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Description
		in := &GetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // Description
		in := &SetDescriptionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDescription(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // LastGeneratedFileNamePrefix
		in := &GetLastGeneratedFileNamePrefixRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLastGeneratedFileNamePrefix(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // GetFilter
		in := &GetFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // SetFilter
		in := &SetFilterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
