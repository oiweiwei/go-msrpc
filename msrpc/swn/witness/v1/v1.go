package witness

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	swn "github.com/oiweiwei/go-msrpc/msrpc/swn"
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
	_ = swn.GoPackage
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "swn"
)

var (
	// Syntax UUID
	WitnessSyntaxUUID = &uuid.UUID{TimeLow: 0xccd8c074, TimeMid: 0xd0e5, TimeHiAndVersion: 0x4a40, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xb4, Node: [6]uint8{0xd0, 0x74, 0xfa, 0xa6, 0xba, 0x28}}
	// Syntax ID
	WitnessSyntaxV1_1 = &dcerpc.SyntaxID{IfUUID: WitnessSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 1}
)

// Witness interface.
type WitnessClient interface {

	// The WitnessrGetInterfaceList method returns information about the interfaces to which
	// witness client connections can be made.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------+
	//	|               RETURN               |                                                              |
	//	|             VALUE/CODE             |                         DESCRIPTION                          |
	//	|                                    |                                                              |
	//	+------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The operation completed successfully.                        |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                            |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                                  |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS     | No more data is available.                                   |
	//	+------------------------------------+--------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | There is not enough storage space to complete the operation. |
	//	+------------------------------------+--------------------------------------------------------------+
	//
	// If there are no entries in the InterfaceList, the server MUST fail the request and
	// return the error code ERROR_NO_MORE_ITEMS.
	//
	// If no entry in the InterfaceList has a State of AVAILABLE, the server MUST wait until
	// at least one entry enters that State, as specified in section 3.1.6.1.
	//
	// For each Interface in the InterfaceList, the server MUST construct a WITNESS_INTERFACE_INFO
	// structure as follows:
	//
	// * The *InterfaceGroupName* field of the WITNESS_INTERFACE_INFO structure MUST be
	// set to *Interface.InterfaceGroupName*.
	//
	// * The *State* field MUST be set to *Interface.State*.
	//
	// * The *Version* field MUST be set to *WitnessServiceVersion*.
	//
	// * If *Interface.IPv4Address* is not empty, the *IPV4* field MUST be set to *Interface.IPv4Address*
	// , and IPv4 flag MUST be set in the *Flags* field.
	//
	// * If *Interface.IPv6Address* is not empty, the *IPV6* field MUST be set to *Interface.IPv6Address*
	// , and IPv6 flag MUST be set in the *Flags* field.
	//
	// * In an implementation-dependent manner, the server MUST determine if the *IPv4Address*
	// or *IPv6Address* match any interface which is hosted on the server and the server
	// is also running this Witness Service instance. If the address is not hosted on the
	// local server, the INTERFACE_WITNESS flag MUST be set in the *Flags* field. Otherwise,
	// the flag MUST NOT be set.
	GetInterfaceList(context.Context, *GetInterfaceListRequest, ...dcerpc.CallOption) (*GetInterfaceListResponse, error)

	// The WitnessrRegister method allows the witness client to register for resource state
	// change notifications of a NetName and IPAddress. The client can subsequently call
	// the WitnessrAsyncNotify method to receive notifications when there is a state change
	// on any of these resources.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER   | The parameter is incorrect.                                            |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000139F ERROR_INVALID_STATE       | The specified resource state is invalid.                               |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000051A ERROR_REVISION_MISMATCH   | The client request contains an invalid Witness protocol version.       |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//
	// If the Version field of the request is not 0x00010001, the server MUST stop processing
	// the request and return the error code ERROR_REVISION_MISMATCH.
	//
	// If NetName, IpAddress or ClientComputerName is NULL, the server MUST fail the request
	// and return the error code ERROR_INVALID_PARAMETER.
	//
	// If the NetName parameter is not equal to ServerGlobalName, the server MUST fail the
	// request and return the error code ERROR_INVALID_PARAMETER.
	Register(context.Context, *RegisterRequest, ...dcerpc.CallOption) (*RegisterResponse, error)

	// The WitnessrUnRegister method allows the client to unregister for notifications from
	// the server. The Witness Service removes its internal state of the registration and
	// no longer notifies the client in the event of any resource state changes.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------+--------------------------------------------+
	//	|             RETURN             |                                            |
	//	|           VALUE/CODE           |                DESCRIPTION                 |
	//	|                                |                                            |
	//	+--------------------------------+--------------------------------------------+
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The operation completed successfully.      |
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                          |
	//	+--------------------------------+--------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND     | The specified CONTEXT_HANDLE is not found. |
	//	+--------------------------------+--------------------------------------------+
	Unregister(context.Context, *UnregisterRequest, ...dcerpc.CallOption) (*UnregisterResponse, error)

	// The WitnessrAsyncNotify method is used by the client to request notification of registered
	// resource changes from the server.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND           | The specified resource name is not found.                              |
	//	+--------------------------------------+------------------------------------------------------------------------+
	AsyncNotify(context.Context, *AsyncNotifyRequest, ...dcerpc.CallOption) (*AsyncNotifyResponse, error)

	// The WitnessrRegisterEx method allows the witness client to register for resource
	// state change notifications of a NetName, ShareName and multiple IPAddresses. The
	// client can subsequently call the WitnessrAsyncNotify method to receive notifications
	// when there is a state change on any of these resources.
	//
	// Return Values: Returns 0x00000000 (ERROR_SUCCESS) on success or a nonzero error code,
	// as specified in [MS-ERREF] section 2.2. The most common error codes are listed in
	// the following table.
	//
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	|                RETURN                |                                                                        |
	//	|              VALUE/CODE              |                              DESCRIPTION                               |
	//	|                                      |                                                                        |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The operation completed successfully.                                  |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | Access is denied.                                                      |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x000005AA ERROR_NO_SYSTEM_RESOURCES | Insufficient system resources exist to complete the requested service. |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER   | The parameter is incorrect.                                            |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000139F ERROR_INVALID_STATE       | The specified resource state is invalid.                               |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//	| 0x0000051A ERROR_REVISION_MISMATCH   | The client request contains an invalid Witness protocol version.       |
	//	+--------------------------------------+------------------------------------------------------------------------+
	//
	// This opnum is applicable only to servers that implement Witness protocol version
	// 2.
	//
	// If the Version field of the request is not 0x00020000, the server MUST stop processing
	// the request and return the error code ERROR_REVISION_MISMATCH.
	//
	// If NetName, IpAddress, or ClientComputerName is NULL, the server MUST fail the request
	// and return the error code ERROR_INVALID_PARAMETER.
	//
	// If the NetName parameter is not equal to ServerGlobalName, the server MUST fail the
	// request and return the error code ERROR_INVALID_PARAMETER.
	//
	// If ShareName is not NULL, the server MUST enumerate the shares by calling NetrShareEnum
	// as specified in [MS-SRVS] section 3.1.4.8. If the enumeration fails or if no shares
	// are returned, the server MUST return the error code ERROR_INVALID_STATE.
	//
	// If none of the shares in the list has shi*_type set to STYPE_CLUSTER_SOFS as specified
	// in [MS-SRVS] section 3.1.4.8, the server MUST ignore ShareName.
	//
	// Otherwise, the server MUST fail the request with the error code ERROR_INVALID_STATE
	// for the following:
	//
	// * *ShareName* does not exist in the enumerated list.
	//
	// * The server MUST search for an *Interface* in *InterfaceList* , where *Interface.IPv4Address*
	// or *Interface.IPv6Address* matches the IpAddress parameter based on its format. If
	// no matching entry is found and *ShareName* has shi*_type set to STYPE_CLUSTER_SOFS,
	// as specified in [MS-SRVS] section 2.2.2.4 ( ../ms-srvs/6069f8c0-c93f-43a0-a5b4-7ed447eb4b84
	// ) , the server MUST fail the request with ERROR_INVALID_STATE. **
	RegisterEx(context.Context, *RegisterExRequest, ...dcerpc.CallOption) (*RegisterExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultWitnessClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWitnessClient) GetInterfaceList(ctx context.Context, in *GetInterfaceListRequest, opts ...dcerpc.CallOption) (*GetInterfaceListResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInterfaceListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWitnessClient) Register(ctx context.Context, in *RegisterRequest, opts ...dcerpc.CallOption) (*RegisterResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWitnessClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...dcerpc.CallOption) (*UnregisterResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnregisterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWitnessClient) AsyncNotify(ctx context.Context, in *AsyncNotifyRequest, opts ...dcerpc.CallOption) (*AsyncNotifyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AsyncNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWitnessClient) RegisterEx(ctx context.Context, in *RegisterExRequest, opts ...dcerpc.CallOption) (*RegisterExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWitnessClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWitnessClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewWitnessClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WitnessClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WitnessSyntaxV1_1))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWitnessClient{cc: cc}, nil
}

// xxx_GetInterfaceListOperation structure represents the WitnessrGetInterfaceList operation
type xxx_GetInterfaceListOperation struct {
	InterfaceList *swn.InterfaceList `idl:"name:InterfaceList" json:"interface_list"`
	Return        uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInterfaceListOperation) OpNum() int { return 0 }

func (o *xxx_GetInterfaceListOperation) OpName() string {
	return "/Witness/v1.1/WitnessrGetInterfaceList"
}

func (o *xxx_GetInterfaceListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetInterfaceListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetInterfaceListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InterfaceList {out} (1:{pointer=ref}*(2))(2:{alias=PWITNESS_INTERFACE_LIST}*(1))(3:{alias=WITNESS_INTERFACE_LIST}(struct))
	{
		if o.InterfaceList != nil {
			_ptr_InterfaceList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InterfaceList != nil {
					if err := o.InterfaceList.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&swn.InterfaceList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceList, _ptr_InterfaceList); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInterfaceListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InterfaceList {out} (1:{pointer=ref}*(2))(2:{alias=PWITNESS_INTERFACE_LIST,pointer=ref}*(1))(3:{alias=WITNESS_INTERFACE_LIST}(struct))
	{
		_ptr_InterfaceList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InterfaceList == nil {
				o.InterfaceList = &swn.InterfaceList{}
			}
			if err := o.InterfaceList.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_InterfaceList := func(ptr interface{}) { o.InterfaceList = *ptr.(**swn.InterfaceList) }
		if err := w.ReadPointer(&o.InterfaceList, _s_InterfaceList, _ptr_InterfaceList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInterfaceListRequest structure represents the WitnessrGetInterfaceList operation request
type GetInterfaceListRequest struct {
}

func (o *GetInterfaceListRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInterfaceListOperation) *xxx_GetInterfaceListOperation {
	if op == nil {
		op = &xxx_GetInterfaceListOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetInterfaceListRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInterfaceListOperation) {
	if o == nil {
		return
	}
}
func (o *GetInterfaceListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInterfaceListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInterfaceListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInterfaceListResponse structure represents the WitnessrGetInterfaceList operation response
type GetInterfaceListResponse struct {
	// InterfaceList:  A pointer to a PWITNESS_INTERFACE_LIST, as specified in section
	// 2.2.2.6.
	InterfaceList *swn.InterfaceList `idl:"name:InterfaceList" json:"interface_list"`
	// Return: The WitnessrGetInterfaceList return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInterfaceListResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInterfaceListOperation) *xxx_GetInterfaceListOperation {
	if op == nil {
		op = &xxx_GetInterfaceListOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceList = o.InterfaceList
	op.Return = o.Return
	return op
}

func (o *GetInterfaceListResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInterfaceListOperation) {
	if o == nil {
		return
	}
	o.InterfaceList = op.InterfaceList
	o.Return = op.Return
}
func (o *GetInterfaceListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInterfaceListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInterfaceListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterOperation structure represents the WitnessrRegister operation
type xxx_RegisterOperation struct {
	Context            *dcetypes.ContextHandle `idl:"name:ppContext" json:"context"`
	Version            uint32                  `idl:"name:Version" json:"version"`
	NetName            string                  `idl:"name:NetName;string;pointer:unique" json:"net_name"`
	IPAddress          string                  `idl:"name:IpAddress;string;pointer:unique" json:"ip_address"`
	ClientComputerName string                  `idl:"name:ClientComputerName;string;pointer:unique" json:"client_computer_name"`
	Return             uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterOperation) OpNum() int { return 1 }

func (o *xxx_RegisterOperation) OpName() string { return "/Witness/v1.1/WitnessrRegister" }

func (o *xxx_RegisterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Version {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// NetName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NetName != "" {
			_ptr_NetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NetName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NetName, _ptr_NetName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// IpAddress {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.IPAddress != "" {
			_ptr_IpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.IPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.IPAddress, _ptr_IpAddress); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClientComputerName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ClientComputerName != "" {
			_ptr_ClientComputerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ClientComputerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientComputerName, _ptr_ClientComputerName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Version {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// NetName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_NetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NetName); err != nil {
				return err
			}
			return nil
		})
		_s_NetName := func(ptr interface{}) { o.NetName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NetName, _s_NetName, _ptr_NetName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// IpAddress {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_IpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.IPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_IpAddress := func(ptr interface{}) { o.IPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.IPAddress, _s_IpAddress, _ptr_IpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientComputerName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ClientComputerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComputerName); err != nil {
				return err
			}
			return nil
		})
		_s_ClientComputerName := func(ptr interface{}) { o.ClientComputerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ClientComputerName, _s_ClientComputerName, _ptr_ClientComputerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppContext {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppContext {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &dcetypes.ContextHandle{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterRequest structure represents the WitnessrRegister operation request
type RegisterRequest struct {
	// Version:  The version of the Witness protocol currently in use by the client.
	Version uint32 `idl:"name:Version" json:"version"`
	// NetName:  A pointer to a null-terminated string that specifies the name of the resource
	// for which the client requires notifications.
	NetName string `idl:"name:NetName;string;pointer:unique" json:"net_name"`
	// IpAddress:  A pointer to a null-terminated string that specifies the IP address
	// to which the client application connection is established.
	IPAddress string `idl:"name:IpAddress;string;pointer:unique" json:"ip_address"`
	// ClientComputerName:  A pointer to a null-terminated string that is used to identify
	// the Witness client.
	ClientComputerName string `idl:"name:ClientComputerName;string;pointer:unique" json:"client_computer_name"`
}

func (o *RegisterRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterOperation) *xxx_RegisterOperation {
	if op == nil {
		op = &xxx_RegisterOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.NetName = o.NetName
	op.IPAddress = o.IPAddress
	op.ClientComputerName = o.ClientComputerName
	return op
}

func (o *RegisterRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.NetName = op.NetName
	o.IPAddress = op.IPAddress
	o.ClientComputerName = op.ClientComputerName
}
func (o *RegisterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterResponse structure represents the WitnessrRegister operation response
type RegisterResponse struct {
	// ppContext: A context handle of type PPCONTEXT_HANDLE, as specified in section 2.2.1.2,
	// that identifies the client on the server.
	Context *dcetypes.ContextHandle `idl:"name:ppContext" json:"context"`
	// Return: The WitnessrRegister return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RegisterResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterOperation) *xxx_RegisterOperation {
	if op == nil {
		op = &xxx_RegisterOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *RegisterResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *RegisterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnregisterOperation structure represents the WitnessrUnRegister operation
type xxx_UnregisterOperation struct {
	Context *dcetypes.ContextHandle `idl:"name:pContext" json:"context"`
	Return  uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_UnregisterOperation) OpNum() int { return 2 }

func (o *xxx_UnregisterOperation) OpName() string { return "/Witness/v1.1/WitnessrUnRegister" }

func (o *xxx_UnregisterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pContext {in} (1:{alias=PCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UnregisterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pContext {in} (1:{alias=PCONTEXT_HANDLE,pointer=ref}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &dcetypes.ContextHandle{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnregisterRequest structure represents the WitnessrUnRegister operation request
type UnregisterRequest struct {
	// pContext: A context handle of type PCONTEXT_HANDLE, specified in section 2.2.1.1,
	// that identifies the client on the server.
	Context *dcetypes.ContextHandle `idl:"name:pContext" json:"context"`
}

func (o *UnregisterRequest) xxx_ToOp(ctx context.Context, op *xxx_UnregisterOperation) *xxx_UnregisterOperation {
	if op == nil {
		op = &xxx_UnregisterOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *UnregisterRequest) xxx_FromOp(ctx context.Context, op *xxx_UnregisterOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *UnregisterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnregisterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnregisterResponse structure represents the WitnessrUnRegister operation response
type UnregisterResponse struct {
	// Return: The WitnessrUnRegister return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UnregisterResponse) xxx_ToOp(ctx context.Context, op *xxx_UnregisterOperation) *xxx_UnregisterOperation {
	if op == nil {
		op = &xxx_UnregisterOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *UnregisterResponse) xxx_FromOp(ctx context.Context, op *xxx_UnregisterOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UnregisterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnregisterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AsyncNotifyOperation structure represents the WitnessrAsyncNotify operation
type xxx_AsyncNotifyOperation struct {
	Context  *swn.Shared              `idl:"name:pContext" json:"context"`
	Response *swn.ResponseAsyncNotify `idl:"name:pResp" json:"response"`
	Return   uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_AsyncNotifyOperation) OpNum() int { return 3 }

func (o *xxx_AsyncNotifyOperation) OpName() string { return "/Witness/v1.1/WitnessrAsyncNotify" }

func (o *xxx_AsyncNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SHARED, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&swn.Shared{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AsyncNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_SHARED, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &swn.Shared{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResp {out} (1:{pointer=ref}*(2))(2:{alias=PRESP_ASYNC_NOTIFY}*(1))(3:{alias=RESP_ASYNC_NOTIFY}(struct))
	{
		if o.Response != nil {
			_ptr_pResp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Response != nil {
					if err := o.Response.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&swn.ResponseAsyncNotify{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Response, _ptr_pResp); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AsyncNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResp {out} (1:{pointer=ref}*(2))(2:{alias=PRESP_ASYNC_NOTIFY,pointer=ref}*(1))(3:{alias=RESP_ASYNC_NOTIFY}(struct))
	{
		_ptr_pResp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Response == nil {
				o.Response = &swn.ResponseAsyncNotify{}
			}
			if err := o.Response.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pResp := func(ptr interface{}) { o.Response = *ptr.(**swn.ResponseAsyncNotify) }
		if err := w.ReadPointer(&o.Response, _s_pResp, _ptr_pResp); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AsyncNotifyRequest structure represents the WitnessrAsyncNotify operation request
type AsyncNotifyRequest struct {
	// pContext: A context handle of type PCONTEXT_HANDLE_SHARED, as specified in section
	// 2.2.1.3, that identifies the client on the server.
	Context *swn.Shared `idl:"name:pContext" json:"context"`
}

func (o *AsyncNotifyRequest) xxx_ToOp(ctx context.Context, op *xxx_AsyncNotifyOperation) *xxx_AsyncNotifyOperation {
	if op == nil {
		op = &xxx_AsyncNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	return op
}

func (o *AsyncNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_AsyncNotifyOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *AsyncNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AsyncNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AsyncNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AsyncNotifyResponse structure represents the WitnessrAsyncNotify operation response
type AsyncNotifyResponse struct {
	// pResp:  A pointer to a PRESP_ASYNC_NOTIFY structure, as specified in section 2.2.2.4.
	Response *swn.ResponseAsyncNotify `idl:"name:pResp" json:"response"`
	// Return: The WitnessrAsyncNotify return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AsyncNotifyResponse) xxx_ToOp(ctx context.Context, op *xxx_AsyncNotifyOperation) *xxx_AsyncNotifyOperation {
	if op == nil {
		op = &xxx_AsyncNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.Response = o.Response
	op.Return = o.Return
	return op
}

func (o *AsyncNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_AsyncNotifyOperation) {
	if o == nil {
		return
	}
	o.Response = op.Response
	o.Return = op.Return
}
func (o *AsyncNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AsyncNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AsyncNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterExOperation structure represents the WitnessrRegisterEx operation
type xxx_RegisterExOperation struct {
	Context            *dcetypes.ContextHandle `idl:"name:ppContext" json:"context"`
	Version            uint32                  `idl:"name:Version" json:"version"`
	NetName            string                  `idl:"name:NetName;string;pointer:unique" json:"net_name"`
	ShareName          string                  `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	IPAddress          string                  `idl:"name:IpAddress;string;pointer:unique" json:"ip_address"`
	ClientComputerName string                  `idl:"name:ClientComputerName;string;pointer:unique" json:"client_computer_name"`
	Flags              uint32                  `idl:"name:Flags" json:"flags"`
	KeepAliveTimeout   uint32                  `idl:"name:KeepAliveTimeout" json:"keep_alive_timeout"`
	Return             uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RegisterExOperation) OpNum() int { return 4 }

func (o *xxx_RegisterExOperation) OpName() string { return "/Witness/v1.1/WitnessrRegisterEx" }

func (o *xxx_RegisterExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Version {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// NetName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NetName != "" {
			_ptr_NetName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NetName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NetName, _ptr_NetName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// IpAddress {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.IPAddress != "" {
			_ptr_IpAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.IPAddress); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.IPAddress, _ptr_IpAddress); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClientComputerName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ClientComputerName != "" {
			_ptr_ClientComputerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ClientComputerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientComputerName, _ptr_ClientComputerName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// KeepAliveTimeout {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.KeepAliveTimeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Version {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// NetName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_NetName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NetName); err != nil {
				return err
			}
			return nil
		})
		_s_NetName := func(ptr interface{}) { o.NetName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NetName, _s_NetName, _ptr_NetName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// IpAddress {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_IpAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.IPAddress); err != nil {
				return err
			}
			return nil
		})
		_s_IpAddress := func(ptr interface{}) { o.IPAddress = *ptr.(*string) }
		if err := w.ReadPointer(&o.IPAddress, _s_IpAddress, _ptr_IpAddress); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientComputerName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ClientComputerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ClientComputerName); err != nil {
				return err
			}
			return nil
		})
		_s_ClientComputerName := func(ptr interface{}) { o.ClientComputerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ClientComputerName, _s_ClientComputerName, _ptr_ClientComputerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// KeepAliveTimeout {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.KeepAliveTimeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppContext {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcetypes.ContextHandle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppContext {out} (1:{pointer=ref, alias=PPCONTEXT_HANDLE}*(1))(2:{alias=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &dcetypes.ContextHandle{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RegisterExRequest structure represents the WitnessrRegisterEx operation request
type RegisterExRequest struct {
	// Version:  The version of the Witness protocol currently in use by the client.
	Version uint32 `idl:"name:Version" json:"version"`
	// NetName:  A pointer to a null-terminated string that specifies the name of the resource
	// for which the client requires notifications.
	NetName string `idl:"name:NetName;string;pointer:unique" json:"net_name"`
	// ShareName:  A pointer to a null-terminated string that specifies the name of the
	// share resource for which the client requires notifications.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// IpAddress:  A pointer to a null-terminated string that specifies the IP address
	// to which the client application connection is established.
	IPAddress string `idl:"name:IpAddress;string;pointer:unique" json:"ip_address"`
	// ClientComputerName:  A pointer to a null-terminated string that is used to identify
	// the Witness client.
	ClientComputerName string `idl:"name:ClientComputerName;string;pointer:unique" json:"client_computer_name"`
	// Flags:  The type of Witness registration. This field MUST be set to one of the following
	// values:
	//
	//	+---------------------------------------------+--------------------------------------------------------------------------------+
	//	|                                             |                                                                                |
	//	|                    VALUE                    |                                    MEANING                                     |
	//	|                                             |                                                                                |
	//	+---------------------------------------------+--------------------------------------------------------------------------------+
	//	+---------------------------------------------+--------------------------------------------------------------------------------+
	//	| WITNESS_REGISTER_NONE 0x00000000            | If set, the client requests notifications only for the registered IP address.  |
	//	+---------------------------------------------+--------------------------------------------------------------------------------+
	//	| WITNESS_REGISTER_IP_NOTIFICATION 0x00000001 | If set, the client requests notifications of any eligible server IP addresses. |
	//	+---------------------------------------------+--------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// KeepAliveTimeout: The maximum number of seconds for any notification response from
	// the server.
	KeepAliveTimeout uint32 `idl:"name:KeepAliveTimeout" json:"keep_alive_timeout"`
}

func (o *RegisterExRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterExOperation) *xxx_RegisterExOperation {
	if op == nil {
		op = &xxx_RegisterExOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.NetName = o.NetName
	op.ShareName = o.ShareName
	op.IPAddress = o.IPAddress
	op.ClientComputerName = o.ClientComputerName
	op.Flags = o.Flags
	op.KeepAliveTimeout = o.KeepAliveTimeout
	return op
}

func (o *RegisterExRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterExOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.NetName = op.NetName
	o.ShareName = op.ShareName
	o.IPAddress = op.IPAddress
	o.ClientComputerName = op.ClientComputerName
	o.Flags = op.Flags
	o.KeepAliveTimeout = op.KeepAliveTimeout
}
func (o *RegisterExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RegisterExResponse structure represents the WitnessrRegisterEx operation response
type RegisterExResponse struct {
	// ppContext:  A context handle of type PPCONTEXT_HANDLE, as specified in section 2.2.1.2,
	// that identifies the client on the server.
	Context *dcetypes.ContextHandle `idl:"name:ppContext" json:"context"`
	// Return: The WitnessrRegisterEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RegisterExResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterExOperation) *xxx_RegisterExOperation {
	if op == nil {
		op = &xxx_RegisterExOperation{}
	}
	if o == nil {
		return op
	}
	op.Context = o.Context
	op.Return = o.Return
	return op
}

func (o *RegisterExResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterExOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *RegisterExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
