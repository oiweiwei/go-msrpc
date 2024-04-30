// The adts package implements the ADTS client protocol.
//
// # Introduction
//
// This is the primary specification for Active Directory, both Active Directory Domain
// Services (AD DS) and Active Directory Lightweight Directory Services (AD LDS). When
// the specification does not refer specifically to AD DS or AD LDS, it applies to both.
// The state model for this specification is prerequisite to the other specifications
// for Active Directory: [MS-DRSR] and [MS-SRPL].
//
// When no operating system version information is specified, information in this document
// applies to all relevant versions of Windows. Similarly, when no DC functional level
// is specified, information in this document applies to all DC functional levels.
//
// The information in this specification is applicable to the following Microsoft products
// or supplemental software. References to product versions include released service
// packs.
//
// # Overview
//
// This is the primary specification for Active Directory. The state model for this
// specification is prerequisite to the other specifications for Active Directory: [MS-DRSR]
// and [MS-SRPL].
//
// Active Directory is either deployed as AD DS or as AD LDS. This document describes
// both forms. When the specification does not refer specifically to AD DS or AD LDS,
// it applies to both.
//
// The remainder of this section describes the structure of this document.
//
// The basic state model is specified in section 3.1.1.1. The basic state model is prerequisite
// to the remainder of the document. Section 3.1.1.1 also includes descriptive content
// to introduce key concepts and refer to places in the document where the full specification
// is given.
//
// The schema completes the state model and is specified in section 3.1.1.2. The schema
// is prerequisite to the remainder of the document.
//
// Active Directory is a server for LDAP. Section 3.1.1.3 specifies the extensions and
// variations of LDAP that are supported by Active Directory.
//
// LDAP is an access protocol that determines very little about the behavior of the
// data being accessed. Section 3.1.1.4 specifies read (LDAP Search) behaviors, and
// section 3.1.1.5 specifies update (LDAP Add, Modify, Modify DN, Delete) behaviors.
// Section 3.1.1.6 specifies background tasks required due to write operations, to the
// extent that those tasks are exposed by protocols.
//
// One of the update behaviors is the maintenance of the change log for use by Windows
// NT 4.0 operating system backup domain controller (BDC) replication [MS-NRPC] section
// 3.6. The maintenance of this change log is specified in section 3.1.1.7.
//
// The security services that Active Directory offers clients of LDAP are specified
// in section 5.1.
//
// Active Directory contains a number of objects, visible through LDAP, that have special
// significance to the system. Section 6.1 specifies these objects.
//
// A server running Active Directory is part of a distributed system that performs replication.
// The Knowledge Consistency Checker (KCC) is a component that is used to create spanning
// trees for DC-to-DC replication and is specified in section 6.2.
//
// A server running Active Directory is responsible for publishing the services that
// it offers, in order to eliminate the administrative burden of configuring clients
// to use particular servers running Active Directory. A server running Active Directory
// also implements the server side of the LDAP ping and mailslot ping protocols to aid
// clients in selecting among all the servers offering the same service. Section 6.3
// specifies how a server running Active Directory publishes its services, and how a
// client needing some service can use this publication plus the LDAP ping or mailslot
// ping to locate a suitable server.
//
// Computers in a network with Active Directory can be put into a state called "domain
// joined"; when in this state, the computer can authenticate itself. Section 6.4 specifies
// both the state in Active Directory and the state on a computer required for the domain
// joined state.
//
// Each type of data stored in Active Directory has an associated function that compares
// two values to determine if they are equal and, if not, which is greater. Section
// 3.1.1.2 specifies all but one of these functions; the methodology for comparing two
// Unicode strings is specified in section 6.5.
package adts

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	claims "github.com/oiweiwei/go-msrpc/msrpc/adts/claims/claims/v1"
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
	_ = dtyp.GoPackage
	_ = claims.GoPackage
)

var (
	// import guard
	GoPackage = "adts"
)

// PACInfoBuffer structure represents PAC_INFO_BUFFER RPC structure.
type PACInfoBuffer struct {
	Type         uint32 `idl:"name:ulType" json:"type"`
	BufferLength uint32 `idl:"name:cbBufferSize" json:"buffer_length"`
	Offset       uint64 `idl:"name:Offset" json:"offset"`
}

func (o *PACInfoBuffer) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACInfoBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if err := w.WriteData(o.Offset); err != nil {
		return err
	}
	return nil
}
func (o *PACInfoBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.Offset); err != nil {
		return err
	}
	return nil
}

// PACType structure represents PACTYPE RPC structure.
type PACType struct {
	BuffersCount uint32           `idl:"name:cBuffers" json:"buffers_count"`
	Version      uint32           `idl:"name:Version" json:"version"`
	Buffers      []*PACInfoBuffer `idl:"name:Buffers" json:"buffers"`
}

func (o *PACType) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.BuffersCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	for i1 := range o.Buffers {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if o.Buffers[i1] != nil {
			if err := o.Buffers[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&PACInfoBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Buffers); uint64(i1) < 1; i1++ {
		if err := (&PACInfoBuffer{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.BuffersCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	o.Buffers = make([]*PACInfoBuffer, 1)
	for i1 := range o.Buffers {
		i1 := i1
		if o.Buffers[i1] == nil {
			o.Buffers[i1] = &PACInfoBuffer{}
		}
		if err := o.Buffers[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// PACCredentialInfo structure represents PAC_CREDENTIAL_INFO RPC structure.
type PACCredentialInfo struct {
	Version        uint32 `idl:"name:Version" json:"version"`
	EncryptionType uint32 `idl:"name:EncryptionType" json:"encryption_type"`
	SerializedData []byte `idl:"name:SerializedData" json:"serialized_data"`
}

func (o *PACCredentialInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.EncryptionType); err != nil {
		return err
	}
	for i1 := range o.SerializedData {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.SerializedData[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.SerializedData); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.EncryptionType); err != nil {
		return err
	}
	o.SerializedData = make([]byte, 1)
	for i1 := range o.SerializedData {
		i1 := i1
		if err := w.ReadData(&o.SerializedData[i1]); err != nil {
			return err
		}
	}
	return nil
}

// SecurityPackageSupplementalCred structure represents SECPKG_SUPPLEMENTAL_CRED RPC structure.
type SecurityPackageSupplementalCred struct {
	PackageName    *dtyp.UnicodeString `idl:"name:PackageName" json:"package_name"`
	CredentialSize uint32              `idl:"name:CredentialSize" json:"credential_size"`
	Credentials    []byte              `idl:"name:Credentials;size_is:(CredentialSize)" json:"credentials"`
}

func (o *SecurityPackageSupplementalCred) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialSize == 0 {
		o.CredentialSize = uint32(len(o.Credentials))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityPackageSupplementalCred) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.PackageName != nil {
		if err := o.PackageName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CredentialSize); err != nil {
		return err
	}
	if o.Credentials != nil || o.CredentialSize > 0 {
		_ptr_Credentials := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CredentialSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Credentials {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Credentials[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Credentials, _ptr_Credentials); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityPackageSupplementalCred) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.PackageName == nil {
		o.PackageName = &dtyp.UnicodeString{}
	}
	if err := o.PackageName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.CredentialSize); err != nil {
		return err
	}
	_ptr_Credentials := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CredentialSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CredentialSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
		}
		o.Credentials = make([]byte, sizeInfo[0])
		for i1 := range o.Credentials {
			i1 := i1
			if err := w.ReadData(&o.Credentials[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Credentials := func(ptr interface{}) { o.Credentials = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Credentials, _s_Credentials, _ptr_Credentials); err != nil {
		return err
	}
	return nil
}

// PACCredentialData structure represents PAC_CREDENTIAL_DATA RPC structure.
type PACCredentialData struct {
	CredentialCount uint32                             `idl:"name:CredentialCount" json:"credential_count"`
	Credentials     []*SecurityPackageSupplementalCred `idl:"name:Credentials;size_is:(CredentialCount)" json:"credentials"`
}

func (o *PACCredentialData) xxx_PreparePayload(ctx context.Context) error {
	if o.Credentials != nil && o.CredentialCount == 0 {
		o.CredentialCount = uint32(len(o.Credentials))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialData) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.CredentialCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PACCredentialData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.CredentialCount); err != nil {
		return err
	}
	for i1 := range o.Credentials {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Credentials[i1] != nil {
			if err := o.Credentials[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityPackageSupplementalCred{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Credentials); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SecurityPackageSupplementalCred{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACCredentialData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.CredentialCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.CredentialCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.CredentialCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Credentials", sizeInfo[0])
	}
	o.Credentials = make([]*SecurityPackageSupplementalCred, sizeInfo[0])
	for i1 := range o.Credentials {
		i1 := i1
		if o.Credentials[i1] == nil {
			o.Credentials[i1] = &SecurityPackageSupplementalCred{}
		}
		if err := o.Credentials[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// PACClientInfo structure represents PAC_CLIENT_INFO RPC structure.
type PACClientInfo struct {
	ClientID   *dtyp.Filetime `idl:"name:ClientId" json:"client_id"`
	NameLength uint16         `idl:"name:NameLength" json:"name_length"`
	Name       []uint16       `idl:"name:Name" json:"name"`
}

func (o *PACClientInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.ClientID != nil {
		if err := o.ClientID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NameLength); err != nil {
		return err
	}
	for i1 := range o.Name {
		i1 := i1
		if uint64(i1) >= 1 {
			break
		}
		if err := w.WriteData(o.Name[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Name); uint64(i1) < 1; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.ClientID == nil {
		o.ClientID = &dtyp.Filetime{}
	}
	if err := o.ClientID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameLength); err != nil {
		return err
	}
	o.Name = make([]uint16, 1)
	for i1 := range o.Name {
		i1 := i1
		if err := w.ReadData(&o.Name[i1]); err != nil {
			return err
		}
	}
	return nil
}

// NTLMSupplementalCredential structure represents NTLM_SUPPLEMENTAL_CREDENTIAL RPC structure.
type NTLMSupplementalCredential struct {
	Version    uint32 `idl:"name:Version" json:"version"`
	Flags      uint32 `idl:"name:Flags" json:"flags"`
	LMPassword []byte `idl:"name:LmPassword" json:"lm_password"`
	NTPassword []byte `idl:"name:NtPassword" json:"nt_password"`
}

func (o *NTLMSupplementalCredential) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMSupplementalCredential) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Version); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	for i1 := range o.LMPassword {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.LMPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LMPassword); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.NTPassword {
		i1 := i1
		if uint64(i1) >= 16 {
			break
		}
		if err := w.WriteData(o.NTPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.NTPassword); uint64(i1) < 16; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *NTLMSupplementalCredential) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Version); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	o.LMPassword = make([]byte, 16)
	for i1 := range o.LMPassword {
		i1 := i1
		if err := w.ReadData(&o.LMPassword[i1]); err != nil {
			return err
		}
	}
	o.NTPassword = make([]byte, 16)
	for i1 := range o.NTPassword {
		i1 := i1
		if err := w.ReadData(&o.NTPassword[i1]); err != nil {
			return err
		}
	}
	return nil
}

// CypherBlock structure represents CYPHER_BLOCK RPC structure.
type CypherBlock struct {
	Data []byte `idl:"name:data" json:"data"`
}

func (o *CypherBlock) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CypherBlock) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CypherBlock) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]byte, 8)
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// UserSessionKey structure represents USER_SESSION_KEY RPC structure.
type UserSessionKey struct {
	Data []*CypherBlock `idl:"name:data" json:"data"`
}

func (o *UserSessionKey) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserSessionKey) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= 2 {
			break
		}
		if o.Data[i1] != nil {
			if err := o.Data[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&CypherBlock{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Data); uint64(i1) < 2; i1++ {
		if err := (&CypherBlock{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserSessionKey) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Data = make([]*CypherBlock, 2)
	for i1 := range o.Data {
		i1 := i1
		if o.Data[i1] == nil {
			o.Data[i1] = &CypherBlock{}
		}
		if err := o.Data[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// KerberosSIDAndAttributes structure represents KERB_SID_AND_ATTRIBUTES RPC structure.
type KerberosSIDAndAttributes struct {
	SID        *dtyp.SID `idl:"name:Sid" json:"sid"`
	Attributes uint32    `idl:"name:Attributes" json:"attributes"`
}

func (o *KerberosSIDAndAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosSIDAndAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *KerberosSIDAndAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// GroupMembership structure represents GROUP_MEMBERSHIP RPC structure.
type GroupMembership struct {
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	Attributes uint32 `idl:"name:Attributes" json:"attributes"`
}

func (o *GroupMembership) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GroupMembership) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeID); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	return nil
}
func (o *GroupMembership) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	return nil
}

// DomainGroupMembership structure represents DOMAIN_GROUP_MEMBERSHIP RPC structure.
type DomainGroupMembership struct {
	DomainID   *dtyp.SID          `idl:"name:DomainId" json:"domain_id"`
	GroupCount uint32             `idl:"name:GroupCount" json:"group_count"`
	GroupIDs   []*GroupMembership `idl:"name:GroupIds;size_is:(GroupCount)" json:"group_ids"`
}

func (o *DomainGroupMembership) xxx_PreparePayload(ctx context.Context) error {
	if o.GroupIDs != nil && o.GroupCount == 0 {
		o.GroupCount = uint32(len(o.GroupIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGroupMembership) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.DomainID != nil {
		_ptr_DomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DomainID != nil {
				if err := o.DomainID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainID, _ptr_DomainId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	if o.GroupIDs != nil || o.GroupCount > 0 {
		_ptr_GroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.GroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.GroupIDs[i1] != nil {
					if err := o.GroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.GroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupIDs, _ptr_GroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DomainGroupMembership) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_DomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DomainID == nil {
			o.DomainID = &dtyp.SID{}
		}
		if err := o.DomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DomainId := func(ptr interface{}) { o.DomainID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.DomainID, _s_DomainId, _ptr_DomainId); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	_ptr_GroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.GroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.GroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GroupIDs", sizeInfo[0])
		}
		o.GroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.GroupIDs {
			i1 := i1
			if o.GroupIDs[i1] == nil {
				o.GroupIDs[i1] = &GroupMembership{}
			}
			if err := o.GroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_GroupIds := func(ptr interface{}) { o.GroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.GroupIDs, _s_GroupIds, _ptr_GroupIds); err != nil {
		return err
	}
	return nil
}

// KerberosValidationInfo structure represents KERB_VALIDATION_INFO RPC structure.
type KerberosValidationInfo struct {
	LogonTime              *dtyp.Filetime              `idl:"name:LogonTime" json:"logon_time"`
	LogoffTime             *dtyp.Filetime              `idl:"name:LogoffTime" json:"logoff_time"`
	KickOffTime            *dtyp.Filetime              `idl:"name:KickOffTime" json:"kick_off_time"`
	PasswordLastSet        *dtyp.Filetime              `idl:"name:PasswordLastSet" json:"password_last_set"`
	PasswordCanChange      *dtyp.Filetime              `idl:"name:PasswordCanChange" json:"password_can_change"`
	PasswordMustChange     *dtyp.Filetime              `idl:"name:PasswordMustChange" json:"password_must_change"`
	EffectiveName          *dtyp.UnicodeString         `idl:"name:EffectiveName" json:"effective_name"`
	FullName               *dtyp.UnicodeString         `idl:"name:FullName" json:"full_name"`
	LogonScript            *dtyp.UnicodeString         `idl:"name:LogonScript" json:"logon_script"`
	ProfilePath            *dtyp.UnicodeString         `idl:"name:ProfilePath" json:"profile_path"`
	HomeDirectory          *dtyp.UnicodeString         `idl:"name:HomeDirectory" json:"home_directory"`
	HomeDirectoryDrive     *dtyp.UnicodeString         `idl:"name:HomeDirectoryDrive" json:"home_directory_drive"`
	LogonCount             uint16                      `idl:"name:LogonCount" json:"logon_count"`
	BadPasswordCount       uint16                      `idl:"name:BadPasswordCount" json:"bad_password_count"`
	UserID                 uint32                      `idl:"name:UserId" json:"user_id"`
	PrimaryGroupID         uint32                      `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	GroupCount             uint32                      `idl:"name:GroupCount" json:"group_count"`
	GroupIDs               []*GroupMembership          `idl:"name:GroupIds;size_is:(GroupCount)" json:"group_ids"`
	UserFlags              uint32                      `idl:"name:UserFlags" json:"user_flags"`
	UserSessionKey         *UserSessionKey             `idl:"name:UserSessionKey" json:"user_session_key"`
	LogonServer            *dtyp.UnicodeString         `idl:"name:LogonServer" json:"logon_server"`
	LogonDomainName        *dtyp.UnicodeString         `idl:"name:LogonDomainName" json:"logon_domain_name"`
	LogonDomainID          *dtyp.SID                   `idl:"name:LogonDomainId" json:"logon_domain_id"`
	_                      []uint32                    `idl:"name:Reserved1"`
	UserAccountControl     uint32                      `idl:"name:UserAccountControl" json:"user_account_control"`
	_                      []uint32                    `idl:"name:Reserved3"`
	SIDCount               uint32                      `idl:"name:SidCount" json:"sid_count"`
	ExtraSIDs              []*KerberosSIDAndAttributes `idl:"name:ExtraSids;size_is:(SidCount)" json:"extra_sids"`
	ResourceGroupDomainSID *dtyp.SID                   `idl:"name:ResourceGroupDomainSid" json:"resource_group_domain_sid"`
	ResourceGroupCount     uint32                      `idl:"name:ResourceGroupCount" json:"resource_group_count"`
	ResourceGroupIDs       []*GroupMembership          `idl:"name:ResourceGroupIds;size_is:(ResourceGroupCount)" json:"resource_group_ids"`
}

func (o *KerberosValidationInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.GroupIDs != nil && o.GroupCount == 0 {
		o.GroupCount = uint32(len(o.GroupIDs))
	}
	if o.ExtraSIDs != nil && o.SIDCount == 0 {
		o.SIDCount = uint32(len(o.ExtraSIDs))
	}
	if o.ResourceGroupIDs != nil && o.ResourceGroupCount == 0 {
		o.ResourceGroupCount = uint32(len(o.ResourceGroupIDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosValidationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.LogonTime != nil {
		if err := o.LogonTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogoffTime != nil {
		if err := o.LogoffTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.KickOffTime != nil {
		if err := o.KickOffTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordLastSet != nil {
		if err := o.PasswordLastSet.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordCanChange != nil {
		if err := o.PasswordCanChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PasswordMustChange != nil {
		if err := o.PasswordMustChange.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.EffectiveName != nil {
		if err := o.EffectiveName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.FullName != nil {
		if err := o.FullName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonScript != nil {
		if err := o.LogonScript.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ProfilePath != nil {
		if err := o.ProfilePath.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectory != nil {
		if err := o.HomeDirectory.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.HomeDirectoryDrive != nil {
		if err := o.HomeDirectoryDrive.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LogonCount); err != nil {
		return err
	}
	if err := w.WriteData(o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.WriteData(o.GroupCount); err != nil {
		return err
	}
	if o.GroupIDs != nil || o.GroupCount > 0 {
		_ptr_GroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.GroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.GroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.GroupIDs[i1] != nil {
					if err := o.GroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.GroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GroupIDs, _ptr_GroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserFlags); err != nil {
		return err
	}
	if o.UserSessionKey != nil {
		if err := o.UserSessionKey.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UserSessionKey{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonServer != nil {
		if err := o.LogonServer.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonDomainName != nil {
		if err := o.LogonDomainName.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.LogonDomainID != nil {
		_ptr_LogonDomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogonDomainID != nil {
				if err := o.LogonDomainID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogonDomainID, _ptr_LogonDomainId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	// reserved Reserved1
	for i1 := 0; uint64(i1) < 2; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.UserAccountControl); err != nil {
		return err
	}
	// reserved Reserved3
	for i1 := 0; uint64(i1) < 7; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SIDCount); err != nil {
		return err
	}
	if o.ExtraSIDs != nil || o.SIDCount > 0 {
		_ptr_ExtraSids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SIDCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ExtraSIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ExtraSIDs[i1] != nil {
					if err := o.ExtraSIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ExtraSIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtraSIDs, _ptr_ExtraSids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ResourceGroupDomainSID != nil {
		_ptr_ResourceGroupDomainSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ResourceGroupDomainSID != nil {
				if err := o.ResourceGroupDomainSID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResourceGroupDomainSID, _ptr_ResourceGroupDomainSid); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ResourceGroupCount); err != nil {
		return err
	}
	if o.ResourceGroupIDs != nil || o.ResourceGroupCount > 0 {
		_ptr_ResourceGroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ResourceGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ResourceGroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ResourceGroupIDs[i1] != nil {
					if err := o.ResourceGroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ResourceGroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ResourceGroupIDs, _ptr_ResourceGroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *KerberosValidationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.LogonTime == nil {
		o.LogonTime = &dtyp.Filetime{}
	}
	if err := o.LogonTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogoffTime == nil {
		o.LogoffTime = &dtyp.Filetime{}
	}
	if err := o.LogoffTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.KickOffTime == nil {
		o.KickOffTime = &dtyp.Filetime{}
	}
	if err := o.KickOffTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordLastSet == nil {
		o.PasswordLastSet = &dtyp.Filetime{}
	}
	if err := o.PasswordLastSet.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordCanChange == nil {
		o.PasswordCanChange = &dtyp.Filetime{}
	}
	if err := o.PasswordCanChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PasswordMustChange == nil {
		o.PasswordMustChange = &dtyp.Filetime{}
	}
	if err := o.PasswordMustChange.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.EffectiveName == nil {
		o.EffectiveName = &dtyp.UnicodeString{}
	}
	if err := o.EffectiveName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.FullName == nil {
		o.FullName = &dtyp.UnicodeString{}
	}
	if err := o.FullName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonScript == nil {
		o.LogonScript = &dtyp.UnicodeString{}
	}
	if err := o.LogonScript.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ProfilePath == nil {
		o.ProfilePath = &dtyp.UnicodeString{}
	}
	if err := o.ProfilePath.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectory == nil {
		o.HomeDirectory = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectory.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.HomeDirectoryDrive == nil {
		o.HomeDirectoryDrive = &dtyp.UnicodeString{}
	}
	if err := o.HomeDirectoryDrive.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.BadPasswordCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	if err := w.ReadData(&o.GroupCount); err != nil {
		return err
	}
	_ptr_GroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.GroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.GroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.GroupIDs", sizeInfo[0])
		}
		o.GroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.GroupIDs {
			i1 := i1
			if o.GroupIDs[i1] == nil {
				o.GroupIDs[i1] = &GroupMembership{}
			}
			if err := o.GroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_GroupIds := func(ptr interface{}) { o.GroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.GroupIDs, _s_GroupIds, _ptr_GroupIds); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserFlags); err != nil {
		return err
	}
	if o.UserSessionKey == nil {
		o.UserSessionKey = &UserSessionKey{}
	}
	if err := o.UserSessionKey.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonServer == nil {
		o.LogonServer = &dtyp.UnicodeString{}
	}
	if err := o.LogonServer.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.LogonDomainName == nil {
		o.LogonDomainName = &dtyp.UnicodeString{}
	}
	if err := o.LogonDomainName.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_LogonDomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogonDomainID == nil {
			o.LogonDomainID = &dtyp.SID{}
		}
		if err := o.LogonDomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_LogonDomainId := func(ptr interface{}) { o.LogonDomainID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.LogonDomainID, _s_LogonDomainId, _ptr_LogonDomainId); err != nil {
		return err
	}
	// reserved Reserved1
	var _Reserved1 []uint32
	_Reserved1 = make([]uint32, 2)
	for i1 := range _Reserved1 {
		i1 := i1
		if err := w.ReadData(&_Reserved1[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.UserAccountControl); err != nil {
		return err
	}
	// reserved Reserved3
	var _Reserved3 []uint32
	_Reserved3 = make([]uint32, 7)
	for i1 := range _Reserved3 {
		i1 := i1
		if err := w.ReadData(&_Reserved3[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.SIDCount); err != nil {
		return err
	}
	_ptr_ExtraSids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SIDCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SIDCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ExtraSIDs", sizeInfo[0])
		}
		o.ExtraSIDs = make([]*KerberosSIDAndAttributes, sizeInfo[0])
		for i1 := range o.ExtraSIDs {
			i1 := i1
			if o.ExtraSIDs[i1] == nil {
				o.ExtraSIDs[i1] = &KerberosSIDAndAttributes{}
			}
			if err := o.ExtraSIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ExtraSids := func(ptr interface{}) { o.ExtraSIDs = *ptr.(*[]*KerberosSIDAndAttributes) }
	if err := w.ReadPointer(&o.ExtraSIDs, _s_ExtraSids, _ptr_ExtraSids); err != nil {
		return err
	}
	_ptr_ResourceGroupDomainSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ResourceGroupDomainSID == nil {
			o.ResourceGroupDomainSID = &dtyp.SID{}
		}
		if err := o.ResourceGroupDomainSID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ResourceGroupDomainSid := func(ptr interface{}) { o.ResourceGroupDomainSID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.ResourceGroupDomainSID, _s_ResourceGroupDomainSid, _ptr_ResourceGroupDomainSid); err != nil {
		return err
	}
	if err := w.ReadData(&o.ResourceGroupCount); err != nil {
		return err
	}
	_ptr_ResourceGroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ResourceGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ResourceGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ResourceGroupIDs", sizeInfo[0])
		}
		o.ResourceGroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.ResourceGroupIDs {
			i1 := i1
			if o.ResourceGroupIDs[i1] == nil {
				o.ResourceGroupIDs[i1] = &GroupMembership{}
			}
			if err := o.ResourceGroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ResourceGroupIds := func(ptr interface{}) { o.ResourceGroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.ResourceGroupIDs, _s_ResourceGroupIds, _ptr_ResourceGroupIds); err != nil {
		return err
	}
	return nil
}

// S4UDelegationInfo structure represents S4U_DELEGATION_INFO RPC structure.
type S4UDelegationInfo struct {
	S4u2proxyTarget      *dtyp.UnicodeString   `idl:"name:S4U2proxyTarget" json:"s4u2proxy_target"`
	TransitedListSize    uint32                `idl:"name:TransitedListSize" json:"transited_list_size"`
	S4UTransitedServices []*dtyp.UnicodeString `idl:"name:S4UTransitedServices;size_is:(TransitedListSize)" json:"s4u_transited_services"`
}

func (o *S4UDelegationInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.S4UTransitedServices != nil && o.TransitedListSize == 0 {
		o.TransitedListSize = uint32(len(o.S4UTransitedServices))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *S4UDelegationInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.S4u2proxyTarget != nil {
		if err := o.S4u2proxyTarget.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TransitedListSize); err != nil {
		return err
	}
	if o.S4UTransitedServices != nil || o.TransitedListSize > 0 {
		_ptr_S4UTransitedServices := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.TransitedListSize)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.S4UTransitedServices {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.S4UTransitedServices[i1] != nil {
					if err := o.S4UTransitedServices[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.S4UTransitedServices); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.S4UTransitedServices, _ptr_S4UTransitedServices); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *S4UDelegationInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.S4u2proxyTarget == nil {
		o.S4u2proxyTarget = &dtyp.UnicodeString{}
	}
	if err := o.S4u2proxyTarget.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TransitedListSize); err != nil {
		return err
	}
	_ptr_S4UTransitedServices := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.TransitedListSize > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.TransitedListSize)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.S4UTransitedServices", sizeInfo[0])
		}
		o.S4UTransitedServices = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.S4UTransitedServices {
			i1 := i1
			if o.S4UTransitedServices[i1] == nil {
				o.S4UTransitedServices[i1] = &dtyp.UnicodeString{}
			}
			if err := o.S4UTransitedServices[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_S4UTransitedServices := func(ptr interface{}) { o.S4UTransitedServices = *ptr.(*[]*dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.S4UTransitedServices, _s_S4UTransitedServices, _ptr_S4UTransitedServices); err != nil {
		return err
	}
	return nil
}

// UPNDNSInfo structure represents UPN_DNS_INFO RPC structure.
type UPNDNSInfo struct {
	UPNLength           uint16 `idl:"name:UpnLength" json:"upn_length"`
	UPNOffset           uint16 `idl:"name:UpnOffset" json:"upn_offset"`
	DNSDomainNameLength uint16 `idl:"name:DnsDomainNameLength" json:"dns_domain_name_length"`
	DNSDomainNameOffset uint16 `idl:"name:DnsDomainNameOffset" json:"dns_domain_name_offset"`
	Flags               uint32 `idl:"name:Flags" json:"flags"`
}

func (o *UPNDNSInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UPNDNSInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.UPNLength); err != nil {
		return err
	}
	if err := w.WriteData(o.UPNOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DNSDomainNameLength); err != nil {
		return err
	}
	if err := w.WriteData(o.DNSDomainNameOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *UPNDNSInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.UPNLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.UPNOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSDomainNameLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.DNSDomainNameOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// PACClientClaimsInfo structure represents PAC_CLIENT_CLAIMS_INFO RPC structure.
type PACClientClaimsInfo struct {
	Claims *claims.ClaimsSetMetadata `idl:"name:Claims" json:"claims"`
}

func (o *PACClientClaimsInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientClaimsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Claims != nil {
		_ptr_Claims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Claims != nil {
				if err := o.Claims.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&claims.ClaimsSetMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Claims, _ptr_Claims); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACClientClaimsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Claims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Claims == nil {
			o.Claims = &claims.ClaimsSetMetadata{}
		}
		if err := o.Claims.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Claims := func(ptr interface{}) { o.Claims = *ptr.(**claims.ClaimsSetMetadata) }
	if err := w.ReadPointer(&o.Claims, _s_Claims, _ptr_Claims); err != nil {
		return err
	}
	return nil
}

// PACDeviceInfo structure represents PAC_DEVICE_INFO RPC structure.
type PACDeviceInfo struct {
	UserID            uint32                      `idl:"name:UserId" json:"user_id"`
	PrimaryGroupID    uint32                      `idl:"name:PrimaryGroupId" json:"primary_group_id"`
	AccountDomainID   *dtyp.SID                   `idl:"name:AccountDomainId" json:"account_domain_id"`
	AccountGroupCount uint32                      `idl:"name:AccountGroupCount" json:"account_group_count"`
	AccountGroupIDs   []*GroupMembership          `idl:"name:AccountGroupIds;size_is:(AccountGroupCount)" json:"account_group_ids"`
	SIDCount          uint32                      `idl:"name:SidCount" json:"sid_count"`
	ExtraSIDs         []*KerberosSIDAndAttributes `idl:"name:ExtraSids;size_is:(SidCount)" json:"extra_sids"`
	DomainGroupCount  uint32                      `idl:"name:DomainGroupCount" json:"domain_group_count"`
	DomainGroup       []*DomainGroupMembership    `idl:"name:DomainGroup;size_is:(DomainGroupCount)" json:"domain_group"`
}

func (o *PACDeviceInfo) xxx_PreparePayload(ctx context.Context) error {
	if o.AccountGroupIDs != nil && o.AccountGroupCount == 0 {
		o.AccountGroupCount = uint32(len(o.AccountGroupIDs))
	}
	if o.ExtraSIDs != nil && o.SIDCount == 0 {
		o.SIDCount = uint32(len(o.ExtraSIDs))
	}
	if o.DomainGroup != nil && o.DomainGroupCount == 0 {
		o.DomainGroupCount = uint32(len(o.DomainGroup))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if err := w.WriteData(o.PrimaryGroupID); err != nil {
		return err
	}
	if o.AccountDomainID != nil {
		_ptr_AccountDomainId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.AccountDomainID != nil {
				if err := o.AccountDomainID.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccountDomainID, _ptr_AccountDomainId); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AccountGroupCount); err != nil {
		return err
	}
	if o.AccountGroupIDs != nil || o.AccountGroupCount > 0 {
		_ptr_AccountGroupIds := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.AccountGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.AccountGroupIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.AccountGroupIDs[i1] != nil {
					if err := o.AccountGroupIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.AccountGroupIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&GroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.AccountGroupIDs, _ptr_AccountGroupIds); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.SIDCount); err != nil {
		return err
	}
	if o.ExtraSIDs != nil || o.SIDCount > 0 {
		_ptr_ExtraSids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SIDCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ExtraSIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ExtraSIDs[i1] != nil {
					if err := o.ExtraSIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ExtraSIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&KerberosSIDAndAttributes{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ExtraSIDs, _ptr_ExtraSids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DomainGroupCount); err != nil {
		return err
	}
	if o.DomainGroup != nil || o.DomainGroupCount > 0 {
		_ptr_DomainGroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.DomainGroupCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.DomainGroup {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.DomainGroup[i1] != nil {
					if err := o.DomainGroup[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&DomainGroupMembership{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.DomainGroup); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&DomainGroupMembership{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainGroup, _ptr_DomainGroup); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PrimaryGroupID); err != nil {
		return err
	}
	_ptr_AccountDomainId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.AccountDomainID == nil {
			o.AccountDomainID = &dtyp.SID{}
		}
		if err := o.AccountDomainID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_AccountDomainId := func(ptr interface{}) { o.AccountDomainID = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.AccountDomainID, _s_AccountDomainId, _ptr_AccountDomainId); err != nil {
		return err
	}
	if err := w.ReadData(&o.AccountGroupCount); err != nil {
		return err
	}
	_ptr_AccountGroupIds := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.AccountGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.AccountGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.AccountGroupIDs", sizeInfo[0])
		}
		o.AccountGroupIDs = make([]*GroupMembership, sizeInfo[0])
		for i1 := range o.AccountGroupIDs {
			i1 := i1
			if o.AccountGroupIDs[i1] == nil {
				o.AccountGroupIDs[i1] = &GroupMembership{}
			}
			if err := o.AccountGroupIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_AccountGroupIds := func(ptr interface{}) { o.AccountGroupIDs = *ptr.(*[]*GroupMembership) }
	if err := w.ReadPointer(&o.AccountGroupIDs, _s_AccountGroupIds, _ptr_AccountGroupIds); err != nil {
		return err
	}
	if err := w.ReadData(&o.SIDCount); err != nil {
		return err
	}
	_ptr_ExtraSids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SIDCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SIDCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ExtraSIDs", sizeInfo[0])
		}
		o.ExtraSIDs = make([]*KerberosSIDAndAttributes, sizeInfo[0])
		for i1 := range o.ExtraSIDs {
			i1 := i1
			if o.ExtraSIDs[i1] == nil {
				o.ExtraSIDs[i1] = &KerberosSIDAndAttributes{}
			}
			if err := o.ExtraSIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ExtraSids := func(ptr interface{}) { o.ExtraSIDs = *ptr.(*[]*KerberosSIDAndAttributes) }
	if err := w.ReadPointer(&o.ExtraSIDs, _s_ExtraSids, _ptr_ExtraSids); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainGroupCount); err != nil {
		return err
	}
	_ptr_DomainGroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.DomainGroupCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.DomainGroupCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DomainGroup", sizeInfo[0])
		}
		o.DomainGroup = make([]*DomainGroupMembership, sizeInfo[0])
		for i1 := range o.DomainGroup {
			i1 := i1
			if o.DomainGroup[i1] == nil {
				o.DomainGroup[i1] = &DomainGroupMembership{}
			}
			if err := o.DomainGroup[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_DomainGroup := func(ptr interface{}) { o.DomainGroup = *ptr.(*[]*DomainGroupMembership) }
	if err := w.ReadPointer(&o.DomainGroup, _s_DomainGroup, _ptr_DomainGroup); err != nil {
		return err
	}
	return nil
}

// PACDeviceClaimsInfo structure represents PAC_DEVICE_CLAIMS_INFO RPC structure.
type PACDeviceClaimsInfo struct {
	Claims *claims.ClaimsSetMetadata `idl:"name:Claims" json:"claims"`
}

func (o *PACDeviceClaimsInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceClaimsInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Claims != nil {
		_ptr_Claims := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Claims != nil {
				if err := o.Claims.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&claims.ClaimsSetMetadata{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Claims, _ptr_Claims); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *PACDeviceClaimsInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Claims := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Claims == nil {
			o.Claims = &claims.ClaimsSetMetadata{}
		}
		if err := o.Claims.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Claims := func(ptr interface{}) { o.Claims = *ptr.(**claims.ClaimsSetMetadata) }
	if err := w.ReadPointer(&o.Claims, _s_Claims, _ptr_Claims); err != nil {
		return err
	}
	return nil
}
