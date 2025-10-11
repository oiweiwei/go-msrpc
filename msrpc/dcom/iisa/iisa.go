// The iisa package implements the IISA client protocol.
//
// # Introduction
//
// This document specifies the Internet Information Services (IIS) Application Host
// COM Protocol. This protocol is a client-to-server protocol that enables remote read/write
// access to server data. The server data can be used to define administration, configuration,
// and operational parameters to an application server service, which can be a web server.
//
// # Overview
//
// This protocol is intended to provide read/write access to administrative configuration
// data that is located on a remote server computer. The administrative configuration
// data is implementation-specific for each server.
package iisa

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
	GoPackage = "dcom/iisa"
)

// AppHostPropertySchema structure represents IAppHostPropertySchema RPC structure.
//
// The IAppHostPropertySchema interface provides methods that access the schema and
// constraints for the corresponding IAppHostProperty object.
//
// The IAppHostPropertySchema interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+-------------+
//	|                     |             |
//	|       METHOD        | DESCRIPTION |
//	|                     |             |
//	+---------------------+-------------+
//	+---------------------+-------------+
//	| Name                | Opnum: 3    |
//	+---------------------+-------------+
//	| Type                | Opnum: 4    |
//	+---------------------+-------------+
//	| DefaultValue        | Opnum: 5    |
//	+---------------------+-------------+
//	| IsRequired          | Opnum: 6    |
//	+---------------------+-------------+
//	| IsUniqueKey         | Opnum: 7    |
//	+---------------------+-------------+
//	| IsCombinedKey       | Opnum: 8    |
//	+---------------------+-------------+
//	| IsExpanded          | Opnum: 9    |
//	+---------------------+-------------+
//	| ValidationType      | Opnum: 10   |
//	+---------------------+-------------+
//	| ValidationParameter | Opnum: 11   |
//	+---------------------+-------------+
//	| GetMetadata         | Opnum: 12   |
//	+---------------------+-------------+
//	| IsCaseSensitive     | Opnum: 13   |
//	+---------------------+-------------+
//	| PossibleValues      | Opnum: 14   |
//	+---------------------+-------------+
//	| DoesAllowInfinite   | Opnum: 15   |
//	+---------------------+-------------+
//	| IsEncrypted         | Opnum: 16   |
//	+---------------------+-------------+
//	| TimeSpanFormat      | Opnum: 17   |
//	+---------------------+-------------+
type AppHostPropertySchema dcom.InterfacePointer

func (o *AppHostPropertySchema) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostPropertySchema) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostPropertySchema) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostPropertySchema) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostPropertySchema) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostPropertySchemaCollection structure represents IAppHostPropertySchemaCollection RPC structure.
//
// The IAppHostPropertySchemaCollection interface provides methods that access a collection
// of IAppHostPropertySchema objects.
//
// The IAppHostPropertySchemaCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostPropertySchemaCollection dcom.InterfacePointer

func (o *AppHostPropertySchemaCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostPropertySchemaCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostPropertySchemaCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostPropertySchemaCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostPropertySchemaCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostSectionDefinition structure represents IAppHostSectionDefinition RPC structure.
//
// The IAppHostSectionDefinition interface provides methods that access a declaration
// of the IAppHostElement object that is supported by the administration system. A declaration
// is distinct from the existence of an IAppHostElement in the administration system.
//
// The IAppHostSectionDefinition interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+--------------------+
//	|                     |                    |
//	|       METHOD        |    DESCRIPTION     |
//	|                     |                    |
//	+---------------------+--------------------+
//	+---------------------+--------------------+
//	| Name                | Opnum: 3           |
//	+---------------------+--------------------+
//	| Type                | "getter" Opnum: 4  |
//	+---------------------+--------------------+
//	| Type                | "setter" Opnum: 5  |
//	+---------------------+--------------------+
//	| OverrideModeDefault | "getter" Opnum: 6  |
//	+---------------------+--------------------+
//	| OverrideModeDefault | "setter" Opnum: 7  |
//	+---------------------+--------------------+
//	| AllowDefinition     | "getter" Opnum: 8  |
//	+---------------------+--------------------+
//	| AllowDefinition     | "setter" Opnum: 9  |
//	+---------------------+--------------------+
//	| AllowLocation       | "getter" Opnum: 10 |
//	+---------------------+--------------------+
//	| AllowLocation       | "setter" Opnum: 11 |
//	+---------------------+--------------------+
type AppHostSectionDefinition dcom.InterfacePointer

func (o *AppHostSectionDefinition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostSectionDefinition) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostSectionDefinition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostSectionDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostSectionDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConfigLocation structure represents IAppHostConfigLocation RPC structure.
//
// The IAppHostConfigLocation interface provides methods that access the IAppHostElement
// objects that are contained in a specific hierarchy subpath in a specified IAppHostConfigFile.
// IAppHostConfigFile maps to a specific hierarchy path in the administration system,
// as specified in section 3.1.4.6.
//
// Each IAppHostConfigFile can optionally contain subpaths within it. Each subpath is
// represented by an IAppHostConfigLocation object. The object contains a collection
// of IAppHostElement objects, with the guarantee that each IAppHostConfigLocation contains
// at most one IAppHostElement object of the same name (in other words, the IAppHostElement
// object name is a key into the collection).
//
// The IAppHostConfigLocation interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+-------------+
//	|                     |             |
//	|       METHOD        | DESCRIPTION |
//	|                     |             |
//	+---------------------+-------------+
//	+---------------------+-------------+
//	| Path                | Opnum: 3    |
//	+---------------------+-------------+
//	| Count               | Opnum: 4    |
//	+---------------------+-------------+
//	| Item                | Opnum: 5    |
//	+---------------------+-------------+
//	| AddConfigSection    | Opnum: 6    |
//	+---------------------+-------------+
//	| DeleteConfigSection | Opnum: 7    |
//	+---------------------+-------------+
type AppHostConfigLocation dcom.InterfacePointer

func (o *AppHostConfigLocation) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConfigLocation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConfigLocation) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConfigLocation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConfigLocation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostProperty structure represents IAppHostProperty RPC structure.
//
// The IAppHostProperty interface provides methods that access properties that can be
// contained under an IAppHostElement object.
//
// IAppHostElement objects can contain zero or more IAppHostProperty objects. These
// IAppHostProperty objects typically represent a single setting or attribute set on
// an IAppHostElement but can be multiple settings or sets.
//
// The IAppHostProperty interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+-------------+-------------------+
//	|             |                   |
//	|   METHOD    |    DESCRIPTION    |
//	|             |                   |
//	+-------------+-------------------+
//	+-------------+-------------------+
//	| Name        | Opnum: 3          |
//	+-------------+-------------------+
//	| Value       | "getter" Opnum: 4 |
//	+-------------+-------------------+
//	| Value       | "setter" Opnum: 5 |
//	+-------------+-------------------+
//	| Clear       | Opnum: 6          |
//	+-------------+-------------------+
//	| StringValue | Opnum: 7          |
//	+-------------+-------------------+
//	| Exception   | Opnum: 8          |
//	+-------------+-------------------+
//	| GetMetadata | Opnum: 9          |
//	+-------------+-------------------+
//	| SetMetadata | Opnum: 10         |
//	+-------------+-------------------+
//	| Schema      | Opnum: 11         |
//	+-------------+-------------------+
type AppHostProperty dcom.InterfacePointer

func (o *AppHostProperty) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostProperty) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostProperty) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostMethod structure represents IAppHostMethod RPC structure.
//
// The IAppHostMethod interface provides methods that access a custom method that is
// optionally supported on a specific IAppHostElement object.
//
// An IAppHostElement object provides a means for an administration system to support
// custom-defined methods that can be executed against a specific IAppHostElement object.
// The methods are executed on the server side and the implementation of these custom
// methods is not exposed to the client.
//
// The IAppHostMethod interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+----------------+-------------+
//	|                |             |
//	|     METHOD     | DESCRIPTION |
//	|                |             |
//	+----------------+-------------+
//	+----------------+-------------+
//	| Name           | Opnum: 3    |
//	+----------------+-------------+
//	| Schema         | Opnum: 4    |
//	+----------------+-------------+
//	| CreateInstance | Opnum: 5    |
//	+----------------+-------------+
type AppHostMethod dcom.InterfacePointer

func (o *AppHostMethod) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *AppHostMethod) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostMethod) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostMethod) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostMethod) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostMethodCollection structure represents IAppHostMethodCollection RPC structure.
//
// The IAppHostMethodCollection interface provides methods that access a collection
// of supported IAppHostMethod objects.
//
// The IAppHostMethodCollection interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostMethodCollection dcom.InterfacePointer

func (o *AppHostMethodCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostMethodCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostMethodCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostMethodCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostMethodCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostElementCollection structure represents IAppHostElementCollection RPC structure.
//
// The IAppHostElementCollection interface provides methods that access a collection
// of "collection IAppHostElements".
//
// "Collection IAppHostElements" are a special class of child IAppHostElement objects
// where all objects typically share the same name but contain different IAppHostProperty
// objects. This IAppHostElementCollection is a collection of these special objects.
//
// The IAppHostElementCollection interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+------------------+----------------------------------------------------------------------------------+
//	|                  |                                                                                  |
//	|      METHOD      |                                   DESCRIPTION                                    |
//	|                  |                                                                                  |
//	+------------------+----------------------------------------------------------------------------------+
//	+------------------+----------------------------------------------------------------------------------+
//	| Count            | Returns the count of elements. Opnum: 3                                          |
//	+------------------+----------------------------------------------------------------------------------+
//	| Item             | Returns the value that is associated with the element at the specified index.    |
//	|                  | Opnum: 4                                                                         |
//	+------------------+----------------------------------------------------------------------------------+
//	| AddElement       | Adds an element to the collection. Opnum: 5                                      |
//	+------------------+----------------------------------------------------------------------------------+
//	| DeleteElement    | Deletes an element from the collection. Opnum: 6                                 |
//	+------------------+----------------------------------------------------------------------------------+
//	| Clear            | Clears an element from the collection. Opnum: 7                                  |
//	+------------------+----------------------------------------------------------------------------------+
//	| CreateNewElement | Creates a new element in the collection. Opnum: 8                                |
//	+------------------+----------------------------------------------------------------------------------+
//	| Schema           | Returns the Document Object Model (DOM) description for the collection. Opnum: 9 |
//	+------------------+----------------------------------------------------------------------------------+
type AppHostElementCollection dcom.InterfacePointer

func (o *AppHostElementCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostElementCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostElementCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostElementCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostElementCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostSectionDefinitionCollection structure represents IAppHostSectionDefinitionCollection RPC structure.
//
// The IAppHostSectionDefinitionCollection interface provides methods that access a
// collection of IAppHostSectionDefinition objects.
//
// The IAppHostSectionDefinitionCollection interface inherits opnums 0–2 from the
// IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------+-------------+
//	|               |             |
//	|    METHOD     | DESCRIPTION |
//	|               |             |
//	+---------------+-------------+
//	+---------------+-------------+
//	| Count         | Opnum: 3    |
//	+---------------+-------------+
//	| Item          | Opnum: 4    |
//	+---------------+-------------+
//	| AddSection    | Opnum: 5    |
//	+---------------+-------------+
//	| DeleteSection | Opnum: 6    |
//	+---------------+-------------+
type AppHostSectionDefinitionCollection dcom.InterfacePointer

func (o *AppHostSectionDefinitionCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostSectionDefinitionCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostSectionDefinitionCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostSectionDefinitionCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostSectionDefinitionCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostElement structure represents IAppHostElement RPC structure.
//
// The IAppHostElement interface provides methods to the base administration system
// object of this whole project.
//
// The administration system is a container of IAppHostElement objects. IAppHostElement
// objects are primarily an n-way-tree object, with each IAppHostElement object primarily
// containing:
//
// * Zero or more IAppHostProperty ( 564b4c1f-13a0-40c4-974b-1cbbd41764b3 ) objects.
// Each has its own unique fixed name.
//
// * Zero or more child IAppHostElement objects (hence the n -way-tree description).
// Each has its own unique fixed name.
//
// * An optional collection of zero or more special child IAppHostElement objects that
// are called collection objects. Each typically has the same fixed name.
//
// * Zero or more IAppHostMethod ( 9154a520-ce4f-43e1-95c2-63590cfea206 ) objects, which
// are additional extension methods that can be executed on the specified IAppHostElement
// object.
//
// The IAppHostElement interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+-------------------+-------------+
//	|                   |             |
//	|      METHOD       | DESCRIPTION |
//	|                   |             |
//	+-------------------+-------------+
//	+-------------------+-------------+
//	| Name              | Opnum: 3    |
//	+-------------------+-------------+
//	| Collection        | Opnum: 4    |
//	+-------------------+-------------+
//	| Properties        | Opnum: 5    |
//	+-------------------+-------------+
//	| ChildElements     | Opnum: 6    |
//	+-------------------+-------------+
//	| GetMetadata       | Opnum: 7    |
//	+-------------------+-------------+
//	| SetMetadata       | Opnum: 8    |
//	+-------------------+-------------+
//	| Schema            | Opnum: 9    |
//	+-------------------+-------------+
//	| GetElementByName  | Opnum: 10   |
//	+-------------------+-------------+
//	| GetPropertyByName | Opnum: 11   |
//	+-------------------+-------------+
//	| Clear             | Opnum: 12   |
//	+-------------------+-------------+
//	| Methods           | Opnum: 13   |
//	+-------------------+-------------+
type AppHostElement dcom.InterfacePointer

func (o *AppHostElement) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostElement) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostElement) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostElement) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostElement) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConstantValueCollection structure represents IAppHostConstantValueCollection RPC structure.
//
// The IAppHostConstantValueCollection interface provides methods that access a collection
// of constant values.
//
// The IAppHostConstantValueCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostConstantValueCollection dcom.InterfacePointer

func (o *AppHostConstantValueCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConstantValueCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConstantValueCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConstantValueCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConstantValueCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConfigException structure represents IAppHostConfigException RPC structure.
//
// Methods that receive an error when they access the administration system can retrieve
// more specific information about why the error occurred by using the IAppHostConfigException
// interface.
//
// The IAppHostConfigException interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------+-------------+
//	|               |             |
//	|    METHOD     | DESCRIPTION |
//	|               |             |
//	+---------------+-------------+
//	+---------------+-------------+
//	| LineNumber    | Opnum: 3    |
//	+---------------+-------------+
//	| FileName      | Opnum: 4    |
//	+---------------+-------------+
//	| ConfigPath    | Opnum: 5    |
//	+---------------+-------------+
//	| ErrorLine     | Opnum: 6    |
//	+---------------+-------------+
//	| PreErrorLine  | Opnum: 7    |
//	+---------------+-------------+
//	| PostErrorLine | Opnum: 8    |
//	+---------------+-------------+
//	| ErrorString   | Opnum: 9    |
//	+---------------+-------------+
type AppHostConfigException dcom.InterfacePointer

func (o *AppHostConfigException) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConfigException) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConfigException) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConfigException) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConfigException) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostMappingExtension structure represents IAppHostMappingExtension RPC structure.
//
// The IAppHostMappingExtension interface provides methods that access the path hierarchy
// mapping system of the administration system.
//
// The administration system implements a path hierarchy system that maps paths to potential
// IAppHostElement containers. Some details of the path hierarchy are exposed to the
// user in this IAppHostMappingExtension.
//
// The IAppHostMappingExtension interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+--------------------------+-------------+
//	|                          |             |
//	|          METHOD          | DESCRIPTION |
//	|                          |             |
//	+--------------------------+-------------+
//	+--------------------------+-------------+
//	| GetSiteNameFromSiteId    | Opnum: 3    |
//	+--------------------------+-------------+
//	| GetSiteIdFromSiteName    | Opnum: 4    |
//	+--------------------------+-------------+
//	| GetSiteElementFromSiteId | Opnum: 5    |
//	+--------------------------+-------------+
//	| MapPath                  | Opnum: 6    |
//	+--------------------------+-------------+
type AppHostMappingExtension dcom.InterfacePointer

func (o *AppHostMappingExtension) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostMappingExtension) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostMappingExtension) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostMappingExtension) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostMappingExtension) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostPathMapper structure represents IAppHostPathMapper RPC structure.
//
// The IAppHostPathMapper interface provides methods that are called by the server implementation
// when the server informs the client about hierarchy mapping decisions.
//
// To receive incoming remote calls for this interface, the client MUST implement a
// UUID (e7927575-5cc3-403b-822e-328a6b904bee). It MUST then specify an object that
// implements this interface to the IAppHostAdminManager::SetMetadata() method with
// a bstrMetadataName of "pathMapper".
//
// As an administration system maps hierarchy paths to physical paths on the server,
// it optionally calls this client-supplied object that implements the IAppHostPathMapper
// interface. The implementer of this interface receives details of all mappings and
// can change the results of each mapping if required.
//
// The IAppHostPathMapper interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------+-------------+
//	|         |             |
//	| METHOD  | DESCRIPTION |
//	|         |             |
//	+---------+-------------+
//	+---------+-------------+
//	| MapPath | Opnum: 3    |
//	+---------+-------------+
type AppHostPathMapper dcom.InterfacePointer

func (o *AppHostPathMapper) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostPathMapper) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostPathMapper) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostPathMapper) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostPathMapper) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostPropertyException structure represents IAppHostPropertyException RPC structure.
//
// The IAppHostPropertyException interface provides methods that access the exception
// information that the administration system encountered when processing the corresponding
// IAppHostProperty. The administration system can indicate errors as encountered by
// filling in this exception. This behavior is defined by the IAppHostAdminManager metadata
// "ignoreInvalidAttributes".
//
// The IAppHostPropertyException interface inherits opnums 0–9 from the IAppHostConfigException
// interface.
//
// Methods in RPC Opnum Order
//
//	+-----------------------------+-------------+
//	|                             |             |
//	|           METHOD            | DESCRIPTION |
//	|                             |             |
//	+-----------------------------+-------------+
//	+-----------------------------+-------------+
//	| InvalidValue                | Opnum: 10   |
//	+-----------------------------+-------------+
//	| ValidationFailureReason     | Opnum: 11   |
//	+-----------------------------+-------------+
//	| ValidationFailureParameters | Opnum: 12   |
//	+-----------------------------+-------------+
type AppHostPropertyException dcom.InterfacePointer

func (o *AppHostPropertyException) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostPropertyException) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostPropertyException) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostPropertyException) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostPropertyException) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConstantValue structure represents IAppHostConstantValue RPC structure.
//
// The IAppHostConstantValue interface provides methods that access the string names
// of a specific constant and its corresponding value.
//
// The IAppHostConstantValue interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Name   | Opnum: 3    |
//	+--------+-------------+
//	| Value  | Opnum: 4    |
//	+--------+-------------+
type AppHostConstantValue dcom.InterfacePointer

func (o *AppHostConstantValue) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConstantValue) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConstantValue) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConstantValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConstantValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConfigLocationCollection structure represents IAppHostConfigLocationCollection RPC structure.
//
// The IAppHostConfigLocationCollection interface provides methods that access the collection
// of subpaths that are available in the specified IAppHostConfigFile.
//
// The IAppHostConfigLocationCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+----------------+-------------+
//	|                |             |
//	|     METHOD     | DESCRIPTION |
//	|                |             |
//	+----------------+-------------+
//	+----------------+-------------+
//	| Count          | Opnum: 3    |
//	+----------------+-------------+
//	| Item           | Opnum: 4    |
//	+----------------+-------------+
//	| AddLocation    | Opnum: 5    |
//	+----------------+-------------+
//	| DeleteLocation | Opnum: 6    |
//	+----------------+-------------+
type AppHostConfigLocationCollection dcom.InterfacePointer

func (o *AppHostConfigLocationCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConfigLocationCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConfigLocationCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConfigLocationCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConfigLocationCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostElementSchemaCollection structure represents IAppHostElementSchemaCollection RPC structure.
//
// The IAppHostElementSchemaCollection interface provides methods that access a collection
// of schema and constraints for child IAppHostElement objects that are supported by
// the corresponding IAppHostElement object.
//
// The IAppHostElementSchemaCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostElementSchemaCollection dcom.InterfacePointer

func (o *AppHostElementSchemaCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostElementSchemaCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostElementSchemaCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostElementSchemaCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostElementSchemaCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostMethodInstance structure represents IAppHostMethodInstance RPC structure.
//
// The IAppHostMethodInstance interface provides methods that access a specific invocation
// instance of the corresponding IAppHostMethod. The caller sets parameters and then
// executes the instance of the method.
//
// The IAppHostMethodInstance interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+-------------+-------------+
//	|             |             |
//	|   METHOD    | DESCRIPTION |
//	|             |             |
//	+-------------+-------------+
//	+-------------+-------------+
//	| Input       | Opnum: 3    |
//	+-------------+-------------+
//	| Output      | Opnum: 4    |
//	+-------------+-------------+
//	| Execute     | Opnum: 5    |
//	+-------------+-------------+
//	| GetMetadata | Opnum: 6    |
//	+-------------+-------------+
//	| SetMetadata | Opnum: 7    |
//	+-------------+-------------+
type AppHostMethodInstance dcom.InterfacePointer

func (o *AppHostMethodInstance) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostMethodInstance) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostMethodInstance) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostMethodInstance) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostMethodInstance) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConfigManager structure represents IAppHostConfigManager RPC structure.
//
// The IAppHostConfigManager interface provides methods that allow access to the available
// hierarchy paths and the IAppHostElement objects that are defined within.
//
// The IAppHostConfigManager interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+-------------+
//	|                     |             |
//	|       METHOD        | DESCRIPTION |
//	|                     |             |
//	+---------------------+-------------+
//	+---------------------+-------------+
//	| GetConfigFile       | Opnum: 3    |
//	+---------------------+-------------+
//	| GetUniqueConfigPath | Opnum: 4    |
//	+---------------------+-------------+
type AppHostConfigManager dcom.InterfacePointer

func (o *AppHostConfigManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConfigManager) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConfigManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConfigManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConfigManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostElementSchema structure represents IAppHostElementSchema RPC structure.
//
// The IAppHostElementSchema interface provides methods that access the schema and constraints
// of a specific IAppHostElement object.
//
// The IAppHostElementSchema interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+----------------------------------+--------------------------------------------------------------------+
//	|                                  |                                                                    |
//	|              METHOD              |                            DESCRIPTION                             |
//	|                                  |                                                                    |
//	+----------------------------------+--------------------------------------------------------------------+
//	+----------------------------------+--------------------------------------------------------------------+
//	| Name                             | Returns the name of the element entry. Opnum: 3                    |
//	+----------------------------------+--------------------------------------------------------------------+
//	| DoesAllowUnschematizedProperties | Determines if the section allows unrecognized attributes. Opnum: 4 |
//	+----------------------------------+--------------------------------------------------------------------+
//	| GetMetadata                      | Used to get a metadata property. Opnum: 5                          |
//	+----------------------------------+--------------------------------------------------------------------+
//	| CollectionSchema                 | Opnum: 6                                                           |
//	+----------------------------------+--------------------------------------------------------------------+
//	| ChildElementSchemas              | Opnum: 7                                                           |
//	+----------------------------------+--------------------------------------------------------------------+
//	| PropertySchemas                  | Opnum: 8                                                           |
//	+----------------------------------+--------------------------------------------------------------------+
//	| IsCollectionDefault              | Opnum: 9                                                           |
//	+----------------------------------+--------------------------------------------------------------------+
type AppHostElementSchema dcom.InterfacePointer

func (o *AppHostElementSchema) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostElementSchema) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostElementSchema) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostElementSchema) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostElementSchema) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostAdminManager structure represents IAppHostAdminManager RPC structure.
//
// The IAppHostAdminManager interface represents a read-only interface to an administration
// system that is implemented on the server.
//
// The administration system consists primarily of a set of administration objects of
// varying complexity, which are accessed through the IAppHostElement interface (section
// 3.1.4.12) and exist at one or more paths that are exposed by the server. The administration
// system allows access to individual IAppHostElement objects that are available at
// specific paths and also provides access to merged IAppHostElement objects that consist
// of the merged contents of individual IAppHostElement objects.
//
// Secondarily, the administration system provides access to, and the setting of, specific
// system behaviors that are available for administration objects, which are represented
// by the term "metadata". Metadata allows a caller to modify and inspect the behavior
// of the administration system.
//
// The IAppHostAdminManager interface is used by clients to access (for read-only purposes)
// the contents of the administration system without altering the contents of the system.
// A tool that seeks to display the administration objects that are contained in the
// system is an example of a consumer of the IAppHostAdminManager interface.
//
// The IAppHostAdminManager interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+-----------------+-------------+
//	|                 |             |
//	|     METHOD      | DESCRIPTION |
//	|                 |             |
//	+-----------------+-------------+
//	+-----------------+-------------+
//	| GetAdminSection | Opnum: 3    |
//	+-----------------+-------------+
//	| GetMetadata     | Opnum: 4    |
//	+-----------------+-------------+
//	| SetMetadata     | Opnum: 5    |
//	+-----------------+-------------+
//	| ConfigManager   | Opnum: 6    |
//	+-----------------+-------------+
type AppHostAdminManager dcom.InterfacePointer

func (o *AppHostAdminManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostAdminManager) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostAdminManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostAdminManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostAdminManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostConfigFile structure represents IAppHostConfigFile RPC structure.
//
// The IAppHostConfigFile interface provides methods for direct access to a container
// of IAppHostElement objects in the administration system. The administration system
// implements a path hierarchy that allows for IAppHostElement objects to be defined
// at multiple places. This path hierarchy is conceptually split into two levels. One
// is the level that is represented by IAppHostConfigFile.
//
// Each IAppHostConfigFile maps to a specific path (for example, "MACHINE/WEBROOT").
// Within each IAppHostConfigFile, individual IAppHostElement objects exist, which are
// the base objects of the administration system. Optionally, the IAppHostConfigFile
// can also support subpaths within it.
//
// The IAppHostConfigFile interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+----------------------+-------------+
//	|                      |             |
//	|        METHOD        | DESCRIPTION |
//	|                      |             |
//	+----------------------+-------------+
//	+----------------------+-------------+
//	| ConfigPath           | Opnum: 3    |
//	+----------------------+-------------+
//	| FilePath             | Opnum: 4    |
//	+----------------------+-------------+
//	| Locations            | Opnum: 5    |
//	+----------------------+-------------+
//	| GetAdminSection      | Opnum: 6    |
//	+----------------------+-------------+
//	| GetMetadata          | Opnum: 7    |
//	+----------------------+-------------+
//	| SetMetadata          | Opnum: 8    |
//	+----------------------+-------------+
//	| ClearInvalidSections | Opnum: 9    |
//	+----------------------+-------------+
//	| RootSectionGroup     | Opnum: 10   |
//	+----------------------+-------------+
type AppHostConfigFile dcom.InterfacePointer

func (o *AppHostConfigFile) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostConfigFile) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostConfigFile) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostConfigFile) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostConfigFile) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostPropertyCollection structure represents IAppHostPropertyCollection RPC structure.
//
// The IAppHostPropertyCollection interface provides methods that access a collection
// of IAppHostProperty objects that are supported by a corresponding IAppHostElement.
//
// The IAppHostPropertyCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostPropertyCollection dcom.InterfacePointer

func (o *AppHostPropertyCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostPropertyCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostPropertyCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostPropertyCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostPropertyCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostChangeHandler structure represents IAppHostChangeHandler RPC structure.
//
// The IAppHostChangeHandler describes an interface that clients can implement and that
// is called when the administration system has detected a change in a part of its path
// hierarchy.
//
// To receive incoming remote calls for this interface, the client MUST implement this
// interface (09829352-87c2-418d-8d79-4133969a489d). It MUST then specify an object
// that implements this interface to the IAppHostAdminManager::SetMetadata() method
// by using a bstrMetadataName of "changeHandler".
//
// The server then calls this object when the administration system detects a change.
// The administration system is free to determine the supported time period during which
// the changes are detected and conveyed through this interface. The time period is
// either the lifetime of the administration system or a shorter time period.
//
// The IAppHostChangeHandler interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+------------------+-------------+
//	|                  |             |
//	|      METHOD      | DESCRIPTION |
//	|                  |             |
//	+------------------+-------------+
//	+------------------+-------------+
//	| OnSectionChanges | Opnum: 3    |
//	+------------------+-------------+
type AppHostChangeHandler dcom.InterfacePointer

func (o *AppHostChangeHandler) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostChangeHandler) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostChangeHandler) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostChangeHandler) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostChangeHandler) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostMethodSchema structure represents IAppHostMethodSchema RPC structure.
//
// The IAppHostMethodSchema interface provides methods that access the schema and constraints
// of the corresponding IAppHostMethod.
//
// The IAppHostMethodSchema interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+--------------+-------------+
//	|              |             |
//	|    METHOD    | DESCRIPTION |
//	|              |             |
//	+--------------+-------------+
//	+--------------+-------------+
//	| Name         | Opnum: 3    |
//	+--------------+-------------+
//	| InputSchema  | Opnum: 4    |
//	+--------------+-------------+
//	| OutputSchema | Opnum: 5    |
//	+--------------+-------------+
//	| GetMetadata  | Opnum: 6    |
//	+--------------+-------------+
type AppHostMethodSchema dcom.InterfacePointer

func (o *AppHostMethodSchema) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostMethodSchema) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostMethodSchema) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostMethodSchema) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostMethodSchema) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostCollectionSchema structure represents IAppHostCollectionSchema RPC structure.
//
// The IAppHostCollectionSchema provides methods that describe the schema and constraints
// that apply to a specific IAppHostElementCollection from which this object was retrieved.
//
// The IAppHostCollectionSchema interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------------+-------------+
//	|                     |             |
//	|       METHOD        | DESCRIPTION |
//	|                     |             |
//	+---------------------+-------------+
//	+---------------------+-------------+
//	| AddElementNames     | Opnum: 3    |
//	+---------------------+-------------+
//	| GetAddElementSchema | Opnum: 4    |
//	+---------------------+-------------+
//	| RemoveElementSchema | Opnum: 5    |
//	+---------------------+-------------+
//	| ClearElementSchema  | Opnum: 6    |
//	+---------------------+-------------+
//	| IsMergeAppend       | Opnum: 7    |
//	+---------------------+-------------+
//	| GetMetadata         | Opnum: 8    |
//	+---------------------+-------------+
//	| DoesAllowDuplicates | Opnum: 9    |
//	+---------------------+-------------+
type AppHostCollectionSchema dcom.InterfacePointer

func (o *AppHostCollectionSchema) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostCollectionSchema) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostCollectionSchema) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostCollectionSchema) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostCollectionSchema) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostWritableAdminManager structure represents IAppHostWritableAdminManager RPC structure.
//
// The IAppHostWritableAdminManager interface provides methods that access a writable
// version of an administration system. It extends the IAppHostAdminManager, which is
// a read-only interface. The IAppHostWritableAdminManager adds methods to allow writing
// to the administration system, the most important of which is the CommitChanges method,
// which instructs the administration system to persist any in-memory changes that it
// has accumulated.
//
// The IAppHostWritableAdminManager interface inherits opnums 0–6 from the IAppHostAdminManager
// interface, as defined in this protocol specification and the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+---------------+-------------------+
//	|               |                   |
//	|    METHOD     |    DESCRIPTION    |
//	|               |                   |
//	+---------------+-------------------+
//	+---------------+-------------------+
//	| CommitChanges | Opnum: 7          |
//	+---------------+-------------------+
//	| CommitPath    | "getter" Opnum: 8 |
//	+---------------+-------------------+
//	| CommitPath    | "setter" Opnum: 9 |
//	+---------------+-------------------+
type AppHostWritableAdminManager dcom.InterfacePointer

func (o *AppHostWritableAdminManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostWritableAdminManager) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostWritableAdminManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostWritableAdminManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostWritableAdminManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostSectionGroup structure represents IAppHostSectionGroup RPC structure.
//
// The IAppHostSectionGroup interface provides methods that access a group of section
// definitions that have a common prefix name.
//
// The IAppHostSectionGroup interface inherits opnums 0–2 from the IUnknown interface.
//
// Methods in RPC Opnum Order
//
//	+--------------------+--------------------+
//	|                    |                    |
//	|       METHOD       |    DESCRIPTION     |
//	|                    |                    |
//	+--------------------+--------------------+
//	+--------------------+--------------------+
//	| Count              | Opnum: 3           |
//	+--------------------+--------------------+
//	| Item               | Opnum: 4           |
//	+--------------------+--------------------+
//	| Sections           | Opnum: 5           |
//	+--------------------+--------------------+
//	| AddSectionGroup    | Opnum: 6           |
//	+--------------------+--------------------+
//	| DeleteSectionGroup | Opnum: 7           |
//	+--------------------+--------------------+
//	| Name               | Opnum: 8           |
//	+--------------------+--------------------+
//	| Type               | "getter" Opnum: 9  |
//	+--------------------+--------------------+
//	| Type               | "setter" Opnum: 10 |
//	+--------------------+--------------------+
type AppHostSectionGroup dcom.InterfacePointer

func (o *AppHostSectionGroup) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostSectionGroup) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostSectionGroup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostSectionGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostSectionGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// AppHostChildElementCollection structure represents IAppHostChildElementCollection RPC structure.
//
// The IAppHostChildElementCollection interface provides methods that allow access to
// any fixed child IAppHostElement objects that are contained by the specific IAppHostElement
// object that provides this interface through IAppHostElement::ChildElements().
//
// An IAppHostElement can contain any number of fixed, uniquely named child IAppHostElement
// objects. This IAppHostChildElementCollection provides access to these child elements.
//
// The IAppHostChildElementCollection interface inherits opnums 0–2 from the IUnknown
// interface.
//
// Methods in RPC Opnum Order
//
//	+--------+-------------+
//	|        |             |
//	| METHOD | DESCRIPTION |
//	|        |             |
//	+--------+-------------+
//	+--------+-------------+
//	| Count  | Opnum: 3    |
//	+--------+-------------+
//	| Item   | Opnum: 4    |
//	+--------+-------------+
type AppHostChildElementCollection dcom.InterfacePointer

func (o *AppHostChildElementCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AppHostChildElementCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AppHostChildElementCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AppHostChildElementCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *AppHostChildElementCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}
