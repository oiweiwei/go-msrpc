package lsarpc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// lsarpc server interface.
type LsarpcServer interface {

	// The LsarClose method frees the resources held by a context handle that was opened
	// earlier. After response, the context handle is no longer usable, and any subsequent
	// uses of this handle MUST fail.
	Close(context.Context, *CloseRequest) (*CloseResponse, error)

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
	OpenPolicy(context.Context, *OpenPolicyRequest) (*OpenPolicyResponse, error)

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
	LookupNames(context.Context, *LookupNamesRequest) (*LookupNamesResponse, error)

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
	LookupSIDs(context.Context, *LookupSIDsRequest) (*LookupSIDsResponse, error)

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
	OpenPolicy2(context.Context, *OpenPolicy2Request) (*OpenPolicy2Response, error)

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
	GetUserName(context.Context, *GetUserNameRequest) (*GetUserNameResponse, error)

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
	LookupSids2(context.Context, *LookupSids2Request) (*LookupSids2Response, error)

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
	LookupNames2(context.Context, *LookupNames2Request) (*LookupNames2Response, error)

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
	LookupNames3(context.Context, *LookupNames3Request) (*LookupNames3Response, error)

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
	LookupSids3(context.Context, *LookupSids3Request) (*LookupSids3Response, error)

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
	LookupNames4(context.Context, *LookupNames4Request) (*LookupNames4Response, error)
}

func RegisterLsarpcServer(conn dcerpc.Conn, o LsarpcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewLsarpcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(LsarpcSyntaxV0_0))...)
}

func NewLsarpcServerHandle(o LsarpcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return LsarpcServerHandle(ctx, o, opNum, r)
	}
}

func LsarpcServerHandle(ctx context.Context, o LsarpcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // LsarClose
		op := &xxx_CloseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Close(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Lsar_LSA_DP_2
		// Lsar_LSA_DP_2
		return nil, nil
	case 3: // Lsar_LSA_DP_3
		// Lsar_LSA_DP_3
		return nil, nil
	case 4: // Lsar_LSA_DP_4
		// Lsar_LSA_DP_4
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // LsarOpenPolicy
		op := &xxx_OpenPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Lsar_LSA_DP_7
		// Lsar_LSA_DP_7
		return nil, nil
	case 8: // Lsar_LSA_DP_8
		// Lsar_LSA_DP_8
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // Lsar_LSA_DP_10
		// Lsar_LSA_DP_10
		return nil, nil
	case 11: // Lsar_LSA_DP_11
		// Lsar_LSA_DP_11
		return nil, nil
	case 12: // Lsar_LSA_DP_12
		// Lsar_LSA_DP_12
		return nil, nil
	case 13: // Lsar_LSA_DP_13
		// Lsar_LSA_DP_13
		return nil, nil
	case 14: // LsarLookupNames
		op := &xxx_LookupNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // LsarLookupSids
		op := &xxx_LookupSIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupSIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupSIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Lsar_LSA_DP_16
		// Lsar_LSA_DP_16
		return nil, nil
	case 17: // Lsar_LSA_DP_17
		// Lsar_LSA_DP_17
		return nil, nil
	case 18: // Lsar_LSA_DP_18
		// Lsar_LSA_DP_18
		return nil, nil
	case 19: // Lsar_LSA_DP_19
		// Lsar_LSA_DP_19
		return nil, nil
	case 20: // Lsar_LSA_DP_20
		// Lsar_LSA_DP_20
		return nil, nil
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // Opnum22NotUsedOnWire
		// Opnum22NotUsedOnWire
		return nil, nil
	case 23: // Lsar_LSA_DP_23
		// Lsar_LSA_DP_23
		return nil, nil
	case 24: // Lsar_LSA_DP_24
		// Lsar_LSA_DP_24
		return nil, nil
	case 25: // Lsar_LSA_DP_25
		// Lsar_LSA_DP_25
		return nil, nil
	case 26: // Lsar_LSA_DP_26
		// Lsar_LSA_DP_26
		return nil, nil
	case 27: // Lsar_LSA_DP_27
		// Lsar_LSA_DP_27
		return nil, nil
	case 28: // Lsar_LSA_DP_28
		// Lsar_LSA_DP_28
		return nil, nil
	case 29: // Lsar_LSA_DP_29
		// Lsar_LSA_DP_29
		return nil, nil
	case 30: // Lsar_LSA_DP_30
		// Lsar_LSA_DP_30
		return nil, nil
	case 31: // Lsar_LSA_DP_31
		// Lsar_LSA_DP_31
		return nil, nil
	case 32: // Lsar_LSA_DP_32
		// Lsar_LSA_DP_32
		return nil, nil
	case 33: // Lsar_LSA_DP_33
		// Lsar_LSA_DP_33
		return nil, nil
	case 34: // Lsar_LSA_DP_34
		// Lsar_LSA_DP_34
		return nil, nil
	case 35: // Lsar_LSA_DP_35
		// Lsar_LSA_DP_35
		return nil, nil
	case 36: // Lsar_LSA_DP_36
		// Lsar_LSA_DP_36
		return nil, nil
	case 37: // Lsar_LSA_DP_37
		// Lsar_LSA_DP_37
		return nil, nil
	case 38: // Lsar_LSA_DP_38
		// Lsar_LSA_DP_38
		return nil, nil
	case 39: // Lsar_LSA_DP_39
		// Lsar_LSA_DP_39
		return nil, nil
	case 40: // Lsar_LSA_DP_40
		// Lsar_LSA_DP_40
		return nil, nil
	case 41: // Lsar_LSA_DP_41
		// Lsar_LSA_DP_41
		return nil, nil
	case 42: // Lsar_LSA_DP_42
		// Lsar_LSA_DP_42
		return nil, nil
	case 43: // Lsar_LSA_DP_43
		// Lsar_LSA_DP_43
		return nil, nil
	case 44: // LsarOpenPolicy2
		op := &xxx_OpenPolicy2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPolicy2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPolicy2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // LsarGetUserName
		op := &xxx_GetUserNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUserNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUserName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // Lsar_LSA_DP_46
		// Lsar_LSA_DP_46
		return nil, nil
	case 47: // Lsar_LSA_DP_47
		// Lsar_LSA_DP_47
		return nil, nil
	case 48: // Lsar_LSA_DP_48
		// Lsar_LSA_DP_48
		return nil, nil
	case 49: // Lsar_LSA_DP_49
		// Lsar_LSA_DP_49
		return nil, nil
	case 50: // Lsar_LSA_DP_50
		// Lsar_LSA_DP_50
		return nil, nil
	case 51: // Lsar_LSA_DP_51
		// Lsar_LSA_DP_51
		return nil, nil
	case 52: // Opnum52NotUsedOnWire
		// Opnum52NotUsedOnWire
		return nil, nil
	case 53: // Lsar_LSA_DP_53
		// Lsar_LSA_DP_53
		return nil, nil
	case 54: // Lsar_LSA_DP_54
		// Lsar_LSA_DP_54
		return nil, nil
	case 55: // Lsar_LSA_DP_55
		// Lsar_LSA_DP_55
		return nil, nil
	case 56: // Opnum56NotUsedOnWire
		// Opnum56NotUsedOnWire
		return nil, nil
	case 57: // LsarLookupSids2
		op := &xxx_LookupSids2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupSids2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupSids2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // LsarLookupNames2
		op := &xxx_LookupNames2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupNames2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupNames2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // Lsar_LSA_DP_59
		// Lsar_LSA_DP_59
		return nil, nil
	case 60: // Opnum60NotUsedOnWire
		// Opnum60NotUsedOnWire
		return nil, nil
	case 61: // Opnum61NotUsedOnWire
		// Opnum61NotUsedOnWire
		return nil, nil
	case 62: // Opnum62NotUsedOnWire
		// Opnum62NotUsedOnWire
		return nil, nil
	case 63: // Opnum63NotUsedOnWire
		// Opnum63NotUsedOnWire
		return nil, nil
	case 64: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	case 65: // Opnum65NotUsedOnWire
		// Opnum65NotUsedOnWire
		return nil, nil
	case 66: // Opnum66NotUsedOnWire
		// Opnum66NotUsedOnWire
		return nil, nil
	case 67: // Opnum67NotUsedOnWire
		// Opnum67NotUsedOnWire
		return nil, nil
	case 68: // LsarLookupNames3
		op := &xxx_LookupNames3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupNames3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupNames3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // Opnum69NotUsedOnWire
		// Opnum69NotUsedOnWire
		return nil, nil
	case 70: // Opnum70NotUsedOnWire
		// Opnum70NotUsedOnWire
		return nil, nil
	case 71: // Opnum71NotUsedOnWire
		// Opnum71NotUsedOnWire
		return nil, nil
	case 72: // Opnum72NotUsedOnWire
		// Opnum72NotUsedOnWire
		return nil, nil
	case 73: // Lsar_LSA_DP_73
		// Lsar_LSA_DP_73
		return nil, nil
	case 74: // Lsar_LSA_DP_74
		// Lsar_LSA_DP_74
		return nil, nil
	case 75: // Opnum75NotUsedOnWire
		// Opnum75NotUsedOnWire
		return nil, nil
	case 76: // LsarLookupSids3
		op := &xxx_LookupSids3Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupSids3Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupSids3(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // LsarLookupNames4
		op := &xxx_LookupNames4Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupNames4Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupNames4(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented lsarpc
type UnimplementedLsarpcServer struct {
}

func (UnimplementedLsarpcServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) OpenPolicy(context.Context, *OpenPolicyRequest) (*OpenPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupNames(context.Context, *LookupNamesRequest) (*LookupNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupSIDs(context.Context, *LookupSIDsRequest) (*LookupSIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) OpenPolicy2(context.Context, *OpenPolicy2Request) (*OpenPolicy2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) GetUserName(context.Context, *GetUserNameRequest) (*GetUserNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupSids2(context.Context, *LookupSids2Request) (*LookupSids2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupNames2(context.Context, *LookupNames2Request) (*LookupNames2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupNames3(context.Context, *LookupNames3Request) (*LookupNames3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupSids3(context.Context, *LookupSids3Request) (*LookupSids3Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedLsarpcServer) LookupNames4(context.Context, *LookupNames4Request) (*LookupNames4Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ LsarpcServer = (*UnimplementedLsarpcServer)(nil)
