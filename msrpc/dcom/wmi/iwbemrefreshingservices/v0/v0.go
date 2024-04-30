package iwbemrefreshingservices

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
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = wmi.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemRefreshingServices interface identifier 2c9273e0-1dc3-11d3-b364-00105a1f8177
	RefreshingServicesIID = &dcom.IID{Data1: 0x2c9273e0, Data2: 0x1dc3, Data3: 0x11d3, Data4: []byte{0xb3, 0x64, 0x00, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	// Syntax UUID
	RefreshingServicesSyntaxUUID = &uuid.UUID{TimeLow: 0x2c9273e0, TimeMid: 0x1dc3, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	// Syntax ID
	RefreshingServicesSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RefreshingServicesSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemRefreshingServices interface.
type RefreshingServicesClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemRefreshingServices::AddObjectToRefresher method MUST add a CIM instance,
	// which is identified by its CIM path, to the list of CIM instances that can be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddObjectToRefresher(context.Context, *AddObjectToRefresherRequest, ...dcerpc.CallOption) (*AddObjectToRefresherResponse, error)

	// The IWbemRefreshingServices::AddObjectToRefresherByTemplate method MUST add a CIM
	// instance, which is identified by its CIM object instance, to the list of CIM instances
	// to be refreshed.
	//
	// The AddObjectToRefresherByTemplate method opnum equals 4.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddObjectToRefresherByTemplate(context.Context, *AddObjectToRefresherByTemplateRequest, ...dcerpc.CallOption) (*AddObjectToRefresherByTemplateResponse, error)

	// The IWbemRefreshingServices::AddEnumToRefresher method MUST add all CIM instances
	// that are identified by the CIM class name to the list of CIM instances to be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	AddEnumToRefresher(context.Context, *AddEnumToRefresherRequest, ...dcerpc.CallOption) (*AddEnumToRefresherResponse, error)

	// The IWbemRefreshingServices::RemoveObjectFromRefresher method MUST remove a CIM instance,
	// which is identified by its CIM path, from the list of CIM instances that can be refreshed.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. If there are no failures, the server MUST always return WBEM_E_NOT_AVAILABLE.<68>
	RemoveObjectFromRefresher(context.Context, *RemoveObjectFromRefresherRequest, ...dcerpc.CallOption) (*RemoveObjectFromRefresherResponse, error)

	// The IWbemRefreshingServices::GetRemoteRefresher method MUST return an IWbemRemoteRefresher
	// interface pointer. This pointer is needed by the client to refresh objects and enumerations.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// In case of failure, the server MUST return an HRESULT whose S (severity) bit is set
	// as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	GetRemoteRefresher(context.Context, *GetRemoteRefresherRequest, ...dcerpc.CallOption) (*GetRemoteRefresherResponse, error)

	// The IWbemRefreshingServices::ReconnectRemoteRefresher method MUST restore a set of
	// CIM instances and enumerations that are passed in apReconnectInfo to a refresher.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR, as specified in section
	// 2.2.11, to indicate the successful completion of the method.
	ReconnectRemoteRefresher(context.Context, *ReconnectRemoteRefresherRequest, ...dcerpc.CallOption) (*ReconnectRemoteRefresherResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RefreshingServicesClient
}

type xxx_DefaultRefreshingServicesClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRefreshingServicesClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRefreshingServicesClient) AddObjectToRefresher(ctx context.Context, in *AddObjectToRefresherRequest, opts ...dcerpc.CallOption) (*AddObjectToRefresherResponse, error) {
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
	out := &AddObjectToRefresherResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) AddObjectToRefresherByTemplate(ctx context.Context, in *AddObjectToRefresherByTemplateRequest, opts ...dcerpc.CallOption) (*AddObjectToRefresherByTemplateResponse, error) {
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
	out := &AddObjectToRefresherByTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) AddEnumToRefresher(ctx context.Context, in *AddEnumToRefresherRequest, opts ...dcerpc.CallOption) (*AddEnumToRefresherResponse, error) {
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
	out := &AddEnumToRefresherResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) RemoveObjectFromRefresher(ctx context.Context, in *RemoveObjectFromRefresherRequest, opts ...dcerpc.CallOption) (*RemoveObjectFromRefresherResponse, error) {
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
	out := &RemoveObjectFromRefresherResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) GetRemoteRefresher(ctx context.Context, in *GetRemoteRefresherRequest, opts ...dcerpc.CallOption) (*GetRemoteRefresherResponse, error) {
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
	out := &GetRemoteRefresherResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) ReconnectRemoteRefresher(ctx context.Context, in *ReconnectRemoteRefresherRequest, opts ...dcerpc.CallOption) (*ReconnectRemoteRefresherResponse, error) {
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
	out := &ReconnectRemoteRefresherResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRefreshingServicesClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRefreshingServicesClient) IPID(ctx context.Context, ipid *dcom.IPID) RefreshingServicesClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRefreshingServicesClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewRefreshingServicesClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RefreshingServicesClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RefreshingServicesSyntaxV0_0))...)
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
	return &xxx_DefaultRefreshingServicesClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_AddObjectToRefresherOperation structure represents the AddObjectToRefresher operation
type xxx_AddObjectToRefresherOperation struct {
	This                   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	Path                   string           `idl:"name:wszPath;string" json:"path"`
	Flags                  int32            `idl:"name:lFlags" json:"flags"`
	Context                *wmi.Context     `idl:"name:pContext" json:"context"`
	ClientRefresherVersion uint32           `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	Info                   *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	ServerRefresherVersion uint32           `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_AddObjectToRefresherOperation) OpNum() int { return 3 }

func (o *xxx_AddObjectToRefresherOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/AddObjectToRefresher"
}

func (o *xxx_AddObjectToRefresherOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// wszPath {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Context, _ptr_pContext); err != nil {
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
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wszPath {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pContext := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pContext, _ptr_pContext); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefreshInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_AddObjectToRefresherOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info == nil {
			o.Info = &wmi.RefreshInfo{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// AddObjectToRefresherRequest structure represents the AddObjectToRefresher operation request
type AddObjectToRefresherRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, which identifies the client that is requesting refreshing services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// wszPath: MUST be a string that MUST contain the CIM path of the CIM instance. This
	// parameter MUST NOT be NULL.
	Path string `idl:"name:wszPath;string" json:"path"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pContext: MUST be a pointer to an IWbemContext interface object, which MUST contain
	// additional information for the server refresher. If pContext is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pContext" json:"context"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<64>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
}

func (o *AddObjectToRefresherRequest) xxx_ToOp(ctx context.Context) *xxx_AddObjectToRefresherOperation {
	if o == nil {
		return &xxx_AddObjectToRefresherOperation{}
	}
	return &xxx_AddObjectToRefresherOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		Path:                   o.Path,
		Flags:                  o.Flags,
		Context:                o.Context,
		ClientRefresherVersion: o.ClientRefresherVersion,
	}
}

func (o *AddObjectToRefresherRequest) xxx_FromOp(ctx context.Context, op *xxx_AddObjectToRefresherOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.Path = op.Path
	o.Flags = op.Flags
	o.Context = op.Context
	o.ClientRefresherVersion = op.ClientRefresherVersion
}
func (o *AddObjectToRefresherRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddObjectToRefresherRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddObjectToRefresherOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddObjectToRefresherResponse structure represents the AddObjectToRefresher operation response
type AddObjectToRefresherResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pInfo: MUST be an output parameter that MUST return a _WBEM_REFRESH_INFO structure,
	// as specified in section 2.2.20, which MUST contain refresher information about the
	// CIM instance in wszPath. It MUST NOT be NULL.
	Info *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	// pdwSvrRefrVersion: MUST be an output parameter that MUST be the version of the server
	// refresher. The value of this parameter SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The AddObjectToRefresher return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddObjectToRefresherResponse) xxx_ToOp(ctx context.Context) *xxx_AddObjectToRefresherOperation {
	if o == nil {
		return &xxx_AddObjectToRefresherOperation{}
	}
	return &xxx_AddObjectToRefresherOperation{
		That:                   o.That,
		Info:                   o.Info,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *AddObjectToRefresherResponse) xxx_FromOp(ctx context.Context, op *xxx_AddObjectToRefresherOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *AddObjectToRefresherResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddObjectToRefresherResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddObjectToRefresherOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddObjectToRefresherByTemplateOperation structure represents the AddObjectToRefresherByTemplate operation
type xxx_AddObjectToRefresherByTemplateOperation struct {
	This                   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	Template               *wmi.ClassObject `idl:"name:pTemplate" json:"template"`
	Flags                  int32            `idl:"name:lFlags" json:"flags"`
	Context                *wmi.Context     `idl:"name:pContext" json:"context"`
	ClientRefresherVersion uint32           `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	Info                   *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	ServerRefresherVersion uint32           `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) OpNum() int { return 4 }

func (o *xxx_AddObjectToRefresherByTemplateOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/AddObjectToRefresherByTemplate"
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pTemplate {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		if o.Template != nil {
			_ptr_pTemplate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Template != nil {
					if err := o.Template.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.ClassObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Template, _ptr_pTemplate); err != nil {
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
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Context, _ptr_pContext); err != nil {
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
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pTemplate {in} (1:{pointer=ref}*(1))(2:{alias=IWbemClassObject}(interface))
	{
		_ptr_pTemplate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Template == nil {
				o.Template = &wmi.ClassObject{}
			}
			if err := o.Template.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pTemplate := func(ptr interface{}) { o.Template = *ptr.(**wmi.ClassObject) }
		if err := w.ReadPointer(&o.Template, _s_pTemplate, _ptr_pTemplate); err != nil {
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
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pContext := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pContext, _ptr_pContext); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddObjectToRefresherByTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefreshInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_AddObjectToRefresherByTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info == nil {
			o.Info = &wmi.RefreshInfo{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// AddObjectToRefresherByTemplateRequest structure represents the AddObjectToRefresherByTemplate operation request
type AddObjectToRefresherByTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, which identifies the client that is requesting refreshing services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// pTemplate: MUST be a pointer to an IWbemClassObject interface CIM instance that MUST
	// be a template for the CIM instances to be refreshed by the refresher. This parameter
	// MUST NOT be NULL.
	Template *wmi.ClassObject `idl:"name:pTemplate" json:"template"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pContext: MUST be a pointer to an IWbemContext interface object, which MUST contain
	// additional information for the server refresher. If pContext is NULL, the parameter
	// MUST be ignored.
	Context *wmi.Context `idl:"name:pContext" json:"context"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<65>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
}

func (o *AddObjectToRefresherByTemplateRequest) xxx_ToOp(ctx context.Context) *xxx_AddObjectToRefresherByTemplateOperation {
	if o == nil {
		return &xxx_AddObjectToRefresherByTemplateOperation{}
	}
	return &xxx_AddObjectToRefresherByTemplateOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		Template:               o.Template,
		Flags:                  o.Flags,
		Context:                o.Context,
		ClientRefresherVersion: o.ClientRefresherVersion,
	}
}

func (o *AddObjectToRefresherByTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_AddObjectToRefresherByTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.Template = op.Template
	o.Flags = op.Flags
	o.Context = op.Context
	o.ClientRefresherVersion = op.ClientRefresherVersion
}
func (o *AddObjectToRefresherByTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddObjectToRefresherByTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddObjectToRefresherByTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddObjectToRefresherByTemplateResponse structure represents the AddObjectToRefresherByTemplate operation response
type AddObjectToRefresherByTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pInfo: MUST be an output parameter that returns a _WBEM_REFRESH_INFO structure, as
	// specified in section 2.2.20, which MUST contain refresher information about the CIM
	// instance in wszPath. This parameter MUST NOT be NULL.
	Info *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	// pdwSvrRefrVersion: MUST be an output parameter that MUST be the version of the server
	// refresher. The value of this parameter SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The AddObjectToRefresherByTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddObjectToRefresherByTemplateResponse) xxx_ToOp(ctx context.Context) *xxx_AddObjectToRefresherByTemplateOperation {
	if o == nil {
		return &xxx_AddObjectToRefresherByTemplateOperation{}
	}
	return &xxx_AddObjectToRefresherByTemplateOperation{
		That:                   o.That,
		Info:                   o.Info,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *AddObjectToRefresherByTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_AddObjectToRefresherByTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *AddObjectToRefresherByTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddObjectToRefresherByTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddObjectToRefresherByTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddEnumToRefresherOperation structure represents the AddEnumToRefresher operation
type xxx_AddEnumToRefresherOperation struct {
	This                   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	Class                  string           `idl:"name:wszClass;string" json:"class"`
	Flags                  int32            `idl:"name:lFlags" json:"flags"`
	Context                *wmi.Context     `idl:"name:pContext" json:"context"`
	ClientRefresherVersion uint32           `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	Info                   *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	ServerRefresherVersion uint32           `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_AddEnumToRefresherOperation) OpNum() int { return 5 }

func (o *xxx_AddEnumToRefresherOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/AddEnumToRefresher"
}

func (o *xxx_AddEnumToRefresherOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddEnumToRefresherOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// wszClass {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Class); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		if o.Context != nil {
			_ptr_pContext := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Context, _ptr_pContext); err != nil {
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
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddEnumToRefresherOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// wszClass {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Class); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pContext {in} (1:{pointer=ref}*(1))(2:{alias=IWbemContext}(interface))
	{
		_ptr_pContext := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Context == nil {
				o.Context = &wmi.Context{}
			}
			if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pContext := func(ptr interface{}) { o.Context = *ptr.(**wmi.Context) }
		if err := w.ReadPointer(&o.Context, _s_pContext, _ptr_pContext); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddEnumToRefresherOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddEnumToRefresherOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefreshInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_AddEnumToRefresherOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pInfo {out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESH_INFO}(struct))
	{
		if o.Info == nil {
			o.Info = &wmi.RefreshInfo{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// AddEnumToRefresherRequest structure represents the AddEnumToRefresher operation request
type AddEnumToRefresherRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, which identifies the client that is requesting refreshing services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// wszClass: MUST be a string that MUST contain the enumeration CIM class name. This
	// parameter MUST NOT be NULL.
	Class string `idl:"name:wszClass;string" json:"class"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// pContext: MUST be a pointer to an IWbemContext interface object, which MUST contain
	// additional information for the server refresher. If pContext is NULL, the parameter
	// is ignored.
	Context *wmi.Context `idl:"name:pContext" json:"context"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<66>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
}

func (o *AddEnumToRefresherRequest) xxx_ToOp(ctx context.Context) *xxx_AddEnumToRefresherOperation {
	if o == nil {
		return &xxx_AddEnumToRefresherOperation{}
	}
	return &xxx_AddEnumToRefresherOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		Class:                  o.Class,
		Flags:                  o.Flags,
		Context:                o.Context,
		ClientRefresherVersion: o.ClientRefresherVersion,
	}
}

func (o *AddEnumToRefresherRequest) xxx_FromOp(ctx context.Context, op *xxx_AddEnumToRefresherOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.Class = op.Class
	o.Flags = op.Flags
	o.Context = op.Context
	o.ClientRefresherVersion = op.ClientRefresherVersion
}
func (o *AddEnumToRefresherRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddEnumToRefresherRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddEnumToRefresherOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddEnumToRefresherResponse structure represents the AddEnumToRefresher operation response
type AddEnumToRefresherResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pInfo: MUST be an output parameter that returns a _WBEM_REFRESH_INFO structure, as
	// specified in section 2.2.20, which MUST contain refresher information about the CIM
	// instance in wszPath. This parameter MUST NOT be NULL.
	Info *wmi.RefreshInfo `idl:"name:pInfo" json:"info"`
	// pdwSvrRefrVersion: MUST be an output parameter, which MUST be the version of the
	// server refresher. The value of this parameter SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The AddEnumToRefresher return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddEnumToRefresherResponse) xxx_ToOp(ctx context.Context) *xxx_AddEnumToRefresherOperation {
	if o == nil {
		return &xxx_AddEnumToRefresherOperation{}
	}
	return &xxx_AddEnumToRefresherOperation{
		That:                   o.That,
		Info:                   o.Info,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *AddEnumToRefresherResponse) xxx_FromOp(ctx context.Context, op *xxx_AddEnumToRefresherOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Info = op.Info
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *AddEnumToRefresherResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddEnumToRefresherResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddEnumToRefresherOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveObjectFromRefresherOperation structure represents the RemoveObjectFromRefresher operation
type xxx_RemoveObjectFromRefresherOperation struct {
	This                   *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	ID                     int32            `idl:"name:lId" json:"id"`
	Flags                  int32            `idl:"name:lFlags" json:"flags"`
	ClientRefresherVersion uint32           `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	ServerRefresherVersion uint32           `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveObjectFromRefresherOperation) OpNum() int { return 6 }

func (o *xxx_RemoveObjectFromRefresherOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/RemoveObjectFromRefresher"
}

func (o *xxx_RemoveObjectFromRefresherOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveObjectFromRefresherOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lId {in} (1:(int32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveObjectFromRefresherOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lId {in} (1:(int32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveObjectFromRefresherOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveObjectFromRefresherOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_RemoveObjectFromRefresherOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// RemoveObjectFromRefresherRequest structure represents the RemoveObjectFromRefresher operation request
type RemoveObjectFromRefresherRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, that identifies the client that is requesting refreshing services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// lId: This parameter MUST be an identifier to the object that is being removed. This
	// parameter MUST NOT be NULL.
	ID int32 `idl:"name:lId" json:"id"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<67>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
}

func (o *RemoveObjectFromRefresherRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveObjectFromRefresherOperation {
	if o == nil {
		return &xxx_RemoveObjectFromRefresherOperation{}
	}
	return &xxx_RemoveObjectFromRefresherOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		ID:                     o.ID,
		Flags:                  o.Flags,
		ClientRefresherVersion: o.ClientRefresherVersion,
	}
}

func (o *RemoveObjectFromRefresherRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveObjectFromRefresherOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.ID = op.ID
	o.Flags = op.Flags
	o.ClientRefresherVersion = op.ClientRefresherVersion
}
func (o *RemoveObjectFromRefresherRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveObjectFromRefresherRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveObjectFromRefresherOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveObjectFromRefresherResponse structure represents the RemoveObjectFromRefresher operation response
type RemoveObjectFromRefresherResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwSvrRefrVersion: MUST be an output parameter, which MUST be the version of the
	// server refresher. This value SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The RemoveObjectFromRefresher return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveObjectFromRefresherResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveObjectFromRefresherOperation {
	if o == nil {
		return &xxx_RemoveObjectFromRefresherOperation{}
	}
	return &xxx_RemoveObjectFromRefresherOperation{
		That:                   o.That,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *RemoveObjectFromRefresherResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveObjectFromRefresherOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *RemoveObjectFromRefresherResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveObjectFromRefresherResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveObjectFromRefresherOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRemoteRefresherOperation structure represents the GetRemoteRefresher operation
type xxx_GetRemoteRefresherOperation struct {
	This                   *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat       `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID     `idl:"name:pRefresherId" json:"refresher_id"`
	Flags                  int32                `idl:"name:lFlags" json:"flags"`
	ClientRefresherVersion uint32               `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	RemoteRefresher        *wmi.RemoteRefresher `idl:"name:ppRemRefresher" json:"remote_refresher"`
	GUID                   *dtyp.GUID           `idl:"name:pGuid" json:"guid"`
	ServerRefresherVersion uint32               `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRemoteRefresherOperation) OpNum() int { return 7 }

func (o *xxx_GetRemoteRefresherOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/GetRemoteRefresher"
}

func (o *xxx_GetRemoteRefresherOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteRefresherOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
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
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteRefresherOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
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
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteRefresherOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRemoteRefresherOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppRemRefresher {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemRemoteRefresher}(interface))
	{
		if o.RemoteRefresher != nil {
			_ptr_ppRemRefresher := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RemoteRefresher != nil {
					if err := o.RemoteRefresher.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&wmi.RemoteRefresher{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RemoteRefresher, _ptr_ppRemRefresher); err != nil {
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
	// pGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_GetRemoteRefresherOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppRemRefresher {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IWbemRemoteRefresher}(interface))
	{
		_ptr_ppRemRefresher := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RemoteRefresher == nil {
				o.RemoteRefresher = &wmi.RemoteRefresher{}
			}
			if err := o.RemoteRefresher.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRemRefresher := func(ptr interface{}) { o.RemoteRefresher = *ptr.(**wmi.RemoteRefresher) }
		if err := w.ReadPointer(&o.RemoteRefresher, _s_ppRemRefresher, _ptr_ppRemRefresher); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pGuid {out} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// GetRemoteRefresherRequest structure represents the GetRemoteRefresher operation request
type GetRemoteRefresherRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, that identifies the client that is requesting refreshing services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<69>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
}

func (o *GetRemoteRefresherRequest) xxx_ToOp(ctx context.Context) *xxx_GetRemoteRefresherOperation {
	if o == nil {
		return &xxx_GetRemoteRefresherOperation{}
	}
	return &xxx_GetRemoteRefresherOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		Flags:                  o.Flags,
		ClientRefresherVersion: o.ClientRefresherVersion,
	}
}

func (o *GetRemoteRefresherRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteRefresherOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.Flags = op.Flags
	o.ClientRefresherVersion = op.ClientRefresherVersion
}
func (o *GetRemoteRefresherRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRemoteRefresherRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteRefresherOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRemoteRefresherResponse structure represents the GetRemoteRefresher operation response
type GetRemoteRefresherResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppRemRefresher: MUST be a pointer to an IWbemRemoteRefresher interface pointer that
	// the client can use to call the IWbemRemoteRefresher::RemoteRefresh method to refresh
	// CIM instances and enumerations. This parameter MUST NOT be NULL.
	RemoteRefresher *wmi.RemoteRefresher `idl:"name:ppRemRefresher" json:"remote_refresher"`
	// pGuid: MUST be an output parameter that MUST be a pointer to a GUID value that MUST
	// identify the returned refresher object. This parameter MUST NOT be NULL.
	GUID *dtyp.GUID `idl:"name:pGuid" json:"guid"`
	// pdwSvrRefrVersion: MUST be an output parameter that MUST be the version of the server
	// refresher. The value of this parameter SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The GetRemoteRefresher return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRemoteRefresherResponse) xxx_ToOp(ctx context.Context) *xxx_GetRemoteRefresherOperation {
	if o == nil {
		return &xxx_GetRemoteRefresherOperation{}
	}
	return &xxx_GetRemoteRefresherOperation{
		That:                   o.That,
		RemoteRefresher:        o.RemoteRefresher,
		GUID:                   o.GUID,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *GetRemoteRefresherResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRemoteRefresherOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RemoteRefresher = op.RemoteRefresher
	o.GUID = op.GUID
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *GetRemoteRefresherResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRemoteRefresherResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRemoteRefresherOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReconnectRemoteRefresherOperation structure represents the ReconnectRemoteRefresher operation
type xxx_ReconnectRemoteRefresherOperation struct {
	This                   *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat          `idl:"name:That" json:"that"`
	RefresherID            *wmi.RefresherID        `idl:"name:pRefresherId" json:"refresher_id"`
	Flags                  int32                   `idl:"name:lFlags" json:"flags"`
	ObjectsLength          int32                   `idl:"name:lNumObjects" json:"objects_length"`
	ClientRefresherVersion uint32                  `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	ReconnectInfo          []*wmi.ReconnectInfo    `idl:"name:apReconnectInfo;size_is:(lNumObjects)" json:"reconnect_info"`
	ReconnectResults       []*wmi.ReconnectResults `idl:"name:apReconnectResults;size_is:(lNumObjects)" json:"reconnect_results"`
	ServerRefresherVersion uint32                  `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	Return                 int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_ReconnectRemoteRefresherOperation) OpNum() int { return 8 }

func (o *xxx_ReconnectRemoteRefresherOperation) OpName() string {
	return "/IWbemRefreshingServices/v0/ReconnectRemoteRefresher"
}

func (o *xxx_ReconnectRemoteRefresherOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ReconnectInfo != nil && o.ObjectsLength == 0 {
		o.ObjectsLength = int32(len(o.ReconnectInfo))
	}
	if o.ReconnectResults != nil && o.ObjectsLength == 0 {
		o.ObjectsLength = int32(len(o.ReconnectResults))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReconnectRemoteRefresherOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID != nil {
			if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&wmi.RefresherID{}).MarshalNDR(ctx, w); err != nil {
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
	// lNumObjects {in} (1:(int32))
	{
		if err := w.WriteData(o.ObjectsLength); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	// apReconnectInfo {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_INFO}[dim:0,size_is=lNumObjects](struct))
	{
		dimSize1 := uint64(o.ObjectsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ReconnectInfo {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ReconnectInfo[i1] != nil {
				if err := o.ReconnectInfo[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&wmi.ReconnectInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ReconnectInfo); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&wmi.ReconnectInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// apReconnectResults {in, out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_RESULTS}[dim:0,size_is=lNumObjects](struct))
	{
		dimSize1 := uint64(o.ObjectsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ReconnectResults {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ReconnectResults[i1] != nil {
				if err := o.ReconnectResults[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&wmi.ReconnectResults{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ReconnectResults); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&wmi.ReconnectResults{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReconnectRemoteRefresherOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pRefresherId {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_REFRESHER_ID}(struct))
	{
		if o.RefresherID == nil {
			o.RefresherID = &wmi.RefresherID{}
		}
		if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
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
	// lNumObjects {in} (1:(int32))
	{
		if err := w.ReadData(&o.ObjectsLength); err != nil {
			return err
		}
	}
	// dwClientRefrVersion {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientRefresherVersion); err != nil {
			return err
		}
	}
	// apReconnectInfo {in} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_INFO}[dim:0,size_is=lNumObjects](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ReconnectInfo", sizeInfo[0])
		}
		o.ReconnectInfo = make([]*wmi.ReconnectInfo, sizeInfo[0])
		for i1 := range o.ReconnectInfo {
			i1 := i1
			if o.ReconnectInfo[i1] == nil {
				o.ReconnectInfo[i1] = &wmi.ReconnectInfo{}
			}
			if err := o.ReconnectInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// apReconnectResults {in, out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_RESULTS}[dim:0,size_is=lNumObjects](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ReconnectResults", sizeInfo[0])
		}
		o.ReconnectResults = make([]*wmi.ReconnectResults, sizeInfo[0])
		for i1 := range o.ReconnectResults {
			i1 := i1
			if o.ReconnectResults[i1] == nil {
				o.ReconnectResults[i1] = &wmi.ReconnectResults{}
			}
			if err := o.ReconnectResults[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReconnectRemoteRefresherOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReconnectRemoteRefresherOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// apReconnectResults {in, out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_RESULTS}[dim:0,size_is=lNumObjects](struct))
	{
		dimSize1 := uint64(o.ObjectsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ReconnectResults {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ReconnectResults[i1] != nil {
				if err := o.ReconnectResults[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&wmi.ReconnectResults{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ReconnectResults); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&wmi.ReconnectResults{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerRefresherVersion); err != nil {
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

func (o *xxx_ReconnectRemoteRefresherOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// apReconnectResults {in, out} (1:{pointer=ref}*(1))(2:{alias=_WBEM_RECONNECT_RESULTS}[dim:0,size_is=lNumObjects](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ReconnectResults", sizeInfo[0])
		}
		o.ReconnectResults = make([]*wmi.ReconnectResults, sizeInfo[0])
		for i1 := range o.ReconnectResults {
			i1 := i1
			if o.ReconnectResults[i1] == nil {
				o.ReconnectResults[i1] = &wmi.ReconnectResults{}
			}
			if err := o.ReconnectResults[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pdwSvrRefrVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerRefresherVersion); err != nil {
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

// ReconnectRemoteRefresherRequest structure represents the ReconnectRemoteRefresher operation request
type ReconnectRemoteRefresherRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pRefresherId: MUST be a pointer to the _WBEM_REFRESHER_ID structure, as specified
	// in section 2.2.21, which identifies the client that is requesting refresh services.
	// This parameter MUST NOT be NULL.
	RefresherID *wmi.RefresherID `idl:"name:pRefresherId" json:"refresher_id"`
	// lFlags: This parameter is not used, and its value SHOULD be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
	// lNumObjects: MUST be the number of CIM instances that are contained in the apReconnectInfo
	// array.
	ObjectsLength int32 `idl:"name:lNumObjects" json:"objects_length"`
	// dwClientRefrVersion: MUST be the version of the client refresher. This value SHOULD<70>
	// be 0x2. The server MUST allow all client versions.
	ClientRefresherVersion uint32 `idl:"name:dwClientRefrVersion" json:"client_refresher_version"`
	// apReconnectInfo: MUST be a pointer to the _WBEM_RECONNECT_INFO structure array (specified
	// in section 2.2.22) that contains a type and a CIM path to the refresher objects.
	// This parameter MUST NOT be NULL.
	ReconnectInfo []*wmi.ReconnectInfo `idl:"name:apReconnectInfo;size_is:(lNumObjects)" json:"reconnect_info"`
	// apReconnectResults: MUST be a pointer to the _WBEM_RECONNECT_RESULTS structure array,
	// which MUST contain the identifier for each CIM instance and enumeration, and the
	// success or failure status of the reconnection. This parameter MUST NOT be NULL.
	ReconnectResults []*wmi.ReconnectResults `idl:"name:apReconnectResults;size_is:(lNumObjects)" json:"reconnect_results"`
}

func (o *ReconnectRemoteRefresherRequest) xxx_ToOp(ctx context.Context) *xxx_ReconnectRemoteRefresherOperation {
	if o == nil {
		return &xxx_ReconnectRemoteRefresherOperation{}
	}
	return &xxx_ReconnectRemoteRefresherOperation{
		This:                   o.This,
		RefresherID:            o.RefresherID,
		Flags:                  o.Flags,
		ObjectsLength:          o.ObjectsLength,
		ClientRefresherVersion: o.ClientRefresherVersion,
		ReconnectInfo:          o.ReconnectInfo,
		ReconnectResults:       o.ReconnectResults,
	}
}

func (o *ReconnectRemoteRefresherRequest) xxx_FromOp(ctx context.Context, op *xxx_ReconnectRemoteRefresherOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RefresherID = op.RefresherID
	o.Flags = op.Flags
	o.ObjectsLength = op.ObjectsLength
	o.ClientRefresherVersion = op.ClientRefresherVersion
	o.ReconnectInfo = op.ReconnectInfo
	o.ReconnectResults = op.ReconnectResults
}
func (o *ReconnectRemoteRefresherRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReconnectRemoteRefresherRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReconnectRemoteRefresherOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReconnectRemoteRefresherResponse structure represents the ReconnectRemoteRefresher operation response
type ReconnectRemoteRefresherResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// apReconnectResults: MUST be a pointer to the _WBEM_RECONNECT_RESULTS structure array,
	// which MUST contain the identifier for each CIM instance and enumeration, and the
	// success or failure status of the reconnection. This parameter MUST NOT be NULL.
	ReconnectResults []*wmi.ReconnectResults `idl:"name:apReconnectResults;size_is:(lNumObjects)" json:"reconnect_results"`
	// pdwSvrRefrVersion: MUST be an output parameter that is the version of the server
	// refresher. This value SHOULD be 0x1.
	ServerRefresherVersion uint32 `idl:"name:pdwSvrRefrVersion" json:"server_refresher_version"`
	// Return: The ReconnectRemoteRefresher return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReconnectRemoteRefresherResponse) xxx_ToOp(ctx context.Context) *xxx_ReconnectRemoteRefresherOperation {
	if o == nil {
		return &xxx_ReconnectRemoteRefresherOperation{}
	}
	return &xxx_ReconnectRemoteRefresherOperation{
		That:                   o.That,
		ReconnectResults:       o.ReconnectResults,
		ServerRefresherVersion: o.ServerRefresherVersion,
		Return:                 o.Return,
	}
}

func (o *ReconnectRemoteRefresherResponse) xxx_FromOp(ctx context.Context, op *xxx_ReconnectRemoteRefresherOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReconnectResults = op.ReconnectResults
	o.ServerRefresherVersion = op.ServerRefresherVersion
	o.Return = op.Return
}
func (o *ReconnectRemoteRefresherResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReconnectRemoteRefresherResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReconnectRemoteRefresherOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
