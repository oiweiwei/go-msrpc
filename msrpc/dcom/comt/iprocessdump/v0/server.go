package iprocessdump

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = idispatch.GoPackage
)

// IProcessDump server interface.
type ProcessDumpServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// This method is called by a client to determine whether or not the COMT server supports
	// process dump.
	//
	// This method has no parameters.
	//
	// Return Values: This method returns S_OK (0x00000000) if the COMT server supports
	// process dump, and MUST return S_FALSE (0x00000001) if not.
	IsSupported(context.Context, *IsSupportedRequest) (*IsSupportedResponse, error)

	// This method is called by a client to request a process dump for the process containing
	// a particular instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	DumpProcess(context.Context, *DumpProcessRequest) (*DumpProcessResponse, error)
}

func RegisterProcessDumpServer(conn dcerpc.Conn, o ProcessDumpServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewProcessDumpServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ProcessDumpSyntaxV0_0))...)
}

func NewProcessDumpServerHandle(o ProcessDumpServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ProcessDumpServerHandle(ctx, o, opNum, r)
	}
}

func ProcessDumpServerHandle(ctx context.Context, o ProcessDumpServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // IsSupported
		in := &IsSupportedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsSupported(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // DumpProcess
		in := &DumpProcessRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DumpProcess(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
