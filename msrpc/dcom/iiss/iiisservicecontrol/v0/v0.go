package iiisservicecontrol

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iiss"
)

var (
	// IIisServiceControl interface identifier e8fb8620-588f-11d2-9d61-00c04f79c5fe
	IISServiceControlIID = &dcom.IID{Data1: 0xe8fb8620, Data2: 0x588f, Data3: 0x11d2, Data4: []byte{0x9d, 0x61, 0x00, 0xc0, 0x4f, 0x79, 0xc5, 0xfe}}
	// Syntax UUID
	IISServiceControlSyntaxUUID = &uuid.UUID{TimeLow: 0xe8fb8620, TimeMid: 0x588f, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x61, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x79, 0xc5, 0xfe}}
	// Syntax ID
	IISServiceControlSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IISServiceControlSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IIisServiceControl interface.
type IISServiceControlClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	Stop(context.Context, *StopRequest, ...dcerpc.CallOption) (*StopResponse, error)

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
	Start(context.Context, *StartRequest, ...dcerpc.CallOption) (*StartResponse, error)

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
	Reboot(context.Context, *RebootRequest, ...dcerpc.CallOption) (*RebootResponse, error)

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
	Status(context.Context, *StatusRequest, ...dcerpc.CallOption) (*StatusResponse, error)

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
	Kill(context.Context, *KillRequest, ...dcerpc.CallOption) (*KillResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IISServiceControlClient
}

type xxx_DefaultIISServiceControlClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIISServiceControlClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultIISServiceControlClient) Stop(ctx context.Context, in *StopRequest, opts ...dcerpc.CallOption) (*StopResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StopResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISServiceControlClient) Start(ctx context.Context, in *StartRequest, opts ...dcerpc.CallOption) (*StartResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISServiceControlClient) Reboot(ctx context.Context, in *RebootRequest, opts ...dcerpc.CallOption) (*RebootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RebootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISServiceControlClient) Status(ctx context.Context, in *StatusRequest, opts ...dcerpc.CallOption) (*StatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISServiceControlClient) Kill(ctx context.Context, in *KillRequest, opts ...dcerpc.CallOption) (*KillResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &KillResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIISServiceControlClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIISServiceControlClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIISServiceControlClient) IPID(ctx context.Context, ipid *dcom.IPID) IISServiceControlClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIISServiceControlClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewIISServiceControlClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IISServiceControlClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IISServiceControlSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultIISServiceControlClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_StopOperation structure represents the Stop operation
type xxx_StopOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	TimeoutMsecs uint32         `idl:"name:dwTimeoutMsecs" json:"timeout_msecs"`
	Force        uint32         `idl:"name:dwForce" json:"force"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StopOperation) OpNum() int { return 7 }

func (o *xxx_StopOperation) OpName() string { return "/IIisServiceControl/v0/Stop" }

func (o *xxx_StopOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwTimeoutMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TimeoutMsecs); err != nil {
			return err
		}
	}
	// dwForce {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTimeoutMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TimeoutMsecs); err != nil {
			return err
		}
	}
	// dwForce {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StopRequest structure represents the Stop operation request
type StopRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwTimeoutMsecs: Length of time allowed for services to stop. If this time has elapsed,
	// and not all services have stopped, then the conditional behavior that follows SHOULD
	// occur.
	TimeoutMsecs uint32 `idl:"name:dwTimeoutMsecs" json:"timeout_msecs"`
	// dwForce: Boolean value that specifies whether the services will be forced to terminate.
	// If the graceful stopping of any service fails, then the conditional behavior that
	// follows SHOULD occur.
	//
	//	+------------------+-------------------------------------------+
	//	|                  |                                           |
	//	|      VALUE       |                  MEANING                  |
	//	|                  |                                           |
	//	+------------------+-------------------------------------------+
	//	+------------------+-------------------------------------------+
	//	| TRUE 0x00000001  | Services MUST be forced to terminate.     |
	//	+------------------+-------------------------------------------+
	//	| FALSE 0x00000000 | Services MUST NOT be forced to terminate. |
	//	+------------------+-------------------------------------------+
	Force uint32 `idl:"name:dwForce" json:"force"`
}

func (o *StopRequest) xxx_ToOp(ctx context.Context, op *xxx_StopOperation) *xxx_StopOperation {
	if op == nil {
		op = &xxx_StopOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TimeoutMsecs = o.TimeoutMsecs
	op.Force = o.Force
	return op
}

func (o *StopRequest) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TimeoutMsecs = op.TimeoutMsecs
	o.Force = op.Force
}
func (o *StopRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopResponse structure represents the Stop operation response
type StopResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Stop return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopResponse) xxx_ToOp(ctx context.Context, op *xxx_StopOperation) *xxx_StopOperation {
	if op == nil {
		op = &xxx_StopOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *StopResponse) xxx_FromOp(ctx context.Context, op *xxx_StopOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StopResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartOperation structure represents the Start operation
type xxx_StartOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	TimeoutMsecs uint32         `idl:"name:dwTimeoutMsecs" json:"timeout_msecs"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StartOperation) OpNum() int { return 8 }

func (o *xxx_StartOperation) OpName() string { return "/IIisServiceControl/v0/Start" }

func (o *xxx_StartOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwTimeoutMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TimeoutMsecs); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTimeoutMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TimeoutMsecs); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartRequest structure represents the Start operation request
type StartRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwTimeoutMsecs: Length of time, in milliseconds, allowed to start the services. After
	// this time has passed, the server MUST return 0x8000041D (E_ERROR_SERVICE_REQUEST_TIMEOUT).
	TimeoutMsecs uint32 `idl:"name:dwTimeoutMsecs" json:"timeout_msecs"`
}

func (o *StartRequest) xxx_ToOp(ctx context.Context, op *xxx_StartOperation) *xxx_StartOperation {
	if op == nil {
		op = &xxx_StartOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TimeoutMsecs = o.TimeoutMsecs
	return op
}

func (o *StartRequest) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TimeoutMsecs = op.TimeoutMsecs
}
func (o *StartRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartResponse structure represents the Start operation response
type StartResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Start return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartResponse) xxx_ToOp(ctx context.Context, op *xxx_StartOperation) *xxx_StartOperation {
	if op == nil {
		op = &xxx_StartOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *StartResponse) xxx_FromOp(ctx context.Context, op *xxx_StartOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StartResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RebootOperation structure represents the Reboot operation
type xxx_RebootOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	TimeoutMsecs    uint32         `idl:"name:dwTimeouMsecs" json:"timeout_msecs"`
	ForceAppsClosed uint32         `idl:"name:dwForceAppsClosed" json:"force_apps_closed"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RebootOperation) OpNum() int { return 9 }

func (o *xxx_RebootOperation) OpName() string { return "/IIisServiceControl/v0/Reboot" }

func (o *xxx_RebootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwTimeouMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TimeoutMsecs); err != nil {
			return err
		}
	}
	// dwForceAppsClosed {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ForceAppsClosed); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTimeouMsecs {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TimeoutMsecs); err != nil {
			return err
		}
	}
	// dwForceAppsClosed {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ForceAppsClosed); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RebootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RebootRequest structure represents the Reboot operation request
type RebootRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	TimeoutMsecs uint32         `idl:"name:dwTimeouMsecs" json:"timeout_msecs"`
	// dwForceAppsClosed: Boolean value that specifies whether applications will be forced
	// to close.
	//
	//	+------------------+-------------------------------------------+
	//	|                  |                                           |
	//	|      VALUE       |                  MEANING                  |
	//	|                  |                                           |
	//	+------------------+-------------------------------------------+
	//	+------------------+-------------------------------------------+
	//	| TRUE 0x00000001  | Applications MUST be forced to close.     |
	//	+------------------+-------------------------------------------+
	//	| FALSE 0x00000000 | Applications MUST NOT be forced to close. |
	//	+------------------+-------------------------------------------+
	ForceAppsClosed uint32 `idl:"name:dwForceAppsClosed" json:"force_apps_closed"`
}

func (o *RebootRequest) xxx_ToOp(ctx context.Context, op *xxx_RebootOperation) *xxx_RebootOperation {
	if op == nil {
		op = &xxx_RebootOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TimeoutMsecs = o.TimeoutMsecs
	op.ForceAppsClosed = o.ForceAppsClosed
	return op
}

func (o *RebootRequest) xxx_FromOp(ctx context.Context, op *xxx_RebootOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TimeoutMsecs = op.TimeoutMsecs
	o.ForceAppsClosed = op.ForceAppsClosed
}
func (o *RebootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RebootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RebootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RebootResponse structure represents the Reboot operation response
type RebootResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reboot return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RebootResponse) xxx_ToOp(ctx context.Context, op *xxx_RebootOperation) *xxx_RebootOperation {
	if op == nil {
		op = &xxx_RebootOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RebootResponse) xxx_FromOp(ctx context.Context, op *xxx_RebootOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RebootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RebootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RebootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StatusOperation structure represents the Status operation
type xxx_StatusOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	BufferSize         uint32         `idl:"name:dwBufferSize" json:"buffer_size"`
	Buffer             []byte         `idl:"name:pbBuffer;size_is:(dwBufferSize)" json:"buffer"`
	RequiredBufferSize uint32         `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	ServicesLength     uint32         `idl:"name:pdwNumServices" json:"services_length"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StatusOperation) OpNum() int { return 10 }

func (o *xxx_StatusOperation) OpName() string { return "/IIisServiceControl/v0/Status" }

func (o *xxx_StatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwBufferSize](uchar))
	{
		dimSize1 := uint64(o.BufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredBufferSize); err != nil {
			return err
		}
	}
	// pdwNumServices {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServicesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=dwBufferSize](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// pdwMDRequiredBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredBufferSize); err != nil {
			return err
		}
	}
	// pdwNumServices {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServicesLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StatusRequest structure represents the Status operation request
type StatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// dwBufferSize: Size, in bytes, of the pbBuffer parameter. If this parameter is not
	// greater than the amount of data the server wants to return in pbBuffer, the conditional
	// behavior that follows MUST occur.
	//
	// * The pdwMDRequiredBufferSize parameter MUST be set to the number of bytes needed
	// to contain the data that is to be returned.
	//
	// * The pbBuffer parameter MUST be set to zero.
	//
	// * The method MUST be failed with code 0x8007007A (E_ERROR_INSUFFICIENT_BUFFER).
	BufferSize uint32 `idl:"name:dwBufferSize" json:"buffer_size"`
}

func (o *StatusRequest) xxx_ToOp(ctx context.Context, op *xxx_StatusOperation) *xxx_StatusOperation {
	if op == nil {
		op = &xxx_StatusOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.BufferSize = o.BufferSize
	return op
}

func (o *StatusRequest) xxx_FromOp(ctx context.Context, op *xxx_StatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.BufferSize = op.BufferSize
}
func (o *StatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StatusResponse structure represents the Status operation response
type StatusResponse struct {
	// XXX: dwBufferSize is an implicit input depedency for output parameters
	BufferSize uint32 `idl:"name:dwBufferSize" json:"buffer_size"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbBuffer: An array of unsigned chars that will be filled with information about the
	// status of the Internet services. For more information, see section 2.2.2.
	Buffer []byte `idl:"name:pbBuffer;size_is:(dwBufferSize)" json:"buffer"`
	// pdwMDRequiredBufferSize: On return from this method, if this parameter is not null,
	// this parameter points to a DWORD containing the number of bytes that pbBuffer must
	// be able to contain for the method to return the services status information. This
	// field MAY be used.
	RequiredBufferSize uint32 `idl:"name:pdwMDRequiredBufferSize" json:"required_buffer_size"`
	// pdwNumServices: The number of services for which status is returned.
	ServicesLength uint32 `idl:"name:pdwNumServices" json:"services_length"`
	// Return: The Status return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StatusResponse) xxx_ToOp(ctx context.Context, op *xxx_StatusOperation) *xxx_StatusOperation {
	if op == nil {
		op = &xxx_StatusOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.BufferSize == uint32(0) {
		op.BufferSize = o.BufferSize
	}

	op.That = o.That
	op.Buffer = o.Buffer
	op.RequiredBufferSize = o.RequiredBufferSize
	op.ServicesLength = o.ServicesLength
	op.Return = o.Return
	return op
}

func (o *StatusResponse) xxx_FromOp(ctx context.Context, op *xxx_StatusOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.BufferSize = op.BufferSize

	o.That = op.That
	o.Buffer = op.Buffer
	o.RequiredBufferSize = op.RequiredBufferSize
	o.ServicesLength = op.ServicesLength
	o.Return = op.Return
}
func (o *StatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_KillOperation structure represents the Kill operation
type xxx_KillOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_KillOperation) OpNum() int { return 11 }

func (o *xxx_KillOperation) OpName() string { return "/IIisServiceControl/v0/Kill" }

func (o *xxx_KillOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KillOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KillOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KillOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KillOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_KillOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// KillRequest structure represents the Kill operation request
type KillRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *KillRequest) xxx_ToOp(ctx context.Context, op *xxx_KillOperation) *xxx_KillOperation {
	if op == nil {
		op = &xxx_KillOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *KillRequest) xxx_FromOp(ctx context.Context, op *xxx_KillOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *KillRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *KillRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KillOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// KillResponse structure represents the Kill operation response
type KillResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Kill return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *KillResponse) xxx_ToOp(ctx context.Context, op *xxx_KillOperation) *xxx_KillOperation {
	if op == nil {
		op = &xxx_KillOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *KillResponse) xxx_FromOp(ctx context.Context, op *xxx_KillOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *KillResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *KillResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_KillOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
