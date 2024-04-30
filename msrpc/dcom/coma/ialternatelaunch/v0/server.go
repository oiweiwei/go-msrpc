package ialternatelaunch

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

// IAlternateLaunch server interface.
type AlternateLaunchServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to create an alternate launch configuration, as
	// specified in section 3.1.1.4, for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF], section 2.1, on failure. All failure results
	// MUST be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	CreateConfiguration(context.Context, *CreateConfigurationRequest) (*CreateConfigurationResponse, error)

	// This method is called by a client to delete an alternate launch configuration (see
	// section 3.1.1.4) for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	DeleteConfiguration(context.Context, *DeleteConfigurationRequest) (*DeleteConfigurationResponse, error)
}

func RegisterAlternateLaunchServer(conn dcerpc.Conn, o AlternateLaunchServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAlternateLaunchServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AlternateLaunchSyntaxV0_0))...)
}

func NewAlternateLaunchServerHandle(o AlternateLaunchServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AlternateLaunchServerHandle(ctx, o, opNum, r)
	}
}

func AlternateLaunchServerHandle(ctx context.Context, o AlternateLaunchServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CreateConfiguration
		in := &CreateConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // DeleteConfiguration
		in := &DeleteConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
