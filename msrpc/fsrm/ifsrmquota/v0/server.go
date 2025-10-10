package ifsrmquota

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmquotaobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotaobject/v0"
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
	_ = ifsrmquotaobject.GoPackage
)

// IFsrmQuota server interface.
type QuotaServer interface {

	// IFsrmQuotaObject base class.
	ifsrmquotaobject.QuotaObjectServer

	// The QuotaUsed (get) method returns the current, read-only quota usage for this quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------+
	//	|         RETURN          |                             |
	//	|       VALUE/CODE        |         DESCRIPTION         |
	//	|                         |                             |
	//	+-------------------------+-----------------------------+
	//	+-------------------------+-----------------------------+
	//	| 0x80070057 E_INVALIDARG | The used parameter is NULL. |
	//	+-------------------------+-----------------------------+
	GetQuotaUsed(context.Context, *GetQuotaUsedRequest) (*GetQuotaUsedResponse, error)

	// The QuotaPeakUsage (get) method returns the peak quota usage of this quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------+
	//	|         RETURN          |                                  |
	//	|       VALUE/CODE        |           DESCRIPTION            |
	//	|                         |                                  |
	//	+-------------------------+----------------------------------+
	//	+-------------------------+----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The peakUsage parameter is NULL. |
	//	+-------------------------+----------------------------------+
	GetQuotaPeakUsage(context.Context, *GetQuotaPeakUsageRequest) (*GetQuotaPeakUsageResponse, error)

	// The QuotaPeakUsageTime (get) method returns the peak quota usage time stamp of this
	// quota.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+------------------------------------------+
	//	|         RETURN          |                                          |
	//	|       VALUE/CODE        |               DESCRIPTION                |
	//	|                         |                                          |
	//	+-------------------------+------------------------------------------+
	//	+-------------------------+------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The peakUsageDateTime parameter is NULL. |
	//	+-------------------------+------------------------------------------+
	GetQuotaPeakUsageTime(context.Context, *GetQuotaPeakUsageTimeRequest) (*GetQuotaPeakUsageTimeResponse, error)

	// The ResetPeakUsage method resets the peak quota usage of this quota to zero and returns
	// S_OK upon successful completion.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	// There are no parameters for this method.
	ResetPeakUsage(context.Context, *ResetPeakUsageRequest) (*ResetPeakUsageResponse, error)

	// The RefreshUsageProperties method refreshes the quota usage information for the caller's
	// copy of the object.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+-----------------------------------------+
	//	|           RETURN            |                                         |
	//	|         VALUE/CODE          |               DESCRIPTION               |
	//	|                             |                                         |
	//	+-----------------------------+-----------------------------------------+
	//	+-----------------------------+-----------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | The specified quota could not be found. |
	//	+-----------------------------+-----------------------------------------+
	//
	// There are no parameters for this method.
	//
	// If no Persisted Directory Quota exists that has the same Directory Quota.Folder path
	// property as Non-Persisted Directory Quota Instance, the server MUST return FSRM_E_NOT_FOUND.
	//
	// Otherwise, the server MUST reset the quota usage, quota peak usage, and quota peak
	// usage time of the Non-Persisted Directory Quota Instance to the current values stored
	// in the corresponding properties of the Persisted Directory Quota that has the same
	// Directory Quota.Folder path property as this Non-Persisted Directory Quota Instance.
	RefreshUsageProperties(context.Context, *RefreshUsagePropertiesRequest) (*RefreshUsagePropertiesResponse, error)
}

func RegisterQuotaServer(conn dcerpc.Conn, o QuotaServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaSyntaxV0_0))...)
}

func NewQuotaServerHandle(o QuotaServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaServerHandle(ctx, o, opNum, r)
	}
}

func QuotaServerHandle(ctx context.Context, o QuotaServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 28 {
		// IFsrmQuotaObject base method.
		return ifsrmquotaobject.QuotaObjectServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 28: // QuotaUsed
		op := &xxx_GetQuotaUsedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaUsedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuotaUsed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // QuotaPeakUsage
		op := &xxx_GetQuotaPeakUsageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaPeakUsageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuotaPeakUsage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // QuotaPeakUsageTime
		op := &xxx_GetQuotaPeakUsageTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuotaPeakUsageTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuotaPeakUsageTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // ResetPeakUsage
		op := &xxx_ResetPeakUsageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetPeakUsageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResetPeakUsage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RefreshUsageProperties
		op := &xxx_RefreshUsagePropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshUsagePropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshUsageProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuota
type UnimplementedQuotaServer struct {
	ifsrmquotaobject.UnimplementedQuotaObjectServer
}

func (UnimplementedQuotaServer) GetQuotaUsed(context.Context, *GetQuotaUsedRequest) (*GetQuotaUsedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaServer) GetQuotaPeakUsage(context.Context, *GetQuotaPeakUsageRequest) (*GetQuotaPeakUsageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaServer) GetQuotaPeakUsageTime(context.Context, *GetQuotaPeakUsageTimeRequest) (*GetQuotaPeakUsageTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaServer) ResetPeakUsage(context.Context, *ResetPeakUsageRequest) (*ResetPeakUsageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQuotaServer) RefreshUsageProperties(context.Context, *RefreshUsagePropertiesRequest) (*RefreshUsagePropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaServer = (*UnimplementedQuotaServer)(nil)
