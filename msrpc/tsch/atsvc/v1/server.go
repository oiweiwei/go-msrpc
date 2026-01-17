package atsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_JobAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &JobAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.JobAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // NetrJobDel
		op := &xxx_JobDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &JobDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.JobDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NetrJobEnum
		op := &xxx_JobEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &JobEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.JobEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // NetrJobGetInfo
		op := &xxx_JobGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &JobGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.JobGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented atsvc
type UnimplementedATSvcServer struct {
}

func (UnimplementedATSvcServer) JobAdd(context.Context, *JobAddRequest) (*JobAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedATSvcServer) JobDelete(context.Context, *JobDeleteRequest) (*JobDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedATSvcServer) JobEnum(context.Context, *JobEnumRequest) (*JobEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedATSvcServer) JobGetInfo(context.Context, *JobGetInfoRequest) (*JobGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ATSvcServer = (*UnimplementedATSvcServer)(nil)
