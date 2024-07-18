// The comev package implements the COMEV client protocol.
//
// # Introduction
//
// This document specifies the behavior of the Component Object Model Plus (COM+) Event
// System Protocol.
//
// The COM+ Event System Protocol is a protocol that exposes DCOM interfaces for storing
// and managing configuration data for publishers of events and their respective subscribers
// on remote computers. This protocol also specifies how to get specific information
// about a publisher and its subscribers.
package comev

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
	GoPackage = "dcom/comev"
)

// EventSystemInitialize structure represents IEventSystemInitialize RPC structure.
//
// The IEventSystemInitialize interface is used to initialize the server implementing
// this protocol. As this is a DCOM interface, Opnum 0 to Opnum 2 are IUnknown methods,
// which MUST be implemented by means of IRemUnknown, as specified in [MS-DCOM] section
// 3.1.1.5.6. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_EventSystem (see section 1.9) using the UUID
// {a0e8f27a-888c-11d1-b763-00c04fb926af} for this interface .
//
// The interface contains the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+------------------------+----------------------------------+
//	|                        |                                  |
//	|         METHOD         |           DESCRIPTION            |
//	|                        |                                  |
//	+------------------------+----------------------------------+
//	+------------------------+----------------------------------+
//	| SetCOMCatalogBehaviour | Initializes the server. Opnum: 3 |
//	+------------------------+----------------------------------+
type EventSystemInitialize dcom.InterfacePointer

func (o *EventSystemInitialize) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EventSystemInitialize) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSystemInitialize) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSystemInitialize) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSystemInitialize) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventClass structure represents IEventClass RPC structure.
//
// The IEventClass interface provides methods that are used by the client to manipulate
// an event class on the server. The interface inherits Opnums 0 to 6 from IDispatch
// as specified in [MS-OAUT] section 3.1.4. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the class ID CLSID_EventClass (see section 1.9) using the
// UUID {fb2b72a0-7a68-11d1-88f9-0080c7d771bf} for this interface.
//
// The interface includes the following methods beyond those in IDispatch.
//
// Methods in RPC Opnum Order
//
//	+----------------------+--------------------------------------------------------------------+
//	|                      |                                                                    |
//	|        METHOD        |                            DESCRIPTION                             |
//	|                      |                                                                    |
//	+----------------------+--------------------------------------------------------------------+
//	+----------------------+--------------------------------------------------------------------+
//	| EventClassID         | Gets the EventClassID property for the event class. Opnum: 7       |
//	+----------------------+--------------------------------------------------------------------+
//	| EventClassID         | Sets the EventClassID property for the event class. Opnum: 8       |
//	+----------------------+--------------------------------------------------------------------+
//	| EventClassName       | Gets the EventClassName property of the event class. Opnum: 9      |
//	+----------------------+--------------------------------------------------------------------+
//	| EventClassName       | Sets the EventClassName property of the event class. Opnum: 10     |
//	+----------------------+--------------------------------------------------------------------+
//	| OwnerSID             | Gets the OwnerSID property of the event class. Opnum: 11           |
//	+----------------------+--------------------------------------------------------------------+
//	| OwnerSID             | Sets the OwnerSID property of the event class. Opnum: 12           |
//	+----------------------+--------------------------------------------------------------------+
//	| FiringInterfaceID    | Gets the FiringInterfaceID property for the event class. Opnum: 13 |
//	+----------------------+--------------------------------------------------------------------+
//	| FiringInterfaceID    | Sets the FiringInterfaceID property for the event class. Opnum: 14 |
//	+----------------------+--------------------------------------------------------------------+
//	| Description          | Gets the Description property for the event class. Opnum: 15       |
//	+----------------------+--------------------------------------------------------------------+
//	| Description          | Sets the Description property for the event class. Opnum: 16       |
//	+----------------------+--------------------------------------------------------------------+
//	| Opnum17NotUsedOnWire | Reserved for local use. Opnum: 17                                  |
//	+----------------------+--------------------------------------------------------------------+
//	| Opnum18NotUsedOnWire | Reserved for local use. Opnum: 18                                  |
//	+----------------------+--------------------------------------------------------------------+
//	| TypeLib              | Gets the Typelib property of the event class. Opnum: 19            |
//	+----------------------+--------------------------------------------------------------------+
//	| TypeLib              | Sets the Typelib property of the event class. Opnum: 20            |
//	+----------------------+--------------------------------------------------------------------+
//
// In the preceding table, the term "Reserved for local use" means that the client MUST
// NOT send the opnum, and the server behavior is undefined<12> because it does not
// affect interoperability.
type EventClass dcom.InterfacePointer

func (o *EventClass) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EventClass) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventClass) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventClass) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventClass) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EnumEventObject structure represents IEnumEventObject RPC structure.
//
// The IEnumEventObject interface provides methods that are used to enumerate a collection
// of event classes or subscriptions. The version for this interface is 0.0.
//
// A client gets this interface by means of the get_NewEnum (Opnum 9) (section 3.1.4.6.3)
// method of the IEventObjectCollection. As this is a DCOM interface, opnums 0 through
// 2 are IUnknown methods, which MUST be implemented by means of IRemUnknown, as specified
// in [MS-DCOM] section 3.1.1.5.6. The DCOM object implementing this interface MUST
// use the UUID {F4A07D63-2E25-11D1-9964-00C04FBBB345}.
//
// This interface includes the following methods beyond those of IUnknown.
//
// Methods in RPC Opnum Order
//
//	+--------+---------------------------------------------------------------------------------+
//	|        |                                                                                 |
//	| METHOD |                                   DESCRIPTION                                   |
//	|        |                                                                                 |
//	+--------+---------------------------------------------------------------------------------+
//	+--------+---------------------------------------------------------------------------------+
//	| Clone  | Clones the collection into another IEnumEventObject-based DCOM object. Opnum: 3 |
//	+--------+---------------------------------------------------------------------------------+
//	| Next   | Returns the next elements and iterates over them. Opnum: 4                      |
//	+--------+---------------------------------------------------------------------------------+
//	| Reset  | Resets the enumerating object back to the first element. Opnum: 5               |
//	+--------+---------------------------------------------------------------------------------+
//	| Skip   | Skips ahead in the collection. Opnum: 6                                         |
//	+--------+---------------------------------------------------------------------------------+
type EnumEventObject dcom.InterfacePointer

func (o *EnumEventObject) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EnumEventObject) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EnumEventObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumEventObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EnumEventObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventClass2 structure represents IEventClass2 RPC structure.
//
// The IEventClass2 interface provides additional methods that are used by the client
// to manipulate event class properties on the server. This interface inherits opnums
// 0 through 20 from IEventClass as specified in section 3.1.4.2. The version for this
// interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the class ID CLSID_EventClass (see section 1.9) using the
// UUID {fb2b72a1-7a68-11d1-88f9-0080c7d771bf} for this interface.
//
// This interface includes the following methods beyond those in IEventClass.
//
// Methods in RPC Opnum Order
//
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	|                                        |                                                                                  |
//	|                 METHOD                 |                                   DESCRIPTION                                    |
//	|                                        |                                                                                  |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| get_PublisherID                        | Gets the PublisherID property of the event class. Opnum: 21                      |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| put_PublisherID                        | Sets the PublisherID property of the event class. Opnum: 22                      |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| get_MultiInterfacePublisherFilterCLSID | Gets the MultiInterfacePublisherFilterCLSID property of the event class. Opnum:  |
//	|                                        |                                                                               23 |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| put_MultiInterfacePublisherFilterCLSID | Sets the MultiInterfacePublisherFilterCLSID property of the event class. Opnum:  |
//	|                                        |                                                                               24 |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| get_AllowInprocActivation              | Gets the AllowInprocActivation property of the event class. Opnum: 25            |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| put_AllowInprocActivation              | Sets the AllowInprocActivation property of the event class. Opnum: 26            |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| get_FireInParallel                     | Gets the FireInParallel property of the event class. Opnum: 27                   |
//	+----------------------------------------+----------------------------------------------------------------------------------+
//	| put_FireInParallel                     | Sets the FireInParallel property of the event class. Opnum: 28                   |
//	+----------------------------------------+----------------------------------------------------------------------------------+
type EventClass2 dcom.InterfacePointer

func (o *EventClass2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EventClass2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventClass2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventClass2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventClass2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventClass3 structure represents IEventClass3 RPC structure.
//
// The IEventClass3 interface provides additional methods that are used by the client
// to manipulate an event class on the server. This interface inherits opnums 0 through
// 28 from the IEventClass2 interface, as specified in section 3.1.4.3. The version
// for this interface is 0.0.
//
// The server SHOULD support this interface.<13> To receive incoming remote calls for
// this interface, the server MUST implement a DCOM object class with the class ID CLSID_EventClass
// (see section 1.9) using the UUID {7FB7EA43-2D76-4ea8-8CD9-3DECC270295E} for this
// interface.
//
// The interface contains the following methods beyond the ones that are defined for
// IEventClass2.
//
// Methods in RPC Opnum Order
//
//	+-----------------------------+-------------------------------------------------------------------------+
//	|                             |                                                                         |
//	|           METHOD            |                               DESCRIPTION                               |
//	|                             |                                                                         |
//	+-----------------------------+-------------------------------------------------------------------------+
//	+-----------------------------+-------------------------------------------------------------------------+
//	| get_EventClassPartitionID   | Gets the EventClassPartitionID property of the event class. Opnum: 29   |
//	+-----------------------------+-------------------------------------------------------------------------+
//	| put_EventClassPartitionID   | Sets the EventClassPartitionID property of the event class. Opnum: 30   |
//	+-----------------------------+-------------------------------------------------------------------------+
//	| get_EventClassApplicationID | Gets the EventClassApplicationID property of the event class. Opnum: 31 |
//	+-----------------------------+-------------------------------------------------------------------------+
//	| put_EventClassApplicationID | Has no effect. Opnum: 32                                                |
//	+-----------------------------+-------------------------------------------------------------------------+
type EventClass3 dcom.InterfacePointer

func (o *EventClass3) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EventClass3) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventClass3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventClass3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventClass3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventSubscription3 structure represents IEventSubscription3 RPC structure.
//
// The IEventSubscription3 interface provides methods to get or set the properties of
// a subscription.<14> This interface inherits opnums 0 through 44 from IEventSubscription2,
// as specified in section 3.1.4.8. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM [MS-DCOM] object class with the class ID CLSID_Subscription (see section 1.9)
// using the UUID {FBC1D17D-C498-43a0-81AF-423DDD530AF6} for this interface.
//
// The interface contains the following methods beyond those of IEventSubscription2.
//
// Methods in RPC Opnum Order
//
//	+-----------------------------+--------------------------------------------------------------------------+
//	|                             |                                                                          |
//	|           METHOD            |                               DESCRIPTION                                |
//	|                             |                                                                          |
//	+-----------------------------+--------------------------------------------------------------------------+
//	+-----------------------------+--------------------------------------------------------------------------+
//	| get_EventClassPartitionID   | Gets the EventClassPartitionID property of the subscription. Opnum: 45   |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| put_EventClassPartitionID   | Sets the EventClassPartitionID property of the subscription. Opnum: 46   |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| get_EventClassApplicationID | Gets the EventClassApplicationID property of the subscription. Opnum: 47 |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| put_EventClassApplicationID | Has no effect. Opnum: 48                                                 |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| get_SubscriberPartitionID   | Gets the SubscriberPartitionID property of the subscription. Opnum: 49   |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| put_SubscriberPartitionID   | Sets the SubscriberPartitionID property of the subscription. Opnum: 50   |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| get_SubscriberApplicationID | Gets the SubscriberApplicationID property of the subscription. Opnum: 51 |
//	+-----------------------------+--------------------------------------------------------------------------+
//	| put_SubscriberApplicationID | Sets the SubscriberApplicationID property of the subscription. Opnum: 52 |
//	+-----------------------------+--------------------------------------------------------------------------+
type EventSubscription3 dcom.InterfacePointer

func (o *EventSubscription3) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EventSubscription3) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSubscription3) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSubscription3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSubscription3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventSubscription2 structure represents IEventSubscription2 RPC structure.
//
// The IEventSubscription2 interface provides methods to get and set the properties
// of a subscription. This interface inherits opnums 0 through 40 from IEventSubscription
// as specified in section 3.1.4.4. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM [MS-DCOM] object class with the class ID CLSID_Subscription (see section 1.9)
// using the UUID {4A6B0E16-2E38-11D1-9965-00C04FBBB345} for this interface.
//
// The interface contains the following methods beyond the ones for IEventSubscription.
//
// Methods in RPC Opnum Order
//
//	+-----------------------+--------------------------------------------------------------------+
//	|                       |                                                                    |
//	|        METHOD         |                            DESCRIPTION                             |
//	|                       |                                                                    |
//	+-----------------------+--------------------------------------------------------------------+
//	+-----------------------+--------------------------------------------------------------------+
//	| get_FilterCriteria    | Gets the FilterCriteria property of the subscription. Opnum: 41    |
//	+-----------------------+--------------------------------------------------------------------+
//	| put_FilterCriteria    | Sets the FilterCriteria property of the subscription. Opnum: 42    |
//	+-----------------------+--------------------------------------------------------------------+
//	| get_SubscriberMoniker | Gets the SubscriberMoniker property of the subscription. Opnum: 43 |
//	+-----------------------+--------------------------------------------------------------------+
//	| put_SubscriberMoniker | Sets the SubscriberMoniker property of the subscription. Opnum: 44 |
//	+-----------------------+--------------------------------------------------------------------+
type EventSubscription2 dcom.InterfacePointer

func (o *EventSubscription2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EventSubscription2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSubscription2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSubscription2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSubscription2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventObjectCollection structure represents IEventObjectCollection RPC structure.
//
// The IEventObjectCollection interface provides methods that manage and enumerate over
// a collection of objects. The interface inherits opnums 0 through 6 from IDispatch
// as specified in [MS-OAUT] section 3.1.4. The version for this interface is 0.0.
//
// The server returns a DCOM object implementing this interface with UUID {f89ac270-d4eb-11d1-b682-00805fc79216}
// from the following methods:
//
// *
//
// Query (Opnum 7) (section 3.1.4.1.1) ( 2929ddac-3b08-4976-be20-876189da05b3 )
//
// *
//
// GetPublisherPropertyCollection (Opnum 34) (section 3.1.4.4.28) ( 9d995cc7-5efa-47bc-9429-d4c9d2d6d087
// )
//
// *
//
// GetSubscriberPropertyCollection (Opnum 38) (section 3.1.4.4.32) ( 7b3c746c-9ce3-4027-a5fa-64c0701fdaf7
// )
//
// This interface includes the following methods beyond those of IDispatch.
//
// Methods in RPC Opnum Order
//
//	+--------------+----------------------------------------------------------------------------------+
//	|              |                                                                                  |
//	|    METHOD    |                                   DESCRIPTION                                    |
//	|              |                                                                                  |
//	+--------------+----------------------------------------------------------------------------------+
//	+--------------+----------------------------------------------------------------------------------+
//	| get__NewEnum | Returns an enumerator DCOM object. Opnum: 7                                      |
//	+--------------+----------------------------------------------------------------------------------+
//	| get_Item     | Gets the item in the collection based on an ID. Opnum: 8                         |
//	+--------------+----------------------------------------------------------------------------------+
//	| get_NewEnum  | Gets the IEnumEventObject supporting the DCOM based object instance for the      |
//	|              | underlying collection. Opnum: 9                                                  |
//	+--------------+----------------------------------------------------------------------------------+
//	| get_Count    | Gets the number of items in the collection. Opnum: 10                            |
//	+--------------+----------------------------------------------------------------------------------+
//	| Add          | Adds an item to the collection. Opnum: 11                                        |
//	+--------------+----------------------------------------------------------------------------------+
//	| Remove       | Removes an item from the collection. Opnum: 12                                   |
//	+--------------+----------------------------------------------------------------------------------+
type EventObjectCollection dcom.InterfacePointer

func (o *EventObjectCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EventObjectCollection) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventObjectCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventObjectCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventObjectCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventSystem2 structure represents IEventSystem2 RPC structure.
//
// This interface is used to perform version checking and transient subscription verifications
// on the server by the client. IEventSystem2 inherits opnums 0 through 12 from the
// IEventSystem interface, as specified in section 3.1.4.1. The version for this interface
// is 0.0.
//
// The server SHOULD support this interface.<15> To receive incoming remote calls for
// this interface, the server MUST implement a DCOM object class with the CLSID CLSID_EventSystem
// (see section 1.9) using the UUID {99CC098F-A48A-4e9c-8E58-965C0AFC19D5}for this interface.
//
// The interface contains the following methods beyond those of IEventSystem.
//
// Methods in RPC Opnum Order
//
//	+----------------------------+----------------------------------------------------------------+
//	|                            |                                                                |
//	|           METHOD           |                          DESCRIPTION                           |
//	|                            |                                                                |
//	+----------------------------+----------------------------------------------------------------+
//	+----------------------------+----------------------------------------------------------------+
//	| GetVersion                 | Gets the version of the event system implementation. Opnum: 13 |
//	+----------------------------+----------------------------------------------------------------+
//	| VerifyTransientSubscribers | Verifies the state of the transient subscribers. Opnum: 14     |
//	+----------------------------+----------------------------------------------------------------+
type EventSystem2 dcom.InterfacePointer

func (o *EventSystem2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EventSystem2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSystem2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSystem2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSystem2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventSubscription structure represents IEventSubscription RPC structure.
//
// The IEventSubscription interface provides methods to get and set the properties of
// a subscription. This interface inherits opnums 0 through 6 from [MS-OAUT] IDispatch
// as specified in [MS-OAUT] section 3.1.4. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM [MS-DCOM] object class with the class ID CLSID_Subscription (see section 1.9)
// using the UUID {4A6B0E15-2E38-11D1-9965-00C04FBBB345} for this interface.
//
// The interface includes the following methods beyond those in IDispatch.
//
// Methods in RPC Opnum Order
//
//	+---------------------------------+----------------------------------------------------------------------------------+
//	|                                 |                                                                                  |
//	|             METHOD              |                                   DESCRIPTION                                    |
//	|                                 |                                                                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_SubscriptionID              | Gets the SubscriptionID property of the subscription. Opnum: 7                   |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_SubscriptionID              | Sets the SubscriptionID property of the subscription. Opnum: 8                   |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_SubscriptionName            | Gets the SubscriptionName property of the subscription. Opnum: 9                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_SubscriptionName            | Sets the SubscriptionName property of the subscription. Opnum: 10                |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_PublisherID                 | Gets the PublisherID property of the subscription. Opnum: 11                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_PublisherID                 | Sets the PublisherID property of the subscription. Opnum: 12                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_EventClassID                | Gets the EventClassID property of the subscription. Opnum: 13                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_EventClassID                | Sets the EventClassID property of the subscription. Opnum: 14                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_MethodName                  | Gets the MethodName property of the subscription. Opnum: 15                      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_MethodName                  | Sets the MethodName property of the subscription. Opnum: 16                      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_SubscriberCLSID             | Gets the SubscriberCLSID property of the subscription. Opnum: 17                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_SubscriberCLSID             | Sets the SubscriberCLSID property of the subscription. Opnum: 18                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_SubscriberInterface         | Gets the SubscriberInterface property of the subscription. Opnum: 19             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_SubscriberInterface         | Sets the SubscriberInterface property of the subscription. Opnum: 20             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_PerUser                     | Gets the PerUser property of the subscription. Opnum: 21                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_PerUser                     | Sets the PerUser property of the subscription. Opnum: 22                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_OwnerSID                    | Gets the OwnerSID property of the subscription. Opnum: 23                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_OwnerSID                    | Sets the OwnerSID property of the subscription. Opnum: 24                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_Enabled                     | Gets the Enabled property of the subscription. Opnum: 25                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_Enabled                     | Sets the Enabled property of the subscription. Opnum: 26                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_Description                 | Gets the Description property of the subscription. Opnum: 27                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_Description                 | Sets the Description property of the subscription. Opnum: 28                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_MachineName                 | Gets the MachineName property of the subscription. Opnum: 29                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_MachineName                 | Sets the MachineName property of the subscription. Opnum: 30                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| GetPublisherProperty            | Gets the application-specific publisher property for the subscription. Opnum: 31 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| PutPublisherProperty            | Sets the application-specific publisher property for the subscription. Opnum: 32 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| RemovePublisherProperty         | Removes an application-specific publisher property for the subscription. Opnum:  |
//	|                                 |                                                                               33 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| GetPublisherPropertyCollection  | Gets an application-specific publisher properties collection for the             |
//	|                                 | subscription. Opnum: 34                                                          |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| GetSubscriberProperty           | Gets an application-specific subscription property for the subscription. Opnum:  |
//	|                                 |                                                                               35 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| PutSubscriberProperty           | Sets an application-specific subscription property for the subscription. Opnum:  |
//	|                                 |                                                                               36 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| RemoveSubscriberProperty        | Removes an application-specific subscription property for the subscription.      |
//	|                                 | Opnum: 37                                                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| GetSubscriberPropertyCollection | Gets the application-specific subscription properties for the subscription as a  |
//	|                                 | collection. Opnum: 38                                                            |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| get_InterfaceID                 | Gets the InterfaceID property for the subscription. Opnum: 39                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| put_InterfaceID                 | Sets the InterfaceID property for the subscription. Opnum: 40                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
type EventSubscription dcom.InterfacePointer

func (o *EventSubscription) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EventSubscription) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSubscription) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSubscription) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSubscription) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventSystem structure represents IEventSystem RPC structure.
//
// The IEventSystem interface provides methods to create, query, delete, and update
// event classes and subscriptions. The interface inherits opnums 0 through 6 from IDispatch
// as specified in [MS-OAUT] section 3.1.4. The version for this interface is 0.0.
//
// To receive incoming remote calls for this interface, the server MUST implement a
// DCOM object class with the CLSID CLSID_EventSystem (see section 1.9) using the UUID
// {4E14FB9F-2E22-11D1-9964-00C04FBBB345} for this interface.
//
// The interface includes the following methods beyond those in IDispatch.
//
// Methods in RPC Opnum Order
//
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	|                                   |                                                                                  |
//	|              METHOD               |                                   DESCRIPTION                                    |
//	|                                   |                                                                                  |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| Query                             | Queries for a collection of event classes or subscriptions based on a query      |
//	|                                   | string. Opnum: 7                                                                 |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| Store                             | Stores an event class or subscription. Opnum: 8                                  |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| Remove                            | Removes a collection of event classes or subscriptions based on a query. Opnum:  |
//	|                                   |                                                                                9 |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| get_EventObjectChangeEventClassID | Returns the CLSID for the event class that notifies when a subscription or an    |
//	|                                   | event class has changed. Opnum: 10                                               |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| QueryS                            | Queries for a single event class or subscription based on a query string. Opnum: |
//	|                                   |                                                                               11 |
//	+-----------------------------------+----------------------------------------------------------------------------------+
//	| RemoveS                           | Removes a single event class or subscription based on a query. Opnum: 12         |
//	+-----------------------------------+----------------------------------------------------------------------------------+
type EventSystem dcom.InterfacePointer

func (o *EventSystem) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EventSystem) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EventSystem) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EventSystem) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EventSystem) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
