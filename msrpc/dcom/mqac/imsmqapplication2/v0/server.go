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
type ImsmqApplication2Server interface {

	// IMSMQApplication base class.
	imsmqapplication.ImsmqApplicationServer

	// RegisterCertificate operation.
	RegisterCertificate(context.Context, *RegisterCertificateRequest) (*RegisterCertificateResponse, error)

	// MachineNameOfMachineId operation.
	MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest) (*MachineNameOfMachineIDResponse, error)

	// MSMQVersionMajor operation.
	GetMsmqVersionMajor(context.Context, *GetMsmqVersionMajorRequest) (*GetMsmqVersionMajorResponse, error)

	// MSMQVersionMinor operation.
	GetMsmqVersionMinor(context.Context, *GetMsmqVersionMinorRequest) (*GetMsmqVersionMinorResponse, error)

	// MSMQVersionBuild operation.
	GetMsmqVersionBuild(context.Context, *GetMsmqVersionBuildRequest) (*GetMsmqVersionBuildResponse, error)

	// IsDsEnabled operation.
	GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest) (*GetIsDSEnabledResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)
}

func RegisterImsmqApplication2Server(conn dcerpc.Conn, o ImsmqApplication2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqApplication2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqApplication2SyntaxV0_0))...)
}

func NewImsmqApplication2ServerHandle(o ImsmqApplication2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqApplication2ServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqApplication2ServerHandle(ctx context.Context, o ImsmqApplication2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 8 {
		// IMSMQApplication base method.
		return imsmqapplication.ImsmqApplicationServerHandle(ctx, o, opNum, r)
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
		op := &xxx_GetMsmqVersionMajorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMsmqVersionMajorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMsmqVersionMajor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // MSMQVersionMinor
		op := &xxx_GetMsmqVersionMinorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMsmqVersionMinorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMsmqVersionMinor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // MSMQVersionBuild
		op := &xxx_GetMsmqVersionBuildOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMsmqVersionBuildRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMsmqVersionBuild(ctx, req)
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
type UnimplementedImsmqApplication2Server struct {
	imsmqapplication.UnimplementedImsmqApplicationServer
}

func (UnimplementedImsmqApplication2Server) RegisterCertificate(context.Context, *RegisterCertificateRequest) (*RegisterCertificateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest) (*MachineNameOfMachineIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) GetMsmqVersionMajor(context.Context, *GetMsmqVersionMajorRequest) (*GetMsmqVersionMajorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) GetMsmqVersionMinor(context.Context, *GetMsmqVersionMinorRequest) (*GetMsmqVersionMinorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) GetMsmqVersionBuild(context.Context, *GetMsmqVersionBuildRequest) (*GetMsmqVersionBuildResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) GetIsDSEnabled(context.Context, *GetIsDSEnabledRequest) (*GetIsDSEnabledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqApplication2Server) GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqApplication2Server = (*UnimplementedImsmqApplication2Server)(nil)
