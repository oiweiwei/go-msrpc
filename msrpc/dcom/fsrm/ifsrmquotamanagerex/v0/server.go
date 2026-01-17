package ifsrmquotamanagerex

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmquotamanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotamanager/v0"
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
	_ = ifsrmquotamanager.GoPackage
)

// IFsrmQuotaManagerEx server interface.
type QuotaManagerExServer interface {

	// IFsrmQuotaManager base class.
	ifsrmquotamanager.QuotaManagerServer

	// The IsAffectedByQuota method retrieves a value that determines whether a specified
	// path is subject to a Persisted Directory Quota (section 3.2.1.2.1.1).
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
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The affected parameter is NULL. |
	//	|                         | The options parameter is not a valid FsrmEnumOptions (section 2.2.1.2.5) value.  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80004003 E_POINTER    | The path parameter is NULL.                                                      |
	//	+-------------------------+----------------------------------------------------------------------------------+
	IsAffectedByQuota(context.Context, *IsAffectedByQuotaRequest) (*IsAffectedByQuotaResponse, error)
}

func RegisterQuotaManagerExServer(conn dcerpc.Conn, o QuotaManagerExServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQuotaManagerExServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QuotaManagerExSyntaxV0_0))...)
}

func NewQuotaManagerExServerHandle(o QuotaManagerExServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QuotaManagerExServerHandle(ctx, o, opNum, r)
	}
}

func QuotaManagerExServerHandle(ctx context.Context, o QuotaManagerExServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 19 {
		// IFsrmQuotaManager base method.
		return ifsrmquotamanager.QuotaManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 19: // IsAffectedByQuota
		op := &xxx_IsAffectedByQuotaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsAffectedByQuotaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsAffectedByQuota(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmQuotaManagerEx
type UnimplementedQuotaManagerExServer struct {
	ifsrmquotamanager.UnimplementedQuotaManagerServer
}

func (UnimplementedQuotaManagerExServer) IsAffectedByQuota(context.Context, *IsAffectedByQuotaRequest) (*IsAffectedByQuotaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QuotaManagerExServer = (*UnimplementedQuotaManagerExServer)(nil)
