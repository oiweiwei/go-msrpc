package iapitracingdatacollector

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idatacollector "github.com/oiweiwei/go-msrpc/msrpc/dcom/pla/idatacollector/v0"
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
	_ = idatacollector.GoPackage
)

// IApiTracingDataCollector server interface.
type APITracingDataCollectorServer interface {

	// IDataCollector base class.
	idatacollector.DataCollectorServer

	// LogApiNamesOnly operation.
	GetLogAPINamesOnly(context.Context, *GetLogAPINamesOnlyRequest) (*GetLogAPINamesOnlyResponse, error)

	// LogApiNamesOnly operation.
	SetLogAPINamesOnly(context.Context, *SetLogAPINamesOnlyRequest) (*SetLogAPINamesOnlyResponse, error)

	// LogApisRecursively operation.
	GetLogAPIsRecursively(context.Context, *GetLogAPIsRecursivelyRequest) (*GetLogAPIsRecursivelyResponse, error)

	// LogApisRecursively operation.
	SetLogAPIsRecursively(context.Context, *SetLogAPIsRecursivelyRequest) (*SetLogAPIsRecursivelyResponse, error)

	// ExePath operation.
	GetExePath(context.Context, *GetExePathRequest) (*GetExePathResponse, error)

	// ExePath operation.
	SetExePath(context.Context, *SetExePathRequest) (*SetExePathResponse, error)

	// LogFilePath operation.
	GetLogFilePath(context.Context, *GetLogFilePathRequest) (*GetLogFilePathResponse, error)

	// LogFilePath operation.
	SetLogFilePath(context.Context, *SetLogFilePathRequest) (*SetLogFilePathResponse, error)

	// IncludeModules operation.
	GetIncludeModules(context.Context, *GetIncludeModulesRequest) (*GetIncludeModulesResponse, error)

	// IncludeModules operation.
	SetIncludeModules(context.Context, *SetIncludeModulesRequest) (*SetIncludeModulesResponse, error)

	// IncludeApis operation.
	GetIncludeAPIs(context.Context, *GetIncludeAPIsRequest) (*GetIncludeAPIsResponse, error)

	// IncludeApis operation.
	SetIncludeAPIs(context.Context, *SetIncludeAPIsRequest) (*SetIncludeAPIsResponse, error)

	// ExcludeApis operation.
	GetExcludeAPIs(context.Context, *GetExcludeAPIsRequest) (*GetExcludeAPIsResponse, error)

	// ExcludeApis operation.
	SetExcludeAPIs(context.Context, *SetExcludeAPIsRequest) (*SetExcludeAPIsResponse, error)
}

func RegisterAPITracingDataCollectorServer(conn dcerpc.Conn, o APITracingDataCollectorServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAPITracingDataCollectorServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(APITracingDataCollectorSyntaxV0_0))...)
}

func NewAPITracingDataCollectorServerHandle(o APITracingDataCollectorServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return APITracingDataCollectorServerHandle(ctx, o, opNum, r)
	}
}

func APITracingDataCollectorServerHandle(ctx context.Context, o APITracingDataCollectorServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 32 {
		// IDataCollector base method.
		return idatacollector.DataCollectorServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 32: // LogApiNamesOnly
		op := &xxx_GetLogAPINamesOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogAPINamesOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogAPINamesOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // LogApiNamesOnly
		op := &xxx_SetLogAPINamesOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogAPINamesOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogAPINamesOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // LogApisRecursively
		op := &xxx_GetLogAPIsRecursivelyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogAPIsRecursivelyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogAPIsRecursively(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // LogApisRecursively
		op := &xxx_SetLogAPIsRecursivelyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogAPIsRecursivelyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogAPIsRecursively(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // ExePath
		op := &xxx_GetExePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // ExePath
		op := &xxx_SetExePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // LogFilePath
		op := &xxx_GetLogFilePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogFilePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogFilePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // LogFilePath
		op := &xxx_SetLogFilePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogFilePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogFilePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // IncludeModules
		op := &xxx_GetIncludeModulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIncludeModulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIncludeModules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // IncludeModules
		op := &xxx_SetIncludeModulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetIncludeModulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetIncludeModules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // IncludeApis
		op := &xxx_GetIncludeAPIsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIncludeAPIsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIncludeAPIs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // IncludeApis
		op := &xxx_SetIncludeAPIsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetIncludeAPIsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetIncludeAPIs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // ExcludeApis
		op := &xxx_GetExcludeAPIsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExcludeAPIsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExcludeAPIs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // ExcludeApis
		op := &xxx_SetExcludeAPIsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExcludeAPIsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExcludeAPIs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IApiTracingDataCollector
type UnimplementedAPITracingDataCollectorServer struct {
	idatacollector.UnimplementedDataCollectorServer
}

func (UnimplementedAPITracingDataCollectorServer) GetLogAPINamesOnly(context.Context, *GetLogAPINamesOnlyRequest) (*GetLogAPINamesOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetLogAPINamesOnly(context.Context, *SetLogAPINamesOnlyRequest) (*SetLogAPINamesOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetLogAPIsRecursively(context.Context, *GetLogAPIsRecursivelyRequest) (*GetLogAPIsRecursivelyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetLogAPIsRecursively(context.Context, *SetLogAPIsRecursivelyRequest) (*SetLogAPIsRecursivelyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetExePath(context.Context, *GetExePathRequest) (*GetExePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetExePath(context.Context, *SetExePathRequest) (*SetExePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetLogFilePath(context.Context, *GetLogFilePathRequest) (*GetLogFilePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetLogFilePath(context.Context, *SetLogFilePathRequest) (*SetLogFilePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetIncludeModules(context.Context, *GetIncludeModulesRequest) (*GetIncludeModulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetIncludeModules(context.Context, *SetIncludeModulesRequest) (*SetIncludeModulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetIncludeAPIs(context.Context, *GetIncludeAPIsRequest) (*GetIncludeAPIsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetIncludeAPIs(context.Context, *SetIncludeAPIsRequest) (*SetIncludeAPIsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) GetExcludeAPIs(context.Context, *GetExcludeAPIsRequest) (*GetExcludeAPIsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAPITracingDataCollectorServer) SetExcludeAPIs(context.Context, *SetExcludeAPIsRequest) (*SetExcludeAPIsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ APITracingDataCollectorServer = (*UnimplementedAPITracingDataCollectorServer)(nil)
