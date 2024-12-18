package winsi2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	raiw "github.com/oiweiwei/go-msrpc/msrpc/raiw"
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
	_ = raiw.GoPackage
)

var (
	// import guard
	GoPackage = "raiw"
)

var (
	// Syntax UUID
	Winsi2SyntaxUUID = &uuid.UUID{TimeLow: 0x811109bf, TimeMid: 0xa4e1, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x54, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x1e, 0x9b, 0x45}}
	// Syntax ID
	Winsi2SyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: Winsi2SyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// winsi2 interface.
type Winsi2Client interface {

	// The R_WinsTombstoneDbRecs method tombstones records whose version numbers fall within
	// a range of version numbers and are owned by a server with a specified address.
	//
	// Return Values: A 32 bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller does not have sufficient permissions. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	TombstoneDBRecords(context.Context, *TombstoneDBRecordsRequest, ...dcerpc.CallOption) (*TombstoneDBRecordsResponse, error)

	// The R_ WinsCheckAccess method retrieves the level of access the client is granted.<12>
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CheckAccess(context.Context, *CheckAccessRequest, ...dcerpc.CallOption) (*CheckAccessResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultWinsi2Client struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWinsi2Client) TombstoneDBRecords(ctx context.Context, in *TombstoneDBRecordsRequest, opts ...dcerpc.CallOption) (*TombstoneDBRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TombstoneDBRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsi2Client) CheckAccess(ctx context.Context, in *CheckAccessRequest, opts ...dcerpc.CallOption) (*CheckAccessResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CheckAccessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsi2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWinsi2Client) Conn() dcerpc.Conn {
	return o.cc
}

func NewWinsi2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Winsi2Client, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Winsi2SyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWinsi2Client{cc: cc}, nil
}

// xxx_TombstoneDBRecordsOperation structure represents the R_WinsTombstoneDbRecs operation
type xxx_TombstoneDBRecordsOperation struct {
	WINSAddr  *raiw.Addr   `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
	Return    uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_TombstoneDBRecordsOperation) OpNum() int { return 0 }

func (o *xxx_TombstoneDBRecordsOperation) OpName() string { return "/winsi2/v1/R_WinsTombstoneDbRecs" }

func (o *xxx_TombstoneDBRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TombstoneDBRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo != nil {
			if err := o.MinVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo != nil {
			if err := o.MaxVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_TombstoneDBRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo == nil {
			o.MinVersNo = &raiw.VersNo{}
		}
		if err := o.MinVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo == nil {
			o.MaxVersNo = &raiw.VersNo{}
		}
		if err := o.MaxVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TombstoneDBRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TombstoneDBRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TombstoneDBRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TombstoneDBRecordsRequest structure represents the R_WinsTombstoneDbRecs operation request
type TombstoneDBRecordsRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	// MinVersNo: The lower bound on the range of version numbers that identifies the range
	// of records to be tombstoned.
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	// MaxVersNo: The upper bound on the range of version numbers that identifies the range
	// of records to be tombstoned.
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
}

func (o *TombstoneDBRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_TombstoneDBRecordsOperation {
	if o == nil {
		return &xxx_TombstoneDBRecordsOperation{}
	}
	return &xxx_TombstoneDBRecordsOperation{
		WINSAddr:  o.WINSAddr,
		MinVersNo: o.MinVersNo,
		MaxVersNo: o.MaxVersNo,
	}
}

func (o *TombstoneDBRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_TombstoneDBRecordsOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.MinVersNo = op.MinVersNo
	o.MaxVersNo = op.MaxVersNo
}
func (o *TombstoneDBRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TombstoneDBRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TombstoneDBRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TombstoneDBRecordsResponse structure represents the R_WinsTombstoneDbRecs operation response
type TombstoneDBRecordsResponse struct {
	// Return: The R_WinsTombstoneDbRecs return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TombstoneDBRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_TombstoneDBRecordsOperation {
	if o == nil {
		return &xxx_TombstoneDBRecordsOperation{}
	}
	return &xxx_TombstoneDBRecordsOperation{
		Return: o.Return,
	}
}

func (o *TombstoneDBRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_TombstoneDBRecordsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *TombstoneDBRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TombstoneDBRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TombstoneDBRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CheckAccessOperation structure represents the R_WinsCheckAccess operation
type xxx_CheckAccessOperation struct {
	Access uint32 `idl:"name:Access" json:"access"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CheckAccessOperation) OpNum() int { return 1 }

func (o *xxx_CheckAccessOperation) OpName() string { return "/winsi2/v1/R_WinsCheckAccess" }

func (o *xxx_CheckAccessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckAccessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_CheckAccessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_CheckAccessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckAccessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Access {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Access); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckAccessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Access {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Access); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CheckAccessRequest structure represents the R_WinsCheckAccess operation request
type CheckAccessRequest struct {
}

func (o *CheckAccessRequest) xxx_ToOp(ctx context.Context) *xxx_CheckAccessOperation {
	if o == nil {
		return &xxx_CheckAccessOperation{}
	}
	return &xxx_CheckAccessOperation{}
}

func (o *CheckAccessRequest) xxx_FromOp(ctx context.Context, op *xxx_CheckAccessOperation) {
	if o == nil {
		return
	}
}
func (o *CheckAccessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CheckAccessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckAccessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CheckAccessResponse structure represents the R_WinsCheckAccess operation response
type CheckAccessResponse struct {
	// Access: Pointer to the access level value. This value MUST not be NULL. The following
	// values are possible as output.
	//
	//	+----------------------+-------+
	//	|                      |       |
	//	|         NAME         | VALUE |
	//	|                      |       |
	//	+----------------------+-------+
	//	+----------------------+-------+
	//	| No access            |     0 |
	//	+----------------------+-------+
	//	| Control level access |     1 |
	//	+----------------------+-------+
	//	| Query level access   |     2 |
	//	+----------------------+-------+
	Access uint32 `idl:"name:Access" json:"access"`
	// Return: The R_WinsCheckAccess return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CheckAccessResponse) xxx_ToOp(ctx context.Context) *xxx_CheckAccessOperation {
	if o == nil {
		return &xxx_CheckAccessOperation{}
	}
	return &xxx_CheckAccessOperation{
		Access: o.Access,
		Return: o.Return,
	}
}

func (o *CheckAccessResponse) xxx_FromOp(ctx context.Context, op *xxx_CheckAccessOperation) {
	if o == nil {
		return
	}
	o.Access = op.Access
	o.Return = op.Return
}
func (o *CheckAccessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CheckAccessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckAccessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
