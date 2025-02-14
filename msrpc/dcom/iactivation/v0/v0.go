package iactivation

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// Syntax UUID
	ActivationSyntaxUUID = &uuid.UUID{TimeLow: 0x4d9f4ab8, TimeMid: 0x7d1c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x1e, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6e, 0x7c, 0x57}}
	// Syntax ID
	ActivationSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActivationSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IActivation interface.
type ActivationClient interface {

	// The RemoteActivation (Opnum 0) method is used by clients to request the activation
	// of an object. It returns the bindings, the IPID for the Remote Unknown, and the COMVERSION
	// of the object exporter that hosts the object.
	RemoteActivation(context.Context, *RemoteActivationRequest, ...dcerpc.CallOption) (*RemoteActivationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// MaxRequestedInterfaces represents the MAX_REQUESTED_INTERFACES RPC constant
const MaxRequestedInterfaces = 0x00008000

// MaxRequestedProtocolSequences represents the MAX_REQUESTED_PROTSEQS RPC constant
const MaxRequestedProtocolSequences = 0x00008000

type xxx_DefaultActivationClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultActivationClient) RemoteActivation(ctx context.Context, in *RemoteActivationRequest, opts ...dcerpc.CallOption) (*RemoteActivationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteActivationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActivationClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActivationClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewActivationClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActivationClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActivationSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultActivationClient{cc: cc}, nil
}

// xxx_RemoteActivationOperation structure represents the RemoteActivation operation
type xxx_RemoteActivationOperation struct {
	ORPCThis                        *dcom.ORPCThis           `idl:"name:ORPCthis" json:"orpc_this"`
	ORPCThat                        *dcom.ORPCThat           `idl:"name:ORPCthat" json:"orpc_that"`
	ClassID                         *dtyp.GUID               `idl:"name:Clsid" json:"class_id"`
	ObjectName                      string                   `idl:"name:pwszObjectName;string;pointer:unique" json:"object_name"`
	ObjectStorage                   *dcom.InterfacePointer   `idl:"name:pObjectStorage;pointer:unique" json:"object_storage"`
	ClientImpLevel                  uint32                   `idl:"name:ClientImpLevel" json:"client_imp_level"`
	Mode                            uint32                   `idl:"name:Mode" json:"mode"`
	Interfaces                      uint32                   `idl:"name:Interfaces" json:"interfaces"`
	IIDs                            []*dcom.IID              `idl:"name:pIIDs;size_is:(Interfaces);pointer:unique" json:"iids"`
	RequestedProtocolSequencesCount uint16                   `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	RequestedProtocolSequences      []uint16                 `idl:"name:aRequestedProtseqs;size_is:(cRequestedProtseqs)" json:"requested_protocol_sequences"`
	OXID                            uint64                   `idl:"name:pOxid" json:"oxid"`
	OXIDBindings                    *dcom.DualStringArray    `idl:"name:ppdsaOxidBindings" json:"oxid_bindings"`
	RemoteUnknown                   *dcom.IPID               `idl:"name:pipidRemUnknown" json:"remote_unknown"`
	AuthnHint                       uint32                   `idl:"name:pAuthnHint" json:"authn_hint"`
	ServerVersion                   *dcom.COMVersion         `idl:"name:pServerVersion" json:"server_version"`
	HResult                         int32                    `idl:"name:phr" json:"hresult"`
	InterfaceData                   []*dcom.InterfacePointer `idl:"name:ppInterfaceData;size_is:(Interfaces)" json:"interface_data"`
	Results                         []int32                  `idl:"name:pResults;size_is:(Interfaces)" json:"results"`
	Return                          uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteActivationOperation) OpNum() int { return 0 }

func (o *xxx_RemoteActivationOperation) OpName() string { return "/IActivation/v0/RemoteActivation" }

func (o *xxx_RemoteActivationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.IIDs != nil && o.Interfaces == 0 {
		o.Interfaces = uint32(len(o.IIDs))
	}
	if o.RequestedProtocolSequences != nil && o.RequestedProtocolSequencesCount == 0 {
		o.RequestedProtocolSequencesCount = uint16(len(o.RequestedProtocolSequences))
	}
	if o.Interfaces < uint32(1) || o.Interfaces > uint32(32768) {
		return fmt.Errorf("Interfaces is out of range")
	}
	if o.RequestedProtocolSequencesCount > uint16(32768) {
		return fmt.Errorf("RequestedProtocolSequencesCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteActivationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ORPCthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis != nil {
			if err := o.ORPCThis.MarshalNDR(ctx, w); err != nil {
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
	// Clsid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID != nil {
			if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pwszObjectName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ObjectName != "" {
			_ptr_pwszObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ObjectName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectName, _ptr_pwszObjectName); err != nil {
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
	// pObjectStorage {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		if o.ObjectStorage != nil {
			_ptr_pObjectStorage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ObjectStorage != nil {
					if err := o.ObjectStorage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ObjectStorage, _ptr_pObjectStorage); err != nil {
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
	// ClientImpLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientImpLevel); err != nil {
			return err
		}
	}
	// Mode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Mode); err != nil {
			return err
		}
	}
	// Interfaces {in} (1:{range=(1,32768), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interfaces); err != nil {
			return err
		}
	}
	// pIIDs {in} (1:{pointer=unique}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=Interfaces](struct))
	{
		if o.IIDs != nil || o.Interfaces > 0 {
			_ptr_pIIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Interfaces)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.IIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.IIDs[i1] != nil {
						if err := o.IIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.IIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IIDs, _ptr_pIIDs); err != nil {
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
	// cRequestedProtseqs {in} (1:{range=(0,32768)}(uint16))
	{
		if err := w.WriteData(o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// aRequestedProtseqs {in} (1:[dim:0,size_is=cRequestedProtseqs](uint16))
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

func (o *xxx_RemoteActivationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ORPCthis {in} (1:{pointer=ref}*(1))(2:{alias=ORPCTHIS}(struct))
	{
		if o.ORPCThis == nil {
			o.ORPCThis = &dcom.ORPCThis{}
		}
		if err := o.ORPCThis.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Clsid {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ClassID == nil {
			o.ClassID = &dtyp.GUID{}
		}
		if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pwszObjectName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ObjectName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ObjectName, _s_pwszObjectName, _ptr_pwszObjectName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pObjectStorage {in} (1:{pointer=unique}*(1))(2:{alias=MInterfacePointer}(struct))
	{
		_ptr_pObjectStorage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ObjectStorage == nil {
				o.ObjectStorage = &dcom.InterfacePointer{}
			}
			if err := o.ObjectStorage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pObjectStorage := func(ptr interface{}) { o.ObjectStorage = *ptr.(**dcom.InterfacePointer) }
		if err := w.ReadPointer(&o.ObjectStorage, _s_pObjectStorage, _ptr_pObjectStorage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientImpLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientImpLevel); err != nil {
			return err
		}
	}
	// Mode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Mode); err != nil {
			return err
		}
	}
	// Interfaces {in} (1:{range=(1,32768), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interfaces); err != nil {
			return err
		}
	}
	// pIIDs {in} (1:{pointer=unique}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=Interfaces](struct))
	{
		_ptr_pIIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.IIDs", sizeInfo[0])
			}
			o.IIDs = make([]*dcom.IID, sizeInfo[0])
			for i1 := range o.IIDs {
				i1 := i1
				if o.IIDs[i1] == nil {
					o.IIDs[i1] = &dcom.IID{}
				}
				if err := o.IIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pIIDs := func(ptr interface{}) { o.IIDs = *ptr.(*[]*dcom.IID) }
		if err := w.ReadPointer(&o.IIDs, _s_pIIDs, _ptr_pIIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cRequestedProtseqs {in} (1:{range=(0,32768)}(uint16))
	{
		if err := w.ReadData(&o.RequestedProtocolSequencesCount); err != nil {
			return err
		}
	}
	// aRequestedProtseqs {in} (1:[dim:0,size_is=cRequestedProtseqs](uint16))
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

func (o *xxx_RemoteActivationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteActivationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ORPCthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat != nil {
			if err := o.ORPCThat.MarshalNDR(ctx, w); err != nil {
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
	// pOxid {out} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.WriteData(o.OXID); err != nil {
			return err
		}
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
	// pServerVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.ServerVersion != nil {
			if err := o.ServerVersion.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.COMVersion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// phr {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.HResult); err != nil {
			return err
		}
	}
	// ppInterfaceData {out, disable_consistency_check} (1:{pointer=ref}*(1)[dim:0,size_is=Interfaces]*(1))(2:{alias=MInterfacePointer}(struct))
	{
		dimSize1 := uint64(o.Interfaces)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.InterfaceData {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.InterfaceData[i1] != nil {
				_ptr_ppInterfaceData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.InterfaceData[i1] != nil {
						if err := o.InterfaceData[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.InterfaceData[i1], _ptr_ppInterfaceData); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.InterfaceData); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pResults {out, disable_consistency_check} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=Interfaces](int32))
	{
		dimSize1 := uint64(o.Interfaces)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Results {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Results[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Results); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
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

func (o *xxx_RemoteActivationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ORPCthat {out} (1:{pointer=ref}*(1))(2:{alias=ORPCTHAT}(struct))
	{
		if o.ORPCThat == nil {
			o.ORPCThat = &dcom.ORPCThat{}
		}
		if err := o.ORPCThat.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pOxid {out} (1:{pointer=ref}*(1))(2:{alias=OXID}(uint64))
	{
		if err := w.ReadData(&o.OXID); err != nil {
			return err
		}
	}
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
	// pServerVersion {out} (1:{pointer=ref}*(1))(2:{alias=COMVERSION}(struct))
	{
		if o.ServerVersion == nil {
			o.ServerVersion = &dcom.COMVersion{}
		}
		if err := o.ServerVersion.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// phr {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.HResult); err != nil {
			return err
		}
	}
	// ppInterfaceData {out, disable_consistency_check} (1:{pointer=ref}*(1)[dim:0,size_is=Interfaces]*(1))(2:{alias=MInterfacePointer}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceData", sizeInfo[0])
		}
		o.InterfaceData = make([]*dcom.InterfacePointer, sizeInfo[0])
		for i1 := range o.InterfaceData {
			i1 := i1
			_ptr_ppInterfaceData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.InterfaceData[i1] == nil {
					o.InterfaceData[i1] = &dcom.InterfacePointer{}
				}
				if err := o.InterfaceData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppInterfaceData := func(ptr interface{}) { o.InterfaceData[i1] = *ptr.(**dcom.InterfacePointer) }
			if err := w.ReadPointer(&o.InterfaceData[i1], _s_ppInterfaceData, _ptr_ppInterfaceData); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResults {out, disable_consistency_check} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=Interfaces](int32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Results", sizeInfo[0])
		}
		o.Results = make([]int32, sizeInfo[0])
		for i1 := range o.Results {
			i1 := i1
			if err := w.ReadData(&o.Results[i1]); err != nil {
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

// RemoteActivationRequest structure represents the RemoteActivation operation request
type RemoteActivationRequest struct {
	// ORPCthis:  This MUST specify an ORPCTHIS. The COMVERSION field SHOULD contain the
	// negotiated version, as described in section 1.7. The extensions field MUST be set
	// to NULL.
	ORPCThis *dcom.ORPCThis `idl:"name:ORPCthis" json:"orpc_this"`
	// Clsid:  This MUST specify the CLSID of the object to be created.
	ClassID *dtyp.GUID `idl:"name:Clsid" json:"class_id"`
	// pwszObjectName: This MAY contain a string to be used to initialize the object.<68>
	ObjectName string `idl:"name:pwszObjectName;string;pointer:unique" json:"object_name"`
	// pObjectStorage:  This MAY contain a marshaled OBJREF to be used to initialize the
	// object.<69>
	ObjectStorage *dcom.InterfacePointer `idl:"name:pObjectStorage;pointer:unique" json:"object_storage"`
	// ClientImpLevel:  This MUST contain an implementation-specific value that MUST be
	// ignored on receipt.<70>
	ClientImpLevel uint32 `idl:"name:ClientImpLevel" json:"client_imp_level"`
	// Mode: If the activation is for a class factory reference, this parameter MUST be
	// 0xFFFFFFFF. Otherwise, it MUST be 0, except when the client specifies an initialization
	// string in pwszObjectName. If it does, this field MAY contain an implementation-specific
	// value.<71>
	Mode uint32 `idl:"name:Mode" json:"mode"`
	// Interfaces: This MUST contain the number of elements in pIIDs. This value MUST be
	// between 1 and MAX_REQUESTED_INTERFACES; see section 2.2.28.1.
	Interfaces uint32 `idl:"name:Interfaces" json:"interfaces"`
	// pIIDs:  This MUST be an array of requested IIDs on the object to be created.
	IIDs []*dcom.IID `idl:"name:pIIDs;size_is:(Interfaces);pointer:unique" json:"iids"`
	// cRequestedProtseqs: This MUST contain the number of elements in aRequestedProtseqs.
	// This value MUST be between 1 and MAX_REQUESTED_PROTSEQS (see section 2.2.28.1).
	RequestedProtocolSequencesCount uint16 `idl:"name:cRequestedProtseqs" json:"requested_protocol_sequences_count"`
	// aRequestedProtseqs: This MUST specify an array of RPC protocol sequence identifiers
	// that the client supports.
	RequestedProtocolSequences []uint16 `idl:"name:aRequestedProtseqs;size_is:(cRequestedProtseqs)" json:"requested_protocol_sequences"`
}

func (o *RemoteActivationRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteActivationOperation) *xxx_RemoteActivationOperation {
	if op == nil {
		op = &xxx_RemoteActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.ORPCThis = o.ORPCThis
	op.ClassID = o.ClassID
	op.ObjectName = o.ObjectName
	op.ObjectStorage = o.ObjectStorage
	op.ClientImpLevel = o.ClientImpLevel
	op.Mode = o.Mode
	op.Interfaces = o.Interfaces
	op.IIDs = o.IIDs
	op.RequestedProtocolSequencesCount = o.RequestedProtocolSequencesCount
	op.RequestedProtocolSequences = o.RequestedProtocolSequences
	return op
}

func (o *RemoteActivationRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteActivationOperation) {
	if o == nil {
		return
	}
	o.ORPCThis = op.ORPCThis
	o.ClassID = op.ClassID
	o.ObjectName = op.ObjectName
	o.ObjectStorage = op.ObjectStorage
	o.ClientImpLevel = op.ClientImpLevel
	o.Mode = op.Mode
	o.Interfaces = op.Interfaces
	o.IIDs = op.IIDs
	o.RequestedProtocolSequencesCount = op.RequestedProtocolSequencesCount
	o.RequestedProtocolSequences = op.RequestedProtocolSequences
}
func (o *RemoteActivationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteActivationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteActivationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteActivationResponse structure represents the RemoteActivation operation response
type RemoteActivationResponse struct {
	// ORPCthat:  This MUST contain an ORPCTHAT. The extensions field MUST be set to NULL.
	ORPCThat *dcom.ORPCThat `idl:"name:ORPCthat" json:"orpc_that"`
	// pOxid: This MUST contain an OXID value identifying the object exporter containing
	// this object.
	OXID uint64 `idl:"name:pOxid" json:"oxid"`
	// ppdsaOxidBindings:  This MUST contain the string and security bindings supported
	// by the object exporter and MUST NOT be NULL. The returned string bindings SHOULD
	// contain endpoints.
	OXIDBindings *dcom.DualStringArray `idl:"name:ppdsaOxidBindings" json:"oxid_bindings"`
	// pipidRemUnknown:  This MUST contain the IPID of the object exporter Remote Unknown
	// object.
	RemoteUnknown *dcom.IPID `idl:"name:pipidRemUnknown" json:"remote_unknown"`
	// pAuthnHint:  This SHOULD contain an RPC authentication level (see [MS-RPCE] section
	// 2.2.1.1.8) that denotes the minimum authentication level supported by the server.
	// This MAY be ignored by the client.<72>
	AuthnHint uint32 `idl:"name:pAuthnHint" json:"authn_hint"`
	// pServerVersion:  This MUST contain the COMVERSION of the object exporter. For details,
	// see section 2.2.11.
	ServerVersion *dcom.COMVersion `idl:"name:pServerVersion" json:"server_version"`
	// phr: An HRESULT that indicates the result of the activation. Success codes other
	// than 0x00000000 MUST NOT be used.
	HResult int32 `idl:"name:phr" json:"hresult"`
	// ppInterfaceData:  This MUST contain an array of MInterfacePointer structures containing
	// the results for each requested interface.
	InterfaceData []*dcom.InterfacePointer `idl:"name:ppInterfaceData;size_is:(Interfaces)" json:"interface_data"`
	// pResults: If the phr parameter contains 0x00000000, this MUST contain an array of
	// HRESULTs containing the results of querying the object for each interface in pIIDs.
	// Success codes other than 0x00000000 MUST NOT be used in the results array. If the
	// phr parameter contains a failure value, this MUST contain an array of 0x00000000
	// values, one for each interface in pIIDs.
	Results []int32 `idl:"name:pResults;size_is:(Interfaces)" json:"results"`
	// Return: The RemoteActivation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoteActivationResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteActivationOperation) *xxx_RemoteActivationOperation {
	if op == nil {
		op = &xxx_RemoteActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.ORPCThat = o.ORPCThat
	op.OXID = o.OXID
	op.OXIDBindings = o.OXIDBindings
	op.RemoteUnknown = o.RemoteUnknown
	op.AuthnHint = o.AuthnHint
	op.ServerVersion = o.ServerVersion
	op.HResult = o.HResult
	op.InterfaceData = o.InterfaceData
	op.Results = o.Results
	op.Return = o.Return
	return op
}

func (o *RemoteActivationResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteActivationOperation) {
	if o == nil {
		return
	}
	o.ORPCThat = op.ORPCThat
	o.OXID = op.OXID
	o.OXIDBindings = op.OXIDBindings
	o.RemoteUnknown = op.RemoteUnknown
	o.AuthnHint = op.AuthnHint
	o.ServerVersion = op.ServerVersion
	o.HResult = op.HResult
	o.InterfaceData = op.InterfaceData
	o.Results = op.Results
	o.Return = op.Return
}
func (o *RemoteActivationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteActivationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteActivationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
