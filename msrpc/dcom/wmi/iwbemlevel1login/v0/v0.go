package iwbemlevel1login

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
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = wmi.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemLevel1Login interface identifier f309ad18-d86a-11d0-a075-00c04fb68820
	Level1LoginIID = &dcom.IID{Data1: 0xf309ad18, Data2: 0xd86a, Data3: 0x11d0, Data4: []byte{0xa0, 0x75, 0x00, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax UUID
	Level1LoginSyntaxUUID = &uuid.UUID{TimeLow: 0xf309ad18, TimeMid: 0xd86a, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x75, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	// Syntax ID
	Level1LoginSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Level1LoginSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemLevel1Login interface.
type Level1LoginClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemLevel1Login::EstablishPosition method does not perform any action. The return
	// value and output parameter are used in locale negotiation as specified in section
	// 3.2.3.
	//
	// Return Values: The server MUST return one of the following values, based on server
	// behavior for the wszPreferredLocale parameter in IWbemLevel1Login::NTLMLogin.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|        RETURN        |                                                                                  |
	//	|      VALUE/CODE      |                                   DESCRIPTION                                    |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| 0x00 WBEM_S_NO_ERROR | The connection was established and no error occurred.<27>                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| 0x80004001 E_NOTIMPL | The attempted operation is not implemented. The value of this element is as      |
	//	|                      | specified in [MS-ERREF] section 2.1.<28>                                         |
	//	+----------------------+----------------------------------------------------------------------------------+
	EstablishPosition(context.Context, *EstablishPositionRequest, ...dcerpc.CallOption) (*EstablishPositionResponse, error)

	// This method does not perform any action.
	RequestChallenge(context.Context, *RequestChallengeRequest, ...dcerpc.CallOption) (*RequestChallengeResponse, error)

	// This method does not perform any action.
	WBEMLogin(context.Context, *WBEMLoginRequest, ...dcerpc.CallOption) (*WBEMLoginResponse, error)

	// The IWbemLevel1Login::NTLMLogin method MUST connect a user to the management services
	// interface in a specified namespace.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	NTLMLogin(context.Context, *NTLMLoginRequest, ...dcerpc.CallOption) (*NTLMLoginResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Level1LoginClient
}

type xxx_DefaultLevel1LoginClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultLevel1LoginClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultLevel1LoginClient) EstablishPosition(ctx context.Context, in *EstablishPositionRequest, opts ...dcerpc.CallOption) (*EstablishPositionResponse, error) {
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
	out := &EstablishPositionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLevel1LoginClient) RequestChallenge(ctx context.Context, in *RequestChallengeRequest, opts ...dcerpc.CallOption) (*RequestChallengeResponse, error) {
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
	out := &RequestChallengeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLevel1LoginClient) WBEMLogin(ctx context.Context, in *WBEMLoginRequest, opts ...dcerpc.CallOption) (*WBEMLoginResponse, error) {
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
	out := &WBEMLoginResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLevel1LoginClient) NTLMLogin(ctx context.Context, in *NTLMLoginRequest, opts ...dcerpc.CallOption) (*NTLMLoginResponse, error) {
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
	out := &NTLMLoginResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLevel1LoginClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLevel1LoginClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultLevel1LoginClient) IPID(ctx context.Context, ipid *dcom.IPID) Level1LoginClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultLevel1LoginClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewLevel1LoginClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Level1LoginClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Level1LoginSyntaxV0_0))...)
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
	return &xxx_DefaultLevel1LoginClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_EstablishPositionOperation structure represents the EstablishPosition operation
type xxx_EstablishPositionOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	_             string         `idl:"name:reserved1;string;pointer:unique"`
	_             uint32         `idl:"name:reserved2"`
	LocaleVersion uint32         `idl:"name:LocaleVersion" json:"locale_version"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EstablishPositionOperation) OpNum() int { return 3 }

func (o *xxx_EstablishPositionOperation) OpName() string {
	return "/IWbemLevel1Login/v0/EstablishPosition"
}

func (o *xxx_EstablishPositionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishPositionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved reserved2
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishPositionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		var _reserved1 string
		_ptr_reserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &_reserved1); err != nil {
				return err
			}
			return nil
		})
		_s_reserved1 := func(ptr interface{}) { _reserved1 = *ptr.(*string) }
		if err := w.ReadPointer(&_reserved1, _s_reserved1, _ptr_reserved1); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{alias=DWORD}(uint32))
	{
		// reserved reserved2
		var _reserved2 uint32
		if err := w.ReadData(&_reserved2); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishPositionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishPositionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// LocaleVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EstablishPositionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// LocaleVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EstablishPositionRequest structure represents the EstablishPosition operation request
type EstablishPositionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *EstablishPositionRequest) xxx_ToOp(ctx context.Context, op *xxx_EstablishPositionOperation) *xxx_EstablishPositionOperation {
	if op == nil {
		op = &xxx_EstablishPositionOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *EstablishPositionRequest) xxx_FromOp(ctx context.Context, op *xxx_EstablishPositionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *EstablishPositionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EstablishPositionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishPositionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EstablishPositionResponse structure represents the EstablishPosition operation response
type EstablishPositionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LocaleVersion: The server MUST set the value of LocaleVersion based on the server
	// behavior when IWbemLevel1Login::NTLMLogin is passed an unrecognized locale name in
	// the wszPreferredLocale parameter:
	//
	// * If the server ignores an unrecognized locale name in the Locale Name Format, as
	// specified in section 2.2.29 ( 259edd31-d6eb-4bc9-a2c4-2891b78bb51d ) , passed to
	// IWbemLevel1Login::NTLMLogin while all other parameters are valid, and completes the
	// execution of the IWbemLevel1Login::NTLMLogin method, the server MUST set the LocaleVersion
	// parameter to 1.
	//
	// * If the server returns an error for an unrecognized locale name in Locale Name Format,
	// as specified in section 2.2.29, passed to IWbemLevel1Login::NTLMLogin, while all
	// other parameters are valid, the server MUST set the LocaleVersion parameter to 0.
	LocaleVersion uint32 `idl:"name:LocaleVersion" json:"locale_version"`
	// Return: The EstablishPosition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EstablishPositionResponse) xxx_ToOp(ctx context.Context, op *xxx_EstablishPositionOperation) *xxx_EstablishPositionOperation {
	if op == nil {
		op = &xxx_EstablishPositionOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.LocaleVersion = op.LocaleVersion
	o.Return = op.Return
	return op
}

func (o *EstablishPositionResponse) xxx_FromOp(ctx context.Context, op *xxx_EstablishPositionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LocaleVersion = op.LocaleVersion
	o.Return = op.Return
}
func (o *EstablishPositionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EstablishPositionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EstablishPositionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RequestChallengeOperation structure represents the RequestChallenge operation
type xxx_RequestChallengeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	_      string         `idl:"name:reserved1;string;pointer:unique"`
	_      string         `idl:"name:reserved2;string;pointer:unique"`
	_      []byte         `idl:"name:reserved3;size_is:(16);length_is:(16)"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RequestChallengeOperation) OpNum() int { return 4 }

func (o *xxx_RequestChallengeOperation) OpName() string {
	return "/IWbemLevel1Login/v0/RequestChallenge"
}

func (o *xxx_RequestChallengeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestChallengeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved2
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestChallengeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		var _reserved1 string
		_ptr_reserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &_reserved1); err != nil {
				return err
			}
			return nil
		})
		_s_reserved1 := func(ptr interface{}) { _reserved1 = *ptr.(*string) }
		if err := w.ReadPointer(&_reserved1, _s_reserved1, _ptr_reserved1); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved2
		var _reserved2 string
		_ptr_reserved2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &_reserved2); err != nil {
				return err
			}
			return nil
		})
		_s_reserved2 := func(ptr interface{}) { _reserved2 = *ptr.(*string) }
		if err := w.ReadPointer(&_reserved2, _s_reserved2, _ptr_reserved2); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestChallengeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestChallengeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// reserved3 {out} (1:{pointer=ref}*(1)[dim:0,size_is=16,length_is=16](uchar))
	{
		// reserved reserved3
		dimSize1 := uint64(16)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(16)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RequestChallengeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// reserved3 {out} (1:{pointer=ref}*(1)[dim:0,size_is=16,length_is=16](uchar))
	{
		// reserved reserved3
		var _reserved3 []byte
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _reserved3", sizeInfo[0])
		}
		_reserved3 = make([]byte, sizeInfo[0])
		for i1 := range _reserved3 {
			i1 := i1
			if err := w.ReadData(&_reserved3[i1]); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RequestChallengeRequest structure represents the RequestChallenge operation request
type RequestChallengeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RequestChallengeRequest) xxx_ToOp(ctx context.Context, op *xxx_RequestChallengeOperation) *xxx_RequestChallengeOperation {
	if op == nil {
		op = &xxx_RequestChallengeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *RequestChallengeRequest) xxx_FromOp(ctx context.Context, op *xxx_RequestChallengeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RequestChallengeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RequestChallengeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestChallengeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RequestChallengeResponse structure represents the RequestChallenge operation response
type RequestChallengeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RequestChallenge return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RequestChallengeResponse) xxx_ToOp(ctx context.Context, op *xxx_RequestChallengeOperation) *xxx_RequestChallengeOperation {
	if op == nil {
		op = &xxx_RequestChallengeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *RequestChallengeResponse) xxx_FromOp(ctx context.Context, op *xxx_RequestChallengeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RequestChallengeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RequestChallengeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RequestChallengeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WBEMLoginOperation structure represents the WBEMLogin operation
type xxx_WBEMLoginOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	_      string         `idl:"name:reserved1;string;pointer:unique"`
	_      []byte         `idl:"name:reserved2;size_is:(16);length_is:(16);pointer:unique"`
	_      int32          `idl:"name:reserved3"`
	_      *wmi.Context   `idl:"name:reserved4"`
	_      *wmi.Services  `idl:"name:reserved5"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WBEMLoginOperation) OpNum() int { return 5 }

func (o *xxx_WBEMLoginOperation) OpName() string { return "/IWbemLevel1Login/v0/WBEMLogin" }

func (o *xxx_WBEMLoginOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WBEMLoginOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{pointer=unique}*(1)[dim:0,size_is=16,length_is=16](uchar))
	{
		// reserved reserved2
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reserved3 {in} (1:(int32))
	{
		// reserved reserved3
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	}
	// reserved4 {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		// reserved reserved4
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WBEMLoginOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// reserved1 {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		// reserved reserved1
		var _reserved1 string
		_ptr_reserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &_reserved1); err != nil {
				return err
			}
			return nil
		})
		_s_reserved1 := func(ptr interface{}) { _reserved1 = *ptr.(*string) }
		if err := w.ReadPointer(&_reserved1, _s_reserved1, _ptr_reserved1); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reserved2 {in} (1:{pointer=unique}*(1)[dim:0,size_is=16,length_is=16](uchar))
	{
		// reserved reserved2
		var _reserved2 []byte
		_ptr_reserved2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _reserved2", sizeInfo[0])
			}
			_reserved2 = make([]byte, sizeInfo[0])
			for i1 := range _reserved2 {
				i1 := i1
				if err := w.ReadData(&_reserved2[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_reserved2 := func(ptr interface{}) { _reserved2 = *ptr.(*[]byte) }
		if err := w.ReadPointer(&_reserved2, _s_reserved2, _ptr_reserved2); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reserved3 {in} (1:(int32))
	{
		// reserved reserved3
		var _reserved3 int32
		if err := w.ReadData(&_reserved3); err != nil {
			return err
		}
	}
	// reserved4 {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		// reserved reserved4
		var _reserved4 *wmi.Context
		_ptr_reserved4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _reserved4 == nil {
				_reserved4 = &wmi.Context{}
			}
			if err := _reserved4.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reserved4 := func(ptr interface{}) { _reserved4 = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&_reserved4, _s_reserved4, _ptr_reserved4); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WBEMLoginOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WBEMLoginOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// reserved5 {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		// reserved reserved5
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WBEMLoginOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// reserved5 {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		// reserved reserved5
		var _reserved5 *wmi.Services
		_ptr_reserved5 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _reserved5 == nil {
				_reserved5 = &wmi.Services{}
			}
			if err := _reserved5.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reserved5 := func(ptr interface{}) { _reserved5 = *ptr.(**wmi.Services) }
		if err := w.ReadPointer(&_reserved5, _s_reserved5, _ptr_reserved5); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WBEMLoginRequest structure represents the WBEMLogin operation request
type WBEMLoginRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *WBEMLoginRequest) xxx_ToOp(ctx context.Context, op *xxx_WBEMLoginOperation) *xxx_WBEMLoginOperation {
	if op == nil {
		op = &xxx_WBEMLoginOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *WBEMLoginRequest) xxx_FromOp(ctx context.Context, op *xxx_WBEMLoginOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *WBEMLoginRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WBEMLoginRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WBEMLoginOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WBEMLoginResponse structure represents the WBEMLogin operation response
type WBEMLoginResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The WBEMLogin return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WBEMLoginResponse) xxx_ToOp(ctx context.Context, op *xxx_WBEMLoginOperation) *xxx_WBEMLoginOperation {
	if op == nil {
		op = &xxx_WBEMLoginOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *WBEMLoginResponse) xxx_FromOp(ctx context.Context, op *xxx_WBEMLoginOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *WBEMLoginResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WBEMLoginResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WBEMLoginOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NTLMLoginOperation structure represents the NTLMLogin operation
type xxx_NTLMLoginOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	NetworkResource string         `idl:"name:wszNetworkResource;string;pointer:unique" json:"network_resource"`
	PreferredLocale string         `idl:"name:wszPreferredLocale;string;pointer:unique" json:"preferred_locale"`
	Flags           int32          `idl:"name:lFlags" json:"flags"`
	Context         *wmi.Context   `idl:"name:pCtx" json:"context"`
	Namespace       *wmi.Services  `idl:"name:ppNamespace" json:"namespace"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_NTLMLoginOperation) OpNum() int { return 6 }

func (o *xxx_NTLMLoginOperation) OpName() string { return "/IWbemLevel1Login/v0/NTLMLogin" }

func (o *xxx_NTLMLoginOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NTLMLoginOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wszNetworkResource {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NetworkResource != "" {
			_ptr_wszNetworkResource := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NetworkResource); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NetworkResource, _ptr_wszNetworkResource); err != nil {
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
	// wszPreferredLocale {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.PreferredLocale != "" {
			_ptr_wszPreferredLocale := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PreferredLocale); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PreferredLocale, _ptr_wszPreferredLocale); err != nil {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pCtx := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Context != nil {
					if err := o.Context.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Context{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Context, _ptr_pCtx); err != nil {
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

func (o *xxx_NTLMLoginOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wszNetworkResource {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_wszNetworkResource := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NetworkResource); err != nil {
				return err
			}
			return nil
		})
		_s_wszNetworkResource := func(ptr interface{}) { o.NetworkResource = *ptr.(*string) }
		if err := w.ReadPointer(&o.NetworkResource, _s_wszNetworkResource, _ptr_wszNetworkResource); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wszPreferredLocale {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_wszPreferredLocale := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PreferredLocale); err != nil {
				return err
			}
			return nil
		})
		_s_wszPreferredLocale := func(ptr interface{}) { o.PreferredLocale = *ptr.(*string) }
		if err := w.ReadPointer(&o.PreferredLocale, _s_wszPreferredLocale, _ptr_wszPreferredLocale); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pCtx {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pCtx := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pCtx := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pCtx, _ptr_pCtx); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NTLMLoginOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NTLMLoginOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppNamespace {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		if o.Namespace != nil {
			_ptr_ppNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Namespace != nil {
					if err := o.Namespace.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.Services{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Namespace, _ptr_ppNamespace); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NTLMLoginOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppNamespace {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemServices}(interface))
	{
		_ptr_ppNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Namespace == nil {
				o.Namespace = &wmi.Services{}
			}
			if err := o.Namespace.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppNamespace := func(ptr interface{}) { o.Namespace = *ptr.(**wmi.Services) }
		if err := w.ReadPointer(&o.Namespace, _s_ppNamespace, _ptr_ppNamespace); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NTLMLoginRequest structure represents the NTLMLogin operation request
type NTLMLoginRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// wszNetworkResource: The string MUST represent the namespace on the server to which
	// the returned IWbemServices object is associated. This parameter MUST NOT be NULL
	// and MUST match the namespace syntax as specified in section 2.2.2.
	NetworkResource string `idl:"name:wszNetworkResource;string;pointer:unique" json:"network_resource"`
	// wszPreferredLocale: MUST be a pointer to a string that MUST specify the locale values
	// in the preferred order, separated by a comma. If the client does not supply it, the
	// server creates a default list which is implementation-specific.<29> Each locale format
	// SHOULD conform to the WMI locale format, as specified in WMI Locale Formats (section
	// 2.2.29). Any subsequent calls that request CIM localizable information (WBEM_FLAG_USE_AMENDED_QUALIFIERS)
	// SHOULD return the localized information in the order of preference if the information
	// is available in the LCID.<30> The server MUST save this information in ClientPreferredLocales.
	PreferredLocale string `idl:"name:wszPreferredLocale;string;pointer:unique" json:"preferred_locale"`
	// lFlags: MUST be 0. The server SHOULD consider any other value as not valid and return
	// WBEM_E_INVALID_PARAMETER; otherwise, the server behavior is implementation-specific.<31>
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pCtx: MUST be a pointer to an IWbemContext interface, which MUST contain additional
	// information sent by the client. If pCtx is NULL, the parameter MUST be ignored.
	Context *wmi.Context `idl:"name:pCtx" json:"context"`
}

func (o *NTLMLoginRequest) xxx_ToOp(ctx context.Context, op *xxx_NTLMLoginOperation) *xxx_NTLMLoginOperation {
	if op == nil {
		op = &xxx_NTLMLoginOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.NetworkResource = op.NetworkResource
	o.PreferredLocale = op.PreferredLocale
	o.Flags = op.Flags
	o.Context = op.Context
	return op
}

func (o *NTLMLoginRequest) xxx_FromOp(ctx context.Context, op *xxx_NTLMLoginOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NetworkResource = op.NetworkResource
	o.PreferredLocale = op.PreferredLocale
	o.Flags = op.Flags
	o.Context = op.Context
}
func (o *NTLMLoginRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *NTLMLoginRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NTLMLoginOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NTLMLoginResponse structure represents the NTLMLogin operation response
type NTLMLoginResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppNamespace: If the call succeeds, ppNamespace MUST return a pointer to an IWbemServices
	// interface pointer. This parameter MUST be set to NULL when an error occurs.
	Namespace *wmi.Services `idl:"name:ppNamespace" json:"namespace"`
	// Return: The NTLMLogin return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *NTLMLoginResponse) xxx_ToOp(ctx context.Context, op *xxx_NTLMLoginOperation) *xxx_NTLMLoginOperation {
	if op == nil {
		op = &xxx_NTLMLoginOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Namespace = op.Namespace
	o.Return = op.Return
	return op
}

func (o *NTLMLoginResponse) xxx_FromOp(ctx context.Context, op *xxx_NTLMLoginOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Namespace = op.Namespace
	o.Return = op.Return
}
func (o *NTLMLoginResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *NTLMLoginResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NTLMLoginOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
