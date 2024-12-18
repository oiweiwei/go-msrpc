package rfri

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
	GoPackage = "oxabref"
)

var (
	// Syntax UUID
	RfriSyntaxUUID = &uuid.UUID{TimeLow: 0x1544f5e0, TimeMid: 0x613c, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0xdf, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd7, 0xbd, 0x9}}
	// Syntax ID
	RfriSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RfriSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// rfri interface.
type RfriClient interface {

	// The RfrGetNewDSA method returns the name of an NSPI server or a server array.
	//
	// Return Values: The server returns 0 for a successful execution. An error results
	// in an HRESULT or other nonzero error code.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	GetNewDSA(context.Context, *GetNewDSARequest, ...dcerpc.CallOption) (*GetNewDSAResponse, error)

	// The RfrGetFQDNFromServerDN method returns the Domain Name System (DNS) FQDN of the
	// server corresponding to the passed DN.
	//
	// Return Values: The server returns 0 for a successful execution. An error results
	// in an HRESULT or other nonzero error code.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	GetFQDNFromServerDN(context.Context, *GetFQDNFromServerDNRequest, ...dcerpc.CallOption) (*GetFQDNFromServerDNResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultRfriClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRfriClient) GetNewDSA(ctx context.Context, in *GetNewDSARequest, opts ...dcerpc.CallOption) (*GetNewDSAResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNewDSAResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRfriClient) GetFQDNFromServerDN(ctx context.Context, in *GetFQDNFromServerDNRequest, opts ...dcerpc.CallOption) (*GetFQDNFromServerDNResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFQDNFromServerDNResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRfriClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRfriClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRfriClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RfriClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RfriSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRfriClient{cc: cc}, nil
}

// xxx_GetNewDSAOperation structure represents the RfrGetNewDSA operation
type xxx_GetNewDSAOperation struct {
	Flags  uint32 `idl:"name:ulFlags" json:"flags"`
	UserDN string `idl:"name:pUserDN;string" json:"user_dn"`
	_      string `idl:"name:ppszUnused;string;pointer:unique"`
	Server string `idl:"name:ppszServer;string;pointer:unique" json:"server"`
	Return int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNewDSAOperation) OpNum() int { return 0 }

func (o *xxx_GetNewDSAOperation) OpName() string { return "/rfri/v1/RfrGetNewDSA" }

func (o *xxx_GetNewDSAOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewDSAOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pUserDN {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.UserDN); err != nil {
			return err
		}
	}
	// ppszUnused {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		// reserved ppszUnused
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppszServer {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		if o.Server != "" {
			_ptr_ppszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Server != "" {
					_ptr_ppszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteCharNString(ctx, w, o.Server); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Server, _ptr_ppszServer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_ppszServer); err != nil {
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

func (o *xxx_GetNewDSAOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pUserDN {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.UserDN); err != nil {
			return err
		}
	}
	// ppszUnused {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		// reserved ppszUnused
		var _ppszUnused string
		_ptr_ppszUnused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppszUnused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharNString(ctx, w, &_ppszUnused); err != nil {
					return err
				}
				return nil
			})
			_s_ppszUnused := func(ptr interface{}) { _ppszUnused = *ptr.(*string) }
			if err := w.ReadPointer(&_ppszUnused, _s_ppszUnused, _ptr_ppszUnused); err != nil {
				return err
			}
			return nil
		})
		_s_ppszUnused := func(ptr interface{}) { _ppszUnused = *ptr.(*string) }
		if err := w.ReadPointer(&_ppszUnused, _s_ppszUnused, _ptr_ppszUnused); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppszServer {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		_ptr_ppszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharNString(ctx, w, &o.Server); err != nil {
					return err
				}
				return nil
			})
			_s_ppszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
			if err := w.ReadPointer(&o.Server, _s_ppszServer, _ptr_ppszServer); err != nil {
				return err
			}
			return nil
		})
		_s_ppszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_ppszServer, _ptr_ppszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewDSAOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewDSAOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppszUnused {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		// reserved ppszUnused
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppszServer {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		if o.Server != "" {
			_ptr_ppszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Server != "" {
					_ptr_ppszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if err := ndr.WriteCharNString(ctx, w, o.Server); err != nil {
							return err
						}
						return nil
					})
					if err := w.WritePointer(&o.Server, _ptr_ppszServer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_ppszServer); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNewDSAOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppszUnused {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		// reserved ppszUnused
		var _ppszUnused string
		_ptr_ppszUnused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppszUnused := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharNString(ctx, w, &_ppszUnused); err != nil {
					return err
				}
				return nil
			})
			_s_ppszUnused := func(ptr interface{}) { _ppszUnused = *ptr.(*string) }
			if err := w.ReadPointer(&_ppszUnused, _s_ppszUnused, _ptr_ppszUnused); err != nil {
				return err
			}
			return nil
		})
		_s_ppszUnused := func(ptr interface{}) { _ppszUnused = *ptr.(*string) }
		if err := w.ReadPointer(&_ppszUnused, _s_ppszUnused, _ptr_ppszUnused); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppszServer {in, out} (1:{string, pointer=unique}*(2)*(1)[dim:0,string,null](uchar))
	{
		_ptr_ppszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadCharNString(ctx, w, &o.Server); err != nil {
					return err
				}
				return nil
			})
			_s_ppszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
			if err := w.ReadPointer(&o.Server, _s_ppszServer, _ptr_ppszServer); err != nil {
				return err
			}
			return nil
		})
		_s_ppszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_ppszServer, _ptr_ppszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNewDSARequest structure represents the RfrGetNewDSA operation request
type GetNewDSARequest struct {
	// ulFlags: An unsigned long value, containing a set of bit flags. Unused; SHOULD be
	// set to zero. Other values MUST be ignored by the server.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// pUserDN: Optional, a DN indicating the mailbox owned by the client user. The client
	// SHOULD pass this to the server. If supplied, the server SHOULD use that DN to affect
	// which NSPI server is returned to the caller.
	UserDN string `idl:"name:pUserDN;string" json:"user_dn"`
	// ppszServer: A string. If the server does not return an error, ppszServer contains
	// the FQDN of an NSPI server or a server array. On failure, the value is undefined.
	Server string `idl:"name:ppszServer;string;pointer:unique" json:"server"`
}

func (o *GetNewDSARequest) xxx_ToOp(ctx context.Context) *xxx_GetNewDSAOperation {
	if o == nil {
		return &xxx_GetNewDSAOperation{}
	}
	return &xxx_GetNewDSAOperation{
		Flags:  o.Flags,
		UserDN: o.UserDN,
		Server: o.Server,
	}
}

func (o *GetNewDSARequest) xxx_FromOp(ctx context.Context, op *xxx_GetNewDSAOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.UserDN = op.UserDN
	o.Server = op.Server
}
func (o *GetNewDSARequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNewDSARequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNewDSAOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNewDSAResponse structure represents the RfrGetNewDSA operation response
type GetNewDSAResponse struct {
	// ppszServer: A string. If the server does not return an error, ppszServer contains
	// the FQDN of an NSPI server or a server array. On failure, the value is undefined.
	Server string `idl:"name:ppszServer;string;pointer:unique" json:"server"`
	// Return: The RfrGetNewDSA return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNewDSAResponse) xxx_ToOp(ctx context.Context) *xxx_GetNewDSAOperation {
	if o == nil {
		return &xxx_GetNewDSAOperation{}
	}
	return &xxx_GetNewDSAOperation{
		Server: o.Server,
		Return: o.Return,
	}
}

func (o *GetNewDSAResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNewDSAOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Return = op.Return
}
func (o *GetNewDSAResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNewDSAResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNewDSAOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFQDNFromServerDNOperation structure represents the RfrGetFQDNFromServerDN operation
type xxx_GetFQDNFromServerDNOperation struct {
	Flags                 uint32 `idl:"name:ulFlags" json:"flags"`
	MailboxServerDNLength uint32 `idl:"name:cbMailboxServerDN" json:"mailbox_server_dn_length"`
	MailboxServerDN       string `idl:"name:szMailboxServerDN;size_is:(cbMailboxServerDN);string" json:"mailbox_server_dn"`
	ServerFQDN            string `idl:"name:ppszServerFQDN;string;pointer:ref" json:"server_fqdn"`
	Return                int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFQDNFromServerDNOperation) OpNum() int { return 1 }

func (o *xxx_GetFQDNFromServerDNOperation) OpName() string { return "/rfri/v1/RfrGetFQDNFromServerDN" }

func (o *xxx_GetFQDNFromServerDNOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MailboxServerDN != "" && o.MailboxServerDNLength == 0 {
		o.MailboxServerDNLength = uint32(len(o.MailboxServerDN))
	}
	if o.MailboxServerDNLength < uint32(10) || o.MailboxServerDNLength > uint32(1024) {
		return fmt.Errorf("MailboxServerDNLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFQDNFromServerDNOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// cbMailboxServerDN {in} (1:{range=(10,1024)}(uint32))
	{
		if err := w.WriteData(o.MailboxServerDNLength); err != nil {
			return err
		}
	}
	// szMailboxServerDN {in} (1:{string, pointer=ref}*(1)[dim:0,size_is=cbMailboxServerDN,string,null](uchar))
	{
		dimSize1 := uint64(o.MailboxServerDNLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.CharNLen(o.MailboxServerDN)
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
		_MailboxServerDN_buf := []byte(o.MailboxServerDN)
		if uint64(len(_MailboxServerDN_buf)) > sizeInfo[0]-1 {
			_MailboxServerDN_buf = _MailboxServerDN_buf[:sizeInfo[0]-1]
		}
		if o.MailboxServerDN != ndr.ZeroString {
			_MailboxServerDN_buf = append(_MailboxServerDN_buf, byte(0))
		}
		for i1 := range _MailboxServerDN_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_MailboxServerDN_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_MailboxServerDN_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetFQDNFromServerDNOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// cbMailboxServerDN {in} (1:{range=(10,1024)}(uint32))
	{
		if err := w.ReadData(&o.MailboxServerDNLength); err != nil {
			return err
		}
	}
	// szMailboxServerDN {in} (1:{string, pointer=ref}*(1)[dim:0,size_is=cbMailboxServerDN,string,null](uchar))
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
		var _MailboxServerDN_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _MailboxServerDN_buf", sizeInfo[0])
		}
		_MailboxServerDN_buf = make([]byte, sizeInfo[0])
		for i1 := range _MailboxServerDN_buf {
			i1 := i1
			if err := w.ReadData(&_MailboxServerDN_buf[i1]); err != nil {
				return err
			}
		}
		o.MailboxServerDN = strings.TrimRight(string(_MailboxServerDN_buf), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_GetFQDNFromServerDNOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFQDNFromServerDNOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppszServerFQDN {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		if o.ServerFQDN != "" {
			_ptr_ppszServerFQDN := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteCharNString(ctx, w, o.ServerFQDN); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerFQDN, _ptr_ppszServerFQDN); err != nil {
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
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFQDNFromServerDNOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppszServerFQDN {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](uchar))
	{
		_ptr_ppszServerFQDN := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadCharNString(ctx, w, &o.ServerFQDN); err != nil {
				return err
			}
			return nil
		})
		_s_ppszServerFQDN := func(ptr interface{}) { o.ServerFQDN = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerFQDN, _s_ppszServerFQDN, _ptr_ppszServerFQDN); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetFQDNFromServerDNRequest structure represents the RfrGetFQDNFromServerDN operation request
type GetFQDNFromServerDNRequest struct {
	// ulFlags: An unsigned long value, containing a set of bit flags. Unused; SHOULD be
	// set to zero. Other values MUST be ignored by the server.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// cbMailboxServerDN: An unsigned long value containing the number of bytes in the value
	// of the szMailboxServerDN parameter, including terminating NULL character. The value
	// is at least 10, at most 1024.
	MailboxServerDNLength uint32 `idl:"name:cbMailboxServerDN" json:"mailbox_server_dn_length"`
	// szMailboxServerDN: A 5 or 6-element DN identifying a mailbox server, which MUST match
	// the server's implementation of server identities. It follows this format:
	//
	// The CN=" instance-name " element is optional.<3>
	//
	// Note that the client MAY receive a DN identifying a specific database on this server,
	// from sources listed in section 1.6. This DN follows this format:
	//
	// # Or
	//
	// If this is the DN available, it is the client's responsibility to remove the final
	// element before passing the DN to the RfrGetFQDNFromServerDN method.
	MailboxServerDN string `idl:"name:szMailboxServerDN;size_is:(cbMailboxServerDN);string" json:"mailbox_server_dn"`
}

func (o *GetFQDNFromServerDNRequest) xxx_ToOp(ctx context.Context) *xxx_GetFQDNFromServerDNOperation {
	if o == nil {
		return &xxx_GetFQDNFromServerDNOperation{}
	}
	return &xxx_GetFQDNFromServerDNOperation{
		Flags:                 o.Flags,
		MailboxServerDNLength: o.MailboxServerDNLength,
		MailboxServerDN:       o.MailboxServerDN,
	}
}

func (o *GetFQDNFromServerDNRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFQDNFromServerDNOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.MailboxServerDNLength = op.MailboxServerDNLength
	o.MailboxServerDN = op.MailboxServerDN
}
func (o *GetFQDNFromServerDNRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFQDNFromServerDNRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFQDNFromServerDNOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFQDNFromServerDNResponse structure represents the RfrGetFQDNFromServerDN operation response
type GetFQDNFromServerDNResponse struct {
	// ppszServerFQDN: A string. If the server does not return an error, the ppszServerFQDN
	// parameter contains the FQDN of the mailbox server identified by the szMailboxServerDN
	// parameter.
	ServerFQDN string `idl:"name:ppszServerFQDN;string;pointer:ref" json:"server_fqdn"`
	// Return: The RfrGetFQDNFromServerDN return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFQDNFromServerDNResponse) xxx_ToOp(ctx context.Context) *xxx_GetFQDNFromServerDNOperation {
	if o == nil {
		return &xxx_GetFQDNFromServerDNOperation{}
	}
	return &xxx_GetFQDNFromServerDNOperation{
		ServerFQDN: o.ServerFQDN,
		Return:     o.Return,
	}
}

func (o *GetFQDNFromServerDNResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFQDNFromServerDNOperation) {
	if o == nil {
		return
	}
	o.ServerFQDN = op.ServerFQDN
	o.Return = op.Return
}
func (o *GetFQDNFromServerDNResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFQDNFromServerDNResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFQDNFromServerDNOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
