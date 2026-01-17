package ialternatelaunch

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// IAlternateLaunch interface identifier 7f43b400-1a0e-4d57-bbc9-6b0c65f7a889
	AlternateLaunchIID = &dcom.IID{Data1: 0x7f43b400, Data2: 0x1a0e, Data3: 0x4d57, Data4: []byte{0xbb, 0xc9, 0x6b, 0x0c, 0x65, 0xf7, 0xa8, 0x89}}
	// Syntax UUID
	AlternateLaunchSyntaxUUID = &uuid.UUID{TimeLow: 0x7f43b400, TimeMid: 0x1a0e, TimeHiAndVersion: 0x4d57, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xc9, Node: [6]uint8{0x6b, 0xc, 0x65, 0xf7, 0xa8, 0x89}}
	// Syntax ID
	AlternateLaunchSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AlternateLaunchSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAlternateLaunch interface.
type AlternateLaunchClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to create an alternate launch configuration, as
	// specified in section 3.1.1.4, for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF], section 2.1, on failure. All failure results
	// MUST be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	CreateConfiguration(context.Context, *CreateConfigurationRequest, ...dcerpc.CallOption) (*CreateConfigurationResponse, error)

	// This method is called by a client to delete an alternate launch configuration (see
	// section 3.1.1.4) for a conglomeration.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	DeleteConfiguration(context.Context, *DeleteConfigurationRequest, ...dcerpc.CallOption) (*DeleteConfigurationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AlternateLaunchClient
}

type xxx_DefaultAlternateLaunchClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAlternateLaunchClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAlternateLaunchClient) CreateConfiguration(ctx context.Context, in *CreateConfigurationRequest, opts ...dcerpc.CallOption) (*CreateConfigurationResponse, error) {
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
	out := &CreateConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAlternateLaunchClient) DeleteConfiguration(ctx context.Context, in *DeleteConfigurationRequest, opts ...dcerpc.CallOption) (*DeleteConfigurationResponse, error) {
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
	out := &DeleteConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAlternateLaunchClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAlternateLaunchClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAlternateLaunchClient) IPID(ctx context.Context, ipid *dcom.IPID) AlternateLaunchClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAlternateLaunchClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAlternateLaunchClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AlternateLaunchClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AlternateLaunchSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAlternateLaunchClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateConfigurationOperation structure represents the CreateConfiguration operation
type xxx_CreateConfigurationOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID  *dtyp.GUID     `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	ConfigurationName *oaut.String   `idl:"name:bstrConfigurationName" json:"configuration_name"`
	StartType         uint32         `idl:"name:dwStartType" json:"start_type"`
	ErrorControl      uint32         `idl:"name:dwErrorControl" json:"error_control"`
	Dependencies      *oaut.String   `idl:"name:bstrDependencies" json:"dependencies"`
	RunAs             *oaut.String   `idl:"name:bstrRunAs" json:"run_as"`
	Password          *oaut.String   `idl:"name:bstrPassword" json:"password"`
	DesktopOK         int16          `idl:"name:bDesktopOk" json:"desktop_ok"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateConfigurationOperation) OpNum() int { return 3 }

func (o *xxx_CreateConfigurationOperation) OpName() string {
	return "/IAlternateLaunch/v0/CreateConfiguration"
}

func (o *xxx_CreateConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// bstrConfigurationName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigurationName != nil {
			_ptr_bstrConfigurationName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationName != nil {
					if err := o.ConfigurationName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationName, _ptr_bstrConfigurationName); err != nil {
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
	// dwStartType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StartType); err != nil {
			return err
		}
	}
	// dwErrorControl {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ErrorControl); err != nil {
			return err
		}
	}
	// bstrDependencies {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Dependencies != nil {
			_ptr_bstrDependencies := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Dependencies != nil {
					if err := o.Dependencies.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Dependencies, _ptr_bstrDependencies); err != nil {
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
	// bstrRunAs {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.RunAs != nil {
			_ptr_bstrRunAs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RunAs != nil {
					if err := o.RunAs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RunAs, _ptr_bstrRunAs); err != nil {
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
	// bstrPassword {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Password != nil {
			_ptr_bstrPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_bstrPassword); err != nil {
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
	// bDesktopOk {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.DesktopOK); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// bstrConfigurationName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigurationName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationName == nil {
				o.ConfigurationName = &oaut.String{}
			}
			if err := o.ConfigurationName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigurationName := func(ptr interface{}) { o.ConfigurationName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigurationName, _s_bstrConfigurationName, _ptr_bstrConfigurationName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwStartType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StartType); err != nil {
			return err
		}
	}
	// dwErrorControl {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ErrorControl); err != nil {
			return err
		}
	}
	// bstrDependencies {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDependencies := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Dependencies == nil {
				o.Dependencies = &oaut.String{}
			}
			if err := o.Dependencies.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDependencies := func(ptr interface{}) { o.Dependencies = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Dependencies, _s_bstrDependencies, _ptr_bstrDependencies); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrRunAs {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrRunAs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RunAs == nil {
				o.RunAs = &oaut.String{}
			}
			if err := o.RunAs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrRunAs := func(ptr interface{}) { o.RunAs = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.RunAs, _s_bstrRunAs, _ptr_bstrRunAs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPassword {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &oaut.String{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPassword := func(ptr interface{}) { o.Password = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Password, _s_bstrPassword, _ptr_bstrPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bDesktopOk {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.DesktopOK); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateConfigurationRequest structure represents the CreateConfiguration operation request
type CreateConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationIdentifier: The conglomeration identifier of a conglomeration on the
	// server.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	// bstrConfigurationName: A value to be used as the AlternateLaunchName property of
	// the alternate launch configuration.
	ConfigurationName *oaut.String `idl:"name:bstrConfigurationName" json:"configuration_name"`
	// dwStartType: A value to be used as the StartType property of the alternate launch
	// configuration.
	StartType uint32 `idl:"name:dwStartType" json:"start_type"`
	// dwErrorControl: A value to be used as the ErrorControl property of the alternate
	// launch configuration.
	ErrorControl uint32 `idl:"name:dwErrorControl" json:"error_control"`
	// bstrDependencies: A value to be used as the Dependencies property of the alternate
	// launch configuration.
	Dependencies *oaut.String `idl:"name:bstrDependencies" json:"dependencies"`
	// bstrRunAs: A value to be used as the AlternateLaunchRunAs property of the alternate
	// launch configuration.
	RunAs *oaut.String `idl:"name:bstrRunAs" json:"run_as"`
	// bstrPassword: A value to be used as the AlternateLaunchPassword property of the alternate
	// launch configuration.
	Password *oaut.String `idl:"name:bstrPassword" json:"password"`
	// bDesktopOk: A value to be used as the DesktopOk property of the alternate launch
	// configuration.
	DesktopOK int16 `idl:"name:bDesktopOk" json:"desktop_ok"`
}

func (o *CreateConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateConfigurationOperation) *xxx_CreateConfigurationOperation {
	if op == nil {
		op = &xxx_CreateConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConglomerationID = o.ConglomerationID
	op.ConfigurationName = o.ConfigurationName
	op.StartType = o.StartType
	op.ErrorControl = o.ErrorControl
	op.Dependencies = o.Dependencies
	op.RunAs = o.RunAs
	op.Password = o.Password
	op.DesktopOK = o.DesktopOK
	return op
}

func (o *CreateConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
	o.ConfigurationName = op.ConfigurationName
	o.StartType = op.StartType
	o.ErrorControl = op.ErrorControl
	o.Dependencies = op.Dependencies
	o.RunAs = op.RunAs
	o.Password = op.Password
	o.DesktopOK = op.DesktopOK
}
func (o *CreateConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateConfigurationResponse structure represents the CreateConfiguration operation response
type CreateConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateConfigurationOperation) *xxx_CreateConfigurationOperation {
	if op == nil {
		op = &xxx_CreateConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteConfigurationOperation structure represents the DeleteConfiguration operation
type xxx_DeleteConfigurationOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConglomerationID *dtyp.GUID     `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteConfigurationOperation) OpNum() int { return 4 }

func (o *xxx_DeleteConfigurationOperation) OpName() string {
	return "/IAlternateLaunch/v0/DeleteConfiguration"
}

func (o *xxx_DeleteConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID != nil {
			if err := o.ConglomerationID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ConglomerationIdentifier {in} (1:{alias=GUID}(struct))
	{
		if o.ConglomerationID == nil {
			o.ConglomerationID = &dtyp.GUID{}
		}
		if err := o.ConglomerationID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteConfigurationRequest structure represents the DeleteConfiguration operation request
type DeleteConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ConglomerationIdentifier: The conglomeration identifier of a conglomeration on the
	// server.
	ConglomerationID *dtyp.GUID `idl:"name:ConglomerationIdentifier" json:"conglomeration_id"`
}

func (o *DeleteConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteConfigurationOperation) *xxx_DeleteConfigurationOperation {
	if op == nil {
		op = &xxx_DeleteConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ConglomerationID = o.ConglomerationID
	return op
}

func (o *DeleteConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConglomerationID = op.ConglomerationID
}
func (o *DeleteConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteConfigurationResponse structure represents the DeleteConfiguration operation response
type DeleteConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteConfigurationOperation) *xxx_DeleteConfigurationOperation {
	if op == nil {
		op = &xxx_DeleteConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
