package iupdatedownloadcontent2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdatedownloadcontent "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdatedownloadcontent/v0"
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
	_ = iupdatedownloadcontent.GoPackage
)

// IUpdateDownloadContent2 server interface.
type UpdateDownloadContent2Server interface {

	// IUpdateDownloadContent base class.
	iupdatedownloadcontent.UpdateDownloadContentServer

	// The IUpdateDownloadContent2::IsDeltaCompressedContent (opnum 9) method retrieves
	// whether the content is delta-compressed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the IsDeltaCompressedContent ADM element.
	GetIsDeltaCompressedContent(context.Context, *GetIsDeltaCompressedContentRequest) (*GetIsDeltaCompressedContentResponse, error)
}

func RegisterUpdateDownloadContent2Server(conn dcerpc.Conn, o UpdateDownloadContent2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateDownloadContent2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateDownloadContent2SyntaxV0_0))...)
}

func NewUpdateDownloadContent2ServerHandle(o UpdateDownloadContent2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateDownloadContent2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateDownloadContent2ServerHandle(ctx context.Context, o UpdateDownloadContent2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 8 {
		// IUpdateDownloadContent base method.
		return iupdatedownloadcontent.UpdateDownloadContentServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 8: // IsDeltaCompressedContent
		op := &xxx_GetIsDeltaCompressedContentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsDeltaCompressedContentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsDeltaCompressedContent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateDownloadContent2
type UnimplementedUpdateDownloadContent2Server struct {
	iupdatedownloadcontent.UnimplementedUpdateDownloadContentServer
}

func (UnimplementedUpdateDownloadContent2Server) GetIsDeltaCompressedContent(context.Context, *GetIsDeltaCompressedContentRequest) (*GetIsDeltaCompressedContentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateDownloadContent2Server = (*UnimplementedUpdateDownloadContent2Server)(nil)
