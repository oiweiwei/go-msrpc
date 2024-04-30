package trkwks

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

// trkwks server interface.
type TrkwksServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

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

	// The LnkSearchMachine method searches for a file object on the specified computer.
	// If information on the file is found, the method attempts to return it. If the file
	// has been moved, the method returns information about its new location.
	//
	// Return Values: A 32-bit integer that indicates success or failure: A value of 0 or
	// any positive value indicates success; a negative value indicates failure. Some of
	// the possible return codes are listed in the following table.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8DEAD106 TRK_E_POTENTIAL_FILE_FOUND | A file was found by a DLT Workstation server processing a LnkSearchMachine call, |
	//	|                                       | and the file meets some—but not all—of the necessary criteria to be considered   |
	//	|                                       | the correct file. More details are presented later in this section.              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8DEAD101 TRK_E_REFERRAL             | A file could not be found by a DLT Workstation server processing a               |
	//	|                                       | LnkSearchMachine call, and the file meets some—but not all—of the necessary      |
	//	|                                       | criteria to be considered the correct file. More details are presented later in  |
	//	|                                       | this section.                                                                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: None.
	//
	// The server responds to this call by attempting to find the requested file in one
	// of its volumes, and either returns the file's new location (in the form of a UNC),
	// or returns information about the location to which the file has moved (see description
	// of TRK_E_REFERRAL in the following bulleted list).
	//
	// The server SHOULD<10> return a failure value, if the file is found, but the length
	// of the path to be returned in the pstszPath argument exceeds 261 characters (where
	// the 261 characters does not include a string terminator character).
	//
	// For each of the possible return values, the output parameters of this call MUST be
	// set as specified in the following list. In these specifications, the requested ObjectID
	// MUST correspond to the Object field of the pDroidLast parameter, the requested VolumeID
	// MUST correspond to the Volume field of the pDroidLast parameter, and the FileID requested
	// MUST correspond to the value of the pDroidBirthLast parameter. The numeric values
	// corresponding to these returns are defined in section 3.1.1:
	//
	// * Success: A success value MUST be returned if the server finds the file, and the
	// client is authorized to get information about the file via a UNC path. The file MUST
	// be found by searching all volumes on the server computer for a file whose ObjectID
	// is equal to the ObjectID of the request, and whose FileID is equal to the FileID
	// of the request. To perform the authorization check, the server MUST use the client's
	// identity (obtained as specified in [MS-RPCE] ( ../ms-rpce/290c38b1-92fe-4229-91e6-4fc376610c15
	// ) section 3.3.3.4.3 ( ../ms-rpce/29b8217a-0bda-4fdb-a3ea-48560125ae8d ) ) to determine,
	// based on local policy, whether or not the client is authorized to get the UNC of
	// the file.
	//
	// If there is more than one file on the server computer that satisfies these conditions,
	// the file MUST be selected as follows:
	//
	// * If one of the matching files is on the volume whose VolumeID equals the VolumeID
	// of the request, that file is selected.
	//
	// * Otherwise, the behavior is arbitrary.
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the pdroidBirthLast that is specified by the client.
	//
	// pdroidNext : The current FileLocation value for the file.
	//
	// pmcidNext : The MachineID of the server computer.
	//
	// ptszPath : The UNC of the file on the server computer. If the file can be located
	// under multiple UNC paths, one of those UNC paths MUST be returned, but the server
	// MAY return any of those paths. <11> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_11
	// )
	//
	// * TRK_E_REFERRAL: This value MUST be returned if the file is no longer stored on
	// any of the volumes of the server computer, but if the MoveTable of the last volume
	// has an entry for the file. (Note that there is no additional per-file access check
	// in this case.) The last volume MUST be determined by using the volume on the server
	// whose VolumeID equals the VolumeID of the request. The entry in the MoveTable for
	// the file, as defined in section 3.1.1, MUST be that entry whose ObjectID equals the
	// ObjectID of the request.
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the FileID specified by the client in the *pdroidBirthLast*
	// field of the call.
	//
	// pdroidNext : The value of the FileLocation field for the file's entry in the *MoveTable*.
	//
	// pmcidNext : The value of the MachineID field for the file's entry in the *MoveTable*.
	//
	// ptszPath : The server MUST NOT modify this value.
	//
	// * TRK_E_POTENTIAL_FILE_FOUND: This value MUST be returned if there is no file on
	// any of the server's volumes with the requested FileID and ObjectID, and there is
	// no entry for this file's ObjectID in the MoveTable, but if there is a file with the
	// requested ObjectID on one of the server's volumes, and the FileID on that file is
	// all zeros. <12> ( caa74a29-27fb-4f50-afa5-bbf152fe5d27#Appendix_A_12 )
	//
	// The server MUST set the output parameters as follows:
	//
	// pdroidBirthNext : The value of the FileID for the found file.
	//
	// pdroidNext : The FileLocation of the file found by the server.
	//
	// pmcidNext : The MachineID of the server computer.
	//
	// ptszPath : A UNC of the file on the server computer.
	//
	// * Other negative return values: A negative value other than those mentioned above
	// MUST be returned if none of the preceding cases is met. The server MUST NOT modify
	// any of the output parameters.
	SearchMachine(context.Context, *SearchMachineRequest) (*SearchMachineResponse, error)
}

func RegisterTrkwksServer(conn dcerpc.Conn, o TrkwksServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTrkwksServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TrkwksSyntaxV1_2))...)
}

func NewTrkwksServerHandle(o TrkwksServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TrkwksServerHandle(ctx, o, opNum, r)
	}
}

func TrkwksServerHandle(ctx context.Context, o TrkwksServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
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
	case 12: // LnkSearchMachine
		in := &SearchMachineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SearchMachine(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
