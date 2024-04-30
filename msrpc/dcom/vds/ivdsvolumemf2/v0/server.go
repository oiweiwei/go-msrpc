package ivdsvolumemf2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// IVdsVolumeMF2 server interface.
type VolumeMF2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetFileSystemTypeName method retrieves the name of the file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetFileSystemTypeName(context.Context, *GetFileSystemTypeNameRequest) (*GetFileSystemTypeNameResponse, error)

	// The QueryFileSystemFormatSupport method retrieves the properties of the file systems
	// that are supported for formatting a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryFileSystemFormatSupport(context.Context, *QueryFileSystemFormatSupportRequest) (*QueryFileSystemFormatSupportResponse, error)

	// The FormatEx method formats a file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatEx(context.Context, *FormatExRequest) (*FormatExResponse, error)
}

func RegisterVolumeMF2Server(conn dcerpc.Conn, o VolumeMF2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeMF2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeMF2SyntaxV0_0))...)
}

func NewVolumeMF2ServerHandle(o VolumeMF2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeMF2ServerHandle(ctx, o, opNum, r)
	}
}

func VolumeMF2ServerHandle(ctx context.Context, o VolumeMF2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetFileSystemTypeName
		in := &GetFileSystemTypeNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileSystemTypeName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // QueryFileSystemFormatSupport
		in := &QueryFileSystemFormatSupportRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryFileSystemFormatSupport(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // FormatEx
		in := &FormatExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FormatEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
