package iclusterlogex

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csvp "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp"
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
	_ = csvp.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterLogEx interface identifier bd7c23c2-c805-457c-8f86-d17fe6b9d19f
	ClusterLogExIID = &dcom.IID{Data1: 0xbd7c23c2, Data2: 0xc805, Data3: 0x457c, Data4: []byte{0x8f, 0x86, 0xd1, 0x7f, 0xe6, 0xb9, 0xd1, 0x9f}}
	// Syntax UUID
	ClusterLogExSyntaxUUID = &uuid.UUID{TimeLow: 0xbd7c23c2, TimeMid: 0xc805, TimeHiAndVersion: 0x457c, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x86, Node: [6]uint8{0xd1, 0x7f, 0xe6, 0xb9, 0xd1, 0x9f}}
	// Syntax ID
	ClusterLogExSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterLogExSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterLogEx interface.
type ClusterLogExClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GenerateClusterLog method writes a file that contains diagnostic information
	// about failover clusters for the server on which it executes. The content and format
	// of the file are implementation-specific, but SHOULD contain diagnostic information.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	GenerateClusterLog(context.Context, *GenerateClusterLogRequest, ...dcerpc.CallOption) (*GenerateClusterLogResponse, error)

	// The GenerateClusterHealthLog method<47> generates the health log file on cluster
	// nodes. The content and format of the file is implementation-specific but SHOULD contain
	// diagnostic information.
	//
	// Return Values: Return values are the same as the return values for the GenerateClusterLog
	// method specified in section 3.12.4.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	GenerateClusterHealthLog(context.Context, *GenerateClusterHealthLogRequest, ...dcerpc.CallOption) (*GenerateClusterHealthLogResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterLogExClient
}

type xxx_DefaultClusterLogExClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterLogExClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterLogExClient) GenerateClusterLog(ctx context.Context, in *GenerateClusterLogRequest, opts ...dcerpc.CallOption) (*GenerateClusterLogResponse, error) {
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
	out := &GenerateClusterLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterLogExClient) GenerateClusterHealthLog(ctx context.Context, in *GenerateClusterHealthLogRequest, opts ...dcerpc.CallOption) (*GenerateClusterHealthLogResponse, error) {
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
	out := &GenerateClusterHealthLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterLogExClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterLogExClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterLogExClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterLogExClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterLogExClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterLogExClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterLogExClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterLogExSyntaxV0_0))...)
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
	return &xxx_DefaultClusterLogExClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GenerateClusterLogOperation structure represents the GenerateClusterLog operation
type xxx_GenerateClusterLogOperation struct {
	This        *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat        `idl:"name:That" json:"that"`
	SpanMinutes uint32                `idl:"name:SpanMinutes" json:"span_minutes"`
	Flags       csvp.ClusterLogExFlag `idl:"name:flags" json:"flags"`
	LogFilePath *oaut.String          `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterLogOperation) OpNum() int { return 3 }

func (o *xxx_GenerateClusterLogOperation) OpName() string {
	return "/IClusterLogEx/v0/GenerateClusterLog"
}

func (o *xxx_GenerateClusterLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// SpanMinutes {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SpanMinutes); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=ClusterLogExFlag}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// SpanMinutes {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SpanMinutes); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=ClusterLogExFlag}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// LogFilePath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LogFilePath != nil {
			_ptr_LogFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LogFilePath != nil {
					if err := o.LogFilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LogFilePath, _ptr_LogFilePath); err != nil {
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

func (o *xxx_GenerateClusterLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// LogFilePath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_LogFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LogFilePath == nil {
				o.LogFilePath = &oaut.String{}
			}
			if err := o.LogFilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_LogFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LogFilePath, _s_LogFilePath, _ptr_LogFilePath); err != nil {
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

// GenerateClusterLogRequest structure represents the GenerateClusterLog operation request
type GenerateClusterLogRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis        `idl:"name:This" json:"this"`
	SpanMinutes uint32                `idl:"name:SpanMinutes" json:"span_minutes"`
	Flags       csvp.ClusterLogExFlag `idl:"name:flags" json:"flags"`
}

func (o *GenerateClusterLogRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterLogOperation {
	if o == nil {
		return &xxx_GenerateClusterLogOperation{}
	}
	return &xxx_GenerateClusterLogOperation{
		This:        o.This,
		SpanMinutes: o.SpanMinutes,
		Flags:       o.Flags,
	}
}

func (o *GenerateClusterLogRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterLogOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SpanMinutes = op.SpanMinutes
	o.Flags = op.Flags
}
func (o *GenerateClusterLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateClusterLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateClusterLogResponse structure represents the GenerateClusterLog operation response
type GenerateClusterLogResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LogFilePath: Upon successful completion of this method, the server MUST set this
	// parameter to the location where the server has exposed a file containing the diagnostic
	// log data. The path is relative to the machine and starts with a share name. The format
	// is "<share>\<filename>" where <share> is a share name, and <filename> is the name
	// of the file or device. The LogFilePath parameter MUST form a valid UncPath if "\\<servername>\"
	// is prepended to its contents. On unsuccessful completion of this method, the client
	// MUST ignore this value.
	LogFilePath *oaut.String `idl:"name:LogFilePath" json:"log_file_path"`
	// Return: The GenerateClusterLog return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateClusterLogResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterLogOperation {
	if o == nil {
		return &xxx_GenerateClusterLogOperation{}
	}
	return &xxx_GenerateClusterLogOperation{
		That:        o.That,
		LogFilePath: o.LogFilePath,
		Return:      o.Return,
	}
}

func (o *GenerateClusterLogResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterLogOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFilePath = op.LogFilePath
	o.Return = op.Return
}
func (o *GenerateClusterLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateClusterLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GenerateClusterHealthLogOperation structure represents the GenerateClusterHealthLog operation
type xxx_GenerateClusterHealthLogOperation struct {
	This        *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat        `idl:"name:That" json:"that"`
	SpanMinutes uint32                `idl:"name:SpanMinutes" json:"span_minutes"`
	Flags       csvp.ClusterLogExFlag `idl:"name:flags" json:"flags"`
	LogFilePath *oaut.String          `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterHealthLogOperation) OpNum() int { return 4 }

func (o *xxx_GenerateClusterHealthLogOperation) OpName() string {
	return "/IClusterLogEx/v0/GenerateClusterHealthLog"
}

func (o *xxx_GenerateClusterHealthLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterHealthLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// SpanMinutes {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.SpanMinutes); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=ClusterLogExFlag}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterHealthLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// SpanMinutes {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.SpanMinutes); err != nil {
			return err
		}
	}
	// flags {in} (1:{alias=ClusterLogExFlag}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterHealthLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterHealthLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// LogFilePath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LogFilePath != nil {
			_ptr_LogFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LogFilePath != nil {
					if err := o.LogFilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LogFilePath, _ptr_LogFilePath); err != nil {
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

func (o *xxx_GenerateClusterHealthLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// LogFilePath {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_LogFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LogFilePath == nil {
				o.LogFilePath = &oaut.String{}
			}
			if err := o.LogFilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_LogFilePath := func(ptr interface{}) { o.LogFilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LogFilePath, _s_LogFilePath, _ptr_LogFilePath); err != nil {
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

// GenerateClusterHealthLogRequest structure represents the GenerateClusterHealthLog operation request
type GenerateClusterHealthLogRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SpanMinutes: A value, in minutes, that indicates those values that SHOULD be in the
	// log. Events that occurred in the range of Now to (Now – SpanMinutes) MUST be in
	// the log and no others. If ClusterLogFlagLocalTime is set in the Flags field, Now
	// is the GMT on the server; otherwise, it is the local time on the server.
	SpanMinutes uint32                `idl:"name:SpanMinutes" json:"span_minutes"`
	Flags       csvp.ClusterLogExFlag `idl:"name:flags" json:"flags"`
}

func (o *GenerateClusterHealthLogRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterHealthLogOperation {
	if o == nil {
		return &xxx_GenerateClusterHealthLogOperation{}
	}
	return &xxx_GenerateClusterHealthLogOperation{
		This:        o.This,
		SpanMinutes: o.SpanMinutes,
		Flags:       o.Flags,
	}
}

func (o *GenerateClusterHealthLogRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterHealthLogOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SpanMinutes = op.SpanMinutes
	o.Flags = op.Flags
}
func (o *GenerateClusterHealthLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateClusterHealthLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterHealthLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateClusterHealthLogResponse structure represents the GenerateClusterHealthLog operation response
type GenerateClusterHealthLogResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LogFilePath: Has the same meaning as parameter LogFilePath for the GenerateClusterLog
	// method specified in section 3.12.4.1.
	LogFilePath *oaut.String `idl:"name:LogFilePath" json:"log_file_path"`
	// Return: The GenerateClusterHealthLog return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateClusterHealthLogResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterHealthLogOperation {
	if o == nil {
		return &xxx_GenerateClusterHealthLogOperation{}
	}
	return &xxx_GenerateClusterHealthLogOperation{
		That:        o.That,
		LogFilePath: o.LogFilePath,
		Return:      o.Return,
	}
}

func (o *GenerateClusterHealthLogResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterHealthLogOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFilePath = op.LogFilePath
	o.Return = op.Return
}
func (o *GenerateClusterHealthLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateClusterHealthLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterHealthLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
