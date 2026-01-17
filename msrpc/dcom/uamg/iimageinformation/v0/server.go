package iimageinformation

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IImageInformation server interface.
type ImageInformationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IImageInformation::AltText (opnum 8) method retrieves the alternate text for
	// the image.
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
	// The server SHOULD return the value of the AltText ADM element.
	GetAltText(context.Context, *GetAltTextRequest) (*GetAltTextResponse, error)

	// The IImageInformation::Height (opnum 9) method retrieves the height of the image,
	// in pixels.
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
	// This method SHOULD return the value of the Height ADM element.
	GetHeight(context.Context, *GetHeightRequest) (*GetHeightResponse, error)

	// The IImageInformation::Source (opnum 10) method retrieves the source location of
	// the image.
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
	// This method SHOULD return the value of the Source ADM element.
	GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error)

	// The IImageInformation::Width (opnum 11) method retrieves the width of the image,
	// in pixels.
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
	// This method SHOULD return the value of the Width ADM element.
	GetWidth(context.Context, *GetWidthRequest) (*GetWidthResponse, error)
}

func RegisterImageInformationServer(conn dcerpc.Conn, o ImageInformationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImageInformationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImageInformationSyntaxV0_0))...)
}

func NewImageInformationServerHandle(o ImageInformationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImageInformationServerHandle(ctx, o, opNum, r)
	}
}

func ImageInformationServerHandle(ctx context.Context, o ImageInformationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // AltText
		op := &xxx_GetAltTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAltTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAltText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Height
		op := &xxx_GetHeightOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHeightRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHeight(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Source
		op := &xxx_GetSourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Width
		op := &xxx_GetWidthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetWidthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetWidth(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IImageInformation
type UnimplementedImageInformationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedImageInformationServer) GetAltText(context.Context, *GetAltTextRequest) (*GetAltTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImageInformationServer) GetHeight(context.Context, *GetHeightRequest) (*GetHeightResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImageInformationServer) GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImageInformationServer) GetWidth(context.Context, *GetWidthRequest) (*GetWidthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImageInformationServer = (*UnimplementedImageInformationServer)(nil)
