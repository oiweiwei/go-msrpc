package iiisservicecontrol

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IIisServiceControl server interface.
type IISServiceControlServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// This method stops any running Internet services.<3>
	//
	// The server can have all functionality through this interface disabled using actions
	// taken local to the server machine. In this case the function MUST return an error
	// when called (E_ERROR_RESOURCE_DISABLED) and MUST NOT perform any other action.
	//
	// If the interface functionality is not disabled, the following actions SHOULD take
	// place on the server when this method is called:
	//
	// * The method SHOULD first attempt a graceful stop ( 546d32cd-905e-4f34-b023-2be4b5e16413#gt_7ea63e69-1023-48fd-bd40-f3729a350c06
	// ) of the services. If the caller has requested that the services be forced to stop
	// and the code either fails to request the stops or times out (based on the dwTimeoutMsecs
	// parameter) while waiting for the services to stop, it SHOULD terminate the processes
	// to ensure that they stop. This procedure SHOULD use the *Kill* ( 10ffdf93-a56f-4fc8-a3fd-5076135bc33b
	// ) method, as specified in section *3.1.4.5* , to handle the forced termination. <4>
	// ( 5c517f8f-7847-402a-b79e-4dbbf517997e#Appendix_A_4 )
	//
	// HRESULT Stop(
	//
	// DWORD dwTimeoutMsecs,
	//
	// # DWORD dwForce
	//
	// );
	//
	// Return Values: A signed, 32-bit value indicating return status. If the method returns
	// a negative value, it has failed. If the 12-bit facility code (bits 16–27) is set
	// to 0x007, the value contains a Win32 error code in the lower 16 bits. 0 or positive
	// values indicate success, with the lower 16 bits in positive nonzero values containing
	// warnings or flags defined in the method implementation. For more information about
	// HRESULT, see [MS-ERREF] section 2.1.
	//
	// The method MUST return S_OK (0x00000000) upon success.
	//
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                |
	//	|                 VALUE/CODE                 |                                  DESCRIPTION                                   |
	//	|                                            |                                                                                |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                       |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070008 E_ERROR_NOT_ENOUGH_MEMORY       | Not enough memory is available to process this command.                        |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x8007041D E_ERROR_SERVICE_REQUEST_TIMEOUT | A time-out has occurred while waiting for the Internet services to be stopped. |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x800710D5 E_ERROR_RESOURCE_DISABLED       | The IIisServiceControl interface is disabled.                                  |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//
	// If the length of time specified by dwTimeoutMsecs has elapsed and not all services
	// have stopped, and if dwForce is set to 0x00000001 (True), then the remaining services
	// SHOULD be forced to terminate.
	Stop(context.Context, *StopRequest) (*StopResponse, error)

	// This method is used to start the Internet services.
	//
	// The server can have all functionality through this interface disabled using actions
	// taken local to the server. In this case the function MUST return an error when called
	// (E_ERROR_RESOURCE_DISABLED) and MUST NOT perform any other action.
	//
	// If the interface functionality is not disabled, the following SHOULD take place on
	// the server when this method is called:
	//
	// * The method SHOULD <5> ( 5c517f8f-7847-402a-b79e-4dbbf517997e#Appendix_A_5 ) start
	// all Internet services that are marked to start automatically when the computer starts
	// up.
	//
	// HRESULT Start(
	//
	// # DWORD dwTimeoutMsecs
	//
	// );
	//
	// Return Values: A signed, 32-bit value indicating return status. If the method returns
	// a negative value, it has failed. If the 12-bit facility code (bits 16–27) is set
	// to 0x007, the value contains a Win32 error code in the lower 16 bits. 0 or positive
	// values indicate success, with the lower 16 bits in positive nonzero values containing
	// warnings or flags defined in the method implementation. For more information about
	// HRESULT, see [MS-ERREF] section 2.1.
	//
	// The method MUST return S_OK (0x00000000) upon success.
	//
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                |
	//	|                 VALUE/CODE                 |                                  DESCRIPTION                                   |
	//	|                                            |                                                                                |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                       |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x80070008 E_ERROR_NOT_ENOUGH_MEMORY       | Not enough memory is available to process this command.                        |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x8007041D E_ERROR_SERVICE_REQUEST_TIMEOUT | A time-out has occurred while waiting for all Internet services to be started. |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x800710D5 E_ERROR_RESOURCE_DISABLED       | The IIisServiceControl Interface is disabled.                                  |
	//	+--------------------------------------------+--------------------------------------------------------------------------------+
	Start(context.Context, *StartRequest) (*StartResponse, error)

	// This method is used to reboot the computer where the IIS service is running.
	//
	// The server implementation MAY<6> not implement this function. If it does not, then
	// it MUST return E_NOTIMPL.
	//
	// Return Values: A signed, 32-bit value indicating return status. If the method returns
	// a negative value, it has failed. If the 12-bit facility code (bits 16–27) is set
	// to 0x007, the value contains a Win32 error code in the lower 16 bits. 0 or positive
	// values indicate success, with the lower 16 bits in positive nonzero values containing
	// warnings or flags defined in the method implementation. For more information about
	// HRESULT, see [MS-ERREF] section 2.1.
	//
	// The method MUST return S_OK (0x00000000) upon success.
	//
	//	+--------------------------------------+----------------------------------------------------------------+
	//	|                RETURN                |                                                                |
	//	|              VALUE/CODE              |                          DESCRIPTION                           |
	//	|                                      |                                                                |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                       |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| 0x80070008 E_ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.        |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| 0x800710D5 E_ERROR_RESOURCE_DISABLED | The IIisServiceControl interface is disabled.                  |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| 0x80004001 E_NOTIMPL                 | This function is not supported for this version of the server. |
	//	+--------------------------------------+----------------------------------------------------------------+
	Reboot(context.Context, *RebootRequest) (*RebootResponse, error)

	// This method returns the status of the Internet services.
	//
	// The server can have all functionality through this interface disabled using actions
	// taken local to the server machine. In this case the function MUST return an error
	// when called (E_ERROR_RESOURCE_DISABLED) and MUST NOT perform any other action.
	//
	// If the interface functionality is not disabled, the following SHOULD take place on
	// the server when this method is called:
	//
	// * The method SHOULD return a buffer of unsigned chars as described in section 2.2.2
	// ( 096ffe89-76be-4d01-9e4d-f68428a231fc ). This buffer of unsigned chars MUST contain
	// data about the status of the Internet services.
	//
	// * If it is not possible to return all the data in the buffer provided, then the following
	// conditional behavior MUST occur.
	//
	// For more information about the unsigned char buffer returned, see section 2.2.2.
	//
	// Return Values: A signed, 32-bit value indicating return status. If the method returns
	// a negative value, it has failed. If the 12-bit facility code (bits 16–27) is set
	// to 0x007, the value contains a Win32 error code in the lower 16 bits. 0 or positive
	// values indicate success, with the lower 16 bits in positive nonzero values containing
	// warnings or flags defined in the method implementation. For more information about
	// HRESULT, see [MS-ERREF] section 2.1.
	//
	// The method MUST return S_OK (0x00000000) upon success.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                        | The call was successful.                                                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007007A E_ERROR_INSUFFICIENT_BUFFER | The size of the pbBuffer is too small to return the status data based on its     |
	//	|                                        | size being declared in dwBufferSize parameter.                                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070008 E_ERROR_NOT_ENOUGH_MEMORY   | Not enough memory is available to process this command.                          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800710D5 E_ERROR_RESOURCE_DISABLED   | The IIisServiceControl interface is disabled.                                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	Status(context.Context, *StatusRequest) (*StatusResponse, error)

	// This method is used to terminate the Internet services processes. This erases the
	// IIS processes from memory, and is used to recover from failed instances of IIS processes.
	//
	// The server can have all functionality through this interface disabled using actions
	// taken local to the server machine. In this case the function MUST return an error
	// when called (E_ERROR_RESOURCE_DISABLED) and MUST NOT perform any other action.
	//
	// If the interface functionality is not disabled, the following SHOULD take place on
	// the server when this method is called:
	//
	// * The method SHOULD terminate all processes involved in supporting the Internet services
	// on the server.
	//
	// How the processes are terminated is implementation-dependent.<7>
	//
	// This method has no parameters.
	//
	// Return Values: A signed, 32-bit value indicating return status. If the method returns
	// a negative value, it has failed. If the 12-bit facility code (bits 16–27) is set
	// to 0x007, the value contains a Win32 error code in the lower 16 bits. 0 or positive
	// values indicate success, with the lower 16 bits in positive nonzero values containing
	// warnings or flags defined in the method implementation. For more information about
	// HRESULT, see [MS-ERREF] section 2.1.
	//
	// Each of the values that follow where the first byte contains 0x8007 is the HRESULT
	// derived from the Win32 error code with the specified name.
	//
	// The method MUST return S_OK (0x00000000) upon success.
	//
	//	+--------------------------------------+---------------------------------------------------------+
	//	|                RETURN                |                                                         |
	//	|              VALUE/CODE              |                       DESCRIPTION                       |
	//	|                                      |                                                         |
	//	+--------------------------------------+---------------------------------------------------------+
	//	+--------------------------------------+---------------------------------------------------------+
	//	| 0x00000000 S_OK                      | The call was successful.                                |
	//	+--------------------------------------+---------------------------------------------------------+
	//	| 0x80070008 E_ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+--------------------------------------+---------------------------------------------------------+
	//	| 0x800710D5 E_ERROR_RESOURCE_DISABLED | The IIisServiceControl interface is disabled.           |
	//	+--------------------------------------+---------------------------------------------------------+
	Kill(context.Context, *KillRequest) (*KillResponse, error)
}

func RegisterIISServiceControlServer(conn dcerpc.Conn, o IISServiceControlServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIISServiceControlServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IISServiceControlSyntaxV0_0))...)
}

func NewIISServiceControlServerHandle(o IISServiceControlServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IISServiceControlServerHandle(ctx, o, opNum, r)
	}
}

func IISServiceControlServerHandle(ctx context.Context, o IISServiceControlServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Stop
		op := &xxx_StopOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Stop(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Start
		op := &xxx_StartOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Start(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Reboot
		op := &xxx_RebootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RebootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reboot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Status
		op := &xxx_StatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Status(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Kill
		op := &xxx_KillOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &KillRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Kill(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IIisServiceControl
type UnimplementedIISServiceControlServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedIISServiceControlServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISServiceControlServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISServiceControlServer) Reboot(context.Context, *RebootRequest) (*RebootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISServiceControlServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIISServiceControlServer) Kill(context.Context, *KillRequest) (*KillResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IISServiceControlServer = (*UnimplementedIISServiceControlServer)(nil)
