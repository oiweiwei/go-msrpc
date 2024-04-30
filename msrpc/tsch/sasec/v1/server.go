package sasec

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// sasec server interface.
type SasecServer interface {

	// The SASetAccountInformation method MUST set the credentials under which the task
	// MUST run.
	//
	// Return Values: For more information about return codes, see section 2.3.14 or Win32
	// Error Codes in [MS-ERREF] section 2.1.<52>
	SetAccountInformation(context.Context, *SetAccountInformationRequest) (*SetAccountInformationResponse, error)

	// The SASetNSAccountInformation method MUST configure the credentials under which all
	// ATSvc tasks run.
	SetNSAccountInformation(context.Context, *SetNSAccountInformationRequest) (*SetNSAccountInformationResponse, error)

	// The SAGetNSAccountInformation method MUST return the ATSvc account name.
	//
	// Return Values: For more information about return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.
	GetNSAccountInformation(context.Context, *GetNSAccountInformationRequest) (*GetNSAccountInformationResponse, error)

	// The SAGetAccountInformation method MUST retrieve the account name for a specified
	// task.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	GetAccountInformation(context.Context, *GetAccountInformationRequest) (*GetAccountInformationResponse, error)
}

func RegisterSasecServer(conn dcerpc.Conn, o SasecServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSasecServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SasecSyntaxV1_0))...)
}

func NewSasecServerHandle(o SasecServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SasecServerHandle(ctx, o, opNum, r)
	}
}

func SasecServerHandle(ctx context.Context, o SasecServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // SASetAccountInformation
		in := &SetAccountInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAccountInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // SASetNSAccountInformation
		in := &SetNSAccountInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNSAccountInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // SAGetNSAccountInformation
		in := &GetNSAccountInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNSAccountInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // SAGetAccountInformation
		in := &GetAccountInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAccountInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
