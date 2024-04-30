package isdkey

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "gkdi"
)

var (
	// Syntax UUID
	ISDKeySyntaxUUID = &uuid.UUID{TimeLow: 0xb9785960, TimeMid: 0x524f, TimeHiAndVersion: 0x11df, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x6d, Node: [6]uint8{0x83, 0xdc, 0xde, 0xd7, 0x20, 0x85}}
	// Syntax ID
	ISDKeySyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: ISDKeySyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// ISDKey interface.
type ISDKeyClient interface {

	// The syntax for the GetKey (Opnum 0) method consists of the following.
	//
	// Return Values: The server MUST return zero if it successfully processes the message
	// received from the client; otherwise, it MUST return a nonzero value.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetKey(context.Context, *GetKeyRequest, ...dcerpc.CallOption) (*GetKeyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

type xxx_DefaultISDKeyClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultISDKeyClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...dcerpc.CallOption) (*GetKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultISDKeyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewISDKeyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ISDKeyClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ISDKeySyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultISDKeyClient{cc: cc}, nil
}

// xxx_GetKeyOperation structure represents the GetKey operation
type xxx_GetKeyOperation struct {
	TargetSDLength uint32     `idl:"name:cbTargetSD" json:"target_sd_length"`
	TargetSD       []byte     `idl:"name:pbTargetSD;size_is:(cbTargetSD);pointer:ref" json:"target_sd"`
	RootKeyID      *dtyp.GUID `idl:"name:pRootKeyID;pointer:unique" json:"root_key_id"`
	L0KeyID        int32      `idl:"name:L0KeyID" json:"l0_key_id"`
	L1KeyID        int32      `idl:"name:L1KeyID" json:"l1_key_id"`
	L2KeyID        int32      `idl:"name:L2KeyID" json:"l2_key_id"`
	OutLength      uint32     `idl:"name:pcbOut" json:"out_length"`
	Out            []byte     `idl:"name:ppbOut;size_is:(, pcbOut)" json:"out"`
	Return         int32      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetKeyOperation) OpNum() int { return 0 }

func (o *xxx_GetKeyOperation) OpName() string { return "/ISDKey/v1/GetKey" }

func (o *xxx_GetKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.TargetSD != nil && o.TargetSDLength == 0 {
		o.TargetSDLength = uint32(len(o.TargetSD))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// cbTargetSD {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TargetSDLength); err != nil {
			return err
		}
	}
	// pbTargetSD {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTargetSD](char))
	{
		dimSize1 := uint64(o.TargetSDLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.TargetSD {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.TargetSD[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.TargetSD); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// pRootKeyID {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		if o.RootKeyID != nil {
			_ptr_pRootKeyID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootKeyID != nil {
					if err := o.RootKeyID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootKeyID, _ptr_pRootKeyID); err != nil {
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
	// L0KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.L0KeyID); err != nil {
			return err
		}
	}
	// L1KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.L1KeyID); err != nil {
			return err
		}
	}
	// L2KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.L2KeyID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// cbTargetSD {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TargetSDLength); err != nil {
			return err
		}
	}
	// pbTargetSD {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTargetSD](char))
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
			return fmt.Errorf("buffer overflow for size %d of array o.TargetSD", sizeInfo[0])
		}
		o.TargetSD = make([]byte, sizeInfo[0])
		for i1 := range o.TargetSD {
			i1 := i1
			if err := w.ReadData(&o.TargetSD[i1]); err != nil {
				return err
			}
		}
	}
	// pRootKeyID {in} (1:{pointer=unique}*(1))(2:{alias=GUID}(struct))
	{
		_ptr_pRootKeyID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RootKeyID == nil {
				o.RootKeyID = &dtyp.GUID{}
			}
			if err := o.RootKeyID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pRootKeyID := func(ptr interface{}) { o.RootKeyID = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.RootKeyID, _s_pRootKeyID, _ptr_pRootKeyID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// L0KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.L0KeyID); err != nil {
			return err
		}
	}
	// L1KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.L1KeyID); err != nil {
			return err
		}
	}
	// L2KeyID {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.L2KeyID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Out != nil && o.OutLength == 0 {
		o.OutLength = uint32(len(o.Out))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pcbOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OutLength); err != nil {
			return err
		}
	}
	// ppbOut {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbOut](uint8))
	{
		if o.Out != nil || o.OutLength > 0 {
			_ptr_ppbOut := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OutLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Out {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Out[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Out); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Out, _ptr_ppbOut); err != nil {
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

func (o *xxx_GetKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pcbOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OutLength); err != nil {
			return err
		}
	}
	// ppbOut {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbOut](uint8))
	{
		_ptr_ppbOut := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Out", sizeInfo[0])
			}
			o.Out = make([]byte, sizeInfo[0])
			for i1 := range o.Out {
				i1 := i1
				if err := w.ReadData(&o.Out[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppbOut := func(ptr interface{}) { o.Out = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Out, _s_ppbOut, _ptr_ppbOut); err != nil {
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

// GetKeyRequest structure represents the GetKey operation request
type GetKeyRequest struct {
	// cbTargetSD: This parameter is equal to the length, in bytes, of the security descriptor
	// supplied in pbTargetSD.
	TargetSDLength uint32 `idl:"name:cbTargetSD" json:"target_sd_length"`
	// pbTargetSD: This parameter is a pointer to the security descriptor for which the
	// group key is being requested.
	TargetSD []byte `idl:"name:pbTargetSD;size_is:(cbTargetSD);pointer:ref" json:"target_sd"`
	// pRootKeyID: This parameter represents the root key identifier of the requested key.
	// It can be set to NULL.
	RootKeyID *dtyp.GUID `idl:"name:pRootKeyID;pointer:unique" json:"root_key_id"`
	// L0KeyID: This parameter represents the L0 index of the requested group key. It MUST
	// be a signed 32-bit integer greater than or equal to -1.
	L0KeyID int32 `idl:"name:L0KeyID" json:"l0_key_id"`
	// L1KeyID: This parameter represents the L1 index of the requested group key. It MUST
	// be a signed 32-bit integer between -1 and 31 (inclusive).
	L1KeyID int32 `idl:"name:L1KeyID" json:"l1_key_id"`
	// L2KeyID: This parameter represents the L2 index of the requested group key. It MUST
	// be a 32-bit integer between -1 and 31 (inclusive).
	L2KeyID int32 `idl:"name:L2KeyID" json:"l2_key_id"`
}

func (o *GetKeyRequest) xxx_ToOp(ctx context.Context) *xxx_GetKeyOperation {
	if o == nil {
		return &xxx_GetKeyOperation{}
	}
	return &xxx_GetKeyOperation{
		TargetSDLength: o.TargetSDLength,
		TargetSD:       o.TargetSD,
		RootKeyID:      o.RootKeyID,
		L0KeyID:        o.L0KeyID,
		L1KeyID:        o.L1KeyID,
		L2KeyID:        o.L2KeyID,
	}
}

func (o *GetKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetKeyOperation) {
	if o == nil {
		return
	}
	o.TargetSDLength = op.TargetSDLength
	o.TargetSD = op.TargetSD
	o.RootKeyID = op.RootKeyID
	o.L0KeyID = op.L0KeyID
	o.L1KeyID = op.L1KeyID
	o.L2KeyID = op.L2KeyID
}
func (o *GetKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetKeyResponse structure represents the GetKey operation response
type GetKeyResponse struct {
	// pcbOut: This parameter is an unsigned, 32-bit integer. It MUST be equal to the length,
	// in bytes, of the data returned in ppbOut.
	OutLength uint32 `idl:"name:pcbOut" json:"out_length"`
	// ppbOut: On successful processing of a request, the server MUST set this to a pointer
	// that refers to the output key binary large object (BLOB), as specified in section
	// 2.2.4.
	Out []byte `idl:"name:ppbOut;size_is:(, pcbOut)" json:"out"`
	// Return: The GetKey return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetKeyResponse) xxx_ToOp(ctx context.Context) *xxx_GetKeyOperation {
	if o == nil {
		return &xxx_GetKeyOperation{}
	}
	return &xxx_GetKeyOperation{
		OutLength: o.OutLength,
		Out:       o.Out,
		Return:    o.Return,
	}
}

func (o *GetKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetKeyOperation) {
	if o == nil {
		return
	}
	o.OutLength = op.OutLength
	o.Out = op.Out
	o.Return = op.Return
}
func (o *GetKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
