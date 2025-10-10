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
		op := &xxx_GetDataCollectorSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataCollectorSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataCollectorSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // DataCollectorType
		op := &xxx_GetDataCollectorTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDataCollectorTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDataCollectorType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // FileName
		op := &xxx_GetFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // FileName
		op := &xxx_SetFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // FileNameFormat
		op := &xxx_GetFileNameFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileNameFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileNameFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // FileNameFormat
		op := &xxx_SetFileNameFormatOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileNameFormatRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileNameFormat(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FileNameFormatPattern
		op := &xxx_GetFileNameFormatPatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileNameFormatPatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileNameFormatPattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FileNameFormatPattern
		op := &xxx_SetFileNameFormatPatternOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFileNameFormatPatternRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFileNameFormatPattern(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // LatestOutputLocation
		op := &xxx_GetLatestOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLatestOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLatestOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // LatestOutputLocation
		op := &xxx_SetLatestOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLatestOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLatestOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // LogAppend
		op := &xxx_GetLogAppendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogAppendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogAppend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // LogAppend
		op := &xxx_SetLogAppendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogAppendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogAppend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // LogCircular
		op := &xxx_GetLogCircularOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogCircularRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogCircular(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // LogCircular
		op := &xxx_SetLogCircularOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogCircularRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogCircular(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // LogOverwrite
		op := &xxx_GetLogOverwriteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogOverwriteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogOverwrite(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // LogOverwrite
		op := &xxx_SetLogOverwriteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogOverwriteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogOverwrite(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // Name
		op := &xxx_SetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // OutputLocation
		op := &xxx_GetOutputLocationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutputLocationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutputLocation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // Index
		op := &xxx_GetIndexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIndexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIndex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Opnum28NotUsedOnWire
		// Opnum28NotUsedOnWire
		return nil, nil
	case 29: // Xml
		op := &xxx_GetXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // SetXml
		op := &xxx_SetXMLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetXMLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetXML(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // Opnum31NotUsedOnWire
		// Opnum31NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IDataCollector
type UnimplementedDataCollectorServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedDataCollectorServer) GetDataCollectorSet(context.Context, *GetDataCollectorSetRequest) (*GetDataCollectorSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetDataCollectorType(context.Context, *GetDataCollectorTypeRequest) (*GetDataCollectorTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetFileName(context.Context, *GetFileNameRequest) (*GetFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetFileName(context.Context, *SetFileNameRequest) (*SetFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetFileNameFormat(context.Context, *GetFileNameFormatRequest) (*GetFileNameFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetFileNameFormat(context.Context, *SetFileNameFormatRequest) (*SetFileNameFormatResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetFileNameFormatPattern(context.Context, *GetFileNameFormatPatternRequest) (*GetFileNameFormatPatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetFileNameFormatPattern(context.Context, *SetFileNameFormatPatternRequest) (*SetFileNameFormatPatternResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetLatestOutputLocation(context.Context, *GetLatestOutputLocationRequest) (*GetLatestOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetLatestOutputLocation(context.Context, *SetLatestOutputLocationRequest) (*SetLatestOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetLogAppend(context.Context, *GetLogAppendRequest) (*GetLogAppendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetLogAppend(context.Context, *SetLogAppendRequest) (*SetLogAppendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetLogCircular(context.Context, *GetLogCircularRequest) (*GetLogCircularResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetLogCircular(context.Context, *SetLogCircularRequest) (*SetLogCircularResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetLogOverwrite(context.Context, *GetLogOverwriteRequest) (*GetLogOverwriteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetLogOverwrite(context.Context, *SetLogOverwriteRequest) (*SetLogOverwriteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetName(context.Context, *SetNameRequest) (*SetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetOutputLocation(context.Context, *GetOutputLocationRequest) (*GetOutputLocationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetIndex(context.Context, *GetIndexRequest) (*GetIndexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) GetXML(context.Context, *GetXMLRequest) (*GetXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDataCollectorServer) SetXML(context.Context, *SetXMLRequest) (*SetXMLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DataCollectorServer = (*UnimplementedDataCollectorServer)(nil)
