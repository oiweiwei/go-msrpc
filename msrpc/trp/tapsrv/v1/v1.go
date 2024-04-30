package tapsrv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "trp"
)

var (
	// Syntax UUID
	TapsrvSyntaxUUID = &uuid.UUID{TimeLow: 0x2f5f6520, TimeMid: 0xca46, TimeHiAndVersion: 0x1067, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x19, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xda}}
	// Syntax ID
	TapsrvSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TapsrvSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// tapsrv interface.
type TapsrvClient interface {

	// The ClientAttach method is called by the client to establish a binding instance with
	// the server.
	//
	// Return Values: The method returns 0 on success; otherwise, it returns a nonzero error
	// code, as specified in [MS-ERREF]. The following table includes a common error code.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000048 LINEERR_OPERATIONFAILED | Generic error on the server.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                -19 | Requesting administrator access via lProcessId equals 0xFFFFFFFD (-3), but the   |
	//	|                                    | user credentials of the client do not have administrator access on the server.   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 0.
	//
	// Either the client or the server can reject unencrypted packets based on the configuration.<6>
	ClientAttach(context.Context, *ClientAttachRequest, ...dcerpc.CallOption) (*ClientAttachResponse, error)

	// The ClientRequest method is called by the client to submit requests to the server.
	//
	// Return Values: This method has no return values. However, the status of the request
	// is encapsulated within the pBuffer parameter and contained in the TAPI32_MSG.Ack_ReturnValue
	// field.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 1.
	ClientRequest(context.Context, *ClientRequestRequest, ...dcerpc.CallOption) (*ClientRequestResponse, error)

	// The ClientDetach method is called by a client when it finishes using the telephony
	// resources on a server. In response, the server frees the referenced binding instance
	// and releases the allocated resources associated with the client. For connection-oriented
	// clients, the server then calls RemoteSPDetach on the client to release the allocated
	// resources created on the client side.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 2.
	ClientDetach(context.Context, *ClientDetachRequest, ...dcerpc.CallOption) (*ClientDetachResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// Type structure represents PCONTEXT_HANDLE_TYPE RPC structure.
type Type dcetypes.ContextHandle

func (o *Type) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Type) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Type) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Type) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTapsrvClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTapsrvClient) ClientAttach(ctx context.Context, in *ClientAttachRequest, opts ...dcerpc.CallOption) (*ClientAttachResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClientAttachResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTapsrvClient) ClientRequest(ctx context.Context, in *ClientRequestRequest, opts ...dcerpc.CallOption) (*ClientRequestResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClientRequestResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultTapsrvClient) ClientDetach(ctx context.Context, in *ClientDetachRequest, opts ...dcerpc.CallOption) (*ClientDetachResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClientDetachResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultTapsrvClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewTapsrvClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TapsrvClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TapsrvSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTapsrvClient{cc: cc}, nil
}

// xxx_ClientAttachOperation structure represents the ClientAttach operation
type xxx_ClientAttachOperation struct {
	Context          *Type  `idl:"name:pphContext" json:"context"`
	ProcessID        int32  `idl:"name:lProcessID" json:"process_id"`
	AsyncEventsEvent int32  `idl:"name:phAsyncEventsEvent" json:"async_events_event"`
	DomainUser       string `idl:"name:pszDomainUser;string" json:"domain_user"`
	Machine          string `idl:"name:pszMachine;string" json:"machine"`
	Return           int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_ClientAttachOperation) OpNum() int { return 0 }

func (o *xxx_ClientAttachOperation) OpName() string { return "/tapsrv/v1/ClientAttach" }

func (o *xxx_ClientAttachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientAttachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lProcessID {in} (1:(int32))
	{
		if err := w.WriteData(o.ProcessID); err != nil {
			return err
		}
	}
	// pszDomainUser {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DomainUser); err != nil {
			return err
		}
	}
	// pszMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Machine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientAttachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lProcessID {in} (1:(int32))
	{
		if err := w.ReadData(&o.ProcessID); err != nil {
			return err
		}
	}
	// pszDomainUser {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainUser); err != nil {
			return err
		}
	}
	// pszMachine {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Machine); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientAttachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientAttachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// phAsyncEventsEvent {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AsyncEventsEvent); err != nil {
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

func (o *xxx_ClientAttachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// phAsyncEventsEvent {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AsyncEventsEvent); err != nil {
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

// ClientAttachRequest structure represents the ClientAttach operation request
type ClientAttachRequest struct {
	// lProcessID: Identifier of the process on the client that generated the attach request.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFF | Client is a remote instance that wants to control the telephony devices on this  |
	//	|            | server.                                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0xFFFFFFFD | Client is a remote instance that wants to manage and administer the telephony    |
	//	|            | server.                                                                          |
	//	+------------+----------------------------------------------------------------------------------+
	ProcessID int32 `idl:"name:lProcessID" json:"process_id"`
	// pszDomainUser: Pointer to a null-terminated string indicating the user's domain account
	// name. The string is passed on the wire.
	//
	// If lProcessId equals 0xFFFFFFFF (-1), pszDomainUser MUST contain either an empty
	// string ("") or the name of a client mailslot. If a mailslot name is specified and
	// the server is able to successfully open the mailslot, then the client and the server
	// MUST use the "pull" model event scheme. The server MUST inform the client that events
	// are available for retrieval by writing a DWORD value to the client's mailslot, and
	// the client retrieves events via the ClientRequest method. If an empty string is specified
	// or the server is unable to open the client's mailslot, then the "push" model event
	// scheme MUST be used in which the server calls the client's RemoteSPEventProc method.
	DomainUser string `idl:"name:pszDomainUser;string" json:"domain_user"`
	// pszMachine: Pointer to a null-terminated string indicating the client's machine name.
	// The string MUST be passed on the wire.
	//
	// If lProcessId equals 0xFFFFFFFF (-1), the pszMachine string takes the following format:
	// <clientComputerName>"<protocolSequence1>"<endpoint1>"<protSeqN>"<epN>"\0. This allows
	// the client to inform the server of its machine name and what protocols and endpoints
	// are supported by the client on its remotesp interface. Quotation marks MUST be used
	// as delimiting tokens. For instance, CLIENT-PC-NAME"ncacn_ip_tcp"251"ncacn_nb_nb"251"\0.
	Machine string `idl:"name:pszMachine;string" json:"machine"`
}

func (o *ClientAttachRequest) xxx_ToOp(ctx context.Context) *xxx_ClientAttachOperation {
	if o == nil {
		return &xxx_ClientAttachOperation{}
	}
	return &xxx_ClientAttachOperation{
		ProcessID:  o.ProcessID,
		DomainUser: o.DomainUser,
		Machine:    o.Machine,
	}
}

func (o *ClientAttachRequest) xxx_FromOp(ctx context.Context, op *xxx_ClientAttachOperation) {
	if o == nil {
		return
	}
	o.ProcessID = op.ProcessID
	o.DomainUser = op.DomainUser
	o.Machine = op.Machine
}
func (o *ClientAttachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClientAttachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientAttachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClientAttachResponse structure represents the ClientAttach operation response
type ClientAttachResponse struct {
	// pphContext: Pointer to a PCONTEXT_HANDLE_TYPE context handle.
	Context *Type `idl:"name:pphContext" json:"context"`
	// phAsyncEventsEvent: If applicable, a pointer to a packet containing any asynchronous
	// event that was triggered by the client attaching to the server.
	//
	// If lProcessId equals 0xFFFFFFFF (-1) and the server supports the NegotiateAPIVersionForAllDevices
	// request, the server MUST set the value of phAsyncEventsEvent to 0xa5c369a5.
	AsyncEventsEvent int32 `idl:"name:phAsyncEventsEvent" json:"async_events_event"`
	// Return: The ClientAttach return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClientAttachResponse) xxx_ToOp(ctx context.Context) *xxx_ClientAttachOperation {
	if o == nil {
		return &xxx_ClientAttachOperation{}
	}
	return &xxx_ClientAttachOperation{
		Context:          o.Context,
		AsyncEventsEvent: o.AsyncEventsEvent,
		Return:           o.Return,
	}
}

func (o *ClientAttachResponse) xxx_FromOp(ctx context.Context, op *xxx_ClientAttachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.AsyncEventsEvent = op.AsyncEventsEvent
	o.Return = op.Return
}
func (o *ClientAttachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClientAttachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientAttachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClientRequestOperation structure represents the ClientRequest operation
type xxx_ClientRequestOperation struct {
	Context    *Type  `idl:"name:phContext" json:"context"`
	Buffer     []byte `idl:"name:pBuffer;size_is:(lNeededSize);length_is:(plUsedSize)" json:"buffer"`
	NeededSize int32  `idl:"name:lNeededSize" json:"needed_size"`
	UsedSize   int32  `idl:"name:plUsedSize" json:"used_size"`
}

func (o *xxx_ClientRequestOperation) OpNum() int { return 1 }

func (o *xxx_ClientRequestOperation) OpName() string { return "/tapsrv/v1/ClientRequest" }

func (o *xxx_ClientRequestOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.NeededSize == 0 {
		o.NeededSize = int32(len(o.Buffer))
	}
	if o.Buffer != nil && o.UsedSize == 0 {
		o.UsedSize = int32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientRequestOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pBuffer {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=lNeededSize,length_is=plUsedSize](uchar))
	{
		dimSize1 := uint64(o.NeededSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.UsedSize)
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
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lNeededSize {in} (1:(int32))
	{
		if err := w.WriteData(o.NeededSize); err != nil {
			return err
		}
	}
	// plUsedSize {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.UsedSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientRequestOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pBuffer {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=lNeededSize,length_is=plUsedSize](uchar))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// lNeededSize {in} (1:(int32))
	{
		if err := w.ReadData(&o.NeededSize); err != nil {
			return err
		}
	}
	// plUsedSize {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.UsedSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientRequestOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.UsedSize == 0 {
		o.UsedSize = int32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientRequestOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pBuffer {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=lNeededSize,length_is=plUsedSize](uchar))
	{
		dimSize1 := uint64(o.NeededSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.UsedSize)
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
		for i1 := range o.Buffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Buffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// plUsedSize {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.UsedSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientRequestOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pBuffer {in, out} (1:{pointer=ref}*(1)[dim:0,size_is=lNeededSize,length_is=plUsedSize](uchar))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
	}
	// plUsedSize {in, out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.UsedSize); err != nil {
			return err
		}
	}
	return nil
}

// ClientRequestRequest structure represents the ClientRequest operation request
type ClientRequestRequest struct {
	// phContext: Parameter that MUST contain the context handle of type PCONTEXT_HANDLE_TYPE.
	Context *Type `idl:"name:phContext" json:"context"`
	// pBuffer: Packet that MUST contain event packets or function calls. The packet follows
	// the structure of a TAPI32_MSG (section 2.2.5.2) packet. The Req_Func field of this
	// packet contains information about the operation to be performed on the server.
	Buffer []byte `idl:"name:pBuffer;size_is:(lNeededSize);length_is:(plUsedSize)" json:"buffer"`
	// lNeededSize: The size, in bytes, of a valid pBuffer.
	NeededSize int32 `idl:"name:lNeededSize" json:"needed_size"`
	// plUsedSize: The size, in bytes, of a valid pBuffer data. If any variable-length input
	// data is specified, both the size of the input data length and all the padding bytes
	// are included, or else all the padding bytes are excluded.
	UsedSize int32 `idl:"name:plUsedSize" json:"used_size"`
}

func (o *ClientRequestRequest) xxx_ToOp(ctx context.Context) *xxx_ClientRequestOperation {
	if o == nil {
		return &xxx_ClientRequestOperation{}
	}
	return &xxx_ClientRequestOperation{
		Context:    o.Context,
		Buffer:     o.Buffer,
		NeededSize: o.NeededSize,
		UsedSize:   o.UsedSize,
	}
}

func (o *ClientRequestRequest) xxx_FromOp(ctx context.Context, op *xxx_ClientRequestOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Buffer = op.Buffer
	o.NeededSize = op.NeededSize
	o.UsedSize = op.UsedSize
}
func (o *ClientRequestRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClientRequestRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientRequestOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClientRequestResponse structure represents the ClientRequest operation response
type ClientRequestResponse struct {
	// pBuffer: Packet that MUST contain event packets or function calls. The packet follows
	// the structure of a TAPI32_MSG (section 2.2.5.2) packet. The Req_Func field of this
	// packet contains information about the operation to be performed on the server.
	Buffer []byte `idl:"name:pBuffer;size_is:(lNeededSize);length_is:(plUsedSize)" json:"buffer"`
	// plUsedSize: The size, in bytes, of a valid pBuffer data. If any variable-length input
	// data is specified, both the size of the input data length and all the padding bytes
	// are included, or else all the padding bytes are excluded.
	UsedSize int32 `idl:"name:plUsedSize" json:"used_size"`
}

func (o *ClientRequestResponse) xxx_ToOp(ctx context.Context) *xxx_ClientRequestOperation {
	if o == nil {
		return &xxx_ClientRequestOperation{}
	}
	return &xxx_ClientRequestOperation{
		Buffer:   o.Buffer,
		UsedSize: o.UsedSize,
	}
}

func (o *ClientRequestResponse) xxx_FromOp(ctx context.Context, op *xxx_ClientRequestOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.UsedSize = op.UsedSize
}
func (o *ClientRequestResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClientRequestResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientRequestOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClientDetachOperation structure represents the ClientDetach operation
type xxx_ClientDetachOperation struct {
	Context *Type `idl:"name:pphContext" json:"context"`
}

func (o *xxx_ClientDetachOperation) OpNum() int { return 2 }

func (o *xxx_ClientDetachOperation) OpName() string { return "/tapsrv/v1/ClientDetach" }

func (o *xxx_ClientDetachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientDetachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ClientDetachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientDetachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientDetachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ClientDetachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ClientDetachRequest structure represents the ClientDetach operation request
type ClientDetachRequest struct {
	// pphContext: Pointer to a PCONTEXT_HANDLE_TYPE handle to the binding instance being
	// terminated.
	//
	// This method has no return values.
	//
	// On success, the server removes the client from the Client list.
	Context *Type `idl:"name:pphContext" json:"context"`
}

func (o *ClientDetachRequest) xxx_ToOp(ctx context.Context) *xxx_ClientDetachOperation {
	if o == nil {
		return &xxx_ClientDetachOperation{}
	}
	return &xxx_ClientDetachOperation{
		Context: o.Context,
	}
}

func (o *ClientDetachRequest) xxx_FromOp(ctx context.Context, op *xxx_ClientDetachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *ClientDetachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClientDetachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientDetachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClientDetachResponse structure represents the ClientDetach operation response
type ClientDetachResponse struct {
	// pphContext: Pointer to a PCONTEXT_HANDLE_TYPE handle to the binding instance being
	// terminated.
	//
	// This method has no return values.
	//
	// On success, the server removes the client from the Client list.
	Context *Type `idl:"name:pphContext" json:"context"`
}

func (o *ClientDetachResponse) xxx_ToOp(ctx context.Context) *xxx_ClientDetachOperation {
	if o == nil {
		return &xxx_ClientDetachOperation{}
	}
	return &xxx_ClientDetachOperation{
		Context: o.Context,
	}
}

func (o *ClientDetachResponse) xxx_FromOp(ctx context.Context, op *xxx_ClientDetachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *ClientDetachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClientDetachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientDetachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
