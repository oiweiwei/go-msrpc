package ifsrmpathmapper

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmPathMapper interface identifier 6f4dbfff-6920-4821-a6c3-b7e94c1fd60c
	PathMapperIID = &dcom.IID{Data1: 0x6f4dbfff, Data2: 0x6920, Data3: 0x4821, Data4: []byte{0xa6, 0xc3, 0xb7, 0xe9, 0x4c, 0x1f, 0xd6, 0x0c}}
	// Syntax UUID
	PathMapperSyntaxUUID = &uuid.UUID{TimeLow: 0x6f4dbfff, TimeMid: 0x6920, TimeHiAndVersion: 0x4821, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xc3, Node: [6]uint8{0xb7, 0xe9, 0x4c, 0x1f, 0xd6, 0xc}}
	// Syntax ID
	PathMapperSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PathMapperSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmPathMapper interface.
type PathMapperClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// The GetSharePathsForLocalPath method returns all the network share paths that point
	// to the specified local path.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH | The local path to return network shares does not exist or exceeds the maximum    |
	//	|                                | length of 260 characters.                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | This code is returned for the following reasons: The localPath parameter is      |
	//	|                                | empty or NULL. The sharePaths parameter is NULL.                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetSharePathsForLocalPath(context.Context, *GetSharePathsForLocalPathRequest, ...dcerpc.CallOption) (*GetSharePathsForLocalPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PathMapperClient
}

type xxx_DefaultPathMapperClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPathMapperClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultPathMapperClient) GetSharePathsForLocalPath(ctx context.Context, in *GetSharePathsForLocalPathRequest, opts ...dcerpc.CallOption) (*GetSharePathsForLocalPathResponse, error) {
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
	out := &GetSharePathsForLocalPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPathMapperClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPathMapperClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPathMapperClient) IPID(ctx context.Context, ipid *dcom.IPID) PathMapperClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPathMapperClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewPathMapperClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PathMapperClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PathMapperSyntaxV0_0))...)
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
	return &xxx_DefaultPathMapperClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetSharePathsForLocalPathOperation structure represents the GetSharePathsForLocalPath operation
type xxx_GetSharePathsForLocalPathOperation struct {
	This       *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat  `idl:"name:That" json:"that"`
	LocalPath  *oaut.String    `idl:"name:localPath" json:"local_path"`
	SharePaths *oaut.SafeArray `idl:"name:sharePaths" json:"share_paths"`
	Return     int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSharePathsForLocalPathOperation) OpNum() int { return 7 }

func (o *xxx_GetSharePathsForLocalPathOperation) OpName() string {
	return "/IFsrmPathMapper/v0/GetSharePathsForLocalPath"
}

func (o *xxx_GetSharePathsForLocalPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSharePathsForLocalPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// localPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LocalPath != nil {
			_ptr_localPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LocalPath != nil {
					if err := o.LocalPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LocalPath, _ptr_localPath); err != nil {
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

func (o *xxx_GetSharePathsForLocalPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// localPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_localPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LocalPath == nil {
				o.LocalPath = &oaut.String{}
			}
			if err := o.LocalPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_localPath := func(ptr interface{}) { o.LocalPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LocalPath, _s_localPath, _ptr_localPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSharePathsForLocalPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSharePathsForLocalPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// sharePaths {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.SharePaths != nil {
			_ptr_sharePaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SharePaths != nil {
					if err := o.SharePaths.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SharePaths, _ptr_sharePaths); err != nil {
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

func (o *xxx_GetSharePathsForLocalPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// sharePaths {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_sharePaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SharePaths == nil {
				o.SharePaths = &oaut.SafeArray{}
			}
			if err := o.SharePaths.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_sharePaths := func(ptr interface{}) { o.SharePaths = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.SharePaths, _s_sharePaths, _ptr_sharePaths); err != nil {
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

// GetSharePathsForLocalPathRequest structure represents the GetSharePathsForLocalPath operation request
type GetSharePathsForLocalPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// localPath: Contains the local path for which to return network shares for.
	LocalPath *oaut.String `idl:"name:localPath" json:"local_path"`
}

func (o *GetSharePathsForLocalPathRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSharePathsForLocalPathOperation) *xxx_GetSharePathsForLocalPathOperation {
	if op == nil {
		op = &xxx_GetSharePathsForLocalPathOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LocalPath = o.LocalPath
	return op
}

func (o *GetSharePathsForLocalPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSharePathsForLocalPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LocalPath = op.LocalPath
}
func (o *GetSharePathsForLocalPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSharePathsForLocalPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSharePathsForLocalPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSharePathsForLocalPathResponse structure represents the GetSharePathsForLocalPath operation response
type GetSharePathsForLocalPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// sharePaths: Pointer to a SAFEARRAY that upon completion contains all the network
	// share paths that point to the specified path.
	SharePaths *oaut.SafeArray `idl:"name:sharePaths" json:"share_paths"`
	// Return: The GetSharePathsForLocalPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSharePathsForLocalPathResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSharePathsForLocalPathOperation) *xxx_GetSharePathsForLocalPathOperation {
	if op == nil {
		op = &xxx_GetSharePathsForLocalPathOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SharePaths = o.SharePaths
	op.Return = o.Return
	return op
}

func (o *GetSharePathsForLocalPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSharePathsForLocalPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SharePaths = op.SharePaths
	o.Return = op.Return
}
func (o *GetSharePathsForLocalPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSharePathsForLocalPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSharePathsForLocalPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
