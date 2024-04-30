// The rpcl package implements the RPCL client protocol.
//
// # Introduction
//
// The Remote Procedure Call Location Services Extensions is a legacy protocol that
// has been deprecated and is not used by Microsoft. Implementers should use the DCE-based
// RPC Location Services capabilities in [C706].
//
// The Remote Procedure Call Location Services Extensions is a set of extensions/restrictions
// to the distributed computing environment (DCE) remote procedure call (RPC) Location
// Services specified in [C706]. These extensions add new capabilities to the DCE RPC
// Location Services Protocol.
//
// This document specifies a set of extensions and restrictions to the DCE RPC Location
// Services specification as specified in [C706].
//
// # Overview
//
// This specification extends the DCE RPC Location Services specification defined in
// the section "Name Service Interface" in Part 2 of [C706]. These extensions add new
// capabilities to the DCE RPC Location Services Protocol and, in some cases, place
// additional restrictions upon it. This specification adheres to the abstract data
// model as specified in [C706] Part 2, but an implementation of this specification
// will not interoperate with an implementation of [C706] Part 2.
//
// This document refers to the Windows implementation of the DCE RPC Location services
// protocol as "LocToLoc".
//
// This document includes the following:
//
// * An extension to provide RPC Location Services functionality in an environment where
// a centrally accessible directory service like Active Directory ( 55ec267c-87d9-4d97-a9d5-5681f5f283b8#gt_e467d927-17bf-49c9-98d1-96ddf61ddd90
// ) directory service is not available. For more details, see Nondirectory mode in
// section 1.3.2 ( dab3af12-2796-4ddf-8313-acba3e7e0dc7 ).
//
// * An extension defining the implementation of the RPC Location Services specification
// in an Active Directory environment. For more details, see Directory-only mode in
// section 1.3.2.
//
// * An extension enabling interoperable RPC Location Service functionality between
// locators running outside an Active Directory environment, and locators running inside
// an Active Directory environment. For more details, see Directory mode in section
// 1.3.2.
//
// * An extension of the syntax for name service entries ( 55ec267c-87d9-4d97-a9d5-5681f5f283b8#gt_6cfe4abe-94bd-43b6-b666-3618d9093373
// ) to include a domain ( 55ec267c-87d9-4d97-a9d5-5681f5f283b8#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca
// ) name. For more details, see section 2.2.2 ( 99becf64-89ed-4bd8-adf0-470add859d15
// ).
//
// * A restriction requiring profile, group, and server attributes to be defined on
// separate name service entries. These attributes are as specified in section "Name
// Service Attributes" in [C706]. For more details, see section 1.3.3 ( e4d192da-2bf6-4649-9f9d-2b85531f614a
// ).
//
// * A restriction requiring clients to be members of an Active Directory domain to
// support persistently storing exported name service entries. For more details, see
// section 1.3.2.
//
// * A restriction requiring clients to be members of an Active Directory domain to
// support profile and group attributes. For more details, see sections 3.2.1 ( 0e5a2e72-15b0-427f-98c7-52da389e734e
// ) , 3.3.3 ( b265805b-4609-4087-a8eb-e583116a912c ).
package rpcl

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
	_ = dtyp.GoPackage
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "rpcl"
)

// Version structure represents RPC_VERSION RPC structure.
type Version struct {
	MajorVersion uint16 `idl:"name:MajorVersion" json:"major_version"`
	MinorVersion uint16 `idl:"name:MinorVersion" json:"minor_version"`
}

func (o *Version) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Version) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	return nil
}
func (o *Version) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	return nil
}

// SyntaxID structure represents RPC_SYNTAX_IDENTIFIER RPC structure.
//
// This structure MUST contain a GUID and version information ([MS-RPCE] section 2.2.2.7).
// It is identical to the RPC_SYNTAX_IDENTIFIER structure used in the LocToLoc interface
// in section 3.1.4. This structure is used to represent the following:
//
// * Identifier and version of an interface.
//
// * Identifier and version of transfer syntax ( 55ec267c-87d9-4d97-a9d5-5681f5f283b8#gt_01216ea7-ac8a-4cc8-9d19-b901bc424c09
// ) for an interface.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SyntaxGUID (16 bytes)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| SyntaxVersion.MajorVersion                                    | SyntaxVersion.MinorVersion                                    |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type SyntaxID struct {
	// SyntaxGUID (16 bytes): As specified in [MS-RPCE] section 2.2.2.7.
	SyntaxGUID    *dtyp.GUID `idl:"name:SyntaxGUID" json:"syntax_guid"`
	SyntaxVersion *Version   `idl:"name:SyntaxVersion" json:"syntax_version"`
}

func (o *SyntaxID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SyntaxID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.SyntaxGUID != nil {
		if err := o.SyntaxGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SyntaxVersion != nil {
		if err := o.SyntaxVersion.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Version{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SyntaxID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.SyntaxGUID == nil {
		o.SyntaxGUID = &dtyp.GUID{}
	}
	if err := o.SyntaxGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SyntaxVersion == nil {
		o.SyntaxVersion = &Version{}
	}
	if err := o.SyntaxVersion.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Context structure represents NSI_NS_HANDLE_T RPC structure.
type Context dcetypes.ContextHandle

func (o *Context) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Binding structure represents NSI_BINDING_T RPC structure.
//
// The NSI_BINDING_T type defines an association of a binding with a server entry.
type Binding struct {
	// string:  A Unicode [UNICODE] string that contains a string binding. For more information,
	// see section "String Bindings" in [C706] Part 2.
	String string `idl:"name:string" json:"string"`
	// entry_name_syntax:  An unsigned 32-bit integer specifying the syntax of the entry_name
	// field. This value MUST be RPC_C_NS_SYNTAX_DCE.
	EntryNameSyntax uint32 `idl:"name:entry_name_syntax" json:"entry_name_syntax"`
	// entry_name:  A Unicode [UNICODE] string specifying the entry name of the name service
	// entry, using the syntax identified by the entry_name_syntax parameter as specified
	// in section 2.2.2.
	EntryName string `idl:"name:entry_name" json:"entry_name"`
}

func (o *Binding) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Binding) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.String != "" {
		_ptr_string := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.String); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_string); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.EntryNameSyntax); err != nil {
		return err
	}
	if o.EntryName != "" {
		_ptr_entry_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryName, _ptr_entry_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Binding) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_string := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.String); err != nil {
			return err
		}
		return nil
	})
	_s_string := func(ptr interface{}) { o.String = *ptr.(*string) }
	if err := w.ReadPointer(&o.String, _s_string, _ptr_string); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntryNameSyntax); err != nil {
		return err
	}
	_ptr_entry_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryName); err != nil {
			return err
		}
		return nil
	})
	_s_entry_name := func(ptr interface{}) { o.EntryName = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryName, _s_entry_name, _ptr_entry_name); err != nil {
		return err
	}
	return nil
}

// BindingVector structure represents NSI_BINDING_VECTOR_T RPC structure.
//
// The NSI_BINDING_VECTOR_T type is defined to hold an array of binding information
// entries.
type BindingVector struct {
	// count:  MUST specify the number of NSI_BINDING_T elements in the binding array.
	Count uint32 `idl:"name:count" json:"count"`
	// binding:  An array of binding information entries.
	Binding []*Binding `idl:"name:binding;size_is:(count)" json:"binding"`
}

func (o *BindingVector) xxx_PreparePayload(ctx context.Context) error {
	if o.Binding != nil && o.Count == 0 {
		o.Count = uint32(len(o.Binding))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindingVector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Count)
	return []uint64{
		dimSize1,
	}
}
func (o *BindingVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	for i1 := range o.Binding {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Binding[i1] != nil {
			if err := o.Binding[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Binding{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Binding); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&Binding{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *BindingVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Count > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Count)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Binding", sizeInfo[0])
	}
	o.Binding = make([]*Binding, sizeInfo[0])
	for i1 := range o.Binding {
		i1 := i1
		if o.Binding[i1] == nil {
			o.Binding[i1] = &Binding{}
		}
		if err := o.Binding[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// UUIDVector structure represents NSI_UUID_VECTOR_T RPC structure.
//
// The NSI_UUID_VECTOR_T type defines an array of NSI_UUID_P_T structures.
type UUIDVector struct {
	// count:  MUST specify the number of NSI_UUID_P_T elements in the uuid member.
	Count uint32 `idl:"name:count" json:"count"`
	// uuid:  An array of NSI_UUID_P_T entries.
	UUID []*dtyp.GUID `idl:"name:uuid;size_is:(count)" json:"uuid"`
}

func (o *UUIDVector) xxx_PreparePayload(ctx context.Context) error {
	if o.UUID != nil && o.Count == 0 {
		o.Count = uint32(len(o.UUID))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UUIDVector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Count)
	return []uint64{
		dimSize1,
	}
}
func (o *UUIDVector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.Count); err != nil {
		return err
	}
	for i1 := range o.UUID {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.UUID[i1] != nil {
			_ptr_uuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UUID[i1] != nil {
					if err := o.UUID[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UUID[i1], _ptr_uuid); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.UUID); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UUIDVector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.Count); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Count > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Count)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.UUID", sizeInfo[0])
	}
	o.UUID = make([]*dtyp.GUID, sizeInfo[0])
	for i1 := range o.UUID {
		i1 := i1
		_ptr_uuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UUID[i1] == nil {
				o.UUID[i1] = &dtyp.GUID{}
			}
			if err := o.UUID[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_uuid := func(ptr interface{}) { o.UUID[i1] = *ptr.(**dtyp.GUID) }
		if err := w.ReadPointer(&o.UUID[i1], _s_uuid, _ptr_uuid); err != nil {
			return err
		}
	}
	return nil
}
