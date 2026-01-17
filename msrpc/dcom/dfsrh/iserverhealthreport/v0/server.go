package iserverhealthreport

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IServerHealthReport server interface.
type ServerHealthReportServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetReport method retrieves health information for the specified replication group
	// that is hosted on the server in addition to the global health data of the DFS-R service
	// on the server.
	//
	// Return Values: The method MUST return 0 on success; or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)

	// The GetCompressedReport method gets the health information for the specified replication
	// group and the global health data of the DFS-R service on the server. The server MUST
	// encode the report as a field in the format that is specified by the DFS-R compression
	// algorithm in [MS-FRS2] section 3.1.1.1.
	//
	// Return Values: The method MUST return zero on success, or an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetCompressedReport(context.Context, *GetCompressedReportRequest) (*GetCompressedReportResponse, error)

	// The GetRawReportEx method is not currently in use and has never been implemented
	// in any version of the DFS-R Helper Protocol. It is reserved for future use.
	//
	// Return Values: The server MUST return the E_NOTIMPL error code (numeric value 0x80004001)
	// and take no action.<60>
	//
	//	+----------------------+------------------+
	//	|        RETURN        |                  |
	//	|      VALUE/CODE      |   DESCRIPTION    |
	//	|                      |                  |
	//	+----------------------+------------------+
	//	+----------------------+------------------+
	//	| 0x80004001 E_NOTIMPL | Not implemented. |
	//	+----------------------+------------------+
	GetRawReportEx(context.Context, *GetRawReportExRequest) (*GetRawReportExResponse, error)

	// The GetReferenceVersionVectors method gets the version vectors for all replicated
	// folders in the specified replication group.
	//
	// Return Values: The method MUST return zero on success or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	GetReferenceVersionVectors(context.Context, *GetReferenceVersionVectorsRequest) (*GetReferenceVersionVectorsResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The GetReferenceBacklogCounts method gets the outbound backlog for a replicated folder
	// on the member, relative to specific version vectors.
	//
	// Return Values: The method MUST return 0 on success or return an implementation-specific
	// nonzero HRESULT error code, as specified in [MS-ERREF] section 2.1, between 0x80000000
	// and 0xFFFFFFFF on failure. For protocol purposes, all nonzero values MUST be treated
	// as equivalent failures.
	//
	// After the server receives this message, it MUST get the backlog count for each version
	// vector that is supplied in the message parameters. If the server fails to retrieve
	// a backlog count, it returns a special value in the backlogCounts array at an index
	// that corresponds to the index in the flatMemberVersionVectors for the entry that
	// was used as input. The overall method MAY still return success (S_OK). <62> These
	// special values are as follows:
	//
	// * BACKLOG_CONTENT_SET_NOT_PRESENT (0xffffffff): The content set is not present in
	// DFS-R.
	//
	// * BACKLOG_ERROR_GET_BACKLOG_FAILED (0xfffffffe): Describes one or more of the following
	// conditions:
	//
	// * Run-time errors or implementation-specific errors that prevent the calculation
	// of the backlog count.
	//
	// * The flat member version vector could not be decompressed by using DFS-R. The DFS-R
	// decompression algorithm is specified in [MS-FRS2] ( ../ms-frs2/9914bdd4-9579-43a7-9f2d-9efe2e162944
	// ) section 3.1.1.2 ( ../ms-frs2/8cb5bae9-edf3-4833-9f0a-9d7e24218d3d ).
	//
	// * The version vector is empty (has a length of zero).
	//
	// The backlog counts MUST be saved in the backlogCounts output parameter.
	GetReferenceBacklogCounts(context.Context, *GetReferenceBacklogCountsRequest) (*GetReferenceBacklogCountsResponse, error)
}

func RegisterServerHealthReportServer(conn dcerpc.Conn, o ServerHealthReportServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServerHealthReportServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServerHealthReportSyntaxV0_0))...)
}

func NewServerHealthReportServerHandle(o ServerHealthReportServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServerHealthReportServerHandle(ctx, o, opNum, r)
	}
}

func ServerHealthReportServerHandle(ctx context.Context, o ServerHealthReportServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetReport
		op := &xxx_GetReportOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReportRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReport(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetCompressedReport
		op := &xxx_GetCompressedReportOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCompressedReportRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCompressedReport(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetRawReportEx
		op := &xxx_GetRawReportExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRawReportExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRawReportEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetReferenceVersionVectors
		op := &xxx_GetReferenceVersionVectorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReferenceVersionVectorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReferenceVersionVectors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // GetReferenceBacklogCounts
		op := &xxx_GetReferenceBacklogCountsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReferenceBacklogCountsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReferenceBacklogCounts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IServerHealthReport
type UnimplementedServerHealthReportServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedServerHealthReportServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServerHealthReportServer) GetCompressedReport(context.Context, *GetCompressedReportRequest) (*GetCompressedReportResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServerHealthReportServer) GetRawReportEx(context.Context, *GetRawReportExRequest) (*GetRawReportExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServerHealthReportServer) GetReferenceVersionVectors(context.Context, *GetReferenceVersionVectorsRequest) (*GetReferenceVersionVectorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedServerHealthReportServer) GetReferenceBacklogCounts(context.Context, *GetReferenceBacklogCountsRequest) (*GetReferenceBacklogCountsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ServerHealthReportServer = (*UnimplementedServerHealthReportServer)(nil)
