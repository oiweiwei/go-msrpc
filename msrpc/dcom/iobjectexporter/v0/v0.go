package iobjectexporter

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// Syntax UUID
	ObjectExporterSyntaxUUID = &uuid.UUID{TimeLow: 0x99fcfec4, TimeMid: 0x5260, TimeHiAndVersion: 0x101b, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xcb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x21, 0x34, 0x7a}}
	// Syntax ID
	ObjectExporterSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ObjectExporterSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IObjectExporter interface.
type ObjectExporterClient interface {

	// The ResolveOxid method returns the bindings and Remote Unknown IPID for an object
	// exporter.
	ResolveOXID(context.Context, *ResolveOXIDRequest, ...dcerpc.CallOption) (*ResolveOXIDResponse, error)

	// The SimplePing method performs a ping of a previously allocated ping set to maintain
	// the reference counts on the objects referred to by the set.
	SimplePing(context.Context, *SimplePingRequest, ...dcerpc.CallOption) (*SimplePingResponse, error)

	// The ComplexPing (Opnum 2) method is invoked to create or modify a ping set, to ping
	// a ping set, or to perform a combination of these operations in one invocation.
	ComplexPing(context.Context, *ComplexPingRequest, ...dcerpc.CallOption) (*ComplexPingResponse, error)

	// The ServerAlive (Opnum 3) method is used by clients to test the aliveness of the
	// server using a given RPC protocol. If it returns without an error, the server is
	// assumed to be reachable.
	ServerAlive(context.Context, *ServerAliveRequest, ...dcerpc.CallOption) (*ServerAliveResponse, error)

	// The ResolveOxid2 method returns the bindings and Remote Unknown IPID for an object
	// exporter, and the COMVERSION of the object server. This method was introduced with
	// version 5.2 of the DCOM Remote Protocol.
	ResolveOxid2(context.Context, *ResolveOxid2Request, ...dcerpc.CallOption) (*ResolveOxid2Response, error)

	// The ServerAlive2 (Opnum 5) method was introduced with version 5.6 of the DCOM Remote
	// Protocol. This method extends the ServerAlive method. It returns string and security
	// bindings for the object resolver, which allows the client to choose the most appropriate,
	// mutually compatible settings.
	ServerAlive2(context.Context, *ServerAlive2Request, ...dcerpc.CallOption) (*ServerAlive2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultObjectExporterClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultObjectExporterClient) ResolveOXID(ctx context.Context, in *ResolveOXIDRequest, opts ...dcerpc.CallOption) (*ResolveOXIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResolveOXIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) SimplePing(ctx context.Context, in *SimplePingRequest, opts ...dcerpc.CallOption) (*SimplePingResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SimplePingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) ComplexPing(ctx context.Context, in *ComplexPingRequest, opts ...dcerpc.CallOption) (*ComplexPingResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ComplexPingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) ServerAlive(ctx context.Context, in *ServerAliveRequest, opts ...dcerpc.CallOption) (*ServerAliveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerAliveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) ResolveOxid2(ctx context.Context, in *ResolveOxid2Request, opts ...dcerpc.CallOption) (*ResolveOxid2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResolveOxid2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) ServerAlive2(ctx context.Context, in *ServerAlive2Request, opts ...dcerpc.CallOption) (*ServerAlive2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerAlive2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultObjectExporterClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultObjectExporterClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewObjectExporterClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ObjectExporterClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ObjectExporterSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultObjectExporterClient{cc: cc}, nil
}

// xxx_ResolveOXIDOperation structure represents the ResolveOxid operation
type xxx_ResolveOXIDOperation struct {
	OXID                            uint64                `idl:"name:pOxid" json:"oxid"`
	RequestedProtocolSequencesCount uint16                `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	RequestedProtocolSequences      []uint16              `idl:"name:arRequestedProtseqs;size_is:(cRequestedProtseqs);pointer:ref" json:"requested_protocol_sequences"`
	OXIDBindings                    *dcom.DualStringArray `idl:"name:ppdsaOxidBindings;pointer:ref" json:"oxid_bindings"`
	RemoteUnknown                   *dcom.IPID            `idl:"name:pipidRemUnknown;pointer:ref" json:"remote_unknown"`
	AuthnHint                       uint32                `idl:"name:pAuthnHint;pointer:ref" json:"authn_hint"`
	Return                          uint32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ResolveOXIDOperation) OpNum() int { return 0 }

func (o *xxx_ResolveOXIDOperation) OpName() string { return "/IObjectExporter/v0/ResolveOxid" }

func (o *xxx_ResolveOXIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestedProtocolSequences != nil && o.RequestedProtocolSequencesCount == 0 {
		o.RequestedProtocolSequencesCount = uint16(len(o.RequestedProtocolSequences))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveOXIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pOxid {in} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.WriteData(o.OXID); err != nil {
			return err
		}
	}
	// cRequestedProtseqs {in} (1:(uint16))
	{
		if err := w.WriteData(o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// arRequestedProtseqs {in} (1:{pointer=ref}[dim:0,size_is=cRequestedProtseqs](uint16))
	{
		dimSize1 := uint64(o.RequestedProtocolSequencesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.RequestedProtocolSequences {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.RequestedProtocolSequences[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.RequestedProtocolSequences); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ResolveOXIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pOxid {in} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.ReadData(&o.OXID); err != nil {
			return err
		}
	}
	// cRequestedProtseqs {in} (1:(uint16))
	{
		if err := w.ReadData(&o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// arRequestedProtseqs {in} (1:{pointer=ref}[dim:0,size_is=cRequestedProtseqs](uint16))
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
			return fmt.Errorf("buffer overflow for size %d of array o.RequestedProtocolSequences", sizeInfo[0])
		}
		o.RequestedProtocolSequences = make([]uint16, sizeInfo[0])
		for i1 := range o.RequestedProtocolSequences {
			i1 := i1
			if err := w.ReadData(&o.RequestedProtocolSequences[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ResolveOXIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveOXIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppdsaOxidBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		if o.OXIDBindings != nil {
			_ptr_ppdsaOxidBindings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OXIDBindings != nil {
					if err := o.OXIDBindings.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.DualStringArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OXIDBindings, _ptr_ppdsaOxidBindings); err != nil {
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
	// pipidRemUnknown {out} (1:{pointer=ref}*(1))(2:{alias=IPID, names=GUID}(struct))
	{
		if o.RemoteUnknown != nil {
			if err := o.RemoteUnknown.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IPID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pAuthnHint {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AuthnHint); err != nil {
			return err
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

func (o *xxx_ResolveOXIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppdsaOxidBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		_ptr_ppdsaOxidBindings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OXIDBindings == nil {
				o.OXIDBindings = &dcom.DualStringArray{}
			}
			if err := o.OXIDBindings.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdsaOxidBindings := func(ptr interface{}) { o.OXIDBindings = *ptr.(**dcom.DualStringArray) }
		if err := w.ReadPointer(&o.OXIDBindings, _s_ppdsaOxidBindings, _ptr_ppdsaOxidBindings); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pipidRemUnknown {out} (1:{pointer=ref}*(1))(2:{alias=IPID, names=GUID}(struct))
	{
		if o.RemoteUnknown == nil {
			o.RemoteUnknown = &dcom.IPID{}
		}
		if err := o.RemoteUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pAuthnHint {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AuthnHint); err != nil {
			return err
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

// ResolveOXIDRequest structure represents the ResolveOxid operation request
type ResolveOXIDRequest struct {
	// pOxid: This MUST specify an OXID identifying an object exporter.
	OXID uint64 `idl:"name:pOxid" json:"oxid"`
	// cRequestedProtseqs: This MUST contain the number of elements in the arRequestedProtseqs
	// array.
	RequestedProtocolSequencesCount uint16 `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	// arRequestedProtseqs: This MUST contain an array of RPC protocol sequence identifiers
	// supported by the client.
	RequestedProtocolSequences []uint16 `idl:"name:arRequestedProtseqs;size_is:(cRequestedProtseqs);pointer:ref" json:"requested_protocol_sequences"`
}

func (o *ResolveOXIDRequest) xxx_ToOp(ctx context.Context, op *xxx_ResolveOXIDOperation) *xxx_ResolveOXIDOperation {
	if op == nil {
		op = &xxx_ResolveOXIDOperation{}
	}
	if o == nil {
		return op
	}
	o.OXID = op.OXID
	o.RequestedProtocolSequencesCount = op.RequestedProtocolSequencesCount
	o.RequestedProtocolSequences = op.RequestedProtocolSequences
	return op
}

func (o *ResolveOXIDRequest) xxx_FromOp(ctx context.Context, op *xxx_ResolveOXIDOperation) {
	if o == nil {
		return
	}
	o.OXID = op.OXID
	o.RequestedProtocolSequencesCount = op.RequestedProtocolSequencesCount
	o.RequestedProtocolSequences = op.RequestedProtocolSequences
}
func (o *ResolveOXIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResolveOXIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveOXIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResolveOXIDResponse structure represents the ResolveOxid operation response
type ResolveOXIDResponse struct {
	// ppdsaOxidBindings: This MUST contain the string and security bindings supported by
	// the object exporter and MUST NOT be NULL. The returned string bindings SHOULD contain
	// endpoints.
	OXIDBindings *dcom.DualStringArray `idl:"name:ppdsaOxidBindings;pointer:ref" json:"oxid_bindings"`
	// pipidRemUnknown: This MUST contain the IPID of the object exporter Remote Unknown
	// object.
	RemoteUnknown *dcom.IPID `idl:"name:pipidRemUnknown;pointer:ref" json:"remote_unknown"`
	// pAuthnHint: This SHOULD contain an RPC authentication level (see [MS-RPCE] section
	// 2.2.1.1.8) that indicates the minimum authentication level supported by the object
	// exporter, which MAY be ignored by the client.<52>
	AuthnHint uint32 `idl:"name:pAuthnHint;pointer:ref" json:"authn_hint"`
	// Return: The ResolveOxid return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResolveOXIDResponse) xxx_ToOp(ctx context.Context, op *xxx_ResolveOXIDOperation) *xxx_ResolveOXIDOperation {
	if op == nil {
		op = &xxx_ResolveOXIDOperation{}
	}
	if o == nil {
		return op
	}
	o.OXIDBindings = op.OXIDBindings
	o.RemoteUnknown = op.RemoteUnknown
	o.AuthnHint = op.AuthnHint
	o.Return = op.Return
	return op
}

func (o *ResolveOXIDResponse) xxx_FromOp(ctx context.Context, op *xxx_ResolveOXIDOperation) {
	if o == nil {
		return
	}
	o.OXIDBindings = op.OXIDBindings
	o.RemoteUnknown = op.RemoteUnknown
	o.AuthnHint = op.AuthnHint
	o.Return = op.Return
}
func (o *ResolveOXIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResolveOXIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveOXIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SimplePingOperation structure represents the SimplePing operation
type xxx_SimplePingOperation struct {
	SetID  uint64 `idl:"name:pSetId" json:"set_id"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SimplePingOperation) OpNum() int { return 1 }

func (o *xxx_SimplePingOperation) OpName() string { return "/IObjectExporter/v0/SimplePing" }

func (o *xxx_SimplePingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SimplePingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pSetId {in} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.WriteData(o.SetID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SimplePingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pSetId {in} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.ReadData(&o.SetID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SimplePingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SimplePingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SimplePingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SimplePingRequest structure represents the SimplePing operation request
type SimplePingRequest struct {
	// pSetId: MUST specify a SETID of the ping set to ping. This parameter MUST specify
	// a SETID previously returned from a call to IObjectExporter::ComplexPing.
	SetID uint64 `idl:"name:pSetId" json:"set_id"`
}

func (o *SimplePingRequest) xxx_ToOp(ctx context.Context, op *xxx_SimplePingOperation) *xxx_SimplePingOperation {
	if op == nil {
		op = &xxx_SimplePingOperation{}
	}
	if o == nil {
		return op
	}
	o.SetID = op.SetID
	return op
}

func (o *SimplePingRequest) xxx_FromOp(ctx context.Context, op *xxx_SimplePingOperation) {
	if o == nil {
		return
	}
	o.SetID = op.SetID
}
func (o *SimplePingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SimplePingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SimplePingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SimplePingResponse structure represents the SimplePing operation response
type SimplePingResponse struct {
	// Return: The SimplePing return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SimplePingResponse) xxx_ToOp(ctx context.Context, op *xxx_SimplePingOperation) *xxx_SimplePingOperation {
	if op == nil {
		op = &xxx_SimplePingOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SimplePingResponse) xxx_FromOp(ctx context.Context, op *xxx_SimplePingOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SimplePingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SimplePingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SimplePingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ComplexPingOperation structure represents the ComplexPing operation
type xxx_ComplexPingOperation struct {
	SetID              uint64   `idl:"name:pSetId" json:"set_id"`
	SequenceNum        uint16   `idl:"name:SequenceNum" json:"sequence_num"`
	AddToSetCount      uint16   `idl:"name:cAddToSet" json:"add_to_set_count"`
	DeleteFromSetCount uint16   `idl:"name:cDelFromSet" json:"delete_from_set_count"`
	AddToSet           []uint64 `idl:"name:AddToSet;size_is:(cAddToSet);pointer:unique" json:"add_to_set"`
	DeleteFromSet      []uint64 `idl:"name:DelFromSet;size_is:(cDelFromSet);pointer:unique" json:"delete_from_set"`
	PingBackoffFactor  uint16   `idl:"name:pPingBackoffFactor" json:"ping_backoff_factor"`
	Return             uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ComplexPingOperation) OpNum() int { return 2 }

func (o *xxx_ComplexPingOperation) OpName() string { return "/IObjectExporter/v0/ComplexPing" }

func (o *xxx_ComplexPingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.AddToSet != nil && o.AddToSetCount == 0 {
		o.AddToSetCount = uint16(len(o.AddToSet))
	}
	if o.DeleteFromSet != nil && o.DeleteFromSetCount == 0 {
		o.DeleteFromSetCount = uint16(len(o.DeleteFromSet))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexPingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pSetId {in, out} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.WriteData(o.SetID); err != nil {
			return err
		}
	}
	// SequenceNum {in} (1:(uint16))
	{
		if err := w.WriteData(o.SequenceNum); err != nil {
			return err
		}
	}
	// cAddToSet {in} (1:(uint16))
	{
		if err := w.WriteData(o.AddToSetCount); err != nil {
			return err
		}
	}
	// cDelFromSet {in} (1:(uint16))
	{
		if err := w.WriteData(o.DeleteFromSetCount); err != nil {
			return err
		}
	}
	// AddToSet {in} (1:{pointer=unique}[dim:0,size_is=cAddToSet])(2:{alias=OID}(uint64))
	{
		dimSize1 := uint64(o.AddToSetCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.AddToSet {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.AddToSet[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.AddToSet); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint64(0)); err != nil {
				return err
			}
		}
	}
	// DelFromSet {in} (1:{pointer=unique}[dim:0,size_is=cDelFromSet])(2:{alias=OID}(uint64))
	{
		dimSize1 := uint64(o.DeleteFromSetCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DeleteFromSet {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.DeleteFromSet[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.DeleteFromSet); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint64(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ComplexPingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pSetId {in, out} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.ReadData(&o.SetID); err != nil {
			return err
		}
	}
	// SequenceNum {in} (1:(uint16))
	{
		if err := w.ReadData(&o.SequenceNum); err != nil {
			return err
		}
	}
	// cAddToSet {in} (1:(uint16))
	{
		if err := w.ReadData(&o.AddToSetCount); err != nil {
			return err
		}
	}
	// cDelFromSet {in} (1:(uint16))
	{
		if err := w.ReadData(&o.DeleteFromSetCount); err != nil {
			return err
		}
	}
	// AddToSet {in} (1:{pointer=unique}[dim:0,size_is=cAddToSet])(2:{alias=OID}(uint64))
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
			return fmt.Errorf("buffer overflow for size %d of array o.AddToSet", sizeInfo[0])
		}
		o.AddToSet = make([]uint64, sizeInfo[0])
		for i1 := range o.AddToSet {
			i1 := i1
			if err := w.ReadData(&o.AddToSet[i1]); err != nil {
				return err
			}
		}
	}
	// DelFromSet {in} (1:{pointer=unique}[dim:0,size_is=cDelFromSet])(2:{alias=OID}(uint64))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DeleteFromSet", sizeInfo[0])
		}
		o.DeleteFromSet = make([]uint64, sizeInfo[0])
		for i1 := range o.DeleteFromSet {
			i1 := i1
			if err := w.ReadData(&o.DeleteFromSet[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ComplexPingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ComplexPingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pSetId {in, out} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.WriteData(o.SetID); err != nil {
			return err
		}
	}
	// pPingBackoffFactor {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.WriteData(o.PingBackoffFactor); err != nil {
			return err
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

func (o *xxx_ComplexPingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pSetId {in, out} (1:{pointer=ref}*(1))(2:{alias=SETID}(uint64))
	{
		if err := w.ReadData(&o.SetID); err != nil {
			return err
		}
	}
	// pPingBackoffFactor {out} (1:{pointer=ref}*(1)(uint16))
	{
		if err := w.ReadData(&o.PingBackoffFactor); err != nil {
			return err
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

// ComplexPingRequest structure represents the ComplexPing operation request
type ComplexPingRequest struct {
	// pSetId: This MUST specify the SETID of the ping set to ping. If the SETID specified
	// by the client is zero, the object resolver MUST return the SETID of a new ping set
	// containing the OIDs specified in AddToSet. If the SETID specified by the client is
	// not zero, the object exporter MUST not change the SETID.
	SetID uint64 `idl:"name:pSetId" json:"set_id"`
	// SequenceNum: This MUST specify a sequence number shared between the client and the
	// object resolver.
	SequenceNum uint16 `idl:"name:SequenceNum" json:"sequence_num"`
	// cAddToSet: This MUST specify the number of OIDs in the AddToSet array.
	AddToSetCount uint16 `idl:"name:cAddToSet" json:"add_to_set_count"`
	// cDelFromSet: This MUST specify the number of OIDs in the DelFromSet array.
	DeleteFromSetCount uint16 `idl:"name:cDelFromSet" json:"delete_from_set_count"`
	// AddToSet: This MUST specify an array of OIDs to add to the set.
	AddToSet []uint64 `idl:"name:AddToSet;size_is:(cAddToSet);pointer:unique" json:"add_to_set"`
	// DelFromSet: This MUST specify an array of OIDs to remove from the set.
	DeleteFromSet []uint64 `idl:"name:DelFromSet;size_is:(cDelFromSet);pointer:unique" json:"delete_from_set"`
}

func (o *ComplexPingRequest) xxx_ToOp(ctx context.Context, op *xxx_ComplexPingOperation) *xxx_ComplexPingOperation {
	if op == nil {
		op = &xxx_ComplexPingOperation{}
	}
	if o == nil {
		return op
	}
	o.SetID = op.SetID
	o.SequenceNum = op.SequenceNum
	o.AddToSetCount = op.AddToSetCount
	o.DeleteFromSetCount = op.DeleteFromSetCount
	o.AddToSet = op.AddToSet
	o.DeleteFromSet = op.DeleteFromSet
	return op
}

func (o *ComplexPingRequest) xxx_FromOp(ctx context.Context, op *xxx_ComplexPingOperation) {
	if o == nil {
		return
	}
	o.SetID = op.SetID
	o.SequenceNum = op.SequenceNum
	o.AddToSetCount = op.AddToSetCount
	o.DeleteFromSetCount = op.DeleteFromSetCount
	o.AddToSet = op.AddToSet
	o.DeleteFromSet = op.DeleteFromSet
}
func (o *ComplexPingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ComplexPingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexPingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ComplexPingResponse structure represents the ComplexPing operation response
type ComplexPingResponse struct {
	// pSetId: This MUST specify the SETID of the ping set to ping. If the SETID specified
	// by the client is zero, the object resolver MUST return the SETID of a new ping set
	// containing the OIDs specified in AddToSet. If the SETID specified by the client is
	// not zero, the object exporter MUST not change the SETID.
	SetID uint64 `idl:"name:pSetId" json:"set_id"`
	// pPingBackoffFactor: This MUST contain a hint to reduce ping load on the server. Servers
	// MAY set the hint to an implementation-specific value. Clients MAY choose to treat
	// this as zero always.<55>
	PingBackoffFactor uint16 `idl:"name:pPingBackoffFactor" json:"ping_backoff_factor"`
	// Return: The ComplexPing return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ComplexPingResponse) xxx_ToOp(ctx context.Context, op *xxx_ComplexPingOperation) *xxx_ComplexPingOperation {
	if op == nil {
		op = &xxx_ComplexPingOperation{}
	}
	if o == nil {
		return op
	}
	o.SetID = op.SetID
	o.PingBackoffFactor = op.PingBackoffFactor
	o.Return = op.Return
	return op
}

func (o *ComplexPingResponse) xxx_FromOp(ctx context.Context, op *xxx_ComplexPingOperation) {
	if o == nil {
		return
	}
	o.SetID = op.SetID
	o.PingBackoffFactor = op.PingBackoffFactor
	o.Return = op.Return
}
func (o *ComplexPingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ComplexPingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ComplexPingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerAliveOperation structure represents the ServerAlive operation
type xxx_ServerAliveOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerAliveOperation) OpNum() int { return 3 }

func (o *xxx_ServerAliveOperation) OpName() string { return "/IObjectExporter/v0/ServerAlive" }

func (o *xxx_ServerAliveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerAliveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_ServerAliveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_ServerAliveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerAliveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerAliveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerAliveRequest structure represents the ServerAlive operation request
type ServerAliveRequest struct {
}

func (o *ServerAliveRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerAliveOperation) *xxx_ServerAliveOperation {
	if op == nil {
		op = &xxx_ServerAliveOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *ServerAliveRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerAliveOperation) {
	if o == nil {
		return
	}
}
func (o *ServerAliveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerAliveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerAliveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerAliveResponse structure represents the ServerAlive operation response
type ServerAliveResponse struct {
	// Return: The ServerAlive return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerAliveResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerAliveOperation) *xxx_ServerAliveOperation {
	if op == nil {
		op = &xxx_ServerAliveOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *ServerAliveResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerAliveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerAliveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerAliveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerAliveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResolveOxid2Operation structure represents the ResolveOxid2 operation
type xxx_ResolveOxid2Operation struct {
	OXID                            uint64                `idl:"name:pOxid" json:"oxid"`
	RequestedProtocolSequencesCount uint16                `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	RequestedProtocolSequences      []uint16              `idl:"name:arRequestedProtseqs;size_is:(cRequestedProtseqs);pointer:ref" json:"requested_protocol_sequences"`
	OXIDBindings                    *dcom.DualStringArray `idl:"name:ppdsaOxidBindings;pointer:ref" json:"oxid_bindings"`
	RemoteUnknown                   *dcom.IPID            `idl:"name:pipidRemUnknown;pointer:ref" json:"remote_unknown"`
	AuthnHint                       uint32                `idl:"name:pAuthnHint;pointer:ref" json:"authn_hint"`
	COMVersion                      *dcom.COMVersion      `idl:"name:pComVersion;pointer:ref" json:"com_version"`
	Return                          uint32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ResolveOxid2Operation) OpNum() int { return 4 }

func (o *xxx_ResolveOxid2Operation) OpName() string { return "/IObjectExporter/v0/ResolveOxid2" }

func (o *xxx_ResolveOxid2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RequestedProtocolSequences != nil && o.RequestedProtocolSequencesCount == 0 {
		o.RequestedProtocolSequencesCount = uint16(len(o.RequestedProtocolSequences))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveOxid2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pOxid {in} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.WriteData(o.OXID); err != nil {
			return err
		}
	}
	// cRequestedProtseqs {in} (1:(uint16))
	{
		if err := w.WriteData(o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// arRequestedProtseqs {in} (1:{pointer=ref}[dim:0,size_is=cRequestedProtseqs](uint16))
	{
		dimSize1 := uint64(o.RequestedProtocolSequencesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.RequestedProtocolSequences {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.RequestedProtocolSequences[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.RequestedProtocolSequences); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ResolveOxid2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pOxid {in} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.ReadData(&o.OXID); err != nil {
			return err
		}
	}
	// cRequestedProtseqs {in} (1:(uint16))
	{
		if err := w.ReadData(&o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// arRequestedProtseqs {in} (1:{pointer=ref}[dim:0,size_is=cRequestedProtseqs](uint16))
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
			return fmt.Errorf("buffer overflow for size %d of array o.RequestedProtocolSequences", sizeInfo[0])
		}
		o.RequestedProtocolSequences = make([]uint16, sizeInfo[0])
		for i1 := range o.RequestedProtocolSequences {
			i1 := i1
			if err := w.ReadData(&o.RequestedProtocolSequences[i1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ResolveOxid2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResolveOxid2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppdsaOxidBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		if o.OXIDBindings != nil {
			_ptr_ppdsaOxidBindings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OXIDBindings != nil {
					if err := o.OXIDBindings.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.DualStringArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OXIDBindings, _ptr_ppdsaOxidBindings); err != nil {
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
	// pipidRemUnknown {out} (1:{pointer=ref}*(1))(2:{alias=IPID, names=GUID}(struct))
	{
		if o.RemoteUnknown != nil {
			if err := o.RemoteUnknown.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IPID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pAuthnHint {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AuthnHint); err != nil {
			return err
		}
	}
	// pComVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.COMVersion != nil {
			if err := o.COMVersion.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.COMVersion{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ResolveOxid2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppdsaOxidBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		_ptr_ppdsaOxidBindings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OXIDBindings == nil {
				o.OXIDBindings = &dcom.DualStringArray{}
			}
			if err := o.OXIDBindings.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdsaOxidBindings := func(ptr interface{}) { o.OXIDBindings = *ptr.(**dcom.DualStringArray) }
		if err := w.ReadPointer(&o.OXIDBindings, _s_ppdsaOxidBindings, _ptr_ppdsaOxidBindings); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pipidRemUnknown {out} (1:{pointer=ref}*(1))(2:{alias=IPID, names=GUID}(struct))
	{
		if o.RemoteUnknown == nil {
			o.RemoteUnknown = &dcom.IPID{}
		}
		if err := o.RemoteUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pAuthnHint {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AuthnHint); err != nil {
			return err
		}
	}
	// pComVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.COMVersion == nil {
			o.COMVersion = &dcom.COMVersion{}
		}
		if err := o.COMVersion.UnmarshalNDR(ctx, w); err != nil {
			return err
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

// ResolveOxid2Request structure represents the ResolveOxid2 operation request
type ResolveOxid2Request struct {
	// pOxid:  This MUST specify an OXID identifying an object exporter.
	OXID uint64 `idl:"name:pOxid" json:"oxid"`
	// cRequestedProtseqs: This MUST contain the number of elements in the arRequestedProtseqs
	// array.
	RequestedProtocolSequencesCount uint16 `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	// arRequestedProtseqs: This MUST contain an array of RPC protocol sequence identifiers
	// supported by the client.
	RequestedProtocolSequences []uint16 `idl:"name:arRequestedProtseqs;size_is:(cRequestedProtseqs);pointer:ref" json:"requested_protocol_sequences"`
}

func (o *ResolveOxid2Request) xxx_ToOp(ctx context.Context, op *xxx_ResolveOxid2Operation) *xxx_ResolveOxid2Operation {
	if op == nil {
		op = &xxx_ResolveOxid2Operation{}
	}
	if o == nil {
		return op
	}
	o.OXID = op.OXID
	o.RequestedProtocolSequencesCount = op.RequestedProtocolSequencesCount
	o.RequestedProtocolSequences = op.RequestedProtocolSequences
	return op
}

func (o *ResolveOxid2Request) xxx_FromOp(ctx context.Context, op *xxx_ResolveOxid2Operation) {
	if o == nil {
		return
	}
	o.OXID = op.OXID
	o.RequestedProtocolSequencesCount = op.RequestedProtocolSequencesCount
	o.RequestedProtocolSequences = op.RequestedProtocolSequences
}
func (o *ResolveOxid2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResolveOxid2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveOxid2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResolveOxid2Response structure represents the ResolveOxid2 operation response
type ResolveOxid2Response struct {
	// ppdsaOxidBindings:  This MUST contain the string and security bindings supported
	// by the object exporter and MUST NOT be NULL. The returned string bindings SHOULD
	// contain endpoints.
	OXIDBindings *dcom.DualStringArray `idl:"name:ppdsaOxidBindings;pointer:ref" json:"oxid_bindings"`
	// pipidRemUnknown:  This MUST contain the IPID of the object exporter Remote Unknown
	// object.
	RemoteUnknown *dcom.IPID `idl:"name:pipidRemUnknown;pointer:ref" json:"remote_unknown"`
	// pAuthnHint:  This SHOULD contain an RPC authentication level (see [MS-RPCE] section
	// 2.2.1.1.8) that denotes the minimum authentication level supported by the object
	// exporter.<57>
	AuthnHint uint32 `idl:"name:pAuthnHint;pointer:ref" json:"authn_hint"`
	// pComVersion:  This MUST contain the COMVERSION of the object exporter. For details,
	// see section 2.2.11.
	COMVersion *dcom.COMVersion `idl:"name:pComVersion;pointer:ref" json:"com_version"`
	// Return: The ResolveOxid2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResolveOxid2Response) xxx_ToOp(ctx context.Context, op *xxx_ResolveOxid2Operation) *xxx_ResolveOxid2Operation {
	if op == nil {
		op = &xxx_ResolveOxid2Operation{}
	}
	if o == nil {
		return op
	}
	o.OXIDBindings = op.OXIDBindings
	o.RemoteUnknown = op.RemoteUnknown
	o.AuthnHint = op.AuthnHint
	o.COMVersion = op.COMVersion
	o.Return = op.Return
	return op
}

func (o *ResolveOxid2Response) xxx_FromOp(ctx context.Context, op *xxx_ResolveOxid2Operation) {
	if o == nil {
		return
	}
	o.OXIDBindings = op.OXIDBindings
	o.RemoteUnknown = op.RemoteUnknown
	o.AuthnHint = op.AuthnHint
	o.COMVersion = op.COMVersion
	o.Return = op.Return
}
func (o *ResolveOxid2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResolveOxid2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResolveOxid2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerAlive2Operation structure represents the ServerAlive2 operation
type xxx_ServerAlive2Operation struct {
	COMVersion             *dcom.COMVersion      `idl:"name:pComVersion;pointer:ref" json:"com_version"`
	ObjectResolverBindings *dcom.DualStringArray `idl:"name:ppdsaOrBindings;pointer:ref" json:"object_resolver_bindings"`
	_                      uint32                `idl:"name:pReserved;pointer:ref"`
	Return                 uint32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerAlive2Operation) OpNum() int { return 5 }

func (o *xxx_ServerAlive2Operation) OpName() string { return "/IObjectExporter/v0/ServerAlive2" }

func (o *xxx_ServerAlive2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerAlive2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_ServerAlive2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_ServerAlive2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerAlive2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pComVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.COMVersion != nil {
			if err := o.COMVersion.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.COMVersion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ppdsaOrBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		if o.ObjectResolverBindings != nil {
			_ptr_ppdsaOrBindings := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectResolverBindings != nil {
					if err := o.ObjectResolverBindings.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.DualStringArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectResolverBindings, _ptr_ppdsaOrBindings); err != nil {
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
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
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

func (o *xxx_ServerAlive2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pComVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.COMVersion == nil {
			o.COMVersion = &dcom.COMVersion{}
		}
		if err := o.COMVersion.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ppdsaOrBindings {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DUALSTRINGARRAY}(struct))
	{
		_ptr_ppdsaOrBindings := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectResolverBindings == nil {
				o.ObjectResolverBindings = &dcom.DualStringArray{}
			}
			if err := o.ObjectResolverBindings.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdsaOrBindings := func(ptr interface{}) { o.ObjectResolverBindings = *ptr.(**dcom.DualStringArray) }
		if err := w.ReadPointer(&o.ObjectResolverBindings, _s_ppdsaOrBindings, _ptr_ppdsaOrBindings); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		var _pReserved uint32
		if err := w.ReadData(&_pReserved); err != nil {
			return err
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

// ServerAlive2Request structure represents the ServerAlive2 operation request
type ServerAlive2Request struct {
}

func (o *ServerAlive2Request) xxx_ToOp(ctx context.Context, op *xxx_ServerAlive2Operation) *xxx_ServerAlive2Operation {
	if op == nil {
		op = &xxx_ServerAlive2Operation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *ServerAlive2Request) xxx_FromOp(ctx context.Context, op *xxx_ServerAlive2Operation) {
	if o == nil {
		return
	}
}
func (o *ServerAlive2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerAlive2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerAlive2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerAlive2Response structure represents the ServerAlive2 operation response
type ServerAlive2Response struct {
	// pComVersion: This MUST contain the COMVERSION of the object resolver; see section
	// 2.2.11.
	COMVersion *dcom.COMVersion `idl:"name:pComVersion;pointer:ref" json:"com_version"`
	// ppdsaOrBindings:  MUST contain the string and security bindings of the object resolver.
	// The returned string bindings MUST NOT contain endpoints.
	ObjectResolverBindings *dcom.DualStringArray `idl:"name:ppdsaOrBindings;pointer:ref" json:"object_resolver_bindings"`
	// Return: The ServerAlive2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerAlive2Response) xxx_ToOp(ctx context.Context, op *xxx_ServerAlive2Operation) *xxx_ServerAlive2Operation {
	if op == nil {
		op = &xxx_ServerAlive2Operation{}
	}
	if o == nil {
		return op
	}
	o.COMVersion = op.COMVersion
	o.ObjectResolverBindings = op.ObjectResolverBindings
	o.Return = op.Return
	return op
}

func (o *ServerAlive2Response) xxx_FromOp(ctx context.Context, op *xxx_ServerAlive2Operation) {
	if o == nil {
		return
	}
	o.COMVersion = op.COMVersion
	o.ObjectResolverBindings = op.ObjectResolverBindings
	o.Return = op.Return
}
func (o *ServerAlive2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerAlive2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerAlive2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
