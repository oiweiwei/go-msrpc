package backupkey

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "bkrp"
)

var (
	// Syntax UUID
	BackupKeySyntaxUUID = &uuid.UUID{TimeLow: 0x3dde7c30, TimeMid: 0x165d, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x8f, Node: [6]uint8{0x0, 0x80, 0x5f, 0x14, 0xdb, 0x40}}
	// Syntax ID
	BackupKeySyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: BackupKeySyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// BackupKey interface.
type BackupKeyClient interface {

	// This section specifies the BackuprKey method.
	//
	// Return Values: The server MUST return 0 if it successfully processes the message
	// received from the client, and a nonzero value otherwise.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	BackupKey(context.Context, *BackupKeyRequest, ...dcerpc.CallOption) (*BackupKeyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultBackupKeyClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultBackupKeyClient) BackupKey(ctx context.Context, in *BackupKeyRequest, opts ...dcerpc.CallOption) (*BackupKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultBackupKeyClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultBackupKeyClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewBackupKeyClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (BackupKeyClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(BackupKeySyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultBackupKeyClient{cc: cc}, nil
}

// xxx_BackupKeyOperation structure represents the BackuprKey operation
type xxx_BackupKeyOperation struct {
	ActionAgent   *dtyp.GUID `idl:"name:pguidActionAgent" json:"action_agent"`
	DataIn        []byte     `idl:"name:pDataIn;size_is:(cbDataIn)" json:"data_in"`
	DataInLength  uint32     `idl:"name:cbDataIn" json:"data_in_length"`
	DataOut       []byte     `idl:"name:ppDataOut;size_is:(, pcbDataOut)" json:"data_out"`
	DataOutLength uint32     `idl:"name:pcbDataOut" json:"data_out_length"`
	Param         uint32     `idl:"name:dwParam" json:"param"`
	Return        uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupKeyOperation) OpNum() int { return 0 }

func (o *xxx_BackupKeyOperation) OpName() string { return "/BackupKey/v1/BackuprKey" }

func (o *xxx_BackupKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DataIn != nil && o.DataInLength == 0 {
		o.DataInLength = uint32(len(o.DataIn))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pguidActionAgent {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ActionAgent != nil {
			if err := o.ActionAgent.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pDataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
	{
		dimSize1 := uint64(o.DataInLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DataIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DataIn[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DataIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbDataIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataInLength); err != nil {
			return err
		}
	}
	// dwParam {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Param); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pguidActionAgent {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ActionAgent == nil {
			o.ActionAgent = &dtyp.GUID{}
		}
		if err := o.ActionAgent.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pDataIn {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbDataIn](uint8))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DataIn", sizeInfo[0])
		}
		o.DataIn = make([]byte, sizeInfo[0])
		for i1 := range o.DataIn {
			i1 := i1
			if err := w.ReadData(&o.DataIn[i1]); err != nil {
				return err
			}
		}
	}
	// cbDataIn {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataInLength); err != nil {
			return err
		}
	}
	// dwParam {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Param); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DataOut != nil && o.DataOutLength == 0 {
		o.DataOutLength = uint32(len(o.DataOut))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppDataOut {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbDataOut](uint8))
	{
		if o.DataOut != nil || o.DataOutLength > 0 {
			_ptr_ppDataOut := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataOutLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.DataOut {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.DataOut[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.DataOut); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DataOut, _ptr_ppDataOut); err != nil {
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
	// pcbDataOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataOutLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppDataOut {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbDataOut](uint8))
	{
		_ptr_ppDataOut := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.DataOut", sizeInfo[0])
			}
			o.DataOut = make([]byte, sizeInfo[0])
			for i1 := range o.DataOut {
				i1 := i1
				if err := w.ReadData(&o.DataOut[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppDataOut := func(ptr interface{}) { o.DataOut = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.DataOut, _s_ppDataOut, _ptr_ppDataOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbDataOut {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataOutLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupKeyRequest structure represents the BackuprKey operation request
type BackupKeyRequest struct {
	// pguidActionAgent: A GUID RPC structure, as specified in [MS-DTYP] section 2.3.4.
	// This MUST be set to one of the following values. The BACKUPKEY_BACKUP_GUID and BACKUPKEY_RESTORE_GUID_WIN2K
	// values indicate the ServerWrap subprotocol, while the BACKUPKEY_RETRIEVE_BACKUP_KEY_GUID
	// and BACKUPKEY_RESTORE_GUID values indicate the ClientWrap subprotocol. A server MUST
	// support at least one of these subprotocols completely, and all server implementations
	// SHOULD support all four values. In addition, if a server supports the wrapping operation
	// of either subprotocol, it MUST also support the corresponding unwrap operation. Thus,
	// if a server supports BACKUPKEY_BACKUP_GUID, then it MUST also support BACKUPKEY_RESTORE_GUID_WIN2K.
	// Similarly, if a server supports BACKUPKEY_RETRIEVE_BACKUP_KEY_GUID, it MUST also
	// support BACKUPKEY_RESTORE_GUID.<9>
	//
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                                         |                                                                                  |
	//	|                                  VALUE                                  |                                     MEANING                                      |
	//	|                                                                         |                                                                                  |
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| BACKUPKEY_BACKUP_GUID 7F752B10-178E-11D1-AB8F-00805F14DB40              | Requests server-side wrapping. On input, pDataIn MUST point to a BLOB containing |
	//	|                                                                         | the secret to be wrapped. The server MUST treat pDataIn as opaque binary data.   |
	//	|                                                                         | On output, ppDataOut MUST contain the wrapped secret in the format specified in  |
	//	|                                                                         | section 2.2.4. For details, see section 3.1.4.1.1.                               |
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| BACKUPKEY_RESTORE_GUID_WIN2K 7FE94D50-178E-11D1-AB8F-00805F14DB40       | Requests unwrapping of a server-side-wrapped secret. On input, pDataIn MUST      |
	//	|                                                                         | point to a BLOB containing the wrapped key, in the format specified in section   |
	//	|                                                                         | 2.2.4. On output, ppDataOut MUST contain a pointer to the unwrapped secret,      |
	//	|                                                                         | as supplied by the client to the BACKUPKEY_BACKUP_GUID call. For details, see    |
	//	|                                                                         | section 3.1.4.1.2.                                                               |
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| BACKUPKEY_RETRIEVE_BACKUP_KEY_GUID 018FF48A-EABA-40C6-8F6D-72370240E967 | Requests the public key part of the server's ClientWrap key pair. The server     |
	//	|                                                                         | MUST ignore the pDataIn and cbDataIn parameters. On output, ppDataOut MUST       |
	//	|                                                                         | contain a pointer to the server's public key in the format specified in section  |
	//	|                                                                         | 2.2.1. For details, see section 3.1.4.1.3.                                       |
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| BACKUPKEY_RESTORE_GUID 47270C64-2FC7-499B-AC5B-0E37CDCE899A             | Request unwrapping of a secret that was client-side-wrapped with the server's    |
	//	|                                                                         | public key. On input, pDataIn MUST point to a client-side wrapped key, formatted |
	//	|                                                                         | as specified in section 2.2.2. On output, ppDataOut MUST contain a pointer to    |
	//	|                                                                         | the unwrapped secret, formatted as specified in section 2.2.3. For details, see  |
	//	|                                                                         | section 3.1.4.1.4.                                                               |
	//	+-------------------------------------------------------------------------+----------------------------------------------------------------------------------+
	ActionAgent *dtyp.GUID `idl:"name:pguidActionAgent" json:"action_agent"`
	// pDataIn: This is the input data supplied by the client. Its format depends on pguidActionAgent
	// as specified in this section.
	DataIn []byte `idl:"name:pDataIn;size_is:(cbDataIn)" json:"data_in"`
	// cbDataIn: This MUST be an unsigned 32-bit integer, encoded in little-endian format.
	// It MUST be equal to the length, in bytes, of the data supplied in pDataIn.
	DataInLength uint32 `idl:"name:cbDataIn" json:"data_in_length"`
	// dwParam: This parameter is unused. It MUST be ignored by the server.
	Param uint32 `idl:"name:dwParam" json:"param"`
}

func (o *BackupKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BackupKeyOperation) *xxx_BackupKeyOperation {
	if op == nil {
		op = &xxx_BackupKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.ActionAgent = o.ActionAgent
	op.DataIn = o.DataIn
	op.DataInLength = o.DataInLength
	op.Param = o.Param
	return op
}

func (o *BackupKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupKeyOperation) {
	if o == nil {
		return
	}
	o.ActionAgent = op.ActionAgent
	o.DataIn = op.DataIn
	o.DataInLength = op.DataInLength
	o.Param = op.Param
}
func (o *BackupKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BackupKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupKeyResponse structure represents the BackuprKey operation response
type BackupKeyResponse struct {
	// ppDataOut: This is the output data returned to the client. Its format depends on
	// pguidActionAgent as specified in this section.
	DataOut []byte `idl:"name:ppDataOut;size_is:(, pcbDataOut)" json:"data_out"`
	// pcbDataOut: This MUST be an unsigned 32-bit integer, encoded in little-endian format.
	// It MUST be equal to the length, in bytes, of the data returned in pDataOut.
	DataOutLength uint32 `idl:"name:pcbDataOut" json:"data_out_length"`
	// Return: The BackuprKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BackupKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BackupKeyOperation) *xxx_BackupKeyOperation {
	if op == nil {
		op = &xxx_BackupKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.DataOut = o.DataOut
	op.DataOutLength = o.DataOutLength
	op.Return = o.Return
	return op
}

func (o *BackupKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupKeyOperation) {
	if o == nil {
		return
	}
	o.DataOut = op.DataOut
	o.DataOutLength = op.DataOutLength
	o.Return = op.Return
}
func (o *BackupKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BackupKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
