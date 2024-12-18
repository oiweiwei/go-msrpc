package icatalogutils2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	coma "github.com/oiweiwei/go-msrpc/msrpc/dcom/coma"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = oaut.GoPackage
	_ = coma.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/coma"
)

var (
	// ICatalogUtils2 interface identifier c726744e-5735-4f08-8286-c510ee638fb6
	CatalogUtils2IID = &dcom.IID{Data1: 0xc726744e, Data2: 0x5735, Data3: 0x4f08, Data4: []byte{0x82, 0x86, 0xc5, 0x10, 0xee, 0x63, 0x8f, 0xb6}}
	// Syntax UUID
	CatalogUtils2SyntaxUUID = &uuid.UUID{TimeLow: 0xc726744e, TimeMid: 0x5735, TimeHiAndVersion: 0x4f08, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x86, Node: [6]uint8{0xc5, 0x10, 0xee, 0x63, 0x8f, 0xb6}}
	// Syntax ID
	CatalogUtils2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogUtils2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogUtils2 interface.
type CatalogUtils2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to copy one or more conglomerations from one partition
	// to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CopyConglomerations(context.Context, *CopyConglomerationsRequest, ...dcerpc.CallOption) (*CopyConglomerationsResponse, error)

	// This method is called by a client to copy a component full configuration from one
	// conglomeration to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CopyComponentConfiguration(context.Context, *CopyComponentConfigurationRequest, ...dcerpc.CallOption) (*CopyComponentConfigurationResponse, error)

	// This method is called by a client to create an alias component full configuration,
	// a component full configuration of a virtual aliased component equivalent to the original
	// component except in CLSID and ProgID.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	AliasComponent(context.Context, *AliasComponentRequest, ...dcerpc.CallOption) (*AliasComponentResponse, error)

	// This method is called by a client to move a component configuration from one conglomeration
	// to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	MoveComponentConfiguration(context.Context, *MoveComponentConfigurationRequest, ...dcerpc.CallOption) (*MoveComponentConfigurationResponse, error)

	// This method is called by a client to get information about the event classes associated
	// with an IID that are configured in a specified partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetEventClassesForIid2(context.Context, *GetEventClassesForIid2Request, ...dcerpc.CallOption) (*GetEventClassesForIid2Response, error)

	// The server method ICatalogUtils2::IsSafeToDelete (section 3.1.4.18.6) can be used
	// to determine if a file is safe to delete, but it is usually impossible for a server
	// to reliably make such a determination.
	//
	// A COMA client MUST NOT call the ICatalogUtils2::IsSafeToDelete method unless it receives
	// an explicit request from a client application to do so, and MUST return the results
	// of the call unaltered to the client application.
	IsSafeToDelete(context.Context, *IsSafeToDeleteRequest, ...dcerpc.CallOption) (*IsSafeToDeleteResponse, error)

	// This method is called by a client to request that the server clear its local cache
	// of the entries in domain-controlled PartitionRoles (section 3.1.1.3.17), PartitionRoleMembers
	// (section 3.1.1.3.18), and PartitionUsers (section 3.1.1.3.16) tables, if the server
	// does such lookups with an active directory.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	FlushPartitionCache(context.Context, *FlushPartitionCacheRequest, ...dcerpc.CallOption) (*FlushPartitionCacheResponse, error)

	// This method is called by a client to get an enumeration of software restriction policy
	// (see section 3.1.1.1.9) trust levels supported by the server.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	// The server SHOULD, by some implementation-specific mechanism, attempt to translate
	// the names of the software restriction policy levels it supports to the language specified
	// by Locale, and SHOULD fall back to a default language if it cannot.
	//
	// The server then MUST attempt to set the values referenced by the out parameters as
	// follows: The server MUST attempt to set the value referenced by cLevels to the number
	// of software restriction policy levels the server supports, and the value referenced
	// by aSRPLevels to an array of SRPLevelInfo structures, each of which describes a software
	// restriction policy level, and fail the call if it cannot.
	EnumerateSRPLevels(context.Context, *EnumerateSRPLevelsRequest, ...dcerpc.CallOption) (*EnumerateSRPLevelsResponse, error)

	// This method is called by a client to get a list of component full configurations
	// for a component.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetComponentVersions(context.Context, *GetComponentVersionsRequest, ...dcerpc.CallOption) (*GetComponentVersionsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogUtils2Client
}

type xxx_DefaultCatalogUtils2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogUtils2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogUtils2Client) CopyConglomerations(ctx context.Context, in *CopyConglomerationsRequest, opts ...dcerpc.CallOption) (*CopyConglomerationsResponse, error) {
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
	out := &CopyConglomerationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) CopyComponentConfiguration(ctx context.Context, in *CopyComponentConfigurationRequest, opts ...dcerpc.CallOption) (*CopyComponentConfigurationResponse, error) {
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
	out := &CopyComponentConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) AliasComponent(ctx context.Context, in *AliasComponentRequest, opts ...dcerpc.CallOption) (*AliasComponentResponse, error) {
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
	out := &AliasComponentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) MoveComponentConfiguration(ctx context.Context, in *MoveComponentConfigurationRequest, opts ...dcerpc.CallOption) (*MoveComponentConfigurationResponse, error) {
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
	out := &MoveComponentConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) GetEventClassesForIid2(ctx context.Context, in *GetEventClassesForIid2Request, opts ...dcerpc.CallOption) (*GetEventClassesForIid2Response, error) {
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
	out := &GetEventClassesForIid2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) IsSafeToDelete(ctx context.Context, in *IsSafeToDeleteRequest, opts ...dcerpc.CallOption) (*IsSafeToDeleteResponse, error) {
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
	out := &IsSafeToDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) FlushPartitionCache(ctx context.Context, in *FlushPartitionCacheRequest, opts ...dcerpc.CallOption) (*FlushPartitionCacheResponse, error) {
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
	out := &FlushPartitionCacheResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) EnumerateSRPLevels(ctx context.Context, in *EnumerateSRPLevelsRequest, opts ...dcerpc.CallOption) (*EnumerateSRPLevelsResponse, error) {
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
	out := &EnumerateSRPLevelsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) GetComponentVersions(ctx context.Context, in *GetComponentVersionsRequest, opts ...dcerpc.CallOption) (*GetComponentVersionsResponse, error) {
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
	out := &GetComponentVersionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogUtils2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogUtils2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCatalogUtils2Client) IPID(ctx context.Context, ipid *dcom.IPID) CatalogUtils2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogUtils2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewCatalogUtils2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogUtils2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogUtils2SyntaxV0_0))...)
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
	return &xxx_DefaultCatalogUtils2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CopyConglomerationsOperation structure represents the CopyConglomerations operation
type xxx_CopyConglomerationsOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourcePartition          string         `idl:"name:pwszSourcePartition;string" json:"source_partition"`
	DestinationPartition     string         `idl:"name:pwszDestPartition;string" json:"destination_partition"`
	ConglomerationsCount     uint32         `idl:"name:cConglomerations" json:"conglomerations_count"`
	ConglomerationNamesOrIDs []string       `idl:"name:ppwszConglomerationNamesOrIds;size_is:(cConglomerations, );string" json:"conglomeration_names_or_ids"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyConglomerationsOperation) OpNum() int { return 3 }

func (o *xxx_CopyConglomerationsOperation) OpName() string {
	return "/ICatalogUtils2/v0/CopyConglomerations"
}

func (o *xxx_CopyConglomerationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ConglomerationNamesOrIDs != nil && o.ConglomerationsCount == 0 {
		o.ConglomerationsCount = uint32(len(o.ConglomerationNamesOrIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyConglomerationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszSourcePartition {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SourcePartition); err != nil {
			return err
		}
	}
	// pwszDestPartition {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DestinationPartition); err != nil {
			return err
		}
	}
	// cConglomerations {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ConglomerationsCount); err != nil {
			return err
		}
	}
	// ppwszConglomerationNamesOrIds {in} (1:{string, pointer=ref}*(1))(2:{alias=LPCWSTR}[dim:0,size_is=cConglomerations]*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.ConglomerationsCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ConglomerationNamesOrIDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ConglomerationNamesOrIDs[i1] != "" {
				_ptr_ppwszConglomerationNamesOrIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.ConglomerationNamesOrIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.ConglomerationNamesOrIDs[i1], _ptr_ppwszConglomerationNamesOrIds); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ConglomerationNamesOrIDs); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_CopyConglomerationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszSourcePartition {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourcePartition); err != nil {
			return err
		}
	}
	// pwszDestPartition {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationPartition); err != nil {
			return err
		}
	}
	// cConglomerations {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ConglomerationsCount); err != nil {
			return err
		}
	}
	// ppwszConglomerationNamesOrIds {in} (1:{string, pointer=ref}*(1))(2:{alias=LPCWSTR}[dim:0,size_is=cConglomerations]*(1)[dim:0,string,null](wchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ConglomerationNamesOrIDs", sizeInfo[0])
		}
		o.ConglomerationNamesOrIDs = make([]string, sizeInfo[0])
		for i1 := range o.ConglomerationNamesOrIDs {
			i1 := i1
			_ptr_ppwszConglomerationNamesOrIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.ConglomerationNamesOrIDs[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ppwszConglomerationNamesOrIds := func(ptr interface{}) { o.ConglomerationNamesOrIDs[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.ConglomerationNamesOrIDs[i1], _s_ppwszConglomerationNamesOrIds, _ptr_ppwszConglomerationNamesOrIds); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyConglomerationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyConglomerationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CopyConglomerationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CopyConglomerationsRequest structure represents the CopyConglomerations operation request
type CopyConglomerationsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszSourcePartition: Either the Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the partition identifier or the Name property of a partition,
	// from which conglomerations are to be copied.
	SourcePartition string `idl:"name:pwszSourcePartition;string" json:"source_partition"`
	// pwszDestPartition: Either the Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the partition identifier or the Name property of a partition,
	// to which conglomerations are to be copied.
	DestinationPartition string `idl:"name:pwszDestPartition;string" json:"destination_partition"`
	// cConglomerations: The number of elements in ppwszConglomerationNamesOrIds.
	ConglomerationsCount uint32 `idl:"name:cConglomerations" json:"conglomerations_count"`
	// ppwszConglomerationNamesOrIds: An array of values, each of which is either the Curly
	// Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation of the conglomeration
	// identifier or the Name property of a conglomeration to be copied.
	ConglomerationNamesOrIDs []string `idl:"name:ppwszConglomerationNamesOrIds;size_is:(cConglomerations, );string" json:"conglomeration_names_or_ids"`
}

func (o *CopyConglomerationsRequest) xxx_ToOp(ctx context.Context) *xxx_CopyConglomerationsOperation {
	if o == nil {
		return &xxx_CopyConglomerationsOperation{}
	}
	return &xxx_CopyConglomerationsOperation{
		This:                     o.This,
		SourcePartition:          o.SourcePartition,
		DestinationPartition:     o.DestinationPartition,
		ConglomerationsCount:     o.ConglomerationsCount,
		ConglomerationNamesOrIDs: o.ConglomerationNamesOrIDs,
	}
}

func (o *CopyConglomerationsRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyConglomerationsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourcePartition = op.SourcePartition
	o.DestinationPartition = op.DestinationPartition
	o.ConglomerationsCount = op.ConglomerationsCount
	o.ConglomerationNamesOrIDs = op.ConglomerationNamesOrIDs
}
func (o *CopyConglomerationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CopyConglomerationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyConglomerationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyConglomerationsResponse structure represents the CopyConglomerations operation response
type CopyConglomerationsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CopyConglomerations return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyConglomerationsResponse) xxx_ToOp(ctx context.Context) *xxx_CopyConglomerationsOperation {
	if o == nil {
		return &xxx_CopyConglomerationsOperation{}
	}
	return &xxx_CopyConglomerationsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CopyConglomerationsResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyConglomerationsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CopyConglomerationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CopyConglomerationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyConglomerationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CopyComponentConfigurationOperation structure represents the CopyComponentConfiguration operation
type xxx_CopyComponentConfigurationOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceConglomeration      string         `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	Component                 string         `idl:"name:pwszComponent;string" json:"component"`
	DestinationConglomeration string         `idl:"name:pwszDestConglomeration;string" json:"destination_conglomeration"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyComponentConfigurationOperation) OpNum() int { return 4 }

func (o *xxx_CopyComponentConfigurationOperation) OpName() string {
	return "/ICatalogUtils2/v0/CopyComponentConfiguration"
}

func (o *xxx_CopyComponentConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyComponentConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Component); err != nil {
			return err
		}
	}
	// pwszDestConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DestinationConglomeration); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyComponentConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Component); err != nil {
			return err
		}
	}
	// pwszDestConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationConglomeration); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyComponentConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyComponentConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CopyComponentConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CopyComponentConfigurationRequest structure represents the CopyComponentConfiguration operation request
type CopyComponentConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszSourceConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the conglomeration identifier or the Name property of
	// a conglomeration from which the component configuration is to be copied.
	SourceConglomeration string `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	// pwszComponent: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of the CLSID or the ProgID property of a component configured in the conglomeration
	// specified by pwszSourceConglomeration.
	Component string `idl:"name:pwszComponent;string" json:"component"`
	// pwszDestConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3)
	// representation of the conglomeration identifier or the Name property of a conglomeration
	// into which the component configuration is to be copied.
	DestinationConglomeration string `idl:"name:pwszDestConglomeration;string" json:"destination_conglomeration"`
}

func (o *CopyComponentConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_CopyComponentConfigurationOperation {
	if o == nil {
		return &xxx_CopyComponentConfigurationOperation{}
	}
	return &xxx_CopyComponentConfigurationOperation{
		This:                      o.This,
		SourceConglomeration:      o.SourceConglomeration,
		Component:                 o.Component,
		DestinationConglomeration: o.DestinationConglomeration,
	}
}

func (o *CopyComponentConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyComponentConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceConglomeration = op.SourceConglomeration
	o.Component = op.Component
	o.DestinationConglomeration = op.DestinationConglomeration
}
func (o *CopyComponentConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CopyComponentConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyComponentConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyComponentConfigurationResponse structure represents the CopyComponentConfiguration operation response
type CopyComponentConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CopyComponentConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyComponentConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_CopyComponentConfigurationOperation {
	if o == nil {
		return &xxx_CopyComponentConfigurationOperation{}
	}
	return &xxx_CopyComponentConfigurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CopyComponentConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyComponentConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CopyComponentConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CopyComponentConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyComponentConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AliasComponentOperation structure represents the AliasComponent operation
type xxx_AliasComponentOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceConglomeration      string         `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	Component                 string         `idl:"name:pwszComponent;string" json:"component"`
	DestinationConglomeration string         `idl:"name:pwszDestConglomeration;string" json:"destination_conglomeration"`
	NewClassID                *dtyp.GUID     `idl:"name:pNewCLSID" json:"new_class_id"`
	NewProgID                 string         `idl:"name:pwszNewProgID;string" json:"new_prog_id"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AliasComponentOperation) OpNum() int { return 5 }

func (o *xxx_AliasComponentOperation) OpName() string { return "/ICatalogUtils2/v0/AliasComponent" }

func (o *xxx_AliasComponentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AliasComponentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Component); err != nil {
			return err
		}
	}
	// pwszDestConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DestinationConglomeration); err != nil {
			return err
		}
	}
	// pNewCLSID {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.NewClassID != nil {
			if err := o.NewClassID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pwszNewProgID {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewProgID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AliasComponentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Component); err != nil {
			return err
		}
	}
	// pwszDestConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationConglomeration); err != nil {
			return err
		}
	}
	// pNewCLSID {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.NewClassID == nil {
			o.NewClassID = &dtyp.GUID{}
		}
		if err := o.NewClassID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pwszNewProgID {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewProgID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AliasComponentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AliasComponentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AliasComponentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AliasComponentRequest structure represents the AliasComponent operation request
type AliasComponentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszSourceConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the conglomeration identifier or the Name property of
	// a conglomeration from which the component configuration is to be copied.
	SourceConglomeration string `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	// pwszComponent: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of the CLSID or the ProgID property of a component configured in the specified by
	// pwszSourceConglomeration.
	Component string `idl:"name:pwszComponent;string" json:"component"`
	// pwszDestConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3)
	// representation of the conglomeration identifier or the Name property of a conglomeration
	// into which the component configuration is to be copied.
	DestinationConglomeration string `idl:"name:pwszDestConglomeration;string" json:"destination_conglomeration"`
	// pNewCLSID: A GUID to use as the CLSID of the aliased component.
	NewClassID *dtyp.GUID `idl:"name:pNewCLSID" json:"new_class_id"`
	// pwszNewProgID: A string to be used as the ProgID of the aliased component.
	NewProgID string `idl:"name:pwszNewProgID;string" json:"new_prog_id"`
}

func (o *AliasComponentRequest) xxx_ToOp(ctx context.Context) *xxx_AliasComponentOperation {
	if o == nil {
		return &xxx_AliasComponentOperation{}
	}
	return &xxx_AliasComponentOperation{
		This:                      o.This,
		SourceConglomeration:      o.SourceConglomeration,
		Component:                 o.Component,
		DestinationConglomeration: o.DestinationConglomeration,
		NewClassID:                o.NewClassID,
		NewProgID:                 o.NewProgID,
	}
}

func (o *AliasComponentRequest) xxx_FromOp(ctx context.Context, op *xxx_AliasComponentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceConglomeration = op.SourceConglomeration
	o.Component = op.Component
	o.DestinationConglomeration = op.DestinationConglomeration
	o.NewClassID = op.NewClassID
	o.NewProgID = op.NewProgID
}
func (o *AliasComponentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AliasComponentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AliasComponentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AliasComponentResponse structure represents the AliasComponent operation response
type AliasComponentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AliasComponent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AliasComponentResponse) xxx_ToOp(ctx context.Context) *xxx_AliasComponentOperation {
	if o == nil {
		return &xxx_AliasComponentOperation{}
	}
	return &xxx_AliasComponentOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AliasComponentResponse) xxx_FromOp(ctx context.Context, op *xxx_AliasComponentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AliasComponentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AliasComponentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AliasComponentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveComponentConfigurationOperation structure represents the MoveComponentConfiguration operation
type xxx_MoveComponentConfigurationOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	SourceConglomeration      string         `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	Component                 string         `idl:"name:pwszComponent;string" json:"component"`
	DestinationConglomeration string         `idl:"name:pwszDestinationConglomeration;string" json:"destination_conglomeration"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveComponentConfigurationOperation) OpNum() int { return 6 }

func (o *xxx_MoveComponentConfigurationOperation) OpName() string {
	return "/ICatalogUtils2/v0/MoveComponentConfiguration"
}

func (o *xxx_MoveComponentConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveComponentConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Component); err != nil {
			return err
		}
	}
	// pwszDestinationConglomeration {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DestinationConglomeration); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveComponentConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszSourceConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SourceConglomeration); err != nil {
			return err
		}
	}
	// pwszComponent {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Component); err != nil {
			return err
		}
	}
	// pwszDestinationConglomeration {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DestinationConglomeration); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveComponentConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveComponentConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MoveComponentConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MoveComponentConfigurationRequest structure represents the MoveComponentConfiguration operation request
type MoveComponentConfigurationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszSourceConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the conglomeration identifier or the Name property of
	// a conglomeration from which the component configuration is to be moved.
	SourceConglomeration string `idl:"name:pwszSourceConglomeration;string" json:"source_conglomeration"`
	// pwszComponent: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of the CLSID or the ProgID property of a component configured in the conglomeration
	// specified by pwszSourceConglomeration.
	Component string `idl:"name:pwszComponent;string" json:"component"`
	// pwszDestinationConglomeration: The Curly Braced GUID String Syntax ([MS-DTYP] section
	// 2.3.4.3) representation of the conglomeration identifier or the Name property of
	// a conglomeration into which the component configuration is to be moved.
	DestinationConglomeration string `idl:"name:pwszDestinationConglomeration;string" json:"destination_conglomeration"`
}

func (o *MoveComponentConfigurationRequest) xxx_ToOp(ctx context.Context) *xxx_MoveComponentConfigurationOperation {
	if o == nil {
		return &xxx_MoveComponentConfigurationOperation{}
	}
	return &xxx_MoveComponentConfigurationOperation{
		This:                      o.This,
		SourceConglomeration:      o.SourceConglomeration,
		Component:                 o.Component,
		DestinationConglomeration: o.DestinationConglomeration,
	}
}

func (o *MoveComponentConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveComponentConfigurationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SourceConglomeration = op.SourceConglomeration
	o.Component = op.Component
	o.DestinationConglomeration = op.DestinationConglomeration
}
func (o *MoveComponentConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MoveComponentConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveComponentConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveComponentConfigurationResponse structure represents the MoveComponentConfiguration operation response
type MoveComponentConfigurationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MoveComponentConfiguration return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MoveComponentConfigurationResponse) xxx_ToOp(ctx context.Context) *xxx_MoveComponentConfigurationOperation {
	if o == nil {
		return &xxx_MoveComponentConfigurationOperation{}
	}
	return &xxx_MoveComponentConfigurationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *MoveComponentConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveComponentConfigurationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MoveComponentConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MoveComponentConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveComponentConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventClassesForIid2Operation structure represents the GetEventClassesForIID2 operation
type xxx_GetEventClassesForIid2Operation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	IID               string         `idl:"name:wszIID;string;pointer:unique" json:"iid"`
	PartitionID       *dtyp.GUID     `idl:"name:PartitionId" json:"partition_id"`
	ClassesCount      uint32         `idl:"name:pcClasses" json:"classes_count"`
	ClassIDs          []string       `idl:"name:pawszCLSIDs;size_is:(, pcClasses);string" json:"class_ids"`
	ProgIDs           []string       `idl:"name:pawszProgIDs;size_is:(, pcClasses);string" json:"prog_ids"`
	Descriptions      []string       `idl:"name:pawszDescriptions;size_is:(, pcClasses);string" json:"descriptions"`
	ConglomerationIDs []string       `idl:"name:pawszConglomerationIDs;size_is:(, pcClasses);string" json:"conglomeration_ids"`
	IsPrivate         []uint32       `idl:"name:padwIsPrivate;size_is:(, pcClasses)" json:"is_private"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassesForIid2Operation) OpNum() int { return 7 }

func (o *xxx_GetEventClassesForIid2Operation) OpName() string {
	return "/ICatalogUtils2/v0/GetEventClassesForIID2"
}

func (o *xxx_GetEventClassesForIid2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIid2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wszIID {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.IID != "" {
			_ptr_wszIID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.IID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.IID, _ptr_wszIID); err != nil {
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
	// PartitionId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID != nil {
			if err := o.PartitionID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIid2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wszIID {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_wszIID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.IID); err != nil {
				return err
			}
			return nil
		})
		_s_wszIID := func(ptr interface{}) { o.IID = *ptr.(*string) }
		if err := w.ReadPointer(&o.IID, _s_wszIID, _ptr_wszIID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PartitionId {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.PartitionID == nil {
			o.PartitionID = &dtyp.GUID{}
		}
		if err := o.PartitionID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIid2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ClassIDs != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.ClassIDs))
	}
	if o.ProgIDs != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.ProgIDs))
	}
	if o.Descriptions != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.Descriptions))
	}
	if o.ConglomerationIDs != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.ConglomerationIDs))
	}
	if o.IsPrivate != nil && o.ClassesCount == 0 {
		o.ClassesCount = uint32(len(o.IsPrivate))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassesForIid2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pcClasses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClassesCount); err != nil {
			return err
		}
	}
	// pawszCLSIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.ClassIDs != nil || o.ClassesCount > 0 {
			_ptr_pawszCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ClassIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ClassIDs[i1] != "" {
						_ptr_pawszCLSIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ClassIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ClassIDs[i1], _ptr_pawszCLSIDs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ClassIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClassIDs, _ptr_pawszCLSIDs); err != nil {
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
	// pawszProgIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.ProgIDs != nil || o.ClassesCount > 0 {
			_ptr_pawszProgIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ProgIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ProgIDs[i1] != "" {
						_ptr_pawszProgIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ProgIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ProgIDs[i1], _ptr_pawszProgIDs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ProgIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProgIDs, _ptr_pawszProgIDs); err != nil {
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
	// pawszDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.Descriptions != nil || o.ClassesCount > 0 {
			_ptr_pawszDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Descriptions {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Descriptions[i1] != "" {
						_ptr_pawszDescriptions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.Descriptions[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.Descriptions[i1], _ptr_pawszDescriptions); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Descriptions); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Descriptions, _ptr_pawszDescriptions); err != nil {
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
	// pawszConglomerationIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		if o.ConglomerationIDs != nil || o.ClassesCount > 0 {
			_ptr_pawszConglomerationIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ConglomerationIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ConglomerationIDs[i1] != "" {
						_ptr_pawszConglomerationIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.ConglomerationIDs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.ConglomerationIDs[i1], _ptr_pawszConglomerationIDs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ConglomerationIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConglomerationIDs, _ptr_pawszConglomerationIDs); err != nil {
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
	// padwIsPrivate {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcClasses](uint32))
	{
		if o.IsPrivate != nil || o.ClassesCount > 0 {
			_ptr_padwIsPrivate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ClassesCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.IsPrivate {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.IsPrivate[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.IsPrivate); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IsPrivate, _ptr_padwIsPrivate); err != nil {
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

func (o *xxx_GetEventClassesForIid2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pcClasses {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClassesCount); err != nil {
			return err
		}
	}
	// pawszCLSIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ClassIDs", sizeInfo[0])
			}
			o.ClassIDs = make([]string, sizeInfo[0])
			for i1 := range o.ClassIDs {
				i1 := i1
				_ptr_pawszCLSIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ClassIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszCLSIDs := func(ptr interface{}) { o.ClassIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ClassIDs[i1], _s_pawszCLSIDs, _ptr_pawszCLSIDs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszCLSIDs := func(ptr interface{}) { o.ClassIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ClassIDs, _s_pawszCLSIDs, _ptr_pawszCLSIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pawszProgIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszProgIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ProgIDs", sizeInfo[0])
			}
			o.ProgIDs = make([]string, sizeInfo[0])
			for i1 := range o.ProgIDs {
				i1 := i1
				_ptr_pawszProgIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ProgIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszProgIDs := func(ptr interface{}) { o.ProgIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ProgIDs[i1], _s_pawszProgIDs, _ptr_pawszProgIDs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszProgIDs := func(ptr interface{}) { o.ProgIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ProgIDs, _s_pawszProgIDs, _ptr_pawszProgIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pawszDescriptions {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Descriptions", sizeInfo[0])
			}
			o.Descriptions = make([]string, sizeInfo[0])
			for i1 := range o.Descriptions {
				i1 := i1
				_ptr_pawszDescriptions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.Descriptions[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszDescriptions := func(ptr interface{}) { o.Descriptions[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.Descriptions[i1], _s_pawszDescriptions, _ptr_pawszDescriptions); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszDescriptions := func(ptr interface{}) { o.Descriptions = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.Descriptions, _s_pawszDescriptions, _ptr_pawszDescriptions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pawszConglomerationIDs {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pcClasses]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pawszConglomerationIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ConglomerationIDs", sizeInfo[0])
			}
			o.ConglomerationIDs = make([]string, sizeInfo[0])
			for i1 := range o.ConglomerationIDs {
				i1 := i1
				_ptr_pawszConglomerationIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.ConglomerationIDs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pawszConglomerationIDs := func(ptr interface{}) { o.ConglomerationIDs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.ConglomerationIDs[i1], _s_pawszConglomerationIDs, _ptr_pawszConglomerationIDs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pawszConglomerationIDs := func(ptr interface{}) { o.ConglomerationIDs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.ConglomerationIDs, _s_pawszConglomerationIDs, _ptr_pawszConglomerationIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// padwIsPrivate {out} (1:{pointer=ref}*(2)*(1))(2:{alias=DWORD}[dim:0,size_is=pcClasses](uint32))
	{
		_ptr_padwIsPrivate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.IsPrivate", sizeInfo[0])
			}
			o.IsPrivate = make([]uint32, sizeInfo[0])
			for i1 := range o.IsPrivate {
				i1 := i1
				if err := w.ReadData(&o.IsPrivate[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_padwIsPrivate := func(ptr interface{}) { o.IsPrivate = *ptr.(*[]uint32) }
		if err := w.ReadPointer(&o.IsPrivate, _s_padwIsPrivate, _ptr_padwIsPrivate); err != nil {
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

// GetEventClassesForIid2Request structure represents the GetEventClassesForIID2 operation request
type GetEventClassesForIid2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// wszIID: The Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representation
	// of the IID for which event classes will be retrieved, or NULL or an empty (zero-length)
	// string to indicate all event classes.
	IID string `idl:"name:wszIID;string;pointer:unique" json:"iid"`
	// PartitionId: The partition identifier of a partition within which to limit the selection
	// of configurations of event classes.
	PartitionID *dtyp.GUID `idl:"name:PartitionId" json:"partition_id"`
}

func (o *GetEventClassesForIid2Request) xxx_ToOp(ctx context.Context) *xxx_GetEventClassesForIid2Operation {
	if o == nil {
		return &xxx_GetEventClassesForIid2Operation{}
	}
	return &xxx_GetEventClassesForIid2Operation{
		This:        o.This,
		IID:         o.IID,
		PartitionID: o.PartitionID,
	}
}

func (o *GetEventClassesForIid2Request) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassesForIid2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IID = op.IID
	o.PartitionID = op.PartitionID
}
func (o *GetEventClassesForIid2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassesForIid2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassesForIid2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassesForIid2Response structure represents the GetEventClassesForIID2 operation response
type GetEventClassesForIid2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pcClasses: A pointer to a value that, upon successful completion, MUST be set to
	// the number of event classes for which information was returned.
	ClassesCount uint32 `idl:"name:pcClasses" json:"classes_count"`
	// pawszCLSIDs: A pointer to a value that, upon successful completion, MUST be set to
	// an array of Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3) representations
	// of CLSIDs of event classes.
	ClassIDs []string `idl:"name:pawszCLSIDs;size_is:(, pcClasses);string" json:"class_ids"`
	// pawszProgIDs: A pointer to a value that, upon successful completion, MUST be set
	// to an array of ProgIDs of event classes, in the same order as pawszCLSIDs.
	ProgIDs []string `idl:"name:pawszProgIDs;size_is:(, pcClasses);string" json:"prog_ids"`
	// pawszDescriptions: A pointer to a value that, upon successful completion, MUST be
	// set to an array of descriptions of event classes, in the same order as pawszCLSIDs.
	Descriptions []string `idl:"name:pawszDescriptions;size_is:(, pcClasses);string" json:"descriptions"`
	// pawszConglomerationIDs: A pointer to a value that, upon successful completion, MUST
	// be set to an array of Curly Braced GUID String Syntax ([MS-DTYP] section 2.3.4.3)
	// representations of the conglomerations in which the event classes are configuration,
	// in the same order as pawszCLSIDs.
	ConglomerationIDs []string `idl:"name:pawszConglomerationIDs;size_is:(, pcClasses);string" json:"conglomeration_ids"`
	// padwIsPrivate: A pointer to a value that, upon successful completion, MUST be set
	// to an array of value indicating whether the configurations are private, in other
	// words the IsPrivate property has the value TRUE (0x000000001), in the same order
	// as pawszCLSIDs.
	IsPrivate []uint32 `idl:"name:padwIsPrivate;size_is:(, pcClasses)" json:"is_private"`
	// Return: The GetEventClassesForIID2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassesForIid2Response) xxx_ToOp(ctx context.Context) *xxx_GetEventClassesForIid2Operation {
	if o == nil {
		return &xxx_GetEventClassesForIid2Operation{}
	}
	return &xxx_GetEventClassesForIid2Operation{
		That:              o.That,
		ClassesCount:      o.ClassesCount,
		ClassIDs:          o.ClassIDs,
		ProgIDs:           o.ProgIDs,
		Descriptions:      o.Descriptions,
		ConglomerationIDs: o.ConglomerationIDs,
		IsPrivate:         o.IsPrivate,
		Return:            o.Return,
	}
}

func (o *GetEventClassesForIid2Response) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassesForIid2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClassesCount = op.ClassesCount
	o.ClassIDs = op.ClassIDs
	o.ProgIDs = op.ProgIDs
	o.Descriptions = op.Descriptions
	o.ConglomerationIDs = op.ConglomerationIDs
	o.IsPrivate = op.IsPrivate
	o.Return = op.Return
}
func (o *GetEventClassesForIid2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassesForIid2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassesForIid2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsSafeToDeleteOperation structure represents the IsSafeToDelete operation
type xxx_IsSafeToDeleteOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	File   *oaut.String   `idl:"name:bstrFile" json:"file"`
	InUse  int32          `idl:"name:pInUse" json:"in_use"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsSafeToDeleteOperation) OpNum() int { return 8 }

func (o *xxx_IsSafeToDeleteOperation) OpName() string { return "/ICatalogUtils2/v0/IsSafeToDelete" }

func (o *xxx_IsSafeToDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSafeToDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrFile {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.File != nil {
			_ptr_bstrFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.File != nil {
					if err := o.File.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.File, _ptr_bstrFile); err != nil {
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

func (o *xxx_IsSafeToDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrFile {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.File == nil {
				o.File = &oaut.String{}
			}
			if err := o.File.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrFile := func(ptr interface{}) { o.File = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.File, _s_bstrFile, _ptr_bstrFile); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSafeToDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsSafeToDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pInUse {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.InUse); err != nil {
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

func (o *xxx_IsSafeToDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pInUse {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.InUse); err != nil {
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

// IsSafeToDeleteRequest structure represents the IsSafeToDelete operation request
type IsSafeToDeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	File *oaut.String   `idl:"name:bstrFile" json:"file"`
}

func (o *IsSafeToDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_IsSafeToDeleteOperation {
	if o == nil {
		return &xxx_IsSafeToDeleteOperation{}
	}
	return &xxx_IsSafeToDeleteOperation{
		This: o.This,
		File: o.File,
	}
}

func (o *IsSafeToDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_IsSafeToDeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.File = op.File
}
func (o *IsSafeToDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsSafeToDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSafeToDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsSafeToDeleteResponse structure represents the IsSafeToDelete operation response
type IsSafeToDeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	InUse int32          `idl:"name:pInUse" json:"in_use"`
	// Return: The IsSafeToDelete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsSafeToDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_IsSafeToDeleteOperation {
	if o == nil {
		return &xxx_IsSafeToDeleteOperation{}
	}
	return &xxx_IsSafeToDeleteOperation{
		That:   o.That,
		InUse:  o.InUse,
		Return: o.Return,
	}
}

func (o *IsSafeToDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_IsSafeToDeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.InUse = op.InUse
	o.Return = op.Return
}
func (o *IsSafeToDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsSafeToDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsSafeToDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FlushPartitionCacheOperation structure represents the FlushPartitionCache operation
type xxx_FlushPartitionCacheOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FlushPartitionCacheOperation) OpNum() int { return 9 }

func (o *xxx_FlushPartitionCacheOperation) OpName() string {
	return "/ICatalogUtils2/v0/FlushPartitionCache"
}

func (o *xxx_FlushPartitionCacheOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushPartitionCacheOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_FlushPartitionCacheOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_FlushPartitionCacheOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushPartitionCacheOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FlushPartitionCacheOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FlushPartitionCacheRequest structure represents the FlushPartitionCache operation request
type FlushPartitionCacheRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *FlushPartitionCacheRequest) xxx_ToOp(ctx context.Context) *xxx_FlushPartitionCacheOperation {
	if o == nil {
		return &xxx_FlushPartitionCacheOperation{}
	}
	return &xxx_FlushPartitionCacheOperation{
		This: o.This,
	}
}

func (o *FlushPartitionCacheRequest) xxx_FromOp(ctx context.Context, op *xxx_FlushPartitionCacheOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *FlushPartitionCacheRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FlushPartitionCacheRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushPartitionCacheOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FlushPartitionCacheResponse structure represents the FlushPartitionCache operation response
type FlushPartitionCacheResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FlushPartitionCache return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FlushPartitionCacheResponse) xxx_ToOp(ctx context.Context) *xxx_FlushPartitionCacheOperation {
	if o == nil {
		return &xxx_FlushPartitionCacheOperation{}
	}
	return &xxx_FlushPartitionCacheOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *FlushPartitionCacheResponse) xxx_FromOp(ctx context.Context, op *xxx_FlushPartitionCacheOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *FlushPartitionCacheResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FlushPartitionCacheResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushPartitionCacheOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateSRPLevelsOperation structure represents the EnumerateSRPLevels operation
type xxx_EnumerateSRPLevelsOperation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	Locale      uint32               `idl:"name:Locale" json:"locale"`
	LevelsCount int32                `idl:"name:cLevels" json:"levels_count"`
	SRPLevels   []*coma.SRPLevelInfo `idl:"name:aSRPLevels;size_is:(, cLevels)" json:"srp_levels"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateSRPLevelsOperation) OpNum() int { return 10 }

func (o *xxx_EnumerateSRPLevelsOperation) OpName() string {
	return "/ICatalogUtils2/v0/EnumerateSRPLevels"
}

func (o *xxx_EnumerateSRPLevelsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateSRPLevelsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Locale); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateSRPLevelsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Locale {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Locale); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateSRPLevelsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SRPLevels != nil && o.LevelsCount == 0 {
		o.LevelsCount = int32(len(o.SRPLevels))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateSRPLevelsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// cLevels {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.LevelsCount); err != nil {
			return err
		}
	}
	// aSRPLevels {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SRPLevelInfo}[dim:0,size_is=cLevels](struct))
	{
		if o.SRPLevels != nil || o.LevelsCount > 0 {
			_ptr_aSRPLevels := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.LevelsCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SRPLevels {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SRPLevels[i1] != nil {
						if err := o.SRPLevels[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&coma.SRPLevelInfo{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SRPLevels); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&coma.SRPLevelInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SRPLevels, _ptr_aSRPLevels); err != nil {
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

func (o *xxx_EnumerateSRPLevelsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// cLevels {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.LevelsCount); err != nil {
			return err
		}
	}
	// aSRPLevels {out} (1:{pointer=ref}*(2)*(1))(2:{alias=SRPLevelInfo}[dim:0,size_is=cLevels](struct))
	{
		_ptr_aSRPLevels := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SRPLevels", sizeInfo[0])
			}
			o.SRPLevels = make([]*coma.SRPLevelInfo, sizeInfo[0])
			for i1 := range o.SRPLevels {
				i1 := i1
				if o.SRPLevels[i1] == nil {
					o.SRPLevels[i1] = &coma.SRPLevelInfo{}
				}
				if err := o.SRPLevels[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_aSRPLevels := func(ptr interface{}) { o.SRPLevels = *ptr.(*[]*coma.SRPLevelInfo) }
		if err := w.ReadPointer(&o.SRPLevels, _s_aSRPLevels, _ptr_aSRPLevels); err != nil {
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

// EnumerateSRPLevelsRequest structure represents the EnumerateSRPLevels operation request
type EnumerateSRPLevelsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Locale: The language code identifier, as specified in [MS-LCID], for the language
	// into which the descriptive strings for each level are to be translated, if possible.
	Locale uint32 `idl:"name:Locale" json:"locale"`
}

func (o *EnumerateSRPLevelsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateSRPLevelsOperation {
	if o == nil {
		return &xxx_EnumerateSRPLevelsOperation{}
	}
	return &xxx_EnumerateSRPLevelsOperation{
		This:   o.This,
		Locale: o.Locale,
	}
}

func (o *EnumerateSRPLevelsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateSRPLevelsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Locale = op.Locale
}
func (o *EnumerateSRPLevelsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateSRPLevelsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateSRPLevelsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateSRPLevelsResponse structure represents the EnumerateSRPLevels operation response
type EnumerateSRPLevelsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// cLevels: A pointer to a variable that, upon successful completion, MUST be set to
	// the number of elements in aSRPLevels.
	LevelsCount int32 `idl:"name:cLevels" json:"levels_count"`
	// aSRPLevels: A pointer to a variable that, upon successful completion, MUST be set
	// to an array of SRPLevelInfo (section 2.2.6) structures representing the software
	// restriction policy levels that the server defines.
	SRPLevels []*coma.SRPLevelInfo `idl:"name:aSRPLevels;size_is:(, cLevels)" json:"srp_levels"`
	// Return: The EnumerateSRPLevels return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateSRPLevelsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateSRPLevelsOperation {
	if o == nil {
		return &xxx_EnumerateSRPLevelsOperation{}
	}
	return &xxx_EnumerateSRPLevelsOperation{
		That:        o.That,
		LevelsCount: o.LevelsCount,
		SRPLevels:   o.SRPLevels,
		Return:      o.Return,
	}
}

func (o *EnumerateSRPLevelsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateSRPLevelsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LevelsCount = op.LevelsCount
	o.SRPLevels = op.SRPLevels
	o.Return = op.Return
}
func (o *EnumerateSRPLevelsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateSRPLevelsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateSRPLevelsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetComponentVersionsOperation structure represents the GetComponentVersions operation
type xxx_GetComponentVersionsOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ClassIDOrProgID   string         `idl:"name:pwszClsidOrProgId" json:"class_id_or_prog_id"`
	Versions          uint32         `idl:"name:pdwVersions" json:"versions"`
	PartitionIDs      []*dtyp.GUID   `idl:"name:ppPartitionIDs;size_is:(, pdwVersions)" json:"partition_ids"`
	ConglomerationIDs []*dtyp.GUID   `idl:"name:ppConglomerationIDs;size_is:(, pdwVersions)" json:"conglomeration_ids"`
	IsPrivate         []bool         `idl:"name:ppIsPrivate;size_is:(, pdwVersions)" json:"is_private"`
	Bitness           []int32        `idl:"name:ppBitness;size_is:(, pdwVersions)" json:"bitness"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetComponentVersionsOperation) OpNum() int { return 11 }

func (o *xxx_GetComponentVersionsOperation) OpName() string {
	return "/ICatalogUtils2/v0/GetComponentVersions"
}

func (o *xxx_GetComponentVersionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentVersionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszClsidOrProgId {in} (1:{alias=LPCWSTR}*(1)[dim:0,string](wchar))
	{
		if err := ndr.WriteUTF16String(ctx, w, o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentVersionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszClsidOrProgId {in} (1:{alias=LPCWSTR,pointer=ref}*(1)[dim:0,string](wchar))
	{
		if err := ndr.ReadUTF16String(ctx, w, &o.ClassIDOrProgID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentVersionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PartitionIDs != nil && o.Versions == 0 {
		o.Versions = uint32(len(o.PartitionIDs))
	}
	if o.ConglomerationIDs != nil && o.Versions == 0 {
		o.Versions = uint32(len(o.ConglomerationIDs))
	}
	if o.IsPrivate != nil && o.Versions == 0 {
		o.Versions = uint32(len(o.IsPrivate))
	}
	if o.Bitness != nil && o.Versions == 0 {
		o.Versions = uint32(len(o.Bitness))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentVersionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pdwVersions {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Versions); err != nil {
			return err
		}
	}
	// ppPartitionIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwVersions](struct))
	{
		if o.PartitionIDs != nil || o.Versions > 0 {
			_ptr_ppPartitionIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Versions)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PartitionIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PartitionIDs[i1] != nil {
						if err := o.PartitionIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PartitionIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PartitionIDs, _ptr_ppPartitionIDs); err != nil {
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
	// ppConglomerationIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwVersions](struct))
	{
		if o.ConglomerationIDs != nil || o.Versions > 0 {
			_ptr_ppConglomerationIDs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Versions)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ConglomerationIDs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ConglomerationIDs[i1] != nil {
						if err := o.ConglomerationIDs[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ConglomerationIDs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConglomerationIDs, _ptr_ppConglomerationIDs); err != nil {
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
	// ppIsPrivate {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BOOL}[dim:0,size_is=pdwVersions](int32))
	{
		if o.IsPrivate != nil || o.Versions > 0 {
			_ptr_ppIsPrivate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Versions)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.IsPrivate {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if !o.IsPrivate[i1] {
						if err := w.WriteData(int32(0)); err != nil {
							return err
						}
					} else {
						if err := w.WriteData(int32(1)); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.IsPrivate); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(false); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.IsPrivate, _ptr_ppIsPrivate); err != nil {
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
	// ppBitness {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pdwVersions](int32))
	{
		if o.Bitness != nil || o.Versions > 0 {
			_ptr_ppBitness := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Versions)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Bitness {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Bitness[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Bitness); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(int32(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Bitness, _ptr_ppBitness); err != nil {
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

func (o *xxx_GetComponentVersionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pdwVersions {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Versions); err != nil {
			return err
		}
	}
	// ppPartitionIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwVersions](struct))
	{
		_ptr_ppPartitionIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PartitionIDs", sizeInfo[0])
			}
			o.PartitionIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.PartitionIDs {
				i1 := i1
				if o.PartitionIDs[i1] == nil {
					o.PartitionIDs[i1] = &dtyp.GUID{}
				}
				if err := o.PartitionIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppPartitionIDs := func(ptr interface{}) { o.PartitionIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.PartitionIDs, _s_ppPartitionIDs, _ptr_ppPartitionIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppConglomerationIDs {out} (1:{pointer=ref}*(2)*(1))(2:{alias=GUID}[dim:0,size_is=pdwVersions](struct))
	{
		_ptr_ppConglomerationIDs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ConglomerationIDs", sizeInfo[0])
			}
			o.ConglomerationIDs = make([]*dtyp.GUID, sizeInfo[0])
			for i1 := range o.ConglomerationIDs {
				i1 := i1
				if o.ConglomerationIDs[i1] == nil {
					o.ConglomerationIDs[i1] = &dtyp.GUID{}
				}
				if err := o.ConglomerationIDs[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppConglomerationIDs := func(ptr interface{}) { o.ConglomerationIDs = *ptr.(*[]*dtyp.GUID) }
		if err := w.ReadPointer(&o.ConglomerationIDs, _s_ppConglomerationIDs, _ptr_ppConglomerationIDs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppIsPrivate {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BOOL}[dim:0,size_is=pdwVersions](int32))
	{
		_ptr_ppIsPrivate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.IsPrivate", sizeInfo[0])
			}
			o.IsPrivate = make([]bool, sizeInfo[0])
			for i1 := range o.IsPrivate {
				i1 := i1
				var _bIsPrivate_i1 int32
				if err := w.ReadData(&_bIsPrivate_i1); err != nil {
					return err
				}
				o.IsPrivate[i1] = _bIsPrivate_i1 != 0
			}
			return nil
		})
		_s_ppIsPrivate := func(ptr interface{}) { o.IsPrivate = *ptr.(*[]bool) }
		if err := w.ReadPointer(&o.IsPrivate, _s_ppIsPrivate, _ptr_ppIsPrivate); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppBitness {out} (1:{pointer=ref}*(2)*(1))(2:{alias=LONG}[dim:0,size_is=pdwVersions](int32))
	{
		_ptr_ppBitness := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Bitness", sizeInfo[0])
			}
			o.Bitness = make([]int32, sizeInfo[0])
			for i1 := range o.Bitness {
				i1 := i1
				if err := w.ReadData(&o.Bitness[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppBitness := func(ptr interface{}) { o.Bitness = *ptr.(*[]int32) }
		if err := w.ReadPointer(&o.Bitness, _s_ppBitness, _ptr_ppBitness); err != nil {
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

// GetComponentVersionsRequest structure represents the GetComponentVersions operation request
type GetComponentVersionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszClsidOrProgId: A string containing either the Curly Braced GUID String Syntax
	// ([MS-DTYP] section 2.3.4.3) representation of a CLSID or the ProgID property of a
	// component.
	ClassIDOrProgID string `idl:"name:pwszClsidOrProgId" json:"class_id_or_prog_id"`
}

func (o *GetComponentVersionsRequest) xxx_ToOp(ctx context.Context) *xxx_GetComponentVersionsOperation {
	if o == nil {
		return &xxx_GetComponentVersionsOperation{}
	}
	return &xxx_GetComponentVersionsOperation{
		This:            o.This,
		ClassIDOrProgID: o.ClassIDOrProgID,
	}
}

func (o *GetComponentVersionsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetComponentVersionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClassIDOrProgID = op.ClassIDOrProgID
}
func (o *GetComponentVersionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetComponentVersionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentVersionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetComponentVersionsResponse structure represents the GetComponentVersions operation response
type GetComponentVersionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pdwVersions: A pointer to a variable that, upon successful completion, MUST be set
	// to the number of component full configurations that exist for the component.
	Versions uint32 `idl:"name:pdwVersions" json:"versions"`
	// ppPartitionIDs: A pointer to a variable that, upon successful completion, MUST be
	// set to an array of partition identifiers of the partitions in which the component
	// full configurations reside.
	PartitionIDs []*dtyp.GUID `idl:"name:ppPartitionIDs;size_is:(, pdwVersions)" json:"partition_ids"`
	// ppConglomerationIDs: A pointer to a variable that, upon successful completion, MUST
	// be set to an array of conglomeration identifiers of the conglomerations containing
	// the component full configurations, in the same order as ppPartitionIDs.
	ConglomerationIDs []*dtyp.GUID `idl:"name:ppConglomerationIDs;size_is:(, pdwVersions)" json:"conglomeration_ids"`
	// ppIsPrivate: A pointer to a variable that, upon successful completion, MUST be set
	// to an array of the values of the IsPrivate (see section 3.1.1.3.1) property of the
	// component full configurations, in the same order as ppPartitionIDs.
	IsPrivate []bool `idl:"name:ppIsPrivate;size_is:(, pdwVersions)" json:"is_private"`
	// ppBitness: A pointer to a variable that, upon successful completion, MUST be set
	// to an array of the values of the ConfigurationBitness (see section 3.1.1.3.1) property
	// of the component full configurations, in the same order as ppPartitionIDs.
	Bitness []int32 `idl:"name:ppBitness;size_is:(, pdwVersions)" json:"bitness"`
	// Return: The GetComponentVersions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetComponentVersionsResponse) xxx_ToOp(ctx context.Context) *xxx_GetComponentVersionsOperation {
	if o == nil {
		return &xxx_GetComponentVersionsOperation{}
	}
	return &xxx_GetComponentVersionsOperation{
		That:              o.That,
		Versions:          o.Versions,
		PartitionIDs:      o.PartitionIDs,
		ConglomerationIDs: o.ConglomerationIDs,
		IsPrivate:         o.IsPrivate,
		Bitness:           o.Bitness,
		Return:            o.Return,
	}
}

func (o *GetComponentVersionsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetComponentVersionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Versions = op.Versions
	o.PartitionIDs = op.PartitionIDs
	o.ConglomerationIDs = op.ConglomerationIDs
	o.IsPrivate = op.IsPrivate
	o.Bitness = op.Bitness
	o.Return = op.Return
}
func (o *GetComponentVersionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetComponentVersionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentVersionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
