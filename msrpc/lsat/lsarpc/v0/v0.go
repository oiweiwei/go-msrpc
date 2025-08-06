package lsarpc

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
	GoPackage = "lsat"
)

var (
	// Syntax UUID
	LsarpcSyntaxUUID = &uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}}
	// Syntax ID
	LsarpcSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LsarpcSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// lsarpc interface.
type LsarpcClient interface {

	// The LsarClose method frees the resources held by a context handle that was opened
	// earlier. After response, the context handle is no longer usable, and any subsequent
	// uses of this handle MUST fail.
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Lsar_LSA_DP_2 operation.
	// LSADP2

	// Lsar_LSA_DP_3 operation.
	// LSADP3

	// Lsar_LSA_DP_4 operation.
	// LSADP4

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// The LsarOpenPolicy method is exactly the same as LsarOpenPolicy2, except that the
	// SystemName parameter in this method contains only one character instead of a full
	// string. This is because its syntactical definition lacks the [string] RPC annotation
	// present in LsarOpenPolicy2, as specified in [C706]. RPC data types are specified
	// in [MS-RPCE] section 2.2.4.1.
	//
	// The SystemName parameter has no effect on message processing in any environment.
	// It MUST be ignored.
	OpenPolicy(context.Context, *OpenPolicyRequest, ...dcerpc.CallOption) (*OpenPolicyResponse, error)

	// Lsar_LSA_DP_7 operation.
	// LSADP7

	// Lsar_LSA_DP_8 operation.
	// LSADP8

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Lsar_LSA_DP_10 operation.
	// LSADP10

	// Lsar_LSA_DP_11 operation.
	// LSADP11

	// Lsar_LSA_DP_12 operation.
	// LSADP12

	// Lsar_LSA_DP_13 operation.
	// LSADP13

	// The LsarLookupNames method translates a batch of security principal names to their
	// SID form. It also returns the domains that these names are a part of.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED   | Some of the information to be translated has not been translated.   |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED       | None of the information to be translated has been translated.       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupNames(context.Context, *LookupNamesRequest, ...dcerpc.CallOption) (*LookupNamesResponse, error)

	// The LsarLookupSids method translates a batch of security principal SIDs to their
	// name forms. It also returns the domains that these names are a part of.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED   | Some of the information to be translated has not been translated.   |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED       | None of the information to be translated has been translated.       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupSIDs(context.Context, *LookupSIDsRequest, ...dcerpc.CallOption) (*LookupSIDsResponse, error)

	// Lsar_LSA_DP_16 operation.
	// LSADP16

	// Lsar_LSA_DP_17 operation.
	// LSADP17

	// Lsar_LSA_DP_18 operation.
	// LSADP18

	// Lsar_LSA_DP_19 operation.
	// LSADP19

	// Lsar_LSA_DP_20 operation.
	// LSADP20

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// Opnum22NotUsedOnWire operation.
	// Opnum22NotUsedOnWire

	// Lsar_LSA_DP_23 operation.
	// LSADP23

	// Lsar_LSA_DP_24 operation.
	// LSADP24

	// Lsar_LSA_DP_25 operation.
	// LSADP25

	// Lsar_LSA_DP_26 operation.
	// LSADP26

	// Lsar_LSA_DP_27 operation.
	// LSADP27

	// Lsar_LSA_DP_28 operation.
	// LSADP28

	// Lsar_LSA_DP_29 operation.
	// LSADP29

	// Lsar_LSA_DP_30 operation.
	// LSADP30

	// Lsar_LSA_DP_31 operation.
	// LSADP31

	// Lsar_LSA_DP_32 operation.
	// LSADP32

	// Lsar_LSA_DP_33 operation.
	// LSADP33

	// Lsar_LSA_DP_34 operation.
	// LSADP34

	// Lsar_LSA_DP_35 operation.
	// LSADP35

	// Lsar_LSA_DP_36 operation.
	// LSADP36

	// Lsar_LSA_DP_37 operation.
	// LSADP37

	// Lsar_LSA_DP_38 operation.
	// LSADP38

	// Lsar_LSA_DP_39 operation.
	// LSADP39

	// Lsar_LSA_DP_40 operation.
	// LSADP40

	// Lsar_LSA_DP_41 operation.
	// LSADP41

	// Lsar_LSA_DP_42 operation.
	// LSADP42

	// Lsar_LSA_DP_43 operation.
	// LSADP43

	// The LsarOpenPolicy2 method opens a context handle to the RPC server.
	OpenPolicy2(context.Context, *OpenPolicy2Request, ...dcerpc.CallOption) (*OpenPolicy2Response, error)

	// The LsarGetUserName method returns the name and the domain name of the security principal
	// that is invoking the method.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+---------------------------------+---------------------------------------------------------------------+
	//	|             RETURN              |                                                                     |
	//	|           VALUE/CODE            |                             DESCRIPTION                             |
	//	|                                 |                                                                     |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS       | The request was successfully completed.                             |
	//	+---------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED | The caller does not have the permissions to perform this operation. |
	//	+---------------------------------+---------------------------------------------------------------------+
	GetUserName(context.Context, *GetUserNameRequest, ...dcerpc.CallOption) (*GetUserNameResponse, error)

	// Lsar_LSA_DP_46 operation.
	// LSADP46

	// Lsar_LSA_DP_47 operation.
	// LSADP47

	// Lsar_LSA_DP_48 operation.
	// LSADP48

	// Lsar_LSA_DP_49 operation.
	// LSADP49

	// Lsar_LSA_DP_50 operation.
	// LSADP50

	// Lsar_LSA_DP_51 operation.
	// LSADP51

	// Opnum52NotUsedOnWire operation.
	// Opnum52NotUsedOnWire

	// Lsar_LSA_DP_53 operation.
	// LSADP53

	// Lsar_LSA_DP_54 operation.
	// LSADP54

	// Lsar_LSA_DP_55 operation.
	// LSADP55

	// Opnum56NotUsedOnWire operation.
	// Opnum56NotUsedOnWire

	// The LsarLookupSids2 method translates a batch of security principal SIDs to their
	// name forms. It also returns the domains that these names are a part of.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED   | Some of the information to be translated has not been translated.   |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED       | None of the information to be translated has been translated.       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupSids2(context.Context, *LookupSids2Request, ...dcerpc.CallOption) (*LookupSids2Response, error)

	// The LsarLookupNames2 method translates a batch of security principal names to their
	// SID form. It also returns the domains that these names are a part of.<30>
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED   | Some of the information to be translated has not been translated.   |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED       | None of the information to be translated has been translated.       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupNames2(context.Context, *LookupNames2Request, ...dcerpc.CallOption) (*LookupNames2Response, error)

	// Lsar_LSA_DP_59 operation.
	// LSADP59

	// Opnum60NotUsedOnWire operation.
	// Opnum60NotUsedOnWire

	// Opnum61NotUsedOnWire operation.
	// Opnum61NotUsedOnWire

	// Opnum62NotUsedOnWire operation.
	// Opnum62NotUsedOnWire

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// Opnum65NotUsedOnWire operation.
	// Opnum65NotUsedOnWire

	// Opnum66NotUsedOnWire operation.
	// Opnum66NotUsedOnWire

	// Opnum67NotUsedOnWire operation.
	// Opnum67NotUsedOnWire

	// The LsarLookupNames3 method translates a batch of security principal names to their
	// SID form. It also returns the domains that these names are a part of.<28>
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	|               RETURN                |                                                                     |
	//	|             VALUE/CODE              |                             DESCRIPTION                             |
	//	|                                     |                                                                     |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS           | The request was successfully completed.                             |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED   | Some of the information to be translated has not been translated.   |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED     | The caller does not have the permissions to perform this operation. |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER | One of the supplied parameters was invalid.                         |
	//	+-------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED       | None of the information to be translated has been translated.       |
	//	+-------------------------------------+---------------------------------------------------------------------+
	LookupNames3(context.Context, *LookupNames3Request, ...dcerpc.CallOption) (*LookupNames3Response, error)

	// Opnum69NotUsedOnWire operation.
	// Opnum69NotUsedOnWire

	// Opnum70NotUsedOnWire operation.
	// Opnum70NotUsedOnWire

	// Opnum71NotUsedOnWire operation.
	// Opnum71NotUsedOnWire

	// Opnum72NotUsedOnWire operation.
	// Opnum72NotUsedOnWire

	// Lsar_LSA_DP_73 operation.
	// LSADP73

	// Lsar_LSA_DP_74 operation.
	// LSADP74

	// Opnum75NotUsedOnWire operation.
	// Opnum75NotUsedOnWire

	// The LsarLookupSids3 method translates a batch of security principal SIDs to their
	// name forms. It also returns the domains that these names are a part of.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                     |
	//	|               VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                        |                                                                     |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS              | The request was successfully completed.                             |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED      | Some of the information to be translated has not been translated.   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000DC STATUS_INVALID_SERVER_STATE | The RPC server is not a domain controller.                          |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED        | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER    | One of the supplied parameters was invalid.                         |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED          | None of the information to be translated has been translated.       |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//
	// This message is valid only if the RPC server is a domain controller. The RPC server
	// MUST return STATUS_INVALID_SERVER_STATE in the return value if it is not a domain
	// controller.
	LookupSids3(context.Context, *LookupSids3Request, ...dcerpc.CallOption) (*LookupSids3Response, error)

	// The LsarLookupNames4 method translates a batch of security principal names to their
	// SID form. It also returns the domains of which these security principals are a part.
	//
	// Return Values: The following table contains a summary of the return values that an
	// implementation MUST return, as specified by the message processing shown after the
	// table.
	//
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                     |
	//	|               VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                        |                                                                     |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 STATUS_SUCCESS              | The request was successfully completed.                             |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000107 STATUS_SOME_NOT_MAPPED      | Some of the information to be translated has not been translated.   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC00000DC STATUS_INVALID_SERVER_STATE | The RPC server is not a domain controller.                          |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000022 STATUS_ACCESS_DENIED        | The caller does not have the permissions to perform this operation. |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC000000D STATUS_INVALID_PARAMETER    | One of the supplied parameters was invalid.                         |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| 0xC0000073 STATUS_NONE_MAPPED          | None of the information to be translated has been translated.       |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//
	// This message is valid only if the RPC server is a domain controller. The RPC server
	// MUST return STATUS_INVALID_SERVER_STATE in the return value if it is not a domain
	// controller.
	LookupNames4(context.Context, *LookupNames4Request, ...dcerpc.CallOption) (*LookupNames4Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Handle structure represents LSAPR_HANDLE RPC structure.
type Handle dcetypes.ContextHandle

func (o *Handle) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Handle) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
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

// String structure represents STRING RPC structure.
//
// The STRING structure defines a string along with the number of characters in the
// string, as specified in [MS-LSAD] section 2.2.3.1.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.3.1.
type String struct {
	Length        uint16 `idl:"name:Length" json:"length"`
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	Buffer        []byte `idl:"name:Buffer;size_is:(MaximumLength);length_is:(Length)" json:"buffer"`
}

func (o *String) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16(len(o.Buffer))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16(len(o.Buffer))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.MaximumLength > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.MaximumLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.Length)
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
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// ACL structure represents LSAPR_ACL RPC structure.
//
// The LSAPR_ACL structure defines the header of an access control list (ACL) as specified
// in [MS-LSAD] section 2.2.3.2.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.3.2.
type ACL struct {
	ACLRevision uint8  `idl:"name:AclRevision" json:"acl_revision"`
	SBZ1        uint8  `idl:"name:Sbz1" json:"sbz1"`
	ACLSize     uint16 `idl:"name:AclSize" json:"acl_size"`
	_           []byte `idl:"name:Dummy1;size_is:((AclSize-4))"`
}

func (o *ACL) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.ACLSize < 4 {
		o.ACLSize = 4
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ACL) NDRSizeInfo() []uint64 {
	dimSize1 := uint64((o.ACLSize - 4))
	if o.ACLSize < 4 {
		dimSize1 = uint64(0)
	}
	return []uint64{
		dimSize1,
	}
}
func (o *ACL) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(2); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLRevision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.ACLSize); err != nil {
		return err
	}
	// reserved Dummy1
	for i1 := 0; uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ACL) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(2); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLRevision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.ACLSize); err != nil {
		return err
	}
	// reserved Dummy1
	var _Dummy1 []byte
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Dummy1", sizeInfo[0])
	}
	_Dummy1 = make([]byte, sizeInfo[0])
	for i1 := range _Dummy1 {
		i1 := i1
		if err := w.ReadData(&_Dummy1[i1]); err != nil {
			return err
		}
	}
	return nil
}

// SecurityDescriptor structure represents LSAPR_SECURITY_DESCRIPTOR RPC structure.
//
// The LSAPR_SECURITY_DESCRIPTOR structure defines an object's security descriptor as
// specified in [MS-LSAD] section 2.2.3.4.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.3.4.
type SecurityDescriptor struct {
	Revision uint8     `idl:"name:Revision" json:"revision"`
	SBZ1     uint8     `idl:"name:Sbz1" json:"sbz1"`
	Control  uint16    `idl:"name:Control" json:"control"`
	Owner    *dtyp.SID `idl:"name:Owner" json:"owner"`
	Group    *dtyp.SID `idl:"name:Group" json:"group"`
	SACL     *ACL      `idl:"name:Sacl" json:"sacl"`
	DACL     *ACL      `idl:"name:Dacl" json:"dacl"`
}

func (o *SecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Revision); err != nil {
		return err
	}
	if err := w.WriteData(o.SBZ1); err != nil {
		return err
	}
	if err := w.WriteData(o.Control); err != nil {
		return err
	}
	if o.Owner != nil {
		_ptr_Owner := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Owner != nil {
				if err := o.Owner.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Owner, _ptr_Owner); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Group != nil {
		_ptr_Group := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Group != nil {
				if err := o.Group.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.SID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Group, _ptr_Group); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SACL != nil {
		_ptr_Sacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SACL != nil {
				if err := o.SACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SACL, _ptr_Sacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DACL != nil {
		_ptr_Dacl := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DACL != nil {
				if err := o.DACL.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ACL{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DACL, _ptr_Dacl); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Revision); err != nil {
		return err
	}
	if err := w.ReadData(&o.SBZ1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Control); err != nil {
		return err
	}
	_ptr_Owner := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Owner == nil {
			o.Owner = &dtyp.SID{}
		}
		if err := o.Owner.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Owner := func(ptr interface{}) { o.Owner = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.Owner, _s_Owner, _ptr_Owner); err != nil {
		return err
	}
	_ptr_Group := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Group == nil {
			o.Group = &dtyp.SID{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Group := func(ptr interface{}) { o.Group = *ptr.(**dtyp.SID) }
	if err := w.ReadPointer(&o.Group, _s_Group, _ptr_Group); err != nil {
		return err
	}
	_ptr_Sacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SACL == nil {
			o.SACL = &ACL{}
		}
		if err := o.SACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Sacl := func(ptr interface{}) { o.SACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.SACL, _s_Sacl, _ptr_Sacl); err != nil {
		return err
	}
	_ptr_Dacl := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DACL == nil {
			o.DACL = &ACL{}
		}
		if err := o.DACL.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Dacl := func(ptr interface{}) { o.DACL = *ptr.(**ACL) }
	if err := w.ReadPointer(&o.DACL, _s_Dacl, _ptr_Dacl); err != nil {
		return err
	}
	return nil
}

// SecurityImpersonationLevel type represents SECURITY_IMPERSONATION_LEVEL RPC enumeration.
//
// The SECURITY_IMPERSONATION_LEVEL enumeration defines a set of values that specify
// security impersonation levels as specified in [MS-LSAD] section 2.2.3.5.
//
// Individual value semantics are specified in [MS-LSAD] section 2.2.3.5.
type SecurityImpersonationLevel uint16

var (
	SecurityImpersonationLevelAnonymous      SecurityImpersonationLevel = 0
	SecurityImpersonationLevelIdentification SecurityImpersonationLevel = 1
	SecurityImpersonationLevelImpersonation  SecurityImpersonationLevel = 2
	SecurityImpersonationLevelDelegation     SecurityImpersonationLevel = 3
)

func (o SecurityImpersonationLevel) String() string {
	switch o {
	case SecurityImpersonationLevelAnonymous:
		return "SecurityImpersonationLevelAnonymous"
	case SecurityImpersonationLevelIdentification:
		return "SecurityImpersonationLevelIdentification"
	case SecurityImpersonationLevelImpersonation:
		return "SecurityImpersonationLevelImpersonation"
	case SecurityImpersonationLevelDelegation:
		return "SecurityImpersonationLevelDelegation"
	}
	return "Invalid"
}

// SecurityQualityOfService structure represents SECURITY_QUALITY_OF_SERVICE RPC structure.
//
// The SECURITY_QUALITY_OF_SERVICE structure defines information used to support client
// impersonation as specified in [MS-LSAD] section 2.2.3.7.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.3.7.
type SecurityQualityOfService struct {
	Length              uint32                     `idl:"name:Length" json:"length"`
	ImpersonationLevel  SecurityImpersonationLevel `idl:"name:ImpersonationLevel" json:"impersonation_level"`
	ContextTrackingMode uint8                      `idl:"name:ContextTrackingMode" json:"context_tracking_mode"`
	EffectiveOnly       uint8                      `idl:"name:EffectiveOnly" json:"effective_only"`
}

func (o *SecurityQualityOfService) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SecurityQualityOfService) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.ImpersonationLevel)); err != nil {
		return err
	}
	if err := w.WriteData(o.ContextTrackingMode); err != nil {
		return err
	}
	if err := w.WriteData(o.EffectiveOnly); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *SecurityQualityOfService) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.ImpersonationLevel)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ContextTrackingMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.EffectiveOnly); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// ObjectAttributes structure represents LSAPR_OBJECT_ATTRIBUTES RPC structure.
//
// The LSAPR_OBJECT_ATTRIBUTES structure specifies an object and its properties as specified
// in [MS-LSAD] section 2.2.2.4.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.2.4.
type ObjectAttributes struct {
	Length                   uint32                    `idl:"name:Length" json:"length"`
	RootDirectory            []byte                    `idl:"name:RootDirectory" json:"root_directory"`
	ObjectName               *String                   `idl:"name:ObjectName" json:"object_name"`
	Attributes               uint32                    `idl:"name:Attributes" json:"attributes"`
	SecurityDescriptor       *SecurityDescriptor       `idl:"name:SecurityDescriptor" json:"security_descriptor"`
	SecurityQualityOfService *SecurityQualityOfService `idl:"name:SecurityQualityOfService" json:"security_quality_of_service"`
}

func (o *ObjectAttributes) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ObjectAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.RootDirectory != nil {
		_ptr_RootDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(len(o.RootDirectory))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.RootDirectory {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.RootDirectory[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.RootDirectory); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RootDirectory, _ptr_RootDirectory); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ObjectName != nil {
		_ptr_ObjectName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ObjectName != nil {
				if err := o.ObjectName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ObjectName, _ptr_ObjectName); err != nil {
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
	if o.SecurityDescriptor != nil {
		_ptr_SecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecurityDescriptor != nil {
				if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.SecurityQualityOfService != nil {
		_ptr_SecurityQualityOfService := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SecurityQualityOfService != nil {
				if err := o.SecurityQualityOfService.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&SecurityQualityOfService{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityQualityOfService, _ptr_SecurityQualityOfService); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_RootDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.RootDirectory", sizeInfo[0])
		}
		o.RootDirectory = make([]byte, sizeInfo[0])
		for i1 := range o.RootDirectory {
			i1 := i1
			if err := w.ReadData(&o.RootDirectory[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_RootDirectory := func(ptr interface{}) { o.RootDirectory = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.RootDirectory, _s_RootDirectory, _ptr_RootDirectory); err != nil {
		return err
	}
	_ptr_ObjectName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ObjectName == nil {
			o.ObjectName = &String{}
		}
		if err := o.ObjectName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ObjectName := func(ptr interface{}) { o.ObjectName = *ptr.(**String) }
	if err := w.ReadPointer(&o.ObjectName, _s_ObjectName, _ptr_ObjectName); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	_ptr_SecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(**SecurityDescriptor) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_SecurityDescriptor, _ptr_SecurityDescriptor); err != nil {
		return err
	}
	_ptr_SecurityQualityOfService := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.SecurityQualityOfService == nil {
			o.SecurityQualityOfService = &SecurityQualityOfService{}
		}
		if err := o.SecurityQualityOfService.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_SecurityQualityOfService := func(ptr interface{}) { o.SecurityQualityOfService = *ptr.(**SecurityQualityOfService) }
	if err := w.ReadPointer(&o.SecurityQualityOfService, _s_SecurityQualityOfService, _ptr_SecurityQualityOfService); err != nil {
		return err
	}
	return nil
}

// TrustInformation structure represents LSAPR_TRUST_INFORMATION RPC structure.
//
// The LSAPR_TRUST_INFORMATION structure contains information about a domain.
//
// Individual member semantics are specified in [MS-LSAD] section 2.2.7.1.
type TrustInformation struct {
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	SID  *dtyp.SID           `idl:"name:Sid" json:"sid"`
}

func (o *TrustInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TrustInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
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
	return nil
}
func (o *TrustInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
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
	return nil
}

// ReferencedDomainList structure represents LSAPR_REFERENCED_DOMAIN_LIST RPC structure.
//
// The LSAPR_REFERENCED_DOMAIN_LIST structure contains information about the domains
// referenced in a lookup operation.
type ReferencedDomainList struct {
	// Entries:  Contains the number of domains referenced in the lookup operation.
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Domains:  Contains a set of structures that identify domains. If the Entries field
	// in this structure is not 0, this field MUST be non-NULL. If Entries is 0, this field
	// MUST be ignored.
	Domains []*TrustInformation `idl:"name:Domains;size_is:(Entries)" json:"domains"`
	// MaxEntries:  This field MUST be ignored. The content is unspecified, and no requirements
	// are placed on its value since it is never used.
	MaxEntries uint32 `idl:"name:MaxEntries" json:"max_entries"`
}

func (o *ReferencedDomainList) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Domains != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Domains))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *ReferencedDomainList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.Domains != nil || o.Entries > 0 {
		_ptr_Domains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Domains {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Domains[i1] != nil {
					if err := o.Domains[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TrustInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Domains); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TrustInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Domains, _ptr_Domains); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.MaxEntries); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ReferencedDomainList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Domains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Domains", sizeInfo[0])
		}
		o.Domains = make([]*TrustInformation, sizeInfo[0])
		for i1 := range o.Domains {
			i1 := i1
			if o.Domains[i1] == nil {
				o.Domains[i1] = &TrustInformation{}
			}
			if err := o.Domains[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Domains := func(ptr interface{}) { o.Domains = *ptr.(*[]*TrustInformation) }
	if err := w.ReadPointer(&o.Domains, _s_Domains, _ptr_Domains); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxEntries); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// SIDNameUse type represents SID_NAME_USE RPC enumeration.
//
// The SID_NAME_USE enumeration contains values that specify the type of an account.<6>
//
// The SidTypeInvalid and SidTypeComputer enumeration values are not used in this protocol.
// Usage information on the remaining enumeration values is specified in section 3.1.1.
type SIDNameUse uint16

var (
	SIDNameUseTypeUser           SIDNameUse = 1
	SIDNameUseTypeGroup          SIDNameUse = 2
	SIDNameUseTypeDomain         SIDNameUse = 3
	SIDNameUseTypeAlias          SIDNameUse = 4
	SIDNameUseTypeWellKnownGroup SIDNameUse = 5
	SIDNameUseTypeDeletedAccount SIDNameUse = 6
	SIDNameUseTypeInvalid        SIDNameUse = 7
	SIDNameUseTypeUnknown        SIDNameUse = 8
	SIDNameUseTypeComputer       SIDNameUse = 9
	SIDNameUseTypeLabel          SIDNameUse = 10
)

func (o SIDNameUse) String() string {
	switch o {
	case SIDNameUseTypeUser:
		return "SIDNameUseTypeUser"
	case SIDNameUseTypeGroup:
		return "SIDNameUseTypeGroup"
	case SIDNameUseTypeDomain:
		return "SIDNameUseTypeDomain"
	case SIDNameUseTypeAlias:
		return "SIDNameUseTypeAlias"
	case SIDNameUseTypeWellKnownGroup:
		return "SIDNameUseTypeWellKnownGroup"
	case SIDNameUseTypeDeletedAccount:
		return "SIDNameUseTypeDeletedAccount"
	case SIDNameUseTypeInvalid:
		return "SIDNameUseTypeInvalid"
	case SIDNameUseTypeUnknown:
		return "SIDNameUseTypeUnknown"
	case SIDNameUseTypeComputer:
		return "SIDNameUseTypeComputer"
	case SIDNameUseTypeLabel:
		return "SIDNameUseTypeLabel"
	}
	return "Invalid"
}

// TranslatedSID structure represents LSA_TRANSLATED_SID RPC structure.
//
// The LSA_TRANSLATED_SID structure contains information about a security principal
// after translation from a name to a SID. This structure MUST always be accompanied
// by an LSAPR_REFERENCED_DOMAIN_LIST structure when DomainIndex is not -1.
type TranslatedSID struct {
	// Use:  Defines the type of the security principal, as specified in section 2.2.13.
	Use SIDNameUse `idl:"name:Use" json:"use"`
	// RelativeId:  Contains the relative identifier (RID) of the security principal with
	// respect to its domain.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// DomainIndex:  Contains the index into an LSAPR_REFERENCED_DOMAIN_LIST structure that
	// specifies the domain that the security principal is in. A DomainIndex value of -1
	// MUST be used to specify that there are no corresponding domains. Other negative values
	// MUST NOT be returned.
	DomainIndex int32 `idl:"name:DomainIndex" json:"domain_index"`
}

func (o *TranslatedSID) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Use)); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeID); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainIndex); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Use)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainIndex); err != nil {
		return err
	}
	return nil
}

// TranslatedSIDs structure represents LSAPR_TRANSLATED_SIDS RPC structure.
//
// The LSAPR_TRANSLATED_SIDS structure defines a set of translated SIDs.
type TranslatedSIDs struct {
	// Entries:  Contains the number of translated SIDs.<7>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Sids:  Contains a set of structures that contain translated SIDs. If the Entries
	// field in this structure is not 0, this field MUST be non-NULL. If Entries is 0, this
	// field MUST be NULL.
	SIDs []*TranslatedSID `idl:"name:Sids;size_is:(Entries)" json:"sids"`
}

func (o *TranslatedSIDs) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SIDs != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SIDs))
	}
	if o.Entries > uint32(1000) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDs) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.SIDs != nil || o.Entries > 0 {
		_ptr_Sids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDs[i1] != nil {
					if err := o.SIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TranslatedSID{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TranslatedSID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDs, _ptr_Sids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TranslatedSIDs) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Sids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDs", sizeInfo[0])
		}
		o.SIDs = make([]*TranslatedSID, sizeInfo[0])
		for i1 := range o.SIDs {
			i1 := i1
			if o.SIDs[i1] == nil {
				o.SIDs[i1] = &TranslatedSID{}
			}
			if err := o.SIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Sids := func(ptr interface{}) { o.SIDs = *ptr.(*[]*TranslatedSID) }
	if err := w.ReadPointer(&o.SIDs, _s_Sids, _ptr_Sids); err != nil {
		return err
	}
	return nil
}

// LookupLevel type represents LSAP_LOOKUP_LEVEL RPC enumeration.
//
// The LSAP_LOOKUP_LEVEL enumeration defines different scopes for searches during translation.<8>
type LookupLevel uint16

var (
	// LsapLookupWksta: SIDs MUST be searched in the views under the Security Principal
	// SID and Security Principal SID History columns in the following order:
	//
	// * Predefined Translation View, as specified in section 3.1.1.1.1 ( e09da72e-e6c9-4f91-aa64-68b0475719b6
	// ).
	//
	// * Configurable Translation View, as specified in section 3.1.1.1.2 ( 252cb6b0-2edc-4403-9d75-5b2786ba859e
	// ).
	//
	// * Builtin Domain Principal View of the account database on the RPC server ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_ae65dac0-cd24-4e83-a946-6d1097b71553
	// ) , as specified in section 3.1.1.1.3 ( 871df7fe-3470-4b69-9e59-15db70e8d4c6 ).
	//
	// * Account Domain View of account database on that machine, as specified in section
	// 3.1.1.1.6 ( 1777c886-7e47-40a0-979e-da0d76fd6d5c ).
	//
	// * If the machine is not joined to a domain ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca
	// ) , the search ends here.
	//
	// * If the machine is not a domain controller ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_76a05049-3531-4abd-aec8-30e19954b4bd
	// ) : the Account Domain View of the domain to which this machine is joined.
	//
	// * Forest View (section 3.1.1.1.9 ( c69c5834-39c0-430c-9e14-98f3fd7ccfde ) ) of the
	// forest ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_fd104241-4fb3-457c-b2c4-e0c18bb20b62
	// ) of the domain to which this machine is joined, unless ClientRevision is 0x00000001
	// and the machine is joined to a mixed mode domain, as specified in [MS-ADTS] ( ../ms-adts/d2435927-0999-4c62-8c6d-13ba31a52e1a
	// ) section 6.1.4.1 ( ../ms-adts/b40f9606-812d-4b76-924e-55a5401e2bc8 ).
	//
	// * Forest Views of trusted forests ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_3b76a71f-9697-4836-9c69-09899b23c21b
	// ) for the forest of the domain to which this machine is joined, unless ClientRevision
	// is 0x00000001 and the machine is joined to a mixed mode domain, as specified in [MS-ADTS]
	// section 6.1.4.1.
	//
	// * Account Domain Views of externally trusted domains ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_f2f00d47-6cf2-4b32-b8f7-b63e38e2e9c4
	// ) for the domain to which this machine is joined.
	LookupLevelWorkstation LookupLevel = 1
	// LsapLookupPDC: SIDs MUST be searched in the views under the Security Principal SID
	// and Security Principal SID History columns in the following order:
	//
	// * Account Domain View of the domain to which this machine is joined.
	//
	// * Forest View of the forest of the domain to which this machine is joined, unless
	// ClientRevision is 0x00000001 and the machine is joined to a mixed mode domain, as
	// specified in [MS-ADTS] section 6.1.4.1.
	//
	// * Forest Views of trusted forests for the forest of the domain to which this machine
	// is joined, unless ClientRevision is 0x00000001 and the machine is joined to a mixed
	// mode domain, as specified in [MS-ADTS] section 6.1.4.1.
	//
	// * Account Domain Views of externally trusted domains for the domain to which this
	// machine is joined.
	LookupLevelPDC LookupLevel = 2
	// LsapLookupTDL: SIDs MUST be searched in the databases under the Security Principal
	// SID column in the following view:
	//
	// * Account Domain View of the domain NC ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_40a58fa4-953e-4cf3-96c8-57dba60237ef
	// ) for the domain to which this machine is joined.
	LookupLevelTDL LookupLevel = 3
	// LsapLookupGC: SIDs MUST be searched in the databases under the Security Principal
	// SID and Security Principal SID History columns in the following view:
	//
	// * Forest View of the forest of the domain to which this machine is joined.
	LookupLevelGC LookupLevel = 4
	// LsapLookupXForestReferral: SIDs MUST be searched in the databases under the Security
	// Principal SID and Security Principal SID History columns in the following views:
	//
	// * Forest Views of trusted forests for the forest of the domain to which this machine
	// is joined.
	LookupLevelXForestReferral LookupLevel = 5
	// LsapLookupXForestResolve: SIDs MUST be searched in the databases under the Security
	// Principal SID and Security Principal SID History columns in the following view:
	//
	// * Forest View of the forest of the domain to which this machine is joined.
	LookupLevelXForestResolve LookupLevel = 6
	// LsapLookupRODCReferralToFullDC: SIDs MUST be searched in the databases under the
	// Security Principal SID and Security Principal SID History columns in the following
	// order:
	//
	// * Forest Views of trusted forests for the forest of the domain to which this machine
	// is joined, unless ClientRevision is 0x00000001 and the machine is joined to a mixed
	// mode domain, as specified in [MS-ADTS] section 6.1.4.1.
	//
	// * Account Domain Views of externally trusted domains for the domain to which this
	// machine is joined.
	LookupLevelReadOnlyDCReferralToFullDC LookupLevel = 7
)

func (o LookupLevel) String() string {
	switch o {
	case LookupLevelWorkstation:
		return "LookupLevelWorkstation"
	case LookupLevelPDC:
		return "LookupLevelPDC"
	case LookupLevelTDL:
		return "LookupLevelTDL"
	case LookupLevelGC:
		return "LookupLevelGC"
	case LookupLevelXForestReferral:
		return "LookupLevelXForestReferral"
	case LookupLevelXForestResolve:
		return "LookupLevelXForestResolve"
	case LookupLevelReadOnlyDCReferralToFullDC:
		return "LookupLevelReadOnlyDCReferralToFullDC"
	}
	return "Invalid"
}

// SIDInformation structure represents LSAPR_SID_INFORMATION RPC structure.
//
// The LSAPR_SID_INFORMATION structure contains a PRPC_SID value.
type SIDInformation struct {
	// Sid:  Contains the PRPC_SID value, as specified in [MS-DTYP] section 2.4.2.3. This
	// field MUST be non-NULL.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
}

func (o *SIDInformation) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SIDInformation) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
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
	return nil
}
func (o *SIDInformation) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
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
	return nil
}

// SIDEnumBuffer structure represents LSAPR_SID_ENUM_BUFFER RPC structure.
//
// The LSAPR_SID_ENUM_BUFFER structure defines a set of SIDs. This structure is used
// during a translation request for a batch of SIDs to names.
type SIDEnumBuffer struct {
	// Entries:  Contains the number of translated SIDs.<9>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// SidInfo:  Contains a set of structures that contain SIDs, as specified in section
	// 2.2.17. If the Entries field in this structure is not 0, this field MUST be non-NULL.
	// If Entries is 0, this field MUST be ignored.
	SIDInfo []*SIDInformation `idl:"name:SidInfo;size_is:(Entries)" json:"sid_info"`
}

func (o *SIDEnumBuffer) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SIDInfo != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SIDInfo))
	}
	if o.Entries > uint32(20480) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *SIDEnumBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.SIDInfo != nil || o.Entries > 0 {
		_ptr_SidInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDInfo[i1] != nil {
					if err := o.SIDInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&SIDInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDInfo, _ptr_SidInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SIDEnumBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_SidInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDInfo", sizeInfo[0])
		}
		o.SIDInfo = make([]*SIDInformation, sizeInfo[0])
		for i1 := range o.SIDInfo {
			i1 := i1
			if o.SIDInfo[i1] == nil {
				o.SIDInfo[i1] = &SIDInformation{}
			}
			if err := o.SIDInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_SidInfo := func(ptr interface{}) { o.SIDInfo = *ptr.(*[]*SIDInformation) }
	if err := w.ReadPointer(&o.SIDInfo, _s_SidInfo, _ptr_SidInfo); err != nil {
		return err
	}
	return nil
}

// TranslatedName structure represents LSAPR_TRANSLATED_NAME RPC structure.
//
// The LSAPR_TRANSLATED_NAME structure contains information about a security principal,
// along with the human-readable identifier for that security principal. This structure
// MUST always be accompanied by an LSAPR_REFERENCED_DOMAIN_LIST structure when DomainIndex
// is not -1, which contains the domain information for the security principals.
type TranslatedName struct {
	// Use:  Defines the type of the security principal, as specified in section 2.2.13.
	Use SIDNameUse `idl:"name:Use" json:"use"`
	// Name:  Contains the name of the security principal, with syntaxes described in section
	// 3.1.4.5. The RPC_UNICODE_STRING structure is defined in [MS-DTYP] section 2.3.10.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// DomainIndex:  Contains the index into the corresponding LSAPR_REFERENCED_DOMAIN_LIST
	// structure that specifies the domain that the security principal is in. A DomainIndex
	// value of -1 MUST be used to specify that there are no corresponding domains. Other
	// negative values MUST NOT be used.
	DomainIndex int32 `idl:"name:DomainIndex" json:"domain_index"`
}

func (o *TranslatedName) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedName) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Use)); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DomainIndex); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedName) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Use)); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainIndex); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TranslatedNames structure represents LSAPR_TRANSLATED_NAMES RPC structure.
//
// The LSAPR_TRANSLATED_NAMES structure defines a set of translated names. This is used
// in the response to a translation request from SIDs to names.
type TranslatedNames struct {
	// Entries:  Contains the number of translated names.<10>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Names:  Contains a set of translated names, as specified in section 2.2.19. If the
	// Entries field in this structure is not 0, this field MUST be non-NULL. If Entries
	// is 0, this field MUST be ignored.
	Names []*TranslatedName `idl:"name:Names;size_is:(Entries)" json:"names"`
}

func (o *TranslatedNames) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Names != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Names))
	}
	if o.Entries > uint32(20480) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedNames) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.Names != nil || o.Entries > 0 {
		_ptr_Names := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Names {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Names[i1] != nil {
					if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TranslatedName{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TranslatedName{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Names, _ptr_Names); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TranslatedNames) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Names := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*TranslatedName, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &TranslatedName{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Names := func(ptr interface{}) { o.Names = *ptr.(*[]*TranslatedName) }
	if err := w.ReadPointer(&o.Names, _s_Names, _ptr_Names); err != nil {
		return err
	}
	return nil
}

// TranslatedNameEx structure represents LSAPR_TRANSLATED_NAME_EX RPC structure.
//
// The LSAPR_TRANSLATED_NAME_EX structure contains information about a security principal
// along with the human-readable identifier for that security principal. This structure
// MUST always be accompanied by an LSAPR_REFERENCED_DOMAIN_LIST structure when DomainIndex
// is not -1, which contains the domain information for the security principals.
type TranslatedNameEx struct {
	// Use:  Defines the type of the security principal, as specified in section 2.2.13.
	Use SIDNameUse `idl:"name:Use" json:"use"`
	// Name:  Contains the name of the security principal. The RPC_UNICODE_STRING structure
	// is defined in [MS-DTYP] section 2.3.10.
	Name *dtyp.UnicodeString `idl:"name:Name" json:"name"`
	// DomainIndex:  Contains the index into the corresponding LSAPR_REFERENCED_DOMAIN_LIST
	// structure that specifies the domain that the security principal is in. A DomainIndex
	// value of -1 MUST be used to specify that there are no corresponding domains. Other
	// negative values MUST NOT be used.
	DomainIndex int32 `idl:"name:DomainIndex" json:"domain_index"`
	// Flags:  Contains bitmapped values that define the properties of this translation.
	// The value MUST be the logical OR of zero or more of the following flags. These flags
	// communicate the following additional information about how the SID was resolved.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The SID was not found by matching against the security principal SID property.   |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The SID might be found by traversing a forest trust.                             |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The SID was found by matching against the last database view, defined in section |
	//	|            | 3.1.1.1.1.                                                                       |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *TranslatedNameEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedNameEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Use)); err != nil {
		return err
	}
	if o.Name != nil {
		if err := o.Name.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DomainIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedNameEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Use)); err != nil {
		return err
	}
	if o.Name == nil {
		o.Name = &dtyp.UnicodeString{}
	}
	if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TranslatedNamesEx structure represents LSAPR_TRANSLATED_NAMES_EX RPC structure.
//
// The LSAPR_TRANSLATED_NAMES_EX structure contains a set of translated names.
type TranslatedNamesEx struct {
	// Entries:  Contains the number of translated names.<12>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Names:  Contains a set of structures that contain translated names, as specified
	// in section 2.2.21. If the Entries field in this structure is not 0, this field MUST
	// be non-NULL. If Entries is 0, this field MUST be ignored.
	Names []*TranslatedNameEx `idl:"name:Names;size_is:(Entries)" json:"names"`
}

func (o *TranslatedNamesEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Names != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.Names))
	}
	if o.Entries > uint32(20480) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedNamesEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.Names != nil || o.Entries > 0 {
		_ptr_Names := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Names {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Names[i1] != nil {
					if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TranslatedNameEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TranslatedNameEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Names, _ptr_Names); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TranslatedNamesEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Names := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*TranslatedNameEx, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &TranslatedNameEx{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Names := func(ptr interface{}) { o.Names = *ptr.(*[]*TranslatedNameEx) }
	if err := w.ReadPointer(&o.Names, _s_Names, _ptr_Names); err != nil {
		return err
	}
	return nil
}

// TranslatedSIDEx structure represents LSAPR_TRANSLATED_SID_EX RPC structure.
//
// The LSAPR_TRANSLATED_SID_EX structure contains information about a security principal
// after it has been translated into a SID. This structure MUST always be accompanied
// by an LSAPR_REFERENCED_DOMAIN_LIST structure when DomainIndex is not -1, which contains
// the domain information for the security principal.
//
// This structure differs from LSA_TRANSLATED_SID only in that a new Flags field is
// added.
type TranslatedSIDEx struct {
	// Use:  Defines the type of the security principal, as specified in section 2.2.13.
	Use SIDNameUse `idl:"name:Use" json:"use"`
	// RelativeId:  Contains the relative identifier (RID) of the security principal with
	// respect to its domain.
	RelativeID uint32 `idl:"name:RelativeId" json:"relative_id"`
	// DomainIndex:  Contains the index into the corresponding LSAPR_REFERENCED_DOMAIN_LIST
	// structure that specifies the domain that the security principal is in. A DomainIndex
	// value of -1 MUST be used to specify that there are no corresponding domains. Other
	// negative values MUST NOT be used.
	DomainIndex int32 `idl:"name:DomainIndex" json:"domain_index"`
	// Flags:  Contains bitmapped values that define the properties of this translation.
	// The value MUST be the logical OR of zero or more of the following flags. These flags
	// communicate additional information on how the name was resolved.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The name was not found by matching against the Security Principal Name property. |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The name might be found by traversing a forest trust.                            |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The name was found by matching against the last database view, as defined in     |
	//	|            | section 3.1.1.1.1.                                                               |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *TranslatedSIDEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Use)); err != nil {
		return err
	}
	if err := w.WriteData(o.RelativeID); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Use)); err != nil {
		return err
	}
	if err := w.ReadData(&o.RelativeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// TranslatedSIDsEx structure represents LSAPR_TRANSLATED_SIDS_EX RPC structure.
//
// The LSAPR_TRANSLATED_SIDS_EX structure contains a set of translated SIDs.
type TranslatedSIDsEx struct {
	// Entries:  Contains the number of translated SIDs.<14>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Sids:  Contains a set of structures that contain translated SIDs, as specified in
	// section 2.2.23. If the Entries field in this structure is not 0, this field MUST
	// be non-NULL. If Entries is 0, this field MUST be NULL.
	SIDs []*TranslatedSIDEx `idl:"name:Sids;size_is:(Entries)" json:"sids"`
}

func (o *TranslatedSIDsEx) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SIDs != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SIDs))
	}
	if o.Entries > uint32(1000) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDsEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.SIDs != nil || o.Entries > 0 {
		_ptr_Sids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDs[i1] != nil {
					if err := o.SIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TranslatedSIDEx{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TranslatedSIDEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDs, _ptr_Sids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TranslatedSIDsEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Sids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDs", sizeInfo[0])
		}
		o.SIDs = make([]*TranslatedSIDEx, sizeInfo[0])
		for i1 := range o.SIDs {
			i1 := i1
			if o.SIDs[i1] == nil {
				o.SIDs[i1] = &TranslatedSIDEx{}
			}
			if err := o.SIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Sids := func(ptr interface{}) { o.SIDs = *ptr.(*[]*TranslatedSIDEx) }
	if err := w.ReadPointer(&o.SIDs, _s_Sids, _ptr_Sids); err != nil {
		return err
	}
	return nil
}

// TranslatedSIDEx2 structure represents LSAPR_TRANSLATED_SID_EX2 RPC structure.
//
// The LSAPR_TRANSLATED_SID_EX2 structure contains information about a security principal
// after it has been translated into a SID. This structure MUST always be accompanied
// by an LSAPR_REFERENCED_DOMAIN_LIST structure when DomainIndex is not -1.
//
// This structure differs from LSAPR_TRANSLATED_SID_EX only in that a SID is returned
// instead of a RID.
type TranslatedSIDEx2 struct {
	// Use:  Defines the type of the security principal, as specified in section 2.2.13.
	Use SIDNameUse `idl:"name:Use" json:"use"`
	// Sid:  Contains the SID ([MS-DTYP] section 2.4.2.3) of the security principal. This
	// field MUST be treated as a [ref] pointer and hence MUST be non-NULL.
	SID *dtyp.SID `idl:"name:Sid" json:"sid"`
	// DomainIndex:  Contains the index into an LSAPR_REFERENCED_DOMAIN_LIST structure that
	// specifies the domain that the security principal is in. A DomainIndex value of -1
	// MUST be used to specify that there are no corresponding domains. Other negative values
	// MUST NOT be used.
	DomainIndex int32 `idl:"name:DomainIndex" json:"domain_index"`
	// Flags:  Contains bitmapped values that define the properties of this translation.
	// The value MUST be the logical OR of zero or more of the following flags. These flags
	// communicate additional information on how the name was resolved.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The name was not found by matching against the Security Principal Name property. |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The name might be found by traversing a forest trust.                            |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 | The name was found by matching against the last database view (see section       |
	//	|            | 3.1.1.1.1).                                                                      |
	//	+------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *TranslatedSIDEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteEnum(uint16(o.Use)); err != nil {
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
	if err := w.WriteData(o.DomainIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint16)(&o.Use)); err != nil {
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
	if err := w.ReadData(&o.DomainIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TranslatedSIDsEx2 structure represents LSAPR_TRANSLATED_SIDS_EX2 RPC structure.
//
// The LSAPR_TRANSLATED_SIDS_EX2 structure contains a set of translated SIDs.
type TranslatedSIDsEx2 struct {
	// Entries:  Contains the number of translated SIDs.<16>
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// Sids:  Contains a set of structures that define translated SIDs, as specified in
	// section 2.2.25. If the Entries field in this structure is not 0, this field MUST
	// be non-NULL. If Entries is 0, this field MUST be NULL.
	SIDs []*TranslatedSIDEx2 `idl:"name:Sids;size_is:(Entries)" json:"sids"`
}

func (o *TranslatedSIDsEx2) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SIDs != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SIDs))
	}
	if o.Entries > uint32(1000) {
		return fmt.Errorf("Entries is out of range")
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *TranslatedSIDsEx2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.SIDs != nil || o.Entries > 0 {
		_ptr_Sids := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDs[i1] != nil {
					if err := o.SIDs[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&TranslatedSIDEx2{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&TranslatedSIDEx2{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDs, _ptr_Sids); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TranslatedSIDsEx2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_Sids := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDs", sizeInfo[0])
		}
		o.SIDs = make([]*TranslatedSIDEx2, sizeInfo[0])
		for i1 := range o.SIDs {
			i1 := i1
			if o.SIDs[i1] == nil {
				o.SIDs[i1] = &TranslatedSIDEx2{}
			}
			if err := o.SIDs[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Sids := func(ptr interface{}) { o.SIDs = *ptr.(*[]*TranslatedSIDEx2) }
	if err := w.ReadPointer(&o.SIDs, _s_Sids, _ptr_Sids); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultLsarpcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultLsarpcClient) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenPolicy(ctx context.Context, in *OpenPolicyRequest, opts ...dcerpc.CallOption) (*OpenPolicyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPolicyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupNames(ctx context.Context, in *LookupNamesRequest, opts ...dcerpc.CallOption) (*LookupNamesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupSIDs(ctx context.Context, in *LookupSIDsRequest, opts ...dcerpc.CallOption) (*LookupSIDsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupSIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) OpenPolicy2(ctx context.Context, in *OpenPolicy2Request, opts ...dcerpc.CallOption) (*OpenPolicy2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPolicy2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) GetUserName(ctx context.Context, in *GetUserNameRequest, opts ...dcerpc.CallOption) (*GetUserNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUserNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupSids2(ctx context.Context, in *LookupSids2Request, opts ...dcerpc.CallOption) (*LookupSids2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupSids2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupNames2(ctx context.Context, in *LookupNames2Request, opts ...dcerpc.CallOption) (*LookupNames2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNames2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupNames3(ctx context.Context, in *LookupNames3Request, opts ...dcerpc.CallOption) (*LookupNames3Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNames3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupSids3(ctx context.Context, in *LookupSids3Request, opts ...dcerpc.CallOption) (*LookupSids3Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupSids3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) LookupNames4(ctx context.Context, in *LookupNames4Request, opts ...dcerpc.CallOption) (*LookupNames4Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LookupNames4Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLsarpcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLsarpcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewLsarpcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LsarpcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LsarpcSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultLsarpcClient{cc: cc}, nil
}

// xxx_CloseOperation structure represents the LsarClose operation
type xxx_CloseOperation struct {
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	Return int32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 0 }

func (o *xxx_CloseOperation) OpName() string { return "/lsarpc/v0/LsarClose" }

func (o *xxx_CloseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object != nil {
			if err := o.Object.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ObjectHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Object == nil {
			o.Object = &Handle{}
		}
		if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseRequest structure represents the LsarClose operation request
type CloseRequest struct {
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.Object = o.Object
	return op
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
}
func (o *CloseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResponse structure represents the LsarClose operation response
type CloseResponse struct {
	Object *Handle `idl:"name:ObjectHandle" json:"object"`
	// Return: The LsarClose return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.Object = o.Object
	op.Return = o.Return
	return op
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.Object = op.Object
	o.Return = op.Return
}
func (o *CloseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPolicyOperation structure represents the LsarOpenPolicy operation
type xxx_OpenPolicyOperation struct {
	SystemName       string            `idl:"name:SystemName;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
	Policy           *Handle           `idl:"name:PolicyHandle" json:"policy"`
	Return           int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPolicyOperation) OpNum() int { return 6 }

func (o *xxx_OpenPolicyOperation) OpName() string { return "/lsarpc/v0/LsarOpenPolicy" }

func (o *xxx_OpenPolicyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SystemName {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		if o.SystemName != "" {
			_ptr_SystemName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.SystemName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemName, _ptr_SystemName); err != nil {
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
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes != nil {
			if err := o.ObjectAttributes.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SystemName {in} (1:{pointer=unique}*(1)[dim:0,string](wchar))
	{
		_ptr_SystemName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.SystemName); err != nil {
				return err
			}
			return nil
		})
		_s_SystemName := func(ptr interface{}) { o.SystemName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SystemName, _s_SystemName, _ptr_SystemName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES,pointer=ref}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes == nil {
			o.ObjectAttributes = &ObjectAttributes{}
		}
		if err := o.ObjectAttributes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenPolicyRequest structure represents the LsarOpenPolicy operation request
type OpenPolicyRequest struct {
	SystemName       string            `idl:"name:SystemName;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenPolicyRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenPolicyOperation) *xxx_OpenPolicyOperation {
	if op == nil {
		op = &xxx_OpenPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.SystemName = o.SystemName
	op.ObjectAttributes = o.ObjectAttributes
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenPolicyRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicyOperation) {
	if o == nil {
		return
	}
	o.SystemName = op.SystemName
	o.ObjectAttributes = op.ObjectAttributes
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPolicyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPolicyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPolicyResponse structure represents the LsarOpenPolicy operation response
type OpenPolicyResponse struct {
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Return: The LsarOpenPolicy return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenPolicyResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenPolicyOperation) *xxx_OpenPolicyOperation {
	if op == nil {
		op = &xxx_OpenPolicyOperation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.Return = o.Return
	return op
}

func (o *OpenPolicyResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicyOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *OpenPolicyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPolicyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNamesOperation structure represents the LsarLookupNames operation
type xxx_LookupNamesOperation struct {
	Policy            *Handle               `idl:"name:PolicyHandle" json:"policy"`
	Count             uint32                `idl:"name:Count" json:"count"`
	Names             []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedSIDs    *TranslatedSIDs       `idl:"name:TranslatedSids" json:"translated_sids"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNamesOperation) OpNum() int { return 14 }

func (o *xxx_LookupNamesOperation) OpName() string { return "/lsarpc/v0/LsarLookupNames" }

func (o *xxx_LookupNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.Count == 0 {
		o.Count = uint32(len(o.Names))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDs{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDs{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDs{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDs{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupNamesRequest structure represents the LsarLookupNames operation request
type LookupNamesRequest struct {
	// PolicyHandle: Context handle obtained by an LsarOpenPolicy or LsarOpenPolicy2 call.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Count: Number of names in the Names array.<32>
	Count uint32 `idl:"name:Count" json:"count"`
	// Names: Contains the security principal names to translate, as specified in section
	// 3.1.4.5.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDs `idl:"name:TranslatedSids" json:"translated_sids"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their SID forms. This parameter has no effect on message processing
	// in any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
}

func (o *LookupNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupNamesOperation) *xxx_LookupNamesOperation {
	if op == nil {
		op = &xxx_LookupNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.Count = o.Count
	op.Names = o.Names
	op.TranslatedSIDs = o.TranslatedSIDs
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	return op
}

func (o *LookupNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupNamesOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Count = op.Count
	o.Names = op.Names
	o.TranslatedSIDs = op.TranslatedSIDs
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
}
func (o *LookupNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNamesResponse structure represents the LsarLookupNames operation response
type LookupNamesResponse struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDs `idl:"name:TranslatedSids" json:"translated_sids"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their SID forms. This parameter has no effect on message processing
	// in any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupNamesOperation) *xxx_LookupNamesOperation {
	if op == nil {
		op = &xxx_LookupNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedSIDs = o.TranslatedSIDs
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupNamesOperation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedSIDs = op.TranslatedSIDs
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupSIDsOperation structure represents the LsarLookupSids operation
type xxx_LookupSIDsOperation struct {
	Policy            *Handle               `idl:"name:PolicyHandle" json:"policy"`
	SIDEnumBuffer     *SIDEnumBuffer        `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedNames   *TranslatedNames      `idl:"name:TranslatedNames" json:"translated_names"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupSIDsOperation) OpNum() int { return 15 }

func (o *xxx_LookupSIDsOperation) OpName() string { return "/lsarpc/v0/LsarLookupSids" }

func (o *xxx_LookupSIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer != nil {
			if err := o.SIDEnumBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNames{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer == nil {
			o.SIDEnumBuffer = &SIDEnumBuffer{}
		}
		if err := o.SIDEnumBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNames{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNames{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNames{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupSIDsRequest structure represents the LsarLookupSids operation request
type LookupSIDsRequest struct {
	// PolicyHandle: Context handle obtained by an LsarOpenPolicy or LsarOpenPolicy2 call.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// SidEnumBuffer: Contains the SIDs to be translated. The SIDs in this structure can
	// be that of users, groups, computers, Windows-defined well-known security principals,
	// or domains.
	SIDEnumBuffer *SIDEnumBuffer `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	// TranslatedNames: On successful return, contains the corresponding name form for security
	// principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNames `idl:"name:TranslatedNames" json:"translated_names"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their Name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
}

func (o *LookupSIDsRequest) xxx_ToOp(ctx context.Context, op *xxx_LookupSIDsOperation) *xxx_LookupSIDsOperation {
	if op == nil {
		op = &xxx_LookupSIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.SIDEnumBuffer = o.SIDEnumBuffer
	op.TranslatedNames = o.TranslatedNames
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	return op
}

func (o *LookupSIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_LookupSIDsOperation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.SIDEnumBuffer = op.SIDEnumBuffer
	o.TranslatedNames = op.TranslatedNames
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
}
func (o *LookupSIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupSIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupSIDsResponse structure represents the LsarLookupSids operation response
type LookupSIDsResponse struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedNames: On successful return, contains the corresponding name form for security
	// principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNames `idl:"name:TranslatedNames" json:"translated_names"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their Name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupSids return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupSIDsResponse) xxx_ToOp(ctx context.Context, op *xxx_LookupSIDsOperation) *xxx_LookupSIDsOperation {
	if op == nil {
		op = &xxx_LookupSIDsOperation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedNames = o.TranslatedNames
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupSIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_LookupSIDsOperation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedNames = op.TranslatedNames
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupSIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupSIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPolicy2Operation structure represents the LsarOpenPolicy2 operation
type xxx_OpenPolicy2Operation struct {
	SystemName       string            `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
	Policy           *Handle           `idl:"name:PolicyHandle" json:"policy"`
	Return           int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPolicy2Operation) OpNum() int { return 44 }

func (o *xxx_OpenPolicy2Operation) OpName() string { return "/lsarpc/v0/LsarOpenPolicy2" }

func (o *xxx_OpenPolicy2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SystemName != "" {
			_ptr_SystemName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SystemName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemName, _ptr_SystemName); err != nil {
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
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes != nil {
			if err := o.ObjectAttributes.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ObjectAttributes{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_SystemName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SystemName); err != nil {
				return err
			}
			return nil
		})
		_s_SystemName := func(ptr interface{}) { o.SystemName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SystemName, _s_SystemName, _ptr_SystemName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ObjectAttributes {in} (1:{alias=PLSAPR_OBJECT_ATTRIBUTES,pointer=ref}*(1))(2:{alias=LSAPR_OBJECT_ATTRIBUTES}(struct))
	{
		if o.ObjectAttributes == nil {
			o.ObjectAttributes = &ObjectAttributes{}
		}
		if err := o.ObjectAttributes.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DesiredAccess {in} (1:{alias=ACCESS_MASK, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPolicy2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenPolicy2Request structure represents the LsarOpenPolicy2 operation request
type OpenPolicy2Request struct {
	SystemName       string            `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	ObjectAttributes *ObjectAttributes `idl:"name:ObjectAttributes" json:"object_attributes"`
	DesiredAccess    uint32            `idl:"name:DesiredAccess" json:"desired_access"`
}

func (o *OpenPolicy2Request) xxx_ToOp(ctx context.Context, op *xxx_OpenPolicy2Operation) *xxx_OpenPolicy2Operation {
	if op == nil {
		op = &xxx_OpenPolicy2Operation{}
	}
	if o == nil {
		return op
	}
	op.SystemName = o.SystemName
	op.ObjectAttributes = o.ObjectAttributes
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenPolicy2Request) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicy2Operation) {
	if o == nil {
		return
	}
	o.SystemName = op.SystemName
	o.ObjectAttributes = op.ObjectAttributes
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPolicy2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPolicy2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicy2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPolicy2Response structure represents the LsarOpenPolicy2 operation response
type OpenPolicy2Response struct {
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Return: The LsarOpenPolicy2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenPolicy2Response) xxx_ToOp(ctx context.Context, op *xxx_OpenPolicy2Operation) *xxx_OpenPolicy2Operation {
	if op == nil {
		op = &xxx_OpenPolicy2Operation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.Return = o.Return
	return op
}

func (o *OpenPolicy2Response) xxx_FromOp(ctx context.Context, op *xxx_OpenPolicy2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Return = op.Return
}
func (o *OpenPolicy2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPolicy2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPolicy2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserNameOperation structure represents the LsarGetUserName operation
type xxx_GetUserNameOperation struct {
	SystemName string              `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	UserName   *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	DomainName *dtyp.UnicodeString `idl:"name:DomainName;pointer:unique" json:"domain_name"`
	Return     int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserNameOperation) OpNum() int { return 45 }

func (o *xxx_GetUserNameOperation) OpName() string { return "/lsarpc/v0/LsarGetUserName" }

func (o *xxx_GetUserNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.SystemName != "" {
			_ptr_SystemName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.SystemName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.SystemName, _ptr_SystemName); err != nil {
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
	// UserName {in, out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.UserName != nil {
			_ptr_UserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserName != nil {
					if err := o.UserName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_UserName); err != nil {
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
	// DomainName {in, out} (1:{pointer=unique}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.DomainName != nil {
			_ptr_DomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainName != nil {
					_ptr_DomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.DomainName != nil {
							if err := o.DomainName.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.DomainName, _ptr_DomainName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_DomainName); err != nil {
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

func (o *xxx_GetUserNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SystemName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_SystemName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.SystemName); err != nil {
				return err
			}
			return nil
		})
		_s_SystemName := func(ptr interface{}) { o.SystemName = *ptr.(*string) }
		if err := w.ReadPointer(&o.SystemName, _s_SystemName, _ptr_SystemName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserName {in, out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_UserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserName == nil {
				o.UserName = &dtyp.UnicodeString{}
			}
			if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserName := func(ptr interface{}) { o.UserName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.UserName, _s_UserName, _ptr_UserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainName {in, out} (1:{pointer=unique}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_DomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_DomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.DomainName == nil {
					o.DomainName = &dtyp.UnicodeString{}
				}
				if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_DomainName := func(ptr interface{}) { o.DomainName = *ptr.(**dtyp.UnicodeString) }
			if err := w.ReadPointer(&o.DomainName, _s_DomainName, _ptr_DomainName); err != nil {
				return err
			}
			return nil
		})
		_s_DomainName := func(ptr interface{}) { o.DomainName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.DomainName, _s_DomainName, _ptr_DomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserName {in, out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.UserName != nil {
			_ptr_UserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserName != nil {
					if err := o.UserName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_UserName); err != nil {
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
	// DomainName {in, out} (1:{pointer=unique}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.DomainName != nil {
			_ptr_DomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DomainName != nil {
					_ptr_DomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.DomainName != nil {
							if err := o.DomainName.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.DomainName, _ptr_DomainName); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_DomainName); err != nil {
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
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserName {in, out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_UserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserName == nil {
				o.UserName = &dtyp.UnicodeString{}
			}
			if err := o.UserName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_UserName := func(ptr interface{}) { o.UserName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.UserName, _s_UserName, _ptr_UserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainName {in, out} (1:{pointer=unique}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_DomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_DomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.DomainName == nil {
					o.DomainName = &dtyp.UnicodeString{}
				}
				if err := o.DomainName.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_DomainName := func(ptr interface{}) { o.DomainName = *ptr.(**dtyp.UnicodeString) }
			if err := w.ReadPointer(&o.DomainName, _s_DomainName, _ptr_DomainName); err != nil {
				return err
			}
			return nil
		})
		_s_DomainName := func(ptr interface{}) { o.DomainName = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.DomainName, _s_DomainName, _ptr_DomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetUserNameRequest structure represents the LsarGetUserName operation request
type GetUserNameRequest struct {
	// SystemName: This parameter has no effect on message processing in any environment.
	// It MUST be ignored.
	SystemName string `idl:"name:SystemName;string;pointer:unique" json:"system_name"`
	// UserName: On return, contains the name of the security principal that is making the
	// call. The string MUST be of the form sAMAccountName. On input, this parameter MUST
	// be ignored. The RPC_UNICODE_STRING structure is defined in [MS-DTYP] section 2.3.10.
	UserName *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	// DomainName: On return, contains the domain name of the security principal that is
	// invoking the method. This string MUST be a NetBIOS name. On input, this parameter
	// MUST be ignored.
	DomainName *dtyp.UnicodeString `idl:"name:DomainName;pointer:unique" json:"domain_name"`
}

func (o *GetUserNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetUserNameOperation) *xxx_GetUserNameOperation {
	if op == nil {
		op = &xxx_GetUserNameOperation{}
	}
	if o == nil {
		return op
	}
	op.SystemName = o.SystemName
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	return op
}

func (o *GetUserNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserNameOperation) {
	if o == nil {
		return
	}
	o.SystemName = op.SystemName
	o.UserName = op.UserName
	o.DomainName = op.DomainName
}
func (o *GetUserNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetUserNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserNameResponse structure represents the LsarGetUserName operation response
type GetUserNameResponse struct {
	// UserName: On return, contains the name of the security principal that is making the
	// call. The string MUST be of the form sAMAccountName. On input, this parameter MUST
	// be ignored. The RPC_UNICODE_STRING structure is defined in [MS-DTYP] section 2.3.10.
	UserName *dtyp.UnicodeString `idl:"name:UserName" json:"user_name"`
	// DomainName: On return, contains the domain name of the security principal that is
	// invoking the method. This string MUST be a NetBIOS name. On input, this parameter
	// MUST be ignored.
	DomainName *dtyp.UnicodeString `idl:"name:DomainName;pointer:unique" json:"domain_name"`
	// Return: The LsarGetUserName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetUserNameOperation) *xxx_GetUserNameOperation {
	if op == nil {
		op = &xxx_GetUserNameOperation{}
	}
	if o == nil {
		return op
	}
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	op.Return = o.Return
	return op
}

func (o *GetUserNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserNameOperation) {
	if o == nil {
		return
	}
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.Return = op.Return
}
func (o *GetUserNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetUserNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupSids2Operation structure represents the LsarLookupSids2 operation
type xxx_LookupSids2Operation struct {
	Policy            *Handle               `idl:"name:PolicyHandle" json:"policy"`
	SIDEnumBuffer     *SIDEnumBuffer        `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedNames   *TranslatedNamesEx    `idl:"name:TranslatedNames" json:"translated_names"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	LookupOptions     uint32                `idl:"name:LookupOptions" json:"lookup_options"`
	ClientRevision    uint32                `idl:"name:ClientRevision" json:"client_revision"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupSids2Operation) OpNum() int { return 57 }

func (o *xxx_LookupSids2Operation) OpName() string { return "/lsarpc/v0/LsarLookupSids2" }

func (o *xxx_LookupSids2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer != nil {
			if err := o.SIDEnumBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNamesEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer == nil {
			o.SIDEnumBuffer = &SIDEnumBuffer{}
		}
		if err := o.SIDEnumBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNamesEx{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNamesEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNamesEx{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupSids2Request structure represents the LsarLookupSids2 operation request
type LookupSids2Request struct {
	// PolicyHandle: Context handle obtained by an LsarOpenPolicy or LsarOpenPolicy2 call.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// SidEnumBuffer: Contains the SIDs to be translated. The SIDs in this structure can
	// be that of users, groups, computers, Windows-defined well-known security principals,
	// or domains.
	SIDEnumBuffer *SIDEnumBuffer `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	// TranslatedNames: On successful return, contains the corresponding name forms for
	// security principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNamesEx `idl:"name:TranslatedNames" json:"translated_names"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On return, contains the number of names that are translated completely
	// to their name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// LookupOptions: Flags that control the lookup operation. This parameter is reserved
	// for future use and SHOULD<33> be set to 0.
	LookupOptions uint32 `idl:"name:LookupOptions" json:"lookup_options"`
	// ClientRevision: Version of the client, which implies the client's capabilities. For
	// possible values and their meanings, see section 3.1.4.5.
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
}

func (o *LookupSids2Request) xxx_ToOp(ctx context.Context, op *xxx_LookupSids2Operation) *xxx_LookupSids2Operation {
	if op == nil {
		op = &xxx_LookupSids2Operation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.SIDEnumBuffer = o.SIDEnumBuffer
	op.TranslatedNames = o.TranslatedNames
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	op.LookupOptions = o.LookupOptions
	op.ClientRevision = o.ClientRevision
	return op
}

func (o *LookupSids2Request) xxx_FromOp(ctx context.Context, op *xxx_LookupSids2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.SIDEnumBuffer = op.SIDEnumBuffer
	o.TranslatedNames = op.TranslatedNames
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
	o.LookupOptions = op.LookupOptions
	o.ClientRevision = op.ClientRevision
}
func (o *LookupSids2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupSids2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSids2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupSids2Response structure represents the LsarLookupSids2 operation response
type LookupSids2Response struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedNames: On successful return, contains the corresponding name forms for
	// security principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNamesEx `idl:"name:TranslatedNames" json:"translated_names"`
	// MappedCount: On return, contains the number of names that are translated completely
	// to their name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupSids2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupSids2Response) xxx_ToOp(ctx context.Context, op *xxx_LookupSids2Operation) *xxx_LookupSids2Operation {
	if op == nil {
		op = &xxx_LookupSids2Operation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedNames = o.TranslatedNames
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupSids2Response) xxx_FromOp(ctx context.Context, op *xxx_LookupSids2Operation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedNames = op.TranslatedNames
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupSids2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupSids2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSids2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNames2Operation structure represents the LsarLookupNames2 operation
type xxx_LookupNames2Operation struct {
	Policy            *Handle               `idl:"name:PolicyHandle" json:"policy"`
	Count             uint32                `idl:"name:Count" json:"count"`
	Names             []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedSIDs    *TranslatedSIDsEx     `idl:"name:TranslatedSids" json:"translated_sids"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	LookupOptions     uint32                `idl:"name:LookupOptions" json:"lookup_options"`
	ClientRevision    uint32                `idl:"name:ClientRevision" json:"client_revision"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNames2Operation) OpNum() int { return 58 }

func (o *xxx_LookupNames2Operation) OpName() string { return "/lsarpc/v0/LsarLookupNames2" }

func (o *xxx_LookupNames2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.Count == 0 {
		o.Count = uint32(len(o.Names))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupNames2Request structure represents the LsarLookupNames2 operation request
type LookupNames2Request struct {
	// PolicyHandle: Context handle obtained by an LsarOpenPolicy or LsarOpenPolicy2 call.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Count: Number of security principal names to look up.<31>
	Count uint32 `idl:"name:Count" json:"count"`
	// Names: Contains the security principal names to translate, as specified in section
	// 3.1.4.5.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx `idl:"name:TranslatedSids" json:"translated_sids"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter has no effect on message processing in
	// any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// LookupOptions: Flags that control the lookup operation. For possible values and their
	// meanings, see section 3.1.4.5.
	LookupOptions uint32 `idl:"name:LookupOptions" json:"lookup_options"`
	// ClientRevision: Version of the client, which implies the client's capabilities. For
	// possible values and their meanings, see section 3.1.4.5.
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
}

func (o *LookupNames2Request) xxx_ToOp(ctx context.Context, op *xxx_LookupNames2Operation) *xxx_LookupNames2Operation {
	if op == nil {
		op = &xxx_LookupNames2Operation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.Count = o.Count
	op.Names = o.Names
	op.TranslatedSIDs = o.TranslatedSIDs
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	op.LookupOptions = o.LookupOptions
	op.ClientRevision = o.ClientRevision
	return op
}

func (o *LookupNames2Request) xxx_FromOp(ctx context.Context, op *xxx_LookupNames2Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Count = op.Count
	o.Names = op.Names
	o.TranslatedSIDs = op.TranslatedSIDs
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
	o.LookupOptions = op.LookupOptions
	o.ClientRevision = op.ClientRevision
}
func (o *LookupNames2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupNames2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNames2Response structure represents the LsarLookupNames2 operation response
type LookupNames2Response struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx `idl:"name:TranslatedSids" json:"translated_sids"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter has no effect on message processing in
	// any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupNames2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNames2Response) xxx_ToOp(ctx context.Context, op *xxx_LookupNames2Operation) *xxx_LookupNames2Operation {
	if op == nil {
		op = &xxx_LookupNames2Operation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedSIDs = o.TranslatedSIDs
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupNames2Response) xxx_FromOp(ctx context.Context, op *xxx_LookupNames2Operation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedSIDs = op.TranslatedSIDs
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupNames2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupNames2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNames3Operation structure represents the LsarLookupNames3 operation
type xxx_LookupNames3Operation struct {
	Policy            *Handle               `idl:"name:PolicyHandle" json:"policy"`
	Count             uint32                `idl:"name:Count" json:"count"`
	Names             []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedSIDs    *TranslatedSIDsEx2    `idl:"name:TranslatedSids" json:"translated_sids"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	LookupOptions     uint32                `idl:"name:LookupOptions" json:"lookup_options"`
	ClientRevision    uint32                `idl:"name:ClientRevision" json:"client_revision"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNames3Operation) OpNum() int { return 68 }

func (o *xxx_LookupNames3Operation) OpName() string { return "/lsarpc/v0/LsarLookupNames3" }

func (o *xxx_LookupNames3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.Count == 0 {
		o.Count = uint32(len(o.Names))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy != nil {
			if err := o.Policy.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Handle{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PolicyHandle {in} (1:{context_handle, alias=LSAPR_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Policy == nil {
			o.Policy = &Handle{}
		}
		if err := o.Policy.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx2{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx2{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupNames3Request structure represents the LsarLookupNames3 operation request
type LookupNames3Request struct {
	// PolicyHandle: Context handle obtained by an LsarOpenPolicy or LsarOpenPolicy2 call.
	Policy *Handle `idl:"name:PolicyHandle" json:"policy"`
	// Count: Number of security principal names to look up.<29>
	Count uint32 `idl:"name:Count" json:"count"`
	// Names: Contains the security principal names to translate, as specified in section
	// 3.1.4.5.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx2 `idl:"name:TranslatedSids" json:"translated_sids"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter has no effect on message processing in
	// any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// LookupOptions: Flags that control the lookup operation. For possible values and their
	// meanings, see section 3.1.4.5.
	LookupOptions uint32 `idl:"name:LookupOptions" json:"lookup_options"`
	// ClientRevision: Version of the client, which implies the client's capabilities. For
	// possible values and their meanings, see section 3.1.4.5.
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
}

func (o *LookupNames3Request) xxx_ToOp(ctx context.Context, op *xxx_LookupNames3Operation) *xxx_LookupNames3Operation {
	if op == nil {
		op = &xxx_LookupNames3Operation{}
	}
	if o == nil {
		return op
	}
	op.Policy = o.Policy
	op.Count = o.Count
	op.Names = o.Names
	op.TranslatedSIDs = o.TranslatedSIDs
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	op.LookupOptions = o.LookupOptions
	op.ClientRevision = o.ClientRevision
	return op
}

func (o *LookupNames3Request) xxx_FromOp(ctx context.Context, op *xxx_LookupNames3Operation) {
	if o == nil {
		return
	}
	o.Policy = op.Policy
	o.Count = op.Count
	o.Names = op.Names
	o.TranslatedSIDs = op.TranslatedSIDs
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
	o.LookupOptions = op.LookupOptions
	o.ClientRevision = op.ClientRevision
}
func (o *LookupNames3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupNames3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNames3Response structure represents the LsarLookupNames3 operation response
type LookupNames3Response struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedSids: On successful return, contains the corresponding SID forms for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx2 `idl:"name:TranslatedSids" json:"translated_sids"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter has no effect on message processing in
	// any environment. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupNames3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNames3Response) xxx_ToOp(ctx context.Context, op *xxx_LookupNames3Operation) *xxx_LookupNames3Operation {
	if op == nil {
		op = &xxx_LookupNames3Operation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedSIDs = o.TranslatedSIDs
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupNames3Response) xxx_FromOp(ctx context.Context, op *xxx_LookupNames3Operation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedSIDs = op.TranslatedSIDs
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupNames3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupNames3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupSids3Operation structure represents the LsarLookupSids3 operation
type xxx_LookupSids3Operation struct {
	SIDEnumBuffer     *SIDEnumBuffer        `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedNames   *TranslatedNamesEx    `idl:"name:TranslatedNames" json:"translated_names"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	LookupOptions     uint32                `idl:"name:LookupOptions" json:"lookup_options"`
	ClientRevision    uint32                `idl:"name:ClientRevision" json:"client_revision"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupSids3Operation) OpNum() int { return 76 }

func (o *xxx_LookupSids3Operation) OpName() string { return "/lsarpc/v0/LsarLookupSids3" }

func (o *xxx_LookupSids3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer != nil {
			if err := o.SIDEnumBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SIDEnumBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNamesEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SidEnumBuffer {in} (1:{alias=PLSAPR_SID_ENUM_BUFFER,pointer=ref}*(1))(2:{alias=LSAPR_SID_ENUM_BUFFER}(struct))
	{
		if o.SIDEnumBuffer == nil {
			o.SIDEnumBuffer = &SIDEnumBuffer{}
		}
		if err := o.SIDEnumBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNamesEx{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames != nil {
			if err := o.TranslatedNames.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedNamesEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupSids3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedNames {in, out} (1:{alias=PLSAPR_TRANSLATED_NAMES_EX,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_NAMES_EX}(struct))
	{
		if o.TranslatedNames == nil {
			o.TranslatedNames = &TranslatedNamesEx{}
		}
		if err := o.TranslatedNames.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupSids3Request structure represents the LsarLookupSids3 operation request
type LookupSids3Request struct {
	// SidEnumBuffer: Contains the SIDs to be translated. The SIDs in this structure can
	// be that of users, groups, computers, Windows-defined well-known security principals,
	// or domains.
	SIDEnumBuffer *SIDEnumBuffer `idl:"name:SidEnumBuffer" json:"sid_enum_buffer"`
	// TranslatedNames: On successful return, contains the corresponding name forms for
	// security principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNamesEx `idl:"name:TranslatedNames" json:"translated_names"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// LookupOptions: Flags that control the lookup operation. This parameter is reserved
	// for future use; it MUST be set to 0 and ignored on receipt.
	LookupOptions uint32 `idl:"name:LookupOptions" json:"lookup_options"`
	// ClientRevision: Version of the client, which implies the client's capabilities. For
	// possible values and their meanings, see section 3.1.4.5.
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
}

func (o *LookupSids3Request) xxx_ToOp(ctx context.Context, op *xxx_LookupSids3Operation) *xxx_LookupSids3Operation {
	if op == nil {
		op = &xxx_LookupSids3Operation{}
	}
	if o == nil {
		return op
	}
	op.SIDEnumBuffer = o.SIDEnumBuffer
	op.TranslatedNames = o.TranslatedNames
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	op.LookupOptions = o.LookupOptions
	op.ClientRevision = o.ClientRevision
	return op
}

func (o *LookupSids3Request) xxx_FromOp(ctx context.Context, op *xxx_LookupSids3Operation) {
	if o == nil {
		return
	}
	o.SIDEnumBuffer = op.SIDEnumBuffer
	o.TranslatedNames = op.TranslatedNames
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
	o.LookupOptions = op.LookupOptions
	o.ClientRevision = op.ClientRevision
}
func (o *LookupSids3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupSids3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSids3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupSids3Response structure represents the LsarLookupSids3 operation response
type LookupSids3Response struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedNames: On successful return, contains the corresponding name forms for
	// security principal SIDs in the SidEnumBuffer parameter. It MUST be ignored on input.
	TranslatedNames *TranslatedNamesEx `idl:"name:TranslatedNames" json:"translated_names"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to their name forms. It MUST be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupSids3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupSids3Response) xxx_ToOp(ctx context.Context, op *xxx_LookupSids3Operation) *xxx_LookupSids3Operation {
	if op == nil {
		op = &xxx_LookupSids3Operation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedNames = o.TranslatedNames
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupSids3Response) xxx_FromOp(ctx context.Context, op *xxx_LookupSids3Operation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedNames = op.TranslatedNames
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupSids3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupSids3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupSids3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LookupNames4Operation structure represents the LsarLookupNames4 operation
type xxx_LookupNames4Operation struct {
	Count             uint32                `idl:"name:Count" json:"count"`
	Names             []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	TranslatedSIDs    *TranslatedSIDsEx2    `idl:"name:TranslatedSids" json:"translated_sids"`
	LookupLevel       LookupLevel           `idl:"name:LookupLevel" json:"lookup_level"`
	MappedCount       uint32                `idl:"name:MappedCount" json:"mapped_count"`
	LookupOptions     uint32                `idl:"name:LookupOptions" json:"lookup_options"`
	ClientRevision    uint32                `idl:"name:ClientRevision" json:"client_revision"`
	Return            int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_LookupNames4Operation) OpNum() int { return 77 }

func (o *xxx_LookupNames4Operation) OpName() string { return "/lsarpc/v0/LsarLookupNames4" }

func (o *xxx_LookupNames4Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Names != nil && o.Count == 0 {
		o.Count = uint32(len(o.Names))
	}
	if o.Count > uint32(1000) {
		return fmt.Errorf("Count is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames4Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.WriteEnum(uint16(o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.WriteData(o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames4Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Count {in} (1:{range=(0,1000)}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Names {in} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}[dim:0,size_is=Count](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*dtyp.UnicodeString, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			if o.Names[i1] == nil {
				o.Names[i1] = &dtyp.UnicodeString{}
			}
			if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx2{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// LookupLevel {in} (1:{alias=LSAP_LOOKUP_LEVEL}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.LookupLevel)); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// LookupOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LookupOptions); err != nil {
			return err
		}
	}
	// ClientRevision {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ClientRevision); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames4Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames4Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		if o.ReferencedDomains != nil {
			_ptr_ReferencedDomains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReferencedDomains != nil {
					if err := o.ReferencedDomains.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ReferencedDomainList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReferencedDomains, _ptr_ReferencedDomains); err != nil {
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
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs != nil {
			if err := o.TranslatedSIDs.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TranslatedSIDsEx2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LookupNames4Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReferencedDomains {out} (1:{pointer=ref}*(2))(2:{alias=PLSAPR_REFERENCED_DOMAIN_LIST,pointer=ref}*(1))(3:{alias=LSAPR_REFERENCED_DOMAIN_LIST}(struct))
	{
		_ptr_ReferencedDomains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReferencedDomains == nil {
				o.ReferencedDomains = &ReferencedDomainList{}
			}
			if err := o.ReferencedDomains.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReferencedDomains := func(ptr interface{}) { o.ReferencedDomains = *ptr.(**ReferencedDomainList) }
		if err := w.ReadPointer(&o.ReferencedDomains, _s_ReferencedDomains, _ptr_ReferencedDomains); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TranslatedSids {in, out} (1:{alias=PLSAPR_TRANSLATED_SIDS_EX2,pointer=ref}*(1))(2:{alias=LSAPR_TRANSLATED_SIDS_EX2}(struct))
	{
		if o.TranslatedSIDs == nil {
			o.TranslatedSIDs = &TranslatedSIDsEx2{}
		}
		if err := o.TranslatedSIDs.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MappedCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.MappedCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NTSTATUS}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LookupNames4Request structure represents the LsarLookupNames4 operation request
type LookupNames4Request struct {
	// Count: Number of security principal names to look up.<24>
	Count uint32 `idl:"name:Count" json:"count"`
	// Names: Contains the security principal names to translate. The RPC_UNICODE_STRING
	// structure is defined in [MS-DTYP] section 2.3.10.
	//
	// * User principal names (UPNs) ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_9d606f55-b798-4def-bf96-97b878bb92c6
	// ) , such as user_name@example.example.com.
	//
	// * Fully qualified account names based on either DNS or NetBIOS names. For example:
	// example.example.com\user_name or example\user_name, where the generalized form is
	// domain\user account name, and domain is either the fully qualified DNS name ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_102a36e2-f66f-49e2-bee3-558736b2ecd5
	// ) or the NetBIOS name of the trusted domain ( e79f2680-84d9-4d34-bc78-5ab9e1255653#gt_f2f00d47-6cf2-4b32-b8f7-b63e38e2e9c4
	// ).
	//
	// * Unqualified or isolated names, such as user_name.
	Names []*dtyp.UnicodeString `idl:"name:Names;size_is:(Count)" json:"names"`
	// TranslatedSids: On successful return, contains the corresponding SID form for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx2 `idl:"name:TranslatedSids" json:"translated_sids"`
	// LookupLevel: Specifies what scopes are to be used during translation, as specified
	// in section 2.2.16.
	LookupLevel LookupLevel `idl:"name:LookupLevel" json:"lookup_level"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter is left as an input parameter for backward
	// compatibility and has no effect on message processing in any environment. It MUST
	// be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// LookupOptions: Flags specified by the caller that control the lookup operation. The
	// value MUST be one of the following.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Isolated names are searched for even when they are not on the local computer.    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x80000000 | Isolated names (except for user principal names (UPNs)) are searched for only on |
	//	|            | the local account database. UPNs are not searched for in any of the views.       |
	//	+------------+----------------------------------------------------------------------------------+
	LookupOptions uint32 `idl:"name:LookupOptions" json:"lookup_options"`
	// ClientRevision: Version of the client, which implies the client's capabilities. The
	// value MUST be one of the following.<25>
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The client does not understand DNS domain names and is not aware that it might   |
	//	|            | be part of a forest.                                                             |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The client understands DNS domain names and is aware that it might be part of    |
	//	|            | a forest. Error codes returned can include STATUS_TRUSTED_DOMAIN_FAILURE and     |
	//	|            | STATUS_TRUSTED_RELATIONSHIP_FAILURE, which are not returned for ClientRevision   |
	//	|            | of 0x00000001. For more information on error codes, see [MS-ERREF].              |
	//	+------------+----------------------------------------------------------------------------------+
	ClientRevision uint32 `idl:"name:ClientRevision" json:"client_revision"`
}

func (o *LookupNames4Request) xxx_ToOp(ctx context.Context, op *xxx_LookupNames4Operation) *xxx_LookupNames4Operation {
	if op == nil {
		op = &xxx_LookupNames4Operation{}
	}
	if o == nil {
		return op
	}
	op.Count = o.Count
	op.Names = o.Names
	op.TranslatedSIDs = o.TranslatedSIDs
	op.LookupLevel = o.LookupLevel
	op.MappedCount = o.MappedCount
	op.LookupOptions = o.LookupOptions
	op.ClientRevision = o.ClientRevision
	return op
}

func (o *LookupNames4Request) xxx_FromOp(ctx context.Context, op *xxx_LookupNames4Operation) {
	if o == nil {
		return
	}
	o.Count = op.Count
	o.Names = op.Names
	o.TranslatedSIDs = op.TranslatedSIDs
	o.LookupLevel = op.LookupLevel
	o.MappedCount = op.MappedCount
	o.LookupOptions = op.LookupOptions
	o.ClientRevision = op.ClientRevision
}
func (o *LookupNames4Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LookupNames4Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames4Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LookupNames4Response structure represents the LsarLookupNames4 operation response
type LookupNames4Response struct {
	// ReferencedDomains: On successful return, contains the domain information for the
	// domain to which each security principal belongs. The domain information includes
	// a NetBIOS domain name and a domain SID for each entry in the list.
	ReferencedDomains *ReferencedDomainList `idl:"name:ReferencedDomains" json:"referenced_domains"`
	// TranslatedSids: On successful return, contains the corresponding SID form for security
	// principal names in the Names parameter. It MUST be ignored on input.
	TranslatedSIDs *TranslatedSIDsEx2 `idl:"name:TranslatedSids" json:"translated_sids"`
	// MappedCount: On successful return, contains the number of names that are translated
	// completely to the SID form. This parameter is left as an input parameter for backward
	// compatibility and has no effect on message processing in any environment. It MUST
	// be ignored on input.
	MappedCount uint32 `idl:"name:MappedCount" json:"mapped_count"`
	// Return: The LsarLookupNames4 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *LookupNames4Response) xxx_ToOp(ctx context.Context, op *xxx_LookupNames4Operation) *xxx_LookupNames4Operation {
	if op == nil {
		op = &xxx_LookupNames4Operation{}
	}
	if o == nil {
		return op
	}
	op.ReferencedDomains = o.ReferencedDomains
	op.TranslatedSIDs = o.TranslatedSIDs
	op.MappedCount = o.MappedCount
	op.Return = o.Return
	return op
}

func (o *LookupNames4Response) xxx_FromOp(ctx context.Context, op *xxx_LookupNames4Operation) {
	if o == nil {
		return
	}
	o.ReferencedDomains = op.ReferencedDomains
	o.TranslatedSIDs = op.TranslatedSIDs
	o.MappedCount = op.MappedCount
	o.Return = op.Return
}
func (o *LookupNames4Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LookupNames4Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LookupNames4Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
