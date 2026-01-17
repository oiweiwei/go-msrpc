package isinglesignonremotemastersecret

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "ssp"
)

var (
	// Syntax UUID
	SSORemoteMasterSecretSyntaxUUID = &uuid.UUID{TimeLow: 0x9d07ca0d, TimeMid: 0x8f02, TimeHiAndVersion: 0x4ed5, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x27, Node: [6]uint8{0xac, 0xf3, 0x7f, 0xea, 0x5b, 0xbc}}
	// Syntax ID
	SSORemoteMasterSecretSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: SSORemoteMasterSecretSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// ISingleSignonRemoteMasterSecret interface.
type SSORemoteMasterSecretClient interface {

	// RemoteGetMasterSecret operation.
	RemoteGetMasterSecret(context.Context, *RemoteGetMasterSecretRequest, ...dcerpc.CallOption) (*RemoteGetMasterSecretResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultSSORemoteMasterSecretClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultSSORemoteMasterSecretClient) RemoteGetMasterSecret(ctx context.Context, in *RemoteGetMasterSecretRequest, opts ...dcerpc.CallOption) (*RemoteGetMasterSecretResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteGetMasterSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSSORemoteMasterSecretClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSSORemoteMasterSecretClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewSSORemoteMasterSecretClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SSORemoteMasterSecretClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SSORemoteMasterSecretSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultSSORemoteMasterSecretClient{cc: cc}, nil
}

// xxx_RemoteGetMasterSecretOperation structure represents the RemoteGetMasterSecret operation
type xxx_RemoteGetMasterSecretOperation struct {
	SecretLength uint32 `idl:"name:pcbSecret" json:"secret_length"`
	Secret       []byte `idl:"name:pbSecret;size_is:(pcbSecret)" json:"secret"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteGetMasterSecretOperation) OpNum() int { return 0 }

func (o *xxx_RemoteGetMasterSecretOperation) OpName() string {
	return "/ISingleSignonRemoteMasterSecret/v1/RemoteGetMasterSecret"
}

func (o *xxx_RemoteGetMasterSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetMasterSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pcbSecret {in, out} (1:{pointer=ref}*(1))(2:{range=(16,16), alias=SizeOfMasterSecretInBytes}(uint32))
	{
		if err := w.WriteData(o.SecretLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetMasterSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pcbSecret {in, out} (1:{pointer=ref}*(1))(2:{range=(16,16), alias=SizeOfMasterSecretInBytes}(uint32))
	{
		if err := w.ReadData(&o.SecretLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetMasterSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Secret != nil && o.SecretLength == 0 {
		o.SecretLength = uint32(len(o.Secret))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetMasterSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcbSecret {in, out} (1:{pointer=ref}*(1))(2:{range=(16,16), alias=SizeOfMasterSecretInBytes}(uint32))
	{
		if err := w.WriteData(o.SecretLength); err != nil {
			return err
		}
	}
	// pbSecret {out} (1:{pointer=ref}*(1)[dim:0,size_is=pcbSecret](uint8))
	{
		dimSize1 := uint64(o.SecretLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Secret {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Secret[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Secret); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteGetMasterSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcbSecret {in, out} (1:{pointer=ref}*(1))(2:{range=(16,16), alias=SizeOfMasterSecretInBytes}(uint32))
	{
		if err := w.ReadData(&o.SecretLength); err != nil {
			return err
		}
	}
	// pbSecret {out} (1:{pointer=ref}*(1)[dim:0,size_is=pcbSecret](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Secret", sizeInfo[0])
		}
		o.Secret = make([]byte, sizeInfo[0])
		for i1 := range o.Secret {
			i1 := i1
			if err := w.ReadData(&o.Secret[i1]); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteGetMasterSecretRequest structure represents the RemoteGetMasterSecret operation request
type RemoteGetMasterSecretRequest struct {
	SecretLength uint32 `idl:"name:pcbSecret" json:"secret_length"`
}

func (o *RemoteGetMasterSecretRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteGetMasterSecretOperation) *xxx_RemoteGetMasterSecretOperation {
	if op == nil {
		op = &xxx_RemoteGetMasterSecretOperation{}
	}
	if o == nil {
		return op
	}
	op.SecretLength = o.SecretLength
	return op
}

func (o *RemoteGetMasterSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteGetMasterSecretOperation) {
	if o == nil {
		return
	}
	o.SecretLength = op.SecretLength
}
func (o *RemoteGetMasterSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteGetMasterSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteGetMasterSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteGetMasterSecretResponse structure represents the RemoteGetMasterSecret operation response
type RemoteGetMasterSecretResponse struct {
	SecretLength uint32 `idl:"name:pcbSecret" json:"secret_length"`
	Secret       []byte `idl:"name:pbSecret;size_is:(pcbSecret)" json:"secret"`
	// Return: The RemoteGetMasterSecret return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoteGetMasterSecretResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteGetMasterSecretOperation) *xxx_RemoteGetMasterSecretOperation {
	if op == nil {
		op = &xxx_RemoteGetMasterSecretOperation{}
	}
	if o == nil {
		return op
	}
	op.SecretLength = o.SecretLength
	op.Secret = o.Secret
	op.Return = o.Return
	return op
}

func (o *RemoteGetMasterSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteGetMasterSecretOperation) {
	if o == nil {
		return
	}
	o.SecretLength = op.SecretLength
	o.Secret = op.Secret
	o.Return = op.Return
}
func (o *RemoteGetMasterSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteGetMasterSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteGetMasterSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
