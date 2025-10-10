package imsmqapplication2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqapplication "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqapplication/v0"
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
	_ = imsmqapplication.GoPackage
)

// IMSMQApplication2 server interface.
type Application2Server interface {

	// IMSMQApplication base class.
	imsmqapplication.ApplicationServer

	// RegisterCertificate operation.
	RegisterCertificate(context.Context, *RegisterCertificateRequest) (*RegisterCertificateResponse, error)

	// MachineNameOfMachineId operation.
	MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest) (*MachineNameOfMachineIDResponse, error)

	// MSMQVersionMajor operation.
	GetVersionMajor(context.Context, *GetVersionMajorRequest) (*GetVersionMajorResponse, error)

	// MSMQVersionMinor operation.
	GetVersionMinor(context.Context, *GetVersionMinorRequest) (*GetVersionMinorResponse, error)

	// MSMQVersionBuild operation.
	GetVersionBuild(context.Context, *GetVersionBuildRequest) (*GetVersionBuildResponse, error)

	// IsDsEnabled operation.
	GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest) (*GetIsDSEnabledResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterApplication2Server(conn dcerpc.Conn, o Application2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewApplication2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Application2SyntaxV0_0))...)
}

func NewApplication2ServerHandle(o Application2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Application2ServerHandle(ctx, o, opNum, r)
	}
}

func Application2ServerHandle(ctx context.Context, o Application2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 8 {
		// IMSMQApplication base method.
		return imsmqapplication.ApplicationServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 8: // RegisterCertificate
		op := &xxx_RegisterCertificateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterCertificateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterCertificate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // MachineNameOfMachineId
		op := &xxx_MachineNameOfMachineIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MachineNameOfMachineIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MachineNameOfMachineID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // MSMQVersionMajor
		op := &xxx_GetVersionMajorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionMajorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersionMajor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // MSMQVersionMinor
		op := &xxx_GetVersionMinorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionMinorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersionMinor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // MSMQVersionBuild
		op := &xxx_GetVersionBuildOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionBuildRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersionBuild(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IsDsEnabled
		op := &xxx_GetIsDSEnabledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsDSEnabledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsDSEnabled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Properties
		op := &xxx_GetPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQApplication2
type UnimplementedApplication2Server struct {
	imsmqapplication.UnimplementedApplicationServer
}

func (UnimplementedApplication2Server) RegisterCertificate(context.Context, *RegisterCertificateRequest) (*RegisterCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest) (*MachineNameOfMachineIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) GetVersionMajor(context.Context, *GetVersionMajorRequest) (*GetVersionMajorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) GetVersionMinor(context.Context, *GetVersionMinorRequest) (*GetVersionMinorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) GetVersionBuild(context.Context, *GetVersionBuildRequest) (*GetVersionBuildResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest) (*GetIsDSEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedApplication2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Application2Server = (*UnimplementedApplication2Server)(nil)
