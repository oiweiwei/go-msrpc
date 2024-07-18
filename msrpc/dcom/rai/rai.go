// The rai package implements the RAI client protocol.
//
// # Introduction
//
// The Remote Assistance Initiation Protocol is a set of Distributed Component Object
// Model (DCOM) interfaces, as specified in [MS-DCOM], for initiating a Remote Assistance
// connection to another computer in a domain. The Remote Assistance Initiation Protocol
// allows an authorized expert to start Remote Assistance on a remote novice computer
// to retrieve data that is required to make a Remote Assistance connection from the
// expert computer to the novice computer.
//
// # Overview
//
// The Remote Assistance Initiation Protocol provides a set of DCOM interfaces that
// enable an expert to retrieve the Remote Assistance connection-specific data from
// the remote novice computer. This Remote Assistance connection-specific data is subsequently
// used to initiate a Remote Assistance connection as explained in the Remote Assistance
// Initiation Protocol.
//
// The expert needs to have the IP address or fully qualified domain name (FQDN) of
// the novice computer to use this protocol.
//
// The expert is the DCOM client and the novice is the DCOM server.
//
// Before the expert's DCOM call is executed on the novice computer, DCOM performs a
// check to verify that the expert is on the list of authorized Remote Assistance helpers
// on the novice computer.<1>
package rai

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
	GoPackage = "dcom/rai"
)

// DispatchIDPCHBase represents the DISPID_PCH_BASE RPC constant
var DispatchIDPCHBase = 134283264

// DispatchIDPCHBaseColumn represents the DISPID_PCH_BASE_COL RPC constant
var DispatchIDPCHBaseColumn = 134283264

// DispatchIDPCHHelpServiceBase represents the DISPID_PCH_HELPSVC_BASE RPC constant
var DispatchIDPCHHelpServiceBase = 134217728

// DispatchIDPCHHelpServiceBaseService represents the DISPID_PCH_HELPSVC_BASE_SVC RPC constant
var DispatchIDPCHHelpServiceBaseService = 134217728

// DispatchIDPCHServiceRemoteConnectionParams represents the DISPID_PCH_SVC__REMOTECONNECTIONPARMS RPC constant
var DispatchIDPCHServiceRemoteConnectionParams = 134217792

// DispatchIDPCHServiceRemoteUserSessionInfo represents the DISPID_PCH_SVC__REMOTEUSERSESSIONINFO RPC constant
var DispatchIDPCHServiceRemoteUserSessionInfo = 134217793

// DispatchIDPCHColumnCount represents the DISPID_PCH_COL__COUNT RPC constant
var DispatchIDPCHColumnCount = 134283264

// DispatchIDSAFBase represents the DISPID_SAF_BASE RPC constant
var DispatchIDSAFBase = 134348800

// DispatchIDSAFBaseRcd represents the DISPID_SAF_BASE_RCD RPC constant
var DispatchIDSAFBaseRcd = 134351616

// DispatchIDSAFBaseUser represents the DISPID_SAF_BASE_USER RPC constant
var DispatchIDSAFBaseUser = 134351872

// DispatchIDSAFBaseSess represents the DISPID_SAF_BASE_SESS RPC constant
var DispatchIDSAFBaseSess = 134352128

// DispatchIDSAFUserDomainName represents the DISPID_SAF_USER__DOMAINNAME RPC constant
var DispatchIDSAFUserDomainName = 134351888

// DispatchIDSAFUserUserName represents the DISPID_SAF_USER__USERNAME RPC constant
var DispatchIDSAFUserUserName = 134351889

// DispatchIDSAFSessSessionID represents the DISPID_SAF_SESS__SESSIONID RPC constant
var DispatchIDSAFSessSessionID = 134352144

// DispatchIDSAFSessSessionState represents the DISPID_SAF_SESS__SESSIONSTATE RPC constant
var DispatchIDSAFSessSessionState = 134352145

// DispatchIDSAFSessDomainName represents the DISPID_SAF_SESS__DOMAINNAME RPC constant
var DispatchIDSAFSessDomainName = 134352146

// DispatchIDSAFSessUserName represents the DISPID_SAF_SESS__USERNAME RPC constant
var DispatchIDSAFSessUserName = 134352147

// SessionStateEnum type represents SessionStateEnum RPC enumeration.
//
// The SessionStateEnum enumeration defines the states of a terminal services session.
type SessionStateEnum uint16

var (
	// pchActive:  The user is logged on and active.
	SessionStateEnumActive SessionStateEnum = 0
	// pchConnected:  The server is connected to the client.
	SessionStateEnumConnected SessionStateEnum = 1
	// pchConnectQuery:  The server is in the process of connecting to the client.
	SessionStateEnumConnectQuery SessionStateEnum = 2
	// pchShadow:  The session is shadowing another session.
	SessionStateEnumShadow SessionStateEnum = 3
	// pchDisconnected:  The client has disconnected from the session.
	SessionStateEnumDisconnected SessionStateEnum = 4
	// pchIdle:  The session is waiting for a client to connect.
	SessionStateEnumIdle SessionStateEnum = 5
	// pchListen:  The session is listening for a request for a new connection. No user
	// is logged on to a listener session. A listener session cannot be reset, shadowed,
	// or changed to a regular client session.
	SessionStateEnumListen SessionStateEnum = 6
	// pchReset:  The session is being reset.
	SessionStateEnumReset SessionStateEnum = 7
	// pchDown:  The session is down due to an error.
	SessionStateEnumDown SessionStateEnum = 8
	// pchInit:  The session is initializing.
	SessionStateEnumInit SessionStateEnum = 9
	// pchStateInvalid:  The session is in an unknown state.
	SessionStateEnumInvalid SessionStateEnum = 10
)

func (o SessionStateEnum) String() string {
	switch o {
	case SessionStateEnumActive:
		return "SessionStateEnumActive"
	case SessionStateEnumConnected:
		return "SessionStateEnumConnected"
	case SessionStateEnumConnectQuery:
		return "SessionStateEnumConnectQuery"
	case SessionStateEnumShadow:
		return "SessionStateEnumShadow"
	case SessionStateEnumDisconnected:
		return "SessionStateEnumDisconnected"
	case SessionStateEnumIdle:
		return "SessionStateEnumIdle"
	case SessionStateEnumListen:
		return "SessionStateEnumListen"
	case SessionStateEnumReset:
		return "SessionStateEnumReset"
	case SessionStateEnumDown:
		return "SessionStateEnumDown"
	case SessionStateEnumInit:
		return "SessionStateEnumInit"
	case SessionStateEnumInvalid:
		return "SessionStateEnumInvalid"
	}
	return "Invalid"
}

// SAFSession structure represents ISAFSession RPC structure.
type SAFSession dcom.InterfacePointer

func (o *SAFSession) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *SAFSession) xxx_PreparePayload(ctx context.Context) error {
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

func (o *SAFSession) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SAFSession) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *SAFSession) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// RemoteAssistanceServer structure represents IRASrv RPC structure.
//
// The IRASrv interface is implemented by the novice computer to allow the expert computer
// to request a Remote Assistance Connection String.<12>
//
// The UUID for this interface is: "F120A684-B926-447F-9DF4-C966CB785648".
//
// Opnums 3 and 4 are not used across the network. These opnums are reserved and MUST
// NOT be reused by non-Microsoft implementations.<13>
//
// Methods in RPC Opnum Order
//
//	+-------------------+----------------------------------------------------------------------------------+
//	|                   |                                                                                  |
//	|      METHOD       |                                   DESCRIPTION                                    |
//	|                   |                                                                                  |
//	+-------------------+----------------------------------------------------------------------------------+
//	+-------------------+----------------------------------------------------------------------------------+
//	| GetNoviceUserInfo | Retrieves the Remote Assistance Connection String information for a specified    |
//	|                   | terminal services session. Opnum: 7                                              |
//	+-------------------+----------------------------------------------------------------------------------+
//	| GetSessionInfo    | Returns the collection of the terminal services sessions on the remote novice    |
//	|                   | computer. This method also returns the number of terminal services sessions      |
//	|                   | running on the novice computer. All the terminal services session information    |
//	|                   | is returned as a SAFEARRAY. The elements of this SAFEARRAY are strings with      |
//	|                   | information about the DomainName, UserName, and SessionID for each terminal      |
//	|                   | services session. Opnum: 8                                                       |
//	+-------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown_QueryInterface, AddRef, and Release
// methods used by the standard COM IUnknown interface, as specified in [MS-DCOM] section
// 3.1.1.5.8. Opnums 5 and 6 are reserved for the GetIDsOfNames, and Invoke methods
// in the IDispatch interface, as specified in [MS-OAUT] section 3.1.
type RemoteAssistanceServer dcom.InterfacePointer

func (o *RemoteAssistanceServer) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteAssistanceServer) xxx_PreparePayload(ctx context.Context) error {
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

func (o *RemoteAssistanceServer) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteAssistanceServer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *RemoteAssistanceServer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PCHService structure represents IPCHService RPC structure.
//
// The IPCHService interface is implemented by the novice to allow the expert to request
// a Remote Assistance Connection String.<6>
//
// The UUID for this interface is: "833E4200-AFF7-4AC3-AAC2-9F24C1457BCE".
//
// Methods in RPC opnum order:
//
// * Opnums ( 1ef0b4e0-d88e-4b60-bc24-c983bec1e088#gt_e127848e-c66d-427d-b3aa-9f904fa4ada7
// ) 0, 1, and 2 are reserved for the IUnknown_QueryInterface, AddRef, and Release methods
// used by the standard COM IUnknown interface specified in [MS-DCOM] ( ../ms-dcom/4a893f3d-bd29-48cd-9f43-d9777a4415b0
// ) section 3.1.1.5.8 ( ../ms-dcom/2b4db106-fb79-4a67-b45f-63654f19c54c ).
//
// * Opnums 3 and 4 are not used across the network. These opnums are reserved and MUST
// NOT be used. <7> ( 90b2d2e5-7931-4762-8949-04617e1d9089#Appendix_A_7 )
//
// * Opnums 5 and 6 are reserved for the GetIDsOfNames and Invoke methods in the IDispatch
// interface specified in [MS-OAUT] ( ../ms-oaut/bbb05720-f724-45c7-8d17-f83c3d1a3961
// ) section 3.1 ( ../ms-oaut/c2c7dbe2-bafa-49da-93a7-7b75499ef90a ).
//
// * Opnums 7 through 18 and opnum 21 are not used by this protocol.
//
// # Method
//
// # Description
//
// RemoteConnectionParms ( 253af556-3edc-40d0-8ffa-5a6e23e6ecb5 )
//
// Gets the Remote Assistance connection parameters for a specified UserName, DomainName,
// and SessionID triple.
//
// Opnum: 19
//
// RemoteUserSessionInfo ( a7abc4db-6ad1-4aed-9cf7-72428f72ce96 )
//
// Returns the collection of the terminal services ( 1ef0b4e0-d88e-4b60-bc24-c983bec1e088#gt_ffff3f01-8c21-44d3-bbda-0062a1fbda4b
// ) sessions on the remote novice computer. All the terminal services session information
// is returned as a standard IPCHCollection interface. The members of this collection
// are objects of type ISAFSession. ISAFSession includes the DomainName, UserName, SessionID,
// and User SessionState for each session.
//
// Opnum: 20
type PCHService dcom.InterfacePointer

func (o *PCHService) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *PCHService) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PCHService) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PCHService) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PCHService) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PCHCollection structure represents IPCHCollection RPC structure.
type PCHCollection dcom.InterfacePointer

func (o *PCHCollection) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *PCHCollection) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PCHCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PCHCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PCHCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PCHService class identifier 833e4010-aff7-4ac3-aac2-9f24c1457bce
var PCHServiceClassID = &dcom.ClassID{Data1: 0x833e4010, Data2: 0xaff7, Data3: 0x4ac3, Data4: []byte{0xaa, 0xc2, 0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}

// RASrv class identifier 3c3a70a7-a468-49b9-8ada-28e11fccad5d
var RaServerClassID = &dcom.ClassID{Data1: 0x3c3a70a7, Data2: 0xa468, Data3: 0x49b9, Data4: []byte{0x8a, 0xda, 0x28, 0xe1, 0x1f, 0xcc, 0xad, 0x5d}}
