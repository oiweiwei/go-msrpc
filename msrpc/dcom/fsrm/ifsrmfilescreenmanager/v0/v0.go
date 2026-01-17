package ifsrmfilescreenmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmFileScreenManager interface identifier ff4fa04e-5a94-4bda-a3a0-d5b4d3c52eba
	FileScreenManagerIID = &dcom.IID{Data1: 0xff4fa04e, Data2: 0x5a94, Data3: 0x4bda, Data4: []byte{0xa3, 0xa0, 0xd5, 0xb4, 0xd3, 0xc5, 0x2e, 0xba}}
	// Syntax UUID
	FileScreenManagerSyntaxUUID = &uuid.UUID{TimeLow: 0xff4fa04e, TimeMid: 0x5a94, TimeHiAndVersion: 0x4bda, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xa0, Node: [6]uint8{0xd5, 0xb4, 0xd3, 0xc5, 0x2e, 0xba}}
	// Syntax ID
	FileScreenManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: FileScreenManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmFileScreenManager interface.
type FileScreenManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// ActionVariables operation.
	GetActionVariables(context.Context, *GetActionVariablesRequest, ...dcerpc.CallOption) (*GetActionVariablesResponse, error)

	// ActionVariableDescriptions operation.
	GetActionVariableDescriptions(context.Context, *GetActionVariableDescriptionsRequest, ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error)

	// The CreateFileScreen method creates a Non-Persisted File Screen Instance (section
	// 3.2.1.3.1.2) on the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                 |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                   |
	//	|                                  |                                                                                 |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The specified path could not be found.                                          |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters. |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | The fileScreen parameter is NULL.                                               |
	//	+----------------------------------+---------------------------------------------------------------------------------+
	CreateFileScreen(context.Context, *CreateFileScreenRequest, ...dcerpc.CallOption) (*CreateFileScreenResponse, error)

	// The GetFileScreen method returns the file screen from the List of Persisted File
	// Screens (section 3.2.1.3) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The file screen for the specified path could not be found.                       |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | fileScreen parameter is NULL.                                                    |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetFileScreen(context.Context, *GetFileScreenRequest, ...dcerpc.CallOption) (*GetFileScreenResponse, error)

	// The EnumFileScreens method returns all the file screens from the List of Persisted
	// File Screens (section 3.2.1.3) that fall under the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND     | A file screen has not been applied to the specified directories.                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileScreens parameter is NULL.                                               |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)      |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumFileScreens(context.Context, *EnumFileScreensRequest, ...dcerpc.CallOption) (*EnumFileScreensResponse, error)

	// The CreateFileScreenException method creates a Non-Persisted File Screen Exception
	// Instance (section 3.2.1.3.2.2) on the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                 |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                   |
	//	|                                |                                                                                 |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH | The content of the path parameter exceeds the maximum length of 260 characters. |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The fileScreenException parameter is NULL.                                      |
	//	+--------------------------------+---------------------------------------------------------------------------------+
	CreateFileScreenException(context.Context, *CreateFileScreenExceptionRequest, ...dcerpc.CallOption) (*CreateFileScreenExceptionResponse, error)

	// The GetFileScreenException method returns the file screen exception from the List
	// of Persisted File Screen Exceptions (section 3.2.1.3) for the specified path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND      | A file screen exception has not been applied to the specified directory.         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND | The file screen exception for the specified path could not be found.             |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH   | The content of the path parameter exceeds the maximum length of 260 characters.  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG          | This code is returned for the following reasons: The path parameter is NULL. The |
	//	|                                  | fileScreenException parameter is NULL.                                           |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetFileScreenException(context.Context, *GetFileScreenExceptionRequest, ...dcerpc.CallOption) (*GetFileScreenExceptionResponse, error)

	// The EnumFileScreenExceptions method returns all the file screen exceptions from the
	// List of Persisted File Screen Exceptions (section 3.2.1.3) that fall under the specified
	// path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND     | A file screen exception has not been applied to the specified directory.         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The fileScreenExceptions parameter is NULL.                                      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | This options parameter contains invalid FsrmEnumOptions (section 2.2.1.2.5)      |
	//	|                                 | values.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumFileScreenExceptions(context.Context, *EnumFileScreenExceptionsRequest, ...dcerpc.CallOption) (*EnumFileScreenExceptionsResponse, error)

	// The CreateFileScreenCollection method creates an empty collection. This creates a
	// location where callers can add file screens.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-----------------------------------+
	//	|         RETURN          |                                   |
	//	|       VALUE/CODE        |            DESCRIPTION            |
	//	|                         |                                   |
	//	+-------------------------+-----------------------------------+
	//	+-------------------------+-----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The collection parameter is NULL. |
	//	+-------------------------+-----------------------------------+
	CreateFileScreenCollection(context.Context, *CreateFileScreenCollectionRequest, ...dcerpc.CallOption) (*CreateFileScreenCollectionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) FileScreenManagerClient
}

type xxx_DefaultFileScreenManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultFileScreenManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultFileScreenManagerClient) GetActionVariables(ctx context.Context, in *GetActionVariablesRequest, opts ...dcerpc.CallOption) (*GetActionVariablesResponse, error) {
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
	out := &GetActionVariablesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) GetActionVariableDescriptions(ctx context.Context, in *GetActionVariableDescriptionsRequest, opts ...dcerpc.CallOption) (*GetActionVariableDescriptionsResponse, error) {
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
	out := &GetActionVariableDescriptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) CreateFileScreen(ctx context.Context, in *CreateFileScreenRequest, opts ...dcerpc.CallOption) (*CreateFileScreenResponse, error) {
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
	out := &CreateFileScreenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) GetFileScreen(ctx context.Context, in *GetFileScreenRequest, opts ...dcerpc.CallOption) (*GetFileScreenResponse, error) {
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
	out := &GetFileScreenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) EnumFileScreens(ctx context.Context, in *EnumFileScreensRequest, opts ...dcerpc.CallOption) (*EnumFileScreensResponse, error) {
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
	out := &EnumFileScreensResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) CreateFileScreenException(ctx context.Context, in *CreateFileScreenExceptionRequest, opts ...dcerpc.CallOption) (*CreateFileScreenExceptionResponse, error) {
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
	out := &CreateFileScreenExceptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) GetFileScreenException(ctx context.Context, in *GetFileScreenExceptionRequest, opts ...dcerpc.CallOption) (*GetFileScreenExceptionResponse, error) {
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
	out := &GetFileScreenExceptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) EnumFileScreenExceptions(ctx context.Context, in *EnumFileScreenExceptionsRequest, opts ...dcerpc.CallOption) (*EnumFileScreenExceptionsResponse, error) {
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
	out := &EnumFileScreenExceptionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) CreateFileScreenCollection(ctx context.Context, in *CreateFileScreenCollectionRequest, opts ...dcerpc.CallOption) (*CreateFileScreenCollectionResponse, error) {
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
	out := &CreateFileScreenCollectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFileScreenManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFileScreenManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultFileScreenManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) FileScreenManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultFileScreenManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewFileScreenManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FileScreenManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FileScreenManagerSyntaxV0_0))...)
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
	return &xxx_DefaultFileScreenManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetActionVariablesOperation structure represents the ActionVariables operation
type xxx_GetActionVariablesOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariablesOperation) OpNum() int { return 7 }

func (o *xxx_GetActionVariablesOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/ActionVariables"
}

func (o *xxx_GetActionVariablesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariablesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariablesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Variables != nil {
			_ptr_variables := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Variables != nil {
					if err := o.Variables.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Variables, _ptr_variables); err != nil {
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

func (o *xxx_GetActionVariablesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// variables {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_variables := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Variables == nil {
				o.Variables = &oaut.SafeArray{}
			}
			if err := o.Variables.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_variables := func(ptr interface{}) { o.Variables = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Variables, _s_variables, _ptr_variables); err != nil {
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

// GetActionVariablesRequest structure represents the ActionVariables operation request
type GetActionVariablesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariablesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariablesOperation) *xxx_GetActionVariablesOperation {
	if op == nil {
		op = &xxx_GetActionVariablesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetActionVariablesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariablesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariablesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariablesResponse structure represents the ActionVariables operation response
type GetActionVariablesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Variables *oaut.SafeArray `idl:"name:variables" json:"variables"`
	// Return: The ActionVariables return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariablesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariablesOperation) *xxx_GetActionVariablesOperation {
	if op == nil {
		op = &xxx_GetActionVariablesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Variables = o.Variables
	op.Return = o.Return
	return op
}

func (o *GetActionVariablesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariablesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Variables = op.Variables
	o.Return = op.Return
}
func (o *GetActionVariablesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariablesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariablesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetActionVariableDescriptionsOperation structure represents the ActionVariableDescriptions operation
type xxx_GetActionVariableDescriptionsOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetActionVariableDescriptionsOperation) OpNum() int { return 8 }

func (o *xxx_GetActionVariableDescriptionsOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/ActionVariableDescriptions"
}

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetActionVariableDescriptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetActionVariableDescriptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Descriptions != nil {
			_ptr_descriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Descriptions != nil {
					if err := o.Descriptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_descriptions); err != nil {
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

func (o *xxx_GetActionVariableDescriptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// descriptions {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_descriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Descriptions == nil {
				o.Descriptions = &oaut.SafeArray{}
			}
			if err := o.Descriptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_descriptions := func(ptr interface{}) { o.Descriptions = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Descriptions, _s_descriptions, _ptr_descriptions); err != nil {
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

// GetActionVariableDescriptionsRequest structure represents the ActionVariableDescriptions operation request
type GetActionVariableDescriptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetActionVariableDescriptionsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) *xxx_GetActionVariableDescriptionsOperation {
	if op == nil {
		op = &xxx_GetActionVariableDescriptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetActionVariableDescriptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetActionVariableDescriptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetActionVariableDescriptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetActionVariableDescriptionsResponse structure represents the ActionVariableDescriptions operation response
type GetActionVariableDescriptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Descriptions *oaut.SafeArray `idl:"name:descriptions" json:"descriptions"`
	// Return: The ActionVariableDescriptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetActionVariableDescriptionsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) *xxx_GetActionVariableDescriptionsOperation {
	if op == nil {
		op = &xxx_GetActionVariableDescriptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Descriptions = o.Descriptions
	op.Return = o.Return
	return op
}

func (o *GetActionVariableDescriptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetActionVariableDescriptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Descriptions = op.Descriptions
	o.Return = op.Return
}
func (o *GetActionVariableDescriptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetActionVariableDescriptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetActionVariableDescriptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateFileScreenOperation structure represents the CreateFileScreen operation
type xxx_CreateFileScreenOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Path       *oaut.String     `idl:"name:path" json:"path"`
	FileScreen *fsrm.FileScreen `idl:"name:fileScreen" json:"file_screen"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFileScreenOperation) OpNum() int { return 9 }

func (o *xxx_CreateFileScreenOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/CreateFileScreen"
}

func (o *xxx_CreateFileScreenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_CreateFileScreenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreen {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreen}(interface))
	{
		if o.FileScreen != nil {
			_ptr_fileScreen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreen != nil {
					if err := o.FileScreen.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileScreen{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreen, _ptr_fileScreen); err != nil {
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

func (o *xxx_CreateFileScreenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreen {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreen}(interface))
	{
		_ptr_fileScreen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreen == nil {
				o.FileScreen = &fsrm.FileScreen{}
			}
			if err := o.FileScreen.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreen := func(ptr interface{}) { o.FileScreen = *ptr.(**fsrm.FileScreen) }
		if err := w.ReadPointer(&o.FileScreen, _s_fileScreen, _ptr_fileScreen); err != nil {
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

// CreateFileScreenRequest structure represents the CreateFileScreen operation request
type CreateFileScreenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path of the local directory to put the file screen on.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *CreateFileScreenRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenOperation) *xxx_CreateFileScreenOperation {
	if op == nil {
		op = &xxx_CreateFileScreenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	return op
}

func (o *CreateFileScreenRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *CreateFileScreenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateFileScreenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFileScreenResponse structure represents the CreateFileScreen operation response
type CreateFileScreenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreen: Pointer to an IFsrmFileScreen interface pointer (section 3.2.4.2.27)
	// that upon completion contains a pointer to the newly created file screen. To have
	// the file screen added to the server's List of Persisted File Screens (section 3.2.1.3),
	// the caller MUST call Commit (section 3.2.4.2.27.1). The caller MUST release the file
	// screen when it is done with it.
	FileScreen *fsrm.FileScreen `idl:"name:fileScreen" json:"file_screen"`
	// Return: The CreateFileScreen return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFileScreenResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenOperation) *xxx_CreateFileScreenOperation {
	if op == nil {
		op = &xxx_CreateFileScreenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreen = o.FileScreen
	op.Return = o.Return
	return op
}

func (o *CreateFileScreenResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreen = op.FileScreen
	o.Return = op.Return
}
func (o *CreateFileScreenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateFileScreenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileScreenOperation structure represents the GetFileScreen operation
type xxx_GetFileScreenOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Path       *oaut.String     `idl:"name:path" json:"path"`
	FileScreen *fsrm.FileScreen `idl:"name:fileScreen" json:"file_screen"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileScreenOperation) OpNum() int { return 10 }

func (o *xxx_GetFileScreenOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/GetFileScreen"
}

func (o *xxx_GetFileScreenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetFileScreenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreen {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreen}(interface))
	{
		if o.FileScreen != nil {
			_ptr_fileScreen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreen != nil {
					if err := o.FileScreen.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileScreen{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreen, _ptr_fileScreen); err != nil {
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

func (o *xxx_GetFileScreenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreen {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreen}(interface))
	{
		_ptr_fileScreen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreen == nil {
				o.FileScreen = &fsrm.FileScreen{}
			}
			if err := o.FileScreen.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreen := func(ptr interface{}) { o.FileScreen = *ptr.(**fsrm.FileScreen) }
		if err := w.ReadPointer(&o.FileScreen, _s_fileScreen, _ptr_fileScreen); err != nil {
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

// GetFileScreenRequest structure represents the GetFileScreen operation request
type GetFileScreenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path of the file screen to return.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *GetFileScreenRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenOperation) *xxx_GetFileScreenOperation {
	if op == nil {
		op = &xxx_GetFileScreenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	return op
}

func (o *GetFileScreenRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *GetFileScreenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileScreenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileScreenResponse structure represents the GetFileScreen operation response
type GetFileScreenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreen: Pointer to an IFsrmFileScreen interface pointer (section 3.2.4.2.27)
	// that upon completion contains a pointer to the file screen for the specified path.
	// The caller MUST release the file screen when it is done with it.
	FileScreen *fsrm.FileScreen `idl:"name:fileScreen" json:"file_screen"`
	// Return: The GetFileScreen return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileScreenResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenOperation) *xxx_GetFileScreenOperation {
	if op == nil {
		op = &xxx_GetFileScreenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreen = o.FileScreen
	op.Return = o.Return
	return op
}

func (o *GetFileScreenResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreen = op.FileScreen
	o.Return = op.Return
}
func (o *GetFileScreenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileScreenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFileScreensOperation structure represents the EnumFileScreens operation
type xxx_EnumFileScreensOperation struct {
	This        *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Path        *oaut.String                `idl:"name:path" json:"path"`
	Options     fsrm.EnumOptions            `idl:"name:options" json:"options"`
	FileScreens *fsrm.CommittableCollection `idl:"name:fileScreens" json:"file_screens"`
	Return      int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFileScreensOperation) OpNum() int { return 11 }

func (o *xxx_EnumFileScreensOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/EnumFileScreens"
}

func (o *xxx_EnumFileScreensOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreensOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreensOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreensOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreensOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreens {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.FileScreens != nil {
			_ptr_fileScreens := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreens != nil {
					if err := o.FileScreens.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreens, _ptr_fileScreens); err != nil {
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

func (o *xxx_EnumFileScreensOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreens {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_fileScreens := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreens == nil {
				o.FileScreens = &fsrm.CommittableCollection{}
			}
			if err := o.FileScreens.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreens := func(ptr interface{}) { o.FileScreens = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.FileScreens, _s_fileScreens, _ptr_fileScreens); err != nil {
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

// EnumFileScreensRequest structure represents the EnumFileScreens operation request
type EnumFileScreensRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path for which to limit the return of file screens.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the file screens.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumFileScreensRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumFileScreensOperation) *xxx_EnumFileScreensOperation {
	if op == nil {
		op = &xxx_EnumFileScreensOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	op.Options = o.Options
	return op
}

func (o *EnumFileScreensRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFileScreensOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *EnumFileScreensRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumFileScreensRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileScreensOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFileScreensResponse structure represents the EnumFileScreens operation response
type EnumFileScreensResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreens: Pointer to an IFsrmCommittableCollection interface pointer (section
	// 3.2.4.2.3) that upon completion contains pointers to every file screen belonging
	// to a path related to the path specified by the wildcards entered in path. The caller
	// MUST release the collection when the caller is done with it.
	FileScreens *fsrm.CommittableCollection `idl:"name:fileScreens" json:"file_screens"`
	// Return: The EnumFileScreens return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFileScreensResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumFileScreensOperation) *xxx_EnumFileScreensOperation {
	if op == nil {
		op = &xxx_EnumFileScreensOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreens = o.FileScreens
	op.Return = o.Return
	return op
}

func (o *EnumFileScreensResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFileScreensOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreens = op.FileScreens
	o.Return = op.Return
}
func (o *EnumFileScreensResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumFileScreensResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileScreensOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateFileScreenExceptionOperation structure represents the CreateFileScreenException operation
type xxx_CreateFileScreenExceptionOperation struct {
	This                *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Path                *oaut.String              `idl:"name:path" json:"path"`
	FileScreenException *fsrm.FileScreenException `idl:"name:fileScreenException" json:"file_screen_exception"`
	Return              int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFileScreenExceptionOperation) OpNum() int { return 12 }

func (o *xxx_CreateFileScreenExceptionOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/CreateFileScreenException"
}

func (o *xxx_CreateFileScreenExceptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenExceptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_CreateFileScreenExceptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenExceptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenExceptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreenException {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreenException}(interface))
	{
		if o.FileScreenException != nil {
			_ptr_fileScreenException := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreenException != nil {
					if err := o.FileScreenException.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileScreenException{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreenException, _ptr_fileScreenException); err != nil {
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

func (o *xxx_CreateFileScreenExceptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreenException {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreenException}(interface))
	{
		_ptr_fileScreenException := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreenException == nil {
				o.FileScreenException = &fsrm.FileScreenException{}
			}
			if err := o.FileScreenException.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreenException := func(ptr interface{}) { o.FileScreenException = *ptr.(**fsrm.FileScreenException) }
		if err := w.ReadPointer(&o.FileScreenException, _s_fileScreenException, _ptr_fileScreenException); err != nil {
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

// CreateFileScreenExceptionRequest structure represents the CreateFileScreenException operation request
type CreateFileScreenExceptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path of the directory to put the file screen exception on.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *CreateFileScreenExceptionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenExceptionOperation) *xxx_CreateFileScreenExceptionOperation {
	if op == nil {
		op = &xxx_CreateFileScreenExceptionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	return op
}

func (o *CreateFileScreenExceptionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenExceptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *CreateFileScreenExceptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateFileScreenExceptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenExceptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFileScreenExceptionResponse structure represents the CreateFileScreenException operation response
type CreateFileScreenExceptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreenException: Pointer to an IFsrmFileScreenException interface pointer (section
	// 3.2.4.2.28) that upon completion contains a pointer to the newly created file screen
	// exception. To have the file screen exception added to the server's List of Persisted
	// File Screen Exceptions (section 3.2.1.3), the caller MUST call Commit (section 3.2.4.2.28.1).
	// The caller MUST release the file screen exception when it is done with it.
	FileScreenException *fsrm.FileScreenException `idl:"name:fileScreenException" json:"file_screen_exception"`
	// Return: The CreateFileScreenException return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFileScreenExceptionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenExceptionOperation) *xxx_CreateFileScreenExceptionOperation {
	if op == nil {
		op = &xxx_CreateFileScreenExceptionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreenException = o.FileScreenException
	op.Return = o.Return
	return op
}

func (o *CreateFileScreenExceptionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenExceptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreenException = op.FileScreenException
	o.Return = op.Return
}
func (o *CreateFileScreenExceptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateFileScreenExceptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenExceptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFileScreenExceptionOperation structure represents the GetFileScreenException operation
type xxx_GetFileScreenExceptionOperation struct {
	This                *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Path                *oaut.String              `idl:"name:path" json:"path"`
	FileScreenException *fsrm.FileScreenException `idl:"name:fileScreenException" json:"file_screen_exception"`
	Return              int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileScreenExceptionOperation) OpNum() int { return 13 }

func (o *xxx_GetFileScreenExceptionOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/GetFileScreenException"
}

func (o *xxx_GetFileScreenExceptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenExceptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetFileScreenExceptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenExceptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileScreenExceptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreenException {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreenException}(interface))
	{
		if o.FileScreenException != nil {
			_ptr_fileScreenException := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreenException != nil {
					if err := o.FileScreenException.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.FileScreenException{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreenException, _ptr_fileScreenException); err != nil {
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

func (o *xxx_GetFileScreenExceptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreenException {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmFileScreenException}(interface))
	{
		_ptr_fileScreenException := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreenException == nil {
				o.FileScreenException = &fsrm.FileScreenException{}
			}
			if err := o.FileScreenException.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreenException := func(ptr interface{}) { o.FileScreenException = *ptr.(**fsrm.FileScreenException) }
		if err := w.ReadPointer(&o.FileScreenException, _s_fileScreenException, _ptr_fileScreenException); err != nil {
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

// GetFileScreenExceptionRequest structure represents the GetFileScreenException operation request
type GetFileScreenExceptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path of the file screen exception to return.
	Path *oaut.String `idl:"name:path" json:"path"`
}

func (o *GetFileScreenExceptionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenExceptionOperation) *xxx_GetFileScreenExceptionOperation {
	if op == nil {
		op = &xxx_GetFileScreenExceptionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	return op
}

func (o *GetFileScreenExceptionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenExceptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *GetFileScreenExceptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFileScreenExceptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenExceptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileScreenExceptionResponse structure represents the GetFileScreenException operation response
type GetFileScreenExceptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreenException: Pointer to an IFsrmFileScreenException interface pointer (section
	// 3.2.4.2.28) that upon completion contains a pointer to the file screen exception
	// for the specified path. The caller MUST release the file screen exception when it
	// is done with it.
	FileScreenException *fsrm.FileScreenException `idl:"name:fileScreenException" json:"file_screen_exception"`
	// Return: The GetFileScreenException return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileScreenExceptionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFileScreenExceptionOperation) *xxx_GetFileScreenExceptionOperation {
	if op == nil {
		op = &xxx_GetFileScreenExceptionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreenException = o.FileScreenException
	op.Return = o.Return
	return op
}

func (o *GetFileScreenExceptionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileScreenExceptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreenException = op.FileScreenException
	o.Return = op.Return
}
func (o *GetFileScreenExceptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFileScreenExceptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileScreenExceptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFileScreenExceptionsOperation structure represents the EnumFileScreenExceptions operation
type xxx_EnumFileScreenExceptionsOperation struct {
	This                 *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Path                 *oaut.String                `idl:"name:path" json:"path"`
	Options              fsrm.EnumOptions            `idl:"name:options" json:"options"`
	FileScreenExceptions *fsrm.CommittableCollection `idl:"name:fileScreenExceptions" json:"file_screen_exceptions"`
	Return               int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFileScreenExceptionsOperation) OpNum() int { return 14 }

func (o *xxx_EnumFileScreenExceptionsOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/EnumFileScreenExceptions"
}

func (o *xxx_EnumFileScreenExceptionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreenExceptionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreenExceptionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// path {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreenExceptionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFileScreenExceptionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// fileScreenExceptions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.FileScreenExceptions != nil {
			_ptr_fileScreenExceptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileScreenExceptions != nil {
					if err := o.FileScreenExceptions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileScreenExceptions, _ptr_fileScreenExceptions); err != nil {
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

func (o *xxx_EnumFileScreenExceptionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// fileScreenExceptions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_fileScreenExceptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileScreenExceptions == nil {
				o.FileScreenExceptions = &fsrm.CommittableCollection{}
			}
			if err := o.FileScreenExceptions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileScreenExceptions := func(ptr interface{}) { o.FileScreenExceptions = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.FileScreenExceptions, _s_fileScreenExceptions, _ptr_fileScreenExceptions); err != nil {
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

// EnumFileScreenExceptionsRequest structure represents the EnumFileScreenExceptions operation request
type EnumFileScreenExceptionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// path: Contains the path for which to limit the return of file screen exceptions.
	// Supports wildcards.
	Path *oaut.String `idl:"name:path" json:"path"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the file screen exception.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumFileScreenExceptionsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumFileScreenExceptionsOperation) *xxx_EnumFileScreenExceptionsOperation {
	if op == nil {
		op = &xxx_EnumFileScreenExceptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Path = o.Path
	op.Options = o.Options
	return op
}

func (o *EnumFileScreenExceptionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFileScreenExceptionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Options = op.Options
}
func (o *EnumFileScreenExceptionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumFileScreenExceptionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileScreenExceptionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFileScreenExceptionsResponse structure represents the EnumFileScreenExceptions operation response
type EnumFileScreenExceptionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileScreenExceptions: Pointer to an IFsrmCommittableCollection interface pointer
	// (section 3.2.4.2.3) that upon completion contains pointers to every file screen exception
	// belonging to a path that is part of the specified path. The caller MUST release the
	// collection when the caller is done with it.
	FileScreenExceptions *fsrm.CommittableCollection `idl:"name:fileScreenExceptions" json:"file_screen_exceptions"`
	// Return: The EnumFileScreenExceptions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFileScreenExceptionsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumFileScreenExceptionsOperation) *xxx_EnumFileScreenExceptionsOperation {
	if op == nil {
		op = &xxx_EnumFileScreenExceptionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FileScreenExceptions = o.FileScreenExceptions
	op.Return = o.Return
	return op
}

func (o *EnumFileScreenExceptionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFileScreenExceptionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileScreenExceptions = op.FileScreenExceptions
	o.Return = op.Return
}
func (o *EnumFileScreenExceptionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumFileScreenExceptionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFileScreenExceptionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateFileScreenCollectionOperation structure represents the CreateFileScreenCollection operation
type xxx_CreateFileScreenCollectionOperation struct {
	This       *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Collection *fsrm.CommittableCollection `idl:"name:collection" json:"collection"`
	Return     int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateFileScreenCollectionOperation) OpNum() int { return 15 }

func (o *xxx_CreateFileScreenCollectionOperation) OpName() string {
	return "/IFsrmFileScreenManager/v0/CreateFileScreenCollection"
}

func (o *xxx_CreateFileScreenCollectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenCollectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateFileScreenCollectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateFileScreenCollectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateFileScreenCollectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.Collection != nil {
			_ptr_collection := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Collection != nil {
					if err := o.Collection.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Collection, _ptr_collection); err != nil {
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

func (o *xxx_CreateFileScreenCollectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// collection {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_collection := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Collection == nil {
				o.Collection = &fsrm.CommittableCollection{}
			}
			if err := o.Collection.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_collection := func(ptr interface{}) { o.Collection = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.Collection, _s_collection, _ptr_collection); err != nil {
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

// CreateFileScreenCollectionRequest structure represents the CreateFileScreenCollection operation request
type CreateFileScreenCollectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateFileScreenCollectionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenCollectionOperation) *xxx_CreateFileScreenCollectionOperation {
	if op == nil {
		op = &xxx_CreateFileScreenCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CreateFileScreenCollectionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenCollectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateFileScreenCollectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateFileScreenCollectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenCollectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateFileScreenCollectionResponse structure represents the CreateFileScreenCollection operation response
type CreateFileScreenCollectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// collection: Pointer to an IFsrmCommittableCollection interface pointer (section 3.2.4.2.3)
	// that, upon completion, points to an empty IFsrmCommittableCollection specific to
	// file screen objects. A caller MUST release the collection received when the caller
	// is done with it.
	Collection *fsrm.CommittableCollection `idl:"name:collection" json:"collection"`
	// Return: The CreateFileScreenCollection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateFileScreenCollectionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateFileScreenCollectionOperation) *xxx_CreateFileScreenCollectionOperation {
	if op == nil {
		op = &xxx_CreateFileScreenCollectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Collection = o.Collection
	op.Return = o.Return
	return op
}

func (o *CreateFileScreenCollectionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateFileScreenCollectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Collection = op.Collection
	o.Return = op.Return
}
func (o *CreateFileScreenCollectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateFileScreenCollectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateFileScreenCollectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
