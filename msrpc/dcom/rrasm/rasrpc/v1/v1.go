package rasrpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	rrasm "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm"
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
	_ = rrasm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// Syntax UUID
	RasrpcSyntaxUUID = &uuid.UUID{TimeLow: 0x20610036, TimeMid: 0xfa22, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x23, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x11, 0xe5, 0xdf}}
	// Syntax ID
	RasrpcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RasrpcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// rasrpc interface.
type RasrpcClient interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	RASRPCDeleteEntry(context.Context, *RASRPCDeleteEntryRequest, ...dcerpc.CallOption) (*RASRPCDeleteEntryResponse, error)

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	RASRPCGetUserPreferences(context.Context, *RASRPCGetUserPreferencesRequest, ...dcerpc.CallOption) (*RASRPCGetUserPreferencesResponse, error)

	RASRPCSetUserPreferences(context.Context, *RASRPCSetUserPreferencesRequest, ...dcerpc.CallOption) (*RASRPCSetUserPreferencesResponse, error)

	RASRPCGetSystemDirectory(context.Context, *RASRPCGetSystemDirectoryRequest, ...dcerpc.CallOption) (*RASRPCGetSystemDirectoryResponse, error)

	RASRPCSubmitRequest(context.Context, *RASRPCSubmitRequestRequest, ...dcerpc.CallOption) (*RASRPCSubmitRequestResponse, error)

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	RASRPCGetInstalledProtocolsEx(context.Context, *RASRPCGetInstalledProtocolsExRequest, ...dcerpc.CallOption) (*RASRPCGetInstalledProtocolsExResponse, error)

	RASRPCGetVersion(context.Context, *RASRPCGetVersionRequest, ...dcerpc.CallOption) (*RASRPCGetVersionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultRasrpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRasrpcClient) RASRPCDeleteEntry(ctx context.Context, in *RASRPCDeleteEntryRequest, opts ...dcerpc.CallOption) (*RASRPCDeleteEntryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCDeleteEntryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCGetUserPreferences(ctx context.Context, in *RASRPCGetUserPreferencesRequest, opts ...dcerpc.CallOption) (*RASRPCGetUserPreferencesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCGetUserPreferencesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCSetUserPreferences(ctx context.Context, in *RASRPCSetUserPreferencesRequest, opts ...dcerpc.CallOption) (*RASRPCSetUserPreferencesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCSetUserPreferencesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCGetSystemDirectory(ctx context.Context, in *RASRPCGetSystemDirectoryRequest, opts ...dcerpc.CallOption) (*RASRPCGetSystemDirectoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCGetSystemDirectoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCSubmitRequest(ctx context.Context, in *RASRPCSubmitRequestRequest, opts ...dcerpc.CallOption) (*RASRPCSubmitRequestResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCSubmitRequestResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCGetInstalledProtocolsEx(ctx context.Context, in *RASRPCGetInstalledProtocolsExRequest, opts ...dcerpc.CallOption) (*RASRPCGetInstalledProtocolsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCGetInstalledProtocolsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) RASRPCGetVersion(ctx context.Context, in *RASRPCGetVersionRequest, opts ...dcerpc.CallOption) (*RASRPCGetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASRPCGetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRasrpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRasrpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRasrpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RasrpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RasrpcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRasrpcClient{cc: cc}, nil
}

// xxx_RASRPCDeleteEntryOperation structure represents the RasRpcDeleteEntry operation
type xxx_RASRPCDeleteEntryOperation struct {
	Phonebook string `idl:"name:lpszPhonebook;string" json:"phonebook"`
	Entry     string `idl:"name:lpszEntry;string" json:"entry"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCDeleteEntryOperation) OpNum() int { return 5 }

func (o *xxx_RASRPCDeleteEntryOperation) OpName() string { return "/rasrpc/v1/RasRpcDeleteEntry" }

func (o *xxx_RASRPCDeleteEntryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCDeleteEntryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszPhonebook {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Phonebook); err != nil {
			return err
		}
	}
	// lpszEntry {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCDeleteEntryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszPhonebook {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Phonebook); err != nil {
			return err
		}
	}
	// lpszEntry {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Entry); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCDeleteEntryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCDeleteEntryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASRPCDeleteEntryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASRPCDeleteEntryRequest structure represents the RasRpcDeleteEntry operation request
type RASRPCDeleteEntryRequest struct {
	Phonebook string `idl:"name:lpszPhonebook;string" json:"phonebook"`
	Entry     string `idl:"name:lpszEntry;string" json:"entry"`
}

func (o *RASRPCDeleteEntryRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCDeleteEntryOperation) *xxx_RASRPCDeleteEntryOperation {
	if op == nil {
		op = &xxx_RASRPCDeleteEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.Phonebook = o.Phonebook
	op.Entry = o.Entry
	return op
}

func (o *RASRPCDeleteEntryRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCDeleteEntryOperation) {
	if o == nil {
		return
	}
	o.Phonebook = op.Phonebook
	o.Entry = op.Entry
}
func (o *RASRPCDeleteEntryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCDeleteEntryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCDeleteEntryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCDeleteEntryResponse structure represents the RasRpcDeleteEntry operation response
type RASRPCDeleteEntryResponse struct {
	// Return: The RasRpcDeleteEntry return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCDeleteEntryResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCDeleteEntryOperation) *xxx_RASRPCDeleteEntryOperation {
	if op == nil {
		op = &xxx_RASRPCDeleteEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASRPCDeleteEntryResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCDeleteEntryOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASRPCDeleteEntryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCDeleteEntryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCDeleteEntryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCGetUserPreferencesOperation structure represents the RasRpcGetUserPreferences operation
type xxx_RASRPCGetUserPreferencesOperation struct {
	User   *rrasm.RasrpcPbuser `idl:"name:pUser" json:"user"`
	Mode   uint32              `idl:"name:dwMode" json:"mode"`
	Return uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCGetUserPreferencesOperation) OpNum() int { return 9 }

func (o *xxx_RASRPCGetUserPreferencesOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetUserPreferences"
}

func (o *xxx_RASRPCGetUserPreferencesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetUserPreferencesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RasrpcPbuser{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetUserPreferencesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.RasrpcPbuser{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetUserPreferencesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetUserPreferencesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RasrpcPbuser{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
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

func (o *xxx_RASRPCGetUserPreferencesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pUser {in, out} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.RasrpcPbuser{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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

// RASRPCGetUserPreferencesRequest structure represents the RasRpcGetUserPreferences operation request
type RASRPCGetUserPreferencesRequest struct {
	User *rrasm.RasrpcPbuser `idl:"name:pUser" json:"user"`
	Mode uint32              `idl:"name:dwMode" json:"mode"`
}

func (o *RASRPCGetUserPreferencesRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetUserPreferencesOperation) *xxx_RASRPCGetUserPreferencesOperation {
	if op == nil {
		op = &xxx_RASRPCGetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Mode = o.Mode
	return op
}

func (o *RASRPCGetUserPreferencesRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Mode = op.Mode
}
func (o *RASRPCGetUserPreferencesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCGetUserPreferencesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCGetUserPreferencesResponse structure represents the RasRpcGetUserPreferences operation response
type RASRPCGetUserPreferencesResponse struct {
	User *rrasm.RasrpcPbuser `idl:"name:pUser" json:"user"`
	// Return: The RasRpcGetUserPreferences return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCGetUserPreferencesResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetUserPreferencesOperation) *xxx_RASRPCGetUserPreferencesOperation {
	if op == nil {
		op = &xxx_RASRPCGetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Return = o.Return
	return op
}

func (o *RASRPCGetUserPreferencesResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Return = op.Return
}
func (o *RASRPCGetUserPreferencesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCGetUserPreferencesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCSetUserPreferencesOperation structure represents the RasRpcSetUserPreferences operation
type xxx_RASRPCSetUserPreferencesOperation struct {
	User   *rrasm.RasrpcPbuser `idl:"name:pUser" json:"user"`
	Mode   uint32              `idl:"name:dwMode" json:"mode"`
	Return uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCSetUserPreferencesOperation) OpNum() int { return 10 }

func (o *xxx_RASRPCSetUserPreferencesOperation) OpName() string {
	return "/rasrpc/v1/RasRpcSetUserPreferences"
}

func (o *xxx_RASRPCSetUserPreferencesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSetUserPreferencesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pUser {in} (1:{alias=LPRASRPC_PBUSER}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RasrpcPbuser{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSetUserPreferencesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pUser {in} (1:{alias=LPRASRPC_PBUSER,pointer=ref}*(1))(2:{alias=RASRPC_PBUSER}(struct))
	{
		if o.User == nil {
			o.User = &rrasm.RasrpcPbuser{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwMode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSetUserPreferencesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSetUserPreferencesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASRPCSetUserPreferencesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASRPCSetUserPreferencesRequest structure represents the RasRpcSetUserPreferences operation request
type RASRPCSetUserPreferencesRequest struct {
	User *rrasm.RasrpcPbuser `idl:"name:pUser" json:"user"`
	Mode uint32              `idl:"name:dwMode" json:"mode"`
}

func (o *RASRPCSetUserPreferencesRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCSetUserPreferencesOperation) *xxx_RASRPCSetUserPreferencesOperation {
	if op == nil {
		op = &xxx_RASRPCSetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.User = o.User
	op.Mode = o.Mode
	return op
}

func (o *RASRPCSetUserPreferencesRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCSetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Mode = op.Mode
}
func (o *RASRPCSetUserPreferencesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCSetUserPreferencesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCSetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCSetUserPreferencesResponse structure represents the RasRpcSetUserPreferences operation response
type RASRPCSetUserPreferencesResponse struct {
	// Return: The RasRpcSetUserPreferences return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCSetUserPreferencesResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCSetUserPreferencesOperation) *xxx_RASRPCSetUserPreferencesOperation {
	if op == nil {
		op = &xxx_RASRPCSetUserPreferencesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASRPCSetUserPreferencesResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCSetUserPreferencesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASRPCSetUserPreferencesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCSetUserPreferencesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCSetUserPreferencesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCGetSystemDirectoryOperation structure represents the RasRpcGetSystemDirectory operation
type xxx_RASRPCGetSystemDirectoryOperation struct {
	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	Size   uint32 `idl:"name:uSize" json:"size"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) OpNum() int { return 11 }

func (o *xxx_RASRPCGetSystemDirectoryOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetSystemDirectory"
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != "" && o.Size == 0 {
		o.Size = uint32(ndr.UTF16NLen(o.Buffer))
	}
	if o.Size > uint32(260) {
		return fmt.Errorf("Size is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpBuffer {in, out} (1:{string, alias=LPWSTR}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Buffer)
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
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0]-1 {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]-1]
		}
		if o.Buffer != ndr.ZeroString {
			_Buffer_buf = append(_Buffer_buf, uint16(0))
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// uSize {in} (1:{range=(0,260), alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpBuffer {in, out} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
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
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// uSize {in} (1:{range=(0,260), alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpBuffer {in, out} (1:{string, alias=LPWSTR}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Buffer)
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
		_Buffer_buf := utf16.Encode([]rune(o.Buffer))
		if uint64(len(_Buffer_buf)) > sizeInfo[0]-1 {
			_Buffer_buf = _Buffer_buf[:sizeInfo[0]-1]
		}
		if o.Buffer != ndr.ZeroString {
			_Buffer_buf = append(_Buffer_buf, uint16(0))
		}
		for i1 := range _Buffer_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetSystemDirectoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpBuffer {in, out} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,size_is=uSize,string,null](wchar))
	{
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
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
	}
	// Return {out} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASRPCGetSystemDirectoryRequest structure represents the RasRpcGetSystemDirectory operation request
type RASRPCGetSystemDirectoryRequest struct {
	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	Size   uint32 `idl:"name:uSize" json:"size"`
}

func (o *RASRPCGetSystemDirectoryRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetSystemDirectoryOperation) *xxx_RASRPCGetSystemDirectoryOperation {
	if op == nil {
		op = &xxx_RASRPCGetSystemDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	op.Buffer = o.Buffer
	op.Size = o.Size
	return op
}

func (o *RASRPCGetSystemDirectoryRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetSystemDirectoryOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Size = op.Size
}
func (o *RASRPCGetSystemDirectoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCGetSystemDirectoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetSystemDirectoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCGetSystemDirectoryResponse structure represents the RasRpcGetSystemDirectory operation response
type RASRPCGetSystemDirectoryResponse struct {
	// XXX: uSize is an implicit input depedency for output parameters
	Size uint32 `idl:"name:uSize" json:"size"`

	Buffer string `idl:"name:lpBuffer;size_is:(uSize);string" json:"buffer"`
	// Return: The RasRpcGetSystemDirectory return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCGetSystemDirectoryResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetSystemDirectoryOperation) *xxx_RASRPCGetSystemDirectoryOperation {
	if op == nil {
		op = &xxx_RASRPCGetSystemDirectoryOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Size == uint32(0) {
		op.Size = o.Size
	}

	op.Buffer = o.Buffer
	op.Return = o.Return
	return op
}

func (o *RASRPCGetSystemDirectoryResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetSystemDirectoryOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Size = op.Size

	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *RASRPCGetSystemDirectoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCGetSystemDirectoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetSystemDirectoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCSubmitRequestOperation structure represents the RasRpcSubmitRequest operation
type xxx_RASRPCSubmitRequestOperation struct {
	RequestBuffer  []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	DwcbBufferSize uint32 `idl:"name:dwcbBufSize" json:"dwcb_buffer_size"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCSubmitRequestOperation) OpNum() int { return 12 }

func (o *xxx_RASRPCSubmitRequestOperation) OpName() string { return "/rasrpc/v1/RasRpcSubmitRequest" }

func (o *xxx_RASRPCSubmitRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestBuffer != nil && o.DwcbBufferSize == 0 {
		o.DwcbBufferSize = uint32(len(o.RequestBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSubmitRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		if o.RequestBuffer != nil || o.DwcbBufferSize > 0 {
			_ptr_pReqBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DwcbBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.RequestBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.RequestBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestBuffer, _ptr_pReqBuffer); err != nil {
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
	// dwcbBufSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DwcbBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSubmitRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		_ptr_pReqBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestBuffer", sizeInfo[0])
			}
			o.RequestBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.RequestBuffer {
				i1 := i1
				if err := w.ReadData(&o.RequestBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pReqBuffer := func(ptr interface{}) { o.RequestBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RequestBuffer, _s_pReqBuffer, _ptr_pReqBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwcbBufSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DwcbBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSubmitRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSubmitRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		if o.RequestBuffer != nil || o.DwcbBufferSize > 0 {
			_ptr_pReqBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DwcbBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RequestBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.RequestBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.RequestBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RequestBuffer, _ptr_pReqBuffer); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCSubmitRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pReqBuffer {in, out} (1:{pointer=unique, alias=PBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwcbBufSize](uchar))
	{
		_ptr_pReqBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RequestBuffer", sizeInfo[0])
			}
			o.RequestBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.RequestBuffer {
				i1 := i1
				if err := w.ReadData(&o.RequestBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pReqBuffer := func(ptr interface{}) { o.RequestBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RequestBuffer, _s_pReqBuffer, _ptr_pReqBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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

// RASRPCSubmitRequestRequest structure represents the RasRpcSubmitRequest operation request
type RASRPCSubmitRequestRequest struct {
	RequestBuffer  []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	DwcbBufferSize uint32 `idl:"name:dwcbBufSize" json:"dwcb_buffer_size"`
}

func (o *RASRPCSubmitRequestRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCSubmitRequestOperation) *xxx_RASRPCSubmitRequestOperation {
	if op == nil {
		op = &xxx_RASRPCSubmitRequestOperation{}
	}
	if o == nil {
		return op
	}
	op.RequestBuffer = o.RequestBuffer
	op.DwcbBufferSize = o.DwcbBufferSize
	return op
}

func (o *RASRPCSubmitRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCSubmitRequestOperation) {
	if o == nil {
		return
	}
	o.RequestBuffer = op.RequestBuffer
	o.DwcbBufferSize = op.DwcbBufferSize
}
func (o *RASRPCSubmitRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCSubmitRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCSubmitRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCSubmitRequestResponse structure represents the RasRpcSubmitRequest operation response
type RASRPCSubmitRequestResponse struct {
	// XXX: dwcbBufSize is an implicit input depedency for output parameters
	DwcbBufferSize uint32 `idl:"name:dwcbBufSize" json:"dwcb_buffer_size"`

	RequestBuffer []byte `idl:"name:pReqBuffer;size_is:(dwcbBufSize);pointer:unique" json:"request_buffer"`
	// Return: The RasRpcSubmitRequest return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCSubmitRequestResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCSubmitRequestOperation) *xxx_RASRPCSubmitRequestOperation {
	if op == nil {
		op = &xxx_RASRPCSubmitRequestOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.DwcbBufferSize == uint32(0) {
		op.DwcbBufferSize = o.DwcbBufferSize
	}

	op.RequestBuffer = o.RequestBuffer
	op.Return = o.Return
	return op
}

func (o *RASRPCSubmitRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCSubmitRequestOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.DwcbBufferSize = op.DwcbBufferSize

	o.RequestBuffer = op.RequestBuffer
	o.Return = op.Return
}
func (o *RASRPCSubmitRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCSubmitRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCSubmitRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCGetInstalledProtocolsExOperation structure represents the RasRpcGetInstalledProtocolsEx operation
type xxx_RASRPCGetInstalledProtocolsExOperation struct {
	Router    bool   `idl:"name:fRouter" json:"router"`
	RASCli    bool   `idl:"name:fRasCli" json:"ras_cli"`
	RASServer bool   `idl:"name:fRasSrv" json:"ras_server"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) OpNum() int { return 14 }

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) OpName() string {
	return "/rasrpc/v1/RasRpcGetInstalledProtocolsEx"
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fRouter {in} (1:{alias=BOOL}(int32))
	{
		if !o.Router {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// fRasCli {in} (1:{alias=BOOL}(int32))
	{
		if !o.RASCli {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	// fRasSrv {in} (1:{alias=BOOL}(int32))
	{
		if !o.RASServer {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fRouter {in} (1:{alias=BOOL}(int32))
	{
		var _bRouter int32
		if err := w.ReadData(&_bRouter); err != nil {
			return err
		}
		o.Router = _bRouter != 0
	}
	// fRasCli {in} (1:{alias=BOOL}(int32))
	{
		var _bRASCli int32
		if err := w.ReadData(&_bRASCli); err != nil {
			return err
		}
		o.RASCli = _bRASCli != 0
	}
	// fRasSrv {in} (1:{alias=BOOL}(int32))
	{
		var _bRASServer int32
		if err := w.ReadData(&_bRASServer); err != nil {
			return err
		}
		o.RASServer = _bRASServer != 0
	}
	return nil
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASRPCGetInstalledProtocolsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASRPCGetInstalledProtocolsExRequest structure represents the RasRpcGetInstalledProtocolsEx operation request
type RASRPCGetInstalledProtocolsExRequest struct {
	Router    bool `idl:"name:fRouter" json:"router"`
	RASCli    bool `idl:"name:fRasCli" json:"ras_cli"`
	RASServer bool `idl:"name:fRasSrv" json:"ras_server"`
}

func (o *RASRPCGetInstalledProtocolsExRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetInstalledProtocolsExOperation) *xxx_RASRPCGetInstalledProtocolsExOperation {
	if op == nil {
		op = &xxx_RASRPCGetInstalledProtocolsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Router = o.Router
	op.RASCli = o.RASCli
	op.RASServer = o.RASServer
	return op
}

func (o *RASRPCGetInstalledProtocolsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetInstalledProtocolsExOperation) {
	if o == nil {
		return
	}
	o.Router = op.Router
	o.RASCli = op.RASCli
	o.RASServer = op.RASServer
}
func (o *RASRPCGetInstalledProtocolsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCGetInstalledProtocolsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetInstalledProtocolsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCGetInstalledProtocolsExResponse structure represents the RasRpcGetInstalledProtocolsEx operation response
type RASRPCGetInstalledProtocolsExResponse struct {
	// Return: The RasRpcGetInstalledProtocolsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCGetInstalledProtocolsExResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetInstalledProtocolsExOperation) *xxx_RASRPCGetInstalledProtocolsExOperation {
	if op == nil {
		op = &xxx_RASRPCGetInstalledProtocolsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASRPCGetInstalledProtocolsExResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetInstalledProtocolsExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASRPCGetInstalledProtocolsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCGetInstalledProtocolsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetInstalledProtocolsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASRPCGetVersionOperation structure represents the RasRpcGetVersion operation
type xxx_RASRPCGetVersionOperation struct {
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASRPCGetVersionOperation) OpNum() int { return 15 }

func (o *xxx_RASRPCGetVersionOperation) OpName() string { return "/rasrpc/v1/RasRpcGetVersion" }

func (o *xxx_RASRPCGetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASRPCGetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
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

func (o *xxx_RASRPCGetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwVersion {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
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

// RASRPCGetVersionRequest structure represents the RasRpcGetVersion operation request
type RASRPCGetVersionRequest struct {
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
}

func (o *RASRPCGetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetVersionOperation) *xxx_RASRPCGetVersionOperation {
	if op == nil {
		op = &xxx_RASRPCGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	return op
}

func (o *RASRPCGetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
}
func (o *RASRPCGetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASRPCGetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASRPCGetVersionResponse structure represents the RasRpcGetVersion operation response
type RASRPCGetVersionResponse struct {
	Version uint32 `idl:"name:pdwVersion;pointer:ref" json:"version"`
	// Return: The RasRpcGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASRPCGetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_RASRPCGetVersionOperation) *xxx_RASRPCGetVersionOperation {
	if op == nil {
		op = &xxx_RASRPCGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.Return = o.Return
	return op
}

func (o *RASRPCGetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_RASRPCGetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Return = op.Return
}
func (o *RASRPCGetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASRPCGetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASRPCGetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
