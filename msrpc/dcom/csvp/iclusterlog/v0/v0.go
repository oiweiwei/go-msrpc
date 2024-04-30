package iclusterlog

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
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterLog interface identifier 85923ca7-1b6b-4e83-a2e4-f5ba3bfbb8a3
	ClusterLogIID = &dcom.IID{Data1: 0x85923ca7, Data2: 0x1b6b, Data3: 0x4e83, Data4: []byte{0xa2, 0xe4, 0xf5, 0xba, 0x3b, 0xfb, 0xb8, 0xa3}}
	// Syntax UUID
	ClusterLogSyntaxUUID = &uuid.UUID{TimeLow: 0x85923ca7, TimeMid: 0x1b6b, TimeHiAndVersion: 0x4e83, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xe4, Node: [6]uint8{0xf5, 0xba, 0x3b, 0xfb, 0xb8, 0xa3}}
	// Syntax ID
	ClusterLogSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterLogSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterLog interface.
type ClusterLogClient interface {

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

	// The GenerateTimeSpanLog method writes a file that contains diagnostic information
	// about failover clusters for the server on which it executes. The log entries in the
	// file date back only for the specified number of minutes. The content and format of
	// the file is implementation-specific, but SHOULD contain diagnostic information.
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
	GenerateTimeSpanLog(context.Context, *GenerateTimeSpanLogRequest, ...dcerpc.CallOption) (*GenerateTimeSpanLogResponse, error)

	// The GenerateClusterLogInLocalTime method<41> writes a file that contains diagnostic
	// information about failover clusters for the server on which it executes. The file
	// uses local time instead of GMT. The content and format of the file are implementation-specific
	// but SHOULD contain diagnostic information.
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
	// The opnum field value for this method is 5.
	GenerateClusterLogInLocalTime(context.Context, *GenerateClusterLogInLocalTimeRequest, ...dcerpc.CallOption) (*GenerateClusterLogInLocalTimeResponse, error)

	// The GenerateTimeSpanLogInLocalTime method<42> writes a file that contains diagnostic
	// information about failover clusters for the server on which it executes. The log
	// entries in the file date back only for the specified number of minutes. The file
	// uses local time instead of GMT. The content and format of the file is implementation-specific
	// but SHOULD contain diagnostic information.
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
	// The opnum field value for this method is 6.
	GenerateTimeSpanLogInLocalTime(context.Context, *GenerateTimeSpanLogInLocalTimeRequest, ...dcerpc.CallOption) (*GenerateTimeSpanLogInLocalTimeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterLogClient
}

type xxx_DefaultClusterLogClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterLogClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterLogClient) GenerateClusterLog(ctx context.Context, in *GenerateClusterLogRequest, opts ...dcerpc.CallOption) (*GenerateClusterLogResponse, error) {
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

func (o *xxx_DefaultClusterLogClient) GenerateTimeSpanLog(ctx context.Context, in *GenerateTimeSpanLogRequest, opts ...dcerpc.CallOption) (*GenerateTimeSpanLogResponse, error) {
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
	out := &GenerateTimeSpanLogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterLogClient) GenerateClusterLogInLocalTime(ctx context.Context, in *GenerateClusterLogInLocalTimeRequest, opts ...dcerpc.CallOption) (*GenerateClusterLogInLocalTimeResponse, error) {
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
	out := &GenerateClusterLogInLocalTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterLogClient) GenerateTimeSpanLogInLocalTime(ctx context.Context, in *GenerateTimeSpanLogInLocalTimeRequest, opts ...dcerpc.CallOption) (*GenerateTimeSpanLogInLocalTimeResponse, error) {
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
	out := &GenerateTimeSpanLogInLocalTimeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterLogClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterLogClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterLogClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterLogClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewClusterLogClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterLogClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterLogSyntaxV0_0))...)
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
	return &xxx_DefaultClusterLogClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GenerateClusterLogOperation structure represents the GenerateClusterLog operation
type xxx_GenerateClusterLogOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogFilePath *oaut.String   `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterLogOperation) OpNum() int { return 3 }

func (o *xxx_GenerateClusterLogOperation) OpName() string {
	return "/IClusterLog/v0/GenerateClusterLog"
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
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GenerateClusterLogRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterLogOperation {
	if o == nil {
		return &xxx_GenerateClusterLogOperation{}
	}
	return &xxx_GenerateClusterLogOperation{
		This: o.This,
	}
}

func (o *GenerateClusterLogRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterLogOperation) {
	if o == nil {
		return
	}
	o.This = op.This
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

// xxx_GenerateTimeSpanLogOperation structure represents the GenerateTimeSpanLog operation
type xxx_GenerateTimeSpanLogOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	SpanMinutes uint32         `idl:"name:SpanMinutes" json:"span_minutes"`
	LogFilePath *oaut.String   `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateTimeSpanLogOperation) OpNum() int { return 4 }

func (o *xxx_GenerateTimeSpanLogOperation) OpName() string {
	return "/IClusterLog/v0/GenerateTimeSpanLog"
}

func (o *xxx_GenerateTimeSpanLogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// SpanMinutes {in} (1:(uint32))
	{
		if err := w.WriteData(o.SpanMinutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// SpanMinutes {in} (1:(uint32))
	{
		if err := w.ReadData(&o.SpanMinutes); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GenerateTimeSpanLogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GenerateTimeSpanLogRequest structure represents the GenerateTimeSpanLog operation request
type GenerateTimeSpanLogRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SpanMinutes: A value, in minutes, that indicates those values that SHOULD be in the
	// log. Events that occurred in the range of Now to (Now - SpanMinutes) MUST be in the
	// log and no others. Now is the GMT on the server.
	SpanMinutes uint32 `idl:"name:SpanMinutes" json:"span_minutes"`
}

func (o *GenerateTimeSpanLogRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateTimeSpanLogOperation {
	if o == nil {
		return &xxx_GenerateTimeSpanLogOperation{}
	}
	return &xxx_GenerateTimeSpanLogOperation{
		This:        o.This,
		SpanMinutes: o.SpanMinutes,
	}
}

func (o *GenerateTimeSpanLogRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateTimeSpanLogOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SpanMinutes = op.SpanMinutes
}
func (o *GenerateTimeSpanLogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateTimeSpanLogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateTimeSpanLogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateTimeSpanLogResponse structure represents the GenerateTimeSpanLog operation response
type GenerateTimeSpanLogResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LogFilePath: Has the same meaning as parameter LogFilePath for the GenerateClusterLog
	// method specified in section 3.12.4.1.
	LogFilePath *oaut.String `idl:"name:LogFilePath" json:"log_file_path"`
	// Return: The GenerateTimeSpanLog return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateTimeSpanLogResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateTimeSpanLogOperation {
	if o == nil {
		return &xxx_GenerateTimeSpanLogOperation{}
	}
	return &xxx_GenerateTimeSpanLogOperation{
		That:        o.That,
		LogFilePath: o.LogFilePath,
		Return:      o.Return,
	}
}

func (o *GenerateTimeSpanLogResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateTimeSpanLogOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFilePath = op.LogFilePath
	o.Return = op.Return
}
func (o *GenerateTimeSpanLogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateTimeSpanLogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateTimeSpanLogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GenerateClusterLogInLocalTimeOperation structure represents the GenerateClusterLogInLocalTime operation
type xxx_GenerateClusterLogInLocalTimeOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	LogFilePath *oaut.String   `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterLogInLocalTimeOperation) OpNum() int { return 5 }

func (o *xxx_GenerateClusterLogInLocalTimeOperation) OpName() string {
	return "/IClusterLog/v0/GenerateClusterLogInLocalTime"
}

func (o *xxx_GenerateClusterLogInLocalTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogInLocalTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GenerateClusterLogInLocalTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GenerateClusterLogInLocalTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterLogInLocalTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GenerateClusterLogInLocalTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GenerateClusterLogInLocalTimeRequest structure represents the GenerateClusterLogInLocalTime operation request
type GenerateClusterLogInLocalTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GenerateClusterLogInLocalTimeRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterLogInLocalTimeOperation {
	if o == nil {
		return &xxx_GenerateClusterLogInLocalTimeOperation{}
	}
	return &xxx_GenerateClusterLogInLocalTimeOperation{
		This: o.This,
	}
}

func (o *GenerateClusterLogInLocalTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterLogInLocalTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GenerateClusterLogInLocalTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateClusterLogInLocalTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterLogInLocalTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateClusterLogInLocalTimeResponse structure represents the GenerateClusterLogInLocalTime operation response
type GenerateClusterLogInLocalTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LogFilePath: Upon successful completion of this method, the server MUST set this
	// parameter to the location where the server has exposed a file containing the diagnostic
	// log data. The path is relative to the machine and starts with a share name. The format
	// is "<share>\<filename>" where <share> is a share name and <filename> is the name
	// of the file or device. The LogFilePath parameter MUST form a valid UncPath if \\<servername>\
	// is prepended to its contents. On unsuccessful completion of this method, the client
	// MUST ignore this value.
	LogFilePath *oaut.String `idl:"name:LogFilePath" json:"log_file_path"`
	// Return: The GenerateClusterLogInLocalTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateClusterLogInLocalTimeResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterLogInLocalTimeOperation {
	if o == nil {
		return &xxx_GenerateClusterLogInLocalTimeOperation{}
	}
	return &xxx_GenerateClusterLogInLocalTimeOperation{
		That:        o.That,
		LogFilePath: o.LogFilePath,
		Return:      o.Return,
	}
}

func (o *GenerateClusterLogInLocalTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterLogInLocalTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFilePath = op.LogFilePath
	o.Return = op.Return
}
func (o *GenerateClusterLogInLocalTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateClusterLogInLocalTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterLogInLocalTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GenerateTimeSpanLogInLocalTimeOperation structure represents the GenerateTimeSpanLogInLocalTime operation
type xxx_GenerateTimeSpanLogInLocalTimeOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	SpanMinutes uint32         `idl:"name:SpanMinutes" json:"span_minutes"`
	LogFilePath *oaut.String   `idl:"name:LogFilePath" json:"log_file_path"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) OpNum() int { return 6 }

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) OpName() string {
	return "/IClusterLog/v0/GenerateTimeSpanLogInLocalTime"
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GenerateTimeSpanLogInLocalTimeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GenerateTimeSpanLogInLocalTimeRequest structure represents the GenerateTimeSpanLogInLocalTime operation request
type GenerateTimeSpanLogInLocalTimeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SpanMinutes: A value, in minutes, that indicates those values that SHOULD be in the
	// log. Events that occurred in the range of Now to (Now - SpanMinutes) MUST be in the
	// log and no others. Now is the local time on the server.
	SpanMinutes uint32 `idl:"name:SpanMinutes" json:"span_minutes"`
}

func (o *GenerateTimeSpanLogInLocalTimeRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateTimeSpanLogInLocalTimeOperation {
	if o == nil {
		return &xxx_GenerateTimeSpanLogInLocalTimeOperation{}
	}
	return &xxx_GenerateTimeSpanLogInLocalTimeOperation{
		This:        o.This,
		SpanMinutes: o.SpanMinutes,
	}
}

func (o *GenerateTimeSpanLogInLocalTimeRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateTimeSpanLogInLocalTimeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SpanMinutes = op.SpanMinutes
}
func (o *GenerateTimeSpanLogInLocalTimeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateTimeSpanLogInLocalTimeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateTimeSpanLogInLocalTimeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateTimeSpanLogInLocalTimeResponse structure represents the GenerateTimeSpanLogInLocalTime operation response
type GenerateTimeSpanLogInLocalTimeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// LogFilePath: Has the same meaning as parameter LogFilePath for the GenerateClusterLog
	// method specified in section 3.12.4.1.
	LogFilePath *oaut.String `idl:"name:LogFilePath" json:"log_file_path"`
	// Return: The GenerateTimeSpanLogInLocalTime return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateTimeSpanLogInLocalTimeResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateTimeSpanLogInLocalTimeOperation {
	if o == nil {
		return &xxx_GenerateTimeSpanLogInLocalTimeOperation{}
	}
	return &xxx_GenerateTimeSpanLogInLocalTimeOperation{
		That:        o.That,
		LogFilePath: o.LogFilePath,
		Return:      o.Return,
	}
}

func (o *GenerateTimeSpanLogInLocalTimeResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateTimeSpanLogInLocalTimeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LogFilePath = op.LogFilePath
	o.Return = op.Return
}
func (o *GenerateTimeSpanLogInLocalTimeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateTimeSpanLogInLocalTimeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateTimeSpanLogInLocalTimeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
