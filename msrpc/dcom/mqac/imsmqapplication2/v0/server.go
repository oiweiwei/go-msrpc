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

	// The RegisterCertificate method MUST register an MQUSERSIGNCERT ([MS-MQMQ] section
	// 2.2.22) in User.Certificates. Implementations of this protocol use certificates to
	// verify the sender for messages that are requesting authentication and to ensure message
	// integrity.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<10>
	RegisterCertificate(context.Context, *RegisterCertificateRequest) (*RegisterCertificateResponse, error)

	// The MachineNameOfMachineId method is received by the server in an RPC_REQUEST packet.
	// In response, the server MUST return a string that contains the QueueManager.ComputerName
	// of the QueueManager that is identified by the bstrGuid input parameter.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.<12>
	MachineNameOfMachineID(context.Context, *MachineNameOfMachineIDRequest) (*MachineNameOfMachineIDResponse, error)

	// The MSMQVersionMajor method is received by the server in an RPC_REQUEST packet. In
	// response, the serverMUST return the major version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetVersionMajor(context.Context, *GetVersionMajorRequest) (*GetVersionMajorResponse, error)

	// The MSMQVersionMinor method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the minor version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// When the server processes this call, it MUST follow these guidelines:
	//
	// * If the *ComputerName* instance variable is not NULL, return MQ_ERROR_INVALID_PARAMETER
	// (0xC00E0006) and take no further action.
	//
	// * Set the psMSMQVersionMinor output variable to the minor version number <14> ( 71c359c3-e9ec-4fe6-a101-aab1eabecdcf#Appendix_A_14
	// ) of the server.
	GetVersionMinor(context.Context, *GetVersionMinorRequest) (*GetVersionMinorResponse, error)

	// The MSMQVersionBuild method is received by the server in an RPC_REQUEST packet. In
	// response, the server MUST return the build version number of the server.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	GetVersionBuild(context.Context, *GetVersionBuildRequest) (*GetVersionBuildResponse, error)

	// The IsDsEnabled method is received by the server in an RPC_REQUEST packet. In response,
	// the server MUST return a BOOLEAN value that indicates whether the local QueueManager
	// is configured to use the directory.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
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
