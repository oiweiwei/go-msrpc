package bitspeerauth

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

var (
	// import guard
	GoPackage = "bpau"
)

var (
	// Syntax UUID
	BitsPeerAuthSyntaxUUID = &uuid.UUID{TimeLow: 0xe3d0d746, TimeMid: 0xd2af, TimeHiAndVersion: 0x40fd, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x7a, Node: [6]uint8{0xd, 0x70, 0x78, 0xbb, 0x70, 0x92}}
	// Syntax ID
	BitsPeerAuthSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: BitsPeerAuthSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// BitsPeerAuth interface.
type BitsPeerAuthClient interface {

	// Return Values: An HRESULT indicating return status. See [MS-ERREF] for details of
	// the HRESULT type.
	//
	// ERROR_SUCCESS (0x00000000)
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ExchangePublicKeys(context.Context, *ExchangePublicKeysRequest, ...dcerpc.CallOption) (*ExchangePublicKeysResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultBitsPeerAuthClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultBitsPeerAuthClient) ExchangePublicKeys(ctx context.Context, in *ExchangePublicKeysRequest, opts ...dcerpc.CallOption) (*ExchangePublicKeysResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ExchangePublicKeysResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBitsPeerAuthClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBitsPeerAuthClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewBitsPeerAuthClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BitsPeerAuthClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BitsPeerAuthSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultBitsPeerAuthClient{cc: cc}, nil
}

// xxx_ExchangePublicKeysOperation structure represents the ExchangePublicKeys operation
type xxx_ExchangePublicKeysOperation struct {
	ClientKeyLength uint32 `idl:"name:ClientKeyLength" json:"client_key_length"`
	ClientKey       []byte `idl:"name:ClientKey;size_is:(ClientKeyLength);pointer:unique" json:"client_key"`
	ServerKeyLength uint32 `idl:"name:pServerKeyLength;pointer:ref" json:"server_key_length"`
	ServerKey       []byte `idl:"name:pServerKey;size_is:(, pServerKeyLength);pointer:ref" json:"server_key"`
	Return          int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_ExchangePublicKeysOperation) OpNum() int { return 0 }

func (o *xxx_ExchangePublicKeysOperation) OpName() string {
	return "/BitsPeerAuth/v1/ExchangePublicKeys"
}

func (o *xxx_ExchangePublicKeysOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ClientKey != nil && o.ClientKeyLength == 0 {
		o.ClientKeyLength = uint32(len(o.ClientKey))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExchangePublicKeysOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ClientKeyLength {in} (1:{range=(0,65536), alias=KEY_LENGTH, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientKeyLength); err != nil {
			return err
		}
	}
	// ClientKey {in} (1:{pointer=unique}*(1)[dim:0,size_is=ClientKeyLength](uint8))
	{
		if o.ClientKey != nil || o.ClientKeyLength > 0 {
			_ptr_ClientKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClientKeyLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ClientKey {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ClientKey[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ClientKey); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientKey, _ptr_ClientKey); err != nil {
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

func (o *xxx_ExchangePublicKeysOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ClientKeyLength {in} (1:{range=(0,65536), alias=KEY_LENGTH, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientKeyLength); err != nil {
			return err
		}
	}
	// ClientKey {in} (1:{pointer=unique}*(1)[dim:0,size_is=ClientKeyLength](uint8))
	{
		_ptr_ClientKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ClientKey", sizeInfo[0])
			}
			o.ClientKey = make([]byte, sizeInfo[0])
			for i1 := range o.ClientKey {
				i1 := i1
				if err := w.ReadData(&o.ClientKey[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ClientKey := func(ptr interface{}) { o.ClientKey = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ClientKey, _s_ClientKey, _ptr_ClientKey); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExchangePublicKeysOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ServerKey != nil && o.ServerKeyLength == 0 {
		o.ServerKeyLength = uint32(len(o.ServerKey))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExchangePublicKeysOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pServerKeyLength {out} (1:{pointer=ref}*(1))(2:{range=(0,65536), alias=KEY_LENGTH, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerKeyLength); err != nil {
			return err
		}
	}
	// pServerKey {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pServerKeyLength](uint8))
	{
		if o.ServerKey != nil || o.ServerKeyLength > 0 {
			_ptr_pServerKey := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ServerKeyLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ServerKey {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.ServerKey[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.ServerKey); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerKey, _ptr_pServerKey); err != nil {
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

func (o *xxx_ExchangePublicKeysOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pServerKeyLength {out} (1:{pointer=ref}*(1))(2:{range=(0,65536), alias=KEY_LENGTH, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerKeyLength); err != nil {
			return err
		}
	}
	// pServerKey {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pServerKeyLength](uint8))
	{
		_ptr_pServerKey := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ServerKey", sizeInfo[0])
			}
			o.ServerKey = make([]byte, sizeInfo[0])
			for i1 := range o.ServerKey {
				i1 := i1
				if err := w.ReadData(&o.ServerKey[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pServerKey := func(ptr interface{}) { o.ServerKey = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.ServerKey, _s_pServerKey, _ptr_pServerKey); err != nil {
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

// ExchangePublicKeysRequest structure represents the ExchangePublicKeys operation request
type ExchangePublicKeysRequest struct {
	// ClientKeyLength: Length of the client's local certificate, or zero if ClientKey is
	// NULL.
	ClientKeyLength uint32 `idl:"name:ClientKeyLength" json:"client_key_length"`
	// ClientKey: The client's local certificate, encoded as a CERTIFICATE_BLOB (section
	// 2.2.2). If NULL, the client is choosing not to send a certificate.
	ClientKey []byte `idl:"name:ClientKey;size_is:(ClientKeyLength);pointer:unique" json:"client_key"`
}

func (o *ExchangePublicKeysRequest) xxx_ToOp(ctx context.Context, op *xxx_ExchangePublicKeysOperation) *xxx_ExchangePublicKeysOperation {
	if op == nil {
		op = &xxx_ExchangePublicKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.ClientKeyLength = o.ClientKeyLength
	op.ClientKey = o.ClientKey
	return op
}

func (o *ExchangePublicKeysRequest) xxx_FromOp(ctx context.Context, op *xxx_ExchangePublicKeysOperation) {
	if o == nil {
		return
	}
	o.ClientKeyLength = op.ClientKeyLength
	o.ClientKey = op.ClientKey
}
func (o *ExchangePublicKeysRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExchangePublicKeysRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExchangePublicKeysOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExchangePublicKeysResponse structure represents the ExchangePublicKeys operation response
type ExchangePublicKeysResponse struct {
	// pServerKeyLength: On output, the length of the server's local certificate. Set to
	// zero if the return value is nonzero, or if pServerKey is NULL.
	ServerKeyLength uint32 `idl:"name:pServerKeyLength;pointer:ref" json:"server_key_length"`
	// pServerKey: On output, the server's local certificate, encoded as a CERTIFICATE_BLOB
	// (section 2.2.2). If NULL, the server is choosing not to return a certificate. Set
	// to NULL if the return value is nonzero. Ignored on the client if the method returns
	// an error or throws an exception.
	ServerKey []byte `idl:"name:pServerKey;size_is:(, pServerKeyLength);pointer:ref" json:"server_key"`
	// Return: The ExchangePublicKeys return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExchangePublicKeysResponse) xxx_ToOp(ctx context.Context, op *xxx_ExchangePublicKeysOperation) *xxx_ExchangePublicKeysOperation {
	if op == nil {
		op = &xxx_ExchangePublicKeysOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerKeyLength = o.ServerKeyLength
	op.ServerKey = o.ServerKey
	op.Return = o.Return
	return op
}

func (o *ExchangePublicKeysResponse) xxx_FromOp(ctx context.Context, op *xxx_ExchangePublicKeysOperation) {
	if o == nil {
		return
	}
	o.ServerKeyLength = op.ServerKeyLength
	o.ServerKey = op.ServerKey
	o.Return = op.Return
}
func (o *ExchangePublicKeysResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExchangePublicKeysResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExchangePublicKeysOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
