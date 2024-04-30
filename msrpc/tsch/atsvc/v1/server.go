package atsvc

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

// atsvc server interface.
type ATSvcServer interface {

	// The NetrJobAdd method MUST add a single AT task to the server's task store.
	//
	// Return Values: For more information on return codes, see section 2.3.14 or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobAdd(context.Context, *JobAddRequest) (*JobAddResponse, error)

	// The NetrJobDel method MUST delete a specified range of tasks from the task store.
	// The method is capable of deleting all AT tasks or just a subset of the tasks, as
	// determined by the values of the MinJobId and MaxJobId parameters.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	//
	// To delete all tasks, specify MinJobId as 0 and MaxJobId as 0xFFFFFFFF.
	JobDelete(context.Context, *JobDeleteRequest) (*JobDeleteResponse, error)

	// The NetrJobEnum method MUST return an enumeration of all AT tasks on the specified
	// server.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobEnum(context.Context, *JobEnumRequest) (*JobEnumResponse, error)

	// The NetrJobGetInfo method MUST return information for a specified ATSvc task. The
	// task identifier MUST be used to locate the task configuration.
	//
	// Return Values: For more information on return codes, see section 2.3.14, or Win32
	// Error Codes in [MS-ERREF] section 2.1.
	JobGetInfo(context.Context, *JobGetInfoRequest) (*JobGetInfoResponse, error)
}

func RegisterATSvcServer(conn dcerpc.Conn, o ATSvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewATSvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ATSvcSyntaxV1_0))...)
}

func NewATSvcServerHandle(o ATSvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ATSvcServerHandle(ctx, o, opNum, r)
	}
}

func ATSvcServerHandle(ctx context.Context, o ATSvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrJobAdd
		in := &JobAddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.JobAdd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // NetrJobDel
		in := &JobDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.JobDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // NetrJobEnum
		in := &JobEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.JobEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // NetrJobGetInfo
		in := &JobGetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.JobGetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
