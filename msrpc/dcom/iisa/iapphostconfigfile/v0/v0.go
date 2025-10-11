package iapphostconfigfile

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
	// IAppHostConfigFile interface identifier ada4e6fb-e025-401e-a5d0-c3134a281f07
	AppHostConfigFileIID = &dcom.IID{Data1: 0xada4e6fb, Data2: 0xe025, Data3: 0x401e, Data4: []byte{0xa5, 0xd0, 0xc3, 0x13, 0x4a, 0x28, 0x1f, 0x07}}
	// Syntax UUID
	AppHostConfigFileSyntaxUUID = &uuid.UUID{TimeLow: 0xada4e6fb, TimeMid: 0xe025, TimeHiAndVersion: 0x401e, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0xd0, Node: [6]uint8{0xc3, 0x13, 0x4a, 0x28, 0x1f, 0x7}}
	// Syntax ID
	AppHostConfigFileSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostConfigFileSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostConfigFile interface.
type AppHostConfigFileClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// ConfigPath operation.
	GetConfigPath(context.Context, *GetConfigPathRequest, ...dcerpc.CallOption) (*GetConfigPathResponse, error)

	// The FilePath method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the server operating system file path that corresponds to the
	// specific IAppHostConfigFile, if that path applies to the administration system implementation.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *pbstrFilePath is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
	GetFilePath(context.Context, *GetFilePathRequest, ...dcerpc.CallOption) (*GetFilePathResponse, error)

	// The Locations method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns any hierarchy subpaths that exist in the specified IAppHostConfigFile.
	// These subpaths are returned in the form of a collection object that implements the
	// IAppHostConfigLocationCollection.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppLocations is not NULL. If processing
	// fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF]. The
	// following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	|               RETURN               |                                                                         |
	//	|             VALUE/CODE             |                               DESCRIPTION                               |
	//	|                                    |                                                                         |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                                   |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.                           |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command.                 |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070002 ERROR_FILE_NOT_FOUND    | A server resource (for example, a file on a disk) could not be found.   |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070005 ERROR_ACCESS_DENIED     | Access to a server resource (for example, a file on a disk) was denied. |
	//	+------------------------------------+-------------------------------------------------------------------------+
	//	| 0X80070013 ERROR_INVALID_DATA      | Configuration data or schema on the server are malformed or corrupted.  |
	//	+------------------------------------+-------------------------------------------------------------------------+
	GetLocations(context.Context, *GetLocationsRequest, ...dcerpc.CallOption) (*GetLocationsResponse, error)

	// GetAdminSection operation.
	GetAdminSection(context.Context, *GetAdminSectionRequest, ...dcerpc.CallOption) (*GetAdminSectionResponse, error)

	// GetMetadata operation.
	GetMetadata(context.Context, *GetMetadataRequest, ...dcerpc.CallOption) (*GetMetadataResponse, error)

	// SetMetadata operation.
	SetMetadata(context.Context, *SetMetadataRequest, ...dcerpc.CallOption) (*SetMetadataResponse, error)

	// The ClearInvalidSections method is received by the server in an RPC_REQUEST packet.
	// In response, the server MAY clear any invalid IAppHostElement objects that exist
	// in the specific IAppHostConfigFile.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result (as specified in [MS-ERREF] section 2) on failure. All failure results MUST
	// be treated identically. The following table specifies failure results specific to
	// this method.
	//
	//	+--------------------------------+-------------------------------+
	//	|             RETURN             |                               |
	//	|           VALUE/CODE           |          DESCRIPTION          |
	//	|                                |                               |
	//	+--------------------------------+-------------------------------+
	//	+--------------------------------+-------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED | The request is not supported. |
	//	+--------------------------------+-------------------------------+
	ClearInvalidSections(context.Context, *ClearInvalidSectionsRequest, ...dcerpc.CallOption) (*ClearInvalidSectionsResponse, error)

	// The RootSectionGroup method is received by the server in an RPC_REQUEST packet. In
	// response, the server returns an IAppHostSectionGroup object, which represents a declaration
	// of IAppHostElement objects that apply to the specified IAppHostConfigFile and that
	// potentially apply to other IAppHostConfigFile at deeper hierarchy paths than the
	// current file.
	//
	// A declaration means that an IAppHostElement of the name in the declaration MAY exist
	// in the specified IAppHostConfigFile or potentially in deeper paths. A declaration
	// is NOT a definition/instance. A declaration only controls whether an actual IAppHostElement
	// instance is supported.
	//
	// This function returns the section group object in the configSections section group
	// that defines the root section group for the current section.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// that is received from the client. In this case, *ppSectionGroups is not NULL. If
	// processing fails, the server MUST return a nonzero HRESULT code as defined in [MS-ERREF].
	// The following table describes the error conditions that MUST be handled and the corresponding
	// error codes. A server MAY return additional implementation-specific error codes.
	//
	//	+------------------------------------+---------------------------------------------------------+
	//	|               RETURN               |                                                         |
	//	|             VALUE/CODE             |                       DESCRIPTION                       |
	//	|                                    |                                                         |
	//	+------------------------------------+---------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000000 NO_ERROR                | The operation completed successfully.                   |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X80070057 ERROR_INVALID_PARAMETER | One or more parameters are incorrect or null.           |
	//	+------------------------------------+---------------------------------------------------------+
	//	| 0X00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough memory is available to process this command. |
	//	+------------------------------------+---------------------------------------------------------+
	GetRootSectionGroup(context.Context, *GetRootSectionGroupRequest, ...dcerpc.CallOption) (*GetRootSectionGroupResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostConfigFileClient
}

type xxx_DefaultAppHostConfigFileClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostConfigFileClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAppHostConfigFileClient) GetConfigPath(ctx context.Context, in *GetConfigPathRequest, opts ...dcerpc.CallOption) (*GetConfigPathResponse, error) {
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
	out := &GetConfigPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) GetFilePath(ctx context.Context, in *GetFilePathRequest, opts ...dcerpc.CallOption) (*GetFilePathResponse, error) {
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
	out := &GetFilePathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) GetLocations(ctx context.Context, in *GetLocationsRequest, opts ...dcerpc.CallOption) (*GetLocationsResponse, error) {
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
	out := &GetLocationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) GetAdminSection(ctx context.Context, in *GetAdminSectionRequest, opts ...dcerpc.CallOption) (*GetAdminSectionResponse, error) {
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
	out := &GetAdminSectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...dcerpc.CallOption) (*GetMetadataResponse, error) {
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
	out := &GetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...dcerpc.CallOption) (*SetMetadataResponse, error) {
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
	out := &SetMetadataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) ClearInvalidSections(ctx context.Context, in *ClearInvalidSectionsRequest, opts ...dcerpc.CallOption) (*ClearInvalidSectionsResponse, error) {
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
	out := &ClearInvalidSectionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) GetRootSectionGroup(ctx context.Context, in *GetRootSectionGroupRequest, opts ...dcerpc.CallOption) (*GetRootSectionGroupResponse, error) {
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
	out := &GetRootSectionGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostConfigFileClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostConfigFileClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostConfigFileClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostConfigFileClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostConfigFileClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAppHostConfigFileClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostConfigFileClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostConfigFileSyntaxV0_0))...)
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
	return &xxx_DefaultAppHostConfigFileClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetConfigPathOperation structure represents the ConfigPath operation
type xxx_GetConfigPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConfigPath *oaut.String   `idl:"name:pbstrConfigPath" json:"config_path"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigPathOperation) OpNum() int { return 3 }

func (o *xxx_GetConfigPathOperation) OpName() string { return "/IAppHostConfigFile/v0/ConfigPath" }

func (o *xxx_GetConfigPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetConfigPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetConfigPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrConfigPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ConfigPath != nil {
			_ptr_pbstrConfigPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.ConfigPath, _ptr_pbstrConfigPath); err != nil {
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

func (o *xxx_GetConfigPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrConfigPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrConfigPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigPath == nil {
				o.ConfigPath = &oaut.String{}
			}
			if err := o.ConfigPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrConfigPath := func(ptr interface{}) { o.ConfigPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ConfigPath, _s_pbstrConfigPath, _ptr_pbstrConfigPath); err != nil {
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

// GetConfigPathRequest structure represents the ConfigPath operation request
type GetConfigPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetConfigPathRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConfigPathOperation) *xxx_GetConfigPathOperation {
	if op == nil {
		op = &xxx_GetConfigPathOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetConfigPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetConfigPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigPathResponse structure represents the ConfigPath operation response
type GetConfigPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	ConfigPath *oaut.String   `idl:"name:pbstrConfigPath" json:"config_path"`
	// Return: The ConfigPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigPathResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConfigPathOperation) *xxx_GetConfigPathOperation {
	if op == nil {
		op = &xxx_GetConfigPathOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ConfigPath = o.ConfigPath
	op.Return = o.Return
	return op
}

func (o *GetConfigPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ConfigPath = op.ConfigPath
	o.Return = op.Return
}
func (o *GetConfigPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFilePathOperation structure represents the FilePath operation
type xxx_GetFilePathOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilePath *oaut.String   `idl:"name:pbstrFilePath" json:"file_path"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFilePathOperation) OpNum() int { return 4 }

func (o *xxx_GetFilePathOperation) OpName() string { return "/IAppHostConfigFile/v0/FilePath" }

func (o *xxx_GetFilePathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFilePathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFilePathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilePath != nil {
			_ptr_pbstrFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilePath != nil {
					if err := o.FilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilePath, _ptr_pbstrFilePath); err != nil {
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

func (o *xxx_GetFilePathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilePath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilePath == nil {
				o.FilePath = &oaut.String{}
			}
			if err := o.FilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilePath := func(ptr interface{}) { o.FilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilePath, _s_pbstrFilePath, _ptr_pbstrFilePath); err != nil {
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

// GetFilePathRequest structure represents the FilePath operation request
type GetFilePathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFilePathRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFilePathOperation) *xxx_GetFilePathOperation {
	if op == nil {
		op = &xxx_GetFilePathOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFilePathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFilePathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFilePathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFilePathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilePathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFilePathResponse structure represents the FilePath operation response
type GetFilePathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrFilePath: Contains the file path of the IAppHostConfigFile.
	FilePath *oaut.String `idl:"name:pbstrFilePath" json:"file_path"`
	// Return: The FilePath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFilePathResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFilePathOperation) *xxx_GetFilePathOperation {
	if op == nil {
		op = &xxx_GetFilePathOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FilePath = o.FilePath
	op.Return = o.Return
	return op
}

func (o *GetFilePathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFilePathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FilePath = op.FilePath
	o.Return = op.Return
}
func (o *GetFilePathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFilePathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilePathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLocationsOperation structure represents the Locations operation
type xxx_GetLocationsOperation struct {
	This      *dcom.ORPCThis                        `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat                        `idl:"name:That" json:"that"`
	Locations *iisa.AppHostConfigLocationCollection `idl:"name:ppLocations" json:"locations"`
	Return    int32                                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLocationsOperation) OpNum() int { return 5 }

func (o *xxx_GetLocationsOperation) OpName() string { return "/IAppHostConfigFile/v0/Locations" }

func (o *xxx_GetLocationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLocationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetLocationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetLocationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLocationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppLocations {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigLocationCollection}(interface))
	{
		if o.Locations != nil {
			_ptr_ppLocations := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Locations != nil {
					if err := o.Locations.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostConfigLocationCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Locations, _ptr_ppLocations); err != nil {
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

func (o *xxx_GetLocationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppLocations {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostConfigLocationCollection}(interface))
	{
		_ptr_ppLocations := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Locations == nil {
				o.Locations = &iisa.AppHostConfigLocationCollection{}
			}
			if err := o.Locations.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppLocations := func(ptr interface{}) { o.Locations = *ptr.(**iisa.AppHostConfigLocationCollection) }
		if err := w.ReadPointer(&o.Locations, _s_ppLocations, _ptr_ppLocations); err != nil {
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

// GetLocationsRequest structure represents the Locations operation request
type GetLocationsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLocationsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLocationsOperation) *xxx_GetLocationsOperation {
	if op == nil {
		op = &xxx_GetLocationsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetLocationsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLocationsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLocationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLocationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLocationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLocationsResponse structure represents the Locations operation response
type GetLocationsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppLocations: Contains a collection object that contains all the subpaths that are
	// available in the specified IAppHostConfigFile.
	Locations *iisa.AppHostConfigLocationCollection `idl:"name:ppLocations" json:"locations"`
	// Return: The Locations return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLocationsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLocationsOperation) *xxx_GetLocationsOperation {
	if op == nil {
		op = &xxx_GetLocationsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Locations = o.Locations
	op.Return = o.Return
	return op
}

func (o *GetLocationsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLocationsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Locations = op.Locations
	o.Return = op.Return
}
func (o *GetLocationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLocationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLocationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAdminSectionOperation structure represents the GetAdminSection operation
type xxx_GetAdminSectionOperation struct {
	This         *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	SectionName  *oaut.String         `idl:"name:bstrSectionName" json:"section_name"`
	Path         *oaut.String         `idl:"name:bstrPath" json:"path"`
	AdminSection *iisa.AppHostElement `idl:"name:ppAdminSection" json:"admin_section"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminSectionOperation) OpNum() int { return 6 }

func (o *xxx_GetAdminSectionOperation) OpName() string {
	return "/IAppHostConfigFile/v0/GetAdminSection"
}

func (o *xxx_GetAdminSectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SectionName != nil {
			_ptr_bstrSectionName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SectionName != nil {
					if err := o.SectionName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SectionName, _ptr_bstrSectionName); err != nil {
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
	// bstrPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_bstrPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_bstrPath); err != nil {
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

func (o *xxx_GetAdminSectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSectionName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSectionName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SectionName == nil {
				o.SectionName = &oaut.String{}
			}
			if err := o.SectionName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSectionName := func(ptr interface{}) { o.SectionName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SectionName, _s_bstrSectionName, _ptr_bstrSectionName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPath := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_bstrPath, _ptr_bstrPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminSectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAdminSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		if o.AdminSection != nil {
			_ptr_ppAdminSection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.AdminSection != nil {
					if err := o.AdminSection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostElement{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.AdminSection, _ptr_ppAdminSection); err != nil {
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

func (o *xxx_GetAdminSectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAdminSection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostElement}(interface))
	{
		_ptr_ppAdminSection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.AdminSection == nil {
				o.AdminSection = &iisa.AppHostElement{}
			}
			if err := o.AdminSection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAdminSection := func(ptr interface{}) { o.AdminSection = *ptr.(**iisa.AppHostElement) }
		if err := w.ReadPointer(&o.AdminSection, _s_ppAdminSection, _ptr_ppAdminSection); err != nil {
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

// GetAdminSectionRequest structure represents the GetAdminSection operation request
type GetAdminSectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	SectionName *oaut.String   `idl:"name:bstrSectionName" json:"section_name"`
	Path        *oaut.String   `idl:"name:bstrPath" json:"path"`
}

func (o *GetAdminSectionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAdminSectionOperation) *xxx_GetAdminSectionOperation {
	if op == nil {
		op = &xxx_GetAdminSectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SectionName = o.SectionName
	op.Path = o.Path
	return op
}

func (o *GetAdminSectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAdminSectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SectionName = op.SectionName
	o.Path = op.Path
}
func (o *GetAdminSectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminSectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminSectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminSectionResponse structure represents the GetAdminSection operation response
type GetAdminSectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat       `idl:"name:That" json:"that"`
	AdminSection *iisa.AppHostElement `idl:"name:ppAdminSection" json:"admin_section"`
	// Return: The GetAdminSection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminSectionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAdminSectionOperation) *xxx_GetAdminSectionOperation {
	if op == nil {
		op = &xxx_GetAdminSectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AdminSection = o.AdminSection
	op.Return = o.Return
	return op
}

func (o *GetAdminSectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAdminSectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AdminSection = op.AdminSection
	o.Return = op.Return
}
func (o *GetAdminSectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminSectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminSectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMetadataOperation structure represents the GetMetadata operation
type xxx_GetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:pValue" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMetadataOperation) OpNum() int { return 7 }

func (o *xxx_GetMetadataOperation) OpName() string { return "/IAppHostConfigFile/v0/GetMetadata" }

func (o *xxx_GetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			_ptr_pValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_pValue); err != nil {
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

func (o *xxx_GetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pValue {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pValue := func(ptr interface{}) { o.Value = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Value, _s_pValue, _ptr_pValue); err != nil {
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

// GetMetadataRequest structure represents the GetMetadata operation request
type GetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
}

func (o *GetMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMetadataOperation) *xxx_GetMetadataOperation {
	if op == nil {
		op = &xxx_GetMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MetadataType = o.MetadataType
	return op
}

func (o *GetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
}
func (o *GetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMetadataResponse structure represents the GetMetadata operation response
type GetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value *oaut.Variant  `idl:"name:pValue" json:"value"`
	// Return: The GetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMetadataOperation) *xxx_GetMetadataOperation {
	if op == nil {
		op = &xxx_GetMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Value = o.Value
	op.Return = o.Return
	return op
}

func (o *GetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Value = op.Value
	o.Return = op.Return
}
func (o *GetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMetadataOperation structure represents the SetMetadata operation
type xxx_SetMetadataOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:value" json:"value"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMetadataOperation) OpNum() int { return 8 }

func (o *xxx_SetMetadataOperation) OpName() string { return "/IAppHostConfigFile/v0/SetMetadata" }

func (o *xxx_SetMetadataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MetadataType != nil {
			_ptr_bstrMetadataType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MetadataType != nil {
					if err := o.MetadataType.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MetadataType, _ptr_bstrMetadataType); err != nil {
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
	// value {in} (1:{alias=VARIANT}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Value != nil {
			if err := o.Value.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMetadataType {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMetadataType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MetadataType == nil {
				o.MetadataType = &oaut.String{}
			}
			if err := o.MetadataType.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMetadataType := func(ptr interface{}) { o.MetadataType = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MetadataType, _s_bstrMetadataType, _ptr_bstrMetadataType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT,pointer=ref}*(1))(2:{alias=_VARIANT}(struct))
	{
		if o.Value == nil {
			o.Value = &oaut.Variant{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMetadataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMetadataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMetadataRequest structure represents the SetMetadata operation request
type SetMetadataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	MetadataType *oaut.String   `idl:"name:bstrMetadataType" json:"metadata_type"`
	Value        *oaut.Variant  `idl:"name:value" json:"value"`
}

func (o *SetMetadataRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMetadataOperation) *xxx_SetMetadataOperation {
	if op == nil {
		op = &xxx_SetMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MetadataType = o.MetadataType
	op.Value = o.Value
	return op
}

func (o *SetMetadataRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMetadataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MetadataType = op.MetadataType
	o.Value = op.Value
}
func (o *SetMetadataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMetadataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMetadataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMetadataResponse structure represents the SetMetadata operation response
type SetMetadataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetMetadata return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMetadataResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMetadataOperation) *xxx_SetMetadataOperation {
	if op == nil {
		op = &xxx_SetMetadataOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMetadataResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMetadataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMetadataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMetadataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMetadataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearInvalidSectionsOperation structure represents the ClearInvalidSections operation
type xxx_ClearInvalidSectionsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearInvalidSectionsOperation) OpNum() int { return 9 }

func (o *xxx_ClearInvalidSectionsOperation) OpName() string {
	return "/IAppHostConfigFile/v0/ClearInvalidSections"
}

func (o *xxx_ClearInvalidSectionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearInvalidSectionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearInvalidSectionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ClearInvalidSectionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearInvalidSectionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearInvalidSectionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearInvalidSectionsRequest structure represents the ClearInvalidSections operation request
type ClearInvalidSectionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ClearInvalidSectionsRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearInvalidSectionsOperation) *xxx_ClearInvalidSectionsOperation {
	if op == nil {
		op = &xxx_ClearInvalidSectionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ClearInvalidSectionsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearInvalidSectionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ClearInvalidSectionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearInvalidSectionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearInvalidSectionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearInvalidSectionsResponse structure represents the ClearInvalidSections operation response
type ClearInvalidSectionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearInvalidSections return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearInvalidSectionsResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearInvalidSectionsOperation) *xxx_ClearInvalidSectionsOperation {
	if op == nil {
		op = &xxx_ClearInvalidSectionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ClearInvalidSectionsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearInvalidSectionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearInvalidSectionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearInvalidSectionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearInvalidSectionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRootSectionGroupOperation structure represents the RootSectionGroup operation
type xxx_GetRootSectionGroupOperation struct {
	This          *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat            `idl:"name:That" json:"that"`
	SectionGroups *iisa.AppHostSectionGroup `idl:"name:ppSectionGroups" json:"section_groups"`
	Return        int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRootSectionGroupOperation) OpNum() int { return 10 }

func (o *xxx_GetRootSectionGroupOperation) OpName() string {
	return "/IAppHostConfigFile/v0/RootSectionGroup"
}

func (o *xxx_GetRootSectionGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootSectionGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetRootSectionGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetRootSectionGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootSectionGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppSectionGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostSectionGroup}(interface))
	{
		if o.SectionGroups != nil {
			_ptr_ppSectionGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SectionGroups != nil {
					if err := o.SectionGroups.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&iisa.AppHostSectionGroup{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SectionGroups, _ptr_ppSectionGroups); err != nil {
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

func (o *xxx_GetRootSectionGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppSectionGroups {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IAppHostSectionGroup}(interface))
	{
		_ptr_ppSectionGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SectionGroups == nil {
				o.SectionGroups = &iisa.AppHostSectionGroup{}
			}
			if err := o.SectionGroups.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSectionGroups := func(ptr interface{}) { o.SectionGroups = *ptr.(**iisa.AppHostSectionGroup) }
		if err := w.ReadPointer(&o.SectionGroups, _s_ppSectionGroups, _ptr_ppSectionGroups); err != nil {
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

// GetRootSectionGroupRequest structure represents the RootSectionGroup operation request
type GetRootSectionGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetRootSectionGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRootSectionGroupOperation) *xxx_GetRootSectionGroupOperation {
	if op == nil {
		op = &xxx_GetRootSectionGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetRootSectionGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRootSectionGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetRootSectionGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRootSectionGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootSectionGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRootSectionGroupResponse structure represents the RootSectionGroup operation response
type GetRootSectionGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppSectionGroups: Contains an IAppHostSectionGroup that contains declarations.
	SectionGroups *iisa.AppHostSectionGroup `idl:"name:ppSectionGroups" json:"section_groups"`
	// Return: The RootSectionGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRootSectionGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRootSectionGroupOperation) *xxx_GetRootSectionGroupOperation {
	if op == nil {
		op = &xxx_GetRootSectionGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SectionGroups = o.SectionGroups
	op.Return = o.Return
	return op
}

func (o *GetRootSectionGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRootSectionGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SectionGroups = op.SectionGroups
	o.Return = op.Return
}
func (o *GetRootSectionGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRootSectionGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootSectionGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
