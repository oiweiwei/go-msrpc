package ivdsserviceiscsi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = iunknown.GoPackage
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsServiceIscsi interface identifier 14fbe036-3ed7-4e10-90e9-a5ff991aff01
	ServiceISCSIIID = &dcom.IID{Data1: 0x14fbe036, Data2: 0x3ed7, Data3: 0x4e10, Data4: []byte{0x90, 0xe9, 0xa5, 0xff, 0x99, 0x1a, 0xff, 0x01}}
	// Syntax UUID
	ServiceISCSISyntaxUUID = &uuid.UUID{TimeLow: 0x14fbe036, TimeMid: 0x3ed7, TimeHiAndVersion: 0x4e10, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xe9, Node: [6]uint8{0xa5, 0xff, 0x99, 0x1a, 0xff, 0x1}}
	// Syntax ID
	ServiceISCSISyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceISCSISyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceIscsi interface.
type ServiceISCSIClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetInitiatorName method returns the iSCSI name of the initiator service.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetInitiatorName(context.Context, *GetInitiatorNameRequest, ...dcerpc.CallOption) (*GetInitiatorNameResponse, error)

	// The QueryInitiatorAdapters method returns an object that enumerates the iSCSI initiator
	// adapters of the initiator.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryInitiatorAdapters(context.Context, *QueryInitiatorAdaptersRequest, ...dcerpc.CallOption) (*QueryInitiatorAdaptersResponse, error)

	// Opnum05NotUsedOnWire operation.
	// Opnum05NotUsedOnWire

	// Opnum06NotUsedOnWire operation.
	// Opnum06NotUsedOnWire

	// Opnum07NotUsedOnWire operation.
	// Opnum07NotUsedOnWire

	// The SetInitiatorSharedSecret method sets the initiator CHAP shared secret that is
	// used for mutual CHAP authentication when the initiator authenticates the target.
	// For more information on CHAP, see [MS-CHAP].<74>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetInitiatorSharedSecret(context.Context, *SetInitiatorSharedSecretRequest, ...dcerpc.CallOption) (*SetInitiatorSharedSecretResponse, error)

	// Opnum09NotUsedOnWire operation.
	// Opnum09NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceISCSIClient
}

type xxx_DefaultServiceISCSIClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceISCSIClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceISCSIClient) GetInitiatorName(ctx context.Context, in *GetInitiatorNameRequest, opts ...dcerpc.CallOption) (*GetInitiatorNameResponse, error) {
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
	out := &GetInitiatorNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceISCSIClient) QueryInitiatorAdapters(ctx context.Context, in *QueryInitiatorAdaptersRequest, opts ...dcerpc.CallOption) (*QueryInitiatorAdaptersResponse, error) {
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
	out := &QueryInitiatorAdaptersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceISCSIClient) SetInitiatorSharedSecret(ctx context.Context, in *SetInitiatorSharedSecretRequest, opts ...dcerpc.CallOption) (*SetInitiatorSharedSecretResponse, error) {
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
	out := &SetInitiatorSharedSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceISCSIClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceISCSIClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceISCSIClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceISCSIClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewServiceISCSIClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceISCSIClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceISCSISyntaxV0_0))...)
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
	return &xxx_DefaultServiceISCSIClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetInitiatorNameOperation structure represents the GetInitiatorName operation
type xxx_GetInitiatorNameOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ISCSIName string         `idl:"name:ppwszIscsiName;string" json:"iscsi_name"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInitiatorNameOperation) OpNum() int { return 3 }

func (o *xxx_GetInitiatorNameOperation) OpName() string {
	return "/IVdsServiceIscsi/v0/GetInitiatorName"
}

func (o *xxx_GetInitiatorNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInitiatorNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetInitiatorNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetInitiatorNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInitiatorNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppwszIscsiName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ISCSIName != "" {
			_ptr_ppwszIscsiName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ISCSIName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ISCSIName, _ptr_ppwszIscsiName); err != nil {
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

func (o *xxx_GetInitiatorNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppwszIscsiName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ppwszIscsiName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ISCSIName); err != nil {
				return err
			}
			return nil
		})
		_s_ppwszIscsiName := func(ptr interface{}) { o.ISCSIName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ISCSIName, _s_ppwszIscsiName, _ptr_ppwszIscsiName); err != nil {
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

// GetInitiatorNameRequest structure represents the GetInitiatorName operation request
type GetInitiatorNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetInitiatorNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetInitiatorNameOperation {
	if o == nil {
		return &xxx_GetInitiatorNameOperation{}
	}
	return &xxx_GetInitiatorNameOperation{
		This: o.This,
	}
}

func (o *GetInitiatorNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInitiatorNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetInitiatorNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetInitiatorNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInitiatorNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInitiatorNameResponse structure represents the GetInitiatorName operation response
type GetInitiatorNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszIscsiName: A pointer that, if the operation is successfully completed, receives
	// a null-terminated Unicode string with the iSCSI name.
	ISCSIName string `idl:"name:ppwszIscsiName;string" json:"iscsi_name"`
	// Return: The GetInitiatorName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetInitiatorNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetInitiatorNameOperation {
	if o == nil {
		return &xxx_GetInitiatorNameOperation{}
	}
	return &xxx_GetInitiatorNameOperation{
		That:      o.That,
		ISCSIName: o.ISCSIName,
		Return:    o.Return,
	}
}

func (o *GetInitiatorNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInitiatorNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ISCSIName = op.ISCSIName
	o.Return = op.Return
}
func (o *GetInitiatorNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetInitiatorNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInitiatorNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInitiatorAdaptersOperation structure represents the QueryInitiatorAdapters operation
type xxx_QueryInitiatorAdaptersOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInitiatorAdaptersOperation) OpNum() int { return 4 }

func (o *xxx_QueryInitiatorAdaptersOperation) OpName() string {
	return "/IVdsServiceIscsi/v0/QueryInitiatorAdapters"
}

func (o *xxx_QueryInitiatorAdaptersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInitiatorAdaptersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryInitiatorAdaptersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryInitiatorAdaptersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInitiatorAdaptersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryInitiatorAdaptersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryInitiatorAdaptersRequest structure represents the QueryInitiatorAdapters operation request
type QueryInitiatorAdaptersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryInitiatorAdaptersRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInitiatorAdaptersOperation {
	if o == nil {
		return &xxx_QueryInitiatorAdaptersOperation{}
	}
	return &xxx_QueryInitiatorAdaptersOperation{
		This: o.This,
	}
}

func (o *QueryInitiatorAdaptersRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInitiatorAdaptersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryInitiatorAdaptersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInitiatorAdaptersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInitiatorAdaptersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInitiatorAdaptersResponse structure represents the QueryInitiatorAdapters operation response
type QueryInitiatorAdaptersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object that contains an enumeration
	// of initiator adapter objects on the server. Callers MUST release the interface when
	// they are done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryInitiatorAdapters return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryInitiatorAdaptersResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInitiatorAdaptersOperation {
	if o == nil {
		return &xxx_QueryInitiatorAdaptersOperation{}
	}
	return &xxx_QueryInitiatorAdaptersOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *QueryInitiatorAdaptersResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInitiatorAdaptersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryInitiatorAdaptersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInitiatorAdaptersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInitiatorAdaptersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInitiatorSharedSecretOperation structure represents the SetInitiatorSharedSecret operation
type xxx_SetInitiatorSharedSecretOperation struct {
	This                  *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat         `idl:"name:That" json:"that"`
	InitiatorSharedSecret *vds.ISCSISharedSecret `idl:"name:pInitiatorSharedSecret;pointer:unique" json:"initiator_shared_secret"`
	TargetID              *vds.ObjectID          `idl:"name:targetId" json:"target_id"`
	Return                int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInitiatorSharedSecretOperation) OpNum() int { return 8 }

func (o *xxx_SetInitiatorSharedSecretOperation) OpName() string {
	return "/IVdsServiceIscsi/v0/SetInitiatorSharedSecret"
}

func (o *xxx_SetInitiatorSharedSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInitiatorSharedSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pInitiatorSharedSecret {in} (1:{pointer=unique}*(1))(2:{alias=VDS_ISCSI_SHARED_SECRET}(struct))
	{
		if o.InitiatorSharedSecret != nil {
			_ptr_pInitiatorSharedSecret := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.InitiatorSharedSecret != nil {
					if err := o.InitiatorSharedSecret.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.ISCSISharedSecret{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InitiatorSharedSecret, _ptr_pInitiatorSharedSecret); err != nil {
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
	// targetId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.TargetID != nil {
			if err := o.TargetID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetInitiatorSharedSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pInitiatorSharedSecret {in} (1:{pointer=unique}*(1))(2:{alias=VDS_ISCSI_SHARED_SECRET}(struct))
	{
		_ptr_pInitiatorSharedSecret := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.InitiatorSharedSecret == nil {
				o.InitiatorSharedSecret = &vds.ISCSISharedSecret{}
			}
			if err := o.InitiatorSharedSecret.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pInitiatorSharedSecret := func(ptr interface{}) { o.InitiatorSharedSecret = *ptr.(**vds.ISCSISharedSecret) }
		if err := w.ReadPointer(&o.InitiatorSharedSecret, _s_pInitiatorSharedSecret, _ptr_pInitiatorSharedSecret); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// targetId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.TargetID == nil {
			o.TargetID = &vds.ObjectID{}
		}
		if err := o.TargetID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInitiatorSharedSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInitiatorSharedSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetInitiatorSharedSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetInitiatorSharedSecretRequest structure represents the SetInitiatorSharedSecret operation request
type SetInitiatorSharedSecretRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pInitiatorSharedSecret: A pointer to a VDS_ISCSI_SHARED_SECRET structure that contains
	// the CHAP shared secret that is used for mutual CHAP authentication when the initiator
	// authenticates the target.
	InitiatorSharedSecret *vds.ISCSISharedSecret `idl:"name:pInitiatorSharedSecret;pointer:unique" json:"initiator_shared_secret"`
	// targetId: This parameter is reserved and not used by the protocol. Callers MUST pass
	// in GUID_NULL. Callers MUST pass in GUID_NULL.
	TargetID *vds.ObjectID `idl:"name:targetId" json:"target_id"`
}

func (o *SetInitiatorSharedSecretRequest) xxx_ToOp(ctx context.Context) *xxx_SetInitiatorSharedSecretOperation {
	if o == nil {
		return &xxx_SetInitiatorSharedSecretOperation{}
	}
	return &xxx_SetInitiatorSharedSecretOperation{
		This:                  o.This,
		InitiatorSharedSecret: o.InitiatorSharedSecret,
		TargetID:              o.TargetID,
	}
}

func (o *SetInitiatorSharedSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInitiatorSharedSecretOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InitiatorSharedSecret = op.InitiatorSharedSecret
	o.TargetID = op.TargetID
}
func (o *SetInitiatorSharedSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInitiatorSharedSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInitiatorSharedSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInitiatorSharedSecretResponse structure represents the SetInitiatorSharedSecret operation response
type SetInitiatorSharedSecretResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetInitiatorSharedSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetInitiatorSharedSecretResponse) xxx_ToOp(ctx context.Context) *xxx_SetInitiatorSharedSecretOperation {
	if o == nil {
		return &xxx_SetInitiatorSharedSecretOperation{}
	}
	return &xxx_SetInitiatorSharedSecretOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetInitiatorSharedSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInitiatorSharedSecretOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetInitiatorSharedSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInitiatorSharedSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInitiatorSharedSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
