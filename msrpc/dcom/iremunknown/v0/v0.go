package iremunknown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom"
)

var (
	// IRemUnknown interface identifier 00000131-0000-0000-c000-000000000046
	RemoteUnknownIID = &dcom.IID{Data1: 0x00000131, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	RemoteUnknownSyntaxUUID = &uuid.UUID{TimeLow: 0x131, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	RemoteUnknownSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteUnknownSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IRemUnknown interface.
type RemoteUnknownClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This RemQueryInterface (Opnum 3) method acquires standard object references (see
	// section 2.2.18.1) to additional interfaces on the object.
	RemoteQueryInterface(context.Context, *RemoteQueryInterfaceRequest, ...dcerpc.CallOption) (*RemoteQueryInterfaceResponse, error)

	// The RemAddRef (Opnum 4) method requests that a specified number of reference counts
	// be incremented on a specified number of interfaces on the object.
	RemoteAddReference(context.Context, *RemoteAddReferenceRequest, ...dcerpc.CallOption) (*RemoteAddReferenceResponse, error)

	// The RemRelease (Opnum 5) method requests that a specified number of reference counts
	// be decremented on a specified number of interfaces on an object.
	RemoteRelease(context.Context, *RemoteReleaseRequest, ...dcerpc.CallOption) (*RemoteReleaseResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteUnknownClient
}

type xxx_DefaultRemoteUnknownClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteUnknownClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteUnknownClient) RemoteQueryInterface(ctx context.Context, in *RemoteQueryInterfaceRequest, opts ...dcerpc.CallOption) (*RemoteQueryInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteQueryInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteUnknownClient) RemoteAddReference(ctx context.Context, in *RemoteAddReferenceRequest, opts ...dcerpc.CallOption) (*RemoteAddReferenceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteAddReferenceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteUnknownClient) RemoteRelease(ctx context.Context, in *RemoteReleaseRequest, opts ...dcerpc.CallOption) (*RemoteReleaseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoteReleaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteUnknownClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteUnknownClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteUnknownClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteUnknownClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewRemoteUnknownClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteUnknownClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteUnknownSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultRemoteUnknownClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RemoteQueryInterfaceOperation structure represents the RemQueryInterface operation
type xxx_RemoteQueryInterfaceOperation struct {
	This                  *dcom.ORPCThis                     `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat                     `idl:"name:That" json:"that"`
	IPID                  *dtyp.GUID                         `idl:"name:ripid" json:"ipid"`
	ReferencesCount       uint32                             `idl:"name:cRefs" json:"references_count"`
	IIDsCount             uint16                             `idl:"name:cIids" json:"iids_count"`
	IIDs                  []*dcom.IID                        `idl:"name:iids;size_is:(cIids)" json:"iids"`
	QueryInterfaceResults []*dcom.RemoteQueryInterfaceResult `idl:"name:ppQIResults;size_is:(, cIids)" json:"query_interface_results"`
	Return                int32                              `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteQueryInterfaceOperation) OpNum() int { return 3 }

func (o *xxx_RemoteQueryInterfaceOperation) OpName() string {
	return "/IRemUnknown/v0/RemQueryInterface"
}

func (o *xxx_RemoteQueryInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.IIDs != nil && o.IIDsCount == 0 {
		o.IIDsCount = uint16(len(o.IIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
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
	// ripid {in} (1:{alias=REFIPID, names=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.IPID != nil {
			if err := o.IPID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// cRefs {in} (1:(uint32))
	{
		if err := w.WriteData(o.ReferencesCount); err != nil {
			return err
		}
	}
	// cIids {in} (1:(uint16))
	{
		if err := w.WriteData(o.IIDsCount); err != nil {
			return err
		}
	}
	// iids {in} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=cIids](struct))
	{
		dimSize1 := uint64(o.IIDsCount)
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
	}
	return nil
}

func (o *xxx_RemoteQueryInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ripid {in} (1:{alias=REFIPID, names=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.IPID == nil {
			o.IPID = &dtyp.GUID{}
		}
		if err := o.IPID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// cRefs {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ReferencesCount); err != nil {
			return err
		}
	}
	// cIids {in} (1:(uint16))
	{
		if err := w.ReadData(&o.IIDsCount); err != nil {
			return err
		}
	}
	// iids {in} (1:{pointer=ref}*(1))(2:{alias=IID, names=GUID}[dim:0,size_is=cIids](struct))
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
	}
	return nil
}

func (o *xxx_RemoteQueryInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteQueryInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
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
	// ppQIResults {out} (1:{pointer=ref}*(2))(2:{disable_consistency_check, alias=PREMQIRESULT}*(1))(3:{alias=REMQIRESULT}[dim:0,size_is=cIids](struct))
	{
		if o.QueryInterfaceResults != nil || o.IIDsCount > 0 {
			_ptr_ppQIResults := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.IIDsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryInterfaceResults {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.QueryInterfaceResults[i1] != nil {
						if err := o.QueryInterfaceResults[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcom.RemoteQueryInterfaceResult{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.QueryInterfaceResults); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dcom.RemoteQueryInterfaceResult{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryInterfaceResults, _ptr_ppQIResults); err != nil {
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

func (o *xxx_RemoteQueryInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppQIResults {out} (1:{pointer=ref}*(2))(2:{disable_consistency_check, alias=PREMQIRESULT,pointer=ref}*(1))(3:{alias=REMQIRESULT}[dim:0,size_is=cIids](struct))
	{
		_ptr_ppQIResults := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryInterfaceResults", sizeInfo[0])
			}
			o.QueryInterfaceResults = make([]*dcom.RemoteQueryInterfaceResult, sizeInfo[0])
			for i1 := range o.QueryInterfaceResults {
				i1 := i1
				if o.QueryInterfaceResults[i1] == nil {
					o.QueryInterfaceResults[i1] = &dcom.RemoteQueryInterfaceResult{}
				}
				if err := o.QueryInterfaceResults[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppQIResults := func(ptr interface{}) { o.QueryInterfaceResults = *ptr.(*[]*dcom.RemoteQueryInterfaceResult) }
		if err := w.ReadPointer(&o.QueryInterfaceResults, _s_ppQIResults, _ptr_ppQIResults); err != nil {
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

// RemoteQueryInterfaceRequest structure represents the RemQueryInterface operation request
type RemoteQueryInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ripid: This MUST specify an IPID that identifies the interface on the object to be
	// queried for more interfaces.
	IPID *dtyp.GUID `idl:"name:ripid" json:"ipid"`
	// cRefs: This MUST specify the number of public reference counts requested on the new
	// interfaces.
	ReferencesCount uint32 `idl:"name:cRefs" json:"references_count"`
	// cIids:  This MUST specify the number of IIDs supplied in the iids parameter and
	// returned in the ppQIResults parameter.
	IIDsCount uint16 `idl:"name:cIids" json:"iids_count"`
	// iids:  This MUST specify an array of IIDs for which the client requests object references.
	IIDs []*dcom.IID `idl:"name:iids;size_is:(cIids)" json:"iids"`
}

func (o *RemoteQueryInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_RemoteQueryInterfaceOperation {
	if o == nil {
		return &xxx_RemoteQueryInterfaceOperation{}
	}
	return &xxx_RemoteQueryInterfaceOperation{
		This:            o.This,
		IPID:            o.IPID,
		ReferencesCount: o.ReferencesCount,
		IIDsCount:       o.IIDsCount,
		IIDs:            o.IIDs,
	}
}

func (o *RemoteQueryInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteQueryInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IPID = op.IPID
	o.ReferencesCount = op.ReferencesCount
	o.IIDsCount = op.IIDsCount
	o.IIDs = op.IIDs
}
func (o *RemoteQueryInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteQueryInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteQueryInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteQueryInterfaceResponse structure represents the RemQueryInterface operation response
type RemoteQueryInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppQIResults: This MUST contain an array of REMQIRESULT structures containing the
	// results of the QueryInterface on the identified object.
	QueryInterfaceResults []*dcom.RemoteQueryInterfaceResult `idl:"name:ppQIResults;size_is:(, cIids)" json:"query_interface_results"`
	// Return: The RemQueryInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteQueryInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_RemoteQueryInterfaceOperation {
	if o == nil {
		return &xxx_RemoteQueryInterfaceOperation{}
	}
	return &xxx_RemoteQueryInterfaceOperation{
		That:                  o.That,
		QueryInterfaceResults: o.QueryInterfaceResults,
		Return:                o.Return,
	}
}

func (o *RemoteQueryInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteQueryInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QueryInterfaceResults = op.QueryInterfaceResults
	o.Return = op.Return
}
func (o *RemoteQueryInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteQueryInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteQueryInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteAddReferenceOperation structure represents the RemAddRef operation
type xxx_RemoteAddReferenceOperation struct {
	This                     *dcom.ORPCThis                   `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat                   `idl:"name:That" json:"that"`
	InterfaceReferencesCount uint16                           `idl:"name:cInterfaceRefs" json:"interface_references_count"`
	InterfaceReferences      []*dcom.RemoteInterfaceReference `idl:"name:InterfaceRefs;size_is:(cInterfaceRefs)" json:"interface_references"`
	Results                  []int32                          `idl:"name:pResults;size_is:(cInterfaceRefs)" json:"results"`
	Return                   int32                            `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteAddReferenceOperation) OpNum() int { return 4 }

func (o *xxx_RemoteAddReferenceOperation) OpName() string { return "/IRemUnknown/v0/RemAddRef" }

func (o *xxx_RemoteAddReferenceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InterfaceReferences != nil && o.InterfaceReferencesCount == 0 {
		o.InterfaceReferencesCount = uint16(len(o.InterfaceReferences))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteAddReferenceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
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
	// cInterfaceRefs {in} (1:(uint16))
	{
		if err := w.WriteData(o.InterfaceReferencesCount); err != nil {
			return err
		}
	}
	// InterfaceRefs {in} (1:[dim:0,size_is=cInterfaceRefs])(2:{alias=REMINTERFACEREF}(struct))
	{
		dimSize1 := uint64(o.InterfaceReferencesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.InterfaceReferences {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.InterfaceReferences[i1] != nil {
				if err := o.InterfaceReferences[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.RemoteInterfaceReference{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.InterfaceReferences); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dcom.RemoteInterfaceReference{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteAddReferenceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cInterfaceRefs {in} (1:(uint16))
	{
		if err := w.ReadData(&o.InterfaceReferencesCount); err != nil {
			return err
		}
	}
	// InterfaceRefs {in} (1:[dim:0,size_is=cInterfaceRefs])(2:{alias=REMINTERFACEREF}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceReferences", sizeInfo[0])
		}
		o.InterfaceReferences = make([]*dcom.RemoteInterfaceReference, sizeInfo[0])
		for i1 := range o.InterfaceReferences {
			i1 := i1
			if o.InterfaceReferences[i1] == nil {
				o.InterfaceReferences[i1] = &dcom.RemoteInterfaceReference{}
			}
			if err := o.InterfaceReferences[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteAddReferenceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteAddReferenceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
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
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=cInterfaceRefs](int32))
	{
		dimSize1 := uint64(o.InterfaceReferencesCount)
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteAddReferenceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=cInterfaceRefs](int32))
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteAddReferenceRequest structure represents the RemAddRef operation request
type RemoteAddReferenceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// cInterfaceRefs:  This MUST specify the number of elements sent in the InterfaceRefs
	// parameter and returned in the pResults parameter.
	InterfaceReferencesCount uint16 `idl:"name:cInterfaceRefs" json:"interface_references_count"`
	// InterfaceRefs:  This MUST specify an array of REMINTERFACEREF structures, each of
	// which specifies the number of public and private references to be added to the interface
	// identified by the IPID.
	InterfaceReferences []*dcom.RemoteInterfaceReference `idl:"name:InterfaceRefs;size_is:(cInterfaceRefs)" json:"interface_references"`
}

func (o *RemoteAddReferenceRequest) xxx_ToOp(ctx context.Context) *xxx_RemoteAddReferenceOperation {
	if o == nil {
		return &xxx_RemoteAddReferenceOperation{}
	}
	return &xxx_RemoteAddReferenceOperation{
		This:                     o.This,
		InterfaceReferencesCount: o.InterfaceReferencesCount,
		InterfaceReferences:      o.InterfaceReferences,
	}
}

func (o *RemoteAddReferenceRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteAddReferenceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InterfaceReferencesCount = op.InterfaceReferencesCount
	o.InterfaceReferences = op.InterfaceReferences
}
func (o *RemoteAddReferenceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteAddReferenceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteAddReferenceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteAddReferenceResponse structure represents the RemAddRef operation response
type RemoteAddReferenceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pResults:  This MUST contain an array of HRESULTs specifying the respective success
	// or failure of the RemAddRef operation for each REMINTERFACEREF element.
	Results []int32 `idl:"name:pResults;size_is:(cInterfaceRefs)" json:"results"`
	// Return: The RemAddRef return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteAddReferenceResponse) xxx_ToOp(ctx context.Context) *xxx_RemoteAddReferenceOperation {
	if o == nil {
		return &xxx_RemoteAddReferenceOperation{}
	}
	return &xxx_RemoteAddReferenceOperation{
		That:    o.That,
		Results: o.Results,
		Return:  o.Return,
	}
}

func (o *RemoteAddReferenceResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteAddReferenceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Results = op.Results
	o.Return = op.Return
}
func (o *RemoteAddReferenceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteAddReferenceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteAddReferenceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoteReleaseOperation structure represents the RemRelease operation
type xxx_RemoteReleaseOperation struct {
	This                     *dcom.ORPCThis                   `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat                   `idl:"name:That" json:"that"`
	InterfaceReferencesCount uint16                           `idl:"name:cInterfaceRefs" json:"interface_references_count"`
	InterfaceReferences      []*dcom.RemoteInterfaceReference `idl:"name:InterfaceRefs;size_is:(cInterfaceRefs)" json:"interface_references"`
	Return                   int32                            `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteReleaseOperation) OpNum() int { return 5 }

func (o *xxx_RemoteReleaseOperation) OpName() string { return "/IRemUnknown/v0/RemRelease" }

func (o *xxx_RemoteReleaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InterfaceReferences != nil && o.InterfaceReferencesCount == 0 {
		o.InterfaceReferencesCount = uint16(len(o.InterfaceReferences))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteReleaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
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
	// cInterfaceRefs {in} (1:(uint16))
	{
		if err := w.WriteData(o.InterfaceReferencesCount); err != nil {
			return err
		}
	}
	// InterfaceRefs {in} (1:[dim:0,size_is=cInterfaceRefs])(2:{alias=REMINTERFACEREF}(struct))
	{
		dimSize1 := uint64(o.InterfaceReferencesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.InterfaceReferences {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.InterfaceReferences[i1] != nil {
				if err := o.InterfaceReferences[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.RemoteInterfaceReference{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.InterfaceReferences); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dcom.RemoteInterfaceReference{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteReleaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cInterfaceRefs {in} (1:(uint16))
	{
		if err := w.ReadData(&o.InterfaceReferencesCount); err != nil {
			return err
		}
	}
	// InterfaceRefs {in} (1:[dim:0,size_is=cInterfaceRefs])(2:{alias=REMINTERFACEREF}(struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.InterfaceReferences", sizeInfo[0])
		}
		o.InterfaceReferences = make([]*dcom.RemoteInterfaceReference, sizeInfo[0])
		for i1 := range o.InterfaceReferences {
			i1 := i1
			if o.InterfaceReferences[i1] == nil {
				o.InterfaceReferences[i1] = &dcom.RemoteInterfaceReference{}
			}
			if err := o.InterfaceReferences[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoteReleaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteReleaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteReleaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// RemoteReleaseRequest structure represents the RemRelease operation request
type RemoteReleaseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// cInterfaceRefs:  This MUST specify the number of elements in the InterfaceRefs parameter.
	InterfaceReferencesCount uint16 `idl:"name:cInterfaceRefs" json:"interface_references_count"`
	// InterfaceRefs:  This MUST specify an array of REMINTERFACEREF structures, each of
	// which specifies the number of public and private references to be released on the
	// interface identified by the IPID.
	InterfaceReferences []*dcom.RemoteInterfaceReference `idl:"name:InterfaceRefs;size_is:(cInterfaceRefs)" json:"interface_references"`
}

func (o *RemoteReleaseRequest) xxx_ToOp(ctx context.Context) *xxx_RemoteReleaseOperation {
	if o == nil {
		return &xxx_RemoteReleaseOperation{}
	}
	return &xxx_RemoteReleaseOperation{
		This:                     o.This,
		InterfaceReferencesCount: o.InterfaceReferencesCount,
		InterfaceReferences:      o.InterfaceReferences,
	}
}

func (o *RemoteReleaseRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteReleaseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InterfaceReferencesCount = op.InterfaceReferencesCount
	o.InterfaceReferences = op.InterfaceReferences
}
func (o *RemoteReleaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoteReleaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteReleaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteReleaseResponse structure represents the RemRelease operation response
type RemoteReleaseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemRelease return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteReleaseResponse) xxx_ToOp(ctx context.Context) *xxx_RemoteReleaseOperation {
	if o == nil {
		return &xxx_RemoteReleaseOperation{}
	}
	return &xxx_RemoteReleaseOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RemoteReleaseResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteReleaseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoteReleaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoteReleaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteReleaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
