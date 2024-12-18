package iapphostconfigmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iisa "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
	_ = iisa.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostConfigManager interface identifier 8f6d760f-f0cb-4d69-b5f6-848b33e9bdc6
	AppHostConfigManagerIID = &dcom.IID{Data1: 0x8f6d760f, Data2: 0xf0cb, Data3: 0x4d69, Data4: []byte{0xb5, 0xf6, 0x84, 0x8b, 0x33, 0xe9, 0xbd, 0xc6}}
	// Syntax UUID
	AppHostConfigManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x8f6d760f, TimeMid: 0xf0cb, TimeHiAndVersion: 0x4d69, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0xf6, Node: [6]uint8{0x84, 0x8b, 0x33, 0xe9, 0xbd, 0xc6}}
	// Syntax ID
	AppHostConfigManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostConfigManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostConfigManager interface.
type AppHostConfigManagerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetConfigFile operation.
	GetConfigFile(context.Context, *GetConfigFileRequest, ...dcerpc.CallOption) (*GetConfigFileResponse, error)

	// GetUniqueConfigPath operation.
	GetUniqueConfigPath(context.Context, *GetUniqueConfigPathRequest, ...dcerpc.CallOption) (*GetUniqueConfigPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostConfigManagerClient
}

type xxx_DefaultAppHostConfigManagerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostConfigManagerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostConfigManagerClient) GetConfigFile(ctx context.Context, in *GetConfigFileRequest, opts ...dcerpc.CallOption) (*GetConfigFileResponse, error) {
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
	out := &GetConfigFileResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigManagerClient) GetUniqueConfigPath(ctx context.Context, in *GetUniqueConfigPathRequest, opts ...dcerpc.CallOption) (*GetUniqueConfigPathResponse, error) {
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
	out := &GetUniqueConfigPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostConfigManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostConfigManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostConfigManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostConfigManagerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostConfigManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostConfigManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostConfigManagerSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostConfigManagerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetConfigFileOperation structure represents the GetConfigFile operation
type xxx_GetConfigFileOperation struct {
	This       *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat          `idl:"name:That" json:"that"`
	ConfigPath *oaut.String            `idl:"name:bstrConfigPath" json:"config_path"`
	ConfigFile *iisa.AppHostConfigFile `idl:"name:ppConfigFile" json:"config_file"`
	Return     int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigFileOperation) OpNum() int { return 3 }

func (o *xxx_GetConfigFileOperation) OpName() string {
	return "/IAppHostConfigManager/v0/GetConfigFile"
}

func (o *xxx_GetConfigFileOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigFileOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigPath != nil {
			_ptr_bstrConfigPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigPath != nil {
					if err := o.ConfigPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigPath, _ptr_bstrConfigPath); err != nil {
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

func (o *xxx_GetConfigFileOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigPath == nil {
				o.ConfigPath = &oaut.String{}
			}
			if err := o.ConfigPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigPath := func(ptr interface{}) { o.ConfigPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigPath, _s_bstrConfigPath, _ptr_bstrConfigPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigFileOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigFileOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppConfigFile {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigFile}(interface))
	{
		if o.ConfigFile != nil {
			_ptr_ppConfigFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigFile != nil {
					if err := o.ConfigFile.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostConfigFile{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigFile, _ptr_ppConfigFile); err != nil {
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

func (o *xxx_GetConfigFileOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppConfigFile {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigFile}(interface))
	{
		_ptr_ppConfigFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigFile == nil {
				o.ConfigFile = &iisa.AppHostConfigFile{}
			}
			if err := o.ConfigFile.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConfigFile := func(ptr interface{}) { o.ConfigFile = *ptr.(**iisa.AppHostConfigFile) }
		if err := w.ReadPointer(&o.ConfigFile, _s_ppConfigFile, _ptr_ppConfigFile); err != nil {
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

// GetConfigFileRequest structure represents the GetConfigFile operation request
type GetConfigFileRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConfigPath *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
}

func (o *GetConfigFileRequest) xxx_ToOp(ctx context.Context) *xxx_GetConfigFileOperation {
	if o == nil {
		return &xxx_GetConfigFileOperation{}
	}
	return &xxx_GetConfigFileOperation{
		This:       o.This,
		ConfigPath: o.ConfigPath,
	}
}

func (o *GetConfigFileRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigFileOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConfigPath = op.ConfigPath
}
func (o *GetConfigFileRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigFileRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigFileOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigFileResponse structure represents the GetConfigFile operation response
type GetConfigFileResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat          `idl:"name:That" json:"that"`
	ConfigFile *iisa.AppHostConfigFile `idl:"name:ppConfigFile" json:"config_file"`
	// Return: The GetConfigFile return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigFileResponse) xxx_ToOp(ctx context.Context) *xxx_GetConfigFileOperation {
	if o == nil {
		return &xxx_GetConfigFileOperation{}
	}
	return &xxx_GetConfigFileOperation{
		That:       o.That,
		ConfigFile: o.ConfigFile,
		Return:     o.Return,
	}
}

func (o *GetConfigFileResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigFileOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConfigFile = op.ConfigFile
	o.Return = op.Return
}
func (o *GetConfigFileResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigFileResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigFileOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUniqueConfigPathOperation structure represents the GetUniqueConfigPath operation
type xxx_GetUniqueConfigPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConfigPath *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
	UniquePath *oaut.String   `idl:"name:pbstrUniquePath" json:"unique_path"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUniqueConfigPathOperation) OpNum() int { return 4 }

func (o *xxx_GetUniqueConfigPathOperation) OpName() string {
	return "/IAppHostConfigManager/v0/GetUniqueConfigPath"
}

func (o *xxx_GetUniqueConfigPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueConfigPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigPath != nil {
			_ptr_bstrConfigPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigPath != nil {
					if err := o.ConfigPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigPath, _ptr_bstrConfigPath); err != nil {
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

func (o *xxx_GetUniqueConfigPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrConfigPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrConfigPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigPath == nil {
				o.ConfigPath = &oaut.String{}
			}
			if err := o.ConfigPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrConfigPath := func(ptr interface{}) { o.ConfigPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigPath, _s_bstrConfigPath, _ptr_bstrConfigPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueConfigPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueConfigPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrUniquePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UniquePath != nil {
			_ptr_pbstrUniquePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UniquePath != nil {
					if err := o.UniquePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UniquePath, _ptr_pbstrUniquePath); err != nil {
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

func (o *xxx_GetUniqueConfigPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrUniquePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrUniquePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UniquePath == nil {
				o.UniquePath = &oaut.String{}
			}
			if err := o.UniquePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrUniquePath := func(ptr interface{}) { o.UniquePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UniquePath, _s_pbstrUniquePath, _ptr_pbstrUniquePath); err != nil {
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

// GetUniqueConfigPathRequest structure represents the GetUniqueConfigPath operation request
type GetUniqueConfigPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	ConfigPath *oaut.String   `idl:"name:bstrConfigPath" json:"config_path"`
}

func (o *GetUniqueConfigPathRequest) xxx_ToOp(ctx context.Context) *xxx_GetUniqueConfigPathOperation {
	if o == nil {
		return &xxx_GetUniqueConfigPathOperation{}
	}
	return &xxx_GetUniqueConfigPathOperation{
		This:       o.This,
		ConfigPath: o.ConfigPath,
	}
}

func (o *GetUniqueConfigPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueConfigPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ConfigPath = op.ConfigPath
}
func (o *GetUniqueConfigPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUniqueConfigPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueConfigPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUniqueConfigPathResponse structure represents the GetUniqueConfigPath operation response
type GetUniqueConfigPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	UniquePath *oaut.String   `idl:"name:pbstrUniquePath" json:"unique_path"`
	// Return: The GetUniqueConfigPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUniqueConfigPathResponse) xxx_ToOp(ctx context.Context) *xxx_GetUniqueConfigPathOperation {
	if o == nil {
		return &xxx_GetUniqueConfigPathOperation{}
	}
	return &xxx_GetUniqueConfigPathOperation{
		That:       o.That,
		UniquePath: o.UniquePath,
		Return:     o.Return,
	}
}

func (o *GetUniqueConfigPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueConfigPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UniquePath = op.UniquePath
	o.Return = op.Return
}
func (o *GetUniqueConfigPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUniqueConfigPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueConfigPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
