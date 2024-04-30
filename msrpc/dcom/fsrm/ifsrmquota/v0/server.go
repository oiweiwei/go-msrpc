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
		in := &GetQuotaUsedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuotaUsed(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // QuotaPeakUsage
		in := &GetQuotaPeakUsageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuotaPeakUsage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // QuotaPeakUsageTime
		in := &GetQuotaPeakUsageTimeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuotaPeakUsageTime(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // ResetPeakUsage
		in := &ResetPeakUsageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResetPeakUsage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // RefreshUsageProperties
		in := &RefreshUsagePropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RefreshUsageProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
