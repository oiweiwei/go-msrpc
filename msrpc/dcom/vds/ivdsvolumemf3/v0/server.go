package ivdsvolumemf3

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

// IVdsVolumeMF3 server interface.
type VolumeMF3Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryVolumeGuidPathnames method returns a volume's volume GUID path names.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVolumeGUIDPathnames(context.Context, *QueryVolumeGUIDPathnamesRequest) (*QueryVolumeGUIDPathnamesResponse, error)

	// The FormatEx2 method formats a file system on a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatEx2(context.Context, *FormatEx2Request) (*FormatEx2Response, error)

	// The OfflineVolume method offlines a volume. An offline volume will fail data IO.
	// The volume can be opened for configuration.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	OfflineVolume(context.Context, *OfflineVolumeRequest) (*OfflineVolumeResponse, error)
}

func RegisterVolumeMF3Server(conn dcerpc.Conn, o VolumeMF3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVolumeMF3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VolumeMF3SyntaxV0_0))...)
}

func NewVolumeMF3ServerHandle(o VolumeMF3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VolumeMF3ServerHandle(ctx, o, opNum, r)
	}
}

func VolumeMF3ServerHandle(ctx context.Context, o VolumeMF3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryVolumeGuidPathnames
		in := &QueryVolumeGUIDPathnamesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryVolumeGUIDPathnames(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // FormatEx2
		in := &FormatEx2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FormatEx2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // OfflineVolume
		in := &OfflineVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OfflineVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
