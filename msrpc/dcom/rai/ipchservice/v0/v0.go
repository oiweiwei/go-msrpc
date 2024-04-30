package ipchservice

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	rai "github.com/oiweiwei/go-msrpc/msrpc/dcom/rai"
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
	_ = oaut.GoPackage
	_ = rai.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rai"
)

var (
	// IPCHService interface identifier 833e4200-aff7-4ac3-aac2-9f24c1457bce
	PCHServiceIID = &dcom.IID{Data1: 0x833e4200, Data2: 0xaff7, Data3: 0x4ac3, Data4: []byte{0xaa, 0xc2, 0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}
	// Syntax UUID
	PCHServiceSyntaxUUID = &uuid.UUID{TimeLow: 0x833e4200, TimeMid: 0xaff7, TimeHiAndVersion: 0x4ac3, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc2, Node: [6]uint8{0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}
	// Syntax ID
	PCHServiceSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PCHServiceSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IPCHService interface.
type PCHServiceClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Opnum7NotUsedByProtocol operation.
	// Opnum7NotUsedByProtocol

	// Opnum8NotUsedByProtocol operation.
	// Opnum8NotUsedByProtocol

	// Opnum9NotUsedByProtocol operation.
	// Opnum9NotUsedByProtocol

	// Opnum10NotUsedByProtocol operation.
	// Opnum10NotUsedByProtocol

	// Opnum11NotUsedByProtocol operation.
	// Opnum11NotUsedByProtocol

	// Opnum12NotUsedByProtocol operation.
	// Opnum12NotUsedByProtocol

	// Opnum13NotUsedByProtocol operation.
	// Opnum13NotUsedByProtocol

	// Opnum14NotUsedByProtocol operation.
	// Opnum14NotUsedByProtocol

	// Opnum15NotUsedByProtocol operation.
	// Opnum15NotUsedByProtocol

	// Opnum16NotUsedByProtocol operation.
	// Opnum16NotUsedByProtocol

	// Opnum17NotUsedByProtocol operation.
	// Opnum17NotUsedByProtocol

	// Opnum18NotUsedByProtocol operation.
	// Opnum18NotUsedByProtocol

	// The RemoteConnectionParms method gets the Remote Assistance connection parameters
	// for a specific UserName, DomainName, and SessionID triple.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure. If the UserName and DomainName are valid BSTRs, the return
	// code is one listed in the following table. If the UserName and DomainName are invalid
	// BSTRs, the HRESULT value returned is the corresponding HRESULT to the system error
	// code ERROR_NONE_MAPPED.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                  | General access denied error. <8>                                                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY                   | Out of memory.                                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704EC ERROR_ACCESS_DISABLED_BY_POLICY | The program cannot be opened because of a software restriction policy. For more  |
	//	|                                            | information, contact the system administrator.                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	RemoteConnectionParameters(context.Context, *RemoteConnectionParametersRequest, ...dcerpc.CallOption) (*RemoteConnectionParametersResponse, error)

	// The RemoteUserSessionInfo method returns the collection of the terminal services
	// sessions on the remote novice machine. All the terminal services session information
	// is returned as a standard IPCHCollection interface. The members of this collection
	// are objects of type ISAFSession. ISAFSession includes the DomainName, SessionID,
	// SessionState, and UserName for each session.
	//
	// Return Values: A signed 32-bit value indicating return status. This method MUST return
	// zero to indicate success, or an HRESULT error value (as specified in [MS-ERREF] section
	// 2.1.1) to indicate failure.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | The call was successful.                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070005 E_ACCESSDENIED                  | General access denied error. <9>                                                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8007000E E_OUTOFMEMORY                   | Out of memory.                                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x800704EC ERROR_ACCESS_DISABLED_BY_POLICY | The program cannot be opened because of a software restriction policy. For more  |
	//	|                                            | information, contact the system administrator.                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol [MS-RPCE].
	RemoteUserSessionInfo(context.Context, *RemoteUserSessionInfoRequest, ...dcerpc.CallOption) (*RemoteUserSessionInfoResponse, error)

	// Opnum21NotUsedByProtocol operation.
	// Opnum21NotUsedByProtocol

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PCHServiceClient
}

type xxx_DefaultPCHServiceClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPCHServiceClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultPCHServiceClient) RemoteConnectionParameters(ctx context.Context, in *RemoteConnectionParametersRequest, opts ...dcerpc.CallOption) (*RemoteConnectionParametersResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &RemoteConnectionParametersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPCHServiceClient) RemoteUserSessionInfo(ctx context.Context, in *RemoteUserSessionInfoRequest, opts ...dcerpc.CallOption) (*RemoteUserSessionInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &RemoteUserSessionInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPCHServiceClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPCHServiceClient) IPID(ctx context.Context, ipid *dcom.IPID) PCHServiceClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPCHServiceClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewPCHServiceClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PCHServiceClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PCHServiceSyntaxV0_0))...)
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
	return &xxx_DefaultPCHServiceClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_RemoteConnectionParametersOperation structure represents the RemoteConnectionParms operation
type xxx_RemoteConnectionParametersOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserName         *oaut.String   `idl:"name:bstrUserName" json:"user_name"`
	DomainName       *oaut.String   `idl:"name:bstrDomainName" json:"domain_name"`
	SessionID        int32          `idl:"name:lSessionID" json:"session_id"`
	UserHelpBlob     *oaut.String   `idl:"name:bstrUserHelpBlob" json:"user_help_blob"`
	ConnectionString *oaut.String   `idl:"name:pbstrConnectionString" json:"connection_string"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteConnectionParametersOperation) OpNum() int { return 19 }

func (o *xxx_RemoteConnectionParametersOperation) OpName() string {
	return "/IPCHService/v0/RemoteConnectionParms"
}

func (o *xxx_RemoteConnectionParametersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteConnectionParametersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrUserName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserName != nil {
			_ptr_bstrUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserName != nil {
					if err := o.UserName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_bstrUserName); err != nil {
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
	// bstrDomainName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DomainName != nil {
			_ptr_bstrDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainName != nil {
					if err := o.DomainName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_bstrDomainName); err != nil {
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
	// lSessionID {in} (1:(int32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// bstrUserHelpBlob {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserHelpBlob != nil {
			_ptr_bstrUserHelpBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserHelpBlob != nil {
					if err := o.UserHelpBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserHelpBlob, _ptr_bstrUserHelpBlob); err != nil {
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

func (o *xxx_RemoteConnectionParametersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrUserName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserName == nil {
				o.UserName = &oaut.String{}
			}
			if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrUserName := func(ptr interface{}) { o.UserName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserName, _s_bstrUserName, _ptr_bstrUserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrDomainName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DomainName == nil {
				o.DomainName = &oaut.String{}
			}
			if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDomainName := func(ptr interface{}) { o.DomainName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DomainName, _s_bstrDomainName, _ptr_bstrDomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lSessionID {in} (1:(int32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// bstrUserHelpBlob {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrUserHelpBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserHelpBlob == nil {
				o.UserHelpBlob = &oaut.String{}
			}
			if err := o.UserHelpBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrUserHelpBlob := func(ptr interface{}) { o.UserHelpBlob = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserHelpBlob, _s_bstrUserHelpBlob, _ptr_bstrUserHelpBlob); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteConnectionParametersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteConnectionParametersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrConnectionString {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConnectionString != nil {
			_ptr_pbstrConnectionString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConnectionString != nil {
					if err := o.ConnectionString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConnectionString, _ptr_pbstrConnectionString); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteConnectionParametersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrConnectionString {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConnectionString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConnectionString == nil {
				o.ConnectionString = &oaut.String{}
			}
			if err := o.ConnectionString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConnectionString := func(ptr interface{}) { o.ConnectionString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConnectionString, _s_pbstrConnectionString, _ptr_pbstrConnectionString); err != nil {
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

// RemoteConnectionParametersRequest structure represents the RemoteConnectionParms operation request
type RemoteConnectionParametersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bstrUserName: The UserName part of the DomainName\UserName string corresponding to
	// the terminal services session for which the client is requesting a Remote Assistance
	// Connection String.
	UserName *oaut.String `idl:"name:bstrUserName" json:"user_name"`
	// bstrDomainName: The DomainName part of the DomainName\UserName string corresponding
	// to the terminal services session for which the client is requesting a Remote Assistance
	// Connection String.
	DomainName *oaut.String `idl:"name:bstrDomainName" json:"domain_name"`
	// lSessionID: Identifier of the terminal services session for which the client is requesting
	// a Remote Assistance Connection String.
	SessionID int32 `idl:"name:lSessionID" json:"session_id"`
	// bstrUserHelpBlob: A semicolon-delimited string that contains the domain and user
	// names of the expert requesting a Remote Assistance Connection String. The format
	// of the string is as follows.
	//
	// The following is an example.
	UserHelpBlob *oaut.String `idl:"name:bstrUserHelpBlob" json:"user_help_blob"`
}

func (o *RemoteConnectionParametersRequest) xxx_ToOp(ctx context.Context) *xxx_RemoteConnectionParametersOperation {
	if o == nil {
		return &xxx_RemoteConnectionParametersOperation{}
	}
	return &xxx_RemoteConnectionParametersOperation{
		This:         o.This,
		UserName:     o.UserName,
		DomainName:   o.DomainName,
		SessionID:    o.SessionID,
		UserHelpBlob: o.UserHelpBlob,
	}
}

func (o *RemoteConnectionParametersRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteConnectionParametersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.SessionID = op.SessionID
	o.UserHelpBlob = op.UserHelpBlob
}
func (o *RemoteConnectionParametersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteConnectionParametersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteConnectionParametersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteConnectionParametersResponse structure represents the RemoteConnectionParms operation response
type RemoteConnectionParametersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrConnectionString: A pointer to a Remote Assistance Connection String for the
	// requested session.
	ConnectionString *oaut.String `idl:"name:pbstrConnectionString" json:"connection_string"`
	// Return: The RemoteConnectionParms return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteConnectionParametersResponse) xxx_ToOp(ctx context.Context) *xxx_RemoteConnectionParametersOperation {
	if o == nil {
		return &xxx_RemoteConnectionParametersOperation{}
	}
	return &xxx_RemoteConnectionParametersOperation{
		That:             o.That,
		ConnectionString: o.ConnectionString,
		Return:           o.Return,
	}
}

func (o *RemoteConnectionParametersResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteConnectionParametersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConnectionString = op.ConnectionString
	o.Return = op.Return
}
func (o *RemoteConnectionParametersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteConnectionParametersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteConnectionParametersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteUserSessionInfoOperation structure represents the RemoteUserSessionInfo operation
type xxx_RemoteUserSessionInfoOperation struct {
	This   *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Value  *rai.PCHCollection `idl:"name:pVal" json:"value"`
	Return int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteUserSessionInfoOperation) OpNum() int { return 20 }

func (o *xxx_RemoteUserSessionInfoOperation) OpName() string {
	return "/IPCHService/v0/RemoteUserSessionInfo"
}

func (o *xxx_RemoteUserSessionInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteUserSessionInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoteUserSessionInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RemoteUserSessionInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteUserSessionInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVal {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IPCHCollection}(interface))
	{
		if o.Value != nil {
			_ptr_pVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rai.PCHCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pVal); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteUserSessionInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVal {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IPCHCollection}(interface))
	{
		_ptr_pVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &rai.PCHCollection{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVal := func(ptr interface{}) { o.Value = *ptr.(**rai.PCHCollection) }
		if err := w.ReadPointer(&o.Value, _s_pVal, _ptr_pVal); err != nil {
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

// RemoteUserSessionInfoRequest structure represents the RemoteUserSessionInfo operation request
type RemoteUserSessionInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RemoteUserSessionInfoRequest) xxx_ToOp(ctx context.Context) *xxx_RemoteUserSessionInfoOperation {
	if o == nil {
		return &xxx_RemoteUserSessionInfoOperation{}
	}
	return &xxx_RemoteUserSessionInfoOperation{
		This: o.This,
	}
}

func (o *RemoteUserSessionInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteUserSessionInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RemoteUserSessionInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteUserSessionInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteUserSessionInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteUserSessionInfoResponse structure represents the RemoteUserSessionInfo operation response
type RemoteUserSessionInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVal: A pointer to an IPCHCollection interface containing terminal services sessions
	// information on the server.
	Value *rai.PCHCollection `idl:"name:pVal" json:"value"`
	// Return: The RemoteUserSessionInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteUserSessionInfoResponse) xxx_ToOp(ctx context.Context) *xxx_RemoteUserSessionInfoOperation {
	if o == nil {
		return &xxx_RemoteUserSessionInfoOperation{}
	}
	return &xxx_RemoteUserSessionInfoOperation{
		That:   o.That,
		Value:  o.Value,
		Return: o.Return,
	}
}

func (o *RemoteUserSessionInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteUserSessionInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *RemoteUserSessionInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteUserSessionInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteUserSessionInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
