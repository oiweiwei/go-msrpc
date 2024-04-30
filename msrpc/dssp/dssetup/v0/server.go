package dssetup

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

// dssetup server interface.
type DssetupServer interface {

	// The DsRolerGetPrimaryDomainInformation (Opnum 0) method returns the requested information
	// about the current configuration or state of the computer on which the server is running.
	//
	// Return Values: The method returns 0 if successful; if failed, it returns a nonzero
	// error code as specified in [MS-ERREF]. Specifically, in addition to any other error
	// codes, the server MUST return the following error codes for the following error conditions.
	// Any other values transmitted in this field are implementation-specific. All nonzero
	// values MUST be treated the same for protocol purposes.
	//
	//	+------------------------------------+---------------------------------------+
	//	|               RETURN               |                                       |
	//	|             VALUE/CODE             |              DESCRIPTION              |
	//	|                                    |                                       |
	//	+------------------------------------+---------------------------------------+
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | One or more parameters are invalid.   |
	//	+------------------------------------+---------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | A memory allocation failure occurred. |
	//	+------------------------------------+---------------------------------------+
	GetPrimaryDomainInformation(context.Context, *GetPrimaryDomainInformationRequest) (*GetPrimaryDomainInformationResponse, error)

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire
}

func RegisterDssetupServer(conn dcerpc.Conn, o DssetupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDssetupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DssetupSyntaxV0_0))...)
}

func NewDssetupServerHandle(o DssetupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DssetupServerHandle(ctx, o, opNum, r)
	}
}

func DssetupServerHandle(ctx context.Context, o DssetupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // DsRolerGetPrimaryDomainInformation
		in := &GetPrimaryDomainInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrimaryDomainInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
