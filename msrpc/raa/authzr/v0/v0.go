package authzr

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
	GoPackage = "raa"
)

var (
	// Syntax UUID
	AuthzrSyntaxUUID = &uuid.UUID{TimeLow: 0xb1c2170, TimeMid: 0x5732, TimeHiAndVersion: 0x4e0e, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xd3, Node: [6]uint8{0xd9, 0xb1, 0x6f, 0x3b, 0x84, 0xd7}}
	// Syntax ID
	AuthzrSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AuthzrSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// authzr interface.
type AuthzrClient interface {

	// The AuthzrFreeContext method (opnum 0) frees all remote structures and memory associated
	// with the client context identified by the ContextHandle parameter.
	//
	// Return Values:
	//
	// If the function succeeds, it MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero 32-bit error code.
	//
	// When a remote authorization server receives this message, it MUST look up the ClientContext
	// structure in the ClientContextTable ADM element and free all structures and memory
	// associated with the ClientContext.
	FreeContext(context.Context, *FreeContextRequest, ...dcerpc.CallOption) (*FreeContextResponse, error)

	// The AuthzrInitializeContextFromSid method (opnum 1) creates a client context from
	// a given security identifier (SID). For domain SIDs, token group and claim attributes
	// will be retrieved from Active Directory through Kerberos.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code.
	InitializeContextFromSID(context.Context, *InitializeContextFromSIDRequest, ...dcerpc.CallOption) (*InitializeContextFromSIDResponse, error)

	// The AuthzrInitializeCompoundContext method (opnum 2) creates a compound context from
	// two specified context handles.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000. If the function fails,
	// it MUST return a nonzero value.
	InitializeCompoundContext(context.Context, *InitializeCompoundContextRequest, ...dcerpc.CallOption) (*InitializeCompoundContextResponse, error)

	// The AuthzrAccessCheck method (opnum 3) determines which access bits can be granted
	// to a client for a given set of security descriptors. The AUTHZR_ACCESS_REPLY structure
	// returns an array of granted access masks and error status.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code.
	AccessCheck(context.Context, *AccessCheckRequest, ...dcerpc.CallOption) (*AccessCheckResponse, error)

	// The AuthzGetInformationFromContext method (opnum 4) returns information about the
	// identified client context.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	GetInformationFromContext(context.Context, *GetInformationFromContextRequest, ...dcerpc.CallOption) (*GetInformationFromContextResponse, error)

	// The AuthzrModifyClaims method (opnum 5) modifies information about the identified
	// client context.
	//
	// Return Values:
	//
	// If the function succeeds, the function MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	ModifyClaims(context.Context, *ModifyClaimsRequest, ...dcerpc.CallOption) (*ModifyClaimsResponse, error)

	// The AuthzrModifySids method (opnum 6) modifies the list of SIDs associated with the
	// identified client context.
	//
	// Return Values:
	//
	// If the function succeeds, it MUST return 0x00000000.
	//
	// If the function fails, it MUST return a nonzero error code value.
	ModifySIDs(context.Context, *ModifySIDsRequest, ...dcerpc.CallOption) (*ModifySIDsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Handle structure represents AUTHZR_HANDLE RPC structure.
type Handle dcetypes.ContextHandle

func (o *Handle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Handle) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Handle) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Handle) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AccessRequest structure represents AUTHZR_ACCESS_REQUEST RPC structure.
//
// The AUTHZR_ACCESS_REQUEST structure defines the contents of a remote access check
// request.
type AccessRequest struct {
	// DesiredAccess:  The type of access to test.
	DesiredAccess uint32 `idl:"name:DesiredAccess" json:"desired_access"`
	// PrincipalSelfSid:  A pointer to the security identifier (SID) to use for the principal
	// self SID in the access control list (ACL).
	PrincipalSelfSID *dtyp.SID `idl:"name:PrincipalSelfSid" json:"principal_self_sid"`
	// ObjectTypeListLength:  The number of elements in the ObjectTypeList array.
	ObjectTypeListLength uint32 `idl:"name:ObjectTypeListLength" json:"object_type_list_length"`
	// ObjectTypeList:  A pointer to an array of OBJECT_TYPE_LIST structures in the object
	// tree for the object.
	ObjectTypeList []*dtyp.ObjectTypeList `idl:"name:ObjectTypeList;size_is:(ObjectTypeListLength)" json:"object_type_list"`
}

func (o *AccessRequest) xxx_PreparePayload(ctx context.Context) error {
	if o.ObjectTypeList != nil && o.ObjectTypeListLength == 0 {
		o.ObjectTypeListLength = uint32(len(o.ObjectTypeList))
	}
	if o.ObjectTypeListLength > uint32(256) {
		return fmt.Errorf("ObjectTypeListLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DesiredAccess); err != nil {
		return err
	}
	if o.PrincipalSelfSID != nil {
		_ptr_PrincipalSelfSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.PrincipalSelfSID != nil {
				if err := o.PrincipalSelfSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.PrincipalSelfSID, _ptr_PrincipalSelfSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ObjectTypeListLength); err != nil {
		return err
	}
	if o.ObjectTypeList != nil || o.ObjectTypeListLength > 0 {
		_ptr_ObjectTypeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ObjectTypeListLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ObjectTypeList {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ObjectTypeList[i1] != nil {
					if err := o.ObjectTypeList[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.ObjectTypeList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ObjectTypeList); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.ObjectTypeList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectTypeList, _ptr_ObjectTypeList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessRequest) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DesiredAccess); err != nil {
		return err
	}
	_ptr_PrincipalSelfSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.PrincipalSelfSID == nil {
			o.PrincipalSelfSID = &dtyp.SID{}
		}
		if err := o.PrincipalSelfSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_PrincipalSelfSid := func(ptr interface{}) { o.PrincipalSelfSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.PrincipalSelfSID, _s_PrincipalSelfSid, _ptr_PrincipalSelfSid); err != nil {
		return err
	}
	if err := w.ReadData(&o.ObjectTypeListLength); err != nil {
		return err
	}
	_ptr_ObjectTypeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ObjectTypeListLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ObjectTypeListLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ObjectTypeList", sizeInfo[0])
		}
		o.ObjectTypeList = make([]*dtyp.ObjectTypeList, sizeInfo[0])
		for i1 := range o.ObjectTypeList {
			i1 := i1
			if o.ObjectTypeList[i1] == nil {
				o.ObjectTypeList[i1] = &dtyp.ObjectTypeList{}
			}
			if err := o.ObjectTypeList[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ObjectTypeList := func(ptr interface{}) { o.ObjectTypeList = *ptr.(*[]*dtyp.ObjectTypeList) }
	if err := w.ReadPointer(&o.ObjectTypeList, _s_ObjectTypeList, _ptr_ObjectTypeList); err != nil {
		return err
	}
	return nil
}

// SelfRelativeSecurityDescriptor structure represents SR_SD RPC structure.
//
// The SR_SD structure defines a self-relative security descriptor. A self-relative
// security descriptor contains the security descriptor structure itself and the necessary
// security information associated with the descriptor.
type SelfRelativeSecurityDescriptor struct {
	// dwLength: The length, in bytes, of the data pointed to in the pSrSd member.
	Length uint32 `idl:"name:dwLength" json:"length"`
	// pSrSd: A pointer to a self-relative security descriptor.
	SelfRelativeSecurityDescriptor []byte `idl:"name:pSrSd;size_is:(dwLength)" json:"self_relative_security_descriptor"`
}

func (o *SelfRelativeSecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if o.SelfRelativeSecurityDescriptor != nil && o.Length == 0 {
		o.Length = uint32(len(o.SelfRelativeSecurityDescriptor))
	}
	if o.Length < uint32(20) || o.Length > uint32(131228) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SelfRelativeSecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.SelfRelativeSecurityDescriptor != nil || o.Length > 0 {
		_ptr_pSrSd := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SelfRelativeSecurityDescriptor {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SelfRelativeSecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SelfRelativeSecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SelfRelativeSecurityDescriptor, _ptr_pSrSd); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SelfRelativeSecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pSrSd := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SelfRelativeSecurityDescriptor", sizeInfo[0])
		}
		o.SelfRelativeSecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SelfRelativeSecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SelfRelativeSecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pSrSd := func(ptr interface{}) { o.SelfRelativeSecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SelfRelativeSecurityDescriptor, _s_pSrSd, _ptr_pSrSd); err != nil {
		return err
	}
	return nil
}

// AccessReply structure represents AUTHZR_ACCESS_REPLY RPC structure.
//
// The AUTHZR_ACCESS_REPLY structure defines the contents of a remote access check reply.
type AccessReply struct {
	// ResultListLength:  The number of elements in the GrantedAccessMask and Error arrays.
	// This number matches the number of entries in the object type list structure used
	// in the access check. The length MUST be between 1 and 256. If no object type is used
	// to represent the object, ResultListLength MUST be set to 1.
	ResultListLength uint32 `idl:"name:ResultListLength" json:"result_list_length"`
	// GrantedAccessMask:  A pointer to an array of granted access masks.
	GrantedAccessMask []uint32 `idl:"name:GrantedAccessMask;size_is:(ResultListLength)" json:"granted_access_mask"`
	// Error:  A pointer to an array of DWORD error code results for each element of the
	// array.
	Error []uint32 `idl:"name:Error;size_is:(ResultListLength)" json:"error"`
}

func (o *AccessReply) xxx_PreparePayload(ctx context.Context) error {
	if o.GrantedAccessMask != nil && o.ResultListLength == 0 {
		o.ResultListLength = uint32(len(o.GrantedAccessMask))
	}
	if o.Error != nil && o.ResultListLength == 0 {
		o.ResultListLength = uint32(len(o.Error))
	}
	if o.ResultListLength > uint32(256) {
		return fmt.Errorf("ResultListLength is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessReply) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ResultListLength); err != nil {
		return err
	}
	if o.GrantedAccessMask != nil || o.ResultListLength > 0 {
		_ptr_GrantedAccessMask := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResultListLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GrantedAccessMask {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.GrantedAccessMask[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.GrantedAccessMask); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GrantedAccessMask, _ptr_GrantedAccessMask); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Error != nil || o.ResultListLength > 0 {
		_ptr_Error := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResultListLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Error {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Error[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Error); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Error, _ptr_Error); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *AccessReply) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResultListLength); err != nil {
		return err
	}
	_ptr_GrantedAccessMask := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResultListLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResultListLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GrantedAccessMask", sizeInfo[0])
		}
		o.GrantedAccessMask = make([]uint32, sizeInfo[0])
		for i1 := range o.GrantedAccessMask {
			i1 := i1
			if err := w.ReadData(&o.GrantedAccessMask[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_GrantedAccessMask := func(ptr interface{}) { o.GrantedAccessMask = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.GrantedAccessMask, _s_GrantedAccessMask, _ptr_GrantedAccessMask); err != nil {
		return err
	}
	_ptr_Error := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResultListLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResultListLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Error", sizeInfo[0])
		}
		o.Error = make([]uint32, sizeInfo[0])
		for i1 := range o.Error {
			i1 := i1
			if err := w.ReadData(&o.Error[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Error := func(ptr interface{}) { o.Error = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Error, _s_Error, _ptr_Error); err != nil {
		return err
	}
	return nil
}

// ContextInformationClass type represents AUTHZ_CONTEXT_INFORMATION_CLASS RPC enumeration.
//
// The AUTHZ_CONTEXT_INFORMATION_CLASS enumeration is used to indicate security attributes
// of a principal represented by an AUTHZR_HANDLE.
type ContextInformationClass uint16

var (
	// AuthzContextInfoUserSid: Indicates the principal's user SID and its attribute.
	ContextInformationClassUserSID ContextInformationClass = 1
	// AuthzContextInfoGroupsSids: Indicates the groups to which the principal belongs and
	// their attributes.
	ContextInformationClassGroupsSIDs ContextInformationClass = 2
	// AuthzContextInfoRestrictedSids: Indicates the restricted SIDs in the principal's
	// security context and their attributes.
	ContextInformationClassRestrictedSIDs ContextInformationClass = 3
	// ReservedEnumValue4: Not used.
	ContextInformationClassReservedEnumValue4 ContextInformationClass = 4
	// ReservedEnumValue5: Not used.
	ContextInformationClassReservedEnumValue5 ContextInformationClass = 5
	// ReservedEnumValue6: Not used.
	ContextInformationClassReservedEnumValue6 ContextInformationClass = 6
	// ReservedEnumValue7: Not used.
	ContextInformationClassReservedEnumValue7 ContextInformationClass = 7
	// ReservedEnumValue8: Not used.
	ContextInformationClassReservedEnumValue8 ContextInformationClass = 8
	// ReservedEnumValue9: Not used.
	ContextInformationClassReservedEnumValue9 ContextInformationClass = 9
	// ReservedEnumValue10: Not used.
	ContextInformationClassReservedEnumValue10 ContextInformationClass = 10
	// ReservedEnumValue11: Not used.
	ContextInformationClassReservedEnumValue11 ContextInformationClass = 11
	// AuthzContextInfoDeviceSids: Indicates the groups to which the device principal belongs
	// and their attributes.
	ContextInformationClassDeviceSIDs ContextInformationClass = 12
	// AuthzContextInfoUserClaims: Indicates the user's security attributes information.
	ContextInformationClassUserClaims ContextInformationClass = 13
	// AuthzContextInfoDeviceClaims: Indicates the device's security attributes information.
	ContextInformationClassDeviceClaims ContextInformationClass = 14
	// ReservedEnumValue15: Not used.
	ContextInformationClassReservedEnumValue15 ContextInformationClass = 15
	// ReservedEnumValue16: Not used.
	ContextInformationClassReservedEnumValue16 ContextInformationClass = 16
)

func (o ContextInformationClass) String() string {
	switch o {
	case ContextInformationClassUserSID:
		return "ContextInformationClassUserSID"
	case ContextInformationClassGroupsSIDs:
		return "ContextInformationClassGroupsSIDs"
	case ContextInformationClassRestrictedSIDs:
		return "ContextInformationClassRestrictedSIDs"
	case ContextInformationClassReservedEnumValue4:
		return "ContextInformationClassReservedEnumValue4"
	case ContextInformationClassReservedEnumValue5:
		return "ContextInformationClassReservedEnumValue5"
	case ContextInformationClassReservedEnumValue6:
		return "ContextInformationClassReservedEnumValue6"
	case ContextInformationClassReservedEnumValue7:
		return "ContextInformationClassReservedEnumValue7"
	case ContextInformationClassReservedEnumValue8:
		return "ContextInformationClassReservedEnumValue8"
	case ContextInformationClassReservedEnumValue9:
		return "ContextInformationClassReservedEnumValue9"
	case ContextInformationClassReservedEnumValue10:
		return "ContextInformationClassReservedEnumValue10"
	case ContextInformationClassReservedEnumValue11:
		return "ContextInformationClassReservedEnumValue11"
	case ContextInformationClassDeviceSIDs:
		return "ContextInformationClassDeviceSIDs"
	case ContextInformationClassUserClaims:
		return "ContextInformationClassUserClaims"
	case ContextInformationClassDeviceClaims:
		return "ContextInformationClassDeviceClaims"
	case ContextInformationClassReservedEnumValue15:
		return "ContextInformationClassReservedEnumValue15"
	case ContextInformationClassReservedEnumValue16:
		return "ContextInformationClassReservedEnumValue16"
	}
	return "Invalid"
}

// SIDAndAttributes structure represents AUTHZR_SID_AND_ATTRIBUTES RPC structure.
//
// The AUTHZR_SID_AND_ATTRIBUTES structure contains information about the security identifiers
// (SIDs) in a token.
type SIDAndAttributes struct {
	// Sid:  A SID structure, as specified in [MS-DTYP] section 2.4.2.3. This is a pass-through
	// value and SHOULD NOT be interpreted by the RAZA protocol.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// Attributes:  Specifies attributes associated with the SID. This is a pass-through
	// value and SHOULD NOT be interpreted by the RAZA protocol.
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *SIDAndAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDAndAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SID != nil {
		_ptr_Sid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SID != nil {
				if err := o.SID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SID, _ptr_Sid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *SIDAndAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_Sid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sid := func(ptr interface{}) { o.SID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.SID, _s_Sid, _ptr_Sid); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// TokenUser structure represents AUTHZR_TOKEN_USER RPC structure.
//
// The AUTHZR_TOKEN_USER structure identifies the user associated with a token.
type TokenUser struct {
	// User:  Contains an AUTHZR_SID_AND_ATTRIBUTES structure (section 2.2.3.8) representing
	// the user associated with the access token.
	User *SIDAndAttributes `idl:"name:User" json:"user"`
}

func (o *TokenUser) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TokenUser) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.User != nil {
		if err := o.User.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TokenUser) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.User == nil {
		o.User = &SIDAndAttributes{}
	}
	if err := o.User.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// TokenGroups structure represents AUTHZR_TOKEN_GROUPS RPC structure.
//
// The AUTHZR_TOKEN_GROUPS structure represents a security identifier (SID) and its
// attributes.
type TokenGroups struct {
	// GroupCount:  Indicates the number of structures in the Groups array.
	GroupCount uint32 `idl:"name:GroupCount" json:"group_count"`
	// Groups:  An array of AUTHZR_SID_AND_ATTRIBUTES structures (section 2.2.3.8) representing
	// groups associated with the token.
	Groups []*SIDAndAttributes `idl:"name:Groups;size_is:(GroupCount)" json:"groups"`
}

func (o *TokenGroups) xxx_PreparePayload(ctx context.Context) error {
	if o.Groups != nil && o.GroupCount == 0 {
		o.GroupCount = uint32(len(o.Groups))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *TokenGroups) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.GroupCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TokenGroups) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	for i1 := range o.Groups {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Groups[i1] != nil {
			if err := o.Groups[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Groups); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *TokenGroups) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.GroupCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.GroupCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Groups", sizeInfo[0])
	}
	o.Groups = make([]*SIDAndAttributes, sizeInfo[0])
	for i1 := range o.Groups {
		i1 := i1
		if o.Groups[i1] == nil {
			o.Groups[i1] = &SIDAndAttributes{}
		}
		if err := o.Groups[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// SecurityAttributeStringValue structure represents AUTHZR_SECURITY_ATTRIBUTE_STRING_VALUE RPC structure.
//
// The AUTHZR_SECURITY_ATTRIBUTE_STRING_VALUE structure contains the string value of
// a claim.
type SecurityAttributeStringValue struct {
	// Length: The length of the string in the Value parameter.
	Length uint32 `idl:"name:Length" json:"length"`
	// Value: A Unicode string containing the pass-through string value of the claim.
	Value string `idl:"name:Value;size_is:(Length);string" json:"value"`
}

func (o *SecurityAttributeStringValue) xxx_PreparePayload(ctx context.Context) error {
	if o.Value != "" && o.Length == 0 {
		o.Length = uint32(len(o.Value))
	}
	if o.Length < uint32(2) || o.Length > uint32(32768) {
		return fmt.Errorf("Length is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeStringValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Value != "" || o.Length > 0 {
		_ptr_Value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.Value)
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
			_Value_buf := utf16.Encode([]rune(o.Value))
			if uint64(len(_Value_buf)) > sizeInfo[0]-1 {
				_Value_buf = _Value_buf[:sizeInfo[0]-1]
			}
			if o.Value != ndr.ZeroString {
				_Value_buf = append(_Value_buf, uint16(0))
			}
			for i1 := range _Value_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Value_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Value_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_Value); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeStringValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_Value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Value_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Value_buf", sizeInfo[0])
		}
		_Value_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Value_buf {
			i1 := i1
			if err := w.ReadData(&_Value_buf[i1]); err != nil {
				return err
			}
		}
		o.Value = strings.TrimRight(string(utf16.Decode(_Value_buf)), ndr.ZeroString)
		return nil
	})
	_s_Value := func(ptr interface{}) { o.Value = *ptr.(*string) }
	if err := w.ReadPointer(&o.Value, _s_Value, _ptr_Value); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeV1Value structure represents AUTHZR_SECURITY_ATTRIBUTE_V1_VALUE RPC structure.
//
// The AUTHZR_SECURITY_ATTRIBUTE_V1_VALUE structure defines a claim.
type SecurityAttributeV1Value struct {
	// ValueType: Identifies the type of the AttributeUnion member.
	//
	//	+----------------+----------------------------------------------------------------------------------+
	//	|                |                                                                                  |
	//	|     VALUE      |                                     MEANING                                      |
	//	|                |                                                                                  |
	//	+----------------+----------------------------------------------------------------------------------+
	//	+----------------+----------------------------------------------------------------------------------+
	//	| 0x0001         | AttributeUnion contains a LONG64 value.                                          |
	//	+----------------+----------------------------------------------------------------------------------+
	//	| 0x0002, 0x0006 | AttributeUnion contains a ULONG64 value.                                         |
	//	+----------------+----------------------------------------------------------------------------------+
	//	| 0x0003         | AttributeUnion contains an AUTHZR_SECURITY_ATTRIBUTE_STRING_VALUE structure, as  |
	//	|                | specified in section 2.2.3.4.                                                    |
	//	+----------------+----------------------------------------------------------------------------------+
	ValueType uint16 `idl:"name:ValueType" json:"value_type"`
	// AttributeUnion: A LONG64, ULONG64, or AUTHZR_SECURITY_ATTRIBUTE_STRING_VALUE, depending
	// on the value of ValueType.
	AttributeUnion *SecurityAttributeV1Value_AttributeUnion `idl:"name:AttributeUnion;switch_is:ValueType" json:"attribute_union"`
}

func (o *SecurityAttributeV1Value) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeV1Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueType); err != nil {
		return err
	}
	_swAttributeUnion := uint16(o.ValueType)
	if o.AttributeUnion != nil {
		if err := o.AttributeUnion.MarshalUnionNDR(ctx, w, _swAttributeUnion); err != nil {
			return err
		}
	} else {
		if err := (&SecurityAttributeV1Value_AttributeUnion{}).MarshalUnionNDR(ctx, w, _swAttributeUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeV1Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueType); err != nil {
		return err
	}
	if o.AttributeUnion == nil {
		o.AttributeUnion = &SecurityAttributeV1Value_AttributeUnion{}
	}
	_swAttributeUnion := uint16(o.ValueType)
	if err := o.AttributeUnion.UnmarshalUnionNDR(ctx, w, _swAttributeUnion); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeV1Value_AttributeUnion structure represents AUTHZR_SECURITY_ATTRIBUTE_V1_VALUE union anonymous member.
//
// The AUTHZR_SECURITY_ATTRIBUTE_V1_VALUE structure defines a claim.
type SecurityAttributeV1Value_AttributeUnion struct {
	// Types that are assignable to Value
	//
	// *SecurityAttributeV1Value_AttributeUnion_Int64
	// *SecurityAttributeV1Value_AttributeUnion_Uint64
	// *SecurityAttributeV1Value_AttributeUnion_String
	Value is_SecurityAttributeV1Value_AttributeUnion `json:"value"`
}

func (o *SecurityAttributeV1Value_AttributeUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SecurityAttributeV1Value_AttributeUnion_Int64:
		if value != nil {
			return value.Int64
		}
	case *SecurityAttributeV1Value_AttributeUnion_Uint64:
		if value != nil {
			return value.Uint64
		}
	case *SecurityAttributeV1Value_AttributeUnion_String:
		if value != nil {
			return value.String
		}
	}
	return nil
}

type is_SecurityAttributeV1Value_AttributeUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SecurityAttributeV1Value_AttributeUnion()
}

func (o *SecurityAttributeV1Value_AttributeUnion) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *SecurityAttributeV1Value_AttributeUnion_Int64:
		return uint16(1)
	case *SecurityAttributeV1Value_AttributeUnion_Uint64:
		switch sw {
		case uint16(2),
			uint16(6):
			return sw
		}
		return uint16(2)
	case *SecurityAttributeV1Value_AttributeUnion_String:
		return uint16(3)
	}
	return uint16(0)
}

func (o *SecurityAttributeV1Value_AttributeUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*SecurityAttributeV1Value_AttributeUnion_Int64)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityAttributeV1Value_AttributeUnion_Int64{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2),
		uint16(6):
		_o, _ := o.Value.(*SecurityAttributeV1Value_AttributeUnion_Uint64)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityAttributeV1Value_AttributeUnion_Uint64{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(3):
		_o, _ := o.Value.(*SecurityAttributeV1Value_AttributeUnion_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityAttributeV1Value_AttributeUnion_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *SecurityAttributeV1Value_AttributeUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &SecurityAttributeV1Value_AttributeUnion_Int64{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2),
		uint16(6):
		o.Value = &SecurityAttributeV1Value_AttributeUnion_Uint64{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(3):
		o.Value = &SecurityAttributeV1Value_AttributeUnion_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// SecurityAttributeV1Value_AttributeUnion_Int64 structure represents SecurityAttributeV1Value_AttributeUnion RPC union arm.
//
// It has following labels: 1
type SecurityAttributeV1Value_AttributeUnion_Int64 struct {
	Int64 int64 `idl:"name:Int64" json:"int64"`
}

func (*SecurityAttributeV1Value_AttributeUnion_Int64) is_SecurityAttributeV1Value_AttributeUnion() {}

func (o *SecurityAttributeV1Value_AttributeUnion_Int64) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int64); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributeV1Value_AttributeUnion_Int64) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int64); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeV1Value_AttributeUnion_Uint64 structure represents SecurityAttributeV1Value_AttributeUnion RPC union arm.
//
// It has following labels: 2, 6
type SecurityAttributeV1Value_AttributeUnion_Uint64 struct {
	Uint64 uint64 `idl:"name:Uint64" json:"uint64"`
}

func (*SecurityAttributeV1Value_AttributeUnion_Uint64) is_SecurityAttributeV1Value_AttributeUnion() {}

func (o *SecurityAttributeV1Value_AttributeUnion_Uint64) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint64); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributeV1Value_AttributeUnion_Uint64) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint64); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeV1Value_AttributeUnion_String structure represents SecurityAttributeV1Value_AttributeUnion RPC union arm.
//
// It has following labels: 3
type SecurityAttributeV1Value_AttributeUnion_String struct {
	String *SecurityAttributeStringValue `idl:"name:String" json:"string"`
}

func (*SecurityAttributeV1Value_AttributeUnion_String) is_SecurityAttributeV1Value_AttributeUnion() {}

func (o *SecurityAttributeV1Value_AttributeUnion_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != nil {
		if err := o.String.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SecurityAttributeStringValue{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeV1Value_AttributeUnion_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.String == nil {
		o.String = &SecurityAttributeStringValue{}
	}
	if err := o.String.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeV1 structure represents AUTHZR_SECURITY_ATTRIBUTE_V1 RPC structure.
//
// The AUTHZR_SECURITY_ATTRIBUTE_V1 structure specifies one or more security attribute
// and value pairs that are associated with a remote authorization context.
type SecurityAttributeV1 struct {
	// Length: The length of the Value parameter, in bytes. MUST be between 2 and 256.
	Length uint32 `idl:"name:Length" json:"length"`
	// Value: A Unicode string containing the security value. This string MUST be between
	// 2 and 256 bytes in length, inclusive.
	Value string `idl:"name:Value;size_is:(Length);string" json:"value"`
	// ValueType:  A union tag value indicating the type of information contained in Values
	// member.
	ValueType uint16 `idl:"name:ValueType" json:"value_type"`
	// Reserved: Reserved. This member MUST be set to zero when sent and MUST be ignored
	// when received.
	_ uint16 `idl:"name:Reserved"`
	// Flags:  MUST be zero or a combination of one or more of the following values.
	//
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                          |                                                                                  |
	//	|                          VALUE                           |                                   DESCRIPTION                                    |
	//	|                                                          |                                                                                  |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| AUTHZ_SECURITY_ATTRIBUTE_NON_INHERITABLE 0x00000001      | This security attribute is not inherited across processes.                       |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| AUTHZ_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE 0x00000002 | The value of the attribute is case sensitive. This flag is valid for values that |
	//	|                                                          | contain string types.                                                            |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ValueCount:  The number of attribute and value pairs pointed to by the Values member.
	// The number of attribute and value pairs MUST be between 0 and 1,024, inclusive.
	ValueCount uint32 `idl:"name:ValueCount" json:"value_count"`
	// Values: An array of AUTHZR_SECURITY_ATTRIBUTE_V1_VALUE structures, as defined in
	// section 2.2.3.6. Each structure contains a security attribute and value pair.
	Values []*SecurityAttributeV1Value `idl:"name:Values;size_is:(ValueCount)" json:"values"`
}

func (o *SecurityAttributeV1) xxx_PreparePayload(ctx context.Context) error {
	if o.Value != "" && o.Length == 0 {
		o.Length = uint32(len(o.Value))
	}
	if o.Values != nil && o.ValueCount == 0 {
		o.ValueCount = uint32(len(o.Values))
	}
	if o.Length < uint32(2) || o.Length > uint32(256) {
		return fmt.Errorf("Length is out of range")
	}
	if o.ValueCount > uint32(1024) {
		return fmt.Errorf("ValueCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeV1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Value != "" || o.Length > 0 {
		_ptr_Value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := ndr.UTF16NLen(o.Value)
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
			_Value_buf := utf16.Encode([]rune(o.Value))
			if uint64(len(_Value_buf)) > sizeInfo[0]-1 {
				_Value_buf = _Value_buf[:sizeInfo[0]-1]
			}
			if o.Value != ndr.ZeroString {
				_Value_buf = append(_Value_buf, uint16(0))
			}
			for i1 := range _Value_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Value_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Value_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_Value); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ValueType); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueCount); err != nil {
		return err
	}
	if o.Values != nil || o.ValueCount > 0 {
		_ptr_Values := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValueCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Values {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Values[i1] != nil {
					if err := o.Values[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributeV1Value{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Values); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SecurityAttributeV1Value{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Values, _ptr_Values); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributeV1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_Value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		var _Value_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Value_buf", sizeInfo[0])
		}
		_Value_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Value_buf {
			i1 := i1
			if err := w.ReadData(&_Value_buf[i1]); err != nil {
				return err
			}
		}
		o.Value = strings.TrimRight(string(utf16.Decode(_Value_buf)), ndr.ZeroString)
		return nil
	})
	_s_Value := func(ptr interface{}) { o.Value = *ptr.(*string) }
	if err := w.ReadPointer(&o.Value, _s_Value, _ptr_Value); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueType); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint16
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueCount); err != nil {
		return err
	}
	_ptr_Values := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValueCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValueCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Values", sizeInfo[0])
		}
		o.Values = make([]*SecurityAttributeV1Value, sizeInfo[0])
		for i1 := range o.Values {
			i1 := i1
			if o.Values[i1] == nil {
				o.Values[i1] = &SecurityAttributeV1Value{}
			}
			if err := o.Values[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Values := func(ptr interface{}) { o.Values = *ptr.(*[]*SecurityAttributeV1Value) }
	if err := w.ReadPointer(&o.Values, _s_Values, _ptr_Values); err != nil {
		return err
	}
	return nil
}

// SecurityAttributesInformation structure represents AUTHZR_SECURITY_ATTRIBUTES_INFORMATION RPC structure.
//
// The AUTHZR_SECURITY_ATTRUBUTES_INFORMATION structure specifies one or more security
// attributes.
type SecurityAttributesInformation struct {
	// Version:  The version of this structure. This value MUST be set to 0x0001.
	Version uint16 `idl:"name:Version" json:"version"`
	// Reserved:  Reserved.  This member MUST be set to zero when sent and MUST be ignored
	// when received.
	_ uint16 `idl:"name:Reserved"`
	// AttributeCount:  The number of attributes specified by the Attribute member. The
	// number of attributes MUST be between zero and 1,024, inclusive.
	AttributeCount uint32 `idl:"name:AttributeCount" json:"attribute_count"`
	// Attributes: A pointer to an array of AUTHZR_SECURITY_ATTRIBUTE_V1 structures, defined
	// in section 2.2.3.5.
	Attributes []*SecurityAttributeV1 `idl:"name:Attributes;size_is:(AttributeCount)" json:"attributes"`
}

func (o *SecurityAttributesInformation) xxx_PreparePayload(ctx context.Context) error {
	if o.Attributes != nil && o.AttributeCount == 0 {
		o.AttributeCount = uint32(len(o.Attributes))
	}
	if o.AttributeCount > uint32(1024) {
		return fmt.Errorf("AttributeCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributesInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.AttributeCount); err != nil {
		return err
	}
	if o.Attributes != nil || o.AttributeCount > 0 {
		_ptr_Attributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AttributeCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Attributes {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Attributes[i1] != nil {
					if err := o.Attributes[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributeV1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Attributes); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SecurityAttributeV1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Attributes, _ptr_Attributes); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributesInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint16
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.AttributeCount); err != nil {
		return err
	}
	_ptr_Attributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AttributeCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AttributeCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Attributes", sizeInfo[0])
		}
		o.Attributes = make([]*SecurityAttributeV1, sizeInfo[0])
		for i1 := range o.Attributes {
			i1 := i1
			if o.Attributes[i1] == nil {
				o.Attributes[i1] = &SecurityAttributeV1{}
			}
			if err := o.Attributes[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Attributes := func(ptr interface{}) { o.Attributes = *ptr.(*[]*SecurityAttributeV1) }
	if err := w.ReadPointer(&o.Attributes, _s_Attributes, _ptr_Attributes); err != nil {
		return err
	}
	return nil
}

// ContextInformation structure represents AUTHZR_CONTEXT_INFORMATION RPC structure.
//
// The AUTHZR_CONTEXT_INFORMATION structure contains security information about a principal.
type ContextInformation struct {
	// ValueType: Identifies the type of the ContextInfoUnion member.
	//
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                                   |                                                                                  |
	//	|                               VALUE                               |                                     MEANING                                      |
	//	|                                                                   |                                                                                  |
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0001 (user)                                                     | ContextInfoUnion contains an AUTHZR_TOKEN_USER structure, as specified in        |
	//	|                                                                   | section 2.2.3.10.                                                                |
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0002 (groups) 0x0003 (restricted groups) 0x000C (device groups) | ContextInfoUnion contains an AUTHZR_TOKEN_GROUPS structure, as specified in      |
	//	|                                                                   | section 2.2.3.9.                                                                 |
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000D (user claim) 0x000E (device claim)                         | ContextInfoUnion contains an AUTHZR_SECURITY_ATTRIBUTES_INFORMATION structure,   |
	//	|                                                                   | as specified in section 2.2.3.7.                                                 |
	//	+-------------------------------------------------------------------+----------------------------------------------------------------------------------+
	ValueType uint16 `idl:"name:ValueType" json:"value_type"`
	// ContextInfoUnion: A pointer to an AUTHZR_TOKEN_USER, AUTHZR_TOKEN_GROUPS, or AUTHZR_SECURITY_ATTRIBUTES_INFORMATION
	// structure, depending on the value of ValueType.
	ContextInfoUnion *ContextInformation_ContextInfoUnion `idl:"name:ContextInfoUnion;switch_is:ValueType" json:"context_info_union"`
}

func (o *ContextInformation) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueType); err != nil {
		return err
	}
	_swContextInfoUnion := uint16(o.ValueType)
	if o.ContextInfoUnion != nil {
		if err := o.ContextInfoUnion.MarshalUnionNDR(ctx, w, _swContextInfoUnion); err != nil {
			return err
		}
	} else {
		if err := (&ContextInformation_ContextInfoUnion{}).MarshalUnionNDR(ctx, w, _swContextInfoUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueType); err != nil {
		return err
	}
	if o.ContextInfoUnion == nil {
		o.ContextInfoUnion = &ContextInformation_ContextInfoUnion{}
	}
	_swContextInfoUnion := uint16(o.ValueType)
	if err := o.ContextInfoUnion.UnmarshalUnionNDR(ctx, w, _swContextInfoUnion); err != nil {
		return err
	}
	return nil
}

// ContextInformation_ContextInfoUnion structure represents AUTHZR_CONTEXT_INFORMATION union anonymous member.
//
// The AUTHZR_CONTEXT_INFORMATION structure contains security information about a principal.
type ContextInformation_ContextInfoUnion struct {
	// Types that are assignable to Value
	//
	// *ContextInformation_ContextInfoUnion_TokenUser
	// *ContextInformation_ContextInfoUnion_TokenGroups
	// *ContextInformation_ContextInfoUnion_TokenClaims
	Value is_ContextInformation_ContextInfoUnion `json:"value"`
}

func (o *ContextInformation_ContextInfoUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *ContextInformation_ContextInfoUnion_TokenUser:
		if value != nil {
			return value.TokenUser
		}
	case *ContextInformation_ContextInfoUnion_TokenGroups:
		if value != nil {
			return value.TokenGroups
		}
	case *ContextInformation_ContextInfoUnion_TokenClaims:
		if value != nil {
			return value.TokenClaims
		}
	}
	return nil
}

type is_ContextInformation_ContextInfoUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_ContextInformation_ContextInfoUnion()
}

func (o *ContextInformation_ContextInfoUnion) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *ContextInformation_ContextInfoUnion_TokenUser:
		return uint16(1)
	case *ContextInformation_ContextInfoUnion_TokenGroups:
		switch sw {
		case uint16(2),
			uint16(3),
			uint16(12):
			return sw
		}
		return uint16(2)
	case *ContextInformation_ContextInfoUnion_TokenClaims:
		switch sw {
		case uint16(13),
			uint16(14):
			return sw
		}
		return uint16(13)
	}
	return uint16(0)
}

func (o *ContextInformation_ContextInfoUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(7); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		_o, _ := o.Value.(*ContextInformation_ContextInfoUnion_TokenUser)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ContextInformation_ContextInfoUnion_TokenUser{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(2),
		uint16(3),
		uint16(12):
		_o, _ := o.Value.(*ContextInformation_ContextInfoUnion_TokenGroups)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ContextInformation_ContextInfoUnion_TokenGroups{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(13),
		uint16(14):
		_o, _ := o.Value.(*ContextInformation_ContextInfoUnion_TokenClaims)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ContextInformation_ContextInfoUnion_TokenClaims{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *ContextInformation_ContextInfoUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(7); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	switch sw {
	case uint16(1):
		o.Value = &ContextInformation_ContextInfoUnion_TokenUser{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(2),
		uint16(3),
		uint16(12):
		o.Value = &ContextInformation_ContextInfoUnion_TokenGroups{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(13),
		uint16(14):
		o.Value = &ContextInformation_ContextInfoUnion_TokenClaims{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// ContextInformation_ContextInfoUnion_TokenUser structure represents ContextInformation_ContextInfoUnion RPC union arm.
//
// It has following labels: 1
type ContextInformation_ContextInfoUnion_TokenUser struct {
	TokenUser *TokenUser `idl:"name:pTokenUser" json:"token_user"`
}

func (*ContextInformation_ContextInfoUnion_TokenUser) is_ContextInformation_ContextInfoUnion() {}

func (o *ContextInformation_ContextInfoUnion_TokenUser) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TokenUser != nil {
		_ptr_pTokenUser := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TokenUser != nil {
				if err := o.TokenUser.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&TokenUser{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TokenUser, _ptr_pTokenUser); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextInformation_ContextInfoUnion_TokenUser) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pTokenUser := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TokenUser == nil {
			o.TokenUser = &TokenUser{}
		}
		if err := o.TokenUser.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pTokenUser := func(ptr interface{}) { o.TokenUser = *ptr.(**TokenUser) }
	if err := w.ReadPointer(&o.TokenUser, _s_pTokenUser, _ptr_pTokenUser); err != nil {
		return err
	}
	return nil
}

// ContextInformation_ContextInfoUnion_TokenGroups structure represents ContextInformation_ContextInfoUnion RPC union arm.
//
// It has following labels: 2, 3, 12
type ContextInformation_ContextInfoUnion_TokenGroups struct {
	TokenGroups *TokenGroups `idl:"name:pTokenGroups" json:"token_groups"`
}

func (*ContextInformation_ContextInfoUnion_TokenGroups) is_ContextInformation_ContextInfoUnion() {}

func (o *ContextInformation_ContextInfoUnion_TokenGroups) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TokenGroups != nil {
		_ptr_pTokenGroups := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TokenGroups != nil {
				if err := o.TokenGroups.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&TokenGroups{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TokenGroups, _ptr_pTokenGroups); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextInformation_ContextInfoUnion_TokenGroups) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pTokenGroups := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TokenGroups == nil {
			o.TokenGroups = &TokenGroups{}
		}
		if err := o.TokenGroups.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pTokenGroups := func(ptr interface{}) { o.TokenGroups = *ptr.(**TokenGroups) }
	if err := w.ReadPointer(&o.TokenGroups, _s_pTokenGroups, _ptr_pTokenGroups); err != nil {
		return err
	}
	return nil
}

// ContextInformation_ContextInfoUnion_TokenClaims structure represents ContextInformation_ContextInfoUnion RPC union arm.
//
// It has following labels: 13, 14
type ContextInformation_ContextInfoUnion_TokenClaims struct {
	TokenClaims *SecurityAttributesInformation `idl:"name:pTokenClaims" json:"token_claims"`
}

func (*ContextInformation_ContextInfoUnion_TokenClaims) is_ContextInformation_ContextInfoUnion() {}

func (o *ContextInformation_ContextInfoUnion_TokenClaims) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TokenClaims != nil {
		_ptr_pTokenClaims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TokenClaims != nil {
				if err := o.TokenClaims.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SecurityAttributesInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TokenClaims, _ptr_pTokenClaims); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ContextInformation_ContextInfoUnion_TokenClaims) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pTokenClaims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TokenClaims == nil {
			o.TokenClaims = &SecurityAttributesInformation{}
		}
		if err := o.TokenClaims.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pTokenClaims := func(ptr interface{}) { o.TokenClaims = *ptr.(**SecurityAttributesInformation) }
	if err := w.ReadPointer(&o.TokenClaims, _s_pTokenClaims, _ptr_pTokenClaims); err != nil {
		return err
	}
	return nil
}

// SecurityAttributeOperation type represents AUTHZ_SECURITY_ATTRIBUTE_OPERATION RPC enumeration.
//
// The AUTHZ_SECURITY_ATTRIBUTE_OPERATION enumeration structure is used with the AuthzrModifyClaims
// operation (section 3.1.4.6) to identify operation types on a client context object.
type SecurityAttributeOperation uint16

var (
	// AUTHZ_SECURITY_ATTRIBUTE_OPERATION_NONE: No operation will be performed.
	SecurityAttributeOperationNone SecurityAttributeOperation = 0
	// AUTHZ_SECURITY_ATTRIBUTE_OPERATION_REPLACE_ALL: The ImpersonationAccessToken on the
	// specified client context will be replaced.
	SecurityAttributeOperationReplaceAll SecurityAttributeOperation = 1
	// AUTHZ_SECURITY_ATTRIBUTE_OPERATION_ADD: A new claim will be added to the server's
	// ImpersonationAccessToken associated with the specified client context.
	SecurityAttributeOperationAdd SecurityAttributeOperation = 2
	// AUTHZ_SECURITY_ATTRIBUTE_OPERATION_DELETE: An existing claim will be deleted from
	// the ImpersonationAccessToken array associated with the specified client context if
	// it is present in that array.
	SecurityAttributeOperationDelete SecurityAttributeOperation = 3
	// AUTHZ_SECURITY_ATTRIBUTE_OPERATION_REPLACE: An existing claim will be replaced in
	// the ImpersonationAccessToken array associated with the specified client context if
	// it is present in the array.
	SecurityAttributeOperationReplace SecurityAttributeOperation = 4
)

func (o SecurityAttributeOperation) String() string {
	switch o {
	case SecurityAttributeOperationNone:
		return "SecurityAttributeOperationNone"
	case SecurityAttributeOperationReplaceAll:
		return "SecurityAttributeOperationReplaceAll"
	case SecurityAttributeOperationAdd:
		return "SecurityAttributeOperationAdd"
	case SecurityAttributeOperationDelete:
		return "SecurityAttributeOperationDelete"
	case SecurityAttributeOperationReplace:
		return "SecurityAttributeOperationReplace"
	}
	return "Invalid"
}

// SIDOperation type represents AUTHZ_SID_OPERATION RPC enumeration.
//
// The AUTHZ_SID_OPERATION enumeration indicates the type of SID operations that can
// be made by a call to the AuthzrModifySids operation (section 3.1.4.7).
type SIDOperation uint16

var (
	// AUTHZ_SID_OPERATION_NONE: Do not modify anything.
	SIDOperationNone SIDOperation = 0
	// AUTHZ_SID_OPERATION_REPLACE_ALL: Replace the existing SIDs with the specified SIDs.
	// If replacement SIDs are not specified, delete the existing SIDs. This operation can
	// be specified only once and must be the only operation specified.
	SIDOperationReplaceAll SIDOperation = 1
	// AUTHZ_SID_OPERATION_ADD: Add a new SID. If the SID already exists, fail the call.
	SIDOperationAdd SIDOperation = 2
	// AUTHZ_SID_OPERATION_DELETE: Delete the specified SID. If the specified SID is not
	// found, fail the call without taking action.
	SIDOperationDelete SIDOperation = 3
	// AUTHZ_SID_OPERATION_REPLACE: Replace the existing SID with the specified SID. If
	// the SID does not exist, add the specified SID.
	SIDOperationReplace SIDOperation = 4
)

func (o SIDOperation) String() string {
	switch o {
	case SIDOperationNone:
		return "SIDOperationNone"
	case SIDOperationReplaceAll:
		return "SIDOperationReplaceAll"
	case SIDOperationAdd:
		return "SIDOperationAdd"
	case SIDOperationDelete:
		return "SIDOperationDelete"
	case SIDOperationReplace:
		return "SIDOperationReplace"
	}
	return "Invalid"
}

type xxx_DefaultAuthzrClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultAuthzrClient) FreeContext(ctx context.Context, in *FreeContextRequest, opts ...dcerpc.CallOption) (*FreeContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FreeContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) InitializeContextFromSID(ctx context.Context, in *InitializeContextFromSIDRequest, opts ...dcerpc.CallOption) (*InitializeContextFromSIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitializeContextFromSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) InitializeCompoundContext(ctx context.Context, in *InitializeCompoundContextRequest, opts ...dcerpc.CallOption) (*InitializeCompoundContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitializeCompoundContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) AccessCheck(ctx context.Context, in *AccessCheckRequest, opts ...dcerpc.CallOption) (*AccessCheckResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AccessCheckResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) GetInformationFromContext(ctx context.Context, in *GetInformationFromContextRequest, opts ...dcerpc.CallOption) (*GetInformationFromContextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInformationFromContextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) ModifyClaims(ctx context.Context, in *ModifyClaimsRequest, opts ...dcerpc.CallOption) (*ModifyClaimsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ModifyClaimsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) ModifySIDs(ctx context.Context, in *ModifySIDsRequest, opts ...dcerpc.CallOption) (*ModifySIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ModifySIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAuthzrClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAuthzrClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewAuthzrClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AuthzrClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AuthzrSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultAuthzrClient{cc: cc}, nil
}

// xxx_FreeContextOperation structure represents the AuthzrFreeContext operation
type xxx_FreeContextOperation struct {
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_FreeContextOperation) OpNum() int { return 0 }

func (o *xxx_FreeContextOperation) OpName() string { return "/authzr/v0/AuthzrFreeContext" }

func (o *xxx_FreeContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FreeContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FreeContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FreeContextRequest structure represents the AuthzrFreeContext operation request
type FreeContextRequest struct {
	// ContextHandle: A pointer to an AUTHZR_HANDLE structure, as defined in section 2.2.1.1.
	// This handle indicates the client context to be freed.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
}

func (o *FreeContextRequest) xxx_ToOp(ctx context.Context, op *xxx_FreeContextOperation) *xxx_FreeContextOperation {
	if op == nil {
		op = &xxx_FreeContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	return op
}

func (o *FreeContextRequest) xxx_FromOp(ctx context.Context, op *xxx_FreeContextOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
}
func (o *FreeContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FreeContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FreeContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FreeContextResponse structure represents the AuthzrFreeContext operation response
type FreeContextResponse struct {
	// ContextHandle: A pointer to an AUTHZR_HANDLE structure, as defined in section 2.2.1.1.
	// This handle indicates the client context to be freed.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// Return: The AuthzrFreeContext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FreeContextResponse) xxx_ToOp(ctx context.Context, op *xxx_FreeContextOperation) *xxx_FreeContextOperation {
	if op == nil {
		op = &xxx_FreeContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
	return op
}

func (o *FreeContextResponse) xxx_FromOp(ctx context.Context, op *xxx_FreeContextOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
}
func (o *FreeContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FreeContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FreeContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeContextFromSIDOperation structure represents the AuthzrInitializeContextFromSid operation
type xxx_InitializeContextFromSIDOperation struct {
	Flags          uint32             `idl:"name:Flags" json:"flags"`
	SID            *dtyp.SID          `idl:"name:Sid" json:"sid"`
	ExpirationTime *dtyp.LargeInteger `idl:"name:pExpirationTime;pointer:unique" json:"expiration_time"`
	ID             *dtyp.LUID         `idl:"name:Identifier" json:"id"`
	ContextHandle  *Handle            `idl:"name:ContextHandle" json:"context_handle"`
	Return         uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeContextFromSIDOperation) OpNum() int { return 1 }

func (o *xxx_InitializeContextFromSIDOperation) OpName() string {
	return "/authzr/v0/AuthzrInitializeContextFromSid"
}

func (o *xxx_InitializeContextFromSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeContextFromSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// Sid {in} (1:{pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.SID != nil {
			if err := o.SID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pExpirationTime {in} (1:{pointer=unique}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		if o.ExpirationTime != nil {
			_ptr_pExpirationTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ExpirationTime != nil {
					if err := o.ExpirationTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExpirationTime, _ptr_pExpirationTime); err != nil {
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
	// Identifier {in} (1:{alias=LUID}(struct))
	{
		if o.ID != nil {
			if err := o.ID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.LUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_InitializeContextFromSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// Sid {in} (1:{pointer=ref}*(1))(2:{alias=RPC_SID}(struct))
	{
		if o.SID == nil {
			o.SID = &dtyp.SID{}
		}
		if err := o.SID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pExpirationTime {in} (1:{pointer=unique}*(1))(2:{alias=LARGE_INTEGER}(struct))
	{
		_ptr_pExpirationTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ExpirationTime == nil {
				o.ExpirationTime = &dtyp.LargeInteger{}
			}
			if err := o.ExpirationTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pExpirationTime := func(ptr interface{}) { o.ExpirationTime = *ptr.(**dtyp.LargeInteger) }
		if err := w.ReadPointer(&o.ExpirationTime, _s_pExpirationTime, _ptr_pExpirationTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Identifier {in} (1:{alias=LUID}(struct))
	{
		if o.ID == nil {
			o.ID = &dtyp.LUID{}
		}
		if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeContextFromSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeContextFromSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ContextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeContextFromSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InitializeContextFromSIDRequest structure represents the AuthzrInitializeContextFromSid operation request
type InitializeContextFromSIDRequest struct {
	// Flags:  Indicates the type of logon behavior when initializing the client context.
	// The following flags are defined.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                 VALUE                 |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                            | When no flags are set, AuthzInitializeContextFromSid attempts to retrieve the    |
	//	|                                       | user's token group information by performing an S4U logon.                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| AUTHZ_COMPUTE_PRIVILEGES (0x00000008) | AuthzInitializeContextFromSid retrieves privileges for the new context. If       |
	//	|                                       | this function performs an S4U logon, it retrieves privileges from the token.     |
	//	|                                       | Otherwise, it retrieves privileges from all SIDs in the context.                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// Sid:  A pointer to the SID of the principal for whom a remote client context will
	// be created. This MUST be a valid user or computer account.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// pExpirationTime:  Reserved. This parameter MUST be set to NULL when sent and MUST
	// be ignored when received.
	ExpirationTime *dtyp.LargeInteger `idl:"name:pExpirationTime;pointer:unique" json:"expiration_time"`
	// Identifier:  Reserved. This parameter MUST be set to zero when sent and MUST be
	// ignored when received.
	ID *dtyp.LUID `idl:"name:Identifier" json:"id"`
}

func (o *InitializeContextFromSIDRequest) xxx_ToOp(ctx context.Context, op *xxx_InitializeContextFromSIDOperation) *xxx_InitializeContextFromSIDOperation {
	if op == nil {
		op = &xxx_InitializeContextFromSIDOperation{}
	}
	if o == nil {
		return op
	}
	o.Flags = op.Flags
	o.SID = op.SID
	o.ExpirationTime = op.ExpirationTime
	o.ID = op.ID
	return op
}

func (o *InitializeContextFromSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeContextFromSIDOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.SID = op.SID
	o.ExpirationTime = op.ExpirationTime
	o.ID = op.ID
}
func (o *InitializeContextFromSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitializeContextFromSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeContextFromSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeContextFromSIDResponse structure represents the AuthzrInitializeContextFromSid operation response
type InitializeContextFromSIDResponse struct {
	// ContextHandle: A pointer to an AUTHZR_HANDLE structure, as defined in section 2.2.1.1.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// Return: The AuthzrInitializeContextFromSid return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *InitializeContextFromSIDResponse) xxx_ToOp(ctx context.Context, op *xxx_InitializeContextFromSIDOperation) *xxx_InitializeContextFromSIDOperation {
	if op == nil {
		op = &xxx_InitializeContextFromSIDOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
	return op
}

func (o *InitializeContextFromSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeContextFromSIDOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Return = op.Return
}
func (o *InitializeContextFromSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitializeContextFromSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeContextFromSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_InitializeCompoundContextOperation structure represents the AuthzrInitializeCompoundContext operation
type xxx_InitializeCompoundContextOperation struct {
	User     *Handle `idl:"name:UserContextHandle" json:"user"`
	Device   *Handle `idl:"name:DeviceContextHandle" json:"device"`
	Compound *Handle `idl:"name:CompoundContextHandle" json:"compound"`
	Return   uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_InitializeCompoundContextOperation) OpNum() int { return 2 }

func (o *xxx_InitializeCompoundContextOperation) OpName() string {
	return "/authzr/v0/AuthzrInitializeCompoundContext"
}

func (o *xxx_InitializeCompoundContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeCompoundContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// UserContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.User != nil {
			if err := o.User.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// DeviceContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Device != nil {
			if err := o.Device.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_InitializeCompoundContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// UserContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.User == nil {
			o.User = &Handle{}
		}
		if err := o.User.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// DeviceContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Device == nil {
			o.Device = &Handle{}
		}
		if err := o.Device.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeCompoundContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeCompoundContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// CompoundContextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Compound != nil {
			if err := o.Compound.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitializeCompoundContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// CompoundContextHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Compound == nil {
			o.Compound = &Handle{}
		}
		if err := o.Compound.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InitializeCompoundContextRequest structure represents the AuthzrInitializeCompoundContext operation request
type InitializeCompoundContextRequest struct {
	// UserContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1, that
	// represents the user context for the compound context.
	User *Handle `idl:"name:UserContextHandle" json:"user"`
	// DeviceContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1, that
	// represents the device context for the compound context.
	Device *Handle `idl:"name:DeviceContextHandle" json:"device"`
}

func (o *InitializeCompoundContextRequest) xxx_ToOp(ctx context.Context, op *xxx_InitializeCompoundContextOperation) *xxx_InitializeCompoundContextOperation {
	if op == nil {
		op = &xxx_InitializeCompoundContextOperation{}
	}
	if o == nil {
		return op
	}
	o.User = op.User
	o.Device = op.Device
	return op
}

func (o *InitializeCompoundContextRequest) xxx_FromOp(ctx context.Context, op *xxx_InitializeCompoundContextOperation) {
	if o == nil {
		return
	}
	o.User = op.User
	o.Device = op.Device
}
func (o *InitializeCompoundContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitializeCompoundContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeCompoundContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitializeCompoundContextResponse structure represents the AuthzrInitializeCompoundContext operation response
type InitializeCompoundContextResponse struct {
	// CompoundContextHandle: A pointer to an AUTHZR_HANDLE structure, as defined in section
	// 2.2.1.1.
	Compound *Handle `idl:"name:CompoundContextHandle" json:"compound"`
	// Return: The AuthzrInitializeCompoundContext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *InitializeCompoundContextResponse) xxx_ToOp(ctx context.Context, op *xxx_InitializeCompoundContextOperation) *xxx_InitializeCompoundContextOperation {
	if op == nil {
		op = &xxx_InitializeCompoundContextOperation{}
	}
	if o == nil {
		return op
	}
	o.Compound = op.Compound
	o.Return = op.Return
	return op
}

func (o *InitializeCompoundContextResponse) xxx_FromOp(ctx context.Context, op *xxx_InitializeCompoundContextOperation) {
	if o == nil {
		return
	}
	o.Compound = op.Compound
	o.Return = op.Return
}
func (o *InitializeCompoundContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitializeCompoundContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitializeCompoundContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AccessCheckOperation structure represents the AuthzrAccessCheck operation
type xxx_AccessCheckOperation struct {
	ContextHandle           *Handle                           `idl:"name:ContextHandle" json:"context_handle"`
	Flags                   uint32                            `idl:"name:Flags" json:"flags"`
	Request                 *AccessRequest                    `idl:"name:pRequest" json:"request"`
	SecurityDescriptorCount uint32                            `idl:"name:SecurityDescriptorCount" json:"security_descriptor_count"`
	SecurityDescriptors     []*SelfRelativeSecurityDescriptor `idl:"name:pSecurityDescriptors;size_is:(SecurityDescriptorCount)" json:"security_descriptors"`
	Reply                   *AccessReply                      `idl:"name:pReply" json:"reply"`
	Return                  uint32                            `idl:"name:Return" json:"return"`
}

func (o *xxx_AccessCheckOperation) OpNum() int { return 3 }

func (o *xxx_AccessCheckOperation) OpName() string { return "/authzr/v0/AuthzrAccessCheck" }

func (o *xxx_AccessCheckOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SecurityDescriptors != nil && o.SecurityDescriptorCount == 0 {
		o.SecurityDescriptorCount = uint32(len(o.SecurityDescriptors))
	}
	if o.SecurityDescriptorCount < uint32(1) || o.SecurityDescriptorCount > uint32(16) {
		return fmt.Errorf("SecurityDescriptorCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pRequest {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REQUEST}(struct))
	{
		if o.Request != nil {
			if err := o.Request.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AccessRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// SecurityDescriptorCount {in} (1:{range=(1,16), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorCount); err != nil {
			return err
		}
	}
	// pSecurityDescriptors {in} (1:{pointer=ref}*(1))(2:{alias=SR_SD}[dim:0,size_is=SecurityDescriptorCount](struct))
	{
		dimSize1 := uint64(o.SecurityDescriptorCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SecurityDescriptors {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.SecurityDescriptors[i1] != nil {
				if err := o.SecurityDescriptors[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SelfRelativeSecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.SecurityDescriptors); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&SelfRelativeSecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pReply {in, out} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REPLY}(struct))
	{
		if o.Reply != nil {
			if err := o.Reply.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AccessReply{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pRequest {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REQUEST}(struct))
	{
		if o.Request == nil {
			o.Request = &AccessRequest{}
		}
		if err := o.Request.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SecurityDescriptorCount {in} (1:{range=(1,16), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorCount); err != nil {
			return err
		}
	}
	// pSecurityDescriptors {in} (1:{pointer=ref}*(1))(2:{alias=SR_SD}[dim:0,size_is=SecurityDescriptorCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptors", sizeInfo[0])
		}
		o.SecurityDescriptors = make([]*SelfRelativeSecurityDescriptor, sizeInfo[0])
		for i1 := range o.SecurityDescriptors {
			i1 := i1
			if o.SecurityDescriptors[i1] == nil {
				o.SecurityDescriptors[i1] = &SelfRelativeSecurityDescriptor{}
			}
			if err := o.SecurityDescriptors[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReply {in, out} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REPLY}(struct))
	{
		if o.Reply == nil {
			o.Reply = &AccessReply{}
		}
		if err := o.Reply.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pReply {in, out} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REPLY}(struct))
	{
		if o.Reply != nil {
			if err := o.Reply.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&AccessReply{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pReply {in, out} (1:{pointer=ref}*(1))(2:{alias=AUTHZR_ACCESS_REPLY}(struct))
	{
		if o.Reply == nil {
			o.Reply = &AccessReply{}
		}
		if err := o.Reply.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AccessCheckRequest structure represents the AuthzrAccessCheck operation request
type AccessCheckRequest struct {
	// ContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1, containing
	// the client context handle.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// Flags:   Reserved.  This parameter MUST be set to zero.
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// pRequest: A pointer to an AUTHZR_ACCESS_REQUEST structure, as defined in section
	// 2.2.3.2. This structure contains the body of the "what-if" access check request.
	Request *AccessRequest `idl:"name:pRequest" json:"request"`
	// SecurityDescriptorCount:  The number of security descriptors in the pSecurityDescriptors
	// parameter, not including the primary security descriptor.
	SecurityDescriptorCount uint32 `idl:"name:SecurityDescriptorCount" json:"security_descriptor_count"`
	// pSecurityDescriptors: A pointer to an array of SR_SD structures, as defined in section
	// 2.2.3.11. The first entry in this array is the primary security descriptor, and it
	// will be used as the security descriptor for the AccessCheck evaluation.
	SecurityDescriptors []*SelfRelativeSecurityDescriptor `idl:"name:pSecurityDescriptors;size_is:(SecurityDescriptorCount)" json:"security_descriptors"`
	// pReply: A pointer to an AUTHZR_ACCESS_REPLY structure, as defined in section 2.2.3.1.
	// This parameter will contain the body of the access check response.
	Reply *AccessReply `idl:"name:pReply" json:"reply"`
}

func (o *AccessCheckRequest) xxx_ToOp(ctx context.Context, op *xxx_AccessCheckOperation) *xxx_AccessCheckOperation {
	if op == nil {
		op = &xxx_AccessCheckOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.Flags = op.Flags
	o.Request = op.Request
	o.SecurityDescriptorCount = op.SecurityDescriptorCount
	o.SecurityDescriptors = op.SecurityDescriptors
	o.Reply = op.Reply
	return op
}

func (o *AccessCheckRequest) xxx_FromOp(ctx context.Context, op *xxx_AccessCheckOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.Flags = op.Flags
	o.Request = op.Request
	o.SecurityDescriptorCount = op.SecurityDescriptorCount
	o.SecurityDescriptors = op.SecurityDescriptors
	o.Reply = op.Reply
}
func (o *AccessCheckRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AccessCheckRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessCheckOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AccessCheckResponse structure represents the AuthzrAccessCheck operation response
type AccessCheckResponse struct {
	// pReply: A pointer to an AUTHZR_ACCESS_REPLY structure, as defined in section 2.2.3.1.
	// This parameter will contain the body of the access check response.
	Reply *AccessReply `idl:"name:pReply" json:"reply"`
	// Return: The AuthzrAccessCheck return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AccessCheckResponse) xxx_ToOp(ctx context.Context, op *xxx_AccessCheckOperation) *xxx_AccessCheckOperation {
	if op == nil {
		op = &xxx_AccessCheckOperation{}
	}
	if o == nil {
		return op
	}
	o.Reply = op.Reply
	o.Return = op.Return
	return op
}

func (o *AccessCheckResponse) xxx_FromOp(ctx context.Context, op *xxx_AccessCheckOperation) {
	if o == nil {
		return
	}
	o.Reply = op.Reply
	o.Return = op.Return
}
func (o *AccessCheckResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AccessCheckResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessCheckOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInformationFromContextOperation structure represents the AuthzGetInformationFromContext operation
type xxx_GetInformationFromContextOperation struct {
	ContextHandle      *Handle                 `idl:"name:ContextHandle" json:"context_handle"`
	InfoClass          ContextInformationClass `idl:"name:InfoClass" json:"info_class"`
	ContextInformation *ContextInformation     `idl:"name:ppContextInformation" json:"context_information"`
	Return             uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInformationFromContextOperation) OpNum() int { return 4 }

func (o *xxx_GetInformationFromContextOperation) OpName() string {
	return "/authzr/v0/AuthzGetInformationFromContext"
}

func (o *xxx_GetInformationFromContextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInformationFromContextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InfoClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.InfoClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInformationFromContextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InfoClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.InfoClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInformationFromContextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInformationFromContextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppContextInformation {out} (1:{pointer=ref}*(2)*(1))(2:{alias=AUTHZR_CONTEXT_INFORMATION}(struct))
	{
		if o.ContextInformation != nil {
			_ptr_ppContextInformation := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ContextInformation != nil {
					if err := o.ContextInformation.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ContextInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContextInformation, _ptr_ppContextInformation); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInformationFromContextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppContextInformation {out} (1:{pointer=ref}*(2)*(1))(2:{alias=AUTHZR_CONTEXT_INFORMATION}(struct))
	{
		_ptr_ppContextInformation := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ContextInformation == nil {
				o.ContextInformation = &ContextInformation{}
			}
			if err := o.ContextInformation.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppContextInformation := func(ptr interface{}) { o.ContextInformation = *ptr.(**ContextInformation) }
		if err := w.ReadPointer(&o.ContextInformation, _s_ppContextInformation, _ptr_ppContextInformation); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInformationFromContextRequest structure represents the AuthzGetInformationFromContext operation request
type GetInformationFromContextRequest struct {
	// ContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1. Represents
	// the client context to retrieve information from.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// InfoClass: An AUTHZ_CONTEXT_INFORMATION_CLASS enumeration, as defined in section
	// 2.2.2.1. Possible values for this field are specified in section 2.2.2.1.
	InfoClass ContextInformationClass `idl:"name:InfoClass" json:"info_class"`
}

func (o *GetInformationFromContextRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInformationFromContextOperation) *xxx_GetInformationFromContextOperation {
	if op == nil {
		op = &xxx_GetInformationFromContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.InfoClass = op.InfoClass
	return op
}

func (o *GetInformationFromContextRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInformationFromContextOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.InfoClass = op.InfoClass
}
func (o *GetInformationFromContextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInformationFromContextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInformationFromContextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInformationFromContextResponse structure represents the AuthzGetInformationFromContext operation response
type GetInformationFromContextResponse struct {
	// ppContextInformation: A two-layer pointer to an AUTHZR_CONTEXT_INFORMATION structure,
	// as defined in section 2.2.3.3. Used to return the context information.
	ContextInformation *ContextInformation `idl:"name:ppContextInformation" json:"context_information"`
	// Return: The AuthzGetInformationFromContext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInformationFromContextResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInformationFromContextOperation) *xxx_GetInformationFromContextOperation {
	if op == nil {
		op = &xxx_GetInformationFromContextOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextInformation = op.ContextInformation
	o.Return = op.Return
	return op
}

func (o *GetInformationFromContextResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInformationFromContextOperation) {
	if o == nil {
		return
	}
	o.ContextInformation = op.ContextInformation
	o.Return = op.Return
}
func (o *GetInformationFromContextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInformationFromContextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInformationFromContextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyClaimsOperation structure represents the AuthzrModifyClaims operation
type xxx_ModifyClaimsOperation struct {
	ContextHandle   *Handle                        `idl:"name:ContextHandle" json:"context_handle"`
	ClaimClass      ContextInformationClass        `idl:"name:ClaimClass" json:"claim_class"`
	OperationCount  uint32                         `idl:"name:OperationCount" json:"operation_count"`
	ClaimOperations []SecurityAttributeOperation   `idl:"name:pClaimOperations;size_is:(OperationCount)" json:"claim_operations"`
	Claims          *SecurityAttributesInformation `idl:"name:pClaims;pointer:unique" json:"claims"`
	Return          uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyClaimsOperation) OpNum() int { return 5 }

func (o *xxx_ModifyClaimsOperation) OpName() string { return "/authzr/v0/AuthzrModifyClaims" }

func (o *xxx_ModifyClaimsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ClaimOperations != nil && o.OperationCount == 0 {
		o.OperationCount = uint32(len(o.ClaimOperations))
	}
	if o.OperationCount < uint32(1) || o.OperationCount > uint32(65535) {
		return fmt.Errorf("OperationCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyClaimsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ClaimClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.ClaimClass)); err != nil {
			return err
		}
	}
	// OperationCount {in} (1:{range=(1,65535), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OperationCount); err != nil {
			return err
		}
	}
	// pClaimOperations {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZ_SECURITY_ATTRIBUTE_OPERATION}[dim:0,size_is=OperationCount](enum))
	{
		dimSize1 := uint64(o.OperationCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.ClaimOperations {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteEnum(uint16(o.ClaimOperations[i1])); err != nil {
				return err
			}
		}
		for i1 := len(o.ClaimOperations); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pClaims {in} (1:{pointer=unique}*(1))(2:{alias=AUTHZR_SECURITY_ATTRIBUTES_INFORMATION}(struct))
	{
		if o.Claims != nil {
			_ptr_pClaims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Claims != nil {
					if err := o.Claims.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributesInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Claims, _ptr_pClaims); err != nil {
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

func (o *xxx_ModifyClaimsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ClaimClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ClaimClass)); err != nil {
			return err
		}
	}
	// OperationCount {in} (1:{range=(1,65535), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OperationCount); err != nil {
			return err
		}
	}
	// pClaimOperations {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZ_SECURITY_ATTRIBUTE_OPERATION}[dim:0,size_is=OperationCount](enum))
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
			return fmt.Errorf("buffer overflow for size %d of array o.ClaimOperations", sizeInfo[0])
		}
		o.ClaimOperations = make([]SecurityAttributeOperation, sizeInfo[0])
		for i1 := range o.ClaimOperations {
			i1 := i1
			if err := w.ReadEnum((*uint16)(&o.ClaimOperations[i1])); err != nil {
				return err
			}
		}
	}
	// pClaims {in} (1:{pointer=unique}*(1))(2:{alias=AUTHZR_SECURITY_ATTRIBUTES_INFORMATION}(struct))
	{
		_ptr_pClaims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Claims == nil {
				o.Claims = &SecurityAttributesInformation{}
			}
			if err := o.Claims.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pClaims := func(ptr interface{}) { o.Claims = *ptr.(**SecurityAttributesInformation) }
		if err := w.ReadPointer(&o.Claims, _s_pClaims, _ptr_pClaims); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyClaimsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyClaimsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyClaimsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ModifyClaimsRequest structure represents the AuthzrModifyClaims operation request
type ModifyClaimsRequest struct {
	// ContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1. Represents
	// the client context to modify.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// ClaimClass: An AUTHZ_CONTEXT_INFORMATION_CLASS enumeration, as defined in section
	// 2.2.2.1. Indicates the claim class.
	ClaimClass ContextInformationClass `idl:"name:ClaimClass" json:"claim_class"`
	// OperationCount: The number of operations to be performed.
	OperationCount uint32 `idl:"name:OperationCount" json:"operation_count"`
	// pClaimOperations: A pointer to an array of AUTHZ_SECURITY_ATTRIBUTE_OPERATION enumerations,
	// as defined in section 2.2.2.2. Specifies the operations to be performed on each claim.
	ClaimOperations []SecurityAttributeOperation `idl:"name:pClaimOperations;size_is:(OperationCount)" json:"claim_operations"`
	// pClaims: A pointer to an array of AUTHZR_SECURITY_ATTRIBUTES_INFORMATION structures,
	// as defined in section 2.2.3.7. Contains the claim(s) used to modify the client context.
	Claims *SecurityAttributesInformation `idl:"name:pClaims;pointer:unique" json:"claims"`
}

func (o *ModifyClaimsRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyClaimsOperation) *xxx_ModifyClaimsOperation {
	if op == nil {
		op = &xxx_ModifyClaimsOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.ClaimClass = op.ClaimClass
	o.OperationCount = op.OperationCount
	o.ClaimOperations = op.ClaimOperations
	o.Claims = op.Claims
	return op
}

func (o *ModifyClaimsRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyClaimsOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.ClaimClass = op.ClaimClass
	o.OperationCount = op.OperationCount
	o.ClaimOperations = op.ClaimOperations
	o.Claims = op.Claims
}
func (o *ModifyClaimsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyClaimsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyClaimsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyClaimsResponse structure represents the AuthzrModifyClaims operation response
type ModifyClaimsResponse struct {
	// Return: The AuthzrModifyClaims return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ModifyClaimsResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyClaimsOperation) *xxx_ModifyClaimsOperation {
	if op == nil {
		op = &xxx_ModifyClaimsOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *ModifyClaimsResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyClaimsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ModifyClaimsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyClaimsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyClaimsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifySIDsOperation structure represents the AuthzrModifySids operation
type xxx_ModifySIDsOperation struct {
	ContextHandle  *Handle                 `idl:"name:ContextHandle" json:"context_handle"`
	SIDClass       ContextInformationClass `idl:"name:SidClass" json:"sid_class"`
	OperationCount uint32                  `idl:"name:OperationCount" json:"operation_count"`
	SIDOperations  []SIDOperation          `idl:"name:pSidOperations;size_is:(OperationCount)" json:"sid_operations"`
	SIDs           *TokenGroups            `idl:"name:pSids;pointer:unique" json:"sids"`
	Return         uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifySIDsOperation) OpNum() int { return 6 }

func (o *xxx_ModifySIDsOperation) OpName() string { return "/authzr/v0/AuthzrModifySids" }

func (o *xxx_ModifySIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SIDOperations != nil && o.OperationCount == 0 {
		o.OperationCount = uint32(len(o.SIDOperations))
	}
	if o.OperationCount < uint32(1) || o.OperationCount > uint32(65535) {
		return fmt.Errorf("OperationCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifySIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle != nil {
			if err := o.ContextHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SidClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.WriteEnum(uint16(o.SIDClass)); err != nil {
			return err
		}
	}
	// OperationCount {in} (1:{range=(1,65535), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OperationCount); err != nil {
			return err
		}
	}
	// pSidOperations {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZ_SID_OPERATION}[dim:0,size_is=OperationCount](enum))
	{
		dimSize1 := uint64(o.OperationCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SIDOperations {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteEnum(uint16(o.SIDOperations[i1])); err != nil {
				return err
			}
		}
		for i1 := len(o.SIDOperations); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	// pSids {in} (1:{pointer=unique}*(1))(2:{alias=AUTHZR_TOKEN_GROUPS}(struct))
	{
		if o.SIDs != nil {
			_ptr_pSids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SIDs != nil {
					if err := o.SIDs.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TokenGroups{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SIDs, _ptr_pSids); err != nil {
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

func (o *xxx_ModifySIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ContextHandle {in} (1:{context_handle, alias=AUTHZR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ContextHandle == nil {
			o.ContextHandle = &Handle{}
		}
		if err := o.ContextHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SidClass {in} (1:{alias=AUTHZ_CONTEXT_INFORMATION_CLASS}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.SIDClass)); err != nil {
			return err
		}
	}
	// OperationCount {in} (1:{range=(1,65535), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OperationCount); err != nil {
			return err
		}
	}
	// pSidOperations {in} (1:{pointer=ref}*(1))(2:{alias=AUTHZ_SID_OPERATION}[dim:0,size_is=OperationCount](enum))
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
			return fmt.Errorf("buffer overflow for size %d of array o.SIDOperations", sizeInfo[0])
		}
		o.SIDOperations = make([]SIDOperation, sizeInfo[0])
		for i1 := range o.SIDOperations {
			i1 := i1
			if err := w.ReadEnum((*uint16)(&o.SIDOperations[i1])); err != nil {
				return err
			}
		}
	}
	// pSids {in} (1:{pointer=unique}*(1))(2:{alias=AUTHZR_TOKEN_GROUPS}(struct))
	{
		_ptr_pSids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SIDs == nil {
				o.SIDs = &TokenGroups{}
			}
			if err := o.SIDs.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSids := func(ptr interface{}) { o.SIDs = *ptr.(**TokenGroups) }
		if err := w.ReadPointer(&o.SIDs, _s_pSids, _ptr_pSids); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifySIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifySIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifySIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ModifySIDsRequest structure represents the AuthzrModifySids operation request
type ModifySIDsRequest struct {
	// ContextHandle: An AUTHZR_HANDLE structure, as defined in section 2.2.1.1, representing
	// the client context to be modified.
	ContextHandle *Handle `idl:"name:ContextHandle" json:"context_handle"`
	// SidClass: An AUTHZ_CONTEXT_INFORMATION_CLASS enumeration value, as defined in section
	// 2.2.2.1, indicating the SID class.
	SIDClass ContextInformationClass `idl:"name:SidClass" json:"sid_class"`
	// OperationCount: The number of operations to be performed.
	OperationCount uint32 `idl:"name:OperationCount" json:"operation_count"`
	// pSidOperations: A pointer to an array of AUTHZ_SID_OPERATION enumeration values that
	// specify the group modifications to be made.
	SIDOperations []SIDOperation `idl:"name:pSidOperations;size_is:(OperationCount)" json:"sid_operations"`
	// pSids: A pointer to an AUTHZR_TOKEN_GROUPS structure, as defined in section 2.2.3.9,
	// specifying the groups to be modified.
	SIDs *TokenGroups `idl:"name:pSids;pointer:unique" json:"sids"`
}

func (o *ModifySIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifySIDsOperation) *xxx_ModifySIDsOperation {
	if op == nil {
		op = &xxx_ModifySIDsOperation{}
	}
	if o == nil {
		return op
	}
	o.ContextHandle = op.ContextHandle
	o.SIDClass = op.SIDClass
	o.OperationCount = op.OperationCount
	o.SIDOperations = op.SIDOperations
	o.SIDs = op.SIDs
	return op
}

func (o *ModifySIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifySIDsOperation) {
	if o == nil {
		return
	}
	o.ContextHandle = op.ContextHandle
	o.SIDClass = op.SIDClass
	o.OperationCount = op.OperationCount
	o.SIDOperations = op.SIDOperations
	o.SIDs = op.SIDs
}
func (o *ModifySIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifySIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifySIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifySIDsResponse structure represents the AuthzrModifySids operation response
type ModifySIDsResponse struct {
	// Return: The AuthzrModifySids return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ModifySIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifySIDsOperation) *xxx_ModifySIDsOperation {
	if op == nil {
		op = &xxx_ModifySIDsOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *ModifySIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifySIDsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ModifySIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifySIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifySIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
