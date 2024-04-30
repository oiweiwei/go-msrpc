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
		in := &GetLogAPINamesOnlyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogAPINamesOnly(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // LogApiNamesOnly
		in := &SetLogAPINamesOnlyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogAPINamesOnly(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // LogApisRecursively
		in := &GetLogAPIsRecursivelyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogAPIsRecursively(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // LogApisRecursively
		in := &SetLogAPIsRecursivelyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogAPIsRecursively(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // ExePath
		in := &GetExePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // ExePath
		in := &SetExePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // LogFilePath
		in := &GetLogFilePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogFilePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // LogFilePath
		in := &SetLogFilePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogFilePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // IncludeModules
		in := &GetIncludeModulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIncludeModules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // IncludeModules
		in := &SetIncludeModulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetIncludeModules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // IncludeApis
		in := &GetIncludeAPIsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIncludeAPIs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // IncludeApis
		in := &SetIncludeAPIsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetIncludeAPIs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // ExcludeApis
		in := &GetExcludeAPIsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExcludeAPIs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // ExcludeApis
		in := &SetExcludeAPIsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExcludeAPIs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
