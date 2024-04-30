package idatacollector

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

// IDataCollector server interface.
type DataCollectorServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The DataCollectorSet (Get) method retrieves the DataCollectorSet property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataCollectorSet(context.Context, *GetDataCollectorSetRequest) (*GetDataCollectorSetResponse, error)

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// The DataCollectorType enumeration defines the data collector types.
	//
	// The DataCollectorType (Get) method retrieves the DataCollectorType property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetDataCollectorType(context.Context, *GetDataCollectorTypeRequest) (*GetDataCollectorTypeResponse, error)

	// FileName operation.
	GetFileName(context.Context, *GetFileNameRequest) (*GetFileNameResponse, error)

	// FileName operation.
	SetFileName(context.Context, *SetFileNameRequest) (*SetFileNameResponse, error)

	// FileNameFormat operation.
	GetFileNameFormat(context.Context, *GetFileNameFormatRequest) (*GetFileNameFormatResponse, error)

	// FileNameFormat operation.
	SetFileNameFormat(context.Context, *SetFileNameFormatRequest) (*SetFileNameFormatResponse, error)

	// FileNameFormatPattern operation.
	GetFileNameFormatPattern(context.Context, *GetFileNameFormatPatternRequest) (*GetFileNameFormatPatternResponse, error)

	// FileNameFormatPattern operation.
	SetFileNameFormatPattern(context.Context, *SetFileNameFormatPatternRequest) (*SetFileNameFormatPatternResponse, error)

	// LatestOutputLocation operation.
	GetLatestOutputLocation(context.Context, *GetLatestOutputLocationRequest) (*GetLatestOutputLocationResponse, error)

	// LatestOutputLocation operation.
	SetLatestOutputLocation(context.Context, *SetLatestOutputLocationRequest) (*SetLatestOutputLocationResponse, error)

	// LogAppend operation.
	GetLogAppend(context.Context, *GetLogAppendRequest) (*GetLogAppendResponse, error)

	// LogAppend operation.
	SetLogAppend(context.Context, *SetLogAppendRequest) (*SetLogAppendResponse, error)

	// LogCircular operation.
	GetLogCircular(context.Context, *GetLogCircularRequest) (*GetLogCircularResponse, error)

	// LogCircular operation.
	SetLogCircular(context.Context, *SetLogCircularRequest) (*SetLogCircularResponse, error)

	// LogOverwrite operation.
	GetLogOverwrite(context.Context, *GetLogOverwriteRequest) (*GetLogOverwriteResponse, error)

	// LogOverwrite operation.
	SetLogOverwrite(context.Context, *SetLogOverwriteRequest) (*SetLogOverwriteResponse, error)

	// Name operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest) (*SetNameResponse, error)

	// OutputLocation operation.
	GetOutputLocation(context.Context, *GetOutputLocationRequest) (*GetOutputLocationResponse, error)

	// The Index (Get) method retrieves the Index property.
	//
	// Return Values: This method MUST return an HRESULT with the severity bit clear on
	// success as specified in [MS-ERREF]; otherwise, it MUST return one of the errors as
	// defined in 2.2.1 or one of the errors as defined in [MS-ERREF] section 2.1.
	GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error)

	// Opnum28NotUsedOnWire operation.
	// Opnum28NotUsedOnWire

	// Xml operation.
	GetXML(context.Context, *GetXMLRequest) (*GetXMLResponse, error)

	// SetXml operation.
	SetXML(context.Context, *SetXMLRequest) (*SetXMLResponse, error)

	// Opnum31NotUsedOnWire operation.
	// Opnum31NotUsedOnWire
}

func RegisterDataCollectorServer(conn dcerpc.Conn, o DataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataCollectorSyntaxV0_0))...)
}

func NewDataCollectorServerHandle(o DataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func DataCollectorServerHandle(ctx context.Context, o DataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DataCollectorSet
		in := &GetDataCollectorSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDataCollectorSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // DataCollectorType
		in := &GetDataCollectorTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDataCollectorType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // FileName
		in := &GetFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // FileName
		in := &SetFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // FileNameFormat
		in := &GetFileNameFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileNameFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // FileNameFormat
		in := &SetFileNameFormatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileNameFormat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // FileNameFormatPattern
		in := &GetFileNameFormatPatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileNameFormatPattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // FileNameFormatPattern
		in := &SetFileNameFormatPatternRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileNameFormatPattern(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // LatestOutputLocation
		in := &GetLatestOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLatestOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // LatestOutputLocation
		in := &SetLatestOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLatestOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // LogAppend
		in := &GetLogAppendRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogAppend(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // LogAppend
		in := &SetLogAppendRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogAppend(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // LogCircular
		in := &GetLogCircularRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogCircular(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // LogCircular
		in := &SetLogCircularRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogCircular(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // LogOverwrite
		in := &GetLogOverwriteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogOverwrite(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // LogOverwrite
		in := &SetLogOverwriteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogOverwrite(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // Name
		in := &GetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // Name
		in := &SetNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // OutputLocation
		in := &GetOutputLocationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOutputLocation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // Index
		in := &GetIndexRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIndex(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // Opnum28NotUsedOnWire
		// Opnum28NotUsedOnWire
		return nil, nil
	case 29: // Xml
		in := &GetXMLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetXML(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // SetXml
		in := &SetXMLRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetXML(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // Opnum31NotUsedOnWire
		// Opnum31NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
