package icatalogtablewrite

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
	GoPackage = "dcom/coma"
)

var (
	// ICatalogTableWrite interface identifier 0e3d6631-b46b-11d1-9d2d-006008b0e5ca
	CatalogTableWriteIID = &dcom.IID{Data1: 0x0e3d6631, Data2: 0xb46b, Data3: 0x11d1, Data4: []byte{0x9d, 0x2d, 0x00, 0x60, 0x08, 0xb0, 0xe5, 0xca}}
	// Syntax UUID
	CatalogTableWriteSyntaxUUID = &uuid.UUID{TimeLow: 0xe3d6631, TimeMid: 0xb46b, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	// Syntax ID
	CatalogTableWriteSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CatalogTableWriteSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICatalogTableWrite interface.
type CatalogTableWriteClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// This method is called by a client to write entries to a catalog table.
	//
	// Return Values:  This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically, with the exception of E_DETAILEDERRORS (0x80110802).
	WriteTable(context.Context, *WriteTableRequest, ...dcerpc.CallOption) (*WriteTableResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CatalogTableWriteClient
}

type xxx_DefaultCatalogTableWriteClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCatalogTableWriteClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultCatalogTableWriteClient) WriteTable(ctx context.Context, in *WriteTableRequest, opts ...dcerpc.CallOption) (*WriteTableResponse, error) {
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
	out := &WriteTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCatalogTableWriteClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCatalogTableWriteClient) IPID(ctx context.Context, ipid *dcom.IPID) CatalogTableWriteClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCatalogTableWriteClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewCatalogTableWriteClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CatalogTableWriteClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CatalogTableWriteSyntaxV0_0))...)
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
	return &xxx_DefaultCatalogTableWriteClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_WriteTableOperation structure represents the WriteTable operation
type xxx_WriteTableOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	CatalogID                 *dtyp.GUID     `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	TableID                   *dtyp.GUID     `idl:"name:pTableIdentifier" json:"table_id"`
	TableFlags                uint32         `idl:"name:tableFlags" json:"table_flags"`
	QueryCellArray            []byte         `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	QueryCellArrayLength      uint32         `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	QueryComparison           []byte         `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	QueryComparisonLength     uint32         `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	QueryFormat               uint32         `idl:"name:eQueryFormat" json:"query_format"`
	TableDataFixedWrite       []byte         `idl:"name:pTableDataFixedWrite;size_is:(cbTableDataFixedWrite)" json:"table_data_fixed_write"`
	TableDataFixedWriteLength uint32         `idl:"name:cbTableDataFixedWrite" json:"table_data_fixed_write_length"`
	TableDataVariable         []byte         `idl:"name:pTableDataVariable;size_is:(cbTableDataVariable)" json:"table_data_variable"`
	TableDataVariableLength   uint32         `idl:"name:cbTableDataVariable" json:"table_data_variable_length"`
	_                         []byte         `idl:"name:pReserved1;size_is:(cbReserved1)"`
	_                         uint32         `idl:"name:cbReserved1"`
	_                         []byte         `idl:"name:pReserved2;size_is:(cbReserved2)"`
	_                         uint32         `idl:"name:cbReserved2"`
	_                         []byte         `idl:"name:pReserved3;size_is:(cbReserved3)"`
	_                         uint32         `idl:"name:cbReserved3"`
	TableDetailedErrors       []byte         `idl:"name:ppTableDetailedErrors;size_is:(, pcbTableDetailedErrors)" json:"table_detailed_errors"`
	TableDetailedErrorsLength uint32         `idl:"name:pcbTableDetailedErrors" json:"table_detailed_errors_length"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WriteTableOperation) OpNum() int { return 3 }

func (o *xxx_WriteTableOperation) OpName() string { return "/ICatalogTableWrite/v0/WriteTable" }

func (o *xxx_WriteTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.QueryCellArray != nil && o.QueryCellArrayLength == 0 {
		o.QueryCellArrayLength = uint32(len(o.QueryCellArray))
	}
	if o.QueryComparison != nil && o.QueryComparisonLength == 0 {
		o.QueryComparisonLength = uint32(len(o.QueryComparison))
	}
	if o.TableDataFixedWrite != nil && o.TableDataFixedWriteLength == 0 {
		o.TableDataFixedWriteLength = uint32(len(o.TableDataFixedWrite))
	}
	if o.TableDataVariable != nil && o.TableDataVariableLength == 0 {
		o.TableDataVariableLength = uint32(len(o.TableDataVariable))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pCatalogIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CatalogID != nil {
			if err := o.CatalogID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pTableIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.TableID != nil {
			if err := o.TableID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// tableFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TableFlags); err != nil {
			return err
		}
	}
	// pQueryCellArray {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryCellArray](char))
	{
		if o.QueryCellArray != nil || o.QueryCellArrayLength > 0 {
			_ptr_pQueryCellArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryCellArrayLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryCellArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.QueryCellArray[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.QueryCellArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryCellArray, _ptr_pQueryCellArray); err != nil {
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
	// cbQueryCellArray {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.QueryCellArrayLength); err != nil {
			return err
		}
	}
	// pQueryComparison {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryComparison](char))
	{
		if o.QueryComparison != nil || o.QueryComparisonLength > 0 {
			_ptr_pQueryComparison := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.QueryComparisonLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.QueryComparison {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.QueryComparison[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.QueryComparison); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueryComparison, _ptr_pQueryComparison); err != nil {
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
	// cbQueryComparison {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.QueryComparisonLength); err != nil {
			return err
		}
	}
	// eQueryFormat {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.QueryFormat); err != nil {
			return err
		}
	}
	// pTableDataFixedWrite {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTableDataFixedWrite](char))
	{
		dimSize1 := uint64(o.TableDataFixedWriteLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.TableDataFixedWrite {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.TableDataFixedWrite[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.TableDataFixedWrite); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbTableDataFixedWrite {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDataFixedWriteLength); err != nil {
			return err
		}
	}
	// pTableDataVariable {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTableDataVariable](char))
	{
		dimSize1 := uint64(o.TableDataVariableLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.TableDataVariable {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.TableDataVariable[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.TableDataVariable); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbTableDataVariable {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDataVariableLength); err != nil {
			return err
		}
	}
	// pReserved1 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved1](char))
	{
		// reserved pReserved1
		dimSize1 := uint64(0)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbReserved1 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved1
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pReserved2 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved2](char))
	{
		// reserved pReserved2
		dimSize1 := uint64(0)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbReserved2 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved2
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pReserved3 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved3](char))
	{
		// reserved pReserved3
		dimSize1 := uint64(0)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbReserved3 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved3
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pCatalogIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.CatalogID == nil {
			o.CatalogID = &dtyp.GUID{}
		}
		if err := o.CatalogID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pTableIdentifier {in} (1:{pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.TableID == nil {
			o.TableID = &dtyp.GUID{}
		}
		if err := o.TableID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// tableFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TableFlags); err != nil {
			return err
		}
	}
	// pQueryCellArray {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryCellArray](char))
	{
		_ptr_pQueryCellArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryCellArray", sizeInfo[0])
			}
			o.QueryCellArray = make([]byte, sizeInfo[0])
			for i1 := range o.QueryCellArray {
				i1 := i1
				if err := w.ReadData(&o.QueryCellArray[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pQueryCellArray := func(ptr interface{}) { o.QueryCellArray = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.QueryCellArray, _s_pQueryCellArray, _ptr_pQueryCellArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbQueryCellArray {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.QueryCellArrayLength); err != nil {
			return err
		}
	}
	// pQueryComparison {in} (1:{pointer=unique}*(1)[dim:0,size_is=cbQueryComparison](char))
	{
		_ptr_pQueryComparison := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.QueryComparison", sizeInfo[0])
			}
			o.QueryComparison = make([]byte, sizeInfo[0])
			for i1 := range o.QueryComparison {
				i1 := i1
				if err := w.ReadData(&o.QueryComparison[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pQueryComparison := func(ptr interface{}) { o.QueryComparison = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.QueryComparison, _s_pQueryComparison, _ptr_pQueryComparison); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbQueryComparison {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.QueryComparisonLength); err != nil {
			return err
		}
	}
	// eQueryFormat {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.QueryFormat); err != nil {
			return err
		}
	}
	// pTableDataFixedWrite {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTableDataFixedWrite](char))
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
			return fmt.Errorf("buffer overflow for size %d of array o.TableDataFixedWrite", sizeInfo[0])
		}
		o.TableDataFixedWrite = make([]byte, sizeInfo[0])
		for i1 := range o.TableDataFixedWrite {
			i1 := i1
			if err := w.ReadData(&o.TableDataFixedWrite[i1]); err != nil {
				return err
			}
		}
	}
	// cbTableDataFixedWrite {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDataFixedWriteLength); err != nil {
			return err
		}
	}
	// pTableDataVariable {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbTableDataVariable](char))
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
			return fmt.Errorf("buffer overflow for size %d of array o.TableDataVariable", sizeInfo[0])
		}
		o.TableDataVariable = make([]byte, sizeInfo[0])
		for i1 := range o.TableDataVariable {
			i1 := i1
			if err := w.ReadData(&o.TableDataVariable[i1]); err != nil {
				return err
			}
		}
	}
	// cbTableDataVariable {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDataVariableLength); err != nil {
			return err
		}
	}
	// pReserved1 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved1](char))
	{
		// reserved pReserved1
		var _pReserved1 []byte
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _pReserved1", sizeInfo[0])
		}
		_pReserved1 = make([]byte, sizeInfo[0])
		for i1 := range _pReserved1 {
			i1 := i1
			if err := w.ReadData(&_pReserved1[i1]); err != nil {
				return err
			}
		}
	}
	// cbReserved1 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved1
		var _cbReserved1 uint32
		if err := w.ReadData(&_cbReserved1); err != nil {
			return err
		}
	}
	// pReserved2 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved2](char))
	{
		// reserved pReserved2
		var _pReserved2 []byte
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _pReserved2", sizeInfo[0])
		}
		_pReserved2 = make([]byte, sizeInfo[0])
		for i1 := range _pReserved2 {
			i1 := i1
			if err := w.ReadData(&_pReserved2[i1]); err != nil {
				return err
			}
		}
	}
	// cbReserved2 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved2
		var _cbReserved2 uint32
		if err := w.ReadData(&_cbReserved2); err != nil {
			return err
		}
	}
	// pReserved3 {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbReserved3](char))
	{
		// reserved pReserved3
		var _pReserved3 []byte
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _pReserved3", sizeInfo[0])
		}
		_pReserved3 = make([]byte, sizeInfo[0])
		for i1 := range _pReserved3 {
			i1 := i1
			if err := w.ReadData(&_pReserved3[i1]); err != nil {
				return err
			}
		}
	}
	// cbReserved3 {in} (1:{alias=ULONG}(uint32))
	{
		// reserved cbReserved3
		var _cbReserved3 uint32
		if err := w.ReadData(&_cbReserved3); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.TableDetailedErrors != nil && o.TableDetailedErrorsLength == 0 {
		o.TableDetailedErrorsLength = uint32(len(o.TableDetailedErrors))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WriteTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppTableDetailedErrors {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDetailedErrors](char))
	{
		if o.TableDetailedErrors != nil || o.TableDetailedErrorsLength > 0 {
			_ptr_ppTableDetailedErrors := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TableDetailedErrorsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.TableDetailedErrors {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.TableDetailedErrors[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.TableDetailedErrors); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TableDetailedErrors, _ptr_ppTableDetailedErrors); err != nil {
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
	// pcbTableDetailedErrors {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TableDetailedErrorsLength); err != nil {
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

func (o *xxx_WriteTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppTableDetailedErrors {out} (1:{pointer=ref}*(2)*(1)[dim:0,size_is=pcbTableDetailedErrors](char))
	{
		_ptr_ppTableDetailedErrors := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.TableDetailedErrors", sizeInfo[0])
			}
			o.TableDetailedErrors = make([]byte, sizeInfo[0])
			for i1 := range o.TableDetailedErrors {
				i1 := i1
				if err := w.ReadData(&o.TableDetailedErrors[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppTableDetailedErrors := func(ptr interface{}) { o.TableDetailedErrors = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.TableDetailedErrors, _s_ppTableDetailedErrors, _ptr_ppTableDetailedErrors); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbTableDetailedErrors {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TableDetailedErrorsLength); err != nil {
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

// WriteTableRequest structure represents the WriteTable operation request
type WriteTableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pCatalogIdentifier:  The catalog identifier of the COMA catalog. MUST be {6E38D3C4-C2A7-11D1-8DEC-00C04FC2E0C7}.
	CatalogID *dtyp.GUID `idl:"name:pCatalogIdentifier" json:"catalog_id"`
	// pTableIdentifier:  The table identifier for one of the tables defined in section
	// 3.1.1.3 for the negotiated catalog version.
	TableID *dtyp.GUID `idl:"name:pTableIdentifier" json:"table_id"`
	// tableFlags:  An fTableFlags (section 2.2.1.1) value supported (see section 3.1.1.2.3)
	// by the table identified by pTableIdentifier.
	TableFlags uint32 `idl:"name:tableFlags" json:"table_flags"`
	// pQueryCellArray:  A QueryCellArray structure, marshaled in the negotiated format
	// as specified in section 2.2.1.5, for a supported query (see section 3.1.1.2.2) on
	// the table identified by pTableIdentifier.
	QueryCellArray []byte `idl:"name:pQueryCellArray;size_is:(cbQueryCellArray);pointer:unique" json:"query_cell_array"`
	// cbQueryCellArray:  The size in bytes of pQueryCellArray.
	QueryCellArrayLength uint32 `idl:"name:cbQueryCellArray" json:"query_cell_array_length"`
	// pQueryComparison:  A QueryComparisonData (section 2.2.1.6) structure for a supported
	// query (see section 3.1.1.2) on the table identified by pTableIdentifier.
	QueryComparison []byte `idl:"name:pQueryComparison;size_is:(cbQueryComparison);pointer:unique" json:"query_comparison"`
	// cbQueryComparison:  The size in bytes of pQueryComparison.
	QueryComparisonLength uint32 `idl:"name:cbQueryComparison" json:"query_comparison_length"`
	// eQueryFormat:  MUST be set to eQUERYFORMAT_1 (0x00000001).
	QueryFormat uint32 `idl:"name:eQueryFormat" json:"query_format"`
	// pTableDataFixedWrite:  A TableDataFixedWrite structure, marshaled as specified in
	// section 2.2.1.13.
	TableDataFixedWrite []byte `idl:"name:pTableDataFixedWrite;size_is:(cbTableDataFixedWrite)" json:"table_data_fixed_write"`
	// cbTableDataFixedWrite:  The length in bytes of the TableDataFixedWrite structure
	// passed in pTableDataFixedWrite.
	TableDataFixedWriteLength uint32 `idl:"name:cbTableDataFixedWrite" json:"table_data_fixed_write_length"`
	// pTableDataVariable:  A TableDataVariable structure, marshaled as specified in section
	// 2.2.1.15.
	TableDataVariable []byte `idl:"name:pTableDataVariable;size_is:(cbTableDataVariable)" json:"table_data_variable"`
	// cbTableDataVariable:  The length in bytes of the TableDataVariable structure passed
	// in pTableDataVariable.
	TableDataVariableLength uint32 `idl:"name:cbTableDataVariable" json:"table_data_variable_length"`
}

func (o *WriteTableRequest) xxx_ToOp(ctx context.Context) *xxx_WriteTableOperation {
	if o == nil {
		return &xxx_WriteTableOperation{}
	}
	return &xxx_WriteTableOperation{
		This:                      o.This,
		CatalogID:                 o.CatalogID,
		TableID:                   o.TableID,
		TableFlags:                o.TableFlags,
		QueryCellArray:            o.QueryCellArray,
		QueryCellArrayLength:      o.QueryCellArrayLength,
		QueryComparison:           o.QueryComparison,
		QueryComparisonLength:     o.QueryComparisonLength,
		QueryFormat:               o.QueryFormat,
		TableDataFixedWrite:       o.TableDataFixedWrite,
		TableDataFixedWriteLength: o.TableDataFixedWriteLength,
		TableDataVariable:         o.TableDataVariable,
		TableDataVariableLength:   o.TableDataVariableLength,
	}
}

func (o *WriteTableRequest) xxx_FromOp(ctx context.Context, op *xxx_WriteTableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CatalogID = op.CatalogID
	o.TableID = op.TableID
	o.TableFlags = op.TableFlags
	o.QueryCellArray = op.QueryCellArray
	o.QueryCellArrayLength = op.QueryCellArrayLength
	o.QueryComparison = op.QueryComparison
	o.QueryComparisonLength = op.QueryComparisonLength
	o.QueryFormat = op.QueryFormat
	o.TableDataFixedWrite = op.TableDataFixedWrite
	o.TableDataFixedWriteLength = op.TableDataFixedWriteLength
	o.TableDataVariable = op.TableDataVariable
	o.TableDataVariableLength = op.TableDataVariableLength
}
func (o *WriteTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WriteTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WriteTableResponse structure represents the WriteTable operation response
type WriteTableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTableDetailedErrors:  A pointer to a variable that, upon successful completion,
	// MUST be set to NULL, and that, upon partial failure, MAY<297> be set to a TableDetailedErrorArray
	// structure, marshaled as specified in section 2.2.1.17.
	TableDetailedErrors []byte `idl:"name:ppTableDetailedErrors;size_is:(, pcbTableDetailedErrors)" json:"table_detailed_errors"`
	// pcbTableDetailedErrors:  A pointer to a variable that, upon completion, MUST be
	// set to the length in bytes of the TableDetailedErrorArray structure returned in ppTableDetailedErrors
	// if such a structure was returned and MUST be set to zero otherwise.
	TableDetailedErrorsLength uint32 `idl:"name:pcbTableDetailedErrors" json:"table_detailed_errors_length"`
	// Return: The WriteTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WriteTableResponse) xxx_ToOp(ctx context.Context) *xxx_WriteTableOperation {
	if o == nil {
		return &xxx_WriteTableOperation{}
	}
	return &xxx_WriteTableOperation{
		That:                      o.That,
		TableDetailedErrors:       o.TableDetailedErrors,
		TableDetailedErrorsLength: o.TableDetailedErrorsLength,
		Return:                    o.Return,
	}
}

func (o *WriteTableResponse) xxx_FromOp(ctx context.Context, op *xxx_WriteTableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TableDetailedErrors = op.TableDetailedErrors
	o.TableDetailedErrorsLength = op.TableDetailedErrorsLength
	o.Return = op.Return
}
func (o *WriteTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WriteTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WriteTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
